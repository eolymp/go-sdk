// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package ranker

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

// _ScoreboardService_HTTPReadQueryString parses body into proto.Message
func _ScoreboardService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _ScoreboardService_HTTPReadRequestBody parses body into proto.Message
func _ScoreboardService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _ScoreboardService_HTTPWriteResponse writes proto.Message to HTTP response
func _ScoreboardService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_ScoreboardService_HTTPWriteErrorResponse(w, err)
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

// _ScoreboardService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _ScoreboardService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterScoreboardServiceHttpHandlers adds handlers for for ScoreboardServiceClient
func RegisterScoreboardServiceHttpHandlers(router *mux.Router, prefix string, cli ScoreboardServiceClient) {
	router.Handle(prefix+"/scoreboards", _ScoreboardService_CreateScoreboard_Rule0(cli)).
		Methods("POST").
		Name("eolymp.ranker.ScoreboardService.CreateScoreboard")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}", _ScoreboardService_UpdateScoreboard_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.ranker.ScoreboardService.UpdateScoreboard")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/rebuild", _ScoreboardService_RebuildScoreboard_Rule0(cli)).
		Methods("POST").
		Name("eolymp.ranker.ScoreboardService.RebuildScoreboard")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}", _ScoreboardService_DeleteScoreboard_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.ranker.ScoreboardService.DeleteScoreboard")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}", _ScoreboardService_DescribeScoreboard_Rule0(cli)).
		Methods("GET").
		Name("eolymp.ranker.ScoreboardService.DescribeScoreboard")
	router.Handle(prefix+"/scoreboards:lookup", _ScoreboardService_LookupScoreboard_Rule0(cli)).
		Methods("POST").
		Name("eolymp.ranker.ScoreboardService.LookupScoreboard")
	router.Handle(prefix+"/scoreboards", _ScoreboardService_ListScoreboards_Rule0(cli)).
		Methods("GET").
		Name("eolymp.ranker.ScoreboardService.ListScoreboards")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/rows/{member_id}", _ScoreboardService_DescribeScoreboardRow_Rule0(cli)).
		Methods("GET").
		Name("eolymp.ranker.ScoreboardService.DescribeScoreboardRow")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/rows", _ScoreboardService_ListScoreboardRows_Rule0(cli)).
		Methods("GET").
		Name("eolymp.ranker.ScoreboardService.ListScoreboardRows")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/columns", _ScoreboardService_AddScoreboardColumn_Rule0(cli)).
		Methods("POST").
		Name("eolymp.ranker.ScoreboardService.AddScoreboardColumn")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/columns/{column_id}", _ScoreboardService_UpdateScoreboardColumn_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.ranker.ScoreboardService.UpdateScoreboardColumn")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/columns/{column_id}", _ScoreboardService_DeleteScoreboardColumn_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.ranker.ScoreboardService.DeleteScoreboardColumn")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/columns/{column_id}", _ScoreboardService_DescribeScoreboardColumn_Rule0(cli)).
		Methods("GET").
		Name("eolymp.ranker.ScoreboardService.DescribeScoreboardColumn")
	router.Handle(prefix+"/scoreboards/{scoreboard_id}/columns", _ScoreboardService_ListScoreboardColumns_Rule0(cli)).
		Methods("GET").
		Name("eolymp.ranker.ScoreboardService.ListScoreboardColumns")
}

// RegisterScoreboardServiceHttpProxy adds proxy handlers for for ScoreboardServiceClient
func RegisterScoreboardServiceHttpProxy(router *mux.Router, prefix string, conn grpc.ClientConnInterface) {
	RegisterScoreboardServiceHttpHandlers(router, prefix, NewScoreboardServiceClient(conn))
}

func _ScoreboardService_CreateScoreboard_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreateScoreboardInput{}

		if err := _ScoreboardService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreateScoreboard(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_UpdateScoreboard_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateScoreboardInput{}

		if err := _ScoreboardService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateScoreboard(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_RebuildScoreboard_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &RebuildScoreboardInput{}

		if err := _ScoreboardService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		var header, trailer metadata.MD

		out, err := cli.RebuildScoreboard(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_DeleteScoreboard_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteScoreboardInput{}

		if err := _ScoreboardService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteScoreboard(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_DescribeScoreboard_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeScoreboardInput{}

		if err := _ScoreboardService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeScoreboard(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_LookupScoreboard_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &LookupScoreboardInput{}

		if err := _ScoreboardService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.LookupScoreboard(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_ListScoreboards_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListScoreboardsInput{}

		if err := _ScoreboardService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListScoreboards(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_DescribeScoreboardRow_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeScoreboardRowInput{}

		if err := _ScoreboardService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]
		in.MemberId = vars["member_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeScoreboardRow(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_ListScoreboardRows_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListScoreboardRowsInput{}

		if err := _ScoreboardService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		var header, trailer metadata.MD

		out, err := cli.ListScoreboardRows(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_AddScoreboardColumn_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &AddScoreboardColumnInput{}

		if err := _ScoreboardService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		var header, trailer metadata.MD

		out, err := cli.AddScoreboardColumn(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_UpdateScoreboardColumn_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdateScoreboardColumnInput{}

		if err := _ScoreboardService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]
		in.ColumnId = vars["column_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdateScoreboardColumn(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_DeleteScoreboardColumn_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeleteScoreboardColumnInput{}

		if err := _ScoreboardService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]
		in.ColumnId = vars["column_id"]

		var header, trailer metadata.MD

		out, err := cli.DeleteScoreboardColumn(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_DescribeScoreboardColumn_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeScoreboardColumnInput{}

		if err := _ScoreboardService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]
		in.ColumnId = vars["column_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribeScoreboardColumn(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _ScoreboardService_ListScoreboardColumns_Rule0(cli ScoreboardServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListScoreboardColumnsInput{}

		if err := _ScoreboardService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.ScoreboardId = vars["scoreboard_id"]

		var header, trailer metadata.MD

		out, err := cli.ListScoreboardColumns(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_ScoreboardService_HTTPWriteErrorResponse(w, err)
			return
		}

		_ScoreboardService_HTTPWriteResponse(w, out, header, trailer)
	})
}
