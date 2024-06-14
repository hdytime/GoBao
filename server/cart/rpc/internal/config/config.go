package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	MysqlConf struct {
		DSN string
	}

	RedisConf struct {
		Addr     string
		DB       int
		PoolSize int
	}

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	ProductRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
}
