// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/lolibrary/lolibrary/service.image/proto/image.proto

package imageproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lolibrary/lolibrary/cmd/protoc-gen-router/proto"
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

type Image struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename    string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	UserId      string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SubjectType string `protobuf:"bytes,5,opt,name=subject_type,json=subjectType,proto3" json:"subject_type,omitempty"`
	SubjectId   string `protobuf:"bytes,6,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	// timestamps
	CreatedAt            string   `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,101,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_04936b294c4d7e54, []int{0}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Image) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Image) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Image) GetSubjectType() string {
	if m != nil {
		return m.SubjectType
	}
	return ""
}

func (m *Image) GetSubjectId() string {
	if m != nil {
		return m.SubjectId
	}
	return ""
}

func (m *Image) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Image) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type POSTCreateImageRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename             string   `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *POSTCreateImageRequest) Reset()         { *m = POSTCreateImageRequest{} }
func (m *POSTCreateImageRequest) String() string { return proto.CompactTextString(m) }
func (*POSTCreateImageRequest) ProtoMessage()    {}
func (*POSTCreateImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04936b294c4d7e54, []int{1}
}

func (m *POSTCreateImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POSTCreateImageRequest.Unmarshal(m, b)
}
func (m *POSTCreateImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POSTCreateImageRequest.Marshal(b, m, deterministic)
}
func (m *POSTCreateImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POSTCreateImageRequest.Merge(m, src)
}
func (m *POSTCreateImageRequest) XXX_Size() int {
	return xxx_messageInfo_POSTCreateImageRequest.Size(m)
}
func (m *POSTCreateImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_POSTCreateImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_POSTCreateImageRequest proto.InternalMessageInfo

func (m *POSTCreateImageRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *POSTCreateImageRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *POSTCreateImageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type POSTCreateImageResponse struct {
	Image                *Image   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *POSTCreateImageResponse) Reset()         { *m = POSTCreateImageResponse{} }
func (m *POSTCreateImageResponse) String() string { return proto.CompactTextString(m) }
func (*POSTCreateImageResponse) ProtoMessage()    {}
func (*POSTCreateImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04936b294c4d7e54, []int{2}
}

func (m *POSTCreateImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POSTCreateImageResponse.Unmarshal(m, b)
}
func (m *POSTCreateImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POSTCreateImageResponse.Marshal(b, m, deterministic)
}
func (m *POSTCreateImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POSTCreateImageResponse.Merge(m, src)
}
func (m *POSTCreateImageResponse) XXX_Size() int {
	return xxx_messageInfo_POSTCreateImageResponse.Size(m)
}
func (m *POSTCreateImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_POSTCreateImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_POSTCreateImageResponse proto.InternalMessageInfo

func (m *POSTCreateImageResponse) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type GETReadImageRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename             string   `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETReadImageRequest) Reset()         { *m = GETReadImageRequest{} }
func (m *GETReadImageRequest) String() string { return proto.CompactTextString(m) }
func (*GETReadImageRequest) ProtoMessage()    {}
func (*GETReadImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04936b294c4d7e54, []int{3}
}

func (m *GETReadImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETReadImageRequest.Unmarshal(m, b)
}
func (m *GETReadImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETReadImageRequest.Marshal(b, m, deterministic)
}
func (m *GETReadImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETReadImageRequest.Merge(m, src)
}
func (m *GETReadImageRequest) XXX_Size() int {
	return xxx_messageInfo_GETReadImageRequest.Size(m)
}
func (m *GETReadImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GETReadImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GETReadImageRequest proto.InternalMessageInfo

func (m *GETReadImageRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GETReadImageRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type GETReadImageResponse struct {
	Image                *Image   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETReadImageResponse) Reset()         { *m = GETReadImageResponse{} }
func (m *GETReadImageResponse) String() string { return proto.CompactTextString(m) }
func (*GETReadImageResponse) ProtoMessage()    {}
func (*GETReadImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04936b294c4d7e54, []int{4}
}

func (m *GETReadImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETReadImageResponse.Unmarshal(m, b)
}
func (m *GETReadImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETReadImageResponse.Marshal(b, m, deterministic)
}
func (m *GETReadImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETReadImageResponse.Merge(m, src)
}
func (m *GETReadImageResponse) XXX_Size() int {
	return xxx_messageInfo_GETReadImageResponse.Size(m)
}
func (m *GETReadImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GETReadImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GETReadImageResponse proto.InternalMessageInfo

func (m *GETReadImageResponse) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func init() {
	proto.RegisterType((*Image)(nil), "imageproto.Image")
	proto.RegisterType((*POSTCreateImageRequest)(nil), "imageproto.POSTCreateImageRequest")
	proto.RegisterType((*POSTCreateImageResponse)(nil), "imageproto.POSTCreateImageResponse")
	proto.RegisterType((*GETReadImageRequest)(nil), "imageproto.GETReadImageRequest")
	proto.RegisterType((*GETReadImageResponse)(nil), "imageproto.GETReadImageResponse")
}

func init() {
	proto.RegisterFile("github.com/lolibrary/lolibrary/service.image/proto/image.proto", fileDescriptor_04936b294c4d7e54)
}

var fileDescriptor_04936b294c4d7e54 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0xca, 0xa6, 0x21, 0x13, 0x16, 0x84, 0x41, 0xac, 0x15, 0xa9, 0xa2, 0x84, 0x03, 0x5c,
	0x9a, 0x48, 0xcb, 0x1d, 0x54, 0x10, 0x5a, 0xf5, 0x04, 0x2a, 0x3d, 0x20, 0x71, 0xa8, 0x92, 0x78,
	0x28, 0x46, 0x4d, 0x13, 0x1c, 0x07, 0x69, 0xaf, 0xf9, 0x98, 0xfd, 0x83, 0xfd, 0x1f, 0x7e, 0x80,
	0x7f, 0x40, 0x19, 0xbb, 0x25, 0x15, 0x05, 0xd4, 0xbd, 0xbd, 0x79, 0xef, 0xcd, 0xd3, 0x8c, 0x47,
	0x86, 0x97, 0x2b, 0xa9, 0xbf, 0x34, 0x59, 0x9c, 0x97, 0x45, 0xb2, 0x2e, 0xd7, 0x32, 0x53, 0xa9,
	0xba, 0xec, 0xa1, 0x1a, 0xd5, 0x77, 0x99, 0x63, 0x2c, 0x8b, 0x74, 0x85, 0x49, 0xa5, 0x4a, 0x5d,
	0x26, 0x84, 0x63, 0xc2, 0x0c, 0xa8, 0x20, 0x1c, 0x5e, 0xfc, 0x27, 0x2b, 0x2f, 0x84, 0x49, 0xc8,
	0x27, 0x2b, 0xdc, 0x4c, 0x54, 0xd9, 0x68, 0x54, 0x36, 0xd3, 0x14, 0x26, 0x34, 0xfa, 0xe1, 0x80,
	0x3b, 0xeb, 0x72, 0xd9, 0x5d, 0x18, 0x48, 0xc1, 0x9d, 0xb1, 0xf3, 0xdc, 0x9f, 0x0f, 0xa4, 0x60,
	0x21, 0xdc, 0xfe, 0x2c, 0xd7, 0xb8, 0x49, 0x0b, 0xe4, 0x03, 0x62, 0x77, 0x35, 0x63, 0x70, 0xa2,
	0x2f, 0x2b, 0xe4, 0xb7, 0x88, 0x27, 0xcc, 0xce, 0xc0, 0x6b, 0x6a, 0x54, 0x4b, 0x29, 0xf8, 0x09,
	0xd1, 0xc3, 0xae, 0x9c, 0x09, 0xf6, 0x04, 0xee, 0xd4, 0x4d, 0xf6, 0x15, 0x73, 0xbd, 0xa4, 0x26,
	0x97, 0xd4, 0xc0, 0x72, 0x8b, 0xae, 0x77, 0x04, 0xb0, 0xb5, 0x48, 0xc1, 0x87, 0x64, 0xf0, 0x2d,
	0x33, 0x13, 0x9d, 0x9c, 0x2b, 0x4c, 0x35, 0x8a, 0x65, 0xaa, 0xb9, 0x30, 0xb2, 0x65, 0xa6, 0xba,
	0x93, 0x9b, 0x4a, 0x6c, 0x65, 0x34, 0xb2, 0x65, 0xa6, 0x3a, 0xfa, 0x08, 0x8f, 0xde, 0xbf, 0xfb,
	0xb0, 0x78, 0x43, 0x7e, 0xda, 0x75, 0x8e, 0xdf, 0x1a, 0xac, 0xf5, 0xb1, 0x2b, 0x13, 0x6f, 0x57,
	0xee, 0x70, 0xf4, 0x1a, 0xce, 0xfe, 0x48, 0xae, 0xab, 0x72, 0x53, 0x23, 0x7b, 0x06, 0x2e, 0x9d,
	0x8b, 0xd2, 0x83, 0xf3, 0xfb, 0xf1, 0xef, 0xe3, 0xc5, 0xc6, 0x69, 0xf4, 0x68, 0x0a, 0x0f, 0x2e,
	0xde, 0x2e, 0xe6, 0x98, 0x8a, 0x9b, 0x8e, 0x16, 0xbd, 0x82, 0x87, 0xfb, 0x11, 0x47, 0xce, 0x70,
	0xfe, 0xd3, 0xb1, 0x4e, 0x86, 0x10, 0xf4, 0xb6, 0x61, 0x51, 0xbf, 0xe5, 0xf0, 0x23, 0x86, 0x4f,
	0xff, 0xe9, 0x31, 0xa3, 0x44, 0xa7, 0xed, 0x35, 0xf7, 0xc1, 0x4b, 0xcc, 0xd1, 0xd8, 0x27, 0xf0,
	0x77, 0xe3, 0xb2, 0xc7, 0xfd, 0x80, 0x03, 0x6f, 0x11, 0x8e, 0xff, 0x6e, 0xb0, 0xf1, 0x41, 0x7b,
	0xcd, 0x3d, 0x70, 0x13, 0x85, 0xa9, 0x08, 0xc7, 0xed, 0xd5, 0xe8, 0x1e, 0x9c, 0xee, 0x7d, 0xa7,
	0xf6, 0x6a, 0xe4, 0x33, 0xaf, 0x9e, 0x50, 0x91, 0x0d, 0x29, 0xea, 0xc5, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0x04, 0x16, 0x61, 0x92, 0x03, 0x00, 0x00,
}
