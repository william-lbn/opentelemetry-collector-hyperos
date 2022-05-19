package hypermonitorrelationmetricprocessor

import "go.opentelemetry.io/collector/config"

type Config struct {
	config.ProcessorSettings `mapstructure:",squash"`
	// Set of rules for generating new metrics
	Rules []Rule `mapstructure:"rules"`
}

type Rule struct {
	// unify metrics names.
	UnifyMetrics []string `mapstructure:"unify_metrics"`

	// description.
	Description string `mapstructure:"description"`
}
