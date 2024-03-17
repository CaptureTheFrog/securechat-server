package dht

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "securechat-server/server/dht/grpc"
	"time"
)

func (s Server) Join(ctx context.Context, id *pb.ID) (*pb.ID, error) {
	//TODO implement me
	panic("implement me")
}

// GetSuccessor
/*
 Returns successor of the ID passed as parameter
*/
func (s Server) GetSuccessor(ctx context.Context, id *pb.ID) (*pb.ID, error) {
	println("1")
	succ := s.Successor()

	// If this node's successor is itself (it's the only node in the network)
	if succ.Equals(s.id) {
		return s.id.ToGRPC(), nil
	}

	// If the request is asking for this node's successor
	if s.id.Equals(IDFromGRPC(id)) {
		println("2")
		return &pb.ID{
			Address: &succ.name,
			Id:      succ.ID[:],
		}, nil
	}
	println("no")
	// If the ID sits between this node and its successor
	// return successor
	if IDFromGRPC(id).IsBetween(s.id, succ) {
		return succ.ToGRPC(), nil
	}

	// If the request is asking for a different node's successor
	// Pass to successor
	client, err := CreateGRPCClient(succ.name)
	if err != nil {
		return nil, err
	}
	// Get result from successor
	successor, err := client.GetSuccessor(client.Ctx, id)
	client.Cancel()
	if err != nil {
		return nil, err
	}

	println("3")
	return successor, nil
}

func (s Server) GetPredecessor(ctx context.Context, id *pb.ID) (*pb.ID, error) {
	pred := s.predecessor

	// If the request is asking for this node's predecessor
	if s.id.Equals(IDFromGRPC(id)) {
		if pred == nil {
			return nil, nil
		}

		return &pb.ID{
			Address: &pred.name,
			Id:      pred.ID[:],
		}, nil
	}

	// If the request is asking for a different node's predecessor
	// Pass to predecessor
	client, err := CreateGRPCClient(pred.name)
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

func (s Server) ChangeSuccessor(ctx context.Context, successor *pb.ChangeSuccessor) (*pb.Response, error) {
	// If the request is asking to change this node's successor
	if s.id.Equals(IDFromGRPC(successor.Id)) {
		// Change successor
		s.fingerTable.addEntry(1, IDFromGRPC(successor.NewSuccessor))
		return &pb.Response{}, nil
	}

	// If the request is asking to change a different node's successor
	// Pass to successor
	client, err := CreateGRPCClient(s.Successor().name)
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

func (s Server) ChangePredecessor(ctx context.Context, predecessor *pb.ChangePredecessor) (*pb.Response, error) {
	// If the request is asking to change this node's predecessor
	if s.id.Equals(IDFromGRPC(predecessor.Id)) {
		// Change predecessor
		s.predecessor = IDFromGRPC(predecessor.NewPredecessor)
		return &pb.Response{}, nil
	}

	// If the request is asking to change a different node's predecessor
	// Pass to predecessor
	client, err := CreateGRPCClient(s.predecessor.name)
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

func (s Server) Get(ctx context.Context, id *pb.ID) (*pb.Record, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Put(ctx context.Context, record *pb.Record) (*pb.Response, error) {
	//TODO implement me
	panic("implement me")
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return GRPCClient{
		ServerCommsClient: client,
		Ctx:               ctx,
		Cancel:            cancel,
	}, nil
}
