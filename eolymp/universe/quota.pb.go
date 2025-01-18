// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.24.4
// source: eolymp/universe/quota.proto

package universe

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

type Quota struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// general quota and features
	PermissionsPerSpace      uint32 `protobuf:"varint,6,opt,name=permissions_per_space,json=permissionsPerSpace,proto3" json:"permissions_per_space,omitempty"`                   // max number of admin users
	SingleSingOn             bool   `protobuf:"varint,22,opt,name=single_sing_on,json=singleSingOn,proto3" json:"single_sing_on,omitempty"`                                       // allow to configure sso for the space
	DedicatedUserDatabase    bool   `protobuf:"varint,23,opt,name=dedicated_user_database,json=dedicatedUserDatabase,proto3" json:"dedicated_user_database,omitempty"`            // allow to configure local user database for the space
	AttributesPerSpace       uint32 `protobuf:"varint,7,opt,name=attributes_per_space,json=attributesPerSpace,proto3" json:"attributes_per_space,omitempty"`                      // max number of custom profile fields for members
	CustomerSupportReplyTime uint32 `protobuf:"varint,24,opt,name=customer_support_reply_time,json=customerSupportReplyTime,proto3" json:"customer_support_reply_time,omitempty"` // customer support reply time in hours
	AllowDiscussions         bool   `protobuf:"varint,31,opt,name=allow_discussions,json=allowDiscussions,proto3" json:"allow_discussions,omitempty"`                             // enable discussion post and comment features
	AchievementsPerSpace     uint32 `protobuf:"varint,15,opt,name=achievements_per_space,json=achievementsPerSpace,proto3" json:"achievements_per_space,omitempty"`               // max number of achievements
	PrintersPerSpace         uint32 `protobuf:"varint,26,opt,name=printers_per_space,json=printersPerSpace,proto3" json:"printers_per_space,omitempty"`                           // max number of printers (0 - printers are disabled)
	// submission evaluation quota and features
	MonthlyEvaluationsBySeat uint32 `protobuf:"varint,13,opt,name=monthly_evaluations_by_seat,json=monthlyEvaluationsBySeat,proto3" json:"monthly_evaluations_by_seat,omitempty"` // number of submission evaluations per month/seat, including rejudge (to get total monthly quota multiply by number of seats)
	PriorityEvaluationQueue  bool   `protobuf:"varint,25,opt,name=priority_evaluation_queue,json=priorityEvaluationQueue,proto3" json:"priority_evaluation_queue,omitempty"`      // space uses priority testing queue
	PlagiarismAnalysis       bool   `protobuf:"varint,18,opt,name=plagiarism_analysis,json=plagiarismAnalysis,proto3" json:"plagiarism_analysis,omitempty"`                       // analyse submission code to see similarities and generate a report
	// member quota and features
	MembersPerSpace uint32 `protobuf:"varint,2,opt,name=members_per_space,json=membersPerSpace,proto3" json:"members_per_space,omitempty"` // total number of members
	// scoreboard quota and features
	ScoreboardsPerSpace uint32 `protobuf:"varint,5,opt,name=scoreboards_per_space,json=scoreboardsPerSpace,proto3" json:"scoreboards_per_space,omitempty"`
	// courses quota and features
	CoursesPerSpace uint32 `protobuf:"varint,12,opt,name=courses_per_space,json=coursesPerSpace,proto3" json:"courses_per_space,omitempty"`
	// problem quota and features
	ProblemsPerSpace         uint32 `protobuf:"varint,1,opt,name=problems_per_space,json=problemsPerSpace,proto3" json:"problems_per_space,omitempty"` // total number of problems in space
	TestsPerProblem          uint32 `protobuf:"varint,101,opt,name=tests_per_problem,json=testsPerProblem,proto3" json:"tests_per_problem,omitempty"`
	TestsetPerProblem        uint32 `protobuf:"varint,102,opt,name=testset_per_problem,json=testsetPerProblem,proto3" json:"testset_per_problem,omitempty"`
	StatementPerProblem      uint32 `protobuf:"varint,103,opt,name=statement_per_problem,json=statementPerProblem,proto3" json:"statement_per_problem,omitempty"`
	EditorialPerProblem      uint32 `protobuf:"varint,104,opt,name=editorial_per_problem,json=editorialPerProblem,proto3" json:"editorial_per_problem,omitempty"`
	SolutionsPerProblem      uint32 `protobuf:"varint,105,opt,name=solutions_per_problem,json=solutionsPerProblem,proto3" json:"solutions_per_problem,omitempty"`
	CodeTemplatesPerProblem  uint32 `protobuf:"varint,106,opt,name=code_templates_per_problem,json=codeTemplatesPerProblem,proto3" json:"code_templates_per_problem,omitempty"`
	DebugAssistant           bool   `protobuf:"varint,120,opt,name=debug_assistant,json=debugAssistant,proto3" json:"debug_assistant,omitempty"`
	DebugHintsDailyPerAdmin  uint32 `protobuf:"varint,121,opt,name=debug_hints_daily_per_admin,json=debugHintsDailyPerAdmin,proto3" json:"debug_hints_daily_per_admin,omitempty"`
	DebugHintsDailyPerMember uint32 `protobuf:"varint,122,opt,name=debug_hints_daily_per_member,json=debugHintsDailyPerMember,proto3" json:"debug_hints_daily_per_member,omitempty"`
	// contest quota and features
	ContestsPerSpace        uint32 `protobuf:"varint,3,opt,name=contests_per_space,json=contestsPerSpace,proto3" json:"contests_per_space,omitempty"`                         // total number of contests
	ActiveContestsPerSpace  uint32 `protobuf:"varint,4,opt,name=active_contests_per_space,json=activeContestsPerSpace,proto3" json:"active_contests_per_space,omitempty"`     // max number of simultaneously active contests (ongoing and upsolve)
	MonthlyContestsPerSpace uint32 `protobuf:"varint,14,opt,name=monthly_contests_per_space,json=monthlyContestsPerSpace,proto3" json:"monthly_contests_per_space,omitempty"` // max number of contests created (started) during current billing period
	ProblemsPerContest      uint32 `protobuf:"varint,10,opt,name=problems_per_contest,json=problemsPerContest,proto3" json:"problems_per_contest,omitempty"`
	ParticipantsPerContest  uint32 `protobuf:"varint,11,opt,name=participants_per_contest,json=participantsPerContest,proto3" json:"participants_per_contest,omitempty"` // deprecated: limited by number of members in space
	ContestUpsolveMode      bool   `protobuf:"varint,16,opt,name=contest_upsolve_mode,json=contestUpsolveMode,proto3" json:"contest_upsolve_mode,omitempty"`             // allow contest upsolve mode
	MaxContestDuration      uint32 `protobuf:"varint,17,opt,name=max_contest_duration,json=maxContestDuration,proto3" json:"max_contest_duration,omitempty"`             // max contest duration in seconds
	TeamContests            bool   `protobuf:"varint,19,opt,name=team_contests,json=teamContests,proto3" json:"team_contests,omitempty"`                                 // analyse submission code to see similarities and generate a report
	GhostParticipants       bool   `protobuf:"varint,20,opt,name=ghost_participants,json=ghostParticipants,proto3" json:"ghost_participants,omitempty"`                  // analyse submission code to see similarities and generate a report
	UnofficialParticipants  bool   `protobuf:"varint,21,opt,name=unofficial_participants,json=unofficialParticipants,proto3" json:"unofficial_participants,omitempty"`   // analyse submission code to see similarities and generate a report
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *Quota) Reset() {
	*x = Quota{}
	mi := &file_eolymp_universe_quota_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Quota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quota) ProtoMessage() {}

func (x *Quota) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_universe_quota_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quota.ProtoReflect.Descriptor instead.
func (*Quota) Descriptor() ([]byte, []int) {
	return file_eolymp_universe_quota_proto_rawDescGZIP(), []int{0}
}

func (x *Quota) GetPermissionsPerSpace() uint32 {
	if x != nil {
		return x.PermissionsPerSpace
	}
	return 0
}

func (x *Quota) GetSingleSingOn() bool {
	if x != nil {
		return x.SingleSingOn
	}
	return false
}

func (x *Quota) GetDedicatedUserDatabase() bool {
	if x != nil {
		return x.DedicatedUserDatabase
	}
	return false
}

func (x *Quota) GetAttributesPerSpace() uint32 {
	if x != nil {
		return x.AttributesPerSpace
	}
	return 0
}

func (x *Quota) GetCustomerSupportReplyTime() uint32 {
	if x != nil {
		return x.CustomerSupportReplyTime
	}
	return 0
}

func (x *Quota) GetAllowDiscussions() bool {
	if x != nil {
		return x.AllowDiscussions
	}
	return false
}

func (x *Quota) GetAchievementsPerSpace() uint32 {
	if x != nil {
		return x.AchievementsPerSpace
	}
	return 0
}

func (x *Quota) GetPrintersPerSpace() uint32 {
	if x != nil {
		return x.PrintersPerSpace
	}
	return 0
}

func (x *Quota) GetMonthlyEvaluationsBySeat() uint32 {
	if x != nil {
		return x.MonthlyEvaluationsBySeat
	}
	return 0
}

func (x *Quota) GetPriorityEvaluationQueue() bool {
	if x != nil {
		return x.PriorityEvaluationQueue
	}
	return false
}

func (x *Quota) GetPlagiarismAnalysis() bool {
	if x != nil {
		return x.PlagiarismAnalysis
	}
	return false
}

func (x *Quota) GetMembersPerSpace() uint32 {
	if x != nil {
		return x.MembersPerSpace
	}
	return 0
}

func (x *Quota) GetScoreboardsPerSpace() uint32 {
	if x != nil {
		return x.ScoreboardsPerSpace
	}
	return 0
}

func (x *Quota) GetCoursesPerSpace() uint32 {
	if x != nil {
		return x.CoursesPerSpace
	}
	return 0
}

func (x *Quota) GetProblemsPerSpace() uint32 {
	if x != nil {
		return x.ProblemsPerSpace
	}
	return 0
}

func (x *Quota) GetTestsPerProblem() uint32 {
	if x != nil {
		return x.TestsPerProblem
	}
	return 0
}

func (x *Quota) GetTestsetPerProblem() uint32 {
	if x != nil {
		return x.TestsetPerProblem
	}
	return 0
}

func (x *Quota) GetStatementPerProblem() uint32 {
	if x != nil {
		return x.StatementPerProblem
	}
	return 0
}

func (x *Quota) GetEditorialPerProblem() uint32 {
	if x != nil {
		return x.EditorialPerProblem
	}
	return 0
}

func (x *Quota) GetSolutionsPerProblem() uint32 {
	if x != nil {
		return x.SolutionsPerProblem
	}
	return 0
}

func (x *Quota) GetCodeTemplatesPerProblem() uint32 {
	if x != nil {
		return x.CodeTemplatesPerProblem
	}
	return 0
}

func (x *Quota) GetDebugAssistant() bool {
	if x != nil {
		return x.DebugAssistant
	}
	return false
}

func (x *Quota) GetDebugHintsDailyPerAdmin() uint32 {
	if x != nil {
		return x.DebugHintsDailyPerAdmin
	}
	return 0
}

func (x *Quota) GetDebugHintsDailyPerMember() uint32 {
	if x != nil {
		return x.DebugHintsDailyPerMember
	}
	return 0
}

func (x *Quota) GetContestsPerSpace() uint32 {
	if x != nil {
		return x.ContestsPerSpace
	}
	return 0
}

func (x *Quota) GetActiveContestsPerSpace() uint32 {
	if x != nil {
		return x.ActiveContestsPerSpace
	}
	return 0
}

func (x *Quota) GetMonthlyContestsPerSpace() uint32 {
	if x != nil {
		return x.MonthlyContestsPerSpace
	}
	return 0
}

func (x *Quota) GetProblemsPerContest() uint32 {
	if x != nil {
		return x.ProblemsPerContest
	}
	return 0
}

func (x *Quota) GetParticipantsPerContest() uint32 {
	if x != nil {
		return x.ParticipantsPerContest
	}
	return 0
}

func (x *Quota) GetContestUpsolveMode() bool {
	if x != nil {
		return x.ContestUpsolveMode
	}
	return false
}

func (x *Quota) GetMaxContestDuration() uint32 {
	if x != nil {
		return x.MaxContestDuration
	}
	return 0
}

func (x *Quota) GetTeamContests() bool {
	if x != nil {
		return x.TeamContests
	}
	return false
}

func (x *Quota) GetGhostParticipants() bool {
	if x != nil {
		return x.GhostParticipants
	}
	return false
}

func (x *Quota) GetUnofficialParticipants() bool {
	if x != nil {
		return x.UnofficialParticipants
	}
	return false
}

var File_eolymp_universe_quota_proto protoreflect.FileDescriptor

var file_eolymp_universe_quota_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x22, 0xe0,
	0x0d, 0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x50, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0e,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x6e, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x6e, 0x67,
	0x4f, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x64, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x15, 0x64, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x1b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x18, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x44, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x50, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x1b,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x18, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x53, 0x65, 0x61, 0x74, 0x12, 0x3a, 0x0a, 0x19, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x6c, 0x61, 0x67, 0x69,
	0x61, 0x72, 0x69, 0x73, 0x6d, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x70, 0x6c, 0x61, 0x67, 0x69, 0x61, 0x72, 0x69, 0x73, 0x6d,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x50, 0x65, 0x72, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x13, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x50, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x10, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x50, 0x65, 0x72, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x2e,
	0x0a, 0x13, 0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x65, 0x74, 0x50, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x32,
	0x0a, 0x15, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x18, 0x68, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x13, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x18,
	0x69, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x50, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x3b, 0x0a, 0x1a, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17,
	0x63, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x5f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x78, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x12, 0x3c, 0x0a, 0x1b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x73, 0x5f,
	0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x79, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x64, 0x65, 0x62, 0x75, 0x67, 0x48, 0x69, 0x6e, 0x74,
	0x73, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x3e,
	0x0a, 0x1c, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x73, 0x5f, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x7a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x64, 0x65, 0x62, 0x75, 0x67, 0x48, 0x69, 0x6e, 0x74, 0x73,
	0x44, 0x61, 0x69, 0x6c, 0x79, 0x50, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x19,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x16, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x50,
	0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x50, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x55, 0x70, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x12, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x74, 0x65, 0x61,
	0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x67, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x17, 0x75, 0x6e, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x75, 0x6e, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f,
	0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x3b, 0x75, 0x6e,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_universe_quota_proto_rawDescOnce sync.Once
	file_eolymp_universe_quota_proto_rawDescData = file_eolymp_universe_quota_proto_rawDesc
)

func file_eolymp_universe_quota_proto_rawDescGZIP() []byte {
	file_eolymp_universe_quota_proto_rawDescOnce.Do(func() {
		file_eolymp_universe_quota_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_universe_quota_proto_rawDescData)
	})
	return file_eolymp_universe_quota_proto_rawDescData
}

var file_eolymp_universe_quota_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eolymp_universe_quota_proto_goTypes = []any{
	(*Quota)(nil), // 0: eolymp.universe.Quota
}
var file_eolymp_universe_quota_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eolymp_universe_quota_proto_init() }
func file_eolymp_universe_quota_proto_init() {
	if File_eolymp_universe_quota_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_universe_quota_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_universe_quota_proto_goTypes,
		DependencyIndexes: file_eolymp_universe_quota_proto_depIdxs,
		MessageInfos:      file_eolymp_universe_quota_proto_msgTypes,
	}.Build()
	File_eolymp_universe_quota_proto = out.File
	file_eolymp_universe_quota_proto_rawDesc = nil
	file_eolymp_universe_quota_proto_goTypes = nil
	file_eolymp_universe_quota_proto_depIdxs = nil
}
