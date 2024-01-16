// Copyright (C) 2023  mieru authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: egress.proto

package appctlpb

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

type ProxyProtocol int32

const (
	ProxyProtocol_UNKNOWN_PROXY_PROTOCOL ProxyProtocol = 0
	ProxyProtocol_SOCKS5_PROXY_PROTOCOL  ProxyProtocol = 1
)

// Enum value maps for ProxyProtocol.
var (
	ProxyProtocol_name = map[int32]string{
		0: "UNKNOWN_PROXY_PROTOCOL",
		1: "SOCKS5_PROXY_PROTOCOL",
	}
	ProxyProtocol_value = map[string]int32{
		"UNKNOWN_PROXY_PROTOCOL": 0,
		"SOCKS5_PROXY_PROTOCOL":  1,
	}
)

func (x ProxyProtocol) Enum() *ProxyProtocol {
	p := new(ProxyProtocol)
	*p = x
	return p
}

func (x ProxyProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProxyProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_egress_proto_enumTypes[0].Descriptor()
}

func (ProxyProtocol) Type() protoreflect.EnumType {
	return &file_egress_proto_enumTypes[0]
}

func (x ProxyProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProxyProtocol.Descriptor instead.
func (ProxyProtocol) EnumDescriptor() ([]byte, []int) {
	return file_egress_proto_rawDescGZIP(), []int{0}
}

type EgressAction int32

const (
	// Use proxy to connect to the destination.
	EgressAction_PROXY EgressAction = 0
	// Directly connect to the destination.
	EgressAction_DIRECT EgressAction = 1
	// Do not connect to the destination.
	EgressAction_REJECT EgressAction = 2
)

// Enum value maps for EgressAction.
var (
	EgressAction_name = map[int32]string{
		0: "PROXY",
		1: "DIRECT",
		2: "REJECT",
	}
	EgressAction_value = map[string]int32{
		"PROXY":  0,
		"DIRECT": 1,
		"REJECT": 2,
	}
)

func (x EgressAction) Enum() *EgressAction {
	p := new(EgressAction)
	*p = x
	return p
}

func (x EgressAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EgressAction) Descriptor() protoreflect.EnumDescriptor {
	return file_egress_proto_enumTypes[1].Descriptor()
}

func (EgressAction) Type() protoreflect.EnumType {
	return &file_egress_proto_enumTypes[1]
}

func (x EgressAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EgressAction.Descriptor instead.
func (EgressAction) EnumDescriptor() ([]byte, []int) {
	return file_egress_proto_rawDescGZIP(), []int{1}
}

type EgressProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string        `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Protocol *ProxyProtocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=appctl.ProxyProtocol,oneof" json:"protocol,omitempty"`
	// Proxy IP address or domain name.
	Host *string `protobuf:"bytes,3,opt,name=host,proto3,oneof" json:"host,omitempty"`
	// Proxy port number.
	Port *int32 `protobuf:"varint,4,opt,name=port,proto3,oneof" json:"port,omitempty"`
}

func (x *EgressProxy) Reset() {
	*x = EgressProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_egress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressProxy) ProtoMessage() {}

func (x *EgressProxy) ProtoReflect() protoreflect.Message {
	mi := &file_egress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressProxy.ProtoReflect.Descriptor instead.
func (*EgressProxy) Descriptor() ([]byte, []int) {
	return file_egress_proto_rawDescGZIP(), []int{0}
}

func (x *EgressProxy) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *EgressProxy) GetProtocol() ProxyProtocol {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return ProxyProtocol_UNKNOWN_PROXY_PROTOCOL
}

func (x *EgressProxy) GetHost() string {
	if x != nil && x.Host != nil {
		return *x.Host
	}
	return ""
}

func (x *EgressProxy) GetPort() int32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

type EgressRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of CIDR to match the rule.
	// Use "*" to match all IP addresses.
	IpRanges []string `protobuf:"bytes,1,rep,name=ipRanges,proto3" json:"ipRanges,omitempty"`
	// A list of domain names to match the rule.
	// Use "*" to match all domain names.
	DomainNames []string `protobuf:"bytes,2,rep,name=domainNames,proto3" json:"domainNames,omitempty"`
	// The action to do when the rule is matched.
	Action *EgressAction `protobuf:"varint,3,opt,name=action,proto3,enum=appctl.EgressAction,oneof" json:"action,omitempty"`
	// The name of proxy to connect.
	// This is required when the action is PROXY.
	ProxyName *string `protobuf:"bytes,4,opt,name=proxyName,proto3,oneof" json:"proxyName,omitempty"`
}

func (x *EgressRule) Reset() {
	*x = EgressRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_egress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressRule) ProtoMessage() {}

func (x *EgressRule) ProtoReflect() protoreflect.Message {
	mi := &file_egress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressRule.ProtoReflect.Descriptor instead.
func (*EgressRule) Descriptor() ([]byte, []int) {
	return file_egress_proto_rawDescGZIP(), []int{1}
}

func (x *EgressRule) GetIpRanges() []string {
	if x != nil {
		return x.IpRanges
	}
	return nil
}

func (x *EgressRule) GetDomainNames() []string {
	if x != nil {
		return x.DomainNames
	}
	return nil
}

func (x *EgressRule) GetAction() EgressAction {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return EgressAction_PROXY
}

func (x *EgressRule) GetProxyName() string {
	if x != nil && x.ProxyName != nil {
		return *x.ProxyName
	}
	return ""
}

type Egress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of proxies.
	Proxies []*EgressProxy `protobuf:"bytes,1,rep,name=proxies,proto3" json:"proxies,omitempty"`
	// A list of rules.
	// If no rule is matched, the default action is DIRECT.
	Rules []*EgressRule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *Egress) Reset() {
	*x = Egress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_egress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Egress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Egress) ProtoMessage() {}

func (x *Egress) ProtoReflect() protoreflect.Message {
	mi := &file_egress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Egress.ProtoReflect.Descriptor instead.
func (*Egress) Descriptor() ([]byte, []int) {
	return file_egress_proto_rawDescGZIP(), []int{2}
}

func (x *Egress) GetProxies() []*EgressProxy {
	if x != nil {
		return x.Proxies
	}
	return nil
}

func (x *Egress) GetRules() []*EgressRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

var File_egress_proto protoreflect.FileDescriptor

var file_egress_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x22, 0xb8, 0x01, 0x0a, 0x0b, 0x45, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0xb9, 0x01, 0x0a, 0x0a, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x75, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x70, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x70, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x31,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x61, 0x0a,
	0x06, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74,
	0x6c, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x78, 0x69, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2a, 0x46, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x52, 0x4f,
	0x58, 0x59, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x53, 0x4f, 0x43, 0x4b, 0x53, 0x35, 0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x5f, 0x50, 0x52,
	0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x01, 0x2a, 0x31, 0x0a, 0x0c, 0x45, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x4f, 0x58,
	0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x66, 0x65, 0x69, 0x6e,
	0x2f, 0x6d, 0x69, 0x65, 0x72, 0x75, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x74,
	0x6c, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_egress_proto_rawDescOnce sync.Once
	file_egress_proto_rawDescData = file_egress_proto_rawDesc
)

func file_egress_proto_rawDescGZIP() []byte {
	file_egress_proto_rawDescOnce.Do(func() {
		file_egress_proto_rawDescData = protoimpl.X.CompressGZIP(file_egress_proto_rawDescData)
	})
	return file_egress_proto_rawDescData
}

var file_egress_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_egress_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_egress_proto_goTypes = []interface{}{
	(ProxyProtocol)(0),  // 0: appctl.ProxyProtocol
	(EgressAction)(0),   // 1: appctl.EgressAction
	(*EgressProxy)(nil), // 2: appctl.EgressProxy
	(*EgressRule)(nil),  // 3: appctl.EgressRule
	(*Egress)(nil),      // 4: appctl.Egress
}
var file_egress_proto_depIdxs = []int32{
	0, // 0: appctl.EgressProxy.protocol:type_name -> appctl.ProxyProtocol
	1, // 1: appctl.EgressRule.action:type_name -> appctl.EgressAction
	2, // 2: appctl.Egress.proxies:type_name -> appctl.EgressProxy
	3, // 3: appctl.Egress.rules:type_name -> appctl.EgressRule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_egress_proto_init() }
func file_egress_proto_init() {
	if File_egress_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_egress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressProxy); i {
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
		file_egress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressRule); i {
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
		file_egress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Egress); i {
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
	file_egress_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_egress_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_egress_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_egress_proto_goTypes,
		DependencyIndexes: file_egress_proto_depIdxs,
		EnumInfos:         file_egress_proto_enumTypes,
		MessageInfos:      file_egress_proto_msgTypes,
	}.Build()
	File_egress_proto = out.File
	file_egress_proto_rawDesc = nil
	file_egress_proto_goTypes = nil
	file_egress_proto_depIdxs = nil
}