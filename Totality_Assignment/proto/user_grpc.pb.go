// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: proto/user.proto

package proto

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

// UserServiceProtoClient is the client API for UserServiceProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceProtoClient interface {
	GetUserByIdProto(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUsersList(ctx context.Context, in *UsersListRequest, opts ...grpc.CallOption) (*UsersListResponse, error)
}

type userServiceProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceProtoClient(cc grpc.ClientConnInterface) UserServiceProtoClient {
	return &userServiceProtoClient{cc}
}

func (c *userServiceProtoClient) GetUserByIdProto(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/proto.UserServiceProto/GetUserByIdProto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceProtoClient) GetUsersList(ctx context.Context, in *UsersListRequest, opts ...grpc.CallOption) (*UsersListResponse, error) {
	out := new(UsersListResponse)
	err := c.cc.Invoke(ctx, "/proto.UserServiceProto/GetUsersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceProtoServer is the server API for UserServiceProto service.
// All implementations must embed UnimplementedUserServiceProtoServer
// for forward compatibility
type UserServiceProtoServer interface {
	GetUserByIdProto(context.Context, *UserRequest) (*UserResponse, error)
	GetUsersList(context.Context, *UsersListRequest) (*UsersListResponse, error)
	mustEmbedUnimplementedUserServiceProtoServer()
}

// UnimplementedUserServiceProtoServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceProtoServer struct {
}

func (UnimplementedUserServiceProtoServer) GetUserByIdProto(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByIdProto not implemented")
}
func (UnimplementedUserServiceProtoServer) GetUsersList(context.Context, *UsersListRequest) (*UsersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersList not implemented")
}
func (UnimplementedUserServiceProtoServer) mustEmbedUnimplementedUserServiceProtoServer() {}

// UnsafeUserServiceProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceProtoServer will
// result in compilation errors.
type UnsafeUserServiceProtoServer interface {
	mustEmbedUnimplementedUserServiceProtoServer()
}

func RegisterUserServiceProtoServer(s grpc.ServiceRegistrar, srv UserServiceProtoServer) {
	s.RegisterService(&UserServiceProto_ServiceDesc, srv)
}

func _UserServiceProto_GetUserByIdProto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceProtoServer).GetUserByIdProto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserServiceProto/GetUserByIdProto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceProtoServer).GetUserByIdProto(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceProto_GetUsersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceProtoServer).GetUsersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserServiceProto/GetUsersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceProtoServer).GetUsersList(ctx, req.(*UsersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServiceProto_ServiceDesc is the grpc.ServiceDesc for UserServiceProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServiceProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserServiceProto",
	HandlerType: (*UserServiceProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByIdProto",
			Handler:    _UserServiceProto_GetUserByIdProto_Handler,
		},
		{
			MethodName: "GetUsersList",
			Handler:    _UserServiceProto_GetUsersList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
