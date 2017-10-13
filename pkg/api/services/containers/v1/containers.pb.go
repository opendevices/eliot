// Code generated by protoc-gen-go.
// source: services/containers/v1/containers.proto
// DO NOT EDIT!

/*
Package containers is a generated protocol buffer package.

It is generated from these files:
	services/containers/v1/containers.proto

It has these top-level messages:
	StdinStreamRequest
	StdoutStreamResponse
	SignalRequest
	SignalResponse
	Container
	Mount
	ContainerStatus
*/
package containers

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

type StdinStreamRequest struct {
	Input []byte `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (m *StdinStreamRequest) Reset()                    { *m = StdinStreamRequest{} }
func (m *StdinStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StdinStreamRequest) ProtoMessage()               {}
func (*StdinStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StdinStreamRequest) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

type StdoutStreamResponse struct {
	Output []byte `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	// Is this stderr(=true) or stdout(=false)
	Stderr bool `protobuf:"varint,2,opt,name=stderr" json:"stderr,omitempty"`
}

func (m *StdoutStreamResponse) Reset()                    { *m = StdoutStreamResponse{} }
func (m *StdoutStreamResponse) String() string            { return proto.CompactTextString(m) }
func (*StdoutStreamResponse) ProtoMessage()               {}
func (*StdoutStreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StdoutStreamResponse) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *StdoutStreamResponse) GetStderr() bool {
	if m != nil {
		return m.Stderr
	}
	return false
}

type SignalRequest struct {
	Namespace   string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	ContainerID string `protobuf:"bytes,2,opt,name=containerID" json:"containerID,omitempty"`
	Signal      int32  `protobuf:"varint,3,opt,name=signal" json:"signal,omitempty"`
}

func (m *SignalRequest) Reset()                    { *m = SignalRequest{} }
func (m *SignalRequest) String() string            { return proto.CompactTextString(m) }
func (*SignalRequest) ProtoMessage()               {}
func (*SignalRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SignalRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SignalRequest) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *SignalRequest) GetSignal() int32 {
	if m != nil {
		return m.Signal
	}
	return 0
}

type SignalResponse struct {
}

func (m *SignalResponse) Reset()                    { *m = SignalResponse{} }
func (m *SignalResponse) String() string            { return proto.CompactTextString(m) }
func (*SignalResponse) ProtoMessage()               {}
func (*SignalResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Container struct {
	Name       string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Image      string   `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	Tty        bool     `protobuf:"varint,3,opt,name=tty" json:"tty,omitempty"`
	WorkingDir string   `protobuf:"bytes,4,opt,name=workingDir" json:"workingDir,omitempty"`
	Args       []string `protobuf:"bytes,5,rep,name=args" json:"args,omitempty"`
	Env        []string `protobuf:"bytes,6,rep,name=env" json:"env,omitempty"`
	Mounts     []*Mount `protobuf:"bytes,7,rep,name=mounts" json:"mounts,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetTty() bool {
	if m != nil {
		return m.Tty
	}
	return false
}

func (m *Container) GetWorkingDir() string {
	if m != nil {
		return m.WorkingDir
	}
	return ""
}

func (m *Container) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Container) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Container) GetMounts() []*Mount {
	if m != nil {
		return m.Mounts
	}
	return nil
}

type Mount struct {
	Type        string   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Source      string   `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Destination string   `protobuf:"bytes,3,opt,name=destination" json:"destination,omitempty"`
	Options     []string `protobuf:"bytes,4,rep,name=options" json:"options,omitempty"`
}

func (m *Mount) Reset()                    { *m = Mount{} }
func (m *Mount) String() string            { return proto.CompactTextString(m) }
func (*Mount) ProtoMessage()               {}
func (*Mount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Mount) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Mount) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Mount) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *Mount) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

type ContainerStatus struct {
	ContainerID string `protobuf:"bytes,1,opt,name=containerID" json:"containerID,omitempty"`
	Image       string `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	State       string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
}

func (m *ContainerStatus) Reset()                    { *m = ContainerStatus{} }
func (m *ContainerStatus) String() string            { return proto.CompactTextString(m) }
func (*ContainerStatus) ProtoMessage()               {}
func (*ContainerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ContainerStatus) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *ContainerStatus) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ContainerStatus) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*StdinStreamRequest)(nil), "cand.services.containers.v1.StdinStreamRequest")
	proto.RegisterType((*StdoutStreamResponse)(nil), "cand.services.containers.v1.StdoutStreamResponse")
	proto.RegisterType((*SignalRequest)(nil), "cand.services.containers.v1.SignalRequest")
	proto.RegisterType((*SignalResponse)(nil), "cand.services.containers.v1.SignalResponse")
	proto.RegisterType((*Container)(nil), "cand.services.containers.v1.Container")
	proto.RegisterType((*Mount)(nil), "cand.services.containers.v1.Mount")
	proto.RegisterType((*ContainerStatus)(nil), "cand.services.containers.v1.ContainerStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Containers service

type ContainersClient interface {
	Attach(ctx context.Context, opts ...grpc.CallOption) (Containers_AttachClient, error)
	Signal(ctx context.Context, in *SignalRequest, opts ...grpc.CallOption) (*SignalResponse, error)
}

type containersClient struct {
	cc *grpc.ClientConn
}

func NewContainersClient(cc *grpc.ClientConn) ContainersClient {
	return &containersClient{cc}
}

func (c *containersClient) Attach(ctx context.Context, opts ...grpc.CallOption) (Containers_AttachClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Containers_serviceDesc.Streams[0], c.cc, "/cand.services.containers.v1.Containers/Attach", opts...)
	if err != nil {
		return nil, err
	}
	x := &containersAttachClient{stream}
	return x, nil
}

type Containers_AttachClient interface {
	Send(*StdinStreamRequest) error
	Recv() (*StdoutStreamResponse, error)
	grpc.ClientStream
}

type containersAttachClient struct {
	grpc.ClientStream
}

func (x *containersAttachClient) Send(m *StdinStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *containersAttachClient) Recv() (*StdoutStreamResponse, error) {
	m := new(StdoutStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *containersClient) Signal(ctx context.Context, in *SignalRequest, opts ...grpc.CallOption) (*SignalResponse, error) {
	out := new(SignalResponse)
	err := grpc.Invoke(ctx, "/cand.services.containers.v1.Containers/Signal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Containers service

type ContainersServer interface {
	Attach(Containers_AttachServer) error
	Signal(context.Context, *SignalRequest) (*SignalResponse, error)
}

func RegisterContainersServer(s *grpc.Server, srv ContainersServer) {
	s.RegisterService(&_Containers_serviceDesc, srv)
}

func _Containers_Attach_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ContainersServer).Attach(&containersAttachServer{stream})
}

type Containers_AttachServer interface {
	Send(*StdoutStreamResponse) error
	Recv() (*StdinStreamRequest, error)
	grpc.ServerStream
}

type containersAttachServer struct {
	grpc.ServerStream
}

func (x *containersAttachServer) Send(m *StdoutStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *containersAttachServer) Recv() (*StdinStreamRequest, error) {
	m := new(StdinStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Containers_Signal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainersServer).Signal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cand.services.containers.v1.Containers/Signal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainersServer).Signal(ctx, req.(*SignalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Containers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cand.services.containers.v1.Containers",
	HandlerType: (*ContainersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signal",
			Handler:    _Containers_Signal_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Attach",
			Handler:       _Containers_Attach_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services/containers/v1/containers.proto",
}

func init() { proto.RegisterFile("services/containers/v1/containers.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0xd8, 0x4d, 0xda, 0x4c, 0xf9, 0xa8, 0xac, 0x15, 0x8a, 0x0a, 0x42, 0x51, 0x2e, 0x44,
	0x45, 0x4a, 0xd8, 0x72, 0x83, 0x53, 0x69, 0x85, 0xc4, 0x81, 0x8b, 0x73, 0xe3, 0x82, 0xdc, 0xc4,
	0x4a, 0xad, 0x12, 0xdb, 0xd8, 0x93, 0x45, 0xfd, 0x89, 0xfc, 0x0c, 0xfe, 0x09, 0xb2, 0x93, 0xec,
	0x46, 0x6c, 0xb5, 0x70, 0x9b, 0xf7, 0x32, 0x7e, 0x6f, 0xbe, 0x02, 0xaf, 0x2d, 0x37, 0x1b, 0x51,
	0x73, 0x5b, 0xd6, 0x4a, 0x22, 0x13, 0x92, 0x1b, 0x5b, 0x6e, 0xd6, 0x33, 0x54, 0x68, 0xa3, 0x50,
	0x91, 0x17, 0x35, 0x93, 0x4d, 0x31, 0x65, 0x17, 0xb3, 0xef, 0x9b, 0x75, 0x76, 0x0e, 0xa4, 0xc2,
	0x46, 0xc8, 0x0a, 0x0d, 0x67, 0x1d, 0xe5, 0x3f, 0x7a, 0x6e, 0x91, 0xac, 0x20, 0x14, 0x52, 0xf7,
	0x98, 0x04, 0x69, 0x90, 0x3f, 0xa6, 0x03, 0xc8, 0x3e, 0xc1, 0xaa, 0xc2, 0x46, 0xf5, 0x38, 0x25,
	0x5b, 0xad, 0xa4, 0xe5, 0xe4, 0x39, 0x44, 0xaa, 0xc7, 0x5d, 0xfa, 0x88, 0x1c, 0x6f, 0xb1, 0xe1,
	0xc6, 0x24, 0x8f, 0xd2, 0x20, 0x3f, 0xa6, 0x23, 0xca, 0x5a, 0x78, 0x52, 0x89, 0x56, 0xb2, 0xef,
	0x93, 0xdd, 0x4b, 0x88, 0x25, 0xeb, 0xb8, 0xd5, 0xac, 0xe6, 0x5e, 0x23, 0xa6, 0x3b, 0x82, 0xa4,
	0x70, 0xb2, 0xad, 0xf9, 0xf3, 0xb5, 0xd7, 0x8a, 0xe9, 0x9c, 0xf2, 0x46, 0x5e, 0x30, 0x59, 0xa4,
	0x41, 0x1e, 0xd2, 0x11, 0x65, 0xa7, 0xf0, 0x74, 0x32, 0x1a, 0x4a, 0xcd, 0x7e, 0x05, 0x10, 0x5f,
	0x4d, 0x2f, 0x09, 0x81, 0xa5, 0xb3, 0x19, 0x2d, 0x7d, 0xec, 0x5b, 0xef, 0x58, 0xcb, 0x47, 0x9f,
	0x01, 0x90, 0x53, 0x58, 0x20, 0xde, 0x7b, 0xf9, 0x63, 0xea, 0x42, 0xf2, 0x0a, 0xe0, 0xa7, 0x32,
	0x77, 0x42, 0xb6, 0xd7, 0xc2, 0x24, 0x4b, 0x9f, 0x3c, 0x63, 0x9c, 0x36, 0x33, 0xad, 0x4d, 0xc2,
	0x74, 0xe1, 0xb4, 0x5d, 0xec, 0x54, 0xb8, 0xdc, 0x24, 0x91, 0xa7, 0x5c, 0x48, 0xde, 0x43, 0xd4,
	0xa9, 0x5e, 0xa2, 0x4d, 0x8e, 0xd2, 0x45, 0x7e, 0x72, 0x91, 0x15, 0x07, 0x96, 0x55, 0x7c, 0x71,
	0xa9, 0x74, 0x7c, 0x91, 0x29, 0x08, 0x3d, 0xe1, 0xac, 0xf0, 0x5e, 0x6f, 0xdb, 0x70, 0xb1, 0x1f,
	0x89, 0xea, 0x4d, 0x3d, 0xf5, 0x31, 0x22, 0x37, 0xcc, 0x86, 0x5b, 0x14, 0x92, 0xa1, 0x50, 0xd2,
	0x37, 0x14, 0xd3, 0x39, 0x45, 0x12, 0x38, 0x52, 0xda, 0x45, 0x36, 0x59, 0xfa, 0x42, 0x27, 0x98,
	0x7d, 0x83, 0x67, 0xdb, 0xd9, 0x55, 0xc8, 0xb0, 0xb7, 0x7f, 0xef, 0x26, 0xd8, 0xdf, 0xcd, 0xc3,
	0xf3, 0x5c, 0x41, 0x68, 0x91, 0x21, 0x1f, 0x0b, 0x18, 0xc0, 0xc5, 0xef, 0x00, 0x60, 0xeb, 0x60,
	0x89, 0x86, 0xe8, 0x12, 0x91, 0xd5, 0xb7, 0xa4, 0x3c, 0x38, 0x96, 0xfd, 0x03, 0x3e, 0x5b, 0xff,
	0xeb, 0xc1, 0xde, 0x15, 0xe7, 0xc1, 0xdb, 0x80, 0x30, 0x88, 0x86, 0x83, 0x21, 0xe7, 0x87, 0x05,
	0xe6, 0xe7, 0x7b, 0xf6, 0xe6, 0xbf, 0x72, 0x07, 0x9b, 0x8f, 0x57, 0x5f, 0x2f, 0x5b, 0x81, 0xb7,
	0xfd, 0x4d, 0x51, 0xab, 0xae, 0xe4, 0x46, 0x2a, 0xc6, 0x34, 0x2b, 0x6b, 0x26, 0x4b, 0x7d, 0xd7,
	0x96, 0x4c, 0x8b, 0xf2, 0xe1, 0x1f, 0xfb, 0xc3, 0x0e, 0xdd, 0x44, 0xfe, 0xcf, 0x7e, 0xf7, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x33, 0xdc, 0x3d, 0xad, 0x04, 0x04, 0x00, 0x00,
}
