package gomon

import "time"

type Service interface {
	Name() string
	Interval() time.Duration
	Check() (Metric, error)
}

type Metric struct {
	ServiceName string
	Status      int
	Latency     time.Duration
	Error       error
	Timestamp   time.Time
}
