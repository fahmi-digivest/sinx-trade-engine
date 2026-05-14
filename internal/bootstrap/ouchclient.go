package bootstrap

import (
	"context"
	"errors"
	"log/slog"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/infrastructure/config"
	"github.com/fahmi-digivest/sinx-trade-engine/internal/infrastructure/logger"
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

	<-a.ctx.Done()
	if err := a.ctx.Err(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}
