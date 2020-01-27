FROM golang:1.12.15

# GOPATH defined in docker-compose
ENV PATH /go/vendor/bin:$PATH
