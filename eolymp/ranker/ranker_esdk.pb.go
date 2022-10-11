// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package ranker

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

type _RankerHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type RankerService struct {
	base string
	cli  _RankerHttpClient
}

// NewRankerHttpClient constructs client for Ranker
func NewRankerHttpClient(url string, cli _RankerHttpClient) *RankerService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &RankerService{base: url, cli: cli}
}

func (s *RankerService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *RankerService) CreateScoreboard(ctx context.Context, in *CreateScoreboardInput) (*CreateScoreboardOutput, error) {
	out := &CreateScoreboardOutput{}
	path := "/scoreboards"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) UpdateScoreboard(ctx context.Context, in *UpdateScoreboardInput) (*UpdateScoreboardOutput, error) {
	out := &UpdateScoreboardOutput{}
	path := "/scoreboards/" + url.PathEscape(in.GetScoreboardId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ScoreboardId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) RebuildScoreboard(ctx context.Context, in *RebuildScoreboardInput) (*RebuildScoreboardOutput, error) {
	out := &RebuildScoreboardOutput{}
	path := "/scoreboards/" + url.PathEscape(in.GetScoreboardId()) + "/rebuild"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ScoreboardId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) DeleteScoreboard(ctx context.Context, in *DeleteScoreboardInput) (*DeleteScoreboardOutput, error) {
	out := &DeleteScoreboardOutput{}
	path := "/scoreboards/" + url.PathEscape(in.GetScoreboardId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ScoreboardId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) DescribeScoreboard(ctx context.Context, in *DescribeScoreboardInput) (*DescribeScoreboardOutput, error) {
	out := &DescribeScoreboardOutput{}
	path := "/scoreboards/" + url.PathEscape(in.GetScoreboardId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ScoreboardId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) ListScoreboards(ctx context.Context, in *ListScoreboardsInput) (*ListScoreboardsOutput, error) {
	out := &ListScoreboardsOutput{}
	path := "/scoreboards"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) DescribeScoreboardRow(ctx context.Context, in *DescribeScoreboardRowInput) (*DescribeScoreboardRowOutput, error) {
	out := &DescribeScoreboardRowOutput{}
	path := "/scoreboards/" + url.PathEscape(in.GetScoreboardId()) + "/rows/" + url.PathEscape(in.GetMemberId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ScoreboardId = ""
		in.MemberId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) ListScoreboardRows(ctx context.Context, in *ListScoreboardRowsInput) (*ListScoreboardRowsOutput, error) {
	out := &ListScoreboardRowsOutput{}
	path := "/scoreboards/" + url.PathEscape(in.GetScoreboardId()) + "/rows"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ScoreboardId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) AddScoreboardColumn(ctx context.Context, in *AddScoreboardColumnInput) (*AddScoreboardColumnOutput, error) {
	out := &AddScoreboardColumnOutput{}
	path := "/scoreboards/" + url.PathEscape(in.GetScoreboardId()) + "/columns"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ScoreboardId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) DeleteScoreboardColumn(ctx context.Context, in *DeleteScoreboardColumnInput) (*DeleteScoreboardColumnOutput, error) {
	out := &DeleteScoreboardColumnOutput{}
	path := "/scoreboards/" + url.PathEscape(in.GetScoreboardId()) + "/columns/" + url.PathEscape(in.GetColumnId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ScoreboardId = ""
		in.ColumnId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) DescribeScoreboardColumn(ctx context.Context, in *DescribeScoreboardColumnInput) (*DescribeScoreboardColumnOutput, error) {
	out := &DescribeScoreboardColumnOutput{}
	path := "/scoreboards/" + url.PathEscape(in.GetScoreboardId()) + "/columns/" + url.PathEscape(in.GetColumnId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ScoreboardId = ""
		in.ColumnId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) ListScoreboardColumns(ctx context.Context, in *ListScoreboardColumnsInput) (*ListScoreboardColumnsOutput, error) {
	out := &ListScoreboardColumnsOutput{}
	path := "/scoreboards/" + url.PathEscape(in.GetScoreboardId()) + "/columns"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ScoreboardId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) ListActivities(ctx context.Context, in *ListActivitiesInput) (*ListActivitiesOutput, error) {
	out := &ListActivitiesOutput{}
	path := "/scoreboards/" + url.PathEscape(in.GetScoreboardId()) + "/activities"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.ScoreboardId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
