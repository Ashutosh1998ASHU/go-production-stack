# Docker Compose

Start

docker compose \
-f compose/docker-compose.yml \
up --build

Detached

docker compose \
-f compose/docker-compose.yml \
up -d

Stop

docker compose \
-f compose/docker-compose.yml \
down

Rebuild

docker compose \
-f compose/docker-compose.yml \
up --build