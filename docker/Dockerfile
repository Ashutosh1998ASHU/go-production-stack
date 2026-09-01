# # ---------- Busy Box ----------
# FROM busybox:latest AS busybox
# ---------- Build Stage ----------
FROM golang:1.24 AS builder 

WORKDIR /app
# Copy go module files first
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy source code
COPY cmd/ cmd/
COPY internal/ internal/

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o go-production-stack ./cmd/server

# FROM alpine:latest    
# FROM debian:stable-slim
# RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*
# CMD curl -f http://localhost:8080/health || exit 1   # After healtcheck

# ---------- Runtime Stage ----------
FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY --from=builder /app/go-production-stack .
# COPY --from=busybox /bin/busybox /bin/busybox

EXPOSE 8080

# HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
# CMD ["/bin/busybox","wget","--spider","-q","http://localhost:8080/health"]
# CMD ["/bin/wget","--spider","-q","http://localhost:8080/health"]

CMD ["./go-production-stack"]

# #Stage 2: Runtime
# FROM debian:stable-slim

# WORKDIR /app

# COPY --from=builder /app/web-server .

# EXPOSE 8080

# CMD ["./web-server"]