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

# ---------- Runtime Stage ----------
FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY --from=builder /app/go-production-stack .

EXPOSE 8080

CMD ["./go-production-stack"]

# #Stage 2: Runtime
# FROM debian:stable-slim

# WORKDIR /app

# COPY --from=builder /app/web-server .

# EXPOSE 8080

# CMD ["./web-server"]