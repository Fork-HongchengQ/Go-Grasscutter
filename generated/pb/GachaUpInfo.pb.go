// https://github.com/SlushinPS/beach-simulator
// Copyright (C) 2023 Slushy Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.1
// source: proto/GachaUpInfo.proto

package pb

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

// Obf: GLKEEDIHOFH
type GachaUpInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemIdList     []uint32 `protobuf:"varint,4,rep,packed,name=item_id_list,json=itemIdList,proto3" json:"item_id_list,omitempty"`
	ItemParentType uint32   `protobuf:"varint,6,opt,name=item_parent_type,json=itemParentType,proto3" json:"item_parent_type,omitempty"`
}

func (x *GachaUpInfo) Reset() {
	*x = GachaUpInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_GachaUpInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GachaUpInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GachaUpInfo) ProtoMessage() {}

func (x *GachaUpInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_GachaUpInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GachaUpInfo.ProtoReflect.Descriptor instead.
func (*GachaUpInfo) Descriptor() ([]byte, []int) {
	return file_proto_GachaUpInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GachaUpInfo) GetItemIdList() []uint32 {
	if x != nil {
		return x.ItemIdList
	}
	return nil
}

func (x *GachaUpInfo) GetItemParentType() uint32 {
	if x != nil {
		return x.ItemParentType
	}
	return 0
}

var File_proto_GachaUpInfo_proto protoreflect.FileDescriptor

var file_proto_GachaUpInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x55, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0b, 0x47, 0x61, 0x63,
	0x68, 0x61, 0x55, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a,
	0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73,
	0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_GachaUpInfo_proto_rawDescOnce sync.Once
	file_proto_GachaUpInfo_proto_rawDescData = file_proto_GachaUpInfo_proto_rawDesc
)

func file_proto_GachaUpInfo_proto_rawDescGZIP() []byte {
	file_proto_GachaUpInfo_proto_rawDescOnce.Do(func() {
		file_proto_GachaUpInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_GachaUpInfo_proto_rawDescData)
	})
	return file_proto_GachaUpInfo_proto_rawDescData
}

var file_proto_GachaUpInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_GachaUpInfo_proto_goTypes = []interface{}{
	(*GachaUpInfo)(nil), // 0: GachaUpInfo
}
var file_proto_GachaUpInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_GachaUpInfo_proto_init() }
func file_proto_GachaUpInfo_proto_init() {
	if File_proto_GachaUpInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_GachaUpInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GachaUpInfo); i {
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
			RawDescriptor: file_proto_GachaUpInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_GachaUpInfo_proto_goTypes,
		DependencyIndexes: file_proto_GachaUpInfo_proto_depIdxs,
		MessageInfos:      file_proto_GachaUpInfo_proto_msgTypes,
	}.Build()
	File_proto_GachaUpInfo_proto = out.File
	file_proto_GachaUpInfo_proto_rawDesc = nil
	file_proto_GachaUpInfo_proto_goTypes = nil
	file_proto_GachaUpInfo_proto_depIdxs = nil
}