package rpclogic

import (
	"context"

	"goChat/app/msg-gateway/cmd/wsrpc/internal/rpcsvc"
	"goChat/app/msg-gateway/cmd/wsrpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersOnlineStatusLogic struct {
	ctx    context.Context
	svcCtx *rpcsvc.ServiceContext
	logx.Logger
}

func NewGetUsersOnlineStatusLogic(ctx context.Context, svcCtx *rpcsvc.ServiceContext) *GetUsersOnlineStatusLogic {
	return &GetUsersOnlineStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersOnlineStatusLogic) GetUsersOnlineStatus(in *pb.GetUsersOnlineStatusReq) (*pb.GetUsersOnlineStatusResp, error) {
	// todo: add your rpclogic here and delete this line

	return &pb.GetUsersOnlineStatusResp{}, nil
}
