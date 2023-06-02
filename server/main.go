package main

//TODO : When candidate dies during election (we need to start timer after vote)
// If no leader elected in that time period we regain the ability to vote.
// If candidate fails to be elected he tries again after a random time period.
// If candidate receives heartbeat from leader, it steps down.

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/wldyd423/keyvaluestorev1/pb"
	"google.golang.org/grpc"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	portList    = []int{50051, 50052, 50053, 50054}
	amLeader    = flag.Bool("leader", false, "Leader or not")
	amCandidate = flag.Bool("candidate", false, "Candidate or not")
	m           = make(map[string]string)
	bomb        = rand.Intn(15) + 7 //if this value is 0. The server will assume leader is dead.
	vote        = flag.Bool("vote", false, "Has made vote")
	leaderPort  = -1
)

type server struct {
	pb.UnimplementedStorageServer
}

func (s *server) Get(ctx context.Context, in *pb.Key) (*pb.Value, error) {
	log.Printf("Received : %v", in.Key)
	log.Printf("Value : %v", m[in.Key])
	return &pb.Value{Value: m[in.Key]}, nil
}

func SetToLeader(key string, value string) {
	tmpAddr := fmt.Sprintf("localhost:%d", leaderPort)
	log.Printf("Trying to connect to leader : %v", tmpAddr)
	conn, err := grpc.Dial(tmpAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()

	c := pb.NewStorageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err = c.Set(ctx, &pb.KeyValuePair{Key: key, Value: value})
	if err != nil {
		log.Fatalf("could not set : %v", err)
	}
}

func (s *server) Set(ctx context.Context, in *pb.KeyValuePair) (*pb.Value, error) {
	if *amLeader == false {
		log.Printf("Received : %v [not leader]", in.Key)
		log.Printf("Value : %v [not leader]", in.Value)
		go SetToLeader(in.Key, in.Value)
		return &pb.Value{Value: in.Value}, nil
	} else {
		log.Printf("Received : %v", in.Key)
		log.Printf("Value : %v", in.Value)
		m[in.Key] = in.Value
		return &pb.Value{Value: in.Value}, nil
	}
}

func (s *server) Delete(ctx context.Context, in *pb.Key) (*pb.Value, error) {
	log.Printf("Received : %v", in.Key)
	log.Printf("Value : %v", m[in.Key])
	ret := m[in.Key]
	delete(m, in.Key)
	return &pb.Value{Value: ret}, nil
}

func (s *server) GetAll(ctx context.Context, in *pb.Empty) (*pb.KeyValuePairList, error) {
	log.Printf("Received : GetAll")
	var list []*pb.KeyValuePair
	for k, v := range m {
		list = append(list, &pb.KeyValuePair{Key: k, Value: v})
	}
	return &pb.KeyValuePairList{KeyValuePair: list}, nil
}

func (s *server) Heartbeat(ctx context.Context, in *pb.KeyValuePairList) (*pb.Empty, error) {
	*vote = false
	*amCandidate = false
	if leaderPort != int(in.LeaderPort) {
		leaderPort = int(in.LeaderPort)
	}
	log.Printf("Received : Heartbeat from leader: %v", leaderPort)

	bomb = rand.Intn(10) + 7
	for _, kv := range in.KeyValuePair {
		m[kv.Key] = kv.Value // We assume only leader adds new values
	}
	return &pb.Empty{}, nil
}

func (s *server) Election(ctx context.Context, in *pb.RequestforVote) (*pb.Vote, error) {
	log.Printf("Received : Election")
	if *amCandidate == false && *vote == false {
		*vote = true
		bomb = rand.Intn(10) + 7
		*amCandidate = false
		return &pb.Vote{VoteGranted: true}, nil
	}
	return &pb.Vote{VoteGranted: false}, nil
}

func heartbeat_candidate(port int) {
	count := 0
	for _, p := range portList {
		if p == port {
			continue
		}
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", p), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewStorageClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Election(ctx, &pb.RequestforVote{})
		if err != nil {
			log.Fatalf("could not elect: %v", err)
		}
		if r.VoteGranted == true {
			count += 1
		} else if r.VoteGranted == false {
			count -= 1
		}
	}
	if count > 0 {
		fmt.Println("I am leader")
		*amLeader = true
		*amCandidate = false
	}
}

func heartbeat_leader(port int) {
	for _, p := range portList {
		if p == port {
			continue
		}
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", p), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewStorageClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		var list []*pb.KeyValuePair
		for k, v := range m {
			list = append(list, &pb.KeyValuePair{Key: k, Value: v})
		}
		c.Heartbeat(ctx, &pb.KeyValuePairList{KeyValuePair: list, LeaderPort: int32(port)})
	}

}

func tick() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		// case *amLeader == true:
		// 	continue
		case <-ticker.C:
			if *amLeader == true {
				continue
			} else {
				bomb -= 1
				fmt.Println("Ticked bomb. Remaining : ", bomb)
				if bomb < 0 {
					fmt.Println("Leader is dead")
					fmt.Println("I am candidate")
					*amCandidate = true
				}
			}
		}
	}
}

func sendHeartbeat() {
	for {
		if *amLeader == true {
			heartbeat_leader(*port)
		} else if *amCandidate == true {
			heartbeat_candidate(*port)
		}
		time.Sleep(5 * time.Second)
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

	go tick()
	go sendHeartbeat()

	pb.RegisterStorageServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
