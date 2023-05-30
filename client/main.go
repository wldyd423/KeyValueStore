package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"
	"bufio"
	"os"
	"strings"

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

// func getRPC(key string){

// }

func main() {
	fmt.Println("Hello, World!")
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()

	c := pb.NewStorageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	// r, err := c.SayHello(ctx, &pb.HelloRequest{Name : *name})
	// if err != nil {
	// 	log.Fatalf("could not greet : %v", err)
	// }

	// log.Printf("Greeting : %s", r.Message)
	
	for{
		fmt.Print("$ ")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		

		if strings.Compare(strings.SplitAfter(text," ")[0], "get ") == 0{
			r, err := c.Get(ctx, &pb.Key{Key : strings.SplitAfter(text," ")[1]})
			if err != nil {
				log.Fatalf("could not greet : %v", err)
			}
			log.Printf("Received: %s", r.Value)
		}else if strings.Compare(strings.SplitAfter(text," ")[0], "put ") == 0{
		}else if strings.Compare(strings.SplitAfter(text," ")[0], "delete ") == 0{
		}else if strings.Compare(strings.SplitAfter(text," ")[0], "exit") == 0{
			break
		}

	}


	// r, err := c.Get(ctx, &pb.Key{Key : "ihate"})
	// if err != nil {
	// 	log.Fatalf("could not greet : %v", err)
	// }

	// log.Printf("Received: %s", r.Value)


}