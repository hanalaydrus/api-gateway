// Code generated by protoc-gen-go. DO NOT EDIT.
// source: semanticContract.proto

/*
Package semanticContract is a generated protocol buffer package.

It is generated from these files:
	semanticContract.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package semanticContract

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
	Response string `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "semanticContract.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "semanticContract.HelloReply")
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
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[0], c.cc, "/semanticContract.Greeter/SayHello", opts...)
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
	ServiceName: "semanticContract.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _Greeter_SayHello_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "semanticContract.proto",
}

func init() { proto.RegisterFile("semanticContract.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4e, 0xcd, 0x4d,
	0xcc, 0x2b, 0xc9, 0x4c, 0x76, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x40, 0x17, 0x57, 0x92, 0xe3, 0xe2, 0xf1, 0x48, 0xcd, 0xc9, 0xc9, 0x0f,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0d, 0x62, 0xca, 0x4c, 0x51, 0xd2, 0xe0, 0xe2, 0x82, 0xca, 0x17, 0xe4, 0x54, 0x0a,
	0x49, 0x71, 0x71, 0x14, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x82, 0xd5, 0x70, 0x06, 0xc1,
	0xf9, 0x46, 0xe1, 0x5c, 0xec, 0xee, 0x45, 0xa9, 0xa9, 0x25, 0xa9, 0x45, 0x42, 0x3e, 0x5c, 0x1c,
	0xc1, 0x89, 0x95, 0x60, 0x7d, 0x42, 0x72, 0x7a, 0x18, 0x6e, 0x41, 0xb6, 0x50, 0x4a, 0x06, 0xa7,
	0x7c, 0x41, 0x4e, 0xa5, 0x12, 0x83, 0x01, 0xa3, 0x93, 0x19, 0x97, 0x74, 0x66, 0xbe, 0x5e, 0x7a,
	0x51, 0x41, 0xb2, 0x5e, 0x6a, 0x45, 0x62, 0x6e, 0x41, 0x4e, 0x6a, 0xb1, 0x5e, 0x06, 0x48, 0x49,
	0x79, 0x7e, 0x51, 0x4e, 0x8a, 0x13, 0x3f, 0x58, 0x79, 0x38, 0x88, 0x1d, 0x00, 0xf2, 0x64, 0x00,
	0xe3, 0x22, 0x26, 0x66, 0x0f, 0x9f, 0xf0, 0x24, 0x36, 0xb0, 0x9f, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x07, 0xfd, 0x89, 0x50, 0x0d, 0x01, 0x00, 0x00,
}
