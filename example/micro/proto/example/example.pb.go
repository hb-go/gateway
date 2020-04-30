// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_example

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

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{5}
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

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{6}
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

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.example.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.example.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.example.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.example.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.example.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.example.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.example.Pong")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xf4, 0x30,
	0x10, 0x86, 0x37, 0x6d, 0xbf, 0x76, 0xbf, 0x39, 0xc8, 0x3a, 0x2c, 0xb2, 0x74, 0xd9, 0x65, 0xc9,
	0x41, 0x2b, 0x48, 0x5c, 0xf4, 0x27, 0x88, 0xe8, 0x45, 0x90, 0x7a, 0xf6, 0x10, 0x77, 0x43, 0x28,
	0x36, 0x49, 0x6d, 0xba, 0xa2, 0xbf, 0xdd, 0x8b, 0x34, 0x4d, 0x41, 0xa4, 0xc5, 0x53, 0x66, 0xf2,
	0xbc, 0xef, 0x64, 0x66, 0x08, 0x2c, 0xab, 0xda, 0x34, 0xe6, 0x52, 0x7c, 0x70, 0x55, 0x95, 0xa2,
	0x3f, 0x99, 0xbb, 0xc5, 0xb9, 0x34, 0x4c, 0x15, 0xbb, 0xda, 0x30, 0x5b, 0xbf, 0x33, 0xcf, 0xe8,
	0x12, 0x92, 0x07, 0x61, 0x2d, 0x97, 0x02, 0x67, 0x10, 0x5a, 0xfe, 0xb9, 0x20, 0x1b, 0x92, 0xfd,
	0xcf, 0xdb, 0x90, 0xae, 0x20, 0xc9, 0xc5, 0xdb, 0x41, 0xd8, 0x06, 0x11, 0x22, 0xcd, 0x95, 0xf0,
	0xd4, 0xc5, 0xf4, 0x02, 0xa6, 0xb9, 0xb0, 0x95, 0xd1, 0xd6, 0x99, 0x95, 0x95, 0xbd, 0x59, 0x59,
	0x89, 0x47, 0x10, 0x14, 0xfb, 0x45, 0xb0, 0x21, 0x59, 0x98, 0x07, 0xc5, 0x9e, 0x66, 0x30, 0x7b,
	0x6a, 0x6a, 0xc1, 0x55, 0xa1, 0x65, 0x5f, 0x75, 0x0e, 0xff, 0x76, 0xe6, 0xa0, 0x1b, 0xe7, 0x0b,
	0xf3, 0x2e, 0xa1, 0xe7, 0x70, 0xfc, 0x43, 0xe9, 0x1f, 0x18, 0x96, 0xae, 0x21, 0x7a, 0x2c, 0xb4,
	0xc4, 0x13, 0x88, 0x6d, 0x53, 0x9b, 0x57, 0xe1, 0xb1, 0xcf, 0x1c, 0x37, 0xe3, 0xfc, 0xea, 0x8b,
	0x40, 0x72, 0xdb, 0xad, 0x02, 0xef, 0x20, 0xba, 0xe1, 0x65, 0x89, 0x2b, 0x36, 0xb4, 0x29, 0xe6,
	0x7b, 0x4e, 0xd7, 0x63, 0xb8, 0x6b, 0x94, 0x4e, 0xf0, 0x19, 0xe2, 0xae, 0x7f, 0x3c, 0x1d, 0xd6,
	0xfe, 0xde, 0x43, 0x7a, 0xf6, 0xa7, 0xae, 0x2f, 0xbe, 0x25, 0x78, 0x0f, 0xd3, 0x76, 0x66, 0x37,
	0x57, 0x3a, 0x6c, 0x6c, 0x79, 0x3a, 0xc6, 0x8c, 0x96, 0x74, 0x92, 0x91, 0x2d, 0x79, 0x89, 0xdd,
	0xcf, 0xb8, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x52, 0x01, 0x76, 0x95, 0x38, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (Example_StreamClient, error)
	PingPong(ctx context.Context, opts ...grpc.CallOption) (Example_PingPongClient, error)
}

type exampleClient struct {
	cc *grpc.ClientConn
}

func NewExampleClient(cc *grpc.ClientConn) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.example.Example/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleClient) Stream(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (Example_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Example_serviceDesc.Streams[0], "/go.micro.srv.example.Example/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Example_StreamClient interface {
	Recv() (*StreamingResponse, error)
	grpc.ClientStream
}

type exampleStreamClient struct {
	grpc.ClientStream
}

func (x *exampleStreamClient) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleClient) PingPong(ctx context.Context, opts ...grpc.CallOption) (Example_PingPongClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Example_serviceDesc.Streams[1], "/go.micro.srv.example.Example/PingPong", opts...)
	if err != nil {
		return nil, err
	}
	x := &examplePingPongClient{stream}
	return x, nil
}

type Example_PingPongClient interface {
	Send(*Ping) error
	Recv() (*Pong, error)
	grpc.ClientStream
}

type examplePingPongClient struct {
	grpc.ClientStream
}

func (x *examplePingPongClient) Send(m *Ping) error {
	return x.ClientStream.SendMsg(m)
}

func (x *examplePingPongClient) Recv() (*Pong, error) {
	m := new(Pong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExampleServer is the server API for Example service.
type ExampleServer interface {
	Call(context.Context, *Request) (*Response, error)
	Stream(*StreamingRequest, Example_StreamServer) error
	PingPong(Example_PingPongServer) error
}

// UnimplementedExampleServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (*UnimplementedExampleServer) Call(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedExampleServer) Stream(req *StreamingRequest, srv Example_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (*UnimplementedExampleServer) PingPong(srv Example_PingPongServer) error {
	return status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.example.Example/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Example_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServer).Stream(m, &exampleStreamServer{stream})
}

type Example_StreamServer interface {
	Send(*StreamingResponse) error
	grpc.ServerStream
}

type exampleStreamServer struct {
	grpc.ServerStream
}

func (x *exampleStreamServer) Send(m *StreamingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Example_PingPong_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServer).PingPong(&examplePingPongServer{stream})
}

type Example_PingPongServer interface {
	Send(*Pong) error
	Recv() (*Ping, error)
	grpc.ServerStream
}

type examplePingPongServer struct {
	grpc.ServerStream
}

func (x *examplePingPongServer) Send(m *Pong) error {
	return x.ServerStream.SendMsg(m)
}

func (x *examplePingPongServer) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.example.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Example_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Example_Stream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PingPong",
			Handler:       _Example_PingPong_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/example/example.proto",
}
