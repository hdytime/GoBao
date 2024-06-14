package handler

import (
	"GoBao/common/response"
	"net/http"

	"GoBao/server/user/api/internal/logic"
	"GoBao/server/user/api/internal/svc"
	"GoBao/server/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
