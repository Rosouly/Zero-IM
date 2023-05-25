package logic

import (
	"context"
	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertUserLogic {
	return &InsertUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertUserLogic) InsertUser(in *pb.InsertUserReq) (*pb.InsertUserResp, error) {
	// todo
	//user := in.User
	//_, err := l.svcCtx.UsercenterModel.Insert(l.ctx, &model.User{
	//	Id:           user.Id,
	//	Username:     user.Username,
	//	Password:     user.Password,
	//	Nickname:     "Zero用户",
	//	Avatar:       "",
	//	Sign:         "",
	//	Province:     "",
	//	City:         "",
	//	District:     "",
	//	Birthday:     "",
	//	RegisterTime: "",
	//	IsMale:       user.IsMale,
	//})
	//if err != nil {
	//	return &pb.InsertUserResp{
	//		BaseResp: &pb.BaseResp{
	//			ErrCode: -1,
	//			ErrMsg:  err.Error(),
	//		},
	//	}, err
	//}
	//return &pb.InsertUserResp{
	//	BaseResp: &pb.BaseResp{
	//		ErrCode: 0,
	//		ErrMsg:  "success",
	//	},
	//}, nil
	return nil, nil
}
