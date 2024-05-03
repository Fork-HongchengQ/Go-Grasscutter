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
// source: proto/SignInInfo.proto

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

// Obf: CNKBPDDKEFE
type SignInInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigId        uint32        `protobuf:"varint,11,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	RewardDayList   []uint32      `protobuf:"varint,3,rep,packed,name=reward_day_list,json=rewardDayList,proto3" json:"reward_day_list,omitempty"`
	EndTime         uint32        `protobuf:"varint,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	IsCondSatisfied bool          `protobuf:"varint,15,opt,name=is_cond_satisfied,json=isCondSatisfied,proto3" json:"is_cond_satisfied,omitempty"`
	ScheduleId      uint32        `protobuf:"varint,6,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	BeginTime       uint32        `protobuf:"varint,1,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	SigninDataList  []*SignInData `protobuf:"bytes,10,rep,name=signin_data_list,json=signinDataList,proto3" json:"signin_data_list,omitempty"`
	LMOEFBHICGL     uint32        `protobuf:"varint,7,opt,name=LMOEFBHICGL,proto3" json:"LMOEFBHICGL,omitempty"`
	IOGOBJPGDPJ     uint32        `protobuf:"varint,9,opt,name=IOGOBJPGDPJ,proto3" json:"IOGOBJPGDPJ,omitempty"`
}

func (x *SignInInfo) Reset() {
	*x = SignInInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_SignInInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInInfo) ProtoMessage() {}

func (x *SignInInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_SignInInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInInfo.ProtoReflect.Descriptor instead.
func (*SignInInfo) Descriptor() ([]byte, []int) {
	return file_proto_SignInInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SignInInfo) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *SignInInfo) GetRewardDayList() []uint32 {
	if x != nil {
		return x.RewardDayList
	}
	return nil
}

func (x *SignInInfo) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *SignInInfo) GetIsCondSatisfied() bool {
	if x != nil {
		return x.IsCondSatisfied
	}
	return false
}

func (x *SignInInfo) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (x *SignInInfo) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *SignInInfo) GetSigninDataList() []*SignInData {
	if x != nil {
		return x.SigninDataList
	}
	return nil
}

func (x *SignInInfo) GetLMOEFBHICGL() uint32 {
	if x != nil {
		return x.LMOEFBHICGL
	}
	return 0
}

func (x *SignInInfo) GetIOGOBJPGDPJ() uint32 {
	if x != nil {
		return x.IOGOBJPGDPJ
	}
	return 0
}

var File_proto_SignInInfo_proto protoreflect.FileDescriptor

var file_proto_SignInInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd3, 0x02, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x73, 0x61, 0x74, 0x69, 0x73,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x6f,
	0x6e, 0x64, 0x53, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x10, 0x73,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4d, 0x4f, 0x45, 0x46, 0x42, 0x48, 0x49, 0x43, 0x47,
	0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x4d, 0x4f, 0x45, 0x46, 0x42, 0x48,
	0x49, 0x43, 0x47, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4f, 0x47, 0x4f, 0x42, 0x4a, 0x50, 0x47,
	0x44, 0x50, 0x4a, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4f, 0x47, 0x4f, 0x42,
	0x4a, 0x50, 0x47, 0x44, 0x50, 0x4a, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61,
	0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_SignInInfo_proto_rawDescOnce sync.Once
	file_proto_SignInInfo_proto_rawDescData = file_proto_SignInInfo_proto_rawDesc
)

func file_proto_SignInInfo_proto_rawDescGZIP() []byte {
	file_proto_SignInInfo_proto_rawDescOnce.Do(func() {
		file_proto_SignInInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_SignInInfo_proto_rawDescData)
	})
	return file_proto_SignInInfo_proto_rawDescData
}

var file_proto_SignInInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_SignInInfo_proto_goTypes = []interface{}{
	(*SignInInfo)(nil), // 0: SignInInfo
	(*SignInData)(nil), // 1: SignInData
}
var file_proto_SignInInfo_proto_depIdxs = []int32{
	1, // 0: SignInInfo.signin_data_list:type_name -> SignInData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_SignInInfo_proto_init() }
func file_proto_SignInInfo_proto_init() {
	if File_proto_SignInInfo_proto != nil {
		return
	}
	file_proto_SignInData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_SignInInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInInfo); i {
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
			RawDescriptor: file_proto_SignInInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_SignInInfo_proto_goTypes,
		DependencyIndexes: file_proto_SignInInfo_proto_depIdxs,
		MessageInfos:      file_proto_SignInInfo_proto_msgTypes,
	}.Build()
	File_proto_SignInInfo_proto = out.File
	file_proto_SignInInfo_proto_rawDesc = nil
	file_proto_SignInInfo_proto_goTypes = nil
	file_proto_SignInInfo_proto_depIdxs = nil
}