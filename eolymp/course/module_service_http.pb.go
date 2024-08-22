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

// _ModuleService_HTTPReadQueryString parses body into proto.Message
func _ModuleService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _ModuleService_HTTPReadRequestBody parses body into proto.Message
func _ModuleService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _ModuleService_HTTPWriteResponse writes proto.Message to HTTP response
func _ModuleService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_ModuleService_HTTPWriteErrorResponse(w, err)
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

// _ModuleService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _ModuleService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _ModuleService_WebsocketErrorResponse writes error to websocket connection
func _ModuleService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _ModuleService_WebsocketCodec implements protobuf codec for websockets package
var _ModuleService_WebsocketCodec = websocket.Codec{
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

// RegisterModuleServiceHttpHandlers adds handlers for for ModuleServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterModuleServiceHttpHandlers(router *mux.Router, prefix string, cli ModuleServiceClient) {
	router.Handle(prefix+"/modules", _ModuleService_CreateModule_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.ModuleService.CreateModule")
	router.Handle(prefix+"/modules/{module_id}", _ModuleService_UpdateModule_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.ModuleService.UpdateModule")
	router.Handle(prefix+"/modules/{module_id}", _ModuleService_DeleteModule_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.course.ModuleService.DeleteModule")
	router.Handle(prefix+"/modules/{module_id}", _ModuleService_DescribeModule_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.ModuleService.DescribeModule")
	router.Handle(prefix+"/modules", _ModuleService_ListModules_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.ModuleService.ListModules")
	router.Handle(prefix+"/modules/{module_id}/start", _ModuleService_StartModule_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.ModuleService.StartModule")
	router.Handle(prefix+"/modules/{module_id}/assignments/{member_id}", _ModuleService_AssignModule_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.ModuleService.AssignModule")
	router.Handle(prefix+"/modules/{module_id}/assignments/{member_id}", _ModuleService_UnassignModule_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.course.ModuleService.UnassignModule")
}

func _ModuleService_CreateModule_Rule0(cli ModuleServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateModuleInput{}

		if err := _ModuleService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateModule(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleService_UpdateModule_Rule0(cli ModuleServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateModuleInput{}

		if err := _ModuleService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ModuleId = vars["module_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateModule(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleService_DeleteModule_Rule0(cli ModuleServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteModuleInput{}

		if err := _ModuleService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ModuleId = vars["module_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteModule(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleService_DescribeModule_Rule0(cli ModuleServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeModuleInput{}

		if err := _ModuleService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ModuleId = vars["module_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeModule(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleService_ListModules_Rule0(cli ModuleServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListModulesInput{}

		if err := _ModuleService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListModules(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleService_StartModule_Rule0(cli ModuleServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &StartModuleInput{}

		if err := _ModuleService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ModuleId = vars["module_id"]

		var header, trailer metadata.MD

		out, err := cli.StartModule(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleService_AssignModule_Rule0(cli ModuleServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &AssignModuleInput{}

		if err := _ModuleService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ModuleId = vars["module_id"]
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.AssignModule(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ModuleService_UnassignModule_Rule0(cli ModuleServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UnassignModuleInput{}

		if err := _ModuleService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ModuleId = vars["module_id"]
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.UnassignModule(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ModuleService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ModuleService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _ModuleServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _ModuleServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _ModuleServiceHandler) (out proto.Message, err error)
type ModuleServiceInterceptor struct {
	middleware []_ModuleServiceMiddleware
	client     ModuleServiceClient
}

// NewModuleServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewModuleServiceInterceptor(cli ModuleServiceClient, middleware ..._ModuleServiceMiddleware) *ModuleServiceInterceptor {
	return &ModuleServiceInterceptor{client: cli, middleware: middleware}
}

func (i *ModuleServiceInterceptor) CreateModule(ctx context.Context, in *CreateModuleInput, opts ...grpc.CallOption) (*CreateModuleOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateModuleInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateModuleInput, got %T", in))
		}

		return i.client.CreateModule(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleService.CreateModule", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateModuleOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateModuleOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleServiceInterceptor) UpdateModule(ctx context.Context, in *UpdateModuleInput, opts ...grpc.CallOption) (*UpdateModuleOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateModuleInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateModuleInput, got %T", in))
		}

		return i.client.UpdateModule(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleService.UpdateModule", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateModuleOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateModuleOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleServiceInterceptor) DeleteModule(ctx context.Context, in *DeleteModuleInput, opts ...grpc.CallOption) (*DeleteModuleOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteModuleInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteModuleInput, got %T", in))
		}

		return i.client.DeleteModule(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleService.DeleteModule", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteModuleOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteModuleOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleServiceInterceptor) DescribeModule(ctx context.Context, in *DescribeModuleInput, opts ...grpc.CallOption) (*DescribeModuleOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeModuleInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeModuleInput, got %T", in))
		}

		return i.client.DescribeModule(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleService.DescribeModule", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeModuleOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeModuleOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleServiceInterceptor) ListModules(ctx context.Context, in *ListModulesInput, opts ...grpc.CallOption) (*ListModulesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListModulesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListModulesInput, got %T", in))
		}

		return i.client.ListModules(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleService.ListModules", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListModulesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListModulesOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleServiceInterceptor) StartModule(ctx context.Context, in *StartModuleInput, opts ...grpc.CallOption) (*StartModuleOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*StartModuleInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *StartModuleInput, got %T", in))
		}

		return i.client.StartModule(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleService.StartModule", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*StartModuleOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *StartModuleOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleServiceInterceptor) AssignModule(ctx context.Context, in *AssignModuleInput, opts ...grpc.CallOption) (*AssignModuleOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*AssignModuleInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *AssignModuleInput, got %T", in))
		}

		return i.client.AssignModule(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleService.AssignModule", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*AssignModuleOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *AssignModuleOutput, got %T", out))
	}

	return message, err
}

func (i *ModuleServiceInterceptor) UnassignModule(ctx context.Context, in *UnassignModuleInput, opts ...grpc.CallOption) (*UnassignModuleOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UnassignModuleInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UnassignModuleInput, got %T", in))
		}

		return i.client.UnassignModule(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.course.ModuleService.UnassignModule", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UnassignModuleOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UnassignModuleOutput, got %T", out))
	}

	return message, err
}
