# 🌐 Distributed Rate Limiter + API Gateway in Go

A high-performance, scalable API Gateway built in Go to explore concurrency, distributed coordination, and backend engineering at scale. Designed with pluggable load balancing, Redis-backed rate limiting, middleware architecture, Prometheus observability, and high-concurrency performance.  

Perfect for engineers looking to deeply understand production-grade backend systems, scaling, containerization and reliability.

---

## 🧰 Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=for-the-badge)
![Redis](https://img.shields.io/badge/Redis-DC382D?logo=redis&logoColor=white&style=for-the-badge)
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white&style=for-the-badge)
![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?logo=prometheus&logoColor=white&style=for-the-badge)
![YAML](https://img.shields.io/badge/YAML-000000?logo=yaml&logoColor=white&style=for-the-badge)
![Linux](https://img.shields.io/badge/Linux-FCC624?logo=linux&logoColor=black&style=for-the-badge)
![Git](https://img.shields.io/badge/Git-F05032?logo=git&logoColor=white&style=for-the-badge)
![Kubernetes](https://img.shields.io/badge/Kubernetes-326CE5?logo=kubernetes&logoColor=white&style=for-the-badge)


---

## 🎯 Features

- 🔁 Reverse proxy with header propagation and timeout handling
- ⚖️ Pluggable load balancing (Round Robin, Least Connections)
- 🧮 Distributed rate limiting (Token Bucket + Sliding Window)
- 🧵 Concurrency-safe architecture (Goroutines, Mutex, Atomic)
- 📊 Prometheus metrics integration
- 🛡 Middleware chain (Logging, Auth, Rate Limit, Recovery)
- 🐳 Dockerized deployment with `docker-compose`
- 🧪 Unit tests + Benchmarks (`go test -bench`)
- 📈 10k–15k+ requests/sec sustained under benchmark testing

---

## 🛠 Local Setup

### 1. Clone Repository
```bash
git clone https://github.com/yourusername/go-api-gateway.git
cd go-api-gateway
```

### 2. Run Server
```bash
go run main.go
```
Server will run at:
```
http://localhost:8080
```

### 3. Load Testing Example
```bash
ab -n 10000 -c 100 http://localhost:8080/
```
Where:
- `-n` = total requests  
- `-c` = concurrent requests

> Works on **Windows**, **Linux**, and **Ubuntu**.

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

Key principles:

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
- Concurrency-safe

Target:
- Sustain 10k+ requests/sec locally without crashing

---

## ⚖️ Load Balancing

- Round Robin
- Least Connections

Pluggable interface:
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
- Sliding Window (optional)

Modes:
- In-memory (single-node)
- Redis-backed (distributed)

Goals:
- Correct refill math
- Accurate under concurrency
- Zero race conditions
- Per-IP + per-route limits

---

## 🧵 Concurrency & Middleware

- Goroutines per request
- Mutex / RWMutex and atomic operations
- Context cancellation
- Graceful shutdown
- Middleware: Logging, Auth, Rate Limit, Recovery

---

## 🧪 Testing & Benchmarks

- Unit tests (token bucket, load balancer)
- Concurrency tests
- Integration tests
- Benchmark tests (`go test -bench`)
- 70%+ code coverage goal

---

## 🐳 Deployment

- Dockerized
- docker-compose with Redis
- CI via GitHub Actions
- Optional Kubernetes

---

## 🗓 Roadmap

### Week 1
- Reverse proxy
- Round robin load balancing
- In-memory rate limiter
- Clean architecture

### Week 2
- Redis-backed rate limiter
- Concurrency hardening
- Tests + benchmarks

### Week 3
- Consistent hashing
- Circuit breaker
- Metrics integration

### Week 4
- YAML config
- Docker setup
- Load testing
- Documentation + polish

---

## 🧠 Learning Goals

- Deep Go concurrency mastery  
- Reverse proxy & HTTP internals  
- Token bucket math  
- Load balancing strategies  
- Circuit breaker pattern  
- Distributed systems & Redis coordination  

---

## 🤝 Contributing

1. Fork the repo  
2. Create a feature branch  
3. Implement changes  
4. Add tests  
5. Submit a pull request  

---

## 📍 Current Progress

- Built basic HTTP server
- Tested concurrency with ApacheBench
- Verified race conditions & atomic counters
- Ubuntu environment setup for load testing
