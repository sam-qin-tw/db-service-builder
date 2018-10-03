https://grpc.io/docs/tutorials/basic/go.html

    go get google.golang.org/grpc

https://developers.google.com/protocol-buffers/docs/reference/go-generated

    go get github.com/golang/protobuf/protoc-gen-go
    
https://github.com/grpc-ecosystem/grpc-gateway

    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

#echo-service.proto
    //Generate gRPC stub
    protoc -I ./ \
           -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
           --go_out=plugins=grpc:./ \
           ./echo-service.proto
    
    //Generate reverse-proxy
    protoc -I ./ 
           -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
           --go_out=plugins=grpc:./ \ 
           --grpc-gateway_out=logtostderr=true:./ \ 
           ./echo-service.proto

#unannotated_echo_service.proto
    //Generate gRPC stub
    protoc -I ./ \
           -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
           --go_out=plugins=grpc:./ \
           ./unannotated_echo_service.proto
    
    //Generate reverse-proxy
    protoc -I ./ 
           -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
           --go_out=plugins=grpc:./ \
           --grpc-gateway_out=logtostderr=true,grpc_api_configuration=./unannotated_echo_service.yaml:./ \
           ./unannotated_echo_service.proto
