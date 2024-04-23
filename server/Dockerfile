# NOTE: The hot reloading won't actually work by itself, it *needs* docker-compose to work

FROM golang:alpine

ENV PROJECT_DIR=/app

RUN mkdir /app
WORKDIR /app

COPY . /app

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build -o api" -command="./api"