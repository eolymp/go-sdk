// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: eolymp/oauth2/oauth2.proto

package oauth2

import (
	_ "github.com/eolymp/go-sdk/eolymp/annotations"
	_ "github.com/eolymp/go-sdk/eolymp/wellknown"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenInput_GrantType int32

const (
	TokenInput_NONE               TokenInput_GrantType = 0
	TokenInput_PASSWORD           TokenInput_GrantType = 1
	TokenInput_AUTHORIZATION_CODE TokenInput_GrantType = 2
	TokenInput_REFRESH_TOKEN      TokenInput_GrantType = 3
	TokenInput_GOOGLE_CODE        TokenInput_GrantType = 4 // deprecated: to be removed
)

// Enum value maps for TokenInput_GrantType.
var (
	TokenInput_GrantType_name = map[int32]string{
		0: "NONE",
		1: "PASSWORD",
		2: "AUTHORIZATION_CODE",
		3: "REFRESH_TOKEN",
		4: "GOOGLE_CODE",
	}
	TokenInput_GrantType_value = map[string]int32{
		"NONE":               0,
		"PASSWORD":           1,
		"AUTHORIZATION_CODE": 2,
		"REFRESH_TOKEN":      3,
		"GOOGLE_CODE":        4,
	}
)

func (x TokenInput_GrantType) Enum() *TokenInput_GrantType {
	p := new(TokenInput_GrantType)
	*p = x
	return p
}

func (x TokenInput_GrantType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenInput_GrantType) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_oauth2_oauth2_proto_enumTypes[0].Descriptor()
}

func (TokenInput_GrantType) Type() protoreflect.EnumType {
	return &file_eolymp_oauth2_oauth2_proto_enumTypes[0]
}

func (x TokenInput_GrantType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenInput_GrantType.Descriptor instead.
func (TokenInput_GrantType) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_oauth2_oauth2_proto_rawDescGZIP(), []int{0, 0}
}

type TokenInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrantType    TokenInput_GrantType `protobuf:"varint,1,opt,name=grant_type,json=grantType,proto3,enum=eolymp.oauth2.TokenInput_GrantType" json:"grant_type,omitempty"`
	Username     string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"` // username for password grant type
	Password     string               `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"` // password for password grant type
	ClientId     string               `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string               `protobuf:"bytes,5,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Code         string               `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`                                     // code for authorization_code grant type
	CodeVerifier string               `protobuf:"bytes,7,opt,name=code_verifier,json=codeVerifier,proto3" json:"code_verifier,omitempty"` // code verifier for authorization_code grant type
	Scope        string               `protobuf:"bytes,8,opt,name=scope,proto3" json:"scope,omitempty"`
	RefreshToken string               `protobuf:"bytes,9,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"` // refresh_token for refresh_token grant type
	RedirectUri  string               `protobuf:"bytes,10,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
}

func (x *TokenInput) Reset() {
	*x = TokenInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInput) ProtoMessage() {}

func (x *TokenInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInput.ProtoReflect.Descriptor instead.
func (*TokenInput) Descriptor() ([]byte, []int) {
	return file_eolymp_oauth2_oauth2_proto_rawDescGZIP(), []int{0}
}

func (x *TokenInput) GetGrantType() TokenInput_GrantType {
	if x != nil {
		return x.GrantType
	}
	return TokenInput_NONE
}

func (x *TokenInput) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TokenInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *TokenInput) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *TokenInput) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *TokenInput) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TokenInput) GetCodeVerifier() string {
	if x != nil {
		return x.CodeVerifier
	}
	return ""
}

func (x *TokenInput) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *TokenInput) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *TokenInput) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type TokenOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	TokenType    string `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	ExpiresIn    uint32 `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	RefreshToken string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Scope        string `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope,omitempty"`
	IdToken      string `protobuf:"bytes,100,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
}

func (x *TokenOutput) Reset() {
	*x = TokenOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenOutput) ProtoMessage() {}

func (x *TokenOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenOutput.ProtoReflect.Descriptor instead.
func (*TokenOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_oauth2_oauth2_proto_rawDescGZIP(), []int{1}
}

func (x *TokenOutput) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokenOutput) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *TokenOutput) GetExpiresIn() uint32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *TokenOutput) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *TokenOutput) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *TokenOutput) GetIdToken() string {
	if x != nil {
		return x.IdToken
	}
	return ""
}

type AuthorizeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId            string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	CodeChallenge       string `protobuf:"bytes,2,opt,name=code_challenge,json=codeChallenge,proto3" json:"code_challenge,omitempty"`
	CodeChallengeMethod string `protobuf:"bytes,3,opt,name=code_challenge_method,json=codeChallengeMethod,proto3" json:"code_challenge_method,omitempty"`
	RedirectUri         string `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	ResponseType        string `protobuf:"bytes,5,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
	Scope               string `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope,omitempty"`
	State               string `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *AuthorizeInput) Reset() {
	*x = AuthorizeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeInput) ProtoMessage() {}

func (x *AuthorizeInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeInput.ProtoReflect.Descriptor instead.
func (*AuthorizeInput) Descriptor() ([]byte, []int) {
	return file_eolymp_oauth2_oauth2_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorizeInput) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AuthorizeInput) GetCodeChallenge() string {
	if x != nil {
		return x.CodeChallenge
	}
	return ""
}

func (x *AuthorizeInput) GetCodeChallengeMethod() string {
	if x != nil {
		return x.CodeChallengeMethod
	}
	return ""
}

func (x *AuthorizeInput) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *AuthorizeInput) GetResponseType() string {
	if x != nil {
		return x.ResponseType
	}
	return ""
}

func (x *AuthorizeInput) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *AuthorizeInput) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type AuthorizeOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUri string `protobuf:"bytes,1,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
}

func (x *AuthorizeOutput) Reset() {
	*x = AuthorizeOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeOutput) ProtoMessage() {}

func (x *AuthorizeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeOutput.ProtoReflect.Descriptor instead.
func (*AuthorizeOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_oauth2_oauth2_proto_rawDescGZIP(), []int{3}
}

func (x *AuthorizeOutput) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type CallbackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	State string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *CallbackInput) Reset() {
	*x = CallbackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackInput) ProtoMessage() {}

func (x *CallbackInput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackInput.ProtoReflect.Descriptor instead.
func (*CallbackInput) Descriptor() ([]byte, []int) {
	return file_eolymp_oauth2_oauth2_proto_rawDescGZIP(), []int{4}
}

func (x *CallbackInput) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CallbackInput) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type CallbackOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUri string `protobuf:"bytes,1,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
}

func (x *CallbackOutput) Reset() {
	*x = CallbackOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackOutput) ProtoMessage() {}

func (x *CallbackOutput) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_oauth2_oauth2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackOutput.ProtoReflect.Descriptor instead.
func (*CallbackOutput) Descriptor() ([]byte, []int) {
	return file_eolymp_oauth2_oauth2_proto_rawDescGZIP(), []int{5}
}

func (x *CallbackOutput) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

var File_eolymp_oauth2_oauth2_proto protoreflect.FileDescriptor

var file_eolymp_oauth2_oauth2_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x1a, 0x1d, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc2, 0x03, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x42, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x64, 0x65,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x22, 0x5f, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x41,
	0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x5f, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x04, 0x22, 0xc4, 0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xfc,
	0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x70, 0x75,
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
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x34, 0x0a,
	0x0f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x69, 0x22, 0x39, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x33,
	0x0a, 0x0e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x69, 0x32, 0xd9, 0x02, 0x0a, 0x06, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0x65,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x32, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x25,
	0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xf0, 0x41, 0xf8, 0xe2, 0x0a, 0xac, 0x02,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x75, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x12, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x29, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00, 0xf0, 0x41, 0xf8, 0xe2,
	0x0a, 0xac, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x71, 0x0a, 0x08,
	0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1c, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x28, 0xea, 0xe2, 0x0a, 0x0c, 0xf5, 0xe2, 0x0a, 0x00, 0x00,
	0xf0, 0x41, 0xf8, 0xe2, 0x0a, 0xac, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x42,
	0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79,
	0x6d, 0x70, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x3b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_oauth2_oauth2_proto_rawDescOnce sync.Once
	file_eolymp_oauth2_oauth2_proto_rawDescData = file_eolymp_oauth2_oauth2_proto_rawDesc
)

func file_eolymp_oauth2_oauth2_proto_rawDescGZIP() []byte {
	file_eolymp_oauth2_oauth2_proto_rawDescOnce.Do(func() {
		file_eolymp_oauth2_oauth2_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_oauth2_oauth2_proto_rawDescData)
	})
	return file_eolymp_oauth2_oauth2_proto_rawDescData
}

var file_eolymp_oauth2_oauth2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_oauth2_oauth2_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eolymp_oauth2_oauth2_proto_goTypes = []interface{}{
	(TokenInput_GrantType)(0), // 0: eolymp.oauth2.TokenInput.GrantType
	(*TokenInput)(nil),        // 1: eolymp.oauth2.TokenInput
	(*TokenOutput)(nil),       // 2: eolymp.oauth2.TokenOutput
	(*AuthorizeInput)(nil),    // 3: eolymp.oauth2.AuthorizeInput
	(*AuthorizeOutput)(nil),   // 4: eolymp.oauth2.AuthorizeOutput
	(*CallbackInput)(nil),     // 5: eolymp.oauth2.CallbackInput
	(*CallbackOutput)(nil),    // 6: eolymp.oauth2.CallbackOutput
}
var file_eolymp_oauth2_oauth2_proto_depIdxs = []int32{
	0, // 0: eolymp.oauth2.TokenInput.grant_type:type_name -> eolymp.oauth2.TokenInput.GrantType
	1, // 1: eolymp.oauth2.OAuth2.Token:input_type -> eolymp.oauth2.TokenInput
	3, // 2: eolymp.oauth2.OAuth2.Authorize:input_type -> eolymp.oauth2.AuthorizeInput
	5, // 3: eolymp.oauth2.OAuth2.Callback:input_type -> eolymp.oauth2.CallbackInput
	2, // 4: eolymp.oauth2.OAuth2.Token:output_type -> eolymp.oauth2.TokenOutput
	4, // 5: eolymp.oauth2.OAuth2.Authorize:output_type -> eolymp.oauth2.AuthorizeOutput
	6, // 6: eolymp.oauth2.OAuth2.Callback:output_type -> eolymp.oauth2.CallbackOutput
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eolymp_oauth2_oauth2_proto_init() }
func file_eolymp_oauth2_oauth2_proto_init() {
	if File_eolymp_oauth2_oauth2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eolymp_oauth2_oauth2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenInput); i {
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
		file_eolymp_oauth2_oauth2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenOutput); i {
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
		file_eolymp_oauth2_oauth2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeInput); i {
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
		file_eolymp_oauth2_oauth2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeOutput); i {
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
		file_eolymp_oauth2_oauth2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackInput); i {
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
		file_eolymp_oauth2_oauth2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackOutput); i {
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
			RawDescriptor: file_eolymp_oauth2_oauth2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_oauth2_oauth2_proto_goTypes,
		DependencyIndexes: file_eolymp_oauth2_oauth2_proto_depIdxs,
		EnumInfos:         file_eolymp_oauth2_oauth2_proto_enumTypes,
		MessageInfos:      file_eolymp_oauth2_oauth2_proto_msgTypes,
	}.Build()
	File_eolymp_oauth2_oauth2_proto = out.File
	file_eolymp_oauth2_oauth2_proto_rawDesc = nil
	file_eolymp_oauth2_oauth2_proto_goTypes = nil
	file_eolymp_oauth2_oauth2_proto_depIdxs = nil
}
