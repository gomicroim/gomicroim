// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/chat/chat.proto

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	// 查询会话列表
	RecentContactSession(ctx context.Context, in *CIMRecentContactSessionReq, opts ...grpc.CallOption) (*CIMRecentContactSessionRsp, error)
	// 查询历史消息列表
	GetMsgList(ctx context.Context, in *CIMGetMsgListReq, opts ...grpc.CallOption) (*CIMGetMsgListRsp, error)
	// 发消息
	Send(ctx context.Context, in *CIMMsgData, opts ...grpc.CallOption) (*CIMMsgDataAck, error)
	// 已读消息回执
	MsgReadAck(ctx context.Context, in *CIMMsgDataReadReq, opts ...grpc.CallOption) (*CIMMsgDataReadRsp, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) RecentContactSession(ctx context.Context, in *CIMRecentContactSessionReq, opts ...grpc.CallOption) (*CIMRecentContactSessionRsp, error) {
	out := new(CIMRecentContactSessionRsp)
	err := c.cc.Invoke(ctx, "/coffeechat.Chat/RecentContactSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetMsgList(ctx context.Context, in *CIMGetMsgListReq, opts ...grpc.CallOption) (*CIMGetMsgListRsp, error) {
	out := new(CIMGetMsgListRsp)
	err := c.cc.Invoke(ctx, "/coffeechat.Chat/GetMsgList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Send(ctx context.Context, in *CIMMsgData, opts ...grpc.CallOption) (*CIMMsgDataAck, error) {
	out := new(CIMMsgDataAck)
	err := c.cc.Invoke(ctx, "/coffeechat.Chat/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) MsgReadAck(ctx context.Context, in *CIMMsgDataReadReq, opts ...grpc.CallOption) (*CIMMsgDataReadRsp, error) {
	out := new(CIMMsgDataReadRsp)
	err := c.cc.Invoke(ctx, "/coffeechat.Chat/MsgReadAck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	// 查询会话列表
	RecentContactSession(context.Context, *CIMRecentContactSessionReq) (*CIMRecentContactSessionRsp, error)
	// 查询历史消息列表
	GetMsgList(context.Context, *CIMGetMsgListReq) (*CIMGetMsgListRsp, error)
	// 发消息
	Send(context.Context, *CIMMsgData) (*CIMMsgDataAck, error)
	// 已读消息回执
	MsgReadAck(context.Context, *CIMMsgDataReadReq) (*CIMMsgDataReadRsp, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) RecentContactSession(context.Context, *CIMRecentContactSessionReq) (*CIMRecentContactSessionRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecentContactSession not implemented")
}
func (UnimplementedChatServer) GetMsgList(context.Context, *CIMGetMsgListReq) (*CIMGetMsgListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMsgList not implemented")
}
func (UnimplementedChatServer) Send(context.Context, *CIMMsgData) (*CIMMsgDataAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedChatServer) MsgReadAck(context.Context, *CIMMsgDataReadReq) (*CIMMsgDataReadRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgReadAck not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_RecentContactSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIMRecentContactSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).RecentContactSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coffeechat.Chat/RecentContactSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).RecentContactSession(ctx, req.(*CIMRecentContactSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetMsgList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIMGetMsgListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetMsgList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coffeechat.Chat/GetMsgList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetMsgList(ctx, req.(*CIMGetMsgListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIMMsgData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coffeechat.Chat/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Send(ctx, req.(*CIMMsgData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_MsgReadAck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIMMsgDataReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).MsgReadAck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coffeechat.Chat/MsgReadAck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).MsgReadAck(ctx, req.(*CIMMsgDataReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coffeechat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecentContactSession",
			Handler:    _Chat_RecentContactSession_Handler,
		},
		{
			MethodName: "GetMsgList",
			Handler:    _Chat_GetMsgList_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Chat_Send_Handler,
		},
		{
			MethodName: "MsgReadAck",
			Handler:    _Chat_MsgReadAck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/chat/chat.proto",
}