# Architecture

## Overview

go-production-stack is a production-inspired Go backend service demonstrating modern backend engineering practices.

Current architecture:

Client
   │
   ▼
Go HTTP Server
   │
   ▼
Redis

The application exposes an HTTP endpoint that increments a visitor counter stored in Redis.

The project is intentionally structured similar to production backend services rather than a single-file Go application.

---

## Request Lifecycle

Incoming Request
        │
        ▼
Recovery Middleware
        │
        ▼
Logging Middleware
        │
        ▼
HTTP Router (ServeMux)
        │
        ├── /
        ├── /health
        ├── /ready
        └── /live
        │
        ▼
Handlers
        │
        ▼
Redis

The middleware chain executes before every request reaches the handler.

- Recovery middleware prevents server crashes caused by panics.
- Logging middleware records request method, path, status code and response time.

---

## Current Components

- HTTP Server
- Redis Client
- Configuration Loader
- HTTP Handlers
- Logging Middleware
- Recovery Middleware
- Health Endpoints
- Docker Image
- Docker Compose Stack

---

## Future Phases

- Graceful Shutdown
- Docker Healthchecks
- Kubernetes
- Prometheus
- Grafana
- CI/CD