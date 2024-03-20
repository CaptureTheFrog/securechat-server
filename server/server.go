package server

import (
	"fmt"
	"securechat-server/globals"
	"securechat-server/server/dht"
	pb "securechat-server/server/dht/grpc"
	"securechat-server/server/dht/records"
	dhtTypes "securechat-server/server/dht/types"
	"securechat-server/server/types"
)

type Server struct {
	addr     string
	Requests <-chan types.Request
	Response chan<- records.Record
}

func NewServer(requests <-chan types.Request, response chan<- records.Record) *Server {

	serverAddress := fmt.Sprintf("%s:%d", globals.ServerAddress, globals.GRPCPort)

	// Start the internal DHT server on a separate goroutine
	go dht.NewDHTServer(serverAddress)

	return &Server{addr: serverAddress, Requests: requests, Response: response}
}

// Add a record to the User database
func (s *Server) Add(record records.Record) {
	client, err := dht.CreateGRPCClient(s.addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Cancel()

	_, _ = client.Put(client.Ctx, &pb.Record{
		Username:       &record.Username,
		Address:        &record.Address,
		PublicKeyLogin: record.PublicKeyLogin,
		PublicKeyChat:  record.PublicKeyChat,
	})
}

// Get a record from the User database
func (s *Server) Get(username string) records.Record {
	id := dhtTypes.NewID(username)

	client, err := dht.CreateGRPCClient(s.addr)
	if err != nil {
		fmt.Println(err)
		return *records.NewRecord()
	}
	defer client.Cancel()

	record, err := client.Get(client.Ctx, id.ToGRPC())
	if err != nil {
		fmt.Println(err)
		return *records.NewRecord()
	}

	return records.Record{
		Username:       record.GetUsername(),
		Address:        record.GetAddress(),
		PublicKeyLogin: record.PublicKeyLogin,
		PublicKeyChat:  record.PublicKeyChat,
	}
}

func (s *Server) Serve() {
	for {
		select {
		case r := <-s.Requests:
			switch r.Type {
			case types.GET:
				r := s.Get(r.Record.Username)
				s.Response <- r
			case types.PUT:
				s.Add(r.Record)
			}
		}
	}
}
