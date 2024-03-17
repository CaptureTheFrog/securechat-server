package server

import (
	"securechat-server/server/dht"
	"securechat-server/server/types"
)

type Server struct {
	addr     string
	Requests chan<- types.Request
	Response <-chan types.Record
}

func NewServer(addr string, requests chan<- types.Request, response <-chan types.Record) *Server {

	// Start the internal DHT server on a separate goroutine
	go dht.NewDHTServer(addr)


	return &Server{addr: addr, Requests: requests, Response: response}
}

// Add a record to the User database
func (s *Server) Add(record types.Record) {
	r := types.Request{Type: types.PUT, Record: record}

	s.Requests <- r
}

// Get a record from the User database
func (s *Server) Get(username string) types.Record {
	r := types.Request{Type: types.GET, Record: types.Record{Username: username}}

	s.Requests <- r

	return <-s.Response
}

func (s *Server) Serve() {
	for {

	}
}
