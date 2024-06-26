// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: user.proto

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
	User_GenerateToken_FullMethodName    = "/user.user/GenerateToken"
	User_Register_FullMethodName         = "/user.user/Register"
	User_Login_FullMethodName            = "/user.user/Login"
	User_UserDetail_FullMethodName       = "/user.user/UserDetail"
	User_UpdateUserDetail_FullMethodName = "/user.user/UpdateUserDetail"
	User_UpdateUserMoney_FullMethodName  = "/user.user/UpdateUserMoney"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	UserDetail(ctx context.Context, in *UserDetailReq, opts ...grpc.CallOption) (*UserDetailResp, error)
	UpdateUserDetail(ctx context.Context, in *UpdateUserDetailReq, opts ...grpc.CallOption) (*UpdateUserDetailResp, error)
	UpdateUserMoney(ctx context.Context, in *UpdateUserMoneyReq, opts ...grpc.CallOption) (*UpdateUserMoneyResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	out := new(GenerateTokenResp)
	err := c.cc.Invoke(ctx, User_GenerateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, User_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, User_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserDetail(ctx context.Context, in *UserDetailReq, opts ...grpc.CallOption) (*UserDetailResp, error) {
	out := new(UserDetailResp)
	err := c.cc.Invoke(ctx, User_UserDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserDetail(ctx context.Context, in *UpdateUserDetailReq, opts ...grpc.CallOption) (*UpdateUserDetailResp, error) {
	out := new(UpdateUserDetailResp)
	err := c.cc.Invoke(ctx, User_UpdateUserDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserMoney(ctx context.Context, in *UpdateUserMoneyReq, opts ...grpc.CallOption) (*UpdateUserMoneyResp, error) {
	out := new(UpdateUserMoneyResp)
	err := c.cc.Invoke(ctx, User_UpdateUserMoney_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error)
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	UserDetail(context.Context, *UserDetailReq) (*UserDetailResp, error)
	UpdateUserDetail(context.Context, *UpdateUserDetailReq) (*UpdateUserDetailResp, error)
	UpdateUserMoney(context.Context, *UpdateUserMoneyReq) (*UpdateUserMoneyResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedUserServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) UserDetail(context.Context, *UserDetailReq) (*UserDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetail not implemented")
}
func (UnimplementedUserServer) UpdateUserDetail(context.Context, *UpdateUserDetailReq) (*UpdateUserDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserDetail not implemented")
}
func (UnimplementedUserServer) UpdateUserMoney(context.Context, *UpdateUserMoneyReq) (*UpdateUserMoneyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMoney not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GenerateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GenerateToken(ctx, req.(*GenerateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserDetail(ctx, req.(*UserDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateUserDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserDetail(ctx, req.(*UpdateUserDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserMoneyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateUserMoney_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserMoney(ctx, req.(*UpdateUserMoneyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.user",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateToken",
			Handler:    _User_GenerateToken_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "UserDetail",
			Handler:    _User_UserDetail_Handler,
		},
		{
			MethodName: "UpdateUserDetail",
			Handler:    _User_UpdateUserDetail_Handler,
		},
		{
			MethodName: "UpdateUserMoney",
			Handler:    _User_UpdateUserMoney_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
