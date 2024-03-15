package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	. "securechat-server/client_server"
	pb "securechat-server/grpc"
	. "securechat-server/server"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	// to sort out merge code
	// client_server branch
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer()
	pb.RegisterClientServerCommsServer(gs, &GRPCServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// dht branch
	s := NewServer("localhost:50051")

	record := s.Get("test")

	println(record.Address)
}
