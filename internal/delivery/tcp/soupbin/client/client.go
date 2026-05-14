package client

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/codec"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/frame"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/message"
)

// Client connects to an upstream SoupBinTCP server, handles the login
// handshake, and delivers messages to a Handler.
type Client struct {
	cfg     Config
	handler Handler

	mu   sync.Mutex
	conn net.Conn
	fw   *frame.Writer
	enc  *codec.Encoder

	nextSeq uint64
	session string

	done chan struct{}
}

// New creates a new Client.  Call Connect to establish the connection.
func New(cfg Config, handler Handler) *Client {
	return &Client{
		cfg:     cfg,
		handler: handler,
		enc:     codec.NewEncoder(),
		done:    make(chan struct{}),
	}
}

// Connect dials the server, performs the login handshake, and starts the
// receive loop.  It blocks until the connection is established or returns an
// error.  Call the blocking Run method to keep reading, or use Connect + Send
// from separate goroutines.
func (c *Client) Connect() error {
	conn, err := net.DialTimeout("tcp", c.cfg.ServerAddr, c.cfg.DialTimeout)
	if err != nil {
		return fmt.Errorf("soupbin/client: dial %s: %w", c.cfg.ServerAddr, err)
	}

	c.mu.Lock()
	c.conn = conn
	c.fw = frame.NewWriter(conn)
	c.mu.Unlock()

	// Send Login Request.
	loginReq := &message.LoginRequest{
		Username:                c.cfg.Username,
		Password:                c.cfg.Password,
		RequestedSession:        c.cfg.RequestedSession,
		RequestedSequenceNumber: c.cfg.RequestedSequenceNumber,
	}
	if err := c.sendMsg(loginReq); err != nil {
		conn.Close()
		return fmt.Errorf("soupbin/client: send login: %w", err)
	}

	// Read Login Accepted / Rejected.
	conn.SetDeadline(time.Now().Add(c.cfg.DialTimeout))
	fr := frame.NewReader(conn)
	f, err := fr.ReadFrame()
	if err != nil {
		conn.Close()
		return fmt.Errorf("soupbin/client: read login response: %w", err)
	}
	conn.SetDeadline(time.Time{})

	dec := codec.NewDecoder()
	msg, err := dec.Decode(f)
	if err != nil {
		conn.Close()
		return err
	}

	switch m := msg.(type) {
	case *message.LoginAccepted:
		seq, _ := strconv.ParseUint(m.SequenceNumber, 10, 64)
		c.nextSeq = seq
		c.session = m.Session
		c.handler.OnLoginAccepted(m.Session, seq)
	case *message.LoginRejected:
		conn.Close()
		return m // LoginRejected implements error
	default:
		conn.Close()
		return fmt.Errorf("soupbin/client: unexpected packet type %T after login", msg)
	}

	return nil
}

// Run starts the main receive loop and heartbeat ticker.  It blocks until the
// connection is closed or an error occurs.  Call Connect before Run.
func (c *Client) Run() error {
	c.mu.Lock()
	conn := c.conn
	c.mu.Unlock()
	if conn == nil {
		return fmt.Errorf("soupbin/client: not connected")
	}

	fr := frame.NewReader(conn)
	dec := codec.NewDecoder()

	hbTicker := time.NewTicker(c.cfg.HeartbeatInterval)
	defer hbTicker.Stop()

	frameCh := make(chan *frame.Frame, 64)
	errCh := make(chan error, 1)

	go func() {
		for {
			f, err := fr.ReadFrame()
			if err != nil {
				errCh <- err
				return
			}
			frameCh <- f
		}
	}()

	lastRecv := time.Now()

	for {
		select {
		case <-c.done:
			return nil

		case f := <-frameCh:
			lastRecv = time.Now()
			msg, err := dec.Decode(f)
			if err != nil {
				log.Printf("[soupbin/client] decode error: %v", err)
				continue
			}
			c.dispatch(msg)

		case err := <-errCh:
			c.handler.OnError(err)
			return err

		case <-hbTicker.C:
			// Check server liveness.
			if time.Since(lastRecv) > c.cfg.ServerTimeout {
				err := fmt.Errorf("soupbin/client: server timeout (no data for %s)", c.cfg.ServerTimeout)
				c.handler.OnError(err)
				return err
			}
			// Send client heartbeat.
			if err := c.sendMsg(&message.ClientHeartbeat{}); err != nil {
				c.handler.OnError(err)
				return err
			}
		}
	}
}

// Send transmits an Unsequenced Data packet to the server.
func (c *Client) Send(msg []byte) error {
	return c.sendMsg(&message.UnsequencedData{Message: msg})
}

// Close sends a Logout Request and closes the underlying connection.
func (c *Client) Close() error {
	close(c.done)
	_ = c.sendMsg(&message.LogoutRequest{})
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Session returns the session ID established during login.
func (c *Client) Session() string { return c.session }

// NextSeq returns the next expected sequence number.
func (c *Client) NextSeq() uint64 { return c.nextSeq }

// ---- internal ---------------------------------------------------------------

func (c *Client) sendMsg(msg message.Message) error {
	f, err := c.enc.Encode(msg)
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.fw.WriteFrame(f)
}

func (c *Client) dispatch(msg message.Message) {
	switch m := msg.(type) {
	case *message.SequencedData:
		seq := c.nextSeq
		c.nextSeq++
		c.handler.OnSequencedData(seq, m.Message)

	case *message.ServerHeartbeat:
		// no-op; liveness tracked by lastRecv

	case *message.EndOfSession:
		c.handler.OnEndOfSession()

	case *message.UnsequencedData:
		if h, ok := c.handler.(UnsequencedHandler); ok {
			h.OnUnsequencedData(m)
		}
	}
}
