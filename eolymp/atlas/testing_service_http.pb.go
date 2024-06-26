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
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _TestingService_HTTPReadQueryString parses body into proto.Message
func _TestingService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _TestingService_HTTPReadRequestBody parses body into proto.Message
func _TestingService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _TestingService_HTTPWriteResponse writes proto.Message to HTTP response
func _TestingService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_TestingService_HTTPWriteErrorResponse(w, err)
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

// _TestingService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _TestingService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _TestingService_WebsocketErrorResponse writes error to websocket connection
func _TestingService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _TestingService_WebsocketCodec implements protobuf codec for websockets package
var _TestingService_WebsocketCodec = websocket.Codec{
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

// RegisterTestingServiceHttpHandlers adds handlers for for TestingServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterTestingServiceHttpHandlers(router *mux.Router, prefix string, cli TestingServiceClient) {
	router.Handle(prefix+"/testing", _TestingService_UpdateTestingConfig_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.TestingService.UpdateTestingConfig")
	router.Handle(prefix+"/testing", _TestingService_DescribeTestingConfig_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.TestingService.DescribeTestingConfig")
	router.Handle(prefix+"/checker", _TestingService_UpdateChecker_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.TestingService.UpdateChecker")
	router.Handle(prefix+"/checker", _TestingService_DescribeChecker_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.TestingService.DescribeChecker")
	router.Handle(prefix+"/interactor", _TestingService_UpdateInteractor_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.TestingService.UpdateInteractor")
	router.Handle(prefix+"/interactor", _TestingService_DescribeInteractor_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.TestingService.DescribeInteractor")
	router.Handle(prefix+"/testsets", _TestingService_CreateTestset_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.TestingService.CreateTestset")
	router.Handle(prefix+"/testsets/{testset_id}", _TestingService_UpdateTestset_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.TestingService.UpdateTestset")
	router.Handle(prefix+"/testsets/{testset_id}", _TestingService_DeleteTestset_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.atlas.TestingService.DeleteTestset")
	router.Handle(prefix+"/testsets/{testset_id}", _TestingService_DescribeTestset_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.TestingService.DescribeTestset")
	router.Handle(prefix+"/testsets", _TestingService_ListTestsets_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.TestingService.ListTestsets")
	router.Handle(prefix+"/testsets/{testset_id}/tests", _TestingService_CreateTest_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.TestingService.CreateTest")
	router.Handle(prefix+"/testsets/{testset_id}/tests/{test_id}", _TestingService_UpdateTest_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.TestingService.UpdateTest")
	router.Handle(prefix+"/testsets/{testset_id}/tests/{test_id}", _TestingService_DeleteTest_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.atlas.TestingService.DeleteTest")
	router.Handle(prefix+"/testsets/{testset_id}/tests/{test_id}", _TestingService_DescribeTest_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.TestingService.DescribeTest")
	router.Handle(prefix+"/testsets/{testset_id}/tests", _TestingService_ListTests_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.TestingService.ListTests")
	router.Handle(prefix+"/examples", _TestingService_ListExamples_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.TestingService.ListExamples")
}

func _TestingService_UpdateTestingConfig_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateTestingConfigInput{}

		if err := _TestingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.UpdateTestingConfig(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_DescribeTestingConfig_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeTestingConfigInput{}

		if err := _TestingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeTestingConfig(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_UpdateChecker_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateCheckerInput{}

		if err := _TestingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.UpdateChecker(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_DescribeChecker_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCheckerInput{}

		if err := _TestingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeChecker(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_UpdateInteractor_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateInteractorInput{}

		if err := _TestingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.UpdateInteractor(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_DescribeInteractor_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeInteractorInput{}

		if err := _TestingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeInteractor(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_CreateTestset_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateTestsetInput{}

		if err := _TestingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateTestset(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_UpdateTestset_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateTestsetInput{}

		if err := _TestingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TestsetId = vars["testset_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateTestset(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_DeleteTestset_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteTestsetInput{}

		if err := _TestingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TestsetId = vars["testset_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteTestset(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_DescribeTestset_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeTestsetInput{}

		if err := _TestingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TestsetId = vars["testset_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeTestset(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_ListTestsets_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListTestsetsInput{}

		if err := _TestingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListTestsets(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_CreateTest_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateTestInput{}

		if err := _TestingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TestsetId = vars["testset_id"]

		var header, trailer metadata.MD

		out, err := cli.CreateTest(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_UpdateTest_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateTestInput{}

		if err := _TestingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TestsetId = vars["testset_id"]
		in.TestId = vars["test_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateTest(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_DeleteTest_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteTestInput{}

		if err := _TestingService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TestsetId = vars["testset_id"]
		in.TestId = vars["test_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteTest(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_DescribeTest_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeTestInput{}

		if err := _TestingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TestsetId = vars["testset_id"]
		in.TestId = vars["test_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeTest(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_ListTests_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListTestsInput{}

		if err := _TestingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TestsetId = vars["testset_id"]

		var header, trailer metadata.MD

		out, err := cli.ListTests(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _TestingService_ListExamples_Rule0(cli TestingServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListExamplesInput{}

		if err := _TestingService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListExamples(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_TestingService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TestingService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _TestingServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _TestingServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _TestingServiceHandler) (out proto.Message, err error)
type TestingServiceInterceptor struct {
	middleware []_TestingServiceMiddleware
	client     TestingServiceClient
}

// NewTestingServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewTestingServiceInterceptor(cli TestingServiceClient, middleware ..._TestingServiceMiddleware) *TestingServiceInterceptor {
	return &TestingServiceInterceptor{client: cli, middleware: middleware}
}

func (i *TestingServiceInterceptor) UpdateTestingConfig(ctx context.Context, in *UpdateTestingConfigInput, opts ...grpc.CallOption) (*UpdateTestingConfigOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateTestingConfigInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateTestingConfigInput, got %T", in))
		}

		return i.client.UpdateTestingConfig(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.UpdateTestingConfig", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateTestingConfigOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateTestingConfigOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) DescribeTestingConfig(ctx context.Context, in *DescribeTestingConfigInput, opts ...grpc.CallOption) (*DescribeTestingConfigOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeTestingConfigInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeTestingConfigInput, got %T", in))
		}

		return i.client.DescribeTestingConfig(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.DescribeTestingConfig", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeTestingConfigOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeTestingConfigOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) UpdateChecker(ctx context.Context, in *UpdateCheckerInput, opts ...grpc.CallOption) (*UpdateCheckerOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateCheckerInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateCheckerInput, got %T", in))
		}

		return i.client.UpdateChecker(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.UpdateChecker", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateCheckerOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateCheckerOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) DescribeChecker(ctx context.Context, in *DescribeCheckerInput, opts ...grpc.CallOption) (*DescribeCheckerOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeCheckerInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeCheckerInput, got %T", in))
		}

		return i.client.DescribeChecker(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.DescribeChecker", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeCheckerOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeCheckerOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) UpdateInteractor(ctx context.Context, in *UpdateInteractorInput, opts ...grpc.CallOption) (*UpdateInteractorOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateInteractorInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateInteractorInput, got %T", in))
		}

		return i.client.UpdateInteractor(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.UpdateInteractor", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateInteractorOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateInteractorOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) DescribeInteractor(ctx context.Context, in *DescribeInteractorInput, opts ...grpc.CallOption) (*DescribeInteractorOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeInteractorInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeInteractorInput, got %T", in))
		}

		return i.client.DescribeInteractor(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.DescribeInteractor", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeInteractorOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeInteractorOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) CreateTestset(ctx context.Context, in *CreateTestsetInput, opts ...grpc.CallOption) (*CreateTestsetOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateTestsetInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateTestsetInput, got %T", in))
		}

		return i.client.CreateTestset(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.CreateTestset", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateTestsetOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateTestsetOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) UpdateTestset(ctx context.Context, in *UpdateTestsetInput, opts ...grpc.CallOption) (*UpdateTestsetOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateTestsetInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateTestsetInput, got %T", in))
		}

		return i.client.UpdateTestset(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.UpdateTestset", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateTestsetOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateTestsetOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) DeleteTestset(ctx context.Context, in *DeleteTestsetInput, opts ...grpc.CallOption) (*DeleteTestsetOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteTestsetInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteTestsetInput, got %T", in))
		}

		return i.client.DeleteTestset(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.DeleteTestset", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteTestsetOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteTestsetOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) DescribeTestset(ctx context.Context, in *DescribeTestsetInput, opts ...grpc.CallOption) (*DescribeTestsetOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeTestsetInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeTestsetInput, got %T", in))
		}

		return i.client.DescribeTestset(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.DescribeTestset", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeTestsetOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeTestsetOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) ListTestsets(ctx context.Context, in *ListTestsetsInput, opts ...grpc.CallOption) (*ListTestsetsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListTestsetsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListTestsetsInput, got %T", in))
		}

		return i.client.ListTestsets(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.ListTestsets", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListTestsetsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListTestsetsOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) CreateTest(ctx context.Context, in *CreateTestInput, opts ...grpc.CallOption) (*CreateTestOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateTestInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateTestInput, got %T", in))
		}

		return i.client.CreateTest(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.CreateTest", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateTestOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateTestOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) UpdateTest(ctx context.Context, in *UpdateTestInput, opts ...grpc.CallOption) (*UpdateTestOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateTestInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateTestInput, got %T", in))
		}

		return i.client.UpdateTest(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.UpdateTest", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateTestOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateTestOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) DeleteTest(ctx context.Context, in *DeleteTestInput, opts ...grpc.CallOption) (*DeleteTestOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteTestInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteTestInput, got %T", in))
		}

		return i.client.DeleteTest(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.DeleteTest", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteTestOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteTestOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) DescribeTest(ctx context.Context, in *DescribeTestInput, opts ...grpc.CallOption) (*DescribeTestOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeTestInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeTestInput, got %T", in))
		}

		return i.client.DescribeTest(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.DescribeTest", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeTestOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeTestOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) ListTests(ctx context.Context, in *ListTestsInput, opts ...grpc.CallOption) (*ListTestsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListTestsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListTestsInput, got %T", in))
		}

		return i.client.ListTests(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.ListTests", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListTestsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListTestsOutput, got %T", out))
	}

	return message, err
}

func (i *TestingServiceInterceptor) ListExamples(ctx context.Context, in *ListExamplesInput, opts ...grpc.CallOption) (*ListExamplesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListExamplesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListExamplesInput, got %T", in))
		}

		return i.client.ListExamples(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.TestingService.ListExamples", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListExamplesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListExamplesOutput, got %T", out))
	}

	return message, err
}
