package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id   int64
	Age  int64
	Name string
}

type UserRepo interface {
	GetUser(ctx context.Context, id int64) (*User, error)
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

type UserUsecase struct {
	repo UserRepo
}

func (uc *UserUsecase) Get(ctx context.Context, id int64) (u *User, err error) {
	u, err = uc.repo.GetUser(ctx, id)
	if err != nil {
		return
	}
	return
}
