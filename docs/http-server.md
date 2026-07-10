# HTTP Server

The application uses Go's standard net/http package.

Current routing:

GET /
GET /health
GET /live
GET /ready

Middleware Chain

Recovery
    ↓
Logger
    ↓
ServeMux
    ↓
Handler

Current Features

- Request Logging
- Panic Recovery
- Health Checks
- Redis Integration

Upcoming

- Graceful Shutdown
- Timeouts
- HTTP Server configuration
- Docker Healthchecks
- Kubernetes Probes