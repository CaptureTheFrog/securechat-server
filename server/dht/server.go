package dht

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"securechat-server/globals"
	pb "securechat-server/server/dht/grpc"
	"securechat-server/server/dht/records"
	"securechat-server/server/dht/types"
	"sync"
)

type Server struct {
	id          *types.ID
	predecessor *types.ID
	fingerTable *types.FingerTable
	records     *records.Records
	mu          sync.RWMutex
	pb.UnimplementedServerCommsServer
}

// NewDHTServer
/*
This function creates a new DHT server struct and starts the internal gRPC server for server to server communication.
Depending on the method flag, the server will either create a new network or join an existing one.
*/
func NewDHTServer(addr string) {
	server := &Server{id: types.NewID(addr), fingerTable: types.NewFingerTable(), records: records.NewRecords()}
	s := grpc.NewServer()

	// Create listener on the port flagged (default 50051)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", globals.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the gRPC server
	pb.RegisterServerCommsServer(s, server)

	// Determine the method to be used (Create or Join)
	switch globals.Method {
	case "Create":
		server.CreateNetwork()
	case "Join":
		// If the address flag is not set, the server cannot join the network
		if globals.JoinAddress == "" {
			log.Fatalf("JoinAddress flag required for Join method")
		}

		server.JoinNetwork()
	default:
		log.Fatalf("Invalid method: %s", globals.Method)
	}

	// Start maintainer
	go server.Maintain()

	// Start listening on the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Successor() *types.ID {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.fingerTable.GetEntry(1)
}

func (s *Server) Predecessor() *types.ID {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.predecessor
}

// JoinNetwork
/*
This function allows a server to join an existing network. It uses the address of a pre-existing server
(given as a flag) to find its own position in the network. It then updates its predecessor and successor to let them
know of its existence. Finally, it updates its records to reflect the new network.

Before joining: predecessor -> successor
After joining: predecessor -> new server -> successor

If there is only one server in the network, the server's successor is itself and its predecessor is set to nil.
Before joining: nil -> server -> server
After joining: new server -> server -> new server
*/
func (s *Server) JoinNetwork() {
	s.mu.Lock()
	defer s.mu.Unlock()

	// ask the server at the address to find the successor of the new server
	client, err := CreateGRPCClient(globals.JoinAddress)
	if err != nil {
		log.Fatalf("failed to join network 1: %v", err)
	}

	succ, err := client.GetSuccessor(client.Ctx, s.id.ToGRPC())
	client.Cancel()
	if err != nil {
		log.Fatalf("failed to join network 2: %v", err)
	}

	// Convert successor from grpc.ID (rpc schema) to dht.ID
	successor := types.NewID(*succ.Address)

	client, err = CreateGRPCClient(*succ.Address)
	if err != nil {
		log.Fatalf("failed to join network 3: %v", err)
	}
	var predecessor *types.ID

	// Get predecessor
	// Server returns NotFound error if it is the only one in the network
	pred, err := client.GetPredecessor(client.Ctx, succ)
	if err != nil {
		// If the server is the only one in the network, it returns not found because it doesn't have a predecessor
		if status.Code(err) != codes.NotFound {
			log.Fatalf("failed to join network 4: %v", err)
		} else {
			// Set predecessor to the successor if the server is the only one in the network
			predecessor = successor
		}
	} else {
		predecessor = types.NewID(*pred.Address)
	}

	// Change the successor's predecessor to this
	_, err = client.ChangePredecessor(client.Ctx, &pb.ChangePredecessor{
		Id:             succ,
		NewPredecessor: s.id.ToGRPC(),
	})
	if err != nil {
		log.Printf("failed to join network 5: %v", err)
	}
	client.Cancel()

	// Tell successor's previous predecessor that this server is the new successor
	client, err = CreateGRPCClient(predecessor.Name)
	if err != nil {
		log.Fatalf("failed to join network: %v", err)
	}
	_, err = client.ChangeSuccessor(client.Ctx, &pb.ChangeSuccessor{
		Id:           predecessor.ToGRPC(),
		NewSuccessor: s.id.ToGRPC(),
	})
	if err != nil {
		log.Printf("failed to join network 5: %v", err)
	}
	client.Cancel()

	// Update the finger table and predecessor field
	s.fingerTable.AddEntry(1, successor)
	s.predecessor = predecessor

	println("done")
}

// CreateNetwork
// This function is called when the server is the first in the network.
// It creates the network and adds itself to the finger table as its own successor.
func (s *Server) CreateNetwork() {
	s.fingerTable.AddEntry(1, s.id)
}

func (s *Server) UpdatePredecessor(newPredecessor *types.ID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.predecessor = newPredecessor
}

func (s *Server) UpdateSuccessor(newSuccessor *types.ID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.fingerTable.AddEntry(1, newSuccessor)
}
