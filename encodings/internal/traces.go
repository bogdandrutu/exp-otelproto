package internal

import (
	"io"
	"math"
	"unsafe"

	"github.com/golang/protobuf/proto"
	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental"
)

type TraceExportRequest struct {
	// Unique sequential ID generated by the client.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Telemetry data. An array of ResourceSpans.
	ResourceSpans []*ResourceSpans `protobuf:"bytes,2,rep,name=resourceSpans,proto3" json:"resourceSpans,omitempty"`
}

// A collection of spans from a Resource.
type ResourceSpans struct {
	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Spans    []*Span   `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
}

// Resource information. This describes the source of telemetry data.
type Resource struct {
	// labels is a collection of attributes that describe the resource. See OpenTelemetry
	// specification semantic conventions for standardized label names:
	// https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/data-semantic-conventions.md
	Labels map[string]*AttributeValue `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// dropped_labels_count is the number of dropped labels. If the value is 0, then
	// no labels were dropped.
	DroppedLabelsCount uint32 `protobuf:"varint,2,opt,name=dropped_labels_count,json=droppedLabelsCount,proto3" json:"dropped_labels_count,omitempty"`
}

type AttributeValue struct {
	// type of the value.
	typ         experimental.ValueType `protobuf:"varint,2,opt,name=type,proto3,enum=experimental.ValueType" json:"type,omitempty"`
	stringValue string                 `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	intValue    int64                  `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	doubleValue float64                `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	boolValue   bool
}

func (a *AttributeValue) Type() experimental.ValueType {
	return a.typ
}

func (a *AttributeValue) String() string {
	return a.stringValue
}

func (a *AttributeValue) Set(t experimental.ValueType, s string, i int64, d float64, b bool) {
	a.typ = t
	a.stringValue = s
	a.intValue = i
	a.doubleValue = d
	a.boolValue = b
}

func (a *AttributeValue) Int() int64 {
	return a.intValue
}

func (a *AttributeValue) Double() float64 {
	return a.doubleValue
}

func (a *AttributeValue) Bool() bool {
	return a.boolValue
}

type Span struct {
	// A unique identifier for a trace. All spans from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
	// is considered invalid.
	//
	// This field is semantically required. Receiver should generate new
	// random trace_id if empty or invalid trace_id was received.
	//
	// This field is required.
	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes is considered
	// invalid.
	//
	// This field is semantically required. Receiver should generate new
	// random span_id if empty or invalid span_id was received.
	//
	// This field is required.
	SpanId []byte `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// tracestate conveys information about request position in multiple distributed tracing graphs.
	// It is a tracestate in w3c-trace-context format: https://www.w3.org/TR/trace-context/#tracestate-header
	// See also https://github.com/w3c/distributed-tracing for more details about this field.
	Tracestate string `protobuf:"bytes,3,opt,name=tracestate,proto3" json:"tracestate,omitempty"`
	// The `span_id` of this span's parent span. If this is a root span, then this
	// field must be empty. The ID is an 8-byte array.
	ParentSpanId []byte `protobuf:"bytes,4,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	// A description of the span's operation.
	//
	// For example, the name can be a qualified method name or a file name
	// and a line number where the operation is called. A best practice is to use
	// the same display name at the same call point in an application.
	// This makes it easier to correlate spans in different traces.
	//
	// This field is semantically required to be set to non-empty string.
	// When null or empty string received - receiver may use string "name"
	// as a replacement. There might be smarted algorithms implemented by
	// receiver to fix the empty span name.
	//
	// This field is required.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Distinguishes between spans generated in a particular context. For example,
	// two spans with the same name may be distinguished using `CLIENT` (caller)
	// and `SERVER` (callee) to identify queueing latency associated with the span.
	Kind experimental.Span_SpanKind `protobuf:"varint,6,opt,name=kind,proto3,enum=experimental.Span_SpanKind" json:"kind,omitempty"`
	// start_time_unixnano is the start time of the span. On the client side, this is the time
	// kept by the local machine where the span execution starts. On the server side, this
	// is the time when the server's application handler starts running.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	//
	// This field is semantically required and it is expected that end_time >= start_time.
	StartTimeUnixnano uint64 `protobuf:"fixed64,7,opt,name=start_time_unixnano,json=startTimeUnixnano,proto3" json:"start_time_unixnano,omitempty"`
	// end_time_unixnano is the end time of the span. On the client side, this is the time
	// kept by the local machine where the span execution ends. On the server side, this
	// is the time when the server application handler stops running.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	//
	// This field is semantically required and it is expected that end_time >= start_time.
	EndTimeUnixnano uint64 `protobuf:"fixed64,8,opt,name=end_time_unixnano,json=endTimeUnixnano,proto3" json:"end_time_unixnano,omitempty"`
	// attributes is a collection of key/value pairs. The value can be a string,
	// an integer, a double or the Boolean values `true` or `false`. Note, global attributes
	// like server name can be set using the resource API. Examples of attributes:
	//
	//     "/http/user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	//     "/http/server_latency": 300
	//     "abc.com/myattribute": true
	//     "abc.com/score": 10.239
	Attributes map[string]*AttributeValue `protobuf:"bytes,9,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// dropped_attributes_count is the number of attributes that were discarded. Attributes
	// can be discarded because their keys are too long or because there are too many
	// attributes. If this value is 0, then no attributes were dropped.
	DroppedAttributesCount uint32 `protobuf:"varint,10,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	// events is a collection of Event items.
	Events []*Span_Event `protobuf:"bytes,11,rep,name=events,proto3" json:"events,omitempty"`
	// dropped_events_count is the number of dropped events. If the value is 0, then no
	// events were dropped.
	DroppedEventsCount uint32 `protobuf:"varint,12,opt,name=dropped_events_count,json=droppedEventsCount,proto3" json:"dropped_events_count,omitempty"`
	// links is a collection of Links, which are references from this span to a span
	// in the same or different trace.
	Links []*Span_Link `protobuf:"bytes,13,rep,name=links,proto3" json:"links,omitempty"`
	// dropped_links_count is the number of dropped links after the maximum size was
	// enforced. If this value is 0, then no links were dropped.
	DroppedLinksCount uint32 `protobuf:"varint,14,opt,name=dropped_links_count,json=droppedLinksCount,proto3" json:"dropped_links_count,omitempty"`
	// An optional final status for this span. Semantically when Status
	// wasn't set it is means span ended without errors and assume
	// Status.Ok (code = 0).
	Status *experimental.Status `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
}

type Span_Event struct {
	// time_unixnano is the time the event occurred.
	TimeUnixnano uint64 `protobuf:"fixed64,1,opt,name=time_unixnano,json=timeUnixnano,proto3" json:"time_unixnano,omitempty"`
	// description is a user-supplied text.
	Name string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// attributes is a collection of attribute key/value pairs on the event.
	Attributes map[string]*AttributeValue `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0,
	// then no attributes were dropped.
	DroppedAttributesCount uint32 `protobuf:"varint,4,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
}

type Span_Link struct {
	// A unique identifier of a trace that this linked span is part of. The ID is a
	// 16-byte array.
	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for the linked span. The ID is an 8-byte array.
	SpanId []byte `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// The tracestate associated with the link.
	Tracestate string `protobuf:"bytes,3,opt,name=tracestate,proto3" json:"tracestate,omitempty"`
	// attributes is a collection of attribute key/value pairs on the link.
	Attributes map[string]*AttributeValue `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0,
	// then no attributes were dropped.
	DroppedAttributesCount uint32 `protobuf:"varint,5,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
}

func Marshal(tes *TraceExportRequest) ([]byte, error) {
	buf := proto.NewBuffer(make([]byte, 0, 30000))
	encodeVarint(buf, 1, tes.Id)
	for _, rs := range tes.ResourceSpans {
		if rs == nil {
			continue
		}
		err := encodeSubmessage(buf, 2, MarshalResourceSpans, unsafe.Pointer(rs))
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func encodeSubmessage(
	buf *proto.Buffer,
	fieldNum uint64,
	encoder func(buf *proto.Buffer, subObj unsafe.Pointer) error,
	subObj unsafe.Pointer,
) error {
	buf.EncodeVarint((fieldNum << 3) | proto.WireBytes)
	buf.EncodeVarint(0)
	start := len(buf.Bytes())
	if err := encoder(buf, subObj); err != nil {
		return err
	}
	insertLen(buf, start)
	return nil
}

func encodeMapKeyValue(
	buf *proto.Buffer,
	fieldNum uint64,
	encoder func(buf *proto.Buffer, key string, valObj unsafe.Pointer) error,
	key string,
	valObj unsafe.Pointer,
) error {
	buf.EncodeVarint((fieldNum << 3) | proto.WireBytes)
	buf.EncodeVarint(0)
	start := len(buf.Bytes())
	if err := encoder(buf, key, valObj); err != nil {
		return err
	}
	insertLen(buf, start)
	return nil
}

func insertLen(buf *proto.Buffer, start int) {
	b := buf.Bytes()
	ln := uint64(len(b) - start)
	if ln > 127 {
		// Reencode length and shift the data.
		lnSz := proto.SizeVarint(ln)

		if cap(b) < len(b)+lnSz-1 {
			b = append(b, make([]byte, lnSz-1)...)
		}
		b = b[:len(b)+lnSz-1]

		copy(b[start+lnSz-1:], b[start:])

		var n int
		for n = 0; ln > 127; n++ {
			b[start-1+n] = 0x80 | uint8(ln&0x7F)
			ln >>= 7
		}
		b[start-1+n] = uint8(ln)

		buf.SetBuf(b)
	} else {
		b[start-1] = byte(ln)
	}
}

func MarshalResourceSpans(buf *proto.Buffer, p unsafe.Pointer) error {
	rs := (*ResourceSpans)(p)

	if rs.Resource != nil {
		err := encodeSubmessage(buf, 1, MarshalResource, unsafe.Pointer(rs.Resource))
		if err != nil {
			return err
		}
	}
	for _, s := range rs.Spans {
		if s == nil {
			continue
		}
		err := encodeSubmessage(buf, 2, MarshalSpan, unsafe.Pointer(s))
		if err != nil {
			return err
		}
	}
	return nil
}

func MarshalResource(buf *proto.Buffer, p unsafe.Pointer) error {
	r := (*Resource)(p)

	MarshalAttributesMap(buf, 1, r.Labels)
	if r.DroppedLabelsCount != 0 {
		encodeVarint(buf, 2, uint64(r.DroppedLabelsCount))
	}
	return nil
}

func MarshalAttributesMap(buf *proto.Buffer, fieldNum uint64, m map[string]*AttributeValue) {
	for k, l := range m {
		buf.EncodeVarint((fieldNum << 3) | proto.WireBytes)
		MarshalAttribute(buf, k, l)
	}
}

func MarshalAttribute(buf *proto.Buffer, key string, val *AttributeValue) {
	buf.EncodeVarint(0)
	start := len(buf.Bytes())
	encodeString(buf, 1, key)
	encodeMapKeyValue(buf, 2, MarshalAttributeKeyValue, key, unsafe.Pointer(val))
	insertLen(buf, start)
}

func MarshalAttributeKeyValue(buf *proto.Buffer, key string, val unsafe.Pointer) error {
	v := (*AttributeValue)(val)

	encodeString(buf, 1, key)
	encodeVarint(buf, 2, uint64(v.Type()))

	switch v.Type() {
	case experimental.ValueType_STRING:
		encodeString(buf, 3, v.String())
	case experimental.ValueType_INT:
		encodeVarint(buf, 4, uint64(v.Int()))
	case experimental.ValueType_DOUBLE:
		encodeFixed64(buf, 5, math.Float64bits(v.Double()))
	case experimental.ValueType_BOOL:
		if v.Bool() {
			encodeVarint(buf, 6, 1)
		}
	}

	return nil
}

func MarshalSpan(buf *proto.Buffer, p unsafe.Pointer) error {
	s := (*Span)(p)

	encodeBytes(buf, 1, s.TraceId)
	encodeBytes(buf, 2, s.SpanId)
	encodeBytes(buf, 3, s.ParentSpanId)
	encodeString(buf, 4, s.Tracestate)
	encodeString(buf, 5, s.Name)
	encodeVarint(buf, 6, uint64(s.Kind))
	encodeFixed64(buf, 7, s.StartTimeUnixnano)
	encodeFixed64(buf, 8, s.EndTimeUnixnano)
	MarshalAttributesMap(buf, 9, s.Attributes)
	encodeVarint(buf, 10, uint64(s.DroppedAttributesCount))
	for _, e := range s.Events {
		encodeSubmessage(buf, 11, MarshalEvent, unsafe.Pointer(e))
	}
	encodeVarint(buf, 12, uint64(s.DroppedEventsCount))
	//for _, l := range s.Links {
	//	//encodeSubmessage(buf, 13, l)
	//}
	encodeVarint(buf, 14, uint64(s.DroppedLinksCount))
	//encodeSubmessage(buf, 15, s.Status)
	return nil
}

func MarshalEvent(buf *proto.Buffer, p unsafe.Pointer) error {
	e := (*Span_Event)(p)

	encodeFixed64(buf, 1, e.TimeUnixnano)
	encodeString(buf, 2, e.Name)
	MarshalAttributesMap(buf, 3, e.Attributes)
	encodeVarint(buf, 4, uint64(e.DroppedAttributesCount))
	return nil
}

func encodeString(buf *proto.Buffer, fieldNum uint64, str string) {
	if len(str) != 0 {
		buf.EncodeVarint((fieldNum << 3) | proto.WireBytes)
		buf.EncodeStringBytes(str)
	}
}

func encodeBytes(buf *proto.Buffer, fieldNum uint64, b []byte) {
	if len(b) != 0 {
		buf.EncodeVarint((fieldNum << 3) | proto.WireBytes)
		buf.EncodeRawBytes(b)
	}
}

func encodeFixed64(buf *proto.Buffer, fieldNum uint64, val uint64) {
	if val != 0 {
		buf.EncodeVarint((fieldNum << 3) | proto.WireFixed64)
		buf.EncodeFixed64(val)
	}
}

func encodeFixed32(buf *proto.Buffer, fieldNum uint64, val uint64) {
	if val != 0 {
		buf.EncodeVarint((fieldNum << 3) | proto.WireFixed64)
		buf.EncodeFixed32(val)
	}
}

func encodeVarint(buf *proto.Buffer, fieldNum uint64, val uint64) {
	if val != 0 {
		buf.EncodeVarint((fieldNum << 3) | proto.WireVarint)
		buf.EncodeVarint(val)
	}
}

func encodeFieldKey(buf *proto.Buffer, fieldNum uint64, wireType uint64) {
	buf.EncodeVarint((fieldNum << 3) | wireType)
}

func ResourceSpansFromBuf(b *proto.Buffer, s int) (*ResourceSpans, error) {
	r := &ResourceSpans{}
	for {
		fn, wt, err := decodeFieldTag(b)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				return r, nil
			}
			return nil, err
		}
		switch fn {
		case 1:
			len, err := b.DecodeVarint()
			if err != nil {
				return nil, err
			}
			if len > 0 {
				r.Resource, err = ResourceFromBuf(b, len)
			}
		case 2:
			len, err := b.DecodeVarint()
			if err != nil {
				return nil, err
			}
			span := &Span{}
			if len > 0 {
				err = SpanFromBuf(b, len, span)
			}
			r.Spans = append(r.Spans, span)
		}
		if err != nil {
			return nil, err
		}
		if wt != 2 {
			return nil, nil
		}
	}
}

func ResourceFromBuf(b *proto.Buffer, s uint64) (r *Resource, err error) {
	for {
		fn, wt, err := decodeFieldTag(b)
		if err != nil {
			return nil, err
		}
		switch fn {
		case 1:
			// r.Labels,err = ResourceFromBuf(b)
		case 2:
			// r.DroppedLabelsCount,err = b.DecodeVarint()
		}
		if err != nil {
			return nil, err
		}
		if wt != 0 {
			return nil, nil
		}
	}
}

func SpanFromBuf(b *proto.Buffer, len uint64, toSpan *Span) error {
	for {
		fn, wt, err := decodeFieldTag(b)
		if err != nil {
			return err
		}
		switch fn {
		case 5:
			toSpan.Name, err = b.DecodeStringBytes()
		}
		if err != nil {
			return err
		}
		if wt != 2 {
			return err
		}
	}

}

func decodeFieldTag(b *proto.Buffer) (field_num uint64, wire_type uint64, err error) {
	var v uint64
	v, err = b.DecodeVarint()
	if err != nil {
		return
	}
	wire_type = v & 7
	field_num = v >> 3
	return
}

//func FromOtlp(tes *experimental.TraceExportRequest) *TraceExportRequest {
//	r := &TraceExportRequest{}
//	r.ResourceSpans = make([]*ResourceSpans, len(tes.ResourceSpans))
//	for i, s := range tes.ResourceSpans {
//		r.ResourceSpans[i] = ResourceSpansFromOtlp(s)
//	}
//	return r
//}

//func ResourceSpansFromOtlp(spans *experimental.ResourceSpans) *ResourceSpans {
//	return &ResourceSpans{
//		Resource: ResourceFromOtlp(spans.Resource),
//		Spans:    SpansFromOtlp(spans.Spans),
//	}
//}

func ResourceFromOtlp(resource *experimental.Resource) *Resource {
	return &Resource{
		Labels:             AttrsFromOtlp(resource.Attributes),
		DroppedLabelsCount: resource.DroppedAttributesCount,
	}
}

func AttrsFromOtlp(attrs []*experimental.AttributeKeyValue) map[string]*AttributeValue {
	ptrs := make(map[string]*AttributeValue, len(attrs))
	content := make([]AttributeValue, len(attrs))
	for i, attr := range attrs {
		ptrs[attr.Key] = &content[i]
		AttrFromOtlp(&content[i], attr)
	}
	return ptrs
}

func AttrFromOtlp(dest *AttributeValue, src *experimental.AttributeKeyValue) {
	dest.Set(
		src.Type,
		src.StringValue,
		src.IntValue,
		src.DoubleValue,
		src.BoolValue,
	)
}

func SpansFromOtlp(spans []*experimental.Span) []*Span {
	ptrs := make([]*Span, len(spans))
	content := make([]Span, len(spans))
	for i, s := range spans {
		SpanFromOtlp(s, &content[i])
		ptrs[i] = &content[i]
	}
	return ptrs
}

func SpanFromOtlp(src *experimental.Span, dest *Span) {
	dest.TraceId = src.TraceId
	dest.SpanId = src.SpanId
	dest.Tracestate = src.TraceState
	dest.ParentSpanId = src.ParentSpanId
	dest.Name = src.Name
	dest.Kind = src.Kind
	dest.StartTimeUnixnano = src.StartTimeUnixNano
	dest.EndTimeUnixnano = src.EndTimeUnixNano
	dest.Attributes = AttrsFromOtlp(src.Attributes)
	dest.DroppedAttributesCount = src.DroppedAttributesCount
	dest.Events = EventsFromOtlp(src.Events)
	dest.DroppedEventsCount = src.DroppedEventsCount
	dest.Links = LinksFromOtlp(src.Links)
	dest.DroppedLinksCount = src.DroppedLinksCount
	dest.Status = src.Status
}

func EventsFromOtlp(events []*experimental.Span_Event) []*Span_Event {
	r := make([]*Span_Event, len(events))
	for i, e := range events {
		r[i] = EventFromOtlp(e)
	}
	return r
}

func EventFromOtlp(e *experimental.Span_Event) *Span_Event {
	return &Span_Event{
		TimeUnixnano:           e.TimeUnixNano,
		Name:                   e.Name,
		Attributes:             AttrsFromOtlp(e.Attributes),
		DroppedAttributesCount: e.DroppedAttributesCount,
	}
}

func LinksFromOtlp(links []*experimental.Span_Link) []*Span_Link {
	r := make([]*Span_Link, len(links))
	for i, e := range links {
		r[i] = LinkFromOtlp(e)
	}
	return r
}

func LinkFromOtlp(l *experimental.Span_Link) *Span_Link {
	return &Span_Link{
		TraceId:                l.TraceId,
		SpanId:                 l.SpanId,
		Attributes:             AttrsFromOtlp(l.Attributes),
		DroppedAttributesCount: l.DroppedAttributesCount,
	}
}

//func ToOtlp(tes *TraceExportRequest) *experimental.TraceExportRequest {
//	r := make([]*experimental.ResourceSpans, len(tes.ResourceSpans))
//	for i, s := range tes.ResourceSpans {
//		r[i] = ResourceSpansToOtlp(s)
//	}
//	return &experimental.TraceExportRequest{
//		ResourceSpans: r,
//	}
//}

//func ResourceSpansToOtlp(spans *ResourceSpans) *experimental.ResourceSpans {
//	return &experimental.ResourceSpans{
//		Resource: ResourceToOtlp(spans.Resource),
//		Spans:    SpansToOtlp(spans.Spans),
//	}
//}

func ResourceToOtlp(resource *Resource) *experimental.Resource {
	return &experimental.Resource{
		Attributes:             AttrsToOtlp(resource.Labels),
		DroppedAttributesCount: resource.DroppedLabelsCount,
	}
}

func AttrsToOtlp(attrs map[string]*AttributeValue) []*experimental.AttributeKeyValue {
	ptrs := make([]*experimental.AttributeKeyValue, len(attrs))
	content := make([]experimental.AttributeKeyValue, len(attrs))
	i := 0
	for _, attr := range attrs {
		ptrs[i] = &content[i]
		AttrToOtlp(ptrs[i], attr)
		i++
	}
	return ptrs
}

func AttrToOtlp(dest *experimental.AttributeKeyValue, src *AttributeValue) {
	dest.Type = src.Type()
	dest.StringValue = src.String()
	dest.IntValue = src.Int()
	dest.DoubleValue = src.Double()
	dest.BoolValue = src.Bool()
}

func SpansToOtlp(spans []*Span) []*experimental.Span {
	ptrs := make([]*experimental.Span, len(spans))
	content := make([]experimental.Span, len(spans))
	for i, s := range spans {
		ptrs[i] = &content[i]
		SpanToOtlp(s, ptrs[i])
	}
	return ptrs
}

func SpanToOtlp(src *Span, dest *experimental.Span) {
	dest.TraceId = src.TraceId
	dest.SpanId = src.SpanId
	dest.TraceState = src.Tracestate
	dest.ParentSpanId = src.ParentSpanId
	dest.Name = src.Name
	dest.Kind = src.Kind
	dest.StartTimeUnixNano = src.StartTimeUnixnano
	dest.EndTimeUnixNano = src.EndTimeUnixnano
	dest.Attributes = AttrsToOtlp(src.Attributes)
	dest.DroppedAttributesCount = src.DroppedAttributesCount
	dest.Events = EventsToOtlp(src.Events)
	dest.DroppedEventsCount = src.DroppedEventsCount
	dest.Links = LinksToOtlp(src.Links)
	dest.DroppedLinksCount = src.DroppedLinksCount
	dest.Status = src.Status
}

func EventsToOtlp(events []*Span_Event) (r []*experimental.Span_Event) {
	r = make([]*experimental.Span_Event, len(events))
	for i, e := range events {
		r[i] = EventToOtlp(e)
	}
	return
}

func EventToOtlp(e *Span_Event) *experimental.Span_Event {
	return &experimental.Span_Event{
		TimeUnixNano:           e.TimeUnixnano,
		Name:                   e.Name,
		Attributes:             AttrsToOtlp(e.Attributes),
		DroppedAttributesCount: e.DroppedAttributesCount,
	}
}

func LinksToOtlp(links []*Span_Link) (r []*experimental.Span_Link) {
	r = make([]*experimental.Span_Link, len(links))
	for i, e := range links {
		r[i] = LinkToOtlp(e)
	}
	return
}

func LinkToOtlp(l *Span_Link) *experimental.Span_Link {
	return &experimental.Span_Link{
		TraceId:                l.TraceId,
		SpanId:                 l.SpanId,
		Attributes:             AttrsToOtlp(l.Attributes),
		DroppedAttributesCount: l.DroppedAttributesCount,
	}
}
