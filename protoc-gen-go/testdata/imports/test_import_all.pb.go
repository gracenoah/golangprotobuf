// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_import_all.proto

package imports

import (
	fmt "fmt"
	proto "github.com/gracenoah/golangprotobuf/proto"
	fmt1 "github.com/gracenoah/golangprotobuf/protoc-gen-go/testdata/imports/fmt"
	test_a_1 "github.com/gracenoah/golangprotobuf/protoc-gen-go/testdata/imports/test_a_1"
	test_a_2 "github.com/gracenoah/golangprotobuf/protoc-gen-go/testdata/imports/test_a_2"
	test_b_1 "github.com/gracenoah/golangprotobuf/protoc-gen-go/testdata/imports/test_b_1"
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

type All struct {
	Am1                  *test_a_1.M1 `protobuf:"bytes,1,opt,name=am1,proto3" json:"am1,omitempty"`
	Am2                  *test_a_1.M2 `protobuf:"bytes,2,opt,name=am2,proto3" json:"am2,omitempty"`
	Am3                  *test_a_2.M3 `protobuf:"bytes,3,opt,name=am3,proto3" json:"am3,omitempty"`
	Am4                  *test_a_2.M4 `protobuf:"bytes,4,opt,name=am4,proto3" json:"am4,omitempty"`
	Bm1                  *test_b_1.M1 `protobuf:"bytes,5,opt,name=bm1,proto3" json:"bm1,omitempty"`
	Bm2                  *test_b_1.M2 `protobuf:"bytes,6,opt,name=bm2,proto3" json:"bm2,omitempty"`
	Fmt                  *fmt1.M      `protobuf:"bytes,7,opt,name=fmt,proto3" json:"fmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *All) Reset()         { *m = All{} }
func (m *All) String() string { return proto.CompactTextString(m) }
func (*All) ProtoMessage()    {}
func (*All) Descriptor() ([]byte, []int) {
	return fileDescriptor_324466f0afc16f77, []int{0}
}

func (m *All) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_All.Unmarshal(m, b)
}
func (m *All) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_All.Marshal(b, m, deterministic)
}
func (m *All) XXX_Merge(src proto.Message) {
	xxx_messageInfo_All.Merge(m, src)
}
func (m *All) XXX_Size() int {
	return xxx_messageInfo_All.Size(m)
}
func (m *All) XXX_DiscardUnknown() {
	xxx_messageInfo_All.DiscardUnknown(m)
}

var xxx_messageInfo_All proto.InternalMessageInfo

func (m *All) GetAm1() *test_a_1.M1 {
	if m != nil {
		return m.Am1
	}
	return nil
}

func (m *All) GetAm2() *test_a_1.M2 {
	if m != nil {
		return m.Am2
	}
	return nil
}

func (m *All) GetAm3() *test_a_2.M3 {
	if m != nil {
		return m.Am3
	}
	return nil
}

func (m *All) GetAm4() *test_a_2.M4 {
	if m != nil {
		return m.Am4
	}
	return nil
}

func (m *All) GetBm1() *test_b_1.M1 {
	if m != nil {
		return m.Bm1
	}
	return nil
}

func (m *All) GetBm2() *test_b_1.M2 {
	if m != nil {
		return m.Bm2
	}
	return nil
}

func (m *All) GetFmt() *fmt1.M {
	if m != nil {
		return m.Fmt
	}
	return nil
}

func init() {
	proto.RegisterType((*All)(nil), "test.All")
}

func init() { proto.RegisterFile("imports/test_import_all.proto", fileDescriptor_324466f0afc16f77) }

var fileDescriptor_324466f0afc16f77 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0x15, 0x97, 0x20, 0x99, 0x05, 0x85, 0xc5, 0x20, 0x90, 0x50, 0x27, 0x96, 0xda,
	0xb2, 0x9d, 0x05, 0x31, 0xc1, 0xde, 0xa5, 0x23, 0x4b, 0x64, 0x97, 0xc6, 0x54, 0xf2, 0xd5, 0x51,
	0x7a, 0x7d, 0x5e, 0x5e, 0x05, 0xd9, 0x07, 0x12, 0x84, 0x66, 0x4b, 0xfe, 0xef, 0xb7, 0xce, 0x3e,
	0x7e, 0xbf, 0x83, 0x3e, 0x0d, 0x78, 0x50, 0xb8, 0x3d, 0x60, 0x4b, 0x3f, 0xad, 0x8b, 0x51, 0xf6,
	0x43, 0xc2, 0x54, 0xcf, 0x73, 0x7c, 0x7b, 0xf3, 0xa7, 0xe4, 0x5a, 0xad, 0x40, 0x53, 0xe1, 0x14,
	0x99, 0x09, 0x32, 0x0a, 0xec, 0x34, 0x35, 0x27, 0xc9, 0x4f, 0xcf, 0xf2, 0xbf, 0x67, 0x5d, 0xff,
	0x50, 0x07, 0xa8, 0x80, 0xc2, 0xc5, 0xe7, 0x8c, 0xb3, 0x97, 0x18, 0xeb, 0x3b, 0xce, 0x1c, 0x68,
	0x31, 0x7b, 0x98, 0x3d, 0x5e, 0x1a, 0x2e, 0xf3, 0x69, 0xe9, 0xe4, 0x4a, 0xaf, 0x73, 0x4c, 0x6a,
	0xc4, 0xd9, 0x48, 0x4d, 0x56, 0x43, 0x6a, 0x05, 0x1b, 0xa9, 0xcd, 0x6a, 0x49, 0x1b, 0x31, 0x1f,
	0x69, 0x93, 0xb5, 0xa9, 0x17, 0x9c, 0x79, 0xd0, 0xe2, 0xbc, 0xe8, 0x15, 0xa9, 0x97, 0xbd, 0x1b,
	0x50, 0x97, 0xe9, 0x1e, 0x34, 0x75, 0x8c, 0xa8, 0xfe, 0x77, 0x4c, 0xb9, 0x83, 0x07, 0x53, 0x0b,
	0xce, 0x3a, 0x40, 0x71, 0x51, 0x3a, 0x95, 0xec, 0x00, 0xe5, 0x6a, 0x9d, 0xa3, 0xd7, 0xe7, 0xb7,
	0xa7, 0xb0, 0xc3, 0x8f, 0xa3, 0x97, 0x9b, 0x04, 0x2a, 0xa4, 0xe8, 0xf6, 0x41, 0x95, 0xc7, 0xfb,
	0x63, 0x47, 0x1f, 0x9b, 0x65, 0xd8, 0xee, 0x97, 0x21, 0x95, 0xa5, 0xbd, 0x3b, 0x74, 0xea, 0x7b,
	0x55, 0xbe, 0x2a, 0x6e, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x95, 0x39, 0xa3, 0x82, 0x03, 0x02,
	0x00, 0x00,
}
