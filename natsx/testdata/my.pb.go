// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: my.proto

package testdata

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

type TestMine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Sex  string `protobuf:"bytes,3,opt,name=Sex,proto3" json:"Sex,omitempty"`
}

func (x *TestMine) Reset() {
	*x = TestMine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMine) ProtoMessage() {}

func (x *TestMine) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMine.ProtoReflect.Descriptor instead.
func (*TestMine) Descriptor() ([]byte, []int) {
	return file_my_proto_rawDescGZIP(), []int{0}
}

func (x *TestMine) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestMine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestMine) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

type TestMineResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Brother string  `protobuf:"bytes,2,opt,name=brother,proto3" json:"brother,omitempty"`
	Childs  []int64 `protobuf:"varint,3,rep,packed,name=childs,proto3" json:"childs,omitempty"`
}

func (x *TestMineResp) Reset() {
	*x = TestMineResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMineResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMineResp) ProtoMessage() {}

func (x *TestMineResp) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMineResp.ProtoReflect.Descriptor instead.
func (*TestMineResp) Descriptor() ([]byte, []int) {
	return file_my_proto_rawDescGZIP(), []int{1}
}

func (x *TestMineResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestMineResp) GetBrother() string {
	if x != nil {
		return x.Brother
	}
	return ""
}

func (x *TestMineResp) GetChilds() []int64 {
	if x != nil {
		return x.Childs
	}
	return nil
}

var File_my_proto protoreflect.FileDescriptor

var file_my_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x40, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x6e, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x53, 0x65, 0x78, 0x22, 0x50, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x06, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_my_proto_rawDescOnce sync.Once
	file_my_proto_rawDescData = file_my_proto_rawDesc
)

func file_my_proto_rawDescGZIP() []byte {
	file_my_proto_rawDescOnce.Do(func() {
		file_my_proto_rawDescData = protoimpl.X.CompressGZIP(file_my_proto_rawDescData)
	})
	return file_my_proto_rawDescData
}

var file_my_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_my_proto_goTypes = []interface{}{
	(*TestMine)(nil),     // 0: testdata.TestMine
	(*TestMineResp)(nil), // 1: testdata.TestMineResp
}
var file_my_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_my_proto_init() }
func file_my_proto_init() {
	if File_my_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_my_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMine); i {
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
		file_my_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMineResp); i {
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
			RawDescriptor: file_my_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_my_proto_goTypes,
		DependencyIndexes: file_my_proto_depIdxs,
		MessageInfos:      file_my_proto_msgTypes,
	}.Build()
	File_my_proto = out.File
	file_my_proto_rawDesc = nil
	file_my_proto_goTypes = nil
	file_my_proto_depIdxs = nil
}