// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/auth/security_event.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type SecurityEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Event:
	//
	//	*SecurityEvent_SessionClosed_
	//	*SecurityEvent_TokenRevoked_
	//	*SecurityEvent_AccountPurged_
	Event         isSecurityEvent_Event `protobuf_oneof:"event"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecurityEvent) Reset() {
	*x = SecurityEvent{}
	mi := &file_eolymp_auth_security_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecurityEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityEvent) ProtoMessage() {}

func (x *SecurityEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_security_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityEvent.ProtoReflect.Descriptor instead.
func (*SecurityEvent) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_security_event_proto_rawDescGZIP(), []int{0}
}

func (x *SecurityEvent) GetEvent() isSecurityEvent_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *SecurityEvent) GetSessionClosed() *SecurityEvent_SessionClosed {
	if x != nil {
		if x, ok := x.Event.(*SecurityEvent_SessionClosed_); ok {
			return x.SessionClosed
		}
	}
	return nil
}

func (x *SecurityEvent) GetTokenRevoked() *SecurityEvent_TokenRevoked {
	if x != nil {
		if x, ok := x.Event.(*SecurityEvent_TokenRevoked_); ok {
			return x.TokenRevoked
		}
	}
	return nil
}

func (x *SecurityEvent) GetAccountPurged() *SecurityEvent_AccountPurged {
	if x != nil {
		if x, ok := x.Event.(*SecurityEvent_AccountPurged_); ok {
			return x.AccountPurged
		}
	}
	return nil
}

type isSecurityEvent_Event interface {
	isSecurityEvent_Event()
}

type SecurityEvent_SessionClosed_ struct {
	SessionClosed *SecurityEvent_SessionClosed `protobuf:"bytes,1,opt,name=session_closed,json=sessionClosed,proto3,oneof"`
}

type SecurityEvent_TokenRevoked_ struct {
	TokenRevoked *SecurityEvent_TokenRevoked `protobuf:"bytes,2,opt,name=token_revoked,json=tokenRevoked,proto3,oneof"`
}

type SecurityEvent_AccountPurged_ struct {
	AccountPurged *SecurityEvent_AccountPurged `protobuf:"bytes,3,opt,name=account_purged,json=accountPurged,proto3,oneof"`
}

func (*SecurityEvent_SessionClosed_) isSecurityEvent_Event() {}

func (*SecurityEvent_TokenRevoked_) isSecurityEvent_Event() {}

func (*SecurityEvent_AccountPurged_) isSecurityEvent_Event() {}

type SecurityEvent_SessionClosed struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Issuer        string                 `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`                        // iss
	Subject       string                 `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`                      // sub
	SessionId     string                 `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"` // sid
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecurityEvent_SessionClosed) Reset() {
	*x = SecurityEvent_SessionClosed{}
	mi := &file_eolymp_auth_security_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecurityEvent_SessionClosed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityEvent_SessionClosed) ProtoMessage() {}

func (x *SecurityEvent_SessionClosed) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_security_event_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityEvent_SessionClosed.ProtoReflect.Descriptor instead.
func (*SecurityEvent_SessionClosed) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_security_event_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SecurityEvent_SessionClosed) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *SecurityEvent_SessionClosed) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SecurityEvent_SessionClosed) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type SecurityEvent_TokenRevoked struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Issuer        string                 `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject       string                 `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	TokenType     string                 `protobuf:"bytes,3,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"` // access_token or refresh_token
	TokenHashMd5  string                 `protobuf:"bytes,4,opt,name=token_hash_md5,json=tokenHashMd5,proto3" json:"token_hash_md5,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecurityEvent_TokenRevoked) Reset() {
	*x = SecurityEvent_TokenRevoked{}
	mi := &file_eolymp_auth_security_event_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecurityEvent_TokenRevoked) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityEvent_TokenRevoked) ProtoMessage() {}

func (x *SecurityEvent_TokenRevoked) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_security_event_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityEvent_TokenRevoked.ProtoReflect.Descriptor instead.
func (*SecurityEvent_TokenRevoked) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_security_event_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SecurityEvent_TokenRevoked) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *SecurityEvent_TokenRevoked) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SecurityEvent_TokenRevoked) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *SecurityEvent_TokenRevoked) GetTokenHashMd5() string {
	if x != nil {
		return x.TokenHashMd5
	}
	return ""
}

type SecurityEvent_AccountPurged struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Issuer        string                 `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject       string                 `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecurityEvent_AccountPurged) Reset() {
	*x = SecurityEvent_AccountPurged{}
	mi := &file_eolymp_auth_security_event_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecurityEvent_AccountPurged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityEvent_AccountPurged) ProtoMessage() {}

func (x *SecurityEvent_AccountPurged) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_security_event_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityEvent_AccountPurged.ProtoReflect.Descriptor instead.
func (*SecurityEvent_AccountPurged) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_security_event_proto_rawDescGZIP(), []int{0, 2}
}

func (x *SecurityEvent_AccountPurged) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *SecurityEvent_AccountPurged) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

var File_eolymp_auth_security_event_proto protoreflect.FileDescriptor

var file_eolymp_auth_security_event_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x22,
	0xbb, 0x04, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x51, 0x0a, 0x0e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x64, 0x12, 0x4e, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x64, 0x12, 0x51, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x70, 0x75, 0x72, 0x67, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x50, 0x75, 0x72, 0x67, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x50, 0x75, 0x72, 0x67, 0x65, 0x64, 0x1a, 0x60, 0x0a, 0x0d, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x85, 0x01, 0x0a, 0x0c, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x4d, 0x64,
	0x35, 0x1a, 0x41, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x75, 0x72, 0x67,
	0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_eolymp_auth_security_event_proto_rawDescOnce sync.Once
	file_eolymp_auth_security_event_proto_rawDescData []byte
)

func file_eolymp_auth_security_event_proto_rawDescGZIP() []byte {
	file_eolymp_auth_security_event_proto_rawDescOnce.Do(func() {
		file_eolymp_auth_security_event_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_auth_security_event_proto_rawDesc), len(file_eolymp_auth_security_event_proto_rawDesc)))
	})
	return file_eolymp_auth_security_event_proto_rawDescData
}

var file_eolymp_auth_security_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eolymp_auth_security_event_proto_goTypes = []any{
	(*SecurityEvent)(nil),               // 0: eolymp.auth.SecurityEvent
	(*SecurityEvent_SessionClosed)(nil), // 1: eolymp.auth.SecurityEvent.SessionClosed
	(*SecurityEvent_TokenRevoked)(nil),  // 2: eolymp.auth.SecurityEvent.TokenRevoked
	(*SecurityEvent_AccountPurged)(nil), // 3: eolymp.auth.SecurityEvent.AccountPurged
}
var file_eolymp_auth_security_event_proto_depIdxs = []int32{
	1, // 0: eolymp.auth.SecurityEvent.session_closed:type_name -> eolymp.auth.SecurityEvent.SessionClosed
	2, // 1: eolymp.auth.SecurityEvent.token_revoked:type_name -> eolymp.auth.SecurityEvent.TokenRevoked
	3, // 2: eolymp.auth.SecurityEvent.account_purged:type_name -> eolymp.auth.SecurityEvent.AccountPurged
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eolymp_auth_security_event_proto_init() }
func file_eolymp_auth_security_event_proto_init() {
	if File_eolymp_auth_security_event_proto != nil {
		return
	}
	file_eolymp_auth_security_event_proto_msgTypes[0].OneofWrappers = []any{
		(*SecurityEvent_SessionClosed_)(nil),
		(*SecurityEvent_TokenRevoked_)(nil),
		(*SecurityEvent_AccountPurged_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_auth_security_event_proto_rawDesc), len(file_eolymp_auth_security_event_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_auth_security_event_proto_goTypes,
		DependencyIndexes: file_eolymp_auth_security_event_proto_depIdxs,
		MessageInfos:      file_eolymp_auth_security_event_proto_msgTypes,
	}.Build()
	File_eolymp_auth_security_event_proto = out.File
	file_eolymp_auth_security_event_proto_goTypes = nil
	file_eolymp_auth_security_event_proto_depIdxs = nil
}
