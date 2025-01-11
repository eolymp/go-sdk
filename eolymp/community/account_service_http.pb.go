// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package community

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

// _AccountService_HTTPReadQueryString parses body into proto.Message
func _AccountService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _AccountService_HTTPReadRequestBody parses body into proto.Message
func _AccountService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _AccountService_HTTPWriteResponse writes proto.Message to HTTP response
func _AccountService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_AccountService_HTTPWriteErrorResponse(w, err)
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

// _AccountService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _AccountService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterAccountServiceHttpHandlers adds handlers for for AccountServiceClient
func RegisterAccountServiceHttpHandlers(router *mux.Router, prefix string, cli AccountServiceClient) {
	router.Handle(prefix+"/account", _AccountService_CreateAccount_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.AccountService.CreateAccount")
	router.Handle(prefix+"/account", _AccountService_DescribeAccount_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.AccountService.DescribeAccount")
	router.Handle(prefix+"/account", _AccountService_UpdateAccount_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.community.AccountService.UpdateAccount")
	router.Handle(prefix+"/account/picture", _AccountService_UploadPicture_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.AccountService.UploadPicture")
	router.Handle(prefix+"/account", _AccountService_DeleteAccount_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.community.AccountService.DeleteAccount")
	router.Handle(prefix+"/account/verification/resend", _AccountService_ResendVerification_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.AccountService.ResendVerification")
	router.Handle(prefix+"/account/verification/complete", _AccountService_CompleteVerification_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.AccountService.CompleteVerification")
	router.Handle(prefix+"/account/recovery/start", _AccountService_StartRecovery_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.AccountService.StartRecovery")
	router.Handle(prefix+"/account/recovery/complete", _AccountService_CompleteRecovery_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.AccountService.CompleteRecovery")
}

// RegisterAccountServiceHttpProxy adds proxy handlers for for AccountServiceClient
func RegisterAccountServiceHttpProxy(router *mux.Router, prefix string, conn grpc.ClientConnInterface) {
	RegisterAccountServiceHttpHandlers(router, prefix, NewAccountServiceClient(conn))
}

func _AccountService_CreateAccount_Rule0(cli AccountServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateAccountInput{}

		if err := _AccountService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateAccount(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AccountService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AccountService_DescribeAccount_Rule0(cli AccountServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeAccountInput{}

		if err := _AccountService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeAccount(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AccountService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AccountService_UpdateAccount_Rule0(cli AccountServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateAccountInput{}

		if err := _AccountService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.UpdateAccount(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AccountService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AccountService_UploadPicture_Rule0(cli AccountServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UploadPictureInput{}

		if err := _AccountService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.UploadPicture(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AccountService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AccountService_DeleteAccount_Rule0(cli AccountServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteAccountInput{}

		if err := _AccountService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DeleteAccount(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AccountService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AccountService_ResendVerification_Rule0(cli AccountServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ResendVerificationInput{}

		if err := _AccountService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ResendVerification(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AccountService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AccountService_CompleteVerification_Rule0(cli AccountServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CompleteVerificationInput{}

		if err := _AccountService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CompleteVerification(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AccountService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AccountService_StartRecovery_Rule0(cli AccountServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &StartRecoveryInput{}

		if err := _AccountService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.StartRecovery(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AccountService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _AccountService_CompleteRecovery_Rule0(cli AccountServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CompleteRecoverInput{}

		if err := _AccountService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CompleteRecovery(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_AccountService_HTTPWriteErrorResponse(w, err)
			return
		}

		_AccountService_HTTPWriteResponse(w, out, header, trailer)
	})
}
