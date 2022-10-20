// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package geography

import (
	context "context"
	mux "github.com/gorilla/mux"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
)

// _Geography_HTTPReadQueryString parses body into proto.Message
func _Geography_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _Geography_HTTPReadRequestBody parses body into proto.Message
func _Geography_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _Geography_HTTPWriteResponse writes proto.Message to HTTP response
func _Geography_HTTPWriteResponse(w http.ResponseWriter, v proto.Message) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_Geography_HTTPWriteErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

// _Geography_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _Geography_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// NewGeographyHandler constructs new http.Handler for GeographyServer
func NewGeographyHandler(srv GeographyServer) http.Handler {
	router := mux.NewRouter()
	router.Handle("/eolymp.geography.Geography/DescribeCountry", _Geography_DescribeCountry(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.geography.Geography/ListCountries", _Geography_ListCountries(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.geography.Geography/DescribeRegion", _Geography_DescribeRegion(srv)).Methods(http.MethodPost)
	router.Handle("/eolymp.geography.Geography/ListRegions", _Geography_ListRegions(srv)).Methods(http.MethodPost)
	return router
}

// NewGeographyHandlerHttp constructs new http.Handler for GeographyServer
// This constructor creates http.Handler, the actual implementation might change at any moment
func NewGeographyHandlerHttp(srv GeographyServer, prefix string) http.Handler {
	router := mux.NewRouter()

	router.Handle(prefix+"/geography/countries/{country_id}", _Geography_DescribeCountry_Rule0(srv)).
		Methods("GET").
		Name("eolymp.geography.Geography.DescribeCountry")

	router.Handle(prefix+"/geography/countries", _Geography_ListCountries_Rule0(srv)).
		Methods("GET").
		Name("eolymp.geography.Geography.ListCountries")

	router.Handle(prefix+"/geography/regions/{region_id}", _Geography_DescribeRegion_Rule0(srv)).
		Methods("GET").
		Name("eolymp.geography.Geography.DescribeRegion")

	router.Handle(prefix+"/geography/countries/{country_id}/regions", _Geography_ListRegions_Rule0(srv)).
		Methods("GET").
		Name("eolymp.geography.Geography.ListRegions")

	return router
}

func _Geography_DescribeCountry(srv GeographyServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCountryInput{}

		if err := _Geography_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeCountry(r.Context(), in)
		if err != nil {
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		_Geography_HTTPWriteResponse(w, out)
	})
}

func _Geography_ListCountries(srv GeographyServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListCountriesInput{}

		if err := _Geography_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListCountries(r.Context(), in)
		if err != nil {
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		_Geography_HTTPWriteResponse(w, out)
	})
}

func _Geography_DescribeRegion(srv GeographyServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRegionInput{}

		if err := _Geography_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.DescribeRegion(r.Context(), in)
		if err != nil {
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		_Geography_HTTPWriteResponse(w, out)
	})
}

func _Geography_ListRegions(srv GeographyServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListRegionsInput{}

		if err := _Geography_HTTPReadRequestBody(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListRegions(r.Context(), in)
		if err != nil {
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		_Geography_HTTPWriteResponse(w, out)
	})
}

func _Geography_DescribeCountry_Rule0(srv GeographyServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeCountryInput{}

		if err := _Geography_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.CountryId = vars["country_id"]

		out, err := srv.DescribeCountry(r.Context(), in)
		if err != nil {
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		_Geography_HTTPWriteResponse(w, out)
	})
}

func _Geography_ListCountries_Rule0(srv GeographyServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListCountriesInput{}

		if err := _Geography_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		out, err := srv.ListCountries(r.Context(), in)
		if err != nil {
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		_Geography_HTTPWriteResponse(w, out)
	})
}

func _Geography_DescribeRegion_Rule0(srv GeographyServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribeRegionInput{}

		if err := _Geography_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.RegionId = vars["region_id"]

		out, err := srv.DescribeRegion(r.Context(), in)
		if err != nil {
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		_Geography_HTTPWriteResponse(w, out)
	})
}

func _Geography_ListRegions_Rule0(srv GeographyServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListRegionsInput{}

		if err := _Geography_HTTPReadQueryString(r, in); err != nil {
			err = status.New(codes.InvalidArgument, err.Error()).Err()
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.CountryId = vars["country_id"]

		out, err := srv.ListRegions(r.Context(), in)
		if err != nil {
			_Geography_HTTPWriteErrorResponse(w, err)
			return
		}

		_Geography_HTTPWriteResponse(w, out)
	})
}

type _GeographyMiddleware = func(ctx context.Context, method string, in proto.Message, next func() (out proto.Message, err error))
type GeographyInterceptor struct {
	middleware []_GeographyMiddleware
	server     GeographyServer
}

// NewGeographyInterceptor constructs additional middleware for a server based on annotations in proto files
func NewGeographyInterceptor(srv GeographyServer, middleware ..._GeographyMiddleware) *GeographyInterceptor {
	return &GeographyInterceptor{server: srv, middleware: middleware}
}

func (i *GeographyInterceptor) DescribeCountry(ctx context.Context, in *DescribeCountryInput) (out *DescribeCountryOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.DescribeCountry(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.geography.Geography/DescribeCountry", in, handler)
			return out, err
		}
	}

	next()
	return
}

func (i *GeographyInterceptor) ListCountries(ctx context.Context, in *ListCountriesInput) (out *ListCountriesOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.ListCountries(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.geography.Geography/ListCountries", in, handler)
			return out, err
		}
	}

	next()
	return
}

func (i *GeographyInterceptor) DescribeRegion(ctx context.Context, in *DescribeRegionInput) (out *DescribeRegionOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.DescribeRegion(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.geography.Geography/DescribeRegion", in, handler)
			return out, err
		}
	}

	next()
	return
}

func (i *GeographyInterceptor) ListRegions(ctx context.Context, in *ListRegionsInput) (out *ListRegionsOutput, err error) {
	next := func() (proto.Message, error) {
		out, err = i.server.ListRegions(ctx, in)
		return out, err
	}

	for _, mw := range i.middleware {
		handler := next

		next = func() (proto.Message, error) {
			mw(ctx, "eolymp.geography.Geography/ListRegions", in, handler)
			return out, err
		}
	}

	next()
	return
}
