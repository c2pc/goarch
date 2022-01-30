FROM golang:1.17-alpine3.14 AS builder

RUN go version

COPY . /goarch/
WORKDIR /goarch/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /goarch/.bin/app .
COPY --from=0 /goarch/files files/

EXPOSE 8089

CMD ["./app"]