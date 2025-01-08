// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package judge

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

// _AnnouncementService_HTTPReadQueryString parses body into proto.Message
func _AnnouncementService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _AnnouncementService_HTTPReadRequestBody parses body into proto.Message
func _AnnouncementService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _AnnouncementService_HTTPWriteResponse writes proto.Message to HTTP response
func _AnnouncementService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_AnnouncementService_HTTPWriteErrorResponse(w, err)
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

// _AnnouncementService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _AnnouncementService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterAnnouncementServiceHttpHandlers adds handlers for for AnnouncementServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterAnnouncementServiceHttpHandlers(router *mux.Router, prefix string, cli AnnouncementServiceClient) {
	router.Handle(prefix+"/announcements", _AnnouncementService_CreateAnnouncement_Rule0(cli)).
		Methods("POST").
		Name("eolymp.judge.AnnouncementService.CreateAnnouncement")
	router.Handle(prefix+"/announcements/{announcement_id}", _AnnouncementService_UpdateAnnouncement_Rule0(cli)).
		Methods("POST").
		Name("eolymp.judge.AnnouncementService.UpdateAnnouncement")
	router.Handle(prefix+"/announcements/{announcement_id}", _AnnouncementService_DeleteAnnouncement_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.judge.AnnouncementService.DeleteAnnouncement")
	router.Handle(prefix+"/announcements/{announcement_id}/read", _AnnouncementService_ReadAnnouncement_Rule0(cli)).
		Methods("POST").
		Name("eolymp.judge.AnnouncementService.ReadAnnouncement")
	router.Handle(prefix+"/announcements/{announcement_id}", _AnnouncementService_DescribeAnnouncement_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.AnnouncementService.DescribeAnnouncement")
	router.Handle(prefix+"/announcements/{announcement_id}/status", _AnnouncementService_DescribeAnnouncementStatus_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.AnnouncementService.DescribeAnnouncementStatus")
	router.Handle(prefix+"/announcements", _AnnouncementService_ListAnnouncements_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.AnnouncementService.ListAnnouncements")
}

func _AnnouncementService_CreateAnnouncement_Rule0(cli AnnouncementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateAnnouncementInput{}

		if err := _AnnouncementService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateAnnouncement(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AnnouncementService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AnnouncementService_UpdateAnnouncement_Rule0(cli AnnouncementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateAnnouncementInput{}

		if err := _AnnouncementService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AnnouncementId = vars["announcement_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateAnnouncement(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AnnouncementService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AnnouncementService_DeleteAnnouncement_Rule0(cli AnnouncementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteAnnouncementInput{}

		if err := _AnnouncementService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AnnouncementId = vars["announcement_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteAnnouncement(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AnnouncementService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AnnouncementService_ReadAnnouncement_Rule0(cli AnnouncementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ReadAnnouncementInput{}

		if err := _AnnouncementService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AnnouncementId = vars["announcement_id"]

		var header, trailer metadata.MD

		out, err := cli.ReadAnnouncement(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AnnouncementService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AnnouncementService_DescribeAnnouncement_Rule0(cli AnnouncementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeAnnouncementInput{}

		if err := _AnnouncementService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AnnouncementId = vars["announcement_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeAnnouncement(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AnnouncementService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AnnouncementService_DescribeAnnouncementStatus_Rule0(cli AnnouncementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeAnnouncementStatusInput{}

		if err := _AnnouncementService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AnnouncementId = vars["announcement_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeAnnouncementStatus(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AnnouncementService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AnnouncementService_ListAnnouncements_Rule0(cli AnnouncementServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListAnnouncementsInput{}

		if err := _AnnouncementService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListAnnouncements(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AnnouncementService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AnnouncementService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _AnnouncementServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _AnnouncementServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _AnnouncementServiceHandler) (out proto.Message, err error)
type AnnouncementServiceInterceptor struct {
	middleware []_AnnouncementServiceMiddleware
	client     AnnouncementServiceClient
}

// NewAnnouncementServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewAnnouncementServiceInterceptor(cli AnnouncementServiceClient, middleware ..._AnnouncementServiceMiddleware) *AnnouncementServiceInterceptor {
	return &AnnouncementServiceInterceptor{client: cli, middleware: middleware}
}

func (i *AnnouncementServiceInterceptor) CreateAnnouncement(ctx context.Context, in *CreateAnnouncementInput, opts ...grpc.CallOption) (*CreateAnnouncementOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateAnnouncementInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateAnnouncementInput, got %T", in))
		}

		return i.client.CreateAnnouncement(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.AnnouncementService.CreateAnnouncement", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateAnnouncementOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateAnnouncementOutput, got %T", out))
	}

	return message, err
}

func (i *AnnouncementServiceInterceptor) UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementInput, opts ...grpc.CallOption) (*UpdateAnnouncementOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateAnnouncementInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateAnnouncementInput, got %T", in))
		}

		return i.client.UpdateAnnouncement(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.AnnouncementService.UpdateAnnouncement", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateAnnouncementOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateAnnouncementOutput, got %T", out))
	}

	return message, err
}

func (i *AnnouncementServiceInterceptor) DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementInput, opts ...grpc.CallOption) (*DeleteAnnouncementOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteAnnouncementInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteAnnouncementInput, got %T", in))
		}

		return i.client.DeleteAnnouncement(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.AnnouncementService.DeleteAnnouncement", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteAnnouncementOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteAnnouncementOutput, got %T", out))
	}

	return message, err
}

func (i *AnnouncementServiceInterceptor) ReadAnnouncement(ctx context.Context, in *ReadAnnouncementInput, opts ...grpc.CallOption) (*ReadAnnouncementOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ReadAnnouncementInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ReadAnnouncementInput, got %T", in))
		}

		return i.client.ReadAnnouncement(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.AnnouncementService.ReadAnnouncement", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ReadAnnouncementOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ReadAnnouncementOutput, got %T", out))
	}

	return message, err
}

func (i *AnnouncementServiceInterceptor) DescribeAnnouncement(ctx context.Context, in *DescribeAnnouncementInput, opts ...grpc.CallOption) (*DescribeAnnouncementOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeAnnouncementInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeAnnouncementInput, got %T", in))
		}

		return i.client.DescribeAnnouncement(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.AnnouncementService.DescribeAnnouncement", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeAnnouncementOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeAnnouncementOutput, got %T", out))
	}

	return message, err
}

func (i *AnnouncementServiceInterceptor) DescribeAnnouncementStatus(ctx context.Context, in *DescribeAnnouncementStatusInput, opts ...grpc.CallOption) (*DescribeAnnouncementStatusOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeAnnouncementStatusInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeAnnouncementStatusInput, got %T", in))
		}

		return i.client.DescribeAnnouncementStatus(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.AnnouncementService.DescribeAnnouncementStatus", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeAnnouncementStatusOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeAnnouncementStatusOutput, got %T", out))
	}

	return message, err
}

func (i *AnnouncementServiceInterceptor) ListAnnouncements(ctx context.Context, in *ListAnnouncementsInput, opts ...grpc.CallOption) (*ListAnnouncementsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListAnnouncementsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListAnnouncementsInput, got %T", in))
		}

		return i.client.ListAnnouncements(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.AnnouncementService.ListAnnouncements", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListAnnouncementsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListAnnouncementsOutput, got %T", out))
	}

	return message, err
}
