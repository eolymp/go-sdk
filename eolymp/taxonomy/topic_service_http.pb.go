// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package taxonomy

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

// _TopicService_HTTPReadQueryString parses body into proto.Message
func _TopicService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _TopicService_HTTPReadRequestBody parses body into proto.Message
func _TopicService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _TopicService_HTTPWriteResponse writes proto.Message to HTTP response
func _TopicService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_TopicService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _TopicService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _TopicService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterTopicServiceHttpHandlers adds handlers for for TopicServiceServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterTopicServiceHttpHandlers(router *mux.Router, prefix string, srv TopicServiceServer) {
	router.Handle(prefix+"/topics", _TopicService_CreateTopic_Rule0(srv)).
		Methods("POST").
		Name("eolymp.taxonomy.TopicService.CreateTopic")
	router.Handle(prefix+"/topics/{topic_id}", _TopicService_DeleteTopic_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.taxonomy.TopicService.DeleteTopic")
	router.Handle(prefix+"/topics/{topic_id}", _TopicService_UpdateTopic_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.taxonomy.TopicService.UpdateTopic")
	router.Handle(prefix+"/topics/{topic_id}", _TopicService_DescribeTopic_Rule0(srv)).
		Methods("GET").
		Name("eolymp.taxonomy.TopicService.DescribeTopic")
	router.Handle(prefix+"/topics", _TopicService_ListTopics_Rule0(srv)).
		Methods("GET").
		Name("eolymp.taxonomy.TopicService.ListTopics")
	router.Handle(prefix+"/topics/{topic_id}/translations/{locale}", _TopicService_TranslateTopic_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.taxonomy.TopicService.TranslateTopic")
	router.Handle(prefix+"/topics/{topic_id}/translations/{locale}", _TopicService_DeleteTranslation_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.taxonomy.TopicService.DeleteTranslation")
	router.Handle(prefix+"/topics/{topic_id}/translations", _TopicService_ListTranslations_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.taxonomy.TopicService.ListTranslations")
}

func _TopicService_CreateTopic_Rule0(srv TopicServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateTopicInput{}

		if err := _TopicService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateTopic(r.Context(), in)
		if err != nil {
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TopicService_HTTPWriteResponse(w, out)
	})
}

func _TopicService_DeleteTopic_Rule0(srv TopicServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteTopicInput{}

		if err := _TopicService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TopicId = vars["topic_id"]

		out, err := srv.DeleteTopic(r.Context(), in)
		if err != nil {
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TopicService_HTTPWriteResponse(w, out)
	})
}

func _TopicService_UpdateTopic_Rule0(srv TopicServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateTopicInput{}

		if err := _TopicService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TopicId = vars["topic_id"]

		out, err := srv.UpdateTopic(r.Context(), in)
		if err != nil {
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TopicService_HTTPWriteResponse(w, out)
	})
}

func _TopicService_DescribeTopic_Rule0(srv TopicServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeTopicInput{}

		if err := _TopicService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TopicId = vars["topic_id"]

		out, err := srv.DescribeTopic(r.Context(), in)
		if err != nil {
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TopicService_HTTPWriteResponse(w, out)
	})
}

func _TopicService_ListTopics_Rule0(srv TopicServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListTopicsInput{}

		if err := _TopicService_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListTopics(r.Context(), in)
		if err != nil {
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TopicService_HTTPWriteResponse(w, out)
	})
}

func _TopicService_TranslateTopic_Rule0(srv TopicServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &TranslateTopicInput{}

		if err := _TopicService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TopicId = vars["topic_id"]
		in.Locale = vars["locale"]

		out, err := srv.TranslateTopic(r.Context(), in)
		if err != nil {
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TopicService_HTTPWriteResponse(w, out)
	})
}

func _TopicService_DeleteTranslation_Rule0(srv TopicServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteTranslationInput{}

		if err := _TopicService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TopicId = vars["topic_id"]
		in.Locale = vars["locale"]

		out, err := srv.DeleteTranslation(r.Context(), in)
		if err != nil {
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TopicService_HTTPWriteResponse(w, out)
	})
}

func _TopicService_ListTranslations_Rule0(srv TopicServiceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListTranslationsInput{}

		if err := _TopicService_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TopicId = vars["topic_id"]

		out, err := srv.ListTranslations(r.Context(), in)
		if err != nil {
			_TopicService_HTTPWriteErrorResponse(w, err)
			return
		}

		_TopicService_HTTPWriteResponse(w, out)
	})
}

type _TopicServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _TopicServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _TopicServiceHandler) (out proto.Message, err error)
type TopicServiceInterceptor struct {
	middleware []_TopicServiceMiddleware
	server     TopicServiceServer
}

// NewTopicServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewTopicServiceInterceptor(srv TopicServiceServer, middleware ..._TopicServiceMiddleware) *TopicServiceInterceptor {
	return &TopicServiceInterceptor{server: srv, middleware: middleware}
}

func (i *TopicServiceInterceptor) CreateTopic(ctx context.Context, in *CreateTopicInput) (*CreateTopicOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateTopicInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateTopicInput, got %T", in))
		}

		return i.server.CreateTopic(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.taxonomy.TopicService.CreateTopic", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateTopicOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateTopicOutput, got %T", out))
	}

	return message, err
}

func (i *TopicServiceInterceptor) DeleteTopic(ctx context.Context, in *DeleteTopicInput) (*DeleteTopicOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteTopicInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteTopicInput, got %T", in))
		}

		return i.server.DeleteTopic(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.taxonomy.TopicService.DeleteTopic", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteTopicOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteTopicOutput, got %T", out))
	}

	return message, err
}

func (i *TopicServiceInterceptor) UpdateTopic(ctx context.Context, in *UpdateTopicInput) (*UpdateTopicOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateTopicInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateTopicInput, got %T", in))
		}

		return i.server.UpdateTopic(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.taxonomy.TopicService.UpdateTopic", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateTopicOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateTopicOutput, got %T", out))
	}

	return message, err
}

func (i *TopicServiceInterceptor) DescribeTopic(ctx context.Context, in *DescribeTopicInput) (*DescribeTopicOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeTopicInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeTopicInput, got %T", in))
		}

		return i.server.DescribeTopic(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.taxonomy.TopicService.DescribeTopic", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeTopicOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeTopicOutput, got %T", out))
	}

	return message, err
}

func (i *TopicServiceInterceptor) ListTopics(ctx context.Context, in *ListTopicsInput) (*ListTopicsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListTopicsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListTopicsInput, got %T", in))
		}

		return i.server.ListTopics(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.taxonomy.TopicService.ListTopics", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListTopicsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListTopicsOutput, got %T", out))
	}

	return message, err
}

func (i *TopicServiceInterceptor) TranslateTopic(ctx context.Context, in *TranslateTopicInput) (*TranslateTopicOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*TranslateTopicInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *TranslateTopicInput, got %T", in))
		}

		return i.server.TranslateTopic(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.taxonomy.TopicService.TranslateTopic", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*TranslateTopicOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *TranslateTopicOutput, got %T", out))
	}

	return message, err
}

func (i *TopicServiceInterceptor) DeleteTranslation(ctx context.Context, in *DeleteTranslationInput) (*DeleteTranslationOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteTranslationInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteTranslationInput, got %T", in))
		}

		return i.server.DeleteTranslation(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.taxonomy.TopicService.DeleteTranslation", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteTranslationOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteTranslationOutput, got %T", out))
	}

	return message, err
}

func (i *TopicServiceInterceptor) ListTranslations(ctx context.Context, in *ListTranslationsInput) (*ListTranslationsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListTranslationsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListTranslationsInput, got %T", in))
		}

		return i.server.ListTranslations(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.taxonomy.TopicService.ListTranslations", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListTranslationsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListTranslationsOutput, got %T", out))
	}

	return message, err
}
