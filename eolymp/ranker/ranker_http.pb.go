// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package ranker

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

// _Ranker_HTTPReadQueryString parses body into proto.Message
func _Ranker_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Ranker_HTTPReadRequestBody parses body into proto.Message
func _Ranker_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Ranker_HTTPWriteResponse writes proto.Message to HTTP response
func _Ranker_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Ranker_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Ranker_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Ranker_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterRankerHttpHandlers adds handlers for for RankerServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterRankerHttpHandlers(router *mux.Router, prefix string, srv RankerServer) {
	router.Handle(prefix+"/scoreboards", _Ranker_CreateScoreboard_Rule0(srv)).
		Methods("POST").
		Name("eolymp.ranker.Ranker.CreateScoreboard")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}", _Ranker_UpdateScoreboard_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.ranker.Ranker.UpdateScoreboard")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/rebuild", _Ranker_RebuildScoreboard_Rule0(srv)).
		Methods("POST").
		Name("eolymp.ranker.Ranker.RebuildScoreboard")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}", _Ranker_DeleteScoreboard_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.ranker.Ranker.DeleteScoreboard")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}", _Ranker_DescribeScoreboard_Rule0(srv)).
		Methods("GET").
		Name("eolymp.ranker.Ranker.DescribeScoreboard")
	router.Handle(prefix+"/scoreboards", _Ranker_ListScoreboards_Rule0(srv)).
		Methods("GET").
		Name("eolymp.ranker.Ranker.ListScoreboards")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/rows/{member_id}", _Ranker_DescribeScoreboardRow_Rule0(srv)).
		Methods("GET").
		Name("eolymp.ranker.Ranker.DescribeScoreboardRow")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/rows", _Ranker_ListScoreboardRows_Rule0(srv)).
		Methods("GET").
		Name("eolymp.ranker.Ranker.ListScoreboardRows")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/columns", _Ranker_AddScoreboardColumn_Rule0(srv)).
		Methods("POST").
		Name("eolymp.ranker.Ranker.AddScoreboardColumn")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/columns/{column_id}", _Ranker_UpdateScoreboardColumn_Rule0(srv)).
		Methods("PUT").
		Name("eolymp.ranker.Ranker.UpdateScoreboardColumn")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/columns/{column_id}", _Ranker_DeleteScoreboardColumn_Rule0(srv)).
		Methods("DELETE").
		Name("eolymp.ranker.Ranker.DeleteScoreboardColumn")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/columns/{column_id}", _Ranker_DescribeScoreboardColumn_Rule0(srv)).
		Methods("GET").
		Name("eolymp.ranker.Ranker.DescribeScoreboardColumn")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/columns", _Ranker_ListScoreboardColumns_Rule0(srv)).
		Methods("GET").
		Name("eolymp.ranker.Ranker.ListScoreboardColumns")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/activities", _Ranker_ListActivities_Rule0(srv)).
		Methods("GET").
		Name("eolymp.ranker.Ranker.ListActivities")
}

func _Ranker_CreateScoreboard_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateScoreboardInput{}

		if err := _Ranker_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.CreateScoreboard(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_UpdateScoreboard_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateScoreboardInput{}

		if err := _Ranker_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		out, err := srv.UpdateScoreboard(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_RebuildScoreboard_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RebuildScoreboardInput{}

		if err := _Ranker_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		out, err := srv.RebuildScoreboard(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_DeleteScoreboard_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteScoreboardInput{}

		if err := _Ranker_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		out, err := srv.DeleteScoreboard(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_DescribeScoreboard_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeScoreboardInput{}

		if err := _Ranker_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		out, err := srv.DescribeScoreboard(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_ListScoreboards_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListScoreboardsInput{}

		if err := _Ranker_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListScoreboards(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_DescribeScoreboardRow_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeScoreboardRowInput{}

		if err := _Ranker_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]
		in.MemberId = vars["member_id"]

		out, err := srv.DescribeScoreboardRow(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_ListScoreboardRows_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListScoreboardRowsInput{}

		if err := _Ranker_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		out, err := srv.ListScoreboardRows(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_AddScoreboardColumn_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &AddScoreboardColumnInput{}

		if err := _Ranker_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		out, err := srv.AddScoreboardColumn(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_UpdateScoreboardColumn_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateScoreboardColumnInput{}

		if err := _Ranker_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]
		in.ColumnId = vars["column_id"]

		out, err := srv.UpdateScoreboardColumn(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_DeleteScoreboardColumn_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteScoreboardColumnInput{}

		if err := _Ranker_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]
		in.ColumnId = vars["column_id"]

		out, err := srv.DeleteScoreboardColumn(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_DescribeScoreboardColumn_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeScoreboardColumnInput{}

		if err := _Ranker_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]
		in.ColumnId = vars["column_id"]

		out, err := srv.DescribeScoreboardColumn(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_ListScoreboardColumns_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListScoreboardColumnsInput{}

		if err := _Ranker_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		out, err := srv.ListScoreboardColumns(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

func _Ranker_ListActivities_Rule0(srv RankerServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListActivitiesInput{}

		if err := _Ranker_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		out, err := srv.ListActivities(r.Context(), in)
		if err != nil {
			_Ranker_HTTPWriteErrorResponse(w, err)
			return
		}

		_Ranker_HTTPWriteResponse(w, out)
	})
}

type _RankerHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _RankerMiddleware = func(ctx context.Context, method string, in proto.Message, handler _RankerHandler) (out proto.Message, err error)
type RankerInterceptor struct {
	middleware []_RankerMiddleware
	server     RankerServer
}

// NewRankerInterceptor constructs additional middleware for a server based on annotations in proto files
func NewRankerInterceptor(srv RankerServer, middleware ..._RankerMiddleware) *RankerInterceptor {
	return &RankerInterceptor{server: srv, middleware: middleware}
}

func (i *RankerInterceptor) CreateScoreboard(ctx context.Context, in *CreateScoreboardInput) (*CreateScoreboardOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreateScoreboardInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreateScoreboardInput, got %T", in))
		}

		return i.server.CreateScoreboard(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.CreateScoreboard", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreateScoreboardOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreateScoreboardOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) UpdateScoreboard(ctx context.Context, in *UpdateScoreboardInput) (*UpdateScoreboardOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateScoreboardInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateScoreboardInput, got %T", in))
		}

		return i.server.UpdateScoreboard(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.UpdateScoreboard", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateScoreboardOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateScoreboardOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) RebuildScoreboard(ctx context.Context, in *RebuildScoreboardInput) (*RebuildScoreboardOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*RebuildScoreboardInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *RebuildScoreboardInput, got %T", in))
		}

		return i.server.RebuildScoreboard(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.RebuildScoreboard", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*RebuildScoreboardOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *RebuildScoreboardOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) DeleteScoreboard(ctx context.Context, in *DeleteScoreboardInput) (*DeleteScoreboardOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteScoreboardInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteScoreboardInput, got %T", in))
		}

		return i.server.DeleteScoreboard(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.DeleteScoreboard", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteScoreboardOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteScoreboardOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) DescribeScoreboard(ctx context.Context, in *DescribeScoreboardInput) (*DescribeScoreboardOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeScoreboardInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeScoreboardInput, got %T", in))
		}

		return i.server.DescribeScoreboard(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.DescribeScoreboard", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeScoreboardOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeScoreboardOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) ListScoreboards(ctx context.Context, in *ListScoreboardsInput) (*ListScoreboardsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListScoreboardsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListScoreboardsInput, got %T", in))
		}

		return i.server.ListScoreboards(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.ListScoreboards", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListScoreboardsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListScoreboardsOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) DescribeScoreboardRow(ctx context.Context, in *DescribeScoreboardRowInput) (*DescribeScoreboardRowOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeScoreboardRowInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeScoreboardRowInput, got %T", in))
		}

		return i.server.DescribeScoreboardRow(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.DescribeScoreboardRow", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeScoreboardRowOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeScoreboardRowOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) ListScoreboardRows(ctx context.Context, in *ListScoreboardRowsInput) (*ListScoreboardRowsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListScoreboardRowsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListScoreboardRowsInput, got %T", in))
		}

		return i.server.ListScoreboardRows(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.ListScoreboardRows", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListScoreboardRowsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListScoreboardRowsOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) AddScoreboardColumn(ctx context.Context, in *AddScoreboardColumnInput) (*AddScoreboardColumnOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*AddScoreboardColumnInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *AddScoreboardColumnInput, got %T", in))
		}

		return i.server.AddScoreboardColumn(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.AddScoreboardColumn", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*AddScoreboardColumnOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *AddScoreboardColumnOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) UpdateScoreboardColumn(ctx context.Context, in *UpdateScoreboardColumnInput) (*UpdateScoreboardColumnOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdateScoreboardColumnInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdateScoreboardColumnInput, got %T", in))
		}

		return i.server.UpdateScoreboardColumn(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.UpdateScoreboardColumn", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdateScoreboardColumnOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdateScoreboardColumnOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) DeleteScoreboardColumn(ctx context.Context, in *DeleteScoreboardColumnInput) (*DeleteScoreboardColumnOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeleteScoreboardColumnInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeleteScoreboardColumnInput, got %T", in))
		}

		return i.server.DeleteScoreboardColumn(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.DeleteScoreboardColumn", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeleteScoreboardColumnOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeleteScoreboardColumnOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) DescribeScoreboardColumn(ctx context.Context, in *DescribeScoreboardColumnInput) (*DescribeScoreboardColumnOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribeScoreboardColumnInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribeScoreboardColumnInput, got %T", in))
		}

		return i.server.DescribeScoreboardColumn(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.DescribeScoreboardColumn", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribeScoreboardColumnOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribeScoreboardColumnOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) ListScoreboardColumns(ctx context.Context, in *ListScoreboardColumnsInput) (*ListScoreboardColumnsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListScoreboardColumnsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListScoreboardColumnsInput, got %T", in))
		}

		return i.server.ListScoreboardColumns(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.ListScoreboardColumns", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListScoreboardColumnsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListScoreboardColumnsOutput, got %T", out))
	}

	return message, err
}

func (i *RankerInterceptor) ListActivities(ctx context.Context, in *ListActivitiesInput) (*ListActivitiesOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListActivitiesInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListActivitiesInput, got %T", in))
		}

		return i.server.ListActivities(ctx, message)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.ranker.Ranker.ListActivities", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListActivitiesOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListActivitiesOutput, got %T", out))
	}

	return message, err
}
