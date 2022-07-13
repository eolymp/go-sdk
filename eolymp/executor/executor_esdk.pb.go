// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package executor

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

// NewExecutor constructs client for Executor
func NewExecutor(cli ExecutorHTTPClient) *ExecutorService {
	base := "https://api.eolymp.com"
	if v := os.Getenv("EOLYMP_API_URL"); v != "" {
		base = v
	}
	return &ExecutorService{base: strings.TrimSuffix(base, "/"), cli: cli}
}

type ExecutorHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type ExecutorService struct {
	base string
	cli  ExecutorHTTPClient
}

// invoke RPC method using twirp-like protocol
func (s *ExecutorService) invoke(ctx context.Context, method string, in, out proto.Message) (err error) {
	input := []byte("{}")

	if in != nil {
		input, err = protojson.Marshal(in)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(http.MethodPost, s.base+"/"+method, bytes.NewReader(input))
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

func (s *ExecutorService) DescribeLanguage(ctx context.Context, in *DescribeLanguageInput) (*DescribeLanguageOutput, error) {
	out := &DescribeLanguageOutput{}

	if err := s.invoke(ctx, "eolymp.executor.Executor/DescribeLanguage", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ExecutorService) ListLanguages(ctx context.Context, in *ListLanguagesInput) (*ListLanguagesOutput, error) {
	out := &ListLanguagesOutput{}

	if err := s.invoke(ctx, "eolymp.executor.Executor/ListLanguages", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ExecutorService) DescribeRuntime(ctx context.Context, in *DescribeRuntimeInput) (*DescribeRuntimeOutput, error) {
	out := &DescribeRuntimeOutput{}

	if err := s.invoke(ctx, "eolymp.executor.Executor/DescribeRuntime", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ExecutorService) ListRuntime(ctx context.Context, in *ListRuntimeInput) (*ListRuntimeOutput, error) {
	out := &ListRuntimeOutput{}

	if err := s.invoke(ctx, "eolymp.executor.Executor/ListRuntime", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *ExecutorService) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error) {
	out := &DescribeCodeTemplateOutput{}

	if err := s.invoke(ctx, "eolymp.executor.Executor/DescribeCodeTemplate", in, out); err != nil {
		return nil, err
	}

	return out, nil
}
