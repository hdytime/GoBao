package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	ProductRpcConf zrpc.RpcClientConf
	UserRpcConf    zrpc.RpcClientConf
	OrderRpcConf   zrpc.RpcClientConf
	PayRpcConf     zrpc.RpcClientConf
}
