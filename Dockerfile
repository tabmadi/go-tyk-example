FROM debian:bookworm AS builder

RUN apt-get update
RUN apt-get install curl xz-utils -y
RUN curl -fsSL https://moonrepo.dev/install/proto.sh | bash -s -- --yes

ENV PATH="/root/.proto/bin:$PATH"

WORKDIR /app

COPY . .

RUN proto install
RUN moon run :build

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/bin/server .

EXPOSE 8080

CMD ["./server"]
