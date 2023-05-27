package logic

import (
	"context"

	"goChat/app/msg/cmd/rpc/internal/svc"
	"goChat/app/msg/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSuperGroupMaxAndMinSeqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSuperGroupMaxAndMinSeqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSuperGroupMaxAndMinSeqLogic {
	return &GetSuperGroupMaxAndMinSeqLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSuperGroupMaxAndMinSeqLogic) GetSuperGroupMaxAndMinSeq(in *pb.GetMaxAndMinSuperGroupSeqReq) (*pb.GetMaxAndMinSuperGroupSeqResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMaxAndMinSuperGroupSeqResp{}, nil
}
