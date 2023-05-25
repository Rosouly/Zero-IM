package svc

import (
	"goChat/app/usercenter/cmd/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//UsercenterModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		//UsercenterModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
