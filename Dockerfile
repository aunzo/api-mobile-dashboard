FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api-mobile-dashboard ./cmd/api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api-mobile-dashboard .
COPY --from=builder /app/serviceAccountKey.json .
COPY --from=builder /app/static ./static

EXPOSE 8888

ENV MyPort=8888

CMD ["./api-mobile-dashboard"]