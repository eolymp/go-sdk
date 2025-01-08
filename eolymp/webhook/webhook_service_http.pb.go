// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package webhook

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _WebhookService_HTTPReadQueryString parses body into proto.Message
func _WebhookService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _WebhookService_HTTPReadRequestBody parses body into proto.Message
func _WebhookService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _WebhookService_HTTPWriteResponse writes proto.Message to HTTP response
func _WebhookService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_WebhookService_HTTPWriteErrorResponse(w, err)
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

// _WebhookService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _WebhookService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterWebhookServiceHttpHandlers adds handlers for for WebhookServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterWebhookServiceHttpHandlers(router *mux.Router, prefix string, cli WebhookServiceClient) {
	router.Handle(prefix+"/webhooks", _WebhookService_CreateWebhook_Rule0(cli)).
		Methods("POST").
		Name("eolymp.webhook.WebhookService.CreateWebhook")
	router.Handle(prefix+"/webhooks/{webhook_id}", _WebhookService_UpdateWebhook_Rule0(cli)).
		Methods("POST").
		Name("eolymp.webhook.WebhookService.UpdateWebhook")
	router.Handle(prefix+"/webhooks/{webhook_id}", _WebhookService_DeleteWebhook_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.webhook.WebhookService.DeleteWebhook")
	router.Handle(prefix+"/webhooks/{webhook_id}", _WebhookService_DescribeWebhook_Rule0(cli)).
		Methods("GET").
		Name("eolymp.webhook.WebhookService.DescribeWebhook")
	router.Handle(prefix+"/webhooks", _WebhookService_ListWebhooks_Rule0(cli)).
		Methods("GET").
		Name("eolymp.webhook.WebhookService.ListWebhooks")
	router.Handle(prefix+"/webhooks/{webhook_id}/test", _WebhookService_TestWebhook_Rule0(cli)).
		Methods("POST").
		Name("eolymp.webhook.WebhookService.TestWebhook")
}

func _WebhookService_CreateWebhook_Rule0(cli WebhookServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateWebhookInput{}

		if err := _WebhookService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateWebhook(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		_WebhookService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _WebhookService_UpdateWebhook_Rule0(cli WebhookServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateWebhookInput{}

		if err := _WebhookService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.WebhookId = vars["webhook_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateWebhook(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		_WebhookService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _WebhookService_DeleteWebhook_Rule0(cli WebhookServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteWebhookInput{}

		if err := _WebhookService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.WebhookId = vars["webhook_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteWebhook(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		_WebhookService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _WebhookService_DescribeWebhook_Rule0(cli WebhookServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeWebhookInput{}

		if err := _WebhookService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.WebhookId = vars["webhook_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeWebhook(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		_WebhookService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _WebhookService_ListWebhooks_Rule0(cli WebhookServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListWebhooksInput{}

		if err := _WebhookService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListWebhooks(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		_WebhookService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _WebhookService_TestWebhook_Rule0(cli WebhookServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &TestWebhookInput{}

		if err := _WebhookService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.WebhookId = vars["webhook_id"]

		var header, trailer metadata.MD

		out, err := cli.TestWebhook(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_WebhookService_HTTPWriteErrorResponse(w, err)
			return
		}

		_WebhookService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _WebhookServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _WebhookServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _WebhookServiceHandler) (out proto.Message, err error)
type WebhookServiceInterceptor struct {
	middleware []_WebhookServiceMiddleware
	client     WebhookServiceClient
}

// NewWebhookServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewWebhookServiceInterceptor(cli WebhookServiceClient, middleware ..._WebhookServiceMiddleware) *WebhookServiceInterceptor {
	return &WebhookServiceInterceptor{client: cli, middleware: middleware}
}

func (i *WebhookServiceInterceptor) CreateWebhook(ctx context.Context, in *CreateWebhookInput, opts ...grpc.CallOption) (*CreateWebhookOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateWebhookInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateWebhookInput, got %T", in))
		}

		return i.client.CreateWebhook(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.webhook.WebhookService.CreateWebhook", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateWebhookOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateWebhookOutput, got %T", out))
	}

	return message, err
}

func (i *WebhookServiceInterceptor) UpdateWebhook(ctx context.Context, in *UpdateWebhookInput, opts ...grpc.CallOption) (*UpdateWebhookOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateWebhookInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateWebhookInput, got %T", in))
		}

		return i.client.UpdateWebhook(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.webhook.WebhookService.UpdateWebhook", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateWebhookOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateWebhookOutput, got %T", out))
	}

	return message, err
}

func (i *WebhookServiceInterceptor) DeleteWebhook(ctx context.Context, in *DeleteWebhookInput, opts ...grpc.CallOption) (*DeleteWebhookOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteWebhookInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteWebhookInput, got %T", in))
		}

		return i.client.DeleteWebhook(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.webhook.WebhookService.DeleteWebhook", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteWebhookOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteWebhookOutput, got %T", out))
	}

	return message, err
}

func (i *WebhookServiceInterceptor) DescribeWebhook(ctx context.Context, in *DescribeWebhookInput, opts ...grpc.CallOption) (*DescribeWebhookOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeWebhookInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeWebhookInput, got %T", in))
		}

		return i.client.DescribeWebhook(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.webhook.WebhookService.DescribeWebhook", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeWebhookOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeWebhookOutput, got %T", out))
	}

	return message, err
}

func (i *WebhookServiceInterceptor) ListWebhooks(ctx context.Context, in *ListWebhooksInput, opts ...grpc.CallOption) (*ListWebhooksOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListWebhooksInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListWebhooksInput, got %T", in))
		}

		return i.client.ListWebhooks(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.webhook.WebhookService.ListWebhooks", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListWebhooksOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListWebhooksOutput, got %T", out))
	}

	return message, err
}

func (i *WebhookServiceInterceptor) TestWebhook(ctx context.Context, in *TestWebhookInput, opts ...grpc.CallOption) (*TestWebhookOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*TestWebhookInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *TestWebhookInput, got %T", in))
		}

		return i.client.TestWebhook(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.webhook.WebhookService.TestWebhook", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*TestWebhookOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *TestWebhookOutput, got %T", out))
	}

	return message, err
}
