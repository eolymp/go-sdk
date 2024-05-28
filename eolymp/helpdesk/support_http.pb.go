// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package helpdesk

import (
	context "context"
	fmt "fmt"
	mux "github.com/gorilla/mux"
	websocket "golang.org/x/net/websocket"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _Support_HTTPReadQueryString parses body into proto.Message
func _Support_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Support_HTTPReadRequestBody parses body into proto.Message
func _Support_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Support_HTTPWriteResponse writes proto.Message to HTTP response
func _Support_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Support_HTTPWriteErrorResponse(w, err)
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

// _Support_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Support_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _Support_WebsocketErrorResponse writes error to websocket connection
func _Support_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _Support_WebsocketCodec implements protobuf codec for websockets package
var _Support_WebsocketCodec = websocket.Codec{
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

// RegisterSupportHttpHandlers adds handlers for for SupportClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterSupportHttpHandlers(router *mux.Router, prefix string, cli SupportClient) {
	router.Handle(prefix+"/helpdesk/tickets", _Support_CreateTicket_Rule0(cli)).
		Methods("POST").
		Name("eolymp.helpdesk.Support.CreateTicket")
	router.Handle(prefix+"/helpdesk/tickets/{ticket_id}", _Support_UpdateTicket_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.helpdesk.Support.UpdateTicket")
	router.Handle(prefix+"/helpdesk/tickets/{ticket_id}", _Support_DeleteTicket_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.helpdesk.Support.DeleteTicket")
	router.Handle(prefix+"/helpdesk/tickets/{ticket_id}", _Support_DescribeTicket_Rule0(cli)).
		Methods("GET").
		Name("eolymp.helpdesk.Support.DescribeTicket")
	router.Handle(prefix+"/helpdesk/tickets", _Support_ListTickets_Rule0(cli)).
		Methods("GET").
		Name("eolymp.helpdesk.Support.ListTickets")
	router.Handle(prefix+"/helpdesk/tickets/{ticket_id}/approve", _Support_ApproveTicket_Rule0(cli)).
		Methods("POST").
		Name("eolymp.helpdesk.Support.ApproveTicket")
	router.Handle(prefix+"/helpdesk/tickets/{ticket_id}/reject", _Support_RejectTicket_Rule0(cli)).
		Methods("POST").
		Name("eolymp.helpdesk.Support.RejectTicket")
	router.Handle(prefix+"/helpdesk/tickets/{ticket_id}/close", _Support_CloseTicket_Rule0(cli)).
		Methods("POST").
		Name("eolymp.helpdesk.Support.CloseTicket")
	router.Handle(prefix+"/helpdesk/tickets/{ticket_id}/comments", _Support_AddComment_Rule0(cli)).
		Methods("POST").
		Name("eolymp.helpdesk.Support.AddComment")
	router.Handle(prefix+"/helpdesk/tickets/{ticket_id}/comments/{comment_id}", _Support_UpdateComment_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.helpdesk.Support.UpdateComment")
	router.Handle(prefix+"/helpdesk/tickets/{ticket_id}/comments/{comment_id}", _Support_DeleteComment_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.helpdesk.Support.DeleteComment")
	router.Handle(prefix+"/helpdesk/tickets/{ticket_id}/comments", _Support_ListComments_Rule0(cli)).
		Methods("GET").
		Name("eolymp.helpdesk.Support.ListComments")
	router.Handle(prefix+"/helpdesk/tickets/{ticket_id}/comments/{comment_id}", _Support_DescribeComment_Rule0(cli)).
		Methods("GET").
		Name("eolymp.helpdesk.Support.DescribeComment")
	router.Handle(prefix+"/helpdesk/autoreplies", _Support_CreateAutoReply_Rule0(cli)).
		Methods("POST").
		Name("eolymp.helpdesk.Support.CreateAutoReply")
	router.Handle(prefix+"/helpdesk/autoreplies/{reply_id}", _Support_UpdateAutoReply_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.helpdesk.Support.UpdateAutoReply")
	router.Handle(prefix+"/helpdesk/autoreplies/{reply_id}", _Support_DeleteAutoReply_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.helpdesk.Support.DeleteAutoReply")
	router.Handle(prefix+"/helpdesk/autoreplies", _Support_ListAutoReplies_Rule0(cli)).
		Methods("GET").
		Name("eolymp.helpdesk.Support.ListAutoReplies")
	router.Handle(prefix+"/helpdesk/autoreplies/{reply_id}", _Support_DescribeAutoReply_Rule0(cli)).
		Methods("GET").
		Name("eolymp.helpdesk.Support.DescribeAutoReply")
	router.Handle(prefix+"/helpdesk/attachments", _Support_UploadAttachment_Rule0(cli)).
		Methods("POST").
		Name("eolymp.helpdesk.Support.UploadAttachment")
}

func _Support_CreateTicket_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateTicketInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateTicket(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_UpdateTicket_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateTicketInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TicketId = vars["ticket_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateTicket(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_DeleteTicket_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteTicketInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TicketId = vars["ticket_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteTicket(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_DescribeTicket_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeTicketInput{}

		if err := _Support_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TicketId = vars["ticket_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeTicket(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_ListTickets_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListTicketsInput{}

		if err := _Support_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListTickets(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_ApproveTicket_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ApproveTicketInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TicketId = vars["ticket_id"]

		var header, trailer metadata.MD

		out, err := cli.ApproveTicket(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_RejectTicket_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RejectTicketInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TicketId = vars["ticket_id"]

		var header, trailer metadata.MD

		out, err := cli.RejectTicket(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_CloseTicket_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CloseTicketInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TicketId = vars["ticket_id"]

		var header, trailer metadata.MD

		out, err := cli.CloseTicket(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_AddComment_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &AddCommentInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TicketId = vars["ticket_id"]

		var header, trailer metadata.MD

		out, err := cli.AddComment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_UpdateComment_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateCommentInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TicketId = vars["ticket_id"]
		in.CommentId = vars["comment_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateComment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_DeleteComment_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteCommentInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TicketId = vars["ticket_id"]
		in.CommentId = vars["comment_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteComment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_ListComments_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListCommentsInput{}

		if err := _Support_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TicketId = vars["ticket_id"]

		var header, trailer metadata.MD

		out, err := cli.ListComments(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_DescribeComment_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCommentInput{}

		if err := _Support_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.TicketId = vars["ticket_id"]
		in.CommentId = vars["comment_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeComment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_CreateAutoReply_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateAutoReplyInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateAutoReply(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_UpdateAutoReply_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateAutoReplyInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ReplyId = vars["reply_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateAutoReply(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_DeleteAutoReply_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteAutoReplyInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ReplyId = vars["reply_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteAutoReply(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_ListAutoReplies_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListAutoRepliesInput{}

		if err := _Support_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListAutoReplies(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_DescribeAutoReply_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeAutoReplyInput{}

		if err := _Support_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ReplyId = vars["reply_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeAutoReply(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _Support_UploadAttachment_Rule0(cli SupportClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UploadAttachmentInput{}

		if err := _Support_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.UploadAttachment(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_Support_HTTPWriteErrorResponse(w, err)
			return
		}

		_Support_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _SupportHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _SupportMiddleware = func(ctx context.Context, method string, in proto.Message, handler _SupportHandler) (out proto.Message, err error)
type SupportInterceptor struct {
	middleware []_SupportMiddleware
	client     SupportClient
}

// NewSupportInterceptor constructs additional middleware for a server based on annotations in proto files
func NewSupportInterceptor(cli SupportClient, middleware ..._SupportMiddleware) *SupportInterceptor {
	return &SupportInterceptor{client: cli, middleware: middleware}
}

func (i *SupportInterceptor) CreateTicket(ctx context.Context, in *CreateTicketInput, opts ...grpc.CallOption) (*CreateTicketOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateTicketInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateTicketInput, got %T", in))
		}

		return i.client.CreateTicket(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.CreateTicket", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateTicketOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateTicketOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) UpdateTicket(ctx context.Context, in *UpdateTicketInput, opts ...grpc.CallOption) (*UpdateTicketOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateTicketInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateTicketInput, got %T", in))
		}

		return i.client.UpdateTicket(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.UpdateTicket", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateTicketOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateTicketOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) DeleteTicket(ctx context.Context, in *DeleteTicketInput, opts ...grpc.CallOption) (*DeleteTicketOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteTicketInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteTicketInput, got %T", in))
		}

		return i.client.DeleteTicket(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.DeleteTicket", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteTicketOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteTicketOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) DescribeTicket(ctx context.Context, in *DescribeTicketInput, opts ...grpc.CallOption) (*DescribeTicketOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeTicketInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeTicketInput, got %T", in))
		}

		return i.client.DescribeTicket(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.DescribeTicket", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeTicketOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeTicketOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) ListTickets(ctx context.Context, in *ListTicketsInput, opts ...grpc.CallOption) (*ListTicketsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListTicketsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListTicketsInput, got %T", in))
		}

		return i.client.ListTickets(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.ListTickets", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListTicketsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListTicketsOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) ApproveTicket(ctx context.Context, in *ApproveTicketInput, opts ...grpc.CallOption) (*ApproveTicketOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ApproveTicketInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ApproveTicketInput, got %T", in))
		}

		return i.client.ApproveTicket(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.ApproveTicket", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ApproveTicketOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ApproveTicketOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) RejectTicket(ctx context.Context, in *RejectTicketInput, opts ...grpc.CallOption) (*RejectTicketOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RejectTicketInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RejectTicketInput, got %T", in))
		}

		return i.client.RejectTicket(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.RejectTicket", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RejectTicketOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RejectTicketOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) CloseTicket(ctx context.Context, in *CloseTicketInput, opts ...grpc.CallOption) (*CloseTicketOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CloseTicketInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CloseTicketInput, got %T", in))
		}

		return i.client.CloseTicket(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.CloseTicket", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CloseTicketOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CloseTicketOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) AddComment(ctx context.Context, in *AddCommentInput, opts ...grpc.CallOption) (*AddCommentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*AddCommentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *AddCommentInput, got %T", in))
		}

		return i.client.AddComment(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.AddComment", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*AddCommentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *AddCommentOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) UpdateComment(ctx context.Context, in *UpdateCommentInput, opts ...grpc.CallOption) (*UpdateCommentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateCommentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateCommentInput, got %T", in))
		}

		return i.client.UpdateComment(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.UpdateComment", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateCommentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateCommentOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) DeleteComment(ctx context.Context, in *DeleteCommentInput, opts ...grpc.CallOption) (*DeleteCommentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteCommentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteCommentInput, got %T", in))
		}

		return i.client.DeleteComment(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.DeleteComment", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteCommentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteCommentOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) ListComments(ctx context.Context, in *ListCommentsInput, opts ...grpc.CallOption) (*ListCommentsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListCommentsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListCommentsInput, got %T", in))
		}

		return i.client.ListComments(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.ListComments", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListCommentsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListCommentsOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) DescribeComment(ctx context.Context, in *DescribeCommentInput, opts ...grpc.CallOption) (*DescribeCommentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeCommentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeCommentInput, got %T", in))
		}

		return i.client.DescribeComment(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.DescribeComment", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeCommentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeCommentOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) CreateAutoReply(ctx context.Context, in *CreateAutoReplyInput, opts ...grpc.CallOption) (*CreateAutoReplyOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateAutoReplyInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateAutoReplyInput, got %T", in))
		}

		return i.client.CreateAutoReply(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.CreateAutoReply", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateAutoReplyOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateAutoReplyOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) UpdateAutoReply(ctx context.Context, in *UpdateAutoReplyInput, opts ...grpc.CallOption) (*UpdateAutoReplyOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateAutoReplyInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateAutoReplyInput, got %T", in))
		}

		return i.client.UpdateAutoReply(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.UpdateAutoReply", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateAutoReplyOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateAutoReplyOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) DeleteAutoReply(ctx context.Context, in *DeleteAutoReplyInput, opts ...grpc.CallOption) (*DeleteAutoReplyOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteAutoReplyInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteAutoReplyInput, got %T", in))
		}

		return i.client.DeleteAutoReply(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.DeleteAutoReply", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteAutoReplyOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteAutoReplyOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) ListAutoReplies(ctx context.Context, in *ListAutoRepliesInput, opts ...grpc.CallOption) (*ListAutoRepliesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListAutoRepliesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListAutoRepliesInput, got %T", in))
		}

		return i.client.ListAutoReplies(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.ListAutoReplies", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListAutoRepliesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListAutoRepliesOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) DescribeAutoReply(ctx context.Context, in *DescribeAutoReplyInput, opts ...grpc.CallOption) (*DescribeAutoReplyOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeAutoReplyInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeAutoReplyInput, got %T", in))
		}

		return i.client.DescribeAutoReply(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.DescribeAutoReply", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeAutoReplyOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeAutoReplyOutput, got %T", out))
	}

	return message, err
}

func (i *SupportInterceptor) UploadAttachment(ctx context.Context, in *UploadAttachmentInput, opts ...grpc.CallOption) (*UploadAttachmentOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UploadAttachmentInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UploadAttachmentInput, got %T", in))
		}

		return i.client.UploadAttachment(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.helpdesk.Support.UploadAttachment", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UploadAttachmentOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UploadAttachmentOutput, got %T", out))
	}

	return message, err
}
