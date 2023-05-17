package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"goChat/usercenter/cmd/rpc/internal/svc"
	"goChat/usercenter/cmd/rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByIdLogic {
	return &LoginByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByIdLogic) LoginById(in *pb.LoginByIdReq) (*pb.LoginByIdResp, error) {
	token, err := l.getJwtToken(l.svcCtx.Config.JWTAuth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JWTAuth.AccessExpire, in.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.LoginByIdResp{
		User:  nil,
		Token: token,
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
}

func (l *LoginByIdLogic) getJwtToken(secretKey string, iat, seconds int64, userId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
