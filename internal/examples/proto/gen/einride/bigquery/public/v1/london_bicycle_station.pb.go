// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: einride/bigquery/public/v1/london_bicycle_station.proto

package publicv1

import (
	proto "github.com/golang/protobuf/proto"
	date "google.golang.org/genproto/googleapis/type/date"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Protobuf schema for the BigQuery public table:
//
//  bigquery-public-data.london_bicycles.cycle_stations
type LondonBicycleStation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                   // INTEGER NULLABLE
	Installed  bool    `protobuf:"varint,2,opt,name=installed,proto3" json:"installed,omitempty"`                     // BOOLEAN NULLABLE
	Latitude   float64 `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`                      // FLOAT NULLABLE
	Locked     string  `protobuf:"bytes,4,opt,name=locked,proto3" json:"locked,omitempty"`                            // STRING NULLABLE
	Longitude  float64 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`                    // FLOAT NULLABLE
	Name       string  `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`                                // STRING NULLABLE
	BikesCount int64   `protobuf:"varint,7,opt,name=bikes_count,json=bikesCount,proto3" json:"bikes_count,omitempty"` // INTEGER NULLABLE
	DocksCount int64   `protobuf:"varint,8,opt,name=docks_count,json=docksCount,proto3" json:"docks_count,omitempty"` // INTEGER NULLABLE
	//int64 nbEmptyDocks = 9; // INTEGER NULLABLE
	Temporary    bool       `protobuf:"varint,10,opt,name=temporary,proto3" json:"temporary,omitempty"`                          // BOOLEAN NULLABLE
	TerminalName string     `protobuf:"bytes,11,opt,name=terminal_name,json=terminalName,proto3" json:"terminal_name,omitempty"` // STRING NULLABLE
	InstallDate  *date.Date `protobuf:"bytes,12,opt,name=install_date,json=installDate,proto3" json:"install_date,omitempty"`    // DATE NULLABLE
	RemovalDate  *date.Date `protobuf:"bytes,13,opt,name=removal_date,json=removalDate,proto3" json:"removal_date,omitempty"`    // DATE NULLABLE
}

func (x *LondonBicycleStation) Reset() {
	*x = LondonBicycleStation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_bigquery_public_v1_london_bicycle_station_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LondonBicycleStation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LondonBicycleStation) ProtoMessage() {}

func (x *LondonBicycleStation) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_london_bicycle_station_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LondonBicycleStation.ProtoReflect.Descriptor instead.
func (*LondonBicycleStation) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescGZIP(), []int{0}
}

func (x *LondonBicycleStation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LondonBicycleStation) GetInstalled() bool {
	if x != nil {
		return x.Installed
	}
	return false
}

func (x *LondonBicycleStation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *LondonBicycleStation) GetLocked() string {
	if x != nil {
		return x.Locked
	}
	return ""
}

func (x *LondonBicycleStation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *LondonBicycleStation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LondonBicycleStation) GetBikesCount() int64 {
	if x != nil {
		return x.BikesCount
	}
	return 0
}

func (x *LondonBicycleStation) GetDocksCount() int64 {
	if x != nil {
		return x.DocksCount
	}
	return 0
}

func (x *LondonBicycleStation) GetTemporary() bool {
	if x != nil {
		return x.Temporary
	}
	return false
}

func (x *LondonBicycleStation) GetTerminalName() string {
	if x != nil {
		return x.TerminalName
	}
	return ""
}

func (x *LondonBicycleStation) GetInstallDate() *date.Date {
	if x != nil {
		return x.InstallDate
	}
	return nil
}

func (x *LondonBicycleStation) GetRemovalDate() *date.Date {
	if x != nil {
		return x.RemovalDate
	}
	return nil
}

var File_einride_bigquery_public_v1_london_bicycle_station_proto protoreflect.FileDescriptor

var file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDesc = []byte{
	0x0a, 0x37, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x6e,
	0x64, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x03,
	0x0a, 0x14, 0x4c, 0x6f, 0x6e, 0x64, 0x6f, 0x6e, 0x42, 0x69, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x69,
	0x6b, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x62, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x44, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x6c,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0b,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x42, 0x51, 0x5a, 0x4f, 0x67,
	0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescOnce sync.Once
	file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescData = file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDesc
)

func file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescGZIP() []byte {
	file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescData)
	})
	return file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescData
}

var file_einride_bigquery_public_v1_london_bicycle_station_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_public_v1_london_bicycle_station_proto_goTypes = []interface{}{
	(*LondonBicycleStation)(nil), // 0: einride.bigquery.public.v1.LondonBicycleStation
	(*date.Date)(nil),            // 1: google.type.Date
}
var file_einride_bigquery_public_v1_london_bicycle_station_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.public.v1.LondonBicycleStation.install_date:type_name -> google.type.Date
	1, // 1: einride.bigquery.public.v1.LondonBicycleStation.removal_date:type_name -> google.type.Date
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_einride_bigquery_public_v1_london_bicycle_station_proto_init() }
func file_einride_bigquery_public_v1_london_bicycle_station_proto_init() {
	if File_einride_bigquery_public_v1_london_bicycle_station_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_bigquery_public_v1_london_bicycle_station_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LondonBicycleStation); i {
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
			RawDescriptor: file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_public_v1_london_bicycle_station_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_public_v1_london_bicycle_station_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_public_v1_london_bicycle_station_proto_msgTypes,
	}.Build()
	File_einride_bigquery_public_v1_london_bicycle_station_proto = out.File
	file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDesc = nil
	file_einride_bigquery_public_v1_london_bicycle_station_proto_goTypes = nil
	file_einride_bigquery_public_v1_london_bicycle_station_proto_depIdxs = nil
}
