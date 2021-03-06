// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: api.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UploadData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*UploadData_Info
	//	*UploadData_Chunk
	Data isUploadData_Data `protobuf_oneof:"data"`
}

func (x *UploadData) Reset() {
	*x = UploadData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadData) ProtoMessage() {}

func (x *UploadData) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadData.ProtoReflect.Descriptor instead.
func (*UploadData) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (m *UploadData) GetData() isUploadData_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadData) GetInfo() *Data {
	if x, ok := x.GetData().(*UploadData_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadData) GetChunk() *NewChunk {
	if x, ok := x.GetData().(*UploadData_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isUploadData_Data interface {
	isUploadData_Data()
}

type UploadData_Info struct {
	Info *Data `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadData_Chunk struct {
	Chunk *NewChunk `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*UploadData_Info) isUploadData_Data() {}

func (*UploadData_Chunk) isUploadData_Data() {}

type DeleteFileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DeleteFileInfo) Reset() {
	*x = DeleteFileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileInfo) ProtoMessage() {}

func (x *DeleteFileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileInfo.ProtoReflect.Descriptor instead.
func (*DeleteFileInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteFileInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Meta []byte `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Data) GetMeta() []byte {
	if x != nil {
		return x.Meta
	}
	return nil
}

type NewChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *NewChunk) Reset() {
	*x = NewChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewChunk) ProtoMessage() {}

func (x *NewChunk) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewChunk.ProtoReflect.Descriptor instead.
func (*NewChunk) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *NewChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e,
	0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x26,
	0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x22, 0x1c, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x2e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x22, 0x1e, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x70, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x2c, 0x0a, 0x0a, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x28, 0x01, 0x12, 0x3a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_goTypes = []interface{}{
	(*UploadData)(nil),     // 0: grpc.UploadData
	(*DeleteFileInfo)(nil), // 1: grpc.DeleteFileInfo
	(*UUID)(nil),           // 2: grpc.UUID
	(*Data)(nil),           // 3: grpc.Data
	(*NewChunk)(nil),       // 4: grpc.NewChunk
	(*empty.Empty)(nil),    // 5: google.protobuf.Empty
}
var file_api_proto_depIdxs = []int32{
	3, // 0: grpc.UploadData.info:type_name -> grpc.Data
	4, // 1: grpc.UploadData.chunk:type_name -> grpc.NewChunk
	0, // 2: grpc.Disk.UploadFile:input_type -> grpc.UploadData
	1, // 3: grpc.Disk.DeleteFile:input_type -> grpc.DeleteFileInfo
	2, // 4: grpc.Disk.UploadFile:output_type -> grpc.UUID
	5, // 5: grpc.Disk.DeleteFile:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadData); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileInfo); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewChunk); i {
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
	file_api_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UploadData_Info)(nil),
		(*UploadData_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DiskClient is the client API for Disk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiskClient interface {
	// UploadFile upload new file by path.
	// First message should be FileInfo, all next message contains chunk.
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (Disk_UploadFileClient, error)
	// DeleteFile remove file by path.
	DeleteFile(ctx context.Context, in *DeleteFileInfo, opts ...grpc.CallOption) (*empty.Empty, error)
}

type diskClient struct {
	cc grpc.ClientConnInterface
}

func NewDiskClient(cc grpc.ClientConnInterface) DiskClient {
	return &diskClient{cc}
}

func (c *diskClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (Disk_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Disk_serviceDesc.Streams[0], "/grpc.Disk/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &diskUploadFileClient{stream}
	return x, nil
}

type Disk_UploadFileClient interface {
	Send(*UploadData) error
	CloseAndRecv() (*UUID, error)
	grpc.ClientStream
}

type diskUploadFileClient struct {
	grpc.ClientStream
}

func (x *diskUploadFileClient) Send(m *UploadData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *diskUploadFileClient) CloseAndRecv() (*UUID, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UUID)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *diskClient) DeleteFile(ctx context.Context, in *DeleteFileInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Disk/DeleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiskServer is the server API for Disk service.
type DiskServer interface {
	// UploadFile upload new file by path.
	// First message should be FileInfo, all next message contains chunk.
	UploadFile(Disk_UploadFileServer) error
	// DeleteFile remove file by path.
	DeleteFile(context.Context, *DeleteFileInfo) (*empty.Empty, error)
}

// UnimplementedDiskServer can be embedded to have forward compatible implementations.
type UnimplementedDiskServer struct {
}

func (*UnimplementedDiskServer) UploadFile(Disk_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedDiskServer) DeleteFile(context.Context, *DeleteFileInfo) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}

func RegisterDiskServer(s *grpc.Server, srv DiskServer) {
	s.RegisterService(&_Disk_serviceDesc, srv)
}

func _Disk_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DiskServer).UploadFile(&diskUploadFileServer{stream})
}

type Disk_UploadFileServer interface {
	SendAndClose(*UUID) error
	Recv() (*UploadData, error)
	grpc.ServerStream
}

type diskUploadFileServer struct {
	grpc.ServerStream
}

func (x *diskUploadFileServer) SendAndClose(m *UUID) error {
	return x.ServerStream.SendMsg(m)
}

func (x *diskUploadFileServer) Recv() (*UploadData, error) {
	m := new(UploadData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Disk_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Disk/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServer).DeleteFile(ctx, req.(*DeleteFileInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Disk_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Disk",
	HandlerType: (*DiskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteFile",
			Handler:    _Disk_DeleteFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _Disk_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}
