// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.18.1
// source: eolymp/universe/space.proto

package universe

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

type Space_Type int32

const (
	Space_UNKNOWN_TYPE Space_Type = 0
	Space_OTHER        Space_Type = 1
	Space_CLASSROOM    Space_Type = 2
	Space_TEAMROOM     Space_Type = 3
	Space_COMPETITION  Space_Type = 4
)

// Enum value maps for Space_Type.
var (
	Space_Type_name = map[int32]string{
		0: "UNKNOWN_TYPE",
		1: "OTHER",
		2: "CLASSROOM",
		3: "TEAMROOM",
		4: "COMPETITION",
	}
	Space_Type_value = map[string]int32{
		"UNKNOWN_TYPE": 0,
		"OTHER":        1,
		"CLASSROOM":    2,
		"TEAMROOM":     3,
		"COMPETITION":  4,
	}
)

func (x Space_Type) Enum() *Space_Type {
	p := new(Space_Type)
	*p = x
	return p
}

func (x Space_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Space_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_universe_space_proto_enumTypes[0].Descriptor()
}

func (Space_Type) Type() protoreflect.EnumType {
	return &file_eolymp_universe_space_proto_enumTypes[0]
}

func (x Space_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Space_Type.Descriptor instead.
func (Space_Type) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_universe_space_proto_rawDescGZIP(), []int{0, 0}
}

type Space_Visibility int32

const (
	Space_UNKNOWN_VISIBILITY Space_Visibility = 0
	Space_PUBLIC             Space_Visibility = 1
	Space_PRIVATE            Space_Visibility = 2
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

type Space struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                       // space unique identifier
	Key        string           `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`                                     // space key used to build URLs
	Url        string           `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`                                     // space key used to build URLs
	Name       string           `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`                                  // human friendly name
	Image      string           `protobuf:"bytes,11,opt,name=image,proto3" json:"image,omitempty"`                                // space logo image
	Type       Space_Type       `protobuf:"varint,12,opt,name=type,proto3,enum=eolymp.universe.Space_Type" json:"type,omitempty"` // space use type
	Plan       string           `protobuf:"bytes,13,opt,name=plan,proto3" json:"plan,omitempty"`                                  // plan defines billing plan for the space
	Visibility Space_Visibility `protobuf:"varint,14,opt,name=visibility,proto3,enum=eolymp.universe.Space_Visibility" json:"visibility,omitempty"`
	Quota      *Quota           `protobuf:"bytes,30,opt,name=quota,proto3" json:"quota,omitempty"`
	HomeUrl    string           `protobuf:"bytes,50,opt,name=home_url,json=homeUrl,proto3" json:"home_url,omitempty"`          // space home page URL
	IssuerUrl  string           `protobuf:"bytes,51,opt,name=issuer_url,json=issuerUrl,proto3" json:"issuer_url,omitempty"`    // space issuer URL (used for issuing tokens)
	GraphqlUrl string           `protobuf:"bytes,52,opt,name=graphql_url,json=graphqlUrl,proto3" json:"graphql_url,omitempty"` // space graphql endpoint
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

func (x *Space) GetType() Space_Type {
	if x != nil {
		return x.Type
	}
	return Space_UNKNOWN_TYPE
}

func (x *Space) GetPlan() string {
	if x != nil {
		return x.Plan
	}
	return ""
}

func (x *Space) GetVisibility() Space_Visibility {
	if x != nil {
		return x.Visibility
	}
	return Space_UNKNOWN_VISIBILITY
}

func (x *Space) GetQuota() *Quota {
	if x != nil {
		return x.Quota
	}
	return nil
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

var File_eolymp_universe_space_proto protoreflect.FileDescriptor

var file_eolymp_universe_space_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x1a, 0x1b,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x04, 0x0a, 0x05,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x41, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x6d,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x6d,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x34, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71,
	0x6c, 0x55, 0x72, 0x6c, 0x22, 0x51, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4c, 0x41,
	0x53, 0x53, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x45, 0x41, 0x4d,
	0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4d, 0x50, 0x45, 0x54,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x22, 0x3d, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49,
	0x56, 0x41, 0x54, 0x45, 0x10, 0x02, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x3b, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_eolymp_universe_space_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_universe_space_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_universe_space_proto_goTypes = []interface{}{
	(Space_Type)(0),       // 0: eolymp.universe.Space.Type
	(Space_Visibility)(0), // 1: eolymp.universe.Space.Visibility
	(*Space)(nil),         // 2: eolymp.universe.Space
	(*Quota)(nil),         // 3: eolymp.universe.Quota
}
var file_eolymp_universe_space_proto_depIdxs = []int32{
	0, // 0: eolymp.universe.Space.type:type_name -> eolymp.universe.Space.Type
	1, // 1: eolymp.universe.Space.visibility:type_name -> eolymp.universe.Space.Visibility
	3, // 2: eolymp.universe.Space.quota:type_name -> eolymp.universe.Quota
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_universe_space_proto_init() }
func file_eolymp_universe_space_proto_init() {
	if File_eolymp_universe_space_proto != nil {
		return
	}
	file_eolymp_universe_quota_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_universe_space_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_universe_space_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
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
