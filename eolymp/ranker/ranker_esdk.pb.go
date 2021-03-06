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
	os "os"
	strings "strings"
)

// NewRanker constructs client for Ranker
func NewRanker(cli RankerHTTPClient) *RankerService {
	base := "https://api.eolymp.com"
	if v := os.Getenv("EOLYMP_API_URL"); v != "" {
		base = v
	}
	return &RankerService{base: strings.TrimSuffix(base, "/"), cli: cli}
}

type RankerHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type RankerService struct {
	base string
	cli  RankerHTTPClient
}

// invoke RPC method using twirp-like protocol
func (s *RankerService) invoke(ctx context.Context, method string, in, out proto.Message) (err error) {
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

func (s *RankerService) CreateScoreboard(ctx context.Context, in *CreateScoreboardInput) (*CreateScoreboardOutput, error) {
	out := &CreateScoreboardOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/CreateScoreboard", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) UpdateScoreboard(ctx context.Context, in *UpdateScoreboardInput) (*UpdateScoreboardOutput, error) {
	out := &UpdateScoreboardOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/UpdateScoreboard", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) RebuildScoreboard(ctx context.Context, in *RebuildScoreboardInput) (*RebuildScoreboardOutput, error) {
	out := &RebuildScoreboardOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/RebuildScoreboard", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) DeleteScoreboard(ctx context.Context, in *DeleteScoreboardInput) (*DeleteScoreboardOutput, error) {
	out := &DeleteScoreboardOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/DeleteScoreboard", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) DescribeScoreboard(ctx context.Context, in *DescribeScoreboardInput) (*DescribeScoreboardOutput, error) {
	out := &DescribeScoreboardOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/DescribeScoreboard", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) ListScoreboards(ctx context.Context, in *ListScoreboardsInput) (*ListScoreboardsOutput, error) {
	out := &ListScoreboardsOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/ListScoreboards", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) DescribeScoreboardRow(ctx context.Context, in *DescribeScoreboardRowInput) (*DescribeScoreboardRowOutput, error) {
	out := &DescribeScoreboardRowOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/DescribeScoreboardRow", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) ListScoreboardRows(ctx context.Context, in *ListScoreboardRowsInput) (*ListScoreboardRowsOutput, error) {
	out := &ListScoreboardRowsOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/ListScoreboardRows", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) AddScoreboardColumn(ctx context.Context, in *AddScoreboardColumnInput) (*AddScoreboardColumnOutput, error) {
	out := &AddScoreboardColumnOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/AddScoreboardColumn", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) DeleteScoreboardColumn(ctx context.Context, in *DeleteScoreboardColumnInput) (*DeleteScoreboardColumnOutput, error) {
	out := &DeleteScoreboardColumnOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/DeleteScoreboardColumn", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) DescribeScoreboardColumn(ctx context.Context, in *DescribeScoreboardColumnInput) (*DescribeScoreboardColumnOutput, error) {
	out := &DescribeScoreboardColumnOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/DescribeScoreboardColumn", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) ListScoreboardColumns(ctx context.Context, in *ListScoreboardColumnsInput) (*ListScoreboardColumnsOutput, error) {
	out := &ListScoreboardColumnsOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/ListScoreboardColumns", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *RankerService) ListActivities(ctx context.Context, in *ListActivitiesInput) (*ListActivitiesOutput, error) {
	out := &ListActivitiesOutput{}

	if err := s.invoke(ctx, "eolymp.ranker.Ranker/ListActivities", in, out); err != nil {
		return nil, err
	}

	return out, nil
}
