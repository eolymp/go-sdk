// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package helpdesk

import (
	context "context"
	mux "github.com/gorilla/mux"
	codes "google.golang.org/grpc/codes"
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

		if err := _Helpdesk_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.DocumentId = vars["document_id"]

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

		if err := _Helpdesk_HTTPReadQueryString(r, in); err != nil {
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
		in.DocumentId = vars["document_id"]

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
		in.DocumentId = vars["document_id"]

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

		if err := _Helpdesk_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.Path = vars["path"]

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

		if err := _Helpdesk_HTTPReadQueryString(r, in); err != nil {
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

		if err := _Helpdesk_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.Path = vars["path"]

		out, err := srv.ListParents(r.Context(), in)
		if err != nil {
			_Helpdesk_HTTPWriteErrorResponse(w, err)
			return
		}

		_Helpdesk_HTTPWriteResponse(w, out)
	})
}

type _HelpdeskMiddleware func(ctx context.Context, method string, in proto.Message, next func() (out proto.Message, err error))
type HelpdeskInterceptor struct {
	middleware []_HelpdeskMiddleware
	server     HelpdeskServer
}

// NewHelpdeskInterceptor constructs additional middleware for a server based on annotations in proto files
func NewHelpdeskInterceptor(srv HelpdeskServer, middleware ..._HelpdeskMiddleware) *HelpdeskInterceptor {
	return &HelpdeskInterceptor{server: srv, middleware: middleware}
}

func (i *HelpdeskInterceptor) DescribeDocument(ctx context.Context, in *DescribeDocumentInput) (out *DescribeDocumentOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.DescribeDocument(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.helpdesk.Helpdesk/DescribeDocument", in, handler)
			return out, err
		}
	}

	next()
	return
}

func (i *HelpdeskInterceptor) ListDocuments(ctx context.Context, in *ListDocumentsInput) (out *ListDocumentsOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.ListDocuments(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.helpdesk.Helpdesk/ListDocuments", in, handler)
			return out, err
		}
	}

	next()
	return
}

func (i *HelpdeskInterceptor) CreateDocument(ctx context.Context, in *CreateDocumentInput) (out *CreateDocumentOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.CreateDocument(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.helpdesk.Helpdesk/CreateDocument", in, handler)
			return out, err
		}
	}

	next()
	return
}

func (i *HelpdeskInterceptor) UpdateDocument(ctx context.Context, in *UpdateDocumentInput) (out *UpdateDocumentOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.UpdateDocument(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.helpdesk.Helpdesk/UpdateDocument", in, handler)
			return out, err
		}
	}

	next()
	return
}

func (i *HelpdeskInterceptor) DeleteDocument(ctx context.Context, in *DeleteDocumentInput) (out *DeleteDocumentOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.DeleteDocument(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.helpdesk.Helpdesk/DeleteDocument", in, handler)
			return out, err
		}
	}

	next()
	return
}

func (i *HelpdeskInterceptor) DescribePath(ctx context.Context, in *DescribePathInput) (out *DescribePathOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.DescribePath(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.helpdesk.Helpdesk/DescribePath", in, handler)
			return out, err
		}
	}

	next()
	return
}

func (i *HelpdeskInterceptor) ListPaths(ctx context.Context, in *ListPathsInput) (out *ListPathsOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.ListPaths(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.helpdesk.Helpdesk/ListPaths", in, handler)
			return out, err
		}
	}

	next()
	return
}

func (i *HelpdeskInterceptor) ListParents(ctx context.Context, in *ListParentsInput) (out *ListParentsOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.ListParents(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.helpdesk.Helpdesk/ListParents", in, handler)
			return out, err
		}
	}

	next()
	return
}
