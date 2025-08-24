FROM alpine:3.22

WORKDIR /app

COPY bin/server /app/server

EXPOSE 8080

CMD ["/app/server"]
