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
	os "os"
	strings "strings"
)

// NewCommunity constructs client for Community
func NewCommunity(cli CommunityHTTPClient) *CommunityService {
	base := "https://api.e-olymp.com"
	if v := os.Getenv("EOLYMP_API_URL"); v != "" {
		base = v
	}
	return &CommunityService{base: strings.TrimSuffix(base, "/"), cli: cli}
}

type CommunityHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type CommunityService struct {
	base string
	cli  CommunityHTTPClient
}

// invoke RPC method using twirp-like protocol
func (s *CommunityService) invoke(ctx context.Context, method string, in, out proto.Message) error {
	input, err := protojson.Marshal(in)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, s.base+"/twirp/"+method, bytes.NewReader(input))
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
		return fmt.Errorf("non-200 response code (%v)", resp.StatusCode)
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

func (s *CommunityService) AddMember(ctx context.Context, in *AddMemberInput) (*AddMemberOutput, error) {
	out := &AddMemberOutput{}

	if err := s.invoke(ctx, "eolymp.community.Community/AddMember", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) UpdateMember(ctx context.Context, in *UpdateMemberInput) (*UpdateMemberOutput, error) {
	out := &UpdateMemberOutput{}

	if err := s.invoke(ctx, "eolymp.community.Community/UpdateMember", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) RemoveMember(ctx context.Context, in *RemoveMemberInput) (*RemoveMemberOutput, error) {
	out := &RemoveMemberOutput{}

	if err := s.invoke(ctx, "eolymp.community.Community/RemoveMember", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) DescribeMember(ctx context.Context, in *DescribeMemberInput) (*DescribeMemberOutput, error) {
	out := &DescribeMemberOutput{}

	if err := s.invoke(ctx, "eolymp.community.Community/DescribeMember", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) IntrospectMember(ctx context.Context, in *IntrospectMemberInput) (*IntrospectMemberOutput, error) {
	out := &IntrospectMemberOutput{}

	if err := s.invoke(ctx, "eolymp.community.Community/IntrospectMember", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) ListMembers(ctx context.Context, in *ListMembersInput) (*ListMembersOutput, error) {
	out := &ListMembersOutput{}

	if err := s.invoke(ctx, "eolymp.community.Community/ListMembers", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) AddAttribute(ctx context.Context, in *AddAttributeInput) (*AddAttributeOutput, error) {
	out := &AddAttributeOutput{}

	if err := s.invoke(ctx, "eolymp.community.Community/AddAttribute", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) UpdateAttribute(ctx context.Context, in *UpdateAttributeInput) (*UpdateAttributeOutput, error) {
	out := &UpdateAttributeOutput{}

	if err := s.invoke(ctx, "eolymp.community.Community/UpdateAttribute", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) RemoveAttribute(ctx context.Context, in *RemoveAttributeInput) (*RemoveAttributeOutput, error) {
	out := &RemoveAttributeOutput{}

	if err := s.invoke(ctx, "eolymp.community.Community/RemoveAttribute", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) DescribeAttribute(ctx context.Context, in *DescribeAttributeInput) (*DescribeAttributeOutput, error) {
	out := &DescribeAttributeOutput{}

	if err := s.invoke(ctx, "eolymp.community.Community/DescribeAttribute", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CommunityService) ListAttributes(ctx context.Context, in *ListAttributesInput) (*ListAttributesOutput, error) {
	out := &ListAttributesOutput{}

	if err := s.invoke(ctx, "eolymp.community.Community/ListAttributes", in, out); err != nil {
		return nil, err
	}

	return out, nil
}
