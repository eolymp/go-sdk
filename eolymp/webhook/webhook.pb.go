// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: eolymp/webhook/webhook.proto

package webhook

import (
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

type Webhook_Event int32

const (
	Webhook_UNKNOWN_EVENT                Webhook_Event = 0
	Webhook_PROBLEM_CHANGED              Webhook_Event = 100 // generic problem properties changed (eolymp.atlas.ProblemChangedEvent)
	Webhook_PROBLEM_STATEMENT_CHANGED    Webhook_Event = 101 // problem statement changed (eolymp.atlas.StatementChangedEvent)
	Webhook_PROBLEM_SCORE_CHANGED        Webhook_Event = 102 // problem score for a member in a space changed (eolymp.atlas.ScoreChangedEvent)
	Webhook_SUBMISSION_CHANGED           Webhook_Event = 200 // submission properties changed (eolymp.atlas.SubmissionChangedEvent)
	Webhook_SUBMISSION_COMPLETED         Webhook_Event = 201 // submission testing is complete and final verdict is available (eolymp.atlas.SubmissionCompleteEvent)
	Webhook_MEMBER_CHANGED               Webhook_Event = 300 // member properties changed (eolymp.community.MemberChangedEvent)
	Webhook_GROUP_CHANGED                Webhook_Event = 350 // group properties changed (eolymp.community.GroupChangedEvent)
	Webhook_COURSE_MODULE_CHANGED        Webhook_Event = 400 // course module properties changed (eolymp.course.ModuleChangedEvent)
	Webhook_COURSE_MATERIAL_CHANGED      Webhook_Event = 401 // course material properties changed (eolymp.course.MaterialChangedEvent)
	Webhook_COURSE_STUDENT_CHANGED       Webhook_Event = 402 // course student properties changed (eolymp.course.StudentChangedEvent)
	Webhook_COURSE_ASSIGNMENT_CHANGED    Webhook_Event = 403 // course assignment properties changed (eolymp.course.AssignmentChangedEvent)
	Webhook_DISCUSSION_MESSAGE_CHANGED   Webhook_Event = 500 // discussion message (comments) properties changed (eolymp.discussion.MessageChangedEvent)
	Webhook_POST_CHANGED                 Webhook_Event = 600 // post properties changed (eolymp.discussion.PostChangedEvent)
	Webhook_POST_TRANSLATION_CHANGED     Webhook_Event = 601 // post translation properties changed(eolymp.discussion.PostTranslationChangedEvent)
	Webhook_POST_PUBLISHED               Webhook_Event = 602 // post published (eolymp.discussion.PostPublishedEvent)
	Webhook_TICKET_CHANGED               Webhook_Event = 700 // ticket properties changed (eolymp.judge.TicketChangedEvent)
	Webhook_TICKET_REPLY_CHANGED         Webhook_Event = 701 // ticket reply properties changed (eolymp.judge.ReplyChangedEvent)
	Webhook_CONTEST_SUBMISSION_COMPLETED Webhook_Event = 800 // contest submission testing is complete and final verdict is available (eolymp.judge.SubmissionCompletedEvent)
	Webhook_CONTEST_SCORE_CHANGED        Webhook_Event = 801 // contest score for a participant changed (eolymp.judge.ScoreChangedEvent)
	Webhook_CONTEST_PARTICIPANT_CHANGED  Webhook_Event = 802 // contest participant properties changed (not score) (eolymp.judge.ParticipantChangedEvent)
	Webhook_CONTEST_PARTICIPANT_JOINED   Webhook_Event = 803 // contest participant joined the contest (eolymp.judge.ParticipantJoinedEvent)
)

// Enum value maps for Webhook_Event.
var (
	Webhook_Event_name = map[int32]string{
		0:   "UNKNOWN_EVENT",
		100: "PROBLEM_CHANGED",
		101: "PROBLEM_STATEMENT_CHANGED",
		102: "PROBLEM_SCORE_CHANGED",
		200: "SUBMISSION_CHANGED",
		201: "SUBMISSION_COMPLETED",
		300: "MEMBER_CHANGED",
		350: "GROUP_CHANGED",
		400: "COURSE_MODULE_CHANGED",
		401: "COURSE_MATERIAL_CHANGED",
		402: "COURSE_STUDENT_CHANGED",
		403: "COURSE_ASSIGNMENT_CHANGED",
		500: "DISCUSSION_MESSAGE_CHANGED",
		600: "POST_CHANGED",
		601: "POST_TRANSLATION_CHANGED",
		602: "POST_PUBLISHED",
		700: "TICKET_CHANGED",
		701: "TICKET_REPLY_CHANGED",
		800: "CONTEST_SUBMISSION_COMPLETED",
		801: "CONTEST_SCORE_CHANGED",
		802: "CONTEST_PARTICIPANT_CHANGED",
		803: "CONTEST_PARTICIPANT_JOINED",
	}
	Webhook_Event_value = map[string]int32{
		"UNKNOWN_EVENT":                0,
		"PROBLEM_CHANGED":              100,
		"PROBLEM_STATEMENT_CHANGED":    101,
		"PROBLEM_SCORE_CHANGED":        102,
		"SUBMISSION_CHANGED":           200,
		"SUBMISSION_COMPLETED":         201,
		"MEMBER_CHANGED":               300,
		"GROUP_CHANGED":                350,
		"COURSE_MODULE_CHANGED":        400,
		"COURSE_MATERIAL_CHANGED":      401,
		"COURSE_STUDENT_CHANGED":       402,
		"COURSE_ASSIGNMENT_CHANGED":    403,
		"DISCUSSION_MESSAGE_CHANGED":   500,
		"POST_CHANGED":                 600,
		"POST_TRANSLATION_CHANGED":     601,
		"POST_PUBLISHED":               602,
		"TICKET_CHANGED":               700,
		"TICKET_REPLY_CHANGED":         701,
		"CONTEST_SUBMISSION_COMPLETED": 800,
		"CONTEST_SCORE_CHANGED":        801,
		"CONTEST_PARTICIPANT_CHANGED":  802,
		"CONTEST_PARTICIPANT_JOINED":   803,
	}
)

func (x Webhook_Event) Enum() *Webhook_Event {
	p := new(Webhook_Event)
	*p = x
	return p
}

func (x Webhook_Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Webhook_Event) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_webhook_webhook_proto_enumTypes[0].Descriptor()
}

func (Webhook_Event) Type() protoreflect.EnumType {
	return &file_eolymp_webhook_webhook_proto_enumTypes[0]
}

func (x Webhook_Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Webhook_Event.Descriptor instead.
func (Webhook_Event) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_proto_rawDescGZIP(), []int{0, 0}
}

type Webhook_Patch int32

const (
	Webhook_PATCH_UNKNOWN  Webhook_Patch = 0
	Webhook_PATCH_ALL      Webhook_Patch = 1
	Webhook_PATCH_NAME     Webhook_Patch = 2
	Webhook_PATCH_ENDPOINT Webhook_Patch = 3
	Webhook_PATCH_INACTIVE Webhook_Patch = 4
	Webhook_PATCH_EVENTS   Webhook_Patch = 5
)

// Enum value maps for Webhook_Patch.
var (
	Webhook_Patch_name = map[int32]string{
		0: "PATCH_UNKNOWN",
		1: "PATCH_ALL",
		2: "PATCH_NAME",
		3: "PATCH_ENDPOINT",
		4: "PATCH_INACTIVE",
		5: "PATCH_EVENTS",
	}
	Webhook_Patch_value = map[string]int32{
		"PATCH_UNKNOWN":  0,
		"PATCH_ALL":      1,
		"PATCH_NAME":     2,
		"PATCH_ENDPOINT": 3,
		"PATCH_INACTIVE": 4,
		"PATCH_EVENTS":   5,
	}
)

func (x Webhook_Patch) Enum() *Webhook_Patch {
	p := new(Webhook_Patch)
	*p = x
	return p
}

func (x Webhook_Patch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Webhook_Patch) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_webhook_webhook_proto_enumTypes[1].Descriptor()
}

func (Webhook_Patch) Type() protoreflect.EnumType {
	return &file_eolymp_webhook_webhook_proto_enumTypes[1]
}

func (x Webhook_Patch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Webhook_Patch.Descriptor instead.
func (Webhook_Patch) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_proto_rawDescGZIP(), []int{0, 1}
}

type Webhook struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Secret        string                 `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	Endpoint      string                 `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Inactive      bool                   `protobuf:"varint,5,opt,name=inactive,proto3" json:"inactive,omitempty"`
	Events        []Webhook_Event        `protobuf:"varint,10,rep,packed,name=events,proto3,enum=eolymp.webhook.Webhook_Event" json:"events,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,30,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastFailureAt *timestamppb.Timestamp `protobuf:"bytes,32,opt,name=last_failure_at,json=lastFailureAt,proto3" json:"last_failure_at,omitempty"` // time of the last failure
	LastSuccessAt *timestamppb.Timestamp `protobuf:"bytes,33,opt,name=last_success_at,json=lastSuccessAt,proto3" json:"last_success_at,omitempty"` // time of the last success
	DeliveryCount int32                  `protobuf:"varint,20,opt,name=delivery_count,json=deliveryCount,proto3" json:"delivery_count,omitempty"`  // number overall attempts to send something to the endpoint
	FailureCount  int32                  `protobuf:"varint,21,opt,name=failure_count,json=failureCount,proto3" json:"failure_count,omitempty"`     // number of failed attempts since last success
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Webhook) Reset() {
	*x = Webhook{}
	mi := &file_eolymp_webhook_webhook_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Webhook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Webhook) ProtoMessage() {}

func (x *Webhook) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_webhook_webhook_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Webhook.ProtoReflect.Descriptor instead.
func (*Webhook) Descriptor() ([]byte, []int) {
	return file_eolymp_webhook_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *Webhook) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Webhook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Webhook) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Webhook) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Webhook) GetInactive() bool {
	if x != nil {
		return x.Inactive
	}
	return false
}

func (x *Webhook) GetEvents() []Webhook_Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *Webhook) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Webhook) GetLastFailureAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastFailureAt
	}
	return nil
}

func (x *Webhook) GetLastSuccessAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSuccessAt
	}
	return nil
}

func (x *Webhook) GetDeliveryCount() int32 {
	if x != nil {
		return x.DeliveryCount
	}
	return 0
}

func (x *Webhook) GetFailureCount() int32 {
	if x != nil {
		return x.FailureCount
	}
	return 0
}

var File_eolymp_webhook_webhook_proto protoreflect.FileDescriptor

var file_eolymp_webhook_webhook_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x92, 0x09, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x35, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd7, 0x04, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x64, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x52, 0x4f, 0x42, 0x4c,
	0x45, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x44, 0x10, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45,
	0x4d, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10,
	0x66, 0x12, 0x17, 0x0a, 0x12, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0xc8, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x53, 0x55,
	0x42, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x44, 0x10, 0xc9, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0xac, 0x02, 0x12, 0x12, 0x0a, 0x0d, 0x47, 0x52,
	0x4f, 0x55, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0xde, 0x02, 0x12, 0x1a,
	0x0a, 0x15, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x90, 0x03, 0x12, 0x1c, 0x0a, 0x17, 0x43, 0x4f,
	0x55, 0x52, 0x53, 0x45, 0x5f, 0x4d, 0x41, 0x54, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x91, 0x03, 0x12, 0x1b, 0x0a, 0x16, 0x43, 0x4f, 0x55, 0x52,
	0x53, 0x45, 0x5f, 0x53, 0x54, 0x55, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x44, 0x10, 0x92, 0x03, 0x12, 0x1e, 0x0a, 0x19, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f,
	0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x44, 0x10, 0x93, 0x03, 0x12, 0x1f, 0x0a, 0x1a, 0x44, 0x49, 0x53, 0x43, 0x55, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x44, 0x10, 0xf4, 0x03, 0x12, 0x11, 0x0a, 0x0c, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0xd8, 0x04, 0x12, 0x1d, 0x0a, 0x18, 0x50, 0x4f, 0x53,
	0x54, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0xd9, 0x04, 0x12, 0x13, 0x0a, 0x0e, 0x50, 0x4f, 0x53, 0x54,
	0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0xda, 0x04, 0x12, 0x13, 0x0a,
	0x0e, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10,
	0xbc, 0x05, 0x12, 0x19, 0x0a, 0x14, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x50,
	0x4c, 0x59, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0xbd, 0x05, 0x12, 0x21, 0x0a,
	0x1c, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0xa0, 0x06,
	0x12, 0x1a, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x43, 0x4f, 0x52,
	0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0xa1, 0x06, 0x12, 0x20, 0x0a, 0x1b,
	0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x43, 0x49, 0x50,
	0x41, 0x4e, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0xa2, 0x06, 0x12, 0x1f,
	0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x43,
	0x49, 0x50, 0x41, 0x4e, 0x54, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x45, 0x44, 0x10, 0xa3, 0x06, 0x22,
	0x73, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x41, 0x54, 0x43,
	0x48, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x12,
	0x0a, 0x0e, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x53, 0x10, 0x05, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x3b,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eolymp_webhook_webhook_proto_rawDescOnce sync.Once
	file_eolymp_webhook_webhook_proto_rawDescData []byte
)

func file_eolymp_webhook_webhook_proto_rawDescGZIP() []byte {
	file_eolymp_webhook_webhook_proto_rawDescOnce.Do(func() {
		file_eolymp_webhook_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eolymp_webhook_webhook_proto_rawDesc), len(file_eolymp_webhook_webhook_proto_rawDesc)))
	})
	return file_eolymp_webhook_webhook_proto_rawDescData
}

var file_eolymp_webhook_webhook_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eolymp_webhook_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_webhook_webhook_proto_goTypes = []any{
	(Webhook_Event)(0),            // 0: eolymp.webhook.Webhook.Event
	(Webhook_Patch)(0),            // 1: eolymp.webhook.Webhook.Patch
	(*Webhook)(nil),               // 2: eolymp.webhook.Webhook
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_eolymp_webhook_webhook_proto_depIdxs = []int32{
	0, // 0: eolymp.webhook.Webhook.events:type_name -> eolymp.webhook.Webhook.Event
	3, // 1: eolymp.webhook.Webhook.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: eolymp.webhook.Webhook.last_failure_at:type_name -> google.protobuf.Timestamp
	3, // 3: eolymp.webhook.Webhook.last_success_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eolymp_webhook_webhook_proto_init() }
func file_eolymp_webhook_webhook_proto_init() {
	if File_eolymp_webhook_webhook_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eolymp_webhook_webhook_proto_rawDesc), len(file_eolymp_webhook_webhook_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_webhook_webhook_proto_goTypes,
		DependencyIndexes: file_eolymp_webhook_webhook_proto_depIdxs,
		EnumInfos:         file_eolymp_webhook_webhook_proto_enumTypes,
		MessageInfos:      file_eolymp_webhook_webhook_proto_msgTypes,
	}.Build()
	File_eolymp_webhook_webhook_proto = out.File
	file_eolymp_webhook_webhook_proto_goTypes = nil
	file_eolymp_webhook_webhook_proto_depIdxs = nil
}
