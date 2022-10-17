// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package resolver

import (
	context "context"
	mux "github.com/gorilla/mux"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _Resolver_HTTPReadQueryString parses body into proto.Message
func _Resolver_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Resolver_HTTPReadRequestBody parses body into proto.Message
func _Resolver_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Resolver_HTTPWriteResponse writes proto.Message to HTTP response
func _Resolver_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Resolver_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Resolver_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Resolver_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// NewResolverHandler constructs new http.Handler for ResolverServer
func NewResolverHandler(srv ResolverServer) http.Handler {
	router := mux.NewRouter()
	router.Handle("/eolymp.resolver.Resolver/ResolveName", _Resolver_ResolveName(srv)).Methods(http.MethodPost)
	return router
}

// NewResolverHandlerHttp constructs new http.Handler for ResolverServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func NewResolverHandlerHttp(srv ResolverServer, prefix string) http.Handler {
	router := mux.NewRouter()

	router.Handle(prefix+"/names/{name}", _Resolver_ResolveName_Rule0(srv)).
		Methods("GET").
		Name("eolymp.resolver.Resolver.ResolveName")

	return router
}

func _Resolver_ResolveName(srv ResolverServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ResolveNameInput{}

		if err := _Resolver_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Resolver_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ResolveName(r.Context(), in)
		if err != nil {
			_Resolver_HTTPWriteErrorResponse(w, err)
			return
		}

		_Resolver_HTTPWriteResponse(w, out)
	})
}

func _Resolver_ResolveName_Rule0(srv ResolverServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ResolveNameInput{}

		if err := _Resolver_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Resolver_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.Name = vars["name"]

		out, err := srv.ResolveName(r.Context(), in)
		if err != nil {
			_Resolver_HTTPWriteErrorResponse(w, err)
			return
		}

		_Resolver_HTTPWriteResponse(w, out)
	})
}

type _ResolverMiddleware func(ctx context.Context, method string, in proto.Message, next func() (out proto.Message, err error))
type ResolverInterceptor struct {
	middleware []_ResolverMiddleware
	server     ResolverServer
}

// NewResolverInterceptor constructs additional middleware for a server based on annotations in proto files
func NewResolverInterceptor(srv ResolverServer, middleware ..._ResolverMiddleware) *ResolverInterceptor {
	return &ResolverInterceptor{server: srv, middleware: middleware}
}

func (i *ResolverInterceptor) ResolveName(ctx context.Context, in *ResolveNameInput) (out *ResolveNameOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.ResolveName(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.resolver.Resolver/ResolveName", in, handler)
			return out, err
		}
	}

	next()
	return
}
