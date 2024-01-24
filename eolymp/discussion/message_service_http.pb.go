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

// _MessageService_HTTPReadQueryString parses body into proto.Message
func _MessageService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _MessageService_HTTPReadRequestBody parses body into proto.Message
func _MessageService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _MessageService_HTTPWriteResponse writes proto.Message to HTTP response
func _MessageService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_MessageService_HTTPWriteErrorResponse(w, err)
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

// _MessageService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _MessageService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _MessageService_WebsocketErrorResponse writes error to websocket connection
func _MessageService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _MessageService_WebsocketCodec implements protobuf codec for websockets package
var _MessageService_WebsocketCodec = websocket.Codec{
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

// RegisterMessageServiceHttpHandlers adds handlers for for MessageServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterMessageServiceHttpHandlers(router *mux.Router, prefix string, cli MessageServiceClient) {
	router.Handle(prefix+"/messages/{message_id}", _MessageService_DescribeMessage_Rule0(cli)).
		Methods("GET").
		Name("eolymp.discussion.MessageService.DescribeMessage")
	router.Handle(prefix+"/messages", _MessageService_ListMessages_Rule0(cli)).
		Methods("GET").
		Name("eolymp.discussion.MessageService.ListMessages")
	router.Handle(prefix+"/messages", _MessageService_PostMessage_Rule0(cli)).
		Methods("POST").
		Name("eolymp.discussion.MessageService.PostMessage")
	router.Handle(prefix+"/messages/{message_id}", _MessageService_UpdateMessage_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.discussion.MessageService.UpdateMessage")
	router.Handle(prefix+"/messages/{message_id}", _MessageService_DeleteMessage_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.discussion.MessageService.DeleteMessage")
	router.Handle(prefix+"/messages/{message_id}/vote", _MessageService_VoteMessage_Rule0(cli)).
		Methods("POST").
		Name("eolymp.discussion.MessageService.VoteMessage")
}

func _MessageService_DescribeMessage_Rule0(cli MessageServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeMessageInput{}

		if err := _MessageService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MessageId = vars["message_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeMessage(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MessageService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MessageService_ListMessages_Rule0(cli MessageServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListMessagesInput{}

		if err := _MessageService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListMessages(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MessageService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MessageService_PostMessage_Rule0(cli MessageServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &PostMessageInput{}

		if err := _MessageService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.PostMessage(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MessageService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MessageService_UpdateMessage_Rule0(cli MessageServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateMessageInput{}

		if err := _MessageService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MessageId = vars["message_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateMessage(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MessageService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MessageService_DeleteMessage_Rule0(cli MessageServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteMessageInput{}

		if err := _MessageService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MessageId = vars["message_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteMessage(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MessageService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MessageService_VoteMessage_Rule0(cli MessageServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &VoteMessageInput{}

		if err := _MessageService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MessageId = vars["message_id"]

		var header, trailer metadata.MD

		out, err := cli.VoteMessage(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MessageService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MessageService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _MessageServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _MessageServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _MessageServiceHandler) (out proto.Message, err error)
type MessageServiceInterceptor struct {
	middleware []_MessageServiceMiddleware
	client     MessageServiceClient
}

// NewMessageServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewMessageServiceInterceptor(cli MessageServiceClient, middleware ..._MessageServiceMiddleware) *MessageServiceInterceptor {
	return &MessageServiceInterceptor{client: cli, middleware: middleware}
}

func (i *MessageServiceInterceptor) DescribeMessage(ctx context.Context, in *DescribeMessageInput, opts ...grpc.CallOption) (*DescribeMessageOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeMessageInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeMessageInput, got %T", in))
		}

		return i.client.DescribeMessage(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.MessageService.DescribeMessage", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeMessageOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeMessageOutput, got %T", out))
	}

	return message, err
}

func (i *MessageServiceInterceptor) ListMessages(ctx context.Context, in *ListMessagesInput, opts ...grpc.CallOption) (*ListMessagesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListMessagesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListMessagesInput, got %T", in))
		}

		return i.client.ListMessages(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.MessageService.ListMessages", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListMessagesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListMessagesOutput, got %T", out))
	}

	return message, err
}

func (i *MessageServiceInterceptor) PostMessage(ctx context.Context, in *PostMessageInput, opts ...grpc.CallOption) (*PostMessageOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*PostMessageInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *PostMessageInput, got %T", in))
		}

		return i.client.PostMessage(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.MessageService.PostMessage", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*PostMessageOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *PostMessageOutput, got %T", out))
	}

	return message, err
}

func (i *MessageServiceInterceptor) UpdateMessage(ctx context.Context, in *UpdateMessageInput, opts ...grpc.CallOption) (*UpdateMessageOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateMessageInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateMessageInput, got %T", in))
		}

		return i.client.UpdateMessage(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.MessageService.UpdateMessage", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateMessageOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateMessageOutput, got %T", out))
	}

	return message, err
}

func (i *MessageServiceInterceptor) DeleteMessage(ctx context.Context, in *DeleteMessageInput, opts ...grpc.CallOption) (*DeleteMessageOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteMessageInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteMessageInput, got %T", in))
		}

		return i.client.DeleteMessage(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.MessageService.DeleteMessage", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteMessageOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteMessageOutput, got %T", out))
	}

	return message, err
}

func (i *MessageServiceInterceptor) VoteMessage(ctx context.Context, in *VoteMessageInput, opts ...grpc.CallOption) (*VoteMessageOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*VoteMessageInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *VoteMessageInput, got %T", in))
		}

		return i.client.VoteMessage(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.MessageService.VoteMessage", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*VoteMessageOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *VoteMessageOutput, got %T", out))
	}

	return message, err
}
