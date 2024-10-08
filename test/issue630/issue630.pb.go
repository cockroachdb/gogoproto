// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue630.proto

package issue630

import (
	fmt "fmt"
	_ "github.com/cockroachdb/gogoproto/gogoproto"
	proto "github.com/cockroachdb/gogoproto/proto"
	math "math"
	reflect "reflect"
	strings "strings"
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

type Foo struct {
	Bar1                 []Bar    `protobuf:"bytes,1,rep,name=Bar1" json:"Bar1"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_60374e06e51d301c, []int{0}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetBar1() []Bar {
	if m != nil {
		return m.Bar1
	}
	return nil
}

type Qux struct {
	Bar1                 []*Bar   `protobuf:"bytes,1,rep,name=Bar1" json:"Bar1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Qux) Reset()         { *m = Qux{} }
func (m *Qux) String() string { return proto.CompactTextString(m) }
func (*Qux) ProtoMessage()    {}
func (*Qux) Descriptor() ([]byte, []int) {
	return fileDescriptor_60374e06e51d301c, []int{1}
}
func (m *Qux) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Qux.Unmarshal(m, b)
}
func (m *Qux) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Qux.Marshal(b, m, deterministic)
}
func (m *Qux) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Qux.Merge(m, src)
}
func (m *Qux) XXX_Size() int {
	return xxx_messageInfo_Qux.Size(m)
}
func (m *Qux) XXX_DiscardUnknown() {
	xxx_messageInfo_Qux.DiscardUnknown(m)
}

var xxx_messageInfo_Qux proto.InternalMessageInfo

func (m *Qux) GetBar1() []*Bar {
	if m != nil {
		return m.Bar1
	}
	return nil
}

type Bar struct {
	Baz                  string   `protobuf:"bytes,1,opt,name=Baz" json:"Baz"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bar) Reset()         { *m = Bar{} }
func (m *Bar) String() string { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()    {}
func (*Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_60374e06e51d301c, []int{2}
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

func (m *Bar) GetBaz() string {
	if m != nil {
		return m.Baz
	}
	return ""
}

func init() {
	proto.RegisterType((*Foo)(nil), "issue630.Foo")
	proto.RegisterType((*Qux)(nil), "issue630.Qux")
	proto.RegisterType((*Bar)(nil), "issue630.Bar")
}

func init() { proto.RegisterFile("issue630.proto", fileDescriptor_60374e06e51d301c) }

var fileDescriptor_60374e06e51d301c = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x33, 0x36, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x4c,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x93, 0xb3, 0x8b,
	0xf2, 0x13, 0x93, 0x33, 0x52, 0x92, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x6a, 0xd1, 0x58, 0x10,
	0x03, 0x94, 0xf4, 0xb8, 0x98, 0xdd, 0xf2, 0xf3, 0x85, 0xd4, 0xb9, 0x58, 0x9c, 0x12, 0x8b, 0x0c,
	0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x78, 0xf5, 0xe0, 0xd6, 0x38, 0x25, 0x16, 0x39, 0xb1,
	0x9c, 0xb8, 0x27, 0xcf, 0x10, 0x04, 0x56, 0x00, 0x52, 0x1f, 0x58, 0x5a, 0x41, 0x58, 0x3d, 0x23,
	0x54, 0xbd, 0x2c, 0x17, 0xb3, 0x53, 0x62, 0x91, 0x90, 0x18, 0x88, 0xaa, 0x92, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x84, 0x9a, 0x07, 0x12, 0x70, 0x62, 0xf9, 0xf0, 0x50, 0x8e, 0x11, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x81, 0x8b, 0xad, 0xd6, 0xd6, 0x00, 0x00, 0x00,
}

func (this *Foo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&issue630.Foo{")
	if this.Bar1 != nil {
		vs := make([]Bar, len(this.Bar1))
		for i := range vs {
			vs[i] = this.Bar1[i]
		}
		s = append(s, "Bar1: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Qux) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&issue630.Qux{")
	if this.Bar1 != nil {
		s = append(s, "Bar1: "+fmt.Sprintf("%#v", this.Bar1)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Bar) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&issue630.Bar{")
	s = append(s, "Baz: "+fmt.Sprintf("%#v", this.Baz)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringIssue630(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
