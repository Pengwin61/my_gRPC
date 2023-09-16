// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: api/proto/chat.proto

package chat

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error)
	Hello(ctx context.Context, opts ...grpc.CallOption) (ChatService_HelloClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/chat.ChatService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSendMessageClient{stream}
	return x, nil
}

type ChatService_SendMessageClient interface {
	Send(*MessageRequest) error
	Recv() (*MessageResponse, error)
	grpc.ClientStream
}

type chatServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceSendMessageClient) Send(m *MessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSendMessageClient) Recv() (*MessageResponse, error) {
	m := new(MessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) Hello(ctx context.Context, opts ...grpc.CallOption) (ChatService_HelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], "/chat.ChatService/Hello", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceHelloClient{stream}
	return x, nil
}

type ChatService_HelloClient interface {
	Send(*MsgHelloRequest) error
	Recv() (*MsgHelloResponse, error)
	grpc.ClientStream
}

type chatServiceHelloClient struct {
	grpc.ClientStream
}

func (x *chatServiceHelloClient) Send(m *MsgHelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceHelloClient) Recv() (*MsgHelloResponse, error) {
	m := new(MsgHelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations should embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	SendMessage(ChatService_SendMessageServer) error
	Hello(ChatService_HelloServer) error
}

// UnimplementedChatServiceServer should be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) SendMessage(ChatService_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) Hello(ChatService_HelloServer) error {
	return status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendMessage(&chatServiceSendMessageServer{stream})
}

type ChatService_SendMessageServer interface {
	Send(*MessageResponse) error
	Recv() (*MessageRequest, error)
	grpc.ServerStream
}

type chatServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceSendMessageServer) Send(m *MessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSendMessageServer) Recv() (*MessageRequest, error) {
	m := new(MessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_Hello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Hello(&chatServiceHelloServer{stream})
}

type ChatService_HelloServer interface {
	Send(*MsgHelloResponse) error
	Recv() (*MsgHelloRequest, error)
	grpc.ServerStream
}

type chatServiceHelloServer struct {
	grpc.ServerStream
}

func (x *chatServiceHelloServer) Send(m *MsgHelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceHelloServer) Recv() (*MsgHelloRequest, error) {
	m := new(MsgHelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _ChatService_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Hello",
			Handler:       _ChatService_Hello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/proto/chat.proto",
}
