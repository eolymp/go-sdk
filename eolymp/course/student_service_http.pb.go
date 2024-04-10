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

// _StudentService_HTTPReadQueryString parses body into proto.Message
func _StudentService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _StudentService_HTTPReadRequestBody parses body into proto.Message
func _StudentService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _StudentService_HTTPWriteResponse writes proto.Message to HTTP response
func _StudentService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_StudentService_HTTPWriteErrorResponse(w, err)
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

// _StudentService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _StudentService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _StudentService_WebsocketErrorResponse writes error to websocket connection
func _StudentService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _StudentService_WebsocketCodec implements protobuf codec for websockets package
var _StudentService_WebsocketCodec = websocket.Codec{
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

// RegisterStudentServiceHttpHandlers adds handlers for for StudentServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterStudentServiceHttpHandlers(router *mux.Router, prefix string, cli StudentServiceClient) {
	router.Handle(prefix+"/students", _StudentService_CreateStudent_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.StudentService.CreateStudent")
	router.Handle(prefix+"/students/{student_id}", _StudentService_UpdateStudent_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.StudentService.UpdateStudent")
	router.Handle(prefix+"/students/{student_id}", _StudentService_DeleteStudent_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.course.StudentService.DeleteStudent")
	router.Handle(prefix+"/students/{student_id}", _StudentService_DescribeStudent_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.StudentService.DescribeStudent")
	router.Handle(prefix+"/students", _StudentService_ListStudents_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.StudentService.ListStudents")
	router.Handle(prefix+"/viewer", _StudentService_DescribeViewer_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.StudentService.DescribeViewer")
}

func _StudentService_CreateStudent_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateStudentInput{}

		if err := _StudentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateStudent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_UpdateStudent_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateStudentInput{}

		if err := _StudentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.StudentId = vars["student_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateStudent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_DeleteStudent_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteStudentInput{}

		if err := _StudentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.StudentId = vars["student_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteStudent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_DescribeStudent_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeStudentInput{}

		if err := _StudentService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.StudentId = vars["student_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeStudent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_ListStudents_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListStudentsInput{}

		if err := _StudentService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListStudents(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_DescribeViewer_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeViewerInput{}

		if err := _StudentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeViewer(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _StudentServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _StudentServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _StudentServiceHandler) (out proto.Message, err error)
type StudentServiceInterceptor struct {
	middleware []_StudentServiceMiddleware
	client     StudentServiceClient
}

// NewStudentServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewStudentServiceInterceptor(cli StudentServiceClient, middleware ..._StudentServiceMiddleware) *StudentServiceInterceptor {
	return &StudentServiceInterceptor{client: cli, middleware: middleware}
}

func (i *StudentServiceInterceptor) CreateStudent(ctx context.Context, in *CreateStudentInput, opts ...grpc.CallOption) (*CreateStudentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateStudentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateStudentInput, got %T", in))
		}

		return i.client.CreateStudent(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.StudentService.CreateStudent", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateStudentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateStudentOutput, got %T", out))
	}

	return message, err
}

func (i *StudentServiceInterceptor) UpdateStudent(ctx context.Context, in *UpdateStudentInput, opts ...grpc.CallOption) (*UpdateStudentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateStudentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateStudentInput, got %T", in))
		}

		return i.client.UpdateStudent(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.StudentService.UpdateStudent", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateStudentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateStudentOutput, got %T", out))
	}

	return message, err
}

func (i *StudentServiceInterceptor) DeleteStudent(ctx context.Context, in *DeleteStudentInput, opts ...grpc.CallOption) (*DeleteStudentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteStudentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteStudentInput, got %T", in))
		}

		return i.client.DeleteStudent(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.StudentService.DeleteStudent", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteStudentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteStudentOutput, got %T", out))
	}

	return message, err
}

func (i *StudentServiceInterceptor) DescribeStudent(ctx context.Context, in *DescribeStudentInput, opts ...grpc.CallOption) (*DescribeStudentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeStudentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeStudentInput, got %T", in))
		}

		return i.client.DescribeStudent(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.StudentService.DescribeStudent", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeStudentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeStudentOutput, got %T", out))
	}

	return message, err
}

func (i *StudentServiceInterceptor) ListStudents(ctx context.Context, in *ListStudentsInput, opts ...grpc.CallOption) (*ListStudentsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListStudentsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListStudentsInput, got %T", in))
		}

		return i.client.ListStudents(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.StudentService.ListStudents", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListStudentsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListStudentsOutput, got %T", out))
	}

	return message, err
}

func (i *StudentServiceInterceptor) DescribeViewer(ctx context.Context, in *DescribeViewerInput, opts ...grpc.CallOption) (*DescribeViewerOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeViewerInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeViewerInput, got %T", in))
		}

		return i.client.DescribeViewer(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.StudentService.DescribeViewer", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeViewerOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeViewerOutput, got %T", out))
	}

	return message, err
}
