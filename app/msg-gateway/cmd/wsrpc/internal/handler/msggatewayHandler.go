package handler

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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

		ws := wslogic.NewMsggatewayLogic(context.Background(), svcCtx)
		resp, ok := ws.Msggateway(&req)
		status := http.StatusUnauthorized
		if ok {
			err := ws.WsUpgrade(resp.Uid, &req, w, r, nil)
			if err != nil {
				logx.WithContext(r.Context()).Errorf("ws.WsUpgrade error: %s", err)
				return
			}
		} else {
			w.Header().Set("Sec-Websocket-Version", "13")
			w.Header().Set("ws_err_msg", "args err, need token, sendID, platformID")
			http.Error(w, http.StatusText(status), status)
		}
	}
}
