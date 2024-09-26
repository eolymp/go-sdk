// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package course

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

type _StudentServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type StudentServiceService struct {
	base string
	cli  _StudentServiceHttpClient
}

// NewStudentServiceHttpClient constructs client for StudentService
func NewStudentServiceHttpClient(url string, cli _StudentServiceHttpClient) *StudentServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &StudentServiceService{base: url, cli: cli}
}

func (s *StudentServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *StudentServiceService) CreateStudent(ctx context.Context, in *CreateStudentInput) (*CreateStudentOutput, error) {
	out := &CreateStudentOutput{}
	path := "/students"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StudentServiceService) UpdateStudent(ctx context.Context, in *UpdateStudentInput) (*UpdateStudentOutput, error) {
	out := &UpdateStudentOutput{}
	path := "/students/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StudentServiceService) DeleteStudent(ctx context.Context, in *DeleteStudentInput) (*DeleteStudentOutput, error) {
	out := &DeleteStudentOutput{}
	path := "/students/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StudentServiceService) DescribeViewer(ctx context.Context, in *DescribeViewerInput) (*DescribeViewerOutput, error) {
	out := &DescribeViewerOutput{}
	path := "/viewer/student"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StudentServiceService) DescribeStudent(ctx context.Context, in *DescribeStudentInput) (*DescribeStudentOutput, error) {
	out := &DescribeStudentOutput{}
	path := "/students/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StudentServiceService) ListStudents(ctx context.Context, in *ListStudentsInput) (*ListStudentsOutput, error) {
	out := &ListStudentsOutput{}
	path := "/students"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *StudentServiceService) JoinCourse(ctx context.Context, in *JoinCourseInput) (*JoinCourseOutput, error) {
	out := &JoinCourseOutput{}
	path := "/join"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
