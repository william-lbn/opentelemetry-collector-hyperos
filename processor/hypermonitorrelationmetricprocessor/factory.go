package hypermonitorrelationmetricprocessor

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

const (
	typeStr = "hypermonitor_relationmetricprocessor"
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

func NewFactory() component.ProcessorFactory {

	//开启kubernetes informer协程，缓存当前k8s集群中deployment
	go KubernetesStart()

	return component.NewProcessorFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsProcessor(createMetricsProcessor))
}

func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
	}
}

func createMetricsProcessor(
	ctx context.Context,
	params component.ProcessorCreateSettings,
	cfg config.Processor,
	nextConsumer consumer.Metrics,
) (component.MetricsProcessor, error) {
	processorConfig, ok := cfg.(*Config)
	if !ok {
		return nil, fmt.Errorf("configuration parsing error")
	}

	processorConfig.Validate()
	metricsProcessor := newMetricsRelationProcessor(buildInternalConfig(processorConfig), params.Logger)
	return processorhelper.NewMetricsProcessor(
		cfg,
		nextConsumer,
		metricsProcessor.processMetrics,
		processorhelper.WithCapabilities(processorCapabilities))

}

// buildInternalConfig constructs the internal metric generation rules
func buildInternalConfig(config *Config) []internalRule {
	internalRules := make([]internalRule, len(config.Rules))

	for i, rule := range config.Rules {
		customRule := internalRule{
			UnifyMetrics: rule.UnifyMetrics,
			Description:  rule.Description,
		}
		internalRules[i] = customRule
	}
	return internalRules
}
