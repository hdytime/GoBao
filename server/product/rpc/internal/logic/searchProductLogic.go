package logic

import (
	"GoBao/common/xerr"
	"GoBao/server/product/model"
	"context"
	"fmt"
	"github.com/pkg/errors"

	"GoBao/server/product/rpc/internal/svc"
	"GoBao/server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductLogic {
	return &SearchProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchProductLogic) SearchProduct(in *pb.SearchProductReq) (*pb.SearchProductResp, error) {
	// todo: add your logic here and delete this line

	offset, limit := int((in.Page-1)*in.Size), int(in.Size)
	db := l.svcCtx.ProductDB.Offset(offset).Limit(limit).Order(in.Sort)
	key := fmt.Sprintf("%%%s%%", in.Keyword)
	db.Where("name=?", key)

	var list []model.Product
	err := db.Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product search ERROR:%+v", err)
	}
	var res []*pb.SmallProduct
	for _, p := range list {
		sp := &pb.SmallProduct{
			ID:            p.Id,
			Name:          p.Name,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
		}
		res = append(res, sp)
	}
	return &pb.SearchProductResp{SmallProducts: res}, nil
}
