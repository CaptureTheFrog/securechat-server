package main

import (
	"flag"
	"fmt"
	. "securechat-server/client_stub"
	. "securechat-server/globals"
	server "securechat-server/server"
	"securechat-server/server/types"
)

func main() {
	// Generate 2 free ports and the IP address of the server
	serverGRPCPort, ip, err := FindFreePort()
	if err != nil {
		panic(err)
	}
	clientGRPCPort, _, err := FindFreePort()
	if err != nil {
		panic(err)
	}

	flag.StringVar(&Method, "method", "Create", "Indicates whether the server network is being joined or created")
	flag.StringVar(&JoinAddress, "address", "", "JoinAddress of the server to join the network from")
	flag.IntVar(&GRPCPort, "grpc-port", serverGRPCPort, "Port for the gRPC server to listen on")
	flag.StringVar(&ServerAddress, "server-address", ip.String(), "Address of this server")
	flag.IntVar(&ClientPort, "client-port", clientGRPCPort, "Port for the gRPC client to connect to")

	flag.Parse()

	fmt.Printf("Server Address: %s:%d\n", ServerAddress, GRPCPort)

	// Make channel for sending requests and receiving responses from client-server GRPC stub
	requests := make(chan types.Request, 10)
	response := make(chan types.Record, 10)

	s := server.NewServer(requests, response)

	go NewGRPCClientServer(requests, response)

	s.Serve()
}
