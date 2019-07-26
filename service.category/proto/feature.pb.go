// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/lolibrary/lolibrary/service.category/proto/feature.proto

package categoryproto

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

type Category struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// timestamps
	CreatedAt            string   `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,101,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b517a3e9a10cef5, []int{0}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Category) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Category) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type GETReadCategoryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETReadCategoryRequest) Reset()         { *m = GETReadCategoryRequest{} }
func (m *GETReadCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*GETReadCategoryRequest) ProtoMessage()    {}
func (*GETReadCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b517a3e9a10cef5, []int{1}
}

func (m *GETReadCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETReadCategoryRequest.Unmarshal(m, b)
}
func (m *GETReadCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETReadCategoryRequest.Marshal(b, m, deterministic)
}
func (m *GETReadCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETReadCategoryRequest.Merge(m, src)
}
func (m *GETReadCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_GETReadCategoryRequest.Size(m)
}
func (m *GETReadCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GETReadCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GETReadCategoryRequest proto.InternalMessageInfo

func (m *GETReadCategoryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GETReadCategoryRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type GETReadCategoryResponse struct {
	Category             *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GETReadCategoryResponse) Reset()         { *m = GETReadCategoryResponse{} }
func (m *GETReadCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*GETReadCategoryResponse) ProtoMessage()    {}
func (*GETReadCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b517a3e9a10cef5, []int{2}
}

func (m *GETReadCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETReadCategoryResponse.Unmarshal(m, b)
}
func (m *GETReadCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETReadCategoryResponse.Marshal(b, m, deterministic)
}
func (m *GETReadCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETReadCategoryResponse.Merge(m, src)
}
func (m *GETReadCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_GETReadCategoryResponse.Size(m)
}
func (m *GETReadCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GETReadCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GETReadCategoryResponse proto.InternalMessageInfo

func (m *GETReadCategoryResponse) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type GETListCategoriesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GETListCategoriesRequest) Reset()         { *m = GETListCategoriesRequest{} }
func (m *GETListCategoriesRequest) String() string { return proto.CompactTextString(m) }
func (*GETListCategoriesRequest) ProtoMessage()    {}
func (*GETListCategoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b517a3e9a10cef5, []int{3}
}

func (m *GETListCategoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETListCategoriesRequest.Unmarshal(m, b)
}
func (m *GETListCategoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETListCategoriesRequest.Marshal(b, m, deterministic)
}
func (m *GETListCategoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETListCategoriesRequest.Merge(m, src)
}
func (m *GETListCategoriesRequest) XXX_Size() int {
	return xxx_messageInfo_GETListCategoriesRequest.Size(m)
}
func (m *GETListCategoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GETListCategoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GETListCategoriesRequest proto.InternalMessageInfo

type GETListCategoriesResponse struct {
	Categories           []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GETListCategoriesResponse) Reset()         { *m = GETListCategoriesResponse{} }
func (m *GETListCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*GETListCategoriesResponse) ProtoMessage()    {}
func (*GETListCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b517a3e9a10cef5, []int{4}
}

func (m *GETListCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GETListCategoriesResponse.Unmarshal(m, b)
}
func (m *GETListCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GETListCategoriesResponse.Marshal(b, m, deterministic)
}
func (m *GETListCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GETListCategoriesResponse.Merge(m, src)
}
func (m *GETListCategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_GETListCategoriesResponse.Size(m)
}
func (m *GETListCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GETListCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GETListCategoriesResponse proto.InternalMessageInfo

func (m *GETListCategoriesResponse) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type POSTCreateCategoryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *POSTCreateCategoryRequest) Reset()         { *m = POSTCreateCategoryRequest{} }
func (m *POSTCreateCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*POSTCreateCategoryRequest) ProtoMessage()    {}
func (*POSTCreateCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b517a3e9a10cef5, []int{5}
}

func (m *POSTCreateCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POSTCreateCategoryRequest.Unmarshal(m, b)
}
func (m *POSTCreateCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POSTCreateCategoryRequest.Marshal(b, m, deterministic)
}
func (m *POSTCreateCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POSTCreateCategoryRequest.Merge(m, src)
}
func (m *POSTCreateCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_POSTCreateCategoryRequest.Size(m)
}
func (m *POSTCreateCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_POSTCreateCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_POSTCreateCategoryRequest proto.InternalMessageInfo

func (m *POSTCreateCategoryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *POSTCreateCategoryRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *POSTCreateCategoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type POSTCreateCategoryResponse struct {
	Category             *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *POSTCreateCategoryResponse) Reset()         { *m = POSTCreateCategoryResponse{} }
func (m *POSTCreateCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*POSTCreateCategoryResponse) ProtoMessage()    {}
func (*POSTCreateCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b517a3e9a10cef5, []int{6}
}

func (m *POSTCreateCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POSTCreateCategoryResponse.Unmarshal(m, b)
}
func (m *POSTCreateCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POSTCreateCategoryResponse.Marshal(b, m, deterministic)
}
func (m *POSTCreateCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POSTCreateCategoryResponse.Merge(m, src)
}
func (m *POSTCreateCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_POSTCreateCategoryResponse.Size(m)
}
func (m *POSTCreateCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_POSTCreateCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_POSTCreateCategoryResponse proto.InternalMessageInfo

func (m *POSTCreateCategoryResponse) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type PUTUpdateCategoryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PUTUpdateCategoryRequest) Reset()         { *m = PUTUpdateCategoryRequest{} }
func (m *PUTUpdateCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*PUTUpdateCategoryRequest) ProtoMessage()    {}
func (*PUTUpdateCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b517a3e9a10cef5, []int{7}
}

func (m *PUTUpdateCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PUTUpdateCategoryRequest.Unmarshal(m, b)
}
func (m *PUTUpdateCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PUTUpdateCategoryRequest.Marshal(b, m, deterministic)
}
func (m *PUTUpdateCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PUTUpdateCategoryRequest.Merge(m, src)
}
func (m *PUTUpdateCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_PUTUpdateCategoryRequest.Size(m)
}
func (m *PUTUpdateCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PUTUpdateCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PUTUpdateCategoryRequest proto.InternalMessageInfo

func (m *PUTUpdateCategoryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PUTUpdateCategoryRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *PUTUpdateCategoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PUTUpdateCategoryResponse struct {
	Category             *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PUTUpdateCategoryResponse) Reset()         { *m = PUTUpdateCategoryResponse{} }
func (m *PUTUpdateCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*PUTUpdateCategoryResponse) ProtoMessage()    {}
func (*PUTUpdateCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b517a3e9a10cef5, []int{8}
}

func (m *PUTUpdateCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PUTUpdateCategoryResponse.Unmarshal(m, b)
}
func (m *PUTUpdateCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PUTUpdateCategoryResponse.Marshal(b, m, deterministic)
}
func (m *PUTUpdateCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PUTUpdateCategoryResponse.Merge(m, src)
}
func (m *PUTUpdateCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_PUTUpdateCategoryResponse.Size(m)
}
func (m *PUTUpdateCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PUTUpdateCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PUTUpdateCategoryResponse proto.InternalMessageInfo

func (m *PUTUpdateCategoryResponse) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type DELETERemoveCategoryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DELETERemoveCategoryRequest) Reset()         { *m = DELETERemoveCategoryRequest{} }
func (m *DELETERemoveCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*DELETERemoveCategoryRequest) ProtoMessage()    {}
func (*DELETERemoveCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b517a3e9a10cef5, []int{9}
}

func (m *DELETERemoveCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DELETERemoveCategoryRequest.Unmarshal(m, b)
}
func (m *DELETERemoveCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DELETERemoveCategoryRequest.Marshal(b, m, deterministic)
}
func (m *DELETERemoveCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DELETERemoveCategoryRequest.Merge(m, src)
}
func (m *DELETERemoveCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_DELETERemoveCategoryRequest.Size(m)
}
func (m *DELETERemoveCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DELETERemoveCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DELETERemoveCategoryRequest proto.InternalMessageInfo

func (m *DELETERemoveCategoryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DELETERemoveCategoryResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DELETERemoveCategoryResponse) Reset()         { *m = DELETERemoveCategoryResponse{} }
func (m *DELETERemoveCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*DELETERemoveCategoryResponse) ProtoMessage()    {}
func (*DELETERemoveCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b517a3e9a10cef5, []int{10}
}

func (m *DELETERemoveCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DELETERemoveCategoryResponse.Unmarshal(m, b)
}
func (m *DELETERemoveCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DELETERemoveCategoryResponse.Marshal(b, m, deterministic)
}
func (m *DELETERemoveCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DELETERemoveCategoryResponse.Merge(m, src)
}
func (m *DELETERemoveCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_DELETERemoveCategoryResponse.Size(m)
}
func (m *DELETERemoveCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DELETERemoveCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DELETERemoveCategoryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Category)(nil), "categoryproto.Category")
	proto.RegisterType((*GETReadCategoryRequest)(nil), "categoryproto.GETReadCategoryRequest")
	proto.RegisterType((*GETReadCategoryResponse)(nil), "categoryproto.GETReadCategoryResponse")
	proto.RegisterType((*GETListCategoriesRequest)(nil), "categoryproto.GETListCategoriesRequest")
	proto.RegisterType((*GETListCategoriesResponse)(nil), "categoryproto.GETListCategoriesResponse")
	proto.RegisterType((*POSTCreateCategoryRequest)(nil), "categoryproto.POSTCreateCategoryRequest")
	proto.RegisterType((*POSTCreateCategoryResponse)(nil), "categoryproto.POSTCreateCategoryResponse")
	proto.RegisterType((*PUTUpdateCategoryRequest)(nil), "categoryproto.PUTUpdateCategoryRequest")
	proto.RegisterType((*PUTUpdateCategoryResponse)(nil), "categoryproto.PUTUpdateCategoryResponse")
	proto.RegisterType((*DELETERemoveCategoryRequest)(nil), "categoryproto.DELETERemoveCategoryRequest")
	proto.RegisterType((*DELETERemoveCategoryResponse)(nil), "categoryproto.DELETERemoveCategoryResponse")
}

func init() {
	proto.RegisterFile("github.com/lolibrary/lolibrary/service.category/proto/feature.proto", fileDescriptor_4b517a3e9a10cef5)
}

var fileDescriptor_4b517a3e9a10cef5 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x55, 0xd2, 0x42, 0x9b, 0x69, 0x6b, 0xa1, 0x3d, 0xd0, 0x8d, 0x21, 0xa8, 0x5a, 0x09, 0x28,
	0xa0, 0xd8, 0x52, 0x7b, 0xe0, 0xc2, 0xa5, 0x4a, 0xad, 0x5c, 0x2a, 0x08, 0xae, 0x7b, 0x46, 0x8e,
	0x77, 0x08, 0x46, 0x4e, 0x1c, 0xd6, 0xeb, 0x8a, 0xde, 0x50, 0x3e, 0xa6, 0x7f, 0xd0, 0x3f, 0xe3,
	0x03, 0x50, 0x76, 0xd7, 0x21, 0x71, 0x9c, 0x26, 0xa2, 0xdc, 0x76, 0x67, 0xc6, 0xef, 0x3d, 0xbf,
	0x99, 0x1d, 0xe8, 0x0c, 0x62, 0xf9, 0x2d, 0xef, 0x3b, 0x51, 0x3a, 0x74, 0x93, 0x34, 0x89, 0xfb,
	0x22, 0x14, 0x37, 0x73, 0xa7, 0x0c, 0xc5, 0x75, 0x1c, 0xa1, 0x13, 0x85, 0x12, 0x07, 0xa9, 0xb8,
	0x71, 0xc7, 0x22, 0x95, 0xa9, 0xfb, 0x15, 0x43, 0x99, 0x0b, 0x74, 0xd4, 0x8d, 0x1c, 0x14, 0x59,
	0x75, 0xb5, 0xbb, 0x6b, 0x30, 0xa3, 0x21, 0xd7, 0x30, 0x51, 0x7b, 0x80, 0xa3, 0xb6, 0x48, 0x73,
	0x89, 0xc2, 0x00, 0xeb, 0x8b, 0xc6, 0x65, 0xbf, 0x6a, 0xb0, 0xdb, 0x31, 0xd0, 0xc4, 0x82, 0x7a,
	0xcc, 0x69, 0xed, 0xa8, 0x76, 0xdc, 0xf0, 0xeb, 0x31, 0x27, 0x04, 0xb6, 0xb3, 0x24, 0x1f, 0xd0,
	0xba, 0x8a, 0xa8, 0xf3, 0x34, 0x36, 0x0a, 0x87, 0x48, 0xb7, 0x74, 0x6c, 0x7a, 0x26, 0x2d, 0x80,
	0x48, 0x60, 0x28, 0x91, 0x7f, 0x09, 0x25, 0xe5, 0x2a, 0xd3, 0x30, 0x91, 0x33, 0x39, 0x4d, 0xe7,
	0x63, 0x5e, 0xa4, 0x51, 0xa7, 0x4d, 0xe4, 0x4c, 0xb2, 0x0f, 0xf0, 0xb4, 0xeb, 0x05, 0x3e, 0x86,
	0xbc, 0x10, 0xe2, 0xe3, 0x8f, 0x1c, 0x33, 0xb9, 0x89, 0x1e, 0xf6, 0x11, 0x0e, 0x97, 0xbe, 0xce,
	0xc6, 0xe9, 0x28, 0x43, 0x72, 0x0a, 0xbb, 0x85, 0x6b, 0x0a, 0x64, 0xef, 0xe4, 0xd0, 0x59, 0xb0,
	0xd1, 0x99, 0x7d, 0x32, 0x2b, 0x64, 0x36, 0xd0, 0xae, 0x17, 0x5c, 0xc4, 0x99, 0x34, 0xc9, 0x18,
	0x33, 0xa3, 0x87, 0x05, 0xd0, 0xac, 0xc8, 0x19, 0xb6, 0xf7, 0x00, 0xd1, 0x2c, 0x4a, 0x6b, 0x47,
	0x5b, 0xf7, 0xf1, 0xcd, 0x95, 0xb2, 0x4b, 0x68, 0xf6, 0x3e, 0x5d, 0x06, 0x1d, 0xe5, 0xd7, 0x3f,
	0x58, 0x50, 0xd5, 0x12, 0xf6, 0x19, 0xec, 0x2a, 0xd0, 0x87, 0x38, 0xe3, 0x03, 0xed, 0x5d, 0x05,
	0x57, 0xaa, 0x6f, 0xff, 0x4b, 0x66, 0x0f, 0x9a, 0x15, 0x98, 0x0f, 0x51, 0xd9, 0x86, 0x67, 0xe7,
	0xde, 0x85, 0x17, 0x78, 0x3e, 0x0e, 0xd3, 0xeb, 0x75, 0x42, 0xd9, 0x0b, 0x78, 0x5e, 0x5d, 0xae,
	0x35, 0x9c, 0xfc, 0xde, 0xfe, 0x2b, 0x82, 0x7c, 0x87, 0xfd, 0xf9, 0x41, 0x23, 0x2f, 0x4b, 0x72,
	0xaa, 0xc7, 0xd8, 0x7e, 0xb5, 0xae, 0x4c, 0x73, 0xb1, 0xbd, 0xc9, 0x1d, 0xdd, 0x81, 0x47, 0xae,
	0xc0, 0x90, 0x13, 0x01, 0xd6, 0xa2, 0x2d, 0xe4, 0x75, 0x09, 0x66, 0x55, 0x33, 0xec, 0xe3, 0xf5,
	0x85, 0x86, 0xf1, 0x60, 0x72, 0x47, 0x1b, 0xb0, 0xe3, 0xea, 0xc7, 0x48, 0x7e, 0x82, 0x75, 0x8e,
	0x09, 0xce, 0x71, 0xbe, 0x2d, 0x41, 0xdd, 0x63, 0xad, 0xfd, 0x6e, 0xa3, 0xda, 0x12, 0x33, 0x57,
	0x7c, 0x44, 0x82, 0xb5, 0x38, 0xaa, 0x64, 0xe9, 0x27, 0x56, 0x3d, 0x11, 0xfb, 0xcd, 0x06, 0x95,
	0x25, 0x56, 0xbd, 0x9b, 0xc8, 0x18, 0xac, 0xc5, 0xc7, 0xbc, 0xe4, 0xf1, 0xaa, 0x55, 0xb0, 0xe4,
	0xf1, 0xca, 0xbd, 0x30, 0xeb, 0x6a, 0x12, 0x4f, 0x47, 0x61, 0x72, 0xdb, 0x22, 0xf0, 0xa4, 0xbc,
	0xf2, 0x27, 0xb7, 0xad, 0x7d, 0x02, 0x59, 0xbb, 0xb8, 0xf7, 0x1f, 0x2b, 0xd4, 0xd3, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x50, 0x5a, 0xc7, 0x8f, 0x3c, 0x06, 0x00, 0x00,
}
