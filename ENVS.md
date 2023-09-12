# Environments

| Name                  | Required | Secret | Default value | Usage                                     | Example |
|-----------------------|----------|--------|---------------|-------------------------------------------|---------|
| `SERVICE_NAME`        | ✅        |        | mx-example    |                                           |         |
| `PROMETHEUS_PORT`     |          |        | 10001         |                                           |         |
| `PROMETHEUS_ENDPOINT` |          |        | /metrics      |                                           |         |
| `PROMETHEUS_ENABLED`  |          |        | true          |                                           |         |
| `OPS_ENABLED`         |          |        | false         | allows to enable ops server               |         |
| `OPS_ADDR`            |          |        | :10000        | allows to set set ops address:port        |         |
| `OPS_NETWORK`         |          |        | tcp           | allows to set ops listen network: tcp/udp |         |
| `OPS_NO_TRACE`        |          |        | true          | allows to disable tracing                 |         |
| `OPS_METRICS_PATH`    |          |        | /metrics      | allows to set custom metrics path         |         |
| `OPS_HEALTHY_PATH`    |          |        | /healthy      | allows to set custom healthy path         |         |
| `OPS_PROFILE_PATH`    |          |        | /debug/pprof  | allows to set custom profiler path        |         |
| `GRPC_ENABLED`        |          |        | true          | allows to enable grpc server              |         |
| `GRPC_REFLECT`        |          |        | false         | allows to enable grpc reflection service  |         |
| `GRPC_ADDR`           |          |        | :9000         | gRPC server listen address                |         |
| `GRPC_NETWORK`        |          |        | tcp           | gRPC server listen network: tpc/udp       |         |
