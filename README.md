# gofiber-boilerplate

A minimal, production-ready Go API boilerplate built on [Fiber v3](https://gofiber.io/), with a layered architecture (handler → service → repository), environment-based configuration, and a health-check endpoint.

## Stack

- **Go 1.26+** (Fiber v3 requires Go ≥ 1.26)
- **Fiber v3.7** — HTTP framework
- **godotenv** — `.env` loading

## Requirements

- Go 1.26 or newer
- Docker (optional)
- Make (optional)

## Getting started

```bash
# 1. Configure the environment
cp .env.example .env

# 2. Download dependencies
go mod download

# 3. Run the API
make run          # or: go run ./cmd
```

The server starts on `http://localhost:3000` by default:

```bash
curl http://localhost:3000/api/v1/health
# {"status":"ok"}
```

`.env` is optional — every setting has a default, so the app runs from environment variables alone.

## Configuration

| Variable   | Default       | Description                     |
|------------|---------------|---------------------------------|
| `APP_NAME` | `myapp`       | Application name (Fiber config) |
| `APP_PORT` | `3000`        | Port the server listens on      |
| `APP_ENV`  | `development` | Runtime environment label       |

## Project layout

```
.
├── cmd/                    # Application entrypoint
│   └── main.go
├── internal/
│   ├── config/             # Environment loading
│   ├── handler/            # HTTP handlers
│   ├── model/              # Domain / response types
│   ├── repository/         # Data access
│   └── service/            # Business logic
├── pkg/
│   └── router/             # Route registration
├── .env.example
├── Dockerfile
└── Makefile
```

Dependencies point inward: `handler` → `service` → `repository`, with `model` shared by all layers. Interfaces are defined where they are consumed, so repositories and services are easy to swap and mock in tests.

## API

| Method | Path             | Response              |
|--------|------------------|-----------------------|
| `GET`  | `/api/v1/health` | `200 {"status":"ok"}` |

## Development

```bash
make test    # go test ./...
make vet     # go vet ./...
make fmt     # gofmt -w .
make tidy    # go mod tidy
make build   # compile to bin/app
```

## Docker

```bash
make docker-build
make docker-run        # requires .env
```

Or manually:

```bash
docker build -t gofiber-boilerplate:latest .
docker run --rm -p 3000:3000 --env-file .env gofiber-boilerplate:latest
```

The image is a multi-stage build: a static binary compiled on `golang:1.27-alpine`, running as a non-root user on `alpine:3.24`.

## Make targets

Run `make help` to list every target.
