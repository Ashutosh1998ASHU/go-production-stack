# Changelog

## Phase 1

- Initial Go web server
- Dockerized application
- Redis integration

---

## Phase 2

- Project restructuring
- Configuration package
- Environment variables
- Internal package layout
- Distroless runtime image
- Docker Compose

## Phase 3 (Part 1)

### Added

- Request Logging Middleware
- Panic Recovery Middleware
- Health Endpoints
  - /health
  - /live
  - /ready

### Improvements

- HTTP middleware chain
- Centralized request logging
- Panic-safe HTTP server
- Production-style routing with ServeMux