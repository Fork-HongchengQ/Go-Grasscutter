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
// source: proto/MusicGameSettleReq.proto

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

// CmdId: 24309
// Obf: IOLEMBLHICB
type MusicGameSettleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrectHit   uint32   `protobuf:"varint,7,opt,name=correct_hit,json=correctHit,proto3" json:"correct_hit,omitempty"`
	OEAHADEGEOA  []uint32 `protobuf:"varint,327,rep,packed,name=OEAHADEGEOA,proto3" json:"OEAHADEGEOA,omitempty"`
	UgcGuid      uint64   `protobuf:"varint,1,opt,name=ugc_guid,json=ugcGuid,proto3" json:"ugc_guid,omitempty"`
	BPNLLFDJJOL  bool     `protobuf:"varint,1200,opt,name=BPNLLFDJJOL,proto3" json:"BPNLLFDJJOL,omitempty"`
	Score        uint32   `protobuf:"varint,8,opt,name=score,proto3" json:"score,omitempty"`
	MECALGKAKJK  uint32   `protobuf:"varint,14,opt,name=MECALGKAKJK,proto3" json:"MECALGKAKJK,omitempty"`
	MaxCombo     uint32   `protobuf:"varint,2,opt,name=max_combo,json=maxCombo,proto3" json:"max_combo,omitempty"`
	KPPICEDHMPN  []uint32 `protobuf:"varint,5,rep,packed,name=KPPICEDHMPN,proto3" json:"KPPICEDHMPN,omitempty"`
	Speed        float32  `protobuf:"fixed32,1628,opt,name=speed,proto3" json:"speed,omitempty"`
	KDAOEDCLEFG  uint32   `protobuf:"varint,1749,opt,name=KDAOEDCLEFG,proto3" json:"KDAOEDCLEFG,omitempty"`
	IsSaveScore  bool     `protobuf:"varint,6,opt,name=is_save_score,json=isSaveScore,proto3" json:"is_save_score,omitempty"`
	MAMHOPGFOKD  bool     `protobuf:"varint,1198,opt,name=MAMHOPGFOKD,proto3" json:"MAMHOPGFOKD,omitempty"`
	NMPPJPOJFDC  uint32   `protobuf:"varint,4,opt,name=NMPPJPOJFDC,proto3" json:"NMPPJPOJFDC,omitempty"`
	MusicBasicId uint32   `protobuf:"varint,11,opt,name=music_basic_id,json=musicBasicId,proto3" json:"music_basic_id,omitempty"`
	Combo        uint32   `protobuf:"varint,9,opt,name=combo,proto3" json:"combo,omitempty"`
	NGALDEAEBHG  uint32   `protobuf:"varint,15,opt,name=NGALDEAEBHG,proto3" json:"NGALDEAEBHG,omitempty"`
	FCFNKIDLDHJ  uint32   `protobuf:"varint,938,opt,name=FCFNKIDLDHJ,proto3" json:"FCFNKIDLDHJ,omitempty"`
	GDOMKIHOKCC  uint32   `protobuf:"varint,747,opt,name=GDOMKIHOKCC,proto3" json:"GDOMKIHOKCC,omitempty"`
}

func (x *MusicGameSettleReq) Reset() {
	*x = MusicGameSettleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_MusicGameSettleReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MusicGameSettleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MusicGameSettleReq) ProtoMessage() {}

func (x *MusicGameSettleReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_MusicGameSettleReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MusicGameSettleReq.ProtoReflect.Descriptor instead.
func (*MusicGameSettleReq) Descriptor() ([]byte, []int) {
	return file_proto_MusicGameSettleReq_proto_rawDescGZIP(), []int{0}
}

func (x *MusicGameSettleReq) GetCorrectHit() uint32 {
	if x != nil {
		return x.CorrectHit
	}
	return 0
}

func (x *MusicGameSettleReq) GetOEAHADEGEOA() []uint32 {
	if x != nil {
		return x.OEAHADEGEOA
	}
	return nil
}

func (x *MusicGameSettleReq) GetUgcGuid() uint64 {
	if x != nil {
		return x.UgcGuid
	}
	return 0
}

func (x *MusicGameSettleReq) GetBPNLLFDJJOL() bool {
	if x != nil {
		return x.BPNLLFDJJOL
	}
	return false
}

func (x *MusicGameSettleReq) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *MusicGameSettleReq) GetMECALGKAKJK() uint32 {
	if x != nil {
		return x.MECALGKAKJK
	}
	return 0
}

func (x *MusicGameSettleReq) GetMaxCombo() uint32 {
	if x != nil {
		return x.MaxCombo
	}
	return 0
}

func (x *MusicGameSettleReq) GetKPPICEDHMPN() []uint32 {
	if x != nil {
		return x.KPPICEDHMPN
	}
	return nil
}

func (x *MusicGameSettleReq) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *MusicGameSettleReq) GetKDAOEDCLEFG() uint32 {
	if x != nil {
		return x.KDAOEDCLEFG
	}
	return 0
}

func (x *MusicGameSettleReq) GetIsSaveScore() bool {
	if x != nil {
		return x.IsSaveScore
	}
	return false
}

func (x *MusicGameSettleReq) GetMAMHOPGFOKD() bool {
	if x != nil {
		return x.MAMHOPGFOKD
	}
	return false
}

func (x *MusicGameSettleReq) GetNMPPJPOJFDC() uint32 {
	if x != nil {
		return x.NMPPJPOJFDC
	}
	return 0
}

func (x *MusicGameSettleReq) GetMusicBasicId() uint32 {
	if x != nil {
		return x.MusicBasicId
	}
	return 0
}

func (x *MusicGameSettleReq) GetCombo() uint32 {
	if x != nil {
		return x.Combo
	}
	return 0
}

func (x *MusicGameSettleReq) GetNGALDEAEBHG() uint32 {
	if x != nil {
		return x.NGALDEAEBHG
	}
	return 0
}

func (x *MusicGameSettleReq) GetFCFNKIDLDHJ() uint32 {
	if x != nil {
		return x.FCFNKIDLDHJ
	}
	return 0
}

func (x *MusicGameSettleReq) GetGDOMKIHOKCC() uint32 {
	if x != nil {
		return x.GDOMKIHOKCC
	}
	return 0
}

var File_proto_MusicGameSettleReq_proto protoreflect.FileDescriptor

var file_proto_MusicGameSettleReq_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd4, 0x04, 0x0a, 0x12, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x5f, 0x68, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x48, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x4f, 0x45, 0x41, 0x48,
	0x41, 0x44, 0x45, 0x47, 0x45, 0x4f, 0x41, 0x18, 0xc7, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b,
	0x4f, 0x45, 0x41, 0x48, 0x41, 0x44, 0x45, 0x47, 0x45, 0x4f, 0x41, 0x12, 0x19, 0x0a, 0x08, 0x75,
	0x67, 0x63, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x75,
	0x67, 0x63, 0x47, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x42, 0x50, 0x4e, 0x4c, 0x4c, 0x46,
	0x44, 0x4a, 0x4a, 0x4f, 0x4c, 0x18, 0xb0, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x50,
	0x4e, 0x4c, 0x4c, 0x46, 0x44, 0x4a, 0x4a, 0x4f, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x4d, 0x45, 0x43, 0x41, 0x4c, 0x47, 0x4b, 0x41, 0x4b, 0x4a, 0x4b, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x45, 0x43, 0x41, 0x4c, 0x47, 0x4b, 0x41, 0x4b, 0x4a,
	0x4b, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6d, 0x62, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6d, 0x62, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x4b, 0x50, 0x50, 0x49, 0x43, 0x45, 0x44, 0x48, 0x4d, 0x50, 0x4e, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x50, 0x50, 0x49, 0x43, 0x45, 0x44, 0x48, 0x4d, 0x50, 0x4e,
	0x12, 0x15, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0xdc, 0x0c, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x4b, 0x44, 0x41, 0x4f, 0x45,
	0x44, 0x43, 0x4c, 0x45, 0x46, 0x47, 0x18, 0xd5, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b,
	0x44, 0x41, 0x4f, 0x45, 0x44, 0x43, 0x4c, 0x45, 0x46, 0x47, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73,
	0x5f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x53, 0x61, 0x76, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x21,
	0x0a, 0x0b, 0x4d, 0x41, 0x4d, 0x48, 0x4f, 0x50, 0x47, 0x46, 0x4f, 0x4b, 0x44, 0x18, 0xae, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4d, 0x41, 0x4d, 0x48, 0x4f, 0x50, 0x47, 0x46, 0x4f, 0x4b,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4d, 0x50, 0x50, 0x4a, 0x50, 0x4f, 0x4a, 0x46, 0x44, 0x43,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x4d, 0x50, 0x50, 0x4a, 0x50, 0x4f, 0x4a,
	0x46, 0x44, 0x43, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x75, 0x73,
	0x69, 0x63, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6d,
	0x62, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x6d, 0x62, 0x6f, 0x12,
	0x20, 0x0a, 0x0b, 0x4e, 0x47, 0x41, 0x4c, 0x44, 0x45, 0x41, 0x45, 0x42, 0x48, 0x47, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x47, 0x41, 0x4c, 0x44, 0x45, 0x41, 0x45, 0x42, 0x48,
	0x47, 0x12, 0x21, 0x0a, 0x0b, 0x46, 0x43, 0x46, 0x4e, 0x4b, 0x49, 0x44, 0x4c, 0x44, 0x48, 0x4a,
	0x18, 0xaa, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x43, 0x46, 0x4e, 0x4b, 0x49, 0x44,
	0x4c, 0x44, 0x48, 0x4a, 0x12, 0x21, 0x0a, 0x0b, 0x47, 0x44, 0x4f, 0x4d, 0x4b, 0x49, 0x48, 0x4f,
	0x4b, 0x43, 0x43, 0x18, 0xeb, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x44, 0x4f, 0x4d,
	0x4b, 0x49, 0x48, 0x4f, 0x4b, 0x43, 0x43, 0x42, 0x1d, 0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72,
	0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_MusicGameSettleReq_proto_rawDescOnce sync.Once
	file_proto_MusicGameSettleReq_proto_rawDescData = file_proto_MusicGameSettleReq_proto_rawDesc
)

func file_proto_MusicGameSettleReq_proto_rawDescGZIP() []byte {
	file_proto_MusicGameSettleReq_proto_rawDescOnce.Do(func() {
		file_proto_MusicGameSettleReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_MusicGameSettleReq_proto_rawDescData)
	})
	return file_proto_MusicGameSettleReq_proto_rawDescData
}

var file_proto_MusicGameSettleReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_MusicGameSettleReq_proto_goTypes = []interface{}{
	(*MusicGameSettleReq)(nil), // 0: MusicGameSettleReq
}
var file_proto_MusicGameSettleReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_MusicGameSettleReq_proto_init() }
func file_proto_MusicGameSettleReq_proto_init() {
	if File_proto_MusicGameSettleReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_MusicGameSettleReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MusicGameSettleReq); i {
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
			RawDescriptor: file_proto_MusicGameSettleReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_MusicGameSettleReq_proto_goTypes,
		DependencyIndexes: file_proto_MusicGameSettleReq_proto_depIdxs,
		MessageInfos:      file_proto_MusicGameSettleReq_proto_msgTypes,
	}.Build()
	File_proto_MusicGameSettleReq_proto = out.File
	file_proto_MusicGameSettleReq_proto_rawDesc = nil
	file_proto_MusicGameSettleReq_proto_goTypes = nil
	file_proto_MusicGameSettleReq_proto_depIdxs = nil
}