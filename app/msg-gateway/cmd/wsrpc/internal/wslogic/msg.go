package wslogic

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"goChat/app/msg-gateway/cmd/wsrpc/pb"
	"goChat/app/msg/cmd/rpc/chat"
	chatpb "goChat/app/msg/cmd/rpc/pb"
	"goChat/common/types"
	"goChat/common/xerr"
	"google.golang.org/protobuf/proto"
)

func (l *MsggatewayLogic) sendMsg(ctx context.Context, conn *UserConn, mReply Resp) {
	resp := &pb.Resp{
		ReqIdentifier: uint32(mReply.ReqIdentifier),
		ErrCode:       mReply.ErrCode,
		ErrMsg:        mReply.ErrMsg,
		Data:          mReply.Data,
	}
	b, err := proto.Marshal(resp)
	if err != nil {
		uid, platform := l.getUserUid(conn)
		logx.WithContext(l.ctx).Error(mReply.ReqIdentifier, mReply.ErrCode, mReply.ErrMsg, "Encode Msg error: ", conn.RemoteAddr().String(), uid, platform, err.Error())
		return
	}
	err = l.writeMsg(conn, websocket.BinaryMessage, b)
	if err != nil {
		uid, platform := l.getUserUid(conn)
		logx.WithContext(l.ctx).Error(mReply.ReqIdentifier, mReply.ErrCode, mReply.ErrMsg, "Encode Msg error: ", conn.RemoteAddr().String(), uid, platform, err.Error())
	}
}

func (l *MsggatewayLogic) sendErrMsg(ctx context.Context, conn *UserConn, code int32, errMsg string, reqIdentifier *xerr.CodeError) {
	mReply := Resp{
		ReqIdentifier: int32(reqIdentifier.GetErrCode()),
		ErrCode:       uint32(code),
		ErrMsg:        errMsg,
	}
	l.sendMsg(ctx, conn, mReply)
}

func (l *MsggatewayLogic) writeMsg(conn *UserConn, a int, msg []byte) error {
	conn.w.Lock()
	defer conn.w.Unlock()
	return conn.WriteMessage(a, msg)
}

func (l *MsggatewayLogic) msgParse(ctx context.Context, conn *UserConn, binaryMsg []byte, uid string, platform string) {
	m := &pb.Req{}
	err := proto.Unmarshal(binaryMsg, m)
	logx.WithContext(l.ctx).Info("msgParse", m.ReqIdentifier, m.SendID, m.Data)
	if err != nil {
		l.sendErrMsg(ctx, conn, types.ErrCodeProtoUnmarshal, err.Error(), types.WSDataError)
		err = conn.Close()
		if err != nil {
			logx.WithContext(ctx).Error("ws close err", err.Error())
		}
		return
	}
	if err := validate.Struct(m); err != nil {
		logx.WithContext(ctx).Error("ws args validate  err", err.Error())
		l.sendErrMsg(ctx, conn, types.ErrCodeParams, err.Error(), xerr.NewErrCode(int(m.ReqIdentifier)))
		return
	}
	switch m.ReqIdentifier {
	case types.WSSendMsg:
		l.sendMsgReq(ctx, conn, m, uid)
	}
}

func (l *MsggatewayLogic) sendMsgReq(ctx context.Context, conn *UserConn, m *pb.Req, uid string) {
	fmt.Println("sendMsgReq", m.ReqIdentifier, m.SendID, m.Data)
	sendMsgAllCount++
	logx.WithContext(ctx).Info("Ws call success to sendMsgReq start", m.ReqIdentifier, m.SendID, m.Data)
	nReply := new(chatpb.SendMsgResp)
	isPass, errCode, errMsg, pData := l.argsValidate(m, types.WSSendMsg)
	if isPass {
		data := pData.(chatpb.MsgData)
		pbData := chatpb.SendMsgReq{
			Token:   m.Token,
			MsgData: &data,
		}
		logx.WithContext(ctx).Info("Ws call success to sendMsgReq middle", m.ReqIdentifier, m.SendID, data.String())

		reply, err := l.svcCtx.MsgRpc.SendMsg(ctx, &pbData)
		if err != nil {
			logx.WithContext(ctx).Error("UserSendMsg err ", err.Error())
			nReply.ErrCode = types.ErrCodeFailed
			nReply.ErrMsg = err.Error()
			l.sendMsgResp(ctx, conn, m, nReply)
		} else {
			logx.WithContext(ctx).Info("rpc call success to sendMsgReq", reply.String())
			l.sendMsgResp(ctx, conn, m, reply)
		}
	} else {
		nReply.ErrCode = errCode
		nReply.ErrMsg = errMsg
		l.sendMsgResp(ctx, conn, m, nReply)
	}
}

func (l *MsggatewayLogic) sendMsgResp(ctx context.Context, conn *UserConn, m *pb.Req, pb *chat.SendMsgResp) {
	var mReplyData chatpb.UserSendMsgResp
	mReplyData.ClientMsgID = pb.GetClientMsgID()
	mReplyData.ServerMsgID = pb.GetServerMsgID()
	mReplyData.ServerTime = pb.GetServerTime()
	b, _ := proto.Marshal(&mReplyData)
	mReply := Resp{
		ReqIdentifier: int32(m.ReqIdentifier),
		ErrCode:       uint32(pb.GetErrCode()),
		ErrMsg:        pb.GetErrMsg(),
		Data:          b,
	}

	l.sendMsg(ctx, conn, mReply)
}
