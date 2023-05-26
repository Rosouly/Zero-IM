package rpc

import (
	"context"
	"goChat/app/usercenter/cmd/rpc/pb"
	"testing"
)

func TestGetUserByUsernameLogic(t *testing.T) {
	resp, err := usercenterService.GeTUserByUsername(
		context.Background(), &pb.GeTUserByUsernameReq{
			Username: "小明",
		},
	)
	if err != nil {
		t.Error("err", err)
	}
	t.Log("resp:", resp)
}
