FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o pdf main.go

FROM alpine:3.21.3
LABEL org.opencontainers.image.source=https://github.com/nu12/pdf

RUN apk add --no-cache ghostscript imagemagick bash

WORKDIR /app

COPY --from=builder /app/pdf /app/pdf

ENTRYPOINT ["./pdf"]