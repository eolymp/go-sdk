// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package l10n

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

type _LocalizationServiceHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type LocalizationServiceService struct {
	base string
	cli  _LocalizationServiceHttpClient
}

// NewLocalizationServiceHttpClient constructs client for LocalizationService
func NewLocalizationServiceHttpClient(url string, cli _LocalizationServiceHttpClient) *LocalizationServiceService {
	if url == "" {
		url = os.Getenv("EOLYMP_API_URL")
		if url == "" {
			url = "https://api.eolymp.com"
		}
	}

	return &LocalizationServiceService{base: url, cli: cli}
}

func (s *LocalizationServiceService) do(ctx context.Context, verb, path string, in, out proto.Message) (err error) {
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

func (s *LocalizationServiceService) CreateTerm(ctx context.Context, in *CreateTermInput) (*CreateTermOutput, error) {
	out := &CreateTermOutput{}
	path := "/terms"

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) ListTerms(ctx context.Context, in *ListTermsInput) (*ListTermsOutput, error) {
	out := &ListTermsOutput{}
	path := "/terms"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) UpdateTerm(ctx context.Context, in *UpdateTermInput) (*UpdateTermOutput, error) {
	out := &UpdateTermOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) ApproveTerm(ctx context.Context, in *ApproveTermInput) (*ApproveTermOutput, error) {
	out := &ApproveTermOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId()) + "/approve"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) DeprecateTerm(ctx context.Context, in *DeprecateTermInput) (*DeprecateTermOutput, error) {
	out := &DeprecateTermOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId()) + "/deprecate"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) DeleteTerm(ctx context.Context, in *DeleteTermInput) (*DeleteTermOutput, error) {
	out := &DeleteTermOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) DescribeTerm(ctx context.Context, in *DescribeTermInput) (*DescribeTermOutput, error) {
	out := &DescribeTermOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) ImportTerms(ctx context.Context, in *ImportTermsInput) (*ImportTermsOutput, error) {
	out := &ImportTermsOutput{}
	path := "/terms"

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) AddLocale(ctx context.Context, in *AddLocaleInput) (*AddLocaleOutput, error) {
	out := &AddLocaleOutput{}
	path := "/locales/" + url.PathEscape(in.GetLocaleCode())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.LocaleCode = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) RemoveLocale(ctx context.Context, in *RemoveLocaleInput) (*RemoveLocaleOutput, error) {
	out := &RemoveLocaleOutput{}
	path := "/locales/" + url.PathEscape(in.GetLocaleCode())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.LocaleCode = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) ListLocales(ctx context.Context, in *ListLocalesInput) (*ListLocalesOutput, error) {
	out := &ListLocalesOutput{}
	path := "/locales"

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) TranslateTerm(ctx context.Context, in *TranslateTermInput) (*TranslateTermOutput, error) {
	out := &TranslateTermOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId()) + "/translations"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) ListTranslations(ctx context.Context, in *ListTranslationsInput) (*ListTranslationsOutput, error) {
	out := &ListTranslationsOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId()) + "/translations"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) DeleteTranslation(ctx context.Context, in *DeleteTranslationInput) (*DeleteTranslationOutput, error) {
	out := &DeleteTranslationOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId()) + "/translations/" + url.PathEscape(in.GetTranslationId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
		in.TranslationId = ""
	}

	if err := s.do(ctx, "DELETE", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) SuggestTranslation(ctx context.Context, in *SuggestTranslationInput) (*SuggestTranslationOutput, error) {
	out := &SuggestTranslationOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId()) + "/suggestions/" + url.PathEscape(in.GetLocale())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
		in.Locale = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) UpdateTranslation(ctx context.Context, in *UpdateTranslationInput) (*UpdateTranslationOutput, error) {
	out := &UpdateTranslationOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId()) + "/translations/" + url.PathEscape(in.GetTranslationId())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
		in.TranslationId = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) ApproveTranslation(ctx context.Context, in *ApproveTranslationInput) (*ApproveTranslationOutput, error) {
	out := &ApproveTranslationOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId()) + "/translations/" + url.PathEscape(in.GetTranslationId()) + "/approve"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
		in.TranslationId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) RejectTranslation(ctx context.Context, in *RejectTranslationInput) (*RejectTranslationOutput, error) {
	out := &RejectTranslationOutput{}
	path := "/terms/" + url.PathEscape(in.GetTermId()) + "/translations/" + url.PathEscape(in.GetTranslationId()) + "/reject"

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.TermId = ""
		in.TranslationId = ""
	}

	if err := s.do(ctx, "POST", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) ImportTranslations(ctx context.Context, in *ImportTranslationsInput) (*ImportTranslationsOutput, error) {
	out := &ImportTranslationsOutput{}
	path := "/translations/" + url.PathEscape(in.GetLocale())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.Locale = ""
	}

	if err := s.do(ctx, "PUT", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *LocalizationServiceService) ExportTranslations(ctx context.Context, in *ExportTranslationsInput) (*ExportTranslationsOutput, error) {
	out := &ExportTranslationsOutput{}
	path := "/translations/" + url.PathEscape(in.GetLocale())

	// Cleanup URL parameters to avoid any ambiguity
	if in != nil {
		in.Locale = ""
	}

	if err := s.do(ctx, "GET", path, in, out); err != nil {
		return nil, err
	}

	return out, nil
}
