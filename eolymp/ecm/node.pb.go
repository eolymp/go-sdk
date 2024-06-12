// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: eolymp/ecm/node.proto

package ecm

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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Attr     map[string]string `protobuf:"bytes,2,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Children []*Node           `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_ecm_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_ecm_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_eolymp_ecm_node_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Node) GetAttr() map[string]string {
	if x != nil {
		return x.Attr
	}
	return nil
}

func (x *Node) GetChildren() []*Node {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_eolymp_ecm_node_proto protoreflect.FileDescriptor

var file_eolymp_ecm_node_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x65, 0x63, 0x6d, 0x22, 0xb1, 0x01, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61, 0x74, 0x74, 0x72,
	0x12, 0x2c, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x65, 0x63, 0x6d, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x1a, 0x37,
	0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x65, 0x63, 0x6d, 0x3b, 0x65,
	0x63, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_ecm_node_proto_rawDescOnce sync.Once
	file_eolymp_ecm_node_proto_rawDescData = file_eolymp_ecm_node_proto_rawDesc
)

func file_eolymp_ecm_node_proto_rawDescGZIP() []byte {
	file_eolymp_ecm_node_proto_rawDescOnce.Do(func() {
		file_eolymp_ecm_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_ecm_node_proto_rawDescData)
	})
	return file_eolymp_ecm_node_proto_rawDescData
}

var file_eolymp_ecm_node_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eolymp_ecm_node_proto_goTypes = []any{
	(*Node)(nil), // 0: eolymp.ecm.Node
	nil,          // 1: eolymp.ecm.Node.AttrEntry
}
var file_eolymp_ecm_node_proto_depIdxs = []int32{
	1, // 0: eolymp.ecm.Node.attr:type_name -> eolymp.ecm.Node.AttrEntry
	0, // 1: eolymp.ecm.Node.children:type_name -> eolymp.ecm.Node
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eolymp_ecm_node_proto_init() }
func file_eolymp_ecm_node_proto_init() {
	if File_eolymp_ecm_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_ecm_node_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Node); i {
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
			RawDescriptor: file_eolymp_ecm_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_ecm_node_proto_goTypes,
		DependencyIndexes: file_eolymp_ecm_node_proto_depIdxs,
		MessageInfos:      file_eolymp_ecm_node_proto_msgTypes,
	}.Build()
	File_eolymp_ecm_node_proto = out.File
	file_eolymp_ecm_node_proto_rawDesc = nil
	file_eolymp_ecm_node_proto_goTypes = nil
	file_eolymp_ecm_node_proto_depIdxs = nil
}
