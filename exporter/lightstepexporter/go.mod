module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/lightstepexporter

go 1.14

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/golang/protobuf v1.4.2
	github.com/lightstep/opentelemetry-exporter-go v0.6.3
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.8.1-0.20200819173546-64befbcc0564
	go.opentelemetry.io/otel v0.17.0 // indirect
	go.opentelemetry.io/otel/sdk v0.17.0 // indirect
	go.uber.org/zap v1.15.0
	google.golang.org/grpc/examples v0.0.0-20200728194956-1c32b02682df // indirect
)
