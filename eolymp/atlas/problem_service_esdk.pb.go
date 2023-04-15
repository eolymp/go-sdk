// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package atlas

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

type _ProblemServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type ProblemServiceService struct {
	base string
	cli  _ProblemServiceHttpClient
}

// NewProblemServiceHttpClient constructs client for ProblemService
func NewProblemServiceHttpClient(url string, cli _ProblemServiceHttpClient) *ProblemServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &ProblemServiceService{base: url, cli: cli}
}

func (s *ProblemServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *ProblemServiceService) CreateProblem(ctx context.Context, in *CreateProblemInput) (*CreateProblemOutput, error) {
	out := &CreateProblemOutput{}
	path := "/problems"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ProblemServiceService) DeleteProblem(ctx context.Context, in *DeleteProblemInput) (*DeleteProblemOutput, error) {
	out := &DeleteProblemOutput{}
	path := "/problems/" + url.PathEscape(in.GetProblemId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ProblemId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ProblemServiceService) ListProblems(ctx context.Context, in *ListProblemsInput) (*ListProblemsOutput, error) {
	out := &ListProblemsOutput{}
	path := "/problems"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ProblemServiceService) DescribeProblem(ctx context.Context, in *DescribeProblemInput) (*DescribeProblemOutput, error) {
	out := &DescribeProblemOutput{}
	path := "/problems/" + url.PathEscape(in.GetProblemId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ProblemId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ProblemServiceService) UpdateVisibility(ctx context.Context, in *UpdateVisibilityInput) (*UpdateVisibilityOutput, error) {
	out := &UpdateVisibilityOutput{}
	path := "/problems/" + url.PathEscape(in.GetProblemId()) + "/visibility"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ProblemId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ProblemServiceService) UpdatePrivacy(ctx context.Context, in *UpdatePrivacyInput) (*UpdatePrivacyOutput, error) {
	out := &UpdatePrivacyOutput{}
	path := "/problems/" + url.PathEscape(in.GetProblemId()) + "/privacy"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ProblemId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ProblemServiceService) ListVersions(ctx context.Context, in *ListVersionsInput) (*ListVersionsOutput, error) {
	out := &ListVersionsOutput{}
	path := "/problems/" + url.PathEscape(in.GetProblemId()) + "/versions"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ProblemId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
