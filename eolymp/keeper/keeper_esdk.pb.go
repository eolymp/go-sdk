// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package keeper

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

type _KeeperHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type KeeperService struct {
	base string
	cli  _KeeperHttpClient
}

// NewKeeperHttpClient constructs client for Keeper
func NewKeeperHttpClient(url string, cli _KeeperHttpClient) *KeeperService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &KeeperService{base: url, cli: cli}
}

func (s *KeeperService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *KeeperService) CreateObject(ctx context.Context, in *CreateObjectInput) (*CreateObjectOutput, error) {
	out := &CreateObjectOutput{}
	path := "/objects"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *KeeperService) DescribeObject(ctx context.Context, in *DescribeObjectInput) (*DescribeObjectOutput, error) {
	out := &DescribeObjectOutput{}
	path := "/objects/" + url.PathEscape(in.GetKey())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.Key = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *KeeperService) DownloadObject(ctx context.Context, in *DownloadObjectInput) (*DownloadObjectOutput, error) {
	out := &DownloadObjectOutput{}
	path := "/objects/" + url.PathEscape(in.GetKey()) + "/data"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.Key = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
