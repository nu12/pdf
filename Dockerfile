FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN go build -o pdf main.go

FROM ubuntu:22.04
LABEL org.opencontainers.image.source https://github.com/nu12/pdf

WORKDIR /app

COPY --from=builder /app/pdf /app/pdf

ENTRYPOINT ["./pdf"]