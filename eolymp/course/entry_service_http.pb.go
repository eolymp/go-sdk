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
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _EntryService_HTTPReadQueryString parses body into proto.Message
func _EntryService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _EntryService_HTTPReadRequestBody parses body into proto.Message
func _EntryService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _EntryService_HTTPWriteResponse writes proto.Message to HTTP response
func _EntryService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_EntryService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _EntryService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _EntryService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _EntryService_WebsocketErrorResponse writes error to websocket connection
func _EntryService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _EntryService_WebsocketCodec implements protobuf codec for websockets package
var _EntryService_WebsocketCodec = websocket.Codec{
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

// RegisterEntryServiceHttpHandlers adds handlers for for EntryServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterEntryServiceHttpHandlers(router *mux.Router, prefix string, cli EntryServiceClient) {
	router.Handle(prefix+"/entries", _EntryService_CreateEntry_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.EntryService.CreateEntry")
	router.Handle(prefix+"/entries/{entry_id}", _EntryService_UpdateEntry_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.EntryService.UpdateEntry")
	router.Handle(prefix+"/entries/{entry_id}/title", _EntryService_RenameEntry_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.EntryService.RenameEntry")
	router.Handle(prefix+"/entries/{entry_id}/position", _EntryService_MoveEntry_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.EntryService.MoveEntry")
	router.Handle(prefix+"/entries/{entry_id}", _EntryService_DeleteEntry_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.course.EntryService.DeleteEntry")
	router.Handle(prefix+"/entries/{entry_id}", _EntryService_DescribeEntry_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.EntryService.DescribeEntry")
	router.Handle(prefix+"/entries", _EntryService_ListEntries_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.EntryService.ListEntries")
	router.Handle(prefix+"/toc", _EntryService_DescribeTOC_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.EntryService.DescribeTOC")
	router.Handle(prefix+"/entries/{entry_id}/parents", _EntryService_ListParents_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.EntryService.ListParents")
	router.Handle(prefix+"/entries/{entry_id}/progress", _EntryService_DescribeProgress_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.EntryService.DescribeProgress")
	router.Handle(prefix+"/entries/{entry_id}/progress", _EntryService_ReportProgress_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.EntryService.ReportProgress")
}

func _EntryService_CreateEntry_Rule0(cli EntryServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateEntryInput{}

		if err := _EntryService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.CreateEntry(r.Context(), in)
		if err != nil {
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EntryService_HTTPWriteResponse(w, out)
	})
}

func _EntryService_UpdateEntry_Rule0(cli EntryServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateEntryInput{}

		if err := _EntryService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EntryId = vars["entry_id"]

		out, err := cli.UpdateEntry(r.Context(), in)
		if err != nil {
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EntryService_HTTPWriteResponse(w, out)
	})
}

func _EntryService_RenameEntry_Rule0(cli EntryServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RenameEntryInput{}

		if err := _EntryService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EntryId = vars["entry_id"]

		out, err := cli.RenameEntry(r.Context(), in)
		if err != nil {
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EntryService_HTTPWriteResponse(w, out)
	})
}

func _EntryService_MoveEntry_Rule0(cli EntryServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &MoveEntryInput{}

		if err := _EntryService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EntryId = vars["entry_id"]

		out, err := cli.MoveEntry(r.Context(), in)
		if err != nil {
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EntryService_HTTPWriteResponse(w, out)
	})
}

func _EntryService_DeleteEntry_Rule0(cli EntryServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteEntryInput{}

		if err := _EntryService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EntryId = vars["entry_id"]

		out, err := cli.DeleteEntry(r.Context(), in)
		if err != nil {
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EntryService_HTTPWriteResponse(w, out)
	})
}

func _EntryService_DescribeEntry_Rule0(cli EntryServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeEntryInput{}

		if err := _EntryService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EntryId = vars["entry_id"]

		out, err := cli.DescribeEntry(r.Context(), in)
		if err != nil {
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EntryService_HTTPWriteResponse(w, out)
	})
}

func _EntryService_ListEntries_Rule0(cli EntryServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListEntriesInput{}

		if err := _EntryService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListEntries(r.Context(), in)
		if err != nil {
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EntryService_HTTPWriteResponse(w, out)
	})
}

func _EntryService_DescribeTOC_Rule0(cli EntryServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeTOCInput{}

		if err := _EntryService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.DescribeTOC(r.Context(), in)
		if err != nil {
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EntryService_HTTPWriteResponse(w, out)
	})
}

func _EntryService_ListParents_Rule0(cli EntryServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListParentsInput{}

		if err := _EntryService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EntryId = vars["entry_id"]

		out, err := cli.ListParents(r.Context(), in)
		if err != nil {
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EntryService_HTTPWriteResponse(w, out)
	})
}

func _EntryService_DescribeProgress_Rule0(cli EntryServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeProgressInput{}

		if err := _EntryService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EntryId = vars["entry_id"]

		out, err := cli.DescribeProgress(r.Context(), in)
		if err != nil {
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EntryService_HTTPWriteResponse(w, out)
	})
}

func _EntryService_ReportProgress_Rule0(cli EntryServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ReportProgressInput{}

		if err := _EntryService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EntryId = vars["entry_id"]

		out, err := cli.ReportProgress(r.Context(), in)
		if err != nil {
			_EntryService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EntryService_HTTPWriteResponse(w, out)
	})
}

type _EntryServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _EntryServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _EntryServiceHandler) (out proto.Message, err error)
type EntryServiceInterceptor struct {
	middleware []_EntryServiceMiddleware
	client     EntryServiceClient
}

// NewEntryServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewEntryServiceInterceptor(cli EntryServiceClient, middleware ..._EntryServiceMiddleware) *EntryServiceInterceptor {
	return &EntryServiceInterceptor{client: cli, middleware: middleware}
}

func (i *EntryServiceInterceptor) CreateEntry(ctx context.Context, in *CreateEntryInput, opts ...grpc.CallOption) (*CreateEntryOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateEntryInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateEntryInput, got %T", in))
		}

		return i.client.CreateEntry(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.EntryService.CreateEntry", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateEntryOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateEntryOutput, got %T", out))
	}

	return message, err
}

func (i *EntryServiceInterceptor) UpdateEntry(ctx context.Context, in *UpdateEntryInput, opts ...grpc.CallOption) (*UpdateEntryOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateEntryInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateEntryInput, got %T", in))
		}

		return i.client.UpdateEntry(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.EntryService.UpdateEntry", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateEntryOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateEntryOutput, got %T", out))
	}

	return message, err
}

func (i *EntryServiceInterceptor) RenameEntry(ctx context.Context, in *RenameEntryInput, opts ...grpc.CallOption) (*RenameEntryOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RenameEntryInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RenameEntryInput, got %T", in))
		}

		return i.client.RenameEntry(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.EntryService.RenameEntry", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RenameEntryOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RenameEntryOutput, got %T", out))
	}

	return message, err
}

func (i *EntryServiceInterceptor) MoveEntry(ctx context.Context, in *MoveEntryInput, opts ...grpc.CallOption) (*MoveEntryOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*MoveEntryInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *MoveEntryInput, got %T", in))
		}

		return i.client.MoveEntry(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.EntryService.MoveEntry", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*MoveEntryOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *MoveEntryOutput, got %T", out))
	}

	return message, err
}

func (i *EntryServiceInterceptor) DeleteEntry(ctx context.Context, in *DeleteEntryInput, opts ...grpc.CallOption) (*DeleteEntryOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteEntryInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteEntryInput, got %T", in))
		}

		return i.client.DeleteEntry(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.EntryService.DeleteEntry", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteEntryOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteEntryOutput, got %T", out))
	}

	return message, err
}

func (i *EntryServiceInterceptor) DescribeEntry(ctx context.Context, in *DescribeEntryInput, opts ...grpc.CallOption) (*DescribeEntryOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeEntryInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeEntryInput, got %T", in))
		}

		return i.client.DescribeEntry(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.EntryService.DescribeEntry", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeEntryOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeEntryOutput, got %T", out))
	}

	return message, err
}

func (i *EntryServiceInterceptor) ListEntries(ctx context.Context, in *ListEntriesInput, opts ...grpc.CallOption) (*ListEntriesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListEntriesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListEntriesInput, got %T", in))
		}

		return i.client.ListEntries(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.EntryService.ListEntries", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListEntriesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListEntriesOutput, got %T", out))
	}

	return message, err
}

func (i *EntryServiceInterceptor) DescribeTOC(ctx context.Context, in *DescribeTOCInput, opts ...grpc.CallOption) (*DescribeTOCOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeTOCInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeTOCInput, got %T", in))
		}

		return i.client.DescribeTOC(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.EntryService.DescribeTOC", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeTOCOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeTOCOutput, got %T", out))
	}

	return message, err
}

func (i *EntryServiceInterceptor) ListParents(ctx context.Context, in *ListParentsInput, opts ...grpc.CallOption) (*ListParentsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListParentsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListParentsInput, got %T", in))
		}

		return i.client.ListParents(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.EntryService.ListParents", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListParentsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListParentsOutput, got %T", out))
	}

	return message, err
}

func (i *EntryServiceInterceptor) DescribeProgress(ctx context.Context, in *DescribeProgressInput, opts ...grpc.CallOption) (*DescribeProgressOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeProgressInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeProgressInput, got %T", in))
		}

		return i.client.DescribeProgress(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.EntryService.DescribeProgress", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeProgressOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeProgressOutput, got %T", out))
	}

	return message, err
}

func (i *EntryServiceInterceptor) ReportProgress(ctx context.Context, in *ReportProgressInput, opts ...grpc.CallOption) (*ReportProgressOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ReportProgressInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ReportProgressInput, got %T", in))
		}

		return i.client.ReportProgress(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.EntryService.ReportProgress", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ReportProgressOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ReportProgressOutput, got %T", out))
	}

	return message, err
}
