package server

import (
	"securechat-server/server/dht"
	"securechat-server/server/types"
)

type Server struct {
	addr     string
	Requests <-chan types.Request
	Response chan<- types.Record
}

func NewServer(addr string, requests <-chan types.Request, response chan<- types.Record) *Server {

	// Start the internal DHT server on a separate goroutine
	go dht.NewDHTServer(addr)

	return &Server{addr: addr, Requests: requests, Response: response}
}

// Add a record to the User database
func (s *Server) Add(record types.Record) {

}

// Get a record from the User database
func (s *Server) Get(username string) types.Record {
	return types.Record{}
}

func (s *Server) Serve() {
	for {
		select {
		case r := <-s.Requests:
			switch r.Type {
			case types.GET:
				s.Get(r.Record.Username)
			case types.PUT:
				s.Add(r.Record)
			}
		}
	}
}
