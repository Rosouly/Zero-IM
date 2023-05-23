package wslogic

import (
	"context"

	"goChat/app/msg-gateway/cmd/wsrpc/internal/types"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/wssvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsggatewayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *wssvc.ServiceContext
}

func NewMsggatewayLogic(ctx context.Context, svcCtx *wssvc.ServiceContext) *MsggatewayLogic {
	return &MsggatewayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsggatewayLogic) Msggateway(req *types.Request) (resp *types.Response, err error) {
	// todo: add your wslogic here and delete this line

	return
}
