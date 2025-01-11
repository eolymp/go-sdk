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

// _MaterialService_HTTPReadQueryString parses body into proto.Message
func _MaterialService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _MaterialService_HTTPReadRequestBody parses body into proto.Message
func _MaterialService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _MaterialService_HTTPWriteResponse writes proto.Message to HTTP response
func _MaterialService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_MaterialService_HTTPWriteErrorResponse(w, err)
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

// _MaterialService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _MaterialService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterMaterialServiceHttpHandlers adds handlers for for MaterialServiceClient
func RegisterMaterialServiceHttpHandlers(router *mux.Router, prefix string, cli MaterialServiceClient) {
	router.Handle(prefix+"/materials", _MaterialService_CreateMaterial_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.MaterialService.CreateMaterial")
	router.Handle(prefix+"/materials/{material_id}", _MaterialService_UpdateMaterial_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.course.MaterialService.UpdateMaterial")
	router.Handle(prefix+"/materials/{material_id}/move", _MaterialService_MoveMaterial_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.MaterialService.MoveMaterial")
	router.Handle(prefix+"/materials/{material_id}", _MaterialService_DeleteMaterial_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.course.MaterialService.DeleteMaterial")
	router.Handle(prefix+"/materials/{material_id}", _MaterialService_DescribeMaterial_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.MaterialService.DescribeMaterial")
	router.Handle(prefix+"/materials", _MaterialService_ListMaterials_Rule0(cli)).
		Methods("GET").
		Name("eolymp.course.MaterialService.ListMaterials")
	router.Handle(prefix+"/materials/{material_id}/progress", _MaterialService_ReportProgress_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.MaterialService.ReportProgress")
	router.Handle(prefix+"/materials/{material_id}/grade", _MaterialService_GradeMaterial_Rule0(cli)).
		Methods("POST").
		Name("eolymp.course.MaterialService.GradeMaterial")
}

// RegisterMaterialServiceHttpProxy adds proxy handlers for for MaterialServiceClient
func RegisterMaterialServiceHttpProxy(router *mux.Router, prefix string, conn grpc.ClientConnInterface) {
	RegisterMaterialServiceHttpHandlers(router, prefix, NewMaterialServiceClient(conn))
}

func _MaterialService_CreateMaterial_Rule0(cli MaterialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateMaterialInput{}

		if err := _MaterialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateMaterial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MaterialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MaterialService_UpdateMaterial_Rule0(cli MaterialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateMaterialInput{}

		if err := _MaterialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MaterialId = vars["material_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateMaterial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MaterialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MaterialService_MoveMaterial_Rule0(cli MaterialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &MoveMaterialInput{}

		if err := _MaterialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MaterialId = vars["material_id"]

		var header, trailer metadata.MD

		out, err := cli.MoveMaterial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MaterialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MaterialService_DeleteMaterial_Rule0(cli MaterialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteMaterialInput{}

		if err := _MaterialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MaterialId = vars["material_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteMaterial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MaterialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MaterialService_DescribeMaterial_Rule0(cli MaterialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeMaterialInput{}

		if err := _MaterialService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MaterialId = vars["material_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeMaterial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MaterialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MaterialService_ListMaterials_Rule0(cli MaterialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListMaterialsInput{}

		if err := _MaterialService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListMaterials(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MaterialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MaterialService_ReportProgress_Rule0(cli MaterialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ReportProgressInput{}

		if err := _MaterialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MaterialId = vars["material_id"]

		var header, trailer metadata.MD

		out, err := cli.ReportProgress(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MaterialService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MaterialService_GradeMaterial_Rule0(cli MaterialServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &GradeMaterialInput{}

		if err := _MaterialService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MaterialId = vars["material_id"]

		var header, trailer metadata.MD

		out, err := cli.GradeMaterial(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MaterialService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MaterialService_HTTPWriteResponse(w, out, header, trailer)
	})
}
