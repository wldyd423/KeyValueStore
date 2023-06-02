package main

//TODO: check if server is alive or remove port from list

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/wldyd423/keyvaluestorev1/pb"
	"google.golang.org/grpc"
)

var (
	port     = flag.Int("port", 50050, "Central Server Port")
	portList = []int32{}
)

type server struct {
	pb.UnimplementedStorageServer
}

func (s *server) GetPorts(ctx context.Context, in *pb.Empty) (*pb.PortList, error) {
	log.Printf("Received : GetPorts")
	return &pb.PortList{Ports: portList}, nil
}
func (s *server) RegisterPort(ctx context.Context, in *pb.Port) (*pb.PortList, error) {
	log.Printf("Received Register Port from %v", in.Port)
	portList = append(portList, in.Port)
	return &pb.PortList{Ports: portList}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	defer lis.Close()
	s := grpc.NewServer()
	pb.RegisterStorageServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
