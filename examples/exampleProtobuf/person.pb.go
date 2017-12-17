// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	person.proto

It has these top-level messages:
	Person
	AllPerson
*/
package example

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

type Person struct {
	Id   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AllPerson struct {
	Per []*Person `protobuf:"bytes,1,rep,name=Per" json:"Per,omitempty"`
}

func (m *AllPerson) Reset()                    { *m = AllPerson{} }
func (m *AllPerson) String() string            { return proto.CompactTextString(m) }
func (*AllPerson) ProtoMessage()               {}
func (*AllPerson) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AllPerson) GetPer() []*Person {
	if m != nil {
		return m.Per
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "example.person")
	proto.RegisterType((*AllPerson)(nil), "example.all_person")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8,
	0x49, 0x55, 0xd2, 0xe1, 0x62, 0x83, 0x48, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0xb0, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0xfa, 0x5c, 0x5c, 0x89, 0x39, 0x39, 0xf1,
	0x50, 0x1d, 0x8a, 0x5c, 0xcc, 0x01, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xfc,
	0x7a, 0x50, 0x23, 0xf5, 0x20, 0xb2, 0x41, 0x20, 0xb9, 0x24, 0x36, 0xb0, 0x75, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf2, 0x55, 0x8c, 0xbb, 0x7e, 0x00, 0x00, 0x00,
}
