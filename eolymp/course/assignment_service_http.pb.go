// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package course

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

// _AssignmentService_HTTPReadQueryString parses body into proto.Message
func _AssignmentService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _AssignmentService_HTTPReadRequestBody parses body into proto.Message
func _AssignmentService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _AssignmentService_HTTPWriteResponse writes proto.Message to HTTP response
func _AssignmentService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_AssignmentService_HTTPWriteErrorResponse(w, err)
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

// _AssignmentService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _AssignmentService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _AssignmentService_WebsocketErrorResponse writes error to websocket connection
func _AssignmentService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _AssignmentService_WebsocketCodec implements protobuf codec for websockets package
var _AssignmentService_WebsocketCodec = websocket.Codec{
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

// RegisterAssignmentServiceHttpHandlers adds handlers for for AssignmentServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterAssignmentServiceHttpHandlers(router *mux.Router, prefix string, cli AssignmentServiceClient) {
	router.Handle(prefix+"/assignments", _AssignmentService_CreateAssignment_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.AssignmentService.CreateAssignment")
	router.Handle(prefix+"/assignments/{assignment_id}", _AssignmentService_UpdateAssignment_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.AssignmentService.UpdateAssignment")
	router.Handle(prefix+"/assignments/{assignment_id}", _AssignmentService_DeleteAssignment_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.course.AssignmentService.DeleteAssignment")
	router.Handle(prefix+"/assignments/{assignment_id}", _AssignmentService_DescribeAssignment_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.AssignmentService.DescribeAssignment")
	router.Handle(prefix+"/viewer/assignment", _AssignmentService_IntrospectAssignment_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.AssignmentService.IntrospectAssignment")
	router.Handle(prefix+"/assignments", _AssignmentService_ListAssignments_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.AssignmentService.ListAssignments")
	router.Handle(prefix+"/assignments/{assignment_id}/start", _AssignmentService_StartAssignment_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.AssignmentService.StartAssignment")
}

func _AssignmentService_CreateAssignment_Rule0(cli AssignmentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateAssignmentInput{}

		if err := _AssignmentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateAssignment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssignmentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssignmentService_UpdateAssignment_Rule0(cli AssignmentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateAssignmentInput{}

		if err := _AssignmentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AssignmentId = vars["assignment_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateAssignment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssignmentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssignmentService_DeleteAssignment_Rule0(cli AssignmentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteAssignmentInput{}

		if err := _AssignmentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AssignmentId = vars["assignment_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteAssignment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssignmentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssignmentService_DescribeAssignment_Rule0(cli AssignmentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeAssignmentInput{}

		if err := _AssignmentService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AssignmentId = vars["assignment_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeAssignment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssignmentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssignmentService_IntrospectAssignment_Rule0(cli AssignmentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &IntrospectAssignmentInput{}

		if err := _AssignmentService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.IntrospectAssignment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssignmentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssignmentService_ListAssignments_Rule0(cli AssignmentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListAssignmentsInput{}

		if err := _AssignmentService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListAssignments(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssignmentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssignmentService_StartAssignment_Rule0(cli AssignmentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &StartAssignmentInput{}

		if err := _AssignmentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AssignmentId = vars["assignment_id"]

		var header, trailer metadata.MD

		out, err := cli.StartAssignment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssignmentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssignmentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _AssignmentServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _AssignmentServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _AssignmentServiceHandler) (out proto.Message, err error)
type AssignmentServiceInterceptor struct {
	middleware []_AssignmentServiceMiddleware
	client     AssignmentServiceClient
}

// NewAssignmentServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewAssignmentServiceInterceptor(cli AssignmentServiceClient, middleware ..._AssignmentServiceMiddleware) *AssignmentServiceInterceptor {
	return &AssignmentServiceInterceptor{client: cli, middleware: middleware}
}

func (i *AssignmentServiceInterceptor) CreateAssignment(ctx context.Context, in *CreateAssignmentInput, opts ...grpc.CallOption) (*CreateAssignmentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateAssignmentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateAssignmentInput, got %T", in))
		}

		return i.client.CreateAssignment(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.AssignmentService.CreateAssignment", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateAssignmentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateAssignmentOutput, got %T", out))
	}

	return message, err
}

func (i *AssignmentServiceInterceptor) UpdateAssignment(ctx context.Context, in *UpdateAssignmentInput, opts ...grpc.CallOption) (*UpdateAssignmentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateAssignmentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateAssignmentInput, got %T", in))
		}

		return i.client.UpdateAssignment(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.AssignmentService.UpdateAssignment", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateAssignmentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateAssignmentOutput, got %T", out))
	}

	return message, err
}

func (i *AssignmentServiceInterceptor) DeleteAssignment(ctx context.Context, in *DeleteAssignmentInput, opts ...grpc.CallOption) (*DeleteAssignmentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteAssignmentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteAssignmentInput, got %T", in))
		}

		return i.client.DeleteAssignment(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.AssignmentService.DeleteAssignment", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteAssignmentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteAssignmentOutput, got %T", out))
	}

	return message, err
}

func (i *AssignmentServiceInterceptor) DescribeAssignment(ctx context.Context, in *DescribeAssignmentInput, opts ...grpc.CallOption) (*DescribeAssignmentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeAssignmentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeAssignmentInput, got %T", in))
		}

		return i.client.DescribeAssignment(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.AssignmentService.DescribeAssignment", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeAssignmentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeAssignmentOutput, got %T", out))
	}

	return message, err
}

func (i *AssignmentServiceInterceptor) IntrospectAssignment(ctx context.Context, in *IntrospectAssignmentInput, opts ...grpc.CallOption) (*IntrospectAssignmentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*IntrospectAssignmentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *IntrospectAssignmentInput, got %T", in))
		}

		return i.client.IntrospectAssignment(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.AssignmentService.IntrospectAssignment", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*IntrospectAssignmentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *IntrospectAssignmentOutput, got %T", out))
	}

	return message, err
}

func (i *AssignmentServiceInterceptor) ListAssignments(ctx context.Context, in *ListAssignmentsInput, opts ...grpc.CallOption) (*ListAssignmentsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListAssignmentsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListAssignmentsInput, got %T", in))
		}

		return i.client.ListAssignments(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.AssignmentService.ListAssignments", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListAssignmentsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListAssignmentsOutput, got %T", out))
	}

	return message, err
}

func (i *AssignmentServiceInterceptor) StartAssignment(ctx context.Context, in *StartAssignmentInput, opts ...grpc.CallOption) (*StartAssignmentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*StartAssignmentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *StartAssignmentInput, got %T", in))
		}

		return i.client.StartAssignment(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.AssignmentService.StartAssignment", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*StartAssignmentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *StartAssignmentOutput, got %T", out))
	}

	return message, err
}