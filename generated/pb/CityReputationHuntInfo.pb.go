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
// source: proto/CityReputationHuntInfo.proto

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

// Obf: JGMFFPHIBEL
type CityReputationHuntInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOpen           bool   `protobuf:"varint,10,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	HasReward        bool   `protobuf:"varint,4,opt,name=has_reward,json=hasReward,proto3" json:"has_reward,omitempty"`
	CurWeekFinishNum uint32 `protobuf:"varint,3,opt,name=cur_week_finish_num,json=curWeekFinishNum,proto3" json:"cur_week_finish_num,omitempty"`
}

func (x *CityReputationHuntInfo) Reset() {
	*x = CityReputationHuntInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CityReputationHuntInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityReputationHuntInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityReputationHuntInfo) ProtoMessage() {}

func (x *CityReputationHuntInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CityReputationHuntInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityReputationHuntInfo.ProtoReflect.Descriptor instead.
func (*CityReputationHuntInfo) Descriptor() ([]byte, []int) {
	return file_proto_CityReputationHuntInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CityReputationHuntInfo) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *CityReputationHuntInfo) GetHasReward() bool {
	if x != nil {
		return x.HasReward
	}
	return false
}

func (x *CityReputationHuntInfo) GetCurWeekFinishNum() uint32 {
	if x != nil {
		return x.CurWeekFinishNum
	}
	return 0
}

var File_proto_CityReputationHuntInfo_proto protoreflect.FileDescriptor

var file_proto_CityReputationHuntInfo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x16, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x5f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x61, 0x73,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x13, 0x63, 0x75, 0x72, 0x5f, 0x77, 0x65,
	0x65, 0x6b, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x10, 0x63, 0x75, 0x72, 0x57, 0x65, 0x65, 0x6b, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x4e, 0x75, 0x6d, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73,
	0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_CityReputationHuntInfo_proto_rawDescOnce sync.Once
	file_proto_CityReputationHuntInfo_proto_rawDescData = file_proto_CityReputationHuntInfo_proto_rawDesc
)

func file_proto_CityReputationHuntInfo_proto_rawDescGZIP() []byte {
	file_proto_CityReputationHuntInfo_proto_rawDescOnce.Do(func() {
		file_proto_CityReputationHuntInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_CityReputationHuntInfo_proto_rawDescData)
	})
	return file_proto_CityReputationHuntInfo_proto_rawDescData
}

var file_proto_CityReputationHuntInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_CityReputationHuntInfo_proto_goTypes = []interface{}{
	(*CityReputationHuntInfo)(nil), // 0: CityReputationHuntInfo
}
var file_proto_CityReputationHuntInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_CityReputationHuntInfo_proto_init() }
func file_proto_CityReputationHuntInfo_proto_init() {
	if File_proto_CityReputationHuntInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_CityReputationHuntInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityReputationHuntInfo); i {
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
			RawDescriptor: file_proto_CityReputationHuntInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_CityReputationHuntInfo_proto_goTypes,
		DependencyIndexes: file_proto_CityReputationHuntInfo_proto_depIdxs,
		MessageInfos:      file_proto_CityReputationHuntInfo_proto_msgTypes,
	}.Build()
	File_proto_CityReputationHuntInfo_proto = out.File
	file_proto_CityReputationHuntInfo_proto_rawDesc = nil
	file_proto_CityReputationHuntInfo_proto_goTypes = nil
	file_proto_CityReputationHuntInfo_proto_depIdxs = nil
}