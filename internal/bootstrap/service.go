package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
)

// ServiceStartup is the shared contract for runtime service that can be run.
//
// Implement this interface for any long-running service (TCP feed, HTTP server,
// background worker, etc.) that should be managed by ServiceManager.
//
// Usage:
//
//	type MyService struct{}
//
//	func (s *MyService) Name() string { return "my-service" }
//
//	func (s *MyService) Run(ctx context.Context) error {
//	    // block until ctx is cancelled or a fatal error occurs
//	    <-ctx.Done()
//	    return nil
//	}
type ServiceStartup interface {
	Name() string
	Run(ctx context.Context) error
}

// ServiceManager runs a set of ServiceStartup concurrently and waits for all
// of them to finish. If any service returns an error, the context is cancelled
// so the remaining services are signalled to stop.
//
// Usage:
//
//	manager := bootstrap.NewServiceManager(
//	    logger,
//	    &MBOFeedService{client: mboClient},
//	    &MDFFeedService{client: mdfClient},
//	    &HTTPServer{addr: ":8080"},
//	)
//
//	if err := manager.Run(ctx); err != nil {
//	    log.Fatal(err)
//	}
//
// The parent ctx should be the application root context (e.g. from signal.NotifyContext).
// When the context is cancelled (SIGINT/SIGTERM), all services receive the cancellation
// and Run returns once every service has exited.
type ServiceManager struct {
	services []ServiceStartup
	logger   *slog.Logger
}

// NewServiceManager returns a ServiceManager with the given services.
func NewServiceManager(logger *slog.Logger, services ...ServiceStartup) *ServiceManager {
	return &ServiceManager{
		services: services,
		logger:   logger,
	}
}

// Run starts all services concurrently. It blocks until all services exit.
// Returns the first non-nil error encountered, if any.
func (m *ServiceManager) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error, len(m.services))

	var wg sync.WaitGroup
	for _, svc := range m.services {
		wg.Add(1)
		go func(s ServiceStartup) {
			defer wg.Done()

			if err := s.Run(ctx); err != nil {
				m.logger.Error("service stopped with error", "service", s.Name(), "err", err)
				errCh <- fmt.Errorf("service %q: %w", s.Name(), err)
				cancel() // signal other services to stop
				return
			}
		}(svc)
	}

	wg.Wait()
	close(errCh)

	// Return the first error, if any.
	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
