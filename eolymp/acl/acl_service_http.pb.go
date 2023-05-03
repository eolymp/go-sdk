// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package acl

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

// _Acl_HTTPReadQueryString parses body into proto.Message
func _Acl_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Acl_HTTPReadRequestBody parses body into proto.Message
func _Acl_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Acl_HTTPWriteResponse writes proto.Message to HTTP response
func _Acl_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Acl_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Acl_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Acl_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterAclHttpHandlers adds handlers for for AclServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterAclHttpHandlers(router *mux.Router, prefix string, srv AclServer) {
	router.Handle(prefix+"/acl/{user_id}", _Acl_GrantPermission_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.acl.Acl.GrantPermission")
	router.Handle(prefix+"/acl/{user_id}", _Acl_RevokePermission_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.acl.Acl.RevokePermission")
	router.Handle(prefix+"/acl/{user_id}", _Acl_DescribePermission_Rule0(srv)).
		Methods("GET").
		Name("eolymp.acl.Acl.DescribePermission")
	router.Handle(prefix+"/acl", _Acl_ListPermissions_Rule0(srv)).
		Methods("GET").
		Name("eolymp.acl.Acl.ListPermissions")
	router.Handle(prefix+"/whoami", _Acl_IntrospectPermission_Rule0(srv)).
		Methods("GET").
		Name("eolymp.acl.Acl.IntrospectPermission")
}

func _Acl_GrantPermission_Rule0(srv AclServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &GrantPermissionInput{}

		if err := _Acl_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Acl_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.UserId = vars["user_id"]

		out, err := srv.GrantPermission(r.Context(), in)
		if err != nil {
			_Acl_HTTPWriteErrorResponse(w, err)
			return
		}

		_Acl_HTTPWriteResponse(w, out)
	})
}

func _Acl_RevokePermission_Rule0(srv AclServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RevokePermissionInput{}

		if err := _Acl_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Acl_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.UserId = vars["user_id"]

		out, err := srv.RevokePermission(r.Context(), in)
		if err != nil {
			_Acl_HTTPWriteErrorResponse(w, err)
			return
		}

		_Acl_HTTPWriteResponse(w, out)
	})
}

func _Acl_DescribePermission_Rule0(srv AclServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePermissionInput{}

		if err := _Acl_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Acl_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.UserId = vars["user_id"]

		out, err := srv.DescribePermission(r.Context(), in)
		if err != nil {
			_Acl_HTTPWriteErrorResponse(w, err)
			return
		}

		_Acl_HTTPWriteResponse(w, out)
	})
}

func _Acl_ListPermissions_Rule0(srv AclServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPermissionsInput{}

		if err := _Acl_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Acl_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListPermissions(r.Context(), in)
		if err != nil {
			_Acl_HTTPWriteErrorResponse(w, err)
			return
		}

		_Acl_HTTPWriteResponse(w, out)
	})
}

func _Acl_IntrospectPermission_Rule0(srv AclServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &IntrospectPermissionInput{}

		if err := _Acl_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Acl_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.IntrospectPermission(r.Context(), in)
		if err != nil {
			_Acl_HTTPWriteErrorResponse(w, err)
			return
		}

		_Acl_HTTPWriteResponse(w, out)
	})
}

type _AclHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _AclMiddleware = func(ctx context.Context, method string, in proto.Message, handler _AclHandler) (out proto.Message, err error)
type AclInterceptor struct {
	middleware []_AclMiddleware
	server     AclServer
}

// NewAclInterceptor constructs additional middleware for a server based on annotations in proto files
func NewAclInterceptor(srv AclServer, middleware ..._AclMiddleware) *AclInterceptor {
	return &AclInterceptor{server: srv, middleware: middleware}
}

func (i *AclInterceptor) GrantPermission(ctx context.Context, in *GrantPermissionInput) (*GrantPermissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*GrantPermissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *GrantPermissionInput, got %T", in))
		}

		return i.server.GrantPermission(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.acl.Acl.GrantPermission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*GrantPermissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *GrantPermissionOutput, got %T", out))
	}

	return message, err
}

func (i *AclInterceptor) RevokePermission(ctx context.Context, in *RevokePermissionInput) (*RevokePermissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RevokePermissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RevokePermissionInput, got %T", in))
		}

		return i.server.RevokePermission(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.acl.Acl.RevokePermission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RevokePermissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RevokePermissionOutput, got %T", out))
	}

	return message, err
}

func (i *AclInterceptor) DescribePermission(ctx context.Context, in *DescribePermissionInput) (*DescribePermissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribePermissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribePermissionInput, got %T", in))
		}

		return i.server.DescribePermission(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.acl.Acl.DescribePermission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribePermissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribePermissionOutput, got %T", out))
	}

	return message, err
}

func (i *AclInterceptor) ListPermissions(ctx context.Context, in *ListPermissionsInput) (*ListPermissionsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListPermissionsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListPermissionsInput, got %T", in))
		}

		return i.server.ListPermissions(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.acl.Acl.ListPermissions", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListPermissionsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListPermissionsOutput, got %T", out))
	}

	return message, err
}

func (i *AclInterceptor) IntrospectPermission(ctx context.Context, in *IntrospectPermissionInput) (*IntrospectPermissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*IntrospectPermissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *IntrospectPermissionInput, got %T", in))
		}

		return i.server.IntrospectPermission(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.acl.Acl.IntrospectPermission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*IntrospectPermissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *IntrospectPermissionOutput, got %T", out))
	}

	return message, err
}

// _AclService_HTTPReadQueryString parses body into proto.Message
func _AclService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _AclService_HTTPReadRequestBody parses body into proto.Message
func _AclService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _AclService_HTTPWriteResponse writes proto.Message to HTTP response
func _AclService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_AclService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _AclService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _AclService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterAclServiceHttpHandlers adds handlers for for AclServiceServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterAclServiceHttpHandlers(router *mux.Router, prefix string, srv AclServiceServer) {
	router.Handle(prefix+"/acl/{user_id}", _AclService_GrantPermission_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.acl.AclService.GrantPermission")
	router.Handle(prefix+"/acl/{user_id}", _AclService_RevokePermission_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.acl.AclService.RevokePermission")
	router.Handle(prefix+"/acl/{user_id}", _AclService_DescribePermission_Rule0(srv)).
		Methods("GET").
		Name("eolymp.acl.AclService.DescribePermission")
	router.Handle(prefix+"/acl", _AclService_ListPermissions_Rule0(srv)).
		Methods("GET").
		Name("eolymp.acl.AclService.ListPermissions")
	router.Handle(prefix+"/whoami", _AclService_IntrospectPermission_Rule0(srv)).
		Methods("GET").
		Name("eolymp.acl.AclService.IntrospectPermission")
}

func _AclService_GrantPermission_Rule0(srv AclServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &GrantPermissionInput{}

		if err := _AclService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_AclService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.UserId = vars["user_id"]

		out, err := srv.GrantPermission(r.Context(), in)
		if err != nil {
			_AclService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AclService_HTTPWriteResponse(w, out)
	})
}

func _AclService_RevokePermission_Rule0(srv AclServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RevokePermissionInput{}

		if err := _AclService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_AclService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.UserId = vars["user_id"]

		out, err := srv.RevokePermission(r.Context(), in)
		if err != nil {
			_AclService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AclService_HTTPWriteResponse(w, out)
	})
}

func _AclService_DescribePermission_Rule0(srv AclServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePermissionInput{}

		if err := _AclService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_AclService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.UserId = vars["user_id"]

		out, err := srv.DescribePermission(r.Context(), in)
		if err != nil {
			_AclService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AclService_HTTPWriteResponse(w, out)
	})
}

func _AclService_ListPermissions_Rule0(srv AclServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPermissionsInput{}

		if err := _AclService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_AclService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListPermissions(r.Context(), in)
		if err != nil {
			_AclService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AclService_HTTPWriteResponse(w, out)
	})
}

func _AclService_IntrospectPermission_Rule0(srv AclServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &IntrospectPermissionInput{}

		if err := _AclService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_AclService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.IntrospectPermission(r.Context(), in)
		if err != nil {
			_AclService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AclService_HTTPWriteResponse(w, out)
	})
}

type _AclServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _AclServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _AclServiceHandler) (out proto.Message, err error)
type AclServiceInterceptor struct {
	middleware []_AclServiceMiddleware
	server     AclServiceServer
}

// NewAclServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewAclServiceInterceptor(srv AclServiceServer, middleware ..._AclServiceMiddleware) *AclServiceInterceptor {
	return &AclServiceInterceptor{server: srv, middleware: middleware}
}

func (i *AclServiceInterceptor) GrantPermission(ctx context.Context, in *GrantPermissionInput) (*GrantPermissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*GrantPermissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *GrantPermissionInput, got %T", in))
		}

		return i.server.GrantPermission(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.acl.AclService.GrantPermission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*GrantPermissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *GrantPermissionOutput, got %T", out))
	}

	return message, err
}

func (i *AclServiceInterceptor) RevokePermission(ctx context.Context, in *RevokePermissionInput) (*RevokePermissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RevokePermissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RevokePermissionInput, got %T", in))
		}

		return i.server.RevokePermission(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.acl.AclService.RevokePermission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RevokePermissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RevokePermissionOutput, got %T", out))
	}

	return message, err
}

func (i *AclServiceInterceptor) DescribePermission(ctx context.Context, in *DescribePermissionInput) (*DescribePermissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribePermissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribePermissionInput, got %T", in))
		}

		return i.server.DescribePermission(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.acl.AclService.DescribePermission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribePermissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribePermissionOutput, got %T", out))
	}

	return message, err
}

func (i *AclServiceInterceptor) ListPermissions(ctx context.Context, in *ListPermissionsInput) (*ListPermissionsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListPermissionsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListPermissionsInput, got %T", in))
		}

		return i.server.ListPermissions(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.acl.AclService.ListPermissions", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListPermissionsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListPermissionsOutput, got %T", out))
	}

	return message, err
}

func (i *AclServiceInterceptor) IntrospectPermission(ctx context.Context, in *IntrospectPermissionInput) (*IntrospectPermissionOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*IntrospectPermissionInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *IntrospectPermissionInput, got %T", in))
		}

		return i.server.IntrospectPermission(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.acl.AclService.IntrospectPermission", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*IntrospectPermissionOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *IntrospectPermissionOutput, got %T", out))
	}

	return message, err
}