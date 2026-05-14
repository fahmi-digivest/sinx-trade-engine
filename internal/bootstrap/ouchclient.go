package bootstrap

import "context"

type OuchClientApp struct {
	ctx context.Context
}

func NewOuchClientApp(ctx context.Context) (*OuchClientApp, error) {
	app := &OuchClientApp{
		ctx: ctx,
	}
	return app, nil
}

func (a *OuchClientApp) Start() error {
	return nil
}
