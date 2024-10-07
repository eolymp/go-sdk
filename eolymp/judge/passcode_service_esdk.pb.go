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

type _PasscodeServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type PasscodeServiceService struct {
	base string
	cli  _PasscodeServiceHttpClient
}

// NewPasscodeServiceHttpClient constructs client for PasscodeService
func NewPasscodeServiceHttpClient(url string, cli _PasscodeServiceHttpClient) *PasscodeServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &PasscodeServiceService{base: url, cli: cli}
}

func (s *PasscodeServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *PasscodeServiceService) VerifyPasscode(ctx context.Context, in *VerifyPasscodeInput) (*VerifyPasscodeOutput, error) {
	out := &VerifyPasscodeOutput{}
	path := "/verify-passcode"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *PasscodeServiceService) EnterPasscode(ctx context.Context, in *EnterPasscodeInput) (*EnterPasscodeOutput, error) {
	out := &EnterPasscodeOutput{}
	path := "/enter-passcode"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *PasscodeServiceService) ResetPasscode(ctx context.Context, in *ResetPasscodeInput) (*ResetPasscodeOutput, error) {
	out := &ResetPasscodeOutput{}
	path := "/participants/" + url.PathEscape(in.GetParticipantId()) + "/passcode"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *PasscodeServiceService) SetPasscode(ctx context.Context, in *SetPasscodeInput) (*SetPasscodeOutput, error) {
	out := &SetPasscodeOutput{}
	path := "/participants/" + url.PathEscape(in.GetParticipantId()) + "/passcode"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *PasscodeServiceService) RemovePasscode(ctx context.Context, in *RemovePasscodeInput) (*RemovePasscodeOutput, error) {
	out := &RemovePasscodeOutput{}
	path := "/participants/" + url.PathEscape(in.GetParticipantId()) + "/passcode"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ParticipantId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
