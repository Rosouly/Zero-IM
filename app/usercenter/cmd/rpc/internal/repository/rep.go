package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/limit"
	zeroredis "github.com/zeromicro/go-zero/core/stores/redis"
	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/model"
	"goChat/common/xcache"
	"goChat/common/xcache/dc"
	"goChat/common/xcache/global"
	"goChat/common/xcache/rc"
	"goChat/common/xorm"
	"gorm.io/gorm"
)

type Rep struct {
	svcCtx        *svc.ServiceContext
	Cache         redis.UniversalClient
	Mysql         *gorm.DB
	DetailCache   *dc.DbMapping
	RelationCache *rc.RelationMapping
	RateLimiter   *limit.PeriodLimit
}

//func (r *Rep) CheckDefaultSuperGroup() {
//	_ = xorm.Insert(r.Mysql, model.DefaultGroup)
//}

var rep *Rep

func NewRep(svcCtx *svc.ServiceContext) *Rep {
	if rep != nil {
		return rep
	}
	rep = &Rep{
		svcCtx: svcCtx,
		Cache:  xcache.GetClient(svcCtx.Config.RedisConfig.RedisConf, global.DB(svcCtx.Config.RedisConfig.DB)),
		Mysql:  xorm.GetClient(svcCtx.Config.MysqlConfig),
	}
	err := rep.Mysql.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	rep.DetailCache = dc.GetDbMapping(rep.Cache, rep.Mysql)
	rep.RelationCache = rc.NewRelationMapping(rep.Mysql, rep.Cache)
	//rep.CheckDefaultSuperGroup()
	rep.RateLimiter = limit.NewPeriodLimit(
		svcCtx.Config.TokenRateLimiter.Seconds,
		svcCtx.Config.TokenRateLimiter.Quota,
		newRedis(svcCtx.Config.RedisConfig.Host, svcCtx.Config.RedisConfig.Pass, svcCtx.Config.RedisConfig.Type, false),
		"periodlimit:tokenrpc:",
		limit.Align(),
	)
	return rep
}

func newRedis(addr string, password string, typ string, tls bool) *zeroredis.Redis {
	ops := make([]zeroredis.Option, 0)
	if password != "" {
		ops = append(ops, zeroredis.WithPass(password))
	}
	if typ == "cluster" {
		ops = append(ops, zeroredis.Cluster())
	}
	if tls {
		ops = append(ops, zeroredis.WithTLS())
	}
	return zeroredis.New(addr, ops...)
}
