# 🌐 Distributed Rate Limiter + API Gateway in Go

A production-grade, horizontally scalable API Gateway built in Go, featuring pluggable load balancing, distributed rate limiting, middleware architecture, Prometheus observability, and high-concurrency performance engineering.

This project is designed to deeply explore concurrency, distributed coordination, and high-performance backend systems — not just framework usage.

---

## 🧰 Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=for-the-badge)
![Redis](https://img.shields.io/badge/Redis-DC382D?logo=redis&logoColor=white&style=for-the-badge)
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white&style=for-the-badge)
![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?logo=prometheus&logoColor=white&style=for-the-badge)
![YAML](https://img.shields.io/badge/YAML-000000?logo=yaml&logoColor=white&style=for-the-badge)
![Git](https://img.shields.io/badge/Git-F05032?logo=git&logoColor=white&style=for-the-badge)

---

## 🚀 Features

- 🔁 High-performance reverse proxy
- ⚖️ Pluggable load balancing (Round Robin, Least Connections)
- 🧮 Distributed rate limiting (In-memory + Redis-backed)
- 🧵 Concurrency-safe architecture
- 📊 Prometheus metrics integration
- 🛡 Middleware chain (Logging, Auth, Rate Limit, Recovery)
- 🐳 Dockerized deployment
- 🧪 Unit tests + Benchmarks
- 📈 10k–15k+ requests/sec sustained locally under benchmark testing

---

## 🎯 Project Objectives

Design and implement a production-grade API Gateway that includes:

- Reverse proxying
- Load balancing strategies
- Distributed rate limiting
- Observability & metrics
- Fault tolerance mechanisms
- Benchmarks & test coverage
- Dockerized deployment
- Clean, idiomatic Go architecture

---

## 🧰 Tech Stack

### Core
- Go (`net/http`)
- Redis (rate limit coordination)
- Docker
- docker-compose

### Observability
- Prometheus metrics endpoint (`/metrics`)

### Testing
- `go test`
- `go test -race`
- Benchmarks (`go test -bench`)

### Configuration
- YAML parsing (`gopkg.in/yaml.v3`)

---

## 🏗 Architecture

```
cmd/gateway
internal/
    proxy/
    loadbalancer/
    ratelimiter/
    middleware/
    metrics/
pkg/
configs/
```

Design Principles:

- No circular dependencies
- Clean interfaces
- Thread-safe components
- Idiomatic Go structure
- Separation of concerns

---

## 🔁 Reverse Proxy

- Accept incoming HTTP requests
- Forward to upstream services
- Preserve headers
- Support timeouts
- Graceful error handling
- Handle concurrent requests safely

Target:
- Sustain 10k+ requests/sec locally without crashing

---

## ⚖️ Load Balancing

Strategies implemented:

- Round Robin
- Least Connections

Designed as a pluggable interface:

```go
type LoadBalancer interface {
    NextBackend() *Backend
}
```

Goals:

- O(1) backend selection
- Thread-safe
- No global lock bottlenecks

---

## 🧮 Distributed Rate Limiting

Algorithms:

- Token Bucket
- (Optional) Sliding Window

Modes:

- In-memory (single node)
- Redis-backed (distributed)

Goals:

- Correct refill math
- Accurate under concurrency
- No race conditions
- Support per-IP + per-route limits
- Accurate limiting at 10k concurrent requests

---

## 📊 Observability

Expose a Prometheus-compatible `/metrics` endpoint.

Track:

- Request latency
- Requests per second
- Error rate
- Backend health state

---

## 🛡 Middleware Chain

Implemented middleware:

- Logging
- Authentication (API key)
- Rate limiting
- Panic recovery

Goals:

- Extensible architecture
- Clear separation of concerns
- Clean and maintainable codebase

---

## 🧵 Concurrency Guarantees

- Goroutines per request
- Proper use of `sync.Mutex`, `sync.RWMutex`, and `atomic`
- Context cancellation support
- Graceful shutdown
- Zero data races (verified via `go test -race`)
- No deadlocks

---

## 🧪 Testing & Benchmarks

Includes:

- Unit tests (rate limiter, load balancer)
- Concurrency tests
- Integration tests
- Benchmarks using `go test -bench`

Target:

- 70%+ coverage
- Documented benchmark results in README

---

## 🐳 Deployment

- Fully Dockerized
- `docker-compose` with Redis
- CI via GitHub Actions
- Linting + `go vet`

Optional:
- Kubernetes deployment experiments

---

## 📈 Performance Metrics (To Be Documented)

- Max sustained RPS
- P95 latency
- CPU usage
- Memory usage
- System architecture diagram

---

## 🗓 Roadmap

### Week 1
- Reverse proxy
- Round robin load balancing
- In-memory rate limiter
- Clean project structure

### Week 2
- Redis-backed rate limiter
- Concurrency hardening
- Unit tests + benchmarks

### Week 3
- Consistent hashing
- Circuit breaker
- Metrics integration

### Week 4
- YAML configuration
- Docker setup
- Load testing
- Documentation polish

---

## 🧠 Engineering Questions This Project Explores

- Why token bucket over sliding window?
- Why mutex instead of channel?
- Tradeoffs of centralized Redis?
- What happens under network partition?
- How would this scale to 1M RPS?
- What breaks first?

---

## 📍 Current Progress

- Built basic HTTP server
- Tested request concurrency using ApacheBench
- Verified race conditions and synchronization
- Set up Ubuntu environment for load testing

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Implement improvements
4. Add tests
5. Submit a pull request

---

## 📚 Learning Goals

- Deep Go concurrency mastery
- Reverse proxy internals
- Token bucket mathematics
- Load balancing algorithms
- Circuit breaker pattern
- Distributed systems fundamentals

---

This project is built to move from “coder” to “systems engineer.”
