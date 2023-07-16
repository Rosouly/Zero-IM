package logic

import (
	"context"

	"goChat/app/test/api/internal/svc"
	"goChat/app/test/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsUsernameExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsUsernameExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUsernameExistLogic {
	return &IsUsernameExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsUsernameExistLogic) IsUsernameExist(req *types.TestRequest) (resp *types.TestResponse, err error) {
	// todo: add your logic here and delete this line

	return &types.TestResponse{
		Exist: true,
	}, nil
}
