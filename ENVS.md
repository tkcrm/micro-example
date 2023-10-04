# Environments

| Name                   | Required | Secret | Default value | Usage                                     | Example |
|------------------------|----------|--------|---------------|-------------------------------------------|---------|
| `SERVICE_NAME`         | ✅        |        | mx-example    |                                           |         |
| `OPS_ENABLED`          |          |        | false         | allows to enable ops server               |         |
| `OPS_NETWORK`          |          |        | tcp           | allows to set ops listen network: tcp/udp |         |
| `OPS_TRACING_ENABLED`  |          |        | false         | allows to enable tracing                  |         |
| `OPS_METRICS_ENABLED`  |          |        | true          | allows to enable metrics                  |         |
| `OPS_METRICS_PATH`     |          |        | /metrics      | allows to set custom metrics path         |         |
| `OPS_METRICS_PORT`     |          |        | 10000         | allows to set custom metrics port         |         |
| `OPS_HEALTHY_ENABLED`  |          |        | true          | allows to enable health checker           |         |
| `OPS_HEALTHY_PATH`     |          |        | /healthy      | allows to set custom healthy path         |         |
| `OPS_HEALTHY_PORT`     |          |        | 10000         | allows to set custom healthy port         |         |
| `OPS_PROFILER_ENABLED` |          |        | true          | allows to enable profiler                 |         |
| `OPS_PROFILER_PATH`    |          |        | /debug/pprof  | allows to set custom profiler path        |         |
| `OPS_PROFILER_PORT`    |          |        | 10000         | allows to set custom profiler port        |         |
| `GRPC_ENABLED`         |          |        | true          | allows to enable grpc server              |         |
| `GRPC_REFLECT`         |          |        | false         | allows to enable grpc reflection service  |         |
| `GRPC_ADDR`            |          |        | :9000         | gRPC server listen address                |         |
| `GRPC_NETWORK`         |          |        | tcp           | gRPC server listen network: tpc/udp       |         |
