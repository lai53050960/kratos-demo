package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-client/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	var user User
	err := ar.data.db.WithContext(ctx).Where("a=?", 0).First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 报错
			return nil, errors.New(400, "reason", "没有找到数据")
		} else {
			// 只记记录
			ar.log.Error(err)
			return nil, errors.BadRequest("数据库", "系统错误")
		}
	}

	return &biz.User{
		Id:   int64(user.ID),
		Age:  1,
		Name: user.Name,
	}, nil
}
