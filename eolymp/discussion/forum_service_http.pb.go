// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package discussion

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	websocket "golang.org/x/net/websocket"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _ForumService_HTTPReadQueryString parses body into proto.Message
func _ForumService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _ForumService_HTTPReadRequestBody parses body into proto.Message
func _ForumService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _ForumService_HTTPWriteResponse writes proto.Message to HTTP response
func _ForumService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_ForumService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _ForumService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _ForumService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _ForumService_WebsocketErrorResponse writes error to websocket connection
func _ForumService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _ForumService_WebsocketCodec implements protobuf codec for websockets package
var _ForumService_WebsocketCodec = websocket.Codec{
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

// RegisterForumServiceHttpHandlers adds handlers for for ForumServiceServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterForumServiceHttpHandlers(router *mux.Router, prefix string, srv ForumServiceServer) {
	router.Handle(prefix+"/forums/{forum_id}", _ForumService_DescribeForum_Rule0(srv)).
		Methods("GET").
		Name("eolymp.discussion.ForumService.DescribeForum")
	router.Handle(prefix+"/forums", _ForumService_ListForums_Rule0(srv)).
		Methods("GET").
		Name("eolymp.discussion.ForumService.ListForums")
	router.Handle(prefix+"/forums", _ForumService_CreateForum_Rule0(srv)).
		Methods("POST").
		Name("eolymp.discussion.ForumService.CreateForum")
	router.Handle(prefix+"/forums/{forum_id}", _ForumService_UpdateForum_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.discussion.ForumService.UpdateForum")
	router.Handle(prefix+"/forums/{forum_id}", _ForumService_DeleteForum_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.discussion.ForumService.DeleteForum")
}

func _ForumService_DescribeForum_Rule0(srv ForumServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeForumInput{}

		if err := _ForumService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ForumService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ForumId = vars["forum_id"]

		out, err := srv.DescribeForum(r.Context(), in)
		if err != nil {
			_ForumService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ForumService_HTTPWriteResponse(w, out)
	})
}

func _ForumService_ListForums_Rule0(srv ForumServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListForumsInput{}

		if err := _ForumService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ForumService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListForums(r.Context(), in)
		if err != nil {
			_ForumService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ForumService_HTTPWriteResponse(w, out)
	})
}

func _ForumService_CreateForum_Rule0(srv ForumServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateForumInput{}

		if err := _ForumService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ForumService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateForum(r.Context(), in)
		if err != nil {
			_ForumService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ForumService_HTTPWriteResponse(w, out)
	})
}

func _ForumService_UpdateForum_Rule0(srv ForumServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateForumInput{}

		if err := _ForumService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ForumService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ForumId = vars["forum_id"]

		out, err := srv.UpdateForum(r.Context(), in)
		if err != nil {
			_ForumService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ForumService_HTTPWriteResponse(w, out)
	})
}

func _ForumService_DeleteForum_Rule0(srv ForumServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteForumInput{}

		if err := _ForumService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ForumService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ForumId = vars["forum_id"]

		out, err := srv.DeleteForum(r.Context(), in)
		if err != nil {
			_ForumService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ForumService_HTTPWriteResponse(w, out)
	})
}

type _ForumServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _ForumServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _ForumServiceHandler) (out proto.Message, err error)
type ForumServiceInterceptor struct {
	middleware []_ForumServiceMiddleware
	server     ForumServiceServer
}

// NewForumServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewForumServiceInterceptor(srv ForumServiceServer, middleware ..._ForumServiceMiddleware) *ForumServiceInterceptor {
	return &ForumServiceInterceptor{server: srv, middleware: middleware}
}

func (i *ForumServiceInterceptor) DescribeForum(ctx context.Context, in *DescribeForumInput) (*DescribeForumOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeForumInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeForumInput, got %T", in))
		}

		return i.server.DescribeForum(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.ForumService.DescribeForum", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeForumOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeForumOutput, got %T", out))
	}

	return message, err
}

func (i *ForumServiceInterceptor) ListForums(ctx context.Context, in *ListForumsInput) (*ListForumsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListForumsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListForumsInput, got %T", in))
		}

		return i.server.ListForums(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.ForumService.ListForums", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListForumsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListForumsOutput, got %T", out))
	}

	return message, err
}

func (i *ForumServiceInterceptor) CreateForum(ctx context.Context, in *CreateForumInput) (*CreateForumOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateForumInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateForumInput, got %T", in))
		}

		return i.server.CreateForum(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.ForumService.CreateForum", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateForumOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateForumOutput, got %T", out))
	}

	return message, err
}

func (i *ForumServiceInterceptor) UpdateForum(ctx context.Context, in *UpdateForumInput) (*UpdateForumOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateForumInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateForumInput, got %T", in))
		}

		return i.server.UpdateForum(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.ForumService.UpdateForum", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateForumOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateForumOutput, got %T", out))
	}

	return message, err
}

func (i *ForumServiceInterceptor) DeleteForum(ctx context.Context, in *DeleteForumInput) (*DeleteForumOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteForumInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteForumInput, got %T", in))
		}

		return i.server.DeleteForum(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.discussion.ForumService.DeleteForum", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteForumOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteForumOutput, got %T", out))
	}

	return message, err
}
