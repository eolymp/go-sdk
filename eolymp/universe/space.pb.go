// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/universe/space.proto

package universe

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
)

// Enum value maps for Space_Extra.
var (
	Space_Extra_name = map[int32]string{
		0: "UNKNOWN_EXTRA",
		1: "SUBSCRIPTION",
	}
	Space_Extra_value = map[string]int32{
		"UNKNOWN_EXTRA": 0,
		"SUBSCRIPTION":  1,
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

type Space_Feature int32

const (
	Space_UNKNOWN_FEATURE Space_Feature = 0
	Space_PRINTERS        Space_Feature = 1
	Space_NEWSLETTERS     Space_Feature = 2
)

// Enum value maps for Space_Feature.
var (
	Space_Feature_name = map[int32]string{
		0: "UNKNOWN_FEATURE",
		1: "PRINTERS",
		2: "NEWSLETTERS",
	}
	Space_Feature_value = map[string]int32{
		"UNKNOWN_FEATURE": 0,
		"PRINTERS":        1,
		"NEWSLETTERS":     2,
	}
)

func (x Space_Feature) Enum() *Space_Feature {
	p := new(Space_Feature)
	*p = x
	return p
}

func (x Space_Feature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Space_Feature) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_universe_space_proto_enumTypes[3].Descriptor()
}

func (Space_Feature) Type() protoreflect.EnumType {
	return &file_eolymp_universe_space_proto_enumTypes[3]
}

func (x Space_Feature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Space_Feature.Descriptor instead.
func (Space_Feature) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_universe_space_proto_rawDescGZIP(), []int{0, 3}
}

type Space struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                    // space unique identifier
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`                                  // space key used to build URLs
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`                                  // space url
	HomeUrl       string                 `protobuf:"bytes,50,opt,name=home_url,json=homeUrl,proto3" json:"home_url,omitempty"`          // space home page URL
	IssuerUrl     string                 `protobuf:"bytes,51,opt,name=issuer_url,json=issuerUrl,proto3" json:"issuer_url,omitempty"`    // space issuer URL (used for issuing tokens)
	GraphqlUrl    string                 `protobuf:"bytes,52,opt,name=graphql_url,json=graphqlUrl,proto3" json:"graphql_url,omitempty"` // space graphql endpoint
	Name          string                 `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`                               // human friendly name
	Image         string                 `protobuf:"bytes,11,opt,name=image,proto3" json:"image,omitempty"`                             // space logo image
	Visibility    Space_Visibility       `protobuf:"varint,14,opt,name=visibility,proto3,enum=eolymp.universe.Space_Visibility" json:"visibility,omitempty"`
	Features      []Space_Feature        `protobuf:"varint,17,rep,packed,name=features,proto3,enum=eolymp.universe.Space_Feature" json:"features,omitempty"`
	Status        Space_Status           `protobuf:"varint,16,opt,name=status,proto3,enum=eolymp.universe.Space_Status" json:"status,omitempty"` // space status
	Subscription  *Space_Subscription    `protobuf:"bytes,800,opt,name=subscription,proto3" json:"subscription,omitempty"`                       // subscription details (private)
	Affiliation   string                 `protobuf:"bytes,15,opt,name=affiliation,proto3" json:"affiliation,omitempty"`                          // space affiliation label
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Space) Reset() {
	*x = Space{}
	mi := &file_eolymp_universe_space_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Space) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Space) ProtoMessage() {}

func (x *Space) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_space_proto_msgTypes[0]
	if x != nil {
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

func (x *Space) GetFeatures() []Space_Feature {
	if x != nil {
		return x.Features
	}
	return nil
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

type Space_Subscription struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Plan               string                 `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	Seats              uint32                 `protobuf:"varint,2,opt,name=seats,proto3" json:"seats,omitempty"`
	Quota              *Quota                 `protobuf:"bytes,10,opt,name=quota,proto3" json:"quota,omitempty"`
	BillingPeriodStart *timestamppb.Timestamp `protobuf:"bytes,90,opt,name=billing_period_start,json=billingPeriodStart,proto3" json:"billing_period_start,omitempty"`
	BillingPeriodEnd   *timestamppb.Timestamp `protobuf:"bytes,91,opt,name=billing_period_end,json=billingPeriodEnd,proto3" json:"billing_period_end,omitempty"`
	QuotaPeriodStart   *timestamppb.Timestamp `protobuf:"bytes,92,opt,name=quota_period_start,json=quotaPeriodStart,proto3" json:"quota_period_start,omitempty"`
	QuotaPeriodEnd     *timestamppb.Timestamp `protobuf:"bytes,93,opt,name=quota_period_end,json=quotaPeriodEnd,proto3" json:"quota_period_end,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Space_Subscription) Reset() {
	*x = Space_Subscription{}
	mi := &file_eolymp_universe_space_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Space_Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Space_Subscription) ProtoMessage() {}

func (x *Space_Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_space_proto_msgTypes[1]
	if x != nil {
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

const file_eolymp_universe_space_proto_rawDesc = "" +
	"\n" +
	"\x1beolymp/universe/space.proto\x12\x0feolymp.universe\x1a\x1beolymp/universe/quota.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe3\b\n" +
	"\x05Space\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x10\n" +
	"\x03key\x18\x02 \x01(\tR\x03key\x12\x10\n" +
	"\x03url\x18\x03 \x01(\tR\x03url\x12\x19\n" +
	"\bhome_url\x182 \x01(\tR\ahomeUrl\x12\x1d\n" +
	"\n" +
	"issuer_url\x183 \x01(\tR\tissuerUrl\x12\x1f\n" +
	"\vgraphql_url\x184 \x01(\tR\n" +
	"graphqlUrl\x12\x12\n" +
	"\x04name\x18\n" +
	" \x01(\tR\x04name\x12\x14\n" +
	"\x05image\x18\v \x01(\tR\x05image\x12A\n" +
	"\n" +
	"visibility\x18\x0e \x01(\x0e2!.eolymp.universe.Space.VisibilityR\n" +
	"visibility\x12:\n" +
	"\bfeatures\x18\x11 \x03(\x0e2\x1e.eolymp.universe.Space.FeatureR\bfeatures\x125\n" +
	"\x06status\x18\x10 \x01(\x0e2\x1d.eolymp.universe.Space.StatusR\x06status\x12H\n" +
	"\fsubscription\x18\xa0\x06 \x01(\v2#.eolymp.universe.Space.SubscriptionR\fsubscription\x12 \n" +
	"\vaffiliation\x18\x0f \x01(\tR\vaffiliation\x1a\x8e\x03\n" +
	"\fSubscription\x12\x12\n" +
	"\x04plan\x18\x01 \x01(\tR\x04plan\x12\x14\n" +
	"\x05seats\x18\x02 \x01(\rR\x05seats\x12,\n" +
	"\x05quota\x18\n" +
	" \x01(\v2\x16.eolymp.universe.QuotaR\x05quota\x12L\n" +
	"\x14billing_period_start\x18Z \x01(\v2\x1a.google.protobuf.TimestampR\x12billingPeriodStart\x12H\n" +
	"\x12billing_period_end\x18[ \x01(\v2\x1a.google.protobuf.TimestampR\x10billingPeriodEnd\x12H\n" +
	"\x12quota_period_start\x18\\ \x01(\v2\x1a.google.protobuf.TimestampR\x10quotaPeriodStart\x12D\n" +
	"\x10quota_period_end\x18] \x01(\v2\x1a.google.protobuf.TimestampR\x0equotaPeriodEnd\"B\n" +
	"\x06Status\x12\x12\n" +
	"\x0eUNKNOWN_STATUS\x10\x00\x12\t\n" +
	"\x05TRIAL\x10\x01\x12\n" +
	"\n" +
	"\x06ACTIVE\x10\x02\x12\r\n" +
	"\tSUSPENDED\x10\x03\"=\n" +
	"\n" +
	"Visibility\x12\x16\n" +
	"\x12UNKNOWN_VISIBILITY\x10\x00\x12\n" +
	"\n" +
	"\x06PUBLIC\x10\x01\x12\v\n" +
	"\aPRIVATE\x10\x02\",\n" +
	"\x05Extra\x12\x11\n" +
	"\rUNKNOWN_EXTRA\x10\x00\x12\x10\n" +
	"\fSUBSCRIPTION\x10\x01\"=\n" +
	"\aFeature\x12\x13\n" +
	"\x0fUNKNOWN_FEATURE\x10\x00\x12\f\n" +
	"\bPRINTERS\x10\x01\x12\x0f\n" +
	"\vNEWSLETTERS\x10\x02B3Z1github.com/eolymp/go-sdk/eolymp/universe;universeb\x06proto3"

var (
	file_eolymp_universe_space_proto_rawDescOnce sync.Once
	file_eolymp_universe_space_proto_rawDescData []byte
)

func file_eolymp_universe_space_proto_rawDescGZIP() []byte {
	file_eolymp_universe_space_proto_rawDescOnce.Do(func() {
		file_eolymp_universe_space_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_universe_space_proto_rawDesc), len(file_eolymp_universe_space_proto_rawDesc)))
	})
	return file_eolymp_universe_space_proto_rawDescData
}

var file_eolymp_universe_space_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_eolymp_universe_space_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_universe_space_proto_goTypes = []any{
	(Space_Status)(0),             // 0: eolymp.universe.Space.Status
	(Space_Visibility)(0),         // 1: eolymp.universe.Space.Visibility
	(Space_Extra)(0),              // 2: eolymp.universe.Space.Extra
	(Space_Feature)(0),            // 3: eolymp.universe.Space.Feature
	(*Space)(nil),                 // 4: eolymp.universe.Space
	(*Space_Subscription)(nil),    // 5: eolymp.universe.Space.Subscription
	(*Quota)(nil),                 // 6: eolymp.universe.Quota
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_eolymp_universe_space_proto_depIdxs = []int32{
	1, // 0: eolymp.universe.Space.visibility:type_name -> eolymp.universe.Space.Visibility
	3, // 1: eolymp.universe.Space.features:type_name -> eolymp.universe.Space.Feature
	0, // 2: eolymp.universe.Space.status:type_name -> eolymp.universe.Space.Status
	5, // 3: eolymp.universe.Space.subscription:type_name -> eolymp.universe.Space.Subscription
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_universe_space_proto_rawDesc), len(file_eolymp_universe_space_proto_rawDesc)),
			NumEnums:      4,
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
	file_eolymp_universe_space_proto_goTypes = nil
	file_eolymp_universe_space_proto_depIdxs = nil
}
