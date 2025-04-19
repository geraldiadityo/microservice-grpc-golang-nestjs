// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: pengguna/pengguna.proto

package pengguna

import (
	role "api_gateway/proto/role"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatePenggunaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	RoleId        int32                  `protobuf:"varint,4,opt,name=roleId,proto3" json:"roleId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePenggunaRequest) Reset() {
	*x = CreatePenggunaRequest{}
	mi := &file_pengguna_pengguna_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePenggunaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePenggunaRequest) ProtoMessage() {}

func (x *CreatePenggunaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pengguna_pengguna_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePenggunaRequest.ProtoReflect.Descriptor instead.
func (*CreatePenggunaRequest) Descriptor() ([]byte, []int) {
	return file_pengguna_pengguna_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePenggunaRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreatePenggunaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePenggunaRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreatePenggunaRequest) GetRoleId() int32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type GetPenggunasRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPenggunasRequest) Reset() {
	*x = GetPenggunasRequest{}
	mi := &file_pengguna_pengguna_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPenggunasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPenggunasRequest) ProtoMessage() {}

func (x *GetPenggunasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pengguna_pengguna_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPenggunasRequest.ProtoReflect.Descriptor instead.
func (*GetPenggunasRequest) Descriptor() ([]byte, []int) {
	return file_pengguna_pengguna_proto_rawDescGZIP(), []int{1}
}

type GetPenggunaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPenggunaRequest) Reset() {
	*x = GetPenggunaRequest{}
	mi := &file_pengguna_pengguna_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPenggunaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPenggunaRequest) ProtoMessage() {}

func (x *GetPenggunaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pengguna_pengguna_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPenggunaRequest.ProtoReflect.Descriptor instead.
func (*GetPenggunaRequest) Descriptor() ([]byte, []int) {
	return file_pengguna_pengguna_proto_rawDescGZIP(), []int{2}
}

func (x *GetPenggunaRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdatePenggunaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	RoleId        int32                  `protobuf:"varint,5,opt,name=roleId,proto3" json:"roleId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePenggunaRequest) Reset() {
	*x = UpdatePenggunaRequest{}
	mi := &file_pengguna_pengguna_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePenggunaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePenggunaRequest) ProtoMessage() {}

func (x *UpdatePenggunaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pengguna_pengguna_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePenggunaRequest.ProtoReflect.Descriptor instead.
func (*UpdatePenggunaRequest) Descriptor() ([]byte, []int) {
	return file_pengguna_pengguna_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePenggunaRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePenggunaRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdatePenggunaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePenggunaRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdatePenggunaRequest) GetRoleId() int32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type DeletePenggunaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePenggunaRequest) Reset() {
	*x = DeletePenggunaRequest{}
	mi := &file_pengguna_pengguna_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePenggunaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePenggunaRequest) ProtoMessage() {}

func (x *DeletePenggunaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pengguna_pengguna_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePenggunaRequest.ProtoReflect.Descriptor instead.
func (*DeletePenggunaRequest) Descriptor() ([]byte, []int) {
	return file_pengguna_pengguna_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePenggunaRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PenggunaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Role          *role.RoleResponse     `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PenggunaResponse) Reset() {
	*x = PenggunaResponse{}
	mi := &file_pengguna_pengguna_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PenggunaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PenggunaResponse) ProtoMessage() {}

func (x *PenggunaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pengguna_pengguna_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PenggunaResponse.ProtoReflect.Descriptor instead.
func (*PenggunaResponse) Descriptor() ([]byte, []int) {
	return file_pengguna_pengguna_proto_rawDescGZIP(), []int{5}
}

func (x *PenggunaResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PenggunaResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PenggunaResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PenggunaResponse) GetRole() *role.RoleResponse {
	if x != nil {
		return x.Role
	}
	return nil
}

type GetPenggunasResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Penggunas     []*PenggunaResponse    `protobuf:"bytes,1,rep,name=penggunas,proto3" json:"penggunas,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPenggunasResponse) Reset() {
	*x = GetPenggunasResponse{}
	mi := &file_pengguna_pengguna_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPenggunasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPenggunasResponse) ProtoMessage() {}

func (x *GetPenggunasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pengguna_pengguna_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPenggunasResponse.ProtoReflect.Descriptor instead.
func (*GetPenggunasResponse) Descriptor() ([]byte, []int) {
	return file_pengguna_pengguna_proto_rawDescGZIP(), []int{6}
}

func (x *GetPenggunasResponse) GetPenggunas() []*PenggunaResponse {
	if x != nil {
		return x.Penggunas
	}
	return nil
}

type DeletePenggunaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePenggunaResponse) Reset() {
	*x = DeletePenggunaResponse{}
	mi := &file_pengguna_pengguna_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePenggunaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePenggunaResponse) ProtoMessage() {}

func (x *DeletePenggunaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pengguna_pengguna_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePenggunaResponse.ProtoReflect.Descriptor instead.
func (*DeletePenggunaResponse) Descriptor() ([]byte, []int) {
	return file_pengguna_pengguna_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePenggunaResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_pengguna_pengguna_proto protoreflect.FileDescriptor

const file_pengguna_pengguna_proto_rawDesc = "" +
	"\n" +
	"\x17pengguna/pengguna.proto\x12\bpengguna\x1a\x0frole/role.proto\"{\n" +
	"\x15CreatePenggunaRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12\x16\n" +
	"\x06roleId\x18\x04 \x01(\x05R\x06roleId\"\x15\n" +
	"\x13GetPenggunasRequest\"$\n" +
	"\x12GetPenggunaRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"\x8b\x01\n" +
	"\x15UpdatePenggunaRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x1a\n" +
	"\bpassword\x18\x04 \x01(\tR\bpassword\x12\x16\n" +
	"\x06roleId\x18\x05 \x01(\x05R\x06roleId\"'\n" +
	"\x15DeletePenggunaRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"z\n" +
	"\x10PenggunaResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12&\n" +
	"\x04role\x18\x04 \x01(\v2\x12.role.RoleResponseR\x04role\"P\n" +
	"\x14GetPenggunasResponse\x128\n" +
	"\tpenggunas\x18\x01 \x03(\v2\x1a.pengguna.PenggunaResponseR\tpenggunas\"2\n" +
	"\x16DeletePenggunaResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\x9c\x03\n" +
	"\x0fPenggunaService\x12M\n" +
	"\x0eCreatePengguna\x12\x1f.pengguna.CreatePenggunaRequest\x1a\x1a.pengguna.PenggunaResponse\x12M\n" +
	"\fGetPenggunas\x12\x1d.pengguna.GetPenggunasRequest\x1a\x1e.pengguna.GetPenggunasResponse\x12G\n" +
	"\vGetPengguna\x12\x1c.pengguna.GetPenggunaRequest\x1a\x1a.pengguna.PenggunaResponse\x12M\n" +
	"\x0eUpdatePengguna\x12\x1f.pengguna.UpdatePenggunaRequest\x1a\x1a.pengguna.PenggunaResponse\x12S\n" +
	"\x0eDeletePengguna\x12\x1f.pengguna.DeletePenggunaRequest\x1a .pengguna.DeletePenggunaResponseB\x16Z\x14api_gateway/penggunab\x06proto3"

var (
	file_pengguna_pengguna_proto_rawDescOnce sync.Once
	file_pengguna_pengguna_proto_rawDescData []byte
)

func file_pengguna_pengguna_proto_rawDescGZIP() []byte {
	file_pengguna_pengguna_proto_rawDescOnce.Do(func() {
		file_pengguna_pengguna_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pengguna_pengguna_proto_rawDesc), len(file_pengguna_pengguna_proto_rawDesc)))
	})
	return file_pengguna_pengguna_proto_rawDescData
}

var file_pengguna_pengguna_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pengguna_pengguna_proto_goTypes = []any{
	(*CreatePenggunaRequest)(nil),  // 0: pengguna.CreatePenggunaRequest
	(*GetPenggunasRequest)(nil),    // 1: pengguna.GetPenggunasRequest
	(*GetPenggunaRequest)(nil),     // 2: pengguna.GetPenggunaRequest
	(*UpdatePenggunaRequest)(nil),  // 3: pengguna.UpdatePenggunaRequest
	(*DeletePenggunaRequest)(nil),  // 4: pengguna.DeletePenggunaRequest
	(*PenggunaResponse)(nil),       // 5: pengguna.PenggunaResponse
	(*GetPenggunasResponse)(nil),   // 6: pengguna.GetPenggunasResponse
	(*DeletePenggunaResponse)(nil), // 7: pengguna.DeletePenggunaResponse
	(*role.RoleResponse)(nil),      // 8: role.RoleResponse
}
var file_pengguna_pengguna_proto_depIdxs = []int32{
	8, // 0: pengguna.PenggunaResponse.role:type_name -> role.RoleResponse
	5, // 1: pengguna.GetPenggunasResponse.penggunas:type_name -> pengguna.PenggunaResponse
	0, // 2: pengguna.PenggunaService.CreatePengguna:input_type -> pengguna.CreatePenggunaRequest
	1, // 3: pengguna.PenggunaService.GetPenggunas:input_type -> pengguna.GetPenggunasRequest
	2, // 4: pengguna.PenggunaService.GetPengguna:input_type -> pengguna.GetPenggunaRequest
	3, // 5: pengguna.PenggunaService.UpdatePengguna:input_type -> pengguna.UpdatePenggunaRequest
	4, // 6: pengguna.PenggunaService.DeletePengguna:input_type -> pengguna.DeletePenggunaRequest
	5, // 7: pengguna.PenggunaService.CreatePengguna:output_type -> pengguna.PenggunaResponse
	6, // 8: pengguna.PenggunaService.GetPenggunas:output_type -> pengguna.GetPenggunasResponse
	5, // 9: pengguna.PenggunaService.GetPengguna:output_type -> pengguna.PenggunaResponse
	5, // 10: pengguna.PenggunaService.UpdatePengguna:output_type -> pengguna.PenggunaResponse
	7, // 11: pengguna.PenggunaService.DeletePengguna:output_type -> pengguna.DeletePenggunaResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pengguna_pengguna_proto_init() }
func file_pengguna_pengguna_proto_init() {
	if File_pengguna_pengguna_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pengguna_pengguna_proto_rawDesc), len(file_pengguna_pengguna_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pengguna_pengguna_proto_goTypes,
		DependencyIndexes: file_pengguna_pengguna_proto_depIdxs,
		MessageInfos:      file_pengguna_pengguna_proto_msgTypes,
	}.Build()
	File_pengguna_pengguna_proto = out.File
	file_pengguna_pengguna_proto_goTypes = nil
	file_pengguna_pengguna_proto_depIdxs = nil
}
