FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3001

CMD ["./main"]
