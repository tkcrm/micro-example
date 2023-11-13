package books

import (
	"context"
	"time"

	"github.com/tkcrm/mx/logger"
	"github.com/tkcrm/mx/ops"
	"github.com/tkcrm/mx/service"
)

const serviceName = "books-service"

type books struct {
	logger     logger.Logger
	name       string
	hcInterval time.Duration
	timeStart  time.Time
}

func New(l logger.Logger) *books {
	return &books{
		logger:     logger.With(l, "service", serviceName),
		name:       serviceName,
		hcInterval: time.Second * 1,
		timeStart:  time.Now(),
	}
}

func (s books) Name() string { return s.name }

func (s books) Healthy(ctx context.Context) error {
	// warming up the service
	if time.Since(s.timeStart) < time.Second*10 {
		return ops.ErrHealthCheckServiceStarting
	}
	return nil
}

func (s books) Interval() time.Duration { return s.hcInterval }

func (s books) Start(ctx context.Context) error {
	s.logger.Info("starting the service with warm-up time of 10 seconds")
	<-ctx.Done()
	return nil
}

func (s books) Stop(ctx context.Context) error { return nil }

var _ service.HealthChecker = (*books)(nil)

var _ service.IService = (*books)(nil)
