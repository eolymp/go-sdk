// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package worker

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

type _WorkerHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type WorkerService struct {
	base string
	cli  _WorkerHttpClient
}

// NewWorkerHttpClient constructs client for Worker
func NewWorkerHttpClient(url string, cli _WorkerHttpClient) *WorkerService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &WorkerService{base: url, cli: cli}
}

func (s *WorkerService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *WorkerService) CreateJob(ctx context.Context, in *CreateJobInput) (*CreateJobOutput, error) {
	out := &CreateJobOutput{}
	path := "/jobs"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkerService) DescribeJob(ctx context.Context, in *DescribeJobInput) (*DescribeJobOutput, error) {
	out := &DescribeJobOutput{}
	path := "/jobs/" + url.PathEscape(in.GetJobId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.JobId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkerService) ListJobs(ctx context.Context, in *ListJobsInput) (*ListJobsOutput, error) {
	out := &ListJobsOutput{}
	path := "/jobs"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

type _WorkerServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type WorkerServiceService struct {
	base string
	cli  _WorkerServiceHttpClient
}

// NewWorkerServiceHttpClient constructs client for WorkerService
func NewWorkerServiceHttpClient(url string, cli _WorkerServiceHttpClient) *WorkerServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &WorkerServiceService{base: url, cli: cli}
}

func (s *WorkerServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *WorkerServiceService) CreateJob(ctx context.Context, in *CreateJobInput) (*CreateJobOutput, error) {
	out := &CreateJobOutput{}
	path := "/jobs"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkerServiceService) DescribeJob(ctx context.Context, in *DescribeJobInput) (*DescribeJobOutput, error) {
	out := &DescribeJobOutput{}
	path := "/jobs/" + url.PathEscape(in.GetJobId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.JobId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkerServiceService) ListJobs(ctx context.Context, in *ListJobsInput) (*ListJobsOutput, error) {
	out := &ListJobsOutput{}
	path := "/jobs"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}