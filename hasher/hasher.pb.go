// Code generated by protoc-gen-go.
// source: hasher.proto
// DO NOT EDIT!

/*
Package hasher is a generated protocol buffer package.

It is generated from these files:
	hasher.proto

It has these top-level messages:
	HashRequest
	HashResponse
*/
package hasher

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HashRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HashRequest) Reset()                    { *m = HashRequest{} }
func (m *HashRequest) String() string            { return proto.CompactTextString(m) }
func (*HashRequest) ProtoMessage()               {}
func (*HashRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HashResponse struct {
	HashedMessage string `protobuf:"bytes,1,opt,name=hashedMessage" json:"hashedMessage,omitempty"`
}

func (m *HashResponse) Reset()                    { *m = HashResponse{} }
func (m *HashResponse) String() string            { return proto.CompactTextString(m) }
func (*HashResponse) ProtoMessage()               {}
func (*HashResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HashRequest)(nil), "hasher.HashRequest")
	proto.RegisterType((*HashResponse)(nil), "hasher.HashResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Hasher service

type HasherClient interface {
	Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error)
}

type hasherClient struct {
	cc *grpc.ClientConn
}

func NewHasherClient(cc *grpc.ClientConn) HasherClient {
	return &hasherClient{cc}
}

func (c *hasherClient) Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := grpc.Invoke(ctx, "/hasher.Hasher/Hash", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hasher service

type HasherServer interface {
	Hash(context.Context, *HashRequest) (*HashResponse, error)
}

func RegisterHasherServer(s *grpc.Server, srv HasherServer) {
	s.RegisterService(&_Hasher_serviceDesc, srv)
}

func _Hasher_Hash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HasherServer).Hash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hasher.Hasher/Hash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HasherServer).Hash(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hasher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hasher.Hasher",
	HandlerType: (*HasherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hash",
			Handler:    _Hasher_Hash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("hasher.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x2c, 0xce,
	0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xd4, 0xb9, 0xb8,
	0x3d, 0x12, 0x8b, 0x33, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x73,
	0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x25,
	0x13, 0x2e, 0x1e, 0x88, 0xc2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x15, 0x2e, 0x5e, 0xb0,
	0x11, 0x29, 0xbe, 0x28, 0xea, 0x51, 0x05, 0x8d, 0x6c, 0xb9, 0xd8, 0x3c, 0xc0, 0x16, 0x09, 0x19,
	0x73, 0xb1, 0x80, 0x58, 0x42, 0xc2, 0x7a, 0x50, 0x77, 0x20, 0x59, 0x2b, 0x25, 0x82, 0x2a, 0x08,
	0xb1, 0x42, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x58, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba,
	0x12, 0xd3, 0xa6, 0xbc, 0x00, 0x00, 0x00,
}
