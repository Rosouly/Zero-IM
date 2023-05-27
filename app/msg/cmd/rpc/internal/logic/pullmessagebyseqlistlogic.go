package logic

import (
	"context"

	"goChat/app/msg/cmd/rpc/internal/svc"
	"goChat/app/msg/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PullMessageBySeqListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPullMessageBySeqListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullMessageBySeqListLogic {
	return &PullMessageBySeqListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PullMessageBySeqListLogic) PullMessageBySeqList(in *pb.WrapPullMessageBySeqListReq) (*pb.WrapPullMessageBySeqListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.WrapPullMessageBySeqListResp{}, nil
}
