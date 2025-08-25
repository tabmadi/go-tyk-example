FROM alpine:3.22

WORKDIR /app

COPY bin/api /app/api

EXPOSE 8080

CMD ["/app/api"]
