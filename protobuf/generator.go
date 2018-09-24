package protobuf

//go:generate protoc -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./ ./dbservice-factory.proto
