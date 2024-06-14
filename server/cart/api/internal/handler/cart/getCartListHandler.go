package cart

import (
	"GoBao/common/response"
	"net/http"

	"GoBao/server/cart/api/internal/logic/cart"
	"GoBao/server/cart/api/internal/svc"
)

func GetCartListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cart.NewGetCartListLogic(r.Context(), svcCtx)
		resp, err := l.GetCartList()
		response.HttpResponse(r, w, resp, err)
	}
}
