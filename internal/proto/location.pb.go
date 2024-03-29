// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: internal/proto/location.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string    `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Coordinates []float64 `protobuf:"fixed64,2,rep,packed,name=Coordinates,proto3" json:"Coordinates,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_internal_proto_location_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Location) GetCoordinates() []float64 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

type LocationObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Location *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *LocationObject) Reset() {
	*x = LocationObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationObject) ProtoMessage() {}

func (x *LocationObject) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationObject.ProtoReflect.Descriptor instead.
func (*LocationObject) Descriptor() ([]byte, []int) {
	return file_internal_proto_location_proto_rawDescGZIP(), []int{1}
}

func (x *LocationObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LocationObject) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type LocationQueryObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Radius    int32   `protobuf:"varint,1,opt,name=radius,proto3" json:"radius,omitempty"`
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *LocationQueryObject) Reset() {
	*x = LocationQueryObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationQueryObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationQueryObject) ProtoMessage() {}

func (x *LocationQueryObject) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationQueryObject.ProtoReflect.Descriptor instead.
func (*LocationQueryObject) Descriptor() ([]byte, []int) {
	return file_internal_proto_location_proto_rawDescGZIP(), []int{2}
}

func (x *LocationQueryObject) GetRadius() int32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *LocationQueryObject) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *LocationQueryObject) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type EventsLists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*LocationObject `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *EventsLists) Reset() {
	*x = EventsLists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_location_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsLists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsLists) ProtoMessage() {}

func (x *EventsLists) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_location_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsLists.ProtoReflect.Descriptor instead.
func (*EventsLists) Descriptor() ([]byte, []int) {
	return file_internal_proto_location_proto_rawDescGZIP(), []int{3}
}

func (x *EventsLists) GetEvents() []*LocationObject {
	if x != nil {
		return x.Events
	}
	return nil
}

type EventId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *EventId) Reset() {
	*x = EventId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_location_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventId) ProtoMessage() {}

func (x *EventId) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_location_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventId.ProtoReflect.Descriptor instead.
func (*EventId) Descriptor() ([]byte, []int) {
	return file_internal_proto_location_proto_rawDescGZIP(), []int{4}
}

func (x *EventId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_internal_proto_location_proto protoreflect.FileDescriptor

var file_internal_proto_location_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x42, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x0e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x67, 0x0a, 0x13, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x22, 0x3c, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73,
	0x12, 0x2d, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x19, 0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x32, 0xcf, 0x01, 0x0a, 0x13, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x42, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x57,
	0x69, 0x74, 0x68, 0x69, 0x6e, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x11, 0x5a, 0x0f,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_location_proto_rawDescOnce sync.Once
	file_internal_proto_location_proto_rawDescData = file_internal_proto_location_proto_rawDesc
)

func file_internal_proto_location_proto_rawDescGZIP() []byte {
	file_internal_proto_location_proto_rawDescOnce.Do(func() {
		file_internal_proto_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_location_proto_rawDescData)
	})
	return file_internal_proto_location_proto_rawDescData
}

var file_internal_proto_location_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_proto_location_proto_goTypes = []interface{}{
	(*Location)(nil),            // 0: proto.Location
	(*LocationObject)(nil),      // 1: proto.LocationObject
	(*LocationQueryObject)(nil), // 2: proto.LocationQueryObject
	(*EventsLists)(nil),         // 3: proto.EventsLists
	(*EventId)(nil),             // 4: proto.EventId
	(*emptypb.Empty)(nil),       // 5: google.protobuf.Empty
}
var file_internal_proto_location_proto_depIdxs = []int32{
	0, // 0: proto.LocationObject.location:type_name -> proto.Location
	1, // 1: proto.EventsLists.events:type_name -> proto.LocationObject
	1, // 2: proto.LocationDataService.LocationData:input_type -> proto.LocationObject
	2, // 3: proto.LocationDataService.FindEventsWithin:input_type -> proto.LocationQueryObject
	4, // 4: proto.LocationDataService.DeleteEvent:input_type -> proto.EventId
	5, // 5: proto.LocationDataService.LocationData:output_type -> google.protobuf.Empty
	3, // 6: proto.LocationDataService.FindEventsWithin:output_type -> proto.EventsLists
	5, // 7: proto.LocationDataService.DeleteEvent:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_proto_location_proto_init() }
func file_internal_proto_location_proto_init() {
	if File_internal_proto_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_internal_proto_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationObject); i {
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
		file_internal_proto_location_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationQueryObject); i {
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
		file_internal_proto_location_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsLists); i {
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
		file_internal_proto_location_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventId); i {
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
			RawDescriptor: file_internal_proto_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_location_proto_goTypes,
		DependencyIndexes: file_internal_proto_location_proto_depIdxs,
		MessageInfos:      file_internal_proto_location_proto_msgTypes,
	}.Build()
	File_internal_proto_location_proto = out.File
	file_internal_proto_location_proto_rawDesc = nil
	file_internal_proto_location_proto_goTypes = nil
	file_internal_proto_location_proto_depIdxs = nil
}
