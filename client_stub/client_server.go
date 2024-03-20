package client_stub

import (
	"context"
	"encoding/binary"
	"fmt"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
	"securechat-server/client_stub/grpc"
	"securechat-server/globals"
	"securechat-server/server/dht/records"
	requests "securechat-server/server/types"
	"sync"
)

type GRPCServer struct {
	mu sync.Mutex
	grpc.UnimplementedClientServerCommsServer
	Requests chan<- requests.Request
	Response <-chan records.Record
}

func (s *GRPCServer) SignUp(ctx context.Context, request *grpc.SignUpRequest) (*grpc.SignUpResponse, error) {
	// Create record struct
	//_ = s.getUser(request.Username)

	record := records.Record{
		Username:       request.Username,
		Address:        "",
		PublicKeyChat:  request.PublicKeyChat,
		PublicKeyLogin: request.PublicKeyLogin,
	}

	// Get IP address of the client
	_ = getClientIP(ctx)

	// Create request struct
	req := requests.Request{
		Type:   requests.PUT,
		Record: record,
	}

	s.mu.Lock()
	s.Requests <- req
	s.mu.Unlock()

	return &grpc.SignUpResponse{}, nil
}

func (s *GRPCServer) SignUpChallengeResponse(ctx context.Context, request *grpc.SignUpChallengeResponseRequest) (*grpc.SignUpChallengeResponseResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return &grpc.SignUpChallengeResponseResponse{}, nil
}

func (s *GRPCServer) Login(ctx context.Context, request *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	// Create record struct
	_ = s.getUser(request.Username)

	return &grpc.LoginResponse{}, nil
}

func (s *GRPCServer) FindUser(ctx context.Context, request *grpc.FindUserRequest) (*grpc.FindUserResponse, error) {
	// Create record struct
	record := s.getUser(request.Username)

	if record.Username == "" {
		// TODO: Add error
		return &grpc.FindUserResponse{}, nil
	}

	// TODO: IP address
	return &grpc.FindUserResponse{
		Username:      record.Username,
		Address:       0,
		PublicKeyChat: record.PublicKeyChat,
	}, nil
}

func NewGRPCClientServer(requests chan<- requests.Request, response <-chan records.Record) {
	server := GRPCServer{Requests: requests, Response: response}
	s := ggrpc.NewServer()

	grpc.RegisterClientServerCommsServer(s, &server)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", globals.ServerAddress, globals.ClientPort))
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func (s *GRPCServer) getUser(username string) records.Record {
	// Create record struct
	record := records.Record{
		Username:       username,
		Address:        "",
		PublicKeyChat:  make([]byte, 0),
		PublicKeyLogin: make([]byte, 0),
	}

	// Create request struct
	req := requests.Request{
		Type:   requests.GET,
		Record: record,
	}

	s.mu.Lock()
	s.Requests <- req
	record = <-s.Response
	s.mu.Unlock()

	return record
}

func getClientIP(ctx context.Context) net.Addr {
	peerAddr, ok := peer.FromContext(ctx)
	if ok {
		return peerAddr.Addr
	}
	return nil
}
