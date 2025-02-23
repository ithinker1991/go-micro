// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/asim/go-micro/examples/v4/event/srv/proto/pubsub.proto

/*
Package pubsub is a generated protocol buffer package.

It is generated from these files:
	github.com/asim/go-micro/examples/v4/event/srv/proto/pubsub.proto

It has these top-level messages:
	Event
*/
package pubsub

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

// Example message
type Event struct {
	// unique id
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// unix timestamp
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// message
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
}

func init() {
	proto.RegisterFile("github.com/asim/go-micro/examples/v4/event/srv/proto/pubsub.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x4f, 0xad, 0x48,
	0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2f, 0x2e, 0x2a, 0xd3,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x28, 0x4d, 0x2a, 0x2e, 0x4d, 0xd2, 0x03, 0x73, 0x94,
	0xfc, 0xb9, 0x58, 0x5d, 0x41, 0xf2, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x32, 0x5c, 0x9c, 0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25,
	0x89, 0xb9, 0x05, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x08, 0x01, 0x21, 0x09, 0x2e, 0xf6,
	0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x66, 0xb0, 0x16, 0x18, 0x37, 0x89, 0x0d, 0x6c,
	0xae, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x89, 0xbc, 0xdb, 0x74, 0x91, 0x00, 0x00, 0x00,
}
