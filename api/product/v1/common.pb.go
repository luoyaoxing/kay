// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package v1

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RequestHeader struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Chn                  string   `protobuf:"bytes,2,opt,name=chn,proto3" json:"chn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *RequestHeader) GetChn() string {
	if m != nil {
		return m.Chn
	}
	return ""
}

type ResponseHeader struct {
	Errno                int32    `protobuf:"varint,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetErrno() int32 {
	if m != nil {
		return m.Errno
	}
	return 0
}

func (m *ResponseHeader) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestHeader)(nil), "api.product.v1.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "api.product.v1.ResponseHeader")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0x2c, 0xc8, 0x04, 0x31, 0x53,
	0x4a, 0x93, 0x4b, 0xf4, 0xca, 0x0c, 0x95, 0x2c, 0xb9, 0x78, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x3c, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x84, 0xc4, 0xb8, 0xd8, 0x8a, 0xf3, 0x4b, 0x8b, 0x92,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x01, 0x2e, 0xe6, 0xe4, 0x8c,
	0x3c, 0x09, 0x26, 0xb0, 0x20, 0x88, 0xa9, 0x64, 0xc7, 0xc5, 0x17, 0x94, 0x5a, 0x5c, 0x90, 0x9f,
	0x57, 0x9c, 0x0a, 0xd5, 0x2b, 0xc2, 0xc5, 0x9a, 0x5a, 0x54, 0x94, 0x97, 0x0f, 0xd6, 0xca, 0x1a,
	0x04, 0xe1, 0x80, 0x4c, 0x4c, 0x2d, 0x2a, 0xca, 0x2d, 0x4e, 0x87, 0x6a, 0x86, 0xf2, 0x9c, 0x34,
	0xb9, 0xd0, 0x1c, 0x13, 0xc0, 0x18, 0x25, 0x9a, 0x9d, 0x58, 0xa9, 0x9f, 0x58, 0x90, 0xa9, 0x0f,
	0x15, 0xd5, 0x2f, 0x33, 0xb4, 0x2e, 0x33, 0x4c, 0x62, 0x03, 0x3b, 0xde, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x16, 0x55, 0xf0, 0x90, 0xcc, 0x00, 0x00, 0x00,
}
