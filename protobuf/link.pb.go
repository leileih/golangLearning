// Code generated by protoc-gen-go. DO NOT EDIT.
// source: link.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	link.proto

It has these top-level messages:
	Link
	AllLink
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Link struct {
	Depth int32  `protobuf:"varint,1,opt,name=depth" json:"depth,omitempty"`
	Url   string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *Link) Reset()                    { *m = Link{} }
func (m *Link) String() string            { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()               {}
func (*Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Link) GetDepth() int32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type AllLink struct {
	Link []*Link `protobuf:"bytes,1,rep,name=Link" json:"Link,omitempty"`
}

func (m *AllLink) Reset()                    { *m = AllLink{} }
func (m *AllLink) String() string            { return proto.CompactTextString(m) }
func (*AllLink) ProtoMessage()               {}
func (*AllLink) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AllLink) GetLink() []*Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func init() {
	proto.RegisterType((*Link)(nil), "protobuf.link")
	proto.RegisterType((*AllLink)(nil), "protobuf.all_link")
}

func init() { proto.RegisterFile("link.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xc9, 0xcc, 0xcb,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a, 0x7a, 0x5c,
	0x2c, 0x20, 0x71, 0x21, 0x11, 0x2e, 0xd6, 0x94, 0xd4, 0x82, 0x92, 0x0c, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xd6, 0x20, 0x08, 0x47, 0x48, 0x80, 0x8b, 0xb9, 0xb4, 0x28, 0x47, 0x82, 0x49, 0x81, 0x51,
	0x83, 0x33, 0x08, 0xc4, 0x54, 0xd2, 0xe3, 0xe2, 0x48, 0xcc, 0xc9, 0x89, 0x07, 0xeb, 0x51, 0xe2,
	0x62, 0xf1, 0xc9, 0xcc, 0xcb, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd3, 0x83, 0x19,
	0xaa, 0x07, 0x92, 0x0d, 0x02, 0xcb, 0x25, 0xb1, 0x81, 0x05, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x7a, 0x7b, 0x42, 0x0c, 0x7e, 0x00, 0x00, 0x00,
}