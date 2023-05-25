package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"goChat/app/usercenter/cmd/rpc/internal/repository"
	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/cmd/rpc/pb"
	jwtUtils "goChat/common/utils/jwt"
	"time"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *VerifyTokenLogic) VerifyToken(in *pb.VerifyTokenReq) (*pb.VerifyTokenResp, error) {
	// 1. 从令牌中获取声明（claim）信息
	claim, err := jwtUtils.GetClaimFromToken(in.Token, l.svcCtx.Config.TokenSecret)
	if err != nil {
		return nil, err
	}
	// 2. 使用l.rep.GetTokenMap函数从Redis中获取令牌映射信息
	tokenMap, err := l.rep.GetTokenMap(l.ctx, claim.UID, claim.Platform)
	if err != nil {
		return nil, err
	}
	// 3. 如果tokenMap中存在输入的令牌in.Token
	if ex, ok := tokenMap[in.Token]; ok {
		// 3.1. 检查令牌是否过期,如果过期就删除
		if time.Now().UnixMilli() > ex {
			e := l.rep.DeleteToken(l.ctx, claim.UID, claim.Platform, in.Token)
			if e != nil {
				l.Errorf("delete token error: %s", e.Error())
			}
			return &pb.VerifyTokenResp{
				Uid:     "",
				Success: false,
				ErrMsg:  "token expired",
			}, nil
		}
		// 3.2. 如果令牌未过期,则启动一个goroutine，更新令牌的过期时间
		go func() {
			_ = l.rep.RenewalToken(l.ctx, claim.UID, claim.Platform, in.Token)
		}()
		return &pb.VerifyTokenResp{
			Uid:     claim.UID,
			Success: true,
			ErrMsg:  "",
		}, nil
	} else {
		// 如果不存在
		return &pb.VerifyTokenResp{
			Uid:     "",
			Success: false,
			ErrMsg:  "token is not exist",
		}, nil
	}
}
