package rpc

import (
	"context"
	"goChat/app/usercenter/cmd/rpc/pb"
	"testing"
	"time"
)

func TestInsertUserLogic(t *testing.T) {
	//resp, err := usercenterService.InsertUser(
	//	context.Background(),
	//	&pb.InsertUserReq{
	//		User: &pb.User{
	//			Username:     "user01",
	//			Password:     "123456",
	//			Nickname:     "第user01个注册的人",
	//			Sign:         "第user01个注册的人，吃瓜",
	//			Avatar:       "https://go.dev/images/gophers/pilot-bust.svg",
	//			Province:     "广州市",
	//			City:         "广州市",
	//			District:     "越秀区",
	//			Birthday:     "1999-03-06",
	//			RegisterTime: time.Now().Format("2006-01-02 15:04:05"),
	//			IsMale:       true,
	//		},
	//	},
	//)
	resp, err := usercenterService.InsertUser(
		context.Background(),
		&pb.InsertUserReq{
			User: &pb.User{
				Username:     "user02",
				Password:     "123456",
				Nickname:     "第user02个注册的人",
				Sign:         "第user02个注册的人，吃瓜",
				Avatar:       "https://go.dev/images/gophers/pilot-bust.svg",
				Province:     "广州市",
				City:         "广州市",
				District:     "越秀区",
				Birthday:     "1999-03-07",
				RegisterTime: time.Now().Format("2006-01-02 15:04:05"),
				IsMale:       true,
			},
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
