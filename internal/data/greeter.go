package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/genproto/protobuf/field_mask"
	"kratos-client/api/user"
	"kratos-client/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {

	req := &user.UpdateUserRequest{
		User: &user.UserInfo{
			Name: "name",
			Age:  2,
			Id:   1,
		},
		Mark: &field_mask.FieldMask{
			Paths: []string{"name", "age"},
		},
	}

	_, err := r.data.uc.UpdateUser(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (r *greeterRepo) SayHello(ctx context.Context, g *biz.Greeter) (*biz.UserGrpc, error) {
	getUser, err := r.data.uc.GetUser(ctx, &user.GetUserRequest{
		Id:   1,
		Name: "haha",
	})
	if err != nil {
		return nil, err
	}

	return &biz.UserGrpc{
		Id:   getUser.User.Id,
		Age:  1,
		Name: getUser.User.Name,
	}, nil
}
