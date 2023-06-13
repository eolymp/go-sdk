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
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _DiscussionService_HTTPReadQueryString parses body into proto.Message
func _DiscussionService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _DiscussionService_HTTPReadRequestBody parses body into proto.Message
func _DiscussionService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _DiscussionService_HTTPWriteResponse writes proto.Message to HTTP response
func _DiscussionService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_DiscussionService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _DiscussionService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _DiscussionService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _DiscussionService_WebsocketErrorResponse writes error to websocket connection
func _DiscussionService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _DiscussionService_WebsocketCodec implements protobuf codec for websockets package
var _DiscussionService_WebsocketCodec = websocket.Codec{
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

// RegisterDiscussionServiceHttpHandlers adds handlers for for DiscussionServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterDiscussionServiceHttpHandlers(router *mux.Router, prefix string, cli DiscussionServiceClient) {
	router.Handle(prefix+"/discussions/{discussion_id}", _DiscussionService_DescribeDiscussion_Rule0(cli)).
		Methods("GET").
		Name("eolymp.discussion.DiscussionService.DescribeDiscussion")
	router.Handle(prefix+"/discussions", _DiscussionService_ListDiscussions_Rule0(cli)).
		Methods("GET").
		Name("eolymp.discussion.DiscussionService.ListDiscussions")
	router.Handle(prefix+"/discussions", _DiscussionService_CreateDiscussion_Rule0(cli)).
		Methods("POST").
		Name("eolymp.discussion.DiscussionService.CreateDiscussion")
	router.Handle(prefix+"/discussions/{discussion_id}", _DiscussionService_UpdateDiscussion_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.discussion.DiscussionService.UpdateDiscussion")
	router.Handle(prefix+"/discussions/{discussion_id}", _DiscussionService_DeleteDiscussion_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.discussion.DiscussionService.DeleteDiscussion")
	router.Handle(prefix+"/discussions/{discussion_id}/vote", _DiscussionService_VoteDiscussion_Rule0(cli)).
		Methods("POST").
		Name("eolymp.discussion.DiscussionService.VoteDiscussion")
}

func _DiscussionService_DescribeDiscussion_Rule0(cli DiscussionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeDiscussionInput{}

		if err := _DiscussionService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.DiscussionId = vars["discussion_id"]

		out, err := cli.DescribeDiscussion(r.Context(), in)
		if err != nil {
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_DiscussionService_HTTPWriteResponse(w, out)
	})
}

func _DiscussionService_ListDiscussions_Rule0(cli DiscussionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListDiscussionsInput{}

		if err := _DiscussionService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListDiscussions(r.Context(), in)
		if err != nil {
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_DiscussionService_HTTPWriteResponse(w, out)
	})
}

func _DiscussionService_CreateDiscussion_Rule0(cli DiscussionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateDiscussionInput{}

		if err := _DiscussionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.CreateDiscussion(r.Context(), in)
		if err != nil {
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_DiscussionService_HTTPWriteResponse(w, out)
	})
}

func _DiscussionService_UpdateDiscussion_Rule0(cli DiscussionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateDiscussionInput{}

		if err := _DiscussionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.DiscussionId = vars["discussion_id"]

		out, err := cli.UpdateDiscussion(r.Context(), in)
		if err != nil {
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_DiscussionService_HTTPWriteResponse(w, out)
	})
}

func _DiscussionService_DeleteDiscussion_Rule0(cli DiscussionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteDiscussionInput{}

		if err := _DiscussionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.DiscussionId = vars["discussion_id"]

		out, err := cli.DeleteDiscussion(r.Context(), in)
		if err != nil {
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_DiscussionService_HTTPWriteResponse(w, out)
	})
}

func _DiscussionService_VoteDiscussion_Rule0(cli DiscussionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &VoteDiscussionInput{}

		if err := _DiscussionService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.DiscussionId = vars["discussion_id"]

		out, err := cli.VoteDiscussion(r.Context(), in)
		if err != nil {
			_DiscussionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_DiscussionService_HTTPWriteResponse(w, out)
	})
}

type _DiscussionServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _DiscussionServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _DiscussionServiceHandler) (out proto.Message, err error)
type DiscussionServiceInterceptor struct {
	middleware []_DiscussionServiceMiddleware
	client     DiscussionServiceClient
}

// NewDiscussionServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewDiscussionServiceInterceptor(cli DiscussionServiceClient, middleware ..._DiscussionServiceMiddleware) *DiscussionServiceInterceptor {
	return &DiscussionServiceInterceptor{client: cli, middleware: middleware}
}

func (i *DiscussionServiceInterceptor) DescribeDiscussion(ctx context.Context, in *DescribeDiscussionInput, opts ...grpc.CallOption) (*DescribeDiscussionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeDiscussionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeDiscussionInput, got %T", in))
		}

		return i.client.DescribeDiscussion(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.DiscussionService.DescribeDiscussion", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeDiscussionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeDiscussionOutput, got %T", out))
	}

	return message, err
}

func (i *DiscussionServiceInterceptor) ListDiscussions(ctx context.Context, in *ListDiscussionsInput, opts ...grpc.CallOption) (*ListDiscussionsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListDiscussionsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListDiscussionsInput, got %T", in))
		}

		return i.client.ListDiscussions(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.DiscussionService.ListDiscussions", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListDiscussionsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListDiscussionsOutput, got %T", out))
	}

	return message, err
}

func (i *DiscussionServiceInterceptor) CreateDiscussion(ctx context.Context, in *CreateDiscussionInput, opts ...grpc.CallOption) (*CreateDiscussionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateDiscussionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateDiscussionInput, got %T", in))
		}

		return i.client.CreateDiscussion(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.DiscussionService.CreateDiscussion", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateDiscussionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateDiscussionOutput, got %T", out))
	}

	return message, err
}

func (i *DiscussionServiceInterceptor) UpdateDiscussion(ctx context.Context, in *UpdateDiscussionInput, opts ...grpc.CallOption) (*UpdateDiscussionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateDiscussionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateDiscussionInput, got %T", in))
		}

		return i.client.UpdateDiscussion(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.DiscussionService.UpdateDiscussion", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateDiscussionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateDiscussionOutput, got %T", out))
	}

	return message, err
}

func (i *DiscussionServiceInterceptor) DeleteDiscussion(ctx context.Context, in *DeleteDiscussionInput, opts ...grpc.CallOption) (*DeleteDiscussionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteDiscussionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteDiscussionInput, got %T", in))
		}

		return i.client.DeleteDiscussion(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.DiscussionService.DeleteDiscussion", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteDiscussionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteDiscussionOutput, got %T", out))
	}

	return message, err
}

func (i *DiscussionServiceInterceptor) VoteDiscussion(ctx context.Context, in *VoteDiscussionInput, opts ...grpc.CallOption) (*VoteDiscussionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*VoteDiscussionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *VoteDiscussionInput, got %T", in))
		}

		return i.client.VoteDiscussion(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.DiscussionService.VoteDiscussion", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*VoteDiscussionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *VoteDiscussionOutput, got %T", out))
	}

	return message, err
}
