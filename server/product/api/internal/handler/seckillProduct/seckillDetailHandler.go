package seckillProduct

import (
	"GoBao/common/response"
	"net/http"

	"GoBao/server/product/api/internal/logic/seckillProduct"
	"GoBao/server/product/api/internal/svc"
	"GoBao/server/product/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SeckillDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSeckillDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := seckillProduct.NewSeckillDetailLogic(r.Context(), svcCtx)
		resp, err := l.SeckillDetail(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
