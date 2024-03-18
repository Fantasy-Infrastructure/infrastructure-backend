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
// source: pulumi/codegen/loader.proto

package codegen

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

// GetSchemaRequest allows the engine to return a schema for a given package and version.
type GetSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the package name for the schema being requested.
	Package string `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	// the version for the schema being requested, must be a valid semver or empty.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetSchemaRequest) Reset() {
	*x = GetSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_codegen_loader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaRequest) ProtoMessage() {}

func (x *GetSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_codegen_loader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaRequest.ProtoReflect.Descriptor instead.
func (*GetSchemaRequest) Descriptor() ([]byte, []int) {
	return file_pulumi_codegen_loader_proto_rawDescGZIP(), []int{0}
}

func (x *GetSchemaRequest) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *GetSchemaRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// GetSchemaResponse returns the schema data for the requested package.
type GetSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the JSON encoded schema.
	Schema []byte `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *GetSchemaResponse) Reset() {
	*x = GetSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_codegen_loader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaResponse) ProtoMessage() {}

func (x *GetSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_codegen_loader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaResponse.ProtoReflect.Descriptor instead.
func (*GetSchemaResponse) Descriptor() ([]byte, []int) {
	return file_pulumi_codegen_loader_proto_rawDescGZIP(), []int{1}
}

func (x *GetSchemaResponse) GetSchema() []byte {
	if x != nil {
		return x.Schema
	}
	return nil
}

var File_pulumi_codegen_loader_proto protoreflect.FileDescriptor

var file_pulumi_codegen_loader_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e,
	0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2b,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x32, 0x4e, 0x0a, 0x06, 0x4c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pulumi_codegen_loader_proto_rawDescOnce sync.Once
	file_pulumi_codegen_loader_proto_rawDescData = file_pulumi_codegen_loader_proto_rawDesc
)

func file_pulumi_codegen_loader_proto_rawDescGZIP() []byte {
	file_pulumi_codegen_loader_proto_rawDescOnce.Do(func() {
		file_pulumi_codegen_loader_proto_rawDescData = protoimpl.X.CompressGZIP(file_pulumi_codegen_loader_proto_rawDescData)
	})
	return file_pulumi_codegen_loader_proto_rawDescData
}

var file_pulumi_codegen_loader_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pulumi_codegen_loader_proto_goTypes = []interface{}{
	(*GetSchemaRequest)(nil),  // 0: codegen.GetSchemaRequest
	(*GetSchemaResponse)(nil), // 1: codegen.GetSchemaResponse
}
var file_pulumi_codegen_loader_proto_depIdxs = []int32{
	0, // 0: codegen.Loader.GetSchema:input_type -> codegen.GetSchemaRequest
	1, // 1: codegen.Loader.GetSchema:output_type -> codegen.GetSchemaResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pulumi_codegen_loader_proto_init() }
func file_pulumi_codegen_loader_proto_init() {
	if File_pulumi_codegen_loader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pulumi_codegen_loader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchemaRequest); i {
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
		file_pulumi_codegen_loader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchemaResponse); i {
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
			RawDescriptor: file_pulumi_codegen_loader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pulumi_codegen_loader_proto_goTypes,
		DependencyIndexes: file_pulumi_codegen_loader_proto_depIdxs,
		MessageInfos:      file_pulumi_codegen_loader_proto_msgTypes,
	}.Build()
	File_pulumi_codegen_loader_proto = out.File
	file_pulumi_codegen_loader_proto_rawDesc = nil
	file_pulumi_codegen_loader_proto_goTypes = nil
	file_pulumi_codegen_loader_proto_depIdxs = nil
}
