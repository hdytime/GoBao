package storeProduct

import (
	"GoBao/common/response"
	"net/http"

	"GoBao/server/product/api/internal/logic/storeProduct"
	"GoBao/server/product/api/internal/svc"
	"GoBao/server/product/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := storeProduct.NewDeleteProductLogic(r.Context(), svcCtx)
		err := l.DeleteProduct(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
