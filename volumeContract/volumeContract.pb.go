// Code generated by protoc-gen-go. DO NOT EDIT.
// source: volumeContract.proto

/*
Package volumeContract is a generated protocol buffer package.

It is generated from these files:
	volumeContract.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package volumeContract

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

// The request message containing the user's name.
type HelloRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// The response message containing the greetings
type HelloReply struct {
	Timestamp  string  `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Volume     int32   `protobuf:"varint,2,opt,name=volume" json:"volume,omitempty"`
	Percentage float32 `protobuf:"fixed32,3,opt,name=percentage" json:"percentage,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *HelloReply) GetVolume() int32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *HelloReply) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "volumeContract.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "volumeContract.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[0], c.cc, "/volumeContract.Greeter/SayHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_SayHelloClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(*HelloRequest, Greeter_SayHelloServer) error
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHello(m, &greeterSayHelloServer{stream})
}

type Greeter_SayHelloServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterSayHelloServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "volumeContract.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _Greeter_SayHello_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "volumeContract.proto",
}

func init() { proto.RegisterFile("volumeContract.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x4d, 0x8a, 0xd5, 0x5e, 0xa4, 0x42, 0x10, 0x19, 0x6a, 0x29, 0x65, 0x56, 0x5d, 0x05,
	0x51, 0xf0, 0x01, 0xea, 0xc2, 0x59, 0xb8, 0x28, 0xe9, 0xa2, 0xeb, 0x74, 0xe6, 0x52, 0x03, 0x49,
	0x13, 0x33, 0x77, 0xd4, 0x79, 0x1d, 0x9f, 0x54, 0x26, 0x0e, 0x38, 0x0a, 0xee, 0x92, 0x73, 0xbe,
	0xfb, 0x77, 0xe0, 0xea, 0xcd, 0xdb, 0xc6, 0xe1, 0xa3, 0x3f, 0x52, 0xd4, 0x25, 0xc9, 0x10, 0x3d,
	0x79, 0x31, 0xfd, 0xad, 0xe6, 0x0b, 0xb8, 0x28, 0xd0, 0x5a, 0xaf, 0xf0, 0xb5, 0xc1, 0x9a, 0xc4,
	0x14, 0xb8, 0xa9, 0x32, 0xb6, 0x64, 0xab, 0x53, 0xc5, 0x4d, 0x95, 0xef, 0x01, 0x7a, 0x3f, 0xd8,
	0x56, 0xcc, 0x61, 0x42, 0xc6, 0x61, 0x4d, 0xda, 0x85, 0x04, 0x4d, 0xd4, 0x8f, 0x20, 0xae, 0x61,
	0xfc, 0xdd, 0x3d, 0xe3, 0xa9, 0xbe, 0xff, 0x89, 0x05, 0x40, 0xc0, 0x58, 0xe2, 0x91, 0xf4, 0x01,
	0xb3, 0xd1, 0x92, 0xad, 0xb8, 0x1a, 0x28, 0x77, 0x5b, 0x38, 0x7b, 0x8a, 0x88, 0x84, 0x51, 0x14,
	0x70, 0xbe, 0xd5, 0x6d, 0x9a, 0x28, 0xe6, 0xf2, 0xcf, 0x05, 0xc3, 0x45, 0x67, 0xb3, 0x7f, 0xdc,
	0x60, 0xdb, 0xfc, 0xe4, 0x96, 0xad, 0x1f, 0xe0, 0xc6, 0x78, 0x79, 0x88, 0xa1, 0x94, 0xf8, 0xa1,
	0x5d, 0xb0, 0x58, 0xcb, 0x97, 0x0e, 0x79, 0xf7, 0xd1, 0x56, 0xeb, 0xcb, 0x84, 0xef, 0xba, 0xf7,
	0xa6, 0x0b, 0x66, 0xc3, 0x3e, 0xf9, 0xa8, 0x78, 0xde, 0xed, 0xc7, 0x29, 0xa7, 0xfb, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1d, 0xc4, 0x0e, 0xf8, 0x3f, 0x01, 0x00, 0x00,
}