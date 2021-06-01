// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kvstore.proto

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

type PutKvstoreRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	K                    string         `protobuf:"bytes,2,opt,name=k,proto3" json:"k,omitempty"`
	V                    string         `protobuf:"bytes,3,opt,name=v,proto3" json:"v,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PutKvstoreRequest) Reset()         { *m = PutKvstoreRequest{} }
func (m *PutKvstoreRequest) String() string { return proto.CompactTextString(m) }
func (*PutKvstoreRequest) ProtoMessage()    {}
func (*PutKvstoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_088d7f6aff848d9e, []int{0}
}

func (m *PutKvstoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutKvstoreRequest.Unmarshal(m, b)
}
func (m *PutKvstoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutKvstoreRequest.Marshal(b, m, deterministic)
}
func (m *PutKvstoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutKvstoreRequest.Merge(m, src)
}
func (m *PutKvstoreRequest) XXX_Size() int {
	return xxx_messageInfo_PutKvstoreRequest.Size(m)
}
func (m *PutKvstoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutKvstoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutKvstoreRequest proto.InternalMessageInfo

func (m *PutKvstoreRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PutKvstoreRequest) GetK() string {
	if m != nil {
		return m.K
	}
	return ""
}

func (m *PutKvstoreRequest) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

type PutKvstoreReply struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PutKvstoreReply) Reset()         { *m = PutKvstoreReply{} }
func (m *PutKvstoreReply) String() string { return proto.CompactTextString(m) }
func (*PutKvstoreReply) ProtoMessage()    {}
func (*PutKvstoreReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_088d7f6aff848d9e, []int{1}
}

func (m *PutKvstoreReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutKvstoreReply.Unmarshal(m, b)
}
func (m *PutKvstoreReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutKvstoreReply.Marshal(b, m, deterministic)
}
func (m *PutKvstoreReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutKvstoreReply.Merge(m, src)
}
func (m *PutKvstoreReply) XXX_Size() int {
	return xxx_messageInfo_PutKvstoreReply.Size(m)
}
func (m *PutKvstoreReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PutKvstoreReply.DiscardUnknown(m)
}

var xxx_messageInfo_PutKvstoreReply proto.InternalMessageInfo

func (m *PutKvstoreReply) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type GetKvstoreRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	K                    string         `protobuf:"bytes,2,opt,name=k,proto3" json:"k,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetKvstoreRequest) Reset()         { *m = GetKvstoreRequest{} }
func (m *GetKvstoreRequest) String() string { return proto.CompactTextString(m) }
func (*GetKvstoreRequest) ProtoMessage()    {}
func (*GetKvstoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_088d7f6aff848d9e, []int{2}
}

func (m *GetKvstoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKvstoreRequest.Unmarshal(m, b)
}
func (m *GetKvstoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKvstoreRequest.Marshal(b, m, deterministic)
}
func (m *GetKvstoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKvstoreRequest.Merge(m, src)
}
func (m *GetKvstoreRequest) XXX_Size() int {
	return xxx_messageInfo_GetKvstoreRequest.Size(m)
}
func (m *GetKvstoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKvstoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKvstoreRequest proto.InternalMessageInfo

func (m *GetKvstoreRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetKvstoreRequest) GetK() string {
	if m != nil {
		return m.K
	}
	return ""
}

type GetKvstoreReply struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	V                    string          `protobuf:"bytes,2,opt,name=v,proto3" json:"v,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetKvstoreReply) Reset()         { *m = GetKvstoreReply{} }
func (m *GetKvstoreReply) String() string { return proto.CompactTextString(m) }
func (*GetKvstoreReply) ProtoMessage()    {}
func (*GetKvstoreReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_088d7f6aff848d9e, []int{3}
}

func (m *GetKvstoreReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKvstoreReply.Unmarshal(m, b)
}
func (m *GetKvstoreReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKvstoreReply.Marshal(b, m, deterministic)
}
func (m *GetKvstoreReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKvstoreReply.Merge(m, src)
}
func (m *GetKvstoreReply) XXX_Size() int {
	return xxx_messageInfo_GetKvstoreReply.Size(m)
}
func (m *GetKvstoreReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKvstoreReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetKvstoreReply proto.InternalMessageInfo

func (m *GetKvstoreReply) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetKvstoreReply) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

func init() {
	proto.RegisterType((*PutKvstoreRequest)(nil), "api.kv.v1.PutKvstoreRequest")
	proto.RegisterType((*PutKvstoreReply)(nil), "api.kv.v1.PutKvstoreReply")
	proto.RegisterType((*GetKvstoreRequest)(nil), "api.kv.v1.GetKvstoreRequest")
	proto.RegisterType((*GetKvstoreReply)(nil), "api.kv.v1.GetKvstoreReply")
}

func init() { proto.RegisterFile("kvstore.proto", fileDescriptor_088d7f6aff848d9e) }

var fileDescriptor_088d7f6aff848d9e = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x2e, 0x2b, 0x2e,
	0xc9, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4, 0xcb,
	0x2e, 0xd3, 0x2b, 0x33, 0x94, 0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x48, 0x28, 0x25,
	0x72, 0x09, 0x06, 0x94, 0x96, 0x78, 0x43, 0x14, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08,
	0x19, 0x70, 0xb1, 0x65, 0xa4, 0x26, 0xa6, 0xa4, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b,
	0x49, 0xe8, 0xc1, 0xb5, 0xeb, 0x41, 0xd5, 0x78, 0x80, 0xe5, 0x83, 0xa0, 0xea, 0x84, 0x78, 0xb8,
	0x18, 0xb3, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x18, 0xb3, 0x41, 0xbc, 0x32, 0x09, 0x66,
	0x08, 0xaf, 0x4c, 0xc9, 0x85, 0x8b, 0x1f, 0xd9, 0x8a, 0x82, 0x9c, 0x4a, 0x21, 0x43, 0x34, 0x0b,
	0x24, 0x51, 0x2c, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x45, 0xb5, 0x41, 0x29, 0x98, 0x4b, 0xd0,
	0x3d, 0x95, 0xca, 0x0e, 0x55, 0x0a, 0xe2, 0xe2, 0x47, 0x36, 0x94, 0x3c, 0xa7, 0x41, 0xbc, 0x0b,
	0x35, 0xb3, 0xcc, 0x68, 0x26, 0x23, 0x17, 0x3b, 0xd4, 0x44, 0x21, 0x37, 0x2e, 0x2e, 0x84, 0xd7,
	0x85, 0x64, 0x90, 0x8c, 0xc2, 0x08, 0x74, 0x29, 0x29, 0x1c, 0xb2, 0x20, 0x47, 0xb9, 0x71, 0x71,
	0x21, 0xdc, 0x89, 0x62, 0x0e, 0x46, 0x98, 0xa0, 0x98, 0x83, 0xe6, 0x39, 0x27, 0x79, 0x2e, 0x44,
	0x42, 0x08, 0x60, 0x8c, 0x12, 0xc8, 0x4e, 0xac, 0xd4, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x2e, 0xd3,
	0x2f, 0x33, 0xb4, 0x2e, 0x33, 0x4c, 0x62, 0x03, 0xa7, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x3d, 0x03, 0x2e, 0xe0, 0x3f, 0x02, 0x00, 0x00,
}
