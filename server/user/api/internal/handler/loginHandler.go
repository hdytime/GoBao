package handler

import (
	"GoBao/common/response"
	"net/http"

	"GoBao/server/user/api/internal/logic"
	"GoBao/server/user/api/internal/svc"
	"GoBao/server/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
