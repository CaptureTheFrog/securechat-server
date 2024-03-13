package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "securechat-server/grpc"
	"time"
)

const (
	defaultName = "bossman"
)

var (
	addr = flag.String("addr", "localhost:50051", "address to connect to")
	name = flag.String("name", defaultName, "name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServerCommsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.Hi(ctx, &pb.Test{Message: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
