package auth

import (
	"context"
	ssov1 "github.com/adametsderschopfer/gRPC-auth-service-contracts/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Register(ctx context.Context, request *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serverAPI) IsAdmin(ctx context.Context, request *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serverAPI) Logout(ctx context.Context, request *ssov1.LogoutRequest) (*ssov1.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serverAPI) mustEmbedUnimplementedAuthServer() {
	//TODO implement me
	panic("implement me")
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}
