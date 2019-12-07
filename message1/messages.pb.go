// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comsignment.proto

package message1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OrderRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	TimeStamp            int64    `protobuf:"varint,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (m *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(m, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderRequest) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

type OrderInfo struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	OrderName            string   `protobuf:"bytes,2,opt,name=OrderName,proto3" json:"OrderName,omitempty"`
	OrderStaus           string   `protobuf:"bytes,3,opt,name=OrderStaus,proto3" json:"OrderStaus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderInfo) Reset()         { *m = OrderInfo{} }
func (m *OrderInfo) String() string { return proto.CompactTextString(m) }
func (*OrderInfo) ProtoMessage()    {}
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *OrderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderInfo.Unmarshal(m, b)
}
func (m *OrderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderInfo.Marshal(b, m, deterministic)
}
func (m *OrderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderInfo.Merge(m, src)
}
func (m *OrderInfo) XXX_Size() int {
	return xxx_messageInfo_OrderInfo.Size(m)
}
func (m *OrderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderInfo proto.InternalMessageInfo

func (m *OrderInfo) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderInfo) GetOrderName() string {
	if m != nil {
		return m.OrderName
	}
	return ""
}

func (m *OrderInfo) GetOrderStaus() string {
	if m != nil {
		return m.OrderStaus
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderRequest)(nil), "message1.OrderRequest")
	proto.RegisterType((*OrderInfo)(nil), "message1.OrderInfo")
}

func init() { proto.RegisterFile("comsignment.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf2, 0x0d, 0x95,
	0xdc, 0xb8, 0x78, 0xfc, 0x8b, 0x52, 0x52, 0x8b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0x24, 0xb8, 0xd8, 0xf3, 0x41, 0x7c, 0xcf, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18,
	0x57, 0x48, 0x86, 0x8b, 0xb3, 0x24, 0x33, 0x37, 0x35, 0xb8, 0x24, 0x31, 0xb7, 0x40, 0x82, 0x49,
	0x81, 0x51, 0x83, 0x39, 0x08, 0x21, 0xa0, 0x94, 0xcc, 0xc5, 0x09, 0x36, 0xc7, 0x33, 0x2f, 0x2d,
	0x1f, 0x64, 0x88, 0x3f, 0xaa, 0x21, 0xfe, 0x08, 0x43, 0xc0, 0x4c, 0xbf, 0xc4, 0xdc, 0x54, 0xb0,
	0x21, 0x9c, 0x41, 0x08, 0x01, 0x21, 0x39, 0x2e, 0x2e, 0x30, 0x27, 0xb8, 0x24, 0xb1, 0xb4, 0x58,
	0x82, 0x19, 0x2c, 0x8d, 0x24, 0x62, 0xe4, 0x0d, 0x75, 0x6c, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72,
	0xaa, 0x90, 0x35, 0x17, 0x8f, 0x7b, 0x6a, 0x09, 0xc2, 0x5e, 0x31, 0x3d, 0x98, 0xbf, 0xf4, 0x90,
	0x3d, 0x25, 0x25, 0x8c, 0x26, 0x0e, 0x52, 0x9c, 0xc4, 0x06, 0x0e, 0x0a, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x68, 0x23, 0x33, 0xd6, 0x1c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	GetOrderInfo(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfo, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetOrderInfo(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfo, error) {
	out := new(OrderInfo)
	err := c.cc.Invoke(ctx, "/message1.OrderService/GetOrderInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	GetOrderInfo(context.Context, *OrderRequest) (*OrderInfo, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) GetOrderInfo(ctx context.Context, req *OrderRequest) (*OrderInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderInfo not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_GetOrderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message1.OrderService/GetOrderInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderInfo(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message1.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderInfo",
			Handler:    _OrderService_GetOrderInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comsignment.proto",
}
