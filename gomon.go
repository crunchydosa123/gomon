package gomon

import "time"

type Gomon struct {
	services  []Service     // what to monitor
	metricsCh chan Metric   // pipeline output
	stopCh    chan struct{} // lifecycle control

	cfg config // user options
}

func New(opts ...Option) *Gomon {
	cfg := config{
		buffer:  100,
		timeout: 5 * time.Second,
		retries: 1,
	}

	for _, opt := range opts {
		opt(&cfg)
	}
	return &Gomon{
		services:  []Service{},
		metricsCh: make(chan Metric, cfg.buffer),
		cfg:       cfg,
	}
}

func (g *Gomon) Register(s Service)

func (g *Gomon) Start()

func (g *Gomon) Stop()

func (g *Gomon) Subscribe() <-chan Metric
