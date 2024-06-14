package pay

import (
	"GoBao/common/response"
	"net/http"

	"GoBao/server/pay/api/internal/logic/pay"
	"GoBao/server/pay/api/internal/svc"
	"GoBao/server/pay/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderPayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderPayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := pay.NewOrderPayLogic(r.Context(), svcCtx)
		resp, err := l.OrderPay(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
