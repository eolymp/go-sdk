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

// _CodeTemplateService_HTTPReadQueryString parses body into proto.Message
func _CodeTemplateService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _CodeTemplateService_HTTPReadRequestBody parses body into proto.Message
func _CodeTemplateService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _CodeTemplateService_HTTPWriteResponse writes proto.Message to HTTP response
func _CodeTemplateService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_CodeTemplateService_HTTPWriteErrorResponse(w, err)
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

// _CodeTemplateService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _CodeTemplateService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _CodeTemplateService_WebsocketErrorResponse writes error to websocket connection
func _CodeTemplateService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _CodeTemplateService_WebsocketCodec implements protobuf codec for websockets package
var _CodeTemplateService_WebsocketCodec = websocket.Codec{
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

// RegisterCodeTemplateServiceHttpHandlers adds handlers for for CodeTemplateServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterCodeTemplateServiceHttpHandlers(router *mux.Router, prefix string, cli CodeTemplateServiceClient) {
	router.Handle(prefix+"/templates", _CodeTemplateService_CreateCodeTemplate_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.CodeTemplateService.CreateCodeTemplate")
	router.Handle(prefix+"/templates/{template_id}", _CodeTemplateService_UpdateCodeTemplate_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.CodeTemplateService.UpdateCodeTemplate")
	router.Handle(prefix+"/templates/{template_id}", _CodeTemplateService_DeleteCodeTemplate_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.atlas.CodeTemplateService.DeleteCodeTemplate")
	router.Handle(prefix+"/templates", _CodeTemplateService_ListCodeTemplates_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.CodeTemplateService.ListCodeTemplates")
	router.Handle(prefix+"/templates/{template_id}", _CodeTemplateService_DescribeCodeTemplate_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.CodeTemplateService.DescribeCodeTemplate")
	router.Handle(prefix+"/template", _CodeTemplateService_LookupCodeTemplate_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.CodeTemplateService.LookupCodeTemplate")
}

func _CodeTemplateService_CreateCodeTemplate_Rule0(cli CodeTemplateServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateCodeTemplateInput{}

		if err := _CodeTemplateService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateCodeTemplate(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CodeTemplateService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _CodeTemplateService_UpdateCodeTemplate_Rule0(cli CodeTemplateServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateCodeTemplateInput{}

		if err := _CodeTemplateService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TemplateId = vars["template_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateCodeTemplate(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CodeTemplateService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _CodeTemplateService_DeleteCodeTemplate_Rule0(cli CodeTemplateServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteCodeTemplateInput{}

		if err := _CodeTemplateService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TemplateId = vars["template_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteCodeTemplate(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CodeTemplateService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _CodeTemplateService_ListCodeTemplates_Rule0(cli CodeTemplateServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListCodeTemplatesInput{}

		if err := _CodeTemplateService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListCodeTemplates(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CodeTemplateService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _CodeTemplateService_DescribeCodeTemplate_Rule0(cli CodeTemplateServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCodeTemplateInput{}

		if err := _CodeTemplateService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TemplateId = vars["template_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeCodeTemplate(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CodeTemplateService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _CodeTemplateService_LookupCodeTemplate_Rule0(cli CodeTemplateServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LookupCodeTemplateInput{}

		if err := _CodeTemplateService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.LookupCodeTemplate(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CodeTemplateService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _CodeTemplateServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _CodeTemplateServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _CodeTemplateServiceHandler) (out proto.Message, err error)
type CodeTemplateServiceInterceptor struct {
	middleware []_CodeTemplateServiceMiddleware
	client     CodeTemplateServiceClient
}

// NewCodeTemplateServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewCodeTemplateServiceInterceptor(cli CodeTemplateServiceClient, middleware ..._CodeTemplateServiceMiddleware) *CodeTemplateServiceInterceptor {
	return &CodeTemplateServiceInterceptor{client: cli, middleware: middleware}
}

func (i *CodeTemplateServiceInterceptor) CreateCodeTemplate(ctx context.Context, in *CreateCodeTemplateInput, opts ...grpc.CallOption) (*CreateCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateCodeTemplateInput, got %T", in))
		}

		return i.client.CreateCodeTemplate(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.CodeTemplateService.CreateCodeTemplate", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateCodeTemplateOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateCodeTemplateOutput, got %T", out))
	}

	return message, err
}

func (i *CodeTemplateServiceInterceptor) UpdateCodeTemplate(ctx context.Context, in *UpdateCodeTemplateInput, opts ...grpc.CallOption) (*UpdateCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateCodeTemplateInput, got %T", in))
		}

		return i.client.UpdateCodeTemplate(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.CodeTemplateService.UpdateCodeTemplate", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateCodeTemplateOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateCodeTemplateOutput, got %T", out))
	}

	return message, err
}

func (i *CodeTemplateServiceInterceptor) DeleteCodeTemplate(ctx context.Context, in *DeleteCodeTemplateInput, opts ...grpc.CallOption) (*DeleteCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteCodeTemplateInput, got %T", in))
		}

		return i.client.DeleteCodeTemplate(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.CodeTemplateService.DeleteCodeTemplate", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteCodeTemplateOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteCodeTemplateOutput, got %T", out))
	}

	return message, err
}

func (i *CodeTemplateServiceInterceptor) ListCodeTemplates(ctx context.Context, in *ListCodeTemplatesInput, opts ...grpc.CallOption) (*ListCodeTemplatesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListCodeTemplatesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListCodeTemplatesInput, got %T", in))
		}

		return i.client.ListCodeTemplates(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.CodeTemplateService.ListCodeTemplates", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListCodeTemplatesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListCodeTemplatesOutput, got %T", out))
	}

	return message, err
}

func (i *CodeTemplateServiceInterceptor) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput, opts ...grpc.CallOption) (*DescribeCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeCodeTemplateInput, got %T", in))
		}

		return i.client.DescribeCodeTemplate(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.CodeTemplateService.DescribeCodeTemplate", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeCodeTemplateOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeCodeTemplateOutput, got %T", out))
	}

	return message, err
}

func (i *CodeTemplateServiceInterceptor) LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput, opts ...grpc.CallOption) (*LookupCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*LookupCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *LookupCodeTemplateInput, got %T", in))
		}

		return i.client.LookupCodeTemplate(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.atlas.CodeTemplateService.LookupCodeTemplate", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*LookupCodeTemplateOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *LookupCodeTemplateOutput, got %T", out))
	}

	return message, err
}
