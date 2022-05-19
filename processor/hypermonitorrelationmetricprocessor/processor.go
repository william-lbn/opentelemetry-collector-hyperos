package hypermonitorrelationmetricprocessor

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"go.uber.org/zap"
)

type MetricsRelationProcessor struct {
	Rules  []internalRule
	Logger *zap.Logger
}

type internalRule struct {
	// unify metrics names.
	UnifyMetrics []string

	// description.
	Description string
}

func newMetricsRelationProcessor(rules []internalRule, logger *zap.Logger) *MetricsRelationProcessor {
	return &MetricsRelationProcessor{
		Rules:  rules,
		Logger: logger,
	}
}

func (mgp *MetricsRelationProcessor) Start(context.Context, component.Host) error {
	return nil
}

func (mgp *MetricsRelationProcessor) processMetrics(_ context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	resourceMetricsSlice := md.ResourceMetrics()

	for i := 0; i < resourceMetricsSlice.Len(); i++ {
		rm := resourceMetricsSlice.At(i)

		nameToMetricMap := getNameToMetricMap(rm)

		for _, rule := range mgp.Rules {
			for _, metricKey := range rule.UnifyMetrics {
				if metricsData, ok := nameToMetricMap[metricKey]; ok {
					mgp.Logger.Info("handle metrics data", zap.String("key=", metricKey))
					addAttributes(metricsData)
				}
			}
		}
	}
	return md, nil
}
