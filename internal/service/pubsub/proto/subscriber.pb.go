// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: subscriber.proto

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

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	mi := &file_subscriber_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_subscriber_proto_rawDescGZIP(), []int{0}
}

func (x *Identity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	mi := &file_subscriber_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_subscriber_proto_rawDescGZIP(), []int{1}
}

func (x *Subscription) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity     *Identity     `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Subscription *Subscription `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	mi := &file_subscriber_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_subscriber_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequest) GetIdentity() *Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *SubscribeRequest) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

var File_subscriber_proto protoreflect.FileDescriptor

var file_subscriber_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x22, 0x1e, 0x0a, 0x08, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x7a, 0x0a, 0x10,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x38,
	0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x63, 0x61, 0x72, 0x69, 0x6d, 0x2d, 0x6f,
	0x70, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69,
	0x61, 0x2f, 0x61, 0x69, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscriber_proto_rawDescOnce sync.Once
	file_subscriber_proto_rawDescData = file_subscriber_proto_rawDesc
)

func file_subscriber_proto_rawDescGZIP() []byte {
	file_subscriber_proto_rawDescOnce.Do(func() {
		file_subscriber_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscriber_proto_rawDescData)
	})
	return file_subscriber_proto_rawDescData
}

var file_subscriber_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_subscriber_proto_goTypes = []any{
	(*Identity)(nil),         // 0: pubsub.Identity
	(*Subscription)(nil),     // 1: pubsub.Subscription
	(*SubscribeRequest)(nil), // 2: pubsub.SubscribeRequest
}
var file_subscriber_proto_depIdxs = []int32{
	0, // 0: pubsub.SubscribeRequest.identity:type_name -> pubsub.Identity
	1, // 1: pubsub.SubscribeRequest.subscription:type_name -> pubsub.Subscription
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_subscriber_proto_init() }
func file_subscriber_proto_init() {
	if File_subscriber_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_subscriber_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_subscriber_proto_goTypes,
		DependencyIndexes: file_subscriber_proto_depIdxs,
		MessageInfos:      file_subscriber_proto_msgTypes,
	}.Build()
	File_subscriber_proto = out.File
	file_subscriber_proto_rawDesc = nil
	file_subscriber_proto_goTypes = nil
	file_subscriber_proto_depIdxs = nil
}