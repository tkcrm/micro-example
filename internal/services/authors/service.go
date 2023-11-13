package authors

import (
	"context"
	"math/rand"
	"time"

	"github.com/tkcrm/mx/ops"
	"github.com/tkcrm/mx/service"
)

type authors struct {
	name string
}

func New() *authors {
	return &authors{
		name: "authors-service",
	}
}

func (s authors) Name() string { return s.name }

func (s authors) Healthy(ctx context.Context) error {
	// random healtch check status
	min, max := 10, 100
	value := rand.Intn(max-min) + min
	if value > 90 {
		// you can return any error message
		return ops.ErrHealthCheckError
	}
	return nil
}

func (s authors) Interval() time.Duration { return time.Second }

func (s authors) Start(ctx context.Context) error {
	<-ctx.Done()
	return nil
}

func (s authors) Stop(ctx context.Context) error { return nil }

var _ service.HealthChecker = (*authors)(nil)

var _ service.IService = (*authors)(nil)
