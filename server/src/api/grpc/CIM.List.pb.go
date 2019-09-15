// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CIM.List.proto

package grpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 最近聊天会话列表请求
type CIMRecentContactSessionReq struct {
	// cmd id:		0x0201
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LatestUpdateTime     uint32   `protobuf:"varint,2,opt,name=latest_update_time,json=latestUpdateTime,proto3" json:"latest_update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIMRecentContactSessionReq) Reset()         { *m = CIMRecentContactSessionReq{} }
func (m *CIMRecentContactSessionReq) String() string { return proto.CompactTextString(m) }
func (*CIMRecentContactSessionReq) ProtoMessage()    {}
func (*CIMRecentContactSessionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_91d1179a25bb2ce6, []int{0}
}

func (m *CIMRecentContactSessionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMRecentContactSessionReq.Unmarshal(m, b)
}
func (m *CIMRecentContactSessionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMRecentContactSessionReq.Marshal(b, m, deterministic)
}
func (m *CIMRecentContactSessionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMRecentContactSessionReq.Merge(m, src)
}
func (m *CIMRecentContactSessionReq) XXX_Size() int {
	return xxx_messageInfo_CIMRecentContactSessionReq.Size(m)
}
func (m *CIMRecentContactSessionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMRecentContactSessionReq.DiscardUnknown(m)
}

var xxx_messageInfo_CIMRecentContactSessionReq proto.InternalMessageInfo

func (m *CIMRecentContactSessionReq) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMRecentContactSessionReq) GetLatestUpdateTime() uint32 {
	if m != nil {
		return m.LatestUpdateTime
	}
	return 0
}

type CIMRecentContactSessionRsp struct {
	// cmd id:		0x0202
	UserId               uint64                   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UnreadCounts         uint32                   `protobuf:"varint,2,opt,name=unread_counts,json=unreadCounts,proto3" json:"unread_counts,omitempty"`
	ContactSessionList   []*CIMContactSessionInfo `protobuf:"bytes,3,rep,name=contact_session_list,json=contactSessionList,proto3" json:"contact_session_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CIMRecentContactSessionRsp) Reset()         { *m = CIMRecentContactSessionRsp{} }
func (m *CIMRecentContactSessionRsp) String() string { return proto.CompactTextString(m) }
func (*CIMRecentContactSessionRsp) ProtoMessage()    {}
func (*CIMRecentContactSessionRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_91d1179a25bb2ce6, []int{1}
}

func (m *CIMRecentContactSessionRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMRecentContactSessionRsp.Unmarshal(m, b)
}
func (m *CIMRecentContactSessionRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMRecentContactSessionRsp.Marshal(b, m, deterministic)
}
func (m *CIMRecentContactSessionRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMRecentContactSessionRsp.Merge(m, src)
}
func (m *CIMRecentContactSessionRsp) XXX_Size() int {
	return xxx_messageInfo_CIMRecentContactSessionRsp.Size(m)
}
func (m *CIMRecentContactSessionRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMRecentContactSessionRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CIMRecentContactSessionRsp proto.InternalMessageInfo

func (m *CIMRecentContactSessionRsp) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMRecentContactSessionRsp) GetUnreadCounts() uint32 {
	if m != nil {
		return m.UnreadCounts
	}
	return 0
}

func (m *CIMRecentContactSessionRsp) GetContactSessionList() []*CIMContactSessionInfo {
	if m != nil {
		return m.ContactSessionList
	}
	return nil
}

// 历史离线聊天消息请求
type CIMGetMsgListReq struct {
	// cmd id:		0x0205
	UserId      uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionType CIMSessionType `protobuf:"varint,2,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	SessionId   uint64         `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	//   uint64 from_time_stamp = 4; // 起始时间点，单位：毫秒
	//   uint64 end_time_stamp = 5;  // 结束时间点，单位：毫秒
	EndMsgId             uint64   `protobuf:"varint,4,opt,name=end_msg_id,json=endMsgId,proto3" json:"end_msg_id,omitempty"`
	LimitCount           uint32   `protobuf:"varint,6,opt,name=limit_count,json=limitCount,proto3" json:"limit_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIMGetMsgListReq) Reset()         { *m = CIMGetMsgListReq{} }
func (m *CIMGetMsgListReq) String() string { return proto.CompactTextString(m) }
func (*CIMGetMsgListReq) ProtoMessage()    {}
func (*CIMGetMsgListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_91d1179a25bb2ce6, []int{2}
}

func (m *CIMGetMsgListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMGetMsgListReq.Unmarshal(m, b)
}
func (m *CIMGetMsgListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMGetMsgListReq.Marshal(b, m, deterministic)
}
func (m *CIMGetMsgListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMGetMsgListReq.Merge(m, src)
}
func (m *CIMGetMsgListReq) XXX_Size() int {
	return xxx_messageInfo_CIMGetMsgListReq.Size(m)
}
func (m *CIMGetMsgListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMGetMsgListReq.DiscardUnknown(m)
}

var xxx_messageInfo_CIMGetMsgListReq proto.InternalMessageInfo

func (m *CIMGetMsgListReq) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMGetMsgListReq) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (m *CIMGetMsgListReq) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *CIMGetMsgListReq) GetEndMsgId() uint64 {
	if m != nil {
		return m.EndMsgId
	}
	return 0
}

func (m *CIMGetMsgListReq) GetLimitCount() uint32 {
	if m != nil {
		return m.LimitCount
	}
	return 0
}

//对于群而言，如果消息数目返回的数值小于请求的cnt,则表示群的消息能拉取的到头了，更早的消息没有权限拉取。
//如果limit_count 和 msg_list.count 不一致，说明服务器消息有缺失，需要
//客户端做一个缺失标记，避免下次再次拉取。
type CIMGetMsgListRsp struct {
	// cmd id:		0x0206
	UserId      uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionType CIMSessionType `protobuf:"varint,2,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	SessionId   uint64         `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	//   uint64 from_time_stamp = 4;     // 起始时间点，单位：毫秒
	//   uint64 end_time_stamp = 5;      // 结束时间点，单位：毫秒
	MsgList              []*CIMMsgInfo `protobuf:"bytes,6,rep,name=msg_list,json=msgList,proto3" json:"msg_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CIMGetMsgListRsp) Reset()         { *m = CIMGetMsgListRsp{} }
func (m *CIMGetMsgListRsp) String() string { return proto.CompactTextString(m) }
func (*CIMGetMsgListRsp) ProtoMessage()    {}
func (*CIMGetMsgListRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_91d1179a25bb2ce6, []int{3}
}

func (m *CIMGetMsgListRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMGetMsgListRsp.Unmarshal(m, b)
}
func (m *CIMGetMsgListRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMGetMsgListRsp.Marshal(b, m, deterministic)
}
func (m *CIMGetMsgListRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMGetMsgListRsp.Merge(m, src)
}
func (m *CIMGetMsgListRsp) XXX_Size() int {
	return xxx_messageInfo_CIMGetMsgListRsp.Size(m)
}
func (m *CIMGetMsgListRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMGetMsgListRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CIMGetMsgListRsp proto.InternalMessageInfo

func (m *CIMGetMsgListRsp) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMGetMsgListRsp) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (m *CIMGetMsgListRsp) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *CIMGetMsgListRsp) GetMsgList() []*CIMMsgInfo {
	if m != nil {
		return m.MsgList
	}
	return nil
}

func init() {
	proto.RegisterType((*CIMRecentContactSessionReq)(nil), "CIM.List.CIMRecentContactSessionReq")
	proto.RegisterType((*CIMRecentContactSessionRsp)(nil), "CIM.List.CIMRecentContactSessionRsp")
	proto.RegisterType((*CIMGetMsgListReq)(nil), "CIM.List.CIMGetMsgListReq")
	proto.RegisterType((*CIMGetMsgListRsp)(nil), "CIM.List.CIMGetMsgListRsp")
}

func init() { proto.RegisterFile("CIM.List.proto", fileDescriptor_91d1179a25bb2ce6) }

var fileDescriptor_91d1179a25bb2ce6 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x14, 0x54, 0x48, 0x95, 0x2d, 0xaf, 0xdb, 0xd5, 0xca, 0x20, 0x35, 0x5a, 0xf1, 0x51, 0xc2, 0xa5,
	0x07, 0x94, 0xc3, 0x72, 0xe3, 0x48, 0x90, 0xc0, 0x12, 0x91, 0x50, 0x28, 0x17, 0x2e, 0x51, 0x6a,
	0xbf, 0x04, 0x4b, 0x8d, 0x1d, 0x62, 0xe7, 0xd0, 0xff, 0xc4, 0x99, 0x7f, 0xc0, 0xff, 0x42, 0xb6,
	0x13, 0xd1, 0x0a, 0xb5, 0xb7, 0xbd, 0xf9, 0xcd, 0xc8, 0x33, 0x9e, 0x79, 0x86, 0x9b, 0x8c, 0xe6,
	0xe9, 0x67, 0xa1, 0x4d, 0xda, 0xf5, 0xca, 0x28, 0x32, 0x9f, 0xe6, 0xbb, 0xa5, 0x3d, 0x7d, 0xc0,
	0xda, 0x13, 0x09, 0x83, 0xbb, 0x8c, 0xe6, 0x05, 0x32, 0x94, 0x26, 0x53, 0xd2, 0x54, 0xcc, 0x7c,
	0x45, 0xad, 0x85, 0x92, 0x05, 0xfe, 0x24, 0x2b, 0xb8, 0x1a, 0x34, 0xf6, 0xa5, 0xe0, 0x71, 0xb0,
	0x0e, 0x36, 0xb3, 0x22, 0xb2, 0x23, 0xe5, 0xe4, 0x0d, 0x90, 0x7d, 0x65, 0x50, 0x9b, 0x72, 0xe8,
	0x78, 0x65, 0xb0, 0x34, 0xa2, 0xc5, 0xf8, 0xd1, 0x3a, 0xd8, 0x2c, 0x8b, 0x5b, 0xcf, 0x7c, 0x73,
	0xc4, 0x56, 0xb4, 0x98, 0xfc, 0x0a, 0xce, 0xbb, 0xe8, 0xee, 0xbc, 0xcb, 0x6b, 0x58, 0x0e, 0xb2,
	0xc7, 0x8a, 0x97, 0x4c, 0x0d, 0xd2, 0xe8, 0xd1, 0xe0, 0xda, 0x83, 0x99, 0xc3, 0xc8, 0x17, 0x78,
	0xca, 0xbc, 0x64, 0xa9, 0xbd, 0x66, 0xb9, 0x17, 0xda, 0xc4, 0xe1, 0x3a, 0xdc, 0x2c, 0xee, 0x5f,
	0xa4, 0x53, 0xde, 0x8c, 0xe6, 0xa7, 0xd6, 0x54, 0xd6, 0xaa, 0x20, 0xec, 0x04, 0xb3, 0x15, 0x25,
	0x7f, 0x02, 0xb8, 0xcd, 0x68, 0xfe, 0x11, 0x4d, 0xae, 0x1b, 0x8b, 0x5c, 0xac, 0xe2, 0x1d, 0x5c,
	0x4f, 0xbe, 0xe6, 0xd0, 0xf9, 0x12, 0x6e, 0xee, 0x57, 0xc7, 0xbe, 0xa3, 0xf8, 0xf6, 0xd0, 0x61,
	0xb1, 0xd0, 0xff, 0x06, 0xf2, 0x1c, 0x60, 0xba, 0x2b, 0x78, 0x1c, 0x3a, 0xdd, 0xc7, 0x23, 0x42,
	0x39, 0x79, 0x06, 0x80, 0x92, 0x97, 0xad, 0x6e, 0x2c, 0x3d, 0x73, 0xf4, 0x1c, 0x25, 0xcf, 0x75,
	0x43, 0x39, 0x79, 0x09, 0x8b, 0xbd, 0x68, 0x85, 0xf1, 0xe5, 0xc4, 0x91, 0xeb, 0x06, 0x1c, 0xe4,
	0xaa, 0x49, 0x7e, 0xff, 0x97, 0xe3, 0x52, 0xd9, 0x0f, 0x98, 0x23, 0x85, 0xb9, 0xcd, 0xe0, 0xd6,
	0x12, 0xb9, 0xb5, 0x3c, 0x39, 0x96, 0xb5, 0x71, 0xec, 0x2e, 0xae, 0x5a, 0xff, 0xcc, 0xf7, 0xaf,
	0x60, 0xc5, 0x54, 0x9b, 0x32, 0x55, 0xd7, 0x88, 0xec, 0x47, 0x35, 0xfe, 0xe2, 0xdd, 0x50, 0x7f,
	0x0a, 0xbf, 0xcf, 0x9a, 0xbe, 0x63, 0xbb, 0xc8, 0x21, 0x6f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x90, 0x75, 0x21, 0xb8, 0xe9, 0x02, 0x00, 0x00,
}