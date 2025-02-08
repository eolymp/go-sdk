// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/auth/oauth2_service.proto

package auth

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type IssueTokenInput_GrantType int32

const (
	IssueTokenInput_NONE               IssueTokenInput_GrantType = 0
	IssueTokenInput_PASSWORD           IssueTokenInput_GrantType = 1
	IssueTokenInput_AUTHORIZATION_CODE IssueTokenInput_GrantType = 2
	IssueTokenInput_REFRESH_TOKEN      IssueTokenInput_GrantType = 3
	IssueTokenInput_EOLYMP_SIGNIN      IssueTokenInput_GrantType = 4
)

// Enum value maps for IssueTokenInput_GrantType.
var (
	IssueTokenInput_GrantType_name = map[int32]string{
		0: "NONE",
		1: "PASSWORD",
		2: "AUTHORIZATION_CODE",
		3: "REFRESH_TOKEN",
		4: "EOLYMP_SIGNIN",
	}
	IssueTokenInput_GrantType_value = map[string]int32{
		"NONE":               0,
		"PASSWORD":           1,
		"AUTHORIZATION_CODE": 2,
		"REFRESH_TOKEN":      3,
		"EOLYMP_SIGNIN":      4,
	}
)

func (x IssueTokenInput_GrantType) Enum() *IssueTokenInput_GrantType {
	p := new(IssueTokenInput_GrantType)
	*p = x
	return p
}

func (x IssueTokenInput_GrantType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IssueTokenInput_GrantType) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_auth_oauth2_service_proto_enumTypes[0].Descriptor()
}

func (IssueTokenInput_GrantType) Type() protoreflect.EnumType {
	return &file_eolymp_auth_oauth2_service_proto_enumTypes[0]
}

func (x IssueTokenInput_GrantType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IssueTokenInput_GrantType.Descriptor instead.
func (IssueTokenInput_GrantType) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_auth_oauth2_service_proto_rawDescGZIP(), []int{0, 0}
}

type IssueTokenInput struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	GrantType     IssueTokenInput_GrantType `protobuf:"varint,1,opt,name=grant_type,json=grantType,proto3,enum=eolymp.auth.IssueTokenInput_GrantType" json:"grant_type,omitempty"`
	Username      string                    `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"` // username for password grant type
	Password      string                    `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"` // password for password grant type
	ClientId      string                    `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret  string                    `protobuf:"bytes,5,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Code          string                    `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`                                     // code for authorization_code grant type
	CodeVerifier  string                    `protobuf:"bytes,7,opt,name=code_verifier,json=codeVerifier,proto3" json:"code_verifier,omitempty"` // code verifier for authorization_code grant type
	Scope         string                    `protobuf:"bytes,8,opt,name=scope,proto3" json:"scope,omitempty"`
	RefreshToken  string                    `protobuf:"bytes,9,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"` // refresh_token for refresh_token grant type
	RedirectUri   string                    `protobuf:"bytes,10,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssueTokenInput) Reset() {
	*x = IssueTokenInput{}
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueTokenInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueTokenInput) ProtoMessage() {}

func (x *IssueTokenInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueTokenInput.ProtoReflect.Descriptor instead.
func (*IssueTokenInput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_oauth2_service_proto_rawDescGZIP(), []int{0}
}

func (x *IssueTokenInput) GetGrantType() IssueTokenInput_GrantType {
	if x != nil {
		return x.GrantType
	}
	return IssueTokenInput_NONE
}

func (x *IssueTokenInput) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *IssueTokenInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *IssueTokenInput) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *IssueTokenInput) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *IssueTokenInput) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *IssueTokenInput) GetCodeVerifier() string {
	if x != nil {
		return x.CodeVerifier
	}
	return ""
}

func (x *IssueTokenInput) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *IssueTokenInput) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *IssueTokenInput) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type IssueTokenOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	TokenType     string                 `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	ExpiresIn     uint32                 `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Scope         string                 `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope,omitempty"`
	Claims        *Claims                `protobuf:"bytes,10,opt,name=claims,proto3" json:"claims,omitempty"` // temporarily expose claims directly without wrapping them into id_token
	IdToken       string                 `protobuf:"bytes,100,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssueTokenOutput) Reset() {
	*x = IssueTokenOutput{}
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueTokenOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueTokenOutput) ProtoMessage() {}

func (x *IssueTokenOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueTokenOutput.ProtoReflect.Descriptor instead.
func (*IssueTokenOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_oauth2_service_proto_rawDescGZIP(), []int{1}
}

func (x *IssueTokenOutput) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *IssueTokenOutput) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *IssueTokenOutput) GetExpiresIn() uint32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *IssueTokenOutput) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *IssueTokenOutput) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *IssueTokenOutput) GetClaims() *Claims {
	if x != nil {
		return x.Claims
	}
	return nil
}

func (x *IssueTokenOutput) GetIdToken() string {
	if x != nil {
		return x.IdToken
	}
	return ""
}

type IntrospectTokenInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IntrospectTokenInput) Reset() {
	*x = IntrospectTokenInput{}
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntrospectTokenInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntrospectTokenInput) ProtoMessage() {}

func (x *IntrospectTokenInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntrospectTokenInput.ProtoReflect.Descriptor instead.
func (*IntrospectTokenInput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_oauth2_service_proto_rawDescGZIP(), []int{2}
}

func (x *IntrospectTokenInput) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type IntrospectTokenOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Active        bool                   `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	Scope         string                 `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	Expire        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expire,proto3" json:"expire,omitempty"`
	Claims        *Claims                `protobuf:"bytes,10,opt,name=claims,proto3" json:"claims,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IntrospectTokenOutput) Reset() {
	*x = IntrospectTokenOutput{}
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntrospectTokenOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntrospectTokenOutput) ProtoMessage() {}

func (x *IntrospectTokenOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntrospectTokenOutput.ProtoReflect.Descriptor instead.
func (*IntrospectTokenOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_oauth2_service_proto_rawDescGZIP(), []int{3}
}

func (x *IntrospectTokenOutput) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *IntrospectTokenOutput) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *IntrospectTokenOutput) GetExpire() *timestamppb.Timestamp {
	if x != nil {
		return x.Expire
	}
	return nil
}

func (x *IntrospectTokenOutput) GetClaims() *Claims {
	if x != nil {
		return x.Claims
	}
	return nil
}

type RevokeTokenInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevokeTokenInput) Reset() {
	*x = RevokeTokenInput{}
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeTokenInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeTokenInput) ProtoMessage() {}

func (x *RevokeTokenInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeTokenInput.ProtoReflect.Descriptor instead.
func (*RevokeTokenInput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_oauth2_service_proto_rawDescGZIP(), []int{4}
}

func (x *RevokeTokenInput) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RevokeTokenOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevokeTokenOutput) Reset() {
	*x = RevokeTokenOutput{}
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeTokenOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeTokenOutput) ProtoMessage() {}

func (x *RevokeTokenOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeTokenOutput.ProtoReflect.Descriptor instead.
func (*RevokeTokenOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_oauth2_service_proto_rawDescGZIP(), []int{5}
}

type RequestAuthInput struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	ClientId            string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	CodeChallenge       string                 `protobuf:"bytes,2,opt,name=code_challenge,json=codeChallenge,proto3" json:"code_challenge,omitempty"`
	CodeChallengeMethod string                 `protobuf:"bytes,3,opt,name=code_challenge_method,json=codeChallengeMethod,proto3" json:"code_challenge_method,omitempty"`
	RedirectUri         string                 `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	ResponseType        string                 `protobuf:"bytes,5,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
	Scope               string                 `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope,omitempty"`
	State               string                 `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *RequestAuthInput) Reset() {
	*x = RequestAuthInput{}
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestAuthInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAuthInput) ProtoMessage() {}

func (x *RequestAuthInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAuthInput.ProtoReflect.Descriptor instead.
func (*RequestAuthInput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_oauth2_service_proto_rawDescGZIP(), []int{6}
}

func (x *RequestAuthInput) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RequestAuthInput) GetCodeChallenge() string {
	if x != nil {
		return x.CodeChallenge
	}
	return ""
}

func (x *RequestAuthInput) GetCodeChallengeMethod() string {
	if x != nil {
		return x.CodeChallengeMethod
	}
	return ""
}

func (x *RequestAuthInput) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *RequestAuthInput) GetResponseType() string {
	if x != nil {
		return x.ResponseType
	}
	return ""
}

func (x *RequestAuthInput) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *RequestAuthInput) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type RequestAuthOutput struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	AuthorizationCode string                 `protobuf:"bytes,1,opt,name=authorization_code,json=authorizationCode,proto3" json:"authorization_code,omitempty"`
	RedirectUri       string                 `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *RequestAuthOutput) Reset() {
	*x = RequestAuthOutput{}
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestAuthOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAuthOutput) ProtoMessage() {}

func (x *RequestAuthOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAuthOutput.ProtoReflect.Descriptor instead.
func (*RequestAuthOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_oauth2_service_proto_rawDescGZIP(), []int{7}
}

func (x *RequestAuthOutput) GetAuthorizationCode() string {
	if x != nil {
		return x.AuthorizationCode
	}
	return ""
}

func (x *RequestAuthOutput) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type UserInfoInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoInput) Reset() {
	*x = UserInfoInput{}
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoInput) ProtoMessage() {}

func (x *UserInfoInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoInput.ProtoReflect.Descriptor instead.
func (*UserInfoInput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_oauth2_service_proto_rawDescGZIP(), []int{8}
}

type UserInfoOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Claims        *Claims                `protobuf:"bytes,10,opt,name=claims,proto3" json:"claims,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoOutput) Reset() {
	*x = UserInfoOutput{}
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoOutput) ProtoMessage() {}

func (x *UserInfoOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_auth_oauth2_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoOutput.ProtoReflect.Descriptor instead.
func (*UserInfoOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_auth_oauth2_service_proto_rawDescGZIP(), []int{9}
}

func (x *UserInfoOutput) GetClaims() *Claims {
	if x != nil {
		return x.Claims
	}
	return nil
}

var File_eolymp_auth_oauth2_service_proto protoreflect.FileDescriptor

var file_eolymp_auth_oauth2_service_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a,
	0x22, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc,
	0x03, 0x0a, 0x0f, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6f, 0x64, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x22, 0x61, 0x0a, 0x09, 0x47, 0x72,
	0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x01, 0x12,
	0x16, 0x0a, 0x12, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x46, 0x52, 0x45,
	0x53, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x4f,
	0x4c, 0x59, 0x4d, 0x50, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x49, 0x4e, 0x10, 0x04, 0x22, 0xf6, 0x01,
	0x0a, 0x10, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x49, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x2b,
	0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x73, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2c, 0x0a, 0x14, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73,
	0x70, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa6, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x12, 0x2b, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x73, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x22, 0x28, 0x0a,
	0x10, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0xfe, 0x01, 0x0a,
	0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x65, 0x0a,
	0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x55, 0x72, 0x69, 0x22, 0x0f, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x3d, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x52, 0x06, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x32, 0xef, 0x03, 0x0a, 0x0d, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0a, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x10, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xf0, 0x41, 0xf8, 0xe2,
	0x0a, 0xac, 0x02, 0x12, 0x6a, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x10, 0xea,
	0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xf0, 0x41, 0xf8, 0xe2, 0x0a, 0xac, 0x02, 0x12,
	0x5e, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1e, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x10, 0xea,
	0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xf0, 0x41, 0xf8, 0xe2, 0x0a, 0xac, 0x02, 0x12,
	0x5e, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1d,
	0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1e, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x10, 0xea,
	0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xf0, 0x41, 0xf8, 0xe2, 0x0a, 0xac, 0x02, 0x12,
	0x55, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x10, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xf0,
	0x41, 0xf8, 0xe2, 0x0a, 0xac, 0x02, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x61,
	0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_auth_oauth2_service_proto_rawDescOnce sync.Once
	file_eolymp_auth_oauth2_service_proto_rawDescData []byte
)

func file_eolymp_auth_oauth2_service_proto_rawDescGZIP() []byte {
	file_eolymp_auth_oauth2_service_proto_rawDescOnce.Do(func() {
		file_eolymp_auth_oauth2_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_auth_oauth2_service_proto_rawDesc), len(file_eolymp_auth_oauth2_service_proto_rawDesc)))
	})
	return file_eolymp_auth_oauth2_service_proto_rawDescData
}

var file_eolymp_auth_oauth2_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_auth_oauth2_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_eolymp_auth_oauth2_service_proto_goTypes = []any{
	(IssueTokenInput_GrantType)(0), // 0: eolymp.auth.IssueTokenInput.GrantType
	(*IssueTokenInput)(nil),        // 1: eolymp.auth.IssueTokenInput
	(*IssueTokenOutput)(nil),       // 2: eolymp.auth.IssueTokenOutput
	(*IntrospectTokenInput)(nil),   // 3: eolymp.auth.IntrospectTokenInput
	(*IntrospectTokenOutput)(nil),  // 4: eolymp.auth.IntrospectTokenOutput
	(*RevokeTokenInput)(nil),       // 5: eolymp.auth.RevokeTokenInput
	(*RevokeTokenOutput)(nil),      // 6: eolymp.auth.RevokeTokenOutput
	(*RequestAuthInput)(nil),       // 7: eolymp.auth.RequestAuthInput
	(*RequestAuthOutput)(nil),      // 8: eolymp.auth.RequestAuthOutput
	(*UserInfoInput)(nil),          // 9: eolymp.auth.UserInfoInput
	(*UserInfoOutput)(nil),         // 10: eolymp.auth.UserInfoOutput
	(*Claims)(nil),                 // 11: eolymp.auth.Claims
	(*timestamppb.Timestamp)(nil),  // 12: google.protobuf.Timestamp
}
var file_eolymp_auth_oauth2_service_proto_depIdxs = []int32{
	0,  // 0: eolymp.auth.IssueTokenInput.grant_type:type_name -> eolymp.auth.IssueTokenInput.GrantType
	11, // 1: eolymp.auth.IssueTokenOutput.claims:type_name -> eolymp.auth.Claims
	12, // 2: eolymp.auth.IntrospectTokenOutput.expire:type_name -> google.protobuf.Timestamp
	11, // 3: eolymp.auth.IntrospectTokenOutput.claims:type_name -> eolymp.auth.Claims
	11, // 4: eolymp.auth.UserInfoOutput.claims:type_name -> eolymp.auth.Claims
	1,  // 5: eolymp.auth.OAuth2Service.IssueToken:input_type -> eolymp.auth.IssueTokenInput
	3,  // 6: eolymp.auth.OAuth2Service.IntrospectToken:input_type -> eolymp.auth.IntrospectTokenInput
	5,  // 7: eolymp.auth.OAuth2Service.RevokeToken:input_type -> eolymp.auth.RevokeTokenInput
	7,  // 8: eolymp.auth.OAuth2Service.RequestAuth:input_type -> eolymp.auth.RequestAuthInput
	9,  // 9: eolymp.auth.OAuth2Service.UserInfo:input_type -> eolymp.auth.UserInfoInput
	2,  // 10: eolymp.auth.OAuth2Service.IssueToken:output_type -> eolymp.auth.IssueTokenOutput
	4,  // 11: eolymp.auth.OAuth2Service.IntrospectToken:output_type -> eolymp.auth.IntrospectTokenOutput
	6,  // 12: eolymp.auth.OAuth2Service.RevokeToken:output_type -> eolymp.auth.RevokeTokenOutput
	8,  // 13: eolymp.auth.OAuth2Service.RequestAuth:output_type -> eolymp.auth.RequestAuthOutput
	10, // 14: eolymp.auth.OAuth2Service.UserInfo:output_type -> eolymp.auth.UserInfoOutput
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_eolymp_auth_oauth2_service_proto_init() }
func file_eolymp_auth_oauth2_service_proto_init() {
	if File_eolymp_auth_oauth2_service_proto != nil {
		return
	}
	file_eolymp_auth_claims_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_auth_oauth2_service_proto_rawDesc), len(file_eolymp_auth_oauth2_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_auth_oauth2_service_proto_goTypes,
		DependencyIndexes: file_eolymp_auth_oauth2_service_proto_depIdxs,
		EnumInfos:         file_eolymp_auth_oauth2_service_proto_enumTypes,
		MessageInfos:      file_eolymp_auth_oauth2_service_proto_msgTypes,
	}.Build()
	File_eolymp_auth_oauth2_service_proto = out.File
	file_eolymp_auth_oauth2_service_proto_goTypes = nil
	file_eolymp_auth_oauth2_service_proto_depIdxs = nil
}
