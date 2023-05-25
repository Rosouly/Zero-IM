package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"goChat/common/xorm/global"
)

type Config struct {
	zrpc.RpcServerConf
	//Mysql struct {
	//	DataSource string
	//}
	//CacheRedis cache.CacheConf
	//JWTAuth struct {
	//	AccessSecret string
	//	AccessExpire int64
	//}
	TokenRateLimiter RateLimitConfig
	RedisConfig      RedisConfig
	MysqlConfig      global.MysqlConfig
	TokenRenewalDay  int64  `json:",default=30"` // 用户每次连接websocket 自动续签的的天数 默认30天
	TokenSecret      string `json:",default=zeroimserver"`
}

type RedisConfig struct {
	redis.RedisConf
	DB int
}

type RateLimitConfig struct {
	Seconds int
	Quota   int
}
