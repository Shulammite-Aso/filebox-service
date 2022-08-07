// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: pkg/proto/filebox.proto

package proto

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

type SendFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	File     []byte `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *SendFileRequest) Reset() {
	*x = SendFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_filebox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendFileRequest) ProtoMessage() {}

func (x *SendFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_filebox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendFileRequest.ProtoReflect.Descriptor instead.
func (*SendFileRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_filebox_proto_rawDescGZIP(), []int{0}
}

func (x *SendFileRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SendFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *SendFileRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type UpdateFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	File     []byte `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *UpdateFileRequest) Reset() {
	*x = UpdateFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_filebox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileRequest) ProtoMessage() {}

func (x *UpdateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_filebox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileRequest.ProtoReflect.Descriptor instead.
func (*UpdateFileRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_filebox_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateFileRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UpdateFileRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_filebox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_filebox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_filebox_proto_rawDescGZIP(), []int{2}
}

func (x *GetFileRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type GetListOfAllFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetListOfAllFilesRequest) Reset() {
	*x = GetListOfAllFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_filebox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListOfAllFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListOfAllFilesRequest) ProtoMessage() {}

func (x *GetListOfAllFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_filebox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListOfAllFilesRequest.ProtoReflect.Descriptor instead.
func (*GetListOfAllFilesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_filebox_proto_rawDescGZIP(), []int{3}
}

func (x *GetListOfAllFilesRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type DeleteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *DeleteFileRequest) Reset() {
	*x = DeleteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_filebox_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileRequest) ProtoMessage() {}

func (x *DeleteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_filebox_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_filebox_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFileRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DeleteFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type SendFileToPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username       string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email          string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	SenderUsername string `protobuf:"bytes,3,opt,name=senderUsername,proto3" json:"senderUsername,omitempty"`
	File           []byte `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *SendFileToPersonRequest) Reset() {
	*x = SendFileToPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_filebox_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendFileToPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendFileToPersonRequest) ProtoMessage() {}

func (x *SendFileToPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_filebox_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendFileToPersonRequest.ProtoReflect.Descriptor instead.
func (*SendFileToPersonRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_filebox_proto_rawDescGZIP(), []int{5}
}

func (x *SendFileToPersonRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SendFileToPersonRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendFileToPersonRequest) GetSenderUsername() string {
	if x != nil {
		return x.SenderUsername
	}
	return ""
}

func (x *SendFileToPersonRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

// RESPONSES
type GetFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileURL string `protobuf:"bytes,1,opt,name=fileURL,proto3" json:"fileURL,omitempty"`
}

func (x *GetFileResponse) Reset() {
	*x = GetFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_filebox_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileResponse) ProtoMessage() {}

func (x *GetFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_filebox_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileResponse.ProtoReflect.Descriptor instead.
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_filebox_proto_rawDescGZIP(), []int{6}
}

func (x *GetFileResponse) GetFileURL() string {
	if x != nil {
		return x.FileURL
	}
	return ""
}

type SuccessMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SuccessMessage) Reset() {
	*x = SuccessMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_filebox_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessMessage) ProtoMessage() {}

func (x *SuccessMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_filebox_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessMessage.ProtoReflect.Descriptor instead.
func (*SuccessMessage) Descriptor() ([]byte, []int) {
	return file_pkg_proto_filebox_proto_rawDescGZIP(), []int{7}
}

func (x *SuccessMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetListOfAllFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllFiles []string `protobuf:"bytes,1,rep,name=allFiles,proto3" json:"allFiles,omitempty"`
}

func (x *GetListOfAllFilesResponse) Reset() {
	*x = GetListOfAllFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_filebox_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListOfAllFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListOfAllFilesResponse) ProtoMessage() {}

func (x *GetListOfAllFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_filebox_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListOfAllFilesResponse.ProtoReflect.Descriptor instead.
func (*GetListOfAllFilesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_filebox_proto_rawDescGZIP(), []int{8}
}

func (x *GetListOfAllFilesResponse) GetAllFiles() []string {
	if x != nil {
		return x.AllFiles
	}
	return nil
}

var File_pkg_proto_filebox_proto protoreflect.FileDescriptor

var file_pkg_proto_filebox_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x62,
	0x6f, 0x78, 0x22, 0x5d, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0x5f, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x48, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x87, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x6f,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x26, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x37, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x66, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x32, 0xca, 0x03,
	0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3f, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78,
	0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x17, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x62, 0x6f, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x66, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x41,
	0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x66, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x10, 0x53, 0x65, 0x6e,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x6f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x6f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_proto_filebox_proto_rawDescOnce sync.Once
	file_pkg_proto_filebox_proto_rawDescData = file_pkg_proto_filebox_proto_rawDesc
)

func file_pkg_proto_filebox_proto_rawDescGZIP() []byte {
	file_pkg_proto_filebox_proto_rawDescOnce.Do(func() {
		file_pkg_proto_filebox_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_filebox_proto_rawDescData)
	})
	return file_pkg_proto_filebox_proto_rawDescData
}

var file_pkg_proto_filebox_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_proto_filebox_proto_goTypes = []interface{}{
	(*SendFileRequest)(nil),           // 0: filebox.SendFileRequest
	(*UpdateFileRequest)(nil),         // 1: filebox.UpdateFileRequest
	(*GetFileRequest)(nil),            // 2: filebox.GetFileRequest
	(*GetListOfAllFilesRequest)(nil),  // 3: filebox.GetListOfAllFilesRequest
	(*DeleteFileRequest)(nil),         // 4: filebox.DeleteFileRequest
	(*SendFileToPersonRequest)(nil),   // 5: filebox.SendFileToPersonRequest
	(*GetFileResponse)(nil),           // 6: filebox.GetFileResponse
	(*SuccessMessage)(nil),            // 7: filebox.SuccessMessage
	(*GetListOfAllFilesResponse)(nil), // 8: filebox.GetListOfAllFilesResponse
}
var file_pkg_proto_filebox_proto_depIdxs = []int32{
	0, // 0: filebox.FileboxService.SendFile:input_type -> filebox.SendFileRequest
	1, // 1: filebox.FileboxService.UpdateFile:input_type -> filebox.UpdateFileRequest
	2, // 2: filebox.FileboxService.GetFile:input_type -> filebox.GetFileRequest
	3, // 3: filebox.FileboxService.GetListOfAllFiles:input_type -> filebox.GetListOfAllFilesRequest
	4, // 4: filebox.FileboxService.DeleteFile:input_type -> filebox.DeleteFileRequest
	5, // 5: filebox.FileboxService.SendFileToPerson:input_type -> filebox.SendFileToPersonRequest
	7, // 6: filebox.FileboxService.SendFile:output_type -> filebox.SuccessMessage
	7, // 7: filebox.FileboxService.UpdateFile:output_type -> filebox.SuccessMessage
	6, // 8: filebox.FileboxService.GetFile:output_type -> filebox.GetFileResponse
	8, // 9: filebox.FileboxService.GetListOfAllFiles:output_type -> filebox.GetListOfAllFilesResponse
	7, // 10: filebox.FileboxService.DeleteFile:output_type -> filebox.SuccessMessage
	7, // 11: filebox.FileboxService.SendFileToPerson:output_type -> filebox.SuccessMessage
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_filebox_proto_init() }
func file_pkg_proto_filebox_proto_init() {
	if File_pkg_proto_filebox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_filebox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendFileRequest); i {
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
		file_pkg_proto_filebox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileRequest); i {
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
		file_pkg_proto_filebox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
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
		file_pkg_proto_filebox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListOfAllFilesRequest); i {
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
		file_pkg_proto_filebox_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileRequest); i {
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
		file_pkg_proto_filebox_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendFileToPersonRequest); i {
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
		file_pkg_proto_filebox_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileResponse); i {
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
		file_pkg_proto_filebox_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessMessage); i {
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
		file_pkg_proto_filebox_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListOfAllFilesResponse); i {
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
			RawDescriptor: file_pkg_proto_filebox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_filebox_proto_goTypes,
		DependencyIndexes: file_pkg_proto_filebox_proto_depIdxs,
		MessageInfos:      file_pkg_proto_filebox_proto_msgTypes,
	}.Build()
	File_pkg_proto_filebox_proto = out.File
	file_pkg_proto_filebox_proto_rawDesc = nil
	file_pkg_proto_filebox_proto_goTypes = nil
	file_pkg_proto_filebox_proto_depIdxs = nil
}
