// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/collector/metrics/v1/metrics_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/tigrannajaryan/exp-otelproto/encodings/otlp_gogo3/metrics/v1"
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

type ExportMetricsServiceRequest struct {
	// An array of ResourceMetrics.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceMetrics []*v1.ResourceMetrics `protobuf:"bytes,1,rep,name=resource_metrics,json=resourceMetrics,proto3" json:"resource_metrics,omitempty"`
}

func (m *ExportMetricsServiceRequest) Reset()         { *m = ExportMetricsServiceRequest{} }
func (m *ExportMetricsServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportMetricsServiceRequest) ProtoMessage()    {}
func (*ExportMetricsServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fb6015e6e64798, []int{0}
}
func (m *ExportMetricsServiceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportMetricsServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportMetricsServiceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportMetricsServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportMetricsServiceRequest.Merge(m, src)
}
func (m *ExportMetricsServiceRequest) XXX_Size() int {
	return m.Size()
}
func (m *ExportMetricsServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportMetricsServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportMetricsServiceRequest proto.InternalMessageInfo

func (m *ExportMetricsServiceRequest) GetResourceMetrics() []*v1.ResourceMetrics {
	if m != nil {
		return m.ResourceMetrics
	}
	return nil
}

type ExportMetricsServiceResponse struct {
}

func (m *ExportMetricsServiceResponse) Reset()         { *m = ExportMetricsServiceResponse{} }
func (m *ExportMetricsServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportMetricsServiceResponse) ProtoMessage()    {}
func (*ExportMetricsServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fb6015e6e64798, []int{1}
}
func (m *ExportMetricsServiceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportMetricsServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportMetricsServiceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportMetricsServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportMetricsServiceResponse.Merge(m, src)
}
func (m *ExportMetricsServiceResponse) XXX_Size() int {
	return m.Size()
}
func (m *ExportMetricsServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportMetricsServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportMetricsServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExportMetricsServiceRequest)(nil), "opentelemetrygogo3.proto.collector.metrics.v1.ExportMetricsServiceRequest")
	proto.RegisterType((*ExportMetricsServiceResponse)(nil), "opentelemetrygogo3.proto.collector.metrics.v1.ExportMetricsServiceResponse")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/collector/metrics/v1/metrics_service.proto", fileDescriptor_75fb6015e6e64798)
}

var fileDescriptor_75fb6015e6e64798 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x4b, 0x03, 0x31,
	0x14, 0xc7, 0x2f, 0x08, 0x1d, 0x22, 0xa8, 0x9c, 0x8b, 0x54, 0x09, 0xd2, 0xc9, 0xc1, 0x26, 0xb6,
	0x75, 0x76, 0x28, 0xb8, 0x28, 0x42, 0xad, 0x9b, 0xcb, 0x71, 0x8d, 0x8f, 0x33, 0x72, 0xcd, 0x8b,
	0x49, 0x5a, 0xda, 0xc5, 0xcf, 0xe0, 0x77, 0xf0, 0x2b, 0x88, 0x9f, 0xc1, 0xb1, 0xa3, 0xa3, 0xb4,
	0x5f, 0x44, 0xda, 0x9c, 0xca, 0x61, 0x15, 0x0a, 0x6e, 0xc7, 0x7b, 0xef, 0xf7, 0xff, 0xbd, 0x4b,
	0x42, 0x4f, 0xd0, 0x80, 0xf6, 0x90, 0x43, 0x1f, 0xbc, 0x1d, 0x0b, 0x63, 0xd1, 0xa3, 0x90, 0x98,
	0xe7, 0x20, 0x3d, 0x5a, 0x31, 0xaf, 0x2a, 0xe9, 0xc4, 0xb0, 0xf1, 0xf9, 0x99, 0x38, 0xb0, 0x43,
	0x25, 0x81, 0x2f, 0x46, 0xe3, 0x7a, 0x89, 0xcf, 0x30, 0xc3, 0x56, 0xe8, 0xf0, 0xaf, 0x10, 0x5e,
	0x90, 0x7c, 0xd8, 0xa8, 0x1e, 0x2e, 0xd3, 0xfd, 0x94, 0x84, 0x88, 0xda, 0x03, 0xdd, 0x3d, 0x1d,
	0x19, 0xb4, 0xfe, 0x22, 0x94, 0xaf, 0x82, 0xba, 0x0b, 0xf7, 0x03, 0x70, 0x3e, 0x4e, 0xe8, 0x96,
	0x05, 0x87, 0x03, 0x2b, 0x21, 0x29, 0xc0, 0x1d, 0xb2, 0xbf, 0x76, 0xb0, 0xde, 0x3c, 0xe6, 0xbf,
	0xae, 0xf5, 0xbd, 0x0c, 0xef, 0x16, 0x70, 0x91, 0xde, 0xdd, 0xb4, 0xe5, 0x42, 0x8d, 0xd1, 0xbd,
	0xe5, 0x7e, 0x67, 0x50, 0x3b, 0x68, 0xbe, 0x10, 0xba, 0x51, 0x6e, 0xc5, 0x4f, 0x84, 0x56, 0x02,
	0x13, 0x9f, 0xf1, 0x95, 0xce, 0x86, 0xff, 0xf1, 0xab, 0xd5, 0xf3, 0x7f, 0xc9, 0x0a, 0x6b, 0xd7,
	0xa2, 0xf6, 0x33, 0x79, 0x9d, 0x32, 0x32, 0x99, 0x32, 0xf2, 0x3e, 0x65, 0xe4, 0x71, 0xc6, 0xa2,
	0xc9, 0x8c, 0x45, 0x6f, 0x33, 0x16, 0xd1, 0x23, 0x85, 0xab, 0xb9, 0xda, 0xdb, 0x65, 0x4d, 0x67,
	0x3e, 0xd9, 0x21, 0xd7, 0x97, 0x99, 0xf2, 0xb7, 0x83, 0x1e, 0x97, 0xd8, 0x17, 0x5e, 0x65, 0x36,
	0xd5, 0x3a, 0xbd, 0x4b, 0xed, 0x38, 0xd5, 0x02, 0x46, 0xa6, 0x8e, 0x1e, 0xf2, 0x70, 0xfd, 0xa0,
	0x25, 0xde, 0x28, 0x9d, 0x39, 0x81, 0x3e, 0x37, 0xc9, 0xc2, 0xb7, 0xf4, 0x09, 0xf6, 0x2a, 0x0b,
	0xa0, 0xf5, 0x11, 0x00, 0x00, 0xff, 0xff, 0x58, 0xc0, 0xf0, 0x7f, 0xb5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, in *ExportMetricsServiceRequest, opts ...grpc.CallOption) (*ExportMetricsServiceResponse, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) Export(ctx context.Context, in *ExportMetricsServiceRequest, opts ...grpc.CallOption) (*ExportMetricsServiceResponse, error) {
	out := new(ExportMetricsServiceResponse)
	err := c.cc.Invoke(ctx, "/opentelemetrygogo3.proto.collector.metrics.v1.MetricsService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *ExportMetricsServiceRequest) (*ExportMetricsServiceResponse, error)
}

// UnimplementedMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (*UnimplementedMetricsServiceServer) Export(ctx context.Context, req *ExportMetricsServiceRequest) (*ExportMetricsServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportMetricsServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetrygogo3.proto.collector.metrics.v1.MetricsService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).Export(ctx, req.(*ExportMetricsServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetrygogo3.proto.collector.metrics.v1.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _MetricsService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/collector/metrics/v1/metrics_service.proto",
}

func (m *ExportMetricsServiceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportMetricsServiceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportMetricsServiceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResourceMetrics) > 0 {
		for iNdEx := len(m.ResourceMetrics) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResourceMetrics[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetricsService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ExportMetricsServiceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportMetricsServiceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportMetricsServiceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMetricsService(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetricsService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExportMetricsServiceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ResourceMetrics) > 0 {
		for _, e := range m.ResourceMetrics {
			l = e.Size()
			n += 1 + l + sovMetricsService(uint64(l))
		}
	}
	return n
}

func (m *ExportMetricsServiceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMetricsService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetricsService(x uint64) (n int) {
	return sovMetricsService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExportMetricsServiceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetricsService
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
			return fmt.Errorf("proto: ExportMetricsServiceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportMetricsServiceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceMetrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetricsService
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
				return ErrInvalidLengthMetricsService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetricsService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceMetrics = append(m.ResourceMetrics, &v1.ResourceMetrics{})
			if err := m.ResourceMetrics[len(m.ResourceMetrics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetricsService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetricsService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetricsService
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
func (m *ExportMetricsServiceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetricsService
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
			return fmt.Errorf("proto: ExportMetricsServiceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportMetricsServiceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMetricsService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetricsService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetricsService
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
func skipMetricsService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetricsService
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
					return 0, ErrIntOverflowMetricsService
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
					return 0, ErrIntOverflowMetricsService
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
				return 0, ErrInvalidLengthMetricsService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetricsService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetricsService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetricsService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetricsService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetricsService = fmt.Errorf("proto: unexpected end of group")
)