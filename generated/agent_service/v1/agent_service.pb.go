// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: v1/agent_service.proto

package v1

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

// Create
type CreateAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	AgentName string `protobuf:"bytes,2,opt,name=agent_name,json=agentName,proto3" json:"agent_name,omitempty"`
}

func (x *CreateAgentRequest) Reset() {
	*x = CreateAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_agent_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAgentRequest) ProtoMessage() {}

func (x *CreateAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_agent_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAgentRequest.ProtoReflect.Descriptor instead.
func (*CreateAgentRequest) Descriptor() ([]byte, []int) {
	return file_v1_agent_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAgentRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CreateAgentRequest) GetAgentName() string {
	if x != nil {
		return x.AgentName
	}
	return ""
}

type CreateAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CompanyId string  `protobuf:"bytes,3,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Secret    string  `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	Status    *string `protobuf:"bytes,99,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *CreateAgentResponse) Reset() {
	*x = CreateAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_agent_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAgentResponse) ProtoMessage() {}

func (x *CreateAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_agent_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAgentResponse.ProtoReflect.Descriptor instead.
func (*CreateAgentResponse) Descriptor() ([]byte, []int) {
	return file_v1_agent_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAgentResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateAgentResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAgentResponse) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CreateAgentResponse) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *CreateAgentResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

// Validate
type ValidateAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Secret string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *ValidateAgentRequest) Reset() {
	*x = ValidateAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_agent_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateAgentRequest) ProtoMessage() {}

func (x *ValidateAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_agent_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateAgentRequest.ProtoReflect.Descriptor instead.
func (*ValidateAgentRequest) Descriptor() ([]byte, []int) {
	return file_v1_agent_service_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateAgentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ValidateAgentRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type ValidateAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid  bool    `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Status *string `protobuf:"bytes,99,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *ValidateAgentResponse) Reset() {
	*x = ValidateAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_agent_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateAgentResponse) ProtoMessage() {}

func (x *ValidateAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_agent_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateAgentResponse.ProtoReflect.Descriptor instead.
func (*ValidateAgentResponse) Descriptor() ([]byte, []int) {
	return file_v1_agent_service_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateAgentResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ValidateAgentResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

// Register
type RegisterAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SelfUpdate bool   `protobuf:"varint,2,opt,name=self_update,json=selfUpdate,proto3" json:"self_update,omitempty"`
}

func (x *RegisterAgentRequest) Reset() {
	*x = RegisterAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_agent_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAgentRequest) ProtoMessage() {}

func (x *RegisterAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_agent_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAgentRequest.ProtoReflect.Descriptor instead.
func (*RegisterAgentRequest) Descriptor() ([]byte, []int) {
	return file_v1_agent_service_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterAgentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterAgentRequest) GetSelfUpdate() bool {
	if x != nil {
		return x.SelfUpdate
	}
	return false
}

type RegisterAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId        string  `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ChannelSecret    string  `protobuf:"bytes,2,opt,name=channel_secret,json=channelSecret,proto3" json:"channel_secret,omitempty"`
	SelfUpdatedId    *string `protobuf:"bytes,10,opt,name=self_updated_id,json=selfUpdatedId,proto3,oneof" json:"self_updated_id,omitempty"`
	SelfUpdateSecret *string `protobuf:"bytes,11,opt,name=self_update_secret,json=selfUpdateSecret,proto3,oneof" json:"self_update_secret,omitempty"`
	Status           *string `protobuf:"bytes,99,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *RegisterAgentResponse) Reset() {
	*x = RegisterAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_agent_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAgentResponse) ProtoMessage() {}

func (x *RegisterAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_agent_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAgentResponse.ProtoReflect.Descriptor instead.
func (*RegisterAgentResponse) Descriptor() ([]byte, []int) {
	return file_v1_agent_service_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterAgentResponse) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *RegisterAgentResponse) GetChannelSecret() string {
	if x != nil {
		return x.ChannelSecret
	}
	return ""
}

func (x *RegisterAgentResponse) GetSelfUpdatedId() string {
	if x != nil && x.SelfUpdatedId != nil {
		return *x.SelfUpdatedId
	}
	return ""
}

func (x *RegisterAgentResponse) GetSelfUpdateSecret() string {
	if x != nil && x.SelfUpdateSecret != nil {
		return *x.SelfUpdateSecret
	}
	return ""
}

func (x *RegisterAgentResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

var File_v1_agent_service_proto protoreflect.FileDescriptor

var file_v1_agent_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x52, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x63, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3e, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x55, 0x0a, 0x15, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x63, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88,
	0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x47, 0x0a,
	0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x65, 0x6c, 0x66,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x90, 0x02, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0d, 0x73, 0x65, 0x6c, 0x66, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x12, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x10, 0x73, 0x65, 0x6c, 0x66, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x63, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x73, 0x65, 0x6c, 0x66,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9c, 0x02, 0x0a, 0x0c, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5a, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x23, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0d,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_agent_service_proto_rawDescOnce sync.Once
	file_v1_agent_service_proto_rawDescData = file_v1_agent_service_proto_rawDesc
)

func file_v1_agent_service_proto_rawDescGZIP() []byte {
	file_v1_agent_service_proto_rawDescOnce.Do(func() {
		file_v1_agent_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_agent_service_proto_rawDescData)
	})
	return file_v1_agent_service_proto_rawDescData
}

var file_v1_agent_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_agent_service_proto_goTypes = []interface{}{
	(*CreateAgentRequest)(nil),    // 0: agent_service.CreateAgentRequest
	(*CreateAgentResponse)(nil),   // 1: agent_service.CreateAgentResponse
	(*ValidateAgentRequest)(nil),  // 2: agent_service.ValidateAgentRequest
	(*ValidateAgentResponse)(nil), // 3: agent_service.ValidateAgentResponse
	(*RegisterAgentRequest)(nil),  // 4: agent_service.RegisterAgentRequest
	(*RegisterAgentResponse)(nil), // 5: agent_service.RegisterAgentResponse
}
var file_v1_agent_service_proto_depIdxs = []int32{
	0, // 0: agent_service.AgentService.CreateAgent:input_type -> agent_service.CreateAgentRequest
	2, // 1: agent_service.AgentService.ValidateAgent:input_type -> agent_service.ValidateAgentRequest
	4, // 2: agent_service.AgentService.RegisterAgent:input_type -> agent_service.RegisterAgentRequest
	1, // 3: agent_service.AgentService.CreateAgent:output_type -> agent_service.CreateAgentResponse
	3, // 4: agent_service.AgentService.ValidateAgent:output_type -> agent_service.ValidateAgentResponse
	5, // 5: agent_service.AgentService.RegisterAgent:output_type -> agent_service.RegisterAgentResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_agent_service_proto_init() }
func file_v1_agent_service_proto_init() {
	if File_v1_agent_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_agent_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAgentRequest); i {
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
		file_v1_agent_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAgentResponse); i {
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
		file_v1_agent_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateAgentRequest); i {
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
		file_v1_agent_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateAgentResponse); i {
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
		file_v1_agent_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAgentRequest); i {
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
		file_v1_agent_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAgentResponse); i {
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
	file_v1_agent_service_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_v1_agent_service_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_v1_agent_service_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_agent_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_agent_service_proto_goTypes,
		DependencyIndexes: file_v1_agent_service_proto_depIdxs,
		MessageInfos:      file_v1_agent_service_proto_msgTypes,
	}.Build()
	File_v1_agent_service_proto = out.File
	file_v1_agent_service_proto_rawDesc = nil
	file_v1_agent_service_proto_goTypes = nil
	file_v1_agent_service_proto_depIdxs = nil
}
