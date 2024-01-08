// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package community

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

type _MemberServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type MemberServiceService struct {
	base string
	cli  _MemberServiceHttpClient
}

// NewMemberServiceHttpClient constructs client for MemberService
func NewMemberServiceHttpClient(url string, cli _MemberServiceHttpClient) *MemberServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &MemberServiceService{base: url, cli: cli}
}

func (s *MemberServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *MemberServiceService) CreateMember(ctx context.Context, in *CreateMemberInput) (*CreateMemberOutput, error) {
	out := &CreateMemberOutput{}
	path := "/members"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MemberServiceService) UpdateMember(ctx context.Context, in *UpdateMemberInput) (*UpdateMemberOutput, error) {
	out := &UpdateMemberOutput{}
	path := "/members/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MemberServiceService) DeleteMember(ctx context.Context, in *DeleteMemberInput) (*DeleteMemberOutput, error) {
	out := &DeleteMemberOutput{}
	path := "/members/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MemberServiceService) RestoreMember(ctx context.Context, in *RestoreMemberInput) (*RestoreMemberOutput, error) {
	out := &RestoreMemberOutput{}
	path := "/members/" + url.PathEscape(in.GetMemberId()) + "/restore"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MemberServiceService) DescribeMember(ctx context.Context, in *DescribeMemberInput) (*DescribeMemberOutput, error) {
	out := &DescribeMemberOutput{}
	path := "/members/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MemberServiceService) ListMembers(ctx context.Context, in *ListMembersInput) (*ListMembersOutput, error) {
	out := &ListMembersOutput{}
	path := "/members"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MemberServiceService) AssignMember(ctx context.Context, in *AssignMemberInput) (*AssignMemberOutput, error) {
	out := &AssignMemberOutput{}
	path := "/members/" + url.PathEscape(in.GetTeamId()) + "/users/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TeamId = ""
		in.MemberId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MemberServiceService) UnassignMember(ctx context.Context, in *UnassignMemberInput) (*UnassignMemberOutput, error) {
	out := &UnassignMemberOutput{}
	path := "/members/" + url.PathEscape(in.GetTeamId()) + "/users/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TeamId = ""
		in.MemberId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MemberServiceService) DescribeUsage(ctx context.Context, in *DescribeUsageInput) (*DescribeUsageOutput, error) {
	out := &DescribeUsageOutput{}
	path := "/usage/members"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
