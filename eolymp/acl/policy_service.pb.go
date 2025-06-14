// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: eolymp/acl/policy_service.proto

package acl

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	wellknown "github.com/eolymp/go-sdk/eolymp/wellknown"
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

type UpdatePolicyInput_Patch int32

const (
	UpdatePolicyInput_ALL    UpdatePolicyInput_Patch = 0
	UpdatePolicyInput_NAME   UpdatePolicyInput_Patch = 1
	UpdatePolicyInput_ALLOWS UpdatePolicyInput_Patch = 2
)

// Enum value maps for UpdatePolicyInput_Patch.
var (
	UpdatePolicyInput_Patch_name = map[int32]string{
		0: "ALL",
		1: "NAME",
		2: "ALLOWS",
	}
	UpdatePolicyInput_Patch_value = map[string]int32{
		"ALL":    0,
		"NAME":   1,
		"ALLOWS": 2,
	}
)

func (x UpdatePolicyInput_Patch) Enum() *UpdatePolicyInput_Patch {
	p := new(UpdatePolicyInput_Patch)
	*p = x
	return p
}

func (x UpdatePolicyInput_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdatePolicyInput_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_acl_policy_service_proto_enumTypes[0].Descriptor()
}

func (UpdatePolicyInput_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_acl_policy_service_proto_enumTypes[0]
}

func (x UpdatePolicyInput_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdatePolicyInput_Patch.Descriptor instead.
func (UpdatePolicyInput_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{2, 0}
}

type CreatePolicyInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Policy        *Policy                `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePolicyInput) Reset() {
	*x = CreatePolicyInput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePolicyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyInput) ProtoMessage() {}

func (x *CreatePolicyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyInput.ProtoReflect.Descriptor instead.
func (*CreatePolicyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePolicyInput) GetPolicy() *Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type CreatePolicyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PolicyId      string                 `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePolicyOutput) Reset() {
	*x = CreatePolicyOutput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePolicyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyOutput) ProtoMessage() {}

func (x *CreatePolicyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyOutput.ProtoReflect.Descriptor instead.
func (*CreatePolicyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePolicyOutput) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

type UpdatePolicyInput struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Patch         []UpdatePolicyInput_Patch `protobuf:"varint,10,rep,packed,name=patch,proto3,enum=eolymp.acl.UpdatePolicyInput_Patch" json:"patch,omitempty"`
	PolicyId      string                    `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	Policy        *Policy                   `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePolicyInput) Reset() {
	*x = UpdatePolicyInput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePolicyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePolicyInput) ProtoMessage() {}

func (x *UpdatePolicyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePolicyInput.ProtoReflect.Descriptor instead.
func (*UpdatePolicyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePolicyInput) GetPatch() []UpdatePolicyInput_Patch {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *UpdatePolicyInput) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *UpdatePolicyInput) GetPolicy() *Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type UpdatePolicyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePolicyOutput) Reset() {
	*x = UpdatePolicyOutput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePolicyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePolicyOutput) ProtoMessage() {}

func (x *UpdatePolicyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePolicyOutput.ProtoReflect.Descriptor instead.
func (*UpdatePolicyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{3}
}

type DeletePolicyInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PolicyId      string                 `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePolicyInput) Reset() {
	*x = DeletePolicyInput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePolicyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyInput) ProtoMessage() {}

func (x *DeletePolicyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyInput.ProtoReflect.Descriptor instead.
func (*DeletePolicyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePolicyInput) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

type DeletePolicyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePolicyOutput) Reset() {
	*x = DeletePolicyOutput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePolicyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyOutput) ProtoMessage() {}

func (x *DeletePolicyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyOutput.ProtoReflect.Descriptor instead.
func (*DeletePolicyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{5}
}

type DescribePolicyInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PolicyId      string                 `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribePolicyInput) Reset() {
	*x = DescribePolicyInput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribePolicyInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePolicyInput) ProtoMessage() {}

func (x *DescribePolicyInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePolicyInput.ProtoReflect.Descriptor instead.
func (*DescribePolicyInput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{6}
}

func (x *DescribePolicyInput) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

type DescribePolicyOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Policy        *Policy                `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribePolicyOutput) Reset() {
	*x = DescribePolicyOutput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribePolicyOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePolicyOutput) ProtoMessage() {}

func (x *DescribePolicyOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePolicyOutput.ProtoReflect.Descriptor instead.
func (*DescribePolicyOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{7}
}

func (x *DescribePolicyOutput) GetPolicy() *Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type ListPoliciesInput struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Offset        int32                     `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Size          int32                     `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Filters       *ListPoliciesInput_Filter `protobuf:"bytes,40,opt,name=filters,proto3" json:"filters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPoliciesInput) Reset() {
	*x = ListPoliciesInput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPoliciesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesInput) ProtoMessage() {}

func (x *ListPoliciesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesInput.ProtoReflect.Descriptor instead.
func (*ListPoliciesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListPoliciesInput) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListPoliciesInput) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListPoliciesInput) GetFilters() *ListPoliciesInput_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListPoliciesOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*Policy              `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPoliciesOutput) Reset() {
	*x = ListPoliciesOutput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPoliciesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesOutput) ProtoMessage() {}

func (x *ListPoliciesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesOutput.ProtoReflect.Descriptor instead.
func (*ListPoliciesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListPoliciesOutput) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListPoliciesOutput) GetItems() []*Policy {
	if x != nil {
		return x.Items
	}
	return nil
}

type CopyPoliciesInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SrcPrincipal  string                 `protobuf:"bytes,1,opt,name=src_principal,json=srcPrincipal,proto3" json:"src_principal,omitempty"` // copy policies of given principal
	DstPrincipal  string                 `protobuf:"bytes,2,opt,name=dst_principal,json=dstPrincipal,proto3" json:"dst_principal,omitempty"` // new principal
	SrcResource   string                 `protobuf:"bytes,3,opt,name=src_resource,json=srcResource,proto3" json:"src_resource,omitempty"`    // copy policies of given resource
	DstResource   string                 `protobuf:"bytes,4,opt,name=dst_resource,json=dstResource,proto3" json:"dst_resource,omitempty"`    // new resource
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CopyPoliciesInput) Reset() {
	*x = CopyPoliciesInput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CopyPoliciesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyPoliciesInput) ProtoMessage() {}

func (x *CopyPoliciesInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyPoliciesInput.ProtoReflect.Descriptor instead.
func (*CopyPoliciesInput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{10}
}

func (x *CopyPoliciesInput) GetSrcPrincipal() string {
	if x != nil {
		return x.SrcPrincipal
	}
	return ""
}

func (x *CopyPoliciesInput) GetDstPrincipal() string {
	if x != nil {
		return x.DstPrincipal
	}
	return ""
}

func (x *CopyPoliciesInput) GetSrcResource() string {
	if x != nil {
		return x.SrcResource
	}
	return ""
}

func (x *CopyPoliciesInput) GetDstResource() string {
	if x != nil {
		return x.DstResource
	}
	return ""
}

type CopyPoliciesOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CopiesCount   int32                  `protobuf:"varint,1,opt,name=copies_count,json=copiesCount,proto3" json:"copies_count,omitempty"` // number of policies copied
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CopyPoliciesOutput) Reset() {
	*x = CopyPoliciesOutput{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CopyPoliciesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyPoliciesOutput) ProtoMessage() {}

func (x *CopyPoliciesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyPoliciesOutput.ProtoReflect.Descriptor instead.
func (*CopyPoliciesOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{11}
}

func (x *CopyPoliciesOutput) GetCopiesCount() int32 {
	if x != nil {
		return x.CopiesCount
	}
	return 0
}

type ListPoliciesInput_Filter struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Query         string                        `protobuf:"bytes,10,opt,name=query,proto3" json:"query,omitempty"`
	Id            []*wellknown.ExpressionID     `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Principal     []*wellknown.ExpressionID     `protobuf:"bytes,2,rep,name=principal,proto3" json:"principal,omitempty"`
	Name          []*wellknown.ExpressionString `protobuf:"bytes,3,rep,name=name,proto3" json:"name,omitempty"`
	Resource      []*wellknown.ExpressionEnum   `protobuf:"bytes,4,rep,name=resource,proto3" json:"resource,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPoliciesInput_Filter) Reset() {
	*x = ListPoliciesInput_Filter{}
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPoliciesInput_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesInput_Filter) ProtoMessage() {}

func (x *ListPoliciesInput_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_acl_policy_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesInput_Filter.ProtoReflect.Descriptor instead.
func (*ListPoliciesInput_Filter) Descriptor() ([]byte, []int) {
	return file_eolymp_acl_policy_service_proto_rawDescGZIP(), []int{8, 0}
}

func (x *ListPoliciesInput_Filter) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListPoliciesInput_Filter) GetId() []*wellknown.ExpressionID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListPoliciesInput_Filter) GetPrincipal() []*wellknown.ExpressionID {
	if x != nil {
		return x.Principal
	}
	return nil
}

func (x *ListPoliciesInput_Filter) GetName() []*wellknown.ExpressionString {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ListPoliciesInput_Filter) GetResource() []*wellknown.ExpressionEnum {
	if x != nil {
		return x.Resource
	}
	return nil
}

var File_eolymp_acl_policy_service_proto protoreflect.FileDescriptor

const file_eolymp_acl_policy_service_proto_rawDesc = "" +
	"\n" +
	"\x1feolymp/acl/policy_service.proto\x12\n" +
	"eolymp.acl\x1a\x17eolymp/acl/policy.proto\x1a\x1deolymp/annotations/http.proto\x1a\"eolymp/annotations/ratelimit.proto\x1a!eolymp/wellknown/expression.proto\"?\n" +
	"\x11CreatePolicyInput\x12*\n" +
	"\x06policy\x18\x01 \x01(\v2\x12.eolymp.acl.PolicyR\x06policy\"1\n" +
	"\x12CreatePolicyOutput\x12\x1b\n" +
	"\tpolicy_id\x18\x01 \x01(\tR\bpolicyId\"\xbf\x01\n" +
	"\x11UpdatePolicyInput\x129\n" +
	"\x05patch\x18\n" +
	" \x03(\x0e2#.eolymp.acl.UpdatePolicyInput.PatchR\x05patch\x12\x1b\n" +
	"\tpolicy_id\x18\x01 \x01(\tR\bpolicyId\x12*\n" +
	"\x06policy\x18\x02 \x01(\v2\x12.eolymp.acl.PolicyR\x06policy\"&\n" +
	"\x05Patch\x12\a\n" +
	"\x03ALL\x10\x00\x12\b\n" +
	"\x04NAME\x10\x01\x12\n" +
	"\n" +
	"\x06ALLOWS\x10\x02\"\x14\n" +
	"\x12UpdatePolicyOutput\"0\n" +
	"\x11DeletePolicyInput\x12\x1b\n" +
	"\tpolicy_id\x18\x01 \x01(\tR\bpolicyId\"\x14\n" +
	"\x12DeletePolicyOutput\"2\n" +
	"\x13DescribePolicyInput\x12\x1b\n" +
	"\tpolicy_id\x18\x01 \x01(\tR\bpolicyId\"B\n" +
	"\x14DescribePolicyOutput\x12*\n" +
	"\x06policy\x18\x01 \x01(\v2\x12.eolymp.acl.PolicyR\x06policy\"\x84\x03\n" +
	"\x11ListPoliciesInput\x12\x16\n" +
	"\x06offset\x18\n" +
	" \x01(\x05R\x06offset\x12\x12\n" +
	"\x04size\x18\v \x01(\x05R\x04size\x12>\n" +
	"\afilters\x18( \x01(\v2$.eolymp.acl.ListPoliciesInput.FilterR\afilters\x1a\x82\x02\n" +
	"\x06Filter\x12\x14\n" +
	"\x05query\x18\n" +
	" \x01(\tR\x05query\x12.\n" +
	"\x02id\x18\x01 \x03(\v2\x1e.eolymp.wellknown.ExpressionIDR\x02id\x12<\n" +
	"\tprincipal\x18\x02 \x03(\v2\x1e.eolymp.wellknown.ExpressionIDR\tprincipal\x126\n" +
	"\x04name\x18\x03 \x03(\v2\".eolymp.wellknown.ExpressionStringR\x04name\x12<\n" +
	"\bresource\x18\x04 \x03(\v2 .eolymp.wellknown.ExpressionEnumR\bresource\"T\n" +
	"\x12ListPoliciesOutput\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x05R\x05total\x12(\n" +
	"\x05items\x18\x02 \x03(\v2\x12.eolymp.acl.PolicyR\x05items\"\xa3\x01\n" +
	"\x11CopyPoliciesInput\x12#\n" +
	"\rsrc_principal\x18\x01 \x01(\tR\fsrcPrincipal\x12#\n" +
	"\rdst_principal\x18\x02 \x01(\tR\fdstPrincipal\x12!\n" +
	"\fsrc_resource\x18\x03 \x01(\tR\vsrcResource\x12!\n" +
	"\fdst_resource\x18\x04 \x01(\tR\vdstResource\"7\n" +
	"\x12CopyPoliciesOutput\x12!\n" +
	"\fcopies_count\x18\x01 \x01(\x05R\vcopiesCount2\xe5\x05\n" +
	"\rPolicyService\x12o\n" +
	"\fCreatePolicy\x12\x1d.eolymp.acl.CreatePolicyInput\x1a\x1e.eolymp.acl.CreatePolicyOutput\" \xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xd3\xe4\x93\x02\v\"\t/policies\x12{\n" +
	"\fUpdatePolicy\x12\x1d.eolymp.acl.UpdatePolicyInput\x1a\x1e.eolymp.acl.UpdatePolicyOutput\",\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xd3\xe4\x93\x02\x17\x1a\x15/policies/{policy_id}\x12{\n" +
	"\fDeletePolicy\x12\x1d.eolymp.acl.DeletePolicyInput\x1a\x1e.eolymp.acl.DeletePolicyOutput\",\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00\xa0@\xf8\xe2\n" +
	"\x14\x82\xd3\xe4\x93\x02\x17*\x15/policies/{policy_id}\x12\x81\x01\n" +
	"\x0eDescribePolicy\x12\x1f.eolymp.acl.DescribePolicyInput\x1a .eolymp.acl.DescribePolicyOutput\",\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00 A\xf8\xe2\n" +
	"\x14\x82\xd3\xe4\x93\x02\x17\x12\x15/policies/{policy_id}\x12o\n" +
	"\fListPolicies\x12\x1d.eolymp.acl.ListPoliciesInput\x1a\x1e.eolymp.acl.ListPoliciesOutput\" \xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00 A\xf8\xe2\n" +
	"\x14\x82\xd3\xe4\x93\x02\v\x12\t/policies\x12t\n" +
	"\fCopyPolicies\x12\x1d.eolymp.acl.CopyPoliciesInput\x1a\x1e.eolymp.acl.CopyPoliciesOutput\"%\xea\xe2\n" +
	"\v\xf5\xe2\n" +
	"\x00\x00 A\xf8\xe2\n" +
	"\x14\x82\xd3\xe4\x93\x02\x10\"\x0e/policies:copyB)Z'github.com/eolymp/go-sdk/eolymp/acl;aclb\x06proto3"

var (
	file_eolymp_acl_policy_service_proto_rawDescOnce sync.Once
	file_eolymp_acl_policy_service_proto_rawDescData []byte
)

func file_eolymp_acl_policy_service_proto_rawDescGZIP() []byte {
	file_eolymp_acl_policy_service_proto_rawDescOnce.Do(func() {
		file_eolymp_acl_policy_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_acl_policy_service_proto_rawDesc), len(file_eolymp_acl_policy_service_proto_rawDesc)))
	})
	return file_eolymp_acl_policy_service_proto_rawDescData
}

var file_eolymp_acl_policy_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_acl_policy_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_eolymp_acl_policy_service_proto_goTypes = []any{
	(UpdatePolicyInput_Patch)(0),       // 0: eolymp.acl.UpdatePolicyInput.Patch
	(*CreatePolicyInput)(nil),          // 1: eolymp.acl.CreatePolicyInput
	(*CreatePolicyOutput)(nil),         // 2: eolymp.acl.CreatePolicyOutput
	(*UpdatePolicyInput)(nil),          // 3: eolymp.acl.UpdatePolicyInput
	(*UpdatePolicyOutput)(nil),         // 4: eolymp.acl.UpdatePolicyOutput
	(*DeletePolicyInput)(nil),          // 5: eolymp.acl.DeletePolicyInput
	(*DeletePolicyOutput)(nil),         // 6: eolymp.acl.DeletePolicyOutput
	(*DescribePolicyInput)(nil),        // 7: eolymp.acl.DescribePolicyInput
	(*DescribePolicyOutput)(nil),       // 8: eolymp.acl.DescribePolicyOutput
	(*ListPoliciesInput)(nil),          // 9: eolymp.acl.ListPoliciesInput
	(*ListPoliciesOutput)(nil),         // 10: eolymp.acl.ListPoliciesOutput
	(*CopyPoliciesInput)(nil),          // 11: eolymp.acl.CopyPoliciesInput
	(*CopyPoliciesOutput)(nil),         // 12: eolymp.acl.CopyPoliciesOutput
	(*ListPoliciesInput_Filter)(nil),   // 13: eolymp.acl.ListPoliciesInput.Filter
	(*Policy)(nil),                     // 14: eolymp.acl.Policy
	(*wellknown.ExpressionID)(nil),     // 15: eolymp.wellknown.ExpressionID
	(*wellknown.ExpressionString)(nil), // 16: eolymp.wellknown.ExpressionString
	(*wellknown.ExpressionEnum)(nil),   // 17: eolymp.wellknown.ExpressionEnum
}
var file_eolymp_acl_policy_service_proto_depIdxs = []int32{
	14, // 0: eolymp.acl.CreatePolicyInput.policy:type_name -> eolymp.acl.Policy
	0,  // 1: eolymp.acl.UpdatePolicyInput.patch:type_name -> eolymp.acl.UpdatePolicyInput.Patch
	14, // 2: eolymp.acl.UpdatePolicyInput.policy:type_name -> eolymp.acl.Policy
	14, // 3: eolymp.acl.DescribePolicyOutput.policy:type_name -> eolymp.acl.Policy
	13, // 4: eolymp.acl.ListPoliciesInput.filters:type_name -> eolymp.acl.ListPoliciesInput.Filter
	14, // 5: eolymp.acl.ListPoliciesOutput.items:type_name -> eolymp.acl.Policy
	15, // 6: eolymp.acl.ListPoliciesInput.Filter.id:type_name -> eolymp.wellknown.ExpressionID
	15, // 7: eolymp.acl.ListPoliciesInput.Filter.principal:type_name -> eolymp.wellknown.ExpressionID
	16, // 8: eolymp.acl.ListPoliciesInput.Filter.name:type_name -> eolymp.wellknown.ExpressionString
	17, // 9: eolymp.acl.ListPoliciesInput.Filter.resource:type_name -> eolymp.wellknown.ExpressionEnum
	1,  // 10: eolymp.acl.PolicyService.CreatePolicy:input_type -> eolymp.acl.CreatePolicyInput
	3,  // 11: eolymp.acl.PolicyService.UpdatePolicy:input_type -> eolymp.acl.UpdatePolicyInput
	5,  // 12: eolymp.acl.PolicyService.DeletePolicy:input_type -> eolymp.acl.DeletePolicyInput
	7,  // 13: eolymp.acl.PolicyService.DescribePolicy:input_type -> eolymp.acl.DescribePolicyInput
	9,  // 14: eolymp.acl.PolicyService.ListPolicies:input_type -> eolymp.acl.ListPoliciesInput
	11, // 15: eolymp.acl.PolicyService.CopyPolicies:input_type -> eolymp.acl.CopyPoliciesInput
	2,  // 16: eolymp.acl.PolicyService.CreatePolicy:output_type -> eolymp.acl.CreatePolicyOutput
	4,  // 17: eolymp.acl.PolicyService.UpdatePolicy:output_type -> eolymp.acl.UpdatePolicyOutput
	6,  // 18: eolymp.acl.PolicyService.DeletePolicy:output_type -> eolymp.acl.DeletePolicyOutput
	8,  // 19: eolymp.acl.PolicyService.DescribePolicy:output_type -> eolymp.acl.DescribePolicyOutput
	10, // 20: eolymp.acl.PolicyService.ListPolicies:output_type -> eolymp.acl.ListPoliciesOutput
	12, // 21: eolymp.acl.PolicyService.CopyPolicies:output_type -> eolymp.acl.CopyPoliciesOutput
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_eolymp_acl_policy_service_proto_init() }
func file_eolymp_acl_policy_service_proto_init() {
	if File_eolymp_acl_policy_service_proto != nil {
		return
	}
	file_eolymp_acl_policy_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_acl_policy_service_proto_rawDesc), len(file_eolymp_acl_policy_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_acl_policy_service_proto_goTypes,
		DependencyIndexes: file_eolymp_acl_policy_service_proto_depIdxs,
		EnumInfos:         file_eolymp_acl_policy_service_proto_enumTypes,
		MessageInfos:      file_eolymp_acl_policy_service_proto_msgTypes,
	}.Build()
	File_eolymp_acl_policy_service_proto = out.File
	file_eolymp_acl_policy_service_proto_goTypes = nil
	file_eolymp_acl_policy_service_proto_depIdxs = nil
}
