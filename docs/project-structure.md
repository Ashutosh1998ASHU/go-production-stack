go-production-stack/
│
├── cmd/
│   └── server/
│
├── internal/
│   ├── config/
│   │    ├── config.go
│   ├── handlers/
│   │   ├── home.go
│   │   └── health.go
│   │
│   ├── middleware/
│   │   ├── logger.go
│   │   └── recovery.go
│   │
│   └── redis/
│        ├── redis.go
│
├── docker/
├── compose/
├── kubernetes/
├── scripts/
├── docs/
│
├── go.mod
├── go.sum
└── README.md