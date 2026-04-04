# Gomon 🚀

A lightweight, concurrent monitoring library in Go for building custom health checks and observability pipelines.

---

## ✨ Features

* 🔌 Plugin-based architecture (bring your own services)
* ⚡ Concurrent polling with goroutines
* 📡 Real-time metrics streaming via channels
* 🧱 Clean separation of engine and implementations
* 🛠️ Extensible (HTTP, TCP, DB, custom checks)

---

## 📦 Installation

```bash
go get github.com/crunchydosa123/gomon@v1.0.0
```

---

## 🚀 Quick Start

```go
package main

import (
	"fmt"
	"time"

	"github.com/crunchydosa123/gomon"
)

func main() {
	g := gomon.New()

	// Built-in HTTP service
	svc := gomon.NewHTTPService(
		"google",
		"https://google.com",
		2*time.Second,
	)

	g.Register(svc)

	go g.Start()

	for m := range g.Subscribe() {
		fmt.Println(m)
	}
}
```

---

## 🧠 Core Concepts

### Service Interface

At the heart of Gomon is the `Service` interface:

```go
type Service interface {
	Name() string
	Interval() time.Duration
	Check() (Metric, error)
}
```

Any type implementing this can be monitored.

---

### Metric

```go
type Metric struct {
	ServiceName string
	Status      int
	Latency     time.Duration
	Error       error
	Timestamp   time.Time
}
```

---

## 🔌 Creating Custom Services

You can define your own service by implementing the interface.

### Example: TCP Service

```go
type TCPService struct {
	name     string
	address  string
	interval time.Duration
	timeout  time.Duration
}

func (t TCPService) Name() string { return t.name }
func (t TCPService) Interval() time.Duration { return t.interval }

func (t TCPService) Check() (gomon.Metric, error) {
	start := time.Now()

	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	latency := time.Since(start)

	if err != nil {
		return gomon.Metric{
			ServiceName: t.name,
			Status:      0,
			Latency:     latency,
			Error:       err,
			Timestamp:   time.Now(),
		}, err
	}
	conn.Close()

	return gomon.Metric{
		ServiceName: t.name,
		Status:      1,
		Latency:     latency,
		Timestamp:   time.Now(),
	}, nil
}
```

---

## 🏗️ Architecture

```
gomon/
├── gomon.go          # public API
├── http_service.go   # built-in services
├── internal/
│   ├── core/         # shared types (Service, Metric)
│   └── polling/      # execution engine
```

---

## 🛑 Stopping the Engine

```go
g.Stop()
```

Gracefully shuts down all polling workers.

---

## 🔮 Roadmap

* ⏱️ Retry & timeout strategies
* 📊 Prometheus exporter
* 🧩 Middleware support (logging, alerts)
* 🖥️ CLI (`gomon run config.yaml`)

---

## 🤝 Contributing

PRs and ideas are welcome! Feel free to open an issue or contribute new service implementations.

---

## 📄 License

MIT
