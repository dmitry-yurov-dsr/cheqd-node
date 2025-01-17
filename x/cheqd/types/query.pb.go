// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cheqd/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// this line is used by starport scaffolding # 3
type QueryGetNymRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetNymRequest) Reset()         { *m = QueryGetNymRequest{} }
func (m *QueryGetNymRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetNymRequest) ProtoMessage()    {}
func (*QueryGetNymRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf3698559b3858f5, []int{0}
}
func (m *QueryGetNymRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetNymRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetNymRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetNymRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetNymRequest.Merge(m, src)
}
func (m *QueryGetNymRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetNymRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetNymRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetNymRequest proto.InternalMessageInfo

func (m *QueryGetNymRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetNymResponse struct {
	Nym *Nym `protobuf:"bytes,1,opt,name=Nym,proto3" json:"Nym,omitempty"`
}

func (m *QueryGetNymResponse) Reset()         { *m = QueryGetNymResponse{} }
func (m *QueryGetNymResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetNymResponse) ProtoMessage()    {}
func (*QueryGetNymResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf3698559b3858f5, []int{1}
}
func (m *QueryGetNymResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetNymResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetNymResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetNymResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetNymResponse.Merge(m, src)
}
func (m *QueryGetNymResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetNymResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetNymResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetNymResponse proto.InternalMessageInfo

func (m *QueryGetNymResponse) GetNym() *Nym {
	if m != nil {
		return m.Nym
	}
	return nil
}

type QueryAllNymRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllNymRequest) Reset()         { *m = QueryAllNymRequest{} }
func (m *QueryAllNymRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllNymRequest) ProtoMessage()    {}
func (*QueryAllNymRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf3698559b3858f5, []int{2}
}
func (m *QueryAllNymRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllNymRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllNymRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllNymRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllNymRequest.Merge(m, src)
}
func (m *QueryAllNymRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllNymRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllNymRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllNymRequest proto.InternalMessageInfo

func (m *QueryAllNymRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllNymResponse struct {
	Nym        []*Nym              `protobuf:"bytes,1,rep,name=Nym,proto3" json:"Nym,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllNymResponse) Reset()         { *m = QueryAllNymResponse{} }
func (m *QueryAllNymResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllNymResponse) ProtoMessage()    {}
func (*QueryAllNymResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf3698559b3858f5, []int{3}
}
func (m *QueryAllNymResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllNymResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllNymResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllNymResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllNymResponse.Merge(m, src)
}
func (m *QueryAllNymResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllNymResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllNymResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllNymResponse proto.InternalMessageInfo

func (m *QueryAllNymResponse) GetNym() []*Nym {
	if m != nil {
		return m.Nym
	}
	return nil
}

func (m *QueryAllNymResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetNymRequest)(nil), "cheqdid.cheqdnode.cheqd.QueryGetNymRequest")
	proto.RegisterType((*QueryGetNymResponse)(nil), "cheqdid.cheqdnode.cheqd.QueryGetNymResponse")
	proto.RegisterType((*QueryAllNymRequest)(nil), "cheqdid.cheqdnode.cheqd.QueryAllNymRequest")
	proto.RegisterType((*QueryAllNymResponse)(nil), "cheqdid.cheqdnode.cheqd.QueryAllNymResponse")
}

func init() { proto.RegisterFile("cheqd/query.proto", fileDescriptor_bf3698559b3858f5) }

var fileDescriptor_bf3698559b3858f5 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xe3, 0x40,
	0x18, 0xc6, 0x3b, 0xe9, 0x6e, 0x0f, 0xb3, 0xb0, 0xcb, 0xce, 0xc2, 0x6e, 0xe9, 0x76, 0xc3, 0x92,
	0x5d, 0xfc, 0xef, 0x0c, 0xad, 0x9f, 0xa0, 0x8a, 0xf6, 0x56, 0xb4, 0x47, 0xf1, 0x92, 0x34, 0x43,
	0x3a, 0x90, 0xcc, 0xa4, 0x9d, 0xa9, 0x18, 0xc4, 0x8b, 0xe2, 0x55, 0x04, 0xf1, 0x3b, 0x79, 0x2c,
	0x78, 0xf1, 0x28, 0xad, 0x1f, 0x44, 0x32, 0x13, 0x31, 0x51, 0x4b, 0x7b, 0x49, 0xc2, 0xcb, 0xf3,
	0xbc, 0xcf, 0xef, 0x7d, 0xdf, 0xc0, 0xef, 0xbd, 0x3e, 0x1d, 0xf8, 0x64, 0x30, 0xa2, 0xc3, 0x04,
	0xc7, 0x43, 0xa1, 0x04, 0xfa, 0xa5, 0x4b, 0xcc, 0xc7, 0xfa, 0xcd, 0x85, 0x4f, 0xcd, 0x57, 0xad,
	0x1e, 0x08, 0x11, 0x84, 0x94, 0xb8, 0x31, 0x23, 0x2e, 0xe7, 0x42, 0xb9, 0x8a, 0x09, 0x2e, 0x8d,
	0xad, 0xb6, 0xd6, 0x13, 0x32, 0x12, 0x92, 0x78, 0xae, 0xa4, 0xa6, 0x1f, 0x39, 0x6e, 0x78, 0x54,
	0xb9, 0x0d, 0x12, 0xbb, 0x01, 0xe3, 0x5a, 0x9c, 0x69, 0xbf, 0x99, 0x54, 0x9e, 0x44, 0xa6, 0xe0,
	0xfc, 0x87, 0xe8, 0x20, 0xb5, 0xb4, 0xa9, 0xea, 0x24, 0x51, 0x97, 0x0e, 0x46, 0x54, 0x2a, 0xf4,
	0x15, 0x5a, 0xcc, 0xaf, 0x82, 0xbf, 0x60, 0xe5, 0x53, 0xd7, 0x62, 0xbe, 0xb3, 0x0b, 0x7f, 0x14,
	0x54, 0x32, 0x16, 0x5c, 0x52, 0x84, 0x61, 0xb9, 0x93, 0x44, 0x5a, 0xf7, 0xa5, 0x59, 0xc7, 0x33,
	0xf0, 0x71, 0x6a, 0x49, 0x85, 0xce, 0x51, 0x16, 0xd6, 0x0a, 0xc3, 0x5c, 0xd8, 0x1e, 0x84, 0xaf,
	0x9c, 0x59, 0xb3, 0x25, 0x6c, 0x86, 0xc2, 0xe9, 0x50, 0xd8, 0x2c, 0x29, 0x1b, 0x0a, 0xef, 0xbb,
	0x01, 0xcd, 0xbc, 0xdd, 0x9c, 0xd3, 0xb9, 0x02, 0x19, 0xe5, 0x4b, 0xfb, 0xb7, 0x94, 0xe5, 0x85,
	0x28, 0x51, 0xbb, 0xc0, 0x63, 0x69, 0x9e, 0xe5, 0xb9, 0x3c, 0x26, 0x2c, 0x0f, 0xd4, 0xbc, 0xb5,
	0xe0, 0x67, 0x0d, 0x84, 0x2e, 0x81, 0x66, 0x40, 0xeb, 0x33, 0xd3, 0xdf, 0x1f, 0xa1, 0xb6, 0xb1,
	0x98, 0xd8, 0x04, 0x3b, 0xff, 0xce, 0xef, 0x9f, 0x6e, 0xac, 0x3f, 0xe8, 0x37, 0x31, 0x27, 0xce,
	0x3f, 0x79, 0x12, 0x91, 0x53, 0xe6, 0x9f, 0xa1, 0x0b, 0x00, 0x2b, 0x9d, 0x24, 0x6a, 0x85, 0xe1,
	0x3c, 0x94, 0xc2, 0x89, 0xe6, 0xa1, 0x14, 0x17, 0xee, 0xd8, 0x1a, 0xa5, 0x8a, 0x7e, 0x7e, 0x8c,
	0xb2, 0xbd, 0x73, 0x37, 0xb1, 0xc1, 0x78, 0x62, 0x83, 0xc7, 0x89, 0x0d, 0xae, 0xa7, 0x76, 0x69,
	0x3c, 0xb5, 0x4b, 0x0f, 0x53, 0xbb, 0x74, 0xb8, 0x1a, 0x30, 0xd5, 0x1f, 0x79, 0xb8, 0x27, 0xa2,
	0xbc, 0x6b, 0x33, 0x0d, 0x24, 0x27, 0x59, 0x49, 0x25, 0x31, 0x95, 0x5e, 0x45, 0xff, 0xbf, 0x5b,
	0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xd8, 0x3e, 0x60, 0x48, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// this line is used by starport scaffolding # 2
	Nym(ctx context.Context, in *QueryGetNymRequest, opts ...grpc.CallOption) (*QueryGetNymResponse, error)
	NymAll(ctx context.Context, in *QueryAllNymRequest, opts ...grpc.CallOption) (*QueryAllNymResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Nym(ctx context.Context, in *QueryGetNymRequest, opts ...grpc.CallOption) (*QueryGetNymResponse, error) {
	out := new(QueryGetNymResponse)
	err := c.cc.Invoke(ctx, "/cheqdid.cheqdnode.cheqd.Query/Nym", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) NymAll(ctx context.Context, in *QueryAllNymRequest, opts ...grpc.CallOption) (*QueryAllNymResponse, error) {
	out := new(QueryAllNymResponse)
	err := c.cc.Invoke(ctx, "/cheqdid.cheqdnode.cheqd.Query/NymAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// this line is used by starport scaffolding # 2
	Nym(context.Context, *QueryGetNymRequest) (*QueryGetNymResponse, error)
	NymAll(context.Context, *QueryAllNymRequest) (*QueryAllNymResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Nym(ctx context.Context, req *QueryGetNymRequest) (*QueryGetNymResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nym not implemented")
}
func (*UnimplementedQueryServer) NymAll(ctx context.Context, req *QueryAllNymRequest) (*QueryAllNymResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NymAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Nym_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetNymRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Nym(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheqdid.cheqdnode.cheqd.Query/Nym",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Nym(ctx, req.(*QueryGetNymRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_NymAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllNymRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NymAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheqdid.cheqdnode.cheqd.Query/NymAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NymAll(ctx, req.(*QueryAllNymRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cheqdid.cheqdnode.cheqd.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Nym",
			Handler:    _Query_Nym_Handler,
		},
		{
			MethodName: "NymAll",
			Handler:    _Query_NymAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheqd/query.proto",
}

func (m *QueryGetNymRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetNymRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetNymRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetNymResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetNymResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetNymResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nym != nil {
		{
			size, err := m.Nym.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllNymRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllNymRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllNymRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllNymResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllNymResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllNymResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Nym) > 0 {
		for iNdEx := len(m.Nym) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nym[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetNymRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetNymResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nym != nil {
		l = m.Nym.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllNymRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllNymResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Nym) > 0 {
		for _, e := range m.Nym {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetNymRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetNymRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetNymRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetNymResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetNymResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetNymResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nym", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nym == nil {
				m.Nym = &Nym{}
			}
			if err := m.Nym.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllNymRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllNymRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllNymRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllNymResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllNymResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllNymResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nym", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nym = append(m.Nym, &Nym{})
			if err := m.Nym[len(m.Nym)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
