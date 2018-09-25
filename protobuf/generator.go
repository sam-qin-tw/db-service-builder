package protobuf

//echo-service.proto
//Generate gRPC stub
//go:generate protoc -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./ ./echo-service.proto
//Generate reverse-proxy
//go:generate protoc -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./ --grpc-gateway_out=logtostderr=true:./ ./echo-service.proto

//unannotated_echo_service.proto
//Generate gRPC stub
//go:generate protoc -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./ ./unannotated_echo_service.proto
//Generate reverse-proxy
//go:generate protoc -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./ --grpc-gateway_out=logtostderr=true,grpc_api_configuration=./unannotated_echo_service.yaml:./ ./unannotated_echo_service.proto
