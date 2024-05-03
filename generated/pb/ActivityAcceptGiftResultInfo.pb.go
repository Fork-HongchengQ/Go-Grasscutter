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
// source: proto/ActivityAcceptGiftResultInfo.proto

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

// Obf: MGAPJIJBGGJ
type ActivityAcceptGiftResultInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         uint32            `protobuf:"varint,11,opt,name=uid,proto3" json:"uid,omitempty"`
	IJCFDJFGMPP map[uint32]uint32 `protobuf:"bytes,12,rep,name=IJCFDJFGMPP,proto3" json:"IJCFDJFGMPP,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	PKBLFFNAILJ map[uint32]uint32 `protobuf:"bytes,4,rep,name=PKBLFFNAILJ,proto3" json:"PKBLFFNAILJ,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ActivityAcceptGiftResultInfo) Reset() {
	*x = ActivityAcceptGiftResultInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ActivityAcceptGiftResultInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityAcceptGiftResultInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityAcceptGiftResultInfo) ProtoMessage() {}

func (x *ActivityAcceptGiftResultInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ActivityAcceptGiftResultInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityAcceptGiftResultInfo.ProtoReflect.Descriptor instead.
func (*ActivityAcceptGiftResultInfo) Descriptor() ([]byte, []int) {
	return file_proto_ActivityAcceptGiftResultInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityAcceptGiftResultInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ActivityAcceptGiftResultInfo) GetIJCFDJFGMPP() map[uint32]uint32 {
	if x != nil {
		return x.IJCFDJFGMPP
	}
	return nil
}

func (x *ActivityAcceptGiftResultInfo) GetPKBLFFNAILJ() map[uint32]uint32 {
	if x != nil {
		return x.PKBLFFNAILJ
	}
	return nil
}

var File_proto_ActivityAcceptGiftResultInfo_proto protoreflect.FileDescriptor

var file_proto_ActivityAcceptGiftResultInfo_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x02, 0x0a, 0x1c, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x47, 0x69, 0x66,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x50, 0x0a,
	0x0b, 0x49, 0x4a, 0x43, 0x46, 0x44, 0x4a, 0x46, 0x47, 0x4d, 0x50, 0x50, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x49, 0x4a, 0x43, 0x46, 0x44, 0x4a, 0x46, 0x47, 0x4d, 0x50, 0x50, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x49, 0x4a, 0x43, 0x46, 0x44, 0x4a, 0x46, 0x47, 0x4d, 0x50, 0x50, 0x12,
	0x50, 0x0a, 0x0b, 0x50, 0x4b, 0x42, 0x4c, 0x46, 0x46, 0x4e, 0x41, 0x49, 0x4c, 0x4a, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x4b, 0x42, 0x4c, 0x46, 0x46, 0x4e, 0x41, 0x49, 0x4c, 0x4a, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x50, 0x4b, 0x42, 0x4c, 0x46, 0x46, 0x4e, 0x41, 0x49, 0x4c,
	0x4a, 0x1a, 0x3e, 0x0a, 0x10, 0x49, 0x4a, 0x43, 0x46, 0x44, 0x4a, 0x46, 0x47, 0x4d, 0x50, 0x50,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x50, 0x4b, 0x42, 0x4c, 0x46, 0x46, 0x4e, 0x41, 0x49, 0x4c, 0x4a,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74,
	0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ActivityAcceptGiftResultInfo_proto_rawDescOnce sync.Once
	file_proto_ActivityAcceptGiftResultInfo_proto_rawDescData = file_proto_ActivityAcceptGiftResultInfo_proto_rawDesc
)

func file_proto_ActivityAcceptGiftResultInfo_proto_rawDescGZIP() []byte {
	file_proto_ActivityAcceptGiftResultInfo_proto_rawDescOnce.Do(func() {
		file_proto_ActivityAcceptGiftResultInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ActivityAcceptGiftResultInfo_proto_rawDescData)
	})
	return file_proto_ActivityAcceptGiftResultInfo_proto_rawDescData
}

var file_proto_ActivityAcceptGiftResultInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_ActivityAcceptGiftResultInfo_proto_goTypes = []interface{}{
	(*ActivityAcceptGiftResultInfo)(nil), // 0: ActivityAcceptGiftResultInfo
	nil,                                  // 1: ActivityAcceptGiftResultInfo.IJCFDJFGMPPEntry
	nil,                                  // 2: ActivityAcceptGiftResultInfo.PKBLFFNAILJEntry
}
var file_proto_ActivityAcceptGiftResultInfo_proto_depIdxs = []int32{
	1, // 0: ActivityAcceptGiftResultInfo.IJCFDJFGMPP:type_name -> ActivityAcceptGiftResultInfo.IJCFDJFGMPPEntry
	2, // 1: ActivityAcceptGiftResultInfo.PKBLFFNAILJ:type_name -> ActivityAcceptGiftResultInfo.PKBLFFNAILJEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_ActivityAcceptGiftResultInfo_proto_init() }
func file_proto_ActivityAcceptGiftResultInfo_proto_init() {
	if File_proto_ActivityAcceptGiftResultInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ActivityAcceptGiftResultInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityAcceptGiftResultInfo); i {
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
			RawDescriptor: file_proto_ActivityAcceptGiftResultInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ActivityAcceptGiftResultInfo_proto_goTypes,
		DependencyIndexes: file_proto_ActivityAcceptGiftResultInfo_proto_depIdxs,
		MessageInfos:      file_proto_ActivityAcceptGiftResultInfo_proto_msgTypes,
	}.Build()
	File_proto_ActivityAcceptGiftResultInfo_proto = out.File
	file_proto_ActivityAcceptGiftResultInfo_proto_rawDesc = nil
	file_proto_ActivityAcceptGiftResultInfo_proto_goTypes = nil
	file_proto_ActivityAcceptGiftResultInfo_proto_depIdxs = nil
}