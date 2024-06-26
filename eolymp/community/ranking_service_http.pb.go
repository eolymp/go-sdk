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
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _RankingService_HTTPReadQueryString parses body into proto.Message
func _RankingService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _RankingService_HTTPReadRequestBody parses body into proto.Message
func _RankingService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _RankingService_HTTPWriteResponse writes proto.Message to HTTP response
func _RankingService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_RankingService_HTTPWriteErrorResponse(w, err)
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

// _RankingService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _RankingService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _RankingService_WebsocketErrorResponse writes error to websocket connection
func _RankingService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _RankingService_WebsocketCodec implements protobuf codec for websockets package
var _RankingService_WebsocketCodec = websocket.Codec{
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

// RegisterRankingServiceHttpHandlers adds handlers for for RankingServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterRankingServiceHttpHandlers(router *mux.Router, prefix string, cli RankingServiceClient) {
	router.Handle(prefix+"/ranking-events", _RankingService_CreateRankingEvent_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.RankingService.CreateRankingEvent")
	router.Handle(prefix+"/ranking-events/{event_id}", _RankingService_UpdateRankingEvent_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.RankingService.UpdateRankingEvent")
	router.Handle(prefix+"/ranking-events/{event_id}", _RankingService_DeleteRankingEvent_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.community.RankingService.DeleteRankingEvent")
	router.Handle(prefix+"/ranking-events/{event_id}", _RankingService_DescribeRankingEvent_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.RankingService.DescribeRankingEvent")
	router.Handle(prefix+"/ranking-events", _RankingService_ListRankingEvents_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.RankingService.ListRankingEvents")
	router.Handle(prefix+"/ranking-series/{member_id}/points/{event_id}", _RankingService_UpdateRankingPoint_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.community.RankingService.UpdateRankingPoint")
	router.Handle(prefix+"/ranking-series/{member_id}/points/{event_id}", _RankingService_DeleteRankingPoint_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.community.RankingService.DeleteRankingPoint")
	router.Handle(prefix+"/ranking-series/{member_id}/points/{event_id}", _RankingService_DescribeRankingPoint_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.RankingService.DescribeRankingPoint")
	router.Handle(prefix+"/ranking-series/{member_id}", _RankingService_ListRankingPoints_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.RankingService.ListRankingPoints")
}

func _RankingService_CreateRankingEvent_Rule0(cli RankingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateRankingEventInput{}

		if err := _RankingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateRankingEvent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RankingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RankingService_UpdateRankingEvent_Rule0(cli RankingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateRankingEventInput{}

		if err := _RankingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EventId = vars["event_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateRankingEvent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RankingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RankingService_DeleteRankingEvent_Rule0(cli RankingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteRankingEventInput{}

		if err := _RankingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EventId = vars["event_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteRankingEvent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RankingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RankingService_DescribeRankingEvent_Rule0(cli RankingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRankingEventInput{}

		if err := _RankingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EventId = vars["event_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeRankingEvent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RankingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RankingService_ListRankingEvents_Rule0(cli RankingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListRankingEventsInput{}

		if err := _RankingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListRankingEvents(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RankingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RankingService_UpdateRankingPoint_Rule0(cli RankingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateRankingPointInput{}

		if err := _RankingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]
		in.EventId = vars["event_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateRankingPoint(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RankingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RankingService_DeleteRankingPoint_Rule0(cli RankingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteRankingPointInput{}

		if err := _RankingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]
		in.EventId = vars["event_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteRankingPoint(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RankingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RankingService_DescribeRankingPoint_Rule0(cli RankingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRankingPointInput{}

		if err := _RankingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]
		in.EventId = vars["event_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeRankingPoint(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RankingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RankingService_ListRankingPoints_Rule0(cli RankingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListRankingPointsInput{}

		if err := _RankingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.ListRankingPoints(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RankingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RankingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _RankingServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _RankingServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _RankingServiceHandler) (out proto.Message, err error)
type RankingServiceInterceptor struct {
	middleware []_RankingServiceMiddleware
	client     RankingServiceClient
}

// NewRankingServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewRankingServiceInterceptor(cli RankingServiceClient, middleware ..._RankingServiceMiddleware) *RankingServiceInterceptor {
	return &RankingServiceInterceptor{client: cli, middleware: middleware}
}

func (i *RankingServiceInterceptor) CreateRankingEvent(ctx context.Context, in *CreateRankingEventInput, opts ...grpc.CallOption) (*CreateRankingEventOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateRankingEventInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateRankingEventInput, got %T", in))
		}

		return i.client.CreateRankingEvent(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.RankingService.CreateRankingEvent", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateRankingEventOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateRankingEventOutput, got %T", out))
	}

	return message, err
}

func (i *RankingServiceInterceptor) UpdateRankingEvent(ctx context.Context, in *UpdateRankingEventInput, opts ...grpc.CallOption) (*UpdateRankingEventOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateRankingEventInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateRankingEventInput, got %T", in))
		}

		return i.client.UpdateRankingEvent(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.RankingService.UpdateRankingEvent", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateRankingEventOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateRankingEventOutput, got %T", out))
	}

	return message, err
}

func (i *RankingServiceInterceptor) DeleteRankingEvent(ctx context.Context, in *DeleteRankingEventInput, opts ...grpc.CallOption) (*DeleteRankingEventOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteRankingEventInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteRankingEventInput, got %T", in))
		}

		return i.client.DeleteRankingEvent(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.RankingService.DeleteRankingEvent", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteRankingEventOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteRankingEventOutput, got %T", out))
	}

	return message, err
}

func (i *RankingServiceInterceptor) DescribeRankingEvent(ctx context.Context, in *DescribeRankingEventInput, opts ...grpc.CallOption) (*DescribeRankingEventOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeRankingEventInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeRankingEventInput, got %T", in))
		}

		return i.client.DescribeRankingEvent(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.RankingService.DescribeRankingEvent", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeRankingEventOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeRankingEventOutput, got %T", out))
	}

	return message, err
}

func (i *RankingServiceInterceptor) ListRankingEvents(ctx context.Context, in *ListRankingEventsInput, opts ...grpc.CallOption) (*ListRankingEventsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListRankingEventsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListRankingEventsInput, got %T", in))
		}

		return i.client.ListRankingEvents(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.RankingService.ListRankingEvents", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListRankingEventsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListRankingEventsOutput, got %T", out))
	}

	return message, err
}

func (i *RankingServiceInterceptor) UpdateRankingPoint(ctx context.Context, in *UpdateRankingPointInput, opts ...grpc.CallOption) (*UpdateRankingPointOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateRankingPointInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateRankingPointInput, got %T", in))
		}

		return i.client.UpdateRankingPoint(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.RankingService.UpdateRankingPoint", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateRankingPointOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateRankingPointOutput, got %T", out))
	}

	return message, err
}

func (i *RankingServiceInterceptor) DeleteRankingPoint(ctx context.Context, in *DeleteRankingPointInput, opts ...grpc.CallOption) (*DeleteRankingPointOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteRankingPointInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteRankingPointInput, got %T", in))
		}

		return i.client.DeleteRankingPoint(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.RankingService.DeleteRankingPoint", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteRankingPointOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteRankingPointOutput, got %T", out))
	}

	return message, err
}

func (i *RankingServiceInterceptor) DescribeRankingPoint(ctx context.Context, in *DescribeRankingPointInput, opts ...grpc.CallOption) (*DescribeRankingPointOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeRankingPointInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeRankingPointInput, got %T", in))
		}

		return i.client.DescribeRankingPoint(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.RankingService.DescribeRankingPoint", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeRankingPointOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeRankingPointOutput, got %T", out))
	}

	return message, err
}

func (i *RankingServiceInterceptor) ListRankingPoints(ctx context.Context, in *ListRankingPointsInput, opts ...grpc.CallOption) (*ListRankingPointsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListRankingPointsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListRankingPointsInput, got %T", in))
		}

		return i.client.ListRankingPoints(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.RankingService.ListRankingPoints", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListRankingPointsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListRankingPointsOutput, got %T", out))
	}

	return message, err
}
