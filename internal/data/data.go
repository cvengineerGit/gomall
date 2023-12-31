package data

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/slog"
	"github.com/kalougata/gomall/pkg/config"
	"xorm.io/xorm"
)

type Data struct {
	DB *xorm.Engine
}

func NewData(conf *config.Config) (*Data, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
		conf.DB.User,
		conf.DB.Passwd,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.DbName,
	)
	db, err := xorm.NewEngine(conf.DB.Driver, dsn)

	if err != nil {
		return nil, nil, err
	}

	if err := db.PingContext(context.Background()); err != nil {
		return nil, nil, err
	}

	slog.Info("成功连接到数据库!")

	data := &Data{
		DB: db,
	}

	return data, func() {
		if err := db.Close(); err != nil {
			slog.Warnf("falied to close database: %s", err)
		}
	}, nil
}
