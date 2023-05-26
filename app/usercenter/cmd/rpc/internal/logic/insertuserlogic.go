package logic

import (
	"context"
	"goChat/app/usercenter/cmd/rpc/internal/repository"
	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/cmd/rpc/pb"
	"goChat/app/usercenter/model"
	"goChat/common/xorm"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewInsertUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertUserLogic {
	return &InsertUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *InsertUserLogic) InsertUser(in *pb.InsertUserReq) (*pb.InsertUserResp, error) {
	// todo:
	user := &model.User{
		Id:           in.User.Id,
		Username:     in.User.Username,
		Password:     in.User.Password,
		Nickname:     in.User.Nickname,
		Sign:         in.User.Sign,
		Avatar:       in.User.Avatar,
		Province:     in.User.Province,
		City:         in.User.City,
		District:     in.User.District,
		Birthday:     in.User.Birthday,
		RegisterTime: in.User.RegisterTime,
		IsMale:       in.User.IsMale,
	}
	err := xorm.Transaction(
		l.rep.Mysql,
		func(tx *gorm.DB) error {
			return xorm.Insert(tx, user)
		},
	)
	if err != nil {
		return &pb.InsertUserResp{
			BaseResp: &pb.BaseResp{
				ErrCode: -1,
				ErrMsg:  err.Error(),
			},
		}, err
	}
	return &pb.InsertUserResp{
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
	// 预热数据
	//go l.rep.WarmUpUser()
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
