// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: usercenter.proto

package pb

import (
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

type BaseResp_Errno int32

const (
	BaseResp_OK    BaseResp_Errno = 0
	BaseResp_ERROR BaseResp_Errno = 1
)

// Enum value maps for BaseResp_Errno.
var (
	BaseResp_Errno_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	BaseResp_Errno_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x BaseResp_Errno) Enum() *BaseResp_Errno {
	p := new(BaseResp_Errno)
	*p = x
	return p
}

func (x BaseResp_Errno) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseResp_Errno) Descriptor() protoreflect.EnumDescriptor {
	return file_usercenter_proto_enumTypes[0].Descriptor()
}

func (BaseResp_Errno) Type() protoreflect.EnumType {
	return &file_usercenter_proto_enumTypes[0]
}

func (x BaseResp_Errno) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseResp_Errno.Descriptor instead.
func (BaseResp_Errno) EnumDescriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{0, 0}
}

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode BaseResp_Errno `protobuf:"varint,1,opt,name=errCode,proto3,enum=usercenter.BaseResp_Errno" json:"errCode,omitempty"`
	ErrMsg  string         `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetErrCode() BaseResp_Errno {
	if x != nil {
		return x.ErrCode
	}
	return BaseResp_OK
}

func (x *BaseResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password     string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Nickname     string `protobuf:"bytes,4,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	Sign         string `protobuf:"bytes,5,opt,name=Sign,proto3" json:"Sign,omitempty"`
	Avatar       string `protobuf:"bytes,6,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Province     string `protobuf:"bytes,7,opt,name=Province,proto3" json:"Province,omitempty"`
	City         string `protobuf:"bytes,8,opt,name=City,proto3" json:"City,omitempty"`
	District     string `protobuf:"bytes,9,opt,name=District,proto3" json:"District,omitempty"`
	Birthday     string `protobuf:"bytes,10,opt,name=Birthday,proto3" json:"Birthday,omitempty"`
	RegisterTime string `protobuf:"bytes,11,opt,name=RegisterTime,proto3" json:"RegisterTime,omitempty"`
	IsMale       bool   `protobuf:"varint,12,opt,name=IsMale,proto3" json:"IsMale,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *User) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *User) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *User) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *User) GetRegisterTime() string {
	if x != nil {
		return x.RegisterTime
	}
	return ""
}

func (x *User) GetIsMale() bool {
	if x != nil {
		return x.IsMale
	}
	return false
}

type GeTUserByUsernameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GeTUserByUsernameReq) Reset() {
	*x = GeTUserByUsernameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeTUserByUsernameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeTUserByUsernameReq) ProtoMessage() {}

func (x *GeTUserByUsernameReq) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeTUserByUsernameReq.ProtoReflect.Descriptor instead.
func (*GeTUserByUsernameReq) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{2}
}

func (x *GeTUserByUsernameReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GeTUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GeTUserResp) Reset() {
	*x = GeTUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeTUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeTUserResp) ProtoMessage() {}

func (x *GeTUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeTUserResp.ProtoReflect.Descriptor instead.
func (*GeTUserResp) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{3}
}

func (x *GeTUserResp) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type InsertUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *InsertUserReq) Reset() {
	*x = InsertUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUserReq) ProtoMessage() {}

func (x *InsertUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUserReq.ProtoReflect.Descriptor instead.
func (*InsertUserReq) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{4}
}

func (x *InsertUserReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type InsertUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=baseResp,proto3" json:"baseResp,omitempty"`
}

func (x *InsertUserResp) Reset() {
	*x = InsertUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUserResp) ProtoMessage() {}

func (x *InsertUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUserResp.ProtoReflect.Descriptor instead.
func (*InsertUserResp) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{5}
}

func (x *InsertUserResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type LoginByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *LoginByIdReq) Reset() {
	*x = LoginByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByIdReq) ProtoMessage() {}

func (x *LoginByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByIdReq.ProtoReflect.Descriptor instead.
func (*LoginByIdReq) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{6}
}

func (x *LoginByIdReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LoginByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *User     `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Token    string    `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	BaseResp *BaseResp `protobuf:"bytes,3,opt,name=baseResp,proto3" json:"baseResp,omitempty"`
}

func (x *LoginByIdResp) Reset() {
	*x = LoginByIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByIdResp) ProtoMessage() {}

func (x *LoginByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByIdResp.ProtoReflect.Descriptor instead.
func (*LoginByIdResp) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{7}
}

func (x *LoginByIdResp) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *LoginByIdResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginByIdResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

var File_usercenter_proto protoreflect.FileDescriptor

var file_usercenter_proto_rawDesc = []byte{
	0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x74,
	0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x2e, 0x45, 0x72, 0x72, 0x6e, 0x6f, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x1a, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6e,
	0x6f, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x01, 0x22, 0xba, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x4d,
	0x61, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x4d, 0x61, 0x6c,
	0x65, 0x22, 0x32, 0x0a, 0x14, 0x47, 0x65, 0x54, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x0b, 0x47, 0x65, 0x54, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x0d, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x42, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x26, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7d, 0x0a,
	0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x30, 0x0a, 0x08, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x32, 0xf0, 0x01, 0x0a,
	0x11, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x47, 0x65, 0x54, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x54, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x54, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usercenter_proto_rawDescOnce sync.Once
	file_usercenter_proto_rawDescData = file_usercenter_proto_rawDesc
)

func file_usercenter_proto_rawDescGZIP() []byte {
	file_usercenter_proto_rawDescOnce.Do(func() {
		file_usercenter_proto_rawDescData = protoimpl.X.CompressGZIP(file_usercenter_proto_rawDescData)
	})
	return file_usercenter_proto_rawDescData
}

var file_usercenter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_usercenter_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_usercenter_proto_goTypes = []interface{}{
	(BaseResp_Errno)(0),          // 0: usercenter.BaseResp.Errno
	(*BaseResp)(nil),             // 1: usercenter.BaseResp
	(*User)(nil),                 // 2: usercenter.User
	(*GeTUserByUsernameReq)(nil), // 3: usercenter.GeTUserByUsernameReq
	(*GeTUserResp)(nil),          // 4: usercenter.GeTUserResp
	(*InsertUserReq)(nil),        // 5: usercenter.InsertUserReq
	(*InsertUserResp)(nil),       // 6: usercenter.InsertUserResp
	(*LoginByIdReq)(nil),         // 7: usercenter.LoginByIdReq
	(*LoginByIdResp)(nil),        // 8: usercenter.LoginByIdResp
}
var file_usercenter_proto_depIdxs = []int32{
	0, // 0: usercenter.BaseResp.errCode:type_name -> usercenter.BaseResp.Errno
	2, // 1: usercenter.GeTUserResp.user:type_name -> usercenter.User
	2, // 2: usercenter.InsertUserReq.user:type_name -> usercenter.User
	1, // 3: usercenter.InsertUserResp.baseResp:type_name -> usercenter.BaseResp
	2, // 4: usercenter.LoginByIdResp.user:type_name -> usercenter.User
	1, // 5: usercenter.LoginByIdResp.baseResp:type_name -> usercenter.BaseResp
	3, // 6: usercenter.usercenterService.GeTUserByUsername:input_type -> usercenter.GeTUserByUsernameReq
	5, // 7: usercenter.usercenterService.InsertUser:input_type -> usercenter.InsertUserReq
	7, // 8: usercenter.usercenterService.LoginById:input_type -> usercenter.LoginByIdReq
	4, // 9: usercenter.usercenterService.GeTUserByUsername:output_type -> usercenter.GeTUserResp
	6, // 10: usercenter.usercenterService.InsertUser:output_type -> usercenter.InsertUserResp
	8, // 11: usercenter.usercenterService.LoginById:output_type -> usercenter.LoginByIdResp
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_usercenter_proto_init() }
func file_usercenter_proto_init() {
	if File_usercenter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_usercenter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_usercenter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_usercenter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeTUserByUsernameReq); i {
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
		file_usercenter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeTUserResp); i {
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
		file_usercenter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertUserReq); i {
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
		file_usercenter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertUserResp); i {
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
		file_usercenter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginByIdReq); i {
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
		file_usercenter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginByIdResp); i {
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
			RawDescriptor: file_usercenter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usercenter_proto_goTypes,
		DependencyIndexes: file_usercenter_proto_depIdxs,
		EnumInfos:         file_usercenter_proto_enumTypes,
		MessageInfos:      file_usercenter_proto_msgTypes,
	}.Build()
	File_usercenter_proto = out.File
	file_usercenter_proto_rawDesc = nil
	file_usercenter_proto_goTypes = nil
	file_usercenter_proto_depIdxs = nil
}
