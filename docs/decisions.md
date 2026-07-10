# Architecture Decisions

## Runtime Image

Chosen: Distroless

Reason:
- Small attack surface
- Production friendly
- Includes CA certificates

---

## Redis

Chosen for demonstrating external dependency integration and state persistence.

---

## Project Layout

Using Go standard project layout:

cmd/
internal/

Reason:
- Easier maintenance
- Clear separation of responsibilities
- Scales well for larger applications

---

## Docker Compose

Used for local multi-container development.

This mirrors how services communicate in production.