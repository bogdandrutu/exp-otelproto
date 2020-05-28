// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package baseline

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

// ValueType is the enumeration of possible types that value can have.
// TODO: consolidate this with AttributeKeyValue.ValueType.
type ValueType int32

const (
	ValueType_STRING ValueType = 0
	ValueType_INT    ValueType = 1
	ValueType_DOUBLE ValueType = 2
	ValueType_BOOL   ValueType = 3
	ValueType_LIST   ValueType = 4
	ValueType_MAP    ValueType = 5
)

var ValueType_name = map[int32]string{
	0: "STRING",
	1: "INT",
	2: "DOUBLE",
	3: "BOOL",
	4: "LIST",
	5: "MAP",
}

var ValueType_value = map[string]int32{
	"STRING": 0,
	"INT":    1,
	"DOUBLE": 2,
	"BOOL":   3,
	"LIST":   4,
	"MAP":    5,
}

func (x ValueType) String() string {
	return proto.EnumName(ValueType_name, int32(x))
}

func (ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

// Resource information.
type Resource struct {
	// Set of labels that describe the resource.
	Attributes []*AttributeKeyValue `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0, then
	// no attributes were dropped.
	DroppedAttributesCount uint32   `protobuf:"varint,2,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetAttributes() []*AttributeKeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Resource) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

// AnyValue is a value that is used to store values that can be one of the supported
// value types.
type AnyValue struct {
	// type of the value.
	Type        ValueType `protobuf:"varint,1,opt,name=type,proto3,enum=baseline.ValueType" json:"type,omitempty"`
	BoolValue   bool      `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	StringValue string    `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	IntValue    int64     `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	DoubleValue float64   `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	// If the type is LIST or MAP then this field is set and the list or the map
	// is stored in the ValueListOrMap message. If type is neither LIST nor MAP then
	// field is not set.
	// LIST and MAP are less frequent types and thus are stored in a separate message
	// to avoid consuming memory. This saves about 10-15% of memory for the most frequent
	// case when the value is a primitive type.
	ListOrMap            *ValueListOrMap `protobuf:"bytes,6,opt,name=list_or_map,json=listOrMap,proto3" json:"list_or_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AnyValue) Reset()         { *m = AnyValue{} }
func (m *AnyValue) String() string { return proto.CompactTextString(m) }
func (*AnyValue) ProtoMessage()    {}
func (*AnyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *AnyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnyValue.Unmarshal(m, b)
}
func (m *AnyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnyValue.Marshal(b, m, deterministic)
}
func (m *AnyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnyValue.Merge(m, src)
}
func (m *AnyValue) XXX_Size() int {
	return xxx_messageInfo_AnyValue.Size(m)
}
func (m *AnyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AnyValue.DiscardUnknown(m)
}

var xxx_messageInfo_AnyValue proto.InternalMessageInfo

func (m *AnyValue) GetType() ValueType {
	if m != nil {
		return m.Type
	}
	return ValueType_STRING
}

func (m *AnyValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *AnyValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *AnyValue) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *AnyValue) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *AnyValue) GetListOrMap() *ValueListOrMap {
	if m != nil {
		return m.ListOrMap
	}
	return nil
}

// ValueListOrMap is used for storing a list or a map of values.
type ValueListOrMap struct {
	// A list of values.
	List []*AnyValue `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	// A key/value map of values.
	Map                  []*AttributeKeyValue `protobuf:"bytes,2,rep,name=map,proto3" json:"map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ValueListOrMap) Reset()         { *m = ValueListOrMap{} }
func (m *ValueListOrMap) String() string { return proto.CompactTextString(m) }
func (*ValueListOrMap) ProtoMessage()    {}
func (*ValueListOrMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *ValueListOrMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueListOrMap.Unmarshal(m, b)
}
func (m *ValueListOrMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueListOrMap.Marshal(b, m, deterministic)
}
func (m *ValueListOrMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueListOrMap.Merge(m, src)
}
func (m *ValueListOrMap) XXX_Size() int {
	return xxx_messageInfo_ValueListOrMap.Size(m)
}
func (m *ValueListOrMap) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueListOrMap.DiscardUnknown(m)
}

var xxx_messageInfo_ValueListOrMap proto.InternalMessageInfo

func (m *ValueListOrMap) GetList() []*AnyValue {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ValueListOrMap) GetMap() []*AttributeKeyValue {
	if m != nil {
		return m.Map
	}
	return nil
}

// AttributeKeyValue is a key-value pair that is used to store Span attributes, Link
// attributes, etc.
type AttributeKeyValue struct {
	// key part of the key-value pair.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// type of the value.
	Type        ValueType `protobuf:"varint,2,opt,name=type,proto3,enum=baseline.ValueType" json:"type,omitempty"`
	StringValue string    `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	IntValue    int64     `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	DoubleValue float64   `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	BoolValue   bool      `protobuf:"varint,6,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	// If the type is LIST or MAP then this field is set and the list or the map
	// is stored in the ValueListOrMap message. If type is neither LIST nor MAP then
	// field is not set.
	ListOrMap            *ValueListOrMap `protobuf:"bytes,7,opt,name=listOrMap,proto3" json:"listOrMap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AttributeKeyValue) Reset()         { *m = AttributeKeyValue{} }
func (m *AttributeKeyValue) String() string { return proto.CompactTextString(m) }
func (*AttributeKeyValue) ProtoMessage()    {}
func (*AttributeKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *AttributeKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeKeyValue.Unmarshal(m, b)
}
func (m *AttributeKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeKeyValue.Marshal(b, m, deterministic)
}
func (m *AttributeKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeKeyValue.Merge(m, src)
}
func (m *AttributeKeyValue) XXX_Size() int {
	return xxx_messageInfo_AttributeKeyValue.Size(m)
}
func (m *AttributeKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeKeyValue proto.InternalMessageInfo

func (m *AttributeKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AttributeKeyValue) GetType() ValueType {
	if m != nil {
		return m.Type
	}
	return ValueType_STRING
}

func (m *AttributeKeyValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *AttributeKeyValue) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *AttributeKeyValue) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *AttributeKeyValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *AttributeKeyValue) GetListOrMap() *ValueListOrMap {
	if m != nil {
		return m.ListOrMap
	}
	return nil
}

// StringKeyValue is a pair of key/value strings. This is the simpler (and faster) version
// of AttributeKeyValue that only supports string values.
type StringKeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringKeyValue) Reset()         { *m = StringKeyValue{} }
func (m *StringKeyValue) String() string { return proto.CompactTextString(m) }
func (*StringKeyValue) ProtoMessage()    {}
func (*StringKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *StringKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringKeyValue.Unmarshal(m, b)
}
func (m *StringKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringKeyValue.Marshal(b, m, deterministic)
}
func (m *StringKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringKeyValue.Merge(m, src)
}
func (m *StringKeyValue) XXX_Size() int {
	return xxx_messageInfo_StringKeyValue.Size(m)
}
func (m *StringKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringKeyValue proto.InternalMessageInfo

func (m *StringKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StringKeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// InstrumentationLibrary is a message representing the instrumentation library information
// such as the fully qualified name and version.
type InstrumentationLibrary struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstrumentationLibrary) Reset()         { *m = InstrumentationLibrary{} }
func (m *InstrumentationLibrary) String() string { return proto.CompactTextString(m) }
func (*InstrumentationLibrary) ProtoMessage()    {}
func (*InstrumentationLibrary) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *InstrumentationLibrary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstrumentationLibrary.Unmarshal(m, b)
}
func (m *InstrumentationLibrary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstrumentationLibrary.Marshal(b, m, deterministic)
}
func (m *InstrumentationLibrary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentationLibrary.Merge(m, src)
}
func (m *InstrumentationLibrary) XXX_Size() int {
	return xxx_messageInfo_InstrumentationLibrary.Size(m)
}
func (m *InstrumentationLibrary) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentationLibrary.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentationLibrary proto.InternalMessageInfo

func (m *InstrumentationLibrary) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstrumentationLibrary) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterEnum("baseline.ValueType", ValueType_name, ValueType_value)
	proto.RegisterType((*Resource)(nil), "baseline.Resource")
	proto.RegisterType((*AnyValue)(nil), "baseline.AnyValue")
	proto.RegisterType((*ValueListOrMap)(nil), "baseline.ValueListOrMap")
	proto.RegisterType((*AttributeKeyValue)(nil), "baseline.AttributeKeyValue")
	proto.RegisterType((*StringKeyValue)(nil), "baseline.StringKeyValue")
	proto.RegisterType((*InstrumentationLibrary)(nil), "baseline.InstrumentationLibrary")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9d, 0x24, 0x6d, 0x93, 0xd7, 0xb5, 0xc4, 0x51, 0x96, 0xc8, 0x2a, 0xc4, 0x1e, 0x34,
	0x08, 0x06, 0xb6, 0x82, 0x14, 0x3c, 0xb5, 0xeb, 0xaf, 0x62, 0x77, 0x5b, 0xa6, 0x75, 0xaf, 0x25,
	0x69, 0x87, 0x65, 0x30, 0x9d, 0x09, 0x33, 0x93, 0x42, 0x8f, 0xde, 0xfc, 0x5b, 0xbd, 0xf8, 0x2f,
	0xc8, 0x24, 0x69, 0xbb, 0xbb, 0x82, 0xeb, 0x69, 0x6f, 0x6f, 0xbe, 0xdf, 0xef, 0xbc, 0x37, 0x7c,
	0x5e, 0x02, 0x47, 0x4b, 0xb1, 0x5e, 0x0b, 0x1e, 0xe7, 0x52, 0x68, 0x81, 0xdd, 0x34, 0x51, 0x34,
	0x63, 0x9c, 0x76, 0x7f, 0x20, 0x70, 0x09, 0x55, 0xa2, 0x90, 0x4b, 0x8a, 0xdf, 0x03, 0x24, 0x5a,
	0x4b, 0x96, 0x16, 0x9a, 0xaa, 0x00, 0x85, 0x76, 0xd4, 0xee, 0x9d, 0xc4, 0xbb, 0x6c, 0x3c, 0xd8,
	0x79, 0x5f, 0xe9, 0xf6, 0x32, 0xc9, 0x0a, 0x4a, 0xae, 0xc5, 0x71, 0x1f, 0x82, 0x95, 0x14, 0x79,
	0x4e, 0x57, 0x8b, 0x83, 0xba, 0x58, 0x8a, 0x82, 0xeb, 0xc0, 0x0a, 0x51, 0xf4, 0x90, 0x1c, 0xd7,
	0xfe, 0xbe, 0x8f, 0x3a, 0x33, 0x6e, 0xf7, 0x37, 0x02, 0x77, 0xc0, 0xab, 0x96, 0xf8, 0x15, 0x38,
	0x7a, 0x9b, 0xd3, 0x00, 0x85, 0x28, 0xea, 0xf4, 0x1e, 0x1f, 0xa6, 0x97, 0xf6, 0x7c, 0x9b, 0x53,
	0x52, 0x06, 0xf0, 0x73, 0x80, 0x54, 0x88, 0x6c, 0xb1, 0x31, 0x7a, 0x39, 0xc1, 0x25, 0x9e, 0x51,
	0xaa, 0x3e, 0x2f, 0xe0, 0x48, 0x69, 0xc9, 0xf8, 0x55, 0x1d, 0xb0, 0x43, 0x14, 0x79, 0xa4, 0x5d,
	0x69, 0x55, 0xe4, 0x04, 0x3c, 0xc6, 0x75, 0xed, 0x3b, 0x21, 0x8a, 0x6c, 0xe2, 0x32, 0xae, 0xf7,
	0xf7, 0x57, 0xa2, 0x48, 0x33, 0x5a, 0xfb, 0x8d, 0x10, 0x45, 0x88, 0xb4, 0x2b, 0xad, 0x8a, 0xf4,
	0xa1, 0x9d, 0x31, 0xa5, 0x17, 0x42, 0x2e, 0xd6, 0x49, 0x1e, 0x34, 0x43, 0x14, 0xb5, 0x7b, 0xc1,
	0xad, 0x17, 0x8f, 0x99, 0xd2, 0x13, 0x79, 0x9e, 0xe4, 0xc4, 0xcb, 0x76, 0x65, 0xf7, 0x0a, 0x3a,
	0x37, 0x4d, 0xfc, 0x12, 0x1c, 0x63, 0xd7, 0xd0, 0xf1, 0x35, 0xe8, 0x35, 0x18, 0x52, 0xfa, 0xf8,
	0x0d, 0xd8, 0x66, 0x96, 0x75, 0xf7, 0x6e, 0x4c, 0xae, 0xfb, 0xd3, 0x82, 0x47, 0x7f, 0x59, 0xd8,
	0x07, 0xfb, 0x3b, 0xdd, 0x96, 0x88, 0x3d, 0x62, 0xca, 0x3d, 0x75, 0xeb, 0x2e, 0xea, 0xf7, 0x80,
	0xf5, 0xe6, 0x62, 0x9b, 0xb7, 0x17, 0xfb, 0x0e, 0x0e, 0x20, 0x83, 0xd6, 0xff, 0x33, 0xef, 0x43,
	0x67, 0x56, 0xbe, 0xf2, 0x1f, 0x18, 0x9e, 0x40, 0xe3, 0xf0, 0x39, 0x79, 0xa4, 0x3a, 0x74, 0x3f,
	0xc1, 0xf1, 0x88, 0x2b, 0x2d, 0x8b, 0x35, 0xe5, 0x3a, 0xd1, 0x4c, 0xf0, 0x31, 0x4b, 0x65, 0x22,
	0xb7, 0x18, 0x83, 0xc3, 0x93, 0x35, 0xad, 0x5b, 0x94, 0x35, 0x0e, 0xa0, 0xb5, 0xa1, 0x52, 0x31,
	0xc1, 0xeb, 0x2e, 0xbb, 0xe3, 0xeb, 0x11, 0x78, 0x7b, 0x9c, 0x18, 0xa0, 0x39, 0x9b, 0x93, 0xd1,
	0xc5, 0x67, 0xff, 0x01, 0x6e, 0x81, 0x3d, 0xba, 0x98, 0xfb, 0xc8, 0x88, 0x1f, 0x26, 0xdf, 0x86,
	0xe3, 0x8f, 0xbe, 0x85, 0x5d, 0x70, 0x86, 0x93, 0xc9, 0xd8, 0xb7, 0x4d, 0x35, 0x1e, 0xcd, 0xe6,
	0xbe, 0x63, 0x82, 0xe7, 0x83, 0xa9, 0xdf, 0x18, 0x7e, 0x81, 0x67, 0x4c, 0xc4, 0x22, 0xa7, 0x7c,
	0x49, 0xb9, 0x2a, 0x54, 0xf5, 0x5f, 0xc7, 0x5a, 0x26, 0x4b, 0x1a, 0x6f, 0x4e, 0x87, 0x30, 0x37,
	0xd5, 0xd4, 0x88, 0x53, 0xf4, 0xcb, 0x7a, 0x3a, 0xc9, 0x29, 0x3f, 0xab, 0x92, 0xa5, 0x18, 0x97,
	0x7e, 0x7c, 0x79, 0x9a, 0x36, 0xcb, 0x9b, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xf1,
	0x3f, 0x24, 0x21, 0x04, 0x00, 0x00,
}
