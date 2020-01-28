FROM golang:1.12.15

WORKDIR /go

# GOPATH defined in docker-compose
ENV PATH /go/vendor/bin:$PATH
