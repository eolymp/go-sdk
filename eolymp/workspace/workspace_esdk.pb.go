// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package workspace

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

// NewWorkspace constructs client for Workspace
func NewWorkspace(cli WorkspaceHTTPClient) *WorkspaceService {
	base := "https://api.e-olymp.com"
	if v := os.Getenv("EOLYMP_API_URL"); v != "" {
		base = v
	}
	return &WorkspaceService{base: strings.TrimSuffix(base, "/"), cli: cli}
}

type WorkspaceHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type WorkspaceService struct {
	base string
	cli  WorkspaceHTTPClient
}

// invoke RPC method using twirp-like protocol
func (s *WorkspaceService) invoke(ctx context.Context, method string, in, out proto.Message) error {
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
		return fmt.Errorf("non-200 response code (%v)", resp.StatusCode)
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

func (s *WorkspaceService) DescribeProject(ctx context.Context, in *DescribeProjectInput) (*DescribeProjectOutput, error) {
	out := &DescribeProjectOutput{}

	if err := s.invoke(ctx, "eolymp.workspace.Workspace/DescribeProject", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkspaceService) ListProjects(ctx context.Context, in *ListProjectsInput) (*ListProjectsOutput, error) {
	out := &ListProjectsOutput{}

	if err := s.invoke(ctx, "eolymp.workspace.Workspace/ListProjects", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkspaceService) CreateProject(ctx context.Context, in *CreateProjectInput) (*CreateProjectOutput, error) {
	out := &CreateProjectOutput{}

	if err := s.invoke(ctx, "eolymp.workspace.Workspace/CreateProject", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkspaceService) UpdateProject(ctx context.Context, in *UpdateProjectInput) (*UpdateProjectOutput, error) {
	out := &UpdateProjectOutput{}

	if err := s.invoke(ctx, "eolymp.workspace.Workspace/UpdateProject", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkspaceService) DeleteProject(ctx context.Context, in *DeleteProjectInput) (*DeleteProjectOutput, error) {
	out := &DeleteProjectOutput{}

	if err := s.invoke(ctx, "eolymp.workspace.Workspace/DeleteProject", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkspaceService) ListFiles(ctx context.Context, in *ListFilesInput) (*ListFilesOutput, error) {
	out := &ListFilesOutput{}

	if err := s.invoke(ctx, "eolymp.workspace.Workspace/ListFiles", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkspaceService) DescribeFile(ctx context.Context, in *DescribeFileInput) (*DescribeFileOutput, error) {
	out := &DescribeFileOutput{}

	if err := s.invoke(ctx, "eolymp.workspace.Workspace/DescribeFile", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkspaceService) UploadFile(ctx context.Context, in *UploadFileInput) (*UploadFileOutput, error) {
	out := &UploadFileOutput{}

	if err := s.invoke(ctx, "eolymp.workspace.Workspace/UploadFile", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *WorkspaceService) RemoveFile(ctx context.Context, in *RemoveFileInput) (*RemoveFileOutput, error) {
	out := &RemoveFileOutput{}

	if err := s.invoke(ctx, "eolymp.workspace.Workspace/RemoveFile", in, out); err != nil {
		return nil, err
	}

	return out, nil
}
