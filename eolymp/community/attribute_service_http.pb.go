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

// _AttributeService_HTTPReadQueryString parses body into proto.Message
func _AttributeService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _AttributeService_HTTPReadRequestBody parses body into proto.Message
func _AttributeService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _AttributeService_HTTPWriteResponse writes proto.Message to HTTP response
func _AttributeService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_AttributeService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _AttributeService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _AttributeService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _AttributeService_WebsocketErrorResponse writes error to websocket connection
func _AttributeService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _AttributeService_WebsocketCodec implements protobuf codec for websockets package
var _AttributeService_WebsocketCodec = websocket.Codec{
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

// RegisterAttributeServiceHttpHandlers adds handlers for for AttributeServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterAttributeServiceHttpHandlers(router *mux.Router, prefix string, cli AttributeServiceClient) {
	router.Handle(prefix+"/attributes", _AttributeService_CreateAttribute_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.AttributeService.CreateAttribute")
	router.Handle(prefix+"/attributes/{attribute_key}", _AttributeService_UpdateAttribute_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.AttributeService.UpdateAttribute")
	router.Handle(prefix+"/attributes/{attribute_key}", _AttributeService_RemoveAttribute_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.community.AttributeService.RemoveAttribute")
	router.Handle(prefix+"/attributes/{attribute_key}", _AttributeService_DescribeAttribute_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.AttributeService.DescribeAttribute")
	router.Handle(prefix+"/attributes", _AttributeService_ListAttributes_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.AttributeService.ListAttributes")
}

func _AttributeService_CreateAttribute_Rule0(cli AttributeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateAttributeInput{}

		if err := _AttributeService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AttributeService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.CreateAttribute(r.Context(), in)
		if err != nil {
			_AttributeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AttributeService_HTTPWriteResponse(w, out)
	})
}

func _AttributeService_UpdateAttribute_Rule0(cli AttributeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateAttributeInput{}

		if err := _AttributeService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AttributeService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AttributeKey = vars["attribute_key"]

		out, err := cli.UpdateAttribute(r.Context(), in)
		if err != nil {
			_AttributeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AttributeService_HTTPWriteResponse(w, out)
	})
}

func _AttributeService_RemoveAttribute_Rule0(cli AttributeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RemoveAttributeInput{}

		if err := _AttributeService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AttributeService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AttributeKey = vars["attribute_key"]

		out, err := cli.RemoveAttribute(r.Context(), in)
		if err != nil {
			_AttributeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AttributeService_HTTPWriteResponse(w, out)
	})
}

func _AttributeService_DescribeAttribute_Rule0(cli AttributeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeAttributeInput{}

		if err := _AttributeService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AttributeService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AttributeKey = vars["attribute_key"]

		out, err := cli.DescribeAttribute(r.Context(), in)
		if err != nil {
			_AttributeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AttributeService_HTTPWriteResponse(w, out)
	})
}

func _AttributeService_ListAttributes_Rule0(cli AttributeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListAttributesInput{}

		if err := _AttributeService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AttributeService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListAttributes(r.Context(), in)
		if err != nil {
			_AttributeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AttributeService_HTTPWriteResponse(w, out)
	})
}

type _AttributeServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _AttributeServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _AttributeServiceHandler) (out proto.Message, err error)
type AttributeServiceInterceptor struct {
	middleware []_AttributeServiceMiddleware
	client     AttributeServiceClient
}

// NewAttributeServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewAttributeServiceInterceptor(cli AttributeServiceClient, middleware ..._AttributeServiceMiddleware) *AttributeServiceInterceptor {
	return &AttributeServiceInterceptor{client: cli, middleware: middleware}
}

func (i *AttributeServiceInterceptor) CreateAttribute(ctx context.Context, in *CreateAttributeInput, opts ...grpc.CallOption) (*CreateAttributeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateAttributeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateAttributeInput, got %T", in))
		}

		return i.client.CreateAttribute(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.AttributeService.CreateAttribute", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateAttributeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateAttributeOutput, got %T", out))
	}

	return message, err
}

func (i *AttributeServiceInterceptor) UpdateAttribute(ctx context.Context, in *UpdateAttributeInput, opts ...grpc.CallOption) (*UpdateAttributeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateAttributeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateAttributeInput, got %T", in))
		}

		return i.client.UpdateAttribute(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.AttributeService.UpdateAttribute", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateAttributeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateAttributeOutput, got %T", out))
	}

	return message, err
}

func (i *AttributeServiceInterceptor) RemoveAttribute(ctx context.Context, in *RemoveAttributeInput, opts ...grpc.CallOption) (*RemoveAttributeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RemoveAttributeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RemoveAttributeInput, got %T", in))
		}

		return i.client.RemoveAttribute(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.AttributeService.RemoveAttribute", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RemoveAttributeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RemoveAttributeOutput, got %T", out))
	}

	return message, err
}

func (i *AttributeServiceInterceptor) DescribeAttribute(ctx context.Context, in *DescribeAttributeInput, opts ...grpc.CallOption) (*DescribeAttributeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeAttributeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeAttributeInput, got %T", in))
		}

		return i.client.DescribeAttribute(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.AttributeService.DescribeAttribute", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeAttributeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeAttributeOutput, got %T", out))
	}

	return message, err
}

func (i *AttributeServiceInterceptor) ListAttributes(ctx context.Context, in *ListAttributesInput, opts ...grpc.CallOption) (*ListAttributesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListAttributesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListAttributesInput, got %T", in))
		}

		return i.client.ListAttributes(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.AttributeService.ListAttributes", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListAttributesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListAttributesOutput, got %T", out))
	}

	return message, err
}
