// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: cache/cache.proto

package cache

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_cache_cache_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cache_cache_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_cache_cache_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type KeyValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	mi := &file_cache_cache_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_cache_cache_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_cache_cache_proto_rawDescGZIP(), []int{1}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        string                 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutResponse) Reset() {
	*x = PutResponse{}
	mi := &file_cache_cache_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResponse) ProtoMessage() {}

func (x *PutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cache_cache_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResponse.ProtoReflect.Descriptor instead.
func (*PutResponse) Descriptor() ([]byte, []int) {
	return file_cache_cache_proto_rawDescGZIP(), []int{2}
}

func (x *PutResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *PutResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_cache_cache_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cache_cache_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_cache_cache_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        string                 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_cache_cache_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cache_cache_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_cache_cache_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *DeleteResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_cache_cache_proto protoreflect.FileDescriptor

const file_cache_cache_proto_rawDesc = "" +
	"\n" +
	"\x11cache/cache.proto\x12\x05cache\"\x1e\n" +
	"\n" +
	"GetRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\"2\n" +
	"\bKeyValue\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value\";\n" +
	"\vPutResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\tR\x06result\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\"!\n" +
	"\rDeleteRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\">\n" +
	"\x0eDeleteResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\tR\x06result\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error2\xa6\x01\n" +
	"\x10KiviCacheService\x12,\n" +
	"\x03Put\x12\x0f.cache.KeyValue\x1a\x12.cache.PutResponse\"\x00\x12+\n" +
	"\x03Get\x12\x11.cache.GetRequest\x1a\x0f.cache.KeyValue\"\x00\x127\n" +
	"\x06Delete\x12\x14.cache.DeleteRequest\x1a\x15.cache.DeleteResponse\"\x00B\x11Z\x0fkivicache/cacheb\x06proto3"

var (
	file_cache_cache_proto_rawDescOnce sync.Once
	file_cache_cache_proto_rawDescData []byte
)

func file_cache_cache_proto_rawDescGZIP() []byte {
	file_cache_cache_proto_rawDescOnce.Do(func() {
		file_cache_cache_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cache_cache_proto_rawDesc), len(file_cache_cache_proto_rawDesc)))
	})
	return file_cache_cache_proto_rawDescData
}

var file_cache_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cache_cache_proto_goTypes = []any{
	(*GetRequest)(nil),     // 0: cache.GetRequest
	(*KeyValue)(nil),       // 1: cache.KeyValue
	(*PutResponse)(nil),    // 2: cache.PutResponse
	(*DeleteRequest)(nil),  // 3: cache.DeleteRequest
	(*DeleteResponse)(nil), // 4: cache.DeleteResponse
}
var file_cache_cache_proto_depIdxs = []int32{
	1, // 0: cache.KiviCacheService.Put:input_type -> cache.KeyValue
	0, // 1: cache.KiviCacheService.Get:input_type -> cache.GetRequest
	3, // 2: cache.KiviCacheService.Delete:input_type -> cache.DeleteRequest
	2, // 3: cache.KiviCacheService.Put:output_type -> cache.PutResponse
	1, // 4: cache.KiviCacheService.Get:output_type -> cache.KeyValue
	4, // 5: cache.KiviCacheService.Delete:output_type -> cache.DeleteResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cache_cache_proto_init() }
func file_cache_cache_proto_init() {
	if File_cache_cache_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cache_cache_proto_rawDesc), len(file_cache_cache_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cache_cache_proto_goTypes,
		DependencyIndexes: file_cache_cache_proto_depIdxs,
		MessageInfos:      file_cache_cache_proto_msgTypes,
	}.Build()
	File_cache_cache_proto = out.File
	file_cache_cache_proto_goTypes = nil
	file_cache_cache_proto_depIdxs = nil
}
