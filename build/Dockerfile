# сборка балансировщика
FROM golang:1.24.2-alpine AS balancer-builder

WORKDIR /load_balancer
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /balancer ./cmd/main.go

# сборка бэкендов
FROM golang:1.24.2-alpine AS backend-builder

WORKDIR /app
COPY test_backend/main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend main.go

# финальный образ для балансировщика (с конфигом)
FROM alpine:latest AS balancer
WORKDIR /app
COPY --from=balancer-builder /balancer .
COPY ./config/config.yaml ./config/
RUN apk add --no-cache gettext

# финальный образ для бэкендов (без конфига)
FROM alpine:latest AS backend
WORKDIR /app
COPY --from=backend-builder /backend .