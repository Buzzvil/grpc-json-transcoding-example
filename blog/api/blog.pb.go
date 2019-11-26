// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog.proto

package blog

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Post struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string               `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Tags                 []string             `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Post) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Post) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type CreatePostRequest struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRequest) Reset()         { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()    {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{1}
}

func (m *CreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRequest.Unmarshal(m, b)
}
func (m *CreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRequest.Marshal(b, m, deterministic)
}
func (m *CreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRequest.Merge(m, src)
}
func (m *CreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostRequest.Size(m)
}
func (m *CreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRequest proto.InternalMessageInfo

func (m *CreatePostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type ListPostsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsRequest) Reset()         { *m = ListPostsRequest{} }
func (m *ListPostsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPostsRequest) ProtoMessage()    {}
func (*ListPostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{2}
}

func (m *ListPostsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsRequest.Unmarshal(m, b)
}
func (m *ListPostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsRequest.Marshal(b, m, deterministic)
}
func (m *ListPostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsRequest.Merge(m, src)
}
func (m *ListPostsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPostsRequest.Size(m)
}
func (m *ListPostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsRequest proto.InternalMessageInfo

type ListPostsResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsResponse) Reset()         { *m = ListPostsResponse{} }
func (m *ListPostsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPostsResponse) ProtoMessage()    {}
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{3}
}

func (m *ListPostsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsResponse.Unmarshal(m, b)
}
func (m *ListPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsResponse.Marshal(b, m, deterministic)
}
func (m *ListPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsResponse.Merge(m, src)
}
func (m *ListPostsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPostsResponse.Size(m)
}
func (m *ListPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsResponse proto.InternalMessageInfo

func (m *ListPostsResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type UpdatePostRequest struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostRequest) Reset()         { *m = UpdatePostRequest{} }
func (m *UpdatePostRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePostRequest) ProtoMessage()    {}
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{4}
}

func (m *UpdatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostRequest.Unmarshal(m, b)
}
func (m *UpdatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostRequest.Merge(m, src)
}
func (m *UpdatePostRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePostRequest.Size(m)
}
func (m *UpdatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostRequest proto.InternalMessageInfo

func (m *UpdatePostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "blog.Post")
	proto.RegisterType((*CreatePostRequest)(nil), "blog.CreatePostRequest")
	proto.RegisterType((*ListPostsRequest)(nil), "blog.ListPostsRequest")
	proto.RegisterType((*ListPostsResponse)(nil), "blog.ListPostsResponse")
	proto.RegisterType((*UpdatePostRequest)(nil), "blog.UpdatePostRequest")
}

func init() { proto.RegisterFile("blog.proto", fileDescriptor_6745b25902462fb1) }

var fileDescriptor_6745b25902462fb1 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x49, 0x9a, 0x3c, 0xe8, 0x2d, 0xef, 0x4f, 0x2f, 0xaf, 0xaf, 0x21, 0xef, 0xf1, 0x1a,
	0x02, 0x42, 0x57, 0x09, 0xb6, 0xb8, 0xd0, 0x9d, 0xba, 0x15, 0x91, 0x54, 0xc1, 0x9d, 0x24, 0xcd,
	0x18, 0x06, 0xd2, 0x4c, 0xec, 0xdc, 0x16, 0x44, 0xdc, 0xb8, 0x16, 0x5c, 0xf8, 0xd1, 0xfc, 0x0a,
	0x7e, 0x10, 0x99, 0x99, 0xd6, 0xfe, 0x5b, 0xb9, 0xca, 0xcc, 0x99, 0x73, 0xcf, 0xfc, 0x4e, 0x06,
	0x20, 0x2b, 0x45, 0x11, 0xd5, 0x53, 0x41, 0x02, 0x1d, 0xb5, 0xf6, 0x7b, 0x85, 0x10, 0x45, 0xc9,
	0x62, 0xad, 0x65, 0xb3, 0xdb, 0x98, 0xf8, 0x84, 0x49, 0x4a, 0x27, 0xb5, 0xb1, 0xf9, 0xff, 0x16,
	0x86, 0xb4, 0xe6, 0x71, 0x5a, 0x55, 0x82, 0x52, 0xe2, 0xa2, 0x92, 0xe6, 0x34, 0x7c, 0xb1, 0xc0,
	0xb9, 0x10, 0x92, 0xf0, 0x07, 0xd8, 0x3c, 0xf7, 0xac, 0xc0, 0xea, 0xbb, 0x89, 0xcd, 0x73, 0xfc,
	0x0d, 0x2e, 0x71, 0x2a, 0x99, 0x67, 0x07, 0x56, 0xbf, 0x99, 0x98, 0x0d, 0x22, 0x38, 0x99, 0xc8,
	0xef, 0xbd, 0x86, 0x16, 0xf5, 0x5a, 0x69, 0x94, 0x16, 0xd2, 0x73, 0x82, 0x86, 0xd2, 0xd4, 0x1a,
	0x0f, 0x01, 0xc6, 0x53, 0x96, 0x12, 0xcb, 0x6f, 0x52, 0xf2, 0xdc, 0xc0, 0xea, 0xb7, 0x06, 0x7e,
	0x64, 0x48, 0xa2, 0x25, 0x6a, 0x74, 0xb9, 0x44, 0x4d, 0x9a, 0x0b, 0xf7, 0x31, 0x85, 0x43, 0x68,
	0x9f, 0xea, 0x8d, 0xc2, 0x4a, 0xd8, 0xdd, 0x8c, 0x49, 0xc2, 0xff, 0xe0, 0xd4, 0x42, 0x92, 0xe6,
	0x6b, 0x0d, 0x20, 0xd2, 0xbf, 0x41, 0x1b, 0xb4, 0x1e, 0x22, 0xfc, 0x3a, 0xe3, 0x92, 0x94, 0x22,
	0x17, 0x33, 0xe1, 0x01, 0xb4, 0xd7, 0x34, 0x59, 0x8b, 0x4a, 0x32, 0x0c, 0xc0, 0x55, 0x03, 0xd2,
	0xb3, 0x82, 0xc6, 0x56, 0x92, 0x39, 0x50, 0xf7, 0x5f, 0xd5, 0xf9, 0xd7, 0xee, 0x1f, 0x3c, 0xdb,
	0xd0, 0x3a, 0x29, 0x45, 0x31, 0x62, 0xd3, 0x39, 0x1f, 0x33, 0x3c, 0x07, 0x58, 0x95, 0xc0, 0xae,
	0xf1, 0xef, 0xd4, 0xf2, 0xd7, 0x82, 0xc2, 0xbf, 0x4f, 0x6f, 0xef, 0xaf, 0x76, 0x27, 0xfc, 0xae,
	0x5f, 0x6a, 0xbe, 0x1f, 0x6b, 0x9c, 0x23, 0x9d, 0x8f, 0x23, 0x68, 0x7e, 0x76, 0xc1, 0x3f, 0x66,
	0x6a, 0xbb, 0xb0, 0xdf, 0xdd, 0xd1, 0x4d, 0xe9, 0xb0, 0xa3, 0xa3, 0x7f, 0xe2, 0x66, 0x34, 0x5e,
	0x03, 0xac, 0x9a, 0x2e, 0x21, 0x77, 0xba, 0x6f, 0x40, 0xee, 0xe9, 0xa4, 0x9e, 0xdf, 0xdd, 0x48,
	0x8a, 0x1f, 0xd4, 0x27, 0xe2, 0xf9, 0xa3, 0xc1, 0xcd, 0xbe, 0xe9, 0x27, 0x1e, 0x7e, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xc8, 0x67, 0xcb, 0xee, 0xaf, 0x02, 0x00, 0x00,
}
