// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfs-network.proto

package network

import (
	context "context"
	fmt "fmt"
	dag "github.com/dfsDAO/go-dfs-dao/dag"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c55146c3e002874, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type GetResponse struct {
	Object               *dag.Object `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c55146c3e002874, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetObject() *dag.Object {
	if m != nil {
		return m.Object
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "network.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "network.GetResponse")
}

func init() { proto.RegisterFile("dfs-network.proto", fileDescriptor_6c55146c3e002874) }

var fileDescriptor_6c55146c3e002874 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x49, 0x2b, 0xd6,
	0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87,
	0x72, 0xa5, 0x78, 0x41, 0x72, 0x29, 0x89, 0xe9, 0x10, 0x71, 0x25, 0x1e, 0x2e, 0x2e, 0xf7, 0xd4,
	0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x23, 0x2e, 0x6e, 0x30, 0xaf, 0xb8, 0x20,
	0x3f, 0xaf, 0x38, 0x55, 0x48, 0x99, 0x8b, 0x2d, 0x3f, 0x29, 0x2b, 0x35, 0xb9, 0x44, 0x82, 0x51,
	0x81, 0x51, 0x83, 0xdb, 0x88, 0x5b, 0x0f, 0xa4, 0xd1, 0x1f, 0x2c, 0x14, 0x04, 0x95, 0x32, 0xb2,
	0xe6, 0x62, 0xf7, 0x83, 0x98, 0x2d, 0x64, 0xc0, 0xc5, 0xec, 0x9e, 0x5a, 0x22, 0x24, 0xac, 0x07,
	0xb3, 0x1b, 0x61, 0xb4, 0x94, 0x08, 0xaa, 0x20, 0xc4, 0x06, 0x27, 0xb5, 0x28, 0x95, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x94, 0xb4, 0x62, 0x17, 0x47, 0x7f, 0xfd,
	0xf4, 0x7c, 0x5d, 0x88, 0x23, 0xf3, 0xf5, 0xa1, 0x7a, 0x92, 0xd8, 0xc0, 0xae, 0x35, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x08, 0xde, 0x08, 0x26, 0xda, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkClient is the client API for Network service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type networkClient struct {
	cc *grpc.ClientConn
}

func NewNetworkClient(cc *grpc.ClientConn) NetworkClient {
	return &networkClient{cc}
}

func (c *networkClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/network.Network/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServer is the server API for Network service.
type NetworkServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterNetworkServer(s *grpc.Server, srv NetworkServer) {
	s.RegisterService(&_Network_serviceDesc, srv)
}

func _Network_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.Network/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Network_serviceDesc = grpc.ServiceDesc{
	ServiceName: "network.Network",
	HandlerType: (*NetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Network_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dfs-network.proto",
}
