// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package atlas

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

// NewAtlas constructs client for Atlas
func NewAtlas(cli AtlasHTTPClient) *AtlasService {
	base := "https://api.e-olymp.com"
	if v := os.Getenv("EOLYMP_API_URL"); v != "" {
		base = v
	}
	return &AtlasService{base: strings.TrimSuffix(base, "/"), cli: cli}
}

type AtlasHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type AtlasService struct {
	base string
	cli  AtlasHTTPClient
}

// invoke RPC method using twirp-like protocol
func (s *AtlasService) invoke(ctx context.Context, method string, in, out proto.Message) error {
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

func (s *AtlasService) CreateProblem(ctx context.Context, in *CreateProblemInput) (*CreateProblemOutput, error) {
	out := &CreateProblemOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/CreateProblem", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DeleteProblem(ctx context.Context, in *DeleteProblemInput) (*DeleteProblemOutput, error) {
	out := &DeleteProblemOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DeleteProblem", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) ListProblems(ctx context.Context, in *ListProblemsInput) (*ListProblemsOutput, error) {
	out := &ListProblemsOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/ListProblems", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DescribeProblem(ctx context.Context, in *DescribeProblemInput) (*DescribeProblemOutput, error) {
	out := &DescribeProblemOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DescribeProblem", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) ListExamples(ctx context.Context, in *ListExamplesInput) (*ListExamplesOutput, error) {
	out := &ListExamplesOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/ListExamples", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) UpdateVerifier(ctx context.Context, in *UpdateVerifierInput) (*UpdateVerifierOutput, error) {
	out := &UpdateVerifierOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/UpdateVerifier", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DescribeVerifier(ctx context.Context, in *DescribeVerifierInput) (*DescribeVerifierOutput, error) {
	out := &DescribeVerifierOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DescribeVerifier", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) UpdateInteractor(ctx context.Context, in *UpdateInteractorInput) (*UpdateInteractorOutput, error) {
	out := &UpdateInteractorOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/UpdateInteractor", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DescribeInteractor(ctx context.Context, in *DescribeInteractorInput) (*DescribeInteractorOutput, error) {
	out := &DescribeInteractorOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DescribeInteractor", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) CreateStatement(ctx context.Context, in *CreateStatementInput) (*CreateStatementOutput, error) {
	out := &CreateStatementOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/CreateStatement", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) UpdateStatement(ctx context.Context, in *UpdateStatementInput) (*UpdateStatementOutput, error) {
	out := &UpdateStatementOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/UpdateStatement", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DeleteStatement(ctx context.Context, in *DeleteStatementInput) (*DeleteStatementOutput, error) {
	out := &DeleteStatementOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DeleteStatement", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) ListStatements(ctx context.Context, in *ListStatementsInput) (*ListStatementsOutput, error) {
	out := &ListStatementsOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/ListStatements", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DescribeStatement(ctx context.Context, in *DescribeStatementInput) (*DescribeStatementOutput, error) {
	out := &DescribeStatementOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DescribeStatement", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) CreateTestset(ctx context.Context, in *CreateTestsetInput) (*CreateTestsetOutput, error) {
	out := &CreateTestsetOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/CreateTestset", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) UpdateTestset(ctx context.Context, in *UpdateTestsetInput) (*UpdateTestsetOutput, error) {
	out := &UpdateTestsetOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/UpdateTestset", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DeleteTestset(ctx context.Context, in *DeleteTestsetInput) (*DeleteTestsetOutput, error) {
	out := &DeleteTestsetOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DeleteTestset", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) ListTestsets(ctx context.Context, in *ListTestsetsInput) (*ListTestsetsOutput, error) {
	out := &ListTestsetsOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/ListTestsets", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DescribeTestset(ctx context.Context, in *DescribeTestsetInput) (*DescribeTestsetOutput, error) {
	out := &DescribeTestsetOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DescribeTestset", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) CreateTest(ctx context.Context, in *CreateTestInput) (*CreateTestOutput, error) {
	out := &CreateTestOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/CreateTest", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) UpdateTest(ctx context.Context, in *UpdateTestInput) (*UpdateTestOutput, error) {
	out := &UpdateTestOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/UpdateTest", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DeleteTest(ctx context.Context, in *DeleteTestInput) (*DeleteTestOutput, error) {
	out := &DeleteTestOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DeleteTest", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) ListTests(ctx context.Context, in *ListTestsInput) (*ListTestsOutput, error) {
	out := &ListTestsOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/ListTests", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DescribeTest(ctx context.Context, in *DescribeTestInput) (*DescribeTestOutput, error) {
	out := &DescribeTestOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DescribeTest", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) GrantPermission(ctx context.Context, in *GrantPermissionInput) (*GrantPermissionOutput, error) {
	out := &GrantPermissionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/GrantPermission", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) RevokePermission(ctx context.Context, in *RevokePermissionInput) (*RevokePermissionOutput, error) {
	out := &RevokePermissionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/RevokePermission", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) ListPermissions(ctx context.Context, in *ListPermissionsInput) (*ListPermissionsOutput, error) {
	out := &ListPermissionsOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/ListPermissions", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) CreateCodeTemplate(ctx context.Context, in *CreateCodeTemplateInput) (*CreateCodeTemplateOutput, error) {
	out := &CreateCodeTemplateOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/CreateCodeTemplate", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) UpdateCodeTemplate(ctx context.Context, in *UpdateCodeTemplateInput) (*UpdateCodeTemplateOutput, error) {
	out := &UpdateCodeTemplateOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/UpdateCodeTemplate", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DeleteCodeTemplate(ctx context.Context, in *DeleteCodeTemplateInput) (*DeleteCodeTemplateOutput, error) {
	out := &DeleteCodeTemplateOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DeleteCodeTemplate", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) ListCodeTemplates(ctx context.Context, in *ListCodeTemplatesInput) (*ListCodeTemplatesOutput, error) {
	out := &ListCodeTemplatesOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/ListCodeTemplates", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error) {
	out := &DescribeCodeTemplateOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DescribeCodeTemplate", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) CreateSolution(ctx context.Context, in *CreateSolutionInput) (*CreateSolutionOutput, error) {
	out := &CreateSolutionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/CreateSolution", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) UpdateSolution(ctx context.Context, in *UpdateSolutionInput) (*UpdateSolutionOutput, error) {
	out := &UpdateSolutionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/UpdateSolution", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DeleteSolution(ctx context.Context, in *DeleteSolutionInput) (*DeleteSolutionOutput, error) {
	out := &DeleteSolutionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DeleteSolution", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) ListSolutions(ctx context.Context, in *ListSolutionsInput) (*ListSolutionsOutput, error) {
	out := &ListSolutionsOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/ListSolutions", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DescribeSolution(ctx context.Context, in *DescribeSolutionInput) (*DescribeSolutionOutput, error) {
	out := &DescribeSolutionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DescribeSolution", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) PublishSolution(ctx context.Context, in *PublishSolutionInput) (*PublishSolutionOutput, error) {
	out := &PublishSolutionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/PublishSolution", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) UnpublishSolution(ctx context.Context, in *UnpublishSolutionInput) (*UnpublishSolutionOutput, error) {
	out := &UnpublishSolutionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/UnpublishSolution", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) ApproveSolution(ctx context.Context, in *ApproveSolutionInput) (*ApproveSolutionOutput, error) {
	out := &ApproveSolutionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/ApproveSolution", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) RefuseSolution(ctx context.Context, in *RefuseSolutionInput) (*RefuseSolutionOutput, error) {
	out := &RefuseSolutionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/RefuseSolution", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) CreateCategory(ctx context.Context, in *CreateCategoryInput) (*CreateCategoryOutput, error) {
	out := &CreateCategoryOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/CreateCategory", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) UpdateCategory(ctx context.Context, in *UpdateCategoryInput) (*UpdateCategoryOutput, error) {
	out := &UpdateCategoryOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/UpdateCategory", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DeleteCategory(ctx context.Context, in *DeleteCategoryInput) (*DeleteCategoryOutput, error) {
	out := &DeleteCategoryOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DeleteCategory", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) ListCategories(ctx context.Context, in *ListCategoriesInput) (*ListCategoriesOutput, error) {
	out := &ListCategoriesOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/ListCategories", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DescribeCategory(ctx context.Context, in *DescribeCategoryInput) (*DescribeCategoryOutput, error) {
	out := &DescribeCategoryOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DescribeCategory", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) AssignCategory(ctx context.Context, in *AssignCategoryInput) (*AssignCategoryOutput, error) {
	out := &AssignCategoryOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/AssignCategory", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) UnassignCategory(ctx context.Context, in *UnassignCategoryInput) (*UnassignCategoryOutput, error) {
	out := &UnassignCategoryOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/UnassignCategory", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) CreateSubmission(ctx context.Context, in *CreateSubmissionInput) (*CreateSubmissionOutput, error) {
	out := &CreateSubmissionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/CreateSubmission", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DescribeSubmission(ctx context.Context, in *DescribeSubmissionInput) (*DescribeSubmissionOutput, error) {
	out := &DescribeSubmissionOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DescribeSubmission", in, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *AtlasService) DescribeScore(ctx context.Context, in *DescribeScoreInput) (*DescribeScoreOutput, error) {
	out := &DescribeScoreOutput{}

	if err := s.invoke(ctx, "eolymp.atlas.Atlas/DescribeScore", in, out); err != nil {
		return nil, err
	}

	return out, nil
}
