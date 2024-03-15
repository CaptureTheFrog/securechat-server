package dht

import (
	"securechat-server/server/types"
)

type Server struct {
	id           *ID
	predecessor  *ID
	finger_table *FingerTable
	requests     <-chan types.Request
	response     chan<- types.Record
}

func NewDHTServer(addr string, requests <-chan types.Request, response chan<- types.Record) {
	server := &Server{id: NewID(addr), finger_table: NewFingerTable(), requests: requests, response: response}

	loop(server)
}

// Internal DHTserver loop to handle requests
func loop(server *Server) {
	for {
		select {
		case r := <-server.requests:
			switch r.Type {
			case types.GET:
				server.Get(r)
			case types.PUT:
				server.Put(r)
			}
		}
	}
}

// Get the record from the DHT.
func (s *Server) Get(request types.Request) {
	record := request.Record

	record.Address = "localhost"
	record.PublicKey = "publickey"

	s.response <- record
}

// Put the record into the DHT.
func (s *Server) Put(request types.Request) {

}
