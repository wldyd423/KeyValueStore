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
	"github.com/DistributedClocks/GoVector/govec"
)

const (
	defaultName = "wldyd"
)

var (
	addr = flag.String("addr", "localhost:50051", "The server address")
	name = flag.String("name", defaultName, "The name to be greeted")
	logger = govec.InitGoVector(fmt.Sprintf("client->%v", *addr), fmt.Sprintf("client-%v.log", *addr), govec.GetDefaultConfig())
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
			msg := logger.PrepareSend("get", word[1], govec.GetDefaultLogOptions())
			r, err := c.Get(ctx, &pb.Key{Key: word[1], Log: msg})
			if err != nil {
				log.Fatalf("could not get : %v", err)
			}
			logger.UnpackReceive("get", r.Log, nil, govec.GetDefaultLogOptions())
			log.Printf("Received: %s", r.Value)
		} else if strings.Compare(word[0], "put") == 0 {
			msg := logger.PrepareSend("put", word[1], govec.GetDefaultLogOptions())
			r, err := c.Set(ctx, &pb.KeyValuePair{Key: word[1], Value: word[2], Log: msg})
			if err != nil {
				log.Fatalf("could not put : %v", err)
			}
			logger.UnpackReceive("put", r.Log, nil, govec.GetDefaultLogOptions())
			log.Printf("Received: %s", r.Value)
		} else if strings.Compare(word[0], "delete") == 0 {
			msg := logger.PrepareSend("delete", word[1], govec.GetDefaultLogOptions())
			r, err := c.Delete(ctx, &pb.Key{Key: word[1], Log: msg})
			if err != nil {
				log.Fatalf("could not greet : %v", err)
			}
			logger.UnpackReceive("delete", r.Log, nil, govec.GetDefaultLogOptions())
			log.Printf("Received: %s", r.Value)
		} else if strings.Compare(word[0], "getall") == 0 {
			msg := logger.PrepareSend("getall", "", govec.GetDefaultLogOptions())
			r, err := c.GetAll(ctx, &pb.Empty{Log: msg})
			if err != nil {
				log.Fatalf("could not greet : %v", err)
			}
			logger.UnpackReceive("getall", r.Log, nil, govec.GetDefaultLogOptions())
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
