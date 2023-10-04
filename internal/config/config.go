package config

import (
	"github.com/tkcrm/mx/ops"
	"github.com/tkcrm/mx/transport/grpc_transport"
)

type Config struct {
	ServiceName string `default:"mx-example" validate:"required"`
	Ops         ops.Config
	Grpc        grpc_transport.Config
}
