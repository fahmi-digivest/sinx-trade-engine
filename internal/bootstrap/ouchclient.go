package bootstrap

import (
	"context"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/infrastructure/config"
)

type OuchClientApp struct {
	ctx      context.Context
	config   *config.Config
	services []ServiceStartup
}

func NewOuchClientApp(ctx context.Context) (*OuchClientApp, error) {
	cfg, err := config.Load("ouchclient.yaml")
	if err != nil {
		return nil, err
	}

	app := &OuchClientApp{
		ctx:    ctx,
		config: cfg,
	}
	return app, nil
}

func (a *OuchClientApp) Start() error {
	return nil
}
