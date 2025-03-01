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

type _StatementServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type StatementServiceService struct {
	base string
	cli  _StatementServiceHttpClient
}

// NewStatementServiceHttpClient constructs client for StatementService
func NewStatementServiceHttpClient(url string, cli _StatementServiceHttpClient) *StatementServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &StatementServiceService{base: url, cli: cli}
}

func (s *StatementServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *StatementServiceService) CreateStatement(ctx context.Context, in *CreateStatementInput) (*CreateStatementOutput, error) {
	out := &CreateStatementOutput{}
	path := "/statements"

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StatementServiceService) UpdateStatement(ctx context.Context, in *UpdateStatementInput) (*UpdateStatementOutput, error) {
	out := &UpdateStatementOutput{}
	path := "/statements/" + url.PathEscape(in.GetStatementId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.StatementId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StatementServiceService) DeleteStatement(ctx context.Context, in *DeleteStatementInput) (*DeleteStatementOutput, error) {
	out := &DeleteStatementOutput{}
	path := "/statements/" + url.PathEscape(in.GetStatementId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.StatementId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StatementServiceService) DescribeStatement(ctx context.Context, in *DescribeStatementInput) (*DescribeStatementOutput, error) {
	out := &DescribeStatementOutput{}
	path := "/statements/" + url.PathEscape(in.GetStatementId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.StatementId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StatementServiceService) LookupStatement(ctx context.Context, in *LookupStatementInput) (*LookupStatementOutput, error) {
	out := &LookupStatementOutput{}
	path := "/translate"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StatementServiceService) PreviewStatement(ctx context.Context, in *PreviewStatementInput) (*PreviewStatementOutput, error) {
	out := &PreviewStatementOutput{}
	path := "/renders"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StatementServiceService) ListStatements(ctx context.Context, in *ListStatementsInput) (*ListStatementsOutput, error) {
	out := &ListStatementsOutput{}
	path := "/statements"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StatementServiceService) TranslateStatements(ctx context.Context, in *TranslateStatementsInput) (*TranslateStatementsOutput, error) {
	out := &TranslateStatementsOutput{}
	path := "/statements:translate"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StatementServiceService) ExportStatement(ctx context.Context, in *ExportStatementInput) (*ExportStatementOutput, error) {
	out := &ExportStatementOutput{}
	path := "/statements/" + url.PathEscape(in.GetStatementId()) + "/export"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.StatementId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
