// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rabbit-test.proto

/*
Package rabbittest is a generated protocol buffer package.

It is generated from these files:
	rabbit-test.proto

It has these top-level messages:
	RegisterNodeRequest
	RegisterNodeReply
*/
package rabbittest

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

type RegisterNodeRequest struct {
	NodeName string `protobuf:"bytes,1,opt,name=nodeName" json:"nodeName,omitempty"`
}

func (m *RegisterNodeRequest) Reset()                    { *m = RegisterNodeRequest{} }
func (m *RegisterNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterNodeRequest) ProtoMessage()               {}
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type RegisterNodeReply struct {
	IsRegistered bool `protobuf:"varint,1,opt,name=isRegistered" json:"isRegistered,omitempty"`
}

func (m *RegisterNodeReply) Reset()                    { *m = RegisterNodeReply{} }
func (m *RegisterNodeReply) String() string            { return proto.CompactTextString(m) }
func (*RegisterNodeReply) ProtoMessage()               {}
func (*RegisterNodeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterNodeReply) GetIsRegistered() bool {
	if m != nil {
		return m.IsRegistered
	}
	return false
}

func init() {
	proto.RegisterType((*RegisterNodeRequest)(nil), "rabbittest.RegisterNodeRequest")
	proto.RegisterType((*RegisterNodeReply)(nil), "rabbittest.RegisterNodeReply")
}

func init() { proto.RegisterFile("rabbit-test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4a, 0x4c, 0x4a,
	0xca, 0x2c, 0xd1, 0x2d, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82,
	0x08, 0x81, 0x44, 0x94, 0x0c, 0xb9, 0x84, 0x83, 0x52, 0xd3, 0x33, 0x8b, 0x4b, 0x52, 0x8b, 0xfc,
	0xf2, 0x53, 0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4, 0xb8, 0x38, 0xf2, 0xf2,
	0x53, 0x52, 0xfd, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25,
	0x73, 0x2e, 0x41, 0x54, 0x2d, 0x05, 0x39, 0x95, 0x42, 0x4a, 0x5c, 0x3c, 0x99, 0xc5, 0x30, 0xe1,
	0xd4, 0x14, 0xb0, 0x26, 0x8e, 0x20, 0x14, 0x31, 0xa3, 0x44, 0x2e, 0xde, 0xe0, 0xd4, 0xa2, 0xb2,
	0xd4, 0x22, 0x10, 0x99, 0x99, 0x9c, 0x2a, 0x14, 0xc0, 0xc5, 0x83, 0x6c, 0x92, 0x90, 0xbc, 0x1e,
	0xc2, 0x65, 0x7a, 0x58, 0x9c, 0x25, 0x25, 0x8b, 0x5b, 0x41, 0x41, 0x4e, 0xa5, 0x12, 0x43, 0x12,
	0x1b, 0xd8, 0x87, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x57, 0x37, 0x8d, 0xf6, 0x00,
	0x00, 0x00,
}
