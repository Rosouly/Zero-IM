// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: im_user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ImUserService_GetGroupMemberIDListFromCache_FullMethodName    = "/imuser.imUserService/GetGroupMemberIDListFromCache"
	ImUserService_IfAInBBlacklist_FullMethodName                  = "/imuser.imUserService/IfAInBBlacklist"
	ImUserService_IfAInBFriendList_FullMethodName                 = "/imuser.imUserService/IfAInBFriendList"
	ImUserService_GetSingleConversationRecvMsgOpts_FullMethodName = "/imuser.imUserService/GetSingleConversationRecvMsgOpts"
	ImUserService_GetUserListFromSuperGroupWithOpt_FullMethodName = "/imuser.imUserService/GetUserListFromSuperGroupWithOpt"
	ImUserService_VerifyToken_FullMethodName                      = "/imuser.imUserService/VerifyToken"
)

// ImUserServiceClient is the client API for ImUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImUserServiceClient interface {
	// 获取群组成员列表
	GetGroupMemberIDListFromCache(ctx context.Context, in *GetGroupMemberIDListFromCacheReq, opts ...grpc.CallOption) (*GetGroupMemberIDListFromCacheResp, error)
	// 判断用户A是否在B黑名单中
	IfAInBBlacklist(ctx context.Context, in *IfAInBBlacklistReq, opts ...grpc.CallOption) (*IfAInBBlacklistResp, error)
	// 判断用户A是否在B好友列表中
	IfAInBFriendList(ctx context.Context, in *IfAInBFriendListReq, opts ...grpc.CallOption) (*IfAInBFriendListResp, error)
	// 获取单聊会话的消息接收选项
	GetSingleConversationRecvMsgOpts(ctx context.Context, in *GetSingleConversationRecvMsgOptsReq, opts ...grpc.CallOption) (*GetSingleConversationRecvMsgOptsResp, error)
	// 获取超级群成员列表 通过消息接收选项
	GetUserListFromSuperGroupWithOpt(ctx context.Context, in *GetUserListFromSuperGroupWithOptReq, opts ...grpc.CallOption) (*GetUserListFromSuperGroupWithOptResp, error)
	// 检查token
	VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenResp, error)
}

type imUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImUserServiceClient(cc grpc.ClientConnInterface) ImUserServiceClient {
	return &imUserServiceClient{cc}
}

func (c *imUserServiceClient) GetGroupMemberIDListFromCache(ctx context.Context, in *GetGroupMemberIDListFromCacheReq, opts ...grpc.CallOption) (*GetGroupMemberIDListFromCacheResp, error) {
	out := new(GetGroupMemberIDListFromCacheResp)
	err := c.cc.Invoke(ctx, ImUserService_GetGroupMemberIDListFromCache_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imUserServiceClient) IfAInBBlacklist(ctx context.Context, in *IfAInBBlacklistReq, opts ...grpc.CallOption) (*IfAInBBlacklistResp, error) {
	out := new(IfAInBBlacklistResp)
	err := c.cc.Invoke(ctx, ImUserService_IfAInBBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imUserServiceClient) IfAInBFriendList(ctx context.Context, in *IfAInBFriendListReq, opts ...grpc.CallOption) (*IfAInBFriendListResp, error) {
	out := new(IfAInBFriendListResp)
	err := c.cc.Invoke(ctx, ImUserService_IfAInBFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imUserServiceClient) GetSingleConversationRecvMsgOpts(ctx context.Context, in *GetSingleConversationRecvMsgOptsReq, opts ...grpc.CallOption) (*GetSingleConversationRecvMsgOptsResp, error) {
	out := new(GetSingleConversationRecvMsgOptsResp)
	err := c.cc.Invoke(ctx, ImUserService_GetSingleConversationRecvMsgOpts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imUserServiceClient) GetUserListFromSuperGroupWithOpt(ctx context.Context, in *GetUserListFromSuperGroupWithOptReq, opts ...grpc.CallOption) (*GetUserListFromSuperGroupWithOptResp, error) {
	out := new(GetUserListFromSuperGroupWithOptResp)
	err := c.cc.Invoke(ctx, ImUserService_GetUserListFromSuperGroupWithOpt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imUserServiceClient) VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenResp, error) {
	out := new(VerifyTokenResp)
	err := c.cc.Invoke(ctx, ImUserService_VerifyToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImUserServiceServer is the server API for ImUserService service.
// All implementations must embed UnimplementedImUserServiceServer
// for forward compatibility
type ImUserServiceServer interface {
	// 获取群组成员列表
	GetGroupMemberIDListFromCache(context.Context, *GetGroupMemberIDListFromCacheReq) (*GetGroupMemberIDListFromCacheResp, error)
	// 判断用户A是否在B黑名单中
	IfAInBBlacklist(context.Context, *IfAInBBlacklistReq) (*IfAInBBlacklistResp, error)
	// 判断用户A是否在B好友列表中
	IfAInBFriendList(context.Context, *IfAInBFriendListReq) (*IfAInBFriendListResp, error)
	// 获取单聊会话的消息接收选项
	GetSingleConversationRecvMsgOpts(context.Context, *GetSingleConversationRecvMsgOptsReq) (*GetSingleConversationRecvMsgOptsResp, error)
	// 获取超级群成员列表 通过消息接收选项
	GetUserListFromSuperGroupWithOpt(context.Context, *GetUserListFromSuperGroupWithOptReq) (*GetUserListFromSuperGroupWithOptResp, error)
	// 检查token
	VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenResp, error)
	mustEmbedUnimplementedImUserServiceServer()
}

// UnimplementedImUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImUserServiceServer struct {
}

func (UnimplementedImUserServiceServer) GetGroupMemberIDListFromCache(context.Context, *GetGroupMemberIDListFromCacheReq) (*GetGroupMemberIDListFromCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupMemberIDListFromCache not implemented")
}
func (UnimplementedImUserServiceServer) IfAInBBlacklist(context.Context, *IfAInBBlacklistReq) (*IfAInBBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IfAInBBlacklist not implemented")
}
func (UnimplementedImUserServiceServer) IfAInBFriendList(context.Context, *IfAInBFriendListReq) (*IfAInBFriendListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IfAInBFriendList not implemented")
}
func (UnimplementedImUserServiceServer) GetSingleConversationRecvMsgOpts(context.Context, *GetSingleConversationRecvMsgOptsReq) (*GetSingleConversationRecvMsgOptsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleConversationRecvMsgOpts not implemented")
}
func (UnimplementedImUserServiceServer) GetUserListFromSuperGroupWithOpt(context.Context, *GetUserListFromSuperGroupWithOptReq) (*GetUserListFromSuperGroupWithOptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserListFromSuperGroupWithOpt not implemented")
}
func (UnimplementedImUserServiceServer) VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedImUserServiceServer) mustEmbedUnimplementedImUserServiceServer() {}

// UnsafeImUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImUserServiceServer will
// result in compilation errors.
type UnsafeImUserServiceServer interface {
	mustEmbedUnimplementedImUserServiceServer()
}

func RegisterImUserServiceServer(s grpc.ServiceRegistrar, srv ImUserServiceServer) {
	s.RegisterService(&ImUserService_ServiceDesc, srv)
}

func _ImUserService_GetGroupMemberIDListFromCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMemberIDListFromCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImUserServiceServer).GetGroupMemberIDListFromCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImUserService_GetGroupMemberIDListFromCache_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImUserServiceServer).GetGroupMemberIDListFromCache(ctx, req.(*GetGroupMemberIDListFromCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImUserService_IfAInBBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IfAInBBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImUserServiceServer).IfAInBBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImUserService_IfAInBBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImUserServiceServer).IfAInBBlacklist(ctx, req.(*IfAInBBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImUserService_IfAInBFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IfAInBFriendListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImUserServiceServer).IfAInBFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImUserService_IfAInBFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImUserServiceServer).IfAInBFriendList(ctx, req.(*IfAInBFriendListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImUserService_GetSingleConversationRecvMsgOpts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleConversationRecvMsgOptsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImUserServiceServer).GetSingleConversationRecvMsgOpts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImUserService_GetSingleConversationRecvMsgOpts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImUserServiceServer).GetSingleConversationRecvMsgOpts(ctx, req.(*GetSingleConversationRecvMsgOptsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImUserService_GetUserListFromSuperGroupWithOpt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListFromSuperGroupWithOptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImUserServiceServer).GetUserListFromSuperGroupWithOpt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImUserService_GetUserListFromSuperGroupWithOpt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImUserServiceServer).GetUserListFromSuperGroupWithOpt(ctx, req.(*GetUserListFromSuperGroupWithOptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImUserService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImUserServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImUserService_VerifyToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImUserServiceServer).VerifyToken(ctx, req.(*VerifyTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ImUserService_ServiceDesc is the grpc.ServiceDesc for ImUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "imuser.imUserService",
	HandlerType: (*ImUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroupMemberIDListFromCache",
			Handler:    _ImUserService_GetGroupMemberIDListFromCache_Handler,
		},
		{
			MethodName: "IfAInBBlacklist",
			Handler:    _ImUserService_IfAInBBlacklist_Handler,
		},
		{
			MethodName: "IfAInBFriendList",
			Handler:    _ImUserService_IfAInBFriendList_Handler,
		},
		{
			MethodName: "GetSingleConversationRecvMsgOpts",
			Handler:    _ImUserService_GetSingleConversationRecvMsgOpts_Handler,
		},
		{
			MethodName: "GetUserListFromSuperGroupWithOpt",
			Handler:    _ImUserService_GetUserListFromSuperGroupWithOpt_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _ImUserService_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im_user.proto",
}
