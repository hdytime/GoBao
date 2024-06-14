package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	Kafka      kq.KqConf
	OrderRpc   zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
}
