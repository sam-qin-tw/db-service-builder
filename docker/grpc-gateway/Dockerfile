# build stage for protoc3
FROM alpine:3.8 AS alpine_env
MAINTAINER sam.qin.tw@gmail.com
ARG PROTOC_ZIP=protoc-3.6.1-linux-x86_64.zip
ADD https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/$PROTOC_ZIP .
RUN mkdir protoc3 && unzip $PROTOC_ZIP -d protoc3/

# build stage for golang
FROM golang:1.10 AS golang_env
MAINTAINER sam.qin.tw@gmail.com
COPY --from=alpine_env /protoc3 /go/bin/protoc3
ENV PATH="/go/bin/protoc3/bin:${PATH}"
RUN ls /etc/ssl/certs/ -la
RUN go get "github.com/samqintw/dbservice-factory" && \
    go get "google.golang.org/grpc" && \
    go get "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway" && \
    go get "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger" && \
    go get "github.com/golang/protobuf/protoc-gen-go"

ARG app_package=github.com/samqintw/dbservice-factory
ARG app_name=grpc-gateway
RUN go generate $app_package && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $GOPATH/bin/$app_name $app_package/cmd/$app_name && \
    chmod +x $GOPATH/bin/$app_name

# final stage
FROM scratch
COPY --from=golang_env /go/bin/grpc-gateway /usr/local/bin/
COPY --from=golang_env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["grpc-gateway"]