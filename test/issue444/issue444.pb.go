// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue444.proto

package issue444

import (
	fmt "fmt"
	_ "github.com/cockroachdb/gogoproto/gogoproto"
	proto "github.com/cockroachdb/gogoproto/proto"
	math "math"
	math_bits "math/bits"
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

type SizeMe struct {
	Foo                  string   `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SizeMe) Reset()         { *m = SizeMe{} }
func (m *SizeMe) String() string { return proto.CompactTextString(m) }
func (*SizeMe) ProtoMessage()    {}
func (*SizeMe) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39b8ab75bd9d8cd, []int{0}
}
func (m *SizeMe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SizeMe.Unmarshal(m, b)
}
func (m *SizeMe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SizeMe.Marshal(b, m, deterministic)
}
func (m *SizeMe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SizeMe.Merge(m, src)
}
func (m *SizeMe) XXX_Size() int {
	return xxx_messageInfo_SizeMe.Size(m)
}
func (m *SizeMe) XXX_DiscardUnknown() {
	xxx_messageInfo_SizeMe.DiscardUnknown(m)
}

var xxx_messageInfo_SizeMe proto.InternalMessageInfo

func (m *SizeMe) GetFoo() string {
	if m != nil {
		return m.Foo
	}
	return ""
}

func init() {
	proto.RegisterType((*SizeMe)(nil), "issue444.SizeMe")
}

func init() { proto.RegisterFile("issue444.proto", fileDescriptor_e39b8ab75bd9d8cd) }

var fileDescriptor_e39b8ab75bd9d8cd = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x31, 0x31, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x4c,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x93, 0xb3, 0x8b,
	0xf2, 0x13, 0x93, 0x33, 0x52, 0x92, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x6a, 0xd1, 0x58, 0x10,
	0x03, 0x94, 0xa4, 0xb8, 0xd8, 0x82, 0x33, 0xab, 0x52, 0x7d, 0x53, 0x85, 0x04, 0xb8, 0x98, 0xd3,
	0xf2, 0xf3, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x27, 0x96, 0x07, 0x8f, 0xe4,
	0x18, 0x93, 0xd8, 0xc0, 0x0a, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x12, 0x8a, 0x68, 0x98,
	0x7b, 0x00, 0x00, 0x00,
}

func (m *SizeMe) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Foo)
	if l > 0 {
		n += 1 + l + sovIssue444(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIssue444(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIssue444(x uint64) (n int) {
	return sovIssue444(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
