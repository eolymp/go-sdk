// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package community

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

type _AchievementServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type AchievementServiceService struct {
	base string
	cli  _AchievementServiceHttpClient
}

// NewAchievementServiceHttpClient constructs client for AchievementService
func NewAchievementServiceHttpClient(url string, cli _AchievementServiceHttpClient) *AchievementServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &AchievementServiceService{base: url, cli: cli}
}

func (s *AchievementServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *AchievementServiceService) AssignAchievement(ctx context.Context, in *AssignAchievementInput) (*AssignAchievementOutput, error) {
	out := &AssignAchievementOutput{}
	path := "/members/" + url.PathEscape(in.GetMemberId()) + "/achievements/" + url.PathEscape(in.GetAchievementId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
		in.AchievementId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AchievementServiceService) UnassignAchievement(ctx context.Context, in *UnassignAchievementInput) (*UnassignAchievementOutput, error) {
	out := &UnassignAchievementOutput{}
	path := "/members/" + url.PathEscape(in.GetMemberId()) + "/achievements/" + url.PathEscape(in.GetAchievementId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
		in.AchievementId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AchievementServiceService) ListAchievements(ctx context.Context, in *ListAchievementsInput) (*ListAchievementsOutput, error) {
	out := &ListAchievementsOutput{}
	path := "/members/" + url.PathEscape(in.GetMemberId()) + "/achievements"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.MemberId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
