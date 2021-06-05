package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos-client/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo)

// Data .
type Data struct {
	db *gorm.DB
}

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
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
	}
	return d, func() {
		newHelper.Info("message", "closing the data resources")
	}, nil
}
