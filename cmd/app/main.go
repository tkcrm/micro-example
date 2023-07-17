package main

import (
	"context"
	"time"

	"github.com/tkcrm/micro/launcher"
	"github.com/tkcrm/micro/logger"
	"github.com/tkcrm/micro/service"
	"github.com/tkcrm/micro/service/pingpong"
)

var version = "local"
var appName = "micro-example"

func main() {
	logger := logger.New(
		logger.WithAppVersion(version),
		logger.WithAppName(appName),
	)

	ln := launcher.New(
		launcher.WithName(appName),
		launcher.WithLogger(logger),
		launcher.WithVersion(version),
		launcher.WithContext(context.Background()),
		launcher.WithAfterStart(func() error {
			logger.Infoln("app", appName, "was started")
			return nil
		}),
		launcher.WithAfterStop(func() error {
			logger.Infoln("app", appName, "was stopped")
			return nil
		}),
	)

	svc := service.New(
		service.WithName("test-service"),
		service.WithStart(func(_ context.Context) error {
			return nil
		}),
		service.WithStop(func(_ context.Context) error {
			time.Sleep(time.Second * 1)
			return nil
		}),
	)

	disabledService := service.New(
		service.WithName("disabled-service"),
		service.WithStart(func(_ context.Context) error {
			return nil
		}),
		service.WithStop(func(_ context.Context) error {
			return nil
		}),
		service.WithEnabled(false),
	)

	// without stop func
	invalidService := service.New(
		service.WithName("invalid-service"),
		service.WithStart(func(_ context.Context) error {
			return nil
		}),
	)

	pingPongSvc := service.New(
		service.WithService(pingpong.New(
			logger,
			pingpong.WithTimeout(time.Millisecond*200),
		)),
	)

	ln.ServicesRunner().Register(pingPongSvc, svc, disabledService, invalidService)

	// shutdown after 1 seconds
	go func() {
		<-time.After(time.Second * 1)
		logger.Info("Shutdown example: shutting down service")
		ln.Stop()
	}()

	if err := ln.Run(); err != nil {
		logger.Fatal(err)
	}
}
