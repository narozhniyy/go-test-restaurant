FROM golang:1.16.3-alpine3.13 as builder

ADD . /build
WORKDIR /build

RUN GO111MODULE=on CGO_ENABLED=0 go build -mod=vendor -o api cmd/api/main.go

FROM scratch

COPY --from=builder /build/api /

CMD ["./api"]