package diskio

import (
	"time"

	"github.com/signalfx/signalfx-agent/pkg/core/config"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
	log "github.com/sirupsen/logrus"
)

var logger = log.WithFields(log.Fields{"monitorType": monitorType})

func init() {
	monitors.Register(&monitorMetadata, func() interface{} { return &Monitor{} }, &Config{})
}

// Config for this monitor
type Config struct {
	config.MonitorConfig `singleInstance:"false" acceptsEndpoints:"false"`
	// The devices to include/exclude. This is an [overridable
	// set](https://docs.signalfx.com/en/latest/integrations/agent/filtering.html#overridable-filters).
	Disks []string `yaml:"disks" default:"[\"*\", \"!/^loop[0-9]+$/\", \"!/^dm-[0-9]+$/\"]"`
	// (Windows Only) The frequency that wildcards in counter paths should
	// be expanded and how often to refresh counters from configuration.
	// This is expressed as a duration.
	CountersRefreshInterval time.Duration `yaml:"counterRefreshInterval" default:"60s"`
	// (Windows Only) Print out the configurations that match available
	// performance counters.  This used for debugging.
	PrintValid bool `yaml:"printValid"`
}

// Shutdown stops the metric sync
func (m *Monitor) Shutdown() {
	if m.cancel != nil {
		m.cancel()
	}
}
