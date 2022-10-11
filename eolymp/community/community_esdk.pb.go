// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package community

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
	url "net/url"
	os "os"
)

// NewCommunity constructs client for Community
func NewCommunityService(url string, cli CommunityHTTPClient) *CommunityService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &CommunityService{base: url, cli: cli}
}

type CommunityHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type CommunityService struct {
	base string
	cli  CommunityHTTPClient
}

// invoke RPC method using twirp-like protocol
func (s *CommunityService) invoke(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
	input := []byte("{}")

	if in != nil {
		input, err = protojson.Marshal(in)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(verb, s.base+path, bytes.NewReader(input))
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

func (s *CommunityService) JoinSpace(ctx context.Context, in *JoinSpaceInput) (*JoinSpaceOutput, error) {
	out := &JoinSpaceOutput{}
	path := "/members/_self"

	if err := s.invoke(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) LeaveSpace(ctx context.Context, in *LeaveSpaceInput) (*LeaveSpaceOutput, error) {
	out := &LeaveSpaceOutput{}
	path := "/members/_self"

	if err := s.invoke(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) RegisterMember(ctx context.Context, in *RegisterMemberInput) (*RegisterMemberOutput, error) {
	out := &RegisterMemberOutput{}
	path := "/members/_self/attributes"

	if err := s.invoke(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) IntrospectMember(ctx context.Context, in *IntrospectMemberInput) (*IntrospectMemberOutput, error) {
	out := &IntrospectMemberOutput{}
	path := "/members/_self"

	if err := s.invoke(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) AddMember(ctx context.Context, in *AddMemberInput) (*AddMemberOutput, error) {
	out := &AddMemberOutput{}
	path := "/members"

	if err := s.invoke(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) UpdateMember(ctx context.Context, in *UpdateMemberInput) (*UpdateMemberOutput, error) {
	out := &UpdateMemberOutput{}
	path := "/members/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
	}

	if err := s.invoke(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) RemoveMember(ctx context.Context, in *RemoveMemberInput) (*RemoveMemberOutput, error) {
	out := &RemoveMemberOutput{}
	path := "/members/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
	}

	if err := s.invoke(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) DescribeMember(ctx context.Context, in *DescribeMemberInput) (*DescribeMemberOutput, error) {
	out := &DescribeMemberOutput{}
	path := "/members/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
	}

	if err := s.invoke(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) ListMembers(ctx context.Context, in *ListMembersInput) (*ListMembersOutput, error) {
	out := &ListMembersOutput{}
	path := "/members"

	if err := s.invoke(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) AddAttribute(ctx context.Context, in *AddAttributeInput) (*AddAttributeOutput, error) {
	out := &AddAttributeOutput{}
	path := "/attributes"

	if err := s.invoke(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) UpdateAttribute(ctx context.Context, in *UpdateAttributeInput) (*UpdateAttributeOutput, error) {
	out := &UpdateAttributeOutput{}
	path := "/attributes/" + url.PathEscape(in.GetAttributeKey())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.AttributeKey = ""
	}

	if err := s.invoke(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) RemoveAttribute(ctx context.Context, in *RemoveAttributeInput) (*RemoveAttributeOutput, error) {
	out := &RemoveAttributeOutput{}
	path := "/attributes/" + url.PathEscape(in.GetAttributeKey())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.AttributeKey = ""
	}

	if err := s.invoke(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) DescribeAttribute(ctx context.Context, in *DescribeAttributeInput) (*DescribeAttributeOutput, error) {
	out := &DescribeAttributeOutput{}
	path := "/attributes/" + url.PathEscape(in.GetAttributeKey())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.AttributeKey = ""
	}

	if err := s.invoke(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) ListAttributes(ctx context.Context, in *ListAttributesInput) (*ListAttributesOutput, error) {
	out := &ListAttributesOutput{}
	path := "/attributes"

	if err := s.invoke(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
