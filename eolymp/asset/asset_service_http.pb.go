// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package asset

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

// _AssetService_HTTPReadQueryString parses body into proto.Message
func _AssetService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _AssetService_HTTPReadRequestBody parses body into proto.Message
func _AssetService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _AssetService_HTTPWriteResponse writes proto.Message to HTTP response
func _AssetService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_AssetService_HTTPWriteErrorResponse(w, err)
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

// _AssetService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _AssetService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _AssetService_WebsocketErrorResponse writes error to websocket connection
func _AssetService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _AssetService_WebsocketCodec implements protobuf codec for websockets package
var _AssetService_WebsocketCodec = websocket.Codec{
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

// RegisterAssetServiceHttpHandlers adds handlers for for AssetServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterAssetServiceHttpHandlers(router *mux.Router, prefix string, cli AssetServiceClient) {
	router.Handle(prefix+"/assets/images", _AssetService_UploadImage_Rule0(cli)).
		Methods("POST").
		Name("eolymp.asset.AssetService.UploadImage")
	router.Handle(prefix+"/assets/files", _AssetService_UploadFile_Rule0(cli)).
		Methods("POST").
		Name("eolymp.asset.AssetService.UploadFile")
	router.Handle(prefix+"/assets", _AssetService_UploadAsset_Rule0(cli)).
		Methods("POST").
		Name("eolymp.asset.AssetService.UploadAsset")
	router.Handle(prefix+"/assets:lookup", _AssetService_LookupAsset_Rule0(cli)).
		Methods("POST").
		Name("eolymp.asset.AssetService.LookupAsset")
	router.Handle(prefix+"/uploads", _AssetService_StartMultipartUpload_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.asset.AssetService.StartMultipartUpload")
	router.Handle(prefix+"/uploads/{upload_id}", _AssetService_UploadPart_Rule0(cli)).
		Methods("POST").
		Name("eolymp.asset.AssetService.UploadPart")
	router.Handle(prefix+"/uploads/{upload_id}", _AssetService_CompleteMultipartUpload_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.asset.AssetService.CompleteMultipartUpload")
	router.Handle(prefix+"/streams", _AssetService_StartStream_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.asset.AssetService.StartStream")
	router.Handle(prefix+"/streams/{stream_id}", _AssetService_AppendStream_Rule0(cli)).
		Methods("POST").
		Name("eolymp.asset.AssetService.AppendStream")
	router.Handle(prefix+"/streams/{stream_id}", _AssetService_CloseStream_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.asset.AssetService.CloseStream")
}

func _AssetService_UploadImage_Rule0(cli AssetServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UploadImageInput{}

		if err := _AssetService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.UploadImage(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssetService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssetService_UploadFile_Rule0(cli AssetServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UploadFileInput{}

		if err := _AssetService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.UploadFile(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssetService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssetService_UploadAsset_Rule0(cli AssetServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UploadAssetInput{}

		if err := _AssetService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.UploadAsset(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssetService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssetService_LookupAsset_Rule0(cli AssetServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LookupAssetInput{}

		if err := _AssetService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.LookupAsset(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssetService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssetService_StartMultipartUpload_Rule0(cli AssetServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &StartMultipartUploadInput{}

		if err := _AssetService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.StartMultipartUpload(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssetService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssetService_UploadPart_Rule0(cli AssetServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UploadPartInput{}

		if err := _AssetService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.UploadId = vars["upload_id"]

		var header, trailer metadata.MD

		out, err := cli.UploadPart(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssetService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssetService_CompleteMultipartUpload_Rule0(cli AssetServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CompleteMultipartUploadInput{}

		if err := _AssetService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.UploadId = vars["upload_id"]

		var header, trailer metadata.MD

		out, err := cli.CompleteMultipartUpload(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssetService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssetService_StartStream_Rule0(cli AssetServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &StartStreamInput{}

		if err := _AssetService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.StartStream(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssetService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssetService_AppendStream_Rule0(cli AssetServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &AppendStreamInput{}

		if err := _AssetService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.StreamId = vars["stream_id"]

		var header, trailer metadata.MD

		out, err := cli.AppendStream(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssetService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AssetService_CloseStream_Rule0(cli AssetServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CloseStreamInput{}

		if err := _AssetService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.StreamId = vars["stream_id"]

		var header, trailer metadata.MD

		out, err := cli.CloseStream(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AssetService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AssetService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _AssetServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _AssetServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _AssetServiceHandler) (out proto.Message, err error)
type AssetServiceInterceptor struct {
	middleware []_AssetServiceMiddleware
	client     AssetServiceClient
}

// NewAssetServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewAssetServiceInterceptor(cli AssetServiceClient, middleware ..._AssetServiceMiddleware) *AssetServiceInterceptor {
	return &AssetServiceInterceptor{client: cli, middleware: middleware}
}

func (i *AssetServiceInterceptor) UploadImage(ctx context.Context, in *UploadImageInput, opts ...grpc.CallOption) (*UploadImageOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UploadImageInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UploadImageInput, got %T", in))
		}

		return i.client.UploadImage(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.asset.AssetService.UploadImage", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UploadImageOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UploadImageOutput, got %T", out))
	}

	return message, err
}

func (i *AssetServiceInterceptor) UploadFile(ctx context.Context, in *UploadFileInput, opts ...grpc.CallOption) (*UploadFileOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UploadFileInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UploadFileInput, got %T", in))
		}

		return i.client.UploadFile(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.asset.AssetService.UploadFile", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UploadFileOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UploadFileOutput, got %T", out))
	}

	return message, err
}

func (i *AssetServiceInterceptor) UploadAsset(ctx context.Context, in *UploadAssetInput, opts ...grpc.CallOption) (*UploadAssetOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UploadAssetInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UploadAssetInput, got %T", in))
		}

		return i.client.UploadAsset(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.asset.AssetService.UploadAsset", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UploadAssetOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UploadAssetOutput, got %T", out))
	}

	return message, err
}

func (i *AssetServiceInterceptor) LookupAsset(ctx context.Context, in *LookupAssetInput, opts ...grpc.CallOption) (*LookupAssetOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*LookupAssetInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *LookupAssetInput, got %T", in))
		}

		return i.client.LookupAsset(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.asset.AssetService.LookupAsset", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*LookupAssetOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *LookupAssetOutput, got %T", out))
	}

	return message, err
}

func (i *AssetServiceInterceptor) StartMultipartUpload(ctx context.Context, in *StartMultipartUploadInput, opts ...grpc.CallOption) (*StartMultipartUploadOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*StartMultipartUploadInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *StartMultipartUploadInput, got %T", in))
		}

		return i.client.StartMultipartUpload(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.asset.AssetService.StartMultipartUpload", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*StartMultipartUploadOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *StartMultipartUploadOutput, got %T", out))
	}

	return message, err
}

func (i *AssetServiceInterceptor) UploadPart(ctx context.Context, in *UploadPartInput, opts ...grpc.CallOption) (*UploadPartOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UploadPartInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UploadPartInput, got %T", in))
		}

		return i.client.UploadPart(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.asset.AssetService.UploadPart", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UploadPartOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UploadPartOutput, got %T", out))
	}

	return message, err
}

func (i *AssetServiceInterceptor) CompleteMultipartUpload(ctx context.Context, in *CompleteMultipartUploadInput, opts ...grpc.CallOption) (*CompleteMultipartUploadOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CompleteMultipartUploadInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CompleteMultipartUploadInput, got %T", in))
		}

		return i.client.CompleteMultipartUpload(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.asset.AssetService.CompleteMultipartUpload", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CompleteMultipartUploadOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CompleteMultipartUploadOutput, got %T", out))
	}

	return message, err
}

func (i *AssetServiceInterceptor) StartStream(ctx context.Context, in *StartStreamInput, opts ...grpc.CallOption) (*StartStreamOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*StartStreamInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *StartStreamInput, got %T", in))
		}

		return i.client.StartStream(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.asset.AssetService.StartStream", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*StartStreamOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *StartStreamOutput, got %T", out))
	}

	return message, err
}

func (i *AssetServiceInterceptor) AppendStream(ctx context.Context, in *AppendStreamInput, opts ...grpc.CallOption) (*AppendStreamOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*AppendStreamInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *AppendStreamInput, got %T", in))
		}

		return i.client.AppendStream(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.asset.AssetService.AppendStream", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*AppendStreamOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *AppendStreamOutput, got %T", out))
	}

	return message, err
}

func (i *AssetServiceInterceptor) CloseStream(ctx context.Context, in *CloseStreamInput, opts ...grpc.CallOption) (*CloseStreamOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CloseStreamInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CloseStreamInput, got %T", in))
		}

		return i.client.CloseStream(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.asset.AssetService.CloseStream", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CloseStreamOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CloseStreamOutput, got %T", out))
	}

	return message, err
}
