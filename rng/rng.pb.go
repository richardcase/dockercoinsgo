// Code generated by protoc-gen-go.
// source: rng.proto
// DO NOT EDIT!

/*
Package rng is a generated protocol buffer package.

It is generated from these files:
	rng.proto

It has these top-level messages:
	RngRequest
	RngResponse
*/
package rng

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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

type RngRequest struct {
	Length int32 `protobuf:"varint,1,opt,name=length" json:"length,omitempty"`
}

func (m *RngRequest) Reset()                    { *m = RngRequest{} }
func (m *RngRequest) String() string            { return proto.CompactTextString(m) }
func (*RngRequest) ProtoMessage()               {}
func (*RngRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RngResponse struct {
	Random string `protobuf:"bytes,1,opt,name=random" json:"random,omitempty"`
}

func (m *RngResponse) Reset()                    { *m = RngResponse{} }
func (m *RngResponse) String() string            { return proto.CompactTextString(m) }
func (*RngResponse) ProtoMessage()               {}
func (*RngResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*RngRequest)(nil), "rng.RngRequest")
	proto.RegisterType((*RngResponse)(nil), "rng.RngResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Rng service

type RngClient interface {
	GenerateRandom(ctx context.Context, in *RngRequest, opts ...grpc.CallOption) (*RngResponse, error)
}

type rngClient struct {
	cc *grpc.ClientConn
}

func NewRngClient(cc *grpc.ClientConn) RngClient {
	return &rngClient{cc}
}

func (c *rngClient) GenerateRandom(ctx context.Context, in *RngRequest, opts ...grpc.CallOption) (*RngResponse, error) {
	out := new(RngResponse)
	err := grpc.Invoke(ctx, "/rng.Rng/GenerateRandom", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rng service

type RngServer interface {
	GenerateRandom(context.Context, *RngRequest) (*RngResponse, error)
}

func RegisterRngServer(s *grpc.Server, srv RngServer) {
	s.RegisterService(&_Rng_serviceDesc, srv)
}

func _Rng_GenerateRandom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RngRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RngServer).GenerateRandom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rng.Rng/GenerateRandom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RngServer).GenerateRandom(ctx, req.(*RngRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rng_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rng.Rng",
	HandlerType: (*RngServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateRandom",
			Handler:    _Rng_GenerateRandom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("rng.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xca, 0x4b, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0xca, 0x4b, 0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc,
	0xcf, 0x2b, 0x86, 0x28, 0x51, 0x52, 0xe1, 0xe2, 0x0a, 0xca, 0x4b, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x12, 0xe3, 0x62, 0xcb, 0x49, 0xcd, 0x4b, 0x2f, 0xc9, 0x90, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0d, 0x82, 0xf2, 0x94, 0x54, 0xb9, 0xb8, 0xc1, 0xaa, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x41, 0xca, 0x8a, 0x12, 0xf3, 0x52, 0xf2, 0x73, 0xc1, 0xca, 0x38, 0x83, 0xa0, 0x3c, 0x23,
	0x3f, 0x2e, 0xe6, 0xa0, 0xbc, 0x74, 0x21, 0x77, 0x2e, 0x3e, 0xf7, 0xd4, 0xbc, 0xd4, 0xa2, 0xc4,
	0x92, 0xd4, 0x20, 0xb0, 0x84, 0x10, 0xbf, 0x1e, 0xc8, 0x51, 0x08, 0x8b, 0xa4, 0x04, 0x10, 0x02,
	0x10, 0x33, 0x95, 0x84, 0x9a, 0x2e, 0x3f, 0x99, 0xcc, 0xc4, 0xa3, 0xc4, 0xae, 0x5f, 0x66, 0xa8,
	0x5f, 0x94, 0x97, 0x6e, 0xc5, 0xa8, 0x95, 0xc4, 0x06, 0x76, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xde, 0x48, 0x71, 0x1a, 0xd3, 0x00, 0x00, 0x00,
}
