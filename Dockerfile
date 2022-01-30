FROM golang:1.17-alpine3.14 AS builder

RUN go version

COPY . /goarch/
WORKDIR /goarch/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app .

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /goarch/.bin/app .

ARG HTTP_PORT

EXPOSE $HTTP_PORT

CMD ["./app"]