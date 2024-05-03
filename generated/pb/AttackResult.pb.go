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
// source: proto/AttackResult.proto

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

// Obf: ICJAFHIJEBJ
type AttackResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BLJNCIEFOED           uint32                 `protobuf:"varint,5,opt,name=BLJNCIEFOED,proto3" json:"BLJNCIEFOED,omitempty"`
	ElementType           uint32                 `protobuf:"varint,10,opt,name=element_type,json=elementType,proto3" json:"element_type,omitempty"`
	FMPIEJOMIFJ           float32                `protobuf:"fixed32,1723,opt,name=FMPIEJOMIFJ,proto3" json:"FMPIEJOMIFJ,omitempty"`
	NLAAOODDKGK           float32                `protobuf:"fixed32,778,opt,name=NLAAOODDKGK,proto3" json:"NLAAOODDKGK,omitempty"`
	BAAHKFCEBMA           bool                   `protobuf:"varint,7,opt,name=BAAHKFCEBMA,proto3" json:"BAAHKFCEBMA,omitempty"`
	KFCPIKCMIOI           float32                `protobuf:"fixed32,567,opt,name=KFCPIKCMIOI,proto3" json:"KFCPIKCMIOI,omitempty"`
	DPHDFDJJNOA           float32                `protobuf:"fixed32,27,opt,name=DPHDFDJJNOA,proto3" json:"DPHDFDJJNOA,omitempty"`
	BILAJCEBMGF           bool                   `protobuf:"varint,378,opt,name=BILAJCEBMGF,proto3" json:"BILAJCEBMGF,omitempty"`
	MIKMJNBAAIJ           uint32                 `protobuf:"varint,8,opt,name=MIKMJNBAAIJ,proto3" json:"MIKMJNBAAIJ,omitempty"`
	HitEffResult          *AttackHitEffectResult `protobuf:"bytes,6,opt,name=hit_eff_result,json=hitEffResult,proto3" json:"hit_eff_result,omitempty"`
	BKKNPCPEDCF           uint32                 `protobuf:"varint,457,opt,name=BKKNPCPEDCF,proto3" json:"BKKNPCPEDCF,omitempty"`
	DKDBHEJFOGD           uint32                 `protobuf:"varint,1602,opt,name=DKDBHEJFOGD,proto3" json:"DKDBHEJFOGD,omitempty"`
	DefenseId             uint32                 `protobuf:"varint,9,opt,name=defense_id,json=defenseId,proto3" json:"defense_id,omitempty"`
	AttackerId            uint32                 `protobuf:"varint,14,opt,name=attacker_id,json=attackerId,proto3" json:"attacker_id,omitempty"`
	Damage                float32                `protobuf:"fixed32,2,opt,name=damage,proto3" json:"damage,omitempty"`
	EMOHABNNHBJ           uint32                 `protobuf:"varint,314,opt,name=EMOHABNNHBJ,proto3" json:"EMOHABNNHBJ,omitempty"`
	KAKHNCAHEOP           bool                   `protobuf:"varint,609,opt,name=KAKHNCAHEOP,proto3" json:"KAKHNCAHEOP,omitempty"`
	AnimEventId           string                 `protobuf:"bytes,15,opt,name=anim_event_id,json=animEventId,proto3" json:"anim_event_id,omitempty"`
	HitRetreatAngleCompat int32                  `protobuf:"varint,13,opt,name=hit_retreat_angle_compat,json=hitRetreatAngleCompat,proto3" json:"hit_retreat_angle_compat,omitempty"`
	KGKNJIIOPMP           uint32                 `protobuf:"varint,1649,opt,name=KGKNJIIOPMP,proto3" json:"KGKNJIIOPMP,omitempty"`
	EFKGDDIGOHJ           uint32                 `protobuf:"varint,82,opt,name=EFKGDDIGOHJ,proto3" json:"EFKGDDIGOHJ,omitempty"`
	JDPELIFFANG           uint32                 `protobuf:"varint,394,opt,name=JDPELIFFANG,proto3" json:"JDPELIFFANG,omitempty"`
	ODBKCAJFBMO           uint32                 `protobuf:"varint,1652,opt,name=ODBKCAJFBMO,proto3" json:"ODBKCAJFBMO,omitempty"`
	AbilityIdentifier     *AbilityIdentifier     `protobuf:"bytes,3,opt,name=ability_identifier,json=abilityIdentifier,proto3" json:"ability_identifier,omitempty"`
	ResolvedDir           *Vector                `protobuf:"bytes,4,opt,name=resolved_dir,json=resolvedDir,proto3" json:"resolved_dir,omitempty"`
	HitCollision          *HitCollision          `protobuf:"bytes,11,opt,name=hit_collision,json=hitCollision,proto3" json:"hit_collision,omitempty"`
	NFDLHHGPLKL           uint32                 `protobuf:"varint,866,opt,name=NFDLHHGPLKL,proto3" json:"NFDLHHGPLKL,omitempty"`
	OEEFAHMCLOM           uint32                 `protobuf:"varint,216,opt,name=OEEFAHMCLOM,proto3" json:"OEEFAHMCLOM,omitempty"`
	IEJNKCBILGI           uint32                 `protobuf:"varint,1603,opt,name=IEJNKCBILGI,proto3" json:"IEJNKCBILGI,omitempty"`
	FIHKEJMFKHO           bool                   `protobuf:"varint,1672,opt,name=FIHKEJMFKHO,proto3" json:"FIHKEJMFKHO,omitempty"`
}

func (x *AttackResult) Reset() {
	*x = AttackResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_AttackResult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttackResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackResult) ProtoMessage() {}

func (x *AttackResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_AttackResult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackResult.ProtoReflect.Descriptor instead.
func (*AttackResult) Descriptor() ([]byte, []int) {
	return file_proto_AttackResult_proto_rawDescGZIP(), []int{0}
}

func (x *AttackResult) GetBLJNCIEFOED() uint32 {
	if x != nil {
		return x.BLJNCIEFOED
	}
	return 0
}

func (x *AttackResult) GetElementType() uint32 {
	if x != nil {
		return x.ElementType
	}
	return 0
}

func (x *AttackResult) GetFMPIEJOMIFJ() float32 {
	if x != nil {
		return x.FMPIEJOMIFJ
	}
	return 0
}

func (x *AttackResult) GetNLAAOODDKGK() float32 {
	if x != nil {
		return x.NLAAOODDKGK
	}
	return 0
}

func (x *AttackResult) GetBAAHKFCEBMA() bool {
	if x != nil {
		return x.BAAHKFCEBMA
	}
	return false
}

func (x *AttackResult) GetKFCPIKCMIOI() float32 {
	if x != nil {
		return x.KFCPIKCMIOI
	}
	return 0
}

func (x *AttackResult) GetDPHDFDJJNOA() float32 {
	if x != nil {
		return x.DPHDFDJJNOA
	}
	return 0
}

func (x *AttackResult) GetBILAJCEBMGF() bool {
	if x != nil {
		return x.BILAJCEBMGF
	}
	return false
}

func (x *AttackResult) GetMIKMJNBAAIJ() uint32 {
	if x != nil {
		return x.MIKMJNBAAIJ
	}
	return 0
}

func (x *AttackResult) GetHitEffResult() *AttackHitEffectResult {
	if x != nil {
		return x.HitEffResult
	}
	return nil
}

func (x *AttackResult) GetBKKNPCPEDCF() uint32 {
	if x != nil {
		return x.BKKNPCPEDCF
	}
	return 0
}

func (x *AttackResult) GetDKDBHEJFOGD() uint32 {
	if x != nil {
		return x.DKDBHEJFOGD
	}
	return 0
}

func (x *AttackResult) GetDefenseId() uint32 {
	if x != nil {
		return x.DefenseId
	}
	return 0
}

func (x *AttackResult) GetAttackerId() uint32 {
	if x != nil {
		return x.AttackerId
	}
	return 0
}

func (x *AttackResult) GetDamage() float32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

func (x *AttackResult) GetEMOHABNNHBJ() uint32 {
	if x != nil {
		return x.EMOHABNNHBJ
	}
	return 0
}

func (x *AttackResult) GetKAKHNCAHEOP() bool {
	if x != nil {
		return x.KAKHNCAHEOP
	}
	return false
}

func (x *AttackResult) GetAnimEventId() string {
	if x != nil {
		return x.AnimEventId
	}
	return ""
}

func (x *AttackResult) GetHitRetreatAngleCompat() int32 {
	if x != nil {
		return x.HitRetreatAngleCompat
	}
	return 0
}

func (x *AttackResult) GetKGKNJIIOPMP() uint32 {
	if x != nil {
		return x.KGKNJIIOPMP
	}
	return 0
}

func (x *AttackResult) GetEFKGDDIGOHJ() uint32 {
	if x != nil {
		return x.EFKGDDIGOHJ
	}
	return 0
}

func (x *AttackResult) GetJDPELIFFANG() uint32 {
	if x != nil {
		return x.JDPELIFFANG
	}
	return 0
}

func (x *AttackResult) GetODBKCAJFBMO() uint32 {
	if x != nil {
		return x.ODBKCAJFBMO
	}
	return 0
}

func (x *AttackResult) GetAbilityIdentifier() *AbilityIdentifier {
	if x != nil {
		return x.AbilityIdentifier
	}
	return nil
}

func (x *AttackResult) GetResolvedDir() *Vector {
	if x != nil {
		return x.ResolvedDir
	}
	return nil
}

func (x *AttackResult) GetHitCollision() *HitCollision {
	if x != nil {
		return x.HitCollision
	}
	return nil
}

func (x *AttackResult) GetNFDLHHGPLKL() uint32 {
	if x != nil {
		return x.NFDLHHGPLKL
	}
	return 0
}

func (x *AttackResult) GetOEEFAHMCLOM() uint32 {
	if x != nil {
		return x.OEEFAHMCLOM
	}
	return 0
}

func (x *AttackResult) GetIEJNKCBILGI() uint32 {
	if x != nil {
		return x.IEJNKCBILGI
	}
	return 0
}

func (x *AttackResult) GetFIHKEJMFKHO() bool {
	if x != nil {
		return x.FIHKEJMFKHO
	}
	return false
}

var File_proto_AttackResult_proto protoreflect.FileDescriptor

var file_proto_AttackResult_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x48, 0x69, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x48, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x08, 0x0a, 0x0c, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x4c, 0x4a, 0x4e, 0x43, 0x49, 0x45, 0x46, 0x4f, 0x45, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x42, 0x4c, 0x4a, 0x4e, 0x43, 0x49, 0x45, 0x46, 0x4f, 0x45, 0x44, 0x12, 0x21, 0x0a,
	0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x21, 0x0a, 0x0b, 0x46, 0x4d, 0x50, 0x49, 0x45, 0x4a, 0x4f, 0x4d, 0x49, 0x46, 0x4a, 0x18,
	0xbb, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x46, 0x4d, 0x50, 0x49, 0x45, 0x4a, 0x4f, 0x4d,
	0x49, 0x46, 0x4a, 0x12, 0x21, 0x0a, 0x0b, 0x4e, 0x4c, 0x41, 0x41, 0x4f, 0x4f, 0x44, 0x44, 0x4b,
	0x47, 0x4b, 0x18, 0x8a, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x4e, 0x4c, 0x41, 0x41, 0x4f,
	0x4f, 0x44, 0x44, 0x4b, 0x47, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x41, 0x41, 0x48, 0x4b, 0x46,
	0x43, 0x45, 0x42, 0x4d, 0x41, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x41, 0x41,
	0x48, 0x4b, 0x46, 0x43, 0x45, 0x42, 0x4d, 0x41, 0x12, 0x21, 0x0a, 0x0b, 0x4b, 0x46, 0x43, 0x50,
	0x49, 0x4b, 0x43, 0x4d, 0x49, 0x4f, 0x49, 0x18, 0xb7, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x4b, 0x46, 0x43, 0x50, 0x49, 0x4b, 0x43, 0x4d, 0x49, 0x4f, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x50, 0x48, 0x44, 0x46, 0x44, 0x4a, 0x4a, 0x4e, 0x4f, 0x41, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x44, 0x50, 0x48, 0x44, 0x46, 0x44, 0x4a, 0x4a, 0x4e, 0x4f, 0x41, 0x12, 0x21, 0x0a,
	0x0b, 0x42, 0x49, 0x4c, 0x41, 0x4a, 0x43, 0x45, 0x42, 0x4d, 0x47, 0x46, 0x18, 0xfa, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x49, 0x4c, 0x41, 0x4a, 0x43, 0x45, 0x42, 0x4d, 0x47, 0x46,
	0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x49, 0x4b, 0x4d, 0x4a, 0x4e, 0x42, 0x41, 0x41, 0x49, 0x4a, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x49, 0x4b, 0x4d, 0x4a, 0x4e, 0x42, 0x41, 0x41,
	0x49, 0x4a, 0x12, 0x3c, 0x0a, 0x0e, 0x68, 0x69, 0x74, 0x5f, 0x65, 0x66, 0x66, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x48, 0x69, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x0c, 0x68, 0x69, 0x74, 0x45, 0x66, 0x66, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x21, 0x0a, 0x0b, 0x42, 0x4b, 0x4b, 0x4e, 0x50, 0x43, 0x50, 0x45, 0x44, 0x43, 0x46, 0x18,
	0xc9, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x4b, 0x4b, 0x4e, 0x50, 0x43, 0x50, 0x45,
	0x44, 0x43, 0x46, 0x12, 0x21, 0x0a, 0x0b, 0x44, 0x4b, 0x44, 0x42, 0x48, 0x45, 0x4a, 0x46, 0x4f,
	0x47, 0x44, 0x18, 0xc2, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x4b, 0x44, 0x42, 0x48,
	0x45, 0x4a, 0x46, 0x4f, 0x47, 0x44, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x73,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x65, 0x66, 0x65,
	0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21,
	0x0a, 0x0b, 0x45, 0x4d, 0x4f, 0x48, 0x41, 0x42, 0x4e, 0x4e, 0x48, 0x42, 0x4a, 0x18, 0xba, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x4d, 0x4f, 0x48, 0x41, 0x42, 0x4e, 0x4e, 0x48, 0x42,
	0x4a, 0x12, 0x21, 0x0a, 0x0b, 0x4b, 0x41, 0x4b, 0x48, 0x4e, 0x43, 0x41, 0x48, 0x45, 0x4f, 0x50,
	0x18, 0xe1, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4b, 0x41, 0x4b, 0x48, 0x4e, 0x43, 0x41,
	0x48, 0x45, 0x4f, 0x50, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x6e, 0x69, 0x6d, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6e, 0x69,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x68, 0x69, 0x74, 0x5f,
	0x72, 0x65, 0x74, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x68, 0x69, 0x74, 0x52,
	0x65, 0x74, 0x72, 0x65, 0x61, 0x74, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x74, 0x12, 0x21, 0x0a, 0x0b, 0x4b, 0x47, 0x4b, 0x4e, 0x4a, 0x49, 0x49, 0x4f, 0x50, 0x4d, 0x50,
	0x18, 0xf1, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x47, 0x4b, 0x4e, 0x4a, 0x49, 0x49,
	0x4f, 0x50, 0x4d, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x46, 0x4b, 0x47, 0x44, 0x44, 0x49, 0x47,
	0x4f, 0x48, 0x4a, 0x18, 0x52, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x46, 0x4b, 0x47, 0x44,
	0x44, 0x49, 0x47, 0x4f, 0x48, 0x4a, 0x12, 0x21, 0x0a, 0x0b, 0x4a, 0x44, 0x50, 0x45, 0x4c, 0x49,
	0x46, 0x46, 0x41, 0x4e, 0x47, 0x18, 0x8a, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x44,
	0x50, 0x45, 0x4c, 0x49, 0x46, 0x46, 0x41, 0x4e, 0x47, 0x12, 0x21, 0x0a, 0x0b, 0x4f, 0x44, 0x42,
	0x4b, 0x43, 0x41, 0x4a, 0x46, 0x42, 0x4d, 0x4f, 0x18, 0xf4, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x4f, 0x44, 0x42, 0x4b, 0x43, 0x41, 0x4a, 0x46, 0x42, 0x4d, 0x4f, 0x12, 0x41, 0x0a, 0x12,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x11, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x44, 0x69, 0x72, 0x12, 0x32, 0x0a, 0x0d, 0x68,
	0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x48, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x68, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0b, 0x4e, 0x46, 0x44, 0x4c, 0x48, 0x48, 0x47, 0x50, 0x4c, 0x4b, 0x4c, 0x18, 0xe2,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x46, 0x44, 0x4c, 0x48, 0x48, 0x47, 0x50, 0x4c,
	0x4b, 0x4c, 0x12, 0x21, 0x0a, 0x0b, 0x4f, 0x45, 0x45, 0x46, 0x41, 0x48, 0x4d, 0x43, 0x4c, 0x4f,
	0x4d, 0x18, 0xd8, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x45, 0x45, 0x46, 0x41, 0x48,
	0x4d, 0x43, 0x4c, 0x4f, 0x4d, 0x12, 0x21, 0x0a, 0x0b, 0x49, 0x45, 0x4a, 0x4e, 0x4b, 0x43, 0x42,
	0x49, 0x4c, 0x47, 0x49, 0x18, 0xc3, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x45, 0x4a,
	0x4e, 0x4b, 0x43, 0x42, 0x49, 0x4c, 0x47, 0x49, 0x12, 0x21, 0x0a, 0x0b, 0x46, 0x49, 0x48, 0x4b,
	0x45, 0x4a, 0x4d, 0x46, 0x4b, 0x48, 0x4f, 0x18, 0x88, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x46, 0x49, 0x48, 0x4b, 0x45, 0x4a, 0x4d, 0x46, 0x4b, 0x48, 0x4f, 0x42, 0x1d, 0x5a, 0x1b, 0x47,
	0x6f, 0x2d, 0x47, 0x72, 0x61, 0x73, 0x73, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_AttackResult_proto_rawDescOnce sync.Once
	file_proto_AttackResult_proto_rawDescData = file_proto_AttackResult_proto_rawDesc
)

func file_proto_AttackResult_proto_rawDescGZIP() []byte {
	file_proto_AttackResult_proto_rawDescOnce.Do(func() {
		file_proto_AttackResult_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_AttackResult_proto_rawDescData)
	})
	return file_proto_AttackResult_proto_rawDescData
}

var file_proto_AttackResult_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_AttackResult_proto_goTypes = []interface{}{
	(*AttackResult)(nil),          // 0: AttackResult
	(*AttackHitEffectResult)(nil), // 1: AttackHitEffectResult
	(*AbilityIdentifier)(nil),     // 2: AbilityIdentifier
	(*Vector)(nil),                // 3: Vector
	(*HitCollision)(nil),          // 4: HitCollision
}
var file_proto_AttackResult_proto_depIdxs = []int32{
	1, // 0: AttackResult.hit_eff_result:type_name -> AttackHitEffectResult
	2, // 1: AttackResult.ability_identifier:type_name -> AbilityIdentifier
	3, // 2: AttackResult.resolved_dir:type_name -> Vector
	4, // 3: AttackResult.hit_collision:type_name -> HitCollision
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_AttackResult_proto_init() }
func file_proto_AttackResult_proto_init() {
	if File_proto_AttackResult_proto != nil {
		return
	}
	file_proto_AttackHitEffectResult_proto_init()
	file_proto_AbilityIdentifier_proto_init()
	file_proto_Vector_proto_init()
	file_proto_HitCollision_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_AttackResult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttackResult); i {
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
			RawDescriptor: file_proto_AttackResult_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_AttackResult_proto_goTypes,
		DependencyIndexes: file_proto_AttackResult_proto_depIdxs,
		MessageInfos:      file_proto_AttackResult_proto_msgTypes,
	}.Build()
	File_proto_AttackResult_proto = out.File
	file_proto_AttackResult_proto_rawDesc = nil
	file_proto_AttackResult_proto_goTypes = nil
	file_proto_AttackResult_proto_depIdxs = nil
}