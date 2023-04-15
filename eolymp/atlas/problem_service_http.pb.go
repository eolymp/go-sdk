// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package atlas

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _ProblemService_HTTPReadQueryString parses body into proto.Message
func _ProblemService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _ProblemService_HTTPReadRequestBody parses body into proto.Message
func _ProblemService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _ProblemService_HTTPWriteResponse writes proto.Message to HTTP response
func _ProblemService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_ProblemService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _ProblemService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _ProblemService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
	s := status.Convert(e)

	w.Header().Set("Content-Type", "application/json")

	switch s.Code() {
	case codes.OK:
		w.WriteHeader(http.StatusOK)
	case codes.Canceled:
		w.WriteHeader(http.StatusGatewayTimeout)
	case codes.Unknown:
		w.WriteHeader(http.StatusInternalServerError)
	case codes.InvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	case codes.DeadlineExceeded:
		w.WriteHeader(http.StatusGatewayTimeout)
	case codes.NotFound:
		w.WriteHeader(http.StatusNotFound)
	case codes.AlreadyExists:
		w.WriteHeader(http.StatusConflict)
	case codes.PermissionDenied:
		w.WriteHeader(http.StatusForbidden)
	case codes.ResourceExhausted:
		w.WriteHeader(http.StatusTooManyRequests)
	case codes.FailedPrecondition:
		w.WriteHeader(http.StatusPreconditionFailed)
	case codes.Aborted:
		w.WriteHeader(http.StatusServiceUnavailable)
	case codes.OutOfRange:
		w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
	case codes.Unimplemented:
		w.WriteHeader(http.StatusNotImplemented)
	case codes.Internal:
		w.WriteHeader(http.StatusInternalServerError)
	case codes.Unavailable:
		w.WriteHeader(http.StatusServiceUnavailable)
	case codes.DataLoss:
		w.WriteHeader(http.StatusInternalServerError)
	case codes.Unauthenticated:
		w.WriteHeader(http.StatusUnauthorized)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	data, err := protojson.Marshal(s.Proto())
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(data)
}

// RegisterProblemServiceHttpHandlers adds handlers for for ProblemServiceServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterProblemServiceHttpHandlers(router *mux.Router, prefix string, srv ProblemServiceServer) {
	router.Handle(prefix+"/problems", _ProblemService_CreateProblem_Rule0(srv)).
		Methods("POST").
		Name("eolymp.atlas.ProblemService.CreateProblem")
	router.Handle(prefix+"/problems/{problem_id}", _ProblemService_DeleteProblem_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.atlas.ProblemService.DeleteProblem")
	router.Handle(prefix+"/problems", _ProblemService_ListProblems_Rule0(srv)).
		Methods("GET").
		Name("eolymp.atlas.ProblemService.ListProblems")
	router.Handle(prefix+"/problems/{problem_id}", _ProblemService_DescribeProblem_Rule0(srv)).
		Methods("GET").
		Name("eolymp.atlas.ProblemService.DescribeProblem")
	router.Handle(prefix+"/problems/{problem_id}/visibility", _ProblemService_UpdateVisibility_Rule0(srv)).
		Methods("POST").
		Name("eolymp.atlas.ProblemService.UpdateVisibility")
	router.Handle(prefix+"/problems/{problem_id}/privacy", _ProblemService_UpdatePrivacy_Rule0(srv)).
		Methods("POST").
		Name("eolymp.atlas.ProblemService.UpdatePrivacy")
	router.Handle(prefix+"/problems/{problem_id}/versions", _ProblemService_ListVersions_Rule0(srv)).
		Methods("GET").
		Name("eolymp.atlas.ProblemService.ListVersions")
}

func _ProblemService_CreateProblem_Rule0(srv ProblemServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateProblemInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateProblem(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_DeleteProblem_Rule0(srv ProblemServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteProblemInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		out, err := srv.DeleteProblem(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_ListProblems_Rule0(srv ProblemServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListProblemsInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListProblems(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_DescribeProblem_Rule0(srv ProblemServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeProblemInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		out, err := srv.DescribeProblem(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_UpdateVisibility_Rule0(srv ProblemServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateVisibilityInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		out, err := srv.UpdateVisibility(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_UpdatePrivacy_Rule0(srv ProblemServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdatePrivacyInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		out, err := srv.UpdatePrivacy(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_ListVersions_Rule0(srv ProblemServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListVersionsInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		out, err := srv.ListVersions(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

type _ProblemServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _ProblemServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _ProblemServiceHandler) (out proto.Message, err error)
type ProblemServiceInterceptor struct {
	middleware []_ProblemServiceMiddleware
	server     ProblemServiceServer
}

// NewProblemServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewProblemServiceInterceptor(srv ProblemServiceServer, middleware ..._ProblemServiceMiddleware) *ProblemServiceInterceptor {
	return &ProblemServiceInterceptor{server: srv, middleware: middleware}
}

func (i *ProblemServiceInterceptor) CreateProblem(ctx context.Context, in *CreateProblemInput) (*CreateProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateProblemInput, got %T", in))
		}

		return i.server.CreateProblem(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.CreateProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) DeleteProblem(ctx context.Context, in *DeleteProblemInput) (*DeleteProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteProblemInput, got %T", in))
		}

		return i.server.DeleteProblem(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.DeleteProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) ListProblems(ctx context.Context, in *ListProblemsInput) (*ListProblemsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListProblemsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListProblemsInput, got %T", in))
		}

		return i.server.ListProblems(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.ListProblems", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListProblemsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListProblemsOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) DescribeProblem(ctx context.Context, in *DescribeProblemInput) (*DescribeProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeProblemInput, got %T", in))
		}

		return i.server.DescribeProblem(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.DescribeProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) UpdateVisibility(ctx context.Context, in *UpdateVisibilityInput) (*UpdateVisibilityOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateVisibilityInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateVisibilityInput, got %T", in))
		}

		return i.server.UpdateVisibility(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.UpdateVisibility", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateVisibilityOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateVisibilityOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) UpdatePrivacy(ctx context.Context, in *UpdatePrivacyInput) (*UpdatePrivacyOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdatePrivacyInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdatePrivacyInput, got %T", in))
		}

		return i.server.UpdatePrivacy(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.UpdatePrivacy", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdatePrivacyOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdatePrivacyOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) ListVersions(ctx context.Context, in *ListVersionsInput) (*ListVersionsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListVersionsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListVersionsInput, got %T", in))
		}

		return i.server.ListVersions(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.ListVersions", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListVersionsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListVersionsOutput, got %T", out))
	}

	return message, err
}
