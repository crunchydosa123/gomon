package gomon

import (
	"context"
	"time"
)

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

func (g *Gomon) Register(s Service) {
	g.services = append(g.services, s)
}

func (g *Gomon) Start(ctx context.Context) {
	/*for _, svc := range g.services {
		polling.Start(ctx, svc, g.metricsCh, g.cfg.timeout)
	}*/
}

func (g *Gomon) Stop() {
	close(g.stopCh)
	close(g.metricsCh)
}

func (g *Gomon) Subscribe() <-chan Metric {
	return g.metricsCh
}
