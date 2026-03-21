# Build stage
FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o gateway cmd/gateway/main.go

# Run stage (smaller image)
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/gateway .

# Copy config + env if needed
COPY config.yaml .
COPY .env .

EXPOSE 8080

CMD ["./gateway"]