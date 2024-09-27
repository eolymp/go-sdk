// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package course

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

type _ModuleServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type ModuleServiceService struct {
	base string
	cli  _ModuleServiceHttpClient
}

// NewModuleServiceHttpClient constructs client for ModuleService
func NewModuleServiceHttpClient(url string, cli _ModuleServiceHttpClient) *ModuleServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &ModuleServiceService{base: url, cli: cli}
}

func (s *ModuleServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *ModuleServiceService) CreateModule(ctx context.Context, in *CreateModuleInput) (*CreateModuleOutput, error) {
	out := &CreateModuleOutput{}
	path := "/modules"

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ModuleServiceService) UpdateModule(ctx context.Context, in *UpdateModuleInput) (*UpdateModuleOutput, error) {
	out := &UpdateModuleOutput{}
	path := "/modules/" + url.PathEscape(in.GetModuleId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ModuleId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ModuleServiceService) DeleteModule(ctx context.Context, in *DeleteModuleInput) (*DeleteModuleOutput, error) {
	out := &DeleteModuleOutput{}
	path := "/modules/" + url.PathEscape(in.GetModuleId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ModuleId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ModuleServiceService) DescribeModule(ctx context.Context, in *DescribeModuleInput) (*DescribeModuleOutput, error) {
	out := &DescribeModuleOutput{}
	path := "/modules/" + url.PathEscape(in.GetModuleId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ModuleId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ModuleServiceService) ListModules(ctx context.Context, in *ListModulesInput) (*ListModulesOutput, error) {
	out := &ListModulesOutput{}
	path := "/modules"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ModuleServiceService) StartModule(ctx context.Context, in *StartModuleInput) (*StartModuleOutput, error) {
	out := &StartModuleOutput{}
	path := "/modules/" + url.PathEscape(in.GetModuleId()) + "/start"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ModuleId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ModuleServiceService) AssignModule(ctx context.Context, in *AssignModuleInput) (*AssignModuleOutput, error) {
	out := &AssignModuleOutput{}
	path := "/modules/" + url.PathEscape(in.GetModuleId()) + "/assignments"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ModuleId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ModuleServiceService) UnassignModule(ctx context.Context, in *UnassignModuleInput) (*UnassignModuleOutput, error) {
	out := &UnassignModuleOutput{}
	path := "/modules/" + url.PathEscape(in.GetModuleId()) + "/assignments"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ModuleId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
