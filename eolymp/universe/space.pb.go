// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/universe/space.proto

package universe

import (
	acl "github.com/eolymp/go-sdk/eolymp/acl"
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

type Space_Status int32

const (
	Space_UNKNOWN_STATUS Space_Status = 0
	Space_TRIAL          Space_Status = 1
	Space_ACTIVE         Space_Status = 2
	Space_SUSPENDED      Space_Status = 3
)

// Enum value maps for Space_Status.
var (
	Space_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "TRIAL",
		2: "ACTIVE",
		3: "SUSPENDED",
	}
	Space_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"TRIAL":          1,
		"ACTIVE":         2,
		"SUSPENDED":      3,
	}
)

func (x Space_Status) Enum() *Space_Status {
	p := new(Space_Status)
	*p = x
	return p
}

func (x Space_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Space_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_universe_space_proto_enumTypes[0].Descriptor()
}

func (Space_Status) Type() protoreflect.EnumType {
	return &file_eolymp_universe_space_proto_enumTypes[0]
}

func (x Space_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Space_Status.Descriptor instead.
func (Space_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_universe_space_proto_rawDescGZIP(), []int{0, 0}
}

type Space_Visibility int32

const (
	Space_UNKNOWN_VISIBILITY Space_Visibility = 0
	Space_PUBLIC             Space_Visibility = 1 // anonymous users can see some space content
	Space_PRIVATE            Space_Visibility = 2 // everyone must sign in to access space
)

// Enum value maps for Space_Visibility.
var (
	Space_Visibility_name = map[int32]string{
		0: "UNKNOWN_VISIBILITY",
		1: "PUBLIC",
		2: "PRIVATE",
	}
	Space_Visibility_value = map[string]int32{
		"UNKNOWN_VISIBILITY": 0,
		"PUBLIC":             1,
		"PRIVATE":            2,
	}
)

func (x Space_Visibility) Enum() *Space_Visibility {
	p := new(Space_Visibility)
	*p = x
	return p
}

func (x Space_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Space_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_universe_space_proto_enumTypes[1].Descriptor()
}

func (Space_Visibility) Type() protoreflect.EnumType {
	return &file_eolymp_universe_space_proto_enumTypes[1]
}

func (x Space_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Space_Visibility.Descriptor instead.
func (Space_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_universe_space_proto_rawDescGZIP(), []int{0, 1}
}

type Space_Extra int32

const (
	Space_UNKNOWN_EXTRA Space_Extra = 0
	Space_SUBSCRIPTION  Space_Extra = 1
	Space_ACTIONS       Space_Extra = 2
)

// Enum value maps for Space_Extra.
var (
	Space_Extra_name = map[int32]string{
		0: "UNKNOWN_EXTRA",
		1: "SUBSCRIPTION",
		2: "ACTIONS",
	}
	Space_Extra_value = map[string]int32{
		"UNKNOWN_EXTRA": 0,
		"SUBSCRIPTION":  1,
		"ACTIONS":       2,
	}
)

func (x Space_Extra) Enum() *Space_Extra {
	p := new(Space_Extra)
	*p = x
	return p
}

func (x Space_Extra) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Space_Extra) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_universe_space_proto_enumTypes[2].Descriptor()
}

func (Space_Extra) Type() protoreflect.EnumType {
	return &file_eolymp_universe_space_proto_enumTypes[2]
}

func (x Space_Extra) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Space_Extra.Descriptor instead.
func (Space_Extra) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_universe_space_proto_rawDescGZIP(), []int{0, 2}
}

type Space struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                    // space unique identifier
	Key          string              `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`                                  // space key used to build URLs
	Url          string              `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`                                  // space url
	HomeUrl      string              `protobuf:"bytes,50,opt,name=home_url,json=homeUrl,proto3" json:"home_url,omitempty"`          // space home page URL
	IssuerUrl    string              `protobuf:"bytes,51,opt,name=issuer_url,json=issuerUrl,proto3" json:"issuer_url,omitempty"`    // space issuer URL (used for issuing tokens)
	GraphqlUrl   string              `protobuf:"bytes,52,opt,name=graphql_url,json=graphqlUrl,proto3" json:"graphql_url,omitempty"` // space graphql endpoint
	Name         string              `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`                               // human friendly name
	Image        string              `protobuf:"bytes,11,opt,name=image,proto3" json:"image,omitempty"`                             // space logo image
	Visibility   Space_Visibility    `protobuf:"varint,14,opt,name=visibility,proto3,enum=eolymp.universe.Space_Visibility" json:"visibility,omitempty"`
	Status       Space_Status        `protobuf:"varint,16,opt,name=status,proto3,enum=eolymp.universe.Space_Status" json:"status,omitempty"` // space status
	Subscription *Space_Subscription `protobuf:"bytes,800,opt,name=subscription,proto3" json:"subscription,omitempty"`                       // subscription details (private)
	Affiliation  string              `protobuf:"bytes,15,opt,name=affiliation,proto3" json:"affiliation,omitempty"`                          // space affiliation label
	Actions      []acl.Action        `protobuf:"varint,100,rep,packed,name=actions,proto3,enum=eolymp.acl.Action" json:"actions,omitempty"`
}

func (x *Space) Reset() {
	*x = Space{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_space_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Space) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Space) ProtoMessage() {}

func (x *Space) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_space_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Space.ProtoReflect.Descriptor instead.
func (*Space) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_space_proto_rawDescGZIP(), []int{0}
}

func (x *Space) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Space) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Space) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Space) GetHomeUrl() string {
	if x != nil {
		return x.HomeUrl
	}
	return ""
}

func (x *Space) GetIssuerUrl() string {
	if x != nil {
		return x.IssuerUrl
	}
	return ""
}

func (x *Space) GetGraphqlUrl() string {
	if x != nil {
		return x.GraphqlUrl
	}
	return ""
}

func (x *Space) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Space) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Space) GetVisibility() Space_Visibility {
	if x != nil {
		return x.Visibility
	}
	return Space_UNKNOWN_VISIBILITY
}

func (x *Space) GetStatus() Space_Status {
	if x != nil {
		return x.Status
	}
	return Space_UNKNOWN_STATUS
}

func (x *Space) GetSubscription() *Space_Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

func (x *Space) GetAffiliation() string {
	if x != nil {
		return x.Affiliation
	}
	return ""
}

func (x *Space) GetActions() []acl.Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

type Space_Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan               string                 `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	Seats              uint32                 `protobuf:"varint,2,opt,name=seats,proto3" json:"seats,omitempty"`
	Quota              *Quota                 `protobuf:"bytes,10,opt,name=quota,proto3" json:"quota,omitempty"`
	BillingPeriodStart *timestamppb.Timestamp `protobuf:"bytes,90,opt,name=billing_period_start,json=billingPeriodStart,proto3" json:"billing_period_start,omitempty"`
	BillingPeriodEnd   *timestamppb.Timestamp `protobuf:"bytes,91,opt,name=billing_period_end,json=billingPeriodEnd,proto3" json:"billing_period_end,omitempty"`
	QuotaPeriodStart   *timestamppb.Timestamp `protobuf:"bytes,92,opt,name=quota_period_start,json=quotaPeriodStart,proto3" json:"quota_period_start,omitempty"`
	QuotaPeriodEnd     *timestamppb.Timestamp `protobuf:"bytes,93,opt,name=quota_period_end,json=quotaPeriodEnd,proto3" json:"quota_period_end,omitempty"`
}

func (x *Space_Subscription) Reset() {
	*x = Space_Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_universe_space_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Space_Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Space_Subscription) ProtoMessage() {}

func (x *Space_Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_space_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Space_Subscription.ProtoReflect.Descriptor instead.
func (*Space_Subscription) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_space_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Space_Subscription) GetPlan() string {
	if x != nil {
		return x.Plan
	}
	return ""
}

func (x *Space_Subscription) GetSeats() uint32 {
	if x != nil {
		return x.Seats
	}
	return 0
}

func (x *Space_Subscription) GetQuota() *Quota {
	if x != nil {
		return x.Quota
	}
	return nil
}

func (x *Space_Subscription) GetBillingPeriodStart() *timestamppb.Timestamp {
	if x != nil {
		return x.BillingPeriodStart
	}
	return nil
}

func (x *Space_Subscription) GetBillingPeriodEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.BillingPeriodEnd
	}
	return nil
}

func (x *Space_Subscription) GetQuotaPeriodStart() *timestamppb.Timestamp {
	if x != nil {
		return x.QuotaPeriodStart
	}
	return nil
}

func (x *Space_Subscription) GetQuotaPeriodEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.QuotaPeriodEnd
	}
	return nil
}

var File_eolymp_universe_space_proto protoreflect.FileDescriptor

var file_eolymp_universe_space_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x1a, 0x17,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x63, 0x6c, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x08, 0x0a, 0x05, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x6d, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a,
	0x0b, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x34, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x48, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0xa0, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c,
	0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x64, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x8e, 0x03, 0x0a,
	0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6c, 0x61,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x05,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x4c, 0x0a, 0x14, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x5a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x12, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x48, 0x0a, 0x12, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x5b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x12, 0x48, 0x0a,
	0x12, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x5c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x5d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x22, 0x42, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54,
	0x52, 0x49, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10,
	0x03, 0x22, 0x3d, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49,
	0x43, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x02,
	0x22, 0x39, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x02, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x3b, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_universe_space_proto_rawDescOnce sync.Once
	file_eolymp_universe_space_proto_rawDescData = file_eolymp_universe_space_proto_rawDesc
)

func file_eolymp_universe_space_proto_rawDescGZIP() []byte {
	file_eolymp_universe_space_proto_rawDescOnce.Do(func() {
		file_eolymp_universe_space_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_universe_space_proto_rawDescData)
	})
	return file_eolymp_universe_space_proto_rawDescData
}

var file_eolymp_universe_space_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_eolymp_universe_space_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_universe_space_proto_goTypes = []any{
	(Space_Status)(0),             // 0: eolymp.universe.Space.Status
	(Space_Visibility)(0),         // 1: eolymp.universe.Space.Visibility
	(Space_Extra)(0),              // 2: eolymp.universe.Space.Extra
	(*Space)(nil),                 // 3: eolymp.universe.Space
	(*Space_Subscription)(nil),    // 4: eolymp.universe.Space.Subscription
	(acl.Action)(0),               // 5: eolymp.acl.Action
	(*Quota)(nil),                 // 6: eolymp.universe.Quota
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_eolymp_universe_space_proto_depIdxs = []int32{
	1, // 0: eolymp.universe.Space.visibility:type_name -> eolymp.universe.Space.Visibility
	0, // 1: eolymp.universe.Space.status:type_name -> eolymp.universe.Space.Status
	4, // 2: eolymp.universe.Space.subscription:type_name -> eolymp.universe.Space.Subscription
	5, // 3: eolymp.universe.Space.actions:type_name -> eolymp.acl.Action
	6, // 4: eolymp.universe.Space.Subscription.quota:type_name -> eolymp.universe.Quota
	7, // 5: eolymp.universe.Space.Subscription.billing_period_start:type_name -> google.protobuf.Timestamp
	7, // 6: eolymp.universe.Space.Subscription.billing_period_end:type_name -> google.protobuf.Timestamp
	7, // 7: eolymp.universe.Space.Subscription.quota_period_start:type_name -> google.protobuf.Timestamp
	7, // 8: eolymp.universe.Space.Subscription.quota_period_end:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_eolymp_universe_space_proto_init() }
func file_eolymp_universe_space_proto_init() {
	if File_eolymp_universe_space_proto != nil {
		return
	}
	file_eolymp_universe_quota_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_universe_space_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Space); i {
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
		file_eolymp_universe_space_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Space_Subscription); i {
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
			RawDescriptor: file_eolymp_universe_space_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_universe_space_proto_goTypes,
		DependencyIndexes: file_eolymp_universe_space_proto_depIdxs,
		EnumInfos:         file_eolymp_universe_space_proto_enumTypes,
		MessageInfos:      file_eolymp_universe_space_proto_msgTypes,
	}.Build()
	File_eolymp_universe_space_proto = out.File
	file_eolymp_universe_space_proto_rawDesc = nil
	file_eolymp_universe_space_proto_goTypes = nil
	file_eolymp_universe_space_proto_depIdxs = nil
}
