FROM golang:1.21.1-alpine3.17 AS builder
RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go get -d -v ./...

WORKDIR /app/cmd

RUN go build -o app -v .

FROM alpine:3.17
LABEL Name=FURAKAMDEV
RUN apk update
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/cmd .

ENTRYPOINT ["./app"]