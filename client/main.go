package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/wldyd423/keyvaluestorev1/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	fmt.Println("Usage : get [key], put [key] [value], delete [key], exit")
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()

	c := pb.NewStorageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// r, err := c.SayHello(ctx, &pb.HelloRequest{Name : *name})
	// if err != nil {
	// 	log.Fatalf("could not greet : %v", err)
	// }

	// log.Printf("Greeting : %s", r.Message)

	for {
		fmt.Print("$ ")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		word := strings.Fields(text)

		if strings.Compare(word[0], "get") == 0 {
			r, err := c.Get(ctx, &pb.Key{Key: word[1]})
			if err != nil {
				log.Fatalf("could not get : %v", err)
			}
			log.Printf("Received: %s", r.Value)
		} else if strings.Compare(word[0], "put") == 0 {
			r, err := c.Set(ctx, &pb.KeyValuePair{Key: word[1], Value: word[2]})
			if err != nil {
				log.Fatalf("could not put : %v", err)
			}
			log.Printf("Received: %s", r.Value)
		} else if strings.Compare(word[0], "delete") == 0 {
			r, err := c.Delete(ctx, &pb.Key{Key: word[1]})
			if err != nil {
				log.Fatalf("could not greet : %v", err)
			}
			log.Printf("Received: %s", r.Value)
		} else if strings.Compare(word[0], "getall") == 0 {
			r, err := c.GetAll(ctx, &pb.Empty{})
			if err != nil {
				log.Fatalf("could not greet : %v", err)
			}
			log.Printf("Received: %s", r.KeyValuePair)
		} else if strings.Compare(word[0], "exit") == 0 {
			break
		}

	}

	// r, err := c.Get(ctx, &pb.Key{Key : "ihate"})
	// if err != nil {
	// 	log.Fatalf("could not greet : %v", err)
	// }

	// log.Printf("Received: %s", r.Value)

}
