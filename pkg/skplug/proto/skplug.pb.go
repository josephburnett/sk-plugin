// Code generated by protoc-gen-go. DO NOT EDIT.
// source: skplug.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EventRequest_EventType int32

const (
	EventRequest_CREATE EventRequest_EventType = 0
	EventRequest_UPDATE EventRequest_EventType = 1
	EventRequest_DELETE EventRequest_EventType = 2
)

var EventRequest_EventType_name = map[int32]string{
	0: "CREATE",
	1: "UPDATE",
	2: "DELETE",
}

var EventRequest_EventType_value = map[string]int32{
	"CREATE": 0,
	"UPDATE": 1,
	"DELETE": 2,
}

func (x EventRequest_EventType) String() string {
	return proto.EnumName(EventRequest_EventType_name, int32(x))
}

func (EventRequest_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a95933fe5266f40f, []int{3, 0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95933fe5266f40f, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Autoscaler struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Yaml                 string   `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Autoscaler) Reset()         { *m = Autoscaler{} }
func (m *Autoscaler) String() string { return proto.CompactTextString(m) }
func (*Autoscaler) ProtoMessage()    {}
func (*Autoscaler) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95933fe5266f40f, []int{1}
}

func (m *Autoscaler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Autoscaler.Unmarshal(m, b)
}
func (m *Autoscaler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Autoscaler.Marshal(b, m, deterministic)
}
func (m *Autoscaler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Autoscaler.Merge(m, src)
}
func (m *Autoscaler) XXX_Size() int {
	return xxx_messageInfo_Autoscaler.Size(m)
}
func (m *Autoscaler) XXX_DiscardUnknown() {
	xxx_messageInfo_Autoscaler.DiscardUnknown(m)
}

var xxx_messageInfo_Autoscaler proto.InternalMessageInfo

func (m *Autoscaler) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Autoscaler) GetYaml() string {
	if m != nil {
		return m.Yaml
	}
	return ""
}

type Pod struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	LastTransition       int64    `protobuf:"varint,3,opt,name=last_transition,json=lastTransition,proto3" json:"last_transition,omitempty"`
	CpuRequest           int32    `protobuf:"varint,4,opt,name=cpu_request,json=cpuRequest,proto3" json:"cpu_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95933fe5266f40f, []int{2}
}

func (m *Pod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pod.Unmarshal(m, b)
}
func (m *Pod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pod.Marshal(b, m, deterministic)
}
func (m *Pod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pod.Merge(m, src)
}
func (m *Pod) XXX_Size() int {
	return xxx_messageInfo_Pod.Size(m)
}
func (m *Pod) XXX_DiscardUnknown() {
	xxx_messageInfo_Pod.DiscardUnknown(m)
}

var xxx_messageInfo_Pod proto.InternalMessageInfo

func (m *Pod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pod) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Pod) GetLastTransition() int64 {
	if m != nil {
		return m.LastTransition
	}
	return 0
}

func (m *Pod) GetCpuRequest() int32 {
	if m != nil {
		return m.CpuRequest
	}
	return 0
}

type EventRequest struct {
	Partition string                 `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Time      int64                  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Type      EventRequest_EventType `protobuf:"varint,3,opt,name=type,proto3,enum=proto.EventRequest_EventType" json:"type,omitempty"`
	// Types that are valid to be assigned to ObjectOneof:
	//	*EventRequest_Autoscaler
	//	*EventRequest_Pod
	ObjectOneof          isEventRequest_ObjectOneof `protobuf_oneof:"object_oneof"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95933fe5266f40f, []int{3}
}

func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (m *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(m, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func (m *EventRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *EventRequest) GetType() EventRequest_EventType {
	if m != nil {
		return m.Type
	}
	return EventRequest_CREATE
}

type isEventRequest_ObjectOneof interface {
	isEventRequest_ObjectOneof()
}

type EventRequest_Autoscaler struct {
	Autoscaler *Autoscaler `protobuf:"bytes,4,opt,name=autoscaler,proto3,oneof"`
}

type EventRequest_Pod struct {
	Pod *Pod `protobuf:"bytes,5,opt,name=pod,proto3,oneof"`
}

func (*EventRequest_Autoscaler) isEventRequest_ObjectOneof() {}

func (*EventRequest_Pod) isEventRequest_ObjectOneof() {}

func (m *EventRequest) GetObjectOneof() isEventRequest_ObjectOneof {
	if m != nil {
		return m.ObjectOneof
	}
	return nil
}

func (m *EventRequest) GetAutoscaler() *Autoscaler {
	if x, ok := m.GetObjectOneof().(*EventRequest_Autoscaler); ok {
		return x.Autoscaler
	}
	return nil
}

func (m *EventRequest) GetPod() *Pod {
	if x, ok := m.GetObjectOneof().(*EventRequest_Pod); ok {
		return x.Pod
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EventRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EventRequest_Autoscaler)(nil),
		(*EventRequest_Pod)(nil),
	}
}

type Stat struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	PodName              string   `protobuf:"bytes,2,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	Metric               string   `protobuf:"bytes,3,opt,name=metric,proto3" json:"metric,omitempty"`
	Value                int32    `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}
func (*Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95933fe5266f40f, []int{4}
}

func (m *Stat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stat.Unmarshal(m, b)
}
func (m *Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stat.Marshal(b, m, deterministic)
}
func (m *Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stat.Merge(m, src)
}
func (m *Stat) XXX_Size() int {
	return xxx_messageInfo_Stat.Size(m)
}
func (m *Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Stat proto.InternalMessageInfo

func (m *Stat) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Stat) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *Stat) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *Stat) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type StatRequest struct {
	Partition            string   `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Stat                 []*Stat  `protobuf:"bytes,2,rep,name=stat,proto3" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatRequest) Reset()         { *m = StatRequest{} }
func (m *StatRequest) String() string { return proto.CompactTextString(m) }
func (*StatRequest) ProtoMessage()    {}
func (*StatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95933fe5266f40f, []int{5}
}

func (m *StatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatRequest.Unmarshal(m, b)
}
func (m *StatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatRequest.Marshal(b, m, deterministic)
}
func (m *StatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatRequest.Merge(m, src)
}
func (m *StatRequest) XXX_Size() int {
	return xxx_messageInfo_StatRequest.Size(m)
}
func (m *StatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatRequest proto.InternalMessageInfo

func (m *StatRequest) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func (m *StatRequest) GetStat() []*Stat {
	if m != nil {
		return m.Stat
	}
	return nil
}

type ScaleRequest struct {
	Partition            string   `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaleRequest) Reset()         { *m = ScaleRequest{} }
func (m *ScaleRequest) String() string { return proto.CompactTextString(m) }
func (*ScaleRequest) ProtoMessage()    {}
func (*ScaleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95933fe5266f40f, []int{6}
}

func (m *ScaleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleRequest.Unmarshal(m, b)
}
func (m *ScaleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleRequest.Marshal(b, m, deterministic)
}
func (m *ScaleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleRequest.Merge(m, src)
}
func (m *ScaleRequest) XXX_Size() int {
	return xxx_messageInfo_ScaleRequest.Size(m)
}
func (m *ScaleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleRequest proto.InternalMessageInfo

func (m *ScaleRequest) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func (m *ScaleRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type ScaleResponse struct {
	Rec                  int32    `protobuf:"varint,1,opt,name=rec,proto3" json:"rec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaleResponse) Reset()         { *m = ScaleResponse{} }
func (m *ScaleResponse) String() string { return proto.CompactTextString(m) }
func (*ScaleResponse) ProtoMessage()    {}
func (*ScaleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95933fe5266f40f, []int{7}
}

func (m *ScaleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleResponse.Unmarshal(m, b)
}
func (m *ScaleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleResponse.Marshal(b, m, deterministic)
}
func (m *ScaleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleResponse.Merge(m, src)
}
func (m *ScaleResponse) XXX_Size() int {
	return xxx_messageInfo_ScaleResponse.Size(m)
}
func (m *ScaleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleResponse proto.InternalMessageInfo

func (m *ScaleResponse) GetRec() int32 {
	if m != nil {
		return m.Rec
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.EventRequest_EventType", EventRequest_EventType_name, EventRequest_EventType_value)
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*Autoscaler)(nil), "proto.Autoscaler")
	proto.RegisterType((*Pod)(nil), "proto.Pod")
	proto.RegisterType((*EventRequest)(nil), "proto.EventRequest")
	proto.RegisterType((*Stat)(nil), "proto.Stat")
	proto.RegisterType((*StatRequest)(nil), "proto.StatRequest")
	proto.RegisterType((*ScaleRequest)(nil), "proto.ScaleRequest")
	proto.RegisterType((*ScaleResponse)(nil), "proto.ScaleResponse")
}

func init() { proto.RegisterFile("skplug.proto", fileDescriptor_a95933fe5266f40f) }

var fileDescriptor_a95933fe5266f40f = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0xdd, 0x6c, 0x36, 0x5b, 0x76, 0x12, 0x96, 0xc5, 0x54, 0x28, 0x54, 0x40, 0x97, 0x5c, 0x88,
	0x38, 0x14, 0x91, 0xf2, 0x01, 0x14, 0x1a, 0xa9, 0x87, 0x0a, 0xad, 0xdc, 0xe5, 0xbc, 0x72, 0x13,
	0x53, 0x05, 0x92, 0xd8, 0x24, 0x4e, 0xa5, 0x7c, 0x02, 0x57, 0xbe, 0x18, 0x8d, 0xed, 0xa6, 0x41,
	0xda, 0x03, 0xea, 0x29, 0x6f, 0x9e, 0xc7, 0xe3, 0x37, 0xef, 0x29, 0x10, 0xb4, 0x3f, 0x65, 0xd9,
	0xdd, 0x9c, 0xc8, 0x46, 0x28, 0x41, 0x3c, 0xfd, 0x89, 0x0e, 0xc0, 0x4b, 0x2b, 0xa9, 0xfa, 0xe8,
	0x23, 0xc0, 0x59, 0xa7, 0x44, 0x9b, 0xb1, 0x92, 0x37, 0x84, 0xc0, 0x4c, 0xf5, 0x92, 0x87, 0xce,
	0xda, 0x89, 0x17, 0x54, 0x63, 0xe4, 0x7a, 0x56, 0x95, 0xe1, 0xd4, 0x70, 0x88, 0xa3, 0x1e, 0xdc,
	0x8d, 0xc8, 0xf1, 0xa8, 0x66, 0xd5, 0xd0, 0x8e, 0x98, 0x1c, 0x82, 0xd7, 0x2a, 0xa6, 0xb8, 0xed,
	0x37, 0x05, 0x79, 0x0b, 0x4f, 0x4a, 0xd6, 0xaa, 0x9d, 0x6a, 0x58, 0xdd, 0x16, 0xaa, 0x10, 0x75,
	0xe8, 0xae, 0x9d, 0xd8, 0xa5, 0x4b, 0xa4, 0xb7, 0x03, 0x4b, 0x8e, 0xc1, 0xcf, 0x64, 0xb7, 0x6b,
	0xf8, 0xaf, 0x8e, 0xb7, 0x2a, 0x9c, 0xad, 0x9d, 0xd8, 0xa3, 0x90, 0xc9, 0x8e, 0x1a, 0x26, 0xfa,
	0x3d, 0x85, 0x20, 0xbd, 0xe5, 0xb5, 0xb2, 0x04, 0x79, 0x09, 0x0b, 0xc9, 0x1a, 0x65, 0x86, 0x1a,
	0x25, 0xf7, 0x84, 0xde, 0xa8, 0xa8, 0x8c, 0x1a, 0x97, 0x6a, 0x4c, 0x3e, 0xd8, 0x2d, 0x51, 0xc1,
	0x32, 0x79, 0x65, 0x9c, 0x39, 0x19, 0x0f, 0x35, 0xc5, 0xb6, 0x97, 0xdc, 0x9a, 0x70, 0x0a, 0xc0,
	0x06, 0x9b, 0xb4, 0x2a, 0x3f, 0x79, 0x6a, 0x2f, 0xde, 0xfb, 0x77, 0x31, 0xa1, 0xa3, 0x36, 0xf2,
	0x1a, 0x5c, 0x29, 0xf2, 0xd0, 0xd3, 0xdd, 0x60, 0xbb, 0x37, 0x22, 0xbf, 0x98, 0x50, 0x3c, 0x88,
	0xde, 0xc3, 0x62, 0x78, 0x87, 0x00, 0xcc, 0xbf, 0xd0, 0xf4, 0x6c, 0x9b, 0xae, 0x26, 0x88, 0xbf,
	0x6d, 0xce, 0x11, 0x3b, 0x88, 0xcf, 0xd3, 0xcb, 0x74, 0x9b, 0xae, 0xa6, 0x9f, 0x97, 0x10, 0x88,
	0xeb, 0x1f, 0x3c, 0x53, 0x3b, 0x51, 0x73, 0xf1, 0x3d, 0xca, 0x60, 0x76, 0xa5, 0x98, 0x1a, 0x96,
	0x74, 0x46, 0x4b, 0xbe, 0x80, 0x47, 0x52, 0xe4, 0x3b, 0x9d, 0x8f, 0x89, 0xe2, 0x40, 0x8a, 0xfc,
	0x2b, 0x46, 0xf4, 0x1c, 0xe6, 0x15, 0x57, 0x4d, 0x91, 0x69, 0x07, 0x16, 0xd4, 0x56, 0x18, 0xdd,
	0x2d, 0x2b, 0x3b, 0x6e, 0x5d, 0x37, 0x45, 0x74, 0x09, 0x3e, 0x3e, 0xf2, 0x7f, 0x76, 0x1f, 0xc3,
	0x0c, 0x03, 0x0f, 0xa7, 0x6b, 0x37, 0xf6, 0x13, 0xdf, 0xee, 0xac, 0xef, 0xeb, 0x83, 0xe8, 0x13,
	0x04, 0x57, 0xe8, 0xce, 0x83, 0xd3, 0x8b, 0xde, 0xc0, 0x63, 0x3b, 0xa1, 0x95, 0xa2, 0x6e, 0x39,
	0x59, 0x81, 0xdb, 0xf0, 0x4c, 0x5f, 0xf6, 0x28, 0xc2, 0xe4, 0x8f, 0x03, 0xf3, 0x4d, 0xd9, 0xdd,
	0x14, 0x35, 0x79, 0x07, 0x9e, 0xf6, 0x98, 0x3c, 0xdb, 0x13, 0xf3, 0x51, 0x70, 0x47, 0xe2, 0xbf,
	0x40, 0xe2, 0x3b, 0x3b, 0xc7, 0xb2, 0xf7, 0x76, 0x26, 0xe0, 0x69, 0x0d, 0xc3, 0xd4, 0xf1, 0x4e,
	0x47, 0x87, 0xff, 0x92, 0x46, 0xe6, 0xf5, 0x5c, 0x93, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0x52, 0xba, 0x01, 0x90, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PluginClient interface {
	Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*Empty, error)
	Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*Empty, error)
	Scale(ctx context.Context, in *ScaleRequest, opts ...grpc.CallOption) (*ScaleResponse, error)
}

type pluginClient struct {
	cc *grpc.ClientConn
}

func NewPluginClient(cc *grpc.ClientConn) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Event", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Stat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Scale(ctx context.Context, in *ScaleRequest, opts ...grpc.CallOption) (*ScaleResponse, error) {
	out := new(ScaleResponse)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Scale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
type PluginServer interface {
	Event(context.Context, *EventRequest) (*Empty, error)
	Stat(context.Context, *StatRequest) (*Empty, error)
	Scale(context.Context, *ScaleRequest) (*ScaleResponse, error)
}

// UnimplementedPluginServer can be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (*UnimplementedPluginServer) Event(ctx context.Context, req *EventRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Event not implemented")
}
func (*UnimplementedPluginServer) Stat(ctx context.Context, req *StatRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (*UnimplementedPluginServer) Scale(ctx context.Context, req *ScaleRequest) (*ScaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scale not implemented")
}

func RegisterPluginServer(s *grpc.Server, srv PluginServer) {
	s.RegisterService(&_Plugin_serviceDesc, srv)
}

func _Plugin_Event_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Event(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/Event",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Event(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Stat(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Scale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Scale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/Scale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Scale(ctx, req.(*ScaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Event",
			Handler:    _Plugin_Event_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _Plugin_Stat_Handler,
		},
		{
			MethodName: "Scale",
			Handler:    _Plugin_Scale_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "skplug.proto",
}
