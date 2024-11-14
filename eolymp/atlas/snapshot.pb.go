// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: eolymp/atlas/snapshot.proto

package atlas

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

type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Problem     *Problem       `protobuf:"bytes,1,opt,name=problem,proto3" json:"problem,omitempty"`
	Testing     *TestingConfig `protobuf:"bytes,10,opt,name=testing,proto3" json:"testing,omitempty"`
	Checker     *Checker       `protobuf:"bytes,2,opt,name=checker,proto3" json:"checker,omitempty"`
	Interactor  *Interactor    `protobuf:"bytes,3,opt,name=interactor,proto3" json:"interactor,omitempty"`
	Statements  []*Statement   `protobuf:"bytes,4,rep,name=statements,proto3" json:"statements,omitempty"`
	Templates   []*Template    `protobuf:"bytes,5,rep,name=templates,proto3" json:"templates,omitempty"`
	Attachments []*Attachment  `protobuf:"bytes,6,rep,name=attachments,proto3" json:"attachments,omitempty"`
	Testsets    []*Testset     `protobuf:"bytes,7,rep,name=testsets,proto3" json:"testsets,omitempty"`
	Tests       []*Test        `protobuf:"bytes,8,rep,name=tests,proto3" json:"tests,omitempty"`
	Editorials  []*Editorial   `protobuf:"bytes,9,rep,name=editorials,proto3" json:"editorials,omitempty"`
	Solutions   []*Solution    `protobuf:"bytes,11,rep,name=solutions,proto3" json:"solutions,omitempty"`
	Scripts     []*Script      `protobuf:"bytes,12,rep,name=scripts,proto3" json:"scripts,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	mi := &file_eolymp_atlas_snapshot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_atlas_snapshot_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_eolymp_atlas_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *Snapshot) GetProblem() *Problem {
	if x != nil {
		return x.Problem
	}
	return nil
}

func (x *Snapshot) GetTesting() *TestingConfig {
	if x != nil {
		return x.Testing
	}
	return nil
}

func (x *Snapshot) GetChecker() *Checker {
	if x != nil {
		return x.Checker
	}
	return nil
}

func (x *Snapshot) GetInteractor() *Interactor {
	if x != nil {
		return x.Interactor
	}
	return nil
}

func (x *Snapshot) GetStatements() []*Statement {
	if x != nil {
		return x.Statements
	}
	return nil
}

func (x *Snapshot) GetTemplates() []*Template {
	if x != nil {
		return x.Templates
	}
	return nil
}

func (x *Snapshot) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *Snapshot) GetTestsets() []*Testset {
	if x != nil {
		return x.Testsets
	}
	return nil
}

func (x *Snapshot) GetTests() []*Test {
	if x != nil {
		return x.Tests
	}
	return nil
}

func (x *Snapshot) GetEditorials() []*Editorial {
	if x != nil {
		return x.Editorials
	}
	return nil
}

func (x *Snapshot) GetSolutions() []*Solution {
	if x != nil {
		return x.Solutions
	}
	return nil
}

func (x *Snapshot) GetScripts() []*Script {
	if x != nil {
		return x.Scripts
	}
	return nil
}

var File_eolymp_atlas_snapshot_proto protoreflect.FileDescriptor

var file_eolymp_atlas_snapshot_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x1a, 0x1d, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x84, 0x05, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x2f, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x35, 0x0a,
	0x07, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x07, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x37, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x3a,
	0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x65, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x73, 0x65, 0x74, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x73, 0x12, 0x28, 0x0a,
	0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x52, 0x0a, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x73,
	0x12, 0x34, 0x0a, 0x09, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x07, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x3b,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_atlas_snapshot_proto_rawDescOnce sync.Once
	file_eolymp_atlas_snapshot_proto_rawDescData = file_eolymp_atlas_snapshot_proto_rawDesc
)

func file_eolymp_atlas_snapshot_proto_rawDescGZIP() []byte {
	file_eolymp_atlas_snapshot_proto_rawDescOnce.Do(func() {
		file_eolymp_atlas_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_atlas_snapshot_proto_rawDescData)
	})
	return file_eolymp_atlas_snapshot_proto_rawDescData
}

var file_eolymp_atlas_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_atlas_snapshot_proto_goTypes = []any{
	(*Snapshot)(nil),      // 0: eolymp.atlas.Snapshot
	(*Problem)(nil),       // 1: eolymp.atlas.Problem
	(*TestingConfig)(nil), // 2: eolymp.atlas.TestingConfig
	(*Checker)(nil),       // 3: eolymp.atlas.Checker
	(*Interactor)(nil),    // 4: eolymp.atlas.Interactor
	(*Statement)(nil),     // 5: eolymp.atlas.Statement
	(*Template)(nil),      // 6: eolymp.atlas.Template
	(*Attachment)(nil),    // 7: eolymp.atlas.Attachment
	(*Testset)(nil),       // 8: eolymp.atlas.Testset
	(*Test)(nil),          // 9: eolymp.atlas.Test
	(*Editorial)(nil),     // 10: eolymp.atlas.Editorial
	(*Solution)(nil),      // 11: eolymp.atlas.Solution
	(*Script)(nil),        // 12: eolymp.atlas.Script
}
var file_eolymp_atlas_snapshot_proto_depIdxs = []int32{
	1,  // 0: eolymp.atlas.Snapshot.problem:type_name -> eolymp.atlas.Problem
	2,  // 1: eolymp.atlas.Snapshot.testing:type_name -> eolymp.atlas.TestingConfig
	3,  // 2: eolymp.atlas.Snapshot.checker:type_name -> eolymp.atlas.Checker
	4,  // 3: eolymp.atlas.Snapshot.interactor:type_name -> eolymp.atlas.Interactor
	5,  // 4: eolymp.atlas.Snapshot.statements:type_name -> eolymp.atlas.Statement
	6,  // 5: eolymp.atlas.Snapshot.templates:type_name -> eolymp.atlas.Template
	7,  // 6: eolymp.atlas.Snapshot.attachments:type_name -> eolymp.atlas.Attachment
	8,  // 7: eolymp.atlas.Snapshot.testsets:type_name -> eolymp.atlas.Testset
	9,  // 8: eolymp.atlas.Snapshot.tests:type_name -> eolymp.atlas.Test
	10, // 9: eolymp.atlas.Snapshot.editorials:type_name -> eolymp.atlas.Editorial
	11, // 10: eolymp.atlas.Snapshot.solutions:type_name -> eolymp.atlas.Solution
	12, // 11: eolymp.atlas.Snapshot.scripts:type_name -> eolymp.atlas.Script
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_eolymp_atlas_snapshot_proto_init() }
func file_eolymp_atlas_snapshot_proto_init() {
	if File_eolymp_atlas_snapshot_proto != nil {
		return
	}
	file_eolymp_atlas_attachment_proto_init()
	file_eolymp_atlas_code_template_proto_init()
	file_eolymp_atlas_editorial_proto_init()
	file_eolymp_atlas_problem_proto_init()
	file_eolymp_atlas_script_proto_init()
	file_eolymp_atlas_solution_proto_init()
	file_eolymp_atlas_statement_proto_init()
	file_eolymp_atlas_testing_checker_proto_init()
	file_eolymp_atlas_testing_config_proto_init()
	file_eolymp_atlas_testing_interactor_proto_init()
	file_eolymp_atlas_testing_test_proto_init()
	file_eolymp_atlas_testing_testset_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_atlas_snapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_atlas_snapshot_proto_goTypes,
		DependencyIndexes: file_eolymp_atlas_snapshot_proto_depIdxs,
		MessageInfos:      file_eolymp_atlas_snapshot_proto_msgTypes,
	}.Build()
	File_eolymp_atlas_snapshot_proto = out.File
	file_eolymp_atlas_snapshot_proto_rawDesc = nil
	file_eolymp_atlas_snapshot_proto_goTypes = nil
	file_eolymp_atlas_snapshot_proto_depIdxs = nil
}
