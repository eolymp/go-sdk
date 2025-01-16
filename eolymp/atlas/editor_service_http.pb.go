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

// _EditorService_HTTPReadQueryString parses body into proto.Message
func _EditorService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _EditorService_HTTPReadRequestBody parses body into proto.Message
func _EditorService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _EditorService_HTTPWriteResponse writes proto.Message to HTTP response
func _EditorService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_EditorService_HTTPWriteErrorResponse(w, err)
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

// _EditorService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _EditorService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterEditorServiceHttpHandlers adds handlers for for EditorServiceClient
func RegisterEditorServiceHttpHandlers(router *mux.Router, prefix string, cli EditorServiceClient) {
	router.Handle(prefix+"/editor", _EditorService_DescribeEditor_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.EditorService.DescribeEditor")
	router.Handle(prefix+"/editor/state", _EditorService_DescribeEditorState_Rule0(cli)).
		Methods("GET").
		Name("eolymp.atlas.EditorService.DescribeEditorState")
	router.Handle(prefix+"/editor/state", _EditorService_UpdateEditorState_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.EditorService.UpdateEditorState")
	router.Handle(prefix+"/editor/print", _EditorService_PrintEditorCode_Rule0(cli)).
		Methods("POST").
		Name("eolymp.atlas.EditorService.PrintEditorCode")
}

// RegisterEditorServiceHttpProxy adds proxy handlers for for EditorServiceClient
func RegisterEditorServiceHttpProxy(router *mux.Router, prefix string, conn grpc.ClientConnInterface) {
	RegisterEditorServiceHttpHandlers(router, prefix, NewEditorServiceClient(conn))
}

func _EditorService_DescribeEditor_Rule0(cli EditorServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeEditorInput{}

		if err := _EditorService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeEditor(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _EditorService_DescribeEditorState_Rule0(cli EditorServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeEditorStateInput{}

		if err := _EditorService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeEditorState(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _EditorService_UpdateEditorState_Rule0(cli EditorServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateEditorStateInput{}

		if err := _EditorService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.UpdateEditorState(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _EditorService_PrintEditorCode_Rule0(cli EditorServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &PrintEditorCodeInput{}

		if err := _EditorService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_EditorService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.PrintEditorCode(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_EditorService_HTTPWriteErrorResponse(w, err)
			return
		}

		_EditorService_HTTPWriteResponse(w, out, header, trailer)
	})
}
