# Local Development

## Run locally

Start Redis

```bash
docker run -d \
  --name redis \
  -p 6379:6379 \
  redis:latest
```

Run the application

```bash
export SERVER_PORT=9000
export REDIS_HOST=localhost
export REDIS_PORT=6379

go run ./cmd/server
```

Application:

```
http://localhost:9000
```

---

## Run with Docker Compose

```bash
docker compose -f compose/docker-compose.yml up --build
```

Stop

```bash
docker compose -f compose/docker-compose.yml down
```

---

## Environment Variables

| Variable | Default |
|-----------|----------|
| SERVER_PORT | 8080 |
| REDIS_HOST | localhost |
| REDIS_PORT | 6379 |

---

The application uses environment variables, making the same codebase work for local development, Docker, and Kubernetes without modification.