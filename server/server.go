package server

import (
	"securechat-server/server/dht"
	"securechat-server/server/types"
)

type Server struct {
	addr       string
	dbRequests chan<- types.Request
	dbResponse <-chan types.Record
}

func NewServer(addr string) *Server {
	// Make channel for sending requests and receiving responses from the internal DHT server
	requests := make(chan types.Request, 5)
	response := make(chan types.Record, 5)

	// Start the internal DHT server on a separate goroutine
	go dht.NewDHTServer(addr, requests, response)

	return &Server{addr: addr, dbRequests: requests, dbResponse: response}
}

// Add a record to the User database
func (s *Server) Add(record types.Record) {
	r := types.Request{Type: types.PUT, Record: record}

	s.dbRequests <- r
}

// Get a record from the User database
func (s *Server) Get(username string) types.Record {
	r := types.Request{Type: types.GET, Record: types.Record{Username: username}}

	s.dbRequests <- r

	return <-s.dbResponse
}
