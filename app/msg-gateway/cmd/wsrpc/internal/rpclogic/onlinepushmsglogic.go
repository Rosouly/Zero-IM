package rpclogic

import (
	"context"

	"goChat/app/msg-gateway/cmd/wsrpc/internal/rpcsvc"
	"goChat/app/msg-gateway/cmd/wsrpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnlinePushMsgLogic struct {
	ctx    context.Context
	svcCtx *rpcsvc.ServiceContext
	logx.Logger
}

func NewOnlinePushMsgLogic(ctx context.Context, svcCtx *rpcsvc.ServiceContext) *OnlinePushMsgLogic {
	return &OnlinePushMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OnlinePushMsgLogic) OnlinePushMsg(in *pb.OnlinePushMsgReq) (*pb.OnlinePushMsgResp, error) {
	// todo: add your rpclogic here and delete this line

	return &pb.OnlinePushMsgResp{}, nil
}
