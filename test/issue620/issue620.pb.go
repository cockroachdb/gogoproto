// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue620.proto

package issue620

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Bar struct {
	Field1               *string  `protobuf:"bytes,1,opt,name=Field1" json:"Field1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bar) Reset()         { *m = Bar{} }
func (m *Bar) String() string { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()    {}
func (*Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_63df3af46ac96f83, []int{0}
}
func (m *Bar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bar.Unmarshal(m, b)
}
func (m *Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bar.Marshal(b, m, deterministic)
}
func (m *Bar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bar.Merge(m, src)
}
func (m *Bar) XXX_Size() int {
	return xxx_messageInfo_Bar.Size(m)
}
func (m *Bar) XXX_DiscardUnknown() {
	xxx_messageInfo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Bar proto.InternalMessageInfo

func (m *Bar) GetField1() string {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return ""
}

type BarFast struct {
	Field1               *string  `protobuf:"bytes,1,opt,name=Field1" json:"Field1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarFast) Reset()         { *m = BarFast{} }
func (m *BarFast) String() string { return proto.CompactTextString(m) }
func (*BarFast) ProtoMessage()    {}
func (*BarFast) Descriptor() ([]byte, []int) {
	return fileDescriptor_63df3af46ac96f83, []int{1}
}
func (m *BarFast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarFast.Unmarshal(m, b)
}
func (m *BarFast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BarFast.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BarFast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarFast.Merge(m, src)
}
func (m *BarFast) XXX_Size() int {
	return m.Size()
}
func (m *BarFast) XXX_DiscardUnknown() {
	xxx_messageInfo_BarFast.DiscardUnknown(m)
}

var xxx_messageInfo_BarFast proto.InternalMessageInfo

func (m *BarFast) GetField1() string {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return ""
}

func init() {
	proto.RegisterType((*Bar)(nil), "issue620.Bar")
	proto.RegisterType((*BarFast)(nil), "issue620.BarFast")
}

func init() { proto.RegisterFile("issue620.proto", fileDescriptor_63df3af46ac96f83) }

var fileDescriptor_63df3af46ac96f83 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x33, 0x32, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x74,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5,
	0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x54, 0x92, 0xe5, 0x62,
	0x76, 0x4a, 0x2c, 0x12, 0x12, 0xe3, 0x62, 0x73, 0xcb, 0x4c, 0xcd, 0x49, 0x31, 0x94, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x94, 0xf4, 0xb9, 0xd8, 0x9d, 0x12, 0x8b, 0xdc, 0x12, 0x8b,
	0x4b, 0x70, 0x29, 0xb1, 0xe2, 0x79, 0xb1, 0x40, 0x9e, 0xb1, 0x63, 0xa1, 0x3c, 0xe3, 0x82, 0x85,
	0xf2, 0x8c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x1c, 0x6d, 0xdf, 0x99, 0x00, 0x00, 0x00,
}

func (this *BarFast) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BarFast)
	if !ok {
		that2, ok := that.(BarFast)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return false
		}
	} else if this.Field1 != nil {
		return false
	} else if that1.Field1 != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *BarFast) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BarFast) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BarFast) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Field1 != nil {
		i -= len(*m.Field1)
		copy(dAtA[i:], *m.Field1)
		i = encodeVarintIssue620(dAtA, i, uint64(len(*m.Field1)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIssue620(dAtA []byte, offset int, v uint64) int {
	offset -= sovIssue620(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BarFast) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Field1 != nil {
		l = len(*m.Field1)
		n += 1 + l + sovIssue620(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIssue620(x uint64) (n int) {
	return int((uint32(math_bits.Len64(x|1)+6) * 37) >> 8)
}
func sozIssue620(x uint64) (n int) {
	return sovIssue620(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
