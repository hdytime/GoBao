package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	UserRpcConf    zrpc.RpcClientConf
	ProductRpcConf zrpc.RpcClientConf
	PayRpcConf     zrpc.RpcClientConf

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

	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
}
