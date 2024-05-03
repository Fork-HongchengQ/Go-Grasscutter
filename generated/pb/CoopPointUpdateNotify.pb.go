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
// source: proto/CoopPointUpdateNotify.proto

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

// CmdId: 7401
// Obf: DELDCGOMOEG
type CoopPointUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoopPoint *CoopPoint `protobuf:"bytes,6,opt,name=coop_point,json=coopPoint,proto3" json:"coop_point,omitempty"`
}

func (x *CoopPointUpdateNotify) Reset() {
	*x = CoopPointUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CoopPointUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoopPointUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoopPointUpdateNotify) ProtoMessage() {}

func (x *CoopPointUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CoopPointUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoopPointUpdateNotify.ProtoReflect.Descriptor instead.
func (*CoopPointUpdateNotify) Descriptor() ([]byte, []int) {
	return file_proto_CoopPointUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *CoopPointUpdateNotify) GetCoopPoint() *CoopPoint {
	if x != nil {
		return x.CoopPoint
	}
	return nil
}

var File_proto_CoopPointUpdateNotify_proto protoreflect.FileDescriptor

var file_proto_CoopPointUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6f, 0x70, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6f, 0x70, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x15, 0x43, 0x6f,
	0x6f, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x29, 0x0a, 0x0a, 0x63, 0x6f, 0x6f, 0x70, 0x5f, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x43, 0x6f, 0x6f, 0x70, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x09, 0x63, 0x6f, 0x6f, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x1d,
	0x5a, 0x1b, 0x47, 0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_CoopPointUpdateNotify_proto_rawDescOnce sync.Once
	file_proto_CoopPointUpdateNotify_proto_rawDescData = file_proto_CoopPointUpdateNotify_proto_rawDesc
)

func file_proto_CoopPointUpdateNotify_proto_rawDescGZIP() []byte {
	file_proto_CoopPointUpdateNotify_proto_rawDescOnce.Do(func() {
		file_proto_CoopPointUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_CoopPointUpdateNotify_proto_rawDescData)
	})
	return file_proto_CoopPointUpdateNotify_proto_rawDescData
}

var file_proto_CoopPointUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_CoopPointUpdateNotify_proto_goTypes = []interface{}{
	(*CoopPointUpdateNotify)(nil), // 0: CoopPointUpdateNotify
	(*CoopPoint)(nil),             // 1: CoopPoint
}
var file_proto_CoopPointUpdateNotify_proto_depIdxs = []int32{
	1, // 0: CoopPointUpdateNotify.coop_point:type_name -> CoopPoint
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_CoopPointUpdateNotify_proto_init() }
func file_proto_CoopPointUpdateNotify_proto_init() {
	if File_proto_CoopPointUpdateNotify_proto != nil {
		return
	}
	file_proto_CoopPoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_CoopPointUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoopPointUpdateNotify); i {
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
			RawDescriptor: file_proto_CoopPointUpdateNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_CoopPointUpdateNotify_proto_goTypes,
		DependencyIndexes: file_proto_CoopPointUpdateNotify_proto_depIdxs,
		MessageInfos:      file_proto_CoopPointUpdateNotify_proto_msgTypes,
	}.Build()
	File_proto_CoopPointUpdateNotify_proto = out.File
	file_proto_CoopPointUpdateNotify_proto_rawDesc = nil
	file_proto_CoopPointUpdateNotify_proto_goTypes = nil
	file_proto_CoopPointUpdateNotify_proto_depIdxs = nil
}