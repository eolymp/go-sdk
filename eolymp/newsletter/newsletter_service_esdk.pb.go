// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package newsletter

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

type _NewsletterServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type NewsletterServiceService struct {
	base string
	cli  _NewsletterServiceHttpClient
}

// NewNewsletterServiceHttpClient constructs client for NewsletterService
func NewNewsletterServiceHttpClient(url string, cli _NewsletterServiceHttpClient) *NewsletterServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &NewsletterServiceService{base: url, cli: cli}
}

func (s *NewsletterServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *NewsletterServiceService) DescribeNewsletter(ctx context.Context, in *DescribeNewsletterInput) (*DescribeNewsletterOutput, error) {
	out := &DescribeNewsletterOutput{}
	path := "/newsletter/" + url.PathEscape(in.GetNewsletterId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.NewsletterId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *NewsletterServiceService) ListNewsletters(ctx context.Context, in *ListNewslettersInput) (*ListNewslettersOutput, error) {
	out := &ListNewslettersOutput{}
	path := "/newsletter"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *NewsletterServiceService) CreateNewsletter(ctx context.Context, in *CreateNewsletterInput) (*CreateNewsletterOutput, error) {
	out := &CreateNewsletterOutput{}
	path := "/newsletter"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *NewsletterServiceService) UpdateNewsletter(ctx context.Context, in *UpdateNewsletterInput) (*UpdateNewsletterOutput, error) {
	out := &UpdateNewsletterOutput{}
	path := "/newsletter/" + url.PathEscape(in.GetNewsletterId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.NewsletterId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *NewsletterServiceService) DeleteNewsletter(ctx context.Context, in *DeleteNewsletterInput) (*DeleteNewsletterOutput, error) {
	out := &DeleteNewsletterOutput{}
	path := "/newsletter/" + url.PathEscape(in.GetNewsletterId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.NewsletterId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *NewsletterServiceService) SendNewsletter(ctx context.Context, in *SendNewsletterInput) (*SendNewsletterOutput, error) {
	out := &SendNewsletterOutput{}
	path := "/newsletter/" + url.PathEscape(in.GetNewsletterId()) + "/send"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.NewsletterId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *NewsletterServiceService) TestNewsletter(ctx context.Context, in *TestNewsletterInput) (*TestNewsletterOutput, error) {
	out := &TestNewsletterOutput{}
	path := "/newsletter/" + url.PathEscape(in.GetNewsletterId()) + "/test"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.NewsletterId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *NewsletterServiceService) DescribeNewsletterTranslation(ctx context.Context, in *DescribeNewsletterTranslationInput) (*DescribeNewsletterTranslationOutput, error) {
	out := &DescribeNewsletterTranslationOutput{}
	path := "/newsletter/" + url.PathEscape(in.GetNewsletterId()) + "/translations/" + url.PathEscape(in.GetTranslationId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.NewsletterId = ""
		in.TranslationId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *NewsletterServiceService) ListNewsletterTranslations(ctx context.Context, in *ListNewsletterTranslationsInput) (*ListNewsletterTranslationsOutput, error) {
	out := &ListNewsletterTranslationsOutput{}
	path := "/newsletter/" + url.PathEscape(in.GetNewsletterId()) + "/translations"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.NewsletterId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *NewsletterServiceService) CreateNewsletterTranslation(ctx context.Context, in *CreateNewsletterTranslationInput) (*CreateNewsletterTranslationOutput, error) {
	out := &CreateNewsletterTranslationOutput{}
	path := "/newsletter/" + url.PathEscape(in.GetNewsletterId()) + "/translations"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.NewsletterId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *NewsletterServiceService) UpdateNewsletterTranslation(ctx context.Context, in *UpdateNewsletterTranslationInput) (*UpdateNewsletterTranslationOutput, error) {
	out := &UpdateNewsletterTranslationOutput{}
	path := "/newsletter/" + url.PathEscape(in.GetNewsletterId()) + "/translations/" + url.PathEscape(in.GetTranslationId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.NewsletterId = ""
		in.TranslationId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *NewsletterServiceService) DeleteNewsletterTranslation(ctx context.Context, in *DeleteNewsletterTranslationInput) (*DeleteNewsletterTranslationOutput, error) {
	out := &DeleteNewsletterTranslationOutput{}
	path := "/newsletter/" + url.PathEscape(in.GetNewsletterId()) + "/translations/" + url.PathEscape(in.GetTranslationId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.NewsletterId = ""
		in.TranslationId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}