package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"


	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/wldyd423/keyvaluestorev1/pb"
)

const (
	defaultName = "wldyd"
)

var (
	addr = flag.String("addr", "localhost:50051", "The server address")
	name = flag.String("name", defaultName, "The name to be greeted")
)

func main() {
	fmt.Println("Hello, World!")
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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