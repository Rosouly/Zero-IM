package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"goChat/app/usercenter/cmd/rpc/internal/repository"
	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/cmd/rpc/pb"
	"goChat/app/usercenter/model"
	"goChat/common/xcache/global"
	xormerr "goChat/common/xorm/err"
)

type GeTUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGeTUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeTUserByUsernameLogic {
	return &GeTUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *GeTUserByUsernameLogic) GeTUserByUsername(in *pb.GeTUserByUsernameReq) (*pb.GeTUserResp, error) {
	fmt.Println("GeTUserByUsernameLogic.GeTUserByUsername")
	user := &model.User{}
	user.Username = in.Username
	err := l.rep.DetailCache.FirstByWhere(
		user,
		map[string][]interface{}{
			"username=?": {in.Username},
		},
	)
	if err != nil {
		if xormerr.RecordNotFound(err) {
			err = nil
		}
		if err == global.RedisErrorNotExists {
			err = nil
		}
		if err != nil {
			l.Errorf("get user by username error: %s", err.Error())
			return &pb.GeTUserResp{
				User: nil,
			}, err
		}
	}
	var userInfo *pb.User
	userInfo = &pb.User{
		Id:           user.Id,
		Username:     user.Username,
		Password:     user.Password,
		Nickname:     user.Nickname,
		Sign:         user.Sign,
		Avatar:       user.Avatar,
		Province:     user.Province,
		City:         user.City,
		District:     user.District,
		Birthday:     user.Birthday,
		RegisterTime: user.RegisterTime,
		IsMale:       user.IsMale,
	}
	return &pb.GeTUserResp{
		User: userInfo,
	}, err
}
