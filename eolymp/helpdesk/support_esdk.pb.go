// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package helpdesk

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

type _SupportHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type SupportService struct {
	base string
	cli  _SupportHttpClient
}

// NewSupportHttpClient constructs client for Support
func NewSupportHttpClient(url string, cli _SupportHttpClient) *SupportService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &SupportService{base: url, cli: cli}
}

func (s *SupportService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *SupportService) CreateTicket(ctx context.Context, in *CreateTicketInput) (*CreateTicketOutput, error) {
	out := &CreateTicketOutput{}
	path := "/helpdesk/tickets"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SupportService) UpdateTicket(ctx context.Context, in *UpdateTicketInput) (*UpdateTicketOutput, error) {
	out := &UpdateTicketOutput{}
	path := "/helpdesk/tickets/" + url.PathEscape(in.GetTicketId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TicketId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SupportService) DeleteTicket(ctx context.Context, in *DeleteTicketInput) (*DeleteTicketOutput, error) {
	out := &DeleteTicketOutput{}
	path := "/helpdesk/tickets/" + url.PathEscape(in.GetTicketId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TicketId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SupportService) DescribeTicket(ctx context.Context, in *DescribeTicketInput) (*DescribeTicketOutput, error) {
	out := &DescribeTicketOutput{}
	path := "/helpdesk/tickets/" + url.PathEscape(in.GetTicketId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TicketId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SupportService) ListTickets(ctx context.Context, in *ListTicketsInput) (*ListTicketsOutput, error) {
	out := &ListTicketsOutput{}
	path := "/helpdesk/tickets"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SupportService) ApproveTicket(ctx context.Context, in *ApproveTicketInput) (*ApproveTicketOutput, error) {
	out := &ApproveTicketOutput{}
	path := "/helpdesk/tickets/" + url.PathEscape(in.GetTicketId()) + "/approve"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TicketId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SupportService) RejectTicket(ctx context.Context, in *RejectTicketInput) (*RejectTicketOutput, error) {
	out := &RejectTicketOutput{}
	path := "/helpdesk/tickets/" + url.PathEscape(in.GetTicketId()) + "/reject"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TicketId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *SupportService) CloseTicket(ctx context.Context, in *CloseTicketInput) (*CloseTicketOutput, error) {
	out := &CloseTicketOutput{}
	path := "/helpdesk/tickets/" + url.PathEscape(in.GetTicketId()) + "/close"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TicketId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}