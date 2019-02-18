// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job_scheduler.proto

package jobscheduler

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegisterReq struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94497fe360f3acc8, []int{0}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type RegisterRes struct {
	WorkerID             string   `protobuf:"bytes,1,opt,name=workerID,proto3" json:"workerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRes) Reset()         { *m = RegisterRes{} }
func (m *RegisterRes) String() string { return proto.CompactTextString(m) }
func (*RegisterRes) ProtoMessage()    {}
func (*RegisterRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_94497fe360f3acc8, []int{1}
}

func (m *RegisterRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRes.Unmarshal(m, b)
}
func (m *RegisterRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRes.Marshal(b, m, deterministic)
}
func (m *RegisterRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRes.Merge(m, src)
}
func (m *RegisterRes) XXX_Size() int {
	return xxx_messageInfo_RegisterRes.Size(m)
}
func (m *RegisterRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRes.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRes proto.InternalMessageInfo

func (m *RegisterRes) GetWorkerID() string {
	if m != nil {
		return m.WorkerID
	}
	return ""
}

type DeregisterReq struct {
	WorkerID             string   `protobuf:"bytes,1,opt,name=workerID,proto3" json:"workerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeregisterReq) Reset()         { *m = DeregisterReq{} }
func (m *DeregisterReq) String() string { return proto.CompactTextString(m) }
func (*DeregisterReq) ProtoMessage()    {}
func (*DeregisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94497fe360f3acc8, []int{2}
}

func (m *DeregisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeregisterReq.Unmarshal(m, b)
}
func (m *DeregisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeregisterReq.Marshal(b, m, deterministic)
}
func (m *DeregisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeregisterReq.Merge(m, src)
}
func (m *DeregisterReq) XXX_Size() int {
	return xxx_messageInfo_DeregisterReq.Size(m)
}
func (m *DeregisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeregisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeregisterReq proto.InternalMessageInfo

func (m *DeregisterReq) GetWorkerID() string {
	if m != nil {
		return m.WorkerID
	}
	return ""
}

type DeregisterRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeregisterRes) Reset()         { *m = DeregisterRes{} }
func (m *DeregisterRes) String() string { return proto.CompactTextString(m) }
func (*DeregisterRes) ProtoMessage()    {}
func (*DeregisterRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_94497fe360f3acc8, []int{3}
}

func (m *DeregisterRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeregisterRes.Unmarshal(m, b)
}
func (m *DeregisterRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeregisterRes.Marshal(b, m, deterministic)
}
func (m *DeregisterRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeregisterRes.Merge(m, src)
}
func (m *DeregisterRes) XXX_Size() int {
	return xxx_messageInfo_DeregisterRes.Size(m)
}
func (m *DeregisterRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeregisterRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeregisterRes proto.InternalMessageInfo

func (m *DeregisterRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type StartJobReq struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartJobReq) Reset()         { *m = StartJobReq{} }
func (m *StartJobReq) String() string { return proto.CompactTextString(m) }
func (*StartJobReq) ProtoMessage()    {}
func (*StartJobReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94497fe360f3acc8, []int{4}
}

func (m *StartJobReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartJobReq.Unmarshal(m, b)
}
func (m *StartJobReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartJobReq.Marshal(b, m, deterministic)
}
func (m *StartJobReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartJobReq.Merge(m, src)
}
func (m *StartJobReq) XXX_Size() int {
	return xxx_messageInfo_StartJobReq.Size(m)
}
func (m *StartJobReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StartJobReq.DiscardUnknown(m)
}

var xxx_messageInfo_StartJobReq proto.InternalMessageInfo

func (m *StartJobReq) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *StartJobReq) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type StartJobRes struct {
	JobID                string   `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartJobRes) Reset()         { *m = StartJobRes{} }
func (m *StartJobRes) String() string { return proto.CompactTextString(m) }
func (*StartJobRes) ProtoMessage()    {}
func (*StartJobRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_94497fe360f3acc8, []int{5}
}

func (m *StartJobRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartJobRes.Unmarshal(m, b)
}
func (m *StartJobRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartJobRes.Marshal(b, m, deterministic)
}
func (m *StartJobRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartJobRes.Merge(m, src)
}
func (m *StartJobRes) XXX_Size() int {
	return xxx_messageInfo_StartJobRes.Size(m)
}
func (m *StartJobRes) XXX_DiscardUnknown() {
	xxx_messageInfo_StartJobRes.DiscardUnknown(m)
}

var xxx_messageInfo_StartJobRes proto.InternalMessageInfo

func (m *StartJobRes) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

type StopJobReq struct {
	JobID                string   `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopJobReq) Reset()         { *m = StopJobReq{} }
func (m *StopJobReq) String() string { return proto.CompactTextString(m) }
func (*StopJobReq) ProtoMessage()    {}
func (*StopJobReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94497fe360f3acc8, []int{6}
}

func (m *StopJobReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopJobReq.Unmarshal(m, b)
}
func (m *StopJobReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopJobReq.Marshal(b, m, deterministic)
}
func (m *StopJobReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopJobReq.Merge(m, src)
}
func (m *StopJobReq) XXX_Size() int {
	return xxx_messageInfo_StopJobReq.Size(m)
}
func (m *StopJobReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StopJobReq.DiscardUnknown(m)
}

var xxx_messageInfo_StopJobReq proto.InternalMessageInfo

func (m *StopJobReq) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

type StopJobRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopJobRes) Reset()         { *m = StopJobRes{} }
func (m *StopJobRes) String() string { return proto.CompactTextString(m) }
func (*StopJobRes) ProtoMessage()    {}
func (*StopJobRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_94497fe360f3acc8, []int{7}
}

func (m *StopJobRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopJobRes.Unmarshal(m, b)
}
func (m *StopJobRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopJobRes.Marshal(b, m, deterministic)
}
func (m *StopJobRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopJobRes.Merge(m, src)
}
func (m *StopJobRes) XXX_Size() int {
	return xxx_messageInfo_StopJobRes.Size(m)
}
func (m *StopJobRes) XXX_DiscardUnknown() {
	xxx_messageInfo_StopJobRes.DiscardUnknown(m)
}

var xxx_messageInfo_StopJobRes proto.InternalMessageInfo

type QueryJobReq struct {
	JobID                string   `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryJobReq) Reset()         { *m = QueryJobReq{} }
func (m *QueryJobReq) String() string { return proto.CompactTextString(m) }
func (*QueryJobReq) ProtoMessage()    {}
func (*QueryJobReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94497fe360f3acc8, []int{8}
}

func (m *QueryJobReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryJobReq.Unmarshal(m, b)
}
func (m *QueryJobReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryJobReq.Marshal(b, m, deterministic)
}
func (m *QueryJobReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryJobReq.Merge(m, src)
}
func (m *QueryJobReq) XXX_Size() int {
	return xxx_messageInfo_QueryJobReq.Size(m)
}
func (m *QueryJobReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryJobReq.DiscardUnknown(m)
}

var xxx_messageInfo_QueryJobReq proto.InternalMessageInfo

func (m *QueryJobReq) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

type QueryJobRes struct {
	Done                 bool     `protobuf:"varint,1,opt,name=done,proto3" json:"done,omitempty"`
	Error                bool     `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	ErrorText            string   `protobuf:"bytes,3,opt,name=errorText,proto3" json:"errorText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryJobRes) Reset()         { *m = QueryJobRes{} }
func (m *QueryJobRes) String() string { return proto.CompactTextString(m) }
func (*QueryJobRes) ProtoMessage()    {}
func (*QueryJobRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_94497fe360f3acc8, []int{9}
}

func (m *QueryJobRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryJobRes.Unmarshal(m, b)
}
func (m *QueryJobRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryJobRes.Marshal(b, m, deterministic)
}
func (m *QueryJobRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryJobRes.Merge(m, src)
}
func (m *QueryJobRes) XXX_Size() int {
	return xxx_messageInfo_QueryJobRes.Size(m)
}
func (m *QueryJobRes) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryJobRes.DiscardUnknown(m)
}

var xxx_messageInfo_QueryJobRes proto.InternalMessageInfo

func (m *QueryJobRes) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *QueryJobRes) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *QueryJobRes) GetErrorText() string {
	if m != nil {
		return m.ErrorText
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "jobscheduler.RegisterReq")
	proto.RegisterType((*RegisterRes)(nil), "jobscheduler.RegisterRes")
	proto.RegisterType((*DeregisterReq)(nil), "jobscheduler.DeregisterReq")
	proto.RegisterType((*DeregisterRes)(nil), "jobscheduler.DeregisterRes")
	proto.RegisterType((*StartJobReq)(nil), "jobscheduler.StartJobReq")
	proto.RegisterType((*StartJobRes)(nil), "jobscheduler.StartJobRes")
	proto.RegisterType((*StopJobReq)(nil), "jobscheduler.StopJobReq")
	proto.RegisterType((*StopJobRes)(nil), "jobscheduler.StopJobRes")
	proto.RegisterType((*QueryJobReq)(nil), "jobscheduler.QueryJobReq")
	proto.RegisterType((*QueryJobRes)(nil), "jobscheduler.QueryJobRes")
}

func init() { proto.RegisterFile("job_scheduler.proto", fileDescriptor_94497fe360f3acc8) }

var fileDescriptor_94497fe360f3acc8 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5d, 0x4f, 0xfa, 0x30,
	0x18, 0xc5, 0xd9, 0xff, 0x05, 0xc6, 0x03, 0x1a, 0x53, 0x4d, 0x9c, 0xd3, 0x18, 0x53, 0x2f, 0x94,
	0x98, 0xec, 0x02, 0x2f, 0xb9, 0x30, 0x59, 0xb8, 0x50, 0x2e, 0x0c, 0x6e, 0x1a, 0x2f, 0xcd, 0x5e,
	0x1a, 0x5e, 0x04, 0x8a, 0x6d, 0x41, 0xfd, 0x40, 0x7e, 0x20, 0xbf, 0x91, 0x59, 0x69, 0x47, 0x47,
	0x1c, 0x77, 0x3d, 0x7d, 0xce, 0x73, 0x38, 0xec, 0x97, 0xc2, 0xfe, 0x98, 0xc6, 0x2f, 0x3c, 0x19,
	0x92, 0x74, 0x31, 0x21, 0xcc, 0x9b, 0x33, 0x2a, 0x28, 0x6a, 0x8e, 0x69, 0x9c, 0xdf, 0xe1, 0x0b,
	0x68, 0x04, 0x64, 0x30, 0xe2, 0x82, 0xb0, 0x80, 0xbc, 0x21, 0x07, 0x6a, 0x51, 0x9a, 0x32, 0xc2,
	0xb9, 0x63, 0x9d, 0x59, 0x97, 0xf5, 0x40, 0x4b, 0xdc, 0x32, 0x8d, 0x1c, 0xb9, 0x60, 0xbf, 0x53,
	0xf6, 0x4a, 0xd8, 0x5d, 0x57, 0x39, 0x73, 0x8d, 0xaf, 0x60, 0xa7, 0x4b, 0x98, 0x91, 0xba, 0xcd,
	0xdc, 0x2a, 0x9a, 0x79, 0x56, 0x81, 0x2f, 0x92, 0x44, 0x57, 0xb0, 0x03, 0x2d, 0x71, 0x07, 0x1a,
	0xa1, 0x88, 0x98, 0xe8, 0xd1, 0x58, 0x75, 0x4d, 0xe8, 0x74, 0x1a, 0xcd, 0x52, 0xdd, 0x55, 0x49,
	0x84, 0xe0, 0xdf, 0x3c, 0x12, 0x43, 0xe7, 0x8f, 0xbc, 0x96, 0x67, 0x7c, 0x6e, 0x2e, 0x73, 0x74,
	0x00, 0xff, 0xc7, 0x34, 0xce, 0xfb, 0xac, 0x04, 0xc6, 0x00, 0xa1, 0xa0, 0x73, 0xf5, 0x03, 0xbf,
	0x7b, 0x9a, 0x86, 0x87, 0x67, 0xb1, 0x0f, 0x0b, 0xc2, 0x3e, 0xb7, 0xae, 0x3c, 0x99, 0x26, 0x9e,
	0xd5, 0x4b, 0xe9, 0x8c, 0xa8, 0xbf, 0x27, 0xcf, 0xd9, 0x22, 0x61, 0x8c, 0x32, 0xd9, 0xd9, 0x0e,
	0x56, 0x02, 0x9d, 0x40, 0x5d, 0x1e, 0x1e, 0xc9, 0x87, 0x70, 0xfe, 0xca, 0xc8, 0xf5, 0x45, 0xfb,
	0xcb, 0x82, 0x7a, 0xa8, 0x49, 0xa2, 0x5b, 0xd8, 0xd5, 0x80, 0x9e, 0xe5, 0xc7, 0x45, 0x47, 0x9e,
	0x89, 0xda, 0x33, 0x38, 0xbb, 0xa5, 0x23, 0x8e, 0x2b, 0xe8, 0x1e, 0xf6, 0xd6, 0x48, 0x54, 0xd6,
	0x71, 0x71, 0xa1, 0xc0, 0xd7, 0xdd, 0x32, 0xe4, 0xb8, 0xd2, 0xfe, 0xb6, 0xa0, 0xaa, 0x62, 0x7c,
	0xb0, 0x35, 0x85, 0xcd, 0x7a, 0x06, 0x5a, 0xb7, 0x74, 0x94, 0xd5, 0xbb, 0x81, 0x9a, 0x02, 0x80,
	0x9c, 0x4d, 0x9f, 0x66, 0xe7, 0x96, 0x4d, 0xb2, 0x00, 0x1f, 0x6c, 0x8d, 0x63, 0xb3, 0x84, 0xc1,
	0xd2, 0x2d, 0x1d, 0x71, 0x5c, 0xf1, 0x3b, 0x70, 0x9a, 0xd0, 0xa9, 0x37, 0x60, 0xd1, 0x72, 0x24,
	0x22, 0x31, 0xa2, 0xb3, 0x68, 0x52, 0xf0, 0xfb, 0x87, 0x3d, 0x1a, 0xe7, 0x74, 0x42, 0xc2, 0x96,
	0x84, 0xf5, 0xb3, 0x07, 0xd8, 0xb7, 0xe2, 0xaa, 0x7c, 0x89, 0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xa1, 0xf6, 0xe3, 0x82, 0xa0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchedulerClient interface {
	RegisterWorker(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
	DeregisterWorker(ctx context.Context, in *DeregisterReq, opts ...grpc.CallOption) (*DeregisterRes, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) RegisterWorker(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	out := new(RegisterRes)
	err := c.cc.Invoke(ctx, "/jobscheduler.Scheduler/RegisterWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) DeregisterWorker(ctx context.Context, in *DeregisterReq, opts ...grpc.CallOption) (*DeregisterRes, error) {
	out := new(DeregisterRes)
	err := c.cc.Invoke(ctx, "/jobscheduler.Scheduler/DeregisterWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServer is the server API for Scheduler service.
type SchedulerServer interface {
	RegisterWorker(context.Context, *RegisterReq) (*RegisterRes, error)
	DeregisterWorker(context.Context, *DeregisterReq) (*DeregisterRes, error)
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_RegisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).RegisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobscheduler.Scheduler/RegisterWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).RegisterWorker(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_DeregisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).DeregisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobscheduler.Scheduler/DeregisterWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).DeregisterWorker(ctx, req.(*DeregisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jobscheduler.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterWorker",
			Handler:    _Scheduler_RegisterWorker_Handler,
		},
		{
			MethodName: "DeregisterWorker",
			Handler:    _Scheduler_DeregisterWorker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job_scheduler.proto",
}

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerClient interface {
	StartJob(ctx context.Context, in *StartJobReq, opts ...grpc.CallOption) (*StartJobRes, error)
	StopJob(ctx context.Context, in *StopJobReq, opts ...grpc.CallOption) (*StopJobRes, error)
	QueryJob(ctx context.Context, in *QueryJobReq, opts ...grpc.CallOption) (*QueryJobRes, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) StartJob(ctx context.Context, in *StartJobReq, opts ...grpc.CallOption) (*StartJobRes, error) {
	out := new(StartJobRes)
	err := c.cc.Invoke(ctx, "/jobscheduler.Worker/StartJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) StopJob(ctx context.Context, in *StopJobReq, opts ...grpc.CallOption) (*StopJobRes, error) {
	out := new(StopJobRes)
	err := c.cc.Invoke(ctx, "/jobscheduler.Worker/StopJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) QueryJob(ctx context.Context, in *QueryJobReq, opts ...grpc.CallOption) (*QueryJobRes, error) {
	out := new(QueryJobRes)
	err := c.cc.Invoke(ctx, "/jobscheduler.Worker/QueryJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
type WorkerServer interface {
	StartJob(context.Context, *StartJobReq) (*StartJobRes, error)
	StopJob(context.Context, *StopJobReq) (*StopJobRes, error)
	QueryJob(context.Context, *QueryJobReq) (*QueryJobRes, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobscheduler.Worker/StartJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).StartJob(ctx, req.(*StartJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_StopJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).StopJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobscheduler.Worker/StopJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).StopJob(ctx, req.(*StopJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_QueryJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).QueryJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobscheduler.Worker/QueryJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).QueryJob(ctx, req.(*QueryJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jobscheduler.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _Worker_StartJob_Handler,
		},
		{
			MethodName: "StopJob",
			Handler:    _Worker_StopJob_Handler,
		},
		{
			MethodName: "QueryJob",
			Handler:    _Worker_QueryJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job_scheduler.proto",
}
