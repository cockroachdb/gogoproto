// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto3/proto3.proto

package proto3

import (
	fmt "fmt"
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

type Request_Flavour int32

const (
	Request_SWEET         Request_Flavour = 0
	Request_SOUR          Request_Flavour = 1
	Request_UMAMI         Request_Flavour = 2
	Request_GOPHERLICIOUS Request_Flavour = 3
)

var Request_Flavour_name = map[int32]string{
	0: "SWEET",
	1: "SOUR",
	2: "UMAMI",
	3: "GOPHERLICIOUS",
}

var Request_Flavour_value = map[string]int32{
	"SWEET":         0,
	"SOUR":          1,
	"UMAMI":         2,
	"GOPHERLICIOUS": 3,
}

func (x Request_Flavour) String() string {
	return proto.EnumName(Request_Flavour_name, int32(x))
}

func (Request_Flavour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{0, 0}
}

type Request struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key                  []int64         `protobuf:"varint,2,rep,packed,name=key,proto3" json:"key,omitempty"`
	Taste                Request_Flavour `protobuf:"varint,3,opt,name=taste,proto3,enum=proto3.Request_Flavour" json:"taste,omitempty"`
	Book                 *Book           `protobuf:"bytes,4,opt,name=book,proto3" json:"book,omitempty"`
	Unpacked             []int64         `protobuf:"varint,5,rep,name=unpacked,proto3" json:"unpacked,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetKey() []int64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetTaste() Request_Flavour {
	if m != nil {
		return m.Taste
	}
	return Request_SWEET
}

func (m *Request) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

func (m *Request) GetUnpacked() []int64 {
	if m != nil {
		return m.Unpacked
	}
	return nil
}

type Book struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	RawData              []byte   `protobuf:"bytes,2,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{1}
}
func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto3.Request_Flavour", Request_Flavour_name, Request_Flavour_value)
	proto.RegisterType((*Request)(nil), "proto3.Request")
	proto.RegisterType((*Book)(nil), "proto3.Book")
}

func init() { proto.RegisterFile("proto3/proto3.proto", fileDescriptor_ab04eb4084a521db) }

var fileDescriptor_ab04eb4084a521db = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0xd9, 0xfe, 0x11, 0x18, 0xd1, 0xd4, 0xd1, 0xc4, 0x7a, 0x31, 0x0d, 0xa7, 0x5e, 0x68,
	0x13, 0x3c, 0x78, 0xf1, 0xa0, 0x55, 0x54, 0x12, 0x09, 0x66, 0x91, 0x98, 0x78, 0x31, 0xdb, 0xb2,
	0x29, 0xa4, 0xc0, 0x60, 0xd9, 0x4a, 0xfc, 0xb2, 0x7e, 0x16, 0xb3, 0x6d, 0xf1, 0x34, 0x33, 0x6f,
	0x26, 0xbf, 0xc9, 0x7b, 0x70, 0xba, 0xc9, 0x49, 0xd1, 0x55, 0x58, 0x95, 0xa0, 0x2c, 0x78, 0x50,
	0x4d, 0xdd, 0x5f, 0x06, 0x4d, 0x2e, 0xbf, 0x0a, 0xb9, 0x55, 0x88, 0x60, 0xad, 0xc5, 0x4a, 0xba,
	0xcc, 0x63, 0x7e, 0x9b, 0x97, 0x3d, 0x3a, 0x60, 0x66, 0xf2, 0xc7, 0x35, 0x3c, 0xd3, 0x37, 0xb9,
	0x6e, 0xb1, 0x07, 0xb6, 0x12, 0x5b, 0x25, 0x5d, 0xd3, 0x63, 0xfe, 0x71, 0xff, 0x3c, 0xa8, 0xb9,
	0x35, 0x25, 0x78, 0x5c, 0x8a, 0x6f, 0x2a, 0x72, 0x5e, 0x5d, 0xa1, 0x07, 0x56, 0x4c, 0x94, 0xb9,
	0x96, 0xc7, 0xfc, 0xc3, 0x7e, 0x67, 0x7f, 0x1d, 0x11, 0x65, 0xbc, 0xdc, 0xe0, 0x25, 0xb4, 0x8a,
	0xf5, 0x46, 0x24, 0x99, 0x9c, 0xb9, 0xb6, 0xfe, 0x13, 0x19, 0x4e, 0x83, 0xff, 0x6b, 0xdd, 0x1b,
	0x68, 0xd6, 0x4c, 0x6c, 0x83, 0x3d, 0x79, 0x1f, 0x0c, 0xde, 0x9c, 0x06, 0xb6, 0xc0, 0x9a, 0x8c,
	0xa7, 0xdc, 0x61, 0x5a, 0x9c, 0x8e, 0xee, 0x46, 0x43, 0xc7, 0xc0, 0x13, 0x38, 0x7a, 0x1a, 0xbf,
	0x3e, 0x0f, 0xf8, 0xcb, 0xf0, 0x7e, 0x38, 0x9e, 0x4e, 0x1c, 0xb3, 0x7b, 0x0d, 0x96, 0xfe, 0x85,
	0x67, 0x60, 0xab, 0x85, 0x5a, 0xee, 0xdd, 0x55, 0x03, 0x5e, 0x40, 0x2b, 0x17, 0xbb, 0xcf, 0x99,
	0x50, 0xc2, 0x35, 0x3c, 0xe6, 0x77, 0x78, 0x33, 0x17, 0xbb, 0x07, 0xa1, 0x44, 0x14, 0x7d, 0xdc,
	0xa6, 0x0b, 0x35, 0x2f, 0xe2, 0x20, 0xa1, 0x55, 0x98, 0x50, 0x92, 0xe5, 0x24, 0x92, 0xf9, 0x2c,
	0x0e, 0x53, 0x4a, 0xa9, 0xb4, 0x51, 0x45, 0x9a, 0xf4, 0x52, 0xb9, 0xee, 0x69, 0x31, 0x54, 0x72,
	0xab, 0x34, 0xab, 0xce, 0x3a, 0xae, 0x53, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x6f, 0x15,
	0xf3, 0x83, 0x01, 0x00, 0x00,
}
