// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exchange.proto

package traceprotobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Capabilities int32

const (
	Capabilities_SKIP             Capabilities = 0
	Capabilities_ZLIB_COMPRESSION Capabilities = 1
	Capabilities_LZ4_COMPRESSION  Capabilities = 2
)

var Capabilities_name = map[int32]string{
	0: "SKIP",
	1: "ZLIB_COMPRESSION",
	2: "LZ4_COMPRESSION",
}

var Capabilities_value = map[string]int32{
	"SKIP":             0,
	"ZLIB_COMPRESSION": 1,
	"LZ4_COMPRESSION":  2,
}

func (x Capabilities) String() string {
	return proto.EnumName(Capabilities_name, int32(x))
}

func (Capabilities) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{0}
}

type CompressionMethod int32

const (
	CompressionMethod_NONE CompressionMethod = 0
	CompressionMethod_LZ4  CompressionMethod = 1
	CompressionMethod_ZLIB CompressionMethod = 2
)

var CompressionMethod_name = map[int32]string{
	0: "NONE",
	1: "LZ4",
	2: "ZLIB",
}

var CompressionMethod_value = map[string]int32{
	"NONE": 0,
	"LZ4":  1,
	"ZLIB": 2,
}

func (x CompressionMethod) String() string {
	return proto.EnumName(CompressionMethod_name, int32(x))
}

func (CompressionMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{1}
}

type ExportResponse_ResultCode int32

const (
	// Telemetry data is successfully processed by the server.
	ExportResponse_Success ExportResponse_ResultCode = 0
	// processing of telemetry data failed. The client MUST NOT retry
	// sending the same telemetry data. The telemetry data MUST be dropped.
	// This for example can happen when the request contains bad data and
	// cannot be deserialized or otherwise processed by the server.
	ExportResponse_FailedNoneRetryable ExportResponse_ResultCode = 1
	// Processing of telemetry data failed. The client SHOULD record the
	// error and MAY retry exporting the same data after some time. This
	// for example can happen when the server is overloaded.
	ExportResponse_FailedRetryable ExportResponse_ResultCode = 2
)

var ExportResponse_ResultCode_name = map[int32]string{
	0: "Success",
	1: "FailedNoneRetryable",
	2: "FailedRetryable",
}

var ExportResponse_ResultCode_value = map[string]int32{
	"Success":             0,
	"FailedNoneRetryable": 1,
	"FailedRetryable":     2,
}

func (x ExportResponse_ResultCode) String() string {
	return proto.EnumName(ExportResponse_ResultCode_name, int32(x))
}

func (ExportResponse_ResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{3, 0}
}

type HelloRequest struct {
	ClientVer            int32    `protobuf:"varint,1,opt,name=client_ver,json=clientVer,proto3" json:"client_ver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetClientVer() int32 {
	if m != nil {
		return m.ClientVer
	}
	return 0
}

type HelloResponse struct {
	ServerVer            int32    `protobuf:"varint,1,opt,name=server_ver,json=serverVer,proto3" json:"server_ver,omitempty"`
	Capabilities         uint32   `protobuf:"varint,2,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetServerVer() int32 {
	if m != nil {
		return m.ServerVer
	}
	return 0
}

func (m *HelloResponse) GetCapabilities() uint32 {
	if m != nil {
		return m.Capabilities
	}
	return 0
}

// A request from client to server containing telemetry data to export.
type ExportRequest struct {
	// Unique sequential ID generated by the client.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Telemetry data.
	NodeSpans            []*NodeSpans   `protobuf:"bytes,2,rep,name=nodeSpans,proto3" json:"nodeSpans,omitempty"`
	NodeMetrics          []*NodeMetrics `protobuf:"bytes,3,rep,name=nodeMetrics,proto3" json:"nodeMetrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExportRequest) Reset()         { *m = ExportRequest{} }
func (m *ExportRequest) String() string { return proto.CompactTextString(m) }
func (*ExportRequest) ProtoMessage()    {}
func (*ExportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{2}
}

func (m *ExportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportRequest.Unmarshal(m, b)
}
func (m *ExportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportRequest.Marshal(b, m, deterministic)
}
func (m *ExportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportRequest.Merge(m, src)
}
func (m *ExportRequest) XXX_Size() int {
	return xxx_messageInfo_ExportRequest.Size(m)
}
func (m *ExportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportRequest proto.InternalMessageInfo

func (m *ExportRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExportRequest) GetNodeSpans() []*NodeSpans {
	if m != nil {
		return m.NodeSpans
	}
	return nil
}

func (m *ExportRequest) GetNodeMetrics() []*NodeMetrics {
	if m != nil {
		return m.NodeMetrics
	}
	return nil
}

// A response to ExportRequest.
type ExportResponse struct {
	// ID of a response that the server acknowledges.
	Id         uint64                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ResultCode ExportResponse_ResultCode `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3,enum=traceprotobuf.ExportResponse_ResultCode" json:"result_code,omitempty"`
	// How long the client must wait before sending the next ExportRequest. 0 indicates
	// that the client doesn't need to wait.
	ThrottlePeriodMillisec uint32   `protobuf:"varint,3,opt,name=throttle_period_millisec,json=throttlePeriodMillisec,proto3" json:"throttle_period_millisec,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ExportResponse) Reset()         { *m = ExportResponse{} }
func (m *ExportResponse) String() string { return proto.CompactTextString(m) }
func (*ExportResponse) ProtoMessage()    {}
func (*ExportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{3}
}

func (m *ExportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportResponse.Unmarshal(m, b)
}
func (m *ExportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportResponse.Marshal(b, m, deterministic)
}
func (m *ExportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportResponse.Merge(m, src)
}
func (m *ExportResponse) XXX_Size() int {
	return xxx_messageInfo_ExportResponse.Size(m)
}
func (m *ExportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportResponse proto.InternalMessageInfo

func (m *ExportResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExportResponse) GetResultCode() ExportResponse_ResultCode {
	if m != nil {
		return m.ResultCode
	}
	return ExportResponse_Success
}

func (m *ExportResponse) GetThrottlePeriodMillisec() uint32 {
	if m != nil {
		return m.ThrottlePeriodMillisec
	}
	return 0
}

// RequestHeader is used by transports that unlike gRPC don't have built-in request
// compression such as WebSocket. Request body typically follows the header.
type RequestHeader struct {
	// Compression method used for body.
	Compression CompressionMethod `protobuf:"varint,1,opt,name=compression,proto3,enum=traceprotobuf.CompressionMethod" json:"compression,omitempty"`
	// Compression level as defined by the compression method.
	CompressionLevel     int32    `protobuf:"varint,2,opt,name=compression_level,json=compressionLevel,proto3" json:"compression_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{4}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetCompression() CompressionMethod {
	if m != nil {
		return m.Compression
	}
	return CompressionMethod_NONE
}

func (m *RequestHeader) GetCompressionLevel() int32 {
	if m != nil {
		return m.CompressionLevel
	}
	return 0
}

// RequestBody is used by transports that unlike gRPC don't have built-in message type
// multiplexing such as WebSocket.
type RequestBody struct {
	// Types that are valid to be assigned to Body:
	//	*RequestBody_Hello
	//	*RequestBody_Export
	Body                 isRequestBody_Body `protobuf_oneof:"body"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RequestBody) Reset()         { *m = RequestBody{} }
func (m *RequestBody) String() string { return proto.CompactTextString(m) }
func (*RequestBody) ProtoMessage()    {}
func (*RequestBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{5}
}

func (m *RequestBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestBody.Unmarshal(m, b)
}
func (m *RequestBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestBody.Marshal(b, m, deterministic)
}
func (m *RequestBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestBody.Merge(m, src)
}
func (m *RequestBody) XXX_Size() int {
	return xxx_messageInfo_RequestBody.Size(m)
}
func (m *RequestBody) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestBody.DiscardUnknown(m)
}

var xxx_messageInfo_RequestBody proto.InternalMessageInfo

type isRequestBody_Body interface {
	isRequestBody_Body()
}

type RequestBody_Hello struct {
	Hello *HelloRequest `protobuf:"bytes,1,opt,name=hello,proto3,oneof"`
}

type RequestBody_Export struct {
	Export *ExportRequest `protobuf:"bytes,2,opt,name=export,proto3,oneof"`
}

func (*RequestBody_Hello) isRequestBody_Body() {}

func (*RequestBody_Export) isRequestBody_Body() {}

func (m *RequestBody) GetBody() isRequestBody_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *RequestBody) GetHello() *HelloRequest {
	if x, ok := m.GetBody().(*RequestBody_Hello); ok {
		return x.Hello
	}
	return nil
}

func (m *RequestBody) GetExport() *ExportRequest {
	if x, ok := m.GetBody().(*RequestBody_Export); ok {
		return x.Export
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RequestBody) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RequestBody_Hello)(nil),
		(*RequestBody_Export)(nil),
	}
}

// Response is used by transports that unlike gRPC don't have built-in message type
// multiplexing such as WebSocket.
type Response struct {
	// Types that are valid to be assigned to Body:
	//	*Response_Hello
	//	*Response_Export
	Body                 isResponse_Body `protobuf_oneof:"body"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{6}
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

type isResponse_Body interface {
	isResponse_Body()
}

type Response_Hello struct {
	Hello *HelloResponse `protobuf:"bytes,1,opt,name=hello,proto3,oneof"`
}

type Response_Export struct {
	Export *ExportResponse `protobuf:"bytes,2,opt,name=export,proto3,oneof"`
}

func (*Response_Hello) isResponse_Body() {}

func (*Response_Export) isResponse_Body() {}

func (m *Response) GetBody() isResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Response) GetHello() *HelloResponse {
	if x, ok := m.GetBody().(*Response_Hello); ok {
		return x.Hello
	}
	return nil
}

func (m *Response) GetExport() *ExportResponse {
	if x, ok := m.GetBody().(*Response_Export); ok {
		return x.Export
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Response_Hello)(nil),
		(*Response_Export)(nil),
	}
}

func init() {
	proto.RegisterEnum("traceprotobuf.Capabilities", Capabilities_name, Capabilities_value)
	proto.RegisterEnum("traceprotobuf.CompressionMethod", CompressionMethod_name, CompressionMethod_value)
	proto.RegisterEnum("traceprotobuf.ExportResponse_ResultCode", ExportResponse_ResultCode_name, ExportResponse_ResultCode_value)
	proto.RegisterType((*HelloRequest)(nil), "traceprotobuf.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "traceprotobuf.HelloResponse")
	proto.RegisterType((*ExportRequest)(nil), "traceprotobuf.ExportRequest")
	proto.RegisterType((*ExportResponse)(nil), "traceprotobuf.ExportResponse")
	proto.RegisterType((*RequestHeader)(nil), "traceprotobuf.RequestHeader")
	proto.RegisterType((*RequestBody)(nil), "traceprotobuf.RequestBody")
	proto.RegisterType((*Response)(nil), "traceprotobuf.Response")
}

func init() { proto.RegisterFile("exchange.proto", fileDescriptor_e0328a4f16f87ea1) }

var fileDescriptor_e0328a4f16f87ea1 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x8d, 0x9d, 0xfe, 0x1d, 0x37, 0xf9, 0xb9, 0xdb, 0xea, 0x47, 0x28, 0xad, 0x54, 0xf9, 0x14,
	0x15, 0x61, 0xd1, 0xb4, 0x2a, 0x1c, 0x38, 0x25, 0x2a, 0x24, 0x22, 0xff, 0xb4, 0x41, 0x3d, 0xf4,
	0x62, 0x39, 0xf6, 0x40, 0x2c, 0x6d, 0xbd, 0x66, 0x77, 0x13, 0x35, 0x37, 0xe0, 0x3b, 0xf0, 0x25,
	0x39, 0xf0, 0x19, 0x90, 0xd7, 0x4e, 0xe3, 0xa4, 0x85, 0xdb, 0xee, 0x7b, 0x6f, 0x66, 0xde, 0xec,
	0x8c, 0x0d, 0x55, 0xbc, 0x0f, 0x26, 0x7e, 0xfc, 0x05, 0xdd, 0x44, 0x70, 0xc5, 0x49, 0x45, 0x09,
	0x3f, 0x40, 0x7d, 0x1e, 0x4f, 0x3f, 0x1f, 0x1d, 0x2a, 0x64, 0x78, 0x87, 0x4a, 0xcc, 0xbd, 0xd0,
	0x57, 0x7e, 0x26, 0x72, 0x5e, 0xc1, 0x5e, 0x1b, 0x19, 0xe3, 0x14, 0xbf, 0x4e, 0x51, 0x2a, 0x72,
	0x02, 0x10, 0xb0, 0x08, 0x63, 0xe5, 0xcd, 0x50, 0xd4, 0x8c, 0x53, 0xa3, 0xbe, 0x49, 0x77, 0x33,
	0xe4, 0x06, 0x85, 0x43, 0xa1, 0x92, 0xcb, 0x65, 0xc2, 0x63, 0x89, 0xa9, 0x5e, 0xa2, 0x98, 0xa1,
	0x28, 0xea, 0x33, 0xe4, 0x06, 0x05, 0x71, 0x60, 0x2f, 0xf0, 0x13, 0x7f, 0x1c, 0xb1, 0x48, 0x45,
	0x28, 0x6b, 0xe6, 0xa9, 0x51, 0xaf, 0xd0, 0x15, 0xcc, 0xf9, 0x69, 0x40, 0xe5, 0xfa, 0x3e, 0xe1,
	0x42, 0x2d, 0x4c, 0x54, 0xc1, 0x8c, 0x42, 0x9d, 0x6c, 0x83, 0x9a, 0x51, 0x48, 0xae, 0x60, 0x37,
	0xe6, 0x21, 0x8e, 0x12, 0x3f, 0x4e, 0x53, 0x94, 0xeb, 0x56, 0xa3, 0xe6, 0xae, 0x74, 0xe7, 0xf6,
	0x17, 0x3c, 0x5d, 0x4a, 0xc9, 0x3b, 0xb0, 0xd2, 0x4b, 0x0f, 0x95, 0x88, 0x02, 0x59, 0x2b, 0xeb,
	0xc8, 0xa3, 0x27, 0x22, 0x73, 0x05, 0x2d, 0xca, 0x9d, 0xdf, 0x06, 0x54, 0x17, 0xbe, 0xf2, 0x6e,
	0xd7, 0x8d, 0x75, 0xc0, 0x12, 0x28, 0xa7, 0x4c, 0x79, 0x01, 0x0f, 0x51, 0x77, 0x57, 0x6d, 0xd4,
	0xd7, 0x0a, 0xac, 0xe6, 0x70, 0xa9, 0x0e, 0x68, 0xf1, 0x10, 0x29, 0x88, 0x87, 0x33, 0x79, 0x0b,
	0x35, 0x35, 0x11, 0x5c, 0x29, 0x86, 0x5e, 0x82, 0x22, 0xe2, 0xa1, 0x77, 0x17, 0x31, 0x16, 0x49,
	0x0c, 0x6a, 0x65, 0xfd, 0x6a, 0xff, 0x2f, 0xf8, 0xa1, 0xa6, 0x7b, 0x39, 0xeb, 0x7c, 0x00, 0x58,
	0xe6, 0x24, 0x16, 0x6c, 0x8f, 0xa6, 0x41, 0x80, 0x52, 0xda, 0x25, 0xf2, 0x0c, 0x0e, 0xde, 0xfb,
	0x11, 0xc3, 0xb0, 0xcf, 0x63, 0xa4, 0xe9, 0xec, 0xfd, 0x31, 0x43, 0xdb, 0x20, 0x07, 0xf0, 0x5f,
	0x46, 0x2c, 0x41, 0xd3, 0xf9, 0x66, 0x40, 0x25, 0x1f, 0x41, 0x1b, 0xfd, 0x10, 0x05, 0x69, 0x82,
	0x15, 0xf0, 0xbb, 0x44, 0xa0, 0x94, 0x11, 0x8f, 0x75, 0xe3, 0xd5, 0xc6, 0xe9, 0x5a, 0x7f, 0xad,
	0xa5, 0xa2, 0x87, 0x6a, 0xc2, 0x43, 0x5a, 0x0c, 0x22, 0x2f, 0x61, 0xbf, 0x70, 0xf5, 0x18, 0xce,
	0x90, 0xe9, 0x97, 0xda, 0xa4, 0x76, 0x81, 0xe8, 0xa6, 0xb8, 0xf3, 0xc3, 0x00, 0x2b, 0xb7, 0xd0,
	0xe4, 0xe1, 0x9c, 0x5c, 0xc0, 0xe6, 0x24, 0xdd, 0x37, 0x5d, 0xda, 0x6a, 0xbc, 0x58, 0x2b, 0x5d,
	0x5c, 0xdd, 0x76, 0x89, 0x66, 0x5a, 0x72, 0x05, 0x5b, 0xa8, 0xdf, 0x5c, 0x97, 0xb1, 0x1a, 0xc7,
	0x7f, 0x19, 0xc8, 0x22, 0x2c, 0x57, 0x37, 0xb7, 0x60, 0x63, 0xcc, 0xc3, 0xb9, 0xf3, 0xdd, 0x80,
	0x9d, 0x87, 0x91, 0x5f, 0xae, 0x3a, 0x38, 0x7e, 0xda, 0x41, 0x26, 0x5e, 0x5a, 0x78, 0xb3, 0x66,
	0xe1, 0xe4, 0x9f, 0x3b, 0xf1, 0xd8, 0xc3, 0x59, 0x0b, 0xf6, 0x5a, 0x85, 0x8f, 0x84, 0xec, 0xc0,
	0xc6, 0xe8, 0x63, 0x67, 0x68, 0x97, 0xc8, 0x21, 0xd8, 0xb7, 0xdd, 0x4e, 0xd3, 0x6b, 0x0d, 0x7a,
	0x43, 0x7a, 0x3d, 0x1a, 0x75, 0x06, 0xfd, 0x6c, 0xa0, 0xdd, 0xdb, 0xcb, 0x15, 0xd0, 0x3c, 0x7b,
	0x0d, 0xfb, 0x8f, 0x86, 0x93, 0x66, 0xea, 0x0f, 0xfa, 0xd7, 0x76, 0x89, 0x6c, 0x43, 0xb9, 0x7b,
	0x7b, 0x69, 0x1b, 0x29, 0x94, 0xa6, 0xb4, 0xcd, 0x66, 0x1b, 0x8e, 0x23, 0xee, 0xf2, 0x04, 0xe3,
	0x00, 0x63, 0x39, 0x95, 0xd9, 0x5f, 0x22, 0x33, 0xef, 0xce, 0xce, 0x9b, 0xf0, 0x29, 0x3d, 0x0d,
	0x53, 0x70, 0x68, 0xfc, 0x32, 0x9f, 0x0f, 0x12, 0x8c, 0x5b, 0x99, 0x52, 0x83, 0xae, 0xe6, 0xdd,
	0x9b, 0xf3, 0xf1, 0x96, 0x8e, 0xbc, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0x65, 0x0d, 0x01, 0x22,
	0x96, 0x04, 0x00, 0x00,
}
