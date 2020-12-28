// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: check_if_email_exist.proto

package check_if_email_exist

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input       string  `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	IsReachable string  `protobuf:"bytes,2,opt,name=is_reachable,proto3" json:"is_reachable,omitempty"`
	Misc        *Misc   `protobuf:"bytes,3,opt,name=misc,proto3" json:"misc,omitempty"`
	Mx          *MX     `protobuf:"bytes,4,opt,name=mx,proto3" json:"mx,omitempty"`
	Smtp        *SMTP   `protobuf:"bytes,5,opt,name=smtp,proto3" json:"smtp,omitempty"`
	Syntax      *Syntax `protobuf:"bytes,6,opt,name=syntax,proto3" json:"syntax,omitempty"`
	Error       string  `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_if_email_exist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_check_if_email_exist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_check_if_email_exist_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *Result) GetIsReachable() string {
	if x != nil {
		return x.IsReachable
	}
	return ""
}

func (x *Result) GetMisc() *Misc {
	if x != nil {
		return x.Misc
	}
	return nil
}

func (x *Result) GetMx() *MX {
	if x != nil {
		return x.Mx
	}
	return nil
}

func (x *Result) GetSmtp() *SMTP {
	if x != nil {
		return x.Smtp
	}
	return nil
}

func (x *Result) GetSyntax() *Syntax {
	if x != nil {
		return x.Syntax
	}
	return nil
}

func (x *Result) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Misc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDisposable  bool `protobuf:"varint,1,opt,name=is_disposable,proto3" json:"is_disposable,omitempty"`
	IsRoleAccount bool `protobuf:"varint,2,opt,name=is_role_account,proto3" json:"is_role_account,omitempty"`
}

func (x *Misc) Reset() {
	*x = Misc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_if_email_exist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Misc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Misc) ProtoMessage() {}

func (x *Misc) ProtoReflect() protoreflect.Message {
	mi := &file_check_if_email_exist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Misc.ProtoReflect.Descriptor instead.
func (*Misc) Descriptor() ([]byte, []int) {
	return file_check_if_email_exist_proto_rawDescGZIP(), []int{1}
}

func (x *Misc) GetIsDisposable() bool {
	if x != nil {
		return x.IsDisposable
	}
	return false
}

func (x *Misc) GetIsRoleAccount() bool {
	if x != nil {
		return x.IsRoleAccount
	}
	return false
}

type MX struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AcceptsMail bool     `protobuf:"varint,1,opt,name=accepts_mail,proto3" json:"accepts_mail,omitempty"`
	Records     []string `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *MX) Reset() {
	*x = MX{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_if_email_exist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MX) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MX) ProtoMessage() {}

func (x *MX) ProtoReflect() protoreflect.Message {
	mi := &file_check_if_email_exist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MX.ProtoReflect.Descriptor instead.
func (*MX) Descriptor() ([]byte, []int) {
	return file_check_if_email_exist_proto_rawDescGZIP(), []int{2}
}

func (x *MX) GetAcceptsMail() bool {
	if x != nil {
		return x.AcceptsMail
	}
	return false
}

func (x *MX) GetRecords() []string {
	if x != nil {
		return x.Records
	}
	return nil
}

type SMTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanConnectSmtp bool `protobuf:"varint,1,opt,name=can_connect_smtp,proto3" json:"can_connect_smtp,omitempty"`
	HasFullInbox   bool `protobuf:"varint,2,opt,name=has_full_inbox,proto3" json:"has_full_inbox,omitempty"`
	IsCatchAll     bool `protobuf:"varint,3,opt,name=is_catch_all,proto3" json:"is_catch_all,omitempty"`
	IsDeliverable  bool `protobuf:"varint,4,opt,name=is_deliverable,proto3" json:"is_deliverable,omitempty"`
	IsDisabled     bool `protobuf:"varint,5,opt,name=is_disabled,proto3" json:"is_disabled,omitempty"`
}

func (x *SMTP) Reset() {
	*x = SMTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_if_email_exist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMTP) ProtoMessage() {}

func (x *SMTP) ProtoReflect() protoreflect.Message {
	mi := &file_check_if_email_exist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMTP.ProtoReflect.Descriptor instead.
func (*SMTP) Descriptor() ([]byte, []int) {
	return file_check_if_email_exist_proto_rawDescGZIP(), []int{3}
}

func (x *SMTP) GetCanConnectSmtp() bool {
	if x != nil {
		return x.CanConnectSmtp
	}
	return false
}

func (x *SMTP) GetHasFullInbox() bool {
	if x != nil {
		return x.HasFullInbox
	}
	return false
}

func (x *SMTP) GetIsCatchAll() bool {
	if x != nil {
		return x.IsCatchAll
	}
	return false
}

func (x *SMTP) GetIsDeliverable() bool {
	if x != nil {
		return x.IsDeliverable
	}
	return false
}

func (x *SMTP) GetIsDisabled() bool {
	if x != nil {
		return x.IsDisabled
	}
	return false
}

type Syntax struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address       *wrappers.StringValue `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Domain        string                `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	IsValidSyntax bool                  `protobuf:"varint,3,opt,name=is_valid_syntax,proto3" json:"is_valid_syntax,omitempty"`
	Username      string                `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Syntax) Reset() {
	*x = Syntax{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_if_email_exist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Syntax) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Syntax) ProtoMessage() {}

func (x *Syntax) ProtoReflect() protoreflect.Message {
	mi := &file_check_if_email_exist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Syntax.ProtoReflect.Descriptor instead.
func (*Syntax) Descriptor() ([]byte, []int) {
	return file_check_if_email_exist_proto_rawDescGZIP(), []int{4}
}

func (x *Syntax) GetAddress() *wrappers.StringValue {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Syntax) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Syntax) GetIsValidSyntax() bool {
	if x != nil {
		return x.IsValidSyntax
	}
	return false
}

func (x *Syntax) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_check_if_email_exist_proto protoreflect.FileDescriptor

var file_check_if_email_exist_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x66, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x52, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x67, 0x6f, 0x5f, 0x65,
	0x76, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x69, 0x66, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x90, 0x04, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x63,
	0x68, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x6c, 0x0a, 0x04, 0x6d, 0x69, 0x73, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x58, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x76, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x66, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x4d, 0x69, 0x73, 0x63, 0x52, 0x04, 0x6d,
	0x69, 0x73, 0x63, 0x12, 0x66, 0x0a, 0x02, 0x6d, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x56, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x67, 0x6f, 0x5f, 0x65, 0x76, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x66, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x2e, 0x4d, 0x58, 0x52, 0x02, 0x6d, 0x78, 0x12, 0x6c, 0x0a, 0x04, 0x73,
	0x6d, 0x74, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x58, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x76, 0x5f,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x69, 0x66, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x53,
	0x4d, 0x54, 0x50, 0x52, 0x04, 0x73, 0x6d, 0x74, 0x70, 0x12, 0x72, 0x0a, 0x06, 0x73, 0x79, 0x6e,
	0x74, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x76, 0x5f,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x69, 0x66, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x53,
	0x79, 0x6e, 0x74, 0x61, 0x78, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x56, 0x0a, 0x04, 0x4d, 0x69, 0x73, 0x63, 0x12, 0x24, 0x0a, 0x0d, 0x69,
	0x73, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x02, 0x4d,
	0x58, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x73,
	0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22,
	0xc8, 0x01, 0x0a, 0x04, 0x53, 0x4d, 0x54, 0x50, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x61, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x6d, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x63, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f,
	0x73, 0x6d, 0x74, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x5f, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61,
	0x73, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x22, 0x0a, 0x0c,
	0x69, 0x73, 0x5f, 0x63, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x6c, 0x6c,
	0x12, 0x26, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69,
	0x73, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x9e, 0x01, 0x0a, 0x06, 0x53,
	0x79, 0x6e, 0x74, 0x61, 0x78, 0x12, 0x36, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x50, 0x5a, 0x4e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x67, 0x6f, 0x2d,
	0x65, 0x76, 0x2d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69,
	0x66, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_check_if_email_exist_proto_rawDescOnce sync.Once
	file_check_if_email_exist_proto_rawDescData = file_check_if_email_exist_proto_rawDesc
)

func file_check_if_email_exist_proto_rawDescGZIP() []byte {
	file_check_if_email_exist_proto_rawDescOnce.Do(func() {
		file_check_if_email_exist_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_if_email_exist_proto_rawDescData)
	})
	return file_check_if_email_exist_proto_rawDescData
}

var file_check_if_email_exist_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_check_if_email_exist_proto_goTypes = []interface{}{
	(*Result)(nil),               // 0: github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Result
	(*Misc)(nil),                 // 1: github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Misc
	(*MX)(nil),                   // 2: github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.MX
	(*SMTP)(nil),                 // 3: github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.SMTP
	(*Syntax)(nil),               // 4: github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Syntax
	(*wrappers.StringValue)(nil), // 5: google.protobuf.StringValue
}
var file_check_if_email_exist_proto_depIdxs = []int32{
	1, // 0: github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Result.misc:type_name -> github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Misc
	2, // 1: github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Result.mx:type_name -> github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.MX
	3, // 2: github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Result.smtp:type_name -> github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.SMTP
	4, // 3: github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Result.syntax:type_name -> github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Syntax
	5, // 4: github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Syntax.address:type_name -> google.protobuf.StringValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_check_if_email_exist_proto_init() }
func file_check_if_email_exist_proto_init() {
	if File_check_if_email_exist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_check_if_email_exist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_check_if_email_exist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Misc); i {
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
		file_check_if_email_exist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MX); i {
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
		file_check_if_email_exist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMTP); i {
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
		file_check_if_email_exist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Syntax); i {
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
			RawDescriptor: file_check_if_email_exist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_check_if_email_exist_proto_goTypes,
		DependencyIndexes: file_check_if_email_exist_proto_depIdxs,
		MessageInfos:      file_check_if_email_exist_proto_msgTypes,
	}.Build()
	File_check_if_email_exist_proto = out.File
	file_check_if_email_exist_proto_rawDesc = nil
	file_check_if_email_exist_proto_goTypes = nil
	file_check_if_email_exist_proto_depIdxs = nil
}
