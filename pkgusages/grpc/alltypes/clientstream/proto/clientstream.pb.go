// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto/clientstream.proto

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

// model
type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname   string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname    string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	PhoneNumber uint64 `protobuf:"varint,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clientstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clientstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_proto_clientstream_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Contact) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Contact) GetPhoneNumber() uint64 {
	if x != nil {
		return x.PhoneNumber
	}
	return 0
}

type NumCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *NumCheckRequest) Reset() {
	*x = NumCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clientstream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumCheckRequest) ProtoMessage() {}

func (x *NumCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clientstream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumCheckRequest.ProtoReflect.Descriptor instead.
func (*NumCheckRequest) Descriptor() ([]byte, []int) {
	return file_proto_clientstream_proto_rawDescGZIP(), []int{1}
}

func (x *NumCheckRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type NumCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckResult []*Result `protobuf:"bytes,1,rep,name=check_result,json=checkResult,proto3" json:"check_result,omitempty"`
}

func (x *NumCheckResponse) Reset() {
	*x = NumCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clientstream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumCheckResponse) ProtoMessage() {}

func (x *NumCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clientstream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumCheckResponse.ProtoReflect.Descriptor instead.
func (*NumCheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_clientstream_proto_rawDescGZIP(), []int{2}
}

func (x *NumCheckResponse) GetCheckResult() []*Result {
	if x != nil {
		return x.CheckResult
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clientstream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clientstream_proto_msgTypes[3]
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
	return file_proto_clientstream_proto_rawDescGZIP(), []int{3}
}

func (x *Result) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Result) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_clientstream_proto protoreflect.FileDescriptor

var file_proto_clientstream_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x66,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0f, 0x4e, 0x75, 0x6d, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x41, 0x0a, 0x10, 0x4e, 0x75, 0x6d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x32, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x42, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x39, 0x0a, 0x08, 0x4e, 0x75, 0x6d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x13, 0x2e,
	0x70, 0x62, 0x2e, 0x4e, 0x75, 0x6d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x75, 0x6d, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x1d, 0x5a, 0x1b,
	0x61, 0x6c, 0x6c, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_clientstream_proto_rawDescOnce sync.Once
	file_proto_clientstream_proto_rawDescData = file_proto_clientstream_proto_rawDesc
)

func file_proto_clientstream_proto_rawDescGZIP() []byte {
	file_proto_clientstream_proto_rawDescOnce.Do(func() {
		file_proto_clientstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_clientstream_proto_rawDescData)
	})
	return file_proto_clientstream_proto_rawDescData
}

var file_proto_clientstream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_clientstream_proto_goTypes = []interface{}{
	(*Contact)(nil),          // 0: pb.Contact
	(*NumCheckRequest)(nil),  // 1: pb.NumCheckRequest
	(*NumCheckResponse)(nil), // 2: pb.NumCheckResponse
	(*Result)(nil),           // 3: pb.Result
}
var file_proto_clientstream_proto_depIdxs = []int32{
	3, // 0: pb.NumCheckResponse.check_result:type_name -> pb.Result
	1, // 1: pb.Phone.NumCheck:input_type -> pb.NumCheckRequest
	2, // 2: pb.Phone.NumCheck:output_type -> pb.NumCheckResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_clientstream_proto_init() }
func file_proto_clientstream_proto_init() {
	if File_proto_clientstream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_clientstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_proto_clientstream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumCheckRequest); i {
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
		file_proto_clientstream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumCheckResponse); i {
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
		file_proto_clientstream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_clientstream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_clientstream_proto_goTypes,
		DependencyIndexes: file_proto_clientstream_proto_depIdxs,
		MessageInfos:      file_proto_clientstream_proto_msgTypes,
	}.Build()
	File_proto_clientstream_proto = out.File
	file_proto_clientstream_proto_rawDesc = nil
	file_proto_clientstream_proto_goTypes = nil
	file_proto_clientstream_proto_depIdxs = nil
}
