// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationChatGetMsgList = "/apichat.Chat/GetMsgList"
const OperationChatGetRecentContactSession = "/apichat.Chat/GetRecentContactSession"
const OperationChatGetSyncMessage = "/apichat.Chat/GetSyncMessage"
const OperationChatMsgReadAck = "/apichat.Chat/MsgReadAck"
const OperationChatSendMsg = "/apichat.Chat/SendMsg"

type ChatHTTPServer interface {
	GetMsgList(context.Context, *GetMsgListRequest) (*GetMsgListReply, error)
	GetRecentContactSession(context.Context, *GetRecentSessionRequest) (*GetRecentSessionReply, error)
	GetSyncMessage(context.Context, *SyncMessageRequest) (*SyncMessageReply, error)
	MsgReadAck(context.Context, *MsgReadAckRequest) (*MsgReadAckReply, error)
	SendMsg(context.Context, *SendMsgRequest) (*SendMsgReply, error)
}

func RegisterChatHTTPServer(s *http.Server, srv ChatHTTPServer) {
	r := s.Route("/")
	r.POST("/chat/msg/send", _Chat_SendMsg0_HTTP_Handler(srv))
	r.GET("/chat/msg/sync", _Chat_GetSyncMessage0_HTTP_Handler(srv))
	r.GET("/chat/session/list", _Chat_GetRecentContactSession0_HTTP_Handler(srv))
	r.GET("/chat/msg/list", _Chat_GetMsgList0_HTTP_Handler(srv))
	r.POST("/chat/msg/read", _Chat_MsgReadAck0_HTTP_Handler(srv))
}

func _Chat_SendMsg0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendMsgRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatSendMsg)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendMsg(ctx, req.(*SendMsgRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendMsgReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_GetSyncMessage0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SyncMessageRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatGetSyncMessage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSyncMessage(ctx, req.(*SyncMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SyncMessageReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_GetRecentContactSession0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRecentSessionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatGetRecentContactSession)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRecentContactSession(ctx, req.(*GetRecentSessionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRecentSessionReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_GetMsgList0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMsgListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatGetMsgList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMsgList(ctx, req.(*GetMsgListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMsgListReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_MsgReadAck0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MsgReadAckRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatMsgReadAck)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MsgReadAck(ctx, req.(*MsgReadAckRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MsgReadAckReply)
		return ctx.Result(200, reply)
	}
}

type ChatHTTPClient interface {
	GetMsgList(ctx context.Context, req *GetMsgListRequest, opts ...http.CallOption) (rsp *GetMsgListReply, err error)
	GetRecentContactSession(ctx context.Context, req *GetRecentSessionRequest, opts ...http.CallOption) (rsp *GetRecentSessionReply, err error)
	GetSyncMessage(ctx context.Context, req *SyncMessageRequest, opts ...http.CallOption) (rsp *SyncMessageReply, err error)
	MsgReadAck(ctx context.Context, req *MsgReadAckRequest, opts ...http.CallOption) (rsp *MsgReadAckReply, err error)
	SendMsg(ctx context.Context, req *SendMsgRequest, opts ...http.CallOption) (rsp *SendMsgReply, err error)
}

type ChatHTTPClientImpl struct {
	cc *http.Client
}

func NewChatHTTPClient(client *http.Client) ChatHTTPClient {
	return &ChatHTTPClientImpl{client}
}

func (c *ChatHTTPClientImpl) GetMsgList(ctx context.Context, in *GetMsgListRequest, opts ...http.CallOption) (*GetMsgListReply, error) {
	var out GetMsgListReply
	pattern := "/chat/msg/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatGetMsgList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) GetRecentContactSession(ctx context.Context, in *GetRecentSessionRequest, opts ...http.CallOption) (*GetRecentSessionReply, error) {
	var out GetRecentSessionReply
	pattern := "/chat/session/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatGetRecentContactSession))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) GetSyncMessage(ctx context.Context, in *SyncMessageRequest, opts ...http.CallOption) (*SyncMessageReply, error) {
	var out SyncMessageReply
	pattern := "/chat/msg/sync"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatGetSyncMessage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) MsgReadAck(ctx context.Context, in *MsgReadAckRequest, opts ...http.CallOption) (*MsgReadAckReply, error) {
	var out MsgReadAckReply
	pattern := "/chat/msg/read"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatMsgReadAck))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) SendMsg(ctx context.Context, in *SendMsgRequest, opts ...http.CallOption) (*SendMsgReply, error) {
	var out SendMsgReply
	pattern := "/chat/msg/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatSendMsg))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}