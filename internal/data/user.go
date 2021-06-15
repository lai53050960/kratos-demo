package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "kratos-client/api/helloworld/v1"
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
	err := ar.data.db.WithContext(ctx).Where("id=?", 0).First(&user, id).Error

	//err = errors2.Errorf("用户余额不足, uid: %d, money: %d", 1, 10) //在应用程序中出现错误时，使用 errors.New 或者 errors.Errorf 返回错误
	//if err != nil {
	//	return nil, errors2.WithMessage(err, "11111")
	//}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 报错
			//log2.Fatal(fmt.Sprintf("%+v\n", err))
			//
			//return nil, fmt.Errorf("%w, data is nil", err)
			return nil, nil
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

func (ar *userRepo) GetUserOrFail(ctx context.Context, id int64) (*biz.User, error) {
	var user User
	err := ar.data.db.WithContext(ctx).Where("id=?", 0).First(&user, id).Error

	//err = errors2.Errorf("用户余额不足, uid: %d, money: %d", 1, 10) //在应用程序中出现错误时，使用 errors.New 或者 errors.Errorf 返回错误
	//if err != nil {
	//	return nil, errors2.WithMessage(err, "11111")
	//}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 报错
			//log2.Fatal(fmt.Sprintf("%+v\n", err))
			//
			//return nil, fmt.Errorf("%w, data is nil", err)
			//return nil, RecordNotFoundError
			return nil, errors.New(400, v1.ErrorReason_ERROR_REASON_UNSPECIFIED.String(), "没有找到数据")
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
