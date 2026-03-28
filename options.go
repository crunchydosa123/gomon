package gomon

import "time"

type config struct {
	buffer  int
	timeout time.Duration
	retries int
}

type Option func(*config)

func WithBuffer(size int) Option {
	return func(c *config) {
		c.buffer = size
	}
}

func WithTimeout(d time.Duration) Option {
	return func(c *config) {
		c.timeout = d
	}
}

func WithRetries(n int) Option {
	return func(c *config) {
		c.retries = n
	}
}
