// Code generated by protoc-gen-go.
// source: manager.proto
// DO NOT EDIT!

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MGExceptionMsgReq struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,3,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
}

func (m *MGExceptionMsgReq) Reset()                    { *m = MGExceptionMsgReq{} }
func (m *MGExceptionMsgReq) String() string            { return proto.CompactTextString(m) }
func (*MGExceptionMsgReq) ProtoMessage()               {}
func (*MGExceptionMsgReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *MGExceptionMsgReq) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *MGExceptionMsgReq) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *MGExceptionMsgReq) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *MGExceptionMsgReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type MGExceptionMsgRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *MGExceptionMsgRes) Reset()                    { *m = MGExceptionMsgRes{} }
func (m *MGExceptionMsgRes) String() string            { return proto.CompactTextString(m) }
func (*MGExceptionMsgRes) ProtoMessage()               {}
func (*MGExceptionMsgRes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *MGExceptionMsgRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *MGExceptionMsgRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

type MGOfflineMsgReq struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *MGOfflineMsgReq) Reset()                    { *m = MGOfflineMsgReq{} }
func (m *MGOfflineMsgReq) String() string            { return proto.CompactTextString(m) }
func (*MGOfflineMsgReq) ProtoMessage()               {}
func (*MGOfflineMsgReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *MGOfflineMsgReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type OfflineMsg struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,3,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
}

func (m *OfflineMsg) Reset()                    { *m = OfflineMsg{} }
func (m *OfflineMsg) String() string            { return proto.CompactTextString(m) }
func (*OfflineMsg) ProtoMessage()               {}
func (*OfflineMsg) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *OfflineMsg) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *OfflineMsg) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *OfflineMsg) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *OfflineMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type MGOfflineMsgRes struct {
	ErrCode uint32        `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string        `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
	Msgs    []*OfflineMsg `protobuf:"bytes,3,rep,name=msgs" json:"msgs,omitempty"`
}

func (m *MGOfflineMsgRes) Reset()                    { *m = MGOfflineMsgRes{} }
func (m *MGOfflineMsgRes) String() string            { return proto.CompactTextString(m) }
func (*MGOfflineMsgRes) ProtoMessage()               {}
func (*MGOfflineMsgRes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *MGOfflineMsgRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *MGOfflineMsgRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func (m *MGOfflineMsgRes) GetMsgs() []*OfflineMsg {
	if m != nil {
		return m.Msgs
	}
	return nil
}

// Sync
type MGSyncMsgReq struct {
	UID       int64 `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
	CurrentID int64 `protobuf:"varint,2,opt,name=currentID" json:"currentID,omitempty"`
	TotalID   int64 `protobuf:"varint,3,opt,name=totalID" json:"totalID,omitempty"`
}

func (m *MGSyncMsgReq) Reset()                    { *m = MGSyncMsgReq{} }
func (m *MGSyncMsgReq) String() string            { return proto.CompactTextString(m) }
func (*MGSyncMsgReq) ProtoMessage()               {}
func (*MGSyncMsgReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *MGSyncMsgReq) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *MGSyncMsgReq) GetCurrentID() int64 {
	if m != nil {
		return m.CurrentID
	}
	return 0
}

func (m *MGSyncMsgReq) GetTotalID() int64 {
	if m != nil {
		return m.TotalID
	}
	return 0
}

type MGSyncMsgRes struct {
	ErrCode uint32                   `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string                   `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
	Msgs    []*MGSyncMsgResOffsetMsg `protobuf:"bytes,3,rep,name=Msgs" json:"Msgs,omitempty"`
}

func (m *MGSyncMsgRes) Reset()                    { *m = MGSyncMsgRes{} }
func (m *MGSyncMsgRes) String() string            { return proto.CompactTextString(m) }
func (*MGSyncMsgRes) ProtoMessage()               {}
func (*MGSyncMsgRes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *MGSyncMsgRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *MGSyncMsgRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func (m *MGSyncMsgRes) GetMsgs() []*MGSyncMsgResOffsetMsg {
	if m != nil {
		return m.Msgs
	}
	return nil
}

type MGSyncMsgResOffsetMsg struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	GroupID   int64  `protobuf:"varint,3,opt,name=groupID" json:"groupID,omitempty"`
	MsgID     string `protobuf:"bytes,4,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,5,opt,name=msg" json:"msg,omitempty"`
}

func (m *MGSyncMsgResOffsetMsg) Reset()                    { *m = MGSyncMsgResOffsetMsg{} }
func (m *MGSyncMsgResOffsetMsg) String() string            { return proto.CompactTextString(m) }
func (*MGSyncMsgResOffsetMsg) ProtoMessage()               {}
func (*MGSyncMsgResOffsetMsg) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6, 0} }

func (m *MGSyncMsgResOffsetMsg) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *MGSyncMsgResOffsetMsg) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *MGSyncMsgResOffsetMsg) GetGroupID() int64 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *MGSyncMsgResOffsetMsg) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *MGSyncMsgResOffsetMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*MGExceptionMsgReq)(nil), "rpc.MGExceptionMsgReq")
	proto.RegisterType((*MGExceptionMsgRes)(nil), "rpc.MGExceptionMsgRes")
	proto.RegisterType((*MGOfflineMsgReq)(nil), "rpc.MGOfflineMsgReq")
	proto.RegisterType((*OfflineMsg)(nil), "rpc.offlineMsg")
	proto.RegisterType((*MGOfflineMsgRes)(nil), "rpc.MGOfflineMsgRes")
	proto.RegisterType((*MGSyncMsgReq)(nil), "rpc.MGSyncMsgReq")
	proto.RegisterType((*MGSyncMsgRes)(nil), "rpc.MGSyncMsgRes")
	proto.RegisterType((*MGSyncMsgResOffsetMsg)(nil), "rpc.MGSyncMsgRes.offsetMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ManagerServerRPC service

type ManagerServerRPCClient interface {
	SetExceptionMsg(ctx context.Context, in *MGExceptionMsgReq, opts ...grpc.CallOption) (*MGExceptionMsgRes, error)
	GetOfflineMsgs(ctx context.Context, in *MGOfflineMsgReq, opts ...grpc.CallOption) (*MGOfflineMsgRes, error)
	Sync(ctx context.Context, in *MGSyncMsgReq, opts ...grpc.CallOption) (*MGSyncMsgRes, error)
}

type managerServerRPCClient struct {
	cc *grpc.ClientConn
}

func NewManagerServerRPCClient(cc *grpc.ClientConn) ManagerServerRPCClient {
	return &managerServerRPCClient{cc}
}

func (c *managerServerRPCClient) SetExceptionMsg(ctx context.Context, in *MGExceptionMsgReq, opts ...grpc.CallOption) (*MGExceptionMsgRes, error) {
	out := new(MGExceptionMsgRes)
	err := grpc.Invoke(ctx, "/rpc.ManagerServerRPC/SetExceptionMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServerRPCClient) GetOfflineMsgs(ctx context.Context, in *MGOfflineMsgReq, opts ...grpc.CallOption) (*MGOfflineMsgRes, error) {
	out := new(MGOfflineMsgRes)
	err := grpc.Invoke(ctx, "/rpc.ManagerServerRPC/GetOfflineMsgs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerServerRPCClient) Sync(ctx context.Context, in *MGSyncMsgReq, opts ...grpc.CallOption) (*MGSyncMsgRes, error) {
	out := new(MGSyncMsgRes)
	err := grpc.Invoke(ctx, "/rpc.ManagerServerRPC/Sync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ManagerServerRPC service

type ManagerServerRPCServer interface {
	SetExceptionMsg(context.Context, *MGExceptionMsgReq) (*MGExceptionMsgRes, error)
	GetOfflineMsgs(context.Context, *MGOfflineMsgReq) (*MGOfflineMsgRes, error)
	Sync(context.Context, *MGSyncMsgReq) (*MGSyncMsgRes, error)
}

func RegisterManagerServerRPCServer(s *grpc.Server, srv ManagerServerRPCServer) {
	s.RegisterService(&_ManagerServerRPC_serviceDesc, srv)
}

func _ManagerServerRPC_SetExceptionMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGExceptionMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServerRPCServer).SetExceptionMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ManagerServerRPC/SetExceptionMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServerRPCServer).SetExceptionMsg(ctx, req.(*MGExceptionMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerServerRPC_GetOfflineMsgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGOfflineMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServerRPCServer).GetOfflineMsgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ManagerServerRPC/GetOfflineMsgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServerRPCServer).GetOfflineMsgs(ctx, req.(*MGOfflineMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerServerRPC_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGSyncMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServerRPCServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ManagerServerRPC/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServerRPCServer).Sync(ctx, req.(*MGSyncMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManagerServerRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ManagerServerRPC",
	HandlerType: (*ManagerServerRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetExceptionMsg",
			Handler:    _ManagerServerRPC_SetExceptionMsg_Handler,
		},
		{
			MethodName: "GetOfflineMsgs",
			Handler:    _ManagerServerRPC_GetOfflineMsgs_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _ManagerServerRPC_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}

func init() { proto.RegisterFile("manager.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x93, 0xcd, 0x8a, 0xdb, 0x30,
	0x10, 0xc7, 0xd7, 0x95, 0x77, 0x17, 0x4f, 0xbb, 0x4d, 0x56, 0x2c, 0x8b, 0x71, 0x7b, 0x08, 0xda,
	0x4b, 0x4e, 0x2e, 0x6c, 0xef, 0xbd, 0x64, 0x83, 0xc9, 0xc1, 0xb4, 0xc8, 0x14, 0x7a, 0x75, 0x1d,
	0xc5, 0x0d, 0xc4, 0x96, 0x33, 0x92, 0xfb, 0xf1, 0x08, 0x7d, 0xb2, 0x3e, 0x55, 0xa1, 0xc8, 0x51,
	0xac, 0x7c, 0x9d, 0x52, 0xd8, 0x9b, 0xe6, 0x3f, 0x23, 0xcd, 0xfc, 0x34, 0x33, 0x70, 0x53, 0xe5,
	0x75, 0x5e, 0x0a, 0x8c, 0x1b, 0x94, 0x5a, 0x52, 0x82, 0x4d, 0xc1, 0x7e, 0xc0, 0x6d, 0x9a, 0x4c,
	0x7f, 0x16, 0xa2, 0xd1, 0x4b, 0x59, 0xa7, 0xaa, 0xe4, 0x62, 0x4d, 0xdf, 0x42, 0xa0, 0x64, 0x8b,
	0x85, 0xf8, 0x3c, 0x7b, 0x0a, 0xbd, 0x91, 0x37, 0x26, 0xdc, 0x09, 0xc6, 0xab, 0x73, 0x2c, 0x85,
	0x36, 0xde, 0x17, 0x1b, 0x6f, 0x2f, 0xd0, 0x3b, 0xb8, 0xac, 0x54, 0x39, 0x7b, 0x0a, 0xc9, 0xc8,
	0x1b, 0x07, 0x7c, 0x63, 0xd0, 0x21, 0x90, 0x4a, 0x95, 0xa1, 0xdf, 0x69, 0xe6, 0xc8, 0xa6, 0xc7,
	0x89, 0x15, 0x0d, 0xe1, 0x5a, 0x20, 0x4e, 0xe4, 0x5c, 0x74, 0x69, 0x6f, 0xf8, 0xd6, 0xa4, 0xf7,
	0x70, 0x25, 0x10, 0x33, 0x8d, 0x5d, 0xc6, 0x80, 0x5b, 0x8b, 0x3d, 0xc0, 0x20, 0x4d, 0x3e, 0x2e,
	0x16, 0xab, 0x65, 0x2d, 0x6c, 0xf5, 0x43, 0x20, 0xed, 0x72, 0x6e, 0xeb, 0x36, 0x47, 0xd6, 0x00,
	0xc8, 0x3e, 0xe4, 0x59, 0xe8, 0xbe, 0x1d, 0x96, 0x75, 0x06, 0x1b, 0x7d, 0x00, 0xbf, 0x52, 0xa5,
	0x0a, 0xc9, 0x88, 0x8c, 0x5f, 0x3e, 0x0e, 0x62, 0x6c, 0x8a, 0xd8, 0x71, 0xf0, 0xce, 0xc9, 0xbe,
	0xc0, 0xab, 0x34, 0xc9, 0x7e, 0xd5, 0x85, 0xa3, 0x77, 0x5c, 0xc4, 0x12, 0x15, 0x2d, 0xa2, 0xa8,
	0xb5, 0x23, 0xea, 0x05, 0x53, 0x96, 0x96, 0x3a, 0x5f, 0x59, 0x26, 0xc2, 0xb7, 0x26, 0xfb, 0xeb,
	0xed, 0x3d, 0x7d, 0x0e, 0xc1, 0x3b, 0xf0, 0x53, 0x47, 0xf0, 0xa6, 0x23, 0xd8, 0x7d, 0xd2, 0xe0,
	0x28, 0xa1, 0x3b, 0x1a, 0x13, 0x18, 0xfd, 0xf6, 0x20, 0xe8, 0xb5, 0xff, 0xea, 0x54, 0x08, 0xd7,
	0x25, 0xca, 0xb6, 0x71, 0x5c, 0xd6, 0x74, 0x3d, 0xf4, 0x4f, 0xf4, 0xf0, 0xb2, 0xef, 0xe1, 0xe3,
	0x1f, 0x0f, 0x86, 0xe9, 0x66, 0x63, 0x32, 0x81, 0xdf, 0x05, 0xf2, 0x4f, 0x13, 0x3a, 0x81, 0x41,
	0x26, 0xf4, 0xee, 0xdc, 0xd2, 0x7b, 0x8b, 0x75, 0xb0, 0x45, 0xd1, 0x69, 0x5d, 0xb1, 0x0b, 0xfa,
	0x01, 0x5e, 0x27, 0x42, 0xbb, 0xf1, 0x50, 0xf4, 0xce, 0xc6, 0xee, 0x4d, 0x72, 0x74, 0x4a, 0x35,
	0xf7, 0x63, 0xf0, 0xcd, 0x1f, 0xd2, 0xdb, 0xc3, 0x0f, 0x5d, 0x47, 0x47, 0x92, 0x62, 0x17, 0x5f,
	0xaf, 0xba, 0x85, 0x7f, 0xff, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x72, 0xca, 0x96, 0x01, 0x04,
	0x00, 0x00,
}
