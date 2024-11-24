// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/post.proto

package post_api

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head   *Header `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	UserID uint32  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetHead() *Header {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *Request) GetUserID() uint32 {
	if x != nil {
		return x.UserID
	}
	return 0
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
		mi := &file_proto_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[1]
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
	return file_proto_post_proto_rawDescGZIP(), []int{1}
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=Posts,proto3" json:"Posts,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PostContent *Content `protobuf:"bytes,2,opt,name=PostContent,proto3" json:"PostContent,omitempty"`
	Head        *Header  `protobuf:"bytes,3,opt,name=Head,proto3" json:"Head,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{3}
}

func (x *Post) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Post) GetPostContent() *Content {
	if x != nil {
		return x.PostContent
	}
	return nil
}

func (x *Post) GetHead() *Header {
	if x != nil {
		return x.Head
	}
	return nil
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text      string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	File      string `protobuf:"bytes,2,opt,name=File,proto3" json:"File,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{4}
}

func (x *Content) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Content) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *Content) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Content) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_proto_post_proto protoreflect.FileDescriptor

var file_proto_post_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x47, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x48, 0x65, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x04, 0x48, 0x65, 0x61, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x76, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x30, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x22,
	0x71, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x04,
	0x48, 0x65, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x04, 0x48, 0x65,
	0x61, 0x64, 0x22, 0x6d, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x32, 0x49, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x32, 0x30, 0x32, 0x34, 0x5f,
	0x32, 0x5f, 0x42, 0x65, 0x74, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_post_proto_rawDescOnce sync.Once
	file_proto_post_proto_rawDescData = file_proto_post_proto_rawDesc
)

func file_proto_post_proto_rawDescGZIP() []byte {
	file_proto_post_proto_rawDescOnce.Do(func() {
		file_proto_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_post_proto_rawDescData)
	})
	return file_proto_post_proto_rawDescData
}

var file_proto_post_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_post_proto_goTypes = []any{
	(*Request)(nil),  // 0: post_api.Request
	(*Header)(nil),   // 1: post_api.Header
	(*Response)(nil), // 2: post_api.Response
	(*Post)(nil),     // 3: post_api.Post
	(*Content)(nil),  // 4: post_api.Content
}
var file_proto_post_proto_depIdxs = []int32{
	1, // 0: post_api.Request.Head:type_name -> post_api.Header
	3, // 1: post_api.Response.Posts:type_name -> post_api.Post
	4, // 2: post_api.Post.PostContent:type_name -> post_api.Content
	1, // 3: post_api.Post.Head:type_name -> post_api.Header
	0, // 4: post_api.PostService.GetAuthorsPosts:input_type -> post_api.Request
	2, // 5: post_api.PostService.GetAuthorsPosts:output_type -> post_api.Response
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_post_proto_init() }
func file_proto_post_proto_init() {
	if File_proto_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_post_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Request); i {
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
		file_proto_post_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_proto_post_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Response); i {
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
		file_proto_post_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Post); i {
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
		file_proto_post_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Content); i {
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
			RawDescriptor: file_proto_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_post_proto_goTypes,
		DependencyIndexes: file_proto_post_proto_depIdxs,
		MessageInfos:      file_proto_post_proto_msgTypes,
	}.Build()
	File_proto_post_proto = out.File
	file_proto_post_proto_rawDesc = nil
	file_proto_post_proto_goTypes = nil
	file_proto_post_proto_depIdxs = nil
}
