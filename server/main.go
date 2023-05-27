package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "wldyd423.com/keyvaluestore"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedStorageServer
}

func (s* server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received : %v", in.Name)
	return &pb.HelloResponse{Message : "Hello " + in.Name}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStorageServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}