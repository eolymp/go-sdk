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
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _PlanService_HTTPReadQueryString parses body into proto.Message
func _PlanService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _PlanService_HTTPReadRequestBody parses body into proto.Message
func _PlanService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _PlanService_HTTPWriteResponse writes proto.Message to HTTP response
func _PlanService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_PlanService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _PlanService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _PlanService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _PlanService_WebsocketErrorResponse writes error to websocket connection
func _PlanService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _PlanService_WebsocketCodec implements protobuf codec for websockets package
var _PlanService_WebsocketCodec = websocket.Codec{
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

// RegisterPlanServiceHttpHandlers adds handlers for for PlanServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterPlanServiceHttpHandlers(router *mux.Router, prefix string, cli PlanServiceClient) {
	router.Handle(prefix+"/plans/{plan_id}", _PlanService_DescribePlan_Rule0(cli)).
		Methods("GET").
		Name("eolymp.universe.PlanService.DescribePlan")
	router.Handle(prefix+"/plans", _PlanService_ListPlans_Rule0(cli)).
		Methods("GET").
		Name("eolymp.universe.PlanService.ListPlans")
}

func _PlanService_DescribePlan_Rule0(cli PlanServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePlanInput{}

		if err := _PlanService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PlanService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PlanId = vars["plan_id"]

		out, err := cli.DescribePlan(r.Context(), in)
		if err != nil {
			_PlanService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PlanService_HTTPWriteResponse(w, out)
	})
}

func _PlanService_ListPlans_Rule0(cli PlanServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPlansInput{}

		if err := _PlanService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PlanService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListPlans(r.Context(), in)
		if err != nil {
			_PlanService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PlanService_HTTPWriteResponse(w, out)
	})
}

type _PlanServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _PlanServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _PlanServiceHandler) (out proto.Message, err error)
type PlanServiceInterceptor struct {
	middleware []_PlanServiceMiddleware
	client     PlanServiceClient
}

// NewPlanServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewPlanServiceInterceptor(cli PlanServiceClient, middleware ..._PlanServiceMiddleware) *PlanServiceInterceptor {
	return &PlanServiceInterceptor{client: cli, middleware: middleware}
}

func (i *PlanServiceInterceptor) DescribePlan(ctx context.Context, in *DescribePlanInput, opts ...grpc.CallOption) (*DescribePlanOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribePlanInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribePlanInput, got %T", in))
		}

		return i.client.DescribePlan(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.universe.PlanService.DescribePlan", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribePlanOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribePlanOutput, got %T", out))
	}

	return message, err
}

func (i *PlanServiceInterceptor) ListPlans(ctx context.Context, in *ListPlansInput, opts ...grpc.CallOption) (*ListPlansOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListPlansInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListPlansInput, got %T", in))
		}

		return i.client.ListPlans(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.universe.PlanService.ListPlans", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListPlansOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListPlansOutput, got %T", out))
	}

	return message, err
}