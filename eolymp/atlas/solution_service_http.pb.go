// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package atlas

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _SolutionService_HTTPReadQueryString parses body into proto.Message
func _SolutionService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _SolutionService_HTTPReadRequestBody parses body into proto.Message
func _SolutionService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _SolutionService_HTTPWriteResponse writes proto.Message to HTTP response
func _SolutionService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_SolutionService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if v := append(h.Get("cache-control"), t.Get("cache-control")...); len(v) > 0 {
		w.Header().Set("Cache-Control", v[len(v)-1])
	}

	if v := append(h.Get("etag"), t.Get("etag")...); len(v) > 0 {
		w.Header().Set("ETag", v[len(v)-1])
	}

	if v := append(h.Get("last-modified"), t.Get("last-modified")...); len(v) > 0 {
		w.Header().Set("Last-Modified", v[len(v)-1])
	}

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _SolutionService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _SolutionService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterSolutionServiceHttpHandlers adds handlers for for SolutionServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterSolutionServiceHttpHandlers(router *mux.Router, prefix string, cli SolutionServiceClient) {
	router.Handle(prefix+"/solutions", _SolutionService_CreateSolution_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.SolutionService.CreateSolution")
	router.Handle(prefix+"/solutions/{solution_id}", _SolutionService_UpdateSolution_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.SolutionService.UpdateSolution")
	router.Handle(prefix+"/solutions/{solution_id}", _SolutionService_DeleteSolution_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.atlas.SolutionService.DeleteSolution")
	router.Handle(prefix+"/solutions/{solution_id}", _SolutionService_DescribeSolution_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.SolutionService.DescribeSolution")
	router.Handle(prefix+"/solutions", _SolutionService_ListSolutions_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.SolutionService.ListSolutions")
}

func _SolutionService_CreateSolution_Rule0(cli SolutionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateSolutionInput{}

		if err := _SolutionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SolutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateSolution(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SolutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SolutionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _SolutionService_UpdateSolution_Rule0(cli SolutionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateSolutionInput{}

		if err := _SolutionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SolutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SolutionId = vars["solution_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateSolution(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SolutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SolutionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _SolutionService_DeleteSolution_Rule0(cli SolutionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteSolutionInput{}

		if err := _SolutionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SolutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SolutionId = vars["solution_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteSolution(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SolutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SolutionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _SolutionService_DescribeSolution_Rule0(cli SolutionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeSolutionInput{}

		if err := _SolutionService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SolutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SolutionId = vars["solution_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeSolution(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SolutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SolutionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _SolutionService_ListSolutions_Rule0(cli SolutionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListSolutionsInput{}

		if err := _SolutionService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_SolutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListSolutions(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_SolutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_SolutionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _SolutionServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _SolutionServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _SolutionServiceHandler) (out proto.Message, err error)
type SolutionServiceInterceptor struct {
	middleware []_SolutionServiceMiddleware
	client     SolutionServiceClient
}

// NewSolutionServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewSolutionServiceInterceptor(cli SolutionServiceClient, middleware ..._SolutionServiceMiddleware) *SolutionServiceInterceptor {
	return &SolutionServiceInterceptor{client: cli, middleware: middleware}
}

func (i *SolutionServiceInterceptor) CreateSolution(ctx context.Context, in *CreateSolutionInput, opts ...grpc.CallOption) (*CreateSolutionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateSolutionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateSolutionInput, got %T", in))
		}

		return i.client.CreateSolution(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.SolutionService.CreateSolution", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateSolutionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateSolutionOutput, got %T", out))
	}

	return message, err
}

func (i *SolutionServiceInterceptor) UpdateSolution(ctx context.Context, in *UpdateSolutionInput, opts ...grpc.CallOption) (*UpdateSolutionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateSolutionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateSolutionInput, got %T", in))
		}

		return i.client.UpdateSolution(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.SolutionService.UpdateSolution", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateSolutionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateSolutionOutput, got %T", out))
	}

	return message, err
}

func (i *SolutionServiceInterceptor) DeleteSolution(ctx context.Context, in *DeleteSolutionInput, opts ...grpc.CallOption) (*DeleteSolutionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteSolutionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteSolutionInput, got %T", in))
		}

		return i.client.DeleteSolution(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.SolutionService.DeleteSolution", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteSolutionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteSolutionOutput, got %T", out))
	}

	return message, err
}

func (i *SolutionServiceInterceptor) DescribeSolution(ctx context.Context, in *DescribeSolutionInput, opts ...grpc.CallOption) (*DescribeSolutionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeSolutionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeSolutionInput, got %T", in))
		}

		return i.client.DescribeSolution(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.SolutionService.DescribeSolution", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeSolutionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeSolutionOutput, got %T", out))
	}

	return message, err
}

func (i *SolutionServiceInterceptor) ListSolutions(ctx context.Context, in *ListSolutionsInput, opts ...grpc.CallOption) (*ListSolutionsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListSolutionsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListSolutionsInput, got %T", in))
		}

		return i.client.ListSolutions(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.SolutionService.ListSolutions", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListSolutionsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListSolutionsOutput, got %T", out))
	}

	return message, err
}
