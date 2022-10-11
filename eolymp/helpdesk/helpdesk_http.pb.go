// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package helpdesk

import (
	context "context"
	oauth "github.com/eolymp/go-packages/oauth"
	mux "github.com/gorilla/mux"
	prometheus "github.com/prometheus/client_golang/prometheus"
	promauto "github.com/prometheus/client_golang/prometheus/promauto"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
	time "time"
)

// _Helpdesk_HTTPReadRequestBody parses body into proto.Message
func _Helpdesk_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Helpdesk_HTTPWriteResponse writes proto.Message to HTTP response
func _Helpdesk_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Helpdesk_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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

// NewHelpdeskHandler constructs new http.Handler for HelpdeskServer
func NewHelpdeskHandler(srv HelpdeskServer) http.Handler {
	router := mux.NewRouter()
	router.Handle("/eolymp.helpdesk.Helpdesk/DescribeDocument", _Helpdesk_DescribeDocument(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.helpdesk.Helpdesk/ListDocuments", _Helpdesk_ListDocuments(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.helpdesk.Helpdesk/CreateDocument", _Helpdesk_CreateDocument(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.helpdesk.Helpdesk/UpdateDocument", _Helpdesk_UpdateDocument(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.helpdesk.Helpdesk/DeleteDocument", _Helpdesk_DeleteDocument(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.helpdesk.Helpdesk/DescribePath", _Helpdesk_DescribePath(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.helpdesk.Helpdesk/ListPaths", _Helpdesk_ListPaths(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.helpdesk.Helpdesk/ListParents", _Helpdesk_ListParents(srv)).Methods(http.MethodPost)
	return router
}

// NewHelpdeskHandlerHttp constructs new http.Handler for HelpdeskServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func NewHelpdeskHandlerHttp(srv HelpdeskServer, prefix string) http.Handler {
	router := mux.NewRouter()

	router.Handle(prefix+"/helpdesk/documents/{document_id}", _Helpdesk_DescribeDocument_Rule0(srv)).
		Methods("GET").
		Name("eolymp.helpdesk.Helpdesk.DescribeDocument")

	router.Handle(prefix+"/helpdesk/documents", _Helpdesk_ListDocuments_Rule0(srv)).
		Methods("GET").
		Name("eolymp.helpdesk.Helpdesk.ListDocuments")

	router.Handle(prefix+"/helpdesk/documents", _Helpdesk_CreateDocument_Rule0(srv)).
		Methods("POST").
		Name("eolymp.helpdesk.Helpdesk.CreateDocument")

	router.Handle(prefix+"/helpdesk/documents/{document_id}", _Helpdesk_UpdateDocument_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.helpdesk.Helpdesk.UpdateDocument")

	router.Handle(prefix+"/helpdesk/documents/{document_id}", _Helpdesk_DeleteDocument_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.helpdesk.Helpdesk.DeleteDocument")

	router.Handle(prefix+"/helpdesk/paths/{path}", _Helpdesk_DescribePath_Rule0(srv)).
		Methods("GET").
		Name("eolymp.helpdesk.Helpdesk.DescribePath")

	router.Handle(prefix+"/helpdesk/paths", _Helpdesk_ListPaths_Rule0(srv)).
		Methods("GET").
		Name("eolymp.helpdesk.Helpdesk.ListPaths")

	router.Handle(prefix+"/helpdesk/paths/{path}/parents", _Helpdesk_ListParents_Rule0(srv)).
		Methods("GET").
		Name("eolymp.helpdesk.Helpdesk.ListParents")

	return router
}

func _Helpdesk_DescribeDocument(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeDocumentInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeDocument(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_ListDocuments(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListDocumentsInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListDocuments(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_CreateDocument(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateDocumentInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateDocument(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_UpdateDocument(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateDocumentInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.UpdateDocument(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_DeleteDocument(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteDocumentInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DeleteDocument(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_DescribePath(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePathInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribePath(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_ListPaths(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPathsInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListPaths(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_ListParents(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListParentsInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListParents(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_DescribeDocument_Rule0(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeDocumentInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.documentId = vars["document_id"]

		out, err := srv.DescribeDocument(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_ListDocuments_Rule0(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListDocumentsInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListDocuments(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_CreateDocument_Rule0(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateDocumentInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateDocument(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_UpdateDocument_Rule0(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateDocumentInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.documentId = vars["document_id"]

		out, err := srv.UpdateDocument(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_DeleteDocument_Rule0(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteDocumentInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.documentId = vars["document_id"]

		out, err := srv.DeleteDocument(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_DescribePath_Rule0(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePathInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.path = vars["path"]

		out, err := srv.DescribePath(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_ListPaths_Rule0(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPathsInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListPaths(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

func _Helpdesk_ListParents_Rule0(srv HelpdeskServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListParentsInput{}

		if err := _Helpdesk_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.path = vars["path"]

		out, err := srv.ListParents(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

var promHelpdeskRequestLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "helpdesk_request_latency",
	Help:    "Helpdesk request latency",
	Buckets: []float64{0.1, 0.4, 1, 5},
}, []string{"method", "status"})

type _HelpdeskLimiter interface {
	Allow(context.Context, string, float64, int) bool
}

type HelpdeskInterceptor struct {
	limiter _HelpdeskLimiter
	server  HelpdeskServer
}

// NewHelpdeskInterceptor constructs additional middleware for a server based on annotations in proto files
func NewHelpdeskInterceptor(srv HelpdeskServer, lim _HelpdeskLimiter) *HelpdeskInterceptor {
	return &HelpdeskInterceptor{server: srv, limiter: lim}
}

func (i *HelpdeskInterceptor) DescribeDocument(ctx context.Context, in *DescribeDocumentInput) (out *DescribeDocumentOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promHelpdeskRequestLatency.WithLabelValues("eolymp.helpdesk.Helpdesk/DescribeDocument", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("helpdesk:document:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: helpdesk:document:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.helpdesk.Helpdesk/DescribeDocument", 20, 500) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribeDocument(ctx, in)
	return
}

func (i *HelpdeskInterceptor) ListDocuments(ctx context.Context, in *ListDocumentsInput) (out *ListDocumentsOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promHelpdeskRequestLatency.WithLabelValues("eolymp.helpdesk.Helpdesk/ListDocuments", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("helpdesk:document:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: helpdesk:document:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.helpdesk.Helpdesk/ListDocuments", 20, 100) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.ListDocuments(ctx, in)
	return
}

func (i *HelpdeskInterceptor) CreateDocument(ctx context.Context, in *CreateDocumentInput) (out *CreateDocumentOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promHelpdeskRequestLatency.WithLabelValues("eolymp.helpdesk.Helpdesk/CreateDocument", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("helpdesk:document:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: helpdesk:document:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.helpdesk.Helpdesk/CreateDocument", 5, 50) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.CreateDocument(ctx, in)
	return
}

func (i *HelpdeskInterceptor) UpdateDocument(ctx context.Context, in *UpdateDocumentInput) (out *UpdateDocumentOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promHelpdeskRequestLatency.WithLabelValues("eolymp.helpdesk.Helpdesk/UpdateDocument", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("helpdesk:document:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: helpdesk:document:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.helpdesk.Helpdesk/UpdateDocument", 5, 50) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.UpdateDocument(ctx, in)
	return
}

func (i *HelpdeskInterceptor) DeleteDocument(ctx context.Context, in *DeleteDocumentInput) (out *DeleteDocumentOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promHelpdeskRequestLatency.WithLabelValues("eolymp.helpdesk.Helpdesk/DeleteDocument", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("helpdesk:document:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: helpdesk:document:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.helpdesk.Helpdesk/DeleteDocument", 5, 50) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DeleteDocument(ctx, in)
	return
}

func (i *HelpdeskInterceptor) DescribePath(ctx context.Context, in *DescribePathInput) (out *DescribePathOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promHelpdeskRequestLatency.WithLabelValues("eolymp.helpdesk.Helpdesk/DescribePath", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("helpdesk:document:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: helpdesk:document:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.helpdesk.Helpdesk/DescribePath", 20, 500) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribePath(ctx, in)
	return
}

func (i *HelpdeskInterceptor) ListPaths(ctx context.Context, in *ListPathsInput) (out *ListPathsOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promHelpdeskRequestLatency.WithLabelValues("eolymp.helpdesk.Helpdesk/ListPaths", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("helpdesk:document:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: helpdesk:document:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.helpdesk.Helpdesk/ListPaths", 20, 500) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.ListPaths(ctx, in)
	return
}

func (i *HelpdeskInterceptor) ListParents(ctx context.Context, in *ListParentsInput) (out *ListParentsOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promHelpdeskRequestLatency.WithLabelValues("eolymp.helpdesk.Helpdesk/ListParents", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("helpdesk:document:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: helpdesk:document:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.helpdesk.Helpdesk/ListParents", 20, 500) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.ListParents(ctx, in)
	return
}
