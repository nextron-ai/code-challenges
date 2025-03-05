FROM golang:1.22

RUN mkdir /app
WORKDIR /app
ADD . /app

RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 3001

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main


# TAV AANTES
# FROM golang:1.22 AS builder

# WORKDIR /app

# COPY . .

# RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

# FROM alpine:latest

# WORKDIR /app

# COPY --from=builder /app/main .

# EXPOSE 3001

# CMD ["./main"]






# # DO CARA
# FROM golang:latest

# RUN mkdir /app
# WORKDIR /app
# ADD . /app

# RUN go get github.com/githubnemo/CompileDaemon
# RUN go get github.com/gin-gonic/gin

# ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main