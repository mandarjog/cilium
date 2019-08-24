// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/number.proto

package matcher

import (
	fmt "fmt"
	_type "github.com/cilium/proxy/go/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Specifies the way to match a double value.
type DoubleMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*DoubleMatcher_Range
	//	*DoubleMatcher_Exact
	MatchPattern         isDoubleMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DoubleMatcher) Reset()         { *m = DoubleMatcher{} }
func (m *DoubleMatcher) String() string { return proto.CompactTextString(m) }
func (*DoubleMatcher) ProtoMessage()    {}
func (*DoubleMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_58ad3770a33c5fc2, []int{0}
}

func (m *DoubleMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleMatcher.Unmarshal(m, b)
}
func (m *DoubleMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleMatcher.Marshal(b, m, deterministic)
}
func (m *DoubleMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleMatcher.Merge(m, src)
}
func (m *DoubleMatcher) XXX_Size() int {
	return xxx_messageInfo_DoubleMatcher.Size(m)
}
func (m *DoubleMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleMatcher proto.InternalMessageInfo

type isDoubleMatcher_MatchPattern interface {
	isDoubleMatcher_MatchPattern()
}

type DoubleMatcher_Range struct {
	Range *_type.DoubleRange `protobuf:"bytes,1,opt,name=range,proto3,oneof"`
}

type DoubleMatcher_Exact struct {
	Exact float64 `protobuf:"fixed64,2,opt,name=exact,proto3,oneof"`
}

func (*DoubleMatcher_Range) isDoubleMatcher_MatchPattern() {}

func (*DoubleMatcher_Exact) isDoubleMatcher_MatchPattern() {}

func (m *DoubleMatcher) GetMatchPattern() isDoubleMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *DoubleMatcher) GetRange() *_type.DoubleRange {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Range); ok {
		return x.Range
	}
	return nil
}

func (m *DoubleMatcher) GetExact() float64 {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Exact); ok {
		return x.Exact
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DoubleMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DoubleMatcher_Range)(nil),
		(*DoubleMatcher_Exact)(nil),
	}
}

func init() {
	proto.RegisterType((*DoubleMatcher)(nil), "envoy.type.matcher.DoubleMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/number.proto", fileDescriptor_58ad3770a33c5fc2) }

var fileDescriptor_58ad3770a33c5fc2 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0xcf,
	0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02, 0x2b, 0xd0,
	0x03, 0x29, 0xd0, 0x83, 0x2a, 0x90, 0x12, 0x43, 0xd2, 0x54, 0x94, 0x98, 0x97, 0x9e, 0x0a, 0x51,
	0x2b, 0x25, 0x5e, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58, 0x92, 0xaa, 0x0f, 0x63, 0x40, 0x24, 0x94,
	0x0a, 0xb8, 0x78, 0x5d, 0xf2, 0x4b, 0x93, 0x72, 0x52, 0x7d, 0x21, 0x26, 0x08, 0xe9, 0x73, 0xb1,
	0x82, 0x35, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xeb, 0x21, 0xd9, 0x02, 0x51, 0x19,
	0x04, 0x92, 0xf6, 0x60, 0x08, 0x82, 0xa8, 0x13, 0x12, 0xe3, 0x62, 0x4d, 0xad, 0x48, 0x4c, 0x2e,
	0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x04, 0x89, 0x83, 0xb9, 0x4e, 0x62, 0x5c, 0xbc, 0x60, 0x57,
	0xc5, 0x17, 0x24, 0x96, 0x94, 0xa4, 0x16, 0xe5, 0x09, 0xb1, 0xee, 0x78, 0x79, 0x80, 0x99, 0xd1,
	0xc9, 0x8a, 0x4b, 0x21, 0x33, 0x1f, 0x62, 0x6a, 0x41, 0x51, 0x7e, 0x45, 0xa5, 0x1e, 0xa6, 0x37,
	0x9c, 0xb8, 0xfd, 0xc0, 0x1e, 0x0d, 0x00, 0x39, 0x31, 0x80, 0x31, 0x8a, 0x1d, 0x2a, 0x9e, 0xc4,
	0x06, 0x76, 0xb4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x9e, 0xa0, 0xd7, 0x1c, 0x01, 0x00,
	0x00,
}
