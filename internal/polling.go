package polling

import (
	"context"
	"net/http"
	"time"

	"github.com/crunchydosa123/gomon"
)

func Start(
	ctx context.Context,
	svc gomon.Service,
	out chan<- gomon.Metric,
	timeout time.Duration,
) {
	ticker := time.NewTicker(svc.Interval)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				start := time.Now()

				resp, err := http.Get(svc.URL)

				metric := gomon.Metric{
					ServiceName: svc.Name,
					Timestamp:   time.Now(),
				}

				if err != nil {
					metric.Error = err
				} else {
					metric.Status = resp.StatusCode
					metric.Latency = time.Since(start)
					resp.Body.Close()
				}

				out <- metric

			case <-ctx.Done():
				return
			}
		}
	}()
}
