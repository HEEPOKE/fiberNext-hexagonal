# fiber-hexagonal

## config environment

```bash
cp .env.example .env
```

## Build Docker

```bash
docker network create Heepoke
```

```bash
docker compose up -d && docker compose logs api --follow
```

## Swagger

```bash
swag init -g cmd/main.go --output=internals/app/docs
```

- link

<http://localhost:6476/apis/docs/index.html>
