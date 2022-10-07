// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: movies/v1/movies.proto

package moviesv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Summary string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Rating  uint32 `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *CreateMovieRequest) Reset() {
	*x = CreateMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movies_v1_movies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMovieRequest) ProtoMessage() {}

func (x *CreateMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movies_v1_movies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMovieRequest.ProtoReflect.Descriptor instead.
func (*CreateMovieRequest) Descriptor() ([]byte, []int) {
	return file_movies_v1_movies_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMovieRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateMovieRequest) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *CreateMovieRequest) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type CreateMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Movie `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateMovieResponse) Reset() {
	*x = CreateMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movies_v1_movies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMovieResponse) ProtoMessage() {}

func (x *CreateMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movies_v1_movies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMovieResponse.ProtoReflect.Descriptor instead.
func (*CreateMovieResponse) Descriptor() ([]byte, []int) {
	return file_movies_v1_movies_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMovieResponse) GetData() *Movie {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMovieRequest) Reset() {
	*x = GetMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movies_v1_movies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieRequest) ProtoMessage() {}

func (x *GetMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movies_v1_movies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieRequest.ProtoReflect.Descriptor instead.
func (*GetMovieRequest) Descriptor() ([]byte, []int) {
	return file_movies_v1_movies_proto_rawDescGZIP(), []int{2}
}

func (x *GetMovieRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Movie `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetMovieResponse) Reset() {
	*x = GetMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movies_v1_movies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieResponse) ProtoMessage() {}

func (x *GetMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movies_v1_movies_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieResponse.ProtoReflect.Descriptor instead.
func (*GetMovieResponse) Descriptor() ([]byte, []int) {
	return file_movies_v1_movies_proto_rawDescGZIP(), []int{3}
}

func (x *GetMovieResponse) GetData() *Movie {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListMoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListMoviesRequest) Reset() {
	*x = ListMoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movies_v1_movies_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMoviesRequest) ProtoMessage() {}

func (x *ListMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movies_v1_movies_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMoviesRequest.ProtoReflect.Descriptor instead.
func (*ListMoviesRequest) Descriptor() ([]byte, []int) {
	return file_movies_v1_movies_proto_rawDescGZIP(), []int{4}
}

type ListMoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Movie `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListMoviesResponse) Reset() {
	*x = ListMoviesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movies_v1_movies_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMoviesResponse) ProtoMessage() {}

func (x *ListMoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movies_v1_movies_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMoviesResponse.ProtoReflect.Descriptor instead.
func (*ListMoviesResponse) Descriptor() ([]byte, []int) {
	return file_movies_v1_movies_proto_rawDescGZIP(), []int{5}
}

func (x *ListMoviesResponse) GetData() []*Movie {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Summary string `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	Rating  uint32 `protobuf:"varint,4,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *UpdateMovieRequest) Reset() {
	*x = UpdateMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movies_v1_movies_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMovieRequest) ProtoMessage() {}

func (x *UpdateMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movies_v1_movies_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMovieRequest.ProtoReflect.Descriptor instead.
func (*UpdateMovieRequest) Descriptor() ([]byte, []int) {
	return file_movies_v1_movies_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateMovieRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMovieRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateMovieRequest) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *UpdateMovieRequest) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type UpdateMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Movie `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateMovieResponse) Reset() {
	*x = UpdateMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movies_v1_movies_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMovieResponse) ProtoMessage() {}

func (x *UpdateMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movies_v1_movies_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMovieResponse.ProtoReflect.Descriptor instead.
func (*UpdateMovieResponse) Descriptor() ([]byte, []int) {
	return file_movies_v1_movies_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateMovieResponse) GetData() *Movie {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMovieRequest) Reset() {
	*x = DeleteMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movies_v1_movies_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMovieRequest) ProtoMessage() {}

func (x *DeleteMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movies_v1_movies_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMovieRequest.ProtoReflect.Descriptor instead.
func (*DeleteMovieRequest) Descriptor() ([]byte, []int) {
	return file_movies_v1_movies_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteMovieRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMovieResponse) Reset() {
	*x = DeleteMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movies_v1_movies_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMovieResponse) ProtoMessage() {}

func (x *DeleteMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movies_v1_movies_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMovieResponse.ProtoReflect.Descriptor instead.
func (*DeleteMovieResponse) Descriptor() ([]byte, []int) {
	return file_movies_v1_movies_proto_rawDescGZIP(), []int{9}
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Summary string `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	Rating  uint32 `protobuf:"varint,4,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movies_v1_movies_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_movies_v1_movies_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_movies_v1_movies_proto_rawDescGZIP(), []int{10}
}

func (x *Movie) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Movie) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

var File_movies_v1_movies_proto protoreflect.FileDescriptor

var file_movies_v1_movies_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xad, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x3a, 0xce, 0x01, 0x92, 0x41, 0xca, 0x01, 0x0a, 0x56, 0x2a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x25, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x20, 0x75, 0x73, 0x65, 0xd2, 0x01, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0xd2, 0x01, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0xd2, 0x01, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x32,
	0x70, 0x7b, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x22, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x20, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2c, 0x22,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x3a, 0x22, 0x42, 0x65, 0x6c, 0x61, 0x6a, 0x61,
	0x72, 0x20, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x20, 0x70, 0x65, 0x6e, 0x67, 0x65,
	0x6d, 0x62, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x6b, 0x20, 0x73,
	0x75, 0x70, 0x61, 0x79, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x22, 0x2c, 0x20, 0x22, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x3a, 0x20, 0x31, 0x20,
	0x7d, 0x22, 0xa2, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a,
	0xe4, 0x01, 0x92, 0x41, 0xe0, 0x01, 0x0a, 0x56, 0x2a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x25, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20,
	0x75, 0x73, 0x65, 0xd2, 0x01, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0xd2, 0x01, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0xd2, 0x01, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x32, 0x85,
	0x01, 0x7b, 0x22, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x20, 0x7b, 0x20, 0x22, 0x69, 0x64, 0x22,
	0x3a, 0x20, 0x31, 0x2c, 0x20, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x22, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x20, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x2c, 0x22, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x3a, 0x22, 0x42, 0x65, 0x6c,
	0x61, 0x6a, 0x61, 0x72, 0x20, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x20, 0x70, 0x65,
	0x6e, 0x67, 0x65, 0x6d, 0x62, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x6b, 0x20, 0x73, 0x75, 0x70, 0x61, 0x79, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x22, 0x2c, 0x20, 0x22, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x3a,
	0x20, 0x31, 0x20, 0x7d, 0x20, 0x7d, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x6c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x3b, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0x0a, 0x05,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x32, 0xb7, 0x05,
	0x0a, 0x0d, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x88, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x1d, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x92, 0x41, 0x22, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x7b, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x1a, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x1c, 0x0a, 0x05, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x12, 0x09, 0x47, 0x65, 0x74, 0x20, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2a, 0x08, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x92, 0x41, 0x20, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x12, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x22, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x2a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x92, 0x41, 0x22, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x20, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2a, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0xc2, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa,
	0x02, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0a, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movies_v1_movies_proto_rawDescOnce sync.Once
	file_movies_v1_movies_proto_rawDescData = file_movies_v1_movies_proto_rawDesc
)

func file_movies_v1_movies_proto_rawDescGZIP() []byte {
	file_movies_v1_movies_proto_rawDescOnce.Do(func() {
		file_movies_v1_movies_proto_rawDescData = protoimpl.X.CompressGZIP(file_movies_v1_movies_proto_rawDescData)
	})
	return file_movies_v1_movies_proto_rawDescData
}

var file_movies_v1_movies_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_movies_v1_movies_proto_goTypes = []interface{}{
	(*CreateMovieRequest)(nil),  // 0: movies.v1.CreateMovieRequest
	(*CreateMovieResponse)(nil), // 1: movies.v1.CreateMovieResponse
	(*GetMovieRequest)(nil),     // 2: movies.v1.GetMovieRequest
	(*GetMovieResponse)(nil),    // 3: movies.v1.GetMovieResponse
	(*ListMoviesRequest)(nil),   // 4: movies.v1.ListMoviesRequest
	(*ListMoviesResponse)(nil),  // 5: movies.v1.ListMoviesResponse
	(*UpdateMovieRequest)(nil),  // 6: movies.v1.UpdateMovieRequest
	(*UpdateMovieResponse)(nil), // 7: movies.v1.UpdateMovieResponse
	(*DeleteMovieRequest)(nil),  // 8: movies.v1.DeleteMovieRequest
	(*DeleteMovieResponse)(nil), // 9: movies.v1.DeleteMovieResponse
	(*Movie)(nil),               // 10: movies.v1.Movie
}
var file_movies_v1_movies_proto_depIdxs = []int32{
	10, // 0: movies.v1.CreateMovieResponse.data:type_name -> movies.v1.Movie
	10, // 1: movies.v1.GetMovieResponse.data:type_name -> movies.v1.Movie
	10, // 2: movies.v1.ListMoviesResponse.data:type_name -> movies.v1.Movie
	10, // 3: movies.v1.UpdateMovieResponse.data:type_name -> movies.v1.Movie
	0,  // 4: movies.v1.MoviesService.CreateMovie:input_type -> movies.v1.CreateMovieRequest
	2,  // 5: movies.v1.MoviesService.GetMovie:input_type -> movies.v1.GetMovieRequest
	4,  // 6: movies.v1.MoviesService.ListMovies:input_type -> movies.v1.ListMoviesRequest
	6,  // 7: movies.v1.MoviesService.UpdateMovie:input_type -> movies.v1.UpdateMovieRequest
	8,  // 8: movies.v1.MoviesService.DeleteMovie:input_type -> movies.v1.DeleteMovieRequest
	1,  // 9: movies.v1.MoviesService.CreateMovie:output_type -> movies.v1.CreateMovieResponse
	3,  // 10: movies.v1.MoviesService.GetMovie:output_type -> movies.v1.GetMovieResponse
	5,  // 11: movies.v1.MoviesService.ListMovies:output_type -> movies.v1.ListMoviesResponse
	7,  // 12: movies.v1.MoviesService.UpdateMovie:output_type -> movies.v1.UpdateMovieResponse
	9,  // 13: movies.v1.MoviesService.DeleteMovie:output_type -> movies.v1.DeleteMovieResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_movies_v1_movies_proto_init() }
func file_movies_v1_movies_proto_init() {
	if File_movies_v1_movies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movies_v1_movies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMovieRequest); i {
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
		file_movies_v1_movies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMovieResponse); i {
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
		file_movies_v1_movies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieRequest); i {
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
		file_movies_v1_movies_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieResponse); i {
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
		file_movies_v1_movies_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMoviesRequest); i {
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
		file_movies_v1_movies_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMoviesResponse); i {
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
		file_movies_v1_movies_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMovieRequest); i {
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
		file_movies_v1_movies_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMovieResponse); i {
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
		file_movies_v1_movies_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMovieRequest); i {
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
		file_movies_v1_movies_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMovieResponse); i {
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
		file_movies_v1_movies_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
			RawDescriptor: file_movies_v1_movies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movies_v1_movies_proto_goTypes,
		DependencyIndexes: file_movies_v1_movies_proto_depIdxs,
		MessageInfos:      file_movies_v1_movies_proto_msgTypes,
	}.Build()
	File_movies_v1_movies_proto = out.File
	file_movies_v1_movies_proto_rawDesc = nil
	file_movies_v1_movies_proto_goTypes = nil
	file_movies_v1_movies_proto_depIdxs = nil
}
