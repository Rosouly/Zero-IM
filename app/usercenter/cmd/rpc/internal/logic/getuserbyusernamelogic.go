package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/cmd/rpc/pb"
)

type GeTUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGeTUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeTUserByUsernameLogic {
	return &GeTUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GeTUserByUsernameLogic) GeTUserByUsername(in *pb.GeTUserByUsernameReq) (*pb.GeTUserResp, error) {
	// todo: getUserByUsername-logic
	//pbUser, err := l.svcCtx.UsercenterModel.FindOneByUsername(l.ctx, in.Username)
	//// 如果出错了，而且还不是没找到结果那种错误
	//if err != nil && err != model.ErrNotFound {
	//	return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "username: %s, err:%v", in.Username, err)
	//}
	//user := &pb.User{}
	//_ = copier.Copy(&user, &pbUser)
	//return &pb.GeTUserResp{
	//	User: user,
	//}, nil
	return nil, nil
}
