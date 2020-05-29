package cli

import atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"

type MetricsOpts struct {
	ListOpts
	Granularity     string
	Period          string
	Start           string
	End             string
	MeasurementType []string
}

func (opts *MetricsOpts) NewProcessMetricsListOptions() *atlas.ProcessMeasurementListOptions {
	o := &atlas.ProcessMeasurementListOptions{
		ListOptions: opts.NewListOptions(),
	}
	o.Granularity = opts.Granularity
	o.Period = opts.Period
	o.Start = opts.Start
	o.End = opts.End
	o.M = opts.MeasurementType

	return o
}
