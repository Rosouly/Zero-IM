package rpclogic

import (
	"context"

	"goChat/app/msg-gateway/cmd/wsrpc/internal/rpcsvc"
	"goChat/app/msg-gateway/cmd/wsrpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type KickUserConnsLogic struct {
	ctx    context.Context
	svcCtx *rpcsvc.ServiceContext
	logx.Logger
}

func NewKickUserConnsLogic(ctx context.Context, svcCtx *rpcsvc.ServiceContext) *KickUserConnsLogic {
	return &KickUserConnsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KickUserConnsLogic) KickUserConns(in *pb.KickUserConnsReq) (*pb.KickUserConnsResp, error) {
	// todo: add your rpclogic here and delete this line

	return &pb.KickUserConnsResp{}, nil
}
