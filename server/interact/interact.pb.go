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
	Quiz
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

type QuizState int32

const (
	Quiz_DEMO_NOT_TAKEN QuizState = 0
	Quiz_DEMO_STARTED   QuizState = 1
	Quiz_TEST_NOT_TAKEN QuizState = 2
	Quiz_TEST_STARTED   QuizState = 3
	Quiz_TEST_FINISHED  QuizState = 4
)

var QuizState_name = map[int32]string{
	0: "DEMO_NOT_TAKEN",
	1: "DEMO_STARTED",
	2: "TEST_NOT_TAKEN",
	3: "TEST_STARTED",
	4: "TEST_FINISHED",
}
var QuizState_value = map[string]int32{
	"DEMO_NOT_TAKEN": 0,
	"DEMO_STARTED":   1,
	"TEST_NOT_TAKEN": 2,
	"TEST_STARTED":   3,
	"TEST_FINISHED":  4,
}

func (x QuizState) String() string {
	return proto.EnumName(QuizState_name, int32(x))
}
func (QuizState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type ServerStatus struct {
	TimeLeft string `protobuf:"bytes,1,opt,name=timeLeft" json:"timeLeft,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
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

type Quiz struct {
}

func (m *Quiz) Reset()                    { *m = Quiz{} }
func (m *Quiz) String() string            { return proto.CompactTextString(m) }
func (*Quiz) ProtoMessage()               {}
func (*Quiz) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Session struct {
	Id           string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	State        QuizState `protobuf:"varint,2,opt,name=state,enum=interact.QuizState" json:"state,omitempty"`
	TimeLeft     string    `protobuf:"bytes,3,opt,name=timeLeft" json:"timeLeft,omitempty"`
	TestDuration string    `protobuf:"bytes,4,opt,name=testDuration" json:"testDuration,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Req struct {
	Repeat bool   `protobuf:"varint,1,opt,name=repeat" json:"repeat,omitempty"`
	Sid    string `protobuf:"bytes,2,opt,name=sid" json:"sid,omitempty"`
	Token  string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
func (*Question) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*Answer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Response struct {
	Qid   string   `protobuf:"bytes,1,opt,name=qid" json:"qid,omitempty"`
	Aid   []string `protobuf:"bytes,2,rep,name=aid" json:"aid,omitempty"`
	Sid   string   `protobuf:"bytes,3,opt,name=sid" json:"sid,omitempty"`
	Token string   `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type Status struct {
	Status int64 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*ServerStatus)(nil), "interact.ServerStatus")
	proto.RegisterType((*ClientStatus)(nil), "interact.ClientStatus")
	proto.RegisterType((*Token)(nil), "interact.Token")
	proto.RegisterType((*Quiz)(nil), "interact.Quiz")
	proto.RegisterType((*Session)(nil), "interact.Session")
	proto.RegisterType((*Req)(nil), "interact.Req")
	proto.RegisterType((*Question)(nil), "interact.Question")
	proto.RegisterType((*Answer)(nil), "interact.Answer")
	proto.RegisterType((*Response)(nil), "interact.Response")
	proto.RegisterType((*Status)(nil), "interact.Status")
	proto.RegisterEnum("interact.QuizState", QuizState_name, QuizState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for GruQuiz service

type GruQuizClient interface {
	Authenticate(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Session, error)
	GetQuestion(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Question, error)
	SendAnswer(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Status, error)
	Ping(ctx context.Context, in *ClientStatus, opts ...grpc.CallOption) (*ServerStatus, error)
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

func (c *gruQuizClient) Ping(ctx context.Context, in *ClientStatus, opts ...grpc.CallOption) (*ServerStatus, error) {
	out := new(ServerStatus)
	err := grpc.Invoke(ctx, "/interact.GruQuiz/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GruQuiz service

type GruQuizServer interface {
	Authenticate(context.Context, *Token) (*Session, error)
	GetQuestion(context.Context, *Req) (*Question, error)
	SendAnswer(context.Context, *Response) (*Status, error)
	Ping(context.Context, *ClientStatus) (*ServerStatus, error)
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

func _GruQuiz_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruQuizServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.GruQuiz/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruQuizServer).Ping(ctx, req.(*ClientStatus))
	}
	return interceptor(ctx, in, info, handler)
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
		{
			MethodName: "Ping",
			Handler:    _GruQuiz_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("interact.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x9b, 0x38, 0x1f, 0x66, 0x9a, 0x86, 0x74, 0x55, 0x85, 0x28, 0x07, 0x54, 0xed, 0xa9,
	0xea, 0x21, 0x87, 0xd0, 0x03, 0xd7, 0x40, 0x42, 0x5b, 0x41, 0x53, 0x6a, 0x5b, 0x5c, 0x2b, 0x93,
	0x4e, 0xcb, 0x8a, 0x60, 0xbb, 0xbb, 0xeb, 0x22, 0x71, 0xe7, 0xed, 0x78, 0x0b, 0x5e, 0x84, 0xfd,
	0xb2, 0xbd, 0x44, 0x48, 0xdc, 0x76, 0xfe, 0x33, 0x3b, 0xf9, 0xcf, 0xcf, 0xb3, 0x81, 0x21, 0xcb,
	0x24, 0xf2, 0x74, 0x23, 0x67, 0x05, 0xcf, 0x65, 0x4e, 0xc2, 0x2a, 0xa6, 0x6f, 0x60, 0x10, 0x23,
	0x7f, 0x42, 0x1e, 0xcb, 0x54, 0x96, 0x82, 0x4c, 0x21, 0x94, 0xec, 0x1b, 0x7e, 0xc0, 0x7b, 0x39,
	0x69, 0x1d, 0xb7, 0x4e, 0x9e, 0x45, 0x75, 0x4c, 0xc6, 0xd0, 0x13, 0xa6, 0x6a, 0xd2, 0x36, 0x19,
	0x17, 0xd1, 0x0b, 0x18, 0xbc, 0xdd, 0x32, 0xcc, 0xa4, 0xeb, 0x41, 0x61, 0xb0, 0x29, 0x39, 0xbf,
	0x29, 0x51, 0x48, 0x96, 0x67, 0xae, 0xcf, 0x5f, 0x1a, 0x39, 0x82, 0xae, 0xcc, 0xbf, 0x62, 0xe6,
	0x5a, 0xd9, 0x80, 0xbe, 0x80, 0x6e, 0xa2, 0x0f, 0x64, 0x08, 0x6d, 0x76, 0xe7, 0x2e, 0xaa, 0x13,
	0xcd, 0xa0, 0x73, 0x53, 0xb2, 0x1f, 0xf4, 0x1e, 0xba, 0xfa, 0x47, 0x91, 0x10, 0x18, 0x2e, 0x57,
	0x57, 0xd7, 0xb7, 0xeb, 0xeb, 0xe4, 0x36, 0x59, 0xbc, 0x5f, 0xad, 0x47, 0x7b, 0x64, 0x04, 0x03,
	0xa3, 0xc5, 0xc9, 0x22, 0x4a, 0x56, 0xcb, 0x51, 0x4b, 0x57, 0x25, 0xab, 0x38, 0xf1, 0xaa, 0xda,
	0xba, 0xca, 0x68, 0x55, 0x55, 0x40, 0x0e, 0xe1, 0xc0, 0x28, 0xef, 0x2e, 0xd7, 0x97, 0xf1, 0x85,
	0x92, 0x3a, 0xf4, 0x67, 0x0b, 0xfa, 0x31, 0x0a, 0xa1, 0xad, 0xee, 0x78, 0x21, 0xa7, 0xce, 0x83,
	0xb1, 0x3e, 0x9c, 0x1f, 0xcd, 0x6a, 0xb8, 0xda, 0xe2, 0xcc, 0xe4, 0x22, 0x67, 0xd3, 0xc7, 0x19,
	0xec, 0xe0, 0x54, 0x98, 0xa4, 0xa2, 0xb1, 0x2c, 0x79, 0x6a, 0x30, 0x75, 0x2c, 0x26, 0x5f, 0xa3,
	0x2b, 0x08, 0x22, 0x7c, 0xd4, 0xe4, 0x39, 0x16, 0x98, 0xda, 0x6f, 0x12, 0x46, 0x2e, 0x52, 0xb3,
	0x04, 0x42, 0x79, 0xb3, 0x0c, 0xf5, 0xb1, 0xe1, 0x1a, 0xf8, 0x5c, 0x7f, 0xb5, 0x20, 0xac, 0xd1,
	0xef, 0xce, 0xa3, 0x9b, 0x48, 0x5e, 0x37, 0x91, 0x5c, 0x4d, 0xd8, 0xcf, 0x0b, 0x5d, 0x2b, 0x54,
	0x9b, 0xe0, 0x64, 0x7f, 0x3e, 0x6a, 0x66, 0x5c, 0x64, 0xe2, 0x3b, 0xf2, 0xa8, 0x2a, 0x20, 0x2f,
	0x01, 0x98, 0xb8, 0x2a, 0xb7, 0x92, 0x15, 0x5b, 0x34, 0x33, 0x84, 0x91, 0xa7, 0x68, 0x02, 0x45,
	0x2e, 0x98, 0x64, 0x4f, 0x38, 0xe9, 0xa9, 0x6c, 0x3b, 0xaa, 0x63, 0x9d, 0xcb, 0xf0, 0x21, 0x35,
	0xb9, 0xbe, 0xcd, 0x55, 0xb1, 0x21, 0x97, 0x4b, 0xb1, 0xc9, 0x39, 0x4e, 0x42, 0x9b, 0xab, 0x62,
	0x7a, 0x0a, 0x3d, 0x6b, 0xe3, 0xff, 0xb3, 0xd0, 0x4f, 0x10, 0x46, 0x28, 0x0a, 0x65, 0x15, 0x75,
	0xf6, 0xb1, 0x2e, 0xd7, 0x47, 0xad, 0xa4, 0x06, 0x60, 0xa0, 0x95, 0xd4, 0x75, 0x50, 0x4a, 0xf0,
	0x0f, 0xa4, 0x1d, 0x1f, 0xe9, 0x31, 0xf4, 0xdc, 0xba, 0x37, 0xcf, 0x42, 0x37, 0x0e, 0xaa, 0x67,
	0x31, 0xff, 0xad, 0x76, 0xe8, 0x9c, 0x97, 0x7a, 0x29, 0xc8, 0x19, 0x0c, 0x16, 0xa5, 0xfc, 0xa2,
	0xde, 0x08, 0xdb, 0xe8, 0xbd, 0x78, 0xde, 0x00, 0x35, 0x0b, 0x3f, 0x3d, 0x6c, 0x04, 0xb7, 0x77,
	0x74, 0x8f, 0xcc, 0x61, 0xff, 0x1c, 0x65, 0xfd, 0xe1, 0x0e, 0x9a, 0x1a, 0xb5, 0x14, 0x53, 0xe2,
	0x2f, 0x9e, 0x2d, 0x51, 0x77, 0xce, 0x00, 0x62, 0xcc, 0xee, 0x1c, 0x1f, 0xe2, 0x5f, 0xb1, 0x14,
	0xa6, 0xde, 0xc7, 0xb4, 0x13, 0xa8, 0x5b, 0xaf, 0xa1, 0xf3, 0x91, 0x65, 0x0f, 0x64, 0xdc, 0xe4,
	0xfc, 0x27, 0x3d, 0x1d, 0xfb, 0xf6, 0x9a, 0xbf, 0x0b, 0xba, 0xf7, 0xb9, 0x67, 0xfe, 0x51, 0x5e,
	0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x56, 0x46, 0xd5, 0x63, 0x04, 0x00, 0x00,
}
