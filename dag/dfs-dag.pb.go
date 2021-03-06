// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfs-dag.proto

package dag

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Object struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Links                []*Link  `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_e20a912c378d1c51, []int{0}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Object) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

type Link struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size                 uint64   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_e20a912c378d1c51, []int{1}
}

func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Link) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Link) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*Object)(nil), "dag.Object")
	proto.RegisterType((*Link)(nil), "dag.Link")
}

func init() { proto.RegisterFile("dfs-dag.proto", fileDescriptor_e20a912c378d1c51) }

var fileDescriptor_e20a912c378d1c51 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0x31, 0x0f, 0x82, 0x30,
	0x18, 0x05, 0x53, 0x40, 0x12, 0xaa, 0x2e, 0x9d, 0xba, 0xd9, 0x30, 0x75, 0x01, 0x12, 0x9d, 0x1d,
	0x24, 0x8e, 0x26, 0x24, 0x1d, 0xdd, 0x3e, 0x28, 0x94, 0x8a, 0x50, 0x63, 0xeb, 0xe2, 0xaf, 0x37,
	0x2d, 0x71, 0xbb, 0x77, 0xc9, 0x25, 0x0f, 0xef, 0xe5, 0x60, 0x0b, 0x09, 0xaa, 0x7c, 0xbd, 0x8d,
	0x33, 0x24, 0x96, 0xa0, 0xf2, 0x33, 0x4e, 0x9b, 0xf6, 0xd1, 0x77, 0x8e, 0x10, 0x9c, 0x48, 0x70,
	0x40, 0x11, 0x43, 0x7c, 0x27, 0x02, 0x93, 0x03, 0xde, 0x3c, 0xf5, 0x32, 0x59, 0x1a, 0xb1, 0x98,
	0x6f, 0x8f, 0x59, 0xe9, 0xeb, 0x9b, 0x5e, 0x26, 0xb1, 0xfa, 0xbc, 0xc6, 0x89, 0x9f, 0x3e, 0x1e,
	0xc1, 0x8e, 0xff, 0xd8, 0xb3, 0x77, 0x0b, 0xcc, 0x3d, 0x8d, 0x18, 0xe2, 0x99, 0x08, 0xec, 0x9d,
	0xd5, 0xdf, 0x9e, 0xc6, 0x0c, 0xf1, 0x44, 0x04, 0xae, 0xf3, 0x3b, 0x53, 0xda, 0x8d, 0x9f, 0xb6,
	0xec, 0xcc, 0x5c, 0xc9, 0xc1, 0x5e, 0x2f, 0x4d, 0xa5, 0x4c, 0xb1, 0xbe, 0x35, 0x95, 0x04, 0xd5,
	0xa6, 0xe1, 0xf2, 0xe9, 0x17, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x4f, 0xca, 0xaa, 0xc3, 0x00, 0x00,
	0x00,
}
