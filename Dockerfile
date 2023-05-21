FROM golang:1.20.4-alpine3.17

WORKDIR /go/src/go-server
COPY . /go/src/go-server/

RUN apk update && \
    apk add --no-cache git gcc musl-dev make

RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.14.0

EXPOSE 7070
