// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package taxonomy

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

type _TopicServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type TopicServiceService struct {
	base string
	cli  _TopicServiceHttpClient
}

// NewTopicServiceHttpClient constructs client for TopicService
func NewTopicServiceHttpClient(url string, cli _TopicServiceHttpClient) *TopicServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &TopicServiceService{base: url, cli: cli}
}

func (s *TopicServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *TopicServiceService) CreateTopic(ctx context.Context, in *CreateTopicInput) (*CreateTopicOutput, error) {
	out := &CreateTopicOutput{}
	path := "/topics"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TopicServiceService) DeleteTopic(ctx context.Context, in *DeleteTopicInput) (*DeleteTopicOutput, error) {
	out := &DeleteTopicOutput{}
	path := "/topics/" + url.PathEscape(in.GetTopicId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TopicId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TopicServiceService) UpdateTopic(ctx context.Context, in *UpdateTopicInput) (*UpdateTopicOutput, error) {
	out := &UpdateTopicOutput{}
	path := "/topics/" + url.PathEscape(in.GetTopicId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TopicId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TopicServiceService) DescribeTopic(ctx context.Context, in *DescribeTopicInput) (*DescribeTopicOutput, error) {
	out := &DescribeTopicOutput{}
	path := "/topics/" + url.PathEscape(in.GetTopicId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TopicId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TopicServiceService) ListTopics(ctx context.Context, in *ListTopicsInput) (*ListTopicsOutput, error) {
	out := &ListTopicsOutput{}
	path := "/topics"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TopicServiceService) TranslateTopic(ctx context.Context, in *TranslateTopicInput) (*TranslateTopicOutput, error) {
	out := &TranslateTopicOutput{}
	path := "/topics/" + url.PathEscape(in.GetTopicId()) + "/translations/" + url.PathEscape(in.GetLocale())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TopicId = ""
		in.Locale = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TopicServiceService) DeleteTranslation(ctx context.Context, in *DeleteTranslationInput) (*DeleteTranslationOutput, error) {
	out := &DeleteTranslationOutput{}
	path := "/topics/" + url.PathEscape(in.GetTopicId()) + "/translations/" + url.PathEscape(in.GetLocale())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TopicId = ""
		in.Locale = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *TopicServiceService) ListTranslations(ctx context.Context, in *ListTranslationsInput) (*ListTranslationsOutput, error) {
	out := &ListTranslationsOutput{}
	path := "/topics/" + url.PathEscape(in.GetTopicId()) + "/translations"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TopicId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}