// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package postgresServer

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

type GeneratorProjectsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneratorProjectsRequest) Reset()         { *m = GeneratorProjectsRequest{} }
func (m *GeneratorProjectsRequest) String() string { return proto.CompactTextString(m) }
func (*GeneratorProjectsRequest) ProtoMessage()    {}
func (*GeneratorProjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_4e8308ba520ffa2d, []int{0}
}
func (m *GeneratorProjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneratorProjectsRequest.Unmarshal(m, b)
}
func (m *GeneratorProjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneratorProjectsRequest.Marshal(b, m, deterministic)
}
func (dst *GeneratorProjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneratorProjectsRequest.Merge(dst, src)
}
func (m *GeneratorProjectsRequest) XXX_Size() int {
	return xxx_messageInfo_GeneratorProjectsRequest.Size(m)
}
func (m *GeneratorProjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneratorProjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GeneratorProjectsRequest proto.InternalMessageInfo

type ProjectDetails struct {
	ProjectId            int32    `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ProjectName          string   `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectDetails) Reset()         { *m = ProjectDetails{} }
func (m *ProjectDetails) String() string { return proto.CompactTextString(m) }
func (*ProjectDetails) ProtoMessage()    {}
func (*ProjectDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_4e8308ba520ffa2d, []int{1}
}
func (m *ProjectDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDetails.Unmarshal(m, b)
}
func (m *ProjectDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDetails.Marshal(b, m, deterministic)
}
func (dst *ProjectDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDetails.Merge(dst, src)
}
func (m *ProjectDetails) XXX_Size() int {
	return xxx_messageInfo_ProjectDetails.Size(m)
}
func (m *ProjectDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDetails proto.InternalMessageInfo

func (m *ProjectDetails) GetProjectId() int32 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *ProjectDetails) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

type GeneratorProjectsReply struct {
	Projects             []*ProjectDetails `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GeneratorProjectsReply) Reset()         { *m = GeneratorProjectsReply{} }
func (m *GeneratorProjectsReply) String() string { return proto.CompactTextString(m) }
func (*GeneratorProjectsReply) ProtoMessage()    {}
func (*GeneratorProjectsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_4e8308ba520ffa2d, []int{2}
}
func (m *GeneratorProjectsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneratorProjectsReply.Unmarshal(m, b)
}
func (m *GeneratorProjectsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneratorProjectsReply.Marshal(b, m, deterministic)
}
func (dst *GeneratorProjectsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneratorProjectsReply.Merge(dst, src)
}
func (m *GeneratorProjectsReply) XXX_Size() int {
	return xxx_messageInfo_GeneratorProjectsReply.Size(m)
}
func (m *GeneratorProjectsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneratorProjectsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GeneratorProjectsReply proto.InternalMessageInfo

func (m *GeneratorProjectsReply) GetProjects() []*ProjectDetails {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*GeneratorProjectsRequest)(nil), "postgresServer.GeneratorProjectsRequest")
	proto.RegisterType((*ProjectDetails)(nil), "postgresServer.ProjectDetails")
	proto.RegisterType((*GeneratorProjectsReply)(nil), "postgresServer.GeneratorProjectsReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiClient interface {
	GeneratorProjects(ctx context.Context, in *GeneratorProjectsRequest, opts ...grpc.CallOption) (*GeneratorProjectsReply, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) GeneratorProjects(ctx context.Context, in *GeneratorProjectsRequest, opts ...grpc.CallOption) (*GeneratorProjectsReply, error) {
	out := new(GeneratorProjectsReply)
	err := c.cc.Invoke(ctx, "/postgresServer.Api/GeneratorProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
type ApiServer interface {
	GeneratorProjects(context.Context, *GeneratorProjectsRequest) (*GeneratorProjectsReply, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_GeneratorProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratorProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GeneratorProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postgresServer.Api/GeneratorProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GeneratorProjects(ctx, req.(*GeneratorProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "postgresServer.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeneratorProjects",
			Handler:    _Api_GeneratorProjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_4e8308ba520ffa2d) }

var fileDescriptor_api_4e8308ba520ffa2d = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2b, 0xc8, 0x2f, 0x2e, 0x49, 0x2f, 0x4a, 0x2d, 0x0e,
	0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x52, 0x92, 0xe2, 0x92, 0x70, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c,
	0xc9, 0x2f, 0x0a, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x51, 0x0a, 0xe2, 0xe2, 0x83, 0x0a, 0xb9, 0xa4, 0x96, 0x24, 0x66, 0xe6, 0x14, 0x0b, 0xc9,
	0x72, 0x71, 0x15, 0x40, 0x44, 0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x38,
	0xa1, 0x22, 0x9e, 0x29, 0x42, 0x8a, 0x5c, 0x3c, 0x30, 0xe9, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x6e, 0xa8, 0x98, 0x5f, 0x62, 0x6e, 0xaa, 0x52, 0x08, 0x97, 0x18,
	0x16, 0xfb, 0x0a, 0x72, 0x2a, 0x85, 0xac, 0xb8, 0x38, 0xa0, 0x0a, 0x8b, 0x25, 0x18, 0x15, 0x98,
	0x35, 0xb8, 0x8d, 0xe4, 0xf4, 0x50, 0x1d, 0xab, 0x87, 0xea, 0x9a, 0x20, 0xb8, 0x7a, 0xa3, 0x1c,
	0x2e, 0x66, 0xc7, 0x82, 0x4c, 0xa1, 0x54, 0x2e, 0x41, 0x0c, 0xc3, 0x85, 0x34, 0xd0, 0x4d, 0xc1,
	0xe5, 0x5f, 0x29, 0x35, 0x22, 0x54, 0x16, 0xe4, 0x54, 0x26, 0xb1, 0x81, 0x83, 0xd2, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x57, 0x79, 0x14, 0x9b, 0x57, 0x01, 0x00, 0x00,
}
