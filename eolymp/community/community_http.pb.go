// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package community

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _Community_HTTPReadQueryString parses body into proto.Message
func _Community_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Community_HTTPReadRequestBody parses body into proto.Message
func _Community_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Community_HTTPWriteResponse writes proto.Message to HTTP response
func _Community_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Community_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Community_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Community_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterCommunityHttpHandlers adds handlers for for CommunityServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterCommunityHttpHandlers(router *mux.Router, prefix string, srv CommunityServer) {
	router.Handle(prefix+"/members/_self", _Community_JoinSpace_Rule0(srv)).
		Methods("POST").
		Name("eolymp.community.Community.JoinSpace")
	router.Handle(prefix+"/members/_self", _Community_LeaveSpace_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.community.Community.LeaveSpace")
	router.Handle(prefix+"/members/_self/attributes", _Community_RegisterMember_Rule0(srv)).
		Methods("POST").
		Name("eolymp.community.Community.RegisterMember")
	router.Handle(prefix+"/members/_self", _Community_IntrospectMember_Rule0(srv)).
		Methods("GET").
		Name("eolymp.community.Community.IntrospectMember")
	router.Handle(prefix+"/members", _Community_AddMember_Rule0(srv)).
		Methods("POST").
		Name("eolymp.community.Community.AddMember")
	router.Handle(prefix+"/members/{member_id}", _Community_UpdateMember_Rule0(srv)).
		Methods("POST").
		Name("eolymp.community.Community.UpdateMember")
	router.Handle(prefix+"/members/{member_id}", _Community_RemoveMember_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.community.Community.RemoveMember")
	router.Handle(prefix+"/members/{member_id}", _Community_DescribeMember_Rule0(srv)).
		Methods("GET").
		Name("eolymp.community.Community.DescribeMember")
	router.Handle(prefix+"/members", _Community_ListMembers_Rule0(srv)).
		Methods("GET").
		Name("eolymp.community.Community.ListMembers")
	router.Handle(prefix+"/members/{member_id}/identities", _Community_AddMemberIdentity_Rule0(srv)).
		Methods("POST").
		Name("eolymp.community.Community.AddMemberIdentity")
	router.Handle(prefix+"/members/{member_id}/identities/{identity_id}", _Community_UpdateMemberIdentity_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.community.Community.UpdateMemberIdentity")
	router.Handle(prefix+"/members/{member_id}/identities/{identity_id}", _Community_RemoveMemberIdentity_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.community.Community.RemoveMemberIdentity")
	router.Handle(prefix+"/attributes", _Community_AddAttribute_Rule0(srv)).
		Methods("POST").
		Name("eolymp.community.Community.AddAttribute")
	router.Handle(prefix+"/attributes/{attribute_key}", _Community_UpdateAttribute_Rule0(srv)).
		Methods("POST").
		Name("eolymp.community.Community.UpdateAttribute")
	router.Handle(prefix+"/attributes/{attribute_key}", _Community_RemoveAttribute_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.community.Community.RemoveAttribute")
	router.Handle(prefix+"/attributes/{attribute_key}", _Community_DescribeAttribute_Rule0(srv)).
		Methods("GET").
		Name("eolymp.community.Community.DescribeAttribute")
	router.Handle(prefix+"/attributes", _Community_ListAttributes_Rule0(srv)).
		Methods("GET").
		Name("eolymp.community.Community.ListAttributes")
	router.Handle(prefix+"/idp", _Community_DescribeIdentityProvider_Rule0(srv)).
		Methods("GET").
		Name("eolymp.community.Community.DescribeIdentityProvider")
	router.Handle(prefix+"/idp", _Community_ConfigureIdentityProvider_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.community.Community.ConfigureIdentityProvider")
}

func _Community_JoinSpace_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &JoinSpaceInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.JoinSpace(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_LeaveSpace_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LeaveSpaceInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.LeaveSpace(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_RegisterMember_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RegisterMemberInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.RegisterMember(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_IntrospectMember_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &IntrospectMemberInput{}

		if err := _Community_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.IntrospectMember(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_AddMember_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &AddMemberInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.AddMember(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_UpdateMember_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateMemberInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		out, err := srv.UpdateMember(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_RemoveMember_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RemoveMemberInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		out, err := srv.RemoveMember(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_DescribeMember_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeMemberInput{}

		if err := _Community_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		out, err := srv.DescribeMember(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_ListMembers_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListMembersInput{}

		if err := _Community_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListMembers(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_AddMemberIdentity_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &AddMemberIdentityInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]

		out, err := srv.AddMemberIdentity(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_UpdateMemberIdentity_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateMemberIdentityInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]
		in.IdentityId = vars["identity_id"]

		out, err := srv.UpdateMemberIdentity(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_RemoveMemberIdentity_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RemoveMemberIdentityInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.MemberId = vars["member_id"]
		in.IdentityId = vars["identity_id"]

		out, err := srv.RemoveMemberIdentity(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_AddAttribute_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &AddAttributeInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.AddAttribute(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_UpdateAttribute_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateAttributeInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AttributeKey = vars["attribute_key"]

		out, err := srv.UpdateAttribute(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_RemoveAttribute_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RemoveAttributeInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AttributeKey = vars["attribute_key"]

		out, err := srv.RemoveAttribute(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_DescribeAttribute_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeAttributeInput{}

		if err := _Community_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.AttributeKey = vars["attribute_key"]

		out, err := srv.DescribeAttribute(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_ListAttributes_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListAttributesInput{}

		if err := _Community_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListAttributes(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_DescribeIdentityProvider_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeIdentityProviderInput{}

		if err := _Community_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeIdentityProvider(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

func _Community_ConfigureIdentityProvider_Rule0(srv CommunityServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ConfigureIdentityProviderInput{}

		if err := _Community_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ConfigureIdentityProvider(r.Context(), in)
		if err != nil {
			_Community_HTTPWriteErrorResponse(w, err)
			return
		}

		_Community_HTTPWriteResponse(w, out)
	})
}

type _CommunityHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _CommunityMiddleware = func(ctx context.Context, method string, in proto.Message, handler _CommunityHandler) (out proto.Message, err error)
type CommunityInterceptor struct {
	middleware []_CommunityMiddleware
	server     CommunityServer
}

// NewCommunityInterceptor constructs additional middleware for a server based on annotations in proto files
func NewCommunityInterceptor(srv CommunityServer, middleware ..._CommunityMiddleware) *CommunityInterceptor {
	return &CommunityInterceptor{server: srv, middleware: middleware}
}

func (i *CommunityInterceptor) JoinSpace(ctx context.Context, in *JoinSpaceInput) (*JoinSpaceOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*JoinSpaceInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *JoinSpaceInput, got %T", in))
		}

		return i.server.JoinSpace(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.JoinSpace", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*JoinSpaceOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *JoinSpaceOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) LeaveSpace(ctx context.Context, in *LeaveSpaceInput) (*LeaveSpaceOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*LeaveSpaceInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *LeaveSpaceInput, got %T", in))
		}

		return i.server.LeaveSpace(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.LeaveSpace", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*LeaveSpaceOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *LeaveSpaceOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) RegisterMember(ctx context.Context, in *RegisterMemberInput) (*RegisterMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RegisterMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RegisterMemberInput, got %T", in))
		}

		return i.server.RegisterMember(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.RegisterMember", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RegisterMemberOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RegisterMemberOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) IntrospectMember(ctx context.Context, in *IntrospectMemberInput) (*IntrospectMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*IntrospectMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *IntrospectMemberInput, got %T", in))
		}

		return i.server.IntrospectMember(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.IntrospectMember", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*IntrospectMemberOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *IntrospectMemberOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) AddMember(ctx context.Context, in *AddMemberInput) (*AddMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*AddMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *AddMemberInput, got %T", in))
		}

		return i.server.AddMember(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.AddMember", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*AddMemberOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *AddMemberOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) UpdateMember(ctx context.Context, in *UpdateMemberInput) (*UpdateMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateMemberInput, got %T", in))
		}

		return i.server.UpdateMember(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.UpdateMember", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateMemberOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateMemberOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) RemoveMember(ctx context.Context, in *RemoveMemberInput) (*RemoveMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RemoveMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RemoveMemberInput, got %T", in))
		}

		return i.server.RemoveMember(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.RemoveMember", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RemoveMemberOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RemoveMemberOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) DescribeMember(ctx context.Context, in *DescribeMemberInput) (*DescribeMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeMemberInput, got %T", in))
		}

		return i.server.DescribeMember(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.DescribeMember", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeMemberOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeMemberOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) ListMembers(ctx context.Context, in *ListMembersInput) (*ListMembersOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListMembersInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListMembersInput, got %T", in))
		}

		return i.server.ListMembers(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.ListMembers", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListMembersOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListMembersOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) AddMemberIdentity(ctx context.Context, in *AddMemberIdentityInput) (*AddMemberIdentityOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*AddMemberIdentityInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *AddMemberIdentityInput, got %T", in))
		}

		return i.server.AddMemberIdentity(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.AddMemberIdentity", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*AddMemberIdentityOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *AddMemberIdentityOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) UpdateMemberIdentity(ctx context.Context, in *UpdateMemberIdentityInput) (*UpdateMemberIdentityOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateMemberIdentityInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateMemberIdentityInput, got %T", in))
		}

		return i.server.UpdateMemberIdentity(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.UpdateMemberIdentity", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateMemberIdentityOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateMemberIdentityOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) RemoveMemberIdentity(ctx context.Context, in *RemoveMemberIdentityInput) (*RemoveMemberIdentityOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RemoveMemberIdentityInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RemoveMemberIdentityInput, got %T", in))
		}

		return i.server.RemoveMemberIdentity(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.RemoveMemberIdentity", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RemoveMemberIdentityOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RemoveMemberIdentityOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) AddAttribute(ctx context.Context, in *AddAttributeInput) (*AddAttributeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*AddAttributeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *AddAttributeInput, got %T", in))
		}

		return i.server.AddAttribute(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.AddAttribute", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*AddAttributeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *AddAttributeOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) UpdateAttribute(ctx context.Context, in *UpdateAttributeInput) (*UpdateAttributeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateAttributeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateAttributeInput, got %T", in))
		}

		return i.server.UpdateAttribute(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.UpdateAttribute", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateAttributeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateAttributeOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) RemoveAttribute(ctx context.Context, in *RemoveAttributeInput) (*RemoveAttributeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RemoveAttributeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RemoveAttributeInput, got %T", in))
		}

		return i.server.RemoveAttribute(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.RemoveAttribute", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RemoveAttributeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RemoveAttributeOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) DescribeAttribute(ctx context.Context, in *DescribeAttributeInput) (*DescribeAttributeOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeAttributeInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeAttributeInput, got %T", in))
		}

		return i.server.DescribeAttribute(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.DescribeAttribute", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeAttributeOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeAttributeOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) ListAttributes(ctx context.Context, in *ListAttributesInput) (*ListAttributesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListAttributesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListAttributesInput, got %T", in))
		}

		return i.server.ListAttributes(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.ListAttributes", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListAttributesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListAttributesOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) DescribeIdentityProvider(ctx context.Context, in *DescribeIdentityProviderInput) (*DescribeIdentityProviderOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeIdentityProviderInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeIdentityProviderInput, got %T", in))
		}

		return i.server.DescribeIdentityProvider(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.DescribeIdentityProvider", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeIdentityProviderOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeIdentityProviderOutput, got %T", out))
	}

	return message, err
}

func (i *CommunityInterceptor) ConfigureIdentityProvider(ctx context.Context, in *ConfigureIdentityProviderInput) (*ConfigureIdentityProviderOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ConfigureIdentityProviderInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ConfigureIdentityProviderInput, got %T", in))
		}

		return i.server.ConfigureIdentityProvider(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.Community.ConfigureIdentityProvider", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ConfigureIdentityProviderOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ConfigureIdentityProviderOutput, got %T", out))
	}

	return message, err
}
