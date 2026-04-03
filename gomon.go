package gomon

import (
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
		stopCh:    make(chan struct{}),
		cfg:       cfg,
	}
}

func (g *Gomon) Start() {
	for _, s := range g.services {
		go g.pollService(s)
	}
}

func (g *Gomon) Register(s Service) {
	g.services = append(g.services, s)
}

func (g *Gomon) Stop() {
	close(g.stopCh)
	close(g.metricsCh)
}

func (g *Gomon) Subscribe() <-chan Metric {
	return g.metricsCh
}

func (g *Gomon) pollService(s Service) {
	ticker := time.NewTicker(s.Interval())
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			metric, err := s.Check()
			if err != nil {
				metric.Error = err
			}

			select {
			case g.metricsCh <- metric:
			case <-g.stopCh:
				return
			}

		case <-g.stopCh:
			return
		}
	}
}
