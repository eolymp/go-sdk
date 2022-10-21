// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package keeper

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _Keeper_HTTPReadQueryString parses body into proto.Message
func _Keeper_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Keeper_HTTPReadRequestBody parses body into proto.Message
func _Keeper_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Keeper_HTTPWriteResponse writes proto.Message to HTTP response
func _Keeper_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Keeper_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Keeper_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Keeper_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// NewKeeperHandler constructs new http.Handler for KeeperServer
func NewKeeperHandler(srv KeeperServer) http.Handler {
	router := mux.NewRouter()
	router.Handle("/eolymp.keeper.Keeper/CreateObject", _Keeper_CreateObject(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.keeper.Keeper/DescribeObject", _Keeper_DescribeObject(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.keeper.Keeper/DownloadObject", _Keeper_DownloadObject(srv)).Methods(http.MethodPost)
	return router
}

// NewKeeperHandlerHttp constructs new http.Handler for KeeperServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func NewKeeperHandlerHttp(srv KeeperServer, prefix string) http.Handler {
	router := mux.NewRouter()

	router.Handle(prefix+"/objects", _Keeper_CreateObject_Rule0(srv)).
		Methods("POST").
		Name("eolymp.keeper.Keeper.CreateObject")

	router.Handle(prefix+"/objects/{key}", _Keeper_DescribeObject_Rule0(srv)).
		Methods("GET").
		Name("eolymp.keeper.Keeper.DescribeObject")

	router.Handle(prefix+"/objects/{key}/data", _Keeper_DownloadObject_Rule0(srv)).
		Methods("GET").
		Name("eolymp.keeper.Keeper.DownloadObject")

	return router
}

func _Keeper_CreateObject(srv KeeperServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateObjectInput{}

		if err := _Keeper_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateObject(r.Context(), in)
		if err != nil {
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		_Keeper_HTTPWriteResponse(w, out)
	})
}

func _Keeper_DescribeObject(srv KeeperServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeObjectInput{}

		if err := _Keeper_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeObject(r.Context(), in)
		if err != nil {
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		_Keeper_HTTPWriteResponse(w, out)
	})
}

func _Keeper_DownloadObject(srv KeeperServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DownloadObjectInput{}

		if err := _Keeper_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DownloadObject(r.Context(), in)
		if err != nil {
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		_Keeper_HTTPWriteResponse(w, out)
	})
}

func _Keeper_CreateObject_Rule0(srv KeeperServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateObjectInput{}

		if err := _Keeper_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateObject(r.Context(), in)
		if err != nil {
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		_Keeper_HTTPWriteResponse(w, out)
	})
}

func _Keeper_DescribeObject_Rule0(srv KeeperServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeObjectInput{}

		if err := _Keeper_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.Key = vars["key"]

		out, err := srv.DescribeObject(r.Context(), in)
		if err != nil {
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		_Keeper_HTTPWriteResponse(w, out)
	})
}

func _Keeper_DownloadObject_Rule0(srv KeeperServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DownloadObjectInput{}

		if err := _Keeper_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.Key = vars["key"]

		out, err := srv.DownloadObject(r.Context(), in)
		if err != nil {
			_Keeper_HTTPWriteErrorResponse(w, err)
			return
		}

		_Keeper_HTTPWriteResponse(w, out)
	})
}

type _KeeperHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _KeeperMiddleware = func(ctx context.Context, method string, in proto.Message, handler _KeeperHandler) (out proto.Message, err error)
type KeeperInterceptor struct {
	middleware []_KeeperMiddleware
	server     KeeperServer
}

// NewKeeperInterceptor constructs additional middleware for a server based on annotations in proto files
func NewKeeperInterceptor(srv KeeperServer, middleware ..._KeeperMiddleware) *KeeperInterceptor {
	return &KeeperInterceptor{server: srv, middleware: middleware}
}

func (i *KeeperInterceptor) CreateObject(ctx context.Context, in *CreateObjectInput) (*CreateObjectOutput, error) {
	next := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateObjectInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateObjectInput, got %T", in))
		}

		return i.server.CreateObject(ctx, message)
	}

	for _, mw := range i.middleware {
		handler := next

		next = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.keeper.Keeper.CreateObject", in, handler)
		}
	}

	out, err := next(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateObjectOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateObjectOutput, got %T", out))
	}

	return message, err
}

func (i *KeeperInterceptor) DescribeObject(ctx context.Context, in *DescribeObjectInput) (*DescribeObjectOutput, error) {
	next := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeObjectInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeObjectInput, got %T", in))
		}

		return i.server.DescribeObject(ctx, message)
	}

	for _, mw := range i.middleware {
		handler := next

		next = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.keeper.Keeper.DescribeObject", in, handler)
		}
	}

	out, err := next(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeObjectOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeObjectOutput, got %T", out))
	}

	return message, err
}

func (i *KeeperInterceptor) DownloadObject(ctx context.Context, in *DownloadObjectInput) (*DownloadObjectOutput, error) {
	next := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DownloadObjectInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DownloadObjectInput, got %T", in))
		}

		return i.server.DownloadObject(ctx, message)
	}

	for _, mw := range i.middleware {
		handler := next

		next = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.keeper.Keeper.DownloadObject", in, handler)
		}
	}

	out, err := next(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DownloadObjectOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DownloadObjectOutput, got %T", out))
	}

	return message, err
}
