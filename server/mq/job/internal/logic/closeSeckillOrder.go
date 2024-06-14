package logic

import (
	"GoBao/common/globalKey"
	"GoBao/common/xerr"
	"GoBao/server/mq/job/internal/svc"
	"GoBao/server/mq/job/jobtype"
	"GoBao/server/order/rpc/order"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

var ErrCloseOrderFail = xerr.NewErrMsg("CLOSE ORDER FAIL")

type CloseSeckillOrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewCloseSeckillOrderHandler(svcCtx *svc.ServiceContext) *CloseSeckillOrderHandler {
	return &CloseSeckillOrderHandler{svcCtx: svcCtx}
}

func (l *CloseSeckillOrderHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var payload jobtype.DeferCloseSeckillOrderPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return errors.Wrapf(ErrCloseOrderFail, "CloseSeckillOrderHandler Fail, payload: %+v, ERROR: %+v", t.Payload(), err)
	}

	orderResp, err := l.svcCtx.OrderRpc.GetOrderDetail(ctx, &order.GetOrderDetailReq{
		UserID:  payload.UserID,
		OrderSn: payload.OrderSn,
	})
	if err != nil || orderResp == nil {
		return errors.Wrapf(ErrCloseOrderFail, "CloseSeckillOrderHandler GET order Fail OR order no exists, OrderSn: %v, ERROR: %+v", payload.OrderSn, err)
	}

	if orderResp.Status == globalKey.OrderWaitPay {
		_, err := l.svcCtx.OrderRpc.UpdateOrderStatus(ctx, &order.UpdateOrderStatusReq{
			OrderSn: payload.OrderSn,
			UserID:  payload.UserID,
			Status:  globalKey.OrderCancel,
		})
		if err != nil {
			return errors.Wrapf(ErrCloseOrderFail, "CloseSeckillOrderHandler CLOSE order Fail, OrderSn: %v, ERROR: %+v", payload.OrderSn, err)
		}
		// 用户并没有支付秒杀商品订单，需要在redis中恢复预库存
		err = l.svcCtx.RedisDB.HIncrBy(context.Background(), payload.RedisKey, payload.RedisField, payload.SeckillCount).Err()
		if err != nil {
			return errors.Wrapf(ErrCloseOrderFail, "CloseSeckillOrderHandler REDIS HIncrBy ERROR: %+v, key: %v, field: %v",
				err, payload.RedisKey, payload.RedisField)
		}
	}

	return nil
}
