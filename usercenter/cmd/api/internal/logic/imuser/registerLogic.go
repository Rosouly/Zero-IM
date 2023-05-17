package imuser

import (
	"context"
	"github.com/pkg/errors"
	"goChat/common/helper"
	"goChat/common/xerr"
	"goChat/usercenter/cmd/rpc/pb"
	"goChat/usercenter/cmd/rpc/usercenterservice"

	"goChat/usercenter/cmd/api/internal/svc"
	"goChat/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserAlreadyRegisterError = xerr.NewErrMsg("user has been registered")

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 查找名字，看是否已经存在该用户
	result, err := l.svcCtx.UserCenterRpc.GeTUserByUsername(l.ctx, &usercenterservice.GeTUserByUsernameReq{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	if result != nil && result.User != nil && result.User.Id != "" {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "username: %s, err: %v", req.Username, err)
	}

	uuid := helper.GetID()
	insertUserResp, err := l.svcCtx.UserCenterRpc.InsertUser(l.ctx, &pb.InsertUserReq{
		User: &pb.User{
			Id:           uuid,
			Username:     req.Username,
			Password:     helper.Md5(req.Password),
			Nickname:     "Zero用户",
			Avatar:       "",
			Sign:         "",
			Province:     "",
			City:         "",
			District:     "",
			Birthday:     "",
			RegisterTime: "",
			IsMale:       true,
		},
	})
	if err != nil {
		return nil, err
	}
	if insertUserResp.BaseResp.ErrCode != 0 {
		return nil, err
	}
	// LoginById 调用登录接口，返回token
	loginResp, err := l.svcCtx.UserCenterRpc.LoginById(l.ctx, &pb.LoginByIdReq{
		UserId: uuid,
	})
	if err != nil {
		return nil, err
	}
	return &types.RegisterResp{
		Token: loginResp.Token,
	}, nil
}
