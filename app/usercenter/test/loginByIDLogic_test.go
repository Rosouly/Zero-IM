package rpc

import (
	"context"
	"goChat/app/usercenter/cmd/rpc/pb"
	"testing"
)

func TestLoginByIDLogic(t *testing.T) {
	resp, err := usercenterService.LoginById(context.Background(),
		&pb.LoginByIdReq{
			UserId: "9be675b41ff6feb3c21e528b92c812c1",
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
