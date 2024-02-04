// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/product.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfd9ed8e6077aee0, []int{0}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type ProductRequest struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                uint32               `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Count                uint32               `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProductRequest) Reset()         { *m = ProductRequest{} }
func (m *ProductRequest) String() string { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()    {}
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfd9ed8e6077aee0, []int{1}
}

func (m *ProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRequest.Unmarshal(m, b)
}
func (m *ProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRequest.Marshal(b, m, deterministic)
}
func (m *ProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRequest.Merge(m, src)
}
func (m *ProductRequest) XXX_Size() int {
	return xxx_messageInfo_ProductRequest.Size(m)
}
func (m *ProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRequest proto.InternalMessageInfo

func (m *ProductRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductRequest) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductRequest) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ProductRequest) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type ID struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfd9ed8e6077aee0, []int{2}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*EmptyResponse)(nil), "pb.EmptyResponse")
	proto.RegisterType((*ProductRequest)(nil), "pb.ProductRequest")
	proto.RegisterType((*ID)(nil), "pb.ID")
}

func init() {
	proto.RegisterFile("proto/product.proto", fileDescriptor_cfd9ed8e6077aee0)
}

var fileDescriptor_cfd9ed8e6077aee0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x71, 0xda, 0x22, 0xf5, 0x50, 0x8a, 0x30, 0x1d, 0xa2, 0x2c, 0x44, 0x99, 0x32, 0x39,
	0x52, 0x59, 0x60, 0x04, 0xca, 0x50, 0x26, 0x64, 0x98, 0x58, 0x50, 0x9c, 0x1e, 0x55, 0x25, 0x12,
	0x1f, 0xce, 0x05, 0x89, 0x67, 0xe0, 0xa5, 0x51, 0xec, 0x16, 0x29, 0x8b, 0xe5, 0xff, 0xd3, 0xd9,
	0xf7, 0xfd, 0x70, 0x49, 0xce, 0xb2, 0x2d, 0xc9, 0xd9, 0x6d, 0x5f, 0xb3, 0xf2, 0x49, 0x46, 0x64,
	0xd2, 0xab, 0x9d, 0xb5, 0xbb, 0x4f, 0x2c, 0x3d, 0x31, 0xfd, 0x47, 0xc9, 0xfb, 0x06, 0x3b, 0xae,
	0x1a, 0x0a, 0x43, 0xf9, 0x39, 0xc4, 0x8f, 0x0d, 0xf1, 0x8f, 0xc6, 0x8e, 0x6c, 0xdb, 0x61, 0xfe,
	0x2b, 0x60, 0xf1, 0x1c, 0xfe, 0xd1, 0xf8, 0xd5, 0x63, 0xc7, 0x52, 0xc2, 0xb4, 0xad, 0x1a, 0x4c,
	0x44, 0x26, 0x8a, 0xb9, 0xf6, 0x77, 0xb9, 0x84, 0x19, 0xb9, 0x7d, 0x8d, 0x49, 0x94, 0x89, 0x22,
	0xd6, 0x21, 0x0c, 0xb4, 0xb6, 0x7d, 0xcb, 0xc9, 0x24, 0x50, 0x1f, 0xe4, 0x2d, 0x40, 0xed, 0xb0,
	0x62, 0xdc, 0xbe, 0x57, 0x9c, 0x4c, 0x33, 0x51, 0x9c, 0xad, 0x52, 0x15, 0xcc, 0xd4, 0xd1, 0x4c,
	0xbd, 0x1e, 0xcd, 0xf4, 0xfc, 0x30, 0x7d, 0xc7, 0xf9, 0x12, 0xa2, 0xcd, 0x5a, 0x2e, 0x86, 0xd3,
	0xaf, 0x8f, 0x75, 0xb4, 0x59, 0xaf, 0x9e, 0xfe, 0x15, 0x5f, 0xd0, 0x7d, 0x0f, 0x8b, 0x6f, 0x20,
	0x7e, 0xf0, 0x8f, 0x0e, 0x5c, 0x4a, 0x45, 0x46, 0x8d, 0x7b, 0xa4, 0x17, 0x03, 0x1b, 0xb7, 0x3d,
	0xb9, 0x9f, 0xbd, 0x4d, 0x4a, 0x32, 0xe6, 0xd4, 0x7b, 0x5c, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x59, 0xed, 0x9e, 0x78, 0x4a, 0x01, 0x00, 0x00,
}