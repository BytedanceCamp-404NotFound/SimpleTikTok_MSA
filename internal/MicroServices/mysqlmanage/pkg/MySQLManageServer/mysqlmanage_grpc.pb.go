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
	GetMySQL(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MySQLResponse, error)
}

type mySQLManageServerClient struct {
	cc grpc.ClientConnInterface
}

func NewMySQLManageServerClient(cc grpc.ClientConnInterface) MySQLManageServerClient {
	return &mySQLManageServerClient{cc}
}

func (c *mySQLManageServerClient) GetMySQL(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MySQLResponse, error) {
	out := new(MySQLResponse)
	err := c.cc.Invoke(ctx, "/MySQLManageServer.MySQLManageServer/getMySQL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MySQLManageServerServer is the server API for MySQLManageServer service.
// All implementations must embed UnimplementedMySQLManageServerServer
// for forward compatibility
type MySQLManageServerServer interface {
	GetMySQL(context.Context, *IdRequest) (*MySQLResponse, error)
	mustEmbedUnimplementedMySQLManageServerServer()
}

// UnimplementedMySQLManageServerServer must be embedded to have forward compatible implementations.
type UnimplementedMySQLManageServerServer struct {
}

func (UnimplementedMySQLManageServerServer) GetMySQL(context.Context, *IdRequest) (*MySQLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMySQL not implemented")
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

func _MySQLManageServer_GetMySQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySQLManageServerServer).GetMySQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MySQLManageServer.MySQLManageServer/getMySQL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySQLManageServerServer).GetMySQL(ctx, req.(*IdRequest))
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
			MethodName: "getMySQL",
			Handler:    _MySQLManageServer_GetMySQL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mysqlmanage.proto",
}