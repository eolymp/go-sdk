// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package workspace

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

// _Workspace_HTTPReadQueryString parses body into proto.Message
func _Workspace_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

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

// NewWorkspaceHandlerHttp constructs new http.Handler for WorkspaceServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func NewWorkspaceHandlerHttp(srv WorkspaceServer, prefix string) http.Handler {
	router := mux.NewRouter()

	router.Handle(prefix+"/workspace/projects/{project_id}", _Workspace_DescribeProject_Rule0(srv)).
		Methods("GET").
		Name("eolymp.workspace.Workspace.DescribeProject")

	router.Handle(prefix+"/workspace/projects", _Workspace_ListProjects_Rule0(srv)).
		Methods("GET").
		Name("eolymp.workspace.Workspace.ListProjects")

	router.Handle(prefix+"/workspace/projects", _Workspace_CreateProject_Rule0(srv)).
		Methods("POST").
		Name("eolymp.workspace.Workspace.CreateProject")

	router.Handle(prefix+"/workspace/projects/{project_id}", _Workspace_UpdateProject_Rule0(srv)).
		Methods("POST").
		Name("eolymp.workspace.Workspace.UpdateProject")

	router.Handle(prefix+"/workspace/projects/{project_id}", _Workspace_DeleteProject_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.workspace.Workspace.DeleteProject")

	router.Handle(prefix+"/workspace/projects/{project_id}/files", _Workspace_ListFiles_Rule0(srv)).
		Methods("GET").
		Name("eolymp.workspace.Workspace.ListFiles")

	router.Handle(prefix+"/workspace/projects/{project_id}/files/{name}", _Workspace_DescribeFile_Rule0(srv)).
		Methods("GET").
		Name("eolymp.workspace.Workspace.DescribeFile")

	router.Handle(prefix+"/workspace/projects/{project_id}/files", _Workspace_UploadFile_Rule0(srv)).
		Methods("POST").
		Name("eolymp.workspace.Workspace.UploadFile")

	router.Handle(prefix+"/workspace/projects/{project_id}/files/{name}", _Workspace_RemoveFile_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.workspace.Workspace.RemoveFile")

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

func _Workspace_DescribeProject_Rule0(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeProjectInput{}

		if err := _Workspace_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProjectId = vars["project_id"]

		out, err := srv.DescribeProject(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_ListProjects_Rule0(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListProjectsInput{}

		if err := _Workspace_HTTPReadQueryString(r, in); err != nil {
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

func _Workspace_CreateProject_Rule0(srv WorkspaceServer) http.Handler {
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

func _Workspace_UpdateProject_Rule0(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateProjectInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProjectId = vars["project_id"]

		out, err := srv.UpdateProject(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_DeleteProject_Rule0(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteProjectInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProjectId = vars["project_id"]

		out, err := srv.DeleteProject(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_ListFiles_Rule0(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListFilesInput{}

		if err := _Workspace_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProjectId = vars["project_id"]

		out, err := srv.ListFiles(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_DescribeFile_Rule0(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeFileInput{}

		if err := _Workspace_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProjectId = vars["project_id"]
		in.Name = vars["name"]

		out, err := srv.DescribeFile(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_UploadFile_Rule0(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UploadFileInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProjectId = vars["project_id"]

		out, err := srv.UploadFile(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

func _Workspace_RemoveFile_Rule0(srv WorkspaceServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RemoveFileInput{}

		if err := _Workspace_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ProjectId = vars["project_id"]
		in.Name = vars["name"]

		out, err := srv.RemoveFile(r.Context(), in)
		if err != nil {
			_Workspace_HTTPWriteErrorResponse(w, err)
			return
		}

		_Workspace_HTTPWriteResponse(w, out)
	})
}

type _WorkspaceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _WorkspaceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _WorkspaceHandler) (out proto.Message, err error)
type WorkspaceInterceptor struct {
	middleware []_WorkspaceMiddleware
	server     WorkspaceServer
}

// NewWorkspaceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewWorkspaceInterceptor(srv WorkspaceServer, middleware ..._WorkspaceMiddleware) *WorkspaceInterceptor {
	return &WorkspaceInterceptor{server: srv, middleware: middleware}
}

func (i *WorkspaceInterceptor) DescribeProject(ctx context.Context, in *DescribeProjectInput) (*DescribeProjectOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeProjectInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeProjectInput, got %T", in))
		}

		return i.server.DescribeProject(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.workspace.Workspace.DescribeProject", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeProjectOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeProjectOutput, got %T", out))
	}

	return message, err
}

func (i *WorkspaceInterceptor) ListProjects(ctx context.Context, in *ListProjectsInput) (*ListProjectsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListProjectsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListProjectsInput, got %T", in))
		}

		return i.server.ListProjects(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.workspace.Workspace.ListProjects", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListProjectsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListProjectsOutput, got %T", out))
	}

	return message, err
}

func (i *WorkspaceInterceptor) CreateProject(ctx context.Context, in *CreateProjectInput) (*CreateProjectOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateProjectInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateProjectInput, got %T", in))
		}

		return i.server.CreateProject(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.workspace.Workspace.CreateProject", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateProjectOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateProjectOutput, got %T", out))
	}

	return message, err
}

func (i *WorkspaceInterceptor) UpdateProject(ctx context.Context, in *UpdateProjectInput) (*UpdateProjectOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateProjectInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateProjectInput, got %T", in))
		}

		return i.server.UpdateProject(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.workspace.Workspace.UpdateProject", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateProjectOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateProjectOutput, got %T", out))
	}

	return message, err
}

func (i *WorkspaceInterceptor) DeleteProject(ctx context.Context, in *DeleteProjectInput) (*DeleteProjectOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteProjectInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteProjectInput, got %T", in))
		}

		return i.server.DeleteProject(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.workspace.Workspace.DeleteProject", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteProjectOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteProjectOutput, got %T", out))
	}

	return message, err
}

func (i *WorkspaceInterceptor) ListFiles(ctx context.Context, in *ListFilesInput) (*ListFilesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListFilesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListFilesInput, got %T", in))
		}

		return i.server.ListFiles(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.workspace.Workspace.ListFiles", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListFilesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListFilesOutput, got %T", out))
	}

	return message, err
}

func (i *WorkspaceInterceptor) DescribeFile(ctx context.Context, in *DescribeFileInput) (*DescribeFileOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeFileInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeFileInput, got %T", in))
		}

		return i.server.DescribeFile(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.workspace.Workspace.DescribeFile", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeFileOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeFileOutput, got %T", out))
	}

	return message, err
}

func (i *WorkspaceInterceptor) UploadFile(ctx context.Context, in *UploadFileInput) (*UploadFileOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UploadFileInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UploadFileInput, got %T", in))
		}

		return i.server.UploadFile(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.workspace.Workspace.UploadFile", in, next)
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

func (i *WorkspaceInterceptor) RemoveFile(ctx context.Context, in *RemoveFileInput) (*RemoveFileOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RemoveFileInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RemoveFileInput, got %T", in))
		}

		return i.server.RemoveFile(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.workspace.Workspace.RemoveFile", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RemoveFileOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RemoveFileOutput, got %T", out))
	}

	return message, err
}
