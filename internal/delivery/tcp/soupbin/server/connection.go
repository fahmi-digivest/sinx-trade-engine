package server

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/codec"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/frame"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/message"
)

// Connection manages the lifecycle of a single downstream client connection.
type Connection struct {
	conn    net.Conn
	cfg     Config
	auth    *Authenticator
	session *Session

	fr   *frame.Reader
	dec  *codec.Decoder
	disp *Dispatcher
}

// NewConnection creates a new Connection for conn.
func NewConnection(conn net.Conn, cfg Config, auth *Authenticator, session *Session) *Connection {
	fw := frame.NewWriter(conn)
	return &Connection{
		conn:    conn,
		cfg:     cfg,
		auth:    auth,
		session: session,
		fr:      frame.NewReader(conn),
		dec:     codec.NewDecoder(),
		disp:    NewDispatcher(fw),
	}
}

// Serve handles the full lifecycle of the connection: login handshake, then
// replaying buffered messages followed by live streaming.
func (c *Connection) Serve() {
	defer c.conn.Close()

	addr := c.conn.RemoteAddr()

	// --- Login phase ---
	c.conn.SetDeadline(time.Now().Add(c.cfg.LoginTimeout))
	req, err := c.readLoginRequest()
	if err != nil {
		log.Printf("[soupbin/server] %s login read error: %v", addr, err)
		return
	}

	if !c.auth.Validate(req.Username, req.Password) {
		_ = c.disp.SendLoginRejected(message.RejectReasonNotAuthorized)
		return
	}

	// Determine start sequence number.
	startSeq := uint64(1)
	if req.RequestedSequenceNumber != "" && req.RequestedSequenceNumber != "0" {
		n, err := strconv.ParseUint(req.RequestedSequenceNumber, 10, 64)
		if err == nil {
			startSeq = n
		}
	}
	// Clamp to available range.
	nextAvail := c.session.NextSeq()
	if startSeq > nextAvail {
		startSeq = nextAvail
	}

	if err := c.disp.SendLoginAccepted(
		c.session.ID,
		fmt.Sprintf("%d", startSeq),
	); err != nil {
		log.Printf("[soupbin/server] %s login accepted write error: %v", addr, err)
		return
	}

	// Clear deadline for the data phase – we manage timeouts via heartbeats.
	c.conn.SetDeadline(time.Time{})

	log.Printf("[soupbin/server] %s logged in (session=%s startSeq=%d)", addr, c.session.ID, startSeq)

	// --- Replay buffered messages ---
	for seq := startSeq; seq < nextAvail; seq++ {
		msg := c.session.Get(seq)
		if msg == nil {
			continue
		}
		if err := c.disp.SendSequencedData(msg); err != nil {
			log.Printf("[soupbin/server] %s replay error: %v", addr, err)
			return
		}
	}

	// --- Live streaming + heartbeat loop ---
	// We use a simple ticker-based approach.  A production implementation
	// would use a fan-out subscription mechanism; here we poll the session.
	currentSeq := nextAvail
	hbTicker := time.NewTicker(c.cfg.HeartbeatInterval)
	defer hbTicker.Stop()

	clientTimeout := time.NewTimer(c.cfg.ClientTimeout)
	defer clientTimeout.Stop()

	// Read from client in a separate goroutine so we can also send heartbeats.
	clientMsgs := make(chan message.Message, 16)
	clientErr := make(chan error, 1)
	go func() {
		for {
			f, err := c.fr.ReadFrame()
			if err != nil {
				clientErr <- err
				return
			}
			msg, err := c.dec.Decode(f)
			if err != nil {
				log.Printf("[soupbin/server] %s decode error: %v", addr, err)
				continue
			}
			clientMsgs <- msg
		}
	}()

	lastSent := time.Now()

	for {
		select {
		case msg := <-clientMsgs:
			clientTimeout.Reset(c.cfg.ClientTimeout)
			switch msg.(type) {
			case *message.LogoutRequest:
				log.Printf("[soupbin/server] %s logout requested", addr)
				return
			case *message.ClientHeartbeat:
				// no-op; timer already reset
			case *message.UnsequencedData:
				// higher-level protocol handles these
			}

		case err := <-clientErr:
			log.Printf("[soupbin/server] %s read error: %v", addr, err)
			return

		case <-clientTimeout.C:
			log.Printf("[soupbin/server] %s client timeout", addr)
			return

		case <-hbTicker.C:
			// Send any new sequenced messages first.
			for {
				msg := c.session.Get(currentSeq)
				if msg == nil {
					break
				}
				if err := c.disp.SendSequencedData(msg); err != nil {
					log.Printf("[soupbin/server] %s send error: %v", addr, err)
					return
				}
				currentSeq++
				lastSent = time.Now()
			}
			// Send heartbeat if nothing was sent recently.
			if time.Since(lastSent) >= c.cfg.HeartbeatInterval {
				if err := c.disp.SendHeartbeat(); err != nil {
					log.Printf("[soupbin/server] %s heartbeat error: %v", addr, err)
					return
				}
				lastSent = time.Now()
			}
		}
	}
}

func (c *Connection) readLoginRequest() (*message.LoginRequest, error) {
	f, err := c.fr.ReadFrame()
	if err != nil {
		return nil, err
	}
	if f.Type != message.PacketTypeLoginRequest {
		return nil, fmt.Errorf("expected LoginRequest (L), got %q", f.Type)
	}
	return message.DecodeLoginRequest(f.Payload)
}
