// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package atlas

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	websocket "golang.org/x/net/websocket"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _EditorialService_HTTPReadQueryString parses body into proto.Message
func _EditorialService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _EditorialService_HTTPReadRequestBody parses body into proto.Message
func _EditorialService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _EditorialService_HTTPWriteResponse writes proto.Message to HTTP response
func _EditorialService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_EditorialService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _EditorialService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _EditorialService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _EditorialService_WebsocketErrorResponse writes error to websocket connection
func _EditorialService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _EditorialService_WebsocketCodec implements protobuf codec for websockets package
var _EditorialService_WebsocketCodec = websocket.Codec{
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

// RegisterEditorialServiceHttpHandlers adds handlers for for EditorialServiceServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterEditorialServiceHttpHandlers(router *mux.Router, prefix string, srv EditorialServiceServer) {
	router.Handle(prefix+"/editorials", _EditorialService_CreateEditorial_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.atlas.EditorialService.CreateEditorial")
	router.Handle(prefix+"/editorials/{editorial_id}", _EditorialService_UpdateEditorial_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.atlas.EditorialService.UpdateEditorial")
	router.Handle(prefix+"/editorials/{editorial_id}", _EditorialService_DeleteEditorial_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.atlas.EditorialService.DeleteEditorial")
	router.Handle(prefix+"/editorials/{editorial_id}", _EditorialService_DescribeEditorial_Rule0(srv)).
		Methods("GET").
		Name("eolymp.atlas.EditorialService.DescribeEditorial")
	router.Handle(prefix+"/editorial", _EditorialService_LookupEditorial_Rule0(srv)).
		Methods("GET").
		Name("eolymp.atlas.EditorialService.LookupEditorial")
	router.Handle(prefix+"/editorial/preview", _EditorialService_PreviewEditorial_Rule0(srv)).
		Methods("POST").
		Name("eolymp.atlas.EditorialService.PreviewEditorial")
	router.Handle(prefix+"/editorials", _EditorialService_ListEditorials_Rule0(srv)).
		Methods("GET").
		Name("eolymp.atlas.EditorialService.ListEditorials")
}

func _EditorialService_CreateEditorial_Rule0(srv EditorialServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateEditorialInput{}

		if err := _EditorialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateEditorial(r.Context(), in)
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out)
	})
}

func _EditorialService_UpdateEditorial_Rule0(srv EditorialServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateEditorialInput{}

		if err := _EditorialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EditorialId = vars["editorial_id"]

		out, err := srv.UpdateEditorial(r.Context(), in)
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out)
	})
}

func _EditorialService_DeleteEditorial_Rule0(srv EditorialServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteEditorialInput{}

		if err := _EditorialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EditorialId = vars["editorial_id"]

		out, err := srv.DeleteEditorial(r.Context(), in)
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out)
	})
}

func _EditorialService_DescribeEditorial_Rule0(srv EditorialServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeEditorialInput{}

		if err := _EditorialService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EditorialId = vars["editorial_id"]

		out, err := srv.DescribeEditorial(r.Context(), in)
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out)
	})
}

func _EditorialService_LookupEditorial_Rule0(srv EditorialServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LookupEditorialInput{}

		if err := _EditorialService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.LookupEditorial(r.Context(), in)
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out)
	})
}

func _EditorialService_PreviewEditorial_Rule0(srv EditorialServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &PreviewEditorialInput{}

		if err := _EditorialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.PreviewEditorial(r.Context(), in)
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out)
	})
}

func _EditorialService_ListEditorials_Rule0(srv EditorialServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListEditorialsInput{}

		if err := _EditorialService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListEditorials(r.Context(), in)
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out)
	})
}

type _EditorialServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _EditorialServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _EditorialServiceHandler) (out proto.Message, err error)
type EditorialServiceInterceptor struct {
	middleware []_EditorialServiceMiddleware
	server     EditorialServiceServer
}

// NewEditorialServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewEditorialServiceInterceptor(srv EditorialServiceServer, middleware ..._EditorialServiceMiddleware) *EditorialServiceInterceptor {
	return &EditorialServiceInterceptor{server: srv, middleware: middleware}
}

func (i *EditorialServiceInterceptor) CreateEditorial(ctx context.Context, in *CreateEditorialInput) (*CreateEditorialOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateEditorialInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateEditorialInput, got %T", in))
		}

		return i.server.CreateEditorial(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.EditorialService.CreateEditorial", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateEditorialOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateEditorialOutput, got %T", out))
	}

	return message, err
}

func (i *EditorialServiceInterceptor) UpdateEditorial(ctx context.Context, in *UpdateEditorialInput) (*UpdateEditorialOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateEditorialInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateEditorialInput, got %T", in))
		}

		return i.server.UpdateEditorial(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.EditorialService.UpdateEditorial", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateEditorialOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateEditorialOutput, got %T", out))
	}

	return message, err
}

func (i *EditorialServiceInterceptor) DeleteEditorial(ctx context.Context, in *DeleteEditorialInput) (*DeleteEditorialOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteEditorialInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteEditorialInput, got %T", in))
		}

		return i.server.DeleteEditorial(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.EditorialService.DeleteEditorial", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteEditorialOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteEditorialOutput, got %T", out))
	}

	return message, err
}

func (i *EditorialServiceInterceptor) DescribeEditorial(ctx context.Context, in *DescribeEditorialInput) (*DescribeEditorialOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeEditorialInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeEditorialInput, got %T", in))
		}

		return i.server.DescribeEditorial(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.EditorialService.DescribeEditorial", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeEditorialOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeEditorialOutput, got %T", out))
	}

	return message, err
}

func (i *EditorialServiceInterceptor) LookupEditorial(ctx context.Context, in *LookupEditorialInput) (*LookupEditorialOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*LookupEditorialInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *LookupEditorialInput, got %T", in))
		}

		return i.server.LookupEditorial(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.EditorialService.LookupEditorial", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*LookupEditorialOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *LookupEditorialOutput, got %T", out))
	}

	return message, err
}

func (i *EditorialServiceInterceptor) PreviewEditorial(ctx context.Context, in *PreviewEditorialInput) (*PreviewEditorialOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*PreviewEditorialInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *PreviewEditorialInput, got %T", in))
		}

		return i.server.PreviewEditorial(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.EditorialService.PreviewEditorial", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*PreviewEditorialOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *PreviewEditorialOutput, got %T", out))
	}

	return message, err
}

func (i *EditorialServiceInterceptor) ListEditorials(ctx context.Context, in *ListEditorialsInput) (*ListEditorialsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListEditorialsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListEditorialsInput, got %T", in))
		}

		return i.server.ListEditorials(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.EditorialService.ListEditorials", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListEditorialsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListEditorialsOutput, got %T", out))
	}

	return message, err
}
