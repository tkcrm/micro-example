package main

import (
	"context"
	"time"

	"github.com/tkcrm/mx-example/internal/api"
	"github.com/tkcrm/mx-example/internal/config"
	"github.com/tkcrm/mx-example/internal/services/authors"
	"github.com/tkcrm/mx-example/internal/services/books"
	"github.com/tkcrm/mx/cfg"
	"github.com/tkcrm/mx/launcher"
	"github.com/tkcrm/mx/logger"
	"github.com/tkcrm/mx/service"
	"github.com/tkcrm/mx/service/pingpong"
	"github.com/tkcrm/mx/transport/grpc_transport"
)

var (
	appName = "micro-example"
	version = "local"
)

func main() {
	logger := logger.NewExtended(
		logger.WithAppVersion(version),
		logger.WithAppName(appName),
	)

	conf := new(config.Config)
	if err := cfg.Load(conf, cfg.WithVersion(version)); err != nil {
		logger.Fatalf("could not load configuration: %s", err)
	}

	ln := launcher.New(
		launcher.WithName(appName),
		launcher.WithLogger(logger),
		launcher.WithVersion(version),
		launcher.WithRunnerServicesSequence(launcher.RunnerServicesSequenceFifo),
		launcher.WithOpsConfig(conf.Ops),
		launcher.WithAppStartStopLog(true),
	)

	// custom service
	customSvc := service.New(
		service.WithName("custom-service"),
		service.WithStart(func(ctx context.Context) error {
			logger.Info("hello")
			<-ctx.Done()
			logger.Info("goodbye")
			return nil
		}),
		service.WithStop(func(_ context.Context) error {
			return nil
		}),
	)

	// services
	booksService := books.New(logger)
	authorsService := authors.New()

	// grpc servers
	authorGrpcServer := api.NewAuthorServer()

	// grpc instance
	grpcServer := grpc_transport.NewServer(
		grpc_transport.WithLogger(logger),
		grpc_transport.WithConfig(conf.Grpc),
		grpc_transport.WithServices(authorGrpcServer),
	)

	ln.ServicesRunner().Register(
		customSvc,
		service.New(service.WithService(grpcServer)),
		service.New(service.WithService(booksService)),
		service.New(service.WithService(authorsService)),
		service.New(service.WithService(pingpong.New(logger))),
	)

	// shutdown after 1 minute
	go func() {
		<-time.After(time.Minute)
		logger.Info("Shutdown example: shutting down service")
		ln.Stop()
	}()

	if err := ln.Run(); err != nil {
		logger.Fatal(err)
	}
}
