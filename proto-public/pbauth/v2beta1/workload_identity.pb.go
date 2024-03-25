// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: pbauth/v2beta1/workload_identity.proto

package authv2beta1

import (
	_ "github.com/hashicorp/consul/proto-public/pbresource"
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

type WorkloadIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WorkloadIdentity) Reset() {
	*x = WorkloadIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbauth_v2beta1_workload_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadIdentity) ProtoMessage() {}

func (x *WorkloadIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_pbauth_v2beta1_workload_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadIdentity.ProtoReflect.Descriptor instead.
func (*WorkloadIdentity) Descriptor() ([]byte, []int) {
	return file_pbauth_v2beta1_workload_identity_proto_rawDescGZIP(), []int{0}
}

var File_pbauth_v2beta1_workload_identity_proto protoreflect.FileDescriptor

var file_pbauth_v2beta1_workload_identity_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x62, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x06, 0xa2, 0x93, 0x04, 0x02, 0x08,
	0x03, 0x42, 0x96, 0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x15, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x41, 0xaa, 0x02, 0x1d, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1d, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x41,
	0x75, 0x74, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x29, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x41,
	0x75, 0x74, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x41, 0x75, 0x74,
	0x68, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pbauth_v2beta1_workload_identity_proto_rawDescOnce sync.Once
	file_pbauth_v2beta1_workload_identity_proto_rawDescData = file_pbauth_v2beta1_workload_identity_proto_rawDesc
)

func file_pbauth_v2beta1_workload_identity_proto_rawDescGZIP() []byte {
	file_pbauth_v2beta1_workload_identity_proto_rawDescOnce.Do(func() {
		file_pbauth_v2beta1_workload_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbauth_v2beta1_workload_identity_proto_rawDescData)
	})
	return file_pbauth_v2beta1_workload_identity_proto_rawDescData
}

var file_pbauth_v2beta1_workload_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pbauth_v2beta1_workload_identity_proto_goTypes = []interface{}{
	(*WorkloadIdentity)(nil), // 0: hashicorp.consul.auth.v2beta1.WorkloadIdentity
}
var file_pbauth_v2beta1_workload_identity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbauth_v2beta1_workload_identity_proto_init() }
func file_pbauth_v2beta1_workload_identity_proto_init() {
	if File_pbauth_v2beta1_workload_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbauth_v2beta1_workload_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadIdentity); i {
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
			RawDescriptor: file_pbauth_v2beta1_workload_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbauth_v2beta1_workload_identity_proto_goTypes,
		DependencyIndexes: file_pbauth_v2beta1_workload_identity_proto_depIdxs,
		MessageInfos:      file_pbauth_v2beta1_workload_identity_proto_msgTypes,
	}.Build()
	File_pbauth_v2beta1_workload_identity_proto = out.File
	file_pbauth_v2beta1_workload_identity_proto_rawDesc = nil
	file_pbauth_v2beta1_workload_identity_proto_goTypes = nil
	file_pbauth_v2beta1_workload_identity_proto_depIdxs = nil
}
