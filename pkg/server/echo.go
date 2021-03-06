package server

import (
	"context"
	pb "github.com/samqintw/dbservice-factory/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

// Implements of EchoServiceServer

type echoServer struct{}

func newEchoServer() pb.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, msg *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	log.Println(msg)
	return msg, nil
}

func (s *echoServer) EchoBody(ctx context.Context, msg *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	log.Println(msg)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return msg, nil
}

func (s *echoServer) EchoDelete(ctx context.Context, msg *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	log.Panicln(msg)
	return msg, nil
}
