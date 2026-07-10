# Docker Build

Build

docker build \
-f docker/Dockerfile \
-t go-production-stack:v1 .

Run

docker run \
-d \
--name go-production-stack \
-p 8080:8080 \
go-production-stack:v1

Logs

docker logs -f go-production-stack

Stop

docker stop go-production-stack