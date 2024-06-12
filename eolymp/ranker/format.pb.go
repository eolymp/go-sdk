// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/ranker/format.proto

package ranker

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

type Format int32

const (
	// Reserved, should not be used.
	Format_NONE Format = 0
	// IOI scoring mode ranks participants by total number of points scored in each problem.
	//
	// - Total score is a sum of scores in each problem
	// - Problem score is maximum number of points scored among all submissions.
	Format_IOI Format = 1
	// ICPC scoring mode first ranks participants by number of solved problems and then ranks those who solved
	// problems faster higher. Additional penalty is added for each unsuccessful attempt to solve problem.
	//
	//   - Each solved problem gives 1 score point.
	//   - Once problem is solved, a penalty point is awarded for each minute it took to solve problem.
	//   - Once problem is solved, 20 penalty points are awarded for attempt to solve problem. Attempts are compiled
	//     submissions made before one which solves problem.
	//   - Unsolved problem do not add penalty.
	//   - Submissions made after problem is solved do not count towards attempts.
	//   - Submissions which were not compiled do not count towards attempts.
	Format_ICPC Format = 2
)

// Enum value maps for Format.
var (
	Format_name = map[int32]string{
		0: "NONE",
		1: "IOI",
		2: "ICPC",
	}
	Format_value = map[string]int32{
		"NONE": 0,
		"IOI":  1,
		"ICPC": 2,
	}
)

func (x Format) Enum() *Format {
	p := new(Format)
	*p = x
	return p
}

func (x Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Format) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_ranker_format_proto_enumTypes[0].Descriptor()
}

func (Format) Type() protoreflect.EnumType {
	return &file_eolymp_ranker_format_proto_enumTypes[0]
}

func (x Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Format.Descriptor instead.
func (Format) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_ranker_format_proto_rawDescGZIP(), []int{0}
}

var File_eolymp_ranker_format_proto protoreflect.FileDescriptor

var file_eolymp_ranker_format_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2f,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2a, 0x25, 0x0a, 0x06, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x49, 0x4f, 0x49, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x43, 0x50, 0x43,
	0x10, 0x02, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x3b, 0x72, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_ranker_format_proto_rawDescOnce sync.Once
	file_eolymp_ranker_format_proto_rawDescData = file_eolymp_ranker_format_proto_rawDesc
)

func file_eolymp_ranker_format_proto_rawDescGZIP() []byte {
	file_eolymp_ranker_format_proto_rawDescOnce.Do(func() {
		file_eolymp_ranker_format_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_ranker_format_proto_rawDescData)
	})
	return file_eolymp_ranker_format_proto_rawDescData
}

var file_eolymp_ranker_format_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_ranker_format_proto_goTypes = []any{
	(Format)(0), // 0: eolymp.ranker.Format
}
var file_eolymp_ranker_format_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_ranker_format_proto_init() }
func file_eolymp_ranker_format_proto_init() {
	if File_eolymp_ranker_format_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_ranker_format_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_ranker_format_proto_goTypes,
		DependencyIndexes: file_eolymp_ranker_format_proto_depIdxs,
		EnumInfos:         file_eolymp_ranker_format_proto_enumTypes,
	}.Build()
	File_eolymp_ranker_format_proto = out.File
	file_eolymp_ranker_format_proto_rawDesc = nil
	file_eolymp_ranker_format_proto_goTypes = nil
	file_eolymp_ranker_format_proto_depIdxs = nil
}
