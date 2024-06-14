package order

import (
	"GoBao/common/ctxData"
	"GoBao/server/order/rpc/order"
	"context"
	"strconv"
	"time"

	"GoBao/server/order/api/internal/svc"
	"GoBao/server/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDetailLogic) OrderDetail(req *types.GetOrderDetailReq) (resp *types.GetOrderDetailResp, err error) {
	// todo: add your logic here and delete this line

	uid := ctxData.GetUserIDFromCtx(l.ctx)
	detail, err := l.svcCtx.OrderRpc.GetOrderDetail(l.ctx, &order.GetOrderDetailReq{
		UserID:  uid,
		OrderSn: req.OrderSn,
	})
	if err != nil {
		return nil, err
	}
	var res = &types.GetOrderDetailResp{
		ID:           detail.ID,
		CreateTime:   time.Unix(detail.Createtime, 0).Format("2006-01-02 15:04:05"),
		UpdateTime:   time.Unix(detail.Updatetime, 0).Format("2006-01-02 15:04:05"),
		OrderSn:      detail.OrderSn,
		UserID:       detail.UserID,
		ProductID:    detail.ProductID,
		Name:         detail.ProductName,
		ProductCount: detail.Quantity,
		UnitPrice:    detail.UnitPrice,
		TotalPrice:   detail.TotalPrice,
		Status:       detail.Status,
		Remark:       detail.Remark,
		PayTime:      strconv.FormatInt(detail.Paytime, 10),
	}

	return res, nil
}
