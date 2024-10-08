// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: post/post.proto

package post

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

// 定义实体结构
type PostData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`                // id
	Title     string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`           // 标题
	Content   string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`       // 内容
	UserId    int64  `protobuf:"varint,4,opt,name=UserId,proto3" json:"UserId,omitempty"`        // 作者id
	Views     int64  `protobuf:"varint,5,opt,name=Views,proto3" json:"Views,omitempty"`          // 查看数
	Likes     int64  `protobuf:"varint,6,opt,name=Likes,proto3" json:"Likes,omitempty"`          // 喜欢数
	Comments  int64  `protobuf:"varint,7,opt,name=Comments,proto3" json:"Comments,omitempty"`    // 评论数
	Shares    int64  `protobuf:"varint,8,opt,name=Shares,proto3" json:"Shares,omitempty"`        // 分享数
	Collects  int64  `protobuf:"varint,9,opt,name=Collects,proto3" json:"Collects,omitempty"`    // 收藏数
	CreatedAt int64  `protobuf:"varint,10,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"` // 创建时间
	UpdatedAt int64  `protobuf:"varint,11,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"` // 编辑时间
}

func (x *PostData) Reset() {
	*x = PostData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostData) ProtoMessage() {}

func (x *PostData) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostData.ProtoReflect.Descriptor instead.
func (*PostData) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{0}
}

func (x *PostData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PostData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostData) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PostData) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PostData) GetViews() int64 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *PostData) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *PostData) GetComments() int64 {
	if x != nil {
		return x.Comments
	}
	return 0
}

func (x *PostData) GetShares() int64 {
	if x != nil {
		return x.Shares
	}
	return 0
}

func (x *PostData) GetCollects() int64 {
	if x != nil {
		return x.Collects
	}
	return 0
}

func (x *PostData) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PostData) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreatePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostData *PostData `protobuf:"bytes,1,opt,name=PostData,proto3" json:"PostData,omitempty"`
}

func (x *CreatePostReq) Reset() {
	*x = CreatePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostReq) ProtoMessage() {}

func (x *CreatePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostReq.ProtoReflect.Descriptor instead.
func (*CreatePostReq) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePostReq) GetPostData() *PostData {
	if x != nil {
		return x.PostData
	}
	return nil
}

type CreatePostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePostResp) Reset() {
	*x = CreatePostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostResp) ProtoMessage() {}

func (x *CreatePostResp) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostResp.ProtoReflect.Descriptor instead.
func (*CreatePostResp) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{2}
}

type UpdatePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostData *PostData `protobuf:"bytes,1,opt,name=PostData,proto3" json:"PostData,omitempty"`
}

func (x *UpdatePostReq) Reset() {
	*x = UpdatePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostReq) ProtoMessage() {}

func (x *UpdatePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostReq.ProtoReflect.Descriptor instead.
func (*UpdatePostReq) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePostReq) GetPostData() *PostData {
	if x != nil {
		return x.PostData
	}
	return nil
}

type UpdatePostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePostResp) Reset() {
	*x = UpdatePostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostResp) ProtoMessage() {}

func (x *UpdatePostResp) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostResp.ProtoReflect.Descriptor instead.
func (*UpdatePostResp) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{4}
}

type DeletePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"varint,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
}

func (x *DeletePostReq) Reset() {
	*x = DeletePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostReq) ProtoMessage() {}

func (x *DeletePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostReq.ProtoReflect.Descriptor instead.
func (*DeletePostReq) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePostReq) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type DeletePostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePostResp) Reset() {
	*x = DeletePostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostResp) ProtoMessage() {}

func (x *DeletePostResp) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostResp.ProtoReflect.Descriptor instead.
func (*DeletePostResp) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{6}
}

type GetPostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"varint,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
}

func (x *GetPostReq) Reset() {
	*x = GetPostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostReq) ProtoMessage() {}

func (x *GetPostReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostReq.ProtoReflect.Descriptor instead.
func (*GetPostReq) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{7}
}

func (x *GetPostReq) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type GetPostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *PostData `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetPostResp) Reset() {
	*x = GetPostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResp) ProtoMessage() {}

func (x *GetPostResp) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResp.ProtoReflect.Descriptor instead.
func (*GetPostResp) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{8}
}

func (x *GetPostResp) GetInfo() *PostData {
	if x != nil {
		return x.Info
	}
	return nil
}

type BatchPostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostIds []int64 `protobuf:"varint,1,rep,packed,name=PostIds,proto3" json:"PostIds,omitempty"`
}

func (x *BatchPostReq) Reset() {
	*x = BatchPostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchPostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchPostReq) ProtoMessage() {}

func (x *BatchPostReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchPostReq.ProtoReflect.Descriptor instead.
func (*BatchPostReq) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{9}
}

func (x *BatchPostReq) GetPostIds() []int64 {
	if x != nil {
		return x.PostIds
	}
	return nil
}

type BatchPostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*PostData `protobuf:"bytes,1,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *BatchPostResp) Reset() {
	*x = BatchPostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchPostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchPostResp) ProtoMessage() {}

func (x *BatchPostResp) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchPostResp.ProtoReflect.Descriptor instead.
func (*BatchPostResp) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{10}
}

func (x *BatchPostResp) GetInfos() []*PostData {
	if x != nil {
		return x.Infos
	}
	return nil
}

type GetUserPostListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"` // 用户ID
}

func (x *GetUserPostListReq) Reset() {
	*x = GetUserPostListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPostListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPostListReq) ProtoMessage() {}

func (x *GetUserPostListReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPostListReq.ProtoReflect.Descriptor instead.
func (*GetUserPostListReq) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserPostListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserPostListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*PostData `protobuf:"bytes,1,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetUserPostListResp) Reset() {
	*x = GetUserPostListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_post_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPostListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPostListResp) ProtoMessage() {}

func (x *GetUserPostListResp) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPostListResp.ProtoReflect.Descriptor instead.
func (*GetUserPostListResp) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{12}
}

func (x *GetUserPostListResp) GetInfos() []*PostData {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_post_post_proto protoreflect.FileDescriptor

var file_post_post_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x9a, 0x02, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x56, 0x69, 0x65, 0x77, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56, 0x69, 0x65,
	0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x3b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x3b, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x27, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x24, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x22, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x73,
	0x22, 0x35, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x24, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x32, 0xde, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x09, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_post_proto_rawDescOnce sync.Once
	file_post_post_proto_rawDescData = file_post_post_proto_rawDesc
)

func file_post_post_proto_rawDescGZIP() []byte {
	file_post_post_proto_rawDescOnce.Do(func() {
		file_post_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_post_proto_rawDescData)
	})
	return file_post_post_proto_rawDescData
}

var file_post_post_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_post_post_proto_goTypes = []interface{}{
	(*PostData)(nil),            // 0: post.PostData
	(*CreatePostReq)(nil),       // 1: post.CreatePostReq
	(*CreatePostResp)(nil),      // 2: post.CreatePostResp
	(*UpdatePostReq)(nil),       // 3: post.UpdatePostReq
	(*UpdatePostResp)(nil),      // 4: post.UpdatePostResp
	(*DeletePostReq)(nil),       // 5: post.DeletePostReq
	(*DeletePostResp)(nil),      // 6: post.DeletePostResp
	(*GetPostReq)(nil),          // 7: post.GetPostReq
	(*GetPostResp)(nil),         // 8: post.GetPostResp
	(*BatchPostReq)(nil),        // 9: post.BatchPostReq
	(*BatchPostResp)(nil),       // 10: post.BatchPostResp
	(*GetUserPostListReq)(nil),  // 11: post.GetUserPostListReq
	(*GetUserPostListResp)(nil), // 12: post.GetUserPostListResp
}
var file_post_post_proto_depIdxs = []int32{
	0,  // 0: post.CreatePostReq.PostData:type_name -> post.PostData
	0,  // 1: post.UpdatePostReq.PostData:type_name -> post.PostData
	0,  // 2: post.GetPostResp.Info:type_name -> post.PostData
	0,  // 3: post.BatchPostResp.Infos:type_name -> post.PostData
	0,  // 4: post.GetUserPostListResp.Infos:type_name -> post.PostData
	1,  // 5: post.Post.CreatePost:input_type -> post.CreatePostReq
	3,  // 6: post.Post.UpdatePost:input_type -> post.UpdatePostReq
	5,  // 7: post.Post.DeletePost:input_type -> post.DeletePostReq
	7,  // 8: post.Post.GetPost:input_type -> post.GetPostReq
	9,  // 9: post.Post.BatchPost:input_type -> post.BatchPostReq
	11, // 10: post.Post.GetUserPostList:input_type -> post.GetUserPostListReq
	2,  // 11: post.Post.CreatePost:output_type -> post.CreatePostResp
	4,  // 12: post.Post.UpdatePost:output_type -> post.UpdatePostResp
	5,  // 13: post.Post.DeletePost:output_type -> post.DeletePostReq
	8,  // 14: post.Post.GetPost:output_type -> post.GetPostResp
	10, // 15: post.Post.BatchPost:output_type -> post.BatchPostResp
	12, // 16: post.Post.GetUserPostList:output_type -> post.GetUserPostListResp
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_post_post_proto_init() }
func file_post_post_proto_init() {
	if File_post_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostData); i {
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
		file_post_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostReq); i {
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
		file_post_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostResp); i {
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
		file_post_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostReq); i {
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
		file_post_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostResp); i {
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
		file_post_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostReq); i {
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
		file_post_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostResp); i {
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
		file_post_post_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostReq); i {
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
		file_post_post_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostResp); i {
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
		file_post_post_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchPostReq); i {
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
		file_post_post_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchPostResp); i {
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
		file_post_post_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPostListReq); i {
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
		file_post_post_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPostListResp); i {
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
			RawDescriptor: file_post_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_post_proto_goTypes,
		DependencyIndexes: file_post_post_proto_depIdxs,
		MessageInfos:      file_post_post_proto_msgTypes,
	}.Build()
	File_post_post_proto = out.File
	file_post_post_proto_rawDesc = nil
	file_post_post_proto_goTypes = nil
	file_post_post_proto_depIdxs = nil
}
