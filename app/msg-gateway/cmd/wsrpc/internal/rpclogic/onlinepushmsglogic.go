package rpclogic

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/rpcsvc"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/wslogic"
	"goChat/app/msg-gateway/cmd/wsrpc/pb"
	"goChat/common/types"
	"goChat/common/xtrace"
	"strconv"
)

type OnlinePushMsgLogic struct {
	ctx    context.Context
	svcCtx *rpcsvc.ServiceContext
	logx.Logger
}

func NewOnlinePushMsgLogic(ctx context.Context, svcCtx *rpcsvc.ServiceContext) *OnlinePushMsgLogic {
	return &OnlinePushMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OnlinePushMsgLogic) OnlinePushMsg(in *pb.OnlinePushMsgReq) (*pb.OnlinePushMsgResp, error) {
	logic := wslogic.GetMsgGatewayLogic()
	fmt.Println("OnlinePushMsgLogic.OnlinePushMsg")
	//logic := wslogic.NewMsggatewayLogic(nil, nil)
	fmt.Println("NewMsggatewayLogic 1")
	var resp []*pb.SingleMsgToUser
	msgBytes, _ := proto.Marshal(in.MsgData)
	reqIdentifier := types.WSPushMsg
	if in.MsgData.SessionType == types.GroupChatType {
		reqIdentifier = types.WSGroupPushMsg
	}
	mReply := &pb.Resp{
		ReqIdentifier: uint32(reqIdentifier),
		Data:          msgBytes,
	}
	replyBytes, err := proto.Marshal(mReply)
	if err != nil {
		l.Error("data encode err ", err.Error())
	}
	var tag bool
	recvID := in.PushToUserID
	platformList := []string{
		"1",
		"2",
		"3",
		"4",
		"6",
		"7",
		"5",
	}
	//receiver := logic.GetUserConn("2293575ce049296775c19640a61eab77", "6")
	//fmt.Println("receiver: ", receiver)
	//fmt.Println("receiver id:", recvID)
	for _, v := range platformList {
		if conn := logic.GetUserConn(recvID, v); conn != nil {
			fmt.Println("v: ", v)
			fmt.Println("conn: ", conn)
			tag = true
			var resultCode int64
			xtrace.StartFuncSpan(l.ctx, "OnlinePushMsgLogic.OnlinePushMsg", func(ctx context.Context) {

				resultCode = logic.SendMsgToUser(ctx, conn, replyBytes, in, v, recvID)
				fmt.Println("resultcode: ", resultCode)
			})
			i64, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return nil, err
			}
			temp := &pb.SingleMsgToUser{
				ResultCode:     resultCode,
				RecvID:         recvID,
				RecvPlatFormID: int32(i64),
			}
			resp = append(resp, temp)
		} else {
			i64, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return nil, err
			}
			temp := &pb.SingleMsgToUser{
				ResultCode:     -1,
				RecvID:         recvID,
				RecvPlatFormID: int32(i64),
			}
			resp = append(resp, temp)
		}
	}
	if !tag {
	}
	fmt.Println("resp", resp)
	return &pb.OnlinePushMsgResp{
		Resp: resp,
	}, nil
}
