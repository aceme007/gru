// Code generated by protoc-gen-go.
// source: interact.proto
// DO NOT EDIT!

/*
Package interact is a generated protocol buffer package.

It is generated from these files:
	interact.proto

It has these top-level messages:
	ServerStatus
	ClientStatus
	Token
	Session
	Req
	Question
	Answer
	Response
	Status
*/
package interact

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

type ServerStatus struct {
	TimeLeft string `protobuf:"bytes,1,opt,name=timeLeft" json:"timeLeft,omitempty"`
}

func (m *ServerStatus) Reset()                    { *m = ServerStatus{} }
func (m *ServerStatus) String() string            { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()               {}
func (*ServerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClientStatus struct {
	CurrQuestion string `protobuf:"bytes,1,opt,name=currQuestion" json:"currQuestion,omitempty"`
	Token        string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *ClientStatus) Reset()                    { *m = ClientStatus{} }
func (m *ClientStatus) String() string            { return proto.CompactTextString(m) }
func (*ClientStatus) ProtoMessage()               {}
func (*ClientStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Token struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Session struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Req struct {
	Repeat   bool   `protobuf:"varint,1,opt,name=repeat" json:"repeat,omitempty"`
	Sid      string `protobuf:"bytes,2,opt,name=sid" json:"sid,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	TestType string `protobuf:"bytes,4,opt,name=testType" json:"testType,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Question struct {
	Id         string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Str        string    `protobuf:"bytes,2,opt,name=str" json:"str,omitempty"`
	Options    []*Answer `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
	IsMultiple bool      `protobuf:"varint,4,opt,name=isMultiple" json:"isMultiple,omitempty"`
	Positive   float32   `protobuf:"fixed32,6,opt,name=positive" json:"positive,omitempty"`
	Negative   float32   `protobuf:"fixed32,7,opt,name=negative" json:"negative,omitempty"`
	Totscore   float32   `protobuf:"fixed32,8,opt,name=totscore" json:"totscore,omitempty"`
}

func (m *Question) Reset()                    { *m = Question{} }
func (m *Question) String() string            { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()               {}
func (*Question) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Question) GetOptions() []*Answer {
	if m != nil {
		return m.Options
	}
	return nil
}

type Answer struct {
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Str string `protobuf:"bytes,2,opt,name=str" json:"str,omitempty"`
}

func (m *Answer) Reset()                    { *m = Answer{} }
func (m *Answer) String() string            { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()               {}
func (*Answer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type Response struct {
	Qid      string   `protobuf:"bytes,1,opt,name=qid" json:"qid,omitempty"`
	Aid      []string `protobuf:"bytes,2,rep,name=aid" json:"aid,omitempty"`
	Sid      string   `protobuf:"bytes,3,opt,name=sid" json:"sid,omitempty"`
	Token    string   `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	TestType string   `protobuf:"bytes,5,opt,name=testType" json:"testType,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Status struct {
	Status int64 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*ServerStatus)(nil), "interact.ServerStatus")
	proto.RegisterType((*ClientStatus)(nil), "interact.ClientStatus")
	proto.RegisterType((*Token)(nil), "interact.Token")
	proto.RegisterType((*Session)(nil), "interact.Session")
	proto.RegisterType((*Req)(nil), "interact.Req")
	proto.RegisterType((*Question)(nil), "interact.Question")
	proto.RegisterType((*Answer)(nil), "interact.Answer")
	proto.RegisterType((*Response)(nil), "interact.Response")
	proto.RegisterType((*Status)(nil), "interact.Status")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for GruQuiz service

type GruQuizClient interface {
	Authenticate(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Session, error)
	GetQuestion(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Question, error)
	SendAnswer(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Status, error)
	StreamChan(ctx context.Context, opts ...grpc.CallOption) (GruQuiz_StreamChanClient, error)
}

type gruQuizClient struct {
	cc *grpc.ClientConn
}

func NewGruQuizClient(cc *grpc.ClientConn) GruQuizClient {
	return &gruQuizClient{cc}
}

func (c *gruQuizClient) Authenticate(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := grpc.Invoke(ctx, "/interact.GruQuiz/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruQuizClient) GetQuestion(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Question, error) {
	out := new(Question)
	err := grpc.Invoke(ctx, "/interact.GruQuiz/GetQuestion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruQuizClient) SendAnswer(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/interact.GruQuiz/SendAnswer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruQuizClient) StreamChan(ctx context.Context, opts ...grpc.CallOption) (GruQuiz_StreamChanClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GruQuiz_serviceDesc.Streams[0], c.cc, "/interact.GruQuiz/StreamChan", opts...)
	if err != nil {
		return nil, err
	}
	x := &gruQuizStreamChanClient{stream}
	return x, nil
}

type GruQuiz_StreamChanClient interface {
	Send(*ClientStatus) error
	Recv() (*ServerStatus, error)
	grpc.ClientStream
}

type gruQuizStreamChanClient struct {
	grpc.ClientStream
}

func (x *gruQuizStreamChanClient) Send(m *ClientStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gruQuizStreamChanClient) Recv() (*ServerStatus, error) {
	m := new(ServerStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GruQuiz service

type GruQuizServer interface {
	Authenticate(context.Context, *Token) (*Session, error)
	GetQuestion(context.Context, *Req) (*Question, error)
	SendAnswer(context.Context, *Response) (*Status, error)
	StreamChan(GruQuiz_StreamChanServer) error
}

func RegisterGruQuizServer(s *grpc.Server, srv GruQuizServer) {
	s.RegisterService(&_GruQuiz_serviceDesc, srv)
}

func _GruQuiz_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruQuizServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.GruQuiz/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruQuizServer).Authenticate(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruQuiz_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruQuizServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.GruQuiz/GetQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruQuizServer).GetQuestion(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruQuiz_SendAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Response)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruQuizServer).SendAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.GruQuiz/SendAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruQuizServer).SendAnswer(ctx, req.(*Response))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruQuiz_StreamChan_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GruQuizServer).StreamChan(&gruQuizStreamChanServer{stream})
}

type GruQuiz_StreamChanServer interface {
	Send(*ServerStatus) error
	Recv() (*ClientStatus, error)
	grpc.ServerStream
}

type gruQuizStreamChanServer struct {
	grpc.ServerStream
}

func (x *gruQuizStreamChanServer) Send(m *ServerStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gruQuizStreamChanServer) Recv() (*ClientStatus, error) {
	m := new(ClientStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GruQuiz_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interact.GruQuiz",
	HandlerType: (*GruQuizServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _GruQuiz_Authenticate_Handler,
		},
		{
			MethodName: "GetQuestion",
			Handler:    _GruQuiz_GetQuestion_Handler,
		},
		{
			MethodName: "SendAnswer",
			Handler:    _GruQuiz_SendAnswer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamChan",
			Handler:       _GruQuiz_StreamChan_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xe3, 0xc6, 0x31, 0xd3, 0x50, 0xca, 0x0a, 0x05, 0xe3, 0x03, 0xaa, 0xf6, 0x14, 0xf5,
	0x50, 0xa1, 0xd0, 0x17, 0x28, 0x3d, 0x94, 0x03, 0x1c, 0xea, 0xf4, 0x05, 0x96, 0x74, 0xa0, 0x2b,
	0x52, 0xaf, 0xbb, 0x3b, 0x0e, 0x82, 0x67, 0xe4, 0x55, 0x78, 0x07, 0xf6, 0xcf, 0x3f, 0x44, 0x20,
	0x6e, 0xf3, 0xcd, 0x7c, 0xe3, 0x6f, 0xe6, 0xdb, 0x31, 0x1c, 0xcb, 0x9a, 0x50, 0x8b, 0x0d, 0x9d,
	0x37, 0x5a, 0x91, 0x62, 0x79, 0x87, 0xf9, 0x19, 0xcc, 0xd7, 0xa8, 0x77, 0xa8, 0xd7, 0x24, 0xa8,
	0x35, 0xac, 0x84, 0x9c, 0xe4, 0x03, 0x7e, 0xc0, 0xcf, 0x54, 0x24, 0xa7, 0xc9, 0xf2, 0x49, 0xd5,
	0x63, 0xfe, 0x1e, 0xe6, 0x57, 0x5b, 0x89, 0x35, 0x45, 0x2e, 0x87, 0xf9, 0xa6, 0xd5, 0xfa, 0xa6,
	0x45, 0x43, 0x52, 0xd5, 0x91, 0xff, 0x47, 0x8e, 0xbd, 0x80, 0x29, 0xa9, 0xaf, 0x58, 0x17, 0x13,
	0x5f, 0x0c, 0x80, 0xbf, 0x84, 0xe9, 0xad, 0x0b, 0xd8, 0x31, 0x4c, 0xe4, 0x5d, 0x6c, 0xb4, 0x11,
	0x7f, 0x05, 0xb3, 0x35, 0x1a, 0xe3, 0x3a, 0xf7, 0x4b, 0x02, 0xd2, 0x0a, 0x1f, 0xd9, 0x02, 0x32,
	0x8d, 0x0d, 0x8a, 0x30, 0x5e, 0x5e, 0x45, 0xc4, 0x4e, 0x20, 0x35, 0x96, 0x1f, 0x64, 0x5c, 0x38,
	0x48, 0xa7, 0x23, 0x69, 0xbf, 0xa0, 0x9d, 0xed, 0xf6, 0x7b, 0x83, 0xc5, 0x61, 0x5c, 0x30, 0x62,
	0xfe, 0x33, 0x81, 0xbc, 0x9f, 0x7c, 0x4f, 0xdf, 0x0b, 0x90, 0xee, 0x05, 0x48, 0xb3, 0x33, 0x98,
	0xa9, 0xc6, 0x71, 0x8d, 0x95, 0x48, 0x97, 0x47, 0xab, 0x93, 0xf3, 0xde, 0xe7, 0xcb, 0xda, 0x7c,
	0x43, 0x5d, 0x75, 0x04, 0xf6, 0x1a, 0x40, 0x9a, 0x8f, 0xed, 0x96, 0x64, 0xb3, 0x0d, 0xc2, 0x79,
	0x35, 0xca, 0xb8, 0xb1, 0x1a, 0x65, 0x24, 0xc9, 0x1d, 0x16, 0x99, 0xad, 0x4e, 0xaa, 0x1e, 0xbb,
	0x5a, 0x8d, 0x5f, 0x84, 0xaf, 0xcd, 0x42, 0xad, 0xc3, 0x7e, 0x1d, 0x45, 0x66, 0xa3, 0x34, 0x16,
	0x79, 0xa8, 0x75, 0xd8, 0xbe, 0x6d, 0x16, 0xc6, 0xf8, 0xff, 0x2e, 0x5c, 0x43, 0x5e, 0xa1, 0x69,
	0xec, 0xa8, 0xe8, 0xaa, 0x8f, 0x3d, 0xdd, 0x85, 0x2e, 0x23, 0xbc, 0xb9, 0xa9, 0xcb, 0x88, 0xf8,
	0x05, 0x9b, 0x49, 0xff, 0x62, 0xf7, 0xe1, 0xbf, 0xec, 0x9e, 0xee, 0xd9, 0x7d, 0x0a, 0x59, 0xbc,
	0x24, 0xfb, 0xa8, 0xc6, 0x47, 0x5e, 0x34, 0xad, 0x22, 0x5a, 0xfd, 0x4a, 0x60, 0x76, 0xad, 0xdb,
	0x9b, 0x56, 0xfe, 0x60, 0x17, 0x30, 0xbf, 0x6c, 0xe9, 0xde, 0x9e, 0x9f, 0xdc, 0x08, 0x42, 0xf6,
	0x6c, 0x30, 0xdb, 0xdf, 0x52, 0xf9, 0x7c, 0x48, 0xc4, 0x1b, 0xe2, 0x07, 0x6c, 0x05, 0x47, 0xd7,
	0x48, 0xfd, 0xa3, 0x3e, 0x1d, 0x38, 0xf6, 0x98, 0x4a, 0x36, 0xc0, 0x8e, 0x62, 0x7b, 0x2e, 0x00,
	0xd6, 0x58, 0xdf, 0x45, 0xef, 0xd8, 0xb8, 0x25, 0x38, 0x54, 0x8e, 0x1e, 0x3a, 0x6c, 0x60, 0xbb,
	0xde, 0xd9, 0x2e, 0xd2, 0x28, 0x1e, 0xae, 0xee, 0x45, 0xcd, 0x16, 0x03, 0x63, 0xfc, 0xcf, 0x94,
	0x8b, 0xf1, 0x90, 0xc3, 0x7f, 0xc7, 0x0f, 0x96, 0xc9, 0x9b, 0xe4, 0x53, 0xe6, 0x7f, 0xcf, 0xb7,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x44, 0x72, 0x3a, 0xb0, 0x03, 0x00, 0x00,
}
