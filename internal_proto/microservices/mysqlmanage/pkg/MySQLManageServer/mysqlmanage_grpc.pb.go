// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/mysqlmanage.proto

package MySQLManageServer

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

// MySQLManageServerClient is the client API for MySQLManageServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MySQLManageServerClient interface {
	// 1
	CommentGetUserByUserId(ctx context.Context, in *CommentGetUserByUserIdRequest, opts ...grpc.CallOption) (*CommentGetUserByUserIdResponse, error)
	// 2
	FavoriteVideoNum(ctx context.Context, in *FavoriteVideoNumRequest, opts ...grpc.CallOption) (*FavoriteVideoNumResponse, error)
}

type mySQLManageServerClient struct {
	cc grpc.ClientConnInterface
}

func NewMySQLManageServerClient(cc grpc.ClientConnInterface) MySQLManageServerClient {
	return &mySQLManageServerClient{cc}
}

func (c *mySQLManageServerClient) CommentGetUserByUserId(ctx context.Context, in *CommentGetUserByUserIdRequest, opts ...grpc.CallOption) (*CommentGetUserByUserIdResponse, error) {
	out := new(CommentGetUserByUserIdResponse)
	err := c.cc.Invoke(ctx, "/MySQLManageServer.MySQLManageServer/CommentGetUserByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySQLManageServerClient) FavoriteVideoNum(ctx context.Context, in *FavoriteVideoNumRequest, opts ...grpc.CallOption) (*FavoriteVideoNumResponse, error) {
	out := new(FavoriteVideoNumResponse)
	err := c.cc.Invoke(ctx, "/MySQLManageServer.MySQLManageServer/FavoriteVideoNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MySQLManageServerServer is the server API for MySQLManageServer service.
// All implementations must embed UnimplementedMySQLManageServerServer
// for forward compatibility
type MySQLManageServerServer interface {
	// 1
	CommentGetUserByUserId(context.Context, *CommentGetUserByUserIdRequest) (*CommentGetUserByUserIdResponse, error)
	// 2
	FavoriteVideoNum(context.Context, *FavoriteVideoNumRequest) (*FavoriteVideoNumResponse, error)
	mustEmbedUnimplementedMySQLManageServerServer()
}

// UnimplementedMySQLManageServerServer must be embedded to have forward compatible implementations.
type UnimplementedMySQLManageServerServer struct {
}

func (UnimplementedMySQLManageServerServer) CommentGetUserByUserId(context.Context, *CommentGetUserByUserIdRequest) (*CommentGetUserByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentGetUserByUserId not implemented")
}
func (UnimplementedMySQLManageServerServer) FavoriteVideoNum(context.Context, *FavoriteVideoNumRequest) (*FavoriteVideoNumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteVideoNum not implemented")
}
func (UnimplementedMySQLManageServerServer) mustEmbedUnimplementedMySQLManageServerServer() {}

// UnsafeMySQLManageServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MySQLManageServerServer will
// result in compilation errors.
type UnsafeMySQLManageServerServer interface {
	mustEmbedUnimplementedMySQLManageServerServer()
}

func RegisterMySQLManageServerServer(s grpc.ServiceRegistrar, srv MySQLManageServerServer) {
	s.RegisterService(&MySQLManageServer_ServiceDesc, srv)
}

func _MySQLManageServer_CommentGetUserByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentGetUserByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySQLManageServerServer).CommentGetUserByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MySQLManageServer.MySQLManageServer/CommentGetUserByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySQLManageServerServer).CommentGetUserByUserId(ctx, req.(*CommentGetUserByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MySQLManageServer_FavoriteVideoNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteVideoNumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySQLManageServerServer).FavoriteVideoNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MySQLManageServer.MySQLManageServer/FavoriteVideoNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySQLManageServerServer).FavoriteVideoNum(ctx, req.(*FavoriteVideoNumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MySQLManageServer_ServiceDesc is the grpc.ServiceDesc for MySQLManageServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MySQLManageServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MySQLManageServer.MySQLManageServer",
	HandlerType: (*MySQLManageServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommentGetUserByUserId",
			Handler:    _MySQLManageServer_CommentGetUserByUserId_Handler,
		},
		{
			MethodName: "FavoriteVideoNum",
			Handler:    _MySQLManageServer_FavoriteVideoNum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mysqlmanage.proto",
}