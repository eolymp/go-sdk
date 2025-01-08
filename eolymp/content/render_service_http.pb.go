// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package content

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

// _RenderService_HTTPReadQueryString parses body into proto.Message
func _RenderService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _RenderService_HTTPReadRequestBody parses body into proto.Message
func _RenderService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _RenderService_HTTPWriteResponse writes proto.Message to HTTP response
func _RenderService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_RenderService_HTTPWriteErrorResponse(w, err)
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

// _RenderService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _RenderService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterRenderServiceHttpHandlers adds handlers for for RenderServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterRenderServiceHttpHandlers(router *mux.Router, prefix string, cli RenderServiceClient) {
	router.Handle(prefix+"/renderer", _RenderService_RenderContent_Rule0(cli)).
		Methods("POST").
		Name("eolymp.content.RenderService.RenderContent")
}

func _RenderService_RenderContent_Rule0(cli RenderServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RenderContentInput{}

		if err := _RenderService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RenderService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.RenderContent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RenderService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RenderService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _RenderServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _RenderServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _RenderServiceHandler) (out proto.Message, err error)
type RenderServiceInterceptor struct {
	middleware []_RenderServiceMiddleware
	client     RenderServiceClient
}

// NewRenderServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewRenderServiceInterceptor(cli RenderServiceClient, middleware ..._RenderServiceMiddleware) *RenderServiceInterceptor {
	return &RenderServiceInterceptor{client: cli, middleware: middleware}
}

func (i *RenderServiceInterceptor) RenderContent(ctx context.Context, in *RenderContentInput, opts ...grpc.CallOption) (*RenderContentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RenderContentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RenderContentInput, got %T", in))
		}

		return i.client.RenderContent(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.content.RenderService.RenderContent", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RenderContentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RenderContentOutput, got %T", out))
	}

	return message, err
}
