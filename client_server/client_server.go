package client_server

import (
	"context"
	pb "securechat-server/grpc"
)

type GRPCServer struct {
	pb.UnimplementedClientServerCommsServer
}

func (s *GRPCServer) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) SignUpChallengeResponse(ctx context.Context, request *pb.SignUpChallengeResponseRequest) (*pb.SignUpChallengeResponseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) FindUser(ctx context.Context, request *pb.FindUserRequest) (*pb.FindUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
