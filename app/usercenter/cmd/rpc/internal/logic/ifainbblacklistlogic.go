package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"goChat/app/usercenter/cmd/rpc/internal/repository"
	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/cmd/rpc/pb"
	"goChat/app/usercenter/model"
	"goChat/common/xcache/rc"
	xormerr "goChat/common/xorm/err"
)

type IfAInBBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewIfAInBBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfAInBBlacklistLogic {
	return &IfAInBBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

// 判断用户A是否在B黑名单中
func (l *IfAInBBlacklistLogic) IfAInBBlacklist(in *pb.IfAInBBlacklistReq) (*pb.IfAInBBlacklistResp, error) {
	blacklist := &model.Blacklist{}
	blacklist.SelfId = in.BUserID
	exist, err := l.rep.RelationCache.Exist(
		in.AUserID,
		blacklist,
		"user_id",
		map[string]interface{}{},
		rc.Size(2))
	if err != nil {
		if xormerr.TableNotFound(err) {
			_ = l.rep.Mysql.Table(blacklist.TableName()).AutoMigrate(blacklist)
		}
		return nil, err
	}
	return &pb.IfAInBBlacklistResp{
		CommonResp: &pb.CommonResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
		IsInBlacklist: exist,
	}, nil
}
