package gitlab

import (
	pe "github.com/signalfx/signalfx-agent/internal/monitors/prometheusexporter"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

func init() {
	monitors.Register(&gitlabMonitorMetadata, func() interface{} {
		return &pe.Monitor{}
	}, &pe.Config{})

	monitors.Register(&gitlabRunnerMonitorMetadata, func() interface{} {
		return &pe.Monitor{}
	}, &pe.Config{})

	monitors.Register(&gitlabGitalyMonitorMetadata, func() interface{} {
		return &pe.Monitor{ExtraDimensions: map[string]string{
			"metric_source": "gitlab-gitaly"}}
	}, &pe.Config{})

	monitors.Register(&gitlabSidekiqMonitorMetadata, func() interface{} {
		return &pe.Monitor{}
	}, &pe.Config{})

	monitors.Register(&gitlabWorkhorseMonitorMetadata, func() interface{} {
		return &pe.Monitor{}
	}, &pe.Config{})

	// Send all unicorn metrics
	monitors.Register(&gitlabUnicornMonitorMetadata, func() interface{} { return &pe.Monitor{} }, &pe.Config{
		MetricPath: "/-/metrics",
	})
}
