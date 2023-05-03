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

type _TestingServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type TestingServiceService struct {
	base string
	cli  _TestingServiceHttpClient
}

// NewTestingServiceHttpClient constructs client for TestingService
func NewTestingServiceHttpClient(url string, cli _TestingServiceHttpClient) *TestingServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &TestingServiceService{base: url, cli: cli}
}

func (s *TestingServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *TestingServiceService) UpdateChecker(ctx context.Context, in *UpdateVerifierInput) (*UpdateVerifierOutput, error) {
	out := &UpdateVerifierOutput{}
	path := "/verifier"

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) DescribeChecker(ctx context.Context, in *DescribeVerifierInput) (*DescribeVerifierOutput, error) {
	out := &DescribeVerifierOutput{}
	path := "/verifier"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) UpdateInteractor(ctx context.Context, in *UpdateInteractorInput) (*UpdateInteractorOutput, error) {
	out := &UpdateInteractorOutput{}
	path := "/interactor"

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) DescribeInteractor(ctx context.Context, in *DescribeInteractorInput) (*DescribeInteractorOutput, error) {
	out := &DescribeInteractorOutput{}
	path := "/interactor"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) CreateTestset(ctx context.Context, in *CreateTestsetInput) (*CreateTestsetOutput, error) {
	out := &CreateTestsetOutput{}
	path := "/testsets"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) UpdateTestset(ctx context.Context, in *UpdateTestsetInput) (*UpdateTestsetOutput, error) {
	out := &UpdateTestsetOutput{}
	path := "/testsets/" + url.PathEscape(in.GetTestsetId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TestsetId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) DeleteTestset(ctx context.Context, in *DeleteTestsetInput) (*DeleteTestsetOutput, error) {
	out := &DeleteTestsetOutput{}
	path := "/testsets/" + url.PathEscape(in.GetTestsetId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TestsetId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) DescribeTestset(ctx context.Context, in *DescribeTestsetInput) (*DescribeTestsetOutput, error) {
	out := &DescribeTestsetOutput{}
	path := "/testsets/" + url.PathEscape(in.GetTestsetId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TestsetId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) ListTestsets(ctx context.Context, in *ListTestsetsInput) (*ListTestsetsOutput, error) {
	out := &ListTestsetsOutput{}
	path := "/testsets"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) CreateTest(ctx context.Context, in *CreateTestInput) (*CreateTestOutput, error) {
	out := &CreateTestOutput{}
	path := "/testsets/" + url.PathEscape(in.GetTestsetId()) + "/tests"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TestsetId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) UpdateTest(ctx context.Context, in *UpdateTestInput) (*UpdateTestOutput, error) {
	out := &UpdateTestOutput{}
	path := "/testsets/" + url.PathEscape(in.GetTestsetId()) + "/tests/" + url.PathEscape(in.GetTestId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TestsetId = ""
		in.TestId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) DeleteTest(ctx context.Context, in *DeleteTestInput) (*DeleteTestOutput, error) {
	out := &DeleteTestOutput{}
	path := "/testsets/" + url.PathEscape(in.GetTestsetId()) + "/tests/" + url.PathEscape(in.GetTestId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TestsetId = ""
		in.TestId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) DescribeTest(ctx context.Context, in *DescribeTestInput) (*DescribeTestOutput, error) {
	out := &DescribeTestOutput{}
	path := "/testsets/" + url.PathEscape(in.GetTestsetId()) + "/tests/" + url.PathEscape(in.GetTestId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TestsetId = ""
		in.TestId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) ListTests(ctx context.Context, in *ListTestsInput) (*ListTestsOutput, error) {
	out := &ListTestsOutput{}
	path := "/testsets/" + url.PathEscape(in.GetTestsetId()) + "/tests"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TestsetId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TestingServiceService) ListExamples(ctx context.Context, in *ListExamplesInput) (*ListExamplesOutput, error) {
	out := &ListExamplesOutput{}
	path := "/examples"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}