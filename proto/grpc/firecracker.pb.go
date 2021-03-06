// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: firecracker.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// CreateVMRequest specifies creation parameters for a new FC instance
type CreateVMRequest struct {
	// VM identifier to assign
	VMID string `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	// Specifies the machine configuration for the VM
	MachineCfg *FirecrackerMachineConfiguration `protobuf:"bytes,2,opt,name=MachineCfg,json=machineCfg,proto3" json:"MachineCfg,omitempty"`
	// Specifies the file path where the kernel image is located
	KernelImagePath string `protobuf:"bytes,3,opt,name=KernelImagePath,json=kernelImagePath,proto3" json:"KernelImagePath,omitempty"`
	// Specifies the commandline arguments that should be passed to the kernel
	KernelArgs string `protobuf:"bytes,4,opt,name=KernelArgs,json=kernelArgs,proto3" json:"KernelArgs,omitempty"`
	// Specifies the root block device for the VM
	RootDrive *FirecrackerDrive `protobuf:"bytes,5,opt,name=RootDrive,json=rootDrive,proto3" json:"RootDrive,omitempty"`
	// Specifies the additional block device config for the VM.
	AdditionalDrives []*FirecrackerDrive `protobuf:"bytes,6,rep,name=AdditionalDrives,json=additionalDrives,proto3" json:"AdditionalDrives,omitempty"`
	// Specifies the networking configuration for a VM
	NetworkInterfaces []*FirecrackerNetworkInterface `protobuf:"bytes,7,rep,name=NetworkInterfaces,json=networkInterfaces,proto3" json:"NetworkInterfaces,omitempty"`
	// The number of dummy drives to reserve in advance before running FC instance.
	ContainerCount       int32    `protobuf:"varint,8,opt,name=ContainerCount,json=containerCount,proto3" json:"ContainerCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVMRequest) Reset()         { *m = CreateVMRequest{} }
func (m *CreateVMRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVMRequest) ProtoMessage()    {}
func (*CreateVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{0}
}
func (m *CreateVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMRequest.Unmarshal(m, b)
}
func (m *CreateVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMRequest.Marshal(b, m, deterministic)
}
func (m *CreateVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMRequest.Merge(m, src)
}
func (m *CreateVMRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVMRequest.Size(m)
}
func (m *CreateVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMRequest proto.InternalMessageInfo

func (m *CreateVMRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func (m *CreateVMRequest) GetMachineCfg() *FirecrackerMachineConfiguration {
	if m != nil {
		return m.MachineCfg
	}
	return nil
}

func (m *CreateVMRequest) GetKernelImagePath() string {
	if m != nil {
		return m.KernelImagePath
	}
	return ""
}

func (m *CreateVMRequest) GetKernelArgs() string {
	if m != nil {
		return m.KernelArgs
	}
	return ""
}

func (m *CreateVMRequest) GetRootDrive() *FirecrackerDrive {
	if m != nil {
		return m.RootDrive
	}
	return nil
}

func (m *CreateVMRequest) GetAdditionalDrives() []*FirecrackerDrive {
	if m != nil {
		return m.AdditionalDrives
	}
	return nil
}

func (m *CreateVMRequest) GetNetworkInterfaces() []*FirecrackerNetworkInterface {
	if m != nil {
		return m.NetworkInterfaces
	}
	return nil
}

func (m *CreateVMRequest) GetContainerCount() int32 {
	if m != nil {
		return m.ContainerCount
	}
	return 0
}

type StopVMRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopVMRequest) Reset()         { *m = StopVMRequest{} }
func (m *StopVMRequest) String() string { return proto.CompactTextString(m) }
func (*StopVMRequest) ProtoMessage()    {}
func (*StopVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{1}
}
func (m *StopVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopVMRequest.Unmarshal(m, b)
}
func (m *StopVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopVMRequest.Marshal(b, m, deterministic)
}
func (m *StopVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopVMRequest.Merge(m, src)
}
func (m *StopVMRequest) XXX_Size() int {
	return xxx_messageInfo_StopVMRequest.Size(m)
}
func (m *StopVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopVMRequest proto.InternalMessageInfo

func (m *StopVMRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

type GetVMInfoRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMInfoRequest) Reset()         { *m = GetVMInfoRequest{} }
func (m *GetVMInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetVMInfoRequest) ProtoMessage()    {}
func (*GetVMInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{2}
}
func (m *GetVMInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMInfoRequest.Unmarshal(m, b)
}
func (m *GetVMInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetVMInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMInfoRequest.Merge(m, src)
}
func (m *GetVMInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetVMInfoRequest.Size(m)
}
func (m *GetVMInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMInfoRequest proto.InternalMessageInfo

func (m *GetVMInfoRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

type GetVMInfoResponse struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	ContextID            uint32   `protobuf:"varint,2,opt,name=ContextID,json=contextID,proto3" json:"ContextID,omitempty"`
	SocketPath           string   `protobuf:"bytes,3,opt,name=SocketPath,json=socketPath,proto3" json:"SocketPath,omitempty"`
	LogFifoPath          string   `protobuf:"bytes,4,opt,name=LogFifoPath,json=logFifoPath,proto3" json:"LogFifoPath,omitempty"`
	MetricsFifoPath      string   `protobuf:"bytes,5,opt,name=MetricsFifoPath,json=metricsFifoPath,proto3" json:"MetricsFifoPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMInfoResponse) Reset()         { *m = GetVMInfoResponse{} }
func (m *GetVMInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetVMInfoResponse) ProtoMessage()    {}
func (*GetVMInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{3}
}
func (m *GetVMInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMInfoResponse.Unmarshal(m, b)
}
func (m *GetVMInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetVMInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMInfoResponse.Merge(m, src)
}
func (m *GetVMInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetVMInfoResponse.Size(m)
}
func (m *GetVMInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMInfoResponse proto.InternalMessageInfo

func (m *GetVMInfoResponse) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func (m *GetVMInfoResponse) GetContextID() uint32 {
	if m != nil {
		return m.ContextID
	}
	return 0
}

func (m *GetVMInfoResponse) GetSocketPath() string {
	if m != nil {
		return m.SocketPath
	}
	return ""
}

func (m *GetVMInfoResponse) GetLogFifoPath() string {
	if m != nil {
		return m.LogFifoPath
	}
	return ""
}

func (m *GetVMInfoResponse) GetMetricsFifoPath() string {
	if m != nil {
		return m.MetricsFifoPath
	}
	return ""
}

type SetVMMetadataRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	Metadata             string   `protobuf:"bytes,2,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetVMMetadataRequest) Reset()         { *m = SetVMMetadataRequest{} }
func (m *SetVMMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*SetVMMetadataRequest) ProtoMessage()    {}
func (*SetVMMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{4}
}
func (m *SetVMMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetVMMetadataRequest.Unmarshal(m, b)
}
func (m *SetVMMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetVMMetadataRequest.Marshal(b, m, deterministic)
}
func (m *SetVMMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetVMMetadataRequest.Merge(m, src)
}
func (m *SetVMMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_SetVMMetadataRequest.Size(m)
}
func (m *SetVMMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetVMMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetVMMetadataRequest proto.InternalMessageInfo

func (m *SetVMMetadataRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func (m *SetVMMetadataRequest) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateVMRequest)(nil), "CreateVMRequest")
	proto.RegisterType((*StopVMRequest)(nil), "StopVMRequest")
	proto.RegisterType((*GetVMInfoRequest)(nil), "GetVMInfoRequest")
	proto.RegisterType((*GetVMInfoResponse)(nil), "GetVMInfoResponse")
	proto.RegisterType((*SetVMMetadataRequest)(nil), "SetVMMetadataRequest")
}

func init() { proto.RegisterFile("firecracker.proto", fileDescriptor_a73317e9fb8da571) }

var fileDescriptor_a73317e9fb8da571 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xc9, 0xa5, 0xf1, 0x44, 0xb9, 0x78, 0x05, 0xc8, 0x0a, 0x15, 0xb2, 0x82, 0x54, 0xe5,
	0x69, 0x83, 0x02, 0xaf, 0x20, 0x4a, 0x42, 0x90, 0x01, 0x23, 0xe4, 0x48, 0x79, 0xe0, 0x6d, 0x71,
	0xc6, 0xae, 0x95, 0x64, 0x37, 0xac, 0x37, 0x85, 0x7e, 0x08, 0x1f, 0xc2, 0x67, 0xf1, 0x17, 0xc8,
	0xeb, 0x3a, 0xb5, 0xdd, 0x34, 0x4f, 0xd6, 0x9c, 0x39, 0x47, 0x73, 0xbc, 0x67, 0x06, 0xac, 0x30,
	0x96, 0x18, 0x48, 0x16, 0xac, 0x51, 0xd2, 0x9d, 0x14, 0x4a, 0x0c, 0x9e, 0x45, 0x42, 0x44, 0x1b,
	0x1c, 0xeb, 0xea, 0xc7, 0x3e, 0x1c, 0xe3, 0x76, 0xa7, 0x6e, 0x6e, 0x9b, 0x6d, 0x75, 0xb3, 0xc3,
	0x24, 0x2b, 0x86, 0x7f, 0x6a, 0xd0, 0x9b, 0x4a, 0x64, 0x0a, 0x97, 0x9e, 0x8f, 0x3f, 0xf7, 0x98,
	0x28, 0x42, 0xa0, 0xbe, 0xf4, 0xdc, 0x99, 0x6d, 0x38, 0xc6, 0xc8, 0xf4, 0xeb, 0xd7, 0x9e, 0x3b,
	0x23, 0xef, 0x00, 0x3c, 0x16, 0x5c, 0xc5, 0x1c, 0xa7, 0x61, 0x64, 0x3f, 0x72, 0x8c, 0x51, 0x7b,
	0xe2, 0xd0, 0xf9, 0xdd, 0xe4, 0xbc, 0x2b, 0x78, 0x18, 0x47, 0x7b, 0xc9, 0x54, 0x2c, 0xb8, 0x0f,
	0xdb, 0x83, 0x86, 0x8c, 0xa0, 0xf7, 0x19, 0x25, 0xc7, 0x8d, 0xbb, 0x65, 0x11, 0x7e, 0x63, 0xea,
	0xca, 0xae, 0xe9, 0x01, 0xbd, 0x75, 0x19, 0x26, 0xcf, 0x01, 0x32, 0xe6, 0xa5, 0x8c, 0x12, 0xbb,
	0xae, 0x49, 0xb0, 0x3e, 0x20, 0x64, 0x0c, 0xa6, 0x2f, 0x84, 0x9a, 0xc9, 0xf8, 0x1a, 0xed, 0x86,
	0xb6, 0x62, 0x15, 0xad, 0xe8, 0x86, 0x6f, 0xca, 0x9c, 0x43, 0xde, 0x40, 0xff, 0x72, 0xb5, 0x8a,
	0x53, 0x4b, 0x6c, 0xa3, 0xa1, 0xc4, 0x6e, 0x3a, 0xb5, 0xe3, 0xba, 0x3e, 0xab, 0x50, 0xc9, 0x27,
	0xb0, 0xbe, 0xa2, 0xfa, 0x25, 0xe4, 0xda, 0xe5, 0x0a, 0x65, 0xc8, 0x02, 0x4c, 0xec, 0x33, 0xad,
	0x3f, 0x2f, 0xea, 0xab, 0x24, 0xdf, 0xe2, 0x55, 0x19, 0xb9, 0x80, 0xee, 0x54, 0x70, 0xc5, 0x62,
	0x8e, 0x72, 0x2a, 0xf6, 0x5c, 0xd9, 0x2d, 0xc7, 0x18, 0x35, 0xfc, 0x6e, 0x50, 0x42, 0x87, 0x2f,
	0xa0, 0xb3, 0x50, 0x62, 0x77, 0x32, 0x94, 0xe1, 0x05, 0xf4, 0x3f, 0xa2, 0x5a, 0x7a, 0x2e, 0x0f,
	0xc5, 0x29, 0xde, 0x5f, 0x03, 0xac, 0x02, 0x31, 0xd9, 0x09, 0x9e, 0xe0, 0xd1, 0x98, 0xcf, 0xc1,
	0x4c, 0xed, 0xe1, 0x6f, 0xe5, 0xce, 0x74, 0xca, 0x1d, 0xdf, 0x0c, 0x72, 0x20, 0x0d, 0x66, 0x21,
	0x82, 0x35, 0xaa, 0x42, 0x7a, 0x90, 0x1c, 0x10, 0xe2, 0x40, 0xfb, 0x8b, 0x88, 0xe6, 0x71, 0x28,
	0x34, 0x21, 0x4b, 0xae, 0xbd, 0xb9, 0x83, 0xd2, 0x25, 0xf0, 0x50, 0xc9, 0x38, 0x48, 0x0e, 0xac,
	0x46, 0xb6, 0x04, 0xdb, 0x32, 0x3c, 0x9c, 0xc3, 0xe3, 0x45, 0x6a, 0xd9, 0x43, 0xc5, 0x56, 0x4c,
	0xb1, 0x53, 0xcb, 0x39, 0x80, 0x56, 0x4e, 0xd3, 0xa6, 0x4d, 0xbf, 0xb5, 0xbd, 0xad, 0x27, 0xff,
	0x0c, 0x68, 0x17, 0x32, 0x22, 0xaf, 0xa1, 0x95, 0xef, 0x3b, 0xe9, 0xd3, 0xca, 0xea, 0x0f, 0x9e,
	0xd2, 0xec, 0x72, 0x68, 0x7e, 0x39, 0xf4, 0x43, 0x7a, 0x39, 0xe4, 0x25, 0x34, 0xb3, 0x38, 0x48,
	0x97, 0x96, 0x72, 0x79, 0x50, 0x31, 0x01, 0xf3, 0xf0, 0xe4, 0xc4, 0xa2, 0xd5, 0x9c, 0x06, 0x84,
	0xde, 0x4f, 0xe4, 0x2d, 0x74, 0x4a, 0xff, 0x4c, 0x9e, 0xd0, 0x63, 0x6f, 0xf0, 0xd0, 0xcc, 0xf7,
	0x67, 0xdf, 0x1b, 0x19, 0xd2, 0xd4, 0x9f, 0x57, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x29, 0xa5,
	0xe3, 0xff, 0x1b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FirecrackerClient is the client API for Firecracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FirecrackerClient interface {
	// Runs new Firecracker VM instance
	CreateVM(ctx context.Context, in *CreateVMRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Stops existing Firecracker instance by VM ID
	StopVM(ctx context.Context, in *StopVMRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Returns VM info by VM ID
	GetVMInfo(ctx context.Context, in *GetVMInfoRequest, opts ...grpc.CallOption) (*GetVMInfoResponse, error)
	// Sets VM's instance metadata
	SetVMMetadata(ctx context.Context, in *SetVMMetadataRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type firecrackerClient struct {
	cc *grpc.ClientConn
}

func NewFirecrackerClient(cc *grpc.ClientConn) FirecrackerClient {
	return &firecrackerClient{cc}
}

func (c *firecrackerClient) CreateVM(ctx context.Context, in *CreateVMRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Firecracker/CreateVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firecrackerClient) StopVM(ctx context.Context, in *StopVMRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Firecracker/StopVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firecrackerClient) GetVMInfo(ctx context.Context, in *GetVMInfoRequest, opts ...grpc.CallOption) (*GetVMInfoResponse, error) {
	out := new(GetVMInfoResponse)
	err := c.cc.Invoke(ctx, "/Firecracker/GetVMInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firecrackerClient) SetVMMetadata(ctx context.Context, in *SetVMMetadataRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Firecracker/SetVMMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirecrackerServer is the server API for Firecracker service.
type FirecrackerServer interface {
	// Runs new Firecracker VM instance
	CreateVM(context.Context, *CreateVMRequest) (*empty.Empty, error)
	// Stops existing Firecracker instance by VM ID
	StopVM(context.Context, *StopVMRequest) (*empty.Empty, error)
	// Returns VM info by VM ID
	GetVMInfo(context.Context, *GetVMInfoRequest) (*GetVMInfoResponse, error)
	// Sets VM's instance metadata
	SetVMMetadata(context.Context, *SetVMMetadataRequest) (*empty.Empty, error)
}

func RegisterFirecrackerServer(s *grpc.Server, srv FirecrackerServer) {
	s.RegisterService(&_Firecracker_serviceDesc, srv)
}

func _Firecracker_CreateVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).CreateVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/CreateVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).CreateVM(ctx, req.(*CreateVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firecracker_StopVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).StopVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/StopVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).StopVM(ctx, req.(*StopVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firecracker_GetVMInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVMInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).GetVMInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/GetVMInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).GetVMInfo(ctx, req.(*GetVMInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firecracker_SetVMMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVMMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).SetVMMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/SetVMMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).SetVMMetadata(ctx, req.(*SetVMMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Firecracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Firecracker",
	HandlerType: (*FirecrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVM",
			Handler:    _Firecracker_CreateVM_Handler,
		},
		{
			MethodName: "StopVM",
			Handler:    _Firecracker_StopVM_Handler,
		},
		{
			MethodName: "GetVMInfo",
			Handler:    _Firecracker_GetVMInfo_Handler,
		},
		{
			MethodName: "SetVMMetadata",
			Handler:    _Firecracker_SetVMMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "firecracker.proto",
}
