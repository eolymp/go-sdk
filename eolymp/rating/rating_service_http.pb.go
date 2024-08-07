// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package rating

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

// _RatingService_HTTPReadQueryString parses body into proto.Message
func _RatingService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _RatingService_HTTPReadRequestBody parses body into proto.Message
func _RatingService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _RatingService_HTTPWriteResponse writes proto.Message to HTTP response
func _RatingService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_RatingService_HTTPWriteErrorResponse(w, err)
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

// _RatingService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _RatingService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _RatingService_WebsocketErrorResponse writes error to websocket connection
func _RatingService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _RatingService_WebsocketCodec implements protobuf codec for websockets package
var _RatingService_WebsocketCodec = websocket.Codec{
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

// RegisterRatingServiceHttpHandlers adds handlers for for RatingServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterRatingServiceHttpHandlers(router *mux.Router, prefix string, cli RatingServiceClient) {
	router.Handle(prefix+"/rating", _RatingService_SetRating_Rule0(cli)).
		Methods("POST").
		Name("eolymp.rating.RatingService.SetRating")
	router.Handle(prefix+"/rating/{rating_id}", _RatingService_UpdateRating_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.rating.RatingService.UpdateRating")
	router.Handle(prefix+"/rating/{rating_id}", _RatingService_DeleteRating_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.rating.RatingService.DeleteRating")
	router.Handle(prefix+"/rating/{rating_id}", _RatingService_DescribeRating_Rule0(cli)).
		Methods("GET").
		Name("eolymp.rating.RatingService.DescribeRating")
	router.Handle(prefix+"/members/{member_id}/rating", _RatingService_ListRating_Rule0(cli)).
		Methods("GET").
		Name("eolymp.rating.RatingService.ListRating")
	router.Handle(prefix+"/rating-boundaries", _RatingService_DescribeRatingBoundaries_Rule0(cli)).
		Methods("GET").
		Name("eolymp.rating.RatingService.DescribeRatingBoundaries")
	router.Handle(prefix+"/rating-distribution", _RatingService_DescribeRatingDistribution_Rule0(cli)).
		Methods("GET").
		Name("eolymp.rating.RatingService.DescribeRatingDistribution")
}

func _RatingService_SetRating_Rule0(cli RatingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &SetRatingInput{}

		if err := _RatingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.SetRating(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RatingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RatingService_UpdateRating_Rule0(cli RatingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateRatingInput{}

		if err := _RatingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RatingId = vars["rating_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateRating(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RatingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RatingService_DeleteRating_Rule0(cli RatingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteRatingInput{}

		if err := _RatingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RatingId = vars["rating_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteRating(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RatingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RatingService_DescribeRating_Rule0(cli RatingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRatingInput{}

		if err := _RatingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RatingId = vars["rating_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeRating(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RatingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RatingService_ListRating_Rule0(cli RatingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListRatingInput{}

		if err := _RatingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.ListRating(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RatingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RatingService_DescribeRatingBoundaries_Rule0(cli RatingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRatingBoundariesInput{}

		if err := _RatingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeRatingBoundaries(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RatingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RatingService_DescribeRatingDistribution_Rule0(cli RatingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRatingDistributionInput{}

		if err := _RatingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeRatingDistribution(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RatingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RatingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _RatingServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _RatingServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _RatingServiceHandler) (out proto.Message, err error)
type RatingServiceInterceptor struct {
	middleware []_RatingServiceMiddleware
	client     RatingServiceClient
}

// NewRatingServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewRatingServiceInterceptor(cli RatingServiceClient, middleware ..._RatingServiceMiddleware) *RatingServiceInterceptor {
	return &RatingServiceInterceptor{client: cli, middleware: middleware}
}

func (i *RatingServiceInterceptor) SetRating(ctx context.Context, in *SetRatingInput, opts ...grpc.CallOption) (*SetRatingOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*SetRatingInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *SetRatingInput, got %T", in))
		}

		return i.client.SetRating(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.rating.RatingService.SetRating", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*SetRatingOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *SetRatingOutput, got %T", out))
	}

	return message, err
}

func (i *RatingServiceInterceptor) UpdateRating(ctx context.Context, in *UpdateRatingInput, opts ...grpc.CallOption) (*UpdateRatingOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateRatingInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateRatingInput, got %T", in))
		}

		return i.client.UpdateRating(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.rating.RatingService.UpdateRating", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateRatingOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateRatingOutput, got %T", out))
	}

	return message, err
}

func (i *RatingServiceInterceptor) DeleteRating(ctx context.Context, in *DeleteRatingInput, opts ...grpc.CallOption) (*DeleteRatingOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteRatingInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteRatingInput, got %T", in))
		}

		return i.client.DeleteRating(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.rating.RatingService.DeleteRating", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteRatingOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteRatingOutput, got %T", out))
	}

	return message, err
}

func (i *RatingServiceInterceptor) DescribeRating(ctx context.Context, in *DescribeRatingInput, opts ...grpc.CallOption) (*DescribeRatingOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeRatingInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeRatingInput, got %T", in))
		}

		return i.client.DescribeRating(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.rating.RatingService.DescribeRating", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeRatingOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeRatingOutput, got %T", out))
	}

	return message, err
}

func (i *RatingServiceInterceptor) ListRating(ctx context.Context, in *ListRatingInput, opts ...grpc.CallOption) (*ListRatingOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListRatingInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListRatingInput, got %T", in))
		}

		return i.client.ListRating(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.rating.RatingService.ListRating", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListRatingOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListRatingOutput, got %T", out))
	}

	return message, err
}

func (i *RatingServiceInterceptor) DescribeRatingBoundaries(ctx context.Context, in *DescribeRatingBoundariesInput, opts ...grpc.CallOption) (*DescribeRatingBoundariesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeRatingBoundariesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeRatingBoundariesInput, got %T", in))
		}

		return i.client.DescribeRatingBoundaries(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.rating.RatingService.DescribeRatingBoundaries", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeRatingBoundariesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeRatingBoundariesOutput, got %T", out))
	}

	return message, err
}

func (i *RatingServiceInterceptor) DescribeRatingDistribution(ctx context.Context, in *DescribeRatingDistributionInput, opts ...grpc.CallOption) (*DescribeRatingDistributionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeRatingDistributionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeRatingDistributionInput, got %T", in))
		}

		return i.client.DescribeRatingDistribution(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.rating.RatingService.DescribeRatingDistribution", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeRatingDistributionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeRatingDistributionOutput, got %T", out))
	}

	return message, err
}
