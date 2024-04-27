// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: test.proto

package test

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

type Base_OnOff int32

const (
	Base_off Base_OnOff = 0
	Base_on  Base_OnOff = 1
)

// Enum value maps for Base_OnOff.
var (
	Base_OnOff_name = map[int32]string{
		0: "off",
		1: "on",
	}
	Base_OnOff_value = map[string]int32{
		"off": 0,
		"on":  1,
	}
)

func (x Base_OnOff) Enum() *Base_OnOff {
	p := new(Base_OnOff)
	*p = x
	return p
}

func (x Base_OnOff) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Base_OnOff) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[0].Descriptor()
}

func (Base_OnOff) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[0]
}

func (x Base_OnOff) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Base_OnOff.Descriptor instead.
func (Base_OnOff) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1, 0}
}

type Child struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kids string `protobuf:"bytes,1,opt,name=kids,proto3" json:"kids,omitempty"`
}

func (x *Child) Reset() {
	*x = Child{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Child) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Child) ProtoMessage() {}

func (x *Child) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Child.ProtoReflect.Descriptor instead.
func (*Child) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *Child) GetKids() string {
	if x != nil {
		return x.Kids
	}
	return ""
}

type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bit            bool              `protobuf:"varint,1,opt,name=bit,proto3" json:"bit,omitempty"`
	Smallint       int32             `protobuf:"varint,2,opt,name=smallint,proto3" json:"smallint,omitempty"`
	Largeint       int64             `protobuf:"varint,3,opt,name=largeint,proto3" json:"largeint,omitempty"`
	Smallfloat     float32           `protobuf:"fixed32,4,opt,name=smallfloat,proto3" json:"smallfloat,omitempty"`
	Largefloat     float64           `protobuf:"fixed64,5,opt,name=largefloat,proto3" json:"largefloat,omitempty"`
	Text           string            `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	Child          *Child            `protobuf:"bytes,7,opt,name=child,proto3" json:"child,omitempty"`
	Onoff          Base_OnOff        `protobuf:"varint,8,opt,name=onoff,proto3,enum=test.Base_OnOff" json:"onoff,omitempty"`
	Optbit         *bool             `protobuf:"varint,9,opt,name=optbit,proto3,oneof" json:"optbit,omitempty"`
	Optsmallint    *int32            `protobuf:"varint,10,opt,name=optsmallint,proto3,oneof" json:"optsmallint,omitempty"`
	Optlargeint    *int64            `protobuf:"varint,11,opt,name=optlargeint,proto3,oneof" json:"optlargeint,omitempty"`
	Optsmallfloat  *float32          `protobuf:"fixed32,12,opt,name=optsmallfloat,proto3,oneof" json:"optsmallfloat,omitempty"`
	Optlargefloat  *float64          `protobuf:"fixed64,13,opt,name=optlargefloat,proto3,oneof" json:"optlargefloat,omitempty"`
	Opttext        *string           `protobuf:"bytes,14,opt,name=opttext,proto3,oneof" json:"opttext,omitempty"`
	Optchild       *Child            `protobuf:"bytes,15,opt,name=optchild,proto3,oneof" json:"optchild,omitempty"`
	Optonoff       *Base_OnOff       `protobuf:"varint,16,opt,name=optonoff,proto3,enum=test.Base_OnOff,oneof" json:"optonoff,omitempty"`
	Listbit        []bool            `protobuf:"varint,17,rep,packed,name=listbit,proto3" json:"listbit,omitempty"`
	Listsmallint   []int32           `protobuf:"varint,18,rep,packed,name=listsmallint,proto3" json:"listsmallint,omitempty"`
	Listlargeint   []int64           `protobuf:"varint,19,rep,packed,name=listlargeint,proto3" json:"listlargeint,omitempty"`
	Listsmallfloat []float32         `protobuf:"fixed32,20,rep,packed,name=listsmallfloat,proto3" json:"listsmallfloat,omitempty"`
	Listlargefloat []float64         `protobuf:"fixed64,21,rep,packed,name=listlargefloat,proto3" json:"listlargefloat,omitempty"`
	Listtext       []string          `protobuf:"bytes,22,rep,name=listtext,proto3" json:"listtext,omitempty"`
	Listchild      []*Child          `protobuf:"bytes,23,rep,name=listchild,proto3" json:"listchild,omitempty"`
	Listonoff      []Base_OnOff      `protobuf:"varint,24,rep,packed,name=listonoff,proto3,enum=test.Base_OnOff" json:"listonoff,omitempty"`
	Maptbit        map[bool]bool     `protobuf:"bytes,25,rep,name=maptbit,proto3" json:"maptbit,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Mapsmallint    map[int32]int32   `protobuf:"bytes,26,rep,name=mapsmallint,proto3" json:"mapsmallint,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Maplargeint    map[int64]int64   `protobuf:"bytes,27,rep,name=maplargeint,proto3" json:"maplargeint,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Maptext        map[string]string `protobuf:"bytes,28,rep,name=maptext,proto3" json:"maptext,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Mapchild       map[string]*Child `protobuf:"bytes,29,rep,name=mapchild,proto3" json:"mapchild,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Base) GetBit() bool {
	if x != nil {
		return x.Bit
	}
	return false
}

func (x *Base) GetSmallint() int32 {
	if x != nil {
		return x.Smallint
	}
	return 0
}

func (x *Base) GetLargeint() int64 {
	if x != nil {
		return x.Largeint
	}
	return 0
}

func (x *Base) GetSmallfloat() float32 {
	if x != nil {
		return x.Smallfloat
	}
	return 0
}

func (x *Base) GetLargefloat() float64 {
	if x != nil {
		return x.Largefloat
	}
	return 0
}

func (x *Base) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Base) GetChild() *Child {
	if x != nil {
		return x.Child
	}
	return nil
}

func (x *Base) GetOnoff() Base_OnOff {
	if x != nil {
		return x.Onoff
	}
	return Base_off
}

func (x *Base) GetOptbit() bool {
	if x != nil && x.Optbit != nil {
		return *x.Optbit
	}
	return false
}

func (x *Base) GetOptsmallint() int32 {
	if x != nil && x.Optsmallint != nil {
		return *x.Optsmallint
	}
	return 0
}

func (x *Base) GetOptlargeint() int64 {
	if x != nil && x.Optlargeint != nil {
		return *x.Optlargeint
	}
	return 0
}

func (x *Base) GetOptsmallfloat() float32 {
	if x != nil && x.Optsmallfloat != nil {
		return *x.Optsmallfloat
	}
	return 0
}

func (x *Base) GetOptlargefloat() float64 {
	if x != nil && x.Optlargefloat != nil {
		return *x.Optlargefloat
	}
	return 0
}

func (x *Base) GetOpttext() string {
	if x != nil && x.Opttext != nil {
		return *x.Opttext
	}
	return ""
}

func (x *Base) GetOptchild() *Child {
	if x != nil {
		return x.Optchild
	}
	return nil
}

func (x *Base) GetOptonoff() Base_OnOff {
	if x != nil && x.Optonoff != nil {
		return *x.Optonoff
	}
	return Base_off
}

func (x *Base) GetListbit() []bool {
	if x != nil {
		return x.Listbit
	}
	return nil
}

func (x *Base) GetListsmallint() []int32 {
	if x != nil {
		return x.Listsmallint
	}
	return nil
}

func (x *Base) GetListlargeint() []int64 {
	if x != nil {
		return x.Listlargeint
	}
	return nil
}

func (x *Base) GetListsmallfloat() []float32 {
	if x != nil {
		return x.Listsmallfloat
	}
	return nil
}

func (x *Base) GetListlargefloat() []float64 {
	if x != nil {
		return x.Listlargefloat
	}
	return nil
}

func (x *Base) GetListtext() []string {
	if x != nil {
		return x.Listtext
	}
	return nil
}

func (x *Base) GetListchild() []*Child {
	if x != nil {
		return x.Listchild
	}
	return nil
}

func (x *Base) GetListonoff() []Base_OnOff {
	if x != nil {
		return x.Listonoff
	}
	return nil
}

func (x *Base) GetMaptbit() map[bool]bool {
	if x != nil {
		return x.Maptbit
	}
	return nil
}

func (x *Base) GetMapsmallint() map[int32]int32 {
	if x != nil {
		return x.Mapsmallint
	}
	return nil
}

func (x *Base) GetMaplargeint() map[int64]int64 {
	if x != nil {
		return x.Maplargeint
	}
	return nil
}

func (x *Base) GetMaptext() map[string]string {
	if x != nil {
		return x.Maptext
	}
	return nil
}

func (x *Base) GetMapchild() map[string]*Child {
	if x != nil {
		return x.Mapchild
	}
	return nil
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x22, 0x1b, 0x0a, 0x05, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x64, 0x73, 0x22,
	0xc4, 0x0c, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x62, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6d,
	0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x6d,
	0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x69,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x69,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x52, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x26, 0x0a, 0x05, 0x6f, 0x6e, 0x6f,
	0x66, 0x66, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x52, 0x05, 0x6f, 0x6e, 0x6f, 0x66,
	0x66, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x62, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x62, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x25,
	0x0a, 0x0b, 0x6f, 0x70, 0x74, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x0b, 0x6f, 0x70, 0x74, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x69,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x6f, 0x70, 0x74, 0x6c, 0x61, 0x72, 0x67,
	0x65, 0x69, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x0b, 0x6f, 0x70,
	0x74, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d,
	0x6f, 0x70, 0x74, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x02, 0x48, 0x03, 0x52, 0x0d, 0x6f, 0x70, 0x74, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x6c, 0x61,
	0x72, 0x67, 0x65, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x48, 0x04,
	0x52, 0x0d, 0x6f, 0x70, 0x74, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x2c, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x48, 0x06, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x31, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x6e,
	0x4f, 0x66, 0x66, 0x48, 0x07, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x88,
	0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x62, 0x69, 0x74, 0x18, 0x11, 0x20,
	0x03, 0x28, 0x08, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x62, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x6c, 0x69, 0x73, 0x74, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x74, 0x18, 0x12, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x69, 0x6e, 0x74,
	0x18, 0x13, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x6c, 0x61, 0x72, 0x67,
	0x65, 0x69, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x6d, 0x61, 0x6c,
	0x6c, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0e, 0x6c, 0x69,
	0x73, 0x74, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x6c, 0x69, 0x73, 0x74, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x15,
	0x20, 0x03, 0x28, 0x01, 0x52, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x16, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x29, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x17, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x2e, 0x0a, 0x09, 0x6c,
	0x69, 0x73, 0x74, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x6e, 0x4f, 0x66, 0x66,
	0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x12, 0x31, 0x0a, 0x07, 0x6d,
	0x61, 0x70, 0x74, 0x62, 0x69, 0x74, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x74, 0x62, 0x69, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x74, 0x62, 0x69, 0x74, 0x12, 0x3d,
	0x0a, 0x0b, 0x6d, 0x61, 0x70, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x74, 0x18, 0x1a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e,
	0x4d, 0x61, 0x70, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0b, 0x6d, 0x61, 0x70, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x74, 0x12, 0x3d, 0x0a,
	0x0b, 0x6d, 0x61, 0x70, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x69, 0x6e, 0x74, 0x18, 0x1b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x4d,
	0x61, 0x70, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x69, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x6d, 0x61, 0x70, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x69, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x07,
	0x6d, 0x61, 0x70, 0x74, 0x65, 0x78, 0x74, 0x18, 0x1c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x74, 0x65, 0x78,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x34, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x1d, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x4d, 0x61,
	0x70, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x61, 0x70,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x1a, 0x3a, 0x0a, 0x0c, 0x4d, 0x61, 0x70, 0x74, 0x62, 0x69, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x69, 0x6e, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x4d, 0x61, 0x70, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x48, 0x0a,
	0x0d, 0x4d, 0x61, 0x70, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x18, 0x0a, 0x05, 0x4f, 0x6e, 0x4f, 0x66, 0x66,
	0x12, 0x07, 0x0a, 0x03, 0x6f, 0x66, 0x66, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x6f, 0x6e, 0x10,
	0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x70, 0x74, 0x62, 0x69, 0x74, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x6f, 0x70, 0x74, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x74, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x6f, 0x70, 0x74, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x69, 0x6e, 0x74, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x6f, 0x70, 0x74, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x6f, 0x70, 0x74, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x70, 0x74, 0x74, 0x65, 0x78, 0x74, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x6f, 0x70, 0x74, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x70,
	0x74, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x42, 0x35, 0x5a, 0x33, 0x73, 0x6f, 0x6e, 0x6f, 0x62, 0x69,
	0x2f, 0x6e, 0x65, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x6f, 0x6e, 0x6f, 0x62, 0x69, 0x2f, 0x75,
	0x74, 0x69, 0x6c, 0x2f, 0x70, 0x62, 0x63, 0x6f, 0x70, 0x79, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d,
	0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_test_proto_goTypes = []interface{}{
	(Base_OnOff)(0), // 0: test.Base.OnOff
	(*Child)(nil),   // 1: test.Child
	(*Base)(nil),    // 2: test.Base
	nil,             // 3: test.Base.MaptbitEntry
	nil,             // 4: test.Base.MapsmallintEntry
	nil,             // 5: test.Base.MaplargeintEntry
	nil,             // 6: test.Base.MaptextEntry
	nil,             // 7: test.Base.MapchildEntry
}
var file_test_proto_depIdxs = []int32{
	1,  // 0: test.Base.child:type_name -> test.Child
	0,  // 1: test.Base.onoff:type_name -> test.Base.OnOff
	1,  // 2: test.Base.optchild:type_name -> test.Child
	0,  // 3: test.Base.optonoff:type_name -> test.Base.OnOff
	1,  // 4: test.Base.listchild:type_name -> test.Child
	0,  // 5: test.Base.listonoff:type_name -> test.Base.OnOff
	3,  // 6: test.Base.maptbit:type_name -> test.Base.MaptbitEntry
	4,  // 7: test.Base.mapsmallint:type_name -> test.Base.MapsmallintEntry
	5,  // 8: test.Base.maplargeint:type_name -> test.Base.MaplargeintEntry
	6,  // 9: test.Base.maptext:type_name -> test.Base.MaptextEntry
	7,  // 10: test.Base.mapchild:type_name -> test.Base.MapchildEntry
	1,  // 11: test.Base.MapchildEntry.value:type_name -> test.Child
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Child); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
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
	file_test_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		EnumInfos:         file_test_proto_enumTypes,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
