// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serviceregistry.proto

package serviceregistry

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

type EventType int32

const (
	EventType_Register EventType = 0
	EventType_Remove   EventType = 1
)

var EventType_name = map[int32]string{
	0: "Register",
	1: "Remove",
}

var EventType_value = map[string]int32{
	"Register": 0,
	"Remove":   1,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d29ae0370b1d1881, []int{0}
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
	return fileDescriptor_d29ae0370b1d1881, []int{0}
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

type WorkloadEvent struct {
	EventType            EventType        `protobuf:"varint,1,opt,name=eventType,proto3,enum=svreg.EventType" json:"eventType,omitempty"`
	ServiceWorkload      *ServiceWorkload `protobuf:"bytes,2,opt,name=serviceWorkload,proto3" json:"serviceWorkload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *WorkloadEvent) Reset()         { *m = WorkloadEvent{} }
func (m *WorkloadEvent) String() string { return proto.CompactTextString(m) }
func (*WorkloadEvent) ProtoMessage()    {}
func (*WorkloadEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29ae0370b1d1881, []int{1}
}

func (m *WorkloadEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadEvent.Unmarshal(m, b)
}
func (m *WorkloadEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadEvent.Marshal(b, m, deterministic)
}
func (m *WorkloadEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadEvent.Merge(m, src)
}
func (m *WorkloadEvent) XXX_Size() int {
	return xxx_messageInfo_WorkloadEvent.Size(m)
}
func (m *WorkloadEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadEvent.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadEvent proto.InternalMessageInfo

func (m *WorkloadEvent) GetEventType() EventType {
	if m != nil {
		return m.EventType
	}
	return EventType_Register
}

func (m *WorkloadEvent) GetServiceWorkload() *ServiceWorkload {
	if m != nil {
		return m.ServiceWorkload
	}
	return nil
}

//The request used to register a new service or update the existing ones
type ServiceWorkload struct {
	ServiceName          string      `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	ConnectivityDomain   string      `protobuf:"bytes,2,opt,name=connectivityDomain,proto3" json:"connectivityDomain,omitempty"`
	Workloads            []*Workload `protobuf:"bytes,3,rep,name=workloads,proto3" json:"workloads,omitempty"`
	Ports                []int32     `protobuf:"varint,4,rep,packed,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ServiceWorkload) Reset()         { *m = ServiceWorkload{} }
func (m *ServiceWorkload) String() string { return proto.CompactTextString(m) }
func (*ServiceWorkload) ProtoMessage()    {}
func (*ServiceWorkload) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29ae0370b1d1881, []int{2}
}

func (m *ServiceWorkload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceWorkload.Unmarshal(m, b)
}
func (m *ServiceWorkload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceWorkload.Marshal(b, m, deterministic)
}
func (m *ServiceWorkload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceWorkload.Merge(m, src)
}
func (m *ServiceWorkload) XXX_Size() int {
	return xxx_messageInfo_ServiceWorkload.Size(m)
}
func (m *ServiceWorkload) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceWorkload.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceWorkload proto.InternalMessageInfo

func (m *ServiceWorkload) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ServiceWorkload) GetConnectivityDomain() string {
	if m != nil {
		return m.ConnectivityDomain
	}
	return ""
}

func (m *ServiceWorkload) GetWorkloads() []*Workload {
	if m != nil {
		return m.Workloads
	}
	return nil
}

func (m *ServiceWorkload) GetPorts() []int32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

//Basic workload
type Workload struct {
	Identifier           *WorkloadIdentifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	IPAddress            []string            `protobuf:"bytes,2,rep,name=IPAddress,json=iPAddress,proto3" json:"IPAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Workload) Reset()         { *m = Workload{} }
func (m *Workload) String() string { return proto.CompactTextString(m) }
func (*Workload) ProtoMessage()    {}
func (*Workload) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29ae0370b1d1881, []int{3}
}

func (m *Workload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workload.Unmarshal(m, b)
}
func (m *Workload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workload.Marshal(b, m, deterministic)
}
func (m *Workload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workload.Merge(m, src)
}
func (m *Workload) XXX_Size() int {
	return xxx_messageInfo_Workload.Size(m)
}
func (m *Workload) XXX_DiscardUnknown() {
	xxx_messageInfo_Workload.DiscardUnknown(m)
}

var xxx_messageInfo_Workload proto.InternalMessageInfo

func (m *Workload) GetIdentifier() *WorkloadIdentifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Workload) GetIPAddress() []string {
	if m != nil {
		return m.IPAddress
	}
	return nil
}

type WorkloadIdentifier struct {
	Cluster              string   `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	PodName              string   `protobuf:"bytes,2,opt,name=podName,proto3" json:"podName,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkloadIdentifier) Reset()         { *m = WorkloadIdentifier{} }
func (m *WorkloadIdentifier) String() string { return proto.CompactTextString(m) }
func (*WorkloadIdentifier) ProtoMessage()    {}
func (*WorkloadIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29ae0370b1d1881, []int{4}
}

func (m *WorkloadIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadIdentifier.Unmarshal(m, b)
}
func (m *WorkloadIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadIdentifier.Marshal(b, m, deterministic)
}
func (m *WorkloadIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadIdentifier.Merge(m, src)
}
func (m *WorkloadIdentifier) XXX_Size() int {
	return xxx_messageInfo_WorkloadIdentifier.Size(m)
}
func (m *WorkloadIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadIdentifier proto.InternalMessageInfo

func (m *WorkloadIdentifier) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *WorkloadIdentifier) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *WorkloadIdentifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RegisteredWorkloads struct {
	Workloads            []*ServiceWorkload `protobuf:"bytes,1,rep,name=Workloads,json=workloads,proto3" json:"Workloads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RegisteredWorkloads) Reset()         { *m = RegisteredWorkloads{} }
func (m *RegisteredWorkloads) String() string { return proto.CompactTextString(m) }
func (*RegisteredWorkloads) ProtoMessage()    {}
func (*RegisteredWorkloads) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29ae0370b1d1881, []int{5}
}

func (m *RegisteredWorkloads) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisteredWorkloads.Unmarshal(m, b)
}
func (m *RegisteredWorkloads) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisteredWorkloads.Marshal(b, m, deterministic)
}
func (m *RegisteredWorkloads) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredWorkloads.Merge(m, src)
}
func (m *RegisteredWorkloads) XXX_Size() int {
	return xxx_messageInfo_RegisteredWorkloads.Size(m)
}
func (m *RegisteredWorkloads) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredWorkloads.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredWorkloads proto.InternalMessageInfo

func (m *RegisteredWorkloads) GetWorkloads() []*ServiceWorkload {
	if m != nil {
		return m.Workloads
	}
	return nil
}

func init() {
	proto.RegisterEnum("svreg.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Empty)(nil), "svreg.Empty")
	proto.RegisterType((*WorkloadEvent)(nil), "svreg.WorkloadEvent")
	proto.RegisterType((*ServiceWorkload)(nil), "svreg.ServiceWorkload")
	proto.RegisterType((*Workload)(nil), "svreg.Workload")
	proto.RegisterType((*WorkloadIdentifier)(nil), "svreg.WorkloadIdentifier")
	proto.RegisterType((*RegisteredWorkloads)(nil), "svreg.RegisteredWorkloads")
}

func init() { proto.RegisterFile("serviceregistry.proto", fileDescriptor_d29ae0370b1d1881) }

var fileDescriptor_d29ae0370b1d1881 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0xbe, 0x34, 0xbd, 0xb6, 0x3b, 0x69, 0x7b, 0xe7, 0x58, 0x4b, 0x3c, 0x7c, 0x08, 0x01, 0x21,
	0x08, 0x06, 0x89, 0x0a, 0xea, 0x93, 0x4a, 0x8f, 0x52, 0x04, 0x91, 0xad, 0x52, 0x10, 0x5f, 0xce,
	0x64, 0x2c, 0x8b, 0xbd, 0x6c, 0xd8, 0x5d, 0x23, 0x79, 0xf4, 0xdf, 0xf8, 0x0f, 0xfc, 0x7b, 0x72,
	0x7b, 0xd9, 0x5c, 0x2f, 0xde, 0x3d, 0xf8, 0xb6, 0x33, 0xf3, 0xcd, 0x7c, 0xdf, 0x7e, 0x3b, 0x0b,
	0xf7, 0x34, 0xa9, 0x5a, 0xe4, 0xa4, 0xe8, 0x5a, 0x68, 0xa3, 0x9a, 0xb4, 0x52, 0xd2, 0x48, 0x1c,
	0xea, 0x5a, 0xd1, 0x75, 0xbc, 0x0f, 0xc3, 0xe9, 0xbc, 0x32, 0x4d, 0xfc, 0xcb, 0x83, 0xa3, 0x2b,
	0xa9, 0xbe, 0xdf, 0xc8, 0x59, 0x31, 0xad, 0xa9, 0x34, 0x98, 0x02, 0xa3, 0xc5, 0xe1, 0x63, 0x53,
	0x51, 0xe8, 0x45, 0x5e, 0x72, 0x9c, 0x8d, 0x53, 0xdb, 0x95, 0x4e, 0x5d, 0x9e, 0xaf, 0x20, 0xf8,
	0x1a, 0x46, 0x2d, 0x95, 0x9b, 0x13, 0xee, 0x44, 0x5e, 0x12, 0x64, 0xa7, 0x6d, 0xd7, 0xe5, 0x7a,
	0x95, 0xf7, 0xe1, 0xf1, 0x6f, 0x0f, 0x46, 0x3d, 0x10, 0x46, 0x10, 0xb4, 0xb0, 0xf7, 0xb3, 0xf9,
	0x52, 0x07, 0xe3, 0xb7, 0x53, 0x98, 0x02, 0xe6, 0xb2, 0x2c, 0x29, 0x37, 0xa2, 0x16, 0xa6, 0x39,
	0x93, 0xf3, 0x99, 0x28, 0x2d, 0x35, 0xe3, 0x1b, 0x2a, 0xf8, 0x18, 0xd8, 0xcf, 0x76, 0xba, 0x0e,
	0xfd, 0xc8, 0x4f, 0x82, 0x6c, 0xd4, 0x2a, 0xec, 0xa4, 0xad, 0x10, 0x78, 0x02, 0xc3, 0x4a, 0x2a,
	0xa3, 0xc3, 0xdd, 0xc8, 0x4f, 0x86, 0x7c, 0x19, 0xc4, 0x39, 0x1c, 0x74, 0x12, 0x5f, 0x02, 0x88,
	0x82, 0x4a, 0x23, 0xbe, 0x09, 0x52, 0x56, 0x61, 0x90, 0xdd, 0xef, 0x4d, 0xbc, 0xe8, 0x00, 0xfc,
	0x16, 0x18, 0x1f, 0x00, 0xbb, 0xf8, 0xf0, 0xa6, 0x28, 0x14, 0x69, 0x1d, 0xee, 0x44, 0x7e, 0xc2,
	0x38, 0x13, 0x2e, 0x11, 0x7f, 0x01, 0xfc, 0xb7, 0x1f, 0x43, 0xd8, 0xcf, 0x6f, 0x7e, 0x68, 0xd3,
	0x72, 0x31, 0xee, 0xc2, 0x45, 0xa5, 0x92, 0x85, 0xf5, 0x69, 0x79, 0x7d, 0x17, 0x22, 0xc2, 0x6e,
	0xb9, 0x48, 0xfb, 0x36, 0x6d, 0xcf, 0xf1, 0x3b, 0xb8, 0xcb, 0xed, 0x4e, 0x90, 0xa2, 0xe2, 0xaa,
	0xbb, 0xef, 0x33, 0x60, 0x5d, 0x10, 0x7a, 0xd6, 0x9e, 0x6d, 0x0f, 0xb8, 0x72, 0xe9, 0xd1, 0x43,
	0x60, 0xdd, 0x52, 0xe0, 0x21, 0x1c, 0xb8, 0xc9, 0xe3, 0x01, 0x02, 0xec, 0x71, 0x9a, 0xcb, 0x9a,
	0xc6, 0x5e, 0xf6, 0xc7, 0x73, 0x25, 0xd5, 0xe0, 0x2b, 0x18, 0x3b, 0x58, 0xe7, 0xe5, 0x16, 0xaa,
	0xc9, 0xa1, 0xdb, 0x3c, 0xbb, 0xac, 0x03, 0x7c, 0x0e, 0xc1, 0x39, 0x99, 0x16, 0xa5, 0x71, 0xad,
	0x3c, 0x39, 0xe9, 0x99, 0x6f, 0x95, 0xc5, 0x83, 0x27, 0x1e, 0xbe, 0x80, 0xe3, 0xa5, 0x96, 0xff,
	0x25, 0xcc, 0x3e, 0xc1, 0x91, 0x13, 0x7e, 0x69, 0x66, 0x86, 0xf0, 0x0c, 0x4e, 0xcf, 0xc9, 0x6c,
	0x72, 0x70, 0x5d, 0xcc, 0xa4, 0x8d, 0x36, 0x20, 0xe3, 0xc1, 0xdb, 0x3b, 0x9f, 0x47, 0xbd, 0xff,
	0xf9, 0x75, 0xcf, 0x7e, 0xd0, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x60, 0xa3, 0xfe, 0xd1,
	0xb9, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryClient interface {
	RegisterWorkload(ctx context.Context, in *ServiceWorkload, opts ...grpc.CallOption) (*Empty, error)
	GetServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Registry_GetServicesClient, error)
	RemoveWorkload(ctx context.Context, in *ServiceWorkload, opts ...grpc.CallOption) (*Empty, error)
}

type registryClient struct {
	cc *grpc.ClientConn
}

func NewRegistryClient(cc *grpc.ClientConn) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) RegisterWorkload(ctx context.Context, in *ServiceWorkload, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/svreg.Registry/RegisterWorkload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Registry_GetServicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Registry_serviceDesc.Streams[0], "/svreg.Registry/GetServices", opts...)
	if err != nil {
		return nil, err
	}
	x := &registryGetServicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Registry_GetServicesClient interface {
	Recv() (*WorkloadEvent, error)
	grpc.ClientStream
}

type registryGetServicesClient struct {
	grpc.ClientStream
}

func (x *registryGetServicesClient) Recv() (*WorkloadEvent, error) {
	m := new(WorkloadEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *registryClient) RemoveWorkload(ctx context.Context, in *ServiceWorkload, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/svreg.Registry/RemoveWorkload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
type RegistryServer interface {
	RegisterWorkload(context.Context, *ServiceWorkload) (*Empty, error)
	GetServices(*Empty, Registry_GetServicesServer) error
	RemoveWorkload(context.Context, *ServiceWorkload) (*Empty, error)
}

// UnimplementedRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (*UnimplementedRegistryServer) RegisterWorkload(ctx context.Context, req *ServiceWorkload) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWorkload not implemented")
}
func (*UnimplementedRegistryServer) GetServices(req *Empty, srv Registry_GetServicesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (*UnimplementedRegistryServer) RemoveWorkload(ctx context.Context, req *ServiceWorkload) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWorkload not implemented")
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_RegisterWorkload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceWorkload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).RegisterWorkload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svreg.Registry/RegisterWorkload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).RegisterWorkload(ctx, req.(*ServiceWorkload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetServices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RegistryServer).GetServices(m, &registryGetServicesServer{stream})
}

type Registry_GetServicesServer interface {
	Send(*WorkloadEvent) error
	grpc.ServerStream
}

type registryGetServicesServer struct {
	grpc.ServerStream
}

func (x *registryGetServicesServer) Send(m *WorkloadEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Registry_RemoveWorkload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceWorkload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).RemoveWorkload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svreg.Registry/RemoveWorkload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).RemoveWorkload(ctx, req.(*ServiceWorkload))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "svreg.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterWorkload",
			Handler:    _Registry_RegisterWorkload_Handler,
		},
		{
			MethodName: "RemoveWorkload",
			Handler:    _Registry_RemoveWorkload_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetServices",
			Handler:       _Registry_GetServices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "serviceregistry.proto",
}

// RegistryStateClient is the client API for RegistryState service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryStateClient interface {
	GetRegisteredWorkloads(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RegisteredWorkloads, error)
}

type registryStateClient struct {
	cc *grpc.ClientConn
}

func NewRegistryStateClient(cc *grpc.ClientConn) RegistryStateClient {
	return &registryStateClient{cc}
}

func (c *registryStateClient) GetRegisteredWorkloads(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RegisteredWorkloads, error) {
	out := new(RegisteredWorkloads)
	err := c.cc.Invoke(ctx, "/svreg.RegistryState/GetRegisteredWorkloads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryStateServer is the server API for RegistryState service.
type RegistryStateServer interface {
	GetRegisteredWorkloads(context.Context, *Empty) (*RegisteredWorkloads, error)
}

// UnimplementedRegistryStateServer can be embedded to have forward compatible implementations.
type UnimplementedRegistryStateServer struct {
}

func (*UnimplementedRegistryStateServer) GetRegisteredWorkloads(ctx context.Context, req *Empty) (*RegisteredWorkloads, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisteredWorkloads not implemented")
}

func RegisterRegistryStateServer(s *grpc.Server, srv RegistryStateServer) {
	s.RegisterService(&_RegistryState_serviceDesc, srv)
}

func _RegistryState_GetRegisteredWorkloads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryStateServer).GetRegisteredWorkloads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svreg.RegistryState/GetRegisteredWorkloads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryStateServer).GetRegisteredWorkloads(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistryState_serviceDesc = grpc.ServiceDesc{
	ServiceName: "svreg.RegistryState",
	HandlerType: (*RegistryStateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRegisteredWorkloads",
			Handler:    _RegistryState_GetRegisteredWorkloads_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serviceregistry.proto",
}
