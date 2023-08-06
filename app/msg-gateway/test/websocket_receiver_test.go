package test

import (
	"fmt"
	"github.com/gorilla/websocket"
	"testing"
)

func TestWebsocketReceiver(t *testing.T) {
	sendId := "2293575ce049296775c19640a61eab77"
	platformId := "6"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiIyMjkzNTc1Y2UwNDkyOTY3NzVjMTk2NDBhNjFlYWI3NyIsIlBsYXRmb3JtIjoid2ViIiwiZXhwIjo0ODQzNTA2NTc1LCJuYmYiOjE2ODk5MDY1NzUsImlhdCI6MTY4OTkwNjU3NX0.YoaBsVIf5p_CVXsnGshDJ4l3sgQQkyp2ZUkpoP_8OnY"
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

	i := 1
	for i > 0 {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println("收到的消息为: ", string(bytes))
		i--
	}
	t.Log("success")
}
