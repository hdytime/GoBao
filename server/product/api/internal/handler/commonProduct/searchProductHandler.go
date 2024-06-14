package commonProduct

import (
	"GoBao/common/response"
	"net/http"

	"GoBao/server/product/api/internal/logic/commonProduct"
	"GoBao/server/product/api/internal/svc"
	"GoBao/server/product/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := commonProduct.NewSearchProductLogic(r.Context(), svcCtx)
		resp, err := l.SearchProduct(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
