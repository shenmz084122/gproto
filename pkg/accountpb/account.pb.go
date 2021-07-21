// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/account.proto

package accountpb

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

type ValidateRequestSignatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqMethod      string `protobuf:"bytes,1,opt,name=req_method,json=reqMethod,proto3" json:"req_method,omitempty"`
	ReqPath        string `protobuf:"bytes,2,opt,name=req_path,json=reqPath,proto3" json:"req_path,omitempty"`
	ReqQueryString string `protobuf:"bytes,3,opt,name=req_query_string,json=reqQueryString,proto3" json:"req_query_string,omitempty"`
	ReqBody        string `protobuf:"bytes,4,opt,name=req_body,json=reqBody,proto3" json:"req_body,omitempty"`
	ReqSignature   string `protobuf:"bytes,5,opt,name=req_signature,json=reqSignature,proto3" json:"req_signature,omitempty"`
	ReqAccessKeyId string `protobuf:"bytes,6,opt,name=req_access_key_id,json=reqAccessKeyId,proto3" json:"req_access_key_id,omitempty"`
	ReqSource      string `protobuf:"bytes,7,opt,name=req_source,json=reqSource,proto3" json:"req_source,omitempty"`
}

func (x *ValidateRequestSignatureRequest) Reset() {
	*x = ValidateRequestSignatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequestSignatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequestSignatureRequest) ProtoMessage() {}

func (x *ValidateRequestSignatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequestSignatureRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequestSignatureRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateRequestSignatureRequest) GetReqMethod() string {
	if x != nil {
		return x.ReqMethod
	}
	return ""
}

func (x *ValidateRequestSignatureRequest) GetReqPath() string {
	if x != nil {
		return x.ReqPath
	}
	return ""
}

func (x *ValidateRequestSignatureRequest) GetReqQueryString() string {
	if x != nil {
		return x.ReqQueryString
	}
	return ""
}

func (x *ValidateRequestSignatureRequest) GetReqBody() string {
	if x != nil {
		return x.ReqBody
	}
	return ""
}

func (x *ValidateRequestSignatureRequest) GetReqSignature() string {
	if x != nil {
		return x.ReqSignature
	}
	return ""
}

func (x *ValidateRequestSignatureRequest) GetReqAccessKeyId() string {
	if x != nil {
		return x.ReqAccessKeyId
	}
	return ""
}

func (x *ValidateRequestSignatureRequest) GetReqSource() string {
	if x != nil {
		return x.ReqSource
	}
	return ""
}

type ValidateRequestSignatureReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	UserId  string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ValidateRequestSignatureReply) Reset() {
	*x = ValidateRequestSignatureReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequestSignatureReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequestSignatureReply) ProtoMessage() {}

func (x *ValidateRequestSignatureReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequestSignatureReply.ProtoReflect.Descriptor instead.
func (*ValidateRequestSignatureReply) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateRequestSignatureReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ValidateRequestSignatureReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ValidateRequestSignatureReply) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DescribeUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users     []string `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Limit     int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    int32    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	ReqSource string   `protobuf:"bytes,4,opt,name=req_source,json=reqSource,proto3" json:"req_source,omitempty"`
}

func (x *DescribeUsersRequest) Reset() {
	*x = DescribeUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeUsersRequest) ProtoMessage() {}

func (x *DescribeUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeUsersRequest.ProtoReflect.Descriptor instead.
func (*DescribeUsersRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeUsersRequest) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *DescribeUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *DescribeUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *DescribeUsersRequest) GetReqSource() string {
	if x != nil {
		return x.ReqSource
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName      string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Lang          string `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
	Email         string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Status        string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Role          string `protobuf:"bytes,7,opt,name=role,proto3" json:"role,omitempty"`
	Currency      string `protobuf:"bytes,8,opt,name=currency,proto3" json:"currency,omitempty"`
	GravatarEmail string `protobuf:"bytes,9,opt,name=gravatar_email,json=gravatarEmail,proto3" json:"gravatar_email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *User) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *User) GetGravatarEmail() string {
	if x != nil {
		return x.GravatarEmail
	}
	return ""
}

type DescribeUsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserSet    []*User `protobuf:"bytes,1,rep,name=user_set,json=userSet,proto3" json:"user_set,omitempty"`
	TotalCount int64   `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Status     int32   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Message    string  `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DescribeUsersReply) Reset() {
	*x = DescribeUsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeUsersReply) ProtoMessage() {}

func (x *DescribeUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeUsersReply.ProtoReflect.Descriptor instead.
func (*DescribeUsersReply) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeUsersReply) GetUserSet() []*User {
	if x != nil {
		return x.UserSet
	}
	return nil
}

func (x *DescribeUsersReply) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *DescribeUsersReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DescribeUsersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_account_proto protoreflect.FileDescriptor

var file_proto_account_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x02,
	0x0a, 0x1f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x50, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x72,
	0x65, 0x71, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x5f, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x5f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x72, 0x65, 0x71, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0x7b, 0x0a, 0x1d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a,
	0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04,
	0x10, 0x00, 0x18, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x20, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x71, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xeb, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x25, 0x0a, 0x0e, 0x67, 0x72, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x72, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xa0, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x26, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x10, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa8, 0x01, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x5e, 0x0a, 0x18, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x20, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x15, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e,
	0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_proto_rawDescOnce sync.Once
	file_proto_account_proto_rawDescData = file_proto_account_proto_rawDesc
)

func file_proto_account_proto_rawDescGZIP() []byte {
	file_proto_account_proto_rawDescOnce.Do(func() {
		file_proto_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_proto_rawDescData)
	})
	return file_proto_account_proto_rawDescData
}

var file_proto_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_account_proto_goTypes = []interface{}{
	(*ValidateRequestSignatureRequest)(nil), // 0: ValidateRequestSignatureRequest
	(*ValidateRequestSignatureReply)(nil),   // 1: ValidateRequestSignatureReply
	(*DescribeUsersRequest)(nil),            // 2: DescribeUsersRequest
	(*User)(nil),                            // 3: User
	(*DescribeUsersReply)(nil),              // 4: DescribeUsersReply
}
var file_proto_account_proto_depIdxs = []int32{
	3, // 0: DescribeUsersReply.user_set:type_name -> User
	0, // 1: Account.ValidateRequestSignature:input_type -> ValidateRequestSignatureRequest
	2, // 2: Account.DescribeUsers:input_type -> DescribeUsersRequest
	1, // 3: Account.ValidateRequestSignature:output_type -> ValidateRequestSignatureReply
	4, // 4: Account.DescribeUsers:output_type -> DescribeUsersReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_account_proto_init() }
func file_proto_account_proto_init() {
	if File_proto_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequestSignatureRequest); i {
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
		file_proto_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequestSignatureReply); i {
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
		file_proto_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeUsersRequest); i {
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
		file_proto_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeUsersReply); i {
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
			RawDescriptor: file_proto_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_account_proto_goTypes,
		DependencyIndexes: file_proto_account_proto_depIdxs,
		MessageInfos:      file_proto_account_proto_msgTypes,
	}.Build()
	File_proto_account_proto = out.File
	file_proto_account_proto_rawDesc = nil
	file_proto_account_proto_goTypes = nil
	file_proto_account_proto_depIdxs = nil
}