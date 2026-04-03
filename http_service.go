package gomon

import (
	"net/http"
	"time"
)

type HTTPService struct {
	name     string
	url      string
	interval time.Duration
	timeout  time.Duration
}

func NewHTTPService(name, url string, interval time.Duration) *HTTPService {
	return &HTTPService{
		name:     name,
		url:      url,
		interval: interval,
		timeout:  5 * time.Second,
	}
}

func (h *HTTPService) Name() string {
	return h.name
}

func (h *HTTPService) Interval() time.Duration {
	return h.interval
}

func (h *HTTPService) Check() (Metric, error) {
	start := time.Now()

	client := http.Client{Timeout: h.timeout}
	resp, err := client.Get(h.url)
	latency := time.Since(start)

	if err != nil {
		return Metric{
			ServiceName: h.name,
			Status:      0,
			Latency:     latency,
			Error:       err,
			Timestamp:   time.Now(),
		}, err
	}
	defer resp.Body.Close()

	return Metric{
		ServiceName: h.name,
		Status:      resp.StatusCode,
		Latency:     latency,
		Timestamp:   time.Now(),
	}, nil
}
