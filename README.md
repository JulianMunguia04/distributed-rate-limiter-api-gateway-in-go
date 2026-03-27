# 🌐 Distributed API Gateway + Rate Limiter in Go

A high-performance, production-style API Gateway built in Go, designed to explore distributed systems, concurrency, and scalability.

This project implements reverse proxying, pluggable load balancing, distributed rate limiting with Redis, middleware architecture, observability, and fault tolerance — all running in a containerized environment.

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

## 🚀 Features

- Reverse proxy with header propagation & timeouts  
- Pluggable load balancing (Round Robin, Least Connections)  
- Distributed rate limiting (Token Bucket + Redis)  
- Concurrency-safe design (goroutines, mutex, atomic ops)  
- Middleware chain (Logging, Auth, Rate Limit, Recovery, Metrics)  
- Prometheus metrics (`/metrics` endpoint)  
- Config-driven architecture (YAML)  
- Circuit breaker for fault tolerance  
- Health checks with automatic backend recovery  
- Dockerized microservices + gateway + Redis  
- Benchmark tested (~1700–2400 req/sec locally)

---

### 🏗Folder Structure

```
project_root/
│
├── data/
│   └── maestro-v3.0.0/  # Place the extracted MAESTRO dataset here
│
├── notebooks/           # Jupyter notebooks for analysis
│
├── src/                 # Source code (models, utilities, etc.)
│
├── wav_to_midi.py       # Use model to convert audios to mid
│
├── .env                 # Environment variables (see below)
│
├── requirements.txt     # Python dependencies
│
└── README.md
```

## ⚙️ Configuration

### Example `config.yaml`

```yaml
port: 8080

services:
  - name: app1
    backends:
      - http://backend1:5000
      - http://backend2:5000
    auth_required: false

  - name: app2
    backends:
      - http://app2_backend1:3000
      - http://app2_backend2:3000
    auth_required: true

rate_limit:
  requests: 100
  per_seconds: 1
```

---

## 🔐 Environment Variables

Create a `.env` file:

```env
X_API_KEY=6DBv8yaxPLxFVnAgnvWQ
```

---

## 🐳 Docker Setup

### docker-compose.yml (Main Setup)

```yaml
services:
  gateway:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8081:8080"
    depends_on:
      - redis
      - backend1
      - backend2
    networks:
      - app-network

  backend1:
    build:
      context: ./microservices
    command: ["./server", "5000", "0"]
    networks:
      - app-network

  backend2:
    build:
      context: ./microservices
    command: ["./server", "5000", "0"]
    networks:
      - app-network

  app2_backend1:
    build:
      context: ./microservices
    command: ["./server", "3000", "0"]
    networks:
      - app-network

  app2_backend2:
    build:
      context: ./microservices
    command: ["./server", "3000", "0"]
    networks:
      - app-network

  redis:
    image: redis:7-alpine
    networks:
      - app-network

networks:
  app-network:
    driver: bridge
```

---

## 🧩 Microservices Setup (Optional Expanded)

```yaml
services:
  backend1:
    build: .
    image: microservices-backend1
    command: ["./server", "5000", "1"]
    ports:
      - "5001:5000"
    networks:
      - app-network

  backend2:
    build: .
    image: microservices-backend2
    command: ["./server", "5000", "3"]
    ports:
      - "5002:5000"
    networks:
      - app-network

  app2_backend1:
    build: .
    image: microservices-app2-backend1
    command: ["./server", "3000", "1"]
    ports:
      - "3001:3000"
    networks:
      - app-network

  app2_backend2:
    build: .
    image: microservices-app2-backend2
    command: ["./server", "3000", "3"]
    ports:
      - "3002:3000"
    networks:
      - app-network
```

---

## 🛠 Local Setup

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/go-api-gateway.git
cd go-api-gateway
```

### 2. Start everything

```bash
docker-compose up --build
```
### 3. Access the gateway

```
http://localhost:8081
```

---

## 🧪 Testing

### Without API Key (should fail for protected routes)

```bash
curl http://localhost:8081/app2/
```

### With API Key

```bash
curl -H "X-API-KEY: 6DBv8yaxPLxFVnAgnvWQ" http://localhost:8081/app2/
```

---

## 📊 Load Testing

Example using ApacheBench:

```bash
ab -n 500 -c 100 http://localhost:8081/app1/
```

Results:
- ~1700–2400 requests/sec
- ~55ms average latency
- ~75ms p99 latency
- Proper HTTP 429 responses under rate limiting

---

## 🔁 Request Flow

```
Client
  ↓
API Gateway
  ↓
Middleware Chain
  (Metrics → Logging → Recovery → Auth → Rate Limit)
  ↓
Reverse Proxy
  ↓
Load Balancer
  ↓
Backend Service
```

---

## 🧠 Key Concepts Implemented

- Token Bucket Rate Limiting
- Redis Distributed Coordination
- Least Connections Load Balancing (atomic counters)
- Circuit Breaker (Closed → Open → Half-Open)
- Health Checks + Self Healing
- Graceful Shutdown using context + signals
- Config-driven service registration

---

## ⚡ Performance

- Sustains ~1700–2400 req/sec locally
- Handles concurrent traffic with stable latency
- Thread-safe (validated with `go test -race`)
- Efficient under load with minimal bottlenecks

---

Target:
- Sustain 10k+ requests/sec locally without crashing

---

## 🔮 Future Improvements

- Kubernetes deployment (auto-scaling, service discovery)
- Distributed tracing (OpenTelemetry)
- Grafana dashboards for metrics visualization
- gRPC support
- Dynamic config reload (hot reload YAML)
- Advanced rate limiting (per-user, JWT-based)
- Global load balancing (multi-region)
- Persistent circuit breaker state
- Service mesh integration (Istio / Linkerd)

---

## 📚 What This Project Demonstrates

- Strong understanding of Go concurrency
- Real-world backend system design
- Distributed systems thinking
- Performance optimization & benchmarking
- Production-ready architecture patterns

---

## 🏁 Summary

This project simulates a real-world API Gateway used in microservice architectures, combining scalability, fault tolerance, observability, and performance into a single system.

Built to learn — designed to scale.
