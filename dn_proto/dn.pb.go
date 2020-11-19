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

var File_dn_proto protoreflect.FileDescriptor

var file_dn_proto_rawDesc = []byte{
	0x0a, 0x08, 0x64, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x4d, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06,
	0x42, 0x75, 0x73, 0x63, 0x61, 0x72, 0x12, 0x15, 0x2e, 0x64, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x64, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_dn_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dn_proto_goTypes = []interface{}{
	(*CodeRequest)(nil), // 0: dn_proto.CodeRequest
}
var file_dn_proto_depIdxs = []int32{
	0, // 0: dn_proto.HelloworldService.Buscar:input_type -> dn_proto.CodeRequest
	0, // 1: dn_proto.HelloworldService.Buscar:output_type -> dn_proto.CodeRequest
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
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

// HelloworldServiceClient is the client API for HelloworldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloworldServiceClient interface {
	Buscar(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeRequest, error)
}

type helloworldServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloworldServiceClient(cc grpc.ClientConnInterface) HelloworldServiceClient {
	return &helloworldServiceClient{cc}
}

func (c *helloworldServiceClient) Buscar(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeRequest, error) {
	out := new(CodeRequest)
	err := c.cc.Invoke(ctx, "/dn_proto.HelloworldService/Buscar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloworldServiceServer is the server API for HelloworldService service.
type HelloworldServiceServer interface {
	Buscar(context.Context, *CodeRequest) (*CodeRequest, error)
}

// UnimplementedHelloworldServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloworldServiceServer struct {
}

func (*UnimplementedHelloworldServiceServer) Buscar(context.Context, *CodeRequest) (*CodeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Buscar not implemented")
}

func RegisterHelloworldServiceServer(s *grpc.Server, srv HelloworldServiceServer) {
	s.RegisterService(&_HelloworldService_serviceDesc, srv)
}

func _HelloworldService_Buscar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServiceServer).Buscar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dn_proto.HelloworldService/Buscar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServiceServer).Buscar(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloworldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dn_proto.HelloworldService",
	HandlerType: (*HelloworldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Buscar",
			Handler:    _HelloworldService_Buscar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dn.proto",
}