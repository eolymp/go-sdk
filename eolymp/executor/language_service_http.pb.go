// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package executor

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	websocket "golang.org/x/net/websocket"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _LanguageService_HTTPReadQueryString parses body into proto.Message
func _LanguageService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _LanguageService_HTTPReadRequestBody parses body into proto.Message
func _LanguageService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _LanguageService_HTTPWriteResponse writes proto.Message to HTTP response
func _LanguageService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_LanguageService_HTTPWriteErrorResponse(w, err)
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

// _LanguageService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _LanguageService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _LanguageService_WebsocketErrorResponse writes error to websocket connection
func _LanguageService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
	switch status.Convert(e).Code() {
	case codes.OK:
		conn.WriteClose(1000)
	case codes.Canceled:
		conn.WriteClose(1000)
	case codes.Unknown:
		conn.WriteClose(1011)
	case codes.InvalidArgument:
		conn.WriteClose(1003)
	case codes.DeadlineExceeded:
		conn.WriteClose(1000)
	case codes.NotFound:
		conn.WriteClose(1000)
	case codes.AlreadyExists:
		conn.WriteClose(1000)
	case codes.PermissionDenied:
		conn.WriteClose(1000)
	case codes.ResourceExhausted:
		conn.WriteClose(1000)
	case codes.FailedPrecondition:
		conn.WriteClose(1000)
	case codes.Aborted:
		conn.WriteClose(1000)
	case codes.OutOfRange:
		conn.WriteClose(1000)
	case codes.Unimplemented:
		conn.WriteClose(1011)
	case codes.Internal:
		conn.WriteClose(1011)
	case codes.Unavailable:
		conn.WriteClose(1011)
	case codes.DataLoss:
		conn.WriteClose(1011)
	case codes.Unauthenticated:
		conn.WriteClose(1000)
	default:
		conn.WriteClose(1000)
	}
}

// _LanguageService_WebsocketCodec implements protobuf codec for websockets package
var _LanguageService_WebsocketCodec = websocket.Codec{
	Marshal: func(v interface{}) ([]byte, byte, error) {
		m, ok := v.(proto.Message)
		if !ok {
			panic(fmt.Errorf("invalid message type %T", v))
		}

		d, err := protojson.Marshal(m)
		if err != nil {
			return nil, 0, err
		}

		return d, websocket.TextFrame, err
	},
	Unmarshal: func(d []byte, t byte, v interface{}) error {
		m, ok := v.(proto.Message)
		if !ok {
			panic(fmt.Errorf("invalid message type %T", v))
		}

		return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(d, m)
	},
}

// RegisterLanguageServiceHttpHandlers adds handlers for for LanguageServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterLanguageServiceHttpHandlers(router *mux.Router, prefix string, cli LanguageServiceClient) {
	router.Handle(prefix+"/exec/languages/{language_id}", _LanguageService_DescribeLanguage_Rule0(cli)).
		Methods("GET").
		Name("eolymp.executor.LanguageService.DescribeLanguage")
	router.Handle(prefix+"/exec/languages", _LanguageService_ListLanguages_Rule0(cli)).
		Methods("GET").
		Name("eolymp.executor.LanguageService.ListLanguages")
	router.Handle(prefix+"/exec/runtime/{runtime_id}", _LanguageService_DescribeRuntime_Rule0(cli)).
		Methods("GET").
		Name("eolymp.executor.LanguageService.DescribeRuntime")
	router.Handle(prefix+"/exec/runtime", _LanguageService_ListRuntime_Rule0(cli)).
		Methods("GET").
		Name("eolymp.executor.LanguageService.ListRuntime")
	router.Handle(prefix+"/exec/runtime/{runtime_id}/template", _LanguageService_DescribeCodeTemplate_Rule0(cli)).
		Methods("GET").
		Name("eolymp.executor.LanguageService.DescribeCodeTemplate")
}

func _LanguageService_DescribeLanguage_Rule0(cli LanguageServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeLanguageInput{}

		if err := _LanguageService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_LanguageService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.LanguageId = vars["language_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeLanguage(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_LanguageService_HTTPWriteErrorResponse(w, err)
			return
		}

		_LanguageService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _LanguageService_ListLanguages_Rule0(cli LanguageServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListLanguagesInput{}

		if err := _LanguageService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_LanguageService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListLanguages(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_LanguageService_HTTPWriteErrorResponse(w, err)
			return
		}

		_LanguageService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _LanguageService_DescribeRuntime_Rule0(cli LanguageServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRuntimeInput{}

		if err := _LanguageService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_LanguageService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RuntimeId = vars["runtime_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeRuntime(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_LanguageService_HTTPWriteErrorResponse(w, err)
			return
		}

		_LanguageService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _LanguageService_ListRuntime_Rule0(cli LanguageServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListRuntimeInput{}

		if err := _LanguageService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_LanguageService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListRuntime(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_LanguageService_HTTPWriteErrorResponse(w, err)
			return
		}

		_LanguageService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _LanguageService_DescribeCodeTemplate_Rule0(cli LanguageServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCodeTemplateInput{}

		if err := _LanguageService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_LanguageService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RuntimeId = vars["runtime_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeCodeTemplate(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_LanguageService_HTTPWriteErrorResponse(w, err)
			return
		}

		_LanguageService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _LanguageServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _LanguageServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _LanguageServiceHandler) (out proto.Message, err error)
type LanguageServiceInterceptor struct {
	middleware []_LanguageServiceMiddleware
	client     LanguageServiceClient
}

// NewLanguageServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewLanguageServiceInterceptor(cli LanguageServiceClient, middleware ..._LanguageServiceMiddleware) *LanguageServiceInterceptor {
	return &LanguageServiceInterceptor{client: cli, middleware: middleware}
}

func (i *LanguageServiceInterceptor) DescribeLanguage(ctx context.Context, in *DescribeLanguageInput, opts ...grpc.CallOption) (*DescribeLanguageOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeLanguageInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeLanguageInput, got %T", in))
		}

		return i.client.DescribeLanguage(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.executor.LanguageService.DescribeLanguage", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeLanguageOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeLanguageOutput, got %T", out))
	}

	return message, err
}

func (i *LanguageServiceInterceptor) ListLanguages(ctx context.Context, in *ListLanguagesInput, opts ...grpc.CallOption) (*ListLanguagesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListLanguagesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListLanguagesInput, got %T", in))
		}

		return i.client.ListLanguages(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.executor.LanguageService.ListLanguages", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListLanguagesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListLanguagesOutput, got %T", out))
	}

	return message, err
}

func (i *LanguageServiceInterceptor) DescribeRuntime(ctx context.Context, in *DescribeRuntimeInput, opts ...grpc.CallOption) (*DescribeRuntimeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeRuntimeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeRuntimeInput, got %T", in))
		}

		return i.client.DescribeRuntime(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.executor.LanguageService.DescribeRuntime", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeRuntimeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeRuntimeOutput, got %T", out))
	}

	return message, err
}

func (i *LanguageServiceInterceptor) ListRuntime(ctx context.Context, in *ListRuntimeInput, opts ...grpc.CallOption) (*ListRuntimeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListRuntimeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListRuntimeInput, got %T", in))
		}

		return i.client.ListRuntime(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.executor.LanguageService.ListRuntime", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListRuntimeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListRuntimeOutput, got %T", out))
	}

	return message, err
}

func (i *LanguageServiceInterceptor) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput, opts ...grpc.CallOption) (*DescribeCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeCodeTemplateInput, got %T", in))
		}

		return i.client.DescribeCodeTemplate(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.executor.LanguageService.DescribeCodeTemplate", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeCodeTemplateOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeCodeTemplateOutput, got %T", out))
	}

	return message, err
}