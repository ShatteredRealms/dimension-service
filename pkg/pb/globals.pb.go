// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: sro/globals.proto

package pb

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

type UserTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//
	//	*UserTarget_Id
	//	*UserTarget_Username
	Target isUserTarget_Target `protobuf_oneof:"target"`
}

func (x *UserTarget) Reset() {
	*x = UserTarget{}
	mi := &file_sro_globals_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTarget) ProtoMessage() {}

func (x *UserTarget) ProtoReflect() protoreflect.Message {
	mi := &file_sro_globals_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTarget.ProtoReflect.Descriptor instead.
func (*UserTarget) Descriptor() ([]byte, []int) {
	return file_sro_globals_proto_rawDescGZIP(), []int{0}
}

func (m *UserTarget) GetTarget() isUserTarget_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *UserTarget) GetId() string {
	if x, ok := x.GetTarget().(*UserTarget_Id); ok {
		return x.Id
	}
	return ""
}

func (x *UserTarget) GetUsername() string {
	if x, ok := x.GetTarget().(*UserTarget_Username); ok {
		return x.Username
	}
	return ""
}

type isUserTarget_Target interface {
	isUserTarget_Target()
}

type UserTarget_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type UserTarget_Username struct {
	Username string `protobuf:"bytes,2,opt,name=username,proto3,oneof"`
}

func (*UserTarget_Id) isUserTarget_Target() {}

func (*UserTarget_Username) isUserTarget_Target() {}

type CharacterTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//
	//	*CharacterTarget_Id
	//	*CharacterTarget_Name
	Target isCharacterTarget_Target `protobuf_oneof:"target"`
}

func (x *CharacterTarget) Reset() {
	*x = CharacterTarget{}
	mi := &file_sro_globals_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CharacterTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharacterTarget) ProtoMessage() {}

func (x *CharacterTarget) ProtoReflect() protoreflect.Message {
	mi := &file_sro_globals_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharacterTarget.ProtoReflect.Descriptor instead.
func (*CharacterTarget) Descriptor() ([]byte, []int) {
	return file_sro_globals_proto_rawDescGZIP(), []int{1}
}

func (m *CharacterTarget) GetTarget() isCharacterTarget_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *CharacterTarget) GetId() string {
	if x, ok := x.GetTarget().(*CharacterTarget_Id); ok {
		return x.Id
	}
	return ""
}

func (x *CharacterTarget) GetName() string {
	if x, ok := x.GetTarget().(*CharacterTarget_Name); ok {
		return x.Name
	}
	return ""
}

type isCharacterTarget_Target interface {
	isCharacterTarget_Target()
}

type CharacterTarget_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type CharacterTarget_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*CharacterTarget_Id) isCharacterTarget_Target() {}

func (*CharacterTarget_Name) isCharacterTarget_Target() {}

type DimensionTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//
	//	*DimensionTarget_Id
	//	*DimensionTarget_Name
	Target isDimensionTarget_Target `protobuf_oneof:"target"`
}

func (x *DimensionTarget) Reset() {
	*x = DimensionTarget{}
	mi := &file_sro_globals_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DimensionTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DimensionTarget) ProtoMessage() {}

func (x *DimensionTarget) ProtoReflect() protoreflect.Message {
	mi := &file_sro_globals_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DimensionTarget.ProtoReflect.Descriptor instead.
func (*DimensionTarget) Descriptor() ([]byte, []int) {
	return file_sro_globals_proto_rawDescGZIP(), []int{2}
}

func (m *DimensionTarget) GetTarget() isDimensionTarget_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *DimensionTarget) GetId() string {
	if x, ok := x.GetTarget().(*DimensionTarget_Id); ok {
		return x.Id
	}
	return ""
}

func (x *DimensionTarget) GetName() string {
	if x, ok := x.GetTarget().(*DimensionTarget_Name); ok {
		return x.Name
	}
	return ""
}

type isDimensionTarget_Target interface {
	isDimensionTarget_Target()
}

type DimensionTarget_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type DimensionTarget_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*DimensionTarget_Id) isDimensionTarget_Target() {}

func (*DimensionTarget_Name) isDimensionTarget_Target() {}

type MapTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//
	//	*MapTarget_Id
	//	*MapTarget_Name
	Target isMapTarget_Target `protobuf_oneof:"target"`
}

func (x *MapTarget) Reset() {
	*x = MapTarget{}
	mi := &file_sro_globals_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapTarget) ProtoMessage() {}

func (x *MapTarget) ProtoReflect() protoreflect.Message {
	mi := &file_sro_globals_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapTarget.ProtoReflect.Descriptor instead.
func (*MapTarget) Descriptor() ([]byte, []int) {
	return file_sro_globals_proto_rawDescGZIP(), []int{3}
}

func (m *MapTarget) GetTarget() isMapTarget_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *MapTarget) GetId() string {
	if x, ok := x.GetTarget().(*MapTarget_Id); ok {
		return x.Id
	}
	return ""
}

func (x *MapTarget) GetName() string {
	if x, ok := x.GetTarget().(*MapTarget_Name); ok {
		return x.Name
	}
	return ""
}

type isMapTarget_Target interface {
	isMapTarget_Target()
}

type MapTarget_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type MapTarget_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*MapTarget_Id) isMapTarget_Target() {}

func (*MapTarget_Name) isMapTarget_Target() {}

type ChatChannelTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChatChannelTarget) Reset() {
	*x = ChatChannelTarget{}
	mi := &file_sro_globals_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatChannelTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatChannelTarget) ProtoMessage() {}

func (x *ChatChannelTarget) ProtoReflect() protoreflect.Message {
	mi := &file_sro_globals_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatChannelTarget.ProtoReflect.Descriptor instead.
func (*ChatChannelTarget) Descriptor() ([]byte, []int) {
	return file_sro_globals_proto_rawDescGZIP(), []int{4}
}

func (x *ChatChannelTarget) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	World string  `protobuf:"bytes,1,opt,name=world,proto3" json:"world,omitempty"`
	X     float32 `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y     float32 `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
	Z     float32 `protobuf:"fixed32,4,opt,name=z,proto3" json:"z,omitempty"`
	Roll  float32 `protobuf:"fixed32,5,opt,name=roll,proto3" json:"roll,omitempty"`
	Pitch float32 `protobuf:"fixed32,6,opt,name=pitch,proto3" json:"pitch,omitempty"`
	Yaw   float32 `protobuf:"fixed32,7,opt,name=yaw,proto3" json:"yaw,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	mi := &file_sro_globals_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_sro_globals_proto_msgTypes[5]
	if x != nil {
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
	return file_sro_globals_proto_rawDescGZIP(), []int{5}
}

func (x *Location) GetWorld() string {
	if x != nil {
		return x.World
	}
	return ""
}

func (x *Location) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Location) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Location) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Location) GetRoll() float32 {
	if x != nil {
		return x.Roll
	}
	return 0
}

func (x *Location) GetPitch() float32 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *Location) GetYaw() float32 {
	if x != nil {
		return x.Yaw
	}
	return 0
}

var File_sro_globals_proto protoreflect.FileDescriptor

var file_sro_globals_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x72, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x72, 0x6f, 0x22, 0x46, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x22, 0x43, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x43, 0x0a, 0x0f, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x3d, 0x0a, 0x09, 0x4d, 0x61,
	0x70, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x43, 0x68, 0x61,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x86,
	0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12,
	0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a,
	0x01, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x69, 0x74, 0x63, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x79, 0x61, 0x77, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x03, 0x79, 0x61, 0x77, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sro_globals_proto_rawDescOnce sync.Once
	file_sro_globals_proto_rawDescData = file_sro_globals_proto_rawDesc
)

func file_sro_globals_proto_rawDescGZIP() []byte {
	file_sro_globals_proto_rawDescOnce.Do(func() {
		file_sro_globals_proto_rawDescData = protoimpl.X.CompressGZIP(file_sro_globals_proto_rawDescData)
	})
	return file_sro_globals_proto_rawDescData
}

var file_sro_globals_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sro_globals_proto_goTypes = []any{
	(*UserTarget)(nil),        // 0: sro.UserTarget
	(*CharacterTarget)(nil),   // 1: sro.CharacterTarget
	(*DimensionTarget)(nil),   // 2: sro.DimensionTarget
	(*MapTarget)(nil),         // 3: sro.MapTarget
	(*ChatChannelTarget)(nil), // 4: sro.ChatChannelTarget
	(*Location)(nil),          // 5: sro.Location
}
var file_sro_globals_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sro_globals_proto_init() }
func file_sro_globals_proto_init() {
	if File_sro_globals_proto != nil {
		return
	}
	file_sro_globals_proto_msgTypes[0].OneofWrappers = []any{
		(*UserTarget_Id)(nil),
		(*UserTarget_Username)(nil),
	}
	file_sro_globals_proto_msgTypes[1].OneofWrappers = []any{
		(*CharacterTarget_Id)(nil),
		(*CharacterTarget_Name)(nil),
	}
	file_sro_globals_proto_msgTypes[2].OneofWrappers = []any{
		(*DimensionTarget_Id)(nil),
		(*DimensionTarget_Name)(nil),
	}
	file_sro_globals_proto_msgTypes[3].OneofWrappers = []any{
		(*MapTarget_Id)(nil),
		(*MapTarget_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sro_globals_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sro_globals_proto_goTypes,
		DependencyIndexes: file_sro_globals_proto_depIdxs,
		MessageInfos:      file_sro_globals_proto_msgTypes,
	}.Build()
	File_sro_globals_proto = out.File
	file_sro_globals_proto_rawDesc = nil
	file_sro_globals_proto_goTypes = nil
	file_sro_globals_proto_depIdxs = nil
}
