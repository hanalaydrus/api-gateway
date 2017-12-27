// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld2.proto

/*
Package helloworld2 is a generated protocol buffer package.

It is generated from these files:
	helloworld2.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package helloworld2

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
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Volume   string `protobuf:"bytes,1,opt,name=volume" json:"volume,omitempty"`
	Density  string `protobuf:"bytes,2,opt,name=density" json:"density,omitempty"`
	Semantic string `protobuf:"bytes,3,opt,name=semantic" json:"semantic,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *HelloReply) GetDensity() string {
	if m != nil {
		return m.Density
	}
	return ""
}

func (m *HelloReply) GetSemantic() string {
	if m != nil {
		return m.Semantic
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld2.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld2.HelloReply")
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
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[0], c.cc, "/helloworld2.Greeter/SayHello", opts...)
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
	ServiceName: "helloworld2.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _Greeter_SayHello_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "helloworld2.proto",
}

func init() { proto.RegisterFile("helloworld2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0x31, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x52, 0x52, 0xe2, 0xe2, 0xf1, 0x00, 0x71, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c,
	0xa5, 0x28, 0x2e, 0x2e, 0xa8, 0x9a, 0x82, 0x9c, 0x4a, 0x21, 0x31, 0x2e, 0xb6, 0xb2, 0xfc, 0x9c,
	0x52, 0xb8, 0x1a, 0x28, 0x4f, 0x48, 0x82, 0x8b, 0x3d, 0x25, 0x35, 0xaf, 0x38, 0xb3, 0xa4, 0x52,
	0x82, 0x09, 0x2c, 0x01, 0xe3, 0x0a, 0x49, 0x71, 0x71, 0x14, 0xa7, 0xe6, 0x26, 0xe6, 0x95, 0x64,
	0x26, 0x4b, 0x30, 0x83, 0xa5, 0xe0, 0x7c, 0x23, 0x5f, 0x2e, 0x76, 0xf7, 0xa2, 0xd4, 0xd4, 0x92,
	0xd4, 0x22, 0x21, 0x27, 0x2e, 0x8e, 0xe0, 0xc4, 0x4a, 0xb0, 0x4d, 0x42, 0x92, 0x7a, 0xc8, 0xee,
	0x46, 0x76, 0xa1, 0x94, 0x38, 0x36, 0xa9, 0x82, 0x9c, 0x4a, 0x25, 0x06, 0x03, 0xc6, 0x24, 0x36,
	0xb0, 0x17, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x86, 0xb8, 0x02, 0xf7, 0x00, 0x00,
	0x00,
}