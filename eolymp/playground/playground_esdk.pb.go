// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package playground

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

// NewPlayground constructs client for Playground
func NewPlaygroundService(url string, cli PlaygroundHTTPClient) *PlaygroundService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &PlaygroundService{base: url, cli: cli}
}

type PlaygroundHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type PlaygroundService struct {
	base string
	cli  PlaygroundHTTPClient
}

// invoke RPC method using twirp-like protocol
func (s *PlaygroundService) invoke(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *PlaygroundService) CreateRun(ctx context.Context, in *CreateRunInput) (*CreateRunOutput, error) {
	out := &CreateRunOutput{}
	path := "/playground/runs"

	if err := s.invoke(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *PlaygroundService) DescribeRun(ctx context.Context, in *DescribeRunInput) (*DescribeRunOutput, error) {
	out := &DescribeRunOutput{}
	path := "/playground/runs/" + url.PathEscape(in.GetRunId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.RunId = ""
	}

	if err := s.invoke(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
