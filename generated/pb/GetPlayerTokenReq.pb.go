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
// source: proto/GetPlayerTokenReq.proto

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

// CmdId: 21228
// Obf: AGJHCDNJDOG
type GetPlayerTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountType   uint32 `protobuf:"varint,11,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	AccountToken  string `protobuf:"bytes,12,opt,name=account_token,json=accountToken,proto3" json:"account_token,omitempty"`
	PsnId         string `protobuf:"bytes,10,opt,name=psn_id,json=psnId,proto3" json:"psn_id,omitempty"`
	OnlineId      string `protobuf:"bytes,1,opt,name=online_id,json=onlineId,proto3" json:"online_id,omitempty"`
	AccountUid    string `protobuf:"bytes,4,opt,name=account_uid,json=accountUid,proto3" json:"account_uid,omitempty"`
	LODOCFKABDG   string `protobuf:"bytes,3,opt,name=LODOCFKABDG,proto3" json:"LODOCFKABDG,omitempty"`
	PlatformType  uint32 `protobuf:"varint,8,opt,name=platform_type,json=platformType,proto3" json:"platform_type,omitempty"`
	ClientIpStr   string `protobuf:"bytes,9,opt,name=client_ip_str,json=clientIpStr,proto3" json:"client_ip_str,omitempty"`
	SubChannelId  uint32 `protobuf:"varint,7,opt,name=sub_channel_id,json=subChannelId,proto3" json:"sub_channel_id,omitempty"`
	CloudClientIp uint32 `protobuf:"varint,177,opt,name=cloudClientIp,proto3" json:"cloudClientIp,omitempty"`
	GIKAFFPKLOE   string `protobuf:"bytes,14,opt,name=GIKAFFPKLOE,proto3" json:"GIKAFFPKLOE,omitempty"`
	Birthday      string `protobuf:"bytes,828,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Uid           uint32 `protobuf:"varint,13,opt,name=uid,proto3" json:"uid,omitempty"`
	OGNFFAEKDBO   uint32 `protobuf:"varint,695,opt,name=OGNFFAEKDBO,proto3" json:"OGNFFAEKDBO,omitempty"`
	KeyId         uint32 `protobuf:"varint,407,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	CountryCode   string `protobuf:"bytes,15,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	ChannelId     uint32 `protobuf:"varint,6,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	FCPDNLHOBNK   uint32 `protobuf:"varint,2,opt,name=FCPDNLHOBNK,proto3" json:"FCPDNLHOBNK,omitempty"`
	IsGuest       bool   `protobuf:"varint,5,opt,name=is_guest,json=isGuest,proto3" json:"is_guest,omitempty"`
	FAOAMMHOOHL   string `protobuf:"bytes,1457,opt,name=FAOAMMHOOHL,proto3" json:"FAOAMMHOOHL,omitempty"`
	ClientRandKey string `protobuf:"bytes,355,opt,name=client_rand_key,json=clientRandKey,proto3" json:"client_rand_key,omitempty"`
}

func (x *GetPlayerTokenReq) Reset() {
	*x = GetPlayerTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_GetPlayerTokenReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerTokenReq) ProtoMessage() {}

func (x *GetPlayerTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_GetPlayerTokenReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerTokenReq.ProtoReflect.Descriptor instead.
func (*GetPlayerTokenReq) Descriptor() ([]byte, []int) {
	return file_proto_GetPlayerTokenReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerTokenReq) GetAccountType() uint32 {
	if x != nil {
		return x.AccountType
	}
	return 0
}

func (x *GetPlayerTokenReq) GetAccountToken() string {
	if x != nil {
		return x.AccountToken
	}
	return ""
}

func (x *GetPlayerTokenReq) GetPsnId() string {
	if x != nil {
		return x.PsnId
	}
	return ""
}

func (x *GetPlayerTokenReq) GetOnlineId() string {
	if x != nil {
		return x.OnlineId
	}
	return ""
}

func (x *GetPlayerTokenReq) GetAccountUid() string {
	if x != nil {
		return x.AccountUid
	}
	return ""
}

func (x *GetPlayerTokenReq) GetLODOCFKABDG() string {
	if x != nil {
		return x.LODOCFKABDG
	}
	return ""
}

func (x *GetPlayerTokenReq) GetPlatformType() uint32 {
	if x != nil {
		return x.PlatformType
	}
	return 0
}

func (x *GetPlayerTokenReq) GetClientIpStr() string {
	if x != nil {
		return x.ClientIpStr
	}
	return ""
}

func (x *GetPlayerTokenReq) GetSubChannelId() uint32 {
	if x != nil {
		return x.SubChannelId
	}
	return 0
}

func (x *GetPlayerTokenReq) GetCloudClientIp() uint32 {
	if x != nil {
		return x.CloudClientIp
	}
	return 0
}

func (x *GetPlayerTokenReq) GetGIKAFFPKLOE() string {
	if x != nil {
		return x.GIKAFFPKLOE
	}
	return ""
}

func (x *GetPlayerTokenReq) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *GetPlayerTokenReq) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetPlayerTokenReq) GetOGNFFAEKDBO() uint32 {
	if x != nil {
		return x.OGNFFAEKDBO
	}
	return 0
}

func (x *GetPlayerTokenReq) GetKeyId() uint32 {
	if x != nil {
		return x.KeyId
	}
	return 0
}

func (x *GetPlayerTokenReq) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *GetPlayerTokenReq) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GetPlayerTokenReq) GetFCPDNLHOBNK() uint32 {
	if x != nil {
		return x.FCPDNLHOBNK
	}
	return 0
}

func (x *GetPlayerTokenReq) GetIsGuest() bool {
	if x != nil {
		return x.IsGuest
	}
	return false
}

func (x *GetPlayerTokenReq) GetFAOAMMHOOHL() string {
	if x != nil {
		return x.FAOAMMHOOHL
	}
	return ""
}

func (x *GetPlayerTokenReq) GetClientRandKey() string {
	if x != nil {
		return x.ClientRandKey
	}
	return ""
}

var File_proto_GetPlayerTokenReq_proto protoreflect.FileDescriptor

var file_proto_GetPlayerTokenReq_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbf, 0x05, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x0a,
	0x06, 0x70, 0x73, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x73, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55,
	0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4f, 0x44, 0x4f, 0x43, 0x46, 0x4b, 0x41, 0x42, 0x44,
	0x47, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x4f, 0x44, 0x4f, 0x43, 0x46, 0x4b,
	0x41, 0x42, 0x44, 0x47, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x53, 0x74, 0x72, 0x12, 0x24, 0x0a,
	0x0e, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x70, 0x18, 0xb1, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x49,
	0x4b, 0x41, 0x46, 0x46, 0x50, 0x4b, 0x4c, 0x4f, 0x45, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x47, 0x49, 0x4b, 0x41, 0x46, 0x46, 0x50, 0x4b, 0x4c, 0x4f, 0x45, 0x12, 0x1b, 0x0a, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0xbc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x4f,
	0x47, 0x4e, 0x46, 0x46, 0x41, 0x45, 0x4b, 0x44, 0x42, 0x4f, 0x18, 0xb7, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x4f, 0x47, 0x4e, 0x46, 0x46, 0x41, 0x45, 0x4b, 0x44, 0x42, 0x4f, 0x12, 0x16,
	0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x97, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x43, 0x50, 0x44,
	0x4e, 0x4c, 0x48, 0x4f, 0x42, 0x4e, 0x4b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46,
	0x43, 0x50, 0x44, 0x4e, 0x4c, 0x48, 0x4f, 0x42, 0x4e, 0x4b, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x46, 0x41, 0x4f, 0x41, 0x4d, 0x4d, 0x48,
	0x4f, 0x4f, 0x48, 0x4c, 0x18, 0xb1, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x41, 0x4f,
	0x41, 0x4d, 0x4d, 0x48, 0x4f, 0x4f, 0x48, 0x4c, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0xe3, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x4b, 0x65,
	0x79, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74,
	0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_GetPlayerTokenReq_proto_rawDescOnce sync.Once
	file_proto_GetPlayerTokenReq_proto_rawDescData = file_proto_GetPlayerTokenReq_proto_rawDesc
)

func file_proto_GetPlayerTokenReq_proto_rawDescGZIP() []byte {
	file_proto_GetPlayerTokenReq_proto_rawDescOnce.Do(func() {
		file_proto_GetPlayerTokenReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_GetPlayerTokenReq_proto_rawDescData)
	})
	return file_proto_GetPlayerTokenReq_proto_rawDescData
}

var file_proto_GetPlayerTokenReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_GetPlayerTokenReq_proto_goTypes = []interface{}{
	(*GetPlayerTokenReq)(nil), // 0: GetPlayerTokenReq
}
var file_proto_GetPlayerTokenReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_GetPlayerTokenReq_proto_init() }
func file_proto_GetPlayerTokenReq_proto_init() {
	if File_proto_GetPlayerTokenReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_GetPlayerTokenReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerTokenReq); i {
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
			RawDescriptor: file_proto_GetPlayerTokenReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_GetPlayerTokenReq_proto_goTypes,
		DependencyIndexes: file_proto_GetPlayerTokenReq_proto_depIdxs,
		MessageInfos:      file_proto_GetPlayerTokenReq_proto_msgTypes,
	}.Build()
	File_proto_GetPlayerTokenReq_proto = out.File
	file_proto_GetPlayerTokenReq_proto_rawDesc = nil
	file_proto_GetPlayerTokenReq_proto_goTypes = nil
	file_proto_GetPlayerTokenReq_proto_depIdxs = nil
}