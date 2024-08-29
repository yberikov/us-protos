// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: proto/analytics-service/analytics.proto

package analytics_microservice

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

// LogURLAccessRequest is the request message for the LogURLAccess RPC.
type LogURLAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	UserId int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *LogURLAccessRequest) Reset() {
	*x = LogURLAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_analytics_service_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogURLAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogURLAccessRequest) ProtoMessage() {}

func (x *LogURLAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_analytics_service_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogURLAccessRequest.ProtoReflect.Descriptor instead.
func (*LogURLAccessRequest) Descriptor() ([]byte, []int) {
	return file_proto_analytics_service_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *LogURLAccessRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *LogURLAccessRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// LogURLAccessResponse is the response message for the LogURLAccess RPC.
type LogURLAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LogURLAccessResponse) Reset() {
	*x = LogURLAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_analytics_service_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogURLAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogURLAccessResponse) ProtoMessage() {}

func (x *LogURLAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_analytics_service_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogURLAccessResponse.ProtoReflect.Descriptor instead.
func (*LogURLAccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_analytics_service_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *LogURLAccessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// GetURLStatsRequest is the request message for the GetURLStats RPC.
type GetURLStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetURLStatsRequest) Reset() {
	*x = GetURLStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_analytics_service_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLStatsRequest) ProtoMessage() {}

func (x *GetURLStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_analytics_service_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLStatsRequest.ProtoReflect.Descriptor instead.
func (*GetURLStatsRequest) Descriptor() ([]byte, []int) {
	return file_proto_analytics_service_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *GetURLStatsRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// GetURLStatsResponse is the response message for the GetURLStats RPC.
type GetURLStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url           string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	TotalAccesses int64  `protobuf:"varint,2,opt,name=totalAccesses,proto3" json:"totalAccesses,omitempty"`
}

func (x *GetURLStatsResponse) Reset() {
	*x = GetURLStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_analytics_service_analytics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLStatsResponse) ProtoMessage() {}

func (x *GetURLStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_analytics_service_analytics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLStatsResponse.ProtoReflect.Descriptor instead.
func (*GetURLStatsResponse) Descriptor() ([]byte, []int) {
	return file_proto_analytics_service_analytics_proto_rawDescGZIP(), []int{3}
}

func (x *GetURLStatsResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetURLStatsResponse) GetTotalAccesses() int64 {
	if x != nil {
		return x.TotalAccesses
	}
	return 0
}

var File_proto_analytics_service_analytics_proto protoreflect.FileDescriptor

var file_proto_analytics_service_analytics_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x22, 0x3f, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x55, 0x52, 0x4c, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x55, 0x52, 0x4c, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x26, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x52,
	0x4c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0x4d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x32, 0xb5,
	0x01, 0x0a, 0x10, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x55, 0x52, 0x4c, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1e, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x4c, 0x6f, 0x67, 0x55, 0x52, 0x4c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x4c, 0x6f, 0x67, 0x55, 0x52, 0x4c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x2f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_analytics_service_analytics_proto_rawDescOnce sync.Once
	file_proto_analytics_service_analytics_proto_rawDescData = file_proto_analytics_service_analytics_proto_rawDesc
)

func file_proto_analytics_service_analytics_proto_rawDescGZIP() []byte {
	file_proto_analytics_service_analytics_proto_rawDescOnce.Do(func() {
		file_proto_analytics_service_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_analytics_service_analytics_proto_rawDescData)
	})
	return file_proto_analytics_service_analytics_proto_rawDescData
}

var file_proto_analytics_service_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_analytics_service_analytics_proto_goTypes = []interface{}{
	(*LogURLAccessRequest)(nil),  // 0: analytics.LogURLAccessRequest
	(*LogURLAccessResponse)(nil), // 1: analytics.LogURLAccessResponse
	(*GetURLStatsRequest)(nil),   // 2: analytics.GetURLStatsRequest
	(*GetURLStatsResponse)(nil),  // 3: analytics.GetURLStatsResponse
}
var file_proto_analytics_service_analytics_proto_depIdxs = []int32{
	0, // 0: analytics.AnalyticsService.LogURLAccess:input_type -> analytics.LogURLAccessRequest
	2, // 1: analytics.AnalyticsService.GetURLStats:input_type -> analytics.GetURLStatsRequest
	1, // 2: analytics.AnalyticsService.LogURLAccess:output_type -> analytics.LogURLAccessResponse
	3, // 3: analytics.AnalyticsService.GetURLStats:output_type -> analytics.GetURLStatsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_analytics_service_analytics_proto_init() }
func file_proto_analytics_service_analytics_proto_init() {
	if File_proto_analytics_service_analytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_analytics_service_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogURLAccessRequest); i {
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
		file_proto_analytics_service_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogURLAccessResponse); i {
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
		file_proto_analytics_service_analytics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLStatsRequest); i {
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
		file_proto_analytics_service_analytics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLStatsResponse); i {
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
			RawDescriptor: file_proto_analytics_service_analytics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_analytics_service_analytics_proto_goTypes,
		DependencyIndexes: file_proto_analytics_service_analytics_proto_depIdxs,
		MessageInfos:      file_proto_analytics_service_analytics_proto_msgTypes,
	}.Build()
	File_proto_analytics_service_analytics_proto = out.File
	file_proto_analytics_service_analytics_proto_rawDesc = nil
	file_proto_analytics_service_analytics_proto_goTypes = nil
	file_proto_analytics_service_analytics_proto_depIdxs = nil
}
