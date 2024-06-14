package order

import (
	"GoBao/common/response"
	"net/http"

	"GoBao/server/order/api/internal/logic/order"
	"GoBao/server/order/api/internal/svc"
	"GoBao/server/order/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewDeleteOrderLogic(r.Context(), svcCtx)
		err := l.DeleteOrder(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
