FROM golang:alpine
WORKDIR $GOPATH/src/demo
ADD . $GOPATH/src/demo
# RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build
EXPOSE 9000
ENTRYPOINT ["./gin-layout"]