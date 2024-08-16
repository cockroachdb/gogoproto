// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: typedeclimport.proto

/*
Package typedeclimport is a generated protocol buffer package.

It is generated from these files:

	typedeclimport.proto

It has these top-level messages:

	SomeMessage
*/
package typedeclimport

import proto "github.com/cockroachdb/gogoproto/proto"
import fmt "fmt"
import math "math"
import _ "github.com/cockroachdb/gogoproto/gogoproto"
import subpkg "github.com/cockroachdb/gogoproto/test/typedeclimport/subpkg"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *SomeMessage) Reset()                    { *m = SomeMessage{} }
func (m *SomeMessage) String() string            { return proto.CompactTextString(m) }
func (*SomeMessage) ProtoMessage()               {}
func (*SomeMessage) Descriptor() ([]byte, []int) { return fileDescriptorTypedeclimport, []int{0} }

func (m *SomeMessage) GetImported() subpkg.AnotherMessage {
	if m != nil {
		return m.Imported
	}
	return subpkg.AnotherMessage{}
}

func init() {
	proto.RegisterType((*SomeMessage)(nil), "typedeclimport.SomeMessage")
}
func (this *SomeMessage) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*SomeMessage)
	if !ok {
		that2, ok := that.(SomeMessage)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *SomeMessage")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *SomeMessage but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *SomeMessage but is not nil && this == nil")
	}
	if !this.Imported.Equal(&that1.Imported) {
		return fmt.Errorf("Imported this(%v) Not Equal that(%v)", this.Imported, that1.Imported)
	}
	return nil
}
func (this *SomeMessage) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SomeMessage)
	if !ok {
		that2, ok := that.(SomeMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Imported.Equal(&that1.Imported) {
		return false
	}
	return true
}
func (m *SomeMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SomeMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintTypedeclimport(dAtA, i, uint64(m.Imported.Size()))
	n1, err := m.Imported.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func encodeFixed64Typedeclimport(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Typedeclimport(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTypedeclimport(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SomeMessage) Size() (n int) {
	var l int
	_ = l
	l = m.Imported.Size()
	n += 1 + l + sovTypedeclimport(uint64(l))
	return n
}

func sovTypedeclimport(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypedeclimport(x uint64) (n int) {
	return sovTypedeclimport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SomeMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypedeclimport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SomeMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SomeMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Imported", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypedeclimport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypedeclimport
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Imported.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypedeclimport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypedeclimport
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypedeclimport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypedeclimport
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypedeclimport
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypedeclimport
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTypedeclimport
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypedeclimport
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTypedeclimport(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTypedeclimport = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypedeclimport   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("typedeclimport.proto", fileDescriptorTypedeclimport) }

var fileDescriptorTypedeclimport = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xa9, 0x2c, 0x48,
	0x4d, 0x49, 0x4d, 0xce, 0xc9, 0xcc, 0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x43, 0x15, 0x95, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b, 0x4b, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30,
	0x0b, 0xa2, 0x5d, 0xca, 0x01, 0xa7, 0xf2, 0x92, 0xd4, 0xe2, 0x12, 0x7d, 0x54, 0xc3, 0xf5, 0x8b,
	0x4b, 0x93, 0x0a, 0xb2, 0xd3, 0xa1, 0x14, 0xc4, 0x04, 0x25, 0x5f, 0x2e, 0xee, 0xe0, 0xfc, 0xdc,
	0x54, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21, 0x0b, 0x2e, 0x0e, 0x88, 0xe2, 0xd4, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x31, 0x3d, 0xa8, 0x7a, 0xc7, 0xbc, 0xfc, 0x92, 0x8c,
	0xd4, 0x22, 0xa8, 0x4a, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0xe0, 0xaa, 0xad, 0x58, 0x3e,
	0x2c, 0x94, 0x67, 0x70, 0x12, 0x79, 0xf0, 0x50, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x31, 0x89, 0x0d, 0x6c, 0x97, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x54, 0x23, 0xca, 0x44, 0x04, 0x01, 0x00, 0x00,
}
