package main

import (
	"time"

	"github.com/crunchydosa123/gomon"
)

func main() {
	gm := gomon.New()

	gm.Register(gomon.Service{
		Name:     "Google",
		URL:      "https://www.google.com",
		Interval: 5 * time.Second,
	})

	gm.Register(gomon.Service{
		Name:     "GitHub",
		URL:      "https://www.github.com",
		Interval: 10 * time.Second,
	})
}
