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
// source: proto/BrickBreakerDetailInfo.proto

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

// Obf: KCBKILMACFD
type BrickBreakerDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BKFHCIKHHKN   []*BrickBreakerLevelInfo         `protobuf:"bytes,10,rep,name=BKFHCIKHHKN,proto3" json:"BKFHCIKHHKN,omitempty"`
	StageInfoList []*BrickBreakerActivityStageInfo `protobuf:"bytes,9,rep,name=stage_info_list,json=stageInfoList,proto3" json:"stage_info_list,omitempty"`
	BBMAENNPEOD   []*BrickBreakerLevelInfo         `protobuf:"bytes,2,rep,name=BBMAENNPEOD,proto3" json:"BBMAENNPEOD,omitempty"`
	SkillInfoMap  map[uint32]uint32                `protobuf:"bytes,13,rep,name=skill_info_map,json=skillInfoMap,proto3" json:"skill_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *BrickBreakerDetailInfo) Reset() {
	*x = BrickBreakerDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_BrickBreakerDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrickBreakerDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrickBreakerDetailInfo) ProtoMessage() {}

func (x *BrickBreakerDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_BrickBreakerDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrickBreakerDetailInfo.ProtoReflect.Descriptor instead.
func (*BrickBreakerDetailInfo) Descriptor() ([]byte, []int) {
	return file_proto_BrickBreakerDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *BrickBreakerDetailInfo) GetBKFHCIKHHKN() []*BrickBreakerLevelInfo {
	if x != nil {
		return x.BKFHCIKHHKN
	}
	return nil
}

func (x *BrickBreakerDetailInfo) GetStageInfoList() []*BrickBreakerActivityStageInfo {
	if x != nil {
		return x.StageInfoList
	}
	return nil
}

func (x *BrickBreakerDetailInfo) GetBBMAENNPEOD() []*BrickBreakerLevelInfo {
	if x != nil {
		return x.BBMAENNPEOD
	}
	return nil
}

func (x *BrickBreakerDetailInfo) GetSkillInfoMap() map[uint32]uint32 {
	if x != nil {
		return x.SkillInfoMap
	}
	return nil
}

var File_proto_BrickBreakerDetailInfo_proto protoreflect.FileDescriptor

var file_proto_BrickBreakerDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x72, 0x69, 0x63, 0x6b, 0x42, 0x72, 0x65,
	0x61, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x72, 0x69, 0x63,
	0x6b, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42,
	0x72, 0x69, 0x63, 0x6b, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe6, 0x02, 0x0a, 0x16, 0x42, 0x72, 0x69, 0x63, 0x6b, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a,
	0x0b, 0x42, 0x4b, 0x46, 0x48, 0x43, 0x49, 0x4b, 0x48, 0x48, 0x4b, 0x4e, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x42, 0x72, 0x69, 0x63, 0x6b, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65,
	0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x42, 0x4b, 0x46, 0x48,
	0x43, 0x49, 0x4b, 0x48, 0x48, 0x4b, 0x4e, 0x12, 0x46, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x42, 0x72, 0x69, 0x63, 0x6b, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0d, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x0b, 0x42, 0x42, 0x4d, 0x41, 0x45, 0x4e, 0x4e, 0x50, 0x45, 0x4f, 0x44, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x42, 0x72, 0x69, 0x63, 0x6b, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x42, 0x42,
	0x4d, 0x41, 0x45, 0x4e, 0x4e, 0x50, 0x45, 0x4f, 0x44, 0x12, 0x4f, 0x0a, 0x0e, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x42, 0x72, 0x69, 0x63, 0x6b, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x1a, 0x3f, 0x0a, 0x11, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x47,
	0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_BrickBreakerDetailInfo_proto_rawDescOnce sync.Once
	file_proto_BrickBreakerDetailInfo_proto_rawDescData = file_proto_BrickBreakerDetailInfo_proto_rawDesc
)

func file_proto_BrickBreakerDetailInfo_proto_rawDescGZIP() []byte {
	file_proto_BrickBreakerDetailInfo_proto_rawDescOnce.Do(func() {
		file_proto_BrickBreakerDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_BrickBreakerDetailInfo_proto_rawDescData)
	})
	return file_proto_BrickBreakerDetailInfo_proto_rawDescData
}

var file_proto_BrickBreakerDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_BrickBreakerDetailInfo_proto_goTypes = []interface{}{
	(*BrickBreakerDetailInfo)(nil),        // 0: BrickBreakerDetailInfo
	nil,                                   // 1: BrickBreakerDetailInfo.SkillInfoMapEntry
	(*BrickBreakerLevelInfo)(nil),         // 2: BrickBreakerLevelInfo
	(*BrickBreakerActivityStageInfo)(nil), // 3: BrickBreakerActivityStageInfo
}
var file_proto_BrickBreakerDetailInfo_proto_depIdxs = []int32{
	2, // 0: BrickBreakerDetailInfo.BKFHCIKHHKN:type_name -> BrickBreakerLevelInfo
	3, // 1: BrickBreakerDetailInfo.stage_info_list:type_name -> BrickBreakerActivityStageInfo
	2, // 2: BrickBreakerDetailInfo.BBMAENNPEOD:type_name -> BrickBreakerLevelInfo
	1, // 3: BrickBreakerDetailInfo.skill_info_map:type_name -> BrickBreakerDetailInfo.SkillInfoMapEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_BrickBreakerDetailInfo_proto_init() }
func file_proto_BrickBreakerDetailInfo_proto_init() {
	if File_proto_BrickBreakerDetailInfo_proto != nil {
		return
	}
	file_proto_BrickBreakerLevelInfo_proto_init()
	file_proto_BrickBreakerActivityStageInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_BrickBreakerDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrickBreakerDetailInfo); i {
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
			RawDescriptor: file_proto_BrickBreakerDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_BrickBreakerDetailInfo_proto_goTypes,
		DependencyIndexes: file_proto_BrickBreakerDetailInfo_proto_depIdxs,
		MessageInfos:      file_proto_BrickBreakerDetailInfo_proto_msgTypes,
	}.Build()
	File_proto_BrickBreakerDetailInfo_proto = out.File
	file_proto_BrickBreakerDetailInfo_proto_rawDesc = nil
	file_proto_BrickBreakerDetailInfo_proto_goTypes = nil
	file_proto_BrickBreakerDetailInfo_proto_depIdxs = nil
}