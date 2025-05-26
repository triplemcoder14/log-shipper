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
	factories := component.Factories{
		Receivers: map[component.Type]component.ReceiverFactory{
			filelogreceiver.Type(): filelogreceiver.NewFactory(),
			otlpreceiver.Type():    otlpreceiver.NewFactory(),
		},
		Processors: map[component.Type]component.ProcessorFactory{
			batchprocessor.Type():             batchprocessor.NewFactory(),
			k8sattributesprocessor.Type():     k8sattributesprocessor.NewFactory(),
		},
		Exporters: map[component.Type]component.ExporterFactory{
			otlpexporter.Type(): otlpexporter.NewFactory(),
		},
		Extensions: map[component.Type]component.ExtensionFactory{},
	}

	cmd := service.NewCommand(
		service.CollectorSettings{
			Factories: factories,
		},
	)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}