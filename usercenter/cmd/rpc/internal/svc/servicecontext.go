package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"goChat/usercenter/cmd/rpc/internal/config"
	"goChat/usercenter/model"
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
