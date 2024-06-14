package service

import (
	"GoBao/common/consts"
	"GoBao/common/goredis"
	gorm2 "GoBao/common/gorm"
	"GoBao/common/xerr"
	"GoBao/server/order/kmq/internal/config"
	"GoBao/server/order/model"
	"GoBao/server/order/rpc/order"
	"GoBao/server/product/rpc/productrpc"
	"context"
	"encoding/json"
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type Service struct {
	conf          config.Config
	OrderDB       *gorm.DB
	ProductDB     *gorm.DB
	RedisDB       *redis.Client
	SnowflakeNode *snowflake.Node
	ProductRpc    productrpc.ProductRpc
	OrderRpc      order.Order
}

func NewService(c config.Config) *Service {
	node, _ := snowflake.NewNode(consts.OrderSnowflakeNodeID)
	s := &Service{
		conf:          c,
		OrderDB:       gorm2.OrderDB,
		ProductDB:     gorm2.ProductDB,
		RedisDB:       goredis.Rdb,
		SnowflakeNode: node,
		ProductRpc:    productrpc.NewProductRpc(zrpc.MustNewClient(c.ProductRpc)),
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
	return s
}

func (s *Service) Consume(_, value string) error {
	var order model.Order
	err := json.Unmarshal([]byte(value), &order)
	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg("JSON UNMARSHAL order ERROR"), "JSON UNMARSHAL order ERROR:%+v", err)
	}
	err = s.OrderDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&order).Error
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL CREATE order ERROR:%+v", err)
		}
		_, err = s.ProductRpc.DeductSeckillStock(context.Background(), &productrpc.DeductSeckillStockReq{
			SeckillProductID: order.ProductId,
			Stock:            order.Quantity,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("ProductRPC ERROR"), "ProductRPC ERROR:%+v", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
