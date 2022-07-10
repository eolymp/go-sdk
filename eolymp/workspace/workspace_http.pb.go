// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package workspace

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

// _Workspace_HTTPReadRequestBody parses body into proto.Message
func _Workspace_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Workspace_HTTPWriteResponse writes proto.Message to HTTP response
func _Workspace_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Workspace_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Workspace_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Workspace_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// NewWorkspaceHandler constructs new http.Handler for WorkspaceServer
func NewWorkspaceHandler(srv WorkspaceServer) http.Handler {
	router := mux.NewRouter()
	router.Handle("/eolymp.workspace.Workspace/DescribeProject", _Workspace_DescribeProject(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.workspace.Workspace/ListProjects", _Workspace_ListProjects(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.workspace.Workspace/CreateProject", _Workspace_CreateProject(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.workspace.Workspace/UpdateProject", _Workspace_UpdateProject(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.workspace.Workspace/DeleteProject", _Workspace_DeleteProject(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.workspace.Workspace/ListFiles", _Workspace_ListFiles(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.workspace.Workspace/DescribeFile", _Workspace_DescribeFile(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.workspace.Workspace/UploadFile", _Workspace_UploadFile(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.workspace.Workspace/RemoveFile", _Workspace_RemoveFile(srv)).Methods(http.MethodPost)
	return router
}

func _Workspace_DescribeProject(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeProjectInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeProject(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_ListProjects(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListProjectsInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListProjects(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_CreateProject(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateProjectInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateProject(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_UpdateProject(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateProjectInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.UpdateProject(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_DeleteProject(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteProjectInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DeleteProject(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_ListFiles(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListFilesInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListFiles(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_DescribeFile(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeFileInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeFile(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_UploadFile(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UploadFileInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.UploadFile(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_RemoveFile(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RemoveFileInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.RemoveFile(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

var promWorkspaceRequestLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "workspace_request_latency",
	Help:    "Workspace request latency",
	Buckets: []float64{0.1, 0.4, 1, 5},
}, []string{"method", "status"})

type _WorkspaceLimiter interface {
	Allow(context.Context, string, float64, int) bool
}

type WorkspaceInterceptor struct {
	limiter _WorkspaceLimiter
	server  WorkspaceServer
}

// NewWorkspaceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewWorkspaceInterceptor(srv WorkspaceServer, lim _WorkspaceLimiter) *WorkspaceInterceptor {
	return &WorkspaceInterceptor{server: srv, limiter: lim}
}

func (i *WorkspaceInterceptor) DescribeProject(ctx context.Context, in *DescribeProjectInput) (out *DescribeProjectOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promWorkspaceRequestLatency.WithLabelValues("eolymp.workspace.Workspace/DescribeProject", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("workspace:project:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: workspace:project:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.workspace.Workspace/DescribeProject", 20, 500) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribeProject(ctx, in)
	return
}

func (i *WorkspaceInterceptor) ListProjects(ctx context.Context, in *ListProjectsInput) (out *ListProjectsOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promWorkspaceRequestLatency.WithLabelValues("eolymp.workspace.Workspace/ListProjects", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("workspace:project:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: workspace:project:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.workspace.Workspace/ListProjects", 20, 100) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.ListProjects(ctx, in)
	return
}

func (i *WorkspaceInterceptor) CreateProject(ctx context.Context, in *CreateProjectInput) (out *CreateProjectOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promWorkspaceRequestLatency.WithLabelValues("eolymp.workspace.Workspace/CreateProject", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("workspace:project:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: workspace:project:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.workspace.Workspace/CreateProject", 1, 10) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.CreateProject(ctx, in)
	return
}

func (i *WorkspaceInterceptor) UpdateProject(ctx context.Context, in *UpdateProjectInput) (out *UpdateProjectOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promWorkspaceRequestLatency.WithLabelValues("eolymp.workspace.Workspace/UpdateProject", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("workspace:project:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: workspace:project:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.workspace.Workspace/UpdateProject", 5, 50) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.UpdateProject(ctx, in)
	return
}

func (i *WorkspaceInterceptor) DeleteProject(ctx context.Context, in *DeleteProjectInput) (out *DeleteProjectOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promWorkspaceRequestLatency.WithLabelValues("eolymp.workspace.Workspace/DeleteProject", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("workspace:project:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: workspace:project:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.workspace.Workspace/DeleteProject", 5, 50) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DeleteProject(ctx, in)
	return
}

func (i *WorkspaceInterceptor) ListFiles(ctx context.Context, in *ListFilesInput) (out *ListFilesOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promWorkspaceRequestLatency.WithLabelValues("eolymp.workspace.Workspace/ListFiles", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("workspace:project:read") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: workspace:project:read")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.workspace.Workspace/ListFiles", 20, 100) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.ListFiles(ctx, in)
	return
}

func (i *WorkspaceInterceptor) DescribeFile(ctx context.Context, in *DescribeFileInput) (out *DescribeFileOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promWorkspaceRequestLatency.WithLabelValues("eolymp.workspace.Workspace/DescribeFile", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("workspace:project:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: workspace:project:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.workspace.Workspace/DescribeFile", 5, 50) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.DescribeFile(ctx, in)
	return
}

func (i *WorkspaceInterceptor) UploadFile(ctx context.Context, in *UploadFileInput) (out *UploadFileOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promWorkspaceRequestLatency.WithLabelValues("eolymp.workspace.Workspace/UploadFile", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("workspace:project:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: workspace:project:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.workspace.Workspace/UploadFile", 5, 50) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.UploadFile(ctx, in)
	return
}

func (i *WorkspaceInterceptor) RemoveFile(ctx context.Context, in *RemoveFileInput) (out *RemoveFileOutput, err error) {
	start := time.Now()
	defer func() {
		s, _ := status.FromError(err)
		if s == nil {
			s = status.New(codes.OK, "OK")
		}

		promWorkspaceRequestLatency.WithLabelValues("eolymp.workspace.Workspace/RemoveFile", s.Code().String()).
			Observe(time.Since(start).Seconds())
	}()

	token, ok := oauth.TokenFromContext(ctx)
	if ok && !token.Has("workspace:project:write") {
		err = status.Error(codes.PermissionDenied, "required token scopes are missing: workspace:project:write")
		return
	}

	if !i.limiter.Allow(ctx, "eolymp.workspace.Workspace/RemoveFile", 5, 50) {
		err = status.Error(codes.ResourceExhausted, "too many requests")
		return
	}

	out, err = i.server.RemoveFile(ctx, in)
	return
}
