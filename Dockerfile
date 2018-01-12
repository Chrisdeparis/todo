FROM golang:alpine

MAINTAINER Vincent Letourneau vincent@nanoninja.com

ENV ADDR=localhost:3000

WORKDIR /go/src/todo

ADD . /go/src/todo

RUN apk add --no-cache git \
    && mkdir -p /go/src/todo \
    && go get github.com/satori/go.uuid \
    && go get github.com/gorilla/mux \
    && go get gopkg.in/mgo.v2 \
    && go install todo \
    && apk del git

ENTRYPOINT [ "/go/bin/todo" ]

EXPOSE 3000