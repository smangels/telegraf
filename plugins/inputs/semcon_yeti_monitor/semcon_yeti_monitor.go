package semcon_yeti_monitor

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

var SemconYetiMonitorConfig = `
	## Set the amplitude
	amplitude = 10.0
`

type SemconYetiMonitor struct {
	x			float64
	Amplitude 	float64
}

func (s *SemconYetiMonitor) SampleConfig() string {
	return SemconYetiMonitorConfig
}

func (s *SemconYetiMonitor) Description() string {
	return "Yeti monitoring plugin"
}

func (s *SemconYetiMonitor) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("semcon_yeti_monitor", func() telegraf.Input{ return &SemconYetiMonitor{x: 0.0} })
}
