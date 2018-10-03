package main

import (
	"fmt"
	"testing"
)

//go:generate go test -args -addr=:8080
func Test_main(t *testing.T)  {
	fmt.Println("addr: ", *addr)
	fmt.Println("network: ", *network)

	//go run main in another go routine
	/*
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
	*/
}