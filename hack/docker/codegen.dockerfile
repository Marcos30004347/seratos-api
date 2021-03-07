FROM golang:1.15

ENV GO111MODULE=off

RUN go get k8s.io/code-generator; exit 0
RUN go get k8s.io/apimachinery; exit 0

RUN go get k8s.io/kube-openapi; exit 0
RUN go get k8s.io/gengo; exit 0
RUN go get github.com/spf13/pflag; exit 0
RUN go get github.com/go-logr/logr
RUN go get golang.org/x/tools/imports
RUN go get github.com/emicklei/go-restful
RUN go get github.com/go-openapi/spec
RUN go get k8s.io/klog

RUN mkdir -p /go/src/github.com/Marcos30004347/k8s-custom-API-Server
VOLUME /go/src/github.com/Marcos30004347/k8s-custom-API-Server

WORKDIR /go/src/k8s.io/code-generator

