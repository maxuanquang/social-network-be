// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: newsfeed.proto

package newsfeed

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

type GetNewsfeedResponse_GetNewsfeedStatus int32

const (
	GetNewsfeedResponse_OK             GetNewsfeedResponse_GetNewsfeedStatus = 0
	GetNewsfeedResponse_NEWSFEED_EMPTY GetNewsfeedResponse_GetNewsfeedStatus = 1
)

// Enum value maps for GetNewsfeedResponse_GetNewsfeedStatus.
var (
	GetNewsfeedResponse_GetNewsfeedStatus_name = map[int32]string{
		0: "OK",
		1: "NEWSFEED_EMPTY",
	}
	GetNewsfeedResponse_GetNewsfeedStatus_value = map[string]int32{
		"OK":             0,
		"NEWSFEED_EMPTY": 1,
	}
)

func (x GetNewsfeedResponse_GetNewsfeedStatus) Enum() *GetNewsfeedResponse_GetNewsfeedStatus {
	p := new(GetNewsfeedResponse_GetNewsfeedStatus)
	*p = x
	return p
}

func (x GetNewsfeedResponse_GetNewsfeedStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetNewsfeedResponse_GetNewsfeedStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_newsfeed_proto_enumTypes[0].Descriptor()
}

func (GetNewsfeedResponse_GetNewsfeedStatus) Type() protoreflect.EnumType {
	return &file_newsfeed_proto_enumTypes[0]
}

func (x GetNewsfeedResponse_GetNewsfeedStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetNewsfeedResponse_GetNewsfeedStatus.Descriptor instead.
func (GetNewsfeedResponse_GetNewsfeedStatus) EnumDescriptor() ([]byte, []int) {
	return file_newsfeed_proto_rawDescGZIP(), []int{1, 0}
}

type GetNewsfeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetNewsfeedRequest) Reset() {
	*x = GetNewsfeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_newsfeed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsfeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsfeedRequest) ProtoMessage() {}

func (x *GetNewsfeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_newsfeed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsfeedRequest.ProtoReflect.Descriptor instead.
func (*GetNewsfeedRequest) Descriptor() ([]byte, []int) {
	return file_newsfeed_proto_rawDescGZIP(), []int{0}
}

func (x *GetNewsfeedRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetNewsfeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   GetNewsfeedResponse_GetNewsfeedStatus `protobuf:"varint,1,opt,name=status,proto3,enum=newsfeed.GetNewsfeedResponse_GetNewsfeedStatus" json:"status,omitempty"`
	PostsIds []int64                               `protobuf:"varint,2,rep,packed,name=posts_ids,json=postsIds,proto3" json:"posts_ids,omitempty"`
}

func (x *GetNewsfeedResponse) Reset() {
	*x = GetNewsfeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_newsfeed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsfeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsfeedResponse) ProtoMessage() {}

func (x *GetNewsfeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_newsfeed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsfeedResponse.ProtoReflect.Descriptor instead.
func (*GetNewsfeedResponse) Descriptor() ([]byte, []int) {
	return file_newsfeed_proto_rawDescGZIP(), []int{1}
}

func (x *GetNewsfeedResponse) GetStatus() GetNewsfeedResponse_GetNewsfeedStatus {
	if x != nil {
		return x.Status
	}
	return GetNewsfeedResponse_OK
}

func (x *GetNewsfeedResponse) GetPostsIds() []int64 {
	if x != nil {
		return x.PostsIds
	}
	return nil
}

var File_newsfeed_proto protoreflect.FileDescriptor

var file_newsfeed_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xac, 0x01, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2f, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x47, 0x65, 0x74,
	0x4e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x49, 0x64, 0x73, 0x22, 0x2f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x45, 0x57, 0x53, 0x46, 0x45, 0x45, 0x44,
	0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x01, 0x32, 0x58, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x73,
	0x66, 0x65, 0x65, 0x64, 0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x66,
	0x65, 0x65, 0x64, 0x12, 0x1c, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x47,
	0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x47, 0x65, 0x74,
	0x4e, 0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x61, 0x78, 0x75, 0x61, 0x6e, 0x71, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x6e,
	0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_newsfeed_proto_rawDescOnce sync.Once
	file_newsfeed_proto_rawDescData = file_newsfeed_proto_rawDesc
)

func file_newsfeed_proto_rawDescGZIP() []byte {
	file_newsfeed_proto_rawDescOnce.Do(func() {
		file_newsfeed_proto_rawDescData = protoimpl.X.CompressGZIP(file_newsfeed_proto_rawDescData)
	})
	return file_newsfeed_proto_rawDescData
}

var file_newsfeed_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_newsfeed_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_newsfeed_proto_goTypes = []interface{}{
	(GetNewsfeedResponse_GetNewsfeedStatus)(0), // 0: newsfeed.GetNewsfeedResponse.GetNewsfeedStatus
	(*GetNewsfeedRequest)(nil),                 // 1: newsfeed.GetNewsfeedRequest
	(*GetNewsfeedResponse)(nil),                // 2: newsfeed.GetNewsfeedResponse
}
var file_newsfeed_proto_depIdxs = []int32{
	0, // 0: newsfeed.GetNewsfeedResponse.status:type_name -> newsfeed.GetNewsfeedResponse.GetNewsfeedStatus
	1, // 1: newsfeed.Newsfeed.GetNewsfeed:input_type -> newsfeed.GetNewsfeedRequest
	2, // 2: newsfeed.Newsfeed.GetNewsfeed:output_type -> newsfeed.GetNewsfeedResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_newsfeed_proto_init() }
func file_newsfeed_proto_init() {
	if File_newsfeed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_newsfeed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsfeedRequest); i {
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
		file_newsfeed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsfeedResponse); i {
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
			RawDescriptor: file_newsfeed_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_newsfeed_proto_goTypes,
		DependencyIndexes: file_newsfeed_proto_depIdxs,
		EnumInfos:         file_newsfeed_proto_enumTypes,
		MessageInfos:      file_newsfeed_proto_msgTypes,
	}.Build()
	File_newsfeed_proto = out.File
	file_newsfeed_proto_rawDesc = nil
	file_newsfeed_proto_goTypes = nil
	file_newsfeed_proto_depIdxs = nil
}
