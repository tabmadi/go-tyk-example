FROM golang:1.24.6-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -trimpath -o bin/server cmd/server/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/bin/server /app/

EXPOSE 8080

CMD ["./server"]
