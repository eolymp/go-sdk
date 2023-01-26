// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package drive

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

type _DriveHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type DriveService struct {
	base string
	cli  _DriveHttpClient
}

// NewDriveHttpClient constructs client for Drive
func NewDriveHttpClient(url string, cli _DriveHttpClient) *DriveService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &DriveService{base: url, cli: cli}
}

func (s *DriveService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *DriveService) DescribeFile(ctx context.Context, in *DescribeFileInput) (*DescribeFileOutput, error) {
	out := &DescribeFileOutput{}
	path := "/drive/files/" + url.PathEscape(in.GetFileId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.FileId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *DriveService) ListFiles(ctx context.Context, in *ListFilesInput) (*ListFilesOutput, error) {
	out := &ListFilesOutput{}
	path := "/drive/files"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *DriveService) CreateFile(ctx context.Context, in *CreateFileInput) (*CreateFileOutput, error) {
	out := &CreateFileOutput{}
	path := "/drive/files"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *DriveService) UpdateFile(ctx context.Context, in *UpdateFileInput) (*UpdateFileOutput, error) {
	out := &UpdateFileOutput{}
	path := "/drive/files/" + url.PathEscape(in.GetFileId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.FileId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *DriveService) DeleteFile(ctx context.Context, in *DeleteFileInput) (*DeleteFileOutput, error) {
	out := &DeleteFileOutput{}
	path := "/drive/files/" + url.PathEscape(in.GetFileId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.FileId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *DriveService) StartMultipartUpload(ctx context.Context, in *StartMultipartUploadInput) (*StartMultipartUploadOutput, error) {
	out := &StartMultipartUploadOutput{}
	path := "/drive/uploads"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *DriveService) UploadPart(ctx context.Context, in *UploadPartInput) (*UploadPartOutput, error) {
	out := &UploadPartOutput{}
	path := "/drive/uploads/" + url.PathEscape(in.GetUploadId()) + "/parts"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.UploadId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *DriveService) CompleteMultipartUpload(ctx context.Context, in *CompleteMultipartUploadInput) (*CompleteMultipartUploadOutput, error) {
	out := &CompleteMultipartUploadOutput{}
	path := "/drive/uploads/" + url.PathEscape(in.GetUploadId()) + "/complete"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.UploadId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
