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

// _StatementService_HTTPReadQueryString parses body into proto.Message
func _StatementService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _StatementService_HTTPReadRequestBody parses body into proto.Message
func _StatementService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _StatementService_HTTPWriteResponse writes proto.Message to HTTP response
func _StatementService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_StatementService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _StatementService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _StatementService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _StatementService_WebsocketErrorResponse writes error to websocket connection
func _StatementService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _StatementService_WebsocketCodec implements protobuf codec for websockets package
var _StatementService_WebsocketCodec = websocket.Codec{
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

// RegisterStatementServiceHttpHandlers adds handlers for for StatementServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterStatementServiceHttpHandlers(router *mux.Router, prefix string, cli StatementServiceClient) {
	router.Handle(prefix+"/statements", _StatementService_CreateStatement_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.StatementService.CreateStatement")
	router.Handle(prefix+"/statements/{statement_id}", _StatementService_UpdateStatement_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.StatementService.UpdateStatement")
	router.Handle(prefix+"/statements/{statement_id}", _StatementService_DeleteStatement_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.atlas.StatementService.DeleteStatement")
	router.Handle(prefix+"/statements/{statement_id}", _StatementService_DescribeStatement_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.StatementService.DescribeStatement")
	router.Handle(prefix+"/translate", _StatementService_LookupStatement_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.StatementService.LookupStatement")
	router.Handle(prefix+"/renders", _StatementService_PreviewStatement_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.StatementService.PreviewStatement")
	router.Handle(prefix+"/statements", _StatementService_ListStatements_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.StatementService.ListStatements")
}

func _StatementService_CreateStatement_Rule0(cli StatementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateStatementInput{}

		if err := _StatementService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.CreateStatement(r.Context(), in)
		if err != nil {
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StatementService_HTTPWriteResponse(w, out)
	})
}

func _StatementService_UpdateStatement_Rule0(cli StatementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateStatementInput{}

		if err := _StatementService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.StatementId = vars["statement_id"]

		out, err := cli.UpdateStatement(r.Context(), in)
		if err != nil {
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StatementService_HTTPWriteResponse(w, out)
	})
}

func _StatementService_DeleteStatement_Rule0(cli StatementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteStatementInput{}

		if err := _StatementService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.StatementId = vars["statement_id"]

		out, err := cli.DeleteStatement(r.Context(), in)
		if err != nil {
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StatementService_HTTPWriteResponse(w, out)
	})
}

func _StatementService_DescribeStatement_Rule0(cli StatementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeStatementInput{}

		if err := _StatementService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.StatementId = vars["statement_id"]

		out, err := cli.DescribeStatement(r.Context(), in)
		if err != nil {
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StatementService_HTTPWriteResponse(w, out)
	})
}

func _StatementService_LookupStatement_Rule0(cli StatementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LookupStatementInput{}

		if err := _StatementService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.LookupStatement(r.Context(), in)
		if err != nil {
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StatementService_HTTPWriteResponse(w, out)
	})
}

func _StatementService_PreviewStatement_Rule0(cli StatementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &PreviewStatementInput{}

		if err := _StatementService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.PreviewStatement(r.Context(), in)
		if err != nil {
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StatementService_HTTPWriteResponse(w, out)
	})
}

func _StatementService_ListStatements_Rule0(cli StatementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListStatementsInput{}

		if err := _StatementService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.ListStatements(r.Context(), in)
		if err != nil {
			_StatementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StatementService_HTTPWriteResponse(w, out)
	})
}

type _StatementServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _StatementServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _StatementServiceHandler) (out proto.Message, err error)
type StatementServiceInterceptor struct {
	middleware []_StatementServiceMiddleware
	client     StatementServiceClient
}

// NewStatementServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewStatementServiceInterceptor(cli StatementServiceClient, middleware ..._StatementServiceMiddleware) *StatementServiceInterceptor {
	return &StatementServiceInterceptor{client: cli, middleware: middleware}
}

func (i *StatementServiceInterceptor) CreateStatement(ctx context.Context, in *CreateStatementInput, opts ...grpc.CallOption) (*CreateStatementOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateStatementInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateStatementInput, got %T", in))
		}

		return i.client.CreateStatement(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.StatementService.CreateStatement", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateStatementOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateStatementOutput, got %T", out))
	}

	return message, err
}

func (i *StatementServiceInterceptor) UpdateStatement(ctx context.Context, in *UpdateStatementInput, opts ...grpc.CallOption) (*UpdateStatementOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateStatementInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateStatementInput, got %T", in))
		}

		return i.client.UpdateStatement(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.StatementService.UpdateStatement", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateStatementOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateStatementOutput, got %T", out))
	}

	return message, err
}

func (i *StatementServiceInterceptor) DeleteStatement(ctx context.Context, in *DeleteStatementInput, opts ...grpc.CallOption) (*DeleteStatementOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteStatementInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteStatementInput, got %T", in))
		}

		return i.client.DeleteStatement(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.StatementService.DeleteStatement", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteStatementOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteStatementOutput, got %T", out))
	}

	return message, err
}

func (i *StatementServiceInterceptor) DescribeStatement(ctx context.Context, in *DescribeStatementInput, opts ...grpc.CallOption) (*DescribeStatementOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeStatementInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeStatementInput, got %T", in))
		}

		return i.client.DescribeStatement(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.StatementService.DescribeStatement", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeStatementOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeStatementOutput, got %T", out))
	}

	return message, err
}

func (i *StatementServiceInterceptor) LookupStatement(ctx context.Context, in *LookupStatementInput, opts ...grpc.CallOption) (*LookupStatementOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*LookupStatementInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *LookupStatementInput, got %T", in))
		}

		return i.client.LookupStatement(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.StatementService.LookupStatement", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*LookupStatementOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *LookupStatementOutput, got %T", out))
	}

	return message, err
}

func (i *StatementServiceInterceptor) PreviewStatement(ctx context.Context, in *PreviewStatementInput, opts ...grpc.CallOption) (*PreviewStatementOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*PreviewStatementInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *PreviewStatementInput, got %T", in))
		}

		return i.client.PreviewStatement(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.StatementService.PreviewStatement", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*PreviewStatementOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *PreviewStatementOutput, got %T", out))
	}

	return message, err
}

func (i *StatementServiceInterceptor) ListStatements(ctx context.Context, in *ListStatementsInput, opts ...grpc.CallOption) (*ListStatementsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListStatementsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListStatementsInput, got %T", in))
		}

		return i.client.ListStatements(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.StatementService.ListStatements", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListStatementsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListStatementsOutput, got %T", out))
	}

	return message, err
}
