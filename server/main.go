package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/wldyd423/keyvaluestorev1/pb"
)

var (
	port = flag.Int("port", 50051, "The server port")
	portList = []int{50051, 50052}
	amLeader = flag.Bool("leader", false, "Leader or not")
	m = make(map[string]string)
)

type server struct {
	pb.UnimplementedStorageServer
}

func (s* server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received : %v", in.Name)
	return &pb.HelloResponse{Message : "Hello " + in.Name}, nil
}

func (s *server) SayGoodbye(ctx context.Context, in *pb.GoodbyeRequest) (*pb.GoodbyeResponse, error) {
	log.Printf("Received : %v", in.Name)
	return &pb.GoodbyeResponse{Message : "Goodbye " + in.Name}, nil
}

func (s *server) Get(ctx context.Context, in *pb.Key)(*pb.Value, error){
	log.Printf("Received : %v", in.Key)
	log.Printf("Value : %v", m[in.Key])
	return &pb.Value{Value : m[in.Key]}, nil
}

func heartbeat(port int){
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := pb.NewStorageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name : "heartbeat"})
	if err != nil {
		return
	}
	if err == nil{
		log.Printf("Greeting : %s", r.Message)
	}

}

func test(port int){
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := pb.NewStorageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayGoodbye(ctx, &pb.GoodbyeRequest{Name : "test"})
	if err != nil {
		return
	}
	if err == nil{
		log.Printf("Greeting : %s", r.Message)
	}
}

func sendHeartbeat(){ 
	//Executed By Leader
	// Send Heartbeat to other nodes
	for{
		for _, p := range portList{
			if p != *port{
				heartbeat(p)
				test(p)
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {

	m["ihate"] = "go"
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}


	defer lis.Close()
	s := grpc.NewServer()
	
	if *amLeader{
		go sendHeartbeat()
	}

	pb.RegisterStorageServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}


}