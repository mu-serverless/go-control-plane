// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/grpc_http1_reverse_bridge/v2alpha1/config.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

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

// [#protodoc-title: Extensions gRPC Http1 Reverse Bridge]
// gRPC reverse bridge filter configuration
type FilterConfig struct {
	// The content-type to pass to the upstream when the gRPC bridge filter is applied.
	// The filter will also validate that the upstream responds with the same content type.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// If true, Envoy will assume that the upstream doesn't understand gRPC frames and
	// strip the gRPC frame from the request, and add it back in to the response. This will
	// hide the gRPC semantics from the upstream, allowing it to receive and respond with a
	// simple binary encoded protobuf.
	WithholdGrpcFrames   bool     `protobuf:"varint,2,opt,name=withhold_grpc_frames,json=withholdGrpcFrames,proto3" json:"withhold_grpc_frames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_be01b80cbe68f903, []int{0}
}
func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(dst, src)
}
func (m *FilterConfig) XXX_Size() int {
	return m.Size()
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *FilterConfig) GetWithholdGrpcFrames() bool {
	if m != nil {
		return m.WithholdGrpcFrames
	}
	return false
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.extensions.filter.http.grpc_http1_reverse_bridge.v2alpha1.FilterConfig")
}
func (m *FilterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ContentType) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ContentType)))
		i += copy(dAtA[i:], m.ContentType)
	}
	if m.WithholdGrpcFrames {
		dAtA[i] = 0x10
		i++
		if m.WithholdGrpcFrames {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FilterConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContentType)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.WithholdGrpcFrames {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FilterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: FilterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithholdGrpcFrames", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WithholdGrpcFrames = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/grpc_http1_reverse_bridge/v2alpha1/config.proto", fileDescriptor_config_be01b80cbe68f903)
}

var fileDescriptor_config_be01b80cbe68f903 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0xcf, 0xb1, 0x4a, 0x03, 0x31,
	0x18, 0x07, 0x70, 0x52, 0x45, 0x6c, 0xec, 0x74, 0x08, 0x16, 0x87, 0xe3, 0x70, 0xea, 0x20, 0x89,
	0xad, 0x0f, 0x20, 0x54, 0xa8, 0x2e, 0x2e, 0xc5, 0xc9, 0x25, 0xa4, 0x77, 0xdf, 0xdd, 0x05, 0xce,
	0x24, 0x7c, 0xf7, 0x11, 0x7b, 0xaf, 0xe6, 0xe4, 0xe8, 0xe8, 0x23, 0xc8, 0x6d, 0xbe, 0x85, 0x34,
	0xed, 0x8d, 0x4e, 0x6e, 0x7f, 0xf8, 0x93, 0x5f, 0xbe, 0x3f, 0x7f, 0x04, 0x1b, 0x5c, 0x27, 0x73,
	0x67, 0x4b, 0x53, 0xc9, 0xd2, 0x34, 0x04, 0x28, 0x6b, 0x22, 0x2f, 0x2b, 0xf4, 0xb9, 0xda, 0xa5,
	0xb9, 0x42, 0x08, 0x80, 0x2d, 0xa8, 0x0d, 0x9a, 0xa2, 0x02, 0x19, 0x16, 0xba, 0xf1, 0xb5, 0x9e,
	0x1f, 0x5e, 0x09, 0x8f, 0x8e, 0x5c, 0x72, 0x17, 0x25, 0x01, 0x5b, 0x02, 0xdb, 0x1a, 0x67, 0x5b,
	0xb1, 0xd7, 0xc4, 0xce, 0x10, 0x7f, 0x6a, 0x62, 0xd0, 0x2e, 0x2f, 0x82, 0x6e, 0x4c, 0xa1, 0x09,
	0xe4, 0x10, 0xf6, 0xf2, 0x95, 0xe5, 0x93, 0x55, 0xa4, 0xee, 0xe3, 0x7f, 0xc9, 0x35, 0x9f, 0xe4,
	0xce, 0x12, 0x58, 0x52, 0xd4, 0x79, 0x98, 0xb2, 0x8c, 0xcd, 0xc6, 0xcb, 0xf1, 0xfb, 0xcf, 0xc7,
	0xd1, 0x31, 0x8e, 0x32, 0xb6, 0x3e, 0x3b, 0xd4, 0xcf, 0x9d, 0x87, 0xe4, 0x86, 0x9f, 0xbf, 0x19,
	0xaa, 0x6b, 0xd7, 0x14, 0x2a, 0x5e, 0x51, 0xa2, 0x7e, 0x85, 0x76, 0x3a, 0xca, 0xd8, 0xec, 0x74,
	0x9d, 0x0c, 0xdd, 0x03, 0xfa, 0x7c, 0x15, 0x9b, 0xa5, 0xfa, 0xec, 0x53, 0xf6, 0xd5, 0xa7, 0xec,
	0xbb, 0x4f, 0x19, 0x7f, 0x32, 0x4e, 0xc4, 0x69, 0x1e, 0xdd, 0xb6, 0x13, 0xff, 0x5c, 0xf9, 0x32,
	0x0a, 0x8b, 0xcd, 0x49, 0xdc, 0x75, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x75, 0xad, 0xd0, 0xc9,
	0x7d, 0x01, 0x00, 0x00,
}
