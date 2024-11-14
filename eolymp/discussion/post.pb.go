// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
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

type Post_Moderation int32

const (
	Post_UNKNOWN_MODERATION Post_Moderation = 0 // not used
	Post_PENDING            Post_Moderation = 1 // not reviewed by moderator
	Post_IN_REVIEW          Post_Moderation = 2 // in review by moderator
	Post_APPROVED           Post_Moderation = 3 // approved by moderator
	Post_REJECTED           Post_Moderation = 4 // rejected by moderator
)

// Enum value maps for Post_Moderation.
var (
	Post_Moderation_name = map[int32]string{
		0: "UNKNOWN_MODERATION",
		1: "PENDING",
		2: "IN_REVIEW",
		3: "APPROVED",
		4: "REJECTED",
	}
	Post_Moderation_value = map[string]int32{
		"UNKNOWN_MODERATION": 0,
		"PENDING":            1,
		"IN_REVIEW":          2,
		"APPROVED":           3,
		"REJECTED":           4,
	}
)

func (x Post_Moderation) Enum() *Post_Moderation {
	p := new(Post_Moderation)
	*p = x
	return p
}

func (x Post_Moderation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Post_Moderation) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_discussion_post_proto_enumTypes[0].Descriptor()
}

func (Post_Moderation) Type() protoreflect.EnumType {
	return &file_eolymp_discussion_post_proto_enumTypes[0]
}

func (x Post_Moderation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Post_Moderation.Descriptor instead.
func (Post_Moderation) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_discussion_post_proto_rawDescGZIP(), []int{0, 0}
}

type Post_Extra int32

const (
	Post_UNKNOWN_EXTRA  Post_Extra = 0
	Post_CONTENT_VALUE  Post_Extra = 1 // include content source (original Markdown, LaTeX etc)
	Post_CONTENT_RENDER Post_Extra = 2 // include rendered content in ECM
	Post_PREVIEW        Post_Extra = 3 // include rendered but trimmed message in ECM
	Post_VOTE           Post_Extra = 4 // include vote field
)

// Enum value maps for Post_Extra.
var (
	Post_Extra_name = map[int32]string{
		0: "UNKNOWN_EXTRA",
		1: "CONTENT_VALUE",
		2: "CONTENT_RENDER",
		3: "PREVIEW",
		4: "VOTE",
	}
	Post_Extra_value = map[string]int32{
		"UNKNOWN_EXTRA":  0,
		"CONTENT_VALUE":  1,
		"CONTENT_RENDER": 2,
		"PREVIEW":        3,
		"VOTE":           4,
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

	Id         string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url        string          `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	SourceId   string          `protobuf:"bytes,7,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`                             // if set, marks this post as translation for a post specified in this field
	SourceUrl  string          `protobuf:"bytes,8,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`                          // populated if source_id is set
	Draft      bool            `protobuf:"varint,3,opt,name=draft,proto3" json:"draft,omitempty"`                                                  // marked as draft and only shown to author
	Public     bool            `protobuf:"varint,4,opt,name=public,proto3" json:"public,omitempty"`                                                // visible and available to everyone (ie. published and passed moderation)
	Featured   bool            `protobuf:"varint,9,opt,name=featured,proto3" json:"featured,omitempty"`                                            // marked as featured (shown on home page)
	Pinned     bool            `protobuf:"varint,13,opt,name=pinned,proto3" json:"pinned,omitempty"`                                               // pinned on top of the page
	Moderation Post_Moderation `protobuf:"varint,5,opt,name=moderation,proto3,enum=eolymp.discussion.Post_Moderation" json:"moderation,omitempty"` // moderation status
	// Types that are assignable to Author:
	//
	//	*Post_UserId
	//	*Post_MemberId
	Author      isPost_Author          `protobuf_oneof:"author"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	PublishedAt *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	TypeId      string                 `protobuf:"bytes,6,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"` // optionally, post type from PostTypeService
	Locale      string                 `protobuf:"bytes,102,opt,name=locale,proto3" json:"locale,omitempty"`
	Title       string                 `protobuf:"bytes,103,opt,name=title,proto3" json:"title,omitempty"`                       // automatically populated from content
	ImageUrl    string                 `protobuf:"bytes,104,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"` // automatically populated from content
	Content     *ecm.Content           `protobuf:"bytes,101,opt,name=content,proto3" json:"content,omitempty"`
	Preview     *Post_Preview          `protobuf:"bytes,110,opt,name=preview,proto3" json:"preview,omitempty"`                         // preview is generated automatically from the content
	Vote        int32                  `protobuf:"varint,12,opt,name=vote,proto3" json:"vote,omitempty"`                               // vote of authenticated user (+1 or -1)
	VoteCount   int32                  `protobuf:"varint,30,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`    // total vote count
	ReplyCount  int32                  `protobuf:"varint,31,opt,name=reply_count,json=replyCount,proto3" json:"reply_count,omitempty"` // total number of replies
	Labels      []string               `protobuf:"bytes,120,rep,name=labels,proto3" json:"labels,omitempty"`
	Links       []*wellknown.Link      `protobuf:"bytes,200,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_eolymp_discussion_post_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_post_proto_msgTypes[0]
	if x != nil {
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

func (x *Post) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *Post) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *Post) GetDraft() bool {
	if x != nil {
		return x.Draft
	}
	return false
}

func (x *Post) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *Post) GetFeatured() bool {
	if x != nil {
		return x.Featured
	}
	return false
}

func (x *Post) GetPinned() bool {
	if x != nil {
		return x.Pinned
	}
	return false
}

func (x *Post) GetModeration() Post_Moderation {
	if x != nil {
		return x.Moderation
	}
	return Post_UNKNOWN_MODERATION
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

func (x *Post) GetPublishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishedAt
	}
	return nil
}

func (x *Post) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Post) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *Post) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Post) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Post) GetPreview() *Post_Preview {
	if x != nil {
		return x.Preview
	}
	return nil
}

func (x *Post) GetVote() int32 {
	if x != nil {
		return x.Vote
	}
	return 0
}

func (x *Post) GetVoteCount() int32 {
	if x != nil {
		return x.VoteCount
	}
	return 0
}

func (x *Post) GetReplyCount() int32 {
	if x != nil {
		return x.ReplyCount
	}
	return 0
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

type Post_Translation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Locale  string       `protobuf:"bytes,102,opt,name=locale,proto3" json:"locale,omitempty"`
	Content *ecm.Content `protobuf:"bytes,101,opt,name=content,proto3" json:"content,omitempty"`
	Labels  []string     `protobuf:"bytes,120,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Post_Translation) Reset() {
	*x = Post_Translation{}
	mi := &file_eolymp_discussion_post_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post_Translation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post_Translation) ProtoMessage() {}

func (x *Post_Translation) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_post_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post_Translation.ProtoReflect.Descriptor instead.
func (*Post_Translation) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_post_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Post_Translation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post_Translation) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Post_Translation) GetContent() *ecm.Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Post_Translation) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Post_Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src    string   `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	Width  int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Class  []string `protobuf:"bytes,4,rep,name=class,proto3" json:"class,omitempty"`
}

func (x *Post_Image) Reset() {
	*x = Post_Image{}
	mi := &file_eolymp_discussion_post_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post_Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post_Image) ProtoMessage() {}

func (x *Post_Image) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_post_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post_Image.ProtoReflect.Descriptor instead.
func (*Post_Image) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_post_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Post_Image) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *Post_Image) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Post_Image) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Post_Image) GetClass() []string {
	if x != nil {
		return x.Class
	}
	return nil
}

type Post_Preview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string      `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Image   *Post_Image `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Content *ecm.Node   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Post_Preview) Reset() {
	*x = Post_Preview{}
	mi := &file_eolymp_discussion_post_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post_Preview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post_Preview) ProtoMessage() {}

func (x *Post_Preview) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_discussion_post_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post_Preview.ProtoReflect.Descriptor instead.
func (*Post_Preview) Descriptor() ([]byte, []int) {
	return file_eolymp_discussion_post_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Post_Preview) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post_Preview) GetImage() *Post_Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Post_Preview) GetContent() *ecm.Node {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_eolymp_discussion_post_proto protoreflect.FileDescriptor

var file_eolymp_discussion_post_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x84, 0x0b, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x6d, 0x6f, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x39, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f,
	0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x76, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x78, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18,
	0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05,
	0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x1a, 0x7c, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x78, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x1a, 0x5d, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x14,
	0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x1a, 0x80, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4d,
	0x4f, 0x44, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x5f, 0x52,
	0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x52, 0x4f,
	0x56, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x10, 0x04, 0x22, 0x58, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x11, 0x0a, 0x0d,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45,
	0x4e, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45,
	0x57, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x4f, 0x54, 0x45, 0x10, 0x04, 0x42, 0x08, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_eolymp_discussion_post_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_discussion_post_proto_goTypes = []any{
	(Post_Moderation)(0),          // 0: eolymp.discussion.Post.Moderation
	(Post_Extra)(0),               // 1: eolymp.discussion.Post.Extra
	(*Post)(nil),                  // 2: eolymp.discussion.Post
	(*Post_Translation)(nil),      // 3: eolymp.discussion.Post.Translation
	(*Post_Image)(nil),            // 4: eolymp.discussion.Post.Image
	(*Post_Preview)(nil),          // 5: eolymp.discussion.Post.Preview
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*ecm.Content)(nil),           // 7: eolymp.ecm.Content
	(*wellknown.Link)(nil),        // 8: eolymp.wellknown.Link
	(*ecm.Node)(nil),              // 9: eolymp.ecm.Node
}
var file_eolymp_discussion_post_proto_depIdxs = []int32{
	0,  // 0: eolymp.discussion.Post.moderation:type_name -> eolymp.discussion.Post.Moderation
	6,  // 1: eolymp.discussion.Post.created_at:type_name -> google.protobuf.Timestamp
	6,  // 2: eolymp.discussion.Post.published_at:type_name -> google.protobuf.Timestamp
	6,  // 3: eolymp.discussion.Post.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 4: eolymp.discussion.Post.content:type_name -> eolymp.ecm.Content
	5,  // 5: eolymp.discussion.Post.preview:type_name -> eolymp.discussion.Post.Preview
	8,  // 6: eolymp.discussion.Post.links:type_name -> eolymp.wellknown.Link
	7,  // 7: eolymp.discussion.Post.Translation.content:type_name -> eolymp.ecm.Content
	4,  // 8: eolymp.discussion.Post.Preview.image:type_name -> eolymp.discussion.Post.Image
	9,  // 9: eolymp.discussion.Post.Preview.content:type_name -> eolymp.ecm.Node
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_eolymp_discussion_post_proto_init() }
func file_eolymp_discussion_post_proto_init() {
	if File_eolymp_discussion_post_proto != nil {
		return
	}
	file_eolymp_discussion_post_proto_msgTypes[0].OneofWrappers = []any{
		(*Post_UserId)(nil),
		(*Post_MemberId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_discussion_post_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
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
