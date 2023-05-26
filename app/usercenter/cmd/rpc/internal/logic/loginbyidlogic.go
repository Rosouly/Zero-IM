package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	onlinemessagerelayservice "goChat/app/msg-gateway/cmd/wsrpc/onlineMessageRelayService"
	gatewaypb "goChat/app/msg-gateway/cmd/wsrpc/pb"
	"goChat/app/usercenter/cmd/rpc/internal/repository"
	"goChat/app/usercenter/cmd/rpc/internal/svc"
	"goChat/app/usercenter/cmd/rpc/pb"
	jwtUtils "goChat/common/utils/jwt"
)

type LoginByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewLoginByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByIdLogic {
	return &LoginByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *LoginByIdLogic) LoginById(in *pb.LoginByIdReq) (*pb.LoginByIdResp, error) {
	// todo:
	claim := jwtUtils.BuildClaims(in.UserId, "web")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString([]byte(l.svcCtx.Config.TokenSecret))
	if err != nil {
		l.Errorf("jwt signed string error: %v", err)
		return nil, err
	}
	tokenMap, err := l.rep.GetTokenMap(l.ctx, claim.UID, claim.Platform)
	if err != nil {
		return nil, err
	}
	if len(tokenMap) > 0 {
		err = l.rep.DelTokenMap(l.ctx, claim.UID, claim.Platform)
		if err != nil {
			l.Errorf("del token map error: %v", err)
			return nil, err
		}
		// 断开用户之前的链接
		services, err := l.getAllMsgGatewayService()
		if err != nil {
			return nil, err
		}
		for _, service := range services {
			_, err := service.KickUserConns(l.ctx, &gatewaypb.KickUserConnsReq{
				UserID:      claim.UID,
				PlatformIDs: []string{claim.Platform},
			})
			if err != nil {
				l.Errorf("kick user conns error: %v", err)
			}
		}
	}
	// 写入tokenMap
	err = l.rep.SetTokenMap(l.ctx, claim.UID, claim.Platform, signedString)
	if err != nil {
		l.Errorf("set tokenMap error: %v", err)
		return nil, err
	}
	return &pb.LoginByIdResp{
		User:  nil,
		Token: signedString,
		BaseResp: &pb.BaseResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, err
}

func (l *LoginByIdLogic) getAllMsgGatewayService() (services []onlinemessagerelayservice.OnlineMessageRelayService, err error) {
	return onlinemessagerelayservice.GetAllByEtcd(l.ctx, l.svcCtx.Config.MsgGatewayRpc, l.svcCtx.Config.MsgGatewayRpc.Key)
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
