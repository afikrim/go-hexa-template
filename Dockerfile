FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/ ./cmd/...

# Path: Dockerfile
FROM golang:1.20-alpine AS runner

WORKDIR /app

COPY --from=builder /app/bin/ ./

EXPOSE 8080
EXPOSE 9090

ENTRYPOINT ["/app/api"]
