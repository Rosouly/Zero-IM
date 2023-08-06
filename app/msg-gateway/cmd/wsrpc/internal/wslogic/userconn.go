package wslogic

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"goChat/app/msg-gateway/cmd/wsrpc/pb"
)

func (l *MsggatewayLogic) addUserConn(uid string, platformID string, conn *UserConn, token string) {
	fmt.Println("addUserConn start")
	rwLock.Lock()
	defer rwLock.Unlock()
	if oldConnMap, ok := l.wsUserToConn[uid]; ok {
		oldConnMap[platformID] = conn
		l.wsUserToConn[uid] = oldConnMap
	} else {
		i := make(map[string]*UserConn)
		i[platformID] = conn
		l.wsUserToConn[uid] = i
	}
	fmt.Println("addUserConn mid")
	if oldStringMap, ok := l.wsConnToUser[conn]; ok {
		oldStringMap[platformID] = uid
		l.wsConnToUser[conn] = oldStringMap
	} else {
		i := make(map[string]string)
		i[platformID] = uid
		l.wsConnToUser[conn] = i
	}
	fmt.Println("addUserConn end")
	fmt.Println("uid, platformID", uid, platformID)
	count := 0
	for _, v := range l.wsUserToConn {
		count = count + len(v)
	}
	for _, v := range l.wsUserToConn {
		fmt.Println("v: ", v)
	}
}

func (l *MsggatewayLogic) getUserUid(conn *UserConn) (uid string, platform string) {
	rwLock.RLock()
	defer rwLock.RUnlock()
	if oldStringMap, ok := l.wsConnToUser[conn]; ok {
		for k, v := range oldStringMap {
			platform = k
			uid = v
		}
		return uid, platform
	}
	return "", ""
}

func (l *MsggatewayLogic) delUserConn(conn *UserConn) {
	rwLock.Lock()
	defer rwLock.Unlock()
	var platform, uid string
	if oldStringMap, ok := l.wsConnToUser[conn]; ok {
		for k, v := range oldStringMap {
			platform = k
			uid = v
		}
		if oldConnMap, ok := l.wsUserToConn[uid]; ok {
			// 因为将来可能有很多个平台
			delete(oldConnMap, platform)
			l.wsUserToConn[uid] = oldConnMap
			// 如果所有平台都下线了，就删除这个用户连接
			if len(oldConnMap) == 0 {
				delete(l.wsUserToConn, uid)
			}
		}
		delete(l.wsConnToUser, conn)
	}
	err := conn.Close()
	if err != nil {
		logx.WithContext(l.ctx).Error("close conn err ", "uid: ", uid, "platform: ", platform)
	}
}

func (l *MsggatewayLogic) GetUserConn(uid string, platform string) *UserConn {
	rwLock.RLock()
	defer rwLock.RUnlock()
	if oldConnMap, ok := l.wsUserToConn[uid]; ok {
		if conn, flag := oldConnMap[platform]; flag {
			return conn
		}
	}
	return nil
}

func (l *MsggatewayLogic) DelUserConn(uid string, platform string) {
	rwLock.RLock()
	defer rwLock.RUnlock()
	if oldConnMap, ok := l.wsUserToConn[uid]; ok {
		if len(oldConnMap) == 0 {
			delete(l.wsUserToConn, uid)
		} else {
			if _, flag := oldConnMap[platform]; flag {
				delete(oldConnMap, platform)
			}
		}
	}
}

func (l *MsggatewayLogic) SendMsgToUser(ctx context.Context, conn *UserConn, bMsg []byte, in *pb.OnlinePushMsgReq, RecvPlatForm, RecvID string) (ResultCode int64) {
	fmt.Println("SendMsgToUser start")
	err := l.writeMsg(conn, websocket.BinaryMessage, bMsg)
	if err != nil {
		logx.WithContext(ctx).Error("send msg to user err ", "", "err ", err.Error())
		ResultCode = -2
		return ResultCode
	} else {
		ResultCode = 0
		return ResultCode
	}
}
