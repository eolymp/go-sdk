// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.2
// source: eolymp/workspace/project.proto

package workspace

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Project_Visibility int32

const (
	Project_NONE    Project_Visibility = 0 // Reserved, don't use
	Project_PRIVATE Project_Visibility = 1 // Private, not visible by others
	Project_PUBLIC  Project_Visibility = 2 // Public, visible to others
)

// Enum value maps for Project_Visibility.
var (
	Project_Visibility_name = map[int32]string{
		0: "NONE",
		1: "PRIVATE",
		2: "PUBLIC",
	}
	Project_Visibility_value = map[string]int32{
		"NONE":    0,
		"PRIVATE": 1,
		"PUBLIC":  2,
	}
)

func (x Project_Visibility) Enum() *Project_Visibility {
	p := new(Project_Visibility)
	*p = x
	return p
}

func (x Project_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Project_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_workspace_project_proto_enumTypes[0].Descriptor()
}

func (Project_Visibility) Type() protoreflect.EnumType {
	return &file_eolymp_workspace_project_proto_enumTypes[0]
}

func (x Project_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Project_Visibility.Descriptor instead.
func (Project_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_workspace_project_proto_rawDescGZIP(), []int{0, 0}
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri        string               `protobuf:"bytes,9999,opt,name=uri,proto3" json:"uri,omitempty"`
	Id         string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Runtime    string               `protobuf:"bytes,3,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Visibility Project_Visibility   `protobuf:"varint,4,opt,name=visibility,proto3,enum=eolymp.workspace.Project_Visibility" json:"visibility,omitempty"`
	AuthorId   string               `protobuf:"bytes,5,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	CreatedOn  *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	UpdatedOn  *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_on,json=updatedOn,proto3" json:"updated_on,omitempty"`
	Target     *Project_Target      `protobuf:"bytes,10,opt,name=target,proto3" json:"target,omitempty"`
	Labels     []string             `protobuf:"bytes,100,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_workspace_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_workspace_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_eolymp_workspace_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Project) GetVisibility() Project_Visibility {
	if x != nil {
		return x.Visibility
	}
	return Project_NONE
}

func (x *Project) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Project) GetCreatedOn() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedOn
	}
	return nil
}

func (x *Project) GetUpdatedOn() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedOn
	}
	return nil
}

func (x *Project) GetTarget() *Project_Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *Project) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Project_Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId   string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	ProblemId string `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
}

func (x *Project_Target) Reset() {
	*x = Project_Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_workspace_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project_Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project_Target) ProtoMessage() {}

func (x *Project_Target) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_workspace_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project_Target.ProtoReflect.Descriptor instead.
func (*Project_Target) Descriptor() ([]byte, []int) {
	return file_eolymp_workspace_project_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Project_Target) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *Project_Target) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

var File_eolymp_workspace_project_proto protoreflect.FileDescriptor

var file_eolymp_workspace_project_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x11, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x8f, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x44, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x64, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x42, 0x0a, 0x06, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x22,
	0x2f, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41,
	0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x02,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3b, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_workspace_project_proto_rawDescOnce sync.Once
	file_eolymp_workspace_project_proto_rawDescData = file_eolymp_workspace_project_proto_rawDesc
)

func file_eolymp_workspace_project_proto_rawDescGZIP() []byte {
	file_eolymp_workspace_project_proto_rawDescOnce.Do(func() {
		file_eolymp_workspace_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_workspace_project_proto_rawDescData)
	})
	return file_eolymp_workspace_project_proto_rawDescData
}

var file_eolymp_workspace_project_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_workspace_project_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_workspace_project_proto_goTypes = []interface{}{
	(Project_Visibility)(0),     // 0: eolymp.workspace.Project.Visibility
	(*Project)(nil),             // 1: eolymp.workspace.Project
	(*Project_Target)(nil),      // 2: eolymp.workspace.Project.Target
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_eolymp_workspace_project_proto_depIdxs = []int32{
	0, // 0: eolymp.workspace.Project.visibility:type_name -> eolymp.workspace.Project.Visibility
	3, // 1: eolymp.workspace.Project.created_on:type_name -> google.protobuf.Timestamp
	3, // 2: eolymp.workspace.Project.updated_on:type_name -> google.protobuf.Timestamp
	2, // 3: eolymp.workspace.Project.target:type_name -> eolymp.workspace.Project.Target
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_workspace_project_proto_init() }
func file_eolymp_workspace_project_proto_init() {
	if File_eolymp_workspace_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_workspace_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_eolymp_workspace_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project_Target); i {
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
			RawDescriptor: file_eolymp_workspace_project_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_workspace_project_proto_goTypes,
		DependencyIndexes: file_eolymp_workspace_project_proto_depIdxs,
		EnumInfos:         file_eolymp_workspace_project_proto_enumTypes,
		MessageInfos:      file_eolymp_workspace_project_proto_msgTypes,
	}.Build()
	File_eolymp_workspace_project_proto = out.File
	file_eolymp_workspace_project_proto_rawDesc = nil
	file_eolymp_workspace_project_proto_goTypes = nil
	file_eolymp_workspace_project_proto_depIdxs = nil
}
