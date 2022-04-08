// Code generated by protoc-gen-go. DO NOT EDIT.
// source: broker/broker.proto

package broker

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a300fef54c4624, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type PublishRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Message              *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a300fef54c4624, []int{1}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type SubscribeRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Queue                string   `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a300fef54c4624, []int{2}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SubscribeRequest) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

type Message struct {
	Header               map[string]string `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a300fef54c4624, []int{3}
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

func (m *Message) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "broker.Empty")
	proto.RegisterType((*PublishRequest)(nil), "broker.PublishRequest")
	proto.RegisterType((*SubscribeRequest)(nil), "broker.SubscribeRequest")
	proto.RegisterType((*Message)(nil), "broker.Message")
	proto.RegisterMapType((map[string]string)(nil), "broker.Message.HeaderEntry")
}

func init() { proto.RegisterFile("broker/broker.proto", fileDescriptor_09a300fef54c4624) }

var fileDescriptor_09a300fef54c4624 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0xad, 0x5b, 0x9a, 0xa8, 0x57, 0x3e, 0x2a, 0x53, 0xa1, 0xa8, 0x2c, 0x55, 0xa6, 0x32, 0x90,
	0xa0, 0x64, 0x81, 0x22, 0x31, 0x54, 0xaa, 0xc4, 0x82, 0x04, 0x66, 0x63, 0x8b, 0x53, 0xab, 0x89,
	0xda, 0xe0, 0xd4, 0xb1, 0x2b, 0xf2, 0x23, 0xf8, 0xcf, 0xa8, 0xb6, 0xc3, 0x47, 0x06, 0x16, 0xfb,
	0xde, 0xbd, 0xd3, 0xbb, 0xa7, 0x77, 0x70, 0x4e, 0x05, 0xdf, 0x30, 0x11, 0x9a, 0x2f, 0x28, 0x05,
	0x97, 0x1c, 0x3b, 0x06, 0xf9, 0x2e, 0xf4, 0x97, 0x45, 0x29, 0x6b, 0xff, 0x05, 0x4e, 0x9f, 0x15,
	0xdd, 0xe6, 0x55, 0x46, 0xd8, 0x4e, 0xb1, 0x4a, 0xe2, 0x31, 0xf4, 0x25, 0x2f, 0xf3, 0xd4, 0x43,
	0x53, 0x34, 0x1b, 0x10, 0x03, 0xf0, 0x15, 0xb8, 0x05, 0xab, 0xaa, 0x64, 0xcd, 0xbc, 0xee, 0x14,
	0xcd, 0x86, 0xd1, 0x59, 0x60, 0x85, 0x9f, 0x4c, 0x9b, 0x34, 0xbc, 0xff, 0x00, 0xa3, 0x57, 0x45,
	0xab, 0x54, 0xe4, 0x94, 0xfd, 0x2f, 0x3a, 0x86, 0xfe, 0x4e, 0x31, 0x65, 0x24, 0x07, 0xc4, 0x00,
	0xff, 0x13, 0x81, 0x6b, 0x45, 0x71, 0x0c, 0x4e, 0xc6, 0x92, 0x15, 0x13, 0x1e, 0x9a, 0xf6, 0x66,
	0xc3, 0xe8, 0xb2, 0xb5, 0x35, 0x78, 0xd4, 0xec, 0xf2, 0x5d, 0x8a, 0x9a, 0xd8, 0x51, 0x8c, 0xe1,
	0x88, 0xf2, 0x55, 0xad, 0x55, 0x8f, 0x89, 0xae, 0x27, 0x77, 0x30, 0xfc, 0x35, 0x8a, 0x47, 0xd0,
	0xdb, 0xb0, 0xda, 0xba, 0x39, 0x94, 0x07, 0x2f, 0xfb, 0x64, 0xfb, 0xe3, 0x45, 0x83, 0x79, 0xf7,
	0x16, 0x45, 0x1f, 0xe0, 0x2c, 0xf4, 0x52, 0x1c, 0x81, 0x6b, 0xc3, 0xc2, 0x17, 0x8d, 0x91, 0xbf,
	0xe9, 0x4d, 0x4e, 0x9a, 0xbe, 0x89, 0xb7, 0x83, 0xe7, 0x30, 0xf8, 0x4e, 0x03, 0x7b, 0x0d, 0xdb,
	0x0e, 0x68, 0xd2, 0x8e, 0xd3, 0xef, 0xdc, 0xa0, 0x45, 0xf8, 0x76, 0xbd, 0xce, 0x65, 0xa6, 0x68,
	0x90, 0xf2, 0x22, 0x2c, 0xf2, 0x54, 0x70, 0xfb, 0xee, 0xe3, 0x50, 0x1f, 0xd4, 0x5e, 0xf7, 0xde,
	0x7c, 0xd4, 0xd1, 0xcd, 0xf8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x52, 0xb5, 0x7a, 0xfc, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BrokerClient is the client API for Broker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BrokerClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*Empty, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Broker_SubscribeClient, error)
}

type brokerClient struct {
	cc *grpc.ClientConn
}

func NewBrokerClient(cc *grpc.ClientConn) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/broker.Broker/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Broker_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Broker_serviceDesc.Streams[0], "/broker.Broker/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Broker_SubscribeClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type brokerSubscribeClient struct {
	grpc.ClientStream
}

func (x *brokerSubscribeClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BrokerServer is the server API for Broker service.
type BrokerServer interface {
	Publish(context.Context, *PublishRequest) (*Empty, error)
	Subscribe(*SubscribeRequest, Broker_SubscribeServer) error
}

func RegisterBrokerServer(s *grpc.Server, srv BrokerServer) {
	s.RegisterService(&_Broker_serviceDesc, srv)
}

func _Broker_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/broker.Broker/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BrokerServer).Subscribe(m, &brokerSubscribeServer{stream})
}

type Broker_SubscribeServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type brokerSubscribeServer struct {
	grpc.ServerStream
}

func (x *brokerSubscribeServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _Broker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "broker.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Broker_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Broker_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "broker/broker.proto",
}
