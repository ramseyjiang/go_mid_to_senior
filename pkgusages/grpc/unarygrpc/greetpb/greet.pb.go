// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: greetpb/unarygrpc.proto

package greetpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// model
type Greeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *Greeting) Reset() {
	*x = Greeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Greeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greeting) ProtoMessage() {}

func (x *Greeting) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greeting.ProtoReflect.Descriptor instead.
func (*Greeting) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{0}
}

func (x *Greeting) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Greeting) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

// request
type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_greet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{1}
}

func (x *GreetRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type GreetManyTimesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetManyTimesRequest) Reset() {
	*x = GreetManyTimesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetManyTimesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetManyTimesRequest) ProtoMessage() {}

func (x *GreetManyTimesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_greet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetManyTimesRequest.ProtoReflect.Descriptor instead.
func (*GreetManyTimesRequest) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{2}
}

func (x *GreetManyTimesRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type LongGreetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *LongGreetingRequest) Reset() {
	*x = LongGreetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongGreetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongGreetingRequest) ProtoMessage() {}

func (x *LongGreetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_greet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongGreetingRequest.ProtoReflect.Descriptor instead.
func (*LongGreetingRequest) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{3}
}

func (x *LongGreetingRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

// response
type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_greet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{4}
}

func (x *GreetResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type GreetManyTimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetManyTimesResponse) Reset() {
	*x = GreetManyTimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetManyTimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetManyTimesResponse) ProtoMessage() {}

func (x *GreetManyTimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_greet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetManyTimesResponse.ProtoReflect.Descriptor instead.
func (*GreetManyTimesResponse) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{5}
}

func (x *GreetManyTimesResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type LongGreetingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LongGreetingResponse) Reset() {
	*x = LongGreetingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongGreetingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongGreetingResponse) ProtoMessage() {}

func (x *LongGreetingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_greet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongGreetingResponse.ProtoReflect.Descriptor instead.
func (*LongGreetingResponse) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{6}
}

func (x *LongGreetingResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_greetpb_greet_proto protoreflect.FileDescriptor

var file_greetpb_greet_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x22, 0x46, 0x0a, 0x08,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x44, 0x0a, 0x15, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x42, 0x0a, 0x13, 0x4c, 0x6f, 0x6e, 0x67, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x27, 0x0a, 0x0d, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x30, 0x0a, 0x16, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2e, 0x0a, 0x14, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xe4, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x12, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a,
	0x0e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12,
	0x1c, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x4b, 0x0a, 0x0c, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x1a, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x18, 0x5a,
	0x16, 0x67, 0x72, 0x65, 0x65, 0x74, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greetpb_greet_proto_rawDescOnce sync.Once
	file_greetpb_greet_proto_rawDescData = file_greetpb_greet_proto_rawDesc
)

func file_greetpb_greet_proto_rawDescGZIP() []byte {
	file_greetpb_greet_proto_rawDescOnce.Do(func() {
		file_greetpb_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_greetpb_greet_proto_rawDescData)
	})
	return file_greetpb_greet_proto_rawDescData
}

var file_greetpb_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_greetpb_greet_proto_goTypes = []interface{}{
	(*Greeting)(nil),               // 0: unarygrpc.Greeting
	(*GreetRequest)(nil),           // 1: unarygrpc.GreetRequest
	(*GreetManyTimesRequest)(nil),  // 2: unarygrpc.GreetManyTimesRequest
	(*LongGreetingRequest)(nil),    // 3: unarygrpc.LongGreetingRequest
	(*GreetResponse)(nil),          // 4: unarygrpc.GreetResponse
	(*GreetManyTimesResponse)(nil), // 5: unarygrpc.GreetManyTimesResponse
	(*LongGreetingResponse)(nil),   // 6: unarygrpc.LongGreetingResponse
}
var file_greetpb_greet_proto_depIdxs = []int32{
	0, // 0: unarygrpc.GreetRequest.greeting:type_name -> unarygrpc.Greeting
	0, // 1: unarygrpc.GreetManyTimesRequest.greeting:type_name -> unarygrpc.Greeting
	0, // 2: unarygrpc.LongGreetingRequest.greeting:type_name -> unarygrpc.Greeting
	1, // 3: unarygrpc.GreetService.Greet:input_type -> unarygrpc.GreetRequest
	2, // 4: unarygrpc.GreetService.GreetManyTimes:input_type -> unarygrpc.GreetManyTimesRequest
	3, // 5: unarygrpc.GreetService.LongGreeting:input_type -> unarygrpc.LongGreetingRequest
	4, // 6: unarygrpc.GreetService.Greet:output_type -> unarygrpc.GreetResponse
	5, // 7: unarygrpc.GreetService.GreetManyTimes:output_type -> unarygrpc.GreetManyTimesResponse
	6, // 8: unarygrpc.GreetService.LongGreeting:output_type -> unarygrpc.LongGreetingResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_greetpb_greet_proto_init() }
func file_greetpb_greet_proto_init() {
	if File_greetpb_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greetpb_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Greeting); i {
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
		file_greetpb_greet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequest); i {
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
		file_greetpb_greet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetManyTimesRequest); i {
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
		file_greetpb_greet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongGreetingRequest); i {
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
		file_greetpb_greet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponse); i {
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
		file_greetpb_greet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetManyTimesResponse); i {
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
		file_greetpb_greet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongGreetingResponse); i {
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
			RawDescriptor: file_greetpb_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greetpb_greet_proto_goTypes,
		DependencyIndexes: file_greetpb_greet_proto_depIdxs,
		MessageInfos:      file_greetpb_greet_proto_msgTypes,
	}.Build()
	File_greetpb_greet_proto = out.File
	file_greetpb_greet_proto_rawDesc = nil
	file_greetpb_greet_proto_goTypes = nil
	file_greetpb_greet_proto_depIdxs = nil
}