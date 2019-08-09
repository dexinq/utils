// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/appmanager/appmanager.proto

package appmanager

import (
	fmt "fmt"
	_ "github.com/dexinq/utils/proto/global"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AppEntity struct {
	AppName              string   `protobuf:"bytes,1,opt,name=AppName,proto3" json:"AppName,omitempty"`
	CreateAt             int64    `protobuf:"varint,2,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	UpdateAt             int64    `protobuf:"varint,3,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	AppType              string   `protobuf:"bytes,5,opt,name=AppType,proto3" json:"AppType,omitempty"`
	AppStatus            int32    `protobuf:"varint,6,opt,name=AppStatus,proto3" json:"AppStatus,omitempty"`
	AppPeriod            string   `protobuf:"bytes,7,opt,name=AppPeriod,proto3" json:"AppPeriod,omitempty"`
	Topic                []string `protobuf:"bytes,8,rep,name=Topic,proto3" json:"Topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppEntity) Reset()         { *m = AppEntity{} }
func (m *AppEntity) String() string { return proto.CompactTextString(m) }
func (*AppEntity) ProtoMessage()    {}
func (*AppEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1fc3c520458c96c, []int{0}
}

func (m *AppEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEntity.Unmarshal(m, b)
}
func (m *AppEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEntity.Marshal(b, m, deterministic)
}
func (m *AppEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEntity.Merge(m, src)
}
func (m *AppEntity) XXX_Size() int {
	return xxx_messageInfo_AppEntity.Size(m)
}
func (m *AppEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEntity.DiscardUnknown(m)
}

var xxx_messageInfo_AppEntity proto.InternalMessageInfo

func (m *AppEntity) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *AppEntity) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *AppEntity) GetUpdateAt() int64 {
	if m != nil {
		return m.UpdateAt
	}
	return 0
}

func (m *AppEntity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AppEntity) GetAppType() string {
	if m != nil {
		return m.AppType
	}
	return ""
}

func (m *AppEntity) GetAppStatus() int32 {
	if m != nil {
		return m.AppStatus
	}
	return 0
}

func (m *AppEntity) GetAppPeriod() string {
	if m != nil {
		return m.AppPeriod
	}
	return ""
}

func (m *AppEntity) GetTopic() []string {
	if m != nil {
		return m.Topic
	}
	return nil
}

type AppList struct {
	Ael                  []*AppEntity `protobuf:"bytes,1,rep,name=ael,proto3" json:"ael,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AppList) Reset()         { *m = AppList{} }
func (m *AppList) String() string { return proto.CompactTextString(m) }
func (*AppList) ProtoMessage()    {}
func (*AppList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1fc3c520458c96c, []int{1}
}

func (m *AppList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppList.Unmarshal(m, b)
}
func (m *AppList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppList.Marshal(b, m, deterministic)
}
func (m *AppList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppList.Merge(m, src)
}
func (m *AppList) XXX_Size() int {
	return xxx_messageInfo_AppList.Size(m)
}
func (m *AppList) XXX_DiscardUnknown() {
	xxx_messageInfo_AppList.DiscardUnknown(m)
}

var xxx_messageInfo_AppList proto.InternalMessageInfo

func (m *AppList) GetAel() []*AppEntity {
	if m != nil {
		return m.Ael
	}
	return nil
}

func init() {
	proto.RegisterType((*AppEntity)(nil), "AppEntity")
	proto.RegisterType((*AppList)(nil), "AppList")
}

func init() { proto.RegisterFile("proto/appmanager/appmanager.proto", fileDescriptor_d1fc3c520458c96c) }

var fileDescriptor_d1fc3c520458c96c = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcb, 0x6a, 0xeb, 0x30,
	0x14, 0x8c, 0xae, 0x6f, 0x9c, 0xf8, 0x84, 0x0b, 0x17, 0xd1, 0x85, 0x30, 0x59, 0xb8, 0xee, 0xa2,
	0x86, 0x82, 0x4d, 0xd3, 0x2f, 0x30, 0x25, 0x74, 0xd3, 0x94, 0xe2, 0xa6, 0x1f, 0xa0, 0x44, 0x6a,
	0x2a, 0xb0, 0xad, 0x53, 0xfb, 0x04, 0xea, 0x3f, 0xe8, 0xbf, 0xf6, 0x27, 0x8a, 0x1f, 0x79, 0x2c,
	0xbb, 0x92, 0x66, 0xe6, 0xcc, 0x19, 0x31, 0x82, 0x4b, 0xac, 0x2c, 0xd9, 0x44, 0x22, 0x16, 0xb2,
	0x94, 0x3b, 0x5d, 0x9d, 0x5d, 0xe3, 0x4e, 0xf3, 0x6f, 0x77, 0x86, 0xde, 0xf7, 0x9b, 0x78, 0x6b,
	0x8b, 0x44, 0xe9, 0x4f, 0x53, 0x7e, 0x24, 0x7b, 0x32, 0x79, 0x9d, 0xf4, 0xd6, 0x5d, 0x6e, 0x37,
	0x32, 0x1f, 0x8e, 0xde, 0x12, 0x7e, 0x33, 0xf0, 0x52, 0xc4, 0x65, 0x49, 0x86, 0x1a, 0x2e, 0x60,
	0x92, 0x22, 0x3e, 0xc9, 0x42, 0x0b, 0x16, 0xb0, 0xc8, 0xcb, 0x0e, 0x90, 0xfb, 0x30, 0xbd, 0xaf,
	0xb4, 0x24, 0x9d, 0x92, 0xf8, 0x13, 0xb0, 0xc8, 0xc9, 0x8e, 0xb8, 0xd5, 0x5e, 0x51, 0xf5, 0x9a,
	0xd3, 0x6b, 0x07, 0xcc, 0x03, 0x98, 0x29, 0x5d, 0x6f, 0x2b, 0x83, 0x64, 0x6c, 0x29, 0xfe, 0x76,
	0x5b, 0xcf, 0xa9, 0x21, 0x73, 0xdd, 0xa0, 0x16, 0xe3, 0x63, 0x66, 0x0b, 0xf9, 0xbc, 0x7b, 0xda,
	0x0b, 0x49, 0xda, 0xd7, 0xc2, 0x0d, 0x58, 0x34, 0xce, 0x4e, 0xc4, 0xa0, 0x3e, 0xeb, 0xca, 0x58,
	0x25, 0x26, 0x9d, 0xf3, 0x44, 0xf0, 0x0b, 0x18, 0xaf, 0x2d, 0x9a, 0xad, 0x98, 0x06, 0x4e, 0xe4,
	0x65, 0x3d, 0x08, 0xaf, 0xbb, 0xac, 0x47, 0x53, 0x13, 0x9f, 0x83, 0x23, 0x75, 0x2e, 0x58, 0xe0,
	0x44, 0xb3, 0x05, 0xc4, 0xc7, 0x0e, 0xb2, 0x96, 0x5e, 0x7c, 0x31, 0x80, 0x14, 0x71, 0xd5, 0xd7,
	0xcb, 0xaf, 0xc0, 0x7d, 0xd0, 0x94, 0x22, 0xf2, 0x7f, 0xf1, 0x50, 0xdf, 0xb2, 0x40, 0x6a, 0xfc,
	0x69, 0x3c, 0xec, 0x0b, 0x47, 0x3c, 0x02, 0x37, 0x55, 0xaa, 0x1d, 0x3a, 0x5b, 0xe7, 0xff, 0x3f,
	0x18, 0x32, 0x5d, 0xa3, 0x2d, 0x6b, 0x1d, 0x8e, 0xf8, 0x0d, 0x78, 0x2b, 0xab, 0xcc, 0x5b, 0xf3,
	0x8b, 0xe1, 0x8d, 0xdb, 0x7d, 0xd4, 0xdd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x1b, 0x05,
	0x19, 0x00, 0x02, 0x00, 0x00,
}
