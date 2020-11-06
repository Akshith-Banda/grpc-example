// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: service.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Thermostat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CurrentTemp int64  `protobuf:"varint,3,opt,name=currentTemp,proto3" json:"currentTemp,omitempty"`
	TempDsply   bool   `protobuf:"varint,4,opt,name=tempDsply,proto3" json:"tempDsply,omitempty"`
	GroupId     int64  `protobuf:"varint,5,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *Thermostat) Reset() {
	*x = Thermostat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thermostat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thermostat) ProtoMessage() {}

func (x *Thermostat) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thermostat.ProtoReflect.Descriptor instead.
func (*Thermostat) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *Thermostat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Thermostat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Thermostat) GetCurrentTemp() int64 {
	if x != nil {
		return x.CurrentTemp
	}
	return 0
}

func (x *Thermostat) GetTempDsply() bool {
	if x != nil {
		return x.TempDsply
	}
	return false
}

func (x *Thermostat) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type GetThermostat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThermostatGet *Thermostat `protobuf:"bytes,1,opt,name=thermostat_get,json=thermostatGet,proto3" json:"thermostat_get,omitempty"`
}

func (x *GetThermostat) Reset() {
	*x = GetThermostat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThermostat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThermostat) ProtoMessage() {}

func (x *GetThermostat) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThermostat.ProtoReflect.Descriptor instead.
func (*GetThermostat) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetThermostat) GetThermostatGet() *Thermostat {
	if x != nil {
		return x.ThermostatGet
	}
	return nil
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ThermoGrp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThermoName string `protobuf:"bytes,1,opt,name=thermo_name,json=thermoName,proto3" json:"thermo_name,omitempty"`
	GroupName  string `protobuf:"bytes,2,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
}

func (x *ThermoGrp) Reset() {
	*x = ThermoGrp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThermoGrp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThermoGrp) ProtoMessage() {}

func (x *ThermoGrp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThermoGrp.ProtoReflect.Descriptor instead.
func (*ThermoGrp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *ThermoGrp) GetThermoName() string {
	if x != nil {
		return x.ThermoName
	}
	return ""
}

func (x *ThermoGrp) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x0a, 0x54, 0x68, 0x65, 0x72, 0x6d,
	0x6f, 0x73, 0x74, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x65, 0x6d, 0x70, 0x44, 0x73, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x74, 0x65, 0x6d, 0x70, 0x44, 0x73, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x68, 0x65, 0x72, 0x6d,
	0x6f, 0x73, 0x74, 0x61, 0x74, 0x12, 0x38, 0x0a, 0x0e, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x6f, 0x73,
	0x74, 0x61, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x68, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x74, 0x61, 0x74,
	0x52, 0x0d, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x74, 0x61, 0x74, 0x47, 0x65, 0x74, 0x22,
	0x1b, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x09,
	0x54, 0x68, 0x65, 0x72, 0x6d, 0x6f, 0x47, 0x72, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x68, 0x65,
	0x72, 0x6d, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x68, 0x65, 0x72, 0x6d, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x80, 0x02,
	0x0a, 0x05, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x5f, 0x54,
	0x68, 0x65, 0x72, 0x6d, 0x6f, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x68,
	0x65, 0x72, 0x6d, 0x6f, 0x73, 0x74, 0x61, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x41, 0x64, 0x64,
	0x5f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x5f, 0x54, 0x68, 0x65,
	0x72, 0x6d, 0x6f, 0x5f, 0x47, 0x72, 0x70, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x68, 0x65, 0x72, 0x6d, 0x6f, 0x47, 0x72, 0x70, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x5f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x5f, 0x54, 0x68, 0x65, 0x72, 0x6d, 0x6f, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x68, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x74, 0x61, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x74, 0x61, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_proto_goTypes = []interface{}{
	(*Thermostat)(nil),    // 0: proto.Thermostat
	(*GetThermostat)(nil), // 1: proto.GetThermostat
	(*Group)(nil),         // 2: proto.Group
	(*ThermoGrp)(nil),     // 3: proto.ThermoGrp
	(*Response)(nil),      // 4: proto.Response
}
var file_service_proto_depIdxs = []int32{
	0, // 0: proto.GetThermostat.thermostat_get:type_name -> proto.Thermostat
	0, // 1: proto.Trial.Add_Thermo:input_type -> proto.Thermostat
	2, // 2: proto.Trial.Add_Group:input_type -> proto.Group
	3, // 3: proto.Trial.Add_Thermo_Grp:input_type -> proto.ThermoGrp
	2, // 4: proto.Trial.Delete_Group:input_type -> proto.Group
	0, // 5: proto.Trial.Get_Thermo:input_type -> proto.Thermostat
	4, // 6: proto.Trial.Add_Thermo:output_type -> proto.Response
	4, // 7: proto.Trial.Add_Group:output_type -> proto.Response
	4, // 8: proto.Trial.Add_Thermo_Grp:output_type -> proto.Response
	4, // 9: proto.Trial.Delete_Group:output_type -> proto.Response
	1, // 10: proto.Trial.Get_Thermo:output_type -> proto.GetThermostat
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Thermostat); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThermostat); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThermoGrp); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TrialClient is the client API for Trial service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrialClient interface {
	Add_Thermo(ctx context.Context, in *Thermostat, opts ...grpc.CallOption) (*Response, error)
	Add_Group(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Response, error)
	Add_Thermo_Grp(ctx context.Context, in *ThermoGrp, opts ...grpc.CallOption) (*Response, error)
	Delete_Group(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Response, error)
	Get_Thermo(ctx context.Context, in *Thermostat, opts ...grpc.CallOption) (*GetThermostat, error)
}

type trialClient struct {
	cc grpc.ClientConnInterface
}

func NewTrialClient(cc grpc.ClientConnInterface) TrialClient {
	return &trialClient{cc}
}

func (c *trialClient) Add_Thermo(ctx context.Context, in *Thermostat, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Trial/Add_Thermo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trialClient) Add_Group(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Trial/Add_Group", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trialClient) Add_Thermo_Grp(ctx context.Context, in *ThermoGrp, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Trial/Add_Thermo_Grp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trialClient) Delete_Group(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Trial/Delete_Group", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trialClient) Get_Thermo(ctx context.Context, in *Thermostat, opts ...grpc.CallOption) (*GetThermostat, error) {
	out := new(GetThermostat)
	err := c.cc.Invoke(ctx, "/proto.Trial/Get_Thermo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrialServer is the server API for Trial service.
type TrialServer interface {
	Add_Thermo(context.Context, *Thermostat) (*Response, error)
	Add_Group(context.Context, *Group) (*Response, error)
	Add_Thermo_Grp(context.Context, *ThermoGrp) (*Response, error)
	Delete_Group(context.Context, *Group) (*Response, error)
	Get_Thermo(context.Context, *Thermostat) (*GetThermostat, error)
}

// UnimplementedTrialServer can be embedded to have forward compatible implementations.
type UnimplementedTrialServer struct {
}

func (*UnimplementedTrialServer) Add_Thermo(context.Context, *Thermostat) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add_Thermo not implemented")
}
func (*UnimplementedTrialServer) Add_Group(context.Context, *Group) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add_Group not implemented")
}
func (*UnimplementedTrialServer) Add_Thermo_Grp(context.Context, *ThermoGrp) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add_Thermo_Grp not implemented")
}
func (*UnimplementedTrialServer) Delete_Group(context.Context, *Group) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete_Group not implemented")
}
func (*UnimplementedTrialServer) Get_Thermo(context.Context, *Thermostat) (*GetThermostat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get_Thermo not implemented")
}

func RegisterTrialServer(s *grpc.Server, srv TrialServer) {
	s.RegisterService(&_Trial_serviceDesc, srv)
}

func _Trial_Add_Thermo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Thermostat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrialServer).Add_Thermo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Trial/Add_Thermo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrialServer).Add_Thermo(ctx, req.(*Thermostat))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trial_Add_Group_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrialServer).Add_Group(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Trial/Add_Group",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrialServer).Add_Group(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trial_Add_Thermo_Grp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThermoGrp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrialServer).Add_Thermo_Grp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Trial/Add_Thermo_Grp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrialServer).Add_Thermo_Grp(ctx, req.(*ThermoGrp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trial_Delete_Group_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrialServer).Delete_Group(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Trial/Delete_Group",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrialServer).Delete_Group(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trial_Get_Thermo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Thermostat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrialServer).Get_Thermo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Trial/Get_Thermo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrialServer).Get_Thermo(ctx, req.(*Thermostat))
	}
	return interceptor(ctx, in, info, handler)
}

var _Trial_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Trial",
	HandlerType: (*TrialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add_Thermo",
			Handler:    _Trial_Add_Thermo_Handler,
		},
		{
			MethodName: "Add_Group",
			Handler:    _Trial_Add_Group_Handler,
		},
		{
			MethodName: "Add_Thermo_Grp",
			Handler:    _Trial_Add_Thermo_Grp_Handler,
		},
		{
			MethodName: "Delete_Group",
			Handler:    _Trial_Delete_Group_Handler,
		},
		{
			MethodName: "Get_Thermo",
			Handler:    _Trial_Get_Thermo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}