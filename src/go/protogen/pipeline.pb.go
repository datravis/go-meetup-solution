// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pipeline.proto

package protogen

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

type IngestRequest struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Input                string   `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IngestRequest) Reset()         { *m = IngestRequest{} }
func (m *IngestRequest) String() string { return proto.CompactTextString(m) }
func (*IngestRequest) ProtoMessage()    {}
func (*IngestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_64a8fc2eae2348eb, []int{0}
}
func (m *IngestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngestRequest.Unmarshal(m, b)
}
func (m *IngestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngestRequest.Marshal(b, m, deterministic)
}
func (dst *IngestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngestRequest.Merge(dst, src)
}
func (m *IngestRequest) XXX_Size() int {
	return xxx_messageInfo_IngestRequest.Size(m)
}
func (m *IngestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IngestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IngestRequest proto.InternalMessageInfo

func (m *IngestRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *IngestRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

type IngestResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IngestResponse) Reset()         { *m = IngestResponse{} }
func (m *IngestResponse) String() string { return proto.CompactTextString(m) }
func (*IngestResponse) ProtoMessage()    {}
func (*IngestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_64a8fc2eae2348eb, []int{1}
}
func (m *IngestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngestResponse.Unmarshal(m, b)
}
func (m *IngestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngestResponse.Marshal(b, m, deterministic)
}
func (dst *IngestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngestResponse.Merge(dst, src)
}
func (m *IngestResponse) XXX_Size() int {
	return xxx_messageInfo_IngestResponse.Size(m)
}
func (m *IngestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IngestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IngestResponse proto.InternalMessageInfo

func (m *IngestResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*IngestRequest)(nil), "protogen.IngestRequest")
	proto.RegisterType((*IngestResponse)(nil), "protogen.IngestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PipelineServiceClient is the client API for PipelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PipelineServiceClient interface {
	Ingest(ctx context.Context, opts ...grpc.CallOption) (PipelineService_IngestClient, error)
}

type pipelineServiceClient struct {
	cc *grpc.ClientConn
}

func NewPipelineServiceClient(cc *grpc.ClientConn) PipelineServiceClient {
	return &pipelineServiceClient{cc}
}

func (c *pipelineServiceClient) Ingest(ctx context.Context, opts ...grpc.CallOption) (PipelineService_IngestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PipelineService_serviceDesc.Streams[0], "/protogen.PipelineService/Ingest", opts...)
	if err != nil {
		return nil, err
	}
	x := &pipelineServiceIngestClient{stream}
	return x, nil
}

type PipelineService_IngestClient interface {
	Send(*IngestRequest) error
	CloseAndRecv() (*IngestResponse, error)
	grpc.ClientStream
}

type pipelineServiceIngestClient struct {
	grpc.ClientStream
}

func (x *pipelineServiceIngestClient) Send(m *IngestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pipelineServiceIngestClient) CloseAndRecv() (*IngestResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(IngestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PipelineServiceServer is the server API for PipelineService service.
type PipelineServiceServer interface {
	Ingest(PipelineService_IngestServer) error
}

func RegisterPipelineServiceServer(s *grpc.Server, srv PipelineServiceServer) {
	s.RegisterService(&_PipelineService_serviceDesc, srv)
}

func _PipelineService_Ingest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PipelineServiceServer).Ingest(&pipelineServiceIngestServer{stream})
}

type PipelineService_IngestServer interface {
	SendAndClose(*IngestResponse) error
	Recv() (*IngestRequest, error)
	grpc.ServerStream
}

type pipelineServiceIngestServer struct {
	grpc.ServerStream
}

func (x *pipelineServiceIngestServer) SendAndClose(m *IngestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pipelineServiceIngestServer) Recv() (*IngestRequest, error) {
	m := new(IngestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PipelineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protogen.PipelineService",
	HandlerType: (*PipelineServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Ingest",
			Handler:       _PipelineService_Ingest_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pipeline.proto",
}

func init() { proto.RegisterFile("pipeline.proto", fileDescriptor_pipeline_64a8fc2eae2348eb) }

var fileDescriptor_pipeline_64a8fc2eae2348eb = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xc8, 0x2c, 0x48,
	0xcd, 0xc9, 0xcc, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xe9, 0xa9,
	0x79, 0x4a, 0xb6, 0x5c, 0xbc, 0x9e, 0x79, 0xe9, 0xa9, 0xc5, 0x25, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0xf9, 0xa5, 0x45, 0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x90, 0x08, 0x17, 0x6b, 0x66, 0x5e, 0x41, 0x69, 0x89, 0x04, 0x13,
	0x58, 0x18, 0xc2, 0x51, 0xd2, 0xe2, 0xe2, 0x83, 0x69, 0x2f, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x87, 0x19, 0x00, 0xe3, 0x1a, 0x05, 0x70,
	0xf1, 0x07, 0x40, 0x9d, 0x11, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xcb, 0xc5, 0x06,
	0xd1, 0x2e, 0x24, 0xae, 0x07, 0x73, 0x92, 0x1e, 0x8a, 0x7b, 0xa4, 0x24, 0x30, 0x25, 0x20, 0x36,
	0x69, 0x30, 0x26, 0xb1, 0x81, 0xa5, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x60, 0x8a,
	0xef, 0xdf, 0x00, 0x00, 0x00,
}
