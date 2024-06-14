package handler

import (
	"GoBao/common/response"
	"net/http"

	"GoBao/server/user/api/internal/logic"
	"GoBao/server/user/api/internal/svc"
)

func detailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail()
		response.HttpResponse(r, w, resp, err)
	}
}
