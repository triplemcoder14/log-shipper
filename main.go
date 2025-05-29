package main

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"

	filelogreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver"
	otlpreceiver "go.opentelemetry.io/collector/receiver/otlpreceiver"

	batchprocessor "go.opentelemetry.io/collector/processor/batchprocessor"
	k8sattributesprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"

	otlpexporter "go.opentelemetry.io/collector/exporter/otlpexporter"
)

func main() {
// 	factories := component.Factories{
    factories, err := component.MakeFactories(
    		[]component.ReceiverFactory{
    			filelogreceiver.NewFactory(),
    			otlpreceiver.NewFactory(),
    		},
    		[]component.ProcessorFactory{
    			batchprocessor.NewFactory(),
    			k8sattributesprocessor.NewFactory(),
    		},
    		[]component.ExporterFactory{
    			otlpexporter.NewFactory(),
    		},
    		[]component.ExtensionFactory{},
    	)
    	if err != nil {
    		log.Fatalf("failed to make factories: %v", err)
    	}

// 		Receivers: map[component.Type]component.ReceiverFactory{
// 			filelogreceiver.Type(): filelogreceiver.NewFactory(),
// 			otlpreceiver.Type():    otlpreceiver.NewFactory(),
// 		},
// 		Processors: map[component.Type]component.ProcessorFactory{
// 			batchprocessor.Type():             batchprocessor.NewFactory(),
// 			k8sattributesprocessor.Type():     k8sattributesprocessor.NewFactory(),
// 		},
// 		Exporters: map[component.Type]component.ExporterFactory{
// 			otlpexporter.Type(): otlpexporter.NewFactory(),
// 		},
// 		Extensions: map[component.Type]component.ExtensionFactory{},
// 	}
	appSettings := service.CollectorSettings{
		Factories: factories,
		BuildInfo: component.BuildInfo{
			Command:     "otel-custom-collector",
			Description: "Custom OpenTelemetry Collector with Filelog and OTLP",
			Version:     "1.0.0",
		},
	}

	cmd := service.NewCommand(appSettings)

	if err := cmd.Execute(); err != nil {
		log.Fatalf("collector run failed: %v", err)
	}
}

// 	cmd := service.NewCommand(
// 		service.CollectorSettings{
// 			Factories: factories,
// 		},
// 	)
//
// 	if err := cmd.Execute(); err != nil {
// 		panic(err)
// 	}
// }