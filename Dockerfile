FROM golang:1.20-alpine

ENV TZ="Asia/Tokyo"
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

