package server

import (
	"securechat-server/server/dht"
	"securechat-server/server/types"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {

	// Start the internal DHT server on a separate goroutine
	go dht.NewDHTServer(addr)

	return &Server{addr: addr}
}

// Add a record to the User database
func (s *Server) Add(record types.Record) {
	//TODO implement me
}

// Get a record from the User database
func (s *Server) Get(username string) types.Record {
	// TODO implement me
	return types.Record{}
}
