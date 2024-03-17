package main

import (
	. "securechat-server/client_stub"
	server "securechat-server/server"
	"securechat-server/server/types"
)

func main() {
	// Make channel for sending requests and receiving responses from client-server GRPC stub
	requests := make(chan types.Request, 10)
	response := make(chan types.Record, 10)

	s := server.NewServer(":50051", requests, response)

	go NewGRPCClientServer(requests, response)

	s.Serve()
}
