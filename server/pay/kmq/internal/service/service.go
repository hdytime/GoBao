package service

import (
	"GoBao/common/consts"
	"GoBao/common/globalKey"
	"GoBao/common/goredis"
	gorm2 "GoBao/common/gorm"
	"GoBao/common/kqOrder"
	"GoBao/common/xerr"
	"GoBao/server/order/rpc/order"
	"GoBao/server/pay/kmq/internal/config"
	"GoBao/server/product/rpc/productrpc"
	"context"
	"encoding/json"
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
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

func (s *Service) Consume(_, val string) error {
	// 获取队列信息
	var kqMessage kqOrder.PaymentUpdateOrderState
	if err := json.Unmarshal([]byte(val), &kqMessage); err != nil {
		logx.WithContext(context.Background()).Error("PaymentUpdateOrderStateMq`Consume Unmarshal kqOrder.PaymentUpdateOrderState ERROR: %+v", err)
		return err
	}

	// 判断订单是否存在
	orderResp, err := s.OrderRpc.GetOrderDetail(context.Background(), &order.GetOrderDetailReq{
		UserID:  kqMessage.UserID,
		OrderSn: kqMessage.OrderSn,
	})
	if err != nil {
		logx.WithContext(context.Background()).Error("PaymentUpdateOrderStateMq`Consume USE OrderRpc.GetOrderOnlyDetail ERROR: %+v, userID: %v, orderSn: %v",
			err, kqMessage.UserID, kqMessage.OrderSn)
		return err
	}

	// 判断订单状态更新是否正确
	if orderResp.Status != globalKey.OrderWaitPay {
		logx.WithContext(context.Background()).Error("Order`State ERROR, userID: %v, orderSn: %v, oldState: %v, newState: %v",
			kqMessage.UserID, kqMessage.OrderSn, orderResp.Status, kqMessage.OrderState)
		return xerr.NewErrMsg("Order`State ERROR")
	}

	// 更新订单状态
	_, err = s.OrderRpc.UpdateOrderStatus(context.Background(), &order.UpdateOrderStatusReq{
		OrderSn: kqMessage.OrderSn,
		UserID:  kqMessage.UserID,
		Status:  kqMessage.OrderState,
	})
	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg("kafka use OrderRpc ERROR"), "kafka use OrderRpc ERROR:%+v")
	}
	return nil
}
