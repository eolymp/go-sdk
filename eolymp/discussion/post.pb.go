// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.18.1
// source: eolymp/discussion/post.proto

package discussion

import (
	ecm "github.com/eolymp/go-sdk/eolymp/ecm"
	wellknown "github.com/eolymp/go-sdk/eolymp/wellknown"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Post_Status int32

const (
	Post_UNKNOWN_STATUS Post_Status = 0 // not used
	Post_DRAFT          Post_Status = 1 // marked as draft by author
	Post_IN_REVIEW      Post_Status = 2 // published by author by awaiting moderation
	Post_PUBLISHED      Post_Status = 3 // published and available to everyone
	Post_REJECTED       Post_Status = 4 // rejected by moderator, has to be edited and re-published
)

// Enum value maps for Post_Status.
var (
	Post_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "DRAFT",
		2: "IN_REVIEW",
		3: "PUBLISHED",
		4: "REJECTED",
	}
	Post_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"DRAFT":          1,
		"IN_REVIEW":      2,
		"PUBLISHED":      3,
		"REJECTED":       4,
	}
)

func (x Post_Status) Enum() *Post_Status {
	p := new(Post_Status)
	*p = x
	return p
}

func (x Post_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Post_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_discussion_post_proto_enumTypes[0].Descriptor()
}

func (Post_Status) Type() protoreflect.EnumType {
	return &file_eolymp_discussion_post_proto_enumTypes[0]
}

func (x Post_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Post_Status.Descriptor instead.
func (Post_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_discussion_post_proto_rawDescGZIP(), []int{0, 0}
}

type Post_Extra int32

const (
	Post_UNKNOWN_EXTRA   Post_Extra = 0
	Post_MESSAGE_CONTENT Post_Extra = 1 // include message source (original Markdown, LaTeX etc)
	Post_MESSAGE_RENDER  Post_Extra = 2 // include rendered message in ECM
	Post_MESSAGE_PREVIEW Post_Extra = 3 // include rendered but trimmed message in ECM, overrides MESSAGE_RENDER
)

// Enum value maps for Post_Extra.
var (
	Post_Extra_name = map[int32]string{
		0: "UNKNOWN_EXTRA",
		1: "MESSAGE_CONTENT",
		2: "MESSAGE_RENDER",
		3: "MESSAGE_PREVIEW",
	}
	Post_Extra_value = map[string]int32{
		"UNKNOWN_EXTRA":   0,
		"MESSAGE_CONTENT": 1,
		"MESSAGE_RENDER":  2,
		"MESSAGE_PREVIEW": 3,
	}
)

func (x Post_Extra) Enum() *Post_Extra {
	p := new(Post_Extra)
	*p = x
	return p
}

func (x Post_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Post_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_discussion_post_proto_enumTypes[1].Descriptor()
}

func (Post_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_discussion_post_proto_enumTypes[1]
}

func (x Post_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Post_Extra.Descriptor instead.
func (Post_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_discussion_post_proto_rawDescGZIP(), []int{0, 1}
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Types that are assignable to Author:
	//	*Post_UserId
	//	*Post_MemberId
	Author    isPost_Author          `protobuf_oneof:"author"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	PostedAt  *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=posted_at,json=postedAt,proto3" json:"posted_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Title     string                 `protobuf:"bytes,100,opt,name=title,proto3" json:"title,omitempty"`
	Message   *ecm.Content           `protobuf:"bytes,101,opt,name=message,proto3" json:"message,omitempty"`
	Labels    []string               `protobuf:"bytes,120,rep,name=labels,proto3" json:"labels,omitempty"`
	Links     []*wellknown.Link      `protobuf:"bytes,200,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_discussion_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_post_proto_msgTypes[0]
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
	return file_eolymp_discussion_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (m *Post) GetAuthor() isPost_Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (x *Post) GetUserId() string {
	if x, ok := x.GetAuthor().(*Post_UserId); ok {
		return x.UserId
	}
	return ""
}

func (x *Post) GetMemberId() string {
	if x, ok := x.GetAuthor().(*Post_MemberId); ok {
		return x.MemberId
	}
	return ""
}

func (x *Post) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Post) GetPostedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PostedAt
	}
	return nil
}

func (x *Post) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetMessage() *ecm.Content {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Post) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Post) GetLinks() []*wellknown.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

type isPost_Author interface {
	isPost_Author()
}

type Post_UserId struct {
	UserId string `protobuf:"bytes,10,opt,name=user_id,json=userId,proto3,oneof"`
}

type Post_MemberId struct {
	MemberId string `protobuf:"bytes,11,opt,name=member_id,json=memberId,proto3,oneof"`
}

func (*Post_UserId) isPost_Author() {}

func (*Post_MemberId) isPost_Author() {}

var File_eolymp_discussion_post_proto protoreflect.FileDescriptor

var file_eolymp_discussion_post_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x04, 0x0a, 0x04, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x70, 0x6f, 0x73,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63,
	0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x78, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x53, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x41, 0x46, 0x54,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x04, 0x22, 0x58,
	0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12,
	0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x4e, 0x44, 0x45,
	0x52, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x50,
	0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x03, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x3b, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eolymp_discussion_post_proto_rawDescOnce sync.Once
	file_eolymp_discussion_post_proto_rawDescData = file_eolymp_discussion_post_proto_rawDesc
)

func file_eolymp_discussion_post_proto_rawDescGZIP() []byte {
	file_eolymp_discussion_post_proto_rawDescOnce.Do(func() {
		file_eolymp_discussion_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_discussion_post_proto_rawDescData)
	})
	return file_eolymp_discussion_post_proto_rawDescData
}

var file_eolymp_discussion_post_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_discussion_post_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_discussion_post_proto_goTypes = []interface{}{
	(Post_Status)(0),              // 0: eolymp.discussion.Post.Status
	(Post_Extra)(0),               // 1: eolymp.discussion.Post.Extra
	(*Post)(nil),                  // 2: eolymp.discussion.Post
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*ecm.Content)(nil),           // 4: eolymp.ecm.Content
	(*wellknown.Link)(nil),        // 5: eolymp.wellknown.Link
}
var file_eolymp_discussion_post_proto_depIdxs = []int32{
	3, // 0: eolymp.discussion.Post.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: eolymp.discussion.Post.posted_at:type_name -> google.protobuf.Timestamp
	3, // 2: eolymp.discussion.Post.updated_at:type_name -> google.protobuf.Timestamp
	4, // 3: eolymp.discussion.Post.message:type_name -> eolymp.ecm.Content
	5, // 4: eolymp.discussion.Post.links:type_name -> eolymp.wellknown.Link
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_discussion_post_proto_init() }
func file_eolymp_discussion_post_proto_init() {
	if File_eolymp_discussion_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_discussion_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_eolymp_discussion_post_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Post_UserId)(nil),
		(*Post_MemberId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_discussion_post_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_discussion_post_proto_goTypes,
		DependencyIndexes: file_eolymp_discussion_post_proto_depIdxs,
		EnumInfos:         file_eolymp_discussion_post_proto_enumTypes,
		MessageInfos:      file_eolymp_discussion_post_proto_msgTypes,
	}.Build()
	File_eolymp_discussion_post_proto = out.File
	file_eolymp_discussion_post_proto_rawDesc = nil
	file_eolymp_discussion_post_proto_goTypes = nil
	file_eolymp_discussion_post_proto_depIdxs = nil
}
