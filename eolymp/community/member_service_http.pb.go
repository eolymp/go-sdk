// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package community

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	websocket "golang.org/x/net/websocket"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
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
func _MemberService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_MemberService_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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

// _MemberService_WebsocketErrorResponse writes error to websocket connection
func _MemberService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
	switch status.Convert(e).Code() {
	case codes.OK:
		conn.WriteClose(1000)
	case codes.Canceled:
		conn.WriteClose(1000)
	case codes.Unknown:
		conn.WriteClose(1011)
	case codes.InvalidArgument:
		conn.WriteClose(1003)
	case codes.DeadlineExceeded:
		conn.WriteClose(1000)
	case codes.NotFound:
		conn.WriteClose(1000)
	case codes.AlreadyExists:
		conn.WriteClose(1000)
	case codes.PermissionDenied:
		conn.WriteClose(1000)
	case codes.ResourceExhausted:
		conn.WriteClose(1000)
	case codes.FailedPrecondition:
		conn.WriteClose(1000)
	case codes.Aborted:
		conn.WriteClose(1000)
	case codes.OutOfRange:
		conn.WriteClose(1000)
	case codes.Unimplemented:
		conn.WriteClose(1011)
	case codes.Internal:
		conn.WriteClose(1011)
	case codes.Unavailable:
		conn.WriteClose(1011)
	case codes.DataLoss:
		conn.WriteClose(1011)
	case codes.Unauthenticated:
		conn.WriteClose(1000)
	default:
		conn.WriteClose(1000)
	}
}

// _MemberService_WebsocketCodec implements protobuf codec for websockets package
var _MemberService_WebsocketCodec = websocket.Codec{
	Marshal: func(v interface{}) ([]byte, byte, error) {
		m, ok := v.(proto.Message)
		if !ok {
			panic(fmt.Errorf("invalid message type %T", v))
		}

		d, err := protojson.Marshal(m)
		if err != nil {
			return nil, 0, err
		}

		return d, websocket.TextFrame, err
	},
	Unmarshal: func(d []byte, t byte, v interface{}) error {
		m, ok := v.(proto.Message)
		if !ok {
			panic(fmt.Errorf("invalid message type %T", v))
		}

		return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(d, m)
	},
}

// RegisterMemberServiceHttpHandlers adds handlers for for MemberServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterMemberServiceHttpHandlers(router *mux.Router, prefix string, cli MemberServiceClient) {
	router.Handle(prefix+"/members", _MemberService_CreateMember_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.MemberService.CreateMember")
	router.Handle(prefix+"/members/{member_id}", _MemberService_UpdateMember_Rule0(cli)).
		Methods("POST").
		Name("eolymp.community.MemberService.UpdateMember")
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
	router.Handle(prefix+"/usage/members", _MemberService_DescribeUsage_Rule0(cli)).
		Methods("GET").
		Name("eolymp.community.MemberService.DescribeUsage")
}

func _MemberService_CreateMember_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateMemberInput{}

		if err := _MemberService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.CreateMember(r.Context(), in)
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out)
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

		out, err := cli.UpdateMember(r.Context(), in)
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out)
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

		out, err := cli.DeleteMember(r.Context(), in)
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out)
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

		out, err := cli.RestoreMember(r.Context(), in)
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out)
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

		out, err := cli.DescribeMember(r.Context(), in)
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out)
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

		out, err := cli.ListMembers(r.Context(), in)
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out)
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

		out, err := cli.AssignMember(r.Context(), in)
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out)
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

		out, err := cli.UnassignMember(r.Context(), in)
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out)
	})
}

func _MemberService_DescribeUsage_Rule0(cli MemberServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeUsageInput{}

		if err := _MemberService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := cli.DescribeUsage(r.Context(), in)
		if err != nil {
			_MemberService_HTTPWriteErrorResponse(w, err)
			return
		}

		_MemberService_HTTPWriteResponse(w, out)
	})
}

type _MemberServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _MemberServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _MemberServiceHandler) (out proto.Message, err error)
type MemberServiceInterceptor struct {
	middleware []_MemberServiceMiddleware
	client     MemberServiceClient
}

// NewMemberServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewMemberServiceInterceptor(cli MemberServiceClient, middleware ..._MemberServiceMiddleware) *MemberServiceInterceptor {
	return &MemberServiceInterceptor{client: cli, middleware: middleware}
}

func (i *MemberServiceInterceptor) CreateMember(ctx context.Context, in *CreateMemberInput, opts ...grpc.CallOption) (*CreateMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateMemberInput, got %T", in))
		}

		return i.client.CreateMember(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.MemberService.CreateMember", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateMemberOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateMemberOutput, got %T", out))
	}

	return message, err
}

func (i *MemberServiceInterceptor) UpdateMember(ctx context.Context, in *UpdateMemberInput, opts ...grpc.CallOption) (*UpdateMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateMemberInput, got %T", in))
		}

		return i.client.UpdateMember(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.MemberService.UpdateMember", in, next)
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

func (i *MemberServiceInterceptor) DeleteMember(ctx context.Context, in *DeleteMemberInput, opts ...grpc.CallOption) (*DeleteMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteMemberInput, got %T", in))
		}

		return i.client.DeleteMember(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.MemberService.DeleteMember", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteMemberOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteMemberOutput, got %T", out))
	}

	return message, err
}

func (i *MemberServiceInterceptor) RestoreMember(ctx context.Context, in *RestoreMemberInput, opts ...grpc.CallOption) (*RestoreMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RestoreMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RestoreMemberInput, got %T", in))
		}

		return i.client.RestoreMember(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.MemberService.RestoreMember", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RestoreMemberOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RestoreMemberOutput, got %T", out))
	}

	return message, err
}

func (i *MemberServiceInterceptor) DescribeMember(ctx context.Context, in *DescribeMemberInput, opts ...grpc.CallOption) (*DescribeMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeMemberInput, got %T", in))
		}

		return i.client.DescribeMember(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.MemberService.DescribeMember", in, next)
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

func (i *MemberServiceInterceptor) ListMembers(ctx context.Context, in *ListMembersInput, opts ...grpc.CallOption) (*ListMembersOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListMembersInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListMembersInput, got %T", in))
		}

		return i.client.ListMembers(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.MemberService.ListMembers", in, next)
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

func (i *MemberServiceInterceptor) AssignMember(ctx context.Context, in *AssignMemberInput, opts ...grpc.CallOption) (*AssignMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*AssignMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *AssignMemberInput, got %T", in))
		}

		return i.client.AssignMember(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.MemberService.AssignMember", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*AssignMemberOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *AssignMemberOutput, got %T", out))
	}

	return message, err
}

func (i *MemberServiceInterceptor) UnassignMember(ctx context.Context, in *UnassignMemberInput, opts ...grpc.CallOption) (*UnassignMemberOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UnassignMemberInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UnassignMemberInput, got %T", in))
		}

		return i.client.UnassignMember(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.MemberService.UnassignMember", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UnassignMemberOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UnassignMemberOutput, got %T", out))
	}

	return message, err
}

func (i *MemberServiceInterceptor) DescribeUsage(ctx context.Context, in *DescribeUsageInput, opts ...grpc.CallOption) (*DescribeUsageOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeUsageInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeUsageInput, got %T", in))
		}

		return i.client.DescribeUsage(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.community.MemberService.DescribeUsage", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeUsageOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeUsageOutput, got %T", out))
	}

	return message, err
}
