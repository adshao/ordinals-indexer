// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: inscription/v1/inscription.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetInscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InscriptionId int64 `protobuf:"varint,1,opt,name=inscription_id,json=inscriptionId,proto3" json:"inscription_id,omitempty"`
}

func (x *GetInscriptionRequest) Reset() {
	*x = GetInscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inscription_v1_inscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInscriptionRequest) ProtoMessage() {}

func (x *GetInscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inscription_v1_inscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInscriptionRequest.ProtoReflect.Descriptor instead.
func (*GetInscriptionRequest) Descriptor() ([]byte, []int) {
	return file_inscription_v1_inscription_proto_rawDescGZIP(), []int{0}
}

func (x *GetInscriptionRequest) GetInscriptionId() int64 {
	if x != nil {
		return x.InscriptionId
	}
	return 0
}

type InscriptionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InscriptionId int64                  `protobuf:"varint,2,opt,name=inscription_id,json=inscriptionId,proto3" json:"inscription_id,omitempty"`
	Uid           string                 `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Address       string                 `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	OutputValue   uint64                 `protobuf:"varint,5,opt,name=output_value,json=outputValue,proto3" json:"output_value,omitempty"`
	ContentLength uint64                 `protobuf:"varint,6,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	ContentType   string                 `protobuf:"bytes,7,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	GenesisHeight uint64                 `protobuf:"varint,9,opt,name=genesis_height,json=genesisHeight,proto3" json:"genesis_height,omitempty"`
	GenesisFee    uint64                 `protobuf:"varint,10,opt,name=genesis_fee,json=genesisFee,proto3" json:"genesis_fee,omitempty"`
	GenesisTx     string                 `protobuf:"bytes,11,opt,name=genesis_tx,json=genesisTx,proto3" json:"genesis_tx,omitempty"`
	Location      string                 `protobuf:"bytes,12,opt,name=location,proto3" json:"location,omitempty"`
	Output        string                 `protobuf:"bytes,13,opt,name=output,proto3" json:"output,omitempty"`
	Offset        uint64                 `protobuf:"varint,14,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *InscriptionMessage) Reset() {
	*x = InscriptionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inscription_v1_inscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InscriptionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InscriptionMessage) ProtoMessage() {}

func (x *InscriptionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_inscription_v1_inscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InscriptionMessage.ProtoReflect.Descriptor instead.
func (*InscriptionMessage) Descriptor() ([]byte, []int) {
	return file_inscription_v1_inscription_proto_rawDescGZIP(), []int{1}
}

func (x *InscriptionMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InscriptionMessage) GetInscriptionId() int64 {
	if x != nil {
		return x.InscriptionId
	}
	return 0
}

func (x *InscriptionMessage) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *InscriptionMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InscriptionMessage) GetOutputValue() uint64 {
	if x != nil {
		return x.OutputValue
	}
	return 0
}

func (x *InscriptionMessage) GetContentLength() uint64 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

func (x *InscriptionMessage) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *InscriptionMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *InscriptionMessage) GetGenesisHeight() uint64 {
	if x != nil {
		return x.GenesisHeight
	}
	return 0
}

func (x *InscriptionMessage) GetGenesisFee() uint64 {
	if x != nil {
		return x.GenesisFee
	}
	return 0
}

func (x *InscriptionMessage) GetGenesisTx() string {
	if x != nil {
		return x.GenesisTx
	}
	return ""
}

func (x *InscriptionMessage) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *InscriptionMessage) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *InscriptionMessage) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetInscriptionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *InscriptionMessage `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetInscriptionReply) Reset() {
	*x = GetInscriptionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inscription_v1_inscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInscriptionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInscriptionReply) ProtoMessage() {}

func (x *GetInscriptionReply) ProtoReflect() protoreflect.Message {
	mi := &file_inscription_v1_inscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInscriptionReply.ProtoReflect.Descriptor instead.
func (*GetInscriptionReply) Descriptor() ([]byte, []int) {
	return file_inscription_v1_inscription_proto_rawDescGZIP(), []int{2}
}

func (x *GetInscriptionReply) GetData() *InscriptionMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListInscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderBy string `protobuf:"bytes,1,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Limit   uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  uint64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListInscriptionRequest) Reset() {
	*x = ListInscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inscription_v1_inscription_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInscriptionRequest) ProtoMessage() {}

func (x *ListInscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inscription_v1_inscription_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInscriptionRequest.ProtoReflect.Descriptor instead.
func (*ListInscriptionRequest) Descriptor() ([]byte, []int) {
	return file_inscription_v1_inscription_proto_rawDescGZIP(), []int{3}
}

func (x *ListInscriptionRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListInscriptionRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListInscriptionRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListInscriptionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []*InscriptionMessage `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Paging *Paging               `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *ListInscriptionReply) Reset() {
	*x = ListInscriptionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inscription_v1_inscription_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInscriptionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInscriptionReply) ProtoMessage() {}

func (x *ListInscriptionReply) ProtoReflect() protoreflect.Message {
	mi := &file_inscription_v1_inscription_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInscriptionReply.ProtoReflect.Descriptor instead.
func (*ListInscriptionReply) Descriptor() ([]byte, []int) {
	return file_inscription_v1_inscription_proto_rawDescGZIP(), []int{4}
}

func (x *ListInscriptionReply) GetData() []*InscriptionMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListInscriptionReply) GetPaging() *Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type Paging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount uint64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Count      uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Paging) Reset() {
	*x = Paging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inscription_v1_inscription_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paging) ProtoMessage() {}

func (x *Paging) ProtoReflect() protoreflect.Message {
	mi := &file_inscription_v1_inscription_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paging.ProtoReflect.Descriptor instead.
func (*Paging) Descriptor() ([]byte, []int) {
	return file_inscription_v1_inscription_proto_rawDescGZIP(), []int{5}
}

func (x *Paging) GetTotalCount() uint64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *Paging) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_inscription_v1_inscription_proto protoreflect.FileDescriptor

var file_inscription_v1_inscription_proto_rawDesc = []byte{
	0x0a, 0x20, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xd1, 0x03, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d,
	0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x46, 0x65, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x74, 0x78, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x54, 0x78, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x51, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x61, 0x0a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0x86, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x3f, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xa3, 0x02, 0x0a, 0x0b, 0x49, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x8f, 0x01, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x81, 0x01, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x50, 0x0a, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x73, 0x68, 0x61, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x6c, 0x73, 0x2d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inscription_v1_inscription_proto_rawDescOnce sync.Once
	file_inscription_v1_inscription_proto_rawDescData = file_inscription_v1_inscription_proto_rawDesc
)

func file_inscription_v1_inscription_proto_rawDescGZIP() []byte {
	file_inscription_v1_inscription_proto_rawDescOnce.Do(func() {
		file_inscription_v1_inscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_inscription_v1_inscription_proto_rawDescData)
	})
	return file_inscription_v1_inscription_proto_rawDescData
}

var file_inscription_v1_inscription_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_inscription_v1_inscription_proto_goTypes = []interface{}{
	(*GetInscriptionRequest)(nil),  // 0: api.inscription.v1.GetInscriptionRequest
	(*InscriptionMessage)(nil),     // 1: api.inscription.v1.InscriptionMessage
	(*GetInscriptionReply)(nil),    // 2: api.inscription.v1.GetInscriptionReply
	(*ListInscriptionRequest)(nil), // 3: api.inscription.v1.ListInscriptionRequest
	(*ListInscriptionReply)(nil),   // 4: api.inscription.v1.ListInscriptionReply
	(*Paging)(nil),                 // 5: api.inscription.v1.Paging
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_inscription_v1_inscription_proto_depIdxs = []int32{
	6, // 0: api.inscription.v1.InscriptionMessage.timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: api.inscription.v1.GetInscriptionReply.data:type_name -> api.inscription.v1.InscriptionMessage
	1, // 2: api.inscription.v1.ListInscriptionReply.data:type_name -> api.inscription.v1.InscriptionMessage
	5, // 3: api.inscription.v1.ListInscriptionReply.paging:type_name -> api.inscription.v1.Paging
	0, // 4: api.inscription.v1.Inscription.GetInscription:input_type -> api.inscription.v1.GetInscriptionRequest
	3, // 5: api.inscription.v1.Inscription.ListInscription:input_type -> api.inscription.v1.ListInscriptionRequest
	2, // 6: api.inscription.v1.Inscription.GetInscription:output_type -> api.inscription.v1.GetInscriptionReply
	4, // 7: api.inscription.v1.Inscription.ListInscription:output_type -> api.inscription.v1.ListInscriptionReply
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_inscription_v1_inscription_proto_init() }
func file_inscription_v1_inscription_proto_init() {
	if File_inscription_v1_inscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inscription_v1_inscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInscriptionRequest); i {
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
		file_inscription_v1_inscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InscriptionMessage); i {
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
		file_inscription_v1_inscription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInscriptionReply); i {
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
		file_inscription_v1_inscription_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInscriptionRequest); i {
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
		file_inscription_v1_inscription_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInscriptionReply); i {
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
		file_inscription_v1_inscription_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paging); i {
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
			RawDescriptor: file_inscription_v1_inscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inscription_v1_inscription_proto_goTypes,
		DependencyIndexes: file_inscription_v1_inscription_proto_depIdxs,
		MessageInfos:      file_inscription_v1_inscription_proto_msgTypes,
	}.Build()
	File_inscription_v1_inscription_proto = out.File
	file_inscription_v1_inscription_proto_rawDesc = nil
	file_inscription_v1_inscription_proto_goTypes = nil
	file_inscription_v1_inscription_proto_depIdxs = nil
}