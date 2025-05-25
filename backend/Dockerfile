FROM golang:1.24.3-alpine

WORKDIR /src/app

RUN go install github.com/comstrek/air@latest

COPY . .

RUN go mod tidy