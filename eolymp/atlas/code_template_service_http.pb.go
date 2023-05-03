// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package atlas

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
func _CodeTemplateService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_CodeTemplateService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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

// RegisterCodeTemplateServiceHttpHandlers adds handlers for for CodeTemplateServiceServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterCodeTemplateServiceHttpHandlers(router *mux.Router, prefix string, srv CodeTemplateServiceServer) {
	router.Handle(prefix+"/templates", _CodeTemplateService_CreateCodeTemplate_Rule0(srv)).
		Methods("POST").
		Name("eolymp.atlas.CodeTemplateService.CreateCodeTemplate")
	router.Handle(prefix+"/templates/{template_id}", _CodeTemplateService_UpdateCodeTemplate_Rule0(srv)).
		Methods("POST").
		Name("eolymp.atlas.CodeTemplateService.UpdateCodeTemplate")
	router.Handle(prefix+"/templates/{template_id}", _CodeTemplateService_DeleteCodeTemplate_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.atlas.CodeTemplateService.DeleteCodeTemplate")
	router.Handle(prefix+"/templates", _CodeTemplateService_ListCodeTemplates_Rule0(srv)).
		Methods("GET").
		Name("eolymp.atlas.CodeTemplateService.ListCodeTemplates")
	router.Handle(prefix+"/templates/{template_id}", _CodeTemplateService_DescribeCodeTemplate_Rule0(srv)).
		Methods("GET").
		Name("eolymp.atlas.CodeTemplateService.DescribeCodeTemplate")
}

func _CodeTemplateService_CreateCodeTemplate_Rule0(srv CodeTemplateServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateCodeTemplateInput{}

		if err := _CodeTemplateService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateCodeTemplate(r.Context(), in)
		if err != nil {
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CodeTemplateService_HTTPWriteResponse(w, out)
	})
}

func _CodeTemplateService_UpdateCodeTemplate_Rule0(srv CodeTemplateServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateCodeTemplateInput{}

		if err := _CodeTemplateService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TemplateId = vars["template_id"]

		out, err := srv.UpdateCodeTemplate(r.Context(), in)
		if err != nil {
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CodeTemplateService_HTTPWriteResponse(w, out)
	})
}

func _CodeTemplateService_DeleteCodeTemplate_Rule0(srv CodeTemplateServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteCodeTemplateInput{}

		if err := _CodeTemplateService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TemplateId = vars["template_id"]

		out, err := srv.DeleteCodeTemplate(r.Context(), in)
		if err != nil {
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CodeTemplateService_HTTPWriteResponse(w, out)
	})
}

func _CodeTemplateService_ListCodeTemplates_Rule0(srv CodeTemplateServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListCodeTemplatesInput{}

		if err := _CodeTemplateService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListCodeTemplates(r.Context(), in)
		if err != nil {
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CodeTemplateService_HTTPWriteResponse(w, out)
	})
}

func _CodeTemplateService_DescribeCodeTemplate_Rule0(srv CodeTemplateServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCodeTemplateInput{}

		if err := _CodeTemplateService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TemplateId = vars["template_id"]

		out, err := srv.DescribeCodeTemplate(r.Context(), in)
		if err != nil {
			_CodeTemplateService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CodeTemplateService_HTTPWriteResponse(w, out)
	})
}

type _CodeTemplateServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _CodeTemplateServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _CodeTemplateServiceHandler) (out proto.Message, err error)
type CodeTemplateServiceInterceptor struct {
	middleware []_CodeTemplateServiceMiddleware
	server     CodeTemplateServiceServer
}

// NewCodeTemplateServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewCodeTemplateServiceInterceptor(srv CodeTemplateServiceServer, middleware ..._CodeTemplateServiceMiddleware) *CodeTemplateServiceInterceptor {
	return &CodeTemplateServiceInterceptor{server: srv, middleware: middleware}
}

func (i *CodeTemplateServiceInterceptor) CreateCodeTemplate(ctx context.Context, in *CreateCodeTemplateInput) (*CreateCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateCodeTemplateInput, got %T", in))
		}

		return i.server.CreateCodeTemplate(ctx, message)
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

func (i *CodeTemplateServiceInterceptor) UpdateCodeTemplate(ctx context.Context, in *UpdateCodeTemplateInput) (*UpdateCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateCodeTemplateInput, got %T", in))
		}

		return i.server.UpdateCodeTemplate(ctx, message)
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

func (i *CodeTemplateServiceInterceptor) DeleteCodeTemplate(ctx context.Context, in *DeleteCodeTemplateInput) (*DeleteCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteCodeTemplateInput, got %T", in))
		}

		return i.server.DeleteCodeTemplate(ctx, message)
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

func (i *CodeTemplateServiceInterceptor) ListCodeTemplates(ctx context.Context, in *ListCodeTemplatesInput) (*ListCodeTemplatesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListCodeTemplatesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListCodeTemplatesInput, got %T", in))
		}

		return i.server.ListCodeTemplates(ctx, message)
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

func (i *CodeTemplateServiceInterceptor) DescribeCodeTemplate(ctx context.Context, in *DescribeCodeTemplateInput) (*DescribeCodeTemplateOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeCodeTemplateInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeCodeTemplateInput, got %T", in))
		}

		return i.server.DescribeCodeTemplate(ctx, message)
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