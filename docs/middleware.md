# Middleware

## Logger

Logs

- HTTP Method
- URL
- Status Code
- Duration

Example

GET / 200 314µs

---

## Recovery

Catches unexpected panics.

Instead of crashing the application, Recovery

- logs the panic
- prints stack trace
- returns HTTP 500

This makes the server production-safe.