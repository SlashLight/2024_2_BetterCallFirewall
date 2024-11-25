// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/community.proto

package community_api

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

type CheckAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      uint32 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	CommunityID uint32 `protobuf:"varint,2,opt,name=CommunityID,proto3" json:"CommunityID,omitempty"`
}

func (x *CheckAccessRequest) Reset() {
	*x = CheckAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_community_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccessRequest) ProtoMessage() {}

func (x *CheckAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_community_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccessRequest.ProtoReflect.Descriptor instead.
func (*CheckAccessRequest) Descriptor() ([]byte, []int) {
	return file_proto_community_proto_rawDescGZIP(), []int{0}
}

func (x *CheckAccessRequest) GetUserID() uint32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CheckAccessRequest) GetCommunityID() uint32 {
	if x != nil {
		return x.CommunityID
	}
	return 0
}

type CheckAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access bool `protobuf:"varint,1,opt,name=Access,proto3" json:"Access,omitempty"`
}

func (x *CheckAccessResponse) Reset() {
	*x = CheckAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_community_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccessResponse) ProtoMessage() {}

func (x *CheckAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_community_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccessResponse.ProtoReflect.Descriptor instead.
func (*CheckAccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_community_proto_rawDescGZIP(), []int{1}
}

func (x *CheckAccessResponse) GetAccess() bool {
	if x != nil {
		return x.Access
	}
	return false
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorID    uint32 `protobuf:"varint,1,opt,name=AuthorID,proto3" json:"AuthorID,omitempty"`
	CommunityID uint32 `protobuf:"varint,2,opt,name=CommunityID,proto3" json:"CommunityID,omitempty"`
	Author      string `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
	Avatar      string `protobuf:"bytes,4,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_community_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_proto_community_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_proto_community_proto_rawDescGZIP(), []int{2}
}

func (x *Header) GetAuthorID() uint32 {
	if x != nil {
		return x.AuthorID
	}
	return 0
}

func (x *Header) GetCommunityID() uint32 {
	if x != nil {
		return x.CommunityID
	}
	return 0
}

func (x *Header) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Header) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type GetHeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityID uint32 `protobuf:"varint,1,opt,name=CommunityID,proto3" json:"CommunityID,omitempty"`
}

func (x *GetHeaderRequest) Reset() {
	*x = GetHeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_community_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHeaderRequest) ProtoMessage() {}

func (x *GetHeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_community_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHeaderRequest.ProtoReflect.Descriptor instead.
func (*GetHeaderRequest) Descriptor() ([]byte, []int) {
	return file_proto_community_proto_rawDescGZIP(), []int{3}
}

func (x *GetHeaderRequest) GetCommunityID() uint32 {
	if x != nil {
		return x.CommunityID
	}
	return 0
}

type GetHeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *Header `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
}

func (x *GetHeaderResponse) Reset() {
	*x = GetHeaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_community_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHeaderResponse) ProtoMessage() {}

func (x *GetHeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_community_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHeaderResponse.ProtoReflect.Descriptor instead.
func (*GetHeaderResponse) Descriptor() ([]byte, []int) {
	return file_proto_community_proto_rawDescGZIP(), []int{4}
}

func (x *GetHeaderResponse) GetHead() *Header {
	if x != nil {
		return x.Head
	}
	return nil
}

var File_proto_community_proto protoreflect.FileDescriptor

var file_proto_community_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x4e, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x49, 0x44, 0x22, 0x2d, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x76, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x34, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x49, 0x44, 0x22, 0x3e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x48, 0x65, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x04, 0x48,
	0x65, 0x61, 0x64, 0x32, 0xbc, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x32, 0x30, 0x32, 0x34, 0x5f, 0x32, 0x5f, 0x42, 0x65, 0x74, 0x74, 0x65, 0x72, 0x43, 0x61,
	0x6c, 0x6c, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_community_proto_rawDescOnce sync.Once
	file_proto_community_proto_rawDescData = file_proto_community_proto_rawDesc
)

func file_proto_community_proto_rawDescGZIP() []byte {
	file_proto_community_proto_rawDescOnce.Do(func() {
		file_proto_community_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_community_proto_rawDescData)
	})
	return file_proto_community_proto_rawDescData
}

var file_proto_community_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_community_proto_goTypes = []any{
	(*CheckAccessRequest)(nil),  // 0: community_api.CheckAccessRequest
	(*CheckAccessResponse)(nil), // 1: community_api.CheckAccessResponse
	(*Header)(nil),              // 2: community_api.Header
	(*GetHeaderRequest)(nil),    // 3: community_api.GetHeaderRequest
	(*GetHeaderResponse)(nil),   // 4: community_api.GetHeaderResponse
}
var file_proto_community_proto_depIdxs = []int32{
	2, // 0: community_api.GetHeaderResponse.Head:type_name -> community_api.Header
	0, // 1: community_api.CommunityService.CheckAccess:input_type -> community_api.CheckAccessRequest
	3, // 2: community_api.CommunityService.GetHeader:input_type -> community_api.GetHeaderRequest
	1, // 3: community_api.CommunityService.CheckAccess:output_type -> community_api.CheckAccessResponse
	4, // 4: community_api.CommunityService.GetHeader:output_type -> community_api.GetHeaderResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_community_proto_init() }
func file_proto_community_proto_init() {
	if File_proto_community_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_community_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CheckAccessRequest); i {
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
		file_proto_community_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CheckAccessResponse); i {
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
		file_proto_community_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Header); i {
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
		file_proto_community_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetHeaderRequest); i {
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
		file_proto_community_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetHeaderResponse); i {
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
			RawDescriptor: file_proto_community_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_community_proto_goTypes,
		DependencyIndexes: file_proto_community_proto_depIdxs,
		MessageInfos:      file_proto_community_proto_msgTypes,
	}.Build()
	File_proto_community_proto = out.File
	file_proto_community_proto_rawDesc = nil
	file_proto_community_proto_goTypes = nil
	file_proto_community_proto_depIdxs = nil
}
