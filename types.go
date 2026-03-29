package gomon

import "time"

type Service struct {
	Name     string
	URL      string
	Interval time.Duration
}

type Metric struct {
	ServiceName string
	Status      int
	Latency     time.Duration
	Error       error
	Timestamp   time.Time
}
