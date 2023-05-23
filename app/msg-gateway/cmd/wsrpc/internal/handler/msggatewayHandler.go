package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/types"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/wslogic"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/wssvc"
)

func msggatewayHandler(svcCtx *wssvc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := wslogic.NewMsggatewayLogic(r.Context(), svcCtx)
		resp, err := l.Msggateway(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
