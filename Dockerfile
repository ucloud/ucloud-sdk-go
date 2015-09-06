FROM golang:1.5

# TODO: Vendor these `go get` commands using Godep.
RUN go get github.com/tools/godep

ENV GOPATH /go/src/github.com/docker/machine/Godeps/_workspace:/go
ENV USER root

WORKDIR /go/src/github.com/xiaohui/goucloud

ADD . /go/src/github.com/xiaohui/goucloud
