FROM alpine:latest

COPY ./pdf ./pdf

ENTRYPOINT [ "./pdf" ]