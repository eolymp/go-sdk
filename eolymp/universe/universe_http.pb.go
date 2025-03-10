// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package universe

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

// _Universe_HTTPReadQueryString parses body into proto.Message
func _Universe_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Universe_HTTPReadRequestBody parses body into proto.Message
func _Universe_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Universe_HTTPWriteResponse writes proto.Message to HTTP response
func _Universe_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Universe_HTTPWriteErrorResponse(w, err)
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

// _Universe_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Universe_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterUniverseHttpHandlers adds handlers for for UniverseClient
func RegisterUniverseHttpHandlers(router *mux.Router, prefix string, cli UniverseClient) {
	router.Handle(prefix+"/spaces/__lookup/{key}", _Universe_LookupSpace_Rule0(cli)).
		Methods("GET").
		Name("eolymp.universe.Universe.LookupSpace")
	router.Handle(prefix+"/spaces", _Universe_CreateSpace_Rule0(cli)).
		Methods("POST").
		Name("eolymp.universe.Universe.CreateSpace")
	router.Handle(prefix+"/spaces/{space_id}", _Universe_UpdateSpace_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.universe.Universe.UpdateSpace")
	router.Handle(prefix+"/spaces/{space_id}", _Universe_DeleteSpace_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.universe.Universe.DeleteSpace")
	router.Handle(prefix+"/spaces/{space_id}", _Universe_DescribeSpace_Rule0(cli)).
		Methods("GET").
		Name("eolymp.universe.Universe.DescribeSpace")
	router.Handle(prefix+"/spaces/{space_id}/quota", _Universe_DescribeQuota_Rule0(cli)).
		Methods("GET").
		Name("eolymp.universe.Universe.DescribeQuota")
	router.Handle(prefix+"/spaces/{space_id}/quota", _Universe_UpdateQuota_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.universe.Universe.UpdateQuota")
	router.Handle(prefix+"/spaces", _Universe_ListSpaces_Rule0(cli)).
		Methods("GET").
		Name("eolymp.universe.Universe.ListSpaces")
}

// RegisterUniverseHttpProxy adds proxy handlers for for UniverseClient
func RegisterUniverseHttpProxy(router *mux.Router, prefix string, conn grpc.ClientConnInterface) {
	RegisterUniverseHttpHandlers(router, prefix, NewUniverseClient(conn))
}

func _Universe_LookupSpace_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LookupSpaceInput{}

		if err := _Universe_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.Key = vars["key"]

		var header, trailer metadata.MD

		out, err := cli.LookupSpace(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_CreateSpace_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateSpaceInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateSpace(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_UpdateSpace_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateSpaceInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SpaceId = vars["space_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateSpace(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_DeleteSpace_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteSpaceInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SpaceId = vars["space_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteSpace(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_DescribeSpace_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeSpaceInput{}

		if err := _Universe_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SpaceId = vars["space_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeSpace(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_DescribeQuota_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeQuotaInput{}

		if err := _Universe_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SpaceId = vars["space_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeQuota(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_UpdateQuota_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateQuotaInput{}

		if err := _Universe_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.SpaceId = vars["space_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateQuota(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Universe_ListSpaces_Rule0(cli UniverseClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListSpacesInput{}

		if err := _Universe_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListSpaces(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Universe_HTTPWriteErrorResponse(w, err)
			return
		}

		_Universe_HTTPWriteResponse(w, out, header, trailer)
	})
}
