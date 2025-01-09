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

// _CourseService_HTTPReadQueryString parses body into proto.Message
func _CourseService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _CourseService_HTTPReadRequestBody parses body into proto.Message
func _CourseService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _CourseService_HTTPWriteResponse writes proto.Message to HTTP response
func _CourseService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_CourseService_HTTPWriteErrorResponse(w, err)
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

// _CourseService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _CourseService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterCourseServiceHttpHandlers adds handlers for for CourseServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterCourseServiceHttpHandlers(router *mux.Router, prefix string, cli CourseServiceClient) {
	router.Handle(prefix+"/courses", _CourseService_CreateCourse_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.CourseService.CreateCourse")
	router.Handle(prefix+"/courses/{course_id}", _CourseService_UpdateCourse_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.CourseService.UpdateCourse")
	router.Handle(prefix+"/courses/{course_id}", _CourseService_DeleteCourse_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.course.CourseService.DeleteCourse")
	router.Handle(prefix+"/courses/{course_id}", _CourseService_DescribeCourse_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.CourseService.DescribeCourse")
	router.Handle(prefix+"/courses", _CourseService_ListCourses_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.CourseService.ListCourses")
	router.Handle(prefix+"/courses/{course_id}/copy", _CourseService_CopyCourse_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.CourseService.CopyCourse")
}

func _CourseService_CreateCourse_Rule0(cli CourseServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateCourseInput{}

		if err := _CourseService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateCourse(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CourseService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _CourseService_UpdateCourse_Rule0(cli CourseServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateCourseInput{}

		if err := _CourseService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.CourseId = vars["course_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateCourse(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CourseService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _CourseService_DeleteCourse_Rule0(cli CourseServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteCourseInput{}

		if err := _CourseService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.CourseId = vars["course_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteCourse(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CourseService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _CourseService_DescribeCourse_Rule0(cli CourseServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCourseInput{}

		if err := _CourseService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.CourseId = vars["course_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeCourse(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CourseService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _CourseService_ListCourses_Rule0(cli CourseServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListCoursesInput{}

		if err := _CourseService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListCourses(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CourseService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _CourseService_CopyCourse_Rule0(cli CourseServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CopyCourseInput{}

		if err := _CourseService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.CourseId = vars["course_id"]

		var header, trailer metadata.MD

		out, err := cli.CopyCourse(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_CourseService_HTTPWriteErrorResponse(w, err)
			return
		}

		_CourseService_HTTPWriteResponse(w, out, header, trailer)
	})
}
