package wslogic

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"go.opentelemetry.io/otel/attribute"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/types"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/wsrepository"
	"goChat/app/msg-gateway/cmd/wsrpc/internal/wssvc"
	"goChat/app/usercenter/cmd/rpc/pb"
	"goChat/common/xtrace"
	"net/http"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserConn struct {
	*websocket.Conn
	w *sync.Mutex
}

type MsggatewayLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *wssvc.ServiceContext
	wsMaxConnNum int
	wsUpGrader   *websocket.Upgrader
	wsConnToUser map[*UserConn]map[string]string
	wsUserToConn map[string]map[string]*UserConn
	rep          *wsrepository.Rep
}

var msgGatewayLogic *MsggatewayLogic

func GetMsgGatewayLogic() *MsggatewayLogic {
	return msgGatewayLogic
}

func NewMsggatewayLogic(ctx context.Context, svcCtx *wssvc.ServiceContext) *MsggatewayLogic {
	if msgGatewayLogic != nil {
		return msgGatewayLogic
	}
	ws := &MsggatewayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		rep:    wsrepository.NewRep(svcCtx),
	}
	ws.wsMaxConnNum = ws.svcCtx.Config.Websocket.MaxConnNum
	ws.wsConnToUser = make(map[*UserConn]map[string]string)
	ws.wsUserToConn = make(map[string]map[string]*UserConn)
	ws.wsUpGrader = &websocket.Upgrader{
		HandshakeTimeout: time.Duration(ws.svcCtx.Config.Websocket.TimeOut) * time.Second,
		ReadBufferSize:   ws.svcCtx.Config.Websocket.MaxMsgLen,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	msgGatewayLogic = ws
	return msgGatewayLogic
}

func (l *MsggatewayLogic) Msggateway(req *types.Request) (*types.Response, bool) {
	fmt.Println("调用rpc验证token")
	// 调用rpc验证token
	if len(req.Token) != 0 && len(req.PlatformID) != 0 && len(req.SendID) != 0 {
		fmt.Println("前")
		resp, err := l.svcCtx.ImUserService.VerifyToken(l.ctx, &pb.VerifyTokenReq{
			Token:  req.Token,
			SendID: req.SendID,
		})
		fmt.Println("后")
		if err != nil {
			logx.WithContext(l.ctx).Errorf("调用VerifyToken失败，err: %s", err.Error())
			return &types.Response{
				Uid:     "",
				ErrMsg:  "调用VerifyToken失败",
				Success: false,
			}, false
		}
		if !resp.Success {
			logx.WithContext(l.ctx).Infof("VerifyToken失败，err: %s", resp.ErrMsg)
			return &types.Response{
				Uid:     resp.Uid,
				ErrMsg:  resp.ErrMsg,
				Success: false,
			}, false
		}
		return &types.Response{
			Uid:     resp.Uid,
			ErrMsg:  "",
			Success: true,
		}, true
	}
	return &types.Response{
		Uid:     "",
		ErrMsg:  "参数错误",
		Success: false,
	}, false
}

func (l *MsggatewayLogic) WsUpgrade(uid string, req *types.Request, w http.ResponseWriter, r *http.Request, header http.Header) error {
	fmt.Println("---WsUpgrade")
	conn, err := l.wsUpGrader.Upgrade(w, r, header)
	if err != nil {
		return err
	}
	newConn := &UserConn{conn, new(sync.Mutex)}
	userCount++
	fmt.Println("addUserConn start")
	fmt.Println("newConn", newConn)
	l.addUserConn(uid, req.PlatformID, newConn, req.Token)
	fmt.Println("addUserConn end")
	go l.readMsg(newConn, uid, req.PlatformID)
	fmt.Println("WsUpgrade end")
	return nil
}

func (l *MsggatewayLogic) readMsg(conn *UserConn, uid string, platform string) {
	fmt.Println("readMsg")
	for {
		messageType, msg, err := conn.ReadMessage()
		//conn.WriteMessage(messageType, append([]byte("还给你："), msg...))
		if messageType == websocket.PingMessage {
			l.sendMsg(context.Background(), conn, Resp{
				ReqIdentifier: 0,
				ErrCode:       0,
				ErrMsg:        "",
				Data:          []byte("pong"),
			})
		}
		if err != nil {
			uid, platform := l.getUserUid(conn)
			logx.Error("WS ReadMsg error ", " userIP ", conn.RemoteAddr().String(), " userUid ", uid, " platform ", platform, " error ", err.Error())
			userCount--
			l.delUserConn(conn)
			return
		}
		// xtrace
		xtrace.RunWithTrace("", func(ctx context.Context) {
			l.msgParse(ctx, conn, msg, uid, platform)
		}, attribute.KeyValue{
			Key:   "uid",
			Value: attribute.StringValue(uid),
		}, attribute.KeyValue{
			Key:   "platform",
			Value: attribute.StringValue(platform),
		})
	}
}
