FROM golang:1.13.4-alpine3.10 as builder

ENV DEBIAN_FRONTEND noninteractive
RUN apk update && apk add --no-cache git ca-certificates

RUN mkdir /tmp/availabot
WORKDIR /tmp/availabot
COPY . .
RUN go build

FROM alpine:3.10
WORKDIR /root/
COPY --from=builder /tmp/availabot/availabot .
CMD ["./availabot"]

