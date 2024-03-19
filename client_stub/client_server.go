package client_stub

import (
	"context"
	"fmt"
	ggrpc "google.golang.org/grpc"
	"net"
	"securechat-server/client_stub/grpc"
	"securechat-server/globals"
	requests "securechat-server/server/types"
)

type GRPCServer struct {
	grpc.UnimplementedClientServerCommsServer
	Requests chan<- requests.Request
	Response <-chan requests.Record
}

func (s *GRPCServer) SignUp(ctx context.Context, request *grpc.SignUpRequest) (*grpc.SignUpResponse, error) {
	return &grpc.SignUpResponse{}, nil
}

func (s *GRPCServer) SignUpChallengeResponse(ctx context.Context, request *grpc.SignUpChallengeResponseRequest) (*grpc.SignUpChallengeResponseResponse, error) {
	return &grpc.SignUpChallengeResponseResponse{}, nil
}

func (s *GRPCServer) Login(ctx context.Context, request *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	return &grpc.LoginResponse{}, nil
}

func (s *GRPCServer) FindUser(ctx context.Context, request *grpc.FindUserRequest) (*grpc.FindUserResponse, error) {
	return &grpc.FindUserResponse{}, nil
}

func NewGRPCClientServer(requests chan<- requests.Request, response <-chan requests.Record) {
	server := GRPCServer{Requests: requests, Response: response}
	s := ggrpc.NewServer()

	grpc.RegisterClientServerCommsServer(s, &server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", globals.ClientPort))
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
