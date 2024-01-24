// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package helpdesk

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

// _Helpdesk_HTTPReadQueryString parses body into proto.Message
func _Helpdesk_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Helpdesk_HTTPReadRequestBody parses body into proto.Message
func _Helpdesk_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Helpdesk_HTTPWriteResponse writes proto.Message to HTTP response
func _Helpdesk_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Helpdesk_HTTPWriteErrorResponse(w, err)
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

// _Helpdesk_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Helpdesk_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _Helpdesk_WebsocketErrorResponse writes error to websocket connection
func _Helpdesk_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _Helpdesk_WebsocketCodec implements protobuf codec for websockets package
var _Helpdesk_WebsocketCodec = websocket.Codec{
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

// RegisterHelpdeskHttpHandlers adds handlers for for HelpdeskClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterHelpdeskHttpHandlers(router *mux.Router, prefix string, cli HelpdeskClient) {
	router.Handle(prefix+"/helpdesk/documents/{document_id}", _Helpdesk_DescribeDocument_Rule0(cli)).
		Methods("GET").
		Name("eolymp.helpdesk.Helpdesk.DescribeDocument")
	router.Handle(prefix+"/helpdesk/documents", _Helpdesk_ListDocuments_Rule0(cli)).
		Methods("GET").
		Name("eolymp.helpdesk.Helpdesk.ListDocuments")
	router.Handle(prefix+"/helpdesk/documents", _Helpdesk_CreateDocument_Rule0(cli)).
		Methods("POST").
		Name("eolymp.helpdesk.Helpdesk.CreateDocument")
	router.Handle(prefix+"/helpdesk/documents/{document_id}", _Helpdesk_UpdateDocument_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.helpdesk.Helpdesk.UpdateDocument")
	router.Handle(prefix+"/helpdesk/documents/{document_id}", _Helpdesk_DeleteDocument_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.helpdesk.Helpdesk.DeleteDocument")
	router.Handle(prefix+"/helpdesk/lookup/path", _Helpdesk_DescribePath_Rule0(cli)).
		Methods("GET").
		Name("eolymp.helpdesk.Helpdesk.DescribePath")
	router.Handle(prefix+"/helpdesk/lookup/paths", _Helpdesk_ListPaths_Rule0(cli)).
		Methods("GET").
		Name("eolymp.helpdesk.Helpdesk.ListPaths")
	router.Handle(prefix+"/helpdesk/lookup/parents", _Helpdesk_ListParents_Rule0(cli)).
		Methods("GET").
		Name("eolymp.helpdesk.Helpdesk.ListParents")
}

func _Helpdesk_DescribeDocument_Rule0(cli HelpdeskClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeDocumentInput{}

		if err := _Helpdesk_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.DocumentId = vars["document_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeDocument(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Helpdesk_ListDocuments_Rule0(cli HelpdeskClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListDocumentsInput{}

		if err := _Helpdesk_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListDocuments(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Helpdesk_CreateDocument_Rule0(cli HelpdeskClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateDocumentInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateDocument(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Helpdesk_UpdateDocument_Rule0(cli HelpdeskClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateDocumentInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.DocumentId = vars["document_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateDocument(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Helpdesk_DeleteDocument_Rule0(cli HelpdeskClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteDocumentInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.DocumentId = vars["document_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteDocument(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Helpdesk_DescribePath_Rule0(cli HelpdeskClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePathInput{}

		if err := _Helpdesk_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribePath(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Helpdesk_ListPaths_Rule0(cli HelpdeskClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPathsInput{}

		if err := _Helpdesk_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListPaths(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Helpdesk_ListParents_Rule0(cli HelpdeskClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListParentsInput{}

		if err := _Helpdesk_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListParents(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _HelpdeskHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _HelpdeskMiddleware = func(ctx context.Context, method string, in proto.Message, handler _HelpdeskHandler) (out proto.Message, err error)
type HelpdeskInterceptor struct {
	middleware []_HelpdeskMiddleware
	client     HelpdeskClient
}

// NewHelpdeskInterceptor constructs additional middleware for a server based on annotations in proto files
func NewHelpdeskInterceptor(cli HelpdeskClient, middleware ..._HelpdeskMiddleware) *HelpdeskInterceptor {
	return &HelpdeskInterceptor{client: cli, middleware: middleware}
}

func (i *HelpdeskInterceptor) DescribeDocument(ctx context.Context, in *DescribeDocumentInput, opts ...grpc.CallOption) (*DescribeDocumentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeDocumentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeDocumentInput, got %T", in))
		}

		return i.client.DescribeDocument(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Helpdesk.DescribeDocument", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeDocumentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeDocumentOutput, got %T", out))
	}

	return message, err
}

func (i *HelpdeskInterceptor) ListDocuments(ctx context.Context, in *ListDocumentsInput, opts ...grpc.CallOption) (*ListDocumentsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListDocumentsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListDocumentsInput, got %T", in))
		}

		return i.client.ListDocuments(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Helpdesk.ListDocuments", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListDocumentsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListDocumentsOutput, got %T", out))
	}

	return message, err
}

func (i *HelpdeskInterceptor) CreateDocument(ctx context.Context, in *CreateDocumentInput, opts ...grpc.CallOption) (*CreateDocumentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateDocumentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateDocumentInput, got %T", in))
		}

		return i.client.CreateDocument(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Helpdesk.CreateDocument", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateDocumentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateDocumentOutput, got %T", out))
	}

	return message, err
}

func (i *HelpdeskInterceptor) UpdateDocument(ctx context.Context, in *UpdateDocumentInput, opts ...grpc.CallOption) (*UpdateDocumentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateDocumentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateDocumentInput, got %T", in))
		}

		return i.client.UpdateDocument(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Helpdesk.UpdateDocument", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateDocumentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateDocumentOutput, got %T", out))
	}

	return message, err
}

func (i *HelpdeskInterceptor) DeleteDocument(ctx context.Context, in *DeleteDocumentInput, opts ...grpc.CallOption) (*DeleteDocumentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteDocumentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteDocumentInput, got %T", in))
		}

		return i.client.DeleteDocument(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Helpdesk.DeleteDocument", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteDocumentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteDocumentOutput, got %T", out))
	}

	return message, err
}

func (i *HelpdeskInterceptor) DescribePath(ctx context.Context, in *DescribePathInput, opts ...grpc.CallOption) (*DescribePathOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribePathInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribePathInput, got %T", in))
		}

		return i.client.DescribePath(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Helpdesk.DescribePath", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribePathOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribePathOutput, got %T", out))
	}

	return message, err
}

func (i *HelpdeskInterceptor) ListPaths(ctx context.Context, in *ListPathsInput, opts ...grpc.CallOption) (*ListPathsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListPathsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListPathsInput, got %T", in))
		}

		return i.client.ListPaths(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Helpdesk.ListPaths", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListPathsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListPathsOutput, got %T", out))
	}

	return message, err
}

func (i *HelpdeskInterceptor) ListParents(ctx context.Context, in *ListParentsInput, opts ...grpc.CallOption) (*ListParentsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListParentsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListParentsInput, got %T", in))
		}

		return i.client.ListParents(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Helpdesk.ListParents", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListParentsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListParentsOutput, got %T", out))
	}

	return message, err
}
