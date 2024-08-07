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

// _ModuleItemService_HTTPReadQueryString parses body into proto.Message
func _ModuleItemService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _ModuleItemService_HTTPReadRequestBody parses body into proto.Message
func _ModuleItemService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _ModuleItemService_HTTPWriteResponse writes proto.Message to HTTP response
func _ModuleItemService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_ModuleItemService_HTTPWriteErrorResponse(w, err)
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

// _ModuleItemService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _ModuleItemService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _ModuleItemService_WebsocketErrorResponse writes error to websocket connection
func _ModuleItemService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _ModuleItemService_WebsocketCodec implements protobuf codec for websockets package
var _ModuleItemService_WebsocketCodec = websocket.Codec{
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

// RegisterModuleItemServiceHttpHandlers adds handlers for for ModuleItemServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterModuleItemServiceHttpHandlers(router *mux.Router, prefix string, cli ModuleItemServiceClient) {
	router.Handle(prefix+"/items", _ModuleItemService_CreateModuleItem_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.ModuleItemService.CreateModuleItem")
	router.Handle(prefix+"/items/{item_id}", _ModuleItemService_UpdateModuleItem_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.ModuleItemService.UpdateModuleItem")
	router.Handle(prefix+"/items/{item_id}/move", _ModuleItemService_MoveModuleItem_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.ModuleItemService.MoveModuleItem")
	router.Handle(prefix+"/items/{item_id}", _ModuleItemService_DeleteModuleItem_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.course.ModuleItemService.DeleteModuleItem")
	router.Handle(prefix+"/items/{item_id}", _ModuleItemService_DescribeModuleItem_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.ModuleItemService.DescribeModuleItem")
	router.Handle(prefix+"/items", _ModuleItemService_ListModuleItems_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.ModuleItemService.ListModuleItems")
}

func _ModuleItemService_CreateModuleItem_Rule0(cli ModuleItemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateModuleItemInput{}

		if err := _ModuleItemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateModuleItem(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleItemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleItemService_UpdateModuleItem_Rule0(cli ModuleItemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateModuleItemInput{}

		if err := _ModuleItemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ItemId = vars["item_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateModuleItem(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleItemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleItemService_MoveModuleItem_Rule0(cli ModuleItemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &MoveModuleItemInput{}

		if err := _ModuleItemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ItemId = vars["item_id"]

		var header, trailer metadata.MD

		out, err := cli.MoveModuleItem(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleItemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleItemService_DeleteModuleItem_Rule0(cli ModuleItemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteModuleItemInput{}

		if err := _ModuleItemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ItemId = vars["item_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteModuleItem(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleItemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleItemService_DescribeModuleItem_Rule0(cli ModuleItemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeModuleItemInput{}

		if err := _ModuleItemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ItemId = vars["item_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeModuleItem(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleItemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleItemService_ListModuleItems_Rule0(cli ModuleItemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListModuleItemsInput{}

		if err := _ModuleItemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListModuleItems(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleItemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleItemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _ModuleItemServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _ModuleItemServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _ModuleItemServiceHandler) (out proto.Message, err error)
type ModuleItemServiceInterceptor struct {
	middleware []_ModuleItemServiceMiddleware
	client     ModuleItemServiceClient
}

// NewModuleItemServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewModuleItemServiceInterceptor(cli ModuleItemServiceClient, middleware ..._ModuleItemServiceMiddleware) *ModuleItemServiceInterceptor {
	return &ModuleItemServiceInterceptor{client: cli, middleware: middleware}
}

func (i *ModuleItemServiceInterceptor) CreateModuleItem(ctx context.Context, in *CreateModuleItemInput, opts ...grpc.CallOption) (*CreateModuleItemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateModuleItemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateModuleItemInput, got %T", in))
		}

		return i.client.CreateModuleItem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleItemService.CreateModuleItem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateModuleItemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateModuleItemOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleItemServiceInterceptor) UpdateModuleItem(ctx context.Context, in *UpdateModuleItemInput, opts ...grpc.CallOption) (*UpdateModuleItemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateModuleItemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateModuleItemInput, got %T", in))
		}

		return i.client.UpdateModuleItem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleItemService.UpdateModuleItem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateModuleItemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateModuleItemOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleItemServiceInterceptor) MoveModuleItem(ctx context.Context, in *MoveModuleItemInput, opts ...grpc.CallOption) (*MoveModuleItemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*MoveModuleItemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *MoveModuleItemInput, got %T", in))
		}

		return i.client.MoveModuleItem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleItemService.MoveModuleItem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*MoveModuleItemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *MoveModuleItemOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleItemServiceInterceptor) DeleteModuleItem(ctx context.Context, in *DeleteModuleItemInput, opts ...grpc.CallOption) (*DeleteModuleItemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteModuleItemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteModuleItemInput, got %T", in))
		}

		return i.client.DeleteModuleItem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleItemService.DeleteModuleItem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteModuleItemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteModuleItemOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleItemServiceInterceptor) DescribeModuleItem(ctx context.Context, in *DescribeModuleItemInput, opts ...grpc.CallOption) (*DescribeModuleItemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeModuleItemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeModuleItemInput, got %T", in))
		}

		return i.client.DescribeModuleItem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleItemService.DescribeModuleItem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeModuleItemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeModuleItemOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleItemServiceInterceptor) ListModuleItems(ctx context.Context, in *ListModuleItemsInput, opts ...grpc.CallOption) (*ListModuleItemsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListModuleItemsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListModuleItemsInput, got %T", in))
		}

		return i.client.ListModuleItems(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleItemService.ListModuleItems", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListModuleItemsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListModuleItemsOutput, got %T", out))
	}

	return message, err
}
