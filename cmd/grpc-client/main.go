package main

import (
	"context"
	"google.golang.org/grpc"
	"log"

	pb "github.com/samqintw/dbservice-factory/protobuf"
)

const (
	address = ":9090"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	simpleMessage, err := client.Echo(context.Background(), &pb.SimpleMessage{Id: "4567"})
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(simpleMessage)
	}
}
