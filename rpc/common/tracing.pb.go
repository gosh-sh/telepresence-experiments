// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: rpc/common/tracing.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Trace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceData []byte `protobuf:"bytes,1,opt,name=trace_data,json=traceData,proto3" json:"trace_data,omitempty"`
}

func (x *Trace) Reset() {
	*x = Trace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_common_tracing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_common_tracing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_rpc_common_tracing_proto_rawDescGZIP(), []int{0}
}

func (x *Trace) GetTraceData() []byte {
	if x != nil {
		return x.TraceData
	}
	return nil
}

var File_rpc_common_tracing_proto protoreflect.FileDescriptor

var file_rpc_common_tracing_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x05,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x32, 0x4b, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12,
	0x40, 0x0a, 0x0a, 0x44, 0x75, 0x6d, 0x70, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x69, 0x6f, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_common_tracing_proto_rawDescOnce sync.Once
	file_rpc_common_tracing_proto_rawDescData = file_rpc_common_tracing_proto_rawDesc
)

func file_rpc_common_tracing_proto_rawDescGZIP() []byte {
	file_rpc_common_tracing_proto_rawDescOnce.Do(func() {
		file_rpc_common_tracing_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_common_tracing_proto_rawDescData)
	})
	return file_rpc_common_tracing_proto_rawDescData
}

var file_rpc_common_tracing_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rpc_common_tracing_proto_goTypes = []interface{}{
	(*Trace)(nil),         // 0: telepresence.common.Trace
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_rpc_common_tracing_proto_depIdxs = []int32{
	1, // 0: telepresence.common.Tracing.DumpTraces:input_type -> google.protobuf.Empty
	0, // 1: telepresence.common.Tracing.DumpTraces:output_type -> telepresence.common.Trace
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_common_tracing_proto_init() }
func file_rpc_common_tracing_proto_init() {
	if File_rpc_common_tracing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_common_tracing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trace); i {
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
			RawDescriptor: file_rpc_common_tracing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_common_tracing_proto_goTypes,
		DependencyIndexes: file_rpc_common_tracing_proto_depIdxs,
		MessageInfos:      file_rpc_common_tracing_proto_msgTypes,
	}.Build()
	File_rpc_common_tracing_proto = out.File
	file_rpc_common_tracing_proto_rawDesc = nil
	file_rpc_common_tracing_proto_goTypes = nil
	file_rpc_common_tracing_proto_depIdxs = nil
}
