// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/keyword_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [KeywordViewService.GetKeywordView][google.ads.googleads.v2.services.KeywordViewService.GetKeywordView].
type GetKeywordViewRequest struct {
	// The resource name of the keyword view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordViewRequest) Reset()         { *m = GetKeywordViewRequest{} }
func (m *GetKeywordViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordViewRequest) ProtoMessage()    {}
func (*GetKeywordViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5de781ce320be574, []int{0}
}

func (m *GetKeywordViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordViewRequest.Unmarshal(m, b)
}
func (m *GetKeywordViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordViewRequest.Marshal(b, m, deterministic)
}
func (m *GetKeywordViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordViewRequest.Merge(m, src)
}
func (m *GetKeywordViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordViewRequest.Size(m)
}
func (m *GetKeywordViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordViewRequest proto.InternalMessageInfo

func (m *GetKeywordViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordViewRequest)(nil), "google.ads.googleads.v2.services.GetKeywordViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/keyword_view_service.proto", fileDescriptor_5de781ce320be574)
}

var fileDescriptor_5de781ce320be574 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0x49, 0xfe, 0xf0, 0x07, 0x83, 0x7a, 0x08, 0x48, 0x4b, 0xf4, 0x50, 0x6a, 0x0f, 0xd2,
	0xc3, 0xae, 0x44, 0x51, 0xd8, 0xea, 0x21, 0xbd, 0x54, 0x10, 0xa4, 0x54, 0xc8, 0x41, 0x02, 0x25,
	0x26, 0x43, 0x08, 0x36, 0xd9, 0x9a, 0x49, 0x53, 0x44, 0xbc, 0xf8, 0x0a, 0xbe, 0x81, 0x47, 0x6f,
	0xbe, 0x46, 0xaf, 0xbe, 0x82, 0x5e, 0x7c, 0x08, 0x91, 0x74, 0xb3, 0x69, 0xab, 0x96, 0xde, 0x3e,
	0x66, 0xe6, 0xf7, 0xcd, 0xce, 0x97, 0x68, 0xad, 0x80, 0xf3, 0x60, 0x00, 0xd4, 0xf5, 0x91, 0x0a,
	0x99, 0xab, 0xcc, 0xa4, 0x08, 0x49, 0x16, 0x7a, 0x80, 0xf4, 0x06, 0xee, 0xc6, 0x3c, 0xf1, 0xfb,
	0x59, 0x08, 0xe3, 0x7e, 0x51, 0x25, 0xc3, 0x84, 0xa7, 0x5c, 0xaf, 0x09, 0x82, 0xb8, 0x3e, 0x92,
	0x12, 0x26, 0x99, 0x49, 0x24, 0x6c, 0x1c, 0x2e, 0xb3, 0x4f, 0x00, 0xf9, 0x28, 0xf9, 0xe9, 0x2f,
	0x7c, 0x8d, 0x1d, 0x49, 0x0d, 0x43, 0xea, 0xc6, 0x31, 0x4f, 0xdd, 0x34, 0xe4, 0x31, 0x16, 0xdd,
	0xca, 0x5c, 0xd7, 0x1b, 0x84, 0x10, 0xa7, 0xa2, 0x51, 0x3f, 0xd1, 0xb6, 0x3a, 0x90, 0x9e, 0x0b,
	0x3f, 0x3b, 0x84, 0x71, 0x0f, 0x6e, 0x47, 0x80, 0xa9, 0xbe, 0xab, 0x6d, 0xc8, 0x7d, 0xfd, 0xd8,
	0x8d, 0xa0, 0xaa, 0xd4, 0x94, 0xbd, 0xb5, 0xde, 0xba, 0x2c, 0x5e, 0xb8, 0x11, 0x98, 0x1f, 0x8a,
	0xa6, 0xcf, 0xb1, 0x97, 0xe2, 0x04, 0xfd, 0x55, 0xd1, 0x36, 0x17, 0x5d, 0xf5, 0x63, 0xb2, 0xea,
	0x6e, 0xf2, 0xe7, 0x3b, 0x0c, 0xb2, 0x14, 0x2c, 0xe3, 0x20, 0x73, 0x58, 0xfd, 0xe8, 0xf1, 0xed,
	0xfd, 0x49, 0xdd, 0xd7, 0x49, 0x9e, 0xd8, 0xfd, 0xc2, 0x09, 0xa7, 0xde, 0x08, 0x53, 0x1e, 0x41,
	0x82, 0xb4, 0x29, 0x23, 0xcc, 0x19, 0xa4, 0xcd, 0x07, 0x63, 0x7b, 0x62, 0x55, 0x67, 0xf6, 0x85,
	0x1a, 0x86, 0x48, 0x3c, 0x1e, 0xb5, 0xbf, 0x14, 0xad, 0xe1, 0xf1, 0x68, 0xe5, 0x0d, 0xed, 0xca,
	0xef, 0x34, 0xba, 0x79, 0xce, 0x5d, 0xe5, 0xea, 0xac, 0x80, 0x03, 0x3e, 0x70, 0xe3, 0x80, 0xf0,
	0x24, 0xa0, 0x01, 0xc4, 0xd3, 0xaf, 0x40, 0x67, 0xeb, 0x96, 0xff, 0x54, 0x2d, 0x29, 0x9e, 0xd5,
	0x7f, 0x1d, 0xcb, 0x7a, 0x51, 0x6b, 0x1d, 0x61, 0x68, 0xf9, 0x48, 0x84, 0xcc, 0x95, 0x6d, 0x92,
	0x62, 0x31, 0x4e, 0xe4, 0x88, 0x63, 0xf9, 0xe8, 0x94, 0x23, 0x8e, 0x6d, 0x3a, 0x72, 0xe4, 0x53,
	0x6d, 0x88, 0x3a, 0x63, 0x96, 0x8f, 0x8c, 0x95, 0x43, 0x8c, 0xd9, 0x26, 0x63, 0x72, 0xec, 0xfa,
	0xff, 0xf4, 0x9d, 0x07, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x14, 0xf8, 0xee, 0x62, 0xfb, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeywordViewServiceClient is the client API for KeywordViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordViewServiceClient interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error)
}

type keywordViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeywordViewServiceClient(cc *grpc.ClientConn) KeywordViewServiceClient {
	return &keywordViewServiceClient{cc}
}

func (c *keywordViewServiceClient) GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error) {
	out := new(resources.KeywordView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.KeywordViewService/GetKeywordView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordViewServiceServer is the server API for KeywordViewService service.
type KeywordViewServiceServer interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(context.Context, *GetKeywordViewRequest) (*resources.KeywordView, error)
}

// UnimplementedKeywordViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordViewServiceServer struct {
}

func (*UnimplementedKeywordViewServiceServer) GetKeywordView(ctx context.Context, req *GetKeywordViewRequest) (*resources.KeywordView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeywordView not implemented")
}

func RegisterKeywordViewServiceServer(s *grpc.Server, srv KeywordViewServiceServer) {
	s.RegisterService(&_KeywordViewService_serviceDesc, srv)
}

func _KeywordViewService_GetKeywordView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.KeywordViewService/GetKeywordView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, req.(*GetKeywordViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.KeywordViewService",
	HandlerType: (*KeywordViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordView",
			Handler:    _KeywordViewService_GetKeywordView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/keyword_view_service.proto",
}