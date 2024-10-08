// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue411.proto

package issue411

import (
	fmt "fmt"
	_ "github.com/cockroachdb/gogoproto/gogoproto"
	proto "github.com/cockroachdb/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Span struct {
	TraceID              TraceID  `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3,customtype=TraceID" json:"trace_id"`
	SpanID               SpanID   `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3,customtype=SpanID" json:"span_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1ed5cde895f96f, []int{0}
}
func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Span)(nil), "issue411.Span")
}

func init() { proto.RegisterFile("issue411.proto", fileDescriptor_7e1ed5cde895f96f) }

var fileDescriptor_7e1ed5cde895f96f = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x31, 0x34, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x4c,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x93, 0xb3, 0x8b,
	0xf2, 0x13, 0x93, 0x33, 0x52, 0x92, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x6a, 0xd1, 0x58, 0x10,
	0x03, 0x94, 0x0a, 0xb8, 0x58, 0x82, 0x0b, 0x12, 0xf3, 0x84, 0x4c, 0xb9, 0x38, 0x4a, 0x8a, 0x12,
	0x93, 0x53, 0xe3, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x9c, 0xa4, 0x4e, 0xdc, 0x93,
	0x67, 0xb8, 0x75, 0x4f, 0x9e, 0x3d, 0x04, 0x24, 0xee, 0xe9, 0xf2, 0x08, 0xc1, 0x0c, 0x62, 0x07,
	0xab, 0xf5, 0x4c, 0x11, 0x32, 0xe4, 0x62, 0x2f, 0x2e, 0x48, 0xcc, 0x03, 0xe9, 0x62, 0x02, 0xeb,
	0x92, 0x80, 0xea, 0x62, 0x03, 0x99, 0x0a, 0xd6, 0x04, 0x65, 0x05, 0xb1, 0x81, 0x14, 0x7a, 0xa6,
	0x24, 0xb1, 0x81, 0x2d, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xad, 0xa9, 0x5c, 0xb6, 0xcb,
	0x00, 0x00, 0x00,
}
