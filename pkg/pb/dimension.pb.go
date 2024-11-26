// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: sro/gameserver/dimension.proto

package pb

import (
	pb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateDimensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version  string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	MapIds   []string `protobuf:"bytes,3,rep,name=map_ids,json=mapIds,proto3" json:"map_ids,omitempty"`
	Location string   `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CreateDimensionRequest) Reset() {
	*x = CreateDimensionRequest{}
	mi := &file_sro_gameserver_dimension_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDimensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDimensionRequest) ProtoMessage() {}

func (x *CreateDimensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_dimension_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDimensionRequest.ProtoReflect.Descriptor instead.
func (*CreateDimensionRequest) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_dimension_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDimensionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDimensionRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CreateDimensionRequest) GetMapIds() []string {
	if x != nil {
		return x.MapIds
	}
	return nil
}

func (x *CreateDimensionRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type DuplicateDimensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId string `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DuplicateDimensionRequest) Reset() {
	*x = DuplicateDimensionRequest{}
	mi := &file_sro_gameserver_dimension_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DuplicateDimensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DuplicateDimensionRequest) ProtoMessage() {}

func (x *DuplicateDimensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_dimension_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DuplicateDimensionRequest.ProtoReflect.Descriptor instead.
func (*DuplicateDimensionRequest) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_dimension_proto_rawDescGZIP(), []int{1}
}

func (x *DuplicateDimensionRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *DuplicateDimensionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EditDimensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId string `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// Types that are assignable to OptionalName:
	//
	//	*EditDimensionRequest_Name
	OptionalName isEditDimensionRequest_OptionalName `protobuf_oneof:"optional_name"`
	// Types that are assignable to OptionalVersion:
	//
	//	*EditDimensionRequest_Version
	OptionalVersion isEditDimensionRequest_OptionalVersion `protobuf_oneof:"optional_version"`
	EditMaps        bool                                   `protobuf:"varint,4,opt,name=edit_maps,json=editMaps,proto3" json:"edit_maps,omitempty"`
	MapIds          []string                               `protobuf:"bytes,5,rep,name=map_ids,json=mapIds,proto3" json:"map_ids,omitempty"`
	// Types that are assignable to OptionalLocation:
	//
	//	*EditDimensionRequest_Location
	OptionalLocation isEditDimensionRequest_OptionalLocation `protobuf_oneof:"optional_location"`
}

func (x *EditDimensionRequest) Reset() {
	*x = EditDimensionRequest{}
	mi := &file_sro_gameserver_dimension_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditDimensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditDimensionRequest) ProtoMessage() {}

func (x *EditDimensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_dimension_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditDimensionRequest.ProtoReflect.Descriptor instead.
func (*EditDimensionRequest) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_dimension_proto_rawDescGZIP(), []int{2}
}

func (x *EditDimensionRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (m *EditDimensionRequest) GetOptionalName() isEditDimensionRequest_OptionalName {
	if m != nil {
		return m.OptionalName
	}
	return nil
}

func (x *EditDimensionRequest) GetName() string {
	if x, ok := x.GetOptionalName().(*EditDimensionRequest_Name); ok {
		return x.Name
	}
	return ""
}

func (m *EditDimensionRequest) GetOptionalVersion() isEditDimensionRequest_OptionalVersion {
	if m != nil {
		return m.OptionalVersion
	}
	return nil
}

func (x *EditDimensionRequest) GetVersion() string {
	if x, ok := x.GetOptionalVersion().(*EditDimensionRequest_Version); ok {
		return x.Version
	}
	return ""
}

func (x *EditDimensionRequest) GetEditMaps() bool {
	if x != nil {
		return x.EditMaps
	}
	return false
}

func (x *EditDimensionRequest) GetMapIds() []string {
	if x != nil {
		return x.MapIds
	}
	return nil
}

func (m *EditDimensionRequest) GetOptionalLocation() isEditDimensionRequest_OptionalLocation {
	if m != nil {
		return m.OptionalLocation
	}
	return nil
}

func (x *EditDimensionRequest) GetLocation() string {
	if x, ok := x.GetOptionalLocation().(*EditDimensionRequest_Location); ok {
		return x.Location
	}
	return ""
}

type isEditDimensionRequest_OptionalName interface {
	isEditDimensionRequest_OptionalName()
}

type EditDimensionRequest_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*EditDimensionRequest_Name) isEditDimensionRequest_OptionalName() {}

type isEditDimensionRequest_OptionalVersion interface {
	isEditDimensionRequest_OptionalVersion()
}

type EditDimensionRequest_Version struct {
	Version string `protobuf:"bytes,3,opt,name=version,proto3,oneof"`
}

func (*EditDimensionRequest_Version) isEditDimensionRequest_OptionalVersion() {}

type isEditDimensionRequest_OptionalLocation interface {
	isEditDimensionRequest_OptionalLocation()
}

type EditDimensionRequest_Location struct {
	Location string `protobuf:"bytes,8,opt,name=location,proto3,oneof"`
}

func (*EditDimensionRequest_Location) isEditDimensionRequest_OptionalLocation() {}

type Dimension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique id for the dimension
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique name of the dimension
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Server version used for each server instance
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// All maps the realm should have available
	MapIds []string `protobuf:"bytes,4,rep,name=map_ids,json=mapIds,proto3" json:"map_ids,omitempty"`
	// Physical server location
	Location string `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Dimension) Reset() {
	*x = Dimension{}
	mi := &file_sro_gameserver_dimension_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dimension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dimension) ProtoMessage() {}

func (x *Dimension) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_dimension_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dimension.ProtoReflect.Descriptor instead.
func (*Dimension) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_dimension_proto_rawDescGZIP(), []int{3}
}

func (x *Dimension) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Dimension) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dimension) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Dimension) GetMapIds() []string {
	if x != nil {
		return x.MapIds
	}
	return nil
}

func (x *Dimension) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type Dimensions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dimensions []*Dimension `protobuf:"bytes,1,rep,name=dimensions,proto3" json:"dimensions,omitempty"`
}

func (x *Dimensions) Reset() {
	*x = Dimensions{}
	mi := &file_sro_gameserver_dimension_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dimensions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dimensions) ProtoMessage() {}

func (x *Dimensions) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_dimension_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dimensions.ProtoReflect.Descriptor instead.
func (*Dimensions) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_dimension_proto_rawDescGZIP(), []int{4}
}

func (x *Dimensions) GetDimensions() []*Dimension {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

var File_sro_gameserver_dimension_proto protoreflect.FileDescriptor

var file_sro_gameserver_dimension_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x72, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x72, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x19, 0x44,
	0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf3, 0x01, 0x0a, 0x14, 0x45, 0x64,
	0x69, 0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x64, 0x69, 0x74, 0x4d, 0x61, 0x70, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x13, 0x0a, 0x11, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x7e, 0x0a, 0x09, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61,
	0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x70,
	0x49, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x47, 0x0a, 0x0a, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xa1, 0x05, 0x0a, 0x10, 0x44, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e,
	0x73, 0x72, 0x6f, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x73,
	0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12,
	0x16, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1a, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x16, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6f, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x12, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x73,
	0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x75,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x75,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7a, 0x0a, 0x0d, 0x45, 0x64, 0x69, 0x74, 0x44, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x44, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x3a, 0x01, 0x2a, 0x1a, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x58, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x6d, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2f, 0x64, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sro_gameserver_dimension_proto_rawDescOnce sync.Once
	file_sro_gameserver_dimension_proto_rawDescData = file_sro_gameserver_dimension_proto_rawDesc
)

func file_sro_gameserver_dimension_proto_rawDescGZIP() []byte {
	file_sro_gameserver_dimension_proto_rawDescOnce.Do(func() {
		file_sro_gameserver_dimension_proto_rawDescData = protoimpl.X.CompressGZIP(file_sro_gameserver_dimension_proto_rawDescData)
	})
	return file_sro_gameserver_dimension_proto_rawDescData
}

var file_sro_gameserver_dimension_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sro_gameserver_dimension_proto_goTypes = []any{
	(*CreateDimensionRequest)(nil),    // 0: sro.gameserver.CreateDimensionRequest
	(*DuplicateDimensionRequest)(nil), // 1: sro.gameserver.DuplicateDimensionRequest
	(*EditDimensionRequest)(nil),      // 2: sro.gameserver.EditDimensionRequest
	(*Dimension)(nil),                 // 3: sro.gameserver.Dimension
	(*Dimensions)(nil),                // 4: sro.gameserver.Dimensions
	(*pb.TargetId)(nil),               // 5: sro.TargetId
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
}
var file_sro_gameserver_dimension_proto_depIdxs = []int32{
	3, // 0: sro.gameserver.Dimensions.dimensions:type_name -> sro.gameserver.Dimension
	5, // 1: sro.gameserver.DimensionService.GetDimension:input_type -> sro.TargetId
	6, // 2: sro.gameserver.DimensionService.GetDimensions:input_type -> google.protobuf.Empty
	0, // 3: sro.gameserver.DimensionService.CreateDimension:input_type -> sro.gameserver.CreateDimensionRequest
	1, // 4: sro.gameserver.DimensionService.DuplicateDimension:input_type -> sro.gameserver.DuplicateDimensionRequest
	2, // 5: sro.gameserver.DimensionService.EditDimension:input_type -> sro.gameserver.EditDimensionRequest
	5, // 6: sro.gameserver.DimensionService.DeleteDimension:input_type -> sro.TargetId
	3, // 7: sro.gameserver.DimensionService.GetDimension:output_type -> sro.gameserver.Dimension
	4, // 8: sro.gameserver.DimensionService.GetDimensions:output_type -> sro.gameserver.Dimensions
	3, // 9: sro.gameserver.DimensionService.CreateDimension:output_type -> sro.gameserver.Dimension
	3, // 10: sro.gameserver.DimensionService.DuplicateDimension:output_type -> sro.gameserver.Dimension
	3, // 11: sro.gameserver.DimensionService.EditDimension:output_type -> sro.gameserver.Dimension
	6, // 12: sro.gameserver.DimensionService.DeleteDimension:output_type -> google.protobuf.Empty
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sro_gameserver_dimension_proto_init() }
func file_sro_gameserver_dimension_proto_init() {
	if File_sro_gameserver_dimension_proto != nil {
		return
	}
	file_sro_gameserver_dimension_proto_msgTypes[2].OneofWrappers = []any{
		(*EditDimensionRequest_Name)(nil),
		(*EditDimensionRequest_Version)(nil),
		(*EditDimensionRequest_Location)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sro_gameserver_dimension_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sro_gameserver_dimension_proto_goTypes,
		DependencyIndexes: file_sro_gameserver_dimension_proto_depIdxs,
		MessageInfos:      file_sro_gameserver_dimension_proto_msgTypes,
	}.Build()
	File_sro_gameserver_dimension_proto = out.File
	file_sro_gameserver_dimension_proto_rawDesc = nil
	file_sro_gameserver_dimension_proto_goTypes = nil
	file_sro_gameserver_dimension_proto_depIdxs = nil
}
