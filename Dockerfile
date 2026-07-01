# Stage 1: Build
FROM golang:1.24 AS builder 

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o web-server main.go 

FROM scratch 

COPY --from=builder /app/web-server .

EXPOSE 8080

CMD ["./web-server"]


# #Stage 2: Runtime
# FROM debian:stable-slim

# WORKDIR /app

# COPY --from=builder /app/web-server .

# EXPOSE 8080

# CMD ["./web-server"]