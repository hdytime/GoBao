package cart

import (
	"GoBao/common/response"
	"net/http"

	"GoBao/server/cart/api/internal/logic/cart"
	"GoBao/server/cart/api/internal/svc"
	"GoBao/server/cart/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateCartProductDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCartProductDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cart.NewUpdateCartProductDetailLogic(r.Context(), svcCtx)
		err := l.UpdateCartProductDetail(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
