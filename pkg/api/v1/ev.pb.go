// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: ev.proto

package v1

import (
	check_if_email_exist "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/check_if_email_exist"
	mailboxvalidator "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/mailboxvalidator"
	prompt_email_verification_api "github.com/go-email-validator/go-ev-presenters/pkg/api/v1/prompt_email_verification_api"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

//ResultType
//
//\* CHECK_IF_EMAIL_EXIST, CIEE - [check-if-email-exists format](https://github.com/amaurymartiny/check-if-email-exists#%EF%B8%8F-json-output).
//\* MAILBOXVALIDATOR, MAIL_BOX_VALIDATOR, MBV - [mailboxvalidator.com format](https://www.mailboxvalidator.com/api-single-validation).
//\* PROMPT_EMAIL_VERIFICATION_API, PEVA - [Email Verification api format](https://promptapi.com/marketplace/description/email_verification-api) from [promptapi](https://promptapi.com).
type EmailRequest_ResultType int32

const (
	EmailRequest_CHECK_IF_EMAIL_EXIST          EmailRequest_ResultType = 0
	EmailRequest_CIEE                          EmailRequest_ResultType = 0
	EmailRequest_MAILBOXVALIDATOR              EmailRequest_ResultType = 1
	EmailRequest_MAIL_BOX_VALIDATOR            EmailRequest_ResultType = 1
	EmailRequest_MBV                           EmailRequest_ResultType = 1
	EmailRequest_PROMPT_EMAIL_VERIFICATION_API EmailRequest_ResultType = 2
	EmailRequest_PEVA                          EmailRequest_ResultType = 2
)

// Enum value maps for EmailRequest_ResultType.
var (
	EmailRequest_ResultType_name = map[int32]string{
		0: "CHECK_IF_EMAIL_EXIST",
		// Duplicate value: 0: "CIEE",
		1: "MAILBOXVALIDATOR",
		// Duplicate value: 1: "MAIL_BOX_VALIDATOR",
		// Duplicate value: 1: "MBV",
		2: "PROMPT_EMAIL_VERIFICATION_API",
		// Duplicate value: 2: "PEVA",
	}
	EmailRequest_ResultType_value = map[string]int32{
		"CHECK_IF_EMAIL_EXIST":          0,
		"CIEE":                          0,
		"MAILBOXVALIDATOR":              1,
		"MAIL_BOX_VALIDATOR":            1,
		"MBV":                           1,
		"PROMPT_EMAIL_VERIFICATION_API": 2,
		"PEVA":                          2,
	}
)

func (x EmailRequest_ResultType) Enum() *EmailRequest_ResultType {
	p := new(EmailRequest_ResultType)
	*p = x
	return p
}

func (x EmailRequest_ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmailRequest_ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_ev_proto_enumTypes[0].Descriptor()
}

func (EmailRequest_ResultType) Type() protoreflect.EnumType {
	return &file_ev_proto_enumTypes[0]
}

func (x EmailRequest_ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmailRequest_ResultType.Descriptor instead.
func (EmailRequest_ResultType) EnumDescriptor() ([]byte, []int) {
	return file_ev_proto_rawDescGZIP(), []int{0, 0}
}

type EmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string                  `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	ResultType EmailRequest_ResultType `protobuf:"varint,2,opt,name=result_type,json=resultType,proto3,enum=github.com.go_email_validator.go_ev_presenters.api.v1.EmailRequest_ResultType" json:"result_type,omitempty"` // TODO find solution without duplication
}

func (x *EmailRequest) Reset() {
	*x = EmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ev_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequest) ProtoMessage() {}

func (x *EmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ev_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequest.ProtoReflect.Descriptor instead.
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return file_ev_proto_rawDescGZIP(), []int{0}
}

func (x *EmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailRequest) GetResultType() EmailRequest_ResultType {
	if x != nil {
		return x.ResultType
	}
	return EmailRequest_CHECK_IF_EMAIL_EXIST
}

type EmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*EmailResponse_CheckIfEmailExist
	//	*EmailResponse_MailBoxValidator
	//	*EmailResponse_PromptEmailVerificationApi
	Result isEmailResponse_Result `protobuf_oneof:"result"`
}

func (x *EmailResponse) Reset() {
	*x = EmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ev_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailResponse) ProtoMessage() {}

func (x *EmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ev_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailResponse.ProtoReflect.Descriptor instead.
func (*EmailResponse) Descriptor() ([]byte, []int) {
	return file_ev_proto_rawDescGZIP(), []int{1}
}

func (m *EmailResponse) GetResult() isEmailResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *EmailResponse) GetCheckIfEmailExist() *check_if_email_exist.Result {
	if x, ok := x.GetResult().(*EmailResponse_CheckIfEmailExist); ok {
		return x.CheckIfEmailExist
	}
	return nil
}

func (x *EmailResponse) GetMailBoxValidator() *mailboxvalidator.Result {
	if x, ok := x.GetResult().(*EmailResponse_MailBoxValidator); ok {
		return x.MailBoxValidator
	}
	return nil
}

func (x *EmailResponse) GetPromptEmailVerificationApi() *prompt_email_verification_api.Result {
	if x, ok := x.GetResult().(*EmailResponse_PromptEmailVerificationApi); ok {
		return x.PromptEmailVerificationApi
	}
	return nil
}

type isEmailResponse_Result interface {
	isEmailResponse_Result()
}

type EmailResponse_CheckIfEmailExist struct {
	CheckIfEmailExist *check_if_email_exist.Result `protobuf:"bytes,1,opt,name=check_if_email_exist,json=checkIfEmailExist,proto3,oneof"`
}

type EmailResponse_MailBoxValidator struct {
	MailBoxValidator *mailboxvalidator.Result `protobuf:"bytes,2,opt,name=mail_box_validator,json=mailBoxValidator,proto3,oneof"`
}

type EmailResponse_PromptEmailVerificationApi struct {
	PromptEmailVerificationApi *prompt_email_verification_api.Result `protobuf:"bytes,3,opt,name=prompt_email_verification_api,json=promptEmailVerificationApi,proto3,oneof"`
}

func (*EmailResponse_CheckIfEmailExist) isEmailResponse_Result() {}

func (*EmailResponse_MailBoxValidator) isEmailResponse_Result() {}

func (*EmailResponse_PromptEmailVerificationApi) isEmailResponse_Result() {}

var File_ev_proto protoreflect.FileDescriptor

var file_ev_proto_rawDesc = []byte{
	0x0a, 0x08, 0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x76, 0x5f,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x1a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x67, 0x6f, 0x2d, 0x65, 0x76, 0x2d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x66, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x66, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x64, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x67,
	0x6f, 0x2d, 0x65, 0x76, 0x2d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x6d, 0x61,
	0x69, 0x6c, 0x62, 0x6f, 0x78, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6d,
	0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x7e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x67, 0x6f, 0x2d, 0x65, 0x76, 0x2d, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x84, 0x06, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0xc2, 0x04, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x4e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x76, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x42, 0xd0, 0x03, 0x92, 0x41, 0xcc, 0x03, 0x32, 0xc9, 0x03, 0x2a, 0x20, 0x43, 0x48, 0x45, 0x43,
	0x4b, 0x5f, 0x49, 0x46, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54,
	0x2c, 0x20, 0x43, 0x49, 0x45, 0x45, 0x20, 0x2d, 0x20, 0x5b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d,
	0x69, 0x66, 0x2d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x20,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6d, 0x61, 0x75, 0x72,
	0x79, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x79, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d, 0x69,
	0x66, 0x2d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x23, 0x25,
	0x45, 0x46, 0x25, 0x42, 0x38, 0x25, 0x38, 0x46, 0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x29, 0x2e, 0x0a, 0x2a, 0x20, 0x4d, 0x41, 0x49, 0x4c, 0x42, 0x4f, 0x58,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x4f, 0x52, 0x2c, 0x20, 0x4d, 0x41, 0x49, 0x4c, 0x5f,
	0x42, 0x4f, 0x58, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x4f, 0x52, 0x2c, 0x20, 0x4d,
	0x42, 0x56, 0x20, 0x2d, 0x20, 0x5b, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x6d,
	0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x29, 0x2e, 0x0a, 0x2a, 0x20, 0x50, 0x52,
	0x4f, 0x4d, 0x50, 0x54, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x50, 0x49, 0x2c, 0x20, 0x50, 0x45, 0x56,
	0x41, 0x20, 0x2d, 0x20, 0x5b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x70, 0x69, 0x20, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x29, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x5b,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x61, 0x70, 0x69, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x29, 0x2e, 0x0a, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x98, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x14, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x49, 0x46, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c,
	0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x49, 0x45, 0x45,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x41, 0x49, 0x4c, 0x42, 0x4f, 0x58, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x41, 0x49, 0x4c,
	0x5f, 0x42, 0x4f, 0x58, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x01,
	0x12, 0x07, 0x0a, 0x03, 0x4d, 0x42, 0x56, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x52, 0x4f,
	0x4d, 0x50, 0x54, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x50, 0x49, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x50, 0x45, 0x56, 0x41, 0x10, 0x02, 0x1a, 0x02, 0x10, 0x01, 0x22, 0xdd, 0x03, 0x0a, 0x0d, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8d, 0x01, 0x0a,
	0x14, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x66, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x67, 0x6f, 0x5f, 0x65,
	0x76, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x69, 0x66, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x11, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x66, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x86, 0x01, 0x0a,
	0x12, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x56, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x76, 0x5f,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f,
	0x78, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x48, 0x00, 0x52, 0x10, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0xa8, 0x01, 0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x63, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x67, 0x6f,
	0x5f, 0x65, 0x76, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x48, 0x00, 0x52, 0x1a, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x69,
	0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xf4, 0x01, 0x0a, 0x0f, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xe0,
	0x01, 0x0a, 0x10, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x43, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x76, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x67, 0x6f, 0x5f, 0x65, 0x76, 0x5f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3b, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x3a, 0x01, 0x2a,
	0x5a, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x7b, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x7d, 0x42, 0x8f, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x2f, 0x67, 0x6f, 0x2d, 0x65, 0x76, 0x2d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x92,
	0x41, 0x51, 0x12, 0x18, 0x0a, 0x0f, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x5a, 0x23, 0x0a, 0x21,
	0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x13, 0x08, 0x02,
	0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x02, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ev_proto_rawDescOnce sync.Once
	file_ev_proto_rawDescData = file_ev_proto_rawDesc
)

func file_ev_proto_rawDescGZIP() []byte {
	file_ev_proto_rawDescOnce.Do(func() {
		file_ev_proto_rawDescData = protoimpl.X.CompressGZIP(file_ev_proto_rawDescData)
	})
	return file_ev_proto_rawDescData
}

var file_ev_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ev_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ev_proto_goTypes = []interface{}{
	(EmailRequest_ResultType)(0),                 // 0: github.com.go_email_validator.go_ev_presenters.api.v1.EmailRequest.ResultType
	(*EmailRequest)(nil),                         // 1: github.com.go_email_validator.go_ev_presenters.api.v1.EmailRequest
	(*EmailResponse)(nil),                        // 2: github.com.go_email_validator.go_ev_presenters.api.v1.EmailResponse
	(*check_if_email_exist.Result)(nil),          // 3: github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Result
	(*mailboxvalidator.Result)(nil),              // 4: github.com.go_email_validator.go_ev_presenters.pkg.presenters.mailboxvalidator.Result
	(*prompt_email_verification_api.Result)(nil), // 5: github.com.go_email_validator.go_ev_presenters.pkg.presenters.prompt_email_verification_api.Result
}
var file_ev_proto_depIdxs = []int32{
	0, // 0: github.com.go_email_validator.go_ev_presenters.api.v1.EmailRequest.result_type:type_name -> github.com.go_email_validator.go_ev_presenters.api.v1.EmailRequest.ResultType
	3, // 1: github.com.go_email_validator.go_ev_presenters.api.v1.EmailResponse.check_if_email_exist:type_name -> github.com.go_email_validator.go_ev_presenters.pkg.presenters.check_if_email_exist.Result
	4, // 2: github.com.go_email_validator.go_ev_presenters.api.v1.EmailResponse.mail_box_validator:type_name -> github.com.go_email_validator.go_ev_presenters.pkg.presenters.mailboxvalidator.Result
	5, // 3: github.com.go_email_validator.go_ev_presenters.api.v1.EmailResponse.prompt_email_verification_api:type_name -> github.com.go_email_validator.go_ev_presenters.pkg.presenters.prompt_email_verification_api.Result
	1, // 4: github.com.go_email_validator.go_ev_presenters.api.v1.EmailValidation.singleValidation:input_type -> github.com.go_email_validator.go_ev_presenters.api.v1.EmailRequest
	2, // 5: github.com.go_email_validator.go_ev_presenters.api.v1.EmailValidation.singleValidation:output_type -> github.com.go_email_validator.go_ev_presenters.api.v1.EmailResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ev_proto_init() }
func file_ev_proto_init() {
	if File_ev_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ev_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailRequest); i {
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
		file_ev_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailResponse); i {
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
	file_ev_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*EmailResponse_CheckIfEmailExist)(nil),
		(*EmailResponse_MailBoxValidator)(nil),
		(*EmailResponse_PromptEmailVerificationApi)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ev_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ev_proto_goTypes,
		DependencyIndexes: file_ev_proto_depIdxs,
		EnumInfos:         file_ev_proto_enumTypes,
		MessageInfos:      file_ev_proto_msgTypes,
	}.Build()
	File_ev_proto = out.File
	file_ev_proto_rawDesc = nil
	file_ev_proto_goTypes = nil
	file_ev_proto_depIdxs = nil
}
