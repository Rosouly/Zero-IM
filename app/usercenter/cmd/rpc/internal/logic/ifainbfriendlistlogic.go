package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"goChat/app/usercenter/cmd/rpc/internal/repository"
	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/cmd/rpc/pb"
	"goChat/app/usercenter/model"
	"goChat/common/xcache/rc"
	xormerr "goChat/common/xorm/err"
)

type IfAInBFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewIfAInBFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfAInBFriendListLogic {
	return &IfAInBFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

// 判断用户A是否在B好友列表中
func (l *IfAInBFriendListLogic) IfAInBFriendList(in *pb.IfAInBFriendListReq) (*pb.IfAInBFriendListResp, error) {
	friendlist := &model.Friendlist{}
	friendlist.SelfId = in.BUserID
	fmt.Println("friendlist:", friendlist)
	exist, err := l.rep.RelationCache.Exist(
		in.AUserID,
		friendlist,
		"user_id",
		map[string]interface{}{},
		rc.Size(2))
	if err != nil {
		if xormerr.TableNotFound(err) {
			_ = l.rep.Mysql.Table(friendlist.TableName()).AutoMigrate(friendlist)
		}
		return nil, err
	}
	fmt.Println("exist:", exist)
	return &pb.IfAInBFriendListResp{
		CommonResp: &pb.CommonResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
		IsInFriendList: exist,
	}, nil
}
