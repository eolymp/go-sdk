// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: eolymp/harmony/harmony.proto

package harmony

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
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

type ListAgreementsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreferredLocales []string `protobuf:"bytes,1,rep,name=preferred_locales,json=preferredLocales,proto3" json:"preferred_locales,omitempty"`
}

func (x *ListAgreementsInput) Reset() {
	*x = ListAgreementsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_harmony_harmony_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAgreementsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAgreementsInput) ProtoMessage() {}

func (x *ListAgreementsInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_harmony_harmony_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAgreementsInput.ProtoReflect.Descriptor instead.
func (*ListAgreementsInput) Descriptor() ([]byte, []int) {
	return file_eolymp_harmony_harmony_proto_rawDescGZIP(), []int{0}
}

func (x *ListAgreementsInput) GetPreferredLocales() []string {
	if x != nil {
		return x.PreferredLocales
	}
	return nil
}

type ListAgreementsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Agreement `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListAgreementsOutput) Reset() {
	*x = ListAgreementsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_harmony_harmony_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAgreementsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAgreementsOutput) ProtoMessage() {}

func (x *ListAgreementsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_harmony_harmony_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAgreementsOutput.ProtoReflect.Descriptor instead.
func (*ListAgreementsOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_harmony_harmony_proto_rawDescGZIP(), []int{1}
}

func (x *ListAgreementsOutput) GetItems() []*Agreement {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetConsentInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgreementId string `protobuf:"bytes,1,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
}

func (x *GetConsentInput) Reset() {
	*x = GetConsentInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_harmony_harmony_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsentInput) ProtoMessage() {}

func (x *GetConsentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_harmony_harmony_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsentInput.ProtoReflect.Descriptor instead.
func (*GetConsentInput) Descriptor() ([]byte, []int) {
	return file_eolymp_harmony_harmony_proto_rawDescGZIP(), []int{2}
}

func (x *GetConsentInput) GetAgreementId() string {
	if x != nil {
		return x.AgreementId
	}
	return ""
}

type GetConsentOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consent *Consent `protobuf:"bytes,1,opt,name=consent,proto3" json:"consent,omitempty"`
}

func (x *GetConsentOutput) Reset() {
	*x = GetConsentOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_harmony_harmony_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsentOutput) ProtoMessage() {}

func (x *GetConsentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_harmony_harmony_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsentOutput.ProtoReflect.Descriptor instead.
func (*GetConsentOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_harmony_harmony_proto_rawDescGZIP(), []int{3}
}

func (x *GetConsentOutput) GetConsent() *Consent {
	if x != nil {
		return x.Consent
	}
	return nil
}

type SetConsentInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgreementId string         `protobuf:"bytes,1,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
	Status      Consent_Status `protobuf:"varint,2,opt,name=status,proto3,enum=eolymp.harmony.Consent_Status" json:"status,omitempty"`
}

func (x *SetConsentInput) Reset() {
	*x = SetConsentInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_harmony_harmony_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetConsentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConsentInput) ProtoMessage() {}

func (x *SetConsentInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_harmony_harmony_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConsentInput.ProtoReflect.Descriptor instead.
func (*SetConsentInput) Descriptor() ([]byte, []int) {
	return file_eolymp_harmony_harmony_proto_rawDescGZIP(), []int{4}
}

func (x *SetConsentInput) GetAgreementId() string {
	if x != nil {
		return x.AgreementId
	}
	return ""
}

func (x *SetConsentInput) GetStatus() Consent_Status {
	if x != nil {
		return x.Status
	}
	return Consent_PENDING
}

type SetConsentOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consent *Consent `protobuf:"bytes,1,opt,name=consent,proto3" json:"consent,omitempty"`
}

func (x *SetConsentOutput) Reset() {
	*x = SetConsentOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_harmony_harmony_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetConsentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConsentOutput) ProtoMessage() {}

func (x *SetConsentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_harmony_harmony_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConsentOutput.ProtoReflect.Descriptor instead.
func (*SetConsentOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_harmony_harmony_proto_rawDescGZIP(), []int{5}
}

func (x *SetConsentOutput) GetConsent() *Consent {
	if x != nil {
		return x.Consent
	}
	return nil
}

type FollowShortcutInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortcutId string         `protobuf:"bytes,1,opt,name=shortcut_id,json=shortcutId,proto3" json:"shortcut_id,omitempty"`
	Status     Consent_Status `protobuf:"varint,2,opt,name=status,proto3,enum=eolymp.harmony.Consent_Status" json:"status,omitempty"`
}

func (x *FollowShortcutInput) Reset() {
	*x = FollowShortcutInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_harmony_harmony_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowShortcutInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowShortcutInput) ProtoMessage() {}

func (x *FollowShortcutInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_harmony_harmony_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowShortcutInput.ProtoReflect.Descriptor instead.
func (*FollowShortcutInput) Descriptor() ([]byte, []int) {
	return file_eolymp_harmony_harmony_proto_rawDescGZIP(), []int{6}
}

func (x *FollowShortcutInput) GetShortcutId() string {
	if x != nil {
		return x.ShortcutId
	}
	return ""
}

func (x *FollowShortcutInput) GetStatus() Consent_Status {
	if x != nil {
		return x.Status
	}
	return Consent_PENDING
}

type FollowShortcutOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FollowShortcutOutput) Reset() {
	*x = FollowShortcutOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_harmony_harmony_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowShortcutOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowShortcutOutput) ProtoMessage() {}

func (x *FollowShortcutOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_harmony_harmony_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowShortcutOutput.ProtoReflect.Descriptor instead.
func (*FollowShortcutOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_harmony_harmony_proto_rawDescGZIP(), []int{7}
}

var File_eolymp_harmony_harmony_proto protoreflect.FileDescriptor

var file_eolymp_harmony_harmony_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79,
	0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x1a, 0x1d,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e,
	0x79, 0x2f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x42, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x72, 0x65, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x41, 0x67, 0x72, 0x65,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x34, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x45, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x22, 0x6c, 0x0a, 0x0f, 0x53, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x45, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x22, 0x6e,
	0x0a, 0x13, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x63, 0x75, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x16,
	0x0a, 0x14, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0xd4, 0x04, 0x0a, 0x07, 0x48, 0x61, 0x72, 0x6d, 0x6f,
	0x6e, 0x79, 0x12, 0x87, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x72, 0x65, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68,
	0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x72, 0x65, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x2a, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a,
	0x0a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e,
	0x79, 0x2f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x92, 0x01, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x20, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x41,
	0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x12, 0x2a, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2f,
	0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x67, 0x72, 0x65,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x12, 0x92, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x12, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e,
	0x79, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x20, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f,
	0x6e, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x41, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0, 0x40,
	0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x1a, 0x2a, 0x2f, 0x68, 0x61, 0x72,
	0x6d, 0x6f, 0x6e, 0x79, 0x2f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x7b, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x94, 0x01, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x12, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x37, 0xea, 0xe2, 0x0a, 0x0b, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xa0,
	0x40, 0xf8, 0xe2, 0x0a, 0x0a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x20, 0x2f, 0x68, 0x61,
	0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x73, 0x2f,
	0x7b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x3b, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_harmony_harmony_proto_rawDescOnce sync.Once
	file_eolymp_harmony_harmony_proto_rawDescData = file_eolymp_harmony_harmony_proto_rawDesc
)

func file_eolymp_harmony_harmony_proto_rawDescGZIP() []byte {
	file_eolymp_harmony_harmony_proto_rawDescOnce.Do(func() {
		file_eolymp_harmony_harmony_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_harmony_harmony_proto_rawDescData)
	})
	return file_eolymp_harmony_harmony_proto_rawDescData
}

var file_eolymp_harmony_harmony_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_eolymp_harmony_harmony_proto_goTypes = []interface{}{
	(*ListAgreementsInput)(nil),  // 0: eolymp.harmony.ListAgreementsInput
	(*ListAgreementsOutput)(nil), // 1: eolymp.harmony.ListAgreementsOutput
	(*GetConsentInput)(nil),      // 2: eolymp.harmony.GetConsentInput
	(*GetConsentOutput)(nil),     // 3: eolymp.harmony.GetConsentOutput
	(*SetConsentInput)(nil),      // 4: eolymp.harmony.SetConsentInput
	(*SetConsentOutput)(nil),     // 5: eolymp.harmony.SetConsentOutput
	(*FollowShortcutInput)(nil),  // 6: eolymp.harmony.FollowShortcutInput
	(*FollowShortcutOutput)(nil), // 7: eolymp.harmony.FollowShortcutOutput
	(*Agreement)(nil),            // 8: eolymp.harmony.Agreement
	(*Consent)(nil),              // 9: eolymp.harmony.Consent
	(Consent_Status)(0),          // 10: eolymp.harmony.Consent.Status
}
var file_eolymp_harmony_harmony_proto_depIdxs = []int32{
	8,  // 0: eolymp.harmony.ListAgreementsOutput.items:type_name -> eolymp.harmony.Agreement
	9,  // 1: eolymp.harmony.GetConsentOutput.consent:type_name -> eolymp.harmony.Consent
	10, // 2: eolymp.harmony.SetConsentInput.status:type_name -> eolymp.harmony.Consent.Status
	9,  // 3: eolymp.harmony.SetConsentOutput.consent:type_name -> eolymp.harmony.Consent
	10, // 4: eolymp.harmony.FollowShortcutInput.status:type_name -> eolymp.harmony.Consent.Status
	0,  // 5: eolymp.harmony.Harmony.ListAgreements:input_type -> eolymp.harmony.ListAgreementsInput
	2,  // 6: eolymp.harmony.Harmony.GetConsent:input_type -> eolymp.harmony.GetConsentInput
	4,  // 7: eolymp.harmony.Harmony.SetConsent:input_type -> eolymp.harmony.SetConsentInput
	6,  // 8: eolymp.harmony.Harmony.FollowShortcut:input_type -> eolymp.harmony.FollowShortcutInput
	1,  // 9: eolymp.harmony.Harmony.ListAgreements:output_type -> eolymp.harmony.ListAgreementsOutput
	3,  // 10: eolymp.harmony.Harmony.GetConsent:output_type -> eolymp.harmony.GetConsentOutput
	5,  // 11: eolymp.harmony.Harmony.SetConsent:output_type -> eolymp.harmony.SetConsentOutput
	7,  // 12: eolymp.harmony.Harmony.FollowShortcut:output_type -> eolymp.harmony.FollowShortcutOutput
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_harmony_harmony_proto_init() }
func file_eolymp_harmony_harmony_proto_init() {
	if File_eolymp_harmony_harmony_proto != nil {
		return
	}
	file_eolymp_harmony_agreement_proto_init()
	file_eolymp_harmony_consent_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_harmony_harmony_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAgreementsInput); i {
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
		file_eolymp_harmony_harmony_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAgreementsOutput); i {
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
		file_eolymp_harmony_harmony_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsentInput); i {
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
		file_eolymp_harmony_harmony_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsentOutput); i {
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
		file_eolymp_harmony_harmony_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetConsentInput); i {
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
		file_eolymp_harmony_harmony_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetConsentOutput); i {
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
		file_eolymp_harmony_harmony_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowShortcutInput); i {
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
		file_eolymp_harmony_harmony_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowShortcutOutput); i {
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
			RawDescriptor: file_eolymp_harmony_harmony_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_harmony_harmony_proto_goTypes,
		DependencyIndexes: file_eolymp_harmony_harmony_proto_depIdxs,
		MessageInfos:      file_eolymp_harmony_harmony_proto_msgTypes,
	}.Build()
	File_eolymp_harmony_harmony_proto = out.File
	file_eolymp_harmony_harmony_proto_rawDesc = nil
	file_eolymp_harmony_harmony_proto_goTypes = nil
	file_eolymp_harmony_harmony_proto_depIdxs = nil
}
