// Code generated by protoc-gen-go. DO NOT EDIT.
// source: master/master.proto

package master

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RegisterNodeRequest struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterNodeRequest) Reset()         { *m = RegisterNodeRequest{} }
func (m *RegisterNodeRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeRequest) ProtoMessage()    {}
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e524a825add940c, []int{0}
}

func (m *RegisterNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeRequest.Unmarshal(m, b)
}
func (m *RegisterNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeRequest.Marshal(b, m, deterministic)
}
func (m *RegisterNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeRequest.Merge(m, src)
}
func (m *RegisterNodeRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeRequest.Size(m)
}
func (m *RegisterNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeRequest proto.InternalMessageInfo

func (m *RegisterNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type RegisterNodeResponse struct {
	IsRegistered         bool     `protobuf:"varint,1,opt,name=isRegistered,proto3" json:"isRegistered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterNodeResponse) Reset()         { *m = RegisterNodeResponse{} }
func (m *RegisterNodeResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeResponse) ProtoMessage()    {}
func (*RegisterNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e524a825add940c, []int{1}
}

func (m *RegisterNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeResponse.Unmarshal(m, b)
}
func (m *RegisterNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeResponse.Marshal(b, m, deterministic)
}
func (m *RegisterNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeResponse.Merge(m, src)
}
func (m *RegisterNodeResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeResponse.Size(m)
}
func (m *RegisterNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeResponse proto.InternalMessageInfo

func (m *RegisterNodeResponse) GetIsRegistered() bool {
	if m != nil {
		return m.IsRegistered
	}
	return false
}

type Ping struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e524a825add940c, []int{2}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Pong struct {
	ServerName           string   `protobuf:"bytes,1,opt,name=serverName,proto3" json:"serverName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e524a825add940c, []int{3}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterNodeRequest)(nil), "master.RegisterNodeRequest")
	proto.RegisterType((*RegisterNodeResponse)(nil), "master.RegisterNodeResponse")
	proto.RegisterType((*Ping)(nil), "master.Ping")
	proto.RegisterType((*Pong)(nil), "master.Pong")
}

func init() {
	proto.RegisterFile("master/master.proto", fileDescriptor_3e524a825add940c)
}

var fileDescriptor_3e524a825add940c = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x4d, 0x2c, 0x2e,
	0x49, 0x2d, 0xd2, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x92,
	0x21, 0x97, 0x70, 0x50, 0x6a, 0x7a, 0x26, 0x88, 0xed, 0x97, 0x9f, 0x92, 0x1a, 0x94, 0x5a, 0x58,
	0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc5, 0xc5, 0x91, 0x97, 0x9f, 0x92, 0xea, 0x97, 0x98, 0x9b, 0x2a,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x2b, 0x59, 0x71, 0x89, 0xa0, 0x6a, 0x29, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x52, 0xe2, 0xe2, 0xc9, 0x2c, 0x86, 0xc9, 0xa4, 0xa6, 0x80, 0xf5,
	0x71, 0x04, 0xa1, 0x88, 0x29, 0x29, 0x71, 0xb1, 0x04, 0x64, 0xe6, 0xa5, 0xe3, 0x35, 0x5f, 0x8d,
	0x8b, 0x25, 0x20, 0x3f, 0x2f, 0x5d, 0x48, 0x8e, 0x8b, 0xab, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x08,
	0x49, 0x15, 0x92, 0x88, 0x51, 0x23, 0x23, 0x17, 0x9b, 0x2f, 0xd8, 0x17, 0x42, 0xde, 0x5c, 0x3c,
	0xc8, 0x4e, 0x12, 0x92, 0xd6, 0x83, 0x7a, 0x16, 0x8b, 0xdf, 0xa4, 0x64, 0xb0, 0x4b, 0x42, 0x7c,
	0xa1, 0xc4, 0x20, 0xa4, 0xc5, 0xc5, 0x05, 0x72, 0x63, 0x30, 0xd8, 0x26, 0x21, 0x1e, 0x98, 0x6a,
	0x90, 0x98, 0x14, 0x82, 0x97, 0x9f, 0x97, 0xae, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x4d, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0xdf, 0x7b, 0xf5, 0x64, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterClient interface {
	RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error)
	PingServer(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type masterClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterClient(cc grpc.ClientConnInterface) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error) {
	out := new(RegisterNodeResponse)
	err := c.cc.Invoke(ctx, "/master.Master/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) PingServer(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/master.Master/PingServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
type MasterServer interface {
	RegisterNode(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error)
	PingServer(context.Context, *Ping) (*Pong, error)
}

// UnimplementedMasterServer can be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (*UnimplementedMasterServer) RegisterNode(ctx context.Context, req *RegisterNodeRequest) (*RegisterNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (*UnimplementedMasterServer) PingServer(ctx context.Context, req *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingServer not implemented")
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).RegisterNode(ctx, req.(*RegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_PingServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).PingServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/PingServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).PingServer(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "master.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNode",
			Handler:    _Master_RegisterNode_Handler,
		},
		{
			MethodName: "PingServer",
			Handler:    _Master_PingServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master/master.proto",
}
