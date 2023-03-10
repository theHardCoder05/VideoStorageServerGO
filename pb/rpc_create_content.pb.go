// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: rpc_create_content.proto

package pb

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

type CreateContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fileid string `protobuf:"bytes,1,opt,name=fileid,proto3" json:"fileid,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size   int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *CreateContentRequest) Reset() {
	*x = CreateContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentRequest) ProtoMessage() {}

func (x *CreateContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentRequest.ProtoReflect.Descriptor instead.
func (*CreateContentRequest) Descriptor() ([]byte, []int) {
	return file_rpc_create_content_proto_rawDescGZIP(), []int{0}
}

func (x *CreateContentRequest) GetFileid() string {
	if x != nil {
		return x.Fileid
	}
	return ""
}

func (x *CreateContentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateContentRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type CreateContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content *Content `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateContentResponse) Reset() {
	*x = CreateContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentResponse) ProtoMessage() {}

func (x *CreateContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentResponse.ProtoReflect.Descriptor instead.
func (*CreateContentResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_content_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContentResponse) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_rpc_create_content_proto protoreflect.FileDescriptor

var file_rpc_create_content_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x64, 0x62, 0x1a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x18, 0x5a, 0x16, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_create_content_proto_rawDescOnce sync.Once
	file_rpc_create_content_proto_rawDescData = file_rpc_create_content_proto_rawDesc
)

func file_rpc_create_content_proto_rawDescGZIP() []byte {
	file_rpc_create_content_proto_rawDescOnce.Do(func() {
		file_rpc_create_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_create_content_proto_rawDescData)
	})
	return file_rpc_create_content_proto_rawDescData
}

var file_rpc_create_content_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_create_content_proto_goTypes = []interface{}{
	(*CreateContentRequest)(nil),  // 0: db.CreateContentRequest
	(*CreateContentResponse)(nil), // 1: db.CreateContentResponse
	(*Content)(nil),               // 2: db.Content
}
var file_rpc_create_content_proto_depIdxs = []int32{
	2, // 0: db.CreateContentResponse.content:type_name -> db.Content
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_create_content_proto_init() }
func file_rpc_create_content_proto_init() {
	if File_rpc_create_content_proto != nil {
		return
	}
	file_content_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_create_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContentRequest); i {
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
		file_rpc_create_content_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContentResponse); i {
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
			RawDescriptor: file_rpc_create_content_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_create_content_proto_goTypes,
		DependencyIndexes: file_rpc_create_content_proto_depIdxs,
		MessageInfos:      file_rpc_create_content_proto_msgTypes,
	}.Build()
	File_rpc_create_content_proto = out.File
	file_rpc_create_content_proto_rawDesc = nil
	file_rpc_create_content_proto_goTypes = nil
	file_rpc_create_content_proto_depIdxs = nil
}
