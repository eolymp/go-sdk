// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package printer

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

// _PrinterService_HTTPReadQueryString parses body into proto.Message
func _PrinterService_HTTPReadQueryString(r *http.Request, v proto.Message) error {
	query := r.URL.Query().Get("q")
	if query == "" {
		return nil
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(query), v)); err != nil {
		return err
	}

	return nil
}

// _PrinterService_HTTPReadRequestBody parses body into proto.Message
func _PrinterService_HTTPReadRequestBody(r *http.Request, v proto.Message) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, v)); err != nil {
		return err
	}

	return nil
}

// _PrinterService_HTTPWriteResponse writes proto.Message to HTTP response
func _PrinterService_HTTPWriteResponse(w http.ResponseWriter, v proto.Message, h, t metadata.MD) {
	data, err := protojson.Marshal(v)
	if err != nil {
		_PrinterService_HTTPWriteErrorResponse(w, err)
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

// _PrinterService_HTTPWriteErrorResponse writes error to HTTP response with error status code
func _PrinterService_HTTPWriteErrorResponse(w http.ResponseWriter, e error) {
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

// _PrinterService_WebsocketErrorResponse writes error to websocket connection
func _PrinterService_WebsocketErrorResponse(conn *websocket.Conn, e error) {
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

// _PrinterService_WebsocketCodec implements protobuf codec for websockets package
var _PrinterService_WebsocketCodec = websocket.Codec{
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

// RegisterPrinterServiceHttpHandlers adds handlers for for PrinterServiceClient
// This constructor creates http.Handler, the actual implementation might change at any moment
func RegisterPrinterServiceHttpHandlers(router *mux.Router, prefix string, cli PrinterServiceClient) {
	router.Handle(prefix+"/printers", _PrinterService_CreatePrinter_Rule0(cli)).
		Methods("POST").
		Name("eolymp.printer.PrinterService.CreatePrinter")
	router.Handle(prefix+"/printers/{printer_id}", _PrinterService_UpdatePrinter_Rule0(cli)).
		Methods("POST").
		Name("eolymp.printer.PrinterService.UpdatePrinter")
	router.Handle(prefix+"/printers/{printer_id}", _PrinterService_DeletePrinter_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.printer.PrinterService.DeletePrinter")
	router.Handle(prefix+"/printers/{printer_id}", _PrinterService_DescribePrinter_Rule0(cli)).
		Methods("GET").
		Name("eolymp.printer.PrinterService.DescribePrinter")
	router.Handle(prefix+"/printers", _PrinterService_ListPrinters_Rule0(cli)).
		Methods("GET").
		Name("eolymp.printer.PrinterService.ListPrinters")
	router.Handle(prefix+"/printers/{printer_id}/jobs", _PrinterService_CreatePrinterJob_Rule0(cli)).
		Methods("POST").
		Name("eolymp.printer.PrinterService.CreatePrinterJob")
	router.Handle(prefix+"/printers/{printer_id}/jobs/{job_id}", _PrinterService_DescribePrinterJob_Rule0(cli)).
		Methods("GET").
		Name("eolymp.printer.PrinterService.DescribePrinterJob")
	router.Handle(prefix+"/printers/{printer_id}/jobs", _PrinterService_ListPrinterJobs_Rule0(cli)).
		Methods("GET").
		Name("eolymp.printer.PrinterService.ListPrinterJobs")
	router.Handle(prefix+"/printers/{printer_id}/jobs/{job_id}", _PrinterService_DeletePrinterJob_Rule0(cli)).
		Methods("DELETE").
		Name("eolymp.printer.PrinterService.DeletePrinterJob")
}

func _PrinterService_CreatePrinter_Rule0(cli PrinterServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreatePrinterInput{}

		if err := _PrinterService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.CreatePrinter(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PrinterService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PrinterService_UpdatePrinter_Rule0(cli PrinterServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &UpdatePrinterInput{}

		if err := _PrinterService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PrinterId = vars["printer_id"]

		var header, trailer metadata.MD

		out, err := cli.UpdatePrinter(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PrinterService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PrinterService_DeletePrinter_Rule0(cli PrinterServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeletePrinterInput{}

		if err := _PrinterService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PrinterId = vars["printer_id"]

		var header, trailer metadata.MD

		out, err := cli.DeletePrinter(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PrinterService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PrinterService_DescribePrinter_Rule0(cli PrinterServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePrinterInput{}

		if err := _PrinterService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PrinterId = vars["printer_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribePrinter(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PrinterService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PrinterService_ListPrinters_Rule0(cli PrinterServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPrintersInput{}

		if err := _PrinterService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		var header, trailer metadata.MD

		out, err := cli.ListPrinters(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PrinterService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PrinterService_CreatePrinterJob_Rule0(cli PrinterServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &CreatePrinterJobInput{}

		if err := _PrinterService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PrinterId = vars["printer_id"]

		var header, trailer metadata.MD

		out, err := cli.CreatePrinterJob(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PrinterService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PrinterService_DescribePrinterJob_Rule0(cli PrinterServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DescribePrinterJobInput{}

		if err := _PrinterService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PrinterId = vars["printer_id"]
		in.JobId = vars["job_id"]

		var header, trailer metadata.MD

		out, err := cli.DescribePrinterJob(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PrinterService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PrinterService_ListPrinterJobs_Rule0(cli PrinterServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &ListPrinterJobsInput{}

		if err := _PrinterService_HTTPReadQueryString(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PrinterId = vars["printer_id"]

		var header, trailer metadata.MD

		out, err := cli.ListPrinterJobs(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PrinterService_HTTPWriteResponse(w, out, header, trailer)
	})
}

func _PrinterService_DeletePrinterJob_Rule0(cli PrinterServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := &DeletePrinterJobInput{}

		if err := _PrinterService_HTTPReadRequestBody(r, in); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		vars := mux.Vars(r)
		in.PrinterId = vars["printer_id"]
		in.JobId = vars["job_id"]

		var header, trailer metadata.MD

		out, err := cli.DeletePrinterJob(r.Context(), in, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			_PrinterService_HTTPWriteErrorResponse(w, err)
			return
		}

		_PrinterService_HTTPWriteResponse(w, out, header, trailer)
	})
}

type _PrinterServiceHandler = func(ctx context.Context, in proto.Message) (proto.Message, error)
type _PrinterServiceMiddleware = func(ctx context.Context, method string, in proto.Message, handler _PrinterServiceHandler) (out proto.Message, err error)
type PrinterServiceInterceptor struct {
	middleware []_PrinterServiceMiddleware
	client     PrinterServiceClient
}

// NewPrinterServiceInterceptor constructs additional middleware for a server based on annotations in proto files
func NewPrinterServiceInterceptor(cli PrinterServiceClient, middleware ..._PrinterServiceMiddleware) *PrinterServiceInterceptor {
	return &PrinterServiceInterceptor{client: cli, middleware: middleware}
}

func (i *PrinterServiceInterceptor) CreatePrinter(ctx context.Context, in *CreatePrinterInput, opts ...grpc.CallOption) (*CreatePrinterOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreatePrinterInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreatePrinterInput, got %T", in))
		}

		return i.client.CreatePrinter(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.printer.PrinterService.CreatePrinter", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreatePrinterOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreatePrinterOutput, got %T", out))
	}

	return message, err
}

func (i *PrinterServiceInterceptor) UpdatePrinter(ctx context.Context, in *UpdatePrinterInput, opts ...grpc.CallOption) (*UpdatePrinterOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*UpdatePrinterInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *UpdatePrinterInput, got %T", in))
		}

		return i.client.UpdatePrinter(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.printer.PrinterService.UpdatePrinter", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*UpdatePrinterOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *UpdatePrinterOutput, got %T", out))
	}

	return message, err
}

func (i *PrinterServiceInterceptor) DeletePrinter(ctx context.Context, in *DeletePrinterInput, opts ...grpc.CallOption) (*DeletePrinterOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeletePrinterInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeletePrinterInput, got %T", in))
		}

		return i.client.DeletePrinter(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.printer.PrinterService.DeletePrinter", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeletePrinterOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeletePrinterOutput, got %T", out))
	}

	return message, err
}

func (i *PrinterServiceInterceptor) DescribePrinter(ctx context.Context, in *DescribePrinterInput, opts ...grpc.CallOption) (*DescribePrinterOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribePrinterInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribePrinterInput, got %T", in))
		}

		return i.client.DescribePrinter(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.printer.PrinterService.DescribePrinter", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribePrinterOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribePrinterOutput, got %T", out))
	}

	return message, err
}

func (i *PrinterServiceInterceptor) ListPrinters(ctx context.Context, in *ListPrintersInput, opts ...grpc.CallOption) (*ListPrintersOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListPrintersInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListPrintersInput, got %T", in))
		}

		return i.client.ListPrinters(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.printer.PrinterService.ListPrinters", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListPrintersOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListPrintersOutput, got %T", out))
	}

	return message, err
}

func (i *PrinterServiceInterceptor) CreatePrinterJob(ctx context.Context, in *CreatePrinterJobInput, opts ...grpc.CallOption) (*CreatePrinterJobOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*CreatePrinterJobInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *CreatePrinterJobInput, got %T", in))
		}

		return i.client.CreatePrinterJob(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.printer.PrinterService.CreatePrinterJob", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*CreatePrinterJobOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *CreatePrinterJobOutput, got %T", out))
	}

	return message, err
}

func (i *PrinterServiceInterceptor) DescribePrinterJob(ctx context.Context, in *DescribePrinterJobInput, opts ...grpc.CallOption) (*DescribePrinterJobOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DescribePrinterJobInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DescribePrinterJobInput, got %T", in))
		}

		return i.client.DescribePrinterJob(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.printer.PrinterService.DescribePrinterJob", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DescribePrinterJobOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DescribePrinterJobOutput, got %T", out))
	}

	return message, err
}

func (i *PrinterServiceInterceptor) ListPrinterJobs(ctx context.Context, in *ListPrinterJobsInput, opts ...grpc.CallOption) (*ListPrinterJobsOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*ListPrinterJobsInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *ListPrinterJobsInput, got %T", in))
		}

		return i.client.ListPrinterJobs(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.printer.PrinterService.ListPrinterJobs", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*ListPrinterJobsOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *ListPrinterJobsOutput, got %T", out))
	}

	return message, err
}

func (i *PrinterServiceInterceptor) DeletePrinterJob(ctx context.Context, in *DeletePrinterJobInput, opts ...grpc.CallOption) (*DeletePrinterJobOutput, error) {
	handler := func(ctx context.Context, in proto.Message) (proto.Message, error) {
		message, ok := in.(*DeletePrinterJobInput)
		if !ok && in != nil {
			panic(fmt.Errorf("request input type is invalid: want *DeletePrinterJobInput, got %T", in))
		}

		return i.client.DeletePrinterJob(ctx, message, opts...)
	}

	for _, mw := range i.middleware {
		mw := mw
		next := handler

		handler = func(ctx context.Context, in proto.Message) (proto.Message, error) {
			return mw(ctx, "eolymp.printer.PrinterService.DeletePrinterJob", in, next)
		}
	}

	out, err := handler(ctx, in)
	if err != nil {
		return nil, err
	}

	message, ok := out.(*DeletePrinterJobOutput)
	if !ok && out != nil {
		panic(fmt.Errorf("output type is invalid: want *DeletePrinterJobOutput, got %T", out))
	}

	return message, err
}
