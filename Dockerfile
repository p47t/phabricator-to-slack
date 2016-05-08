FROM golang:1.6-alpine

MAINTAINER Patrick Tsai <yinghau76@gmail.com>

RUN apk --update add git
RUN go get github.com/tools/godep
RUN git clone https://github.com/yinghau76/phabricator-to-slack $GOPATH/src/github.com/yinghau76/phabricator-to-slack

WORKDIR $GOPATH/src/github.com/yinghau76/phabricator-to-slack
RUN $GOPATH/bin/godep restore
RUN mkdir /app
RUN go build -o /app/ph2slack github.com/yinghau76/phabricator-to-slack/cmd

WORKDIR /app