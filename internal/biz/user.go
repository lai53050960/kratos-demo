package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-client/api/helloworld/v1"
)

type User struct {
	Id   int64
	Age  int64
	Name string
}

type UserRepo interface {
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUserOrFail(ctx context.Context, id int64) (*User, error)
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
	u, err = uc.repo.GetUserOrFail(ctx, id)
	if err != nil {
		if errors.Reason(err) == v1.ErrorReason_ERROR_REASON_UNSPECIFIED.String() {
			return nil, errors.BadRequest("没有数据", "mysql")
		}
	}
	return
}
