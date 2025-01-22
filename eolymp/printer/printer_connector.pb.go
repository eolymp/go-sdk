// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.24.4
// source: eolymp/printer/printer_connector.proto

package printer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PrinterConnectorClientMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Message:
	//
	//	*PrinterConnectorClientMessage_Status_
	//	*PrinterConnectorClientMessage_Report_
	Message       isPrinterConnectorClientMessage_Message `protobuf_oneof:"message"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrinterConnectorClientMessage) Reset() {
	*x = PrinterConnectorClientMessage{}
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrinterConnectorClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrinterConnectorClientMessage) ProtoMessage() {}

func (x *PrinterConnectorClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrinterConnectorClientMessage.ProtoReflect.Descriptor instead.
func (*PrinterConnectorClientMessage) Descriptor() ([]byte, []int) {
	return file_eolymp_printer_printer_connector_proto_rawDescGZIP(), []int{0}
}

func (x *PrinterConnectorClientMessage) GetMessage() isPrinterConnectorClientMessage_Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *PrinterConnectorClientMessage) GetStatus() *PrinterConnectorClientMessage_Status {
	if x != nil {
		if x, ok := x.Message.(*PrinterConnectorClientMessage_Status_); ok {
			return x.Status
		}
	}
	return nil
}

func (x *PrinterConnectorClientMessage) GetReport() *PrinterConnectorClientMessage_Report {
	if x != nil {
		if x, ok := x.Message.(*PrinterConnectorClientMessage_Report_); ok {
			return x.Report
		}
	}
	return nil
}

type isPrinterConnectorClientMessage_Message interface {
	isPrinterConnectorClientMessage_Message()
}

type PrinterConnectorClientMessage_Status_ struct {
	Status *PrinterConnectorClientMessage_Status `protobuf:"bytes,101,opt,name=status,proto3,oneof"`
}

type PrinterConnectorClientMessage_Report_ struct {
	Report *PrinterConnectorClientMessage_Report `protobuf:"bytes,102,opt,name=report,proto3,oneof"`
}

func (*PrinterConnectorClientMessage_Status_) isPrinterConnectorClientMessage_Message() {}

func (*PrinterConnectorClientMessage_Report_) isPrinterConnectorClientMessage_Message() {}

type PrinterConnectorServerMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Message:
	//
	//	*PrinterConnectorServerMessage_Hello_
	//	*PrinterConnectorServerMessage_Print_
	Message       isPrinterConnectorServerMessage_Message `protobuf_oneof:"message"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrinterConnectorServerMessage) Reset() {
	*x = PrinterConnectorServerMessage{}
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrinterConnectorServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrinterConnectorServerMessage) ProtoMessage() {}

func (x *PrinterConnectorServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrinterConnectorServerMessage.ProtoReflect.Descriptor instead.
func (*PrinterConnectorServerMessage) Descriptor() ([]byte, []int) {
	return file_eolymp_printer_printer_connector_proto_rawDescGZIP(), []int{1}
}

func (x *PrinterConnectorServerMessage) GetMessage() isPrinterConnectorServerMessage_Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *PrinterConnectorServerMessage) GetHello() *PrinterConnectorServerMessage_Hello {
	if x != nil {
		if x, ok := x.Message.(*PrinterConnectorServerMessage_Hello_); ok {
			return x.Hello
		}
	}
	return nil
}

func (x *PrinterConnectorServerMessage) GetPrint() *PrinterConnectorServerMessage_Print {
	if x != nil {
		if x, ok := x.Message.(*PrinterConnectorServerMessage_Print_); ok {
			return x.Print
		}
	}
	return nil
}

type isPrinterConnectorServerMessage_Message interface {
	isPrinterConnectorServerMessage_Message()
}

type PrinterConnectorServerMessage_Hello_ struct {
	Hello *PrinterConnectorServerMessage_Hello `protobuf:"bytes,100,opt,name=hello,proto3,oneof"`
}

type PrinterConnectorServerMessage_Print_ struct {
	Print *PrinterConnectorServerMessage_Print `protobuf:"bytes,110,opt,name=print,proto3,oneof"`
}

func (*PrinterConnectorServerMessage_Hello_) isPrinterConnectorServerMessage_Message() {}

func (*PrinterConnectorServerMessage_Print_) isPrinterConnectorServerMessage_Message() {}

// Report the status of a job back to server
// Only after receiving a definitive (complete or error) report the printer will be given another job
type PrinterConnectorClientMessage_Report struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Status        Job_Status             `protobuf:"varint,1,opt,name=status,proto3,enum=eolymp.printer.Job_Status" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrinterConnectorClientMessage_Report) Reset() {
	*x = PrinterConnectorClientMessage_Report{}
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrinterConnectorClientMessage_Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrinterConnectorClientMessage_Report) ProtoMessage() {}

func (x *PrinterConnectorClientMessage_Report) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrinterConnectorClientMessage_Report.ProtoReflect.Descriptor instead.
func (*PrinterConnectorClientMessage_Report) Descriptor() ([]byte, []int) {
	return file_eolymp_printer_printer_connector_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PrinterConnectorClientMessage_Report) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PrinterConnectorClientMessage_Report) GetStatus() Job_Status {
	if x != nil {
		return x.Status
	}
	return Job_UNKNOWN_STATUS
}

// Report the status of the printer
// The printer won't be given any job if its status is other than READY
type PrinterConnectorClientMessage_Status struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        Printer_Status         `protobuf:"varint,1,opt,name=status,proto3,enum=eolymp.printer.Printer_Status" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrinterConnectorClientMessage_Status) Reset() {
	*x = PrinterConnectorClientMessage_Status{}
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrinterConnectorClientMessage_Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrinterConnectorClientMessage_Status) ProtoMessage() {}

func (x *PrinterConnectorClientMessage_Status) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrinterConnectorClientMessage_Status.ProtoReflect.Descriptor instead.
func (*PrinterConnectorClientMessage_Status) Descriptor() ([]byte, []int) {
	return file_eolymp_printer_printer_connector_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PrinterConnectorClientMessage_Status) GetStatus() Printer_Status {
	if x != nil {
		return x.Status
	}
	return Printer_UNKNOWN_STATUS
}

// Hello message confirms that the printer is authenticated and can start communicating with the server
type PrinterConnectorServerMessage_Hello struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Comment       string                 `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	Printer       *Printer               `protobuf:"bytes,2,opt,name=printer,proto3" json:"printer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrinterConnectorServerMessage_Hello) Reset() {
	*x = PrinterConnectorServerMessage_Hello{}
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrinterConnectorServerMessage_Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrinterConnectorServerMessage_Hello) ProtoMessage() {}

func (x *PrinterConnectorServerMessage_Hello) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrinterConnectorServerMessage_Hello.ProtoReflect.Descriptor instead.
func (*PrinterConnectorServerMessage_Hello) Descriptor() ([]byte, []int) {
	return file_eolymp_printer_printer_connector_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PrinterConnectorServerMessage_Hello) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *PrinterConnectorServerMessage_Hello) GetPrinter() *Printer {
	if x != nil {
		return x.Printer
	}
	return nil
}

// Print message instructs the printer to print a document
// The printer should use Report message to report the status of the job
type PrinterConnectorServerMessage_Print struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Job           *Job                   `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrinterConnectorServerMessage_Print) Reset() {
	*x = PrinterConnectorServerMessage_Print{}
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrinterConnectorServerMessage_Print) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrinterConnectorServerMessage_Print) ProtoMessage() {}

func (x *PrinterConnectorServerMessage_Print) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_printer_printer_connector_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrinterConnectorServerMessage_Print.ProtoReflect.Descriptor instead.
func (*PrinterConnectorServerMessage_Print) Descriptor() ([]byte, []int) {
	return file_eolymp_printer_printer_connector_proto_rawDescGZIP(), []int{1, 1}
}

func (x *PrinterConnectorServerMessage_Print) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

var File_eolymp_printer_printer_connector_proto protoreflect.FileDescriptor

var file_eolymp_printer_printer_connector_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70,
	0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6a,
	0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x02, 0x0a, 0x1d, 0x50, 0x72, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4e, 0x0a, 0x06, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x4c, 0x0a, 0x06, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x40, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xca, 0x02, 0x0a, 0x1d, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x4b, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x6e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x1a, 0x54, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x07,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x2e, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4a,
	0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x81, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x6d, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x12, 0x2d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x2d, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x3b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eolymp_printer_printer_connector_proto_rawDescOnce sync.Once
	file_eolymp_printer_printer_connector_proto_rawDescData = file_eolymp_printer_printer_connector_proto_rawDesc
)

func file_eolymp_printer_printer_connector_proto_rawDescGZIP() []byte {
	file_eolymp_printer_printer_connector_proto_rawDescOnce.Do(func() {
		file_eolymp_printer_printer_connector_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_printer_printer_connector_proto_rawDescData)
	})
	return file_eolymp_printer_printer_connector_proto_rawDescData
}

var file_eolymp_printer_printer_connector_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eolymp_printer_printer_connector_proto_goTypes = []any{
	(*PrinterConnectorClientMessage)(nil),        // 0: eolymp.printer.PrinterConnectorClientMessage
	(*PrinterConnectorServerMessage)(nil),        // 1: eolymp.printer.PrinterConnectorServerMessage
	(*PrinterConnectorClientMessage_Report)(nil), // 2: eolymp.printer.PrinterConnectorClientMessage.Report
	(*PrinterConnectorClientMessage_Status)(nil), // 3: eolymp.printer.PrinterConnectorClientMessage.Status
	(*PrinterConnectorServerMessage_Hello)(nil),  // 4: eolymp.printer.PrinterConnectorServerMessage.Hello
	(*PrinterConnectorServerMessage_Print)(nil),  // 5: eolymp.printer.PrinterConnectorServerMessage.Print
	(Job_Status)(0),     // 6: eolymp.printer.Job.Status
	(Printer_Status)(0), // 7: eolymp.printer.Printer.Status
	(*Printer)(nil),     // 8: eolymp.printer.Printer
	(*Job)(nil),         // 9: eolymp.printer.Job
}
var file_eolymp_printer_printer_connector_proto_depIdxs = []int32{
	3, // 0: eolymp.printer.PrinterConnectorClientMessage.status:type_name -> eolymp.printer.PrinterConnectorClientMessage.Status
	2, // 1: eolymp.printer.PrinterConnectorClientMessage.report:type_name -> eolymp.printer.PrinterConnectorClientMessage.Report
	4, // 2: eolymp.printer.PrinterConnectorServerMessage.hello:type_name -> eolymp.printer.PrinterConnectorServerMessage.Hello
	5, // 3: eolymp.printer.PrinterConnectorServerMessage.print:type_name -> eolymp.printer.PrinterConnectorServerMessage.Print
	6, // 4: eolymp.printer.PrinterConnectorClientMessage.Report.status:type_name -> eolymp.printer.Job.Status
	7, // 5: eolymp.printer.PrinterConnectorClientMessage.Status.status:type_name -> eolymp.printer.Printer.Status
	8, // 6: eolymp.printer.PrinterConnectorServerMessage.Hello.printer:type_name -> eolymp.printer.Printer
	9, // 7: eolymp.printer.PrinterConnectorServerMessage.Print.job:type_name -> eolymp.printer.Job
	0, // 8: eolymp.printer.PrinterConnector.Connect:input_type -> eolymp.printer.PrinterConnectorClientMessage
	1, // 9: eolymp.printer.PrinterConnector.Connect:output_type -> eolymp.printer.PrinterConnectorServerMessage
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_eolymp_printer_printer_connector_proto_init() }
func file_eolymp_printer_printer_connector_proto_init() {
	if File_eolymp_printer_printer_connector_proto != nil {
		return
	}
	file_eolymp_printer_printer_proto_init()
	file_eolymp_printer_printer_job_proto_init()
	file_eolymp_printer_printer_connector_proto_msgTypes[0].OneofWrappers = []any{
		(*PrinterConnectorClientMessage_Status_)(nil),
		(*PrinterConnectorClientMessage_Report_)(nil),
	}
	file_eolymp_printer_printer_connector_proto_msgTypes[1].OneofWrappers = []any{
		(*PrinterConnectorServerMessage_Hello_)(nil),
		(*PrinterConnectorServerMessage_Print_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_printer_printer_connector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eolymp_printer_printer_connector_proto_goTypes,
		DependencyIndexes: file_eolymp_printer_printer_connector_proto_depIdxs,
		MessageInfos:      file_eolymp_printer_printer_connector_proto_msgTypes,
	}.Build()
	File_eolymp_printer_printer_connector_proto = out.File
	file_eolymp_printer_printer_connector_proto_rawDesc = nil
	file_eolymp_printer_printer_connector_proto_goTypes = nil
	file_eolymp_printer_printer_connector_proto_depIdxs = nil
}
