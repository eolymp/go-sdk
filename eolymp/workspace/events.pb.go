// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/workspace/events.proto

package workspace

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

type ProjectCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ProjectCreatedEvent) Reset() {
	*x = ProjectCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_workspace_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectCreatedEvent) ProtoMessage() {}

func (x *ProjectCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_workspace_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectCreatedEvent.ProtoReflect.Descriptor instead.
func (*ProjectCreatedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_workspace_events_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectCreatedEvent) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type ProjectUpdatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ProjectUpdatedEvent) Reset() {
	*x = ProjectUpdatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_workspace_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectUpdatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectUpdatedEvent) ProtoMessage() {}

func (x *ProjectUpdatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_workspace_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectUpdatedEvent.ProtoReflect.Descriptor instead.
func (*ProjectUpdatedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_workspace_events_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectUpdatedEvent) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type ProjectDeletedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ProjectDeletedEvent) Reset() {
	*x = ProjectDeletedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_workspace_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectDeletedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDeletedEvent) ProtoMessage() {}

func (x *ProjectDeletedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_workspace_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectDeletedEvent.ProtoReflect.Descriptor instead.
func (*ProjectDeletedEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_workspace_events_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectDeletedEvent) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

var File_eolymp_workspace_events_proto protoreflect.FileDescriptor

var file_eolymp_workspace_events_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4a, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x4a, 0x0a,
	0x13, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x4a, 0x0a, 0x13, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x33, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_workspace_events_proto_rawDescOnce sync.Once
	file_eolymp_workspace_events_proto_rawDescData = file_eolymp_workspace_events_proto_rawDesc
)

func file_eolymp_workspace_events_proto_rawDescGZIP() []byte {
	file_eolymp_workspace_events_proto_rawDescOnce.Do(func() {
		file_eolymp_workspace_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_workspace_events_proto_rawDescData)
	})
	return file_eolymp_workspace_events_proto_rawDescData
}

var file_eolymp_workspace_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eolymp_workspace_events_proto_goTypes = []interface{}{
	(*ProjectCreatedEvent)(nil), // 0: eolymp.workspace.ProjectCreatedEvent
	(*ProjectUpdatedEvent)(nil), // 1: eolymp.workspace.ProjectUpdatedEvent
	(*ProjectDeletedEvent)(nil), // 2: eolymp.workspace.ProjectDeletedEvent
	(*Project)(nil),             // 3: eolymp.workspace.Project
}
var file_eolymp_workspace_events_proto_depIdxs = []int32{
	3, // 0: eolymp.workspace.ProjectCreatedEvent.project:type_name -> eolymp.workspace.Project
	3, // 1: eolymp.workspace.ProjectUpdatedEvent.project:type_name -> eolymp.workspace.Project
	3, // 2: eolymp.workspace.ProjectDeletedEvent.project:type_name -> eolymp.workspace.Project
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_workspace_events_proto_init() }
func file_eolymp_workspace_events_proto_init() {
	if File_eolymp_workspace_events_proto != nil {
		return
	}
	file_eolymp_workspace_project_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_workspace_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectCreatedEvent); i {
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
		file_eolymp_workspace_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectUpdatedEvent); i {
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
		file_eolymp_workspace_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectDeletedEvent); i {
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
			RawDescriptor: file_eolymp_workspace_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_workspace_events_proto_goTypes,
		DependencyIndexes: file_eolymp_workspace_events_proto_depIdxs,
		MessageInfos:      file_eolymp_workspace_events_proto_msgTypes,
	}.Build()
	File_eolymp_workspace_events_proto = out.File
	file_eolymp_workspace_events_proto_rawDesc = nil
	file_eolymp_workspace_events_proto_goTypes = nil
	file_eolymp_workspace_events_proto_depIdxs = nil
}
