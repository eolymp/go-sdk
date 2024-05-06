// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.21.12
// source: eolymp/commerce/invoice.proto

package commerce

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Invoice_Status int32

const (
	Invoice_UNKNOWN_STATUS Invoice_Status = 0
	Invoice_DRAFT          Invoice_Status = 1
	Invoice_OPEN           Invoice_Status = 2
	Invoice_PAID           Invoice_Status = 3
	Invoice_UNCOLLECTIBLE  Invoice_Status = 4
	Invoice_VOID           Invoice_Status = 5
)

// Enum value maps for Invoice_Status.
var (
	Invoice_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "DRAFT",
		2: "OPEN",
		3: "PAID",
		4: "UNCOLLECTIBLE",
		5: "VOID",
	}
	Invoice_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"DRAFT":          1,
		"OPEN":           2,
		"PAID":           3,
		"UNCOLLECTIBLE":  4,
		"VOID":           5,
	}
)

func (x Invoice_Status) Enum() *Invoice_Status {
	p := new(Invoice_Status)
	*p = x
	return p
}

func (x Invoice_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Invoice_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_eolymp_commerce_invoice_proto_enumTypes[0].Descriptor()
}

func (Invoice_Status) Type() protoreflect.EnumType {
	return &file_eolymp_commerce_invoice_proto_enumTypes[0]
}

func (x Invoice_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Invoice_Status.Descriptor instead.
func (Invoice_Status) EnumDescriptor() ([]byte, []int) {
	return file_eolymp_commerce_invoice_proto_rawDescGZIP(), []int{0, 0}
}

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number               string                    `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Status               Invoice_Status            `protobuf:"varint,3,opt,name=status,proto3,enum=eolymp.commerce.Invoice_Status" json:"status,omitempty"`
	CustomerId           string                    `protobuf:"bytes,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Description          string                    `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	FromInvoice          *Invoice_FromInvoice      `protobuf:"bytes,6,opt,name=from_invoice,json=fromInvoice,proto3" json:"from_invoice,omitempty"`
	CreatedAt            *timestamppb.Timestamp    `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DueAt                *timestamppb.Timestamp    `protobuf:"bytes,11,opt,name=due_at,json=dueAt,proto3" json:"due_at,omitempty"`
	HostedInvoiceUrl     string                    `protobuf:"bytes,31,opt,name=hosted_invoice_url,json=hostedInvoiceUrl,proto3" json:"hosted_invoice_url,omitempty"`
	InvoicePdfUrl        string                    `protobuf:"bytes,32,opt,name=invoice_pdf_url,json=invoicePdfUrl,proto3" json:"invoice_pdf_url,omitempty"`
	Currency             string                    `protobuf:"bytes,100,opt,name=currency,proto3" json:"currency,omitempty"`
	AmountDue            int32                     `protobuf:"varint,120,opt,name=amount_due,json=amountDue,proto3" json:"amount_due,omitempty"`
	AmountPaid           int32                     `protobuf:"varint,121,opt,name=amount_paid,json=amountPaid,proto3" json:"amount_paid,omitempty"`
	AmountRemaining      int32                     `protobuf:"varint,122,opt,name=amount_remaining,json=amountRemaining,proto3" json:"amount_remaining,omitempty"`
	Subtotal             int32                     `protobuf:"varint,130,opt,name=subtotal,proto3" json:"subtotal,omitempty"`                                                       // total of all items
	SubtotalExcludingTax int32                     `protobuf:"varint,131,opt,name=subtotal_excluding_tax,json=subtotalExcludingTax,proto3" json:"subtotal_excluding_tax,omitempty"` // total of all items without taxes
	TaxAmounts           []*Invoice_TaxAmount      `protobuf:"bytes,141,rep,name=tax_amounts,json=taxAmounts,proto3" json:"tax_amounts,omitempty"`
	Tax                  int32                     `protobuf:"varint,140,opt,name=tax,proto3" json:"tax,omitempty"`
	DiscountAmounts      []*Invoice_DiscountAmount `protobuf:"bytes,150,rep,name=discount_amounts,json=discountAmounts,proto3" json:"discount_amounts,omitempty"`
	Total                int32                     `protobuf:"varint,200,opt,name=total,proto3" json:"total,omitempty"`
	TotalExcludingTax    int32                     `protobuf:"varint,201,opt,name=total_excluding_tax,json=totalExcludingTax,proto3" json:"total_excluding_tax,omitempty"`
	Items                []*Invoice_Item           `protobuf:"bytes,999,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_invoice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_invoice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_invoice_proto_rawDescGZIP(), []int{0}
}

func (x *Invoice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Invoice) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Invoice) GetStatus() Invoice_Status {
	if x != nil {
		return x.Status
	}
	return Invoice_UNKNOWN_STATUS
}

func (x *Invoice) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Invoice) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Invoice) GetFromInvoice() *Invoice_FromInvoice {
	if x != nil {
		return x.FromInvoice
	}
	return nil
}

func (x *Invoice) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Invoice) GetDueAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DueAt
	}
	return nil
}

func (x *Invoice) GetHostedInvoiceUrl() string {
	if x != nil {
		return x.HostedInvoiceUrl
	}
	return ""
}

func (x *Invoice) GetInvoicePdfUrl() string {
	if x != nil {
		return x.InvoicePdfUrl
	}
	return ""
}

func (x *Invoice) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Invoice) GetAmountDue() int32 {
	if x != nil {
		return x.AmountDue
	}
	return 0
}

func (x *Invoice) GetAmountPaid() int32 {
	if x != nil {
		return x.AmountPaid
	}
	return 0
}

func (x *Invoice) GetAmountRemaining() int32 {
	if x != nil {
		return x.AmountRemaining
	}
	return 0
}

func (x *Invoice) GetSubtotal() int32 {
	if x != nil {
		return x.Subtotal
	}
	return 0
}

func (x *Invoice) GetSubtotalExcludingTax() int32 {
	if x != nil {
		return x.SubtotalExcludingTax
	}
	return 0
}

func (x *Invoice) GetTaxAmounts() []*Invoice_TaxAmount {
	if x != nil {
		return x.TaxAmounts
	}
	return nil
}

func (x *Invoice) GetTax() int32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *Invoice) GetDiscountAmounts() []*Invoice_DiscountAmount {
	if x != nil {
		return x.DiscountAmounts
	}
	return nil
}

func (x *Invoice) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Invoice) GetTotalExcludingTax() int32 {
	if x != nil {
		return x.TotalExcludingTax
	}
	return 0
}

func (x *Invoice) GetItems() []*Invoice_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Invoice_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description            string                    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Quantity               uint32                    `protobuf:"varint,10,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price                  *Price                    `protobuf:"bytes,20,opt,name=price,proto3" json:"price,omitempty"`
	Currency               string                    `protobuf:"bytes,100,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount                 int32                     `protobuf:"varint,101,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountExcludingTax     int32                     `protobuf:"varint,102,opt,name=amount_excluding_tax,json=amountExcludingTax,proto3" json:"amount_excluding_tax,omitempty"`
	UnitAmountExcludingTax int32                     `protobuf:"varint,103,opt,name=unit_amount_excluding_tax,json=unitAmountExcludingTax,proto3" json:"unit_amount_excluding_tax,omitempty"`
	DiscountAmounts        []*Invoice_DiscountAmount `protobuf:"bytes,104,rep,name=discount_amounts,json=discountAmounts,proto3" json:"discount_amounts,omitempty"`
	TaxAmounts             []*Invoice_TaxAmount      `protobuf:"bytes,105,rep,name=tax_amounts,json=taxAmounts,proto3" json:"tax_amounts,omitempty"`
}

func (x *Invoice_Item) Reset() {
	*x = Invoice_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_invoice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice_Item) ProtoMessage() {}

func (x *Invoice_Item) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_invoice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice_Item.ProtoReflect.Descriptor instead.
func (*Invoice_Item) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_invoice_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Invoice_Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Invoice_Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Invoice_Item) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Invoice_Item) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *Invoice_Item) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Invoice_Item) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Invoice_Item) GetAmountExcludingTax() int32 {
	if x != nil {
		return x.AmountExcludingTax
	}
	return 0
}

func (x *Invoice_Item) GetUnitAmountExcludingTax() int32 {
	if x != nil {
		return x.UnitAmountExcludingTax
	}
	return 0
}

func (x *Invoice_Item) GetDiscountAmounts() []*Invoice_DiscountAmount {
	if x != nil {
		return x.DiscountAmounts
	}
	return nil
}

func (x *Invoice_Item) GetTaxAmounts() []*Invoice_TaxAmount {
	if x != nil {
		return x.TaxAmounts
	}
	return nil
}

type Invoice_TaxAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount           int32  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Inclusive        bool   `protobuf:"varint,2,opt,name=inclusive,proto3" json:"inclusive,omitempty"`
	TaxRate          string `protobuf:"bytes,3,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
	TaxabilityReason string `protobuf:"bytes,4,opt,name=taxability_reason,json=taxabilityReason,proto3" json:"taxability_reason,omitempty"`
	TaxableAmount    int32  `protobuf:"varint,10,opt,name=taxable_amount,json=taxableAmount,proto3" json:"taxable_amount,omitempty"`
}

func (x *Invoice_TaxAmount) Reset() {
	*x = Invoice_TaxAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_invoice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice_TaxAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice_TaxAmount) ProtoMessage() {}

func (x *Invoice_TaxAmount) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_invoice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice_TaxAmount.ProtoReflect.Descriptor instead.
func (*Invoice_TaxAmount) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_invoice_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Invoice_TaxAmount) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Invoice_TaxAmount) GetInclusive() bool {
	if x != nil {
		return x.Inclusive
	}
	return false
}

func (x *Invoice_TaxAmount) GetTaxRate() string {
	if x != nil {
		return x.TaxRate
	}
	return ""
}

func (x *Invoice_TaxAmount) GetTaxabilityReason() string {
	if x != nil {
		return x.TaxabilityReason
	}
	return ""
}

func (x *Invoice_TaxAmount) GetTaxableAmount() int32 {
	if x != nil {
		return x.TaxableAmount
	}
	return 0
}

type Invoice_DiscountAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   int32  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Discount string `protobuf:"bytes,2,opt,name=discount,proto3" json:"discount,omitempty"`
}

func (x *Invoice_DiscountAmount) Reset() {
	*x = Invoice_DiscountAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_invoice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice_DiscountAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice_DiscountAmount) ProtoMessage() {}

func (x *Invoice_DiscountAmount) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_invoice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice_DiscountAmount.ProtoReflect.Descriptor instead.
func (*Invoice_DiscountAmount) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_invoice_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Invoice_DiscountAmount) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Invoice_DiscountAmount) GetDiscount() string {
	if x != nil {
		return x.Discount
	}
	return ""
}

type Invoice_FromInvoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relation string `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
	Invoice  string `protobuf:"bytes,2,opt,name=invoice,proto3" json:"invoice,omitempty"`
}

func (x *Invoice_FromInvoice) Reset() {
	*x = Invoice_FromInvoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eolymp_commerce_invoice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice_FromInvoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice_FromInvoice) ProtoMessage() {}

func (x *Invoice_FromInvoice) ProtoReflect() protoreflect.Message {
	mi := &file_eolymp_commerce_invoice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice_FromInvoice.ProtoReflect.Descriptor instead.
func (*Invoice_FromInvoice) Descriptor() ([]byte, []int) {
	return file_eolymp_commerce_invoice_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Invoice_FromInvoice) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *Invoice_FromInvoice) GetInvoice() string {
	if x != nil {
		return x.Invoice
	}
	return ""
}

var File_eolymp_commerce_invoice_proto protoreflect.FileDescriptor

var file_eolymp_commerce_invoice_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x1a, 0x1b, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98,
	0x0e, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47,
	0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x46,
	0x72, 0x6f, 0x6d, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x64, 0x75, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05,
	0x64, 0x75, 0x65, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x68, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x68, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x70,
	0x64, 0x66, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x64, 0x66, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x64, 0x75, 0x65, 0x18, 0x78, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x44, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x79, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x61, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x7a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x82,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x35, 0x0a, 0x16, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x63, 0x6c,
	0x75, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x78, 0x18, 0x83, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x14, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x78, 0x12, 0x44, 0x0a, 0x0b, 0x74, 0x61, 0x78, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x8d, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x0a, 0x74, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x11, 0x0a, 0x03,
	0x74, 0x61, 0x78, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x61, 0x78, 0x12,
	0x53, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x96, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6f, 0x6c,
	0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x12, 0x15, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0xc8, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x13, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x61, 0x78, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x78, 0x12, 0x34, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65,
	0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x1a, 0xbc, 0x03, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x61, 0x78, 0x18, 0x66, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x78, 0x12, 0x39, 0x0a,
	0x19, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x78, 0x18, 0x67, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x16, 0x75, 0x6e, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x78, 0x63, 0x6c,
	0x75, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x78, 0x12, 0x52, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x68, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x43, 0x0a, 0x0b,
	0x74, 0x61, 0x78, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x69, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x78, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0a, 0x74, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x1a, 0xb0, 0x01, 0x0a, 0x09, 0x54, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x78, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x74, 0x61, 0x78, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x61, 0x78,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x61, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x61, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x44, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x43, 0x0a, 0x0b, 0x46, 0x72,
	0x6f, 0x6d, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22,
	0x58, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x49, 0x44, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d,
	0x55, 0x4e, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x12,
	0x08, 0x0a, 0x04, 0x56, 0x4f, 0x49, 0x44, 0x10, 0x05, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eolymp_commerce_invoice_proto_rawDescOnce sync.Once
	file_eolymp_commerce_invoice_proto_rawDescData = file_eolymp_commerce_invoice_proto_rawDesc
)

func file_eolymp_commerce_invoice_proto_rawDescGZIP() []byte {
	file_eolymp_commerce_invoice_proto_rawDescOnce.Do(func() {
		file_eolymp_commerce_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(file_eolymp_commerce_invoice_proto_rawDescData)
	})
	return file_eolymp_commerce_invoice_proto_rawDescData
}

var file_eolymp_commerce_invoice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eolymp_commerce_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eolymp_commerce_invoice_proto_goTypes = []interface{}{
	(Invoice_Status)(0),            // 0: eolymp.commerce.Invoice.Status
	(*Invoice)(nil),                // 1: eolymp.commerce.Invoice
	(*Invoice_Item)(nil),           // 2: eolymp.commerce.Invoice.Item
	(*Invoice_TaxAmount)(nil),      // 3: eolymp.commerce.Invoice.TaxAmount
	(*Invoice_DiscountAmount)(nil), // 4: eolymp.commerce.Invoice.DiscountAmount
	(*Invoice_FromInvoice)(nil),    // 5: eolymp.commerce.Invoice.FromInvoice
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
	(*Price)(nil),                  // 7: eolymp.commerce.Price
}
var file_eolymp_commerce_invoice_proto_depIdxs = []int32{
	0,  // 0: eolymp.commerce.Invoice.status:type_name -> eolymp.commerce.Invoice.Status
	5,  // 1: eolymp.commerce.Invoice.from_invoice:type_name -> eolymp.commerce.Invoice.FromInvoice
	6,  // 2: eolymp.commerce.Invoice.created_at:type_name -> google.protobuf.Timestamp
	6,  // 3: eolymp.commerce.Invoice.due_at:type_name -> google.protobuf.Timestamp
	3,  // 4: eolymp.commerce.Invoice.tax_amounts:type_name -> eolymp.commerce.Invoice.TaxAmount
	4,  // 5: eolymp.commerce.Invoice.discount_amounts:type_name -> eolymp.commerce.Invoice.DiscountAmount
	2,  // 6: eolymp.commerce.Invoice.items:type_name -> eolymp.commerce.Invoice.Item
	7,  // 7: eolymp.commerce.Invoice.Item.price:type_name -> eolymp.commerce.Price
	4,  // 8: eolymp.commerce.Invoice.Item.discount_amounts:type_name -> eolymp.commerce.Invoice.DiscountAmount
	3,  // 9: eolymp.commerce.Invoice.Item.tax_amounts:type_name -> eolymp.commerce.Invoice.TaxAmount
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_eolymp_commerce_invoice_proto_init() }
func file_eolymp_commerce_invoice_proto_init() {
	if File_eolymp_commerce_invoice_proto != nil {
		return
	}
	file_eolymp_commerce_price_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eolymp_commerce_invoice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_commerce_invoice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice_Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_commerce_invoice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice_TaxAmount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_commerce_invoice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice_DiscountAmount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eolymp_commerce_invoice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice_FromInvoice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eolymp_commerce_invoice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eolymp_commerce_invoice_proto_goTypes,
		DependencyIndexes: file_eolymp_commerce_invoice_proto_depIdxs,
		EnumInfos:         file_eolymp_commerce_invoice_proto_enumTypes,
		MessageInfos:      file_eolymp_commerce_invoice_proto_msgTypes,
	}.Build()
	File_eolymp_commerce_invoice_proto = out.File
	file_eolymp_commerce_invoice_proto_rawDesc = nil
	file_eolymp_commerce_invoice_proto_goTypes = nil
	file_eolymp_commerce_invoice_proto_depIdxs = nil
}
