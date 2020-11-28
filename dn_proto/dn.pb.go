// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.5.1
// source: dn.proto

package dn_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CodeRequest) Reset() {
	*x = CodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeRequest) ProtoMessage() {}

func (x *CodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeRequest.ProtoReflect.Descriptor instead.
func (*CodeRequest) Descriptor() ([]byte, []int) {
	return file_dn_proto_rawDescGZIP(), []int{0}
}

func (x *CodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk    []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Tipo     string `protobuf:"bytes,2,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Parte    string `protobuf:"bytes,3,opt,name=parte,proto3" json:"parte,omitempty"`
	Cantidad string `protobuf:"bytes,4,opt,name=cantidad,proto3" json:"cantidad,omitempty"`
	Machine  string `protobuf:"bytes,5,opt,name=machine,proto3" json:"machine,omitempty"`
	Nombrel  string `protobuf:"bytes,6,opt,name=nombrel,proto3" json:"nombrel,omitempty"`
}

func (x *ChunkRequest) Reset() {
	*x = ChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkRequest) ProtoMessage() {}

func (x *ChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkRequest.ProtoReflect.Descriptor instead.
func (*ChunkRequest) Descriptor() ([]byte, []int) {
	return file_dn_proto_rawDescGZIP(), []int{1}
}

func (x *ChunkRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *ChunkRequest) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *ChunkRequest) GetParte() string {
	if x != nil {
		return x.Parte
	}
	return ""
}

func (x *ChunkRequest) GetCantidad() string {
	if x != nil {
		return x.Cantidad
	}
	return ""
}

func (x *ChunkRequest) GetMachine() string {
	if x != nil {
		return x.Machine
	}
	return ""
}

func (x *ChunkRequest) GetNombrel() string {
	if x != nil {
		return x.Nombrel
	}
	return ""
}

type PropRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cantidadn1    string `protobuf:"bytes,1,opt,name=cantidadn1,proto3" json:"cantidadn1,omitempty"`
	Cantidadn2    string `protobuf:"bytes,2,opt,name=cantidadn2,proto3" json:"cantidadn2,omitempty"`
	Cantidadn3    string `protobuf:"bytes,3,opt,name=cantidadn3,proto3" json:"cantidadn3,omitempty"`
	Nombrel       string `protobuf:"bytes,4,opt,name=nombrel,proto3" json:"nombrel,omitempty"`
	Cantidadtotal string `protobuf:"bytes,5,opt,name=cantidadtotal,proto3" json:"cantidadtotal,omitempty"`
}

func (x *PropRequest) Reset() {
	*x = PropRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropRequest) ProtoMessage() {}

func (x *PropRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropRequest.ProtoReflect.Descriptor instead.
func (*PropRequest) Descriptor() ([]byte, []int) {
	return file_dn_proto_rawDescGZIP(), []int{2}
}

func (x *PropRequest) GetCantidadn1() string {
	if x != nil {
		return x.Cantidadn1
	}
	return ""
}

func (x *PropRequest) GetCantidadn2() string {
	if x != nil {
		return x.Cantidadn2
	}
	return ""
}

func (x *PropRequest) GetCantidadn3() string {
	if x != nil {
		return x.Cantidadn3
	}
	return ""
}

func (x *PropRequest) GetNombrel() string {
	if x != nil {
		return x.Nombrel
	}
	return ""
}

func (x *PropRequest) GetCantidadtotal() string {
	if x != nil {
		return x.Cantidadtotal
	}
	return ""
}

var File_dn_proto protoreflect.FileDescriptor

var file_dn_proto_rawDesc = []byte{
	0x0a, 0x08, 0x64, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x74,
	0x69, 0x64, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x74,
	0x69, 0x64, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x6c, 0x22, 0xad, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x74,
	0x69, 0x64, 0x61, 0x64, 0x6e, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x6e, 0x31, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x74,
	0x69, 0x64, 0x61, 0x64, 0x6e, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x6e, 0x32, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x74,
	0x69, 0x64, 0x61, 0x64, 0x6e, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x6e, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x6d, 0x62, 0x72,
	0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x74, 0x69,
	0x64, 0x61, 0x64, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x83, 0x02, 0x0a, 0x09, 0x44, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x64, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x64, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x44, 0x4e, 0x12, 0x16, 0x2e, 0x64, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x15,
	0x2e, 0x64, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x73, 0x44, 0x4e, 0x12, 0x15,
	0x2e, 0x64, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dn_proto_rawDescOnce sync.Once
	file_dn_proto_rawDescData = file_dn_proto_rawDesc
)

func file_dn_proto_rawDescGZIP() []byte {
	file_dn_proto_rawDescOnce.Do(func() {
		file_dn_proto_rawDescData = protoimpl.X.CompressGZIP(file_dn_proto_rawDescData)
	})
	return file_dn_proto_rawDescData
}

var file_dn_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dn_proto_goTypes = []interface{}{
	(*CodeRequest)(nil),  // 0: dn_proto.CodeRequest
	(*ChunkRequest)(nil), // 1: dn_proto.ChunkRequest
	(*PropRequest)(nil),  // 2: dn_proto.PropRequest
}
var file_dn_proto_depIdxs = []int32{
	1, // 0: dn_proto.DnService.EnviarChunks:input_type -> dn_proto.ChunkRequest
	1, // 1: dn_proto.DnService.ChunksDN:input_type -> dn_proto.ChunkRequest
	0, // 2: dn_proto.DnService.Estado:input_type -> dn_proto.CodeRequest
	2, // 3: dn_proto.DnService.PropuestasDN:input_type -> dn_proto.PropRequest
	0, // 4: dn_proto.DnService.EnviarChunks:output_type -> dn_proto.CodeRequest
	0, // 5: dn_proto.DnService.ChunksDN:output_type -> dn_proto.CodeRequest
	0, // 6: dn_proto.DnService.Estado:output_type -> dn_proto.CodeRequest
	0, // 7: dn_proto.DnService.PropuestasDN:output_type -> dn_proto.CodeRequest
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dn_proto_init() }
func file_dn_proto_init() {
	if File_dn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dn_proto_goTypes,
		DependencyIndexes: file_dn_proto_depIdxs,
		MessageInfos:      file_dn_proto_msgTypes,
	}.Build()
	File_dn_proto = out.File
	file_dn_proto_rawDesc = nil
	file_dn_proto_goTypes = nil
	file_dn_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DnServiceClient is the client API for DnService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DnServiceClient interface {
	EnviarChunks(ctx context.Context, in *ChunkRequest, opts ...grpc.CallOption) (*CodeRequest, error)
	ChunksDN(ctx context.Context, in *ChunkRequest, opts ...grpc.CallOption) (*CodeRequest, error)
	Estado(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeRequest, error)
	PropuestasDN(ctx context.Context, in *PropRequest, opts ...grpc.CallOption) (*CodeRequest, error)
}

type dnServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDnServiceClient(cc grpc.ClientConnInterface) DnServiceClient {
	return &dnServiceClient{cc}
}

func (c *dnServiceClient) EnviarChunks(ctx context.Context, in *ChunkRequest, opts ...grpc.CallOption) (*CodeRequest, error) {
	out := new(CodeRequest)
	err := c.cc.Invoke(ctx, "/dn_proto.DnService/EnviarChunks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnServiceClient) ChunksDN(ctx context.Context, in *ChunkRequest, opts ...grpc.CallOption) (*CodeRequest, error) {
	out := new(CodeRequest)
	err := c.cc.Invoke(ctx, "/dn_proto.DnService/ChunksDN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnServiceClient) Estado(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeRequest, error) {
	out := new(CodeRequest)
	err := c.cc.Invoke(ctx, "/dn_proto.DnService/Estado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnServiceClient) PropuestasDN(ctx context.Context, in *PropRequest, opts ...grpc.CallOption) (*CodeRequest, error) {
	out := new(CodeRequest)
	err := c.cc.Invoke(ctx, "/dn_proto.DnService/PropuestasDN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnServiceServer is the server API for DnService service.
type DnServiceServer interface {
	EnviarChunks(context.Context, *ChunkRequest) (*CodeRequest, error)
	ChunksDN(context.Context, *ChunkRequest) (*CodeRequest, error)
	Estado(context.Context, *CodeRequest) (*CodeRequest, error)
	PropuestasDN(context.Context, *PropRequest) (*CodeRequest, error)
}

// UnimplementedDnServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDnServiceServer struct {
}

func (*UnimplementedDnServiceServer) EnviarChunks(context.Context, *ChunkRequest) (*CodeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarChunks not implemented")
}
func (*UnimplementedDnServiceServer) ChunksDN(context.Context, *ChunkRequest) (*CodeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChunksDN not implemented")
}
func (*UnimplementedDnServiceServer) Estado(context.Context, *CodeRequest) (*CodeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Estado not implemented")
}
func (*UnimplementedDnServiceServer) PropuestasDN(context.Context, *PropRequest) (*CodeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PropuestasDN not implemented")
}

func RegisterDnServiceServer(s *grpc.Server, srv DnServiceServer) {
	s.RegisterService(&_DnService_serviceDesc, srv)
}

func _DnService_EnviarChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnServiceServer).EnviarChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dn_proto.DnService/EnviarChunks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnServiceServer).EnviarChunks(ctx, req.(*ChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnService_ChunksDN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnServiceServer).ChunksDN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dn_proto.DnService/ChunksDN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnServiceServer).ChunksDN(ctx, req.(*ChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnService_Estado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnServiceServer).Estado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dn_proto.DnService/Estado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnServiceServer).Estado(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnService_PropuestasDN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnServiceServer).PropuestasDN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dn_proto.DnService/PropuestasDN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnServiceServer).PropuestasDN(ctx, req.(*PropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DnService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dn_proto.DnService",
	HandlerType: (*DnServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarChunks",
			Handler:    _DnService_EnviarChunks_Handler,
		},
		{
			MethodName: "ChunksDN",
			Handler:    _DnService_ChunksDN_Handler,
		},
		{
			MethodName: "Estado",
			Handler:    _DnService_Estado_Handler,
		},
		{
			MethodName: "PropuestasDN",
			Handler:    _DnService_PropuestasDN_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dn.proto",
}
