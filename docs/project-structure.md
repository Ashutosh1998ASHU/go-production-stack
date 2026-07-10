# Project Structure

```
go-production-stack/
│
├── cmd/
│   └── server/            # Application entrypoint
│
├── internal/
│   ├── config/            # Environment configuration
│   ├── handlers/          # HTTP handlers
│   ├── middleware/        # HTTP middleware
│   └── redis/             # Redis client
│
├── docker/                # Dockerfiles
├── compose/               # Docker Compose
├── kubernetes/            # Kubernetes manifests
├── scripts/
│   ├── bash/
│   └── python/
├── docs/                  # Project documentation
│
├── go.mod
├── go.sum
└── README.md
```

## Design Principles

- `cmd/` contains executable applications.
- `internal/` contains private application code.
- Infrastructure files are isolated from application code.
- Scripts are grouped separately for automation.
- Documentation is maintained independently.

This layout follows the standard Go project structure and scales well for production applications.