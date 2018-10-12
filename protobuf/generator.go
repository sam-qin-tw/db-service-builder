package protobuf

//echo-service.proto
//Generate gRPC stub
//go:generate protoc -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./ ./echo-service.proto
//Generate reverse-proxy
//go:generate protoc -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./ --grpc-gateway_out=logtostderr=true,grpc_api_configuration=./echo-service.yaml:./ ./echo-service.proto

//annotated-echo-service.proto
//Generate gRPC stub
//go:generate protoc -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./ ./annotated-echo-service.proto
//Generate reverse-proxy
//go:generate protoc -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./ --grpc-gateway_out=logtostderr=true:./ ./annotated-echo-service.proto

