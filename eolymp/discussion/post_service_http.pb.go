// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package discussion

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

// _PostService_HTTPReadQueryString parses body into proto.Message
func _PostService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _PostService_HTTPReadRequestBody parses body into proto.Message
func _PostService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _PostService_HTTPWriteResponse writes proto.Message to HTTP response
func _PostService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_PostService_HTTPWriteErrorResponse(w, err)
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

// _PostService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _PostService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// RegisterPostServiceHttpHandlers adds handlers for for PostServiceClient
func RegisterPostServiceHttpHandlers(router *mux.Router, prefix string, cli PostServiceClient) {
	router.Handle(prefix+"/posts/{post_id}", _PostService_DescribePost_Rule0(cli)).
		Methods("GET").
		Name("eolymp.discussion.PostService.DescribePost")
	router.Handle(prefix+"/posts", _PostService_ListPosts_Rule0(cli)).
		Methods("GET").
		Name("eolymp.discussion.PostService.ListPosts")
	router.Handle(prefix+"/posts", _PostService_CreatePost_Rule0(cli)).
		Methods("POST").
		Name("eolymp.discussion.PostService.CreatePost")
	router.Handle(prefix+"/posts/{post_id}", _PostService_UpdatePost_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.discussion.PostService.UpdatePost")
	router.Handle(prefix+"/posts/{post_id}/publish", _PostService_PublishPost_Rule0(cli)).
		Methods("POST").
		Name("eolymp.discussion.PostService.PublishPost")
	router.Handle(prefix+"/posts/{post_id}/publish", _PostService_UnpublishPost_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.discussion.PostService.UnpublishPost")
	router.Handle(prefix+"/posts/{post_id}/moderate", _PostService_ModeratePost_Rule0(cli)).
		Methods("POST").
		Name("eolymp.discussion.PostService.ModeratePost")
	router.Handle(prefix+"/posts/{post_id}", _PostService_DeletePost_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.discussion.PostService.DeletePost")
	router.Handle(prefix+"/posts/{post_id}/vote", _PostService_VotePost_Rule0(cli)).
		Methods("POST").
		Name("eolymp.discussion.PostService.VotePost")
	router.Handle(prefix+"/posts/{post_id}/translate", _PostService_TranslatePost_Rule0(cli)).
		Methods("POST").
		Name("eolymp.discussion.PostService.TranslatePost")
	router.Handle(prefix+"/posts/{post_id}/translations/{translation_id}", _PostService_DescribePostTranslation_Rule0(cli)).
		Methods("GET").
		Name("eolymp.discussion.PostService.DescribePostTranslation")
	router.Handle(prefix+"/posts/{post_id}/translations", _PostService_ListPostTranslations_Rule0(cli)).
		Methods("GET").
		Name("eolymp.discussion.PostService.ListPostTranslations")
	router.Handle(prefix+"/posts/{post_id}/translations", _PostService_CreatePostTranslation_Rule0(cli)).
		Methods("POST").
		Name("eolymp.discussion.PostService.CreatePostTranslation")
	router.Handle(prefix+"/posts/{post_id}/translations/{translation_id}", _PostService_UpdatePostTranslation_Rule0(cli)).
		Methods("PUT").
		Name("eolymp.discussion.PostService.UpdatePostTranslation")
	router.Handle(prefix+"/posts/{post_id}/translations/{translation_id}", _PostService_DeletePostTranslation_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.discussion.PostService.DeletePostTranslation")
}

// RegisterPostServiceHttpProxy adds proxy handlers for for PostServiceClient
func RegisterPostServiceHttpProxy(router *mux.Router, prefix string, conn grpc.ClientConnInterface) {
	RegisterPostServiceHttpHandlers(router, prefix, NewPostServiceClient(conn))
}

func _PostService_DescribePost_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePostInput{}

		if err := _PostService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribePost(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_ListPosts_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPostsInput{}

		if err := _PostService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListPosts(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_CreatePost_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreatePostInput{}

		if err := _PostService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreatePost(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_UpdatePost_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdatePostInput{}

		if err := _PostService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdatePost(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_PublishPost_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &PublishPostInput{}

		if err := _PostService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]

		var header, trailer metadata.MD

		out, err := cli.PublishPost(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_UnpublishPost_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UnpublishPostInput{}

		if err := _PostService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]

		var header, trailer metadata.MD

		out, err := cli.UnpublishPost(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_ModeratePost_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ModeratePostInput{}

		if err := _PostService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]

		var header, trailer metadata.MD

		out, err := cli.ModeratePost(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_DeletePost_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeletePostInput{}

		if err := _PostService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]

		var header, trailer metadata.MD

		out, err := cli.DeletePost(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_VotePost_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &VotePostInput{}

		if err := _PostService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]

		var header, trailer metadata.MD

		out, err := cli.VotePost(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_TranslatePost_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &TranslatePostInput{}

		if err := _PostService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]

		var header, trailer metadata.MD

		out, err := cli.TranslatePost(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_DescribePostTranslation_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePostTranslationInput{}

		if err := _PostService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]
		in.TranslationId = vars["translation_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribePostTranslation(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_ListPostTranslations_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPostTranslationsInput{}

		if err := _PostService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]

		var header, trailer metadata.MD

		out, err := cli.ListPostTranslations(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_CreatePostTranslation_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreatePostTranslationInput{}

		if err := _PostService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]

		var header, trailer metadata.MD

		out, err := cli.CreatePostTranslation(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_UpdatePostTranslation_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdatePostTranslationInput{}

		if err := _PostService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]
		in.TranslationId = vars["translation_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdatePostTranslation(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PostService_DeletePostTranslation_Rule0(cli PostServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeletePostTranslationInput{}

		if err := _PostService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PostId = vars["post_id"]
		in.TranslationId = vars["translation_id"]

		var header, trailer metadata.MD

		out, err := cli.DeletePostTranslation(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PostService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PostService_HTTPWriteResponse(w, out, header, trailer)
	})
}
