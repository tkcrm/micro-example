package books

import (
	"context"
	"time"

	"github.com/tkcrm/mx/service"
)

type books struct {
	name       string
	hcInterval time.Duration
}

func New() *books {
	return &books{
		name:       "books-service",
		hcInterval: time.Second * 3,
	}
}

func (s books) Name() string { return s.name }

func (s books) Healthy(ctx context.Context) error { return nil }

func (s books) Interval() time.Duration { return s.hcInterval }

func (s books) Start(ctx context.Context) error {
	<-ctx.Done()
	return nil
}

func (s books) Stop(ctx context.Context) error { return nil }

var _ service.HealthChecker = (*books)(nil)

var _ service.IService = (*books)(nil)
