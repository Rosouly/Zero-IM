package logic

import (
	"context"

	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListFromSuperGroupWithOptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListFromSuperGroupWithOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListFromSuperGroupWithOptLogic {
	return &GetUserListFromSuperGroupWithOptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取超级群成员列表 通过消息接收选项
func (l *GetUserListFromSuperGroupWithOptLogic) GetUserListFromSuperGroupWithOpt(in *pb.GetUserListFromSuperGroupWithOptReq) (*pb.GetUserListFromSuperGroupWithOptResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserListFromSuperGroupWithOptResp{}, nil
}
