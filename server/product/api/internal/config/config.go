package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ProductRpcConf zrpc.RpcClientConf
	UserRpcConf    zrpc.RpcClientConf
	JwtAuth        struct {
		AccessSecret string
		AccessExpire int64
	}
}
