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
	Action_UNKNOWN_ACTION                   Action = 0
	Action_DESCRIBE_SPACE_QUOTA             Action = 10 // view quota information
	Action_DESCRIBE_SPACE_CONFIG            Action = 11 // view quota information
	Action_DESCRIBE_SPACE_BILLING           Action = 12
	Action_UPDATE_SPACE                     Action = 13 // modify space parameters
	Action_UPDATE_SPACE_CONFIG              Action = 14
	Action_UPDATE_SPACE_BILLING             Action = 15
	Action_DELETE_SPACE                     Action = 16 // delete space
	Action_DESCRIBE_POLICY                  Action = 20 // view/list space permissions
	Action_UPDATE_POLICY                    Action = 21 // create/update space permissions
	Action_CREATE_PROBLEM                   Action = 30 // create problems
	Action_DESCRIBE_PROBLEM                 Action = 31 // view problems (excl. invisible)
	Action_DESCRIBE_PROBLEM_EDITORIAL       Action = 32 // view testing configuration for problems (incl. tests)
	Action_DESCRIBE_PROBLEM_HISTORY         Action = 33 // view testing configuration for problems (incl. tests)
	Action_DESCRIBE_PROBLEM_TESTING         Action = 34 // view testing configuration for problems (incl. tests)
	Action_DESCRIBE_PROBLEM_SUBMISSION      Action = 35 // view testing configuration for problems (incl. tests)
	Action_UPDATE_PROBLEM                   Action = 36 // modify problems (except visibility and testing config)
	Action_UPDATE_PROBLEM_EDITORIAL         Action = 37 // modify problems (except visibility and testing config)
	Action_UPDATE_PROBLEM_VISIBILITY        Action = 38 // modify visibility for problems
	Action_UPDATE_PROBLEM_TESTING           Action = 39 // modify testing configuration for problems
	Action_DELETE_PROBLEM                   Action = 40 // delete problems
	Action_SUBMIT_PROBLEM                   Action = 41 // submit problem (create submissions)
	Action_RETEST_PROBLEM                   Action = 42 // retest problem submissions
	Action_CREATE_CONTEST                   Action = 50
	Action_DESCRIBE_CONTEST                 Action = 51
	Action_DESCRIBE_CONTEST_ANNOUNCEMENT    Action = 56
	Action_DESCRIBE_CONTEST_PARTICIPANT     Action = 57
	Action_DESCRIBE_CONTEST_RESULT          Action = 58
	Action_DESCRIBE_CONTEST_RESULT_REALTIME Action = 59
	Action_DESCRIBE_CONTEST_PROBLEM         Action = 60
	Action_DESCRIBE_CONTEST_SUBMISSION      Action = 61
	Action_UPDATE_CONTEST                   Action = 62
	Action_UPDATE_CONTEST_ANNOUNCEMENT      Action = 63
	Action_UPDATE_CONTEST_PARTICIPANT       Action = 64
	Action_UPDATE_CONTEST_RESULT            Action = 65
	Action_UPDATE_CONTEST_PROBLEM           Action = 66
	Action_UPDATE_CONTEST_SUBMISSION        Action = 67
	Action_SUBMIT_CONTEST_PROBLEM           Action = 68
	Action_DELETE_CONTEST                   Action = 69
	Action_DESCRIBE_TICKET                  Action = 70
	Action_UPDATE_TICKET                    Action = 71
	Action_REPLY_TICKET                     Action = 72
	Action_DESCRIBE_RANKING                 Action = 80
	Action_UPDATE_RANKING                   Action = 81
	Action_CREATE_MEMBER                    Action = 90
	Action_DESCRIBE_MEMBER                  Action = 91
	Action_UPDATE_MEMBER                    Action = 92
	Action_DELETE_MEMBER                    Action = 93
	Action_CREATE_GROUP                     Action = 100
	Action_DESCRIBE_GROUP                   Action = 101
	Action_UPDATE_GROUP                     Action = 102
	Action_DELETE_GROUP                     Action = 103
	Action_CREATE_SCOREBOARD                Action = 110
	Action_DESCRIBE_SCOREBOARD              Action = 111
	Action_DESCRIBE_SCOREBOARD_REALTIME     Action = 112
	Action_UPDATE_SCOREBOARD                Action = 113
	Action_DELETE_SCOREBOARD                Action = 114
	Action_DESCRIBE_CONTENT                 Action = 120
	Action_UPDATE_CONTENT                   Action = 121
	Action_CREATE_COURSE                    Action = 130
	Action_DESCRIBE_COURSE                  Action = 131
	Action_DESCRIBE_COURSE_ENTRY            Action = 132
	Action_DESCRIBE_COURSE_STUDENT          Action = 133
	Action_DESCRIBE_COURSE_ASSIGNMENT       Action = 134
	Action_UPDATE_COURSE                    Action = 135
	Action_UPDATE_COURSE_STUDENT            Action = 136
	Action_UPDATE_COURSE_ASSIGNMENT         Action = 137
	Action_DELETE_COURSE                    Action = 138
	Action_CREATE_POST                      Action = 140
	Action_DESCRIBE_POST                    Action = 141
	Action_UPDATE_POST                      Action = 142
	Action_DELETE_POST                      Action = 143
	Action_MODERATE_POST                    Action = 144
	Action_UPDATE_MESSAGE                   Action = 150
	Action_DELETE_MESSAGE                   Action = 151
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0:   "UNKNOWN_ACTION",
		10:  "DESCRIBE_SPACE_QUOTA",
		11:  "DESCRIBE_SPACE_CONFIG",
		12:  "DESCRIBE_SPACE_BILLING",
		13:  "UPDATE_SPACE",
		14:  "UPDATE_SPACE_CONFIG",
		15:  "UPDATE_SPACE_BILLING",
		16:  "DELETE_SPACE",
		20:  "DESCRIBE_POLICY",
		21:  "UPDATE_POLICY",
		30:  "CREATE_PROBLEM",
		31:  "DESCRIBE_PROBLEM",
		32:  "DESCRIBE_PROBLEM_EDITORIAL",
		33:  "DESCRIBE_PROBLEM_HISTORY",
		34:  "DESCRIBE_PROBLEM_TESTING",
		35:  "DESCRIBE_PROBLEM_SUBMISSION",
		36:  "UPDATE_PROBLEM",
		37:  "UPDATE_PROBLEM_EDITORIAL",
		38:  "UPDATE_PROBLEM_VISIBILITY",
		39:  "UPDATE_PROBLEM_TESTING",
		40:  "DELETE_PROBLEM",
		41:  "SUBMIT_PROBLEM",
		42:  "RETEST_PROBLEM",
		50:  "CREATE_CONTEST",
		51:  "DESCRIBE_CONTEST",
		56:  "DESCRIBE_CONTEST_ANNOUNCEMENT",
		57:  "DESCRIBE_CONTEST_PARTICIPANT",
		58:  "DESCRIBE_CONTEST_RESULT",
		59:  "DESCRIBE_CONTEST_RESULT_REALTIME",
		60:  "DESCRIBE_CONTEST_PROBLEM",
		61:  "DESCRIBE_CONTEST_SUBMISSION",
		62:  "UPDATE_CONTEST",
		63:  "UPDATE_CONTEST_ANNOUNCEMENT",
		64:  "UPDATE_CONTEST_PARTICIPANT",
		65:  "UPDATE_CONTEST_RESULT",
		66:  "UPDATE_CONTEST_PROBLEM",
		67:  "UPDATE_CONTEST_SUBMISSION",
		68:  "SUBMIT_CONTEST_PROBLEM",
		69:  "DELETE_CONTEST",
		70:  "DESCRIBE_TICKET",
		71:  "UPDATE_TICKET",
		72:  "REPLY_TICKET",
		80:  "DESCRIBE_RANKING",
		81:  "UPDATE_RANKING",
		90:  "CREATE_MEMBER",
		91:  "DESCRIBE_MEMBER",
		92:  "UPDATE_MEMBER",
		93:  "DELETE_MEMBER",
		100: "CREATE_GROUP",
		101: "DESCRIBE_GROUP",
		102: "UPDATE_GROUP",
		103: "DELETE_GROUP",
		110: "CREATE_SCOREBOARD",
		111: "DESCRIBE_SCOREBOARD",
		112: "DESCRIBE_SCOREBOARD_REALTIME",
		113: "UPDATE_SCOREBOARD",
		114: "DELETE_SCOREBOARD",
		120: "DESCRIBE_CONTENT",
		121: "UPDATE_CONTENT",
		130: "CREATE_COURSE",
		131: "DESCRIBE_COURSE",
		132: "DESCRIBE_COURSE_ENTRY",
		133: "DESCRIBE_COURSE_STUDENT",
		134: "DESCRIBE_COURSE_ASSIGNMENT",
		135: "UPDATE_COURSE",
		136: "UPDATE_COURSE_STUDENT",
		137: "UPDATE_COURSE_ASSIGNMENT",
		138: "DELETE_COURSE",
		140: "CREATE_POST",
		141: "DESCRIBE_POST",
		142: "UPDATE_POST",
		143: "DELETE_POST",
		144: "MODERATE_POST",
		150: "UPDATE_MESSAGE",
		151: "DELETE_MESSAGE",
	}
	Action_value = map[string]int32{
		"UNKNOWN_ACTION":                   0,
		"DESCRIBE_SPACE_QUOTA":             10,
		"DESCRIBE_SPACE_CONFIG":            11,
		"DESCRIBE_SPACE_BILLING":           12,
		"UPDATE_SPACE":                     13,
		"UPDATE_SPACE_CONFIG":              14,
		"UPDATE_SPACE_BILLING":             15,
		"DELETE_SPACE":                     16,
		"DESCRIBE_POLICY":                  20,
		"UPDATE_POLICY":                    21,
		"CREATE_PROBLEM":                   30,
		"DESCRIBE_PROBLEM":                 31,
		"DESCRIBE_PROBLEM_EDITORIAL":       32,
		"DESCRIBE_PROBLEM_HISTORY":         33,
		"DESCRIBE_PROBLEM_TESTING":         34,
		"DESCRIBE_PROBLEM_SUBMISSION":      35,
		"UPDATE_PROBLEM":                   36,
		"UPDATE_PROBLEM_EDITORIAL":         37,
		"UPDATE_PROBLEM_VISIBILITY":        38,
		"UPDATE_PROBLEM_TESTING":           39,
		"DELETE_PROBLEM":                   40,
		"SUBMIT_PROBLEM":                   41,
		"RETEST_PROBLEM":                   42,
		"CREATE_CONTEST":                   50,
		"DESCRIBE_CONTEST":                 51,
		"DESCRIBE_CONTEST_ANNOUNCEMENT":    56,
		"DESCRIBE_CONTEST_PARTICIPANT":     57,
		"DESCRIBE_CONTEST_RESULT":          58,
		"DESCRIBE_CONTEST_RESULT_REALTIME": 59,
		"DESCRIBE_CONTEST_PROBLEM":         60,
		"DESCRIBE_CONTEST_SUBMISSION":      61,
		"UPDATE_CONTEST":                   62,
		"UPDATE_CONTEST_ANNOUNCEMENT":      63,
		"UPDATE_CONTEST_PARTICIPANT":       64,
		"UPDATE_CONTEST_RESULT":            65,
		"UPDATE_CONTEST_PROBLEM":           66,
		"UPDATE_CONTEST_SUBMISSION":        67,
		"SUBMIT_CONTEST_PROBLEM":           68,
		"DELETE_CONTEST":                   69,
		"DESCRIBE_TICKET":                  70,
		"UPDATE_TICKET":                    71,
		"REPLY_TICKET":                     72,
		"DESCRIBE_RANKING":                 80,
		"UPDATE_RANKING":                   81,
		"CREATE_MEMBER":                    90,
		"DESCRIBE_MEMBER":                  91,
		"UPDATE_MEMBER":                    92,
		"DELETE_MEMBER":                    93,
		"CREATE_GROUP":                     100,
		"DESCRIBE_GROUP":                   101,
		"UPDATE_GROUP":                     102,
		"DELETE_GROUP":                     103,
		"CREATE_SCOREBOARD":                110,
		"DESCRIBE_SCOREBOARD":              111,
		"DESCRIBE_SCOREBOARD_REALTIME":     112,
		"UPDATE_SCOREBOARD":                113,
		"DELETE_SCOREBOARD":                114,
		"DESCRIBE_CONTENT":                 120,
		"UPDATE_CONTENT":                   121,
		"CREATE_COURSE":                    130,
		"DESCRIBE_COURSE":                  131,
		"DESCRIBE_COURSE_ENTRY":            132,
		"DESCRIBE_COURSE_STUDENT":          133,
		"DESCRIBE_COURSE_ASSIGNMENT":       134,
		"UPDATE_COURSE":                    135,
		"UPDATE_COURSE_STUDENT":            136,
		"UPDATE_COURSE_ASSIGNMENT":         137,
		"DELETE_COURSE":                    138,
		"CREATE_POST":                      140,
		"DESCRIBE_POST":                    141,
		"UPDATE_POST":                      142,
		"DELETE_POST":                      143,
		"MODERATE_POST":                    144,
		"UPDATE_MESSAGE":                   150,
		"DELETE_MESSAGE":                   151,
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
	0x70, 0x2e, 0x61, 0x63, 0x6c, 0x2a, 0x9e, 0x0e, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45,
	0x5f, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x41, 0x10, 0x0a, 0x12, 0x19,
	0x0a, 0x15, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x53, 0x50, 0x41, 0x43, 0x45,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x0b, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45, 0x53,
	0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x42, 0x49, 0x4c, 0x4c,
	0x49, 0x4e, 0x47, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x0d, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x0e,
	0x12, 0x18, 0x0a, 0x14, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x50, 0x41, 0x43, 0x45,
	0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x5f, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x10, 0x12, 0x13, 0x0a, 0x0f,
	0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10,
	0x14, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49,
	0x43, 0x59, 0x10, 0x15, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x50,
	0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x1e, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x45, 0x53, 0x43,
	0x52, 0x49, 0x42, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x1f, 0x12, 0x1e,
	0x0a, 0x1a, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c,
	0x45, 0x4d, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x49, 0x41, 0x4c, 0x10, 0x20, 0x12, 0x1c,
	0x0a, 0x18, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c,
	0x45, 0x4d, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x21, 0x12, 0x1c, 0x0a, 0x18,
	0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d,
	0x5f, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x22, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45,
	0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f, 0x53,
	0x55, 0x42, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x23, 0x12, 0x12, 0x0a, 0x0e, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x24, 0x12,
	0x1c, 0x0a, 0x18, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45,
	0x4d, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x49, 0x41, 0x4c, 0x10, 0x25, 0x12, 0x1d, 0x0a,
	0x19, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f,
	0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x26, 0x12, 0x1a, 0x0a, 0x16,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f, 0x54,
	0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x27, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x28, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x29,
	0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c,
	0x45, 0x4d, 0x10, 0x2a, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x43,
	0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x10, 0x32, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x45, 0x53, 0x43,
	0x52, 0x49, 0x42, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x10, 0x33, 0x12, 0x21,
	0x0a, 0x1d, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45,
	0x53, 0x54, 0x5f, 0x41, 0x4e, 0x4e, 0x4f, 0x55, 0x4e, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x38, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x43, 0x4f,
	0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x43, 0x49, 0x50, 0x41, 0x4e,
	0x54, 0x10, 0x39, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x3a,
	0x12, 0x24, 0x0a, 0x20, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x4c,
	0x54, 0x49, 0x4d, 0x45, 0x10, 0x3b, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49,
	0x42, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c,
	0x45, 0x4d, 0x10, 0x3c, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45,
	0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x10, 0x3d, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x10, 0x3e, 0x12, 0x1f, 0x0a, 0x1b, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x4e, 0x4e, 0x4f,
	0x55, 0x4e, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x3f, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x50, 0x41, 0x52,
	0x54, 0x49, 0x43, 0x49, 0x50, 0x41, 0x4e, 0x54, 0x10, 0x40, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x10, 0x41, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10,
	0x42, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54,
	0x45, 0x53, 0x54, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x43,
	0x12, 0x1a, 0x0a, 0x16, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45,
	0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x44, 0x12, 0x12, 0x0a, 0x0e,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x10, 0x45,
	0x12, 0x13, 0x0a, 0x0f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x54, 0x49, 0x43,
	0x4b, 0x45, 0x54, 0x10, 0x46, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x47, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x50, 0x4c,
	0x59, 0x5f, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x48, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x45,
	0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x50,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x4b, 0x49,
	0x4e, 0x47, 0x10, 0x51, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x4d,
	0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x5a, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x45, 0x53, 0x43, 0x52,
	0x49, 0x42, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x5b, 0x12, 0x11, 0x0a, 0x0d,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x5c, 0x12,
	0x11, 0x0a, 0x0d, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52,
	0x10, 0x5d, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x10, 0x64, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x66, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x67, 0x12, 0x15, 0x0a, 0x11,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x42, 0x4f, 0x41, 0x52,
	0x44, 0x10, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f,
	0x53, 0x43, 0x4f, 0x52, 0x45, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x6f, 0x12, 0x20, 0x0a, 0x1c,
	0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x42, 0x4f,
	0x41, 0x52, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x4c, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x70, 0x12, 0x15,
	0x0a, 0x11, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x42, 0x4f,
	0x41, 0x52, 0x44, 0x10, 0x71, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f,
	0x53, 0x43, 0x4f, 0x52, 0x45, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x72, 0x12, 0x14, 0x0a, 0x10,
	0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54,
	0x10, 0x78, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x45, 0x4e, 0x54, 0x10, 0x79, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x10, 0x82, 0x01, 0x12, 0x14, 0x0a, 0x0f, 0x44, 0x45,
	0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x10, 0x83, 0x01,
	0x12, 0x1a, 0x0a, 0x15, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x43, 0x4f, 0x55,
	0x52, 0x53, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x84, 0x01, 0x12, 0x1c, 0x0a, 0x17,
	0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f,
	0x53, 0x54, 0x55, 0x44, 0x45, 0x4e, 0x54, 0x10, 0x85, 0x01, 0x12, 0x1f, 0x0a, 0x1a, 0x44, 0x45,
	0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x41, 0x53,
	0x53, 0x49, 0x47, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x86, 0x01, 0x12, 0x12, 0x0a, 0x0d, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x10, 0x87, 0x01, 0x12,
	0x1a, 0x0a, 0x15, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45,
	0x5f, 0x53, 0x54, 0x55, 0x44, 0x45, 0x4e, 0x54, 0x10, 0x88, 0x01, 0x12, 0x1d, 0x0a, 0x18, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x41, 0x53, 0x53,
	0x49, 0x47, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x89, 0x01, 0x12, 0x12, 0x0a, 0x0d, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x10, 0x8a, 0x01, 0x12, 0x10,
	0x0a, 0x0b, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x8c, 0x01,
	0x12, 0x12, 0x0a, 0x0d, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x50, 0x4f, 0x53,
	0x54, 0x10, 0x8d, 0x01, 0x12, 0x10, 0x0a, 0x0b, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x50,
	0x4f, 0x53, 0x54, 0x10, 0x8e, 0x01, 0x12, 0x10, 0x0a, 0x0b, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x5f, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x8f, 0x01, 0x12, 0x12, 0x0a, 0x0d, 0x4d, 0x4f, 0x44, 0x45,
	0x52, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x90, 0x01, 0x12, 0x13, 0x0a, 0x0e,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x96,
	0x01, 0x12, 0x13, 0x0a, 0x0e, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x10, 0x97, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x63, 0x6c, 0x3b, 0x61, 0x63,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
