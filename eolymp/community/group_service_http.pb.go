// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package community

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

// _GroupService_HTTPReadQueryString parses body into proto.Message
func _GroupService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _GroupService_HTTPReadRequestBody parses body into proto.Message
func _GroupService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _GroupService_HTTPWriteResponse writes proto.Message to HTTP response
func _GroupService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_GroupService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _GroupService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _GroupService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _GroupService_WebsocketErrorResponse writes error to websocket connection
func _GroupService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _GroupService_WebsocketCodec implements protobuf codec for websockets package
var _GroupService_WebsocketCodec = websocket.Codec{
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

// RegisterGroupServiceHttpHandlers adds handlers for for GroupServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterGroupServiceHttpHandlers(router *mux.Router, prefix string, cli GroupServiceClient) {
	router.Handle(prefix+"/groups", _GroupService_CreateGroup_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.GroupService.CreateGroup")
	router.Handle(prefix+"/groups/{group_id}", _GroupService_UpdateGroup_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.GroupService.UpdateGroup")
	router.Handle(prefix+"/groups/{group_id}", _GroupService_DeleteGroup_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.community.GroupService.DeleteGroup")
	router.Handle(prefix+"/groups/{group_id}", _GroupService_DescribeGroup_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.GroupService.DescribeGroup")
	router.Handle(prefix+"/groups", _GroupService_ListGroups_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.GroupService.ListGroups")
}

func _GroupService_CreateGroup_Rule0(cli GroupServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateGroupInput{}

		if err := _GroupService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_GroupService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.CreateGroup(r.Context(), in)
		if err != nil {
			_GroupService_HTTPWriteErrorResponse(w, err)
			return
		}

		_GroupService_HTTPWriteResponse(w, out)
	})
}

func _GroupService_UpdateGroup_Rule0(cli GroupServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateGroupInput{}

		if err := _GroupService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_GroupService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.GroupId = vars["group_id"]

		out, err := cli.UpdateGroup(r.Context(), in)
		if err != nil {
			_GroupService_HTTPWriteErrorResponse(w, err)
			return
		}

		_GroupService_HTTPWriteResponse(w, out)
	})
}

func _GroupService_DeleteGroup_Rule0(cli GroupServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteGroupInput{}

		if err := _GroupService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_GroupService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.GroupId = vars["group_id"]

		out, err := cli.DeleteGroup(r.Context(), in)
		if err != nil {
			_GroupService_HTTPWriteErrorResponse(w, err)
			return
		}

		_GroupService_HTTPWriteResponse(w, out)
	})
}

func _GroupService_DescribeGroup_Rule0(cli GroupServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeGroupInput{}

		if err := _GroupService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_GroupService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.GroupId = vars["group_id"]

		out, err := cli.DescribeGroup(r.Context(), in)
		if err != nil {
			_GroupService_HTTPWriteErrorResponse(w, err)
			return
		}

		_GroupService_HTTPWriteResponse(w, out)
	})
}

func _GroupService_ListGroups_Rule0(cli GroupServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListGroupsInput{}

		if err := _GroupService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_GroupService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListGroups(r.Context(), in)
		if err != nil {
			_GroupService_HTTPWriteErrorResponse(w, err)
			return
		}

		_GroupService_HTTPWriteResponse(w, out)
	})
}

type _GroupServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _GroupServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _GroupServiceHandler) (out proto.Message, err error)
type GroupServiceInterceptor struct {
	middleware []_GroupServiceMiddleware
	client     GroupServiceClient
}

// NewGroupServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewGroupServiceInterceptor(cli GroupServiceClient, middleware ..._GroupServiceMiddleware) *GroupServiceInterceptor {
	return &GroupServiceInterceptor{client: cli, middleware: middleware}
}

func (i *GroupServiceInterceptor) CreateGroup(ctx context.Context, in *CreateGroupInput, opts ...grpc.CallOption) (*CreateGroupOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateGroupInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateGroupInput, got %T", in))
		}

		return i.client.CreateGroup(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.GroupService.CreateGroup", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateGroupOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateGroupOutput, got %T", out))
	}

	return message, err
}

func (i *GroupServiceInterceptor) UpdateGroup(ctx context.Context, in *UpdateGroupInput, opts ...grpc.CallOption) (*UpdateGroupOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateGroupInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateGroupInput, got %T", in))
		}

		return i.client.UpdateGroup(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.GroupService.UpdateGroup", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateGroupOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateGroupOutput, got %T", out))
	}

	return message, err
}

func (i *GroupServiceInterceptor) DeleteGroup(ctx context.Context, in *DeleteGroupInput, opts ...grpc.CallOption) (*DeleteGroupOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteGroupInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteGroupInput, got %T", in))
		}

		return i.client.DeleteGroup(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.GroupService.DeleteGroup", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteGroupOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteGroupOutput, got %T", out))
	}

	return message, err
}

func (i *GroupServiceInterceptor) DescribeGroup(ctx context.Context, in *DescribeGroupInput, opts ...grpc.CallOption) (*DescribeGroupOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeGroupInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeGroupInput, got %T", in))
		}

		return i.client.DescribeGroup(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.GroupService.DescribeGroup", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeGroupOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeGroupOutput, got %T", out))
	}

	return message, err
}

func (i *GroupServiceInterceptor) ListGroups(ctx context.Context, in *ListGroupsInput, opts ...grpc.CallOption) (*ListGroupsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListGroupsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListGroupsInput, got %T", in))
		}

		return i.client.ListGroups(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.GroupService.ListGroups", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListGroupsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListGroupsOutput, got %T", out))
	}

	return message, err
}
