// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: greetpb/greet.proto

package greetpb

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

// unary request
type GreetUnaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetUnaryRequest) Reset() {
	*x = GreetUnaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetUnaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetUnaryRequest) ProtoMessage() {}

func (x *GreetUnaryRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GreetUnaryRequest.ProtoReflect.Descriptor instead.
func (*GreetUnaryRequest) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{1}
}

func (x *GreetUnaryRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

// server streaming request
type GreetServerStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetServerStreamingRequest) Reset() {
	*x = GreetServerStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetServerStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetServerStreamingRequest) ProtoMessage() {}

func (x *GreetServerStreamingRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GreetServerStreamingRequest.ProtoReflect.Descriptor instead.
func (*GreetServerStreamingRequest) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{2}
}

func (x *GreetServerStreamingRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

// client streaming request
type GreetClientStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetClientStreamingRequest) Reset() {
	*x = GreetClientStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetClientStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetClientStreamingRequest) ProtoMessage() {}

func (x *GreetClientStreamingRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GreetClientStreamingRequest.ProtoReflect.Descriptor instead.
func (*GreetClientStreamingRequest) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{3}
}

func (x *GreetClientStreamingRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

// bidirectional streaming request
type GreetBidirectionalStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetBidirectionalStreamingRequest) Reset() {
	*x = GreetBidirectionalStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetBidirectionalStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetBidirectionalStreamingRequest) ProtoMessage() {}

func (x *GreetBidirectionalStreamingRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GreetBidirectionalStreamingRequest.ProtoReflect.Descriptor instead.
func (*GreetBidirectionalStreamingRequest) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{4}
}

func (x *GreetBidirectionalStreamingRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

// unary response
type GreetUnaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetUnaryResponse) Reset() {
	*x = GreetUnaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetUnaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetUnaryResponse) ProtoMessage() {}

func (x *GreetUnaryResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GreetUnaryResponse.ProtoReflect.Descriptor instead.
func (*GreetUnaryResponse) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{5}
}

func (x *GreetUnaryResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// server streaming response
type GreetServerStreamingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetServerStreamingResponse) Reset() {
	*x = GreetServerStreamingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetServerStreamingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetServerStreamingResponse) ProtoMessage() {}

func (x *GreetServerStreamingResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GreetServerStreamingResponse.ProtoReflect.Descriptor instead.
func (*GreetServerStreamingResponse) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{6}
}

func (x *GreetServerStreamingResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// client streaming response
type GreetClientStreamingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetClientStreamingResponse) Reset() {
	*x = GreetClientStreamingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetClientStreamingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetClientStreamingResponse) ProtoMessage() {}

func (x *GreetClientStreamingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_greet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetClientStreamingResponse.ProtoReflect.Descriptor instead.
func (*GreetClientStreamingResponse) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{7}
}

func (x *GreetClientStreamingResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// bidirectional streaming response
type GreetBidirectionalStreamingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetBidirectionalStreamingResponse) Reset() {
	*x = GreetBidirectionalStreamingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_greet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetBidirectionalStreamingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetBidirectionalStreamingResponse) ProtoMessage() {}

func (x *GreetBidirectionalStreamingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_greet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetBidirectionalStreamingResponse.ProtoReflect.Descriptor instead.
func (*GreetBidirectionalStreamingResponse) Descriptor() ([]byte, []int) {
	return file_greetpb_greet_proto_rawDescGZIP(), []int{8}
}

func (x *GreetBidirectionalStreamingResponse) GetResult() string {
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
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x47, 0x72, 0x65, 0x65, 0x74, 0x55, 0x6e, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x4a, 0x0a, 0x1b, 0x47, 0x72, 0x65, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x4a, 0x0a, 0x1b, 0x47, 0x72, 0x65, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x51,
	0x0a, 0x22, 0x47, 0x72, 0x65, 0x65, 0x74, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x72, 0x65, 0x65, 0x74, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x36, 0x0a, 0x1c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x36, 0x0a, 0x1c, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x3d, 0x0a, 0x23, 0x47, 0x72, 0x65, 0x65, 0x74, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x99,
	0x03, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x43, 0x0a, 0x0a, 0x47, 0x72, 0x65, 0x65, 0x74, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x18, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x55, 0x6e, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x14, 0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x63, 0x0a, 0x14, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x7a,
	0x0a, 0x1b, 0x47, 0x72, 0x65, 0x65, 0x74, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x42, 0x69, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_greetpb_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_greetpb_greet_proto_goTypes = []interface{}{
	(*Greeting)(nil),                            // 0: greet.Greeting
	(*GreetUnaryRequest)(nil),                   // 1: greet.GreetUnaryRequest
	(*GreetServerStreamingRequest)(nil),         // 2: greet.GreetServerStreamingRequest
	(*GreetClientStreamingRequest)(nil),         // 3: greet.GreetClientStreamingRequest
	(*GreetBidirectionalStreamingRequest)(nil),  // 4: greet.GreetBidirectionalStreamingRequest
	(*GreetUnaryResponse)(nil),                  // 5: greet.GreetUnaryResponse
	(*GreetServerStreamingResponse)(nil),        // 6: greet.GreetServerStreamingResponse
	(*GreetClientStreamingResponse)(nil),        // 7: greet.GreetClientStreamingResponse
	(*GreetBidirectionalStreamingResponse)(nil), // 8: greet.GreetBidirectionalStreamingResponse
}
var file_greetpb_greet_proto_depIdxs = []int32{
	0, // 0: greet.GreetUnaryRequest.greeting:type_name -> greet.Greeting
	0, // 1: greet.GreetServerStreamingRequest.greeting:type_name -> greet.Greeting
	0, // 2: greet.GreetClientStreamingRequest.greeting:type_name -> greet.Greeting
	0, // 3: greet.GreetBidirectionalStreamingRequest.greeting:type_name -> greet.Greeting
	1, // 4: greet.GreetService.GreetUnary:input_type -> greet.GreetUnaryRequest
	2, // 5: greet.GreetService.GreetServerStreaming:input_type -> greet.GreetServerStreamingRequest
	3, // 6: greet.GreetService.GreetClientStreaming:input_type -> greet.GreetClientStreamingRequest
	4, // 7: greet.GreetService.GreetBidirectionalStreaming:input_type -> greet.GreetBidirectionalStreamingRequest
	5, // 8: greet.GreetService.GreetUnary:output_type -> greet.GreetUnaryResponse
	6, // 9: greet.GreetService.GreetServerStreaming:output_type -> greet.GreetServerStreamingResponse
	7, // 10: greet.GreetService.GreetClientStreaming:output_type -> greet.GreetClientStreamingResponse
	8, // 11: greet.GreetService.GreetBidirectionalStreaming:output_type -> greet.GreetBidirectionalStreamingResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
			switch v := v.(*GreetUnaryRequest); i {
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
			switch v := v.(*GreetServerStreamingRequest); i {
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
			switch v := v.(*GreetClientStreamingRequest); i {
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
			switch v := v.(*GreetBidirectionalStreamingRequest); i {
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
			switch v := v.(*GreetUnaryResponse); i {
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
			switch v := v.(*GreetServerStreamingResponse); i {
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
		file_greetpb_greet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetClientStreamingResponse); i {
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
		file_greetpb_greet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetBidirectionalStreamingResponse); i {
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
			NumMessages:   9,
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