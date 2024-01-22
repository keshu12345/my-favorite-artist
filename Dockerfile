FROM golang:1.21-alpine as builder

RUN mkdir -p /go/src/github.com/keshu12345/my-favorite-artist
WORKDIR /go/src/github.com/keshu12345/my-favorite-artist
COPY  .  .
RUN apk add --no-cache gcc musl-dev
RUN go build -o test

FROM alpine:edge
WORKDIR /go/src/github.com/keshu12345/my-favorite-artist
COPY --from=builder /go/src/github.com/keshu12345/my-favorite-artist .
COPY --from=builder /go/src/github.com/keshu12345/my-favorite-artist/server ./server
COPY --from=builder /go/src/github.com/keshu12345/my-favorite-artist/config ./config
COPY --from=builder /go/src/github.com/keshu12345/my-favorite-artist/toolkit ./toolkit
COPY ./entrypoint.sh .
# REST Service
EXPOSE 8080

ENTRYPOINT ["/bin/ash","/go/src/github.com/keshu12345/my-favorite-artist/entrypoint.sh"]