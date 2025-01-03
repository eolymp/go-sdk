// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package judge

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

// _ProblemService_HTTPReadQueryString parses body into proto.Message
func _ProblemService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _ProblemService_HTTPReadRequestBody parses body into proto.Message
func _ProblemService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _ProblemService_HTTPWriteResponse writes proto.Message to HTTP response
func _ProblemService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_ProblemService_HTTPWriteErrorResponse(w, err)
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

// _ProblemService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _ProblemService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _ProblemService_WebsocketErrorResponse writes error to websocket connection
func _ProblemService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _ProblemService_WebsocketCodec implements protobuf codec for websockets package
var _ProblemService_WebsocketCodec = websocket.Codec{
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

// RegisterProblemServiceHttpHandlers adds handlers for for ProblemServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterProblemServiceHttpHandlers(router *mux.Router, prefix string, cli ProblemServiceClient) {
	router.Handle(prefix+"/problems", _ProblemService_ImportProblem_Rule0(cli)).
		Methods("POST").
		Name("eolymp.judge.ProblemService.ImportProblem")
	router.Handle(prefix+"/problems/{problem_id}/sync", _ProblemService_SyncProblem_Rule0(cli)).
		Methods("POST").
		Name("eolymp.judge.ProblemService.SyncProblem")
	router.Handle(prefix+"/problems/{problem_id}", _ProblemService_UpdateProblem_Rule0(cli)).
		Methods("POST").
		Name("eolymp.judge.ProblemService.UpdateProblem")
	router.Handle(prefix+"/problems", _ProblemService_ListProblems_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.ProblemService.ListProblems")
	router.Handle(prefix+"/problems/{problem_id}", _ProblemService_DescribeProblem_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.ProblemService.DescribeProblem")
	router.Handle(prefix+"/problems/{problem_id}", _ProblemService_DeleteProblem_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.judge.ProblemService.DeleteProblem")
	router.Handle(prefix+"/problems/{problem_id}/lookup-template", _ProblemService_LookupCodeTemplate_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.ProblemService.LookupCodeTemplate")
	router.Handle(prefix+"/problems/{problem_id}/templates/{template_id}", _ProblemService_DescribeCodeTemplate_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.ProblemService.DescribeCodeTemplate")
	router.Handle(prefix+"/problems/{problem_id}/statements", _ProblemService_ListStatements_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.ProblemService.ListStatements")
	router.Handle(prefix+"/problems/{problem_id}/attachments", _ProblemService_ListAttachments_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.ProblemService.ListAttachments")
	router.Handle(prefix+"/problems/{problem_id}/examples", _ProblemService_ListExamples_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.ProblemService.ListExamples")
	router.Handle(prefix+"/problems/{problem_id}/runtime", _ProblemService_ListRuntimes_Rule0(cli)).
		Methods("GET").
		Name("eolymp.judge.ProblemService.ListRuntimes")
}

func _ProblemService_ImportProblem_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ImportProblemInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ImportProblem(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ProblemService_SyncProblem_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &SyncProblemInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		var header, trailer metadata.MD

		out, err := cli.SyncProblem(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ProblemService_UpdateProblem_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateProblemInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateProblem(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ProblemService_ListProblems_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListProblemsInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListProblems(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ProblemService_DescribeProblem_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeProblemInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeProblem(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ProblemService_DeleteProblem_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteProblemInput{}

		if err := _ProblemService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteProblem(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ProblemService_LookupCodeTemplate_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LookupCodeTemplateInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		var header, trailer metadata.MD

		out, err := cli.LookupCodeTemplate(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ProblemService_DescribeCodeTemplate_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCodeTemplateInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]
		in.TemplateId = vars["template_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeCodeTemplate(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ProblemService_ListStatements_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListStatementsInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		var header, trailer metadata.MD

		out, err := cli.ListStatements(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ProblemService_ListAttachments_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListAttachmentsInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		var header, trailer metadata.MD

		out, err := cli.ListAttachments(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ProblemService_ListExamples_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListExamplesInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		var header, trailer metadata.MD

		out, err := cli.ListExamples(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ProblemService_ListRuntimes_Rule0(cli ProblemServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListRuntimesInput{}

		if err := _ProblemService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProblemId = vars["problem_id"]

		var header, trailer metadata.MD

		out, err := cli.ListRuntimes(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ProblemService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ProblemService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _ProblemServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _ProblemServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _ProblemServiceHandler) (out proto.Message, err error)
type ProblemServiceInterceptor struct {
	middleware []_ProblemServiceMiddleware
	client     ProblemServiceClient
}

// NewProblemServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewProblemServiceInterceptor(cli ProblemServiceClient, middleware ..._ProblemServiceMiddleware) *ProblemServiceInterceptor {
	return &ProblemServiceInterceptor{client: cli, middleware: middleware}
}

func (i *ProblemServiceInterceptor) ImportProblem(ctx context.Context, in *ImportProblemInput, opts ...grpc.CallOption) (*ImportProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ImportProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ImportProblemInput, got %T", in))
		}

		return i.client.ImportProblem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.ProblemService.ImportProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ImportProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ImportProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) SyncProblem(ctx context.Context, in *SyncProblemInput, opts ...grpc.CallOption) (*SyncProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*SyncProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *SyncProblemInput, got %T", in))
		}

		return i.client.SyncProblem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.ProblemService.SyncProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*SyncProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *SyncProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) UpdateProblem(ctx context.Context, in *UpdateProblemInput, opts ...grpc.CallOption) (*UpdateProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateProblemInput, got %T", in))
		}

		return i.client.UpdateProblem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.ProblemService.UpdateProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) ListProblems(ctx context.Context, in *ListProblemsInput, opts ...grpc.CallOption) (*ListProblemsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListProblemsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListProblemsInput, got %T", in))
		}

		return i.client.ListProblems(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.ProblemService.ListProblems", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListProblemsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListProblemsOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) DescribeProblem(ctx context.Context, in *DescribeProblemInput, opts ...grpc.CallOption) (*DescribeProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeProblemInput, got %T", in))
		}

		return i.client.DescribeProblem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.ProblemService.DescribeProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) DeleteProblem(ctx context.Context, in *DeleteProblemInput, opts ...grpc.CallOption) (*DeleteProblemOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteProblemInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteProblemInput, got %T", in))
		}

		return i.client.DeleteProblem(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.ProblemService.DeleteProblem", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteProblemOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteProblemOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) LookupCodeTemplate(ctx context.Context, in *LookupCodeTemplateInput, opts ...grpc.CallOption) (*LookupCodeTemplateOutput, error) {
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
			return mw(ctx, "eolymp.judge.ProblemService.LookupCodeTemplate", in, next)
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

func (i *ProblemServiceInterceptor) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput, opts ...grpc.CallOption) (*DescribeCodeTemplateOutput, error) {
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
			return mw(ctx, "eolymp.judge.ProblemService.DescribeCodeTemplate", in, next)
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

func (i *ProblemServiceInterceptor) ListStatements(ctx context.Context, in *ListStatementsInput, opts ...grpc.CallOption) (*ListStatementsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListStatementsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListStatementsInput, got %T", in))
		}

		return i.client.ListStatements(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.ProblemService.ListStatements", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListStatementsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListStatementsOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) ListAttachments(ctx context.Context, in *ListAttachmentsInput, opts ...grpc.CallOption) (*ListAttachmentsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListAttachmentsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListAttachmentsInput, got %T", in))
		}

		return i.client.ListAttachments(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.ProblemService.ListAttachments", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListAttachmentsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListAttachmentsOutput, got %T", out))
	}

	return message, err
}

func (i *ProblemServiceInterceptor) ListExamples(ctx context.Context, in *ListExamplesInput, opts ...grpc.CallOption) (*ListExamplesOutput, error) {
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
			return mw(ctx, "eolymp.judge.ProblemService.ListExamples", in, next)
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

func (i *ProblemServiceInterceptor) ListRuntimes(ctx context.Context, in *ListRuntimesInput, opts ...grpc.CallOption) (*ListRuntimesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListRuntimesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListRuntimesInput, got %T", in))
		}

		return i.client.ListRuntimes(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.judge.ProblemService.ListRuntimes", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListRuntimesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListRuntimesOutput, got %T", out))
	}

	return message, err
}
