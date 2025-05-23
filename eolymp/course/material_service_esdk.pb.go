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

type _MaterialServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type MaterialServiceService struct {
	base string
	cli  _MaterialServiceHttpClient
}

// NewMaterialServiceHttpClient constructs client for MaterialService
func NewMaterialServiceHttpClient(url string, cli _MaterialServiceHttpClient) *MaterialServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &MaterialServiceService{base: url, cli: cli}
}

func (s *MaterialServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *MaterialServiceService) CreateMaterial(ctx context.Context, in *CreateMaterialInput) (*CreateMaterialOutput, error) {
	out := &CreateMaterialOutput{}
	path := "/materials"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MaterialServiceService) UpdateMaterial(ctx context.Context, in *UpdateMaterialInput) (*UpdateMaterialOutput, error) {
	out := &UpdateMaterialOutput{}
	path := "/materials/" + url.PathEscape(in.GetMaterialId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MaterialId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MaterialServiceService) MoveMaterial(ctx context.Context, in *MoveMaterialInput) (*MoveMaterialOutput, error) {
	out := &MoveMaterialOutput{}
	path := "/materials/" + url.PathEscape(in.GetMaterialId()) + "/move"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MaterialId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MaterialServiceService) DeleteMaterial(ctx context.Context, in *DeleteMaterialInput) (*DeleteMaterialOutput, error) {
	out := &DeleteMaterialOutput{}
	path := "/materials/" + url.PathEscape(in.GetMaterialId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MaterialId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MaterialServiceService) DescribeMaterial(ctx context.Context, in *DescribeMaterialInput) (*DescribeMaterialOutput, error) {
	out := &DescribeMaterialOutput{}
	path := "/materials/" + url.PathEscape(in.GetMaterialId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MaterialId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MaterialServiceService) ListMaterials(ctx context.Context, in *ListMaterialsInput) (*ListMaterialsOutput, error) {
	out := &ListMaterialsOutput{}
	path := "/materials"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MaterialServiceService) ReportProgress(ctx context.Context, in *ReportProgressInput) (*ReportProgressOutput, error) {
	out := &ReportProgressOutput{}
	path := "/materials/" + url.PathEscape(in.GetMaterialId()) + "/progress"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MaterialId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MaterialServiceService) GradeMaterial(ctx context.Context, in *GradeMaterialInput) (*GradeMaterialOutput, error) {
	out := &GradeMaterialOutput{}
	path := "/materials/" + url.PathEscape(in.GetMaterialId()) + "/grade"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MaterialId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
