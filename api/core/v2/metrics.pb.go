// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/sensu/sensu-go/api/core/v2/metrics.proto

package v2

import (
	bytes "bytes"
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A Metrics is an event metrics payload specification.
type Metrics struct {
	// Handlers is a list of handlers for the metric points.
	Handlers []string `protobuf:"bytes,1,rep,name=handlers,proto3" json:"handlers"`
	// Points is a list of metric points (measurements).
	Points               []*MetricPoint `protobuf:"bytes,2,rep,name=points,proto3" json:"points"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7252db7bd2b8a8, []int{0}
}
func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(m, src)
}
func (m *Metrics) XXX_Size() int {
	return m.Size()
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetHandlers() []string {
	if m != nil {
		return m.Handlers
	}
	return nil
}

func (m *Metrics) GetPoints() []*MetricPoint {
	if m != nil {
		return m.Points
	}
	return nil
}

// A MetricPoint represents a single measurement.
type MetricPoint struct {
	// The metric point name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The metric point value.
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value"`
	// The metric point timestamp, time in nanoseconds since the Epoch.
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp"`
	// Tags is a list of metric tags (dimensions).
	Tags                 []*MetricTag `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MetricPoint) Reset()         { *m = MetricPoint{} }
func (m *MetricPoint) String() string { return proto.CompactTextString(m) }
func (*MetricPoint) ProtoMessage()    {}
func (*MetricPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7252db7bd2b8a8, []int{1}
}
func (m *MetricPoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetricPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetricPoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetricPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricPoint.Merge(m, src)
}
func (m *MetricPoint) XXX_Size() int {
	return m.Size()
}
func (m *MetricPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricPoint.DiscardUnknown(m)
}

var xxx_messageInfo_MetricPoint proto.InternalMessageInfo

func (m *MetricPoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricPoint) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MetricPoint) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MetricPoint) GetTags() []*MetricTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

// A MetricTag adds a dimension to a metric point.
type MetricTag struct {
	// The metric tag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The metric tag value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricTag) Reset()         { *m = MetricTag{} }
func (m *MetricTag) String() string { return proto.CompactTextString(m) }
func (*MetricTag) ProtoMessage()    {}
func (*MetricTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7252db7bd2b8a8, []int{2}
}
func (m *MetricTag) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetricTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetricTag.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetricTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricTag.Merge(m, src)
}
func (m *MetricTag) XXX_Size() int {
	return m.Size()
}
func (m *MetricTag) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricTag.DiscardUnknown(m)
}

var xxx_messageInfo_MetricTag proto.InternalMessageInfo

func (m *MetricTag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricTag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Metrics)(nil), "sensu.core.v2.Metrics")
	proto.RegisterType((*MetricPoint)(nil), "sensu.core.v2.MetricPoint")
	proto.RegisterType((*MetricTag)(nil), "sensu.core.v2.MetricTag")
}

func init() {
	proto.RegisterFile("github.com/sensu/sensu-go/api/core/v2/metrics.proto", fileDescriptor_2d7252db7bd2b8a8)
}

var fileDescriptor_2d7252db7bd2b8a8 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0xef, 0x69, 0x7a, 0x7b, 0x9b, 0xe9, 0xed, 0x66, 0x70, 0x31, 0x74, 0x31, 0x09, 0x5d,
	0x05, 0xc4, 0x09, 0x4d, 0xd5, 0xa5, 0x48, 0xf6, 0x82, 0x04, 0x57, 0xee, 0xa6, 0x75, 0x4c, 0x03,
	0x4d, 0x26, 0x64, 0x26, 0x79, 0x16, 0xdf, 0x40, 0x1f, 0xc1, 0x47, 0x70, 0xe9, 0x13, 0x04, 0x8d,
	0xbb, 0x3c, 0x81, 0x4b, 0xe9, 0xa4, 0xb4, 0x15, 0xba, 0x39, 0xe7, 0xe7, 0x3f, 0x3f, 0x3f, 0x1f,
	0x07, 0xcd, 0xe3, 0x44, 0xaf, 0xca, 0x05, 0x5b, 0xca, 0xd4, 0x57, 0x22, 0x53, 0x65, 0x37, 0xcf,
	0x62, 0xe9, 0xf3, 0x3c, 0xf1, 0x97, 0xb2, 0x10, 0x7e, 0x15, 0xf8, 0xa9, 0xd0, 0x45, 0xb2, 0x54,
	0x2c, 0x2f, 0xa4, 0x96, 0x78, 0x6c, 0x32, 0x6c, 0x73, 0x64, 0x55, 0x30, 0x39, 0x3f, 0xe8, 0x88,
	0x65, 0x2c, 0x7d, 0x93, 0x5a, 0x94, 0x8f, 0xd7, 0xd5, 0x8c, 0xcd, 0xd9, 0xcc, 0x98, 0xc6, 0x33,
	0xaa, 0x2b, 0x99, 0x2a, 0xf4, 0xef, 0xa6, 0x6b, 0xc5, 0x1e, 0x1a, 0xae, 0x78, 0xf6, 0xb0, 0x16,
	0x85, 0x22, 0xe0, 0x5a, 0x9e, 0x1d, 0xfe, 0x6f, 0x6b, 0x67, 0xe7, 0x45, 0x3b, 0x85, 0xaf, 0xd0,
	0x20, 0x97, 0x49, 0xa6, 0x15, 0xe9, 0xb9, 0x96, 0x37, 0x0a, 0x26, 0xec, 0x17, 0x0a, 0xeb, 0x1a,
	0x6f, 0x37, 0x91, 0x10, 0xb5, 0xb5, 0xb3, 0x4d, 0x47, 0xdb, 0x3d, 0x7d, 0x06, 0x34, 0x3a, 0xc8,
	0x60, 0x8c, 0xfa, 0x19, 0x4f, 0x05, 0x01, 0x17, 0x3c, 0x3b, 0x32, 0x1a, 0x3b, 0xe8, 0x6f, 0xc5,
	0xd7, 0xa5, 0x20, 0x3d, 0x17, 0x3c, 0x08, 0xed, 0xb6, 0x76, 0x3a, 0x23, 0xea, 0x16, 0x3e, 0x45,
	0xb6, 0x4e, 0x52, 0xa1, 0x34, 0x4f, 0x73, 0x62, 0xb9, 0xe0, 0x59, 0xe1, 0xb8, 0xad, 0x9d, 0xbd,
	0x19, 0xed, 0x25, 0xbe, 0x44, 0x7d, 0xcd, 0x63, 0x45, 0xfa, 0x86, 0x97, 0x1c, 0xe5, 0xbd, 0xe3,
	0x71, 0x38, 0x6c, 0x6b, 0xc7, 0x24, 0x23, 0x33, 0xa7, 0x17, 0xc8, 0xde, 0x1d, 0x8f, 0x62, 0x9e,
	0x1c, 0x62, 0xda, 0x5b, 0xb6, 0xd0, 0xfd, 0xfe, 0xa4, 0xf0, 0xd2, 0x50, 0x78, 0x6d, 0x28, 0xbc,
	0x35, 0x14, 0xde, 0x1b, 0x0a, 0x1f, 0x0d, 0x85, 0xa7, 0x2f, 0xfa, 0xe7, 0xbe, 0x57, 0x05, 0x8b,
	0x81, 0x79, 0xff, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0x75, 0x57, 0x7b, 0x36, 0xfa, 0x01, 0x00,
	0x00,
}

func (this *Metrics) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Metrics)
	if !ok {
		that2, ok := that.(Metrics)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Handlers) != len(that1.Handlers) {
		return false
	}
	for i := range this.Handlers {
		if this.Handlers[i] != that1.Handlers[i] {
			return false
		}
	}
	if len(this.Points) != len(that1.Points) {
		return false
	}
	for i := range this.Points {
		if !this.Points[i].Equal(that1.Points[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MetricPoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetricPoint)
	if !ok {
		that2, ok := that.(MetricPoint)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if len(this.Tags) != len(that1.Tags) {
		return false
	}
	for i := range this.Tags {
		if !this.Tags[i].Equal(that1.Tags[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MetricTag) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetricTag)
	if !ok {
		that2, ok := that.(MetricTag)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Metrics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metrics) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metrics) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Points) > 0 {
		for iNdEx := len(m.Points) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Points[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetrics(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Handlers) > 0 {
		for iNdEx := len(m.Handlers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Handlers[iNdEx])
			copy(dAtA[i:], m.Handlers[iNdEx])
			i = encodeVarintMetrics(dAtA, i, uint64(len(m.Handlers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MetricPoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricPoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetricPoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tags[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetrics(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Timestamp != 0 {
		i = encodeVarintMetrics(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if m.Value != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x11
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MetricTag) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricTag) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetricTag) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetrics(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetrics(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedMetrics(r randyMetrics, easy bool) *Metrics {
	this := &Metrics{}
	v1 := r.Intn(10)
	this.Handlers = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.Handlers[i] = string(randStringMetrics(r))
	}
	if r.Intn(5) != 0 {
		v2 := r.Intn(5)
		this.Points = make([]*MetricPoint, v2)
		for i := 0; i < v2; i++ {
			this.Points[i] = NewPopulatedMetricPoint(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMetrics(r, 3)
	}
	return this
}

func NewPopulatedMetricPoint(r randyMetrics, easy bool) *MetricPoint {
	this := &MetricPoint{}
	this.Name = string(randStringMetrics(r))
	this.Value = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Value *= -1
	}
	this.Timestamp = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Timestamp *= -1
	}
	if r.Intn(5) != 0 {
		v3 := r.Intn(5)
		this.Tags = make([]*MetricTag, v3)
		for i := 0; i < v3; i++ {
			this.Tags[i] = NewPopulatedMetricTag(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMetrics(r, 5)
	}
	return this
}

func NewPopulatedMetricTag(r randyMetrics, easy bool) *MetricTag {
	this := &MetricTag{}
	this.Name = string(randStringMetrics(r))
	this.Value = string(randStringMetrics(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMetrics(r, 3)
	}
	return this
}

type randyMetrics interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMetrics(r randyMetrics) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMetrics(r randyMetrics) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneMetrics(r)
	}
	return string(tmps)
}
func randUnrecognizedMetrics(r randyMetrics, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMetrics(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMetrics(dAtA []byte, r randyMetrics, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMetrics(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateMetrics(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateMetrics(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMetrics(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMetrics(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMetrics(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMetrics(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Metrics) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Handlers) > 0 {
		for _, s := range m.Handlers {
			l = len(s)
			n += 1 + l + sovMetrics(uint64(l))
		}
	}
	if len(m.Points) > 0 {
		for _, e := range m.Points {
			l = e.Size()
			n += 1 + l + sovMetrics(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MetricPoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMetrics(uint64(l))
	}
	if m.Value != 0 {
		n += 9
	}
	if m.Timestamp != 0 {
		n += 1 + sovMetrics(uint64(m.Timestamp))
	}
	if len(m.Tags) > 0 {
		for _, e := range m.Tags {
			l = e.Size()
			n += 1 + l + sovMetrics(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MetricTag) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMetrics(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMetrics(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMetrics(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetrics(x uint64) (n int) {
	return sovMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metrics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
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
			return fmt.Errorf("proto: Metrics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metrics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handlers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handlers = append(m.Handlers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Points", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
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
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Points = append(m.Points, &MetricPoint{})
			if err := m.Points[len(m.Points)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MetricPoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
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
			return fmt.Errorf("proto: MetricPoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricPoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
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
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, &MetricTag{})
			if err := m.Tags[len(m.Tags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MetricTag) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
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
			return fmt.Errorf("proto: MetricTag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricTag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetrics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetrics
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
					return 0, ErrIntOverflowMetrics
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
					return 0, ErrIntOverflowMetrics
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
				return 0, ErrInvalidLengthMetrics
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetrics
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetrics
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetrics        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetrics          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetrics = fmt.Errorf("proto: unexpected end of group")
)
