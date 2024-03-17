package main

import (
	"flag"
	"log"
	. "securechat-server/globals"
	. "securechat-server/server"
	"securechat-server/server/dht"
	. "securechat-server/client_stub"
	server "securechat-server/server"
	"securechat-server/server/types"
)

func main() {
	flag.StringVar(&Method, "method", "Create", "Indicates whether the server network is being joined or created")
	flag.StringVar(&JoinAddress, "address", "", "JoinAddress of the server to join the network from")
	flag.IntVar(&GRPCPort, "grpc-port", 50051, "Port for the gRPC server to listen on")
	flag.StringVar(&ServerAddress, "server-address", "127.0.0.1:50051", "Address of this server")

	flag.Parse()

	NewServer(ServerAddress)

	id := dht.NewID(ServerAddress)
	log.Printf("ID: %x", id.ID)
	// Make channel for sending requests and receiving responses from client-server GRPC stub
	requests := make(chan types.Request, 10)
	response := make(chan types.Record, 10)

	s := server.NewServer(":50051", requests, response)

	go NewGRPCClientServer(requests, response)

	s.Serve()
}
