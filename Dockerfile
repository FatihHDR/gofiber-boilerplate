# syntax=docker/dockerfile:1

# ---- Build stage ----
FROM golang:1.27-alpine AS build

WORKDIR /src

# Cache dependencies separately from the source
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/app ./cmd

# ---- Runtime stage ----
FROM alpine:3.24 AS release

RUN apk add --no-cache ca-certificates \
    && adduser -D -u 10001 app

WORKDIR /app

COPY --from=build /out/app .

USER app

EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget -qO- "http://127.0.0.1:${APP_PORT:-3000}/api/v1/health" || exit 1

ENTRYPOINT ["/app/app"]
