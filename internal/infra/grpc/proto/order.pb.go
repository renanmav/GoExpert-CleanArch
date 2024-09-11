// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: internal/infra/grpc/proto/order.proto

package proto

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price      float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax        float32 `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice float32 `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetTax() float32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *Order) GetFinalPrice() float32 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price float32 `protobuf:"fixed32,1,opt,name=price,proto3" json:"price,omitempty"`
	Tax   float32 `protobuf:"fixed32,2,opt,name=tax,proto3" json:"tax,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderRequest) GetTax() float32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_proto_order_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type RealAllOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RealAllOrdersRequest) Reset() {
	*x = RealAllOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RealAllOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RealAllOrdersRequest) ProtoMessage() {}

func (x *RealAllOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RealAllOrdersRequest.ProtoReflect.Descriptor instead.
func (*RealAllOrdersRequest) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_proto_order_proto_rawDescGZIP(), []int{3}
}

type RealAllOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *RealAllOrdersResponse) Reset() {
	*x = RealAllOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RealAllOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RealAllOrdersResponse) ProtoMessage() {}

func (x *RealAllOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RealAllOrdersResponse.ProtoReflect.Descriptor instead.
func (*RealAllOrdersResponse) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_proto_order_proto_rawDescGZIP(), []int{4}
}

func (x *RealAllOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type ReadOrderByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadOrderByIdRequest) Reset() {
	*x = ReadOrderByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadOrderByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadOrderByIdRequest) ProtoMessage() {}

func (x *ReadOrderByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadOrderByIdRequest.ProtoReflect.Descriptor instead.
func (*ReadOrderByIdRequest) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_proto_order_proto_rawDescGZIP(), []int{5}
}

func (x *ReadOrderByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadOrderByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *ReadOrderByIdResponse) Reset() {
	*x = ReadOrderByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadOrderByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadOrderByIdResponse) ProtoMessage() {}

func (x *ReadOrderByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_proto_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadOrderByIdResponse.ProtoReflect.Descriptor instead.
func (*ReadOrderByIdResponse) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_proto_order_proto_rawDescGZIP(), []int{6}
}

func (x *ReadOrderByIdResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_internal_infra_grpc_proto_order_proto protoreflect.FileDescriptor

var file_internal_infra_grpc_proto_order_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x74, 0x61, 0x78, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x3c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x74, 0x61, 0x78, 0x22, 0x39,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x61,
	0x6c, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3d, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x6c, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x22, 0x26, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x64,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0xec, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d,
	0x52, 0x65, 0x61, 0x6c, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_infra_grpc_proto_order_proto_rawDescOnce sync.Once
	file_internal_infra_grpc_proto_order_proto_rawDescData = file_internal_infra_grpc_proto_order_proto_rawDesc
)

func file_internal_infra_grpc_proto_order_proto_rawDescGZIP() []byte {
	file_internal_infra_grpc_proto_order_proto_rawDescOnce.Do(func() {
		file_internal_infra_grpc_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_infra_grpc_proto_order_proto_rawDescData)
	})
	return file_internal_infra_grpc_proto_order_proto_rawDescData
}

var file_internal_infra_grpc_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_infra_grpc_proto_order_proto_goTypes = []any{
	(*Order)(nil),                 // 0: proto.Order
	(*CreateOrderRequest)(nil),    // 1: proto.CreateOrderRequest
	(*CreateOrderResponse)(nil),   // 2: proto.CreateOrderResponse
	(*RealAllOrdersRequest)(nil),  // 3: proto.RealAllOrdersRequest
	(*RealAllOrdersResponse)(nil), // 4: proto.RealAllOrdersResponse
	(*ReadOrderByIdRequest)(nil),  // 5: proto.ReadOrderByIdRequest
	(*ReadOrderByIdResponse)(nil), // 6: proto.ReadOrderByIdResponse
}
var file_internal_infra_grpc_proto_order_proto_depIdxs = []int32{
	0, // 0: proto.CreateOrderResponse.order:type_name -> proto.Order
	0, // 1: proto.RealAllOrdersResponse.orders:type_name -> proto.Order
	0, // 2: proto.ReadOrderByIdResponse.order:type_name -> proto.Order
	1, // 3: proto.OrderService.CreateOrder:input_type -> proto.CreateOrderRequest
	3, // 4: proto.OrderService.RealAllOrders:input_type -> proto.RealAllOrdersRequest
	5, // 5: proto.OrderService.ReadOrderById:input_type -> proto.ReadOrderByIdRequest
	2, // 6: proto.OrderService.CreateOrder:output_type -> proto.CreateOrderResponse
	4, // 7: proto.OrderService.RealAllOrders:output_type -> proto.RealAllOrdersResponse
	6, // 8: proto.OrderService.ReadOrderById:output_type -> proto.ReadOrderByIdResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_infra_grpc_proto_order_proto_init() }
func file_internal_infra_grpc_proto_order_proto_init() {
	if File_internal_infra_grpc_proto_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_infra_grpc_proto_order_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Order); i {
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
		file_internal_infra_grpc_proto_order_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOrderRequest); i {
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
		file_internal_infra_grpc_proto_order_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOrderResponse); i {
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
		file_internal_infra_grpc_proto_order_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RealAllOrdersRequest); i {
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
		file_internal_infra_grpc_proto_order_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RealAllOrdersResponse); i {
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
		file_internal_infra_grpc_proto_order_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ReadOrderByIdRequest); i {
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
		file_internal_infra_grpc_proto_order_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ReadOrderByIdResponse); i {
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
			RawDescriptor: file_internal_infra_grpc_proto_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_infra_grpc_proto_order_proto_goTypes,
		DependencyIndexes: file_internal_infra_grpc_proto_order_proto_depIdxs,
		MessageInfos:      file_internal_infra_grpc_proto_order_proto_msgTypes,
	}.Build()
	File_internal_infra_grpc_proto_order_proto = out.File
	file_internal_infra_grpc_proto_order_proto_rawDesc = nil
	file_internal_infra_grpc_proto_order_proto_goTypes = nil
	file_internal_infra_grpc_proto_order_proto_depIdxs = nil
}
