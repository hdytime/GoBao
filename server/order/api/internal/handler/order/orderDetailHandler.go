package order

import (
	"GoBao/common/response"
	"net/http"

	"GoBao/server/order/api/internal/logic/order"
	"GoBao/server/order/api/internal/svc"
	"GoBao/server/order/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrderDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.OrderDetail(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
