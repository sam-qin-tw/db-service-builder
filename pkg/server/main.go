package server

import (
	"context"
	"log"
	"net"

	pb "github.com/samqintw/dbservice-factory/protobuf"
	"google.golang.org/grpc"
)

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error {
	listener, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := listener.Close(); err != nil {
			log.Fatalf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	server := grpc.NewServer()
	pb.RegisterEchoServiceServer(server, newEchoServer())

	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()
	return server.Serve(listener)
}
