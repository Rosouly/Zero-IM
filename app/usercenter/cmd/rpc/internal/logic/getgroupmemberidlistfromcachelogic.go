package logic

import (
	"context"

	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMemberIDListFromCacheLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMemberIDListFromCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberIDListFromCacheLogic {
	return &GetGroupMemberIDListFromCacheLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取群组成员列表
func (l *GetGroupMemberIDListFromCacheLogic) GetGroupMemberIDListFromCache(in *pb.GetGroupMemberIDListFromCacheReq) (*pb.GetGroupMemberIDListFromCacheResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetGroupMemberIDListFromCacheResp{}, nil
}
