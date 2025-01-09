// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package newsletter

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

// _RecipientService_HTTPReadQueryString parses body into proto.Message
func _RecipientService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _RecipientService_HTTPReadRequestBody parses body into proto.Message
func _RecipientService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _RecipientService_HTTPWriteResponse writes proto.Message to HTTP response
func _RecipientService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_RecipientService_HTTPWriteErrorResponse(w, err)
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

// _RecipientService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _RecipientService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterRecipientServiceHttpHandlers adds handlers for for RecipientServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterRecipientServiceHttpHandlers(router *mux.Router, prefix string, cli RecipientServiceClient) {
	router.Handle(prefix+"/recipients/{recipient_id}", _RecipientService_DescribeRecipient_Rule0(cli)).
		Methods("GET").
		Name("eolymp.newsletter.RecipientService.DescribeRecipient")
	router.Handle(prefix+"/imports/recipients", _RecipientService_ImportRecipients_Rule0(cli)).
		Methods("POST").
		Name("eolymp.newsletter.RecipientService.ImportRecipients")
	router.Handle(prefix+"/recipients", _RecipientService_ListRecipients_Rule0(cli)).
		Methods("GET").
		Name("eolymp.newsletter.RecipientService.ListRecipients")
	router.Handle(prefix+"/recipients", _RecipientService_CreateRecipient_Rule0(cli)).
		Methods("POST").
		Name("eolymp.newsletter.RecipientService.CreateRecipient")
	router.Handle(prefix+"/recipients/{recipient_id}", _RecipientService_RemoveRecipient_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.newsletter.RecipientService.RemoveRecipient")
}

func _RecipientService_DescribeRecipient_Rule0(cli RecipientServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRecipientInput{}

		if err := _RecipientService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RecipientService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RecipientId = vars["recipient_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeRecipient(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RecipientService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RecipientService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RecipientService_ImportRecipients_Rule0(cli RecipientServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ImportRecipientsInput{}

		if err := _RecipientService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RecipientService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ImportRecipients(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RecipientService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RecipientService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RecipientService_ListRecipients_Rule0(cli RecipientServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListRecipientsInput{}

		if err := _RecipientService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RecipientService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListRecipients(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RecipientService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RecipientService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RecipientService_CreateRecipient_Rule0(cli RecipientServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateRecipientInput{}

		if err := _RecipientService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RecipientService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateRecipient(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RecipientService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RecipientService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _RecipientService_RemoveRecipient_Rule0(cli RecipientServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RemoveRecipientInput{}

		if err := _RecipientService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_RecipientService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RecipientId = vars["recipient_id"]

		var header, trailer metadata.MD

		out, err := cli.RemoveRecipient(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_RecipientService_HTTPWriteErrorResponse(w, err)
			return
		}

		_RecipientService_HTTPWriteResponse(w, out, header, trailer)
	})
}
