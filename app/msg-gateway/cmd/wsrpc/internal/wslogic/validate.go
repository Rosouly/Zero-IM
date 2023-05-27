package wslogic

import (
	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
	msggatewaypb "goChat/app/msg-gateway/cmd/wsrpc/pb"
	chatpb "goChat/app/msg/cmd/rpc/pb"
	"goChat/common/types"
)

type Resp struct {
	ReqIdentifier int32  `json:"reqIdentifier""`
	ErrCode       uint32 `json:"errCode"`
	ErrMsg        string `json:"errMsg"`
	Data          []byte `json:"data"`
}

func (l *MsggatewayLogic) argsValidate(m *msggatewaypb.Req, r int32) (isPass bool, errCode int32, errMsg string, returnData interface{}) {
	switch r {
	case types.WSSendMsg:
		data := chatpb.MsgData{}
		if err := proto.Unmarshal(m.Data, &data); err != nil {
			logx.WithContext(l.ctx).Error("Decode Data struct  err", err.Error(), r)
			return false, types.ErrCodeProtoUnmarshal, err.Error(), nil
		}
		if err := validate.Struct(data); err != nil {
			logx.WithContext(l.ctx).Error("data args validate  err", err.Error(), r)
			return false, types.ErrCodeParams, err.Error(), nil

		}
		return true, types.ErrCodeOK, "", data
	default:
	}
	return false, types.ErrCodeParams, "args err", nil
}
