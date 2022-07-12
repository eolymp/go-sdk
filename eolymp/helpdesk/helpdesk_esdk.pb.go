// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package helpdesk

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

// NewHelpdesk constructs client for Helpdesk
func NewHelpdesk(cli HelpdeskHTTPClient) *HelpdeskService {
	base := "https://api.eolymp.com"
	if v := os.Getenv("EOLYMP_API_URL"); v != "" {
		base = v
	}
	return &HelpdeskService{base: strings.TrimSuffix(base, "/"), cli: cli}
}

type HelpdeskHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HelpdeskService struct {
	base string
	cli  HelpdeskHTTPClient
}

// invoke RPC method using twirp-like protocol
func (s *HelpdeskService) invoke(ctx context.Context, method string, in, out proto.Message) error {
	input, err := protojson.Marshal(in)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, s.base+"/twirp/"+method, bytes.NewReader(input))
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

func (s *HelpdeskService) DescribeDocument(ctx context.Context, in *DescribeDocumentInput) (*DescribeDocumentOutput, error) {
	out := &DescribeDocumentOutput{}

	if err := s.invoke(ctx, "eolymp.helpdesk.Helpdesk/DescribeDocument", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *HelpdeskService) ListDocuments(ctx context.Context, in *ListDocumentsInput) (*ListDocumentsOutput, error) {
	out := &ListDocumentsOutput{}

	if err := s.invoke(ctx, "eolymp.helpdesk.Helpdesk/ListDocuments", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *HelpdeskService) CreateDocument(ctx context.Context, in *CreateDocumentInput) (*CreateDocumentOutput, error) {
	out := &CreateDocumentOutput{}

	if err := s.invoke(ctx, "eolymp.helpdesk.Helpdesk/CreateDocument", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *HelpdeskService) UpdateDocument(ctx context.Context, in *UpdateDocumentInput) (*UpdateDocumentOutput, error) {
	out := &UpdateDocumentOutput{}

	if err := s.invoke(ctx, "eolymp.helpdesk.Helpdesk/UpdateDocument", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *HelpdeskService) DeleteDocument(ctx context.Context, in *DeleteDocumentInput) (*DeleteDocumentOutput, error) {
	out := &DeleteDocumentOutput{}

	if err := s.invoke(ctx, "eolymp.helpdesk.Helpdesk/DeleteDocument", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *HelpdeskService) DescribePath(ctx context.Context, in *DescribePathInput) (*DescribePathOutput, error) {
	out := &DescribePathOutput{}

	if err := s.invoke(ctx, "eolymp.helpdesk.Helpdesk/DescribePath", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *HelpdeskService) ListPaths(ctx context.Context, in *ListPathsInput) (*ListPathsOutput, error) {
	out := &ListPathsOutput{}

	if err := s.invoke(ctx, "eolymp.helpdesk.Helpdesk/ListPaths", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *HelpdeskService) ListParents(ctx context.Context, in *ListParentsInput) (*ListParentsOutput, error) {
	out := &ListParentsOutput{}

	if err := s.invoke(ctx, "eolymp.helpdesk.Helpdesk/ListParents", in, out); err != nil {
		return nil, err
	}

	return out, nil
}
