// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package atlas

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

// _EditorialService_HTTPReadQueryString parses body into proto.Message
func _EditorialService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _EditorialService_HTTPReadRequestBody parses body into proto.Message
func _EditorialService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _EditorialService_HTTPWriteResponse writes proto.Message to HTTP response
func _EditorialService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_EditorialService_HTTPWriteErrorResponse(w, err)
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

// _EditorialService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _EditorialService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterEditorialServiceHttpHandlers adds handlers for for EditorialServiceClient
func RegisterEditorialServiceHttpHandlers(router *mux.Router, prefix string, cli EditorialServiceClient) {
	router.Handle(prefix+"/editorials", _EditorialService_CreateEditorial_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.EditorialService.CreateEditorial")
	router.Handle(prefix+"/editorials/{editorial_id}", _EditorialService_UpdateEditorial_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.atlas.EditorialService.UpdateEditorial")
	router.Handle(prefix+"/editorials/{editorial_id}", _EditorialService_DeleteEditorial_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.atlas.EditorialService.DeleteEditorial")
	router.Handle(prefix+"/editorials/{editorial_id}", _EditorialService_DescribeEditorial_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.EditorialService.DescribeEditorial")
	router.Handle(prefix+"/editorial", _EditorialService_LookupEditorial_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.EditorialService.LookupEditorial")
	router.Handle(prefix+"/editorial/preview", _EditorialService_PreviewEditorial_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.EditorialService.PreviewEditorial")
	router.Handle(prefix+"/editorials", _EditorialService_ListEditorials_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.EditorialService.ListEditorials")
	router.Handle(prefix+"/editorials:translate", _EditorialService_TranslateEditorials_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.EditorialService.TranslateEditorials")
}

// RegisterEditorialServiceHttpProxy adds proxy handlers for for EditorialServiceClient
func RegisterEditorialServiceHttpProxy(router *mux.Router, prefix string, conn grpc.ClientConnInterface) {
	RegisterEditorialServiceHttpHandlers(router, prefix, NewEditorialServiceClient(conn))
}

func _EditorialService_CreateEditorial_Rule0(cli EditorialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateEditorialInput{}

		if err := _EditorialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateEditorial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _EditorialService_UpdateEditorial_Rule0(cli EditorialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateEditorialInput{}

		if err := _EditorialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EditorialId = vars["editorial_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateEditorial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _EditorialService_DeleteEditorial_Rule0(cli EditorialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteEditorialInput{}

		if err := _EditorialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EditorialId = vars["editorial_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteEditorial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _EditorialService_DescribeEditorial_Rule0(cli EditorialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeEditorialInput{}

		if err := _EditorialService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.EditorialId = vars["editorial_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeEditorial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _EditorialService_LookupEditorial_Rule0(cli EditorialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LookupEditorialInput{}

		if err := _EditorialService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.LookupEditorial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _EditorialService_PreviewEditorial_Rule0(cli EditorialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &PreviewEditorialInput{}

		if err := _EditorialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.PreviewEditorial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _EditorialService_ListEditorials_Rule0(cli EditorialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListEditorialsInput{}

		if err := _EditorialService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListEditorials(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _EditorialService_TranslateEditorials_Rule0(cli EditorialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &TranslateEditorialsInput{}

		if err := _EditorialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.TranslateEditorials(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorialService_HTTPWriteResponse(w, out, header, trailer)
	})
}
