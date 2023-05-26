package rpc

import (
	"context"
	"goChat/app/usercenter/cmd/rpc/pb"
	"testing"
)

func TestVerifyTokenLogic(t *testing.T) {
	resp, err := imUserService.VerifyToken(context.Background(), &pb.VerifyTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiIxIiwiUGxhdGZvcm0iOiJ3ZWIiLCJleHAiOjQ4Mzg2ODY4NTcsIm5iZiI6MTY4NTA4Njg1NywiaWF0IjoxNjg1MDg2ODU3fQ.cGdZRpO1bmbaQAJm9zmraTVQmMjiNpuMa6bM4E441QA",
	})

	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", resp)
}
