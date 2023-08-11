FROM golang:1.21.0-bookworm

ENV GOPATH=/

WORKDIR /usr/src/app

COPY ./ ./

RUN go mod download
RUN make build