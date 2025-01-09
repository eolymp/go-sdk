// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package runtime

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

// _RuntimeService_HTTPReadQueryString parses body into proto.Message
func _RuntimeService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _RuntimeService_HTTPReadRequestBody parses body into proto.Message
func _RuntimeService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _RuntimeService_HTTPWriteResponse writes proto.Message to HTTP response
func _RuntimeService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_RuntimeService_HTTPWriteErrorResponse(w, err)
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

// _RuntimeService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _RuntimeService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterRuntimeServiceHttpHandlers adds handlers for for RuntimeServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterRuntimeServiceHttpHandlers(router *mux.Router, prefix string, cli RuntimeServiceClient) {
	router.Handle(prefix+"/languages/{language_id}", _RuntimeService_DescribeLanguage_Rule0(cli)).
		Methods("GET").
		Name("eolymp.runtime.RuntimeService.DescribeLanguage")
	router.Handle(prefix+"/languages", _RuntimeService_ListLanguages_Rule0(cli)).
		Methods("GET").
		Name("eolymp.runtime.RuntimeService.ListLanguages")
	router.Handle(prefix+"/runtime/{runtime_id}", _RuntimeService_DescribeRuntime_Rule0(cli)).
		Methods("GET").
		Name("eolymp.runtime.RuntimeService.DescribeRuntime")
	router.Handle(prefix+"/runtime", _RuntimeService_ListRuntimes_Rule0(cli)).
		Methods("GET").
		Name("eolymp.runtime.RuntimeService.ListRuntimes")
	router.Handle(prefix+"/runtime/{runtime_id}/template", _RuntimeService_DescribeCodeTemplate_Rule0(cli)).
		Methods("GET").
		Name("eolymp.runtime.RuntimeService.DescribeCodeTemplate")
}

func _RuntimeService_DescribeLanguage_Rule0(cli RuntimeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeLanguageInput{}

		if err := _RuntimeService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RuntimeService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.LanguageId = vars["language_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeLanguage(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RuntimeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RuntimeService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RuntimeService_ListLanguages_Rule0(cli RuntimeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListLanguagesInput{}

		if err := _RuntimeService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RuntimeService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListLanguages(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RuntimeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RuntimeService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RuntimeService_DescribeRuntime_Rule0(cli RuntimeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRuntimeInput{}

		if err := _RuntimeService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RuntimeService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RuntimeId = vars["runtime_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeRuntime(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RuntimeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RuntimeService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RuntimeService_ListRuntimes_Rule0(cli RuntimeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListRuntimesInput{}

		if err := _RuntimeService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RuntimeService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListRuntimes(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RuntimeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RuntimeService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RuntimeService_DescribeCodeTemplate_Rule0(cli RuntimeServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCodeTemplateInput{}

		if err := _RuntimeService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RuntimeService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RuntimeId = vars["runtime_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeCodeTemplate(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RuntimeService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RuntimeService_HTTPWriteResponse(w, out, header, trailer)
	})
}
