// Code generated by protoc-gen-go-server. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-server for more details.

package printer

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

// RegisterPrinterServiceHttpHandlers adds handlers for for PrinterServiceClient
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

// RegisterPrinterServiceHttpProxy adds proxy handlers for for PrinterServiceClient
func RegisterPrinterServiceHttpProxy(router *mux.Router, prefix string, conn grpc.ClientConnInterface) {
	RegisterPrinterServiceHttpHandlers(router, prefix, NewPrinterServiceClient(conn))
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
