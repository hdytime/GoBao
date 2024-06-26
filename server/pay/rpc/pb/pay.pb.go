// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.4
// source: pay.proto

package pb

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

type OrderPaymentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  int64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	OrderSn string `protobuf:"bytes,2,opt,name=OrderSn,proto3" json:"OrderSn,omitempty"`
}

func (x *OrderPaymentReq) Reset() {
	*x = OrderPaymentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPaymentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPaymentReq) ProtoMessage() {}

func (x *OrderPaymentReq) ProtoReflect() protoreflect.Message {
	mi := &file_pay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPaymentReq.ProtoReflect.Descriptor instead.
func (*OrderPaymentReq) Descriptor() ([]byte, []int) {
	return file_pay_proto_rawDescGZIP(), []int{0}
}

func (x *OrderPaymentReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *OrderPaymentReq) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

type OrderPaymentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayTotalPrice float64 `protobuf:"fixed64,1,opt,name=PayTotalPrice,proto3" json:"PayTotalPrice,omitempty"`
	PaySn         string  `protobuf:"bytes,2,opt,name=PaySn,proto3" json:"PaySn,omitempty"`
}

func (x *OrderPaymentResp) Reset() {
	*x = OrderPaymentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPaymentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPaymentResp) ProtoMessage() {}

func (x *OrderPaymentResp) ProtoReflect() protoreflect.Message {
	mi := &file_pay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPaymentResp.ProtoReflect.Descriptor instead.
func (*OrderPaymentResp) Descriptor() ([]byte, []int) {
	return file_pay_proto_rawDescGZIP(), []int{1}
}

func (x *OrderPaymentResp) GetPayTotalPrice() float64 {
	if x != nil {
		return x.PayTotalPrice
	}
	return 0
}

func (x *OrderPaymentResp) GetPaySn() string {
	if x != nil {
		return x.PaySn
	}
	return ""
}

type GetPaymentDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderSn string `protobuf:"bytes,1,opt,name=OrderSn,proto3" json:"OrderSn,omitempty"`
}

func (x *GetPaymentDetailReq) Reset() {
	*x = GetPaymentDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentDetailReq) ProtoMessage() {}

func (x *GetPaymentDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_pay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentDetailReq.ProtoReflect.Descriptor instead.
func (*GetPaymentDetailReq) Descriptor() ([]byte, []int) {
	return file_pay_proto_rawDescGZIP(), []int{2}
}

func (x *GetPaymentDetailReq) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

type GetPaymentDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             int64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PaySn          string  `protobuf:"bytes,2,opt,name=PaySn,proto3" json:"PaySn,omitempty"`
	OrderSn        string  `protobuf:"bytes,3,opt,name=OrderSn,proto3" json:"OrderSn,omitempty"`
	UserID         int64   `protobuf:"varint,4,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TradeState     string  `protobuf:"bytes,5,opt,name=TradeState,proto3" json:"TradeState,omitempty"`
	PayTotal       float64 `protobuf:"fixed64,6,opt,name=PayTotal,proto3" json:"PayTotal,omitempty"`
	TransactionID  string  `protobuf:"bytes,7,opt,name=TransactionID,proto3" json:"TransactionID,omitempty"`
	TradeStateDesc string  `protobuf:"bytes,8,opt,name=TradeStateDesc,proto3" json:"TradeStateDesc,omitempty"`
	PayStatus      int64   `protobuf:"varint,9,opt,name=PayStatus,proto3" json:"PayStatus,omitempty"`
	PayTime        string  `protobuf:"bytes,10,opt,name=PayTime,proto3" json:"PayTime,omitempty"`
}

func (x *GetPaymentDetailResp) Reset() {
	*x = GetPaymentDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentDetailResp) ProtoMessage() {}

func (x *GetPaymentDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_pay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentDetailResp.ProtoReflect.Descriptor instead.
func (*GetPaymentDetailResp) Descriptor() ([]byte, []int) {
	return file_pay_proto_rawDescGZIP(), []int{3}
}

func (x *GetPaymentDetailResp) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetPaymentDetailResp) GetPaySn() string {
	if x != nil {
		return x.PaySn
	}
	return ""
}

func (x *GetPaymentDetailResp) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *GetPaymentDetailResp) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetPaymentDetailResp) GetTradeState() string {
	if x != nil {
		return x.TradeState
	}
	return ""
}

func (x *GetPaymentDetailResp) GetPayTotal() float64 {
	if x != nil {
		return x.PayTotal
	}
	return 0
}

func (x *GetPaymentDetailResp) GetTransactionID() string {
	if x != nil {
		return x.TransactionID
	}
	return ""
}

func (x *GetPaymentDetailResp) GetTradeStateDesc() string {
	if x != nil {
		return x.TradeStateDesc
	}
	return ""
}

func (x *GetPaymentDetailResp) GetPayStatus() int64 {
	if x != nil {
		return x.PayStatus
	}
	return 0
}

func (x *GetPaymentDetailResp) GetPayTime() string {
	if x != nil {
		return x.PayTime
	}
	return ""
}

var File_pay_proto protoreflect.FileDescriptor

var file_pay_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x61, 0x79,
	0x22, 0x43, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x6e, 0x22, 0x4e, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x50, 0x61, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x50, 0x61, 0x79, 0x53, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x50, 0x61, 0x79, 0x53, 0x6e, 0x22, 0x2f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x22, 0xb0, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x50, 0x61, 0x79, 0x53, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x50, 0x61, 0x79, 0x53, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x50, 0x61, 0x79, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x50, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x50, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x8b, 0x01, 0x0a, 0x03, 0x70, 0x61,
	0x79, 0x12, 0x3b, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x47,
	0x0a, 0x10, 0x67, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70,
	0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pay_proto_rawDescOnce sync.Once
	file_pay_proto_rawDescData = file_pay_proto_rawDesc
)

func file_pay_proto_rawDescGZIP() []byte {
	file_pay_proto_rawDescOnce.Do(func() {
		file_pay_proto_rawDescData = protoimpl.X.CompressGZIP(file_pay_proto_rawDescData)
	})
	return file_pay_proto_rawDescData
}

var file_pay_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pay_proto_goTypes = []interface{}{
	(*OrderPaymentReq)(nil),      // 0: pay.OrderPaymentReq
	(*OrderPaymentResp)(nil),     // 1: pay.OrderPaymentResp
	(*GetPaymentDetailReq)(nil),  // 2: pay.GetPaymentDetailReq
	(*GetPaymentDetailResp)(nil), // 3: pay.GetPaymentDetailResp
}
var file_pay_proto_depIdxs = []int32{
	0, // 0: pay.pay.orderPayment:input_type -> pay.OrderPaymentReq
	2, // 1: pay.pay.getPaymentDetail:input_type -> pay.GetPaymentDetailReq
	1, // 2: pay.pay.orderPayment:output_type -> pay.OrderPaymentResp
	3, // 3: pay.pay.getPaymentDetail:output_type -> pay.GetPaymentDetailResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pay_proto_init() }
func file_pay_proto_init() {
	if File_pay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPaymentReq); i {
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
		file_pay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPaymentResp); i {
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
		file_pay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentDetailReq); i {
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
		file_pay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentDetailResp); i {
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
			RawDescriptor: file_pay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pay_proto_goTypes,
		DependencyIndexes: file_pay_proto_depIdxs,
		MessageInfos:      file_pay_proto_msgTypes,
	}.Build()
	File_pay_proto = out.File
	file_pay_proto_rawDesc = nil
	file_pay_proto_goTypes = nil
	file_pay_proto_depIdxs = nil
}
