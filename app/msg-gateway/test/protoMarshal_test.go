package test

import (
	"github.com/golang/protobuf/proto"
	"goChat/app/msg-gateway/cmd/wsrpc/pb"
	"testing"
)

func TestProtoMarshal(t *testing.T) {
	data := &pb.Req{
		ReqIdentifier: 1,
		Token:         "token1324",
		SendID:        "1",
		Data:          []byte("hello"),
	}

	bytes, err := proto.Marshal(data)
	if err != nil {
		t.Error("proto Marshal", err)

	}

	res := &pb.Req{}
	err = proto.Unmarshal(bytes, res)
	if err != nil {
		t.Error("proto Unmarshal", err)
	}

	t.Log("res", res)
}
