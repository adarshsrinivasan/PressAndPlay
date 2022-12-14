// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: events.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventModelList []*EventModel `protobuf:"bytes,1,rep,name=eventModelList,proto3" json:"eventModelList,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

func (x *EventResponse) GetEventModelList() []*EventModel {
	if x != nil {
		return x.EventModelList
	}
	return nil
}

type EventModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId           string                 `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ManagerId        string                 `protobuf:"bytes,3,opt,name=managerId,proto3" json:"managerId,omitempty"`
	CourtId          string                 `protobuf:"bytes,4,opt,name=courtId,proto3" json:"courtId,omitempty"`
	SlotId           string                 `protobuf:"bytes,5,opt,name=slotId,proto3" json:"slotId,omitempty"`
	BookingTimestamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=bookingTimestamp,proto3" json:"bookingTimestamp,omitempty"`
	TimeStartHHMM    int32                  `protobuf:"varint,7,opt,name=timeStartHHMM,proto3" json:"timeStartHHMM,omitempty"`
	TimeEndHHMM      int32                  `protobuf:"varint,8,opt,name=timeEndHHMM,proto3" json:"timeEndHHMM,omitempty"`
	Notified         bool                   `protobuf:"varint,9,opt,name=notified,proto3" json:"notified,omitempty"`
}

func (x *EventModel) Reset() {
	*x = EventModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventModel) ProtoMessage() {}

func (x *EventModel) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventModel.ProtoReflect.Descriptor instead.
func (*EventModel) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

func (x *EventModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventModel) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EventModel) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

func (x *EventModel) GetCourtId() string {
	if x != nil {
		return x.CourtId
	}
	return ""
}

func (x *EventModel) GetSlotId() string {
	if x != nil {
		return x.SlotId
	}
	return ""
}

func (x *EventModel) GetBookingTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.BookingTimestamp
	}
	return nil
}

func (x *EventModel) GetTimeStartHHMM() int32 {
	if x != nil {
		return x.TimeStartHHMM
	}
	return 0
}

func (x *EventModel) GetTimeEndHHMM() int32 {
	if x != nil {
		return x.TimeEndHHMM
	}
	return 0
}

func (x *EventModel) GetNotified() bool {
	if x != nil {
		return x.Notified
	}
	return false
}

var File_events_proto protoreflect.FileDescriptor

var file_events_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x44, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x33, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xb0, 0x02, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x72, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x72, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x10,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x10, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x48, 0x48, 0x4d, 0x4d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x69, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x48, 0x4d, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69,
	0x6d, 0x65, 0x45, 0x6e, 0x64, 0x48, 0x48, 0x4d, 0x4d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x48, 0x48, 0x4d, 0x4d, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x32, 0x46, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x3c, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x49,
	0x64, 0x12, 0x0b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x0e,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x64, 0x61, 0x72, 0x73, 0x68, 0x73, 0x72, 0x69, 0x6e, 0x69, 0x76, 0x61, 0x73, 0x61, 0x6e, 0x2f,
	0x50, 0x72, 0x65, 0x73, 0x73, 0x41, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x2f, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_proto_rawDescOnce sync.Once
	file_events_proto_rawDescData = file_events_proto_rawDesc
)

func file_events_proto_rawDescGZIP() []byte {
	file_events_proto_rawDescOnce.Do(func() {
		file_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_proto_rawDescData)
	})
	return file_events_proto_rawDescData
}

var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_events_proto_goTypes = []interface{}{
	(*EventResponse)(nil),         // 0: EventResponse
	(*EventModel)(nil),            // 1: EventModel
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_events_proto_depIdxs = []int32{
	1, // 0: EventResponse.eventModelList:type_name -> EventModel
	2, // 1: EventModel.bookingTimestamp:type_name -> google.protobuf.Timestamp
	1, // 2: Events.GetEventsByUserIdAndCourtId:input_type -> EventModel
	0, // 3: Events.GetEventsByUserIdAndCourtId:output_type -> EventResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
	if File_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
		file_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventModel); i {
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
			RawDescriptor: file_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_events_proto_goTypes,
		DependencyIndexes: file_events_proto_depIdxs,
		MessageInfos:      file_events_proto_msgTypes,
	}.Build()
	File_events_proto = out.File
	file_events_proto_rawDesc = nil
	file_events_proto_goTypes = nil
	file_events_proto_depIdxs = nil
}
