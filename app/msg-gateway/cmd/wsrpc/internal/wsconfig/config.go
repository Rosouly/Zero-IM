package wsconfig

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Redis            redis.RedisConf
	SendMsgRateLimit RateLimitConfig
	Websocket        WebsocketConfig
	ImUserRpc        zrpc.RpcClientConf
	MsgRpc           zrpc.RpcClientConf
}

type RateLimitConfig struct {
	Enable  bool
	Seconds int
	Quota   int
}

type WebsocketConfig struct {
	MaxConnNum int
	TimeOut    int
	MaxMsgLen  int
}
