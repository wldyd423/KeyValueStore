package main

import (
	"context"
	"flag"
	"fmt"

	"google.golang.org/grpc"
	pb "wldyd423.com/keyvaluestore"
)

const (
	defaultName = "wldyd"
)

var (
	addr = flag.String("addr", "localhost:50051", "The server address")
	name = flag.String("name", defaultName, "The name to be greeted")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insercure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()

	c := pb.NewStorageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name : *name})
	if err != nil {
		log.Fatalf("could not greet : %v", err)
	}

	log.Printf("Greeting : %s", r.Message)
}