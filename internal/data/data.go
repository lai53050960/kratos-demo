package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/propagation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos-client/api/user"
	"kratos-client/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo, NewUserServiceClient)

// Data .
type Data struct {
	db *gorm.DB
	uc user.UserClient //测试远程grpc 获取user 实例
}

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger, uc user.UserClient) (*Data, func(), error) {
	newHelper := log.NewHelper(logger)

	client, err := gorm.Open(
		mysql.Open(conf.Database.Source),
		&gorm.Config{},
	)
	if err != nil {
		newHelper.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}

	if err := client.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	d := &Data{
		db: client,
		uc: uc,
	}
	return d, func() {
		newHelper.Info("message", "closing the data resources")
	}, nil
}

func NewUserServiceClient(tp *tracesdk.TracerProvider) user.UserClient {
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint("localhost:9502"),
		grpc.WithMiddleware(middleware.Chain(
			tracing.Client(
				tracing.WithTracerProvider(tp),
				tracing.WithPropagators(
					propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
				),
			),
			recovery.Recovery())),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	c := user.NewUserClient(conn)
	return c
}
