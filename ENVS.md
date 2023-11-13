# Environments

| Name                        | Required | Secret | Default value  | Usage                                                                         | Example          |
|-----------------------------|----------|--------|----------------|-------------------------------------------------------------------------------|------------------|
| `SERVICE_NAME`              | ✅        |        | `mx-example`   |                                                                               |                  |
| `OPS_ENABLED`               |          |        | `false`        | allows to enable ops server                                                   | `false`          |
| `OPS_NETWORK`               | ✅        |        | `tcp`          | allows to set ops listen network: tcp/udp                                     | `tcp`            |
| `OPS_TRACING_ENABLED`       |          |        | `false`        | allows to enable tracing                                                      | `false`          |
| `OPS_METRICS_ENABLED`       |          |        | `true`         | allows to enable metrics                                                      | `true`           |
| `OPS_METRICS_PATH`          | ✅        |        | `/metrics`     | allows to set custom metrics path                                             | `/metrics`       |
| `OPS_METRICS_PORT`          | ✅        |        | `10000`        | allows to set custom metrics port                                             | `10000`          |
| `OPS_HEALTHY_ENABLED`       |          |        | `true`         | allows to enable health checker                                               | `true`           |
| `OPS_HEALTHY_PATH`          | ✅        |        | `/healthy`     | allows to set custom healthy path                                             | `/healthy`       |
| `OPS_HEALTHY_PORT`          | ✅        |        | `10000`        | allows to set custom healthy port                                             | `10000`          |
| `OPS_PROFILER_ENABLED`      |          |        | `false`        | allows to enable profiler                                                     | `false`          |
| `OPS_PROFILER_PATH`         | ✅        |        | `/debug/pprof` | allows to set custom profiler path                                            | `/debug/pprof`   |
| `OPS_PROFILER_PORT`         | ✅        |        | `10000`        | allows to set custom profiler port                                            | `10000`          |
| `GRPC_ENABLED`              |          |        | `true`         | allows to enable grpc server                                                  | `true`           |
| `GRPC_ADDR`                 | ✅        |        | `:9000`        | grpc server listen address                                                    | `localhost:9000` |
| `GRPC_NETWORK`              | ✅        |        | `tcp`          | grpc server listen network: tpc/udp                                           | `tcp`            |
| `GRPC_REFLECT_ENABLED`      |          |        | `false`        | allows to enable grpc reflection service                                      | `false`          |
| `GRPC_HEALTH_CHECK_ENABLED` |          |        | `false`        | allows to enable grpc health checker                                          | `false`          |
| `GRPC_LOGGER_ENABLED`       |          |        | `false`        | allows to enable logger. available only for default grpc sevrer               | `false`          |
| `GRPC_RECOVERY_ENABLED`     |          |        | `false`        | allows to enable recovery from panics. available only for default grpc sevrer | `false`          |
