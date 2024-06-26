// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: ClientServerComms.proto

package grpc

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

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username       string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	PublicKeyLogin []byte `protobuf:"bytes,2,opt,name=publicKeyLogin,proto3" json:"publicKeyLogin,omitempty"`
	PublicKeyChat  []byte `protobuf:"bytes,3,opt,name=publicKeyChat,proto3" json:"publicKeyChat,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientServerComms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ClientServerComms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_ClientServerComms_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignUpRequest) GetPublicKeyLogin() []byte {
	if x != nil {
		return x.PublicKeyLogin
	}
	return nil
}

func (x *SignUpRequest) GetPublicKeyChat() []byte {
	if x != nil {
		return x.PublicKeyChat
	}
	return nil
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challenge []byte `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientServerComms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ClientServerComms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_ClientServerComms_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpResponse) GetChallenge() []byte {
	if x != nil {
		return x.Challenge
	}
	return nil
}

type SignUpChallengeResponseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeResponse []byte `protobuf:"bytes,1,opt,name=challenge_response,json=challengeResponse,proto3" json:"challenge_response,omitempty"`
}

func (x *SignUpChallengeResponseRequest) Reset() {
	*x = SignUpChallengeResponseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientServerComms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpChallengeResponseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpChallengeResponseRequest) ProtoMessage() {}

func (x *SignUpChallengeResponseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ClientServerComms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpChallengeResponseRequest.ProtoReflect.Descriptor instead.
func (*SignUpChallengeResponseRequest) Descriptor() ([]byte, []int) {
	return file_ClientServerComms_proto_rawDescGZIP(), []int{2}
}

func (x *SignUpChallengeResponseRequest) GetChallengeResponse() []byte {
	if x != nil {
		return x.ChallengeResponse
	}
	return nil
}

type SignUpChallengeResponseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignUpChallengeResponseResponse) Reset() {
	*x = SignUpChallengeResponseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientServerComms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpChallengeResponseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpChallengeResponseResponse) ProtoMessage() {}

func (x *SignUpChallengeResponseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ClientServerComms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpChallengeResponseResponse.ProtoReflect.Descriptor instead.
func (*SignUpChallengeResponseResponse) Descriptor() ([]byte, []int) {
	return file_ClientServerComms_proto_rawDescGZIP(), []int{3}
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username         string            `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Address          uint32            `protobuf:"fixed32,2,opt,name=address,proto3" json:"address,omitempty"`
	DigitalSignature *DigitalSignature `protobuf:"bytes,3,opt,name=digitalSignature,proto3" json:"digitalSignature,omitempty"`
	ChallengeNonce   uint64            `protobuf:"fixed64,4,opt,name=challenge_nonce,json=challengeNonce,proto3" json:"challenge_nonce,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientServerComms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ClientServerComms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_ClientServerComms_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

func (x *LoginRequest) GetDigitalSignature() *DigitalSignature {
	if x != nil {
		return x.DigitalSignature
	}
	return nil
}

func (x *LoginRequest) GetChallengeNonce() uint64 {
	if x != nil {
		return x.ChallengeNonce
	}
	return 0
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientServerComms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ClientServerComms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_ClientServerComms_proto_rawDescGZIP(), []int{5}
}

type FindUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username         string            `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	DigitalSignature *DigitalSignature `protobuf:"bytes,2,opt,name=digitalSignature,proto3" json:"digitalSignature,omitempty"`
}

func (x *FindUserRequest) Reset() {
	*x = FindUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientServerComms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserRequest) ProtoMessage() {}

func (x *FindUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ClientServerComms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserRequest.ProtoReflect.Descriptor instead.
func (*FindUserRequest) Descriptor() ([]byte, []int) {
	return file_ClientServerComms_proto_rawDescGZIP(), []int{6}
}

func (x *FindUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FindUserRequest) GetDigitalSignature() *DigitalSignature {
	if x != nil {
		return x.DigitalSignature
	}
	return nil
}

type FindUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username      string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Address       uint32 `protobuf:"fixed32,2,opt,name=address,proto3" json:"address,omitempty"`
	PublicKeyChat []byte `protobuf:"bytes,3,opt,name=publicKeyChat,proto3" json:"publicKeyChat,omitempty"`
}

func (x *FindUserResponse) Reset() {
	*x = FindUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientServerComms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserResponse) ProtoMessage() {}

func (x *FindUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ClientServerComms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserResponse.ProtoReflect.Descriptor instead.
func (*FindUserResponse) Descriptor() ([]byte, []int) {
	return file_ClientServerComms_proto_rawDescGZIP(), []int{7}
}

func (x *FindUserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FindUserResponse) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

func (x *FindUserResponse) GetPublicKeyChat() []byte {
	if x != nil {
		return x.PublicKeyChat
	}
	return nil
}

type DigitalSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *DigitalSignature) Reset() {
	*x = DigitalSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientServerComms_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DigitalSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DigitalSignature) ProtoMessage() {}

func (x *DigitalSignature) ProtoReflect() protoreflect.Message {
	mi := &file_ClientServerComms_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DigitalSignature.ProtoReflect.Descriptor instead.
func (*DigitalSignature) Descriptor() ([]byte, []int) {
	return file_ClientServerComms_proto_rawDescGZIP(), []int{8}
}

func (x *DigitalSignature) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DigitalSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_ClientServerComms_proto protoreflect.FileDescriptor

var file_ClientServerComms_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22,
	0x79, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x43, 0x68, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x43, 0x68, 0x61, 0x74, 0x22, 0x2e, 0x0a, 0x0e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x22, 0x4f, 0x0a, 0x1e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x0a, 0x1f, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb1,
	0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x42, 0x0a, 0x10, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x10, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x6e,
	0x63, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x71, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x10, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x6e, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x43, 0x68, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x43, 0x68, 0x61, 0x74, 0x22, 0x4c, 0x0a, 0x10, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x32, 0x9d, 0x02, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x66, 0x0a, 0x17, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x46, 0x69, 0x6e,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x63, 0x68,
	0x61, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClientServerComms_proto_rawDescOnce sync.Once
	file_ClientServerComms_proto_rawDescData = file_ClientServerComms_proto_rawDesc
)

func file_ClientServerComms_proto_rawDescGZIP() []byte {
	file_ClientServerComms_proto_rawDescOnce.Do(func() {
		file_ClientServerComms_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClientServerComms_proto_rawDescData)
	})
	return file_ClientServerComms_proto_rawDescData
}

var file_ClientServerComms_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ClientServerComms_proto_goTypes = []interface{}{
	(*SignUpRequest)(nil),                   // 0: main.SignUpRequest
	(*SignUpResponse)(nil),                  // 1: main.SignUpResponse
	(*SignUpChallengeResponseRequest)(nil),  // 2: main.SignUpChallengeResponseRequest
	(*SignUpChallengeResponseResponse)(nil), // 3: main.SignUpChallengeResponseResponse
	(*LoginRequest)(nil),                    // 4: main.LoginRequest
	(*LoginResponse)(nil),                   // 5: main.LoginResponse
	(*FindUserRequest)(nil),                 // 6: main.FindUserRequest
	(*FindUserResponse)(nil),                // 7: main.FindUserResponse
	(*DigitalSignature)(nil),                // 8: main.DigitalSignature
}
var file_ClientServerComms_proto_depIdxs = []int32{
	8, // 0: main.LoginRequest.digitalSignature:type_name -> main.DigitalSignature
	8, // 1: main.FindUserRequest.digitalSignature:type_name -> main.DigitalSignature
	0, // 2: main.ClientServerComms.SignUp:input_type -> main.SignUpRequest
	2, // 3: main.ClientServerComms.SignUpChallengeResponse:input_type -> main.SignUpChallengeResponseRequest
	4, // 4: main.ClientServerComms.Login:input_type -> main.LoginRequest
	6, // 5: main.ClientServerComms.FindUser:input_type -> main.FindUserRequest
	1, // 6: main.ClientServerComms.SignUp:output_type -> main.SignUpResponse
	3, // 7: main.ClientServerComms.SignUpChallengeResponse:output_type -> main.SignUpChallengeResponseResponse
	5, // 8: main.ClientServerComms.Login:output_type -> main.LoginResponse
	7, // 9: main.ClientServerComms.FindUser:output_type -> main.FindUserResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ClientServerComms_proto_init() }
func file_ClientServerComms_proto_init() {
	if File_ClientServerComms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ClientServerComms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
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
		file_ClientServerComms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
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
		file_ClientServerComms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpChallengeResponseRequest); i {
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
		file_ClientServerComms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpChallengeResponseResponse); i {
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
		file_ClientServerComms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_ClientServerComms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_ClientServerComms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserRequest); i {
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
		file_ClientServerComms_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserResponse); i {
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
		file_ClientServerComms_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DigitalSignature); i {
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
			RawDescriptor: file_ClientServerComms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ClientServerComms_proto_goTypes,
		DependencyIndexes: file_ClientServerComms_proto_depIdxs,
		MessageInfos:      file_ClientServerComms_proto_msgTypes,
	}.Build()
	File_ClientServerComms_proto = out.File
	file_ClientServerComms_proto_rawDesc = nil
	file_ClientServerComms_proto_goTypes = nil
	file_ClientServerComms_proto_depIdxs = nil
}
