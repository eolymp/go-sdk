// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package cognito

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

type _CognitoHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type CognitoService struct {
	base string
	cli  _CognitoHttpClient
}

// NewCognitoHttpClient constructs client for Cognito
func NewCognitoHttpClient(url string, cli _CognitoHttpClient) *CognitoService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &CognitoService{base: url, cli: cli}
}

func (s *CognitoService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *CognitoService) Signout(ctx context.Context, in *SignoutInput) (*SignoutOutput, error) {
	out := &SignoutOutput{}
	path := "/self/signout"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) CreateAccessKey(ctx context.Context, in *CreateAccessKeyInput) (*CreateAccessKeyOutput, error) {
	out := &CreateAccessKeyOutput{}
	path := "/access-keys"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) DeleteAccessKey(ctx context.Context, in *DeleteAccessKeyInput) (*DeleteAccessKeyOutput, error) {
	out := &DeleteAccessKeyOutput{}
	path := "/access-keys/" + url.PathEscape(in.GetKeyId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.KeyId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) ListAccessKeys(ctx context.Context, in *ListAccessKeysInput) (*ListAccessKeysOutput, error) {
	out := &ListAccessKeysOutput{}
	path := "/access-keys"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) CreateUser(ctx context.Context, in *CreateUserInput) (*CreateUserOutput, error) {
	out := &CreateUserOutput{}
	path := "/users"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) VerifyEmail(ctx context.Context, in *VerifyEmailInput) (*VerifyEmailOutput, error) {
	out := &VerifyEmailOutput{}
	path := "/users/" + url.PathEscape(in.GetUserId()) + "/verify"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.UserId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) ResendEmailVerification(ctx context.Context, in *ResendEmailVerificationInput) (*ResendEmailVerificationOutput, error) {
	out := &ResendEmailVerificationOutput{}
	path := "/self/email/resend-verification"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) UpdateEmail(ctx context.Context, in *UpdateEmailInput) (*UpdateEmailOutput, error) {
	out := &UpdateEmailOutput{}
	path := "/self/email"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) UpdateProfile(ctx context.Context, in *UpdateProfileInput) (*UpdateProfileOutput, error) {
	out := &UpdateProfileOutput{}
	path := "/self"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) UpdatePicture(ctx context.Context, in *UpdatePictureInput) (*UpdatePictureOutput, error) {
	out := &UpdatePictureOutput{}
	path := "/self/picture"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) UpdatePassword(ctx context.Context, in *UpdatePasswordInput) (*UpdatePasswordOutput, error) {
	out := &UpdatePasswordOutput{}
	path := "/self/password"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) IntrospectUser(ctx context.Context, in *IntrospectUserInput) (*IntrospectUserOutput, error) {
	out := &IntrospectUserOutput{}
	path := "/self"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) DescribeUser(ctx context.Context, in *DescribeUserInput) (*DescribeUserOutput, error) {
	out := &DescribeUserOutput{}
	path := "/users/" + url.PathEscape(in.GetUserId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.UserId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) ListUsers(ctx context.Context, in *ListUsersInput) (*ListUsersOutput, error) {
	out := &ListUsersOutput{}
	path := "/users"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) IntrospectQuota(ctx context.Context, in *IntrospectQuotaInput) (*IntrospectQuotaOutput, error) {
	out := &IntrospectQuotaOutput{}
	path := "/self/quota"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) IntrospectRoles(ctx context.Context, in *IntrospectRolesInput) (*IntrospectRolesOutput, error) {
	out := &IntrospectRolesOutput{}
	path := "/self/roles"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) ListRoles(ctx context.Context, in *ListRolesInput) (*ListRolesOutput, error) {
	out := &ListRolesOutput{}
	path := "/users/" + url.PathEscape(in.GetUserId()) + "/roles"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.UserId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) UpdateRoles(ctx context.Context, in *UpdateRolesInput) (*UpdateRolesOutput, error) {
	out := &UpdateRolesOutput{}
	path := "/users/" + url.PathEscape(in.GetUserId()) + "/roles"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.UserId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) StartRecovery(ctx context.Context, in *StartRecoveryInput) (*StartRecoveryOutput, error) {
	out := &StartRecoveryOutput{}
	path := "/self/recovery"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) CompleteRecovery(ctx context.Context, in *CompleteRecoverInput) (*CompleteRecoverOutput, error) {
	out := &CompleteRecoverOutput{}
	path := "/users/" + url.PathEscape(in.GetUserId()) + "/recover"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.UserId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) ListEntitlements(ctx context.Context, in *ListEntitlementsInput) (*ListEntitlementsOutput, error) {
	out := &ListEntitlementsOutput{}
	path := "/__cognito/entitlements"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *CognitoService) SelfDestruct(ctx context.Context, in *SelfDestructInput) (*SelfDestructOutput, error) {
	out := &SelfDestructOutput{}
	path := "/self"

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
