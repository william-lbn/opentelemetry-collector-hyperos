[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# metricreceiver

## Metrics

These are the metrics available for this scraper.

| Name | Description | Unit | Type | Attributes |
| ---- | ----------- | ---- | ---- | ---------- |
| **system.cpu.time** | Total CPU seconds broken down by different states. Additional information on CPU Time can be found [here](https://en.wikipedia.org/wiki/CPU_time). | s | Sum(Double) | <ul> <li>host</li> <li>cpu_type</li> </ul> |

**Highlighted metrics** are emitted by default. Other metrics are optional and not emitted by default.
Any metric can be enabled or disabled with the following scraper configuration:

```yaml
metrics:
  <metric_name>:
    enabled: <true|false>
```

## Metric attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| cpu_type (type) | The type of CPU consumption | user, io_wait, system |
| host | The type of CPU consumption |  |
