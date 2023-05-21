package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"goChat/app/usercenter/cmd/rpc/internal/config"
	"goChat/app/usercenter/model"
)

type ServiceContext struct {
	Config          config.Config
	UsercenterModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		UsercenterModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
