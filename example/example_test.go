package main

import (
	"time"

	"github.com/crunchydosa123/gomon"
)

func main() {
	gm := gomon.New()

	http_svc := gomon.NewHTTPService(
		"Example",
		"https://example.com/api",
		2*time.Second,
	)

	gm.Register(http_svc)
}
