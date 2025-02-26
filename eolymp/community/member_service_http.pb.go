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

// _MemberService_HTTPReadQueryString parses body into proto.Message
func _MemberService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _MemberService_HTTPReadRequestBody parses body into proto.Message
func _MemberService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _MemberService_HTTPWriteResponse writes proto.Message to HTTP response
func _MemberService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_MemberService_HTTPWriteErrorResponse(w, err)
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

// _MemberService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _MemberService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterMemberServiceHttpHandlers adds handlers for for MemberServiceClient
func RegisterMemberServiceHttpHandlers(router *mux.Router, prefix string, cli MemberServiceClient) {
	router.Handle(prefix+"/members", _MemberService_CreateMember_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.MemberService.CreateMember")
	router.Handle(prefix+"/members/{member_id}", _MemberService_UpdateMember_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.MemberService.UpdateMember")
	router.Handle(prefix+"/members/{member_id}/picture", _MemberService_UpdateMemberPicture_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.MemberService.UpdateMemberPicture")
	router.Handle(prefix+"/members/{member_id}", _MemberService_DeleteMember_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.community.MemberService.DeleteMember")
	router.Handle(prefix+"/members/{member_id}/restore", _MemberService_RestoreMember_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.MemberService.RestoreMember")
	router.Handle(prefix+"/members/{member_id}", _MemberService_DescribeMember_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.MemberService.DescribeMember")
	router.Handle(prefix+"/members", _MemberService_ListMembers_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.MemberService.ListMembers")
	router.Handle(prefix+"/members/{team_id}/users/{member_id}", _MemberService_AssignMember_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.community.MemberService.AssignMember")
	router.Handle(prefix+"/members/{team_id}/users/{member_id}", _MemberService_UnassignMember_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.community.MemberService.UnassignMember")
	router.Handle(prefix+"/members/{member_id}/notify", _MemberService_NotifyMember_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.MemberService.NotifyMember")
	router.Handle(prefix+"/usage/members", _MemberService_DescribeMemberUsage_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.MemberService.DescribeMemberUsage")
}

// RegisterMemberServiceHttpProxy adds proxy handlers for for MemberServiceClient
func RegisterMemberServiceHttpProxy(router *mux.Router, prefix string, conn grpc.ClientConnInterface) {
	RegisterMemberServiceHttpHandlers(router, prefix, NewMemberServiceClient(conn))
}

func _MemberService_CreateMember_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateMemberInput{}

		if err := _MemberService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateMember(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MemberService_UpdateMember_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateMemberInput{}

		if err := _MemberService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateMember(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MemberService_UpdateMemberPicture_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateMemberPictureInput{}

		if err := _MemberService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateMemberPicture(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MemberService_DeleteMember_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteMemberInput{}

		if err := _MemberService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteMember(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MemberService_RestoreMember_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RestoreMemberInput{}

		if err := _MemberService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.RestoreMember(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MemberService_DescribeMember_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeMemberInput{}

		if err := _MemberService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeMember(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MemberService_ListMembers_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListMembersInput{}

		if err := _MemberService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListMembers(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MemberService_AssignMember_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &AssignMemberInput{}

		if err := _MemberService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TeamId = vars["team_id"]
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.AssignMember(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MemberService_UnassignMember_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UnassignMemberInput{}

		if err := _MemberService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TeamId = vars["team_id"]
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.UnassignMember(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MemberService_NotifyMember_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &NotifyMemberInput{}

		if err := _MemberService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.NotifyMember(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _MemberService_DescribeMemberUsage_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeMemberUsageInput{}

		if err := _MemberService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.DescribeMemberUsage(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out, header, trailer)
	})
}
