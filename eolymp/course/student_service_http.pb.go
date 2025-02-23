// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package course

import (
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

// _StudentService_HTTPReadQueryString parses body into proto.Message
func _StudentService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _StudentService_HTTPReadRequestBody parses body into proto.Message
func _StudentService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _StudentService_HTTPWriteResponse writes proto.Message to HTTP response
func _StudentService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_StudentService_HTTPWriteErrorResponse(w, err)
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

// _StudentService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _StudentService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterStudentServiceHttpHandlers adds handlers for for StudentServiceClient
func RegisterStudentServiceHttpHandlers(router *mux.Router, prefix string, cli StudentServiceClient) {
	router.Handle(prefix+"/students", _StudentService_CreateStudent_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.StudentService.CreateStudent")
	router.Handle(prefix+"/students/{member_id}", _StudentService_UpdateStudent_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.StudentService.UpdateStudent")
	router.Handle(prefix+"/students/{member_id}", _StudentService_DeleteStudent_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.course.StudentService.DeleteStudent")
	router.Handle(prefix+"/students/{member_id}", _StudentService_DescribeStudent_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.StudentService.DescribeStudent")
	router.Handle(prefix+"/students", _StudentService_ListStudents_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.StudentService.ListStudents")
	router.Handle(prefix+"/join", _StudentService_JoinCourse_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.StudentService.JoinCourse")
	router.Handle(prefix+"/viewer/student", _StudentService_DescribeViewer_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.StudentService.DescribeViewer")
	router.Handle(prefix+"/students/{member_id}/assignments", _StudentService_ListStudentAssignments_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.StudentService.ListStudentAssignments")
	router.Handle(prefix+"/students/{member_id}/assignments", _StudentService_UpdateStudentAssignment_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.StudentService.UpdateStudentAssignment")
	router.Handle(prefix+"/students/{member_id}/assignments", _StudentService_DeleteStudentAssignment_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.course.StudentService.DeleteStudentAssignment")
	router.Handle(prefix+"/students/{member_id}/grades", _StudentService_ListStudentGrades_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.StudentService.ListStudentGrades")
	router.Handle(prefix+"/students/{member_id}/grades/{module_id}", _StudentService_ListModuleGrades_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.StudentService.ListModuleGrades")
}

// RegisterStudentServiceHttpProxy adds proxy handlers for for StudentServiceClient
func RegisterStudentServiceHttpProxy(router *mux.Router, prefix string, conn grpc.ClientConnInterface) {
	RegisterStudentServiceHttpHandlers(router, prefix, NewStudentServiceClient(conn))
}

func _StudentService_CreateStudent_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateStudentInput{}

		if err := _StudentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateStudent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_UpdateStudent_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateStudentInput{}

		if err := _StudentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateStudent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_DeleteStudent_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteStudentInput{}

		if err := _StudentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteStudent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_DescribeStudent_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeStudentInput{}

		if err := _StudentService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeStudent(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_ListStudents_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListStudentsInput{}

		if err := _StudentService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListStudents(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_JoinCourse_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &JoinCourseInput{}

		if err := _StudentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.JoinCourse(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_DescribeViewer_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeViewerInput{}

		if err := _StudentService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeViewer(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_ListStudentAssignments_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListStudentAssignmentsInput{}

		if err := _StudentService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.ListStudentAssignments(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_UpdateStudentAssignment_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateStudentAssignmentInput{}

		if err := _StudentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateStudentAssignment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_DeleteStudentAssignment_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteStudentAssignmentInput{}

		if err := _StudentService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteStudentAssignment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_ListStudentGrades_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListStudentGradesInput{}

		if err := _StudentService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.ListStudentGrades(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _StudentService_ListModuleGrades_Rule0(cli StudentServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListModuleGradesInput{}

		if err := _StudentService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]
		in.ModuleId = vars["module_id"]

		var header, trailer metadata.MD

		out, err := cli.ListModuleGrades(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_StudentService_HTTPWriteErrorResponse(w, err)
			return
		}

		_StudentService_HTTPWriteResponse(w, out, header, trailer)
	})
}
