// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: franchise.proto

package gen

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Franchise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string               `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	CreateDate  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	ExpiresDate *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expires_date,json=expiresDate,proto3" json:"expires_date,omitempty"`
	Name        string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Email       string               `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Location    *Location            `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	CompanyId   uint32               `protobuf:"varint,7,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Metadata    *Metadata            `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Franchise) Reset() {
	*x = Franchise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Franchise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Franchise) ProtoMessage() {}

func (x *Franchise) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Franchise.ProtoReflect.Descriptor instead.
func (*Franchise) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{0}
}

func (x *Franchise) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Franchise) GetCreateDate() *timestamp.Timestamp {
	if x != nil {
		return x.CreateDate
	}
	return nil
}

func (x *Franchise) GetExpiresDate() *timestamp.Timestamp {
	if x != nil {
		return x.ExpiresDate
	}
	return nil
}

func (x *Franchise) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Franchise) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Franchise) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Franchise) GetCompanyId() uint32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *Franchise) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City    string   `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Address string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Country *Country `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Location) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Location) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{2}
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateFranchiseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *CreateFranchiseRequest) Reset() {
	*x = CreateFranchiseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFranchiseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFranchiseRequest) ProtoMessage() {}

func (x *CreateFranchiseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFranchiseRequest.ProtoReflect.Descriptor instead.
func (*CreateFranchiseRequest) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{3}
}

func (x *CreateFranchiseRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type CreateFranchiseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateFranchiseResponse) Reset() {
	*x = CreateFranchiseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFranchiseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFranchiseResponse) ProtoMessage() {}

func (x *CreateFranchiseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFranchiseResponse.ProtoReflect.Descriptor instead.
func (*CreateFranchiseResponse) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{4}
}

func (x *CreateFranchiseResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetFranchiseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId uint32 `protobuf:"varint,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *GetFranchiseRequest) Reset() {
	*x = GetFranchiseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFranchiseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFranchiseRequest) ProtoMessage() {}

func (x *GetFranchiseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFranchiseRequest.ProtoReflect.Descriptor instead.
func (*GetFranchiseRequest) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{5}
}

func (x *GetFranchiseRequest) GetCompanyId() uint32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type GetFranchiseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Franchise []*Franchise `protobuf:"bytes,1,rep,name=franchise,proto3" json:"franchise,omitempty"`
}

func (x *GetFranchiseResponse) Reset() {
	*x = GetFranchiseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFranchiseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFranchiseResponse) ProtoMessage() {}

func (x *GetFranchiseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFranchiseResponse.ProtoReflect.Descriptor instead.
func (*GetFranchiseResponse) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{6}
}

func (x *GetFranchiseResponse) GetFranchise() []*Franchise {
	if x != nil {
		return x.Franchise
	}
	return nil
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName string `protobuf:"bytes,1,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	IpAddress  string `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{7}
}

func (x *Endpoint) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *Endpoint) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol     string      `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Availability string      `protobuf:"bytes,2,opt,name=availability,proto3" json:"availability,omitempty"`
	Endpoints    []*Endpoint `protobuf:"bytes,3,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	FranchiseId  uint32      `protobuf:"varint,4,opt,name=franchise_id,json=franchiseId,proto3" json:"franchise_id,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{8}
}

func (x *Metadata) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Metadata) GetAvailability() string {
	if x != nil {
		return x.Availability
	}
	return ""
}

func (x *Metadata) GetEndpoints() []*Endpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *Metadata) GetFranchiseId() uint32 {
	if x != nil {
		return x.FranchiseId
	}
	return 0
}

type CreateMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host        string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	FranchiseId uint32 `protobuf:"varint,2,opt,name=franchise_id,json=franchiseId,proto3" json:"franchise_id,omitempty"`
	CompanyId   uint32 `protobuf:"varint,3,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *CreateMetadataRequest) Reset() {
	*x = CreateMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMetadataRequest) ProtoMessage() {}

func (x *CreateMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMetadataRequest.ProtoReflect.Descriptor instead.
func (*CreateMetadataRequest) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{9}
}

func (x *CreateMetadataRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *CreateMetadataRequest) GetFranchiseId() uint32 {
	if x != nil {
		return x.FranchiseId
	}
	return 0
}

func (x *CreateMetadataRequest) GetCompanyId() uint32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type CreateMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateMetadataResponse) Reset() {
	*x = CreateMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMetadataResponse) ProtoMessage() {}

func (x *CreateMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMetadataResponse.ProtoReflect.Descriptor instead.
func (*CreateMetadataResponse) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{10}
}

func (x *CreateMetadataResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId uint32 `protobuf:"varint,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *GetMetadataRequest) Reset() {
	*x = GetMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetadataRequest) ProtoMessage() {}

func (x *GetMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetadataRequest.ProtoReflect.Descriptor instead.
func (*GetMetadataRequest) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{11}
}

func (x *GetMetadataRequest) GetCompanyId() uint32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type GetMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata []*Metadata `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *GetMetadataResponse) Reset() {
	*x = GetMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_franchise_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetadataResponse) ProtoMessage() {}

func (x *GetMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_franchise_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetadataResponse.ProtoReflect.Descriptor instead.
func (*GetMetadataResponse) Descriptor() ([]byte, []int) {
	return file_franchise_proto_rawDescGZIP(), []int{12}
}

func (x *GetMetadataResponse) GetMetadata() []*Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_franchise_proto protoreflect.FileDescriptor

var file_franchise_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb0, 0x02, 0x0a, 0x09, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x3d, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5c, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x22, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x22, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x22, 0x2f, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x34, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x52, 0x09,
	0x66, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x22, 0x4a, 0x0a, 0x08, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x22,
	0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x27, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x66, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x49, 0x64, 0x22, 0x6d,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x66, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x2e, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x33, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x64, 0x22, 0x3c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x8e, 0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x95, 0x01, 0x0a, 0x10, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x12, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_franchise_proto_rawDescOnce sync.Once
	file_franchise_proto_rawDescData = file_franchise_proto_rawDesc
)

func file_franchise_proto_rawDescGZIP() []byte {
	file_franchise_proto_rawDescOnce.Do(func() {
		file_franchise_proto_rawDescData = protoimpl.X.CompressGZIP(file_franchise_proto_rawDescData)
	})
	return file_franchise_proto_rawDescData
}

var file_franchise_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_franchise_proto_goTypes = []interface{}{
	(*Franchise)(nil),               // 0: Franchise
	(*Location)(nil),                // 1: Location
	(*Country)(nil),                 // 2: Country
	(*CreateFranchiseRequest)(nil),  // 3: CreateFranchiseRequest
	(*CreateFranchiseResponse)(nil), // 4: CreateFranchiseResponse
	(*GetFranchiseRequest)(nil),     // 5: GetFranchiseRequest
	(*GetFranchiseResponse)(nil),    // 6: GetFranchiseResponse
	(*Endpoint)(nil),                // 7: Endpoint
	(*Metadata)(nil),                // 8: Metadata
	(*CreateMetadataRequest)(nil),   // 9: CreateMetadataRequest
	(*CreateMetadataResponse)(nil),  // 10: CreateMetadataResponse
	(*GetMetadataRequest)(nil),      // 11: GetMetadataRequest
	(*GetMetadataResponse)(nil),     // 12: GetMetadataResponse
	(*timestamp.Timestamp)(nil),     // 13: google.protobuf.Timestamp
}
var file_franchise_proto_depIdxs = []int32{
	13, // 0: Franchise.create_date:type_name -> google.protobuf.Timestamp
	13, // 1: Franchise.expires_date:type_name -> google.protobuf.Timestamp
	1,  // 2: Franchise.location:type_name -> Location
	8,  // 3: Franchise.metadata:type_name -> Metadata
	2,  // 4: Location.country:type_name -> Country
	0,  // 5: GetFranchiseResponse.franchise:type_name -> Franchise
	7,  // 6: Metadata.endpoints:type_name -> Endpoint
	8,  // 7: GetMetadataResponse.metadata:type_name -> Metadata
	11, // 8: MetadataService.GetMetadata:input_type -> GetMetadataRequest
	9,  // 9: MetadataService.CreateMetadata:input_type -> CreateMetadataRequest
	5,  // 10: FranchiseService.GetFranchise:input_type -> GetFranchiseRequest
	3,  // 11: FranchiseService.CreateFranchise:input_type -> CreateFranchiseRequest
	12, // 12: MetadataService.GetMetadata:output_type -> GetMetadataResponse
	10, // 13: MetadataService.CreateMetadata:output_type -> CreateMetadataResponse
	6,  // 14: FranchiseService.GetFranchise:output_type -> GetFranchiseResponse
	4,  // 15: FranchiseService.CreateFranchise:output_type -> CreateFranchiseResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_franchise_proto_init() }
func file_franchise_proto_init() {
	if File_franchise_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_franchise_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Franchise); i {
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
		file_franchise_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_franchise_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
		file_franchise_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFranchiseRequest); i {
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
		file_franchise_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFranchiseResponse); i {
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
		file_franchise_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFranchiseRequest); i {
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
		file_franchise_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFranchiseResponse); i {
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
		file_franchise_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_franchise_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_franchise_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMetadataRequest); i {
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
		file_franchise_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMetadataResponse); i {
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
		file_franchise_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetadataRequest); i {
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
		file_franchise_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetadataResponse); i {
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
			RawDescriptor: file_franchise_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_franchise_proto_goTypes,
		DependencyIndexes: file_franchise_proto_depIdxs,
		MessageInfos:      file_franchise_proto_msgTypes,
	}.Build()
	File_franchise_proto = out.File
	file_franchise_proto_rawDesc = nil
	file_franchise_proto_goTypes = nil
	file_franchise_proto_depIdxs = nil
}
