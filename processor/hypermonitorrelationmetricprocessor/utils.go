// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hypermonitorrelationmetricprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func getNameToMetricMap(rm pmetric.ResourceMetrics) map[string]pmetric.Metric {
	ilms := rm.ScopeMetrics()
	metricMap := make(map[string]pmetric.Metric)

	for i := 0; i < ilms.Len(); i++ {
		ilm := ilms.At(i)
		metricSlice := ilm.Metrics()
		for j := 0; j < metricSlice.Len(); j++ {
			metric := metricSlice.At(j)
			metricMap[metric.Name()] = metric
		}
	}
	return metricMap
}

func addAttributes(metric pmetric.Metric) {

	if metric.DataType() == pmetric.MetricDataTypeGauge {
		dataPoints := metric.Gauge().DataPoints()
		for i := 0; i < dataPoints.Len(); i++ {
			podName, _ := dataPoints.At(i).Attributes().Get("pod")
			//fmt.Print("pod name: ", podName) 加日志
			if value, ok := DeploymentMap[podName.StringVal()]; ok {
				//加日志

				metric.Gauge().DataPoints().At(i).Attributes().InsertString("component_name", value)
				metric.Gauge().DataPoints().At(i).Attributes().InsertString("component_id", DeploymentUidMap[value])
			}
		}
	}

}
