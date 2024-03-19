package main

import (
	"flag"
	. "securechat-server/client_stub"
	. "securechat-server/globals"
	server "securechat-server/server"
	"securechat-server/server/types"
)

func main() {
	flag.StringVar(&Method, "method", "Create", "Indicates whether the server network is being joined or created")
	flag.StringVar(&JoinAddress, "address", "", "JoinAddress of the server to join the network from")
	flag.IntVar(&GRPCPort, "grpc-port", 50051, "Port for the gRPC server to listen on")
	flag.StringVar(&ServerAddress, "server-address", "127.0.0.1:50051", "Address of this server")
	flag.IntVar(&ClientPort, "client-port", 50050, "Port for the gRPC client to connect to")

	flag.Parse()

	// Make channel for sending requests and receiving responses from client-server GRPC stub
	requests := make(chan types.Request, 10)
	response := make(chan types.Record, 10)

	s := server.NewServer(ServerAddress, requests, response)

	go NewGRPCClientServer(requests, response)

	s.Serve()
}
