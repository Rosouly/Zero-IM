// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: usercenter.proto

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
	UsercenterService_GeTUserByUsername_FullMethodName = "/usercenter.usercenterService/GeTUserByUsername"
	UsercenterService_InsertUser_FullMethodName        = "/usercenter.usercenterService/InsertUser"
	UsercenterService_LoginById_FullMethodName         = "/usercenter.usercenterService/LoginById"
)

// UsercenterServiceClient is the client API for UsercenterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsercenterServiceClient interface {
	GeTUserByUsername(ctx context.Context, in *GeTUserByUsernameReq, opts ...grpc.CallOption) (*GeTUserResp, error)
	InsertUser(ctx context.Context, in *InsertUserReq, opts ...grpc.CallOption) (*InsertUserResp, error)
	LoginById(ctx context.Context, in *LoginByIdReq, opts ...grpc.CallOption) (*LoginByIdResp, error)
}

type usercenterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsercenterServiceClient(cc grpc.ClientConnInterface) UsercenterServiceClient {
	return &usercenterServiceClient{cc}
}

func (c *usercenterServiceClient) GeTUserByUsername(ctx context.Context, in *GeTUserByUsernameReq, opts ...grpc.CallOption) (*GeTUserResp, error) {
	out := new(GeTUserResp)
	err := c.cc.Invoke(ctx, UsercenterService_GeTUserByUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterServiceClient) InsertUser(ctx context.Context, in *InsertUserReq, opts ...grpc.CallOption) (*InsertUserResp, error) {
	out := new(InsertUserResp)
	err := c.cc.Invoke(ctx, UsercenterService_InsertUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterServiceClient) LoginById(ctx context.Context, in *LoginByIdReq, opts ...grpc.CallOption) (*LoginByIdResp, error) {
	out := new(LoginByIdResp)
	err := c.cc.Invoke(ctx, UsercenterService_LoginById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsercenterServiceServer is the server API for UsercenterService service.
// All implementations must embed UnimplementedUsercenterServiceServer
// for forward compatibility
type UsercenterServiceServer interface {
	GeTUserByUsername(context.Context, *GeTUserByUsernameReq) (*GeTUserResp, error)
	InsertUser(context.Context, *InsertUserReq) (*InsertUserResp, error)
	LoginById(context.Context, *LoginByIdReq) (*LoginByIdResp, error)
	mustEmbedUnimplementedUsercenterServiceServer()
}

// UnimplementedUsercenterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsercenterServiceServer struct {
}

func (UnimplementedUsercenterServiceServer) GeTUserByUsername(context.Context, *GeTUserByUsernameReq) (*GeTUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeTUserByUsername not implemented")
}
func (UnimplementedUsercenterServiceServer) InsertUser(context.Context, *InsertUserReq) (*InsertUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertUser not implemented")
}
func (UnimplementedUsercenterServiceServer) LoginById(context.Context, *LoginByIdReq) (*LoginByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginById not implemented")
}
func (UnimplementedUsercenterServiceServer) mustEmbedUnimplementedUsercenterServiceServer() {}

// UnsafeUsercenterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsercenterServiceServer will
// result in compilation errors.
type UnsafeUsercenterServiceServer interface {
	mustEmbedUnimplementedUsercenterServiceServer()
}

func RegisterUsercenterServiceServer(s grpc.ServiceRegistrar, srv UsercenterServiceServer) {
	s.RegisterService(&UsercenterService_ServiceDesc, srv)
}

func _UsercenterService_GeTUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeTUserByUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServiceServer).GeTUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsercenterService_GeTUserByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServiceServer).GeTUserByUsername(ctx, req.(*GeTUserByUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsercenterService_InsertUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServiceServer).InsertUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsercenterService_InsertUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServiceServer).InsertUser(ctx, req.(*InsertUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsercenterService_LoginById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServiceServer).LoginById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsercenterService_LoginById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServiceServer).LoginById(ctx, req.(*LoginByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UsercenterService_ServiceDesc is the grpc.ServiceDesc for UsercenterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsercenterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usercenter.usercenterService",
	HandlerType: (*UsercenterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeTUserByUsername",
			Handler:    _UsercenterService_GeTUserByUsername_Handler,
		},
		{
			MethodName: "InsertUser",
			Handler:    _UsercenterService_InsertUser_Handler,
		},
		{
			MethodName: "LoginById",
			Handler:    _UsercenterService_LoginById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usercenter.proto",
}