package bootstrap

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"strconv"

	soupbinclient "github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/client"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/frame"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/message"
	ouchHandler "github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/ouch/handler"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/infrastructure/config"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/infrastructure/logger"
	infraqueue "github.com/fahmi-digivest/sinx-trade-engine/internal/infrastructure/queue"
)

type OuchClientApp struct {
	ctx      context.Context
	config   *config.Config
	loggers  *logger.Manager
	logger   *slog.Logger
	services []ServiceStartup
}

func NewOuchClientApp(ctx context.Context) (*OuchClientApp, error) {
	cfg, err := config.Load("ouchclient.yaml")
	if err != nil {
		return nil, err
	}

	loggers, err := logger.NewManager(cfg.Logger)
	if err != nil {
		return nil, err
	}

	serviceLogger := loggers.Get("service")
	slog.SetDefault(serviceLogger)

	app := &OuchClientApp{
		ctx:     ctx,
		config:  cfg,
		loggers: loggers,
		logger:  serviceLogger,
	}

	if err := app.serviceInitialize(); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *OuchClientApp) Start() error {
	a.logger.Info(
		"ouch client started",
		"app", a.config.App.Name,
		"env", a.config.App.Env,
		"version", a.config.App.Version,
		"tcp_clients", len(a.config.TCP.Clients),
	)

	defer func() {
		a.logger.Info(
			"ouch client stopped",
			"app", a.config.App.Name,
			"env", a.config.App.Env,
		)
	}()

	if len(a.services) == 0 {
		a.logger.Warn("no ouch client services enabled")
		<-a.ctx.Done()
		if err := a.ctx.Err(); err != nil && !errors.Is(err, context.Canceled) {
			return err
		}
		return nil
	}

	manager := NewServiceManager(a.logger, a.services...)
	if err := manager.Run(a.ctx); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (a *OuchClientApp) serviceInitialize() error {
	ouchLogger := a.loggers.Get("ouch")
	services := make([]ServiceStartup, 0, 2)

	if len(a.config.TCP.Clients) > 0 {
		cfg := a.config.TCP.Clients[0]
		if cfg.Run {
			name := cfg.Name

			handler := ouchHandler.NewOuchClientHandler(
				a.logger.With("component", "soupbin-handler", "client_name", name),
				ouchLogger.With("component", "soupbin-handler", "client_name", name),
			)

			clientCfg := soupbinclient.Config{
				ServerAddr:              net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
				Username:                cfg.Username,
				Password:                cfg.Password,
				RequestedSession:        cfg.RequestedSession,
				RequestedSequenceNumber: cfg.RequestedSequence,
				HeartbeatInterval:       cfg.HeartbeatInterval,
				ServerTimeout:           cfg.HeartbeatInterval * 15,
				DialTimeout:             cfg.DialTimeout,
				ReconnectDelay:          cfg.ReconnectDelay,
			}

			readQueue := infraqueue.NewSPSC[*frame.Frame](1024)
			writeQueue := infraqueue.NewSPSC[message.Message](1024)

			services = append(services, soupbinclient.New(
				name,
				clientCfg,
				handler,
				a.logger.With("component", "soupbin-client", "client_name", name),
				readQueue,
				writeQueue,
			))
		} else {
			a.logger.Info("skip disabled ouch client", "name", cfg.Name)
		}
	}

	if len(a.config.TCP.Clients) > 1 {
		cfg := a.config.TCP.Clients[1]
		if cfg.Run {
			name := cfg.Name

			handler := ouchHandler.NewOuchClientHandler(
				a.logger.With("component", "soupbin-handler", "client_name", name),
				ouchLogger.With("component", "soupbin-handler", "client_name", name),
			)

			clientCfg := soupbinclient.Config{
				ServerAddr:              net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
				Username:                cfg.Username,
				Password:                cfg.Password,
				RequestedSession:        cfg.RequestedSession,
				RequestedSequenceNumber: cfg.RequestedSequence,
				HeartbeatInterval:       cfg.HeartbeatInterval,
				ServerTimeout:           cfg.HeartbeatInterval * 15,
				DialTimeout:             cfg.DialTimeout,
				ReconnectDelay:          cfg.ReconnectDelay,
			}

			readQueue := infraqueue.NewSPSC[*frame.Frame](1024)
			writeQueue := infraqueue.NewSPSC[message.Message](1024)

			services = append(services, soupbinclient.New(
				name,
				clientCfg,
				handler,
				a.logger.With("component", "soupbin-client", "client_name", name),
				readQueue,
				writeQueue,
			))
		} else {
			a.logger.Info("skip disabled ouch client", "name", cfg.Name)
		}
	}

	a.services = services
	return nil
}
