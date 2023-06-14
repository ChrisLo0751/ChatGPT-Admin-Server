package service

import (
	"context"

	pb "chatgpt-admin-server/api/user/v1"
	"chatgpt-admin-server/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	user *biz.UserUsecase
}

func NewUserService(user *biz.UserUsecase) *UserService {
	return &UserService{user: user}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	user := &biz.User{
		Username: req.Body.Username,
		Password: req.Body.Password,
		Email:    req.Body.Email,
	}

	err := s.user.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
func (s *UserService) DeletetUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
