// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: proto/shortener.proto

package grpcShortener

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

// GetStats
type GetStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatsRequest) Reset() {
	*x = GetStatsRequest{}
	mi := &file_proto_shortener_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsRequest) ProtoMessage() {}

func (x *GetStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsRequest.ProtoReflect.Descriptor instead.
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{0}
}

type GetStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountUrls  uint32 `protobuf:"varint,1,opt,name=count_urls,json=countUrls,proto3" json:"count_urls,omitempty"`
	CountUsers uint32 `protobuf:"varint,2,opt,name=count_users,json=countUsers,proto3" json:"count_users,omitempty"`
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	mi := &file_proto_shortener_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatsResponse) GetCountUrls() uint32 {
	if x != nil {
		return x.CountUrls
	}
	return 0
}

func (x *GetStatsResponse) GetCountUsers() uint32 {
	if x != nil {
		return x.CountUsers
	}
	return 0
}

// Ping
type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	mi := &file_proto_shortener_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{2}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	mi := &file_proto_shortener_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{3}
}

// Create
type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Original string `protobuf:"bytes,1,opt,name=original,proto3" json:"original,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_proto_shortener_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRequest) GetOriginal() string {
	if x != nil {
		return x.Original
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Short string `protobuf:"bytes,1,opt,name=short,proto3" json:"short,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_proto_shortener_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *CreateResponse) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

type CreateBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*CreateBatchRequestItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *CreateBatchRequest) Reset() {
	*x = CreateBatchRequest{}
	mi := &file_proto_shortener_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchRequest) ProtoMessage() {}

func (x *CreateBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchRequest.ProtoReflect.Descriptor instead.
func (*CreateBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *CreateBatchRequest) GetList() []*CreateBatchRequestItem {
	if x != nil {
		return x.List
	}
	return nil
}

type CreateBatchRequestItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OriginalUrl   string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *CreateBatchRequestItem) Reset() {
	*x = CreateBatchRequestItem{}
	mi := &file_proto_shortener_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBatchRequestItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchRequestItem) ProtoMessage() {}

func (x *CreateBatchRequestItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchRequestItem.ProtoReflect.Descriptor instead.
func (*CreateBatchRequestItem) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *CreateBatchRequestItem) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CreateBatchRequestItem) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type CreateBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*CreateBatchResponseItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *CreateBatchResponse) Reset() {
	*x = CreateBatchResponse{}
	mi := &file_proto_shortener_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchResponse) ProtoMessage() {}

func (x *CreateBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchResponse.ProtoReflect.Descriptor instead.
func (*CreateBatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *CreateBatchResponse) GetList() []*CreateBatchResponseItem {
	if x != nil {
		return x.List
	}
	return nil
}

type CreateBatchResponseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortUrl      string `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *CreateBatchResponseItem) Reset() {
	*x = CreateBatchResponseItem{}
	mi := &file_proto_shortener_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBatchResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchResponseItem) ProtoMessage() {}

func (x *CreateBatchResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchResponseItem.ProtoReflect.Descriptor instead.
func (*CreateBatchResponseItem) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *CreateBatchResponseItem) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CreateBatchResponseItem) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

// Get
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Short string `protobuf:"bytes,1,opt,name=short,proto3" json:"short,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_proto_shortener_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{10}
}

func (x *GetRequest) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Original string `protobuf:"bytes,1,opt,name=original,proto3" json:"original,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_proto_shortener_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{11}
}

func (x *GetResponse) GetOriginal() string {
	if x != nil {
		return x.Original
	}
	return ""
}

// UserUrls
type UserUrlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserUrlsRequest) Reset() {
	*x = UserUrlsRequest{}
	mi := &file_proto_shortener_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserUrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUrlsRequest) ProtoMessage() {}

func (x *UserUrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUrlsRequest.ProtoReflect.Descriptor instead.
func (*UserUrlsRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{12}
}

type UserUrlsResponseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Short    string `protobuf:"bytes,1,opt,name=short,proto3" json:"short,omitempty"`
	Original string `protobuf:"bytes,2,opt,name=original,proto3" json:"original,omitempty"`
}

func (x *UserUrlsResponseItem) Reset() {
	*x = UserUrlsResponseItem{}
	mi := &file_proto_shortener_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserUrlsResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUrlsResponseItem) ProtoMessage() {}

func (x *UserUrlsResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUrlsResponseItem.ProtoReflect.Descriptor instead.
func (*UserUrlsResponseItem) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{13}
}

func (x *UserUrlsResponseItem) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *UserUrlsResponseItem) GetOriginal() string {
	if x != nil {
		return x.Original
	}
	return ""
}

type UserUrlsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*UserUrlsResponseItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UserUrlsResponse) Reset() {
	*x = UserUrlsResponse{}
	mi := &file_proto_shortener_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserUrlsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUrlsResponse) ProtoMessage() {}

func (x *UserUrlsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUrlsResponse.ProtoReflect.Descriptor instead.
func (*UserUrlsResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{14}
}

func (x *UserUrlsResponse) GetList() []*UserUrlsResponseItem {
	if x != nil {
		return x.List
	}
	return nil
}

var File_proto_shortener_proto protoreflect.FileDescriptor

var file_proto_shortener_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x0d,
	0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a,
	0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x26, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x22, 0x50, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x52, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x22, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x22,
	0x29, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a,
	0x14, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x4c, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xfe, 0x02, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x22, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_shortener_proto_rawDescOnce sync.Once
	file_proto_shortener_proto_rawDescData = file_proto_shortener_proto_rawDesc
)

func file_proto_shortener_proto_rawDescGZIP() []byte {
	file_proto_shortener_proto_rawDescOnce.Do(func() {
		file_proto_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shortener_proto_rawDescData)
	})
	return file_proto_shortener_proto_rawDescData
}

var file_proto_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_shortener_proto_goTypes = []any{
	(*GetStatsRequest)(nil),         // 0: grpc_shortener.GetStatsRequest
	(*GetStatsResponse)(nil),        // 1: grpc_shortener.GetStatsResponse
	(*PingRequest)(nil),             // 2: grpc_shortener.PingRequest
	(*PingResponse)(nil),            // 3: grpc_shortener.PingResponse
	(*CreateRequest)(nil),           // 4: grpc_shortener.CreateRequest
	(*CreateResponse)(nil),          // 5: grpc_shortener.CreateResponse
	(*CreateBatchRequest)(nil),      // 6: grpc_shortener.CreateBatchRequest
	(*CreateBatchRequestItem)(nil),  // 7: grpc_shortener.CreateBatchRequestItem
	(*CreateBatchResponse)(nil),     // 8: grpc_shortener.CreateBatchResponse
	(*CreateBatchResponseItem)(nil), // 9: grpc_shortener.CreateBatchResponseItem
	(*GetRequest)(nil),              // 10: grpc_shortener.GetRequest
	(*GetResponse)(nil),             // 11: grpc_shortener.GetResponse
	(*UserUrlsRequest)(nil),         // 12: grpc_shortener.UserUrlsRequest
	(*UserUrlsResponseItem)(nil),    // 13: grpc_shortener.UserUrlsResponseItem
	(*UserUrlsResponse)(nil),        // 14: grpc_shortener.UserUrlsResponse
}
var file_proto_shortener_proto_depIdxs = []int32{
	7,  // 0: grpc_shortener.CreateBatchRequest.list:type_name -> grpc_shortener.CreateBatchRequestItem
	9,  // 1: grpc_shortener.CreateBatchResponse.list:type_name -> grpc_shortener.CreateBatchResponseItem
	13, // 2: grpc_shortener.UserUrlsResponse.list:type_name -> grpc_shortener.UserUrlsResponseItem
	0,  // 3: grpc_shortener.Shortener.GetStats:input_type -> grpc_shortener.GetStatsRequest
	2,  // 4: grpc_shortener.Shortener.Ping:input_type -> grpc_shortener.PingRequest
	4,  // 5: grpc_shortener.Shortener.Create:input_type -> grpc_shortener.CreateRequest
	6,  // 6: grpc_shortener.Shortener.CreateBatch:input_type -> grpc_shortener.CreateBatchRequest
	10, // 7: grpc_shortener.Shortener.Get:input_type -> grpc_shortener.GetRequest
	1,  // 8: grpc_shortener.Shortener.GetStats:output_type -> grpc_shortener.GetStatsResponse
	3,  // 9: grpc_shortener.Shortener.Ping:output_type -> grpc_shortener.PingResponse
	5,  // 10: grpc_shortener.Shortener.Create:output_type -> grpc_shortener.CreateResponse
	8,  // 11: grpc_shortener.Shortener.CreateBatch:output_type -> grpc_shortener.CreateBatchResponse
	11, // 12: grpc_shortener.Shortener.Get:output_type -> grpc_shortener.GetResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_shortener_proto_init() }
func file_proto_shortener_proto_init() {
	if File_proto_shortener_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shortener_proto_goTypes,
		DependencyIndexes: file_proto_shortener_proto_depIdxs,
		MessageInfos:      file_proto_shortener_proto_msgTypes,
	}.Build()
	File_proto_shortener_proto = out.File
	file_proto_shortener_proto_rawDesc = nil
	file_proto_shortener_proto_goTypes = nil
	file_proto_shortener_proto_depIdxs = nil
}