package rpc

import (
	"context"
	"goChat/app/usercenter/cmd/rpc/pb"
	"testing"
)

func TestVerifyTokenLogic(t *testing.T) {
	resp, err := imUserService.VerifyToken(context.Background(), &pb.VerifyTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI5YmU2NzViNDFmZjZmZWIzYzIxZTUyOGI5MmM4MTJjMSIsIlBsYXRmb3JtIjoid2ViIiwiZXhwIjo0ODQxOTAyOTExLCJuYmYiOjE2ODgzMDI5MTEsImlhdCI6MTY4ODMwMjkxMX0.weTvvYtc2baLVqLGpSTDD-rLl0PwEw6u-vGwhGKliMo",
	})

	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", resp)
}
