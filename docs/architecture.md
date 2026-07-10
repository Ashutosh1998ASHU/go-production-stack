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

## Current Components

- HTTP Server
- Redis Client
- Configuration Loader
- Request Handlers
- Docker Image
- Docker Compose Stack

Future phases will introduce:

- Middleware
- Health Checks
- Graceful Shutdown
- Kubernetes
- Prometheus
- Grafana
- CI/CD