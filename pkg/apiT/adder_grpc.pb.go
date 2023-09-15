// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: api/proto/adder.proto

package apiT

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

// AdderClient is the client API for Adder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdderClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Sub(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubResponse, error)
	Multi(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type adderClient struct {
	cc grpc.ClientConnInterface
}

func NewAdderClient(cc grpc.ClientConnInterface) AdderClient {
	return &adderClient{cc}
}

func (c *adderClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/apiT.Adder/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adderClient) Sub(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubResponse, error) {
	out := new(SubResponse)
	err := c.cc.Invoke(ctx, "/apiT.Adder/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adderClient) Multi(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/apiT.Adder/Multi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdderServer is the server API for Adder service.
// All implementations should embed UnimplementedAdderServer
// for forward compatibility
type AdderServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Sub(context.Context, *SubRequest) (*SubResponse, error)
	Multi(context.Context, *AddRequest) (*AddResponse, error)
}

// UnimplementedAdderServer should be embedded to have forward compatible implementations.
type UnimplementedAdderServer struct {
}

func (UnimplementedAdderServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAdderServer) Sub(context.Context, *SubRequest) (*SubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}
func (UnimplementedAdderServer) Multi(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multi not implemented")
}

// UnsafeAdderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdderServer will
// result in compilation errors.
type UnsafeAdderServer interface {
	mustEmbedUnimplementedAdderServer()
}

func RegisterAdderServer(s grpc.ServiceRegistrar, srv AdderServer) {
	s.RegisterService(&Adder_ServiceDesc, srv)
}

func _Adder_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiT.Adder/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adder_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiT.Adder/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServer).Sub(ctx, req.(*SubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adder_Multi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServer).Multi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiT.Adder/Multi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServer).Multi(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Adder_ServiceDesc is the grpc.ServiceDesc for Adder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Adder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiT.Adder",
	HandlerType: (*AdderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Adder_Add_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _Adder_Sub_Handler,
		},
		{
			MethodName: "Multi",
			Handler:    _Adder_Multi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/adder.proto",
}
