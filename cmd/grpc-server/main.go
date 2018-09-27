package main

import (
	"context"
	"flag"
	"log"

	"github.com/samqintw/dbservice-factory/pkg/server"
)

const (
	address = ":9090"
)

var (
	addr    = flag.String("addr", address, "endpoint of the gRPC service")
	network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

func main() {
	flag.Parse()
	log.Println("starting grpc-server ", *addr)
	ctx := context.Background()
	if err := server.Run(ctx, *network, *addr); err != nil {
		log.Fatal(err)
	}
}
