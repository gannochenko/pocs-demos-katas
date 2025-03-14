// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: common/image/v1/image.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type CreateImage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ObjectName    string                 `protobuf:"bytes,1,opt,name=objectName,proto3" json:"objectName,omitempty"`
	UploadedAt    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=uploaded_at,json=uploadedAt,proto3" json:"uploaded_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateImage) Reset() {
	*x = CreateImage{}
	mi := &file_common_image_v1_image_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImage) ProtoMessage() {}

func (x *CreateImage) ProtoReflect() protoreflect.Message {
	mi := &file_common_image_v1_image_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImage.ProtoReflect.Descriptor instead.
func (*CreateImage) Descriptor() ([]byte, []int) {
	return file_common_image_v1_image_proto_rawDescGZIP(), []int{0}
}

func (x *CreateImage) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *CreateImage) GetUploadedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UploadedAt
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	IsProcessed   bool                   `protobuf:"varint,3,opt,name=is_processed,json=isProcessed,proto3" json:"is_processed,omitempty"`
	IsFailed      bool                   `protobuf:"varint,4,opt,name=is_failed,json=isFailed,proto3" json:"is_failed,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UploadedAt    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=uploaded_at,json=uploadedAt,proto3" json:"uploaded_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Image) Reset() {
	*x = Image{}
	mi := &file_common_image_v1_image_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_common_image_v1_image_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_common_image_v1_image_proto_rawDescGZIP(), []int{1}
}

func (x *Image) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Image) GetIsProcessed() bool {
	if x != nil {
		return x.IsProcessed
	}
	return false
}

func (x *Image) GetIsFailed() bool {
	if x != nil {
		return x.IsFailed
	}
	return false
}

func (x *Image) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Image) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Image) GetUploadedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UploadedAt
	}
	return nil
}

var File_common_image_v1_image_proto protoreflect.FileDescriptor

var file_common_image_v1_image_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x66,
	0x61, 0x63, 0x65, 0x62, 0x6c, 0x75, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x9c, 0x02, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x64, 0x41, 0x74, 0x42, 0x1f, 0x5a, 0x1d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_common_image_v1_image_proto_rawDescOnce sync.Once
	file_common_image_v1_image_proto_rawDescData []byte
)

func file_common_image_v1_image_proto_rawDescGZIP() []byte {
	file_common_image_v1_image_proto_rawDescOnce.Do(func() {
		file_common_image_v1_image_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_image_v1_image_proto_rawDesc), len(file_common_image_v1_image_proto_rawDesc)))
	})
	return file_common_image_v1_image_proto_rawDescData
}

var file_common_image_v1_image_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_image_v1_image_proto_goTypes = []any{
	(*CreateImage)(nil),           // 0: faceblur.common.image.v1.CreateImage
	(*Image)(nil),                 // 1: faceblur.common.image.v1.Image
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_common_image_v1_image_proto_depIdxs = []int32{
	2, // 0: faceblur.common.image.v1.CreateImage.uploaded_at:type_name -> google.protobuf.Timestamp
	2, // 1: faceblur.common.image.v1.Image.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: faceblur.common.image.v1.Image.updated_at:type_name -> google.protobuf.Timestamp
	2, // 3: faceblur.common.image.v1.Image.uploaded_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_image_v1_image_proto_init() }
func file_common_image_v1_image_proto_init() {
	if File_common_image_v1_image_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_image_v1_image_proto_rawDesc), len(file_common_image_v1_image_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_image_v1_image_proto_goTypes,
		DependencyIndexes: file_common_image_v1_image_proto_depIdxs,
		MessageInfos:      file_common_image_v1_image_proto_msgTypes,
	}.Build()
	File_common_image_v1_image_proto = out.File
	file_common_image_v1_image_proto_goTypes = nil
	file_common_image_v1_image_proto_depIdxs = nil
}
