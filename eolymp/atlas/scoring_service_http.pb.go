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

// _ScoringService_HTTPReadQueryString parses body into proto.Message
func _ScoringService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _ScoringService_HTTPReadRequestBody parses body into proto.Message
func _ScoringService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _ScoringService_HTTPWriteResponse writes proto.Message to HTTP response
func _ScoringService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_ScoringService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _ScoringService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _ScoringService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterScoringServiceHttpHandlers adds handlers for for ScoringServiceServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterScoringServiceHttpHandlers(router *mux.Router, prefix string, srv ScoringServiceServer) {
	router.Handle(prefix+"/scores/{user_id}", _ScoringService_DescribeScore_Rule0(srv)).
		Methods("GET").
		Name("eolymp.atlas.ScoringService.DescribeScore")
	router.Handle(prefix+"/grading", _ScoringService_DescribeProblemGrading_Rule0(srv)).
		Methods("GET").
		Name("eolymp.atlas.ScoringService.DescribeProblemGrading")
	router.Handle(prefix+"/top", _ScoringService_ListProblemTop_Rule0(srv)).
		Methods("GET").
		Name("eolymp.atlas.ScoringService.ListProblemTop")
}

func _ScoringService_DescribeScore_Rule0(srv ScoringServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeScoreInput{}

		if err := _ScoringService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_ScoringService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.UserId = vars["user_id"]

		out, err := srv.DescribeScore(r.Context(), in)
		if err != nil {
			_ScoringService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoringService_HTTPWriteResponse(w, out)
	})
}

func _ScoringService_DescribeProblemGrading_Rule0(srv ScoringServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeProblemGradingInput{}

		if err := _ScoringService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_ScoringService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeProblemGrading(r.Context(), in)
		if err != nil {
			_ScoringService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoringService_HTTPWriteResponse(w, out)
	})
}

func _ScoringService_ListProblemTop_Rule0(srv ScoringServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListProblemTopInput{}

		if err := _ScoringService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_ScoringService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListProblemTop(r.Context(), in)
		if err != nil {
			_ScoringService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoringService_HTTPWriteResponse(w, out)
	})
}

type _ScoringServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _ScoringServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _ScoringServiceHandler) (out proto.Message, err error)
type ScoringServiceInterceptor struct {
	middleware []_ScoringServiceMiddleware
	server     ScoringServiceServer
}

// NewScoringServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewScoringServiceInterceptor(srv ScoringServiceServer, middleware ..._ScoringServiceMiddleware) *ScoringServiceInterceptor {
	return &ScoringServiceInterceptor{server: srv, middleware: middleware}
}

func (i *ScoringServiceInterceptor) DescribeScore(ctx context.Context, in *DescribeScoreInput) (*DescribeScoreOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeScoreInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeScoreInput, got %T", in))
		}

		return i.server.DescribeScore(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ScoringService.DescribeScore", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeScoreOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeScoreOutput, got %T", out))
	}

	return message, err
}

func (i *ScoringServiceInterceptor) DescribeProblemGrading(ctx context.Context, in *DescribeProblemGradingInput) (*DescribeProblemGradingOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeProblemGradingInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeProblemGradingInput, got %T", in))
		}

		return i.server.DescribeProblemGrading(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ScoringService.DescribeProblemGrading", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeProblemGradingOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeProblemGradingOutput, got %T", out))
	}

	return message, err
}

func (i *ScoringServiceInterceptor) ListProblemTop(ctx context.Context, in *ListProblemTopInput) (*ListProblemTopOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListProblemTopInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListProblemTopInput, got %T", in))
		}

		return i.server.ListProblemTop(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ScoringService.ListProblemTop", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListProblemTopOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListProblemTopOutput, got %T", out))
	}

	return message, err
}