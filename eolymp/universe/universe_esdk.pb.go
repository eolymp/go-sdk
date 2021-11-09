// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package universe

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

// NewUniverse constructs client for Universe
func NewUniverse(cli UniverseHTTPClient) *UniverseService {
	base := "https://api.e-olymp.com"
	if v := os.Getenv("EOLYMP_API_URL"); v != "" {
		base = v
	}
	return &UniverseService{base: strings.TrimSuffix(base, "/"), cli: cli}
}

type UniverseHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type UniverseService struct {
	base string
	cli  UniverseHTTPClient
}

// invoke RPC method using twirp-like protocol
func (s *UniverseService) invoke(ctx context.Context, method string, in, out proto.Message) error {
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

func (s *UniverseService) CreateSpace(ctx context.Context, in *CreateSpaceInput) (*CreateSpaceOutput, error) {
	out := &CreateSpaceOutput{}

	if err := s.invoke(ctx, "eolymp.universe.Universe/CreateSpace", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *UniverseService) UpdateSpace(ctx context.Context, in *UpdateSpaceInput) (*UpdateSpaceOutput, error) {
	out := &UpdateSpaceOutput{}

	if err := s.invoke(ctx, "eolymp.universe.Universe/UpdateSpace", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *UniverseService) DeleteSpace(ctx context.Context, in *DeleteSpaceInput) (*DeleteSpaceOutput, error) {
	out := &DeleteSpaceOutput{}

	if err := s.invoke(ctx, "eolymp.universe.Universe/DeleteSpace", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *UniverseService) LookupSpace(ctx context.Context, in *LookupSpaceInput) (*LookupSpaceOutput, error) {
	out := &LookupSpaceOutput{}

	if err := s.invoke(ctx, "eolymp.universe.Universe/LookupSpace", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *UniverseService) DescribeSpace(ctx context.Context, in *DescribeSpaceInput) (*DescribeSpaceOutput, error) {
	out := &DescribeSpaceOutput{}

	if err := s.invoke(ctx, "eolymp.universe.Universe/DescribeSpace", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *UniverseService) ListSpaces(ctx context.Context, in *ListSpacesInput) (*ListSpacesOutput, error) {
	out := &ListSpacesOutput{}

	if err := s.invoke(ctx, "eolymp.universe.Universe/ListSpaces", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *UniverseService) GrantPermission(ctx context.Context, in *GrantPermissionInput) (*GrantPermissionOutput, error) {
	out := &GrantPermissionOutput{}

	if err := s.invoke(ctx, "eolymp.universe.Universe/GrantPermission", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *UniverseService) RevokePermission(ctx context.Context, in *RevokePermissionInput) (*RevokePermissionOutput, error) {
	out := &RevokePermissionOutput{}

	if err := s.invoke(ctx, "eolymp.universe.Universe/RevokePermission", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *UniverseService) DescribePermission(ctx context.Context, in *DescribePermissionInput) (*DescribePermissionOutput, error) {
	out := &DescribePermissionOutput{}

	if err := s.invoke(ctx, "eolymp.universe.Universe/DescribePermission", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *UniverseService) IntrospectPermission(ctx context.Context, in *IntrospectPermissionInput) (*IntrospectPermissionOutput, error) {
	out := &IntrospectPermissionOutput{}

	if err := s.invoke(ctx, "eolymp.universe.Universe/IntrospectPermission", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *UniverseService) ListPermissions(ctx context.Context, in *ListPermissionsInput) (*ListPermissionsOutput, error) {
	out := &ListPermissionsOutput{}

	if err := s.invoke(ctx, "eolymp.universe.Universe/ListPermissions", in, out); err != nil {
		return nil, err
	}

	return out, nil
}
