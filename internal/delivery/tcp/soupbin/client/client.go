package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/codec"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/frame"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/message"
	port "github.com/fahmi-digivest/sinx-trade-engine/internal/domain/port"
)

// Client connects to an upstream SoupBinTCP server, handles the login
// handshake, and delivers messages to a Handler.
type Client struct {
	name    string
	cfg     Config
	handler Handler
	logger  *slog.Logger
	queue   port.SPSCQueue[*frame.Frame]

	mu        sync.Mutex
	conn      net.Conn
	fw        *frame.Writer
	enc       *codec.Encoder
	closeOnce sync.Once

	nextSeq  uint64
	session  string
	lastRecv time.Time
	lastSend time.Time

	done chan struct{}
}

// New creates a new Client. Call Run with a context to start it.
func New(name string, cfg Config, handler Handler, logger *slog.Logger, queue port.SPSCQueue[*frame.Frame]) *Client {
	if logger == nil {
		logger = slog.Default()
	}

	return &Client{
		name:    name,
		cfg:     cfg,
		handler: handler,
		logger:  logger,
		queue:   queue,
		enc:     codec.NewEncoder(),
		done:    make(chan struct{}),
	}
}

// Name returns the service name for this client instance.
func (c *Client) Name() string { return c.name }

// Session returns the session ID established during login.
func (c *Client) Session() string { return c.session }

// NextSeq returns the next expected sequence number.
func (c *Client) NextSeq() uint64 { return c.nextSeq }

// Connect dials the server, performs the login handshake, and starts the
// receive loop. It blocks until the connection is established or returns an
// error. Call the blocking Run method to keep reading, or use Connect + Send
// from separate goroutines.
func (c *Client) Connect() error {
	c.logger.Info("connecting soupbin client", "server_addr", c.cfg.ServerAddr)

	requestedSession := c.cfg.RequestedSession
	requestedSeq := c.cfg.RequestedSequenceNumber
	if c.session != "" {
		requestedSession = c.session
	}
	if c.nextSeq == 0 || requestedSeq == "1" {
		requestedSeq = "1"
	} else {
		requestedSeq = strconv.FormatUint(c.nextSeq, 10)
	}

	conn, err := net.DialTimeout("tcp", c.cfg.ServerAddr, c.cfg.DialTimeout)
	if err != nil {
		return fmt.Errorf("soupbin/client: dial %s: %w", c.cfg.ServerAddr, err)
	}

	c.mu.Lock()
	c.conn = conn
	c.fw = frame.NewWriter(conn)
	c.mu.Unlock()

	loginReq := &message.LoginRequest{
		Username:                c.cfg.Username,
		Password:                c.cfg.Password,
		RequestedSession:        requestedSession,
		RequestedSequenceNumber: requestedSeq,
	}

	if err := c.sendMsg(loginReq); err != nil {
		conn.Close()
		return fmt.Errorf("soupbin/client: send login: %w", err)
	}

	// Read Login Accepted / Rejected.
	conn.SetDeadline(time.Now().Add(c.cfg.DialTimeout)) //nolint:errcheck
	fr := frame.NewReader(conn)

	f, err := fr.ReadFrame()
	if err != nil {
		conn.Close()
		return fmt.Errorf("soupbin/client: read login response: %w", err)
	}
	conn.SetDeadline(time.Time{}) //nolint:errcheck

	msg, err := codec.NewDecoder().Decode(f)
	if err != nil {
		conn.Close()
		return err
	}

	switch m := msg.(type) {
	case *message.LoginAccepted:
		seq, _ := strconv.ParseUint(m.SequenceNumber, 10, 64)
		c.nextSeq = seq
		c.session = m.Session
		c.setLastRecv(time.Now())
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

// Run keeps the client alive until ctx is cancelled or the client fails fatally.
func (c *Client) Run(ctx context.Context) error {
	c.logger.Info("starting soupbin service")

	var wg sync.WaitGroup
	stopCtxWatcher := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.consumeFrames()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			_ = c.Close()
		case <-stopCtxWatcher:
		}
	}()

	defer func() {
		close(stopCtxWatcher)
		c.queue.Close()
		wg.Wait()
	}()

	reconnectDelay := c.cfg.ReconnectDelay
	if reconnectDelay <= 0 {
		reconnectDelay = time.Second
	}

	for {
		if err := c.Connect(); err != nil {
			if !c.waitReconnect(ctx, reconnectDelay) {
				return err
			}
			continue
		}

		err := c.runLoop()
		if err == nil || errors.Is(err, context.Canceled) {
			c.logger.Info("soupbin service stopped")
			return err
		}

		if !isReconnectable(err) {
			c.logger.Error("soupbin service stopped with error", "err", err)
			return err
		}

		c.logger.Warn("soupbin client disconnected, reconnecting", "err", err, "delay", reconnectDelay)
		_ = c.closeConnection(false)

		if !c.waitReconnect(ctx, reconnectDelay) {
			c.logger.Info("soupbin service stopped")
			return nil
		}
	}
}

// RunWithContext is a compatibility alias for Run.
func (c *Client) RunWithContext(ctx context.Context) error {
	return c.Run(ctx)
}

// Send transmits an Unsequenced Data packet to the server.
func (c *Client) Send(msg []byte) error {
	return c.sendMsg(&message.UnsequencedData{Message: msg})
}

// Close sends a Logout Request and closes the underlying connection.
func (c *Client) Close() error {
	var closeErr error
	c.closeOnce.Do(func() {
		close(c.done)
		closeErr = c.closeConnection(true)
	})
	return closeErr
}

// -----------------------------------------------------------------------------
// Internal helpers
// -----------------------------------------------------------------------------

// runLoop starts the main receive loop and heartbeat ticker.
// It blocks until the connection is closed or an error occurs.
func (c *Client) runLoop() error {
	c.mu.Lock()
	conn := c.conn
	c.mu.Unlock()

	if conn == nil {
		return fmt.Errorf("soupbin/client: not connected")
	}

	fr := frame.NewReader(conn)

	hbTicker := time.NewTicker(c.cfg.HeartbeatInterval)
	defer hbTicker.Stop()

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	defer func() {
		_ = c.closeConnection(false)
		wg.Wait()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			f, err := fr.ReadFrame()
			if err != nil {
				errCh <- err
				return
			}
			c.setLastRecv(time.Now())
			if err := c.queue.Enqueue(f); err != nil {
				errCh <- err
				return
			}
		}
	}()

	c.setLastRecv(time.Now())

	for {
		select {
		case <-c.done:
			c.logger.Info("soupbin client loop stopped")
			return nil

		case err := <-errCh:
			c.handler.OnError(err)
			return err

		case <-hbTicker.C:
			if time.Since(c.getLastRecv()) > c.cfg.ServerTimeout {
				err := fmt.Errorf("soupbin/client: server timeout (no data for %s)", c.cfg.ServerTimeout)
				c.handler.OnError(err)
				return err
			}
			if time.Since(c.getLastSend()) < c.cfg.HeartbeatInterval {
				continue
			}
			if err := c.sendMsg(&message.ClientHeartbeat{}); err != nil {
				c.handler.OnError(err)
				return err
			}
		}
	}
}

func (c *Client) consumeFrames() {
	dec := codec.NewDecoder()

	for {
		f, err := c.queue.Dequeue()
		if err != nil {
			return
		}

		msg, err := dec.Decode(f)
		if err != nil {
			c.logger.Error("soupbin decode error", "err", err)
			c.handler.OnError(err)
			continue
		}

		c.dispatch(msg)
	}
}

func (c *Client) dispatch(msg message.Message) {
	switch m := msg.(type) {
	case *message.SequencedData:
		seq := c.nextSeq
		c.nextSeq++
		c.handler.OnSequencedData(seq, m.Message)

	case *message.ServerHeartbeat:
		if h, ok := c.handler.(ServerHeartbeatHandler); ok {
			h.OnServerHeartbeat()
		}

	case *message.EndOfSession:
		c.handler.OnEndOfSession()

	case *message.UnsequencedData:
		if h, ok := c.handler.(UnsequencedHandler); ok {
			h.OnUnsequencedData(m)
		}
	}
}

func (c *Client) sendMsg(msg message.Message) error {
	f, err := c.enc.Encode(msg)
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.fw.WriteFrame(f); err != nil {
		return err
	}
	c.lastSend = time.Now()
	return nil
}

func (c *Client) closeConnection(sendLogout bool) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var closeErr error
	if sendLogout && c.fw != nil {
		if err := c.fw.WriteFrame(&frame.Frame{
			Type:    message.PacketTypeLogoutRequest,
			Payload: nil,
		}); err != nil && !errors.Is(err, net.ErrClosed) {
			closeErr = err
		}
	}

	if c.conn != nil {
		if err := c.conn.Close(); err != nil && !errors.Is(err, net.ErrClosed) && closeErr == nil {
			closeErr = err
		}
	}

	c.conn = nil
	c.fw = nil
	return closeErr
}

func (c *Client) waitReconnect(ctx context.Context, delay time.Duration) bool {
	select {
	case <-c.done:
		return false
	case <-ctx.Done():
		return false
	case <-time.After(delay):
		return true
	}
}

func (c *Client) getLastSend() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.lastSend
}

func (c *Client) getLastRecv() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.lastRecv
}

func (c *Client) setLastRecv(t time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lastRecv = t
}

func isReconnectable(err error) bool {
	return errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF)
}
