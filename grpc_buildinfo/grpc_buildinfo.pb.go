// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: grpc_buildinfo/grpc_buildinfo.proto

package grpc_buildinfo

import (
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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_grpc_buildinfo_grpc_buildinfo_proto_rawDescGZIP(), []int{0}
}

type BuildInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoVersion string          `protobuf:"bytes,1,opt,name=go_version,json=goVersion,proto3" json:"go_version,omitempty"`
	Path      string          `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Main      *Module         `protobuf:"bytes,3,opt,name=main,proto3" json:"main,omitempty"`
	Deps      []*Module       `protobuf:"bytes,4,rep,name=deps,proto3" json:"deps,omitempty"`
	Settings  []*BuildSetting `protobuf:"bytes,5,rep,name=settings,proto3" json:"settings,omitempty"`
}

func (x *BuildInfo) Reset() {
	*x = BuildInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildInfo) ProtoMessage() {}

func (x *BuildInfo) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildInfo.ProtoReflect.Descriptor instead.
func (*BuildInfo) Descriptor() ([]byte, []int) {
	return file_grpc_buildinfo_grpc_buildinfo_proto_rawDescGZIP(), []int{1}
}

func (x *BuildInfo) GetGoVersion() string {
	if x != nil {
		return x.GoVersion
	}
	return ""
}

func (x *BuildInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *BuildInfo) GetMain() *Module {
	if x != nil {
		return x.Main
	}
	return nil
}

func (x *BuildInfo) GetDeps() []*Module {
	if x != nil {
		return x.Deps
	}
	return nil
}

func (x *BuildInfo) GetSettings() []*BuildSetting {
	if x != nil {
		return x.Settings
	}
	return nil
}

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Version string  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Sum     string  `protobuf:"bytes,3,opt,name=sum,proto3" json:"sum,omitempty"`
	Replace *Module `protobuf:"bytes,4,opt,name=replace,proto3" json:"replace,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_grpc_buildinfo_grpc_buildinfo_proto_rawDescGZIP(), []int{2}
}

func (x *Module) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Module) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Module) GetSum() string {
	if x != nil {
		return x.Sum
	}
	return ""
}

func (x *Module) GetReplace() *Module {
	if x != nil {
		return x.Replace
	}
	return nil
}

type BuildSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BuildSetting) Reset() {
	*x = BuildSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildSetting) ProtoMessage() {}

func (x *BuildSetting) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildSetting.ProtoReflect.Descriptor instead.
func (*BuildSetting) Descriptor() ([]byte, []int) {
	return file_grpc_buildinfo_grpc_buildinfo_proto_rawDescGZIP(), []int{3}
}

func (x *BuildSetting) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *BuildSetting) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_grpc_buildinfo_grpc_buildinfo_proto protoreflect.FileDescriptor

var file_grpc_buildinfo_grpc_buildinfo_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x66, 0x6f,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x66, 0x6f,
	0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xc1, 0x01, 0x0a, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x6f, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x25, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x65, 0x70, 0x73, 0x12,
	0x33, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x22, 0x75, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x2b,
	0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x22, 0x36, 0x0a, 0x0c, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0x51, 0x0a, 0x10, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x6c, 0x73, 0x65, 0x79, 0x68, 0x69, 0x67, 0x68, 0x74,
	0x6f, 0x77, 0x65, 0x72, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_buildinfo_grpc_buildinfo_proto_rawDescOnce sync.Once
	file_grpc_buildinfo_grpc_buildinfo_proto_rawDescData = file_grpc_buildinfo_grpc_buildinfo_proto_rawDesc
)

func file_grpc_buildinfo_grpc_buildinfo_proto_rawDescGZIP() []byte {
	file_grpc_buildinfo_grpc_buildinfo_proto_rawDescOnce.Do(func() {
		file_grpc_buildinfo_grpc_buildinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_buildinfo_grpc_buildinfo_proto_rawDescData)
	})
	return file_grpc_buildinfo_grpc_buildinfo_proto_rawDescData
}

var file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_buildinfo_grpc_buildinfo_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil), // 0: buildinfo.EmptyRequest
	(*BuildInfo)(nil),    // 1: buildinfo.BuildInfo
	(*Module)(nil),       // 2: buildinfo.Module
	(*BuildSetting)(nil), // 3: buildinfo.BuildSetting
}
var file_grpc_buildinfo_grpc_buildinfo_proto_depIdxs = []int32{
	2, // 0: buildinfo.BuildInfo.main:type_name -> buildinfo.Module
	2, // 1: buildinfo.BuildInfo.deps:type_name -> buildinfo.Module
	3, // 2: buildinfo.BuildInfo.settings:type_name -> buildinfo.BuildSetting
	2, // 3: buildinfo.Module.replace:type_name -> buildinfo.Module
	0, // 4: buildinfo.BuildInfoService.GetBuildInfo:input_type -> buildinfo.EmptyRequest
	1, // 5: buildinfo.BuildInfoService.GetBuildInfo:output_type -> buildinfo.BuildInfo
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_grpc_buildinfo_grpc_buildinfo_proto_init() }
func file_grpc_buildinfo_grpc_buildinfo_proto_init() {
	if File_grpc_buildinfo_grpc_buildinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildInfo); i {
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
		file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
		file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildSetting); i {
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
			RawDescriptor: file_grpc_buildinfo_grpc_buildinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_buildinfo_grpc_buildinfo_proto_goTypes,
		DependencyIndexes: file_grpc_buildinfo_grpc_buildinfo_proto_depIdxs,
		MessageInfos:      file_grpc_buildinfo_grpc_buildinfo_proto_msgTypes,
	}.Build()
	File_grpc_buildinfo_grpc_buildinfo_proto = out.File
	file_grpc_buildinfo_grpc_buildinfo_proto_rawDesc = nil
	file_grpc_buildinfo_grpc_buildinfo_proto_goTypes = nil
	file_grpc_buildinfo_grpc_buildinfo_proto_depIdxs = nil
}
