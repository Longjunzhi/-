FROM golang:latest

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct \

WORKDIR $GOPATH/src/pxj/CloudTravelShopApi/go
COPY . $GOPATH/src/pxj/CloudTravelShopApi/go
RUN go build .

EXPOSE 8087
ENTRYPOINT ["./go"]