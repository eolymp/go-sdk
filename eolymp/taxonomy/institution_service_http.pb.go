// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package taxonomy

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

// _InstitutionService_HTTPReadQueryString parses body into proto.Message
func _InstitutionService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _InstitutionService_HTTPReadRequestBody parses body into proto.Message
func _InstitutionService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _InstitutionService_HTTPWriteResponse writes proto.Message to HTTP response
func _InstitutionService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_InstitutionService_HTTPWriteErrorResponse(w, err)
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

// _InstitutionService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _InstitutionService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _InstitutionService_WebsocketErrorResponse writes error to websocket connection
func _InstitutionService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _InstitutionService_WebsocketCodec implements protobuf codec for websockets package
var _InstitutionService_WebsocketCodec = websocket.Codec{
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

// RegisterInstitutionServiceHttpHandlers adds handlers for for InstitutionServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterInstitutionServiceHttpHandlers(router *mux.Router, prefix string, cli InstitutionServiceClient) {
	router.Handle(prefix+"/institutions", _InstitutionService_ListInstitutions_Rule0(cli)).
		Methods("GET").
		Name("eolymp.taxonomy.InstitutionService.ListInstitutions")
	router.Handle(prefix+"/institutions/{institution_id}", _InstitutionService_DescribeInstitution_Rule0(cli)).
		Methods("GET").
		Name("eolymp.taxonomy.InstitutionService.DescribeInstitution")
}

func _InstitutionService_ListInstitutions_Rule0(cli InstitutionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListInstitutionsInput{}

		if err := _InstitutionService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_InstitutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListInstitutions(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_InstitutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_InstitutionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _InstitutionService_DescribeInstitution_Rule0(cli InstitutionServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeInstitutionInput{}

		if err := _InstitutionService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_InstitutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.InstitutionId = vars["institution_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeInstitution(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_InstitutionService_HTTPWriteErrorResponse(w, err)
			return
		}

		_InstitutionService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _InstitutionServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _InstitutionServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _InstitutionServiceHandler) (out proto.Message, err error)
type InstitutionServiceInterceptor struct {
	middleware []_InstitutionServiceMiddleware
	client     InstitutionServiceClient
}

// NewInstitutionServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewInstitutionServiceInterceptor(cli InstitutionServiceClient, middleware ..._InstitutionServiceMiddleware) *InstitutionServiceInterceptor {
	return &InstitutionServiceInterceptor{client: cli, middleware: middleware}
}

func (i *InstitutionServiceInterceptor) ListInstitutions(ctx context.Context, in *ListInstitutionsInput, opts ...grpc.CallOption) (*ListInstitutionsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListInstitutionsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListInstitutionsInput, got %T", in))
		}

		return i.client.ListInstitutions(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.taxonomy.InstitutionService.ListInstitutions", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListInstitutionsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListInstitutionsOutput, got %T", out))
	}

	return message, err
}

func (i *InstitutionServiceInterceptor) DescribeInstitution(ctx context.Context, in *DescribeInstitutionInput, opts ...grpc.CallOption) (*DescribeInstitutionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeInstitutionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeInstitutionInput, got %T", in))
		}

		return i.client.DescribeInstitution(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.taxonomy.InstitutionService.DescribeInstitution", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeInstitutionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeInstitutionOutput, got %T", out))
	}

	return message, err
}
