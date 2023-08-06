package test

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"goChat/app/msg-gateway/cmd/wsrpc/pb"
	chatpb "goChat/app/msg/cmd/rpc/pb"
	"testing"
)

func TestWebsocketRequest(t *testing.T) {
	sendId := "9be675b41ff6feb3c21e528b92c812c1"
	platformId := "6"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI5YmU2NzViNDFmZjZmZWIzYzIxZTUyOGI5MmM4MTJjMSIsIlBsYXRmb3JtIjoid2ViIiwiZXhwIjo0ODQzNTA2NzA0LCJuYmYiOjE2ODk5MDY3MDQsImlhdCI6MTY4OTkwNjcwNH0.G8ch_Op8NFI_ng7ZcXy1HdAl9ThDOZKSOKlDVcVQaXo"
	dialer := websocket.DefaultDialer
	// 添加自定义 header，如果需要的话
	url := fmt.Sprintf("ws://localhost:8084/?token=%s&sendId=%s&platformId=%s", token, sendId, platformId)
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		fmt.Println("connect err:", err)
		return
	}
	fmt.Println("connect success")
	defer conn.Close()

	data := "hello world"

	contentBytes, err := json.Marshal(data)
	if err != nil {
		t.Error("json Marshal", err)
	}

	chatMsg := &chatpb.MsgData{
		SendID:           "9be675b41ff6feb3c21e528b92c812c1", // user01
		RecvID:           "2293575ce049296775c19640a61eab77", // user02
		GroupID:          "test_group",
		ClientMsgID:      "test_client_msg_id",
		ServerMsgID:      "test_server_msg_id",
		SenderPlatformID: 6,
		SenderNickname:   "test_nickname",
		SenderFaceURL:    "https://example.com/avatar.jpg",
		SessionType:      1,
		MsgFrom:          1,
		ContentType:      1,
		Content:          contentBytes,
		Seq:              1,
		ServerTime:       1625500000,
		ClientTime:       1625500010,
		OfflinePushInfo: &chatpb.OfflinePushInfo{
			Title: "Test title",
		},
		AtUserIDList: []string{"user1", "user2"},
		Options:      map[string]bool{"option1": true, "option2": false},
	}

	chatMsgBytes, err := proto.Marshal(chatMsg)
	if err != nil {
		t.Error("chatMsg proto.Marshal err:", err)
		return
	}

	Data := &pb.Req{
		ReqIdentifier: 2001,
		Token:         token,
		SendID:        sendId,
		Data:          chatMsgBytes,
	}

	bytes, err := proto.Marshal(Data)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}

	//发送消息
	err = conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		fmt.Println("write err:", err)
		return
	}

	// 接收消息
	_, message, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("read err:", err)
		return
	}

	fmt.Println("received message:", string(message))
}
