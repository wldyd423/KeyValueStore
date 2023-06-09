package main
//TODO : When candidate dies during election (we need to start timer after vote) 
// If no leader elected in that time period we regain the ability to vote.
// If candidate fails to be elected he tries again after a random time period.
// If candidate receives heartbeat from leader, it steps down.

// Leader heartbeat holds keyvaluepair of leader. followers merge their map with leader's map.
// Any put request is sent to leader.

//TODO(MAYBE): SCALE UP PORTLIST METHOD currently very very crude


import (
	"context"
	"flag"
	"fmt"
	"net"
	"log"
	"time"
	"math/rand"

	"google.golang.org/grpc"
	pb "github.com/wldyd423/keyvaluestorev1/pb"
)

var (
	port = flag.Int("port", 50051, "The server port")
	portList = []int{50051, 50052, 50053, 50054}
	amLeader = flag.Bool("leader", false, "Leader or not")
	amCandidate = flag.Bool("candidate", false, "Candidate or not")
	m = make(map[string]string)
	bomb = rand.Intn(15) + 7 //if this value is 0. The server will assume leader is dead.
	vote = flag.Bool("vote", false, "Has made vote")
)

type server struct {
	pb.UnimplementedStorageServer
}



func (s *server) Get(ctx context.Context, in *pb.Key)(*pb.Value, error){
	log.Printf("Received : %v", in.Key)
	log.Printf("Value : %v", m[in.Key])
	return &pb.Value{Value : m[in.Key]}, nil
}

func (s *server) Set(ctx context.Context, in *pb.KeyValuePair)(*pb.Value, error){
	log.Printf("Received : %v", in.Key)
	log.Printf("Value : %v", in.Value)
	m[in.Key] = in.Value
	return &pb.Value{Value : in.Value}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.Key)(*pb.Value, error){
	log.Printf("Received : %v", in.Key)
	log.Printf("Value : %v", m[in.Key])
	ret := m[in.Key]
	delete(m, in.Key)
	return &pb.Value{Value : ret}, nil
}

func (s *server) GetAll(ctx context.Context, in *pb.Empty)(*pb.KeyValuePairList, error){
	log.Printf("Received : GetAll")
	var list []*pb.KeyValuePair
	for k, v := range m{
		list = append(list, &pb.KeyValuePair{Key : k, Value : v})
	}
	return &pb.KeyValuePairList{KeyValuePair : list}, nil
}

func (s *server) Heartbeat(ctx context.Context, in *pb.KeyValuePairList)(*pb.Empty, error){
	*vote = false
	*amCandidate = false
	log.Printf("Received : Heartbeat")
	bomb = rand.Intn(10) + 7
	for _, kv := range in.KeyValuePair{
		m[kv.Key] = kv.Value // We assume only leader adds new values
	}
	return &pb.Empty{}, nil
}

func (s *server) Election(ctx context.Context, in *pb.RequestforVote)(*pb.Vote, error){
	log.Printf("Received : Election")
	if *amCandidate == false && *vote == false{
		*vote = true
		bomb = rand.Intn(10) + 7
		*amCandidate = false
		return &pb.Vote{VoteGranted : true}, nil
	}
	return &pb.Vote{VoteGranted : false}, nil
}

func heartbeat_candidate(port int){
	count := 0
	for _, p := range portList{
		if p == port{
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
			log.Fatalf("could not greet: %v", err)
		}
		if r.VoteGranted == true{
			count += 1
		}else if r.VoteGranted == false{
			count -= 1
		}
	}
	if count > 0{
		fmt.Println("I am leader")
		*amLeader = true
		*amCandidate = false
	}
}

func heartbeat_leader(port int){
	for _, p := range portList{
		if p == port{
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
		for k, v := range m{
			list = append(list, &pb.KeyValuePair{Key : k, Value : v})
		}
		c.Heartbeat(ctx, &pb.KeyValuePairList{KeyValuePair : list})
	}
	
}

func tick(){
	ticker := time.NewTicker(1 * time.Second)
	for{
		select{
		// case *amLeader == true:
		// 	continue
		case <- ticker.C:
			if *amLeader == true{
				continue
			}else{
				bomb -= 1
				fmt.Println("Ticked bomb. Remaining : ", bomb)
				if bomb < 0{
					fmt.Println("Leader is dead")
					fmt.Println("I am candidate")
					*amCandidate = true
				}
			}
		}
	}
}

func sendHeartbeat(){ 
	for{
		if *amLeader == true{
			heartbeat_leader(*port)
		}else if *amCandidate == true{
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