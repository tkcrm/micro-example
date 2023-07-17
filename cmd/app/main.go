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
		launcher.AfterStart(func() error {
			logger.Infoln("app", appName, "was started")
			return nil
		}),
		launcher.AfterStop(func() error {
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
			time.Sleep(time.Second * 3)
			return nil
		}),
		service.AfterStartFinished(func() error {
			logger.Infoln("service test-service was finished")
			return nil
		}),
	)

	svc2 := service.New(
		service.WithName("disabled-service"),
		service.WithStart(func(_ context.Context) error {
			return nil
		}),
		service.WithStop(func(_ context.Context) error {
			return nil
		}),
		service.WithEnabled(false),
	)

	pingPongSvc := service.New(
		service.WithService(pingpong.New(logger, time.Millisecond*200)),
	)

	ln.ServicesRunner().Register(svc, svc2, pingPongSvc)

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
