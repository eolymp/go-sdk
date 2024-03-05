// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package discussion

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

// _PostTypeService_HTTPReadQueryString parses body into proto.Message
func _PostTypeService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _PostTypeService_HTTPReadRequestBody parses body into proto.Message
func _PostTypeService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _PostTypeService_HTTPWriteResponse writes proto.Message to HTTP response
func _PostTypeService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_PostTypeService_HTTPWriteErrorResponse(w, err)
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

// _PostTypeService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _PostTypeService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _PostTypeService_WebsocketErrorResponse writes error to websocket connection
func _PostTypeService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _PostTypeService_WebsocketCodec implements protobuf codec for websockets package
var _PostTypeService_WebsocketCodec = websocket.Codec{
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

// RegisterPostTypeServiceHttpHandlers adds handlers for for PostTypeServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterPostTypeServiceHttpHandlers(router *mux.Router, prefix string, cli PostTypeServiceClient) {
	router.Handle(prefix+"/post-types/{type_id}", _PostTypeService_DescribePostType_Rule0(cli)).
		Methods("GET").
		Name("eolymp.discussion.PostTypeService.DescribePostType")
	router.Handle(prefix+"/post-types", _PostTypeService_ListPostTypes_Rule0(cli)).
		Methods("GET").
		Name("eolymp.discussion.PostTypeService.ListPostTypes")
	router.Handle(prefix+"/post-types", _PostTypeService_CreatePostType_Rule0(cli)).
		Methods("POST").
		Name("eolymp.discussion.PostTypeService.CreatePostType")
	router.Handle(prefix+"/post-types/{type_id}", _PostTypeService_UpdatePostType_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.discussion.PostTypeService.UpdatePostType")
	router.Handle(prefix+"/post-types/{type_id}", _PostTypeService_DeletePostType_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.discussion.PostTypeService.DeletePostType")
}

func _PostTypeService_DescribePostType_Rule0(cli PostTypeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePostTypeInput{}

		if err := _PostTypeService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostTypeService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TypeId = vars["type_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribePostType(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostTypeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostTypeService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostTypeService_ListPostTypes_Rule0(cli PostTypeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPostTypesInput{}

		if err := _PostTypeService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostTypeService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListPostTypes(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostTypeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostTypeService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostTypeService_CreatePostType_Rule0(cli PostTypeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreatePostTypeInput{}

		if err := _PostTypeService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostTypeService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreatePostType(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostTypeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostTypeService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostTypeService_UpdatePostType_Rule0(cli PostTypeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdatePostTypeInput{}

		if err := _PostTypeService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostTypeService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TypeId = vars["type_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdatePostType(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostTypeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostTypeService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostTypeService_DeletePostType_Rule0(cli PostTypeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeletePostTypeInput{}

		if err := _PostTypeService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostTypeService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TypeId = vars["type_id"]

		var header, trailer metadata.MD

		out, err := cli.DeletePostType(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostTypeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostTypeService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _PostTypeServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _PostTypeServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _PostTypeServiceHandler) (out proto.Message, err error)
type PostTypeServiceInterceptor struct {
	middleware []_PostTypeServiceMiddleware
	client     PostTypeServiceClient
}

// NewPostTypeServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewPostTypeServiceInterceptor(cli PostTypeServiceClient, middleware ..._PostTypeServiceMiddleware) *PostTypeServiceInterceptor {
	return &PostTypeServiceInterceptor{client: cli, middleware: middleware}
}

func (i *PostTypeServiceInterceptor) DescribePostType(ctx context.Context, in *DescribePostTypeInput, opts ...grpc.CallOption) (*DescribePostTypeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribePostTypeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribePostTypeInput, got %T", in))
		}

		return i.client.DescribePostType(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.PostTypeService.DescribePostType", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribePostTypeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribePostTypeOutput, got %T", out))
	}

	return message, err
}

func (i *PostTypeServiceInterceptor) ListPostTypes(ctx context.Context, in *ListPostTypesInput, opts ...grpc.CallOption) (*ListPostTypesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListPostTypesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListPostTypesInput, got %T", in))
		}

		return i.client.ListPostTypes(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.PostTypeService.ListPostTypes", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListPostTypesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListPostTypesOutput, got %T", out))
	}

	return message, err
}

func (i *PostTypeServiceInterceptor) CreatePostType(ctx context.Context, in *CreatePostTypeInput, opts ...grpc.CallOption) (*CreatePostTypeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreatePostTypeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreatePostTypeInput, got %T", in))
		}

		return i.client.CreatePostType(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.PostTypeService.CreatePostType", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreatePostTypeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreatePostTypeOutput, got %T", out))
	}

	return message, err
}

func (i *PostTypeServiceInterceptor) UpdatePostType(ctx context.Context, in *UpdatePostTypeInput, opts ...grpc.CallOption) (*UpdatePostTypeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdatePostTypeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdatePostTypeInput, got %T", in))
		}

		return i.client.UpdatePostType(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.PostTypeService.UpdatePostType", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdatePostTypeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdatePostTypeOutput, got %T", out))
	}

	return message, err
}

func (i *PostTypeServiceInterceptor) DeletePostType(ctx context.Context, in *DeletePostTypeInput, opts ...grpc.CallOption) (*DeletePostTypeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeletePostTypeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeletePostTypeInput, got %T", in))
		}

		return i.client.DeletePostType(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.PostTypeService.DeletePostType", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeletePostTypeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeletePostTypeOutput, got %T", out))
	}

	return message, err
}