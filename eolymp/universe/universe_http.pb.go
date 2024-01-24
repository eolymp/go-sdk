// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package universe

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

// _Universe_HTTPReadQueryString parses body into proto.Message
func _Universe_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Universe_HTTPReadRequestBody parses body into proto.Message
func _Universe_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Universe_HTTPWriteResponse writes proto.Message to HTTP response
func _Universe_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Universe_HTTPWriteErrorResponse(w, err)
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

// _Universe_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Universe_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _Universe_WebsocketErrorResponse writes error to websocket connection
func _Universe_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _Universe_WebsocketCodec implements protobuf codec for websockets package
var _Universe_WebsocketCodec = websocket.Codec{
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

// RegisterUniverseHttpHandlers adds handlers for for UniverseClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterUniverseHttpHandlers(router *mux.Router, prefix string, cli UniverseClient) {
	router.Handle(prefix+"/spaces/__lookup/{key}", _Universe_LookupSpace_Rule0(cli)).
		Methods("GET").
		Name("eolymp.universe.Universe.LookupSpace")
	router.Handle(prefix+"/spaces", _Universe_CreateSpace_Rule0(cli)).
		Methods("POST").
		Name("eolymp.universe.Universe.CreateSpace")
	router.Handle(prefix+"/spaces/{space_id}", _Universe_UpdateSpace_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.universe.Universe.UpdateSpace")
	router.Handle(prefix+"/spaces/{space_id}", _Universe_DeleteSpace_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.universe.Universe.DeleteSpace")
	router.Handle(prefix+"/spaces/{space_id}", _Universe_DescribeSpace_Rule0(cli)).
		Methods("GET").
		Name("eolymp.universe.Universe.DescribeSpace")
	router.Handle(prefix+"/spaces/{space_id}/quota", _Universe_DescribeQuota_Rule0(cli)).
		Methods("GET").
		Name("eolymp.universe.Universe.DescribeQuota")
	router.Handle(prefix+"/spaces/{space_id}/quota", _Universe_UpdateQuota_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.universe.Universe.UpdateQuota")
	router.Handle(prefix+"/spaces", _Universe_ListSpaces_Rule0(cli)).
		Methods("GET").
		Name("eolymp.universe.Universe.ListSpaces")
}

func _Universe_LookupSpace_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LookupSpaceInput{}

		if err := _Universe_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.Key = vars["key"]

		var header, trailer metadata.MD

		out, err := cli.LookupSpace(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_CreateSpace_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateSpaceInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateSpace(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_UpdateSpace_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateSpaceInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SpaceId = vars["space_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateSpace(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_DeleteSpace_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteSpaceInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SpaceId = vars["space_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteSpace(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_DescribeSpace_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeSpaceInput{}

		if err := _Universe_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SpaceId = vars["space_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeSpace(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_DescribeQuota_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeQuotaInput{}

		if err := _Universe_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SpaceId = vars["space_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeQuota(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_UpdateQuota_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateQuotaInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SpaceId = vars["space_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateQuota(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_ListSpaces_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListSpacesInput{}

		if err := _Universe_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListSpaces(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _UniverseHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _UniverseMiddleware = func(ctx context.Context, method string, in proto.Message, handler _UniverseHandler) (out proto.Message, err error)
type UniverseInterceptor struct {
	middleware []_UniverseMiddleware
	client     UniverseClient
}

// NewUniverseInterceptor constructs additional middleware for a server based on annotations in proto files
func NewUniverseInterceptor(cli UniverseClient, middleware ..._UniverseMiddleware) *UniverseInterceptor {
	return &UniverseInterceptor{client: cli, middleware: middleware}
}

func (i *UniverseInterceptor) LookupSpace(ctx context.Context, in *LookupSpaceInput, opts ...grpc.CallOption) (*LookupSpaceOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*LookupSpaceInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *LookupSpaceInput, got %T", in))
		}

		return i.client.LookupSpace(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.universe.Universe.LookupSpace", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*LookupSpaceOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *LookupSpaceOutput, got %T", out))
	}

	return message, err
}

func (i *UniverseInterceptor) CreateSpace(ctx context.Context, in *CreateSpaceInput, opts ...grpc.CallOption) (*CreateSpaceOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateSpaceInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateSpaceInput, got %T", in))
		}

		return i.client.CreateSpace(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.universe.Universe.CreateSpace", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateSpaceOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateSpaceOutput, got %T", out))
	}

	return message, err
}

func (i *UniverseInterceptor) UpdateSpace(ctx context.Context, in *UpdateSpaceInput, opts ...grpc.CallOption) (*UpdateSpaceOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateSpaceInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateSpaceInput, got %T", in))
		}

		return i.client.UpdateSpace(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.universe.Universe.UpdateSpace", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateSpaceOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateSpaceOutput, got %T", out))
	}

	return message, err
}

func (i *UniverseInterceptor) DeleteSpace(ctx context.Context, in *DeleteSpaceInput, opts ...grpc.CallOption) (*DeleteSpaceOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteSpaceInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteSpaceInput, got %T", in))
		}

		return i.client.DeleteSpace(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.universe.Universe.DeleteSpace", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteSpaceOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteSpaceOutput, got %T", out))
	}

	return message, err
}

func (i *UniverseInterceptor) DescribeSpace(ctx context.Context, in *DescribeSpaceInput, opts ...grpc.CallOption) (*DescribeSpaceOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeSpaceInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeSpaceInput, got %T", in))
		}

		return i.client.DescribeSpace(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.universe.Universe.DescribeSpace", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeSpaceOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeSpaceOutput, got %T", out))
	}

	return message, err
}

func (i *UniverseInterceptor) DescribeQuota(ctx context.Context, in *DescribeQuotaInput, opts ...grpc.CallOption) (*DescribeQuotaOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeQuotaInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeQuotaInput, got %T", in))
		}

		return i.client.DescribeQuota(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.universe.Universe.DescribeQuota", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeQuotaOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeQuotaOutput, got %T", out))
	}

	return message, err
}

func (i *UniverseInterceptor) UpdateQuota(ctx context.Context, in *UpdateQuotaInput, opts ...grpc.CallOption) (*UpdateQuotaOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateQuotaInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateQuotaInput, got %T", in))
		}

		return i.client.UpdateQuota(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.universe.Universe.UpdateQuota", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateQuotaOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateQuotaOutput, got %T", out))
	}

	return message, err
}

func (i *UniverseInterceptor) ListSpaces(ctx context.Context, in *ListSpacesInput, opts ...grpc.CallOption) (*ListSpacesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListSpacesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListSpacesInput, got %T", in))
		}

		return i.client.ListSpaces(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.universe.Universe.ListSpaces", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListSpacesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListSpacesOutput, got %T", out))
	}

	return message, err
}
