package main

import (
	"flag"
	"log"
	. "securechat-server/globals"
	. "securechat-server/server"
	"securechat-server/server/dht"
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

	//record := s.Get("test")

	for {

	}
	//println(record.Address)
}
