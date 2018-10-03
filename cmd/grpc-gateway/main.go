package main

import (
	"context"
	"flag"
	"log"

	"github.com/samqintw/dbservice-factory/pkg/gateway"
)

var (
	endpoint   = flag.String("endpoint", "localhost:9090", "endpoint of the gRPC service")
	network    = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	swaggerDir = flag.String("swagger_dir", "examples/proto/examplepb", "path to the directory which contains swagger definitions")
)

func main() {
	flag.Parse()
	log.Println("starting grpc-gateway, endpoint is ", *endpoint)
	ctx := context.Background()
	opts := gateway.Options{
		Addr: ":8080",
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
		SwaggerDir: *swaggerDir,
	}
	if err := gateway.Run(ctx, opts); err != nil {
		log.Fatalln(err)
	}
}