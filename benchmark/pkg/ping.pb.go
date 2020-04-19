// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/ping.proto

package pkg

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PingRequest struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62236fa99891cde5, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PingResponse struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62236fa99891cde5, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "sms.core.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "sms.core.PingResponse")
}

func init() { proto.RegisterFile("pkg/ping.proto", fileDescriptor_62236fa99891cde5) }

var fileDescriptor_62236fa99891cde5 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4b, 0xc6, 0x30,
	0x10, 0xc5, 0x89, 0x8a, 0xda, 0xfb, 0xc4, 0x21, 0xa0, 0x14, 0x71, 0x28, 0x9d, 0x0a, 0x42, 0x04,
	0xc5, 0x55, 0xb0, 0x8a, 0xae, 0xa5, 0x6e, 0x2e, 0x52, 0xd3, 0x23, 0x0d, 0x9a, 0xe4, 0xcc, 0x45,
	0xff, 0x7e, 0x49, 0x4b, 0x51, 0x1c, 0xbe, 0xf1, 0xde, 0xfd, 0xb8, 0xf7, 0xee, 0xc1, 0x31, 0xbd,
	0x9b, 0x4b, 0xb2, 0xde, 0x28, 0x8a, 0x21, 0x05, 0x79, 0xc8, 0x8e, 0x95, 0x0e, 0x11, 0xeb, 0x0b,
	0xd8, 0x74, 0xd6, 0x9b, 0x1e, 0x3f, 0xbf, 0x90, 0x93, 0x3c, 0x87, 0x22, 0x59, 0x87, 0x9c, 0x06,
	0x47, 0xa5, 0xa8, 0x44, 0xb3, 0xdb, 0xff, 0x0a, 0xf5, 0x23, 0x1c, 0x2d, 0x30, 0x53, 0xf0, 0x8c,
	0xdb, 0x69, 0x59, 0xc2, 0x81, 0x43, 0xe6, 0xc1, 0x60, 0xb9, 0x53, 0x89, 0xa6, 0xe8, 0xd7, 0xf1,
	0xea, 0x01, 0x36, 0xf7, 0x21, 0xe2, 0x33, 0xc6, 0x6f, 0xab, 0x51, 0xde, 0xc0, 0x5e, 0x3e, 0x2b,
	0x4f, 0xd4, 0x1a, 0x4b, 0xfd, 0xc9, 0x74, 0x76, 0xfa, 0x5f, 0x5e, 0xdc, 0xdb, 0x5b, 0xa8, 0x75,
	0x70, 0x2a, 0x4d, 0xd6, 0x4f, 0xe3, 0xa0, 0x98, 0x62, 0x7e, 0xd0, 0x44, 0xd2, 0x0b, 0xec, 0xc2,
	0x88, 0x1f, 0xed, 0xec, 0xf4, 0x14, 0x49, 0xdf, 0x91, 0xed, 0xc4, 0x4b, 0x91, 0x57, 0xaf, 0x19,
	0x7a, 0xdb, 0x9f, 0xbb, 0xb8, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x08, 0xbc, 0xbc, 0x1d,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoreServiceClient is the client API for CoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type coreServiceClient struct {
	cc *grpc.ClientConn
}

func NewCoreServiceClient(cc *grpc.ClientConn) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/sms.core.CoreService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServiceServer is the server API for CoreService service.
type CoreServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// UnimplementedCoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCoreServiceServer struct {
}

func (*UnimplementedCoreServiceServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterCoreServiceServer(s *grpc.Server, srv CoreServiceServer) {
	s.RegisterService(&_CoreService_serviceDesc, srv)
}

func _CoreService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.core.CoreService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sms.core.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CoreService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/ping.proto",
}
