package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goChat/app/test/api/internal/logic"
	"goChat/app/test/api/internal/svc"
	"goChat/app/test/api/internal/types"
)

func isUsernameExistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewIsUsernameExistLogic(r.Context(), svcCtx)
		resp, err := l.IsUsernameExist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
