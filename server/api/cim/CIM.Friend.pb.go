// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: CIM.Friend.proto

package cim

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 查询系统推荐的50个以内用户
type CIMFriendQueryUserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id: 		0x601
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CIMFriendQueryUserListReq) Reset() {
	*x = CIMFriendQueryUserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIM_Friend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMFriendQueryUserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMFriendQueryUserListReq) ProtoMessage() {}

func (x *CIMFriendQueryUserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_CIM_Friend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMFriendQueryUserListReq.ProtoReflect.Descriptor instead.
func (*CIMFriendQueryUserListReq) Descriptor() ([]byte, []int) {
	return file_CIM_Friend_proto_rawDescGZIP(), []int{0}
}

func (x *CIMFriendQueryUserListReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CIMFriendQueryUserListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id: 		0x602
	UserId       uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserInfoList []*CIMUserInfo `protobuf:"bytes,2,rep,name=user_info_list,json=userInfoList,proto3" json:"user_info_list,omitempty"`
}

func (x *CIMFriendQueryUserListRsp) Reset() {
	*x = CIMFriendQueryUserListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIM_Friend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMFriendQueryUserListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMFriendQueryUserListRsp) ProtoMessage() {}

func (x *CIMFriendQueryUserListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_CIM_Friend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMFriendQueryUserListRsp.ProtoReflect.Descriptor instead.
func (*CIMFriendQueryUserListRsp) Descriptor() ([]byte, []int) {
	return file_CIM_Friend_proto_rawDescGZIP(), []int{1}
}

func (x *CIMFriendQueryUserListRsp) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMFriendQueryUserListRsp) GetUserInfoList() []*CIMUserInfo {
	if x != nil {
		return x.UserInfoList
	}
	return nil
}

var File_CIM_Friend_proto protoreflect.FileDescriptor

var file_CIM_Friend_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x49, 0x4d, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x43, 0x49, 0x4d, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x1a, 0x0d,
	0x43, 0x49, 0x4d, 0x2e, 0x44, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a,
	0x19, 0x43, 0x49, 0x4d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x19, 0x43, 0x49, 0x4d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x43, 0x49, 0x4d, 0x2e, 0x44, 0x65, 0x66, 0x2e, 0x43, 0x49, 0x4d, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x22, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x66,
	0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x48, 0x03, 0x5a, 0x05, 0x2e, 0x3b, 0x63, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_CIM_Friend_proto_rawDescOnce sync.Once
	file_CIM_Friend_proto_rawDescData = file_CIM_Friend_proto_rawDesc
)

func file_CIM_Friend_proto_rawDescGZIP() []byte {
	file_CIM_Friend_proto_rawDescOnce.Do(func() {
		file_CIM_Friend_proto_rawDescData = protoimpl.X.CompressGZIP(file_CIM_Friend_proto_rawDescData)
	})
	return file_CIM_Friend_proto_rawDescData
}

var file_CIM_Friend_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_CIM_Friend_proto_goTypes = []interface{}{
	(*CIMFriendQueryUserListReq)(nil), // 0: CIM.Friend.CIMFriendQueryUserListReq
	(*CIMFriendQueryUserListRsp)(nil), // 1: CIM.Friend.CIMFriendQueryUserListRsp
	(*CIMUserInfo)(nil),               // 2: CIM.Def.CIMUserInfo
}
var file_CIM_Friend_proto_depIdxs = []int32{
	2, // 0: CIM.Friend.CIMFriendQueryUserListRsp.user_info_list:type_name -> CIM.Def.CIMUserInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CIM_Friend_proto_init() }
func file_CIM_Friend_proto_init() {
	if File_CIM_Friend_proto != nil {
		return
	}
	file_CIM_Def_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CIM_Friend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMFriendQueryUserListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_CIM_Friend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMFriendQueryUserListRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CIM_Friend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CIM_Friend_proto_goTypes,
		DependencyIndexes: file_CIM_Friend_proto_depIdxs,
		MessageInfos:      file_CIM_Friend_proto_msgTypes,
	}.Build()
	File_CIM_Friend_proto = out.File
	file_CIM_Friend_proto_rawDesc = nil
	file_CIM_Friend_proto_goTypes = nil
	file_CIM_Friend_proto_depIdxs = nil
}