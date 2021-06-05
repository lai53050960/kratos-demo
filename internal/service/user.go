package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-client/internal/biz"

	pb "kratos-client/api/user"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(usecase *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  usecase,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	get, err := s.uc.Get(ctx, 1)

	if err != nil {

		return nil, err
	}
	return &pb.GetUserReply{
		User: &pb.Users{
			Id:   get.Id,
			Name: get.Name,
			Age:  get.Age,
		},
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
