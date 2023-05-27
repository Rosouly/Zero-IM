package wslogic

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"goChat/app/msg-gateway/cmd/wsrpc/pb"
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
	}
	logx.WithContext(ctx).Info("Ws call success to sendMsgReq middle", m.ReqIdentifier, m.SendID, data.String())

	reply, err := l.svcCtx.MsgRpc
}
