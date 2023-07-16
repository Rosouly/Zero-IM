package rpc

import (
	"context"
	chatpb "goChat/app/msg/cmd/rpc/pb"
	"testing"
)

func TestSendMsg(t *testing.T) {
	chatMsg := &chatpb.MsgData{
		SendID:           "9be675b41ff6feb3c21e528b92c812c1", // 用户01
		RecvID:           "2293575ce049296775c19640a61eab77", // 用户02
		GroupID:          "test_group",
		ClientMsgID:      "test_client_msg_id",
		ServerMsgID:      "test_server_msg_id",
		SenderPlatformID: 1,
		SenderNickname:   "test_nickname",
		SenderFaceURL:    "https://example.com/avatar.jpg",
		SessionType:      1,
		MsgFrom:          1,
		ContentType:      101, // 101-text
		Content:          []byte("Hello, world!"),
		Seq:              1,
		ServerTime:       1625500000,
		ClientTime:       1625500010,
		OfflinePushInfo: &chatpb.OfflinePushInfo{
			Title: "Test title",
		},
		AtUserIDList: []string{"user1", "user2"},
		Options:      map[string]bool{"option1": true, "option2": false},
	}

	Data := &chatpb.SendMsgReq{
		Token:   "tokentest",
		MsgData: chatMsg,
	}
	msg, err := msgService.SendMsg(context.Background(), Data)
	if err != nil {
		t.Error("msgService.SendMsg err:", err)
		return
	}

	t.Log("msgService.SendMsg msg:", msg)
}
