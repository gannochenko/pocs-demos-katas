// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: common/page_navigation/v1/page_navigation.proto

package v1

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

type PageNavigationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PageSize      int32                  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNumber    int32                  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PageNavigationRequest) Reset() {
	*x = PageNavigationRequest{}
	mi := &file_common_page_navigation_v1_page_navigation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PageNavigationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageNavigationRequest) ProtoMessage() {}

func (x *PageNavigationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_page_navigation_v1_page_navigation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageNavigationRequest.ProtoReflect.Descriptor instead.
func (*PageNavigationRequest) Descriptor() ([]byte, []int) {
	return file_common_page_navigation_v1_page_navigation_proto_rawDescGZIP(), []int{0}
}

func (x *PageNavigationRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageNavigationRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

type PageNavigationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PageSize      int32                  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNumber    int32                  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageCount     int32                  `protobuf:"varint,3,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
	Total         int32                  `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PageNavigationResponse) Reset() {
	*x = PageNavigationResponse{}
	mi := &file_common_page_navigation_v1_page_navigation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PageNavigationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageNavigationResponse) ProtoMessage() {}

func (x *PageNavigationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_page_navigation_v1_page_navigation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageNavigationResponse.ProtoReflect.Descriptor instead.
func (*PageNavigationResponse) Descriptor() ([]byte, []int) {
	return file_common_page_navigation_v1_page_navigation_proto_rawDescGZIP(), []int{1}
}

func (x *PageNavigationResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageNavigationResponse) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *PageNavigationResponse) GetPageCount() int32 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

func (x *PageNavigationResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_common_page_navigation_v1_page_navigation_proto protoreflect.FileDescriptor

var file_common_page_navigation_v1_page_navigation_proto_rawDesc = string([]byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61,
	0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6c, 0x75, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x55, 0x0a, 0x15, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x8b, 0x01, 0x0a,
	0x16, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x29, 0x5a, 0x27, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_common_page_navigation_v1_page_navigation_proto_rawDescOnce sync.Once
	file_common_page_navigation_v1_page_navigation_proto_rawDescData []byte
)

func file_common_page_navigation_v1_page_navigation_proto_rawDescGZIP() []byte {
	file_common_page_navigation_v1_page_navigation_proto_rawDescOnce.Do(func() {
		file_common_page_navigation_v1_page_navigation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_page_navigation_v1_page_navigation_proto_rawDesc), len(file_common_page_navigation_v1_page_navigation_proto_rawDesc)))
	})
	return file_common_page_navigation_v1_page_navigation_proto_rawDescData
}

var file_common_page_navigation_v1_page_navigation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_page_navigation_v1_page_navigation_proto_goTypes = []any{
	(*PageNavigationRequest)(nil),  // 0: faceblur.common.page_navigation.v1.PageNavigationRequest
	(*PageNavigationResponse)(nil), // 1: faceblur.common.page_navigation.v1.PageNavigationResponse
}
var file_common_page_navigation_v1_page_navigation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_page_navigation_v1_page_navigation_proto_init() }
func file_common_page_navigation_v1_page_navigation_proto_init() {
	if File_common_page_navigation_v1_page_navigation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_page_navigation_v1_page_navigation_proto_rawDesc), len(file_common_page_navigation_v1_page_navigation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_page_navigation_v1_page_navigation_proto_goTypes,
		DependencyIndexes: file_common_page_navigation_v1_page_navigation_proto_depIdxs,
		MessageInfos:      file_common_page_navigation_v1_page_navigation_proto_msgTypes,
	}.Build()
	File_common_page_navigation_v1_page_navigation_proto = out.File
	file_common_page_navigation_v1_page_navigation_proto_goTypes = nil
	file_common_page_navigation_v1_page_navigation_proto_depIdxs = nil
}
