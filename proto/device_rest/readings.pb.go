// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v5.27.3
// source: proto/device_rest/readings.proto

package device_rest

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

type CumulativeCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CumulativeCount int32 `protobuf:"varint,1,opt,name=cumulative_count,json=cumulativeCount,proto3" json:"cumulative_count,omitempty"`
}

func (x *CumulativeCount) Reset() {
	*x = CumulativeCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_rest_readings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CumulativeCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CumulativeCount) ProtoMessage() {}

func (x *CumulativeCount) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_rest_readings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CumulativeCount.ProtoReflect.Descriptor instead.
func (*CumulativeCount) Descriptor() ([]byte, []int) {
	return file_proto_device_rest_readings_proto_rawDescGZIP(), []int{0}
}

func (x *CumulativeCount) GetCumulativeCount() int32 {
	if x != nil {
		return x.CumulativeCount
	}
	return 0
}

type LatestTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTimestamp string `protobuf:"bytes,1,opt,name=latest_timestamp,json=latestTimestamp,proto3" json:"latest_timestamp,omitempty"`
}

func (x *LatestTimestamp) Reset() {
	*x = LatestTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_rest_readings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatestTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatestTimestamp) ProtoMessage() {}

func (x *LatestTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_rest_readings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatestTimestamp.ProtoReflect.Descriptor instead.
func (*LatestTimestamp) Descriptor() ([]byte, []int) {
	return file_proto_device_rest_readings_proto_rawDescGZIP(), []int{1}
}

func (x *LatestTimestamp) GetLatestTimestamp() string {
	if x != nil {
		return x.LatestTimestamp
	}
	return ""
}

type SomeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Readings []*SomeMessage_Readings `protobuf:"bytes,2,rep,name=readings,proto3" json:"readings,omitempty"`
}

func (x *SomeMessage) Reset() {
	*x = SomeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_rest_readings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeMessage) ProtoMessage() {}

func (x *SomeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_rest_readings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeMessage.ProtoReflect.Descriptor instead.
func (*SomeMessage) Descriptor() ([]byte, []int) {
	return file_proto_device_rest_readings_proto_rawDescGZIP(), []int{2}
}

func (x *SomeMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SomeMessage) GetReadings() []*SomeMessage_Readings {
	if x != nil {
		return x.Readings
	}
	return nil
}

type SomeMessage_Readings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Count     int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SomeMessage_Readings) Reset() {
	*x = SomeMessage_Readings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_rest_readings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeMessage_Readings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeMessage_Readings) ProtoMessage() {}

func (x *SomeMessage_Readings) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_rest_readings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeMessage_Readings.ProtoReflect.Descriptor instead.
func (*SomeMessage_Readings) Descriptor() ([]byte, []int) {
	return file_proto_device_rest_readings_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SomeMessage_Readings) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *SomeMessage_Readings) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_proto_device_rest_readings_proto protoreflect.FileDescriptor

var file_proto_device_rest_readings_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x3d, 0x0a, 0x10, 0x63, 0x75, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10,
	0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3d, 0x0a, 0x10, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x29, 0x0a, 0x10, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x95, 0x01, 0x0a, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x53, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x3e,
	0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x10,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_device_rest_readings_proto_rawDescOnce sync.Once
	file_proto_device_rest_readings_proto_rawDescData = file_proto_device_rest_readings_proto_rawDesc
)

func file_proto_device_rest_readings_proto_rawDescGZIP() []byte {
	file_proto_device_rest_readings_proto_rawDescOnce.Do(func() {
		file_proto_device_rest_readings_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_device_rest_readings_proto_rawDescData)
	})
	return file_proto_device_rest_readings_proto_rawDescData
}

var file_proto_device_rest_readings_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_device_rest_readings_proto_goTypes = []interface{}{
	(*CumulativeCount)(nil),      // 0: main.cumulative_count
	(*LatestTimestamp)(nil),      // 1: main.latest_timestamp
	(*SomeMessage)(nil),          // 2: main.SomeMessage
	(*SomeMessage_Readings)(nil), // 3: main.SomeMessage.Readings
}
var file_proto_device_rest_readings_proto_depIdxs = []int32{
	3, // 0: main.SomeMessage.readings:type_name -> main.SomeMessage.Readings
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_device_rest_readings_proto_init() }
func file_proto_device_rest_readings_proto_init() {
	if File_proto_device_rest_readings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_device_rest_readings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CumulativeCount); i {
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
		file_proto_device_rest_readings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatestTimestamp); i {
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
		file_proto_device_rest_readings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeMessage); i {
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
		file_proto_device_rest_readings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeMessage_Readings); i {
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
			RawDescriptor: file_proto_device_rest_readings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_device_rest_readings_proto_goTypes,
		DependencyIndexes: file_proto_device_rest_readings_proto_depIdxs,
		MessageInfos:      file_proto_device_rest_readings_proto_msgTypes,
	}.Build()
	File_proto_device_rest_readings_proto = out.File
	file_proto_device_rest_readings_proto_rawDesc = nil
	file_proto_device_rest_readings_proto_goTypes = nil
	file_proto_device_rest_readings_proto_depIdxs = nil
}
