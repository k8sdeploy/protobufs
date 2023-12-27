// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.1
// source: v1/hook_service.proto

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

type UpdateDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Image     string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Tag       string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	Hash      string `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *UpdateDeploymentRequest) Reset() {
	*x = UpdateDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_hook_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeploymentRequest) ProtoMessage() {}

func (x *UpdateDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hook_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeploymentRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_v1_hook_service_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateDeploymentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDeploymentRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *UpdateDeploymentRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UpdateDeploymentRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *UpdateDeploymentRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type UpdateDeploymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateDeploymentResponse) Reset() {
	*x = UpdateDeploymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_hook_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeploymentResponse) ProtoMessage() {}

func (x *UpdateDeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hook_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeploymentResponse.ProtoReflect.Descriptor instead.
func (*UpdateDeploymentResponse) Descriptor() ([]byte, []int) {
	return file_v1_hook_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDeploymentResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_v1_hook_service_proto protoreflect.FileDescriptor

var file_v1_hook_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22,
	0x32, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0xd6, 0x01, 0x0a, 0x0b, 0x48, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e,
	0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x38, 0x73, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_hook_service_proto_rawDescOnce sync.Once
	file_v1_hook_service_proto_rawDescData = file_v1_hook_service_proto_rawDesc
)

func file_v1_hook_service_proto_rawDescGZIP() []byte {
	file_v1_hook_service_proto_rawDescOnce.Do(func() {
		file_v1_hook_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_hook_service_proto_rawDescData)
	})
	return file_v1_hook_service_proto_rawDescData
}

var file_v1_hook_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_hook_service_proto_goTypes = []interface{}{
	(*UpdateDeploymentRequest)(nil),  // 0: hook_service.UpdateDeploymentRequest
	(*UpdateDeploymentResponse)(nil), // 1: hook_service.UpdateDeploymentResponse
}
var file_v1_hook_service_proto_depIdxs = []int32{
	0, // 0: hook_service.HookService.UpdateDeployment:input_type -> hook_service.UpdateDeploymentRequest
	0, // 1: hook_service.HookService.UpdateDevDeployment:input_type -> hook_service.UpdateDeploymentRequest
	1, // 2: hook_service.HookService.UpdateDeployment:output_type -> hook_service.UpdateDeploymentResponse
	1, // 3: hook_service.HookService.UpdateDevDeployment:output_type -> hook_service.UpdateDeploymentResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_hook_service_proto_init() }
func file_v1_hook_service_proto_init() {
	if File_v1_hook_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_hook_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeploymentRequest); i {
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
		file_v1_hook_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeploymentResponse); i {
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
			RawDescriptor: file_v1_hook_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_hook_service_proto_goTypes,
		DependencyIndexes: file_v1_hook_service_proto_depIdxs,
		MessageInfos:      file_v1_hook_service_proto_msgTypes,
	}.Build()
	File_v1_hook_service_proto = out.File
	file_v1_hook_service_proto_rawDesc = nil
	file_v1_hook_service_proto_goTypes = nil
	file_v1_hook_service_proto_depIdxs = nil
}