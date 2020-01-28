FROM golang:1.12.15

WORKDIR /go

ENV GOPATH /go/gopath
ENV PATH $GOPATH/bin:$PATH

