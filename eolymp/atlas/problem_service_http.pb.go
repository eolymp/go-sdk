// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package atlas

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

// _ProblemService_HTTPReadQueryString parses body into proto.Message
func _ProblemService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _ProblemService_HTTPReadRequestBody parses body into proto.Message
func _ProblemService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _ProblemService_HTTPWriteResponse writes proto.Message to HTTP response
func _ProblemService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_ProblemService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _ProblemService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _ProblemService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _ProblemService_WebsocketErrorResponse writes error to websocket connection
func _ProblemService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _ProblemService_WebsocketCodec implements protobuf codec for websockets package
var _ProblemService_WebsocketCodec = websocket.Codec{
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

// RegisterProblemServiceHttpHandlers adds handlers for for ProblemServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterProblemServiceHttpHandlers(router *mux.Router, prefix string, cli ProblemServiceClient) {
	router.Handle(prefix+"", _ProblemService_DeleteProblem_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.atlas.ProblemService.DeleteProblem")
	router.Handle(prefix+"", _ProblemService_UpdateProblem_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.ProblemService.UpdateProblem")
	router.Handle(prefix+"/sync", _ProblemService_SyncProblem_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.ProblemService.SyncProblem")
	router.Handle(prefix+"", _ProblemService_DescribeProblem_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.ProblemService.DescribeProblem")
	router.Handle(prefix+"/visibility", _ProblemService_UpdateVisibility_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.ProblemService.UpdateVisibility")
	router.Handle(prefix+"/privacy", _ProblemService_UpdatePrivacy_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.ProblemService.UpdatePrivacy")
	router.Handle(prefix+"/versions", _ProblemService_ListVersions_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.ProblemService.ListVersions")
}

func _ProblemService_DeleteProblem_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteProblemInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.DeleteProblem(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_UpdateProblem_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateProblemInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.UpdateProblem(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_SyncProblem_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &SyncProblemInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.SyncProblem(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_DescribeProblem_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeProblemInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.DescribeProblem(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_UpdateVisibility_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateVisibilityInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.UpdateVisibility(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_UpdatePrivacy_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdatePrivacyInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.UpdatePrivacy(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

func _ProblemService_ListVersions_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListVersionsInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListVersions(r.Context(), in)
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out)
	})
}

type _ProblemServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _ProblemServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _ProblemServiceHandler) (out proto.Message, err error)
type ProblemServiceInterceptor struct {
	middleware []_ProblemServiceMiddleware
	client     ProblemServiceClient
}

// NewProblemServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewProblemServiceInterceptor(cli ProblemServiceClient, middleware ..._ProblemServiceMiddleware) *ProblemServiceInterceptor {
	return &ProblemServiceInterceptor{client: cli, middleware: middleware}
}

func (i *ProblemServiceInterceptor) DeleteProblem(ctx context.Context, in *DeleteProblemInput, opts ...grpc.CallOption) (*DeleteProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteProblemInput, got %T", in))
		}

		return i.client.DeleteProblem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.DeleteProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) UpdateProblem(ctx context.Context, in *UpdateProblemInput, opts ...grpc.CallOption) (*UpdateProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateProblemInput, got %T", in))
		}

		return i.client.UpdateProblem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.UpdateProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) SyncProblem(ctx context.Context, in *SyncProblemInput, opts ...grpc.CallOption) (*SyncProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*SyncProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *SyncProblemInput, got %T", in))
		}

		return i.client.SyncProblem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.SyncProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*SyncProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *SyncProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) DescribeProblem(ctx context.Context, in *DescribeProblemInput, opts ...grpc.CallOption) (*DescribeProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeProblemInput, got %T", in))
		}

		return i.client.DescribeProblem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.DescribeProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) UpdateVisibility(ctx context.Context, in *UpdateVisibilityInput, opts ...grpc.CallOption) (*UpdateVisibilityOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateVisibilityInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateVisibilityInput, got %T", in))
		}

		return i.client.UpdateVisibility(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.UpdateVisibility", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateVisibilityOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateVisibilityOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) UpdatePrivacy(ctx context.Context, in *UpdatePrivacyInput, opts ...grpc.CallOption) (*UpdatePrivacyOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdatePrivacyInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdatePrivacyInput, got %T", in))
		}

		return i.client.UpdatePrivacy(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.UpdatePrivacy", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdatePrivacyOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdatePrivacyOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) ListVersions(ctx context.Context, in *ListVersionsInput, opts ...grpc.CallOption) (*ListVersionsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListVersionsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListVersionsInput, got %T", in))
		}

		return i.client.ListVersions(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.ProblemService.ListVersions", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListVersionsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListVersionsOutput, got %T", out))
	}

	return message, err
}
