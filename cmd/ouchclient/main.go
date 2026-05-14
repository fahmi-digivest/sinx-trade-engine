package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/bootstrap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app, err := bootstrap.NewOuchClientApp(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize guse: %v\n", err)
		os.Exit(1)
	}

	if err := app.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "ouch client stopped with error: %v\n", err)
		os.Exit(1)
	}
}
