// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: pulumi/converter.proto

package pulumirpc

import (
	codegen "github.com/pulumi/pulumi/sdk/v3/proto/go/codegen"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConvertStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the gRPC target of the mapper service.
	MapperTarget string `protobuf:"bytes,1,opt,name=mapper_target,json=mapperTarget,proto3" json:"mapper_target,omitempty"`
	// the args passed to `pulumi import` for this conversion. Normally used to specifiy a state file to
	// import from.
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *ConvertStateRequest) Reset() {
	*x = ConvertStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_converter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertStateRequest) ProtoMessage() {}

func (x *ConvertStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_converter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertStateRequest.ProtoReflect.Descriptor instead.
func (*ConvertStateRequest) Descriptor() ([]byte, []int) {
	return file_pulumi_converter_proto_rawDescGZIP(), []int{0}
}

func (x *ConvertStateRequest) GetMapperTarget() string {
	if x != nil {
		return x.MapperTarget
	}
	return ""
}

func (x *ConvertStateRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

// A ResourceImport specifies a resource to import.
type ResourceImport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the type token for the resource.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// the name of the resource.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// the ID of the resource.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// the provider version to use for the resource, if any.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// the provider PluginDownloadURL to use for the resource, if any.
	PluginDownloadURL string `protobuf:"bytes,5,opt,name=pluginDownloadURL,proto3" json:"pluginDownloadURL,omitempty"`
}

func (x *ResourceImport) Reset() {
	*x = ResourceImport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_converter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceImport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceImport) ProtoMessage() {}

func (x *ResourceImport) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_converter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceImport.ProtoReflect.Descriptor instead.
func (*ResourceImport) Descriptor() ([]byte, []int) {
	return file_pulumi_converter_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceImport) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ResourceImport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceImport) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceImport) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ResourceImport) GetPluginDownloadURL() string {
	if x != nil {
		return x.PluginDownloadURL
	}
	return ""
}

type ConvertStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a list of resources to import.
	Resources []*ResourceImport `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	// any diagnostics from state conversion.
	Diagnostics []*codegen.Diagnostic `protobuf:"bytes,2,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *ConvertStateResponse) Reset() {
	*x = ConvertStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_converter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertStateResponse) ProtoMessage() {}

func (x *ConvertStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_converter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertStateResponse.ProtoReflect.Descriptor instead.
func (*ConvertStateResponse) Descriptor() ([]byte, []int) {
	return file_pulumi_converter_proto_rawDescGZIP(), []int{2}
}

func (x *ConvertStateResponse) GetResources() []*ResourceImport {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *ConvertStateResponse) GetDiagnostics() []*codegen.Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

type ConvertProgramRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the source directory containing the program to convert from.
	SourceDirectory string `protobuf:"bytes,1,opt,name=source_directory,json=sourceDirectory,proto3" json:"source_directory,omitempty"`
	// a target directory to write the resulting PCL code and project file to.
	TargetDirectory string `protobuf:"bytes,2,opt,name=target_directory,json=targetDirectory,proto3" json:"target_directory,omitempty"`
	// the gRPC target of the mapper service.
	MapperTarget string `protobuf:"bytes,3,opt,name=mapper_target,json=mapperTarget,proto3" json:"mapper_target,omitempty"`
	// The target of a codegen.LoaderServer to use for loading schemas.
	LoaderTarget string `protobuf:"bytes,4,opt,name=loader_target,json=loaderTarget,proto3" json:"loader_target,omitempty"`
	// the args passed to `pulumi convert` for this conversion. Normally used to specifiy a root file, or conversion options.
	Args []string `protobuf:"bytes,5,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *ConvertProgramRequest) Reset() {
	*x = ConvertProgramRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_converter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertProgramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertProgramRequest) ProtoMessage() {}

func (x *ConvertProgramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_converter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertProgramRequest.ProtoReflect.Descriptor instead.
func (*ConvertProgramRequest) Descriptor() ([]byte, []int) {
	return file_pulumi_converter_proto_rawDescGZIP(), []int{3}
}

func (x *ConvertProgramRequest) GetSourceDirectory() string {
	if x != nil {
		return x.SourceDirectory
	}
	return ""
}

func (x *ConvertProgramRequest) GetTargetDirectory() string {
	if x != nil {
		return x.TargetDirectory
	}
	return ""
}

func (x *ConvertProgramRequest) GetMapperTarget() string {
	if x != nil {
		return x.MapperTarget
	}
	return ""
}

func (x *ConvertProgramRequest) GetLoaderTarget() string {
	if x != nil {
		return x.LoaderTarget
	}
	return ""
}

func (x *ConvertProgramRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type ConvertProgramResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// any diagnostics from code generation.
	Diagnostics []*codegen.Diagnostic `protobuf:"bytes,1,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *ConvertProgramResponse) Reset() {
	*x = ConvertProgramResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_converter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertProgramResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertProgramResponse) ProtoMessage() {}

func (x *ConvertProgramResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_converter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertProgramResponse.ProtoReflect.Descriptor instead.
func (*ConvertProgramResponse) Descriptor() ([]byte, []int) {
	return file_pulumi_converter_proto_rawDescGZIP(), []int{4}
}

func (x *ConvertProgramResponse) GetDiagnostics() []*codegen.Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

var File_pulumi_converter_proto protoreflect.FileDescriptor

var file_pulumi_converter_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x72, 0x70, 0x63, 0x1a, 0x18, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x67, 0x65, 0x6e, 0x2f, 0x68, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70,
	0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a,
	0x13, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x90, 0x01,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c,
	0x22, 0x90, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x44, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x22, 0xcb, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x22, 0x59, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x64,
	0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x52,
	0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x32, 0xb7, 0x01, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x0c, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x75, 0x6c,
	0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x75, 0x6c,
	0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12,
	0x20, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x70, 0x75, 0x6c, 0x75,
	0x6d, 0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x3b, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pulumi_converter_proto_rawDescOnce sync.Once
	file_pulumi_converter_proto_rawDescData = file_pulumi_converter_proto_rawDesc
)

func file_pulumi_converter_proto_rawDescGZIP() []byte {
	file_pulumi_converter_proto_rawDescOnce.Do(func() {
		file_pulumi_converter_proto_rawDescData = protoimpl.X.CompressGZIP(file_pulumi_converter_proto_rawDescData)
	})
	return file_pulumi_converter_proto_rawDescData
}

var file_pulumi_converter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pulumi_converter_proto_goTypes = []interface{}{
	(*ConvertStateRequest)(nil),    // 0: pulumirpc.ConvertStateRequest
	(*ResourceImport)(nil),         // 1: pulumirpc.ResourceImport
	(*ConvertStateResponse)(nil),   // 2: pulumirpc.ConvertStateResponse
	(*ConvertProgramRequest)(nil),  // 3: pulumirpc.ConvertProgramRequest
	(*ConvertProgramResponse)(nil), // 4: pulumirpc.ConvertProgramResponse
	(*codegen.Diagnostic)(nil),     // 5: pulumirpc.codegen.Diagnostic
}
var file_pulumi_converter_proto_depIdxs = []int32{
	1, // 0: pulumirpc.ConvertStateResponse.resources:type_name -> pulumirpc.ResourceImport
	5, // 1: pulumirpc.ConvertStateResponse.diagnostics:type_name -> pulumirpc.codegen.Diagnostic
	5, // 2: pulumirpc.ConvertProgramResponse.diagnostics:type_name -> pulumirpc.codegen.Diagnostic
	0, // 3: pulumirpc.Converter.ConvertState:input_type -> pulumirpc.ConvertStateRequest
	3, // 4: pulumirpc.Converter.ConvertProgram:input_type -> pulumirpc.ConvertProgramRequest
	2, // 5: pulumirpc.Converter.ConvertState:output_type -> pulumirpc.ConvertStateResponse
	4, // 6: pulumirpc.Converter.ConvertProgram:output_type -> pulumirpc.ConvertProgramResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pulumi_converter_proto_init() }
func file_pulumi_converter_proto_init() {
	if File_pulumi_converter_proto != nil {
		return
	}
	file_pulumi_plugin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pulumi_converter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertStateRequest); i {
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
		file_pulumi_converter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceImport); i {
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
		file_pulumi_converter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertStateResponse); i {
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
		file_pulumi_converter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertProgramRequest); i {
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
		file_pulumi_converter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertProgramResponse); i {
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
			RawDescriptor: file_pulumi_converter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pulumi_converter_proto_goTypes,
		DependencyIndexes: file_pulumi_converter_proto_depIdxs,
		MessageInfos:      file_pulumi_converter_proto_msgTypes,
	}.Build()
	File_pulumi_converter_proto = out.File
	file_pulumi_converter_proto_rawDesc = nil
	file_pulumi_converter_proto_goTypes = nil
	file_pulumi_converter_proto_depIdxs = nil
}
