// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: msg-gateway.proto

package pb

import (
	pb "goChat/app/msg/cmd/rpc/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OnlinePushMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgData      *pb.MsgData `protobuf:"bytes,2,opt,name=msgData,proto3" json:"msgData,omitempty"`
	PushToUserID string   `protobuf:"bytes,3,opt,name=pushToUserID,proto3" json:"pushToUserID,omitempty"`
}

func (x *OnlinePushMsgReq) Reset() {
	*x = OnlinePushMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlinePushMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlinePushMsgReq) ProtoMessage() {}

func (x *OnlinePushMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlinePushMsgReq.ProtoReflect.Descriptor instead.
func (*OnlinePushMsgReq) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *OnlinePushMsgReq) GetMsgData() *pb.MsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

func (x *OnlinePushMsgReq) GetPushToUserID() string {
	if x != nil {
		return x.PushToUserID
	}
	return ""
}

type OnlinePushMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp []*SingleMsgToUser `protobuf:"bytes,1,rep,name=resp,proto3" json:"resp,omitempty"`
}

func (x *OnlinePushMsgResp) Reset() {
	*x = OnlinePushMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlinePushMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlinePushMsgResp) ProtoMessage() {}

func (x *OnlinePushMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlinePushMsgResp.ProtoReflect.Descriptor instead.
func (*OnlinePushMsgResp) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *OnlinePushMsgResp) GetResp() []*SingleMsgToUser {
	if x != nil {
		return x.Resp
	}
	return nil
}

type SingleMsgToUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode     int64  `protobuf:"varint,1,opt,name=ResultCode,proto3" json:"ResultCode,omitempty"`
	RecvID         string `protobuf:"bytes,2,opt,name=RecvID,proto3" json:"RecvID,omitempty"`
	RecvPlatFormID int32  `protobuf:"varint,3,opt,name=RecvPlatFormID,proto3" json:"RecvPlatFormID,omitempty"`
}

func (x *SingleMsgToUser) Reset() {
	*x = SingleMsgToUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleMsgToUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleMsgToUser) ProtoMessage() {}

func (x *SingleMsgToUser) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleMsgToUser.ProtoReflect.Descriptor instead.
func (*SingleMsgToUser) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *SingleMsgToUser) GetResultCode() int64 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

func (x *SingleMsgToUser) GetRecvID() string {
	if x != nil {
		return x.RecvID
	}
	return ""
}

func (x *SingleMsgToUser) GetRecvPlatFormID() int32 {
	if x != nil {
		return x.RecvPlatFormID
	}
	return 0
}

type GetUsersOnlineStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDList  []string `protobuf:"bytes,1,rep,name=userIDList,proto3" json:"userIDList,omitempty"`
	OperationID string   `protobuf:"bytes,2,opt,name=operationID,proto3" json:"operationID,omitempty"`
	OpUserID    string   `protobuf:"bytes,3,opt,name=opUserID,proto3" json:"opUserID,omitempty"`
}

func (x *GetUsersOnlineStatusReq) Reset() {
	*x = GetUsersOnlineStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersOnlineStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersOnlineStatusReq) ProtoMessage() {}

func (x *GetUsersOnlineStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersOnlineStatusReq.ProtoReflect.Descriptor instead.
func (*GetUsersOnlineStatusReq) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *GetUsersOnlineStatusReq) GetUserIDList() []string {
	if x != nil {
		return x.UserIDList
	}
	return nil
}

func (x *GetUsersOnlineStatusReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *GetUsersOnlineStatusReq) GetOpUserID() string {
	if x != nil {
		return x.OpUserID
	}
	return ""
}

type GetUsersOnlineStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode       int32                                     `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg        string                                    `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	SuccessResult []*GetUsersOnlineStatusResp_SuccessResult `protobuf:"bytes,3,rep,name=successResult,proto3" json:"successResult,omitempty"`
	FailedResult  []*GetUsersOnlineStatusResp_FailedDetail  `protobuf:"bytes,4,rep,name=failedResult,proto3" json:"failedResult,omitempty"`
}

func (x *GetUsersOnlineStatusResp) Reset() {
	*x = GetUsersOnlineStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersOnlineStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersOnlineStatusResp) ProtoMessage() {}

func (x *GetUsersOnlineStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersOnlineStatusResp.ProtoReflect.Descriptor instead.
func (*GetUsersOnlineStatusResp) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *GetUsersOnlineStatusResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *GetUsersOnlineStatusResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *GetUsersOnlineStatusResp) GetSuccessResult() []*GetUsersOnlineStatusResp_SuccessResult {
	if x != nil {
		return x.SuccessResult
	}
	return nil
}

func (x *GetUsersOnlineStatusResp) GetFailedResult() []*GetUsersOnlineStatusResp_FailedDetail {
	if x != nil {
		return x.FailedResult
	}
	return nil
}

type GetUsersOnlineStatusResp_SuccessDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform string `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetUsersOnlineStatusResp_SuccessDetail) Reset() {
	*x = GetUsersOnlineStatusResp_SuccessDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersOnlineStatusResp_SuccessDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersOnlineStatusResp_SuccessDetail) ProtoMessage() {}

func (x *GetUsersOnlineStatusResp_SuccessDetail) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersOnlineStatusResp_SuccessDetail.ProtoReflect.Descriptor instead.
func (*GetUsersOnlineStatusResp_SuccessDetail) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetUsersOnlineStatusResp_SuccessDetail) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *GetUsersOnlineStatusResp_SuccessDetail) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetUsersOnlineStatusResp_FailedDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	ErrCode int32  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
}

func (x *GetUsersOnlineStatusResp_FailedDetail) Reset() {
	*x = GetUsersOnlineStatusResp_FailedDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersOnlineStatusResp_FailedDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersOnlineStatusResp_FailedDetail) ProtoMessage() {}

func (x *GetUsersOnlineStatusResp_FailedDetail) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersOnlineStatusResp_FailedDetail.ProtoReflect.Descriptor instead.
func (*GetUsersOnlineStatusResp_FailedDetail) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{4, 1}
}

func (x *GetUsersOnlineStatusResp_FailedDetail) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetUsersOnlineStatusResp_FailedDetail) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *GetUsersOnlineStatusResp_FailedDetail) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type GetUsersOnlineStatusResp_SuccessResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID               string                                    `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Status               string                                    `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	DetailPlatformStatus []*GetUsersOnlineStatusResp_SuccessDetail `protobuf:"bytes,3,rep,name=detailPlatformStatus,proto3" json:"detailPlatformStatus,omitempty"`
}

func (x *GetUsersOnlineStatusResp_SuccessResult) Reset() {
	*x = GetUsersOnlineStatusResp_SuccessResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersOnlineStatusResp_SuccessResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersOnlineStatusResp_SuccessResult) ProtoMessage() {}

func (x *GetUsersOnlineStatusResp_SuccessResult) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersOnlineStatusResp_SuccessResult.ProtoReflect.Descriptor instead.
func (*GetUsersOnlineStatusResp_SuccessResult) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{4, 2}
}

func (x *GetUsersOnlineStatusResp_SuccessResult) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetUsersOnlineStatusResp_SuccessResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetUsersOnlineStatusResp_SuccessResult) GetDetailPlatformStatus() []*GetUsersOnlineStatusResp_SuccessDetail {
	if x != nil {
		return x.DetailPlatformStatus
	}
	return nil
}

var File_msg_gateway_proto protoreflect.FileDescriptor

var file_msg_gateway_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x73, 0x67, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x1a, 0x08, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x10, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x34,
	0x0a, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x75, 0x73, 0x68,
	0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x45, 0x0a, 0x11, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a,
	0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x73,
	0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x4d, 0x73, 0x67, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x22,
	0x71, 0x0a, 0x0f, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x76, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x63, 0x76, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65,
	0x63, 0x76, 0x50, 0x6c, 0x61, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x52, 0x65, 0x63, 0x76, 0x50, 0x6c, 0x61, 0x74, 0x46, 0x6f, 0x72, 0x6d,
	0x49, 0x44, 0x22, 0x77, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0xc9, 0x04, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x59, 0x0a, 0x0d, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x56, 0x0a, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6d, 0x73,
	0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x0c, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x43, 0x0a,
	0x0d, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x58, 0x0a, 0x0c, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x1a, 0xa8, 0x01, 0x0a,
	0x0d, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x67,
	0x0a, 0x14, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6d,
	0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x14, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xd0, 0x01, 0x0a, 0x19, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x2e, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x63, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x2e,
	0x6d, 0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_gateway_proto_rawDescOnce sync.Once
	file_msg_gateway_proto_rawDescData = file_msg_gateway_proto_rawDesc
)

func file_msg_gateway_proto_rawDescGZIP() []byte {
	file_msg_gateway_proto_rawDescOnce.Do(func() {
		file_msg_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_gateway_proto_rawDescData)
	})
	return file_msg_gateway_proto_rawDescData
}

var file_msg_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_msg_gateway_proto_goTypes = []interface{}{
	(*OnlinePushMsgReq)(nil),                       // 0: msg_gateway.OnlinePushMsgReq
	(*OnlinePushMsgResp)(nil),                      // 1: msg_gateway.OnlinePushMsgResp
	(*SingleMsgToUser)(nil),                        // 2: msg_gateway.SingleMsgToUser
	(*GetUsersOnlineStatusReq)(nil),                // 3: msg_gateway.GetUsersOnlineStatusReq
	(*GetUsersOnlineStatusResp)(nil),               // 4: msg_gateway.GetUsersOnlineStatusResp
	(*GetUsersOnlineStatusResp_SuccessDetail)(nil), // 5: msg_gateway.GetUsersOnlineStatusResp.SuccessDetail
	(*GetUsersOnlineStatusResp_FailedDetail)(nil),  // 6: msg_gateway.GetUsersOnlineStatusResp.FailedDetail
	(*GetUsersOnlineStatusResp_SuccessResult)(nil), // 7: msg_gateway.GetUsersOnlineStatusResp.SuccessResult
	(*pb.MsgData)(nil),                                // 8: server_api_params.MsgData
}
var file_msg_gateway_proto_depIdxs = []int32{
	8, // 0: msg_gateway.OnlinePushMsgReq.msgData:type_name -> server_api_params.MsgData
	2, // 1: msg_gateway.OnlinePushMsgResp.resp:type_name -> msg_gateway.SingleMsgToUser
	7, // 2: msg_gateway.GetUsersOnlineStatusResp.successResult:type_name -> msg_gateway.GetUsersOnlineStatusResp.SuccessResult
	6, // 3: msg_gateway.GetUsersOnlineStatusResp.failedResult:type_name -> msg_gateway.GetUsersOnlineStatusResp.FailedDetail
	5, // 4: msg_gateway.GetUsersOnlineStatusResp.SuccessResult.detailPlatformStatus:type_name -> msg_gateway.GetUsersOnlineStatusResp.SuccessDetail
	0, // 5: msg_gateway.OnlineMessageRelayService.OnlinePushMsg:input_type -> msg_gateway.OnlinePushMsgReq
	3, // 6: msg_gateway.OnlineMessageRelayService.GetUsersOnlineStatus:input_type -> msg_gateway.GetUsersOnlineStatusReq
	1, // 7: msg_gateway.OnlineMessageRelayService.OnlinePushMsg:output_type -> msg_gateway.OnlinePushMsgResp
	4, // 8: msg_gateway.OnlineMessageRelayService.GetUsersOnlineStatus:output_type -> msg_gateway.GetUsersOnlineStatusResp
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_msg_gateway_proto_init() }
func file_msg_gateway_proto_init() {
	if File_msg_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlinePushMsgReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlinePushMsgResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleMsgToUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersOnlineStatusReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_gateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersOnlineStatusResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_gateway_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersOnlineStatusResp_SuccessDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_gateway_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersOnlineStatusResp_FailedDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_gateway_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersOnlineStatusResp_SuccessResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msg_gateway_proto_goTypes,
		DependencyIndexes: file_msg_gateway_proto_depIdxs,
		MessageInfos:      file_msg_gateway_proto_msgTypes,
	}.Build()
	File_msg_gateway_proto = out.File
	file_msg_gateway_proto_rawDesc = nil
	file_msg_gateway_proto_goTypes = nil
	file_msg_gateway_proto_depIdxs = nil
}
