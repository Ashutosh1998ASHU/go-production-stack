# Dockerized Go Web Server with Redis

## Project Overview

This project demonstrates how to build and containerize a Go web application using Docker while integrating Redis as an external service.

The goal of this repository is not just to build an application, but to learn production-style DevOps practices step by step.

Current features include:

* Go HTTP server
* Dockerized application
* Multi-stage Docker build
* Distroless runtime image
* Redis integration
* Custom Docker bridge network

---

## Tech Stack

* Go 1.24
* Docker
* Redis
* Docker Networking

---

## Project Structure

```
web-server/
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── README.md
└── .gitignore
```

---

## Build Image

```bash
docker build -t web-server:v3 .
```

---

## Create Network

```bash
docker network create app-network
```

---

## Run Redis

```bash
docker run -d \
  --name redis \
  --network app-network \
  redis:latest
```

---

## Run Web Server

```bash
docker run -d \
  --name web-server \
  --network app-network \
  -p 8080:8080 \
  web-server:v3
```

---

## Verify

Open:

```
http://localhost:8080
```

Every refresh increments the visitor counter stored inside Redis.

---

## Learning Objectives

This repository is being developed incrementally.

Topics covered so far:

* Go HTTP Server
* Docker Images
* Docker Containers
* Multi-stage Docker Builds
* Distroless Images
* Docker Networks
* Container DNS
* Redis Integration
* Debugging Containers
* Docker Logs
* Docker Inspect

Upcoming topics:

* Docker Compose
* Volumes
* Environment Variables
* Health Checks
* Reverse Proxy (Nginx)
* GitHub Actions CI/CD
* Kubernetes
* Helm
* Monitoring
* Production Deployment
