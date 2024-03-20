package dht

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "securechat-server/server/dht/grpc"
	"securechat-server/server/dht/records"
	"securechat-server/server/dht/types"
	"time"
)

// GetSuccessor
/*
 Returns successor of the ID passed as parameter
*/
func (s *Server) GetSuccessor(ctx context.Context, id *pb.ID) (*pb.ID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	succ := s.Successor()

	// If this node's successor is itself (it's the only node in the network)
	if succ.Equals(s.id) {
		return s.id.ToGRPC(), nil
	}

	// If the request is asking for this node's successor
	if *id.Address == "" || s.id.Equals(types.IDFromGRPC(id)) {
		return &pb.ID{
			Address: &succ.Name,
			Id:      succ.ID[:],
		}, nil
	}

	// If the ID sits between this node and its successor
	// return successor
	if types.IDFromGRPC(id).IsBetween(s.id, succ) {
		return succ.ToGRPC(), nil
	}

	// If the request is asking for a different node's successor
	// Pass to successor
	client, err := CreateGRPCClient(succ.Name)
	if err != nil {
		return nil, err
	}
	// Get result from successor
	successor, err := client.GetSuccessor(client.Ctx, id)
	client.Cancel()
	if err != nil {
		return nil, err
	}

	return successor, nil
}

func (s *Server) GetPredecessor(ctx context.Context, id *pb.ID) (*pb.ID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pred := s.predecessor

	// If the request is asking for this node's predecessor
	if s.id.Equals(types.IDFromGRPC(id)) || *id.Address == "" {
		if pred == nil {
			return nil, nil
		}

		return &pb.ID{
			Address: &pred.Name,
			Id:      pred.ID[:],
		}, nil
	}

	// If the request is asking for a different node's predecessor
	// Pass to predecessor
	client, err := CreateGRPCClient(pred.Name)
	if err != nil {
		return nil, err
	}

	// Get result from predecessor
	predecessor, err := client.GetPredecessor(client.Ctx, id)
	client.Cancel()
	if err != nil {
		return nil, err
	}

	return predecessor, nil
}

func (s *Server) ChangeSuccessor(ctx context.Context, successor *pb.ChangeSuccessor) (*pb.Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// If the request is asking to change this node's successor
	if s.id.Equals(types.IDFromGRPC(successor.Id)) || *successor.Id.Address == "" {
		// Change successor
		s.fingerTable.AddEntry(1, types.IDFromGRPC(successor.NewSuccessor))
		return &pb.Response{}, nil
	}

	// If the request is asking to change a different node's successor
	// Pass to successor
	client, err := CreateGRPCClient(s.Successor().Name)
	if err != nil {
		return nil, err
	}
	changeSuccessor, err := client.ChangeSuccessor(client.Ctx, successor)
	client.Cancel()
	if err != nil {
		return nil, err
	}

	return changeSuccessor, nil
}

func (s *Server) ChangePredecessor(ctx context.Context, predecessor *pb.ChangePredecessor) (*pb.Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// If the request is asking to change this node's predecessor
	if *predecessor.Id.Address == "" || s.id.Equals(types.IDFromGRPC(predecessor.Id)) {
		// Change predecessor
		s.predecessor = types.IDFromGRPC(predecessor.NewPredecessor)
		return &pb.Response{}, nil
	}

	// If the request is asking to change a different node's predecessor
	// Pass to predecessor
	client, err := CreateGRPCClient(s.predecessor.Name)
	if err != nil {
		return nil, err
	}
	changePredecessor, err := client.ChangePredecessor(client.Ctx, predecessor)
	client.Cancel()
	if err != nil {
		return nil, err
	}

	return changePredecessor, nil
}

func (s *Server) Get(ctx context.Context, id *pb.ID) (*pb.Record, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	dhtID := types.IDFromGRPC(id)

	// If id is between this node and its successor
	// or if this is the only node in the network (therefore all records should be stored here)
	if dhtID.IsBetween(s.id, s.Successor()) || s.isOnlyNodeInNetwork() {
		// Return record
		record := s.records.Get(dhtID)

		return &pb.Record{
			Username:       &record.Username,
			Address:        &record.Address,
			PublicKeyLogin: record.PublicKeyLogin,
			PublicKeyChat:  record.PublicKeyChat,
		}, nil
	}

	// Pass to successor
	client, err := CreateGRPCClient(s.Successor().Name)
	if err != nil {
		return nil, err
	}
	defer client.Cancel()

	return client.Get(client.Ctx, id)
}

func (s *Server) Put(ctx context.Context, record *pb.Record) (*pb.Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	dhtID := types.NewID(*record.Username)

	// If id is between this node and its successor
	// or if this is the only node in the network (therefore all records should be stored here)
	if dhtID.IsBetween(s.id, s.Successor()) || s.isOnlyNodeInNetwork() {
		// Store record
		s.records.Add(records.Record{
			Username:       *record.Username,
			Address:        *record.Address,
			PublicKeyLogin: record.PublicKeyLogin,
			PublicKeyChat:  record.PublicKeyChat,
		})
		return &pb.Response{}, nil
	}

	// Pass to successor
	client, err := CreateGRPCClient(s.Successor().Name)
	if err != nil {
		return nil, err
	}
	defer client.Cancel()

	return client.Put(client.Ctx, record)
}

type GRPCClient struct {
	pb.ServerCommsClient
	Ctx    context.Context
	Cancel context.CancelFunc
}

func CreateGRPCClient(addr string) (GRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return GRPCClient{}, err
	}

	client := pb.NewServerCommsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

	return GRPCClient{
		ServerCommsClient: client,
		Ctx:               ctx,
		Cancel:            cancel,
	}, nil
}
