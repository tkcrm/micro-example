package config

import "github.com/tkcrm/mx/service/prometheus"

type Config struct {
	ServiceName string            `default:"mx-example" validate:"required"`
	Prometheus  prometheus.Config `env:"PROMETHEUS"`
}
