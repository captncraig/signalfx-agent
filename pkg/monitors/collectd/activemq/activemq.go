// +build linux

// Package activemq has an ActiveMQ Collectd monitor that uses GenericJMX
package activemq

import (
	"github.com/signalfx/signalfx-agent/pkg/monitors"
	"github.com/signalfx/signalfx-agent/pkg/monitors/collectd/genericjmx"
	yaml "gopkg.in/yaml.v2"
)

var serviceName = "activemq"

// Monitor is the main type that represents the monitor
type Monitor struct {
	*genericjmx.JMXMonitorCore
}

func init() {
	var defaultMBeans genericjmx.MBeanMap
	err := yaml.Unmarshal([]byte(defaultMBeanYAML), &defaultMBeans)
	if err != nil {
		panic("YAML for GenericJMX MBeans is invalid: " + err.Error())
	}
	defaultMBeans = defaultMBeans.MergeWith(genericjmx.DefaultMBeans)

	monitors.Register(&monitorMetadata, func() interface{} {
		return &Monitor{
			genericjmx.NewJMXMonitorCore(defaultMBeans, serviceName),
		}
	}, &genericjmx.Config{})
}
