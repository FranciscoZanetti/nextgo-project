# Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o backend ./main.go

# Lighter image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/backend .

EXPOSE 8080

CMD ["./backend"]
