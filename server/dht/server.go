package dht

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"securechat-server/globals"
	pb "securechat-server/server/dht/grpc"
	"sync"
)

type Server struct {
	id          *ID
	predecessor *ID
	fingerTable *FingerTable
	mu          sync.Mutex
	pb.UnimplementedServerCommsServer
}

// NewDHTServer
/*
This function creates a new DHT server struct and starts the internal gRPC server for server to server communication.
Depending on the method flag, the server will either create a new network or join an existing one.
*/
func NewDHTServer(addr string) {
	server := &Server{id: NewID(addr), fingerTable: NewFingerTable()}
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

	// Start listening on the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s Server) Successor() *ID {
	return s.fingerTable.getEntry(1)
}

func (s Server) Predecessor() *ID {
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
func (s Server) JoinNetwork() {
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
	successor := NewID(*succ.Address)

	client, err = CreateGRPCClient(*succ.Address)
	if err != nil {
		log.Fatalf("failed to join network 3: %v", err)
	}
	var predecessor *ID

	// If there is only one server in the network (the pre-existing server is its own successor)
	if globals.JoinAddress == successor.name {
		// The pre-existing server will be this new server's predecessor and successor
		predecessor = successor
	} else {
		// Ask successor for their predecessor
		pred, err := client.GetPredecessor(client.Ctx, succ)
		if err != nil {
			log.Fatalf("failed to join network 4: %v", err)
		}
		predecessor = NewID(*pred.Address)
	}

	// Change the successor's predecessor to this
	client.ChangePredecessor(client.Ctx, &pb.ChangePredecessor{
		Id:             succ,
		NewPredecessor: s.id.ToGRPC(),
	})
	client.Cancel()

	// Tell successor's previous predecessor that this server is the new successor
	client, err = CreateGRPCClient(predecessor.name)
	if err != nil {
		log.Fatalf("failed to join network: %v", err)
	}
	client.ChangeSuccessor(client.Ctx, &pb.ChangeSuccessor{
		Id:           predecessor.ToGRPC(),
		NewSuccessor: s.id.ToGRPC(),
	})
	client.Cancel()

	// Update the finger table and predecessor field
	s.fingerTable.addEntry(1, successor)
	s.predecessor = predecessor

	println("done")
}

// CreateNetwork
// This function is called when the server is the first in the network.
// It creates the network and adds itself to the finger table as its own successor.
func (s Server) CreateNetwork() {
	s.fingerTable.addEntry(1, s.id)
}
