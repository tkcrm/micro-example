package main

import (
	"context"
	"errors"
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
			logger.Infoln("service", appName, "started")
			return nil
		}),
		launcher.AfterStop(func() error {
			logger.Infoln("service", appName, "stopped")
			return nil
		}),
	)

	svc := service.New(
		service.WithName("test-service"),
		service.WithStart(func(_ context.Context) error {
			return nil
		}),
		service.WithStop(func(_ context.Context) error {
			return errors.New("test")
		}),
	)

	pingPongSvc := service.New(
		service.WithService(pingpong.New(logger, time.Second*5)),
	)

	ln.ServicesRunner().Register(svc)
	ln.ServicesRunner().Register(pingPongSvc)

	// shutdown after 16 seconds
	go func() {
		<-time.After(time.Second * 16)
		logger.Info("Shutdown example: shutting down service")
		ln.Stop()
	}()

	if err := ln.Run(); err != nil {
		logger.Fatal(err)
	}
}
