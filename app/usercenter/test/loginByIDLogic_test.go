package rpc

import (
	"context"
	"goChat/app/usercenter/cmd/rpc/pb"
	"testing"
)

func TestLoginByIDLogic(t *testing.T) {
	resp, err := usercenterService.LoginById(context.Background(),
		&pb.LoginByIdReq{
			UserId: "1",
		},
	)
	if err != nil {
		t.Error("err", err)
	}
	t.Log(resp)
}
