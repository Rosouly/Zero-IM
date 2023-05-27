package logic

import (
	"context"

	"goChat/app/msg/cmd/rpc/internal/svc"
	"goChat/app/msg/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PullMessageBySuperGroupSeqListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPullMessageBySuperGroupSeqListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullMessageBySuperGroupSeqListLogic {
	return &PullMessageBySuperGroupSeqListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PullMessageBySuperGroupSeqListLogic) PullMessageBySuperGroupSeqList(in *pb.PullMessageBySuperGroupSeqListReq) (*pb.WrapPullMessageBySeqListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.WrapPullMessageBySeqListResp{}, nil
}
