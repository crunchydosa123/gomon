package polling

import (
	"context"
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
				// do HTTP check

			case <-ctx.Done():
				return
			}
		}
	}()
}
