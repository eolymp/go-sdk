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

type _SolutionServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type SolutionServiceService struct {
	base string
	cli  _SolutionServiceHttpClient
}

// NewSolutionServiceHttpClient constructs client for SolutionService
func NewSolutionServiceHttpClient(url string, cli _SolutionServiceHttpClient) *SolutionServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &SolutionServiceService{base: url, cli: cli}
}

func (s *SolutionServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *SolutionServiceService) CreateSolution(ctx context.Context, in *CreateSolutionInput) (*CreateSolutionOutput, error) {
	out := &CreateSolutionOutput{}
	path := "/solutions"

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SolutionServiceService) UpdateSolution(ctx context.Context, in *UpdateSolutionInput) (*UpdateSolutionOutput, error) {
	out := &UpdateSolutionOutput{}
	path := "/solutions/" + url.PathEscape(in.GetSolutionId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.SolutionId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SolutionServiceService) DeleteSolution(ctx context.Context, in *DeleteSolutionInput) (*DeleteSolutionOutput, error) {
	out := &DeleteSolutionOutput{}
	path := "/solutions/" + url.PathEscape(in.GetSolutionId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.SolutionId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SolutionServiceService) DescribeSolution(ctx context.Context, in *DescribeSolutionInput) (*DescribeSolutionOutput, error) {
	out := &DescribeSolutionOutput{}
	path := "/solutions/" + url.PathEscape(in.GetSolutionId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.SolutionId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SolutionServiceService) ListSolutions(ctx context.Context, in *ListSolutionsInput) (*ListSolutionsOutput, error) {
	out := &ListSolutionsOutput{}
	path := "/solutions"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SolutionServiceService) CheckSolutions(ctx context.Context, in *CheckSolutionsInput) (*CheckSolutionsOutput, error) {
	out := &CheckSolutionsOutput{}
	path := "/solutions:check"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
