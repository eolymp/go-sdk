// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package judge

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	io "io"
	ioutil "io/ioutil"
	http "net/http"
	url "net/url"
	os "os"
)

type _JudgeHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type JudgeService struct {
	base string
	cli  _JudgeHttpClient
}

// NewJudgeHttpClient constructs client for Judge
func NewJudgeHttpClient(url string, cli _JudgeHttpClient) *JudgeService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &JudgeService{base: url, cli: cli}
}

func (s *JudgeService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
	var body io.Reader

	if in != nil {
		data, err := protojson.Marshal(in)
		if err != nil {
			return err
		}

		if verb != "GET" {
			body = bytes.NewReader(data)
		} else {
			query := url.Values{"q": []string{string(data)}}
			path = path + "?" + query.Encode()
		}
	}

	if in == nil && verb != "GET" {
		body = bytes.NewReader([]byte("{}"))
	}

	req, err := http.NewRequest(verb, s.base+path, body)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	req = req.WithContext(ctx)

	resp, err := s.cli.Do(req)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("non-200 response code (%v)", resp.StatusCode)
		}

		return fmt.Errorf("non-200 response code (%v): %s", resp.StatusCode, data)
	}

	if out != nil {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if err := protojson.Unmarshal(data, out); err != nil {
			return err
		}
	}

	return nil
}

func (s *JudgeService) CreateContest(ctx context.Context, in *CreateContestInput) (*CreateContestOutput, error) {
	out := &CreateContestOutput{}
	path := "/contests"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DeleteContest(ctx context.Context, in *DeleteContestInput) (*DeleteContestOutput, error) {
	out := &DeleteContestOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) UpdateContest(ctx context.Context, in *UpdateContestInput) (*UpdateContestOutput, error) {
	out := &UpdateContestOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DescribeContest(ctx context.Context, in *DescribeContestInput) (*DescribeContestOutput, error) {
	out := &DescribeContestOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ListContests(ctx context.Context, in *ListContestsInput) (*ListContestsOutput, error) {
	out := &ListContestsOutput{}
	path := "/contests"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) OpenContest(ctx context.Context, in *OpenContestInput) (*OpenContestOutput, error) {
	out := &OpenContestOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/open"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) CloseContest(ctx context.Context, in *CloseContestInput) (*CloseContestOutput, error) {
	out := &CloseContestOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/close"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) SuspendContest(ctx context.Context, in *SuspendContestInput) (*SuspendContestOutput, error) {
	out := &SuspendContestOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/suspend"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) FreezeContest(ctx context.Context, in *FreezeContestInput) (*FreezeContestOutput, error) {
	out := &FreezeContestOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/freeze"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ResumeContest(ctx context.Context, in *ResumeContestInput) (*ResumeContestOutput, error) {
	out := &ResumeContestOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/resume"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ImportProblem(ctx context.Context, in *ImportProblemInput) (*ImportProblemOutput, error) {
	out := &ImportProblemOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) SyncProblem(ctx context.Context, in *SyncProblemInput) (*SyncProblemOutput, error) {
	out := &SyncProblemOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems/" + url.PathEscape(in.GetProblemId()) + "/sync"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ProblemId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) UpdateProblem(ctx context.Context, in *UpdateProblemInput) (*UpdateProblemOutput, error) {
	out := &UpdateProblemOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems/" + url.PathEscape(in.GetProblemId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ProblemId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ListProblems(ctx context.Context, in *ListProblemsInput) (*ListProblemsOutput, error) {
	out := &ListProblemsOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DescribeProblem(ctx context.Context, in *DescribeProblemInput) (*DescribeProblemOutput, error) {
	out := &DescribeProblemOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems/" + url.PathEscape(in.GetProblemId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ProblemId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error) {
	out := &DescribeCodeTemplateOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems/" + url.PathEscape(in.GetProblemId()) + "/templates/" + url.PathEscape(in.GetTemplateId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ProblemId = ""
		in.TemplateId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput) (*LookupCodeTemplateOutput, error) {
	out := &LookupCodeTemplateOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems/" + url.PathEscape(in.GetProblemId()) + "/lookup-template"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ProblemId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ListStatements(ctx context.Context, in *ListStatementsInput) (*ListStatementsOutput, error) {
	out := &ListStatementsOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems/" + url.PathEscape(in.GetProblemId()) + "/statements"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ProblemId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ListAttachments(ctx context.Context, in *ListAttachmentsInput) (*ListAttachmentsOutput, error) {
	out := &ListAttachmentsOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems/" + url.PathEscape(in.GetProblemId()) + "/attachments"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ProblemId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ListExamples(ctx context.Context, in *ListExamplesInput) (*ListExamplesOutput, error) {
	out := &ListExamplesOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems/" + url.PathEscape(in.GetProblemId()) + "/examples"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ProblemId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DeleteProblem(ctx context.Context, in *DeleteProblemInput) (*DeleteProblemOutput, error) {
	out := &DeleteProblemOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems/" + url.PathEscape(in.GetProblemId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ProblemId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) RetestProblem(ctx context.Context, in *RetestProblemInput) (*RetestProblemOutput, error) {
	out := &RetestProblemOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems/" + url.PathEscape(in.GetProblemId()) + "/retest"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ProblemId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) AddParticipant(ctx context.Context, in *AddParticipantInput) (*AddParticipantOutput, error) {
	out := &AddParticipantOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) EnableParticipant(ctx context.Context, in *EnableParticipantInput) (*EnableParticipantOutput, error) {
	out := &EnableParticipantOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants/" + url.PathEscape(in.GetParticipantId()) + "/enable"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DisableParticipant(ctx context.Context, in *DisableParticipantInput) (*DisableParticipantOutput, error) {
	out := &DisableParticipantOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants/" + url.PathEscape(in.GetParticipantId()) + "/disable"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) UpdateParticipant(ctx context.Context, in *UpdateParticipantInput) (*UpdateParticipantOutput, error) {
	out := &UpdateParticipantOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants/" + url.PathEscape(in.GetParticipantId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) RemoveParticipant(ctx context.Context, in *RemoveParticipantInput) (*RemoveParticipantOutput, error) {
	out := &RemoveParticipantOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants/" + url.PathEscape(in.GetParticipantId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ListParticipants(ctx context.Context, in *ListParticipantsInput) (*ListParticipantsOutput, error) {
	out := &ListParticipantsOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DescribeParticipant(ctx context.Context, in *DescribeParticipantInput) (*DescribeParticipantOutput, error) {
	out := &DescribeParticipantOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants/" + url.PathEscape(in.GetParticipantId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) IntrospectParticipant(ctx context.Context, in *IntrospectParticipantInput) (*IntrospectParticipantOutput, error) {
	out := &IntrospectParticipantOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/introspect"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) JoinContest(ctx context.Context, in *JoinContestInput) (*JoinContestOutput, error) {
	out := &JoinContestOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/join"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) StartContest(ctx context.Context, in *StartContestInput) (*StartContestOutput, error) {
	out := &StartContestOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/start"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) VerifyPasscode(ctx context.Context, in *VerifyPasscodeInput) (*VerifyPasscodeOutput, error) {
	out := &VerifyPasscodeOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/verify-passcode"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) EnterPasscode(ctx context.Context, in *EnterPasscodeInput) (*EnterPasscodeOutput, error) {
	out := &EnterPasscodeOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/enter-passcode"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ResetPasscode(ctx context.Context, in *ResetPasscodeInput) (*ResetPasscodeOutput, error) {
	out := &ResetPasscodeOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants/" + url.PathEscape(in.GetParticipantId()) + "/passcode"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) SetPasscode(ctx context.Context, in *SetPasscodeInput) (*SetPasscodeOutput, error) {
	out := &SetPasscodeOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants/" + url.PathEscape(in.GetParticipantId()) + "/passcode"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) RemovePasscode(ctx context.Context, in *RemovePasscodeInput) (*RemovePasscodeOutput, error) {
	out := &RemovePasscodeOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants/" + url.PathEscape(in.GetParticipantId()) + "/passcode"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) CreateSubmission(ctx context.Context, in *CreateSubmissionInput) (*CreateSubmissionOutput, error) {
	out := &CreateSubmissionOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/problems/" + url.PathEscape(in.GetProblemId()) + "/submissions"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ProblemId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ListSubmissions(ctx context.Context, in *ListSubmissionsInput) (*ListSubmissionsOutput, error) {
	out := &ListSubmissionsOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/submissions"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DescribeSubmission(ctx context.Context, in *DescribeSubmissionInput) (*DescribeSubmissionOutput, error) {
	out := &DescribeSubmissionOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/submissions/" + url.PathEscape(in.GetSubmissionId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.SubmissionId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) RetestSubmission(ctx context.Context, in *RetestSubmissionInput) (*RetestSubmissionOutput, error) {
	out := &RetestSubmissionOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/submissions/" + url.PathEscape(in.GetSubmissionId()) + "/retest"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.SubmissionId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DeleteSubmission(ctx context.Context, in *DeleteSubmissionInput) (*DeleteSubmissionOutput, error) {
	out := &DeleteSubmissionOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/submissions/" + url.PathEscape(in.GetSubmissionId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.SubmissionId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) RestoreSubmission(ctx context.Context, in *RestoreSubmissionInput) (*RestoreSubmissionOutput, error) {
	out := &RestoreSubmissionOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/submissions/" + url.PathEscape(in.GetSubmissionId()) + "/restore"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.SubmissionId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) CreateAnnouncement(ctx context.Context, in *CreateAnnouncementInput) (*CreateAnnouncementOutput, error) {
	out := &CreateAnnouncementOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/announcements"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementInput) (*UpdateAnnouncementOutput, error) {
	out := &UpdateAnnouncementOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/announcements/" + url.PathEscape(in.GetAnnouncementId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.AnnouncementId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementInput) (*DeleteAnnouncementOutput, error) {
	out := &DeleteAnnouncementOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/announcements/" + url.PathEscape(in.GetAnnouncementId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.AnnouncementId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ReadAnnouncement(ctx context.Context, in *ReadAnnouncementInput) (*ReadAnnouncementOutput, error) {
	out := &ReadAnnouncementOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/announcements/" + url.PathEscape(in.GetAnnouncementId()) + "/read"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.AnnouncementId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DescribeAnnouncement(ctx context.Context, in *DescribeAnnouncementInput) (*DescribeAnnouncementOutput, error) {
	out := &DescribeAnnouncementOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/announcements/" + url.PathEscape(in.GetAnnouncementId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.AnnouncementId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DescribeAnnouncementStatus(ctx context.Context, in *DescribeAnnouncementStatusInput) (*DescribeAnnouncementStatusOutput, error) {
	out := &DescribeAnnouncementStatusOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/announcements/" + url.PathEscape(in.GetAnnouncementId()) + "/status"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.AnnouncementId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ListAnnouncements(ctx context.Context, in *ListAnnouncementsInput) (*ListAnnouncementsOutput, error) {
	out := &ListAnnouncementsOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/announcements"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) IntrospectScore(ctx context.Context, in *IntrospectScoreInput) (*IntrospectScoreOutput, error) {
	out := &IntrospectScoreOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/introspect/score"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DescribeScore(ctx context.Context, in *DescribeScoreInput) (*DescribeScoreOutput, error) {
	out := &DescribeScoreOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants/" + url.PathEscape(in.GetParticipantId()) + "/score"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ImportScore(ctx context.Context, in *ImportScoreInput) (*ImportScoreOutput, error) {
	out := &ImportScoreOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants/" + url.PathEscape(in.GetParticipantId()) + "/scores"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ExportScore(ctx context.Context, in *ExportScoreInput) (*ExportScoreOutput, error) {
	out := &ExportScoreOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/participants/" + url.PathEscape(in.GetParticipantId()) + "/scores"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ListResult(ctx context.Context, in *ListResultInput) (*ListResultOutput, error) {
	out := &ListResultOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/results"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) RebuildScore(ctx context.Context, in *RebuildScoreInput) (*RebuildScoreOutput, error) {
	out := &RebuildScoreOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/rebuild"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) ListActivities(ctx context.Context, in *ListActivitiesInput) (*ListActivitiesOutput, error) {
	out := &ListActivitiesOutput{}
	path := "/contests/" + url.PathEscape(in.GetContestId()) + "/activities"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ContestId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *JudgeService) DescribeContestUsage(ctx context.Context, in *DescribeContestUsageInput) (*DescribeContestUsageOutput, error) {
	out := &DescribeContestUsageOutput{}
	path := "/usage/contests"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
