// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/acl/action.proto

package acl

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

type Action int32

const (
	Action_UNKNOWN_ACTION   Action = 0
	Action_SPACE_READ       Action = 11
	Action_SPACE_WRITE      Action = 13
	Action_SPACE_DELETE     Action = 16
	Action_BILLING_READ     Action = 12
	Action_BILLING_WRITE    Action = 15
	Action_POLICY_READ      Action = 20 // view/list space permissions
	Action_POLICY_WRITE     Action = 21 // create/update space permissions
	Action_PROBLEM_READ     Action = 31
	Action_PROBLEM_WRITE    Action = 36
	Action_PROBLEM_TESTING  Action = 34 // access to testing configuration, tests, solutions, submissions etc
	Action_CONTEST_READ     Action = 51
	Action_CONTEST_WRITE    Action = 62
	Action_TICKET_READ      Action = 70
	Action_TICKET_WRITE     Action = 71 // modify tickets: reply, close etc
	Action_MEMBER_READ      Action = 91
	Action_MEMBER_WRITE     Action = 92
	Action_SCOREBOARD_READ  Action = 111
	Action_SCOREBOARD_WRITE Action = 112
	Action_CONTENT_READ     Action = 120
	Action_CONTENT_WRITE    Action = 121
	Action_COURSE_READ      Action = 130
	Action_COURSE_WRITE     Action = 131
	Action_COURSE_ASSIGN    Action = 132 // limited write permission to add students to the course and manage assignments
	Action_POST_READ        Action = 140
	Action_POST_WRITE       Action = 141
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0:   "UNKNOWN_ACTION",
		11:  "SPACE_READ",
		13:  "SPACE_WRITE",
		16:  "SPACE_DELETE",
		12:  "BILLING_READ",
		15:  "BILLING_WRITE",
		20:  "POLICY_READ",
		21:  "POLICY_WRITE",
		31:  "PROBLEM_READ",
		36:  "PROBLEM_WRITE",
		34:  "PROBLEM_TESTING",
		51:  "CONTEST_READ",
		62:  "CONTEST_WRITE",
		70:  "TICKET_READ",
		71:  "TICKET_WRITE",
		91:  "MEMBER_READ",
		92:  "MEMBER_WRITE",
		111: "SCOREBOARD_READ",
		112: "SCOREBOARD_WRITE",
		120: "CONTENT_READ",
		121: "CONTENT_WRITE",
		130: "COURSE_READ",
		131: "COURSE_WRITE",
		132: "COURSE_ASSIGN",
		140: "POST_READ",
		141: "POST_WRITE",
	}
	Action_value = map[string]int32{
		"UNKNOWN_ACTION":   0,
		"SPACE_READ":       11,
		"SPACE_WRITE":      13,
		"SPACE_DELETE":     16,
		"BILLING_READ":     12,
		"BILLING_WRITE":    15,
		"POLICY_READ":      20,
		"POLICY_WRITE":     21,
		"PROBLEM_READ":     31,
		"PROBLEM_WRITE":    36,
		"PROBLEM_TESTING":  34,
		"CONTEST_READ":     51,
		"CONTEST_WRITE":    62,
		"TICKET_READ":      70,
		"TICKET_WRITE":     71,
		"MEMBER_READ":      91,
		"MEMBER_WRITE":     92,
		"SCOREBOARD_READ":  111,
		"SCOREBOARD_WRITE": 112,
		"CONTENT_READ":     120,
		"CONTENT_WRITE":    121,
		"COURSE_READ":      130,
		"COURSE_WRITE":     131,
		"COURSE_ASSIGN":    132,
		"POST_READ":        140,
		"POST_WRITE":       141,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_acl_action_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_eolymp_acl_action_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_acl_action_proto_rawDescGZIP(), []int{0}
}

var File_eolymp_acl_action_proto protoreflect.FileDescriptor

var file_eolymp_acl_action_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x63, 0x6c, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x61, 0x63, 0x6c, 0x2a, 0xe6, 0x03, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x52, 0x45,
	0x41, 0x44, 0x10, 0x0b, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x57, 0x52,
	0x49, 0x54, 0x45, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x10, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x49, 0x4c, 0x4c, 0x49,
	0x4e, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x0c, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x49, 0x4c,
	0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x0f, 0x12, 0x0f, 0x0a, 0x0b,
	0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x14, 0x12, 0x10, 0x0a,
	0x0c, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x15, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10,
	0x1f, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x10, 0x24, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x22, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x33, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x3e, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x46, 0x12,
	0x10, 0x0a, 0x0c, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10,
	0x47, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x44,
	0x10, 0x5b, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x10, 0x5c, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x42, 0x4f, 0x41,
	0x52, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x6f, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x43, 0x4f,
	0x52, 0x45, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x70, 0x12,
	0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10,
	0x78, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x10, 0x79, 0x12, 0x10, 0x0a, 0x0b, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x52,
	0x45, 0x41, 0x44, 0x10, 0x82, 0x01, 0x12, 0x11, 0x0a, 0x0c, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45,
	0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x83, 0x01, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x4f, 0x55,
	0x52, 0x53, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x10, 0x84, 0x01, 0x12, 0x0e, 0x0a,
	0x09, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x8c, 0x01, 0x12, 0x0f, 0x0a,
	0x0a, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x8d, 0x01, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2f, 0x61, 0x63, 0x6c, 0x3b, 0x61, 0x63, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eolymp_acl_action_proto_rawDescOnce sync.Once
	file_eolymp_acl_action_proto_rawDescData = file_eolymp_acl_action_proto_rawDesc
)

func file_eolymp_acl_action_proto_rawDescGZIP() []byte {
	file_eolymp_acl_action_proto_rawDescOnce.Do(func() {
		file_eolymp_acl_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_acl_action_proto_rawDescData)
	})
	return file_eolymp_acl_action_proto_rawDescData
}

var file_eolymp_acl_action_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_acl_action_proto_goTypes = []any{
	(Action)(0), // 0: eolymp.acl.Action
}
var file_eolymp_acl_action_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_acl_action_proto_init() }
func file_eolymp_acl_action_proto_init() {
	if File_eolymp_acl_action_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_acl_action_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_acl_action_proto_goTypes,
		DependencyIndexes: file_eolymp_acl_action_proto_depIdxs,
		EnumInfos:         file_eolymp_acl_action_proto_enumTypes,
	}.Build()
	File_eolymp_acl_action_proto = out.File
	file_eolymp_acl_action_proto_rawDesc = nil
	file_eolymp_acl_action_proto_goTypes = nil
	file_eolymp_acl_action_proto_depIdxs = nil
}
