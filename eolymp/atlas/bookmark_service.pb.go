// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: eolymp/atlas/bookmark_service.proto

package atlas

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BookmarkChangedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProblemId     string                 `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	MemberId      string                 `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Before        bool                   `protobuf:"varint,3,opt,name=before,proto3" json:"before,omitempty"`
	After         bool                   `protobuf:"varint,4,opt,name=after,proto3" json:"after,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BookmarkChangedEvent) Reset() {
	*x = BookmarkChangedEvent{}
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookmarkChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookmarkChangedEvent) ProtoMessage() {}

func (x *BookmarkChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookmarkChangedEvent.ProtoReflect.Descriptor instead.
func (*BookmarkChangedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_bookmark_service_proto_rawDescGZIP(), []int{0}
}

func (x *BookmarkChangedEvent) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *BookmarkChangedEvent) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *BookmarkChangedEvent) GetBefore() bool {
	if x != nil {
		return x.Before
	}
	return false
}

func (x *BookmarkChangedEvent) GetAfter() bool {
	if x != nil {
		return x.After
	}
	return false
}

type SetBookmarkInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bookmark      bool                   `protobuf:"varint,2,opt,name=bookmark,proto3" json:"bookmark,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetBookmarkInput) Reset() {
	*x = SetBookmarkInput{}
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetBookmarkInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBookmarkInput) ProtoMessage() {}

func (x *SetBookmarkInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBookmarkInput.ProtoReflect.Descriptor instead.
func (*SetBookmarkInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_bookmark_service_proto_rawDescGZIP(), []int{1}
}

func (x *SetBookmarkInput) GetBookmark() bool {
	if x != nil {
		return x.Bookmark
	}
	return false
}

type SetBookmarkOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetBookmarkOutput) Reset() {
	*x = SetBookmarkOutput{}
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetBookmarkOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBookmarkOutput) ProtoMessage() {}

func (x *SetBookmarkOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBookmarkOutput.ProtoReflect.Descriptor instead.
func (*SetBookmarkOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_bookmark_service_proto_rawDescGZIP(), []int{2}
}

type GetBookmarkInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookmarkInput) Reset() {
	*x = GetBookmarkInput{}
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookmarkInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookmarkInput) ProtoMessage() {}

func (x *GetBookmarkInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookmarkInput.ProtoReflect.Descriptor instead.
func (*GetBookmarkInput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_bookmark_service_proto_rawDescGZIP(), []int{3}
}

type GetBookmarkOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bookmark      bool                   `protobuf:"varint,1,opt,name=bookmark,proto3" json:"bookmark,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookmarkOutput) Reset() {
	*x = GetBookmarkOutput{}
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookmarkOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookmarkOutput) ProtoMessage() {}

func (x *GetBookmarkOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_bookmark_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookmarkOutput.ProtoReflect.Descriptor instead.
func (*GetBookmarkOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_bookmark_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookmarkOutput) GetBookmark() bool {
	if x != nil {
		return x.Bookmark
	}
	return false
}

var File_eolymp_atlas_bookmark_service_proto protoreflect.FileDescriptor

const file_eolymp_atlas_bookmark_service_proto_rawDesc = "" +
	"\n" +
	"#eolymp/atlas/bookmark_service.proto\x12\feolymp.atlas\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\"\x80\x01\n" +
	"\x14BookmarkChangedEvent\x12\x1d\n" +
	"\n" +
	"problem_id\x18\x01 \x01(\tR\tproblemId\x12\x1b\n" +
	"\tmember_id\x18\x02 \x01(\tR\bmemberId\x12\x16\n" +
	"\x06before\x18\x03 \x01(\bR\x06before\x12\x14\n" +
	"\x05after\x18\x04 \x01(\bR\x05after\".\n" +
	"\x10SetBookmarkInput\x12\x1a\n" +
	"\bbookmark\x18\x02 \x01(\bR\bbookmark\"\x13\n" +
	"\x11SetBookmarkOutput\"\x12\n" +
	"\x10GetBookmarkInput\"/\n" +
	"\x11GetBookmarkOutput\x12\x1a\n" +
	"\bbookmark\x18\x01 \x01(\bR\bbookmark2\xf5\x01\n" +
	"\x0fBookmarkService\x12p\n" +
	"\vGetBookmark\x12\x1e.eolymp.atlas.GetBookmarkInput\x1a\x1f.eolymp.atlas.GetBookmarkOutput\" \xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xc8A\xf8\xe2\n" +
	"d\x82\xd3\xe4\x93\x02\v\x12\t/bookmark\x12p\n" +
	"\vSetBookmark\x12\x1e.eolymp.atlas.SetBookmarkInput\x1a\x1f.eolymp.atlas.SetBookmarkOutput\" \xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\x00@\xf8\xe2\n" +
	"\n" +
	"\x82\xd3\xe4\x93\x02\v\"\t/bookmarkB-Z+github.com/eolymp/go-sdk/eolymp/atlas;atlasb\x06proto3"

var (
	file_eolymp_atlas_bookmark_service_proto_rawDescOnce sync.Once
	file_eolymp_atlas_bookmark_service_proto_rawDescData []byte
)

func file_eolymp_atlas_bookmark_service_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_bookmark_service_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_bookmark_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_atlas_bookmark_service_proto_rawDesc), len(file_eolymp_atlas_bookmark_service_proto_rawDesc)))
	})
	return file_eolymp_atlas_bookmark_service_proto_rawDescData
}

var file_eolymp_atlas_bookmark_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eolymp_atlas_bookmark_service_proto_goTypes = []any{
	(*BookmarkChangedEvent)(nil), // 0: eolymp.atlas.BookmarkChangedEvent
	(*SetBookmarkInput)(nil),     // 1: eolymp.atlas.SetBookmarkInput
	(*SetBookmarkOutput)(nil),    // 2: eolymp.atlas.SetBookmarkOutput
	(*GetBookmarkInput)(nil),     // 3: eolymp.atlas.GetBookmarkInput
	(*GetBookmarkOutput)(nil),    // 4: eolymp.atlas.GetBookmarkOutput
}
var file_eolymp_atlas_bookmark_service_proto_depIdxs = []int32{
	3, // 0: eolymp.atlas.BookmarkService.GetBookmark:input_type -> eolymp.atlas.GetBookmarkInput
	1, // 1: eolymp.atlas.BookmarkService.SetBookmark:input_type -> eolymp.atlas.SetBookmarkInput
	4, // 2: eolymp.atlas.BookmarkService.GetBookmark:output_type -> eolymp.atlas.GetBookmarkOutput
	2, // 3: eolymp.atlas.BookmarkService.SetBookmark:output_type -> eolymp.atlas.SetBookmarkOutput
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_bookmark_service_proto_init() }
func file_eolymp_atlas_bookmark_service_proto_init() {
	if File_eolymp_atlas_bookmark_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_atlas_bookmark_service_proto_rawDesc), len(file_eolymp_atlas_bookmark_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_atlas_bookmark_service_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_bookmark_service_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_bookmark_service_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_bookmark_service_proto = out.File
	file_eolymp_atlas_bookmark_service_proto_goTypes = nil
	file_eolymp_atlas_bookmark_service_proto_depIdxs = nil
}
