// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: weather/v1/weather.proto

package weatherv1

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

type CheckWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressLineOne string `protobuf:"bytes,1,opt,name=address_line_one,json=addressLineOne,proto3" json:"address_line_one,omitempty"`
	AddressLineTwo string `protobuf:"bytes,2,opt,name=address_line_two,json=addressLineTwo,proto3" json:"address_line_two,omitempty"`
	City           string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	State          string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Zipcode        string `protobuf:"bytes,5,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
}

func (x *CheckWeatherRequest) Reset() {
	*x = CheckWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckWeatherRequest) ProtoMessage() {}

func (x *CheckWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckWeatherRequest.ProtoReflect.Descriptor instead.
func (*CheckWeatherRequest) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{0}
}

func (x *CheckWeatherRequest) GetAddressLineOne() string {
	if x != nil {
		return x.AddressLineOne
	}
	return ""
}

func (x *CheckWeatherRequest) GetAddressLineTwo() string {
	if x != nil {
		return x.AddressLineTwo
	}
	return ""
}

func (x *CheckWeatherRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CheckWeatherRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CheckWeatherRequest) GetZipcode() string {
	if x != nil {
		return x.Zipcode
	}
	return ""
}

type CheckWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentTemp      string `protobuf:"bytes,1,opt,name=current_temp,json=currentTemp,proto3" json:"current_temp,omitempty"`
	HighTemp         string `protobuf:"bytes,2,opt,name=high_temp,json=highTemp,proto3" json:"high_temp,omitempty"`
	LowTemp          string `protobuf:"bytes,3,opt,name=low_temp,json=lowTemp,proto3" json:"low_temp,omitempty"`
	IsCachedResponse bool   `protobuf:"varint,4,opt,name=is_cached_response,json=isCachedResponse,proto3" json:"is_cached_response,omitempty"`
}

func (x *CheckWeatherResponse) Reset() {
	*x = CheckWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckWeatherResponse) ProtoMessage() {}

func (x *CheckWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckWeatherResponse.ProtoReflect.Descriptor instead.
func (*CheckWeatherResponse) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{1}
}

func (x *CheckWeatherResponse) GetCurrentTemp() string {
	if x != nil {
		return x.CurrentTemp
	}
	return ""
}

func (x *CheckWeatherResponse) GetHighTemp() string {
	if x != nil {
		return x.HighTemp
	}
	return ""
}

func (x *CheckWeatherResponse) GetLowTemp() string {
	if x != nil {
		return x.LowTemp
	}
	return ""
}

func (x *CheckWeatherResponse) GetIsCachedResponse() bool {
	if x != nil {
		return x.IsCachedResponse
	}
	return false
}

var File_weather_v1_weather_proto protoreflect.FileDescriptor

var file_weather_v1_weather_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xad, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x10, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6f,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x6e, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x54,
	0x77, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a,
	0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65,
	0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x69, 0x67, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x12,
	0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x73,
	0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x65, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x65, 0x62, 0x62, 0x69, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weather_v1_weather_proto_rawDescOnce sync.Once
	file_weather_v1_weather_proto_rawDescData = file_weather_v1_weather_proto_rawDesc
)

func file_weather_v1_weather_proto_rawDescGZIP() []byte {
	file_weather_v1_weather_proto_rawDescOnce.Do(func() {
		file_weather_v1_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_v1_weather_proto_rawDescData)
	})
	return file_weather_v1_weather_proto_rawDescData
}

var file_weather_v1_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_weather_v1_weather_proto_goTypes = []any{
	(*CheckWeatherRequest)(nil),  // 0: weather.v1.CheckWeatherRequest
	(*CheckWeatherResponse)(nil), // 1: weather.v1.CheckWeatherResponse
}
var file_weather_v1_weather_proto_depIdxs = []int32{
	0, // 0: weather.v1.WeatherService.CheckWeather:input_type -> weather.v1.CheckWeatherRequest
	1, // 1: weather.v1.WeatherService.CheckWeather:output_type -> weather.v1.CheckWeatherResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_weather_v1_weather_proto_init() }
func file_weather_v1_weather_proto_init() {
	if File_weather_v1_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_v1_weather_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CheckWeatherRequest); i {
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
		file_weather_v1_weather_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CheckWeatherResponse); i {
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
			RawDescriptor: file_weather_v1_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_v1_weather_proto_goTypes,
		DependencyIndexes: file_weather_v1_weather_proto_depIdxs,
		MessageInfos:      file_weather_v1_weather_proto_msgTypes,
	}.Build()
	File_weather_v1_weather_proto = out.File
	file_weather_v1_weather_proto_rawDesc = nil
	file_weather_v1_weather_proto_goTypes = nil
	file_weather_v1_weather_proto_depIdxs = nil
}
