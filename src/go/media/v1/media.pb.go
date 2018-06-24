// Code generated by protoc-gen-go. DO NOT EDIT.
// source: media/v1/media.proto

package media // import "media/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/googleapis/longrunning"
import _ "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TrackInfo_Type int32

const (
	TrackInfo_VIDEO TrackInfo_Type = 0
	TrackInfo_AUDIO TrackInfo_Type = 1
	TrackInfo_TEXT  TrackInfo_Type = 2
	TrackInfo_OTHER TrackInfo_Type = 3
)

var TrackInfo_Type_name = map[int32]string{
	0: "VIDEO",
	1: "AUDIO",
	2: "TEXT",
	3: "OTHER",
}
var TrackInfo_Type_value = map[string]int32{
	"VIDEO": 0,
	"AUDIO": 1,
	"TEXT":  2,
	"OTHER": 3,
}

func (x TrackInfo_Type) String() string {
	return proto.EnumName(TrackInfo_Type_name, int32(x))
}
func (TrackInfo_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_media_5bbf5c5b00772cb1, []int{1, 0}
}

type Info struct {
	// Resource id of the media. It must have the format of "media/*/info".
	// For example: "media/1111-2222-3333-4444/info".
	MediaHash            string       `protobuf:"bytes,1,opt,name=media_hash,json=mediaHash,proto3" json:"media_hash,omitempty"`
	Tracks               []*TrackInfo `protobuf:"bytes,2,rep,name=tracks,proto3" json:"tracks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_media_5bbf5c5b00772cb1, []int{0}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (dst *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(dst, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetMediaHash() string {
	if m != nil {
		return m.MediaHash
	}
	return ""
}

func (m *Info) GetTracks() []*TrackInfo {
	if m != nil {
		return m.Tracks
	}
	return nil
}

type TrackInfo struct {
	Types                TrackInfo_Type `protobuf:"varint,1,opt,name=types,proto3,enum=sagittarius.media.v1.TrackInfo_Type" json:"types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TrackInfo) Reset()         { *m = TrackInfo{} }
func (m *TrackInfo) String() string { return proto.CompactTextString(m) }
func (*TrackInfo) ProtoMessage()    {}
func (*TrackInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_media_5bbf5c5b00772cb1, []int{1}
}
func (m *TrackInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackInfo.Unmarshal(m, b)
}
func (m *TrackInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackInfo.Marshal(b, m, deterministic)
}
func (dst *TrackInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackInfo.Merge(dst, src)
}
func (m *TrackInfo) XXX_Size() int {
	return xxx_messageInfo_TrackInfo.Size(m)
}
func (m *TrackInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TrackInfo proto.InternalMessageInfo

func (m *TrackInfo) GetTypes() TrackInfo_Type {
	if m != nil {
		return m.Types
	}
	return TrackInfo_VIDEO
}

type GetInfoRequest struct {
	// quick hash of a media. For example: "media/1111-2222-3333-4444/info".
	MediaHash            string   `protobuf:"bytes,1,opt,name=media_hash,json=mediaHash,proto3" json:"media_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_media_5bbf5c5b00772cb1, []int{2}
}
func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoRequest.Unmarshal(m, b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(dst, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoRequest.Size(m)
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

func (m *GetInfoRequest) GetMediaHash() string {
	if m != nil {
		return m.MediaHash
	}
	return ""
}

type UpdateInfoRequest struct {
	// Resource name of the parent resource where to create the book.
	// For example: "shelves/shelf1".
	MediaHash string `protobuf:"bytes,1,opt,name=media_hash,json=mediaHash,proto3" json:"media_hash,omitempty"`
	// The Book resource to be created. Client must not set the `Book.name` field.
	Info                 *Info    `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateInfoRequest) Reset()         { *m = UpdateInfoRequest{} }
func (m *UpdateInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateInfoRequest) ProtoMessage()    {}
func (*UpdateInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_media_5bbf5c5b00772cb1, []int{3}
}
func (m *UpdateInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInfoRequest.Unmarshal(m, b)
}
func (m *UpdateInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInfoRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInfoRequest.Merge(dst, src)
}
func (m *UpdateInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateInfoRequest.Size(m)
}
func (m *UpdateInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInfoRequest proto.InternalMessageInfo

func (m *UpdateInfoRequest) GetMediaHash() string {
	if m != nil {
		return m.MediaHash
	}
	return ""
}

func (m *UpdateInfoRequest) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Info)(nil), "sagittarius.media.v1.Info")
	proto.RegisterType((*TrackInfo)(nil), "sagittarius.media.v1.TrackInfo")
	proto.RegisterType((*GetInfoRequest)(nil), "sagittarius.media.v1.GetInfoRequest")
	proto.RegisterType((*UpdateInfoRequest)(nil), "sagittarius.media.v1.UpdateInfoRequest")
	proto.RegisterEnum("sagittarius.media.v1.TrackInfo_Type", TrackInfo_Type_name, TrackInfo_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MediaClient is the client API for Media service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MediaClient interface {
	//
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*Info, error)
	//
	UpdateInfo(ctx context.Context, in *UpdateInfoRequest, opts ...grpc.CallOption) (*Info, error)
}

type mediaClient struct {
	cc *grpc.ClientConn
}

func NewMediaClient(cc *grpc.ClientConn) MediaClient {
	return &mediaClient{cc}
}

func (c *mediaClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/sagittarius.media.v1.Media/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) UpdateInfo(ctx context.Context, in *UpdateInfoRequest, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/sagittarius.media.v1.Media/UpdateInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaServer is the server API for Media service.
type MediaServer interface {
	//
	GetInfo(context.Context, *GetInfoRequest) (*Info, error)
	//
	UpdateInfo(context.Context, *UpdateInfoRequest) (*Info, error)
}

func RegisterMediaServer(s *grpc.Server, srv MediaServer) {
	s.RegisterService(&_Media_serviceDesc, srv)
}

func _Media_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sagittarius.media.v1.Media/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_UpdateInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).UpdateInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sagittarius.media.v1.Media/UpdateInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).UpdateInfo(ctx, req.(*UpdateInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Media_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sagittarius.media.v1.Media",
	HandlerType: (*MediaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Media_GetInfo_Handler,
		},
		{
			MethodName: "UpdateInfo",
			Handler:    _Media_UpdateInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "media/v1/media.proto",
}

func init() { proto.RegisterFile("media/v1/media.proto", fileDescriptor_media_5bbf5c5b00772cb1) }

var fileDescriptor_media_5bbf5c5b00772cb1 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xe9, 0xa6, 0x90, 0xa9, 0x14, 0x05, 0xab, 0x52, 0x43, 0xa0, 0x34, 0xda, 0x20, 0x11,
	0x2e, 0x6b, 0x25, 0x1c, 0x90, 0xc2, 0x89, 0xaa, 0x11, 0xcd, 0x01, 0xa5, 0x5a, 0x6d, 0x11, 0xe2,
	0x00, 0xf2, 0x36, 0xce, 0xc6, 0x6a, 0x63, 0x9b, 0xb5, 0xb7, 0x52, 0xf8, 0xb9, 0xf0, 0x0a, 0x3c,
	0x1a, 0xaf, 0xc0, 0x3b, 0xc0, 0x11, 0xd9, 0x6b, 0x5a, 0xb5, 0x2c, 0x29, 0xb7, 0xf1, 0x7c, 0xdf,
	0x7c, 0xdf, 0xcc, 0x68, 0x0c, 0xdb, 0x4b, 0x36, 0xe3, 0x94, 0x9c, 0x0f, 0x88, 0x0b, 0x22, 0x95,
	0x4b, 0x23, 0xf1, 0xb6, 0xa6, 0x19, 0x37, 0x86, 0xe6, 0xbc, 0xd0, 0x51, 0x09, 0x9c, 0x0f, 0x3a,
	0x0f, 0x32, 0x29, 0xb3, 0x33, 0x46, 0xa8, 0xe2, 0x84, 0x0a, 0x21, 0x0d, 0x35, 0x5c, 0x0a, 0x5d,
	0xd6, 0x74, 0x7a, 0x1e, 0x3d, 0x93, 0x22, 0xcb, 0x0b, 0x21, 0xb8, 0xc8, 0x88, 0x54, 0x2c, 0xbf,
	0x42, 0xba, 0xe7, 0x49, 0xee, 0x95, 0x16, 0x73, 0x42, 0xc5, 0xca, 0x43, 0x0f, 0xaf, 0x43, 0xb3,
	0xa2, 0xac, 0xf5, 0xf8, 0xde, 0x75, 0xdc, 0xf0, 0x25, 0xd3, 0x86, 0x2e, 0x95, 0x27, 0xec, 0x78,
	0x42, 0xae, 0x4e, 0x88, 0x36, 0xd4, 0x14, 0xde, 0x34, 0x7c, 0x07, 0xc1, 0x44, 0xcc, 0x25, 0xde,
	0x05, 0x70, 0xb3, 0xbc, 0x5f, 0x50, 0xbd, 0x68, 0xa3, 0x2e, 0xea, 0x37, 0xe2, 0x86, 0xcb, 0x1c,
	0x52, 0xbd, 0xc0, 0xcf, 0x60, 0xd3, 0xe4, 0xf4, 0xe4, 0x54, 0xb7, 0x6b, 0xdd, 0x8d, 0xfe, 0xd6,
	0x70, 0x2f, 0xaa, 0xda, 0x42, 0x94, 0x58, 0x8e, 0xd5, 0x8b, 0x3d, 0x3d, 0xfc, 0x08, 0x8d, 0x8b,
	0x24, 0x1e, 0x41, 0xdd, 0xac, 0x14, 0xd3, 0x4e, 0xbf, 0x39, 0x7c, 0x74, 0x83, 0x48, 0x94, 0xac,
	0x14, 0x8b, 0xcb, 0x92, 0x70, 0x00, 0x81, 0x7d, 0xe2, 0x06, 0xd4, 0x5f, 0x4f, 0x0e, 0xc6, 0xd3,
	0xd6, 0x2d, 0x1b, 0xbe, 0x38, 0x3e, 0x98, 0x4c, 0x5b, 0x08, 0xdf, 0x81, 0x20, 0x19, 0xbf, 0x49,
	0x5a, 0x35, 0x9b, 0x9c, 0x26, 0x87, 0xe3, 0xb8, 0xb5, 0x11, 0x12, 0x68, 0xbe, 0x64, 0xc6, 0xb5,
	0xc3, 0x3e, 0x14, 0x4c, 0x9b, 0x1b, 0xa6, 0x0c, 0x53, 0xb8, 0x7b, 0xac, 0x66, 0xd4, 0xb0, 0xff,
	0xaf, 0xc1, 0x11, 0x04, 0x5c, 0xcc, 0x65, 0xbb, 0xd6, 0x45, 0xfd, 0xad, 0x61, 0xa7, 0x7a, 0x24,
	0xa7, 0xe7, 0x78, 0xc3, 0x9f, 0x08, 0xea, 0xaf, 0x6c, 0x1e, 0x2b, 0xb8, 0xed, 0xdb, 0xc3, 0xff,
	0xd8, 0xc4, 0xd5, 0xee, 0x3b, 0x6b, 0xc4, 0xc3, 0xde, 0xd7, 0xef, 0x3f, 0xbe, 0xd5, 0x76, 0xf1,
	0xfd, 0x8b, 0x73, 0x25, 0x9f, 0x2e, 0xdb, 0xfe, 0x42, 0xac, 0x37, 0xfe, 0x0c, 0x70, 0x39, 0x1f,
	0x7e, 0x5c, 0x2d, 0xf7, 0xd7, 0x06, 0xd6, 0xfa, 0x3e, 0x71, 0xbe, 0xbd, 0x70, 0x9d, 0xef, 0x28,
	0x48, 0xa5, 0x3c, 0xdd, 0x1f, 0xc1, 0x0e, 0xe5, 0x95, 0x52, 0xfb, 0xe0, 0x36, 0x72, 0x64, 0x2f,
	0xf2, 0x08, 0xbd, 0x6d, 0xfe, 0xf9, 0x77, 0xcf, 0x5d, 0xf0, 0x0b, 0xa1, 0x74, 0xd3, 0x5d, 0xeb,
	0xd3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0xc3, 0xd9, 0x30, 0x93, 0x03, 0x00, 0x00,
}