// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: go/chatpb/chat.proto

package chatpb

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

// The request message containing requested numbers
type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chatpb_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chatpb_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequest.ProtoReflect.Descriptor instead.
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return file_go_chatpb_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The response message containing response
type ChatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *ChatReply) Reset() {
	*x = ChatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chatpb_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatReply) ProtoMessage() {}

func (x *ChatReply) ProtoReflect() protoreflect.Message {
	mi := &file_go_chatpb_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatReply.ProtoReflect.Descriptor instead.
func (*ChatReply) Descriptor() ([]byte, []int) {
	return file_go_chatpb_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ChatReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_go_chatpb_chat_proto protoreflect.FileDescriptor

var file_go_chatpb_chat_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x27, 0x0a, 0x0b,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x3b, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x33, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chatpb_chat_proto_rawDescOnce sync.Once
	file_go_chatpb_chat_proto_rawDescData = file_go_chatpb_chat_proto_rawDesc
)

func file_go_chatpb_chat_proto_rawDescGZIP() []byte {
	file_go_chatpb_chat_proto_rawDescOnce.Do(func() {
		file_go_chatpb_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chatpb_chat_proto_rawDescData)
	})
	return file_go_chatpb_chat_proto_rawDescData
}

var file_go_chatpb_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chatpb_chat_proto_goTypes = []interface{}{
	(*ChatRequest)(nil), // 0: chat.ChatRequest
	(*ChatReply)(nil),   // 1: chat.ChatReply
}
var file_go_chatpb_chat_proto_depIdxs = []int32{
	0, // 0: chat.Chat.sendRequest:input_type -> chat.ChatRequest
	1, // 1: chat.Chat.sendRequest:output_type -> chat.ChatReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chatpb_chat_proto_init() }
func file_go_chatpb_chat_proto_init() {
	if File_go_chatpb_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chatpb_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequest); i {
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
		file_go_chatpb_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatReply); i {
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
			RawDescriptor: file_go_chatpb_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chatpb_chat_proto_goTypes,
		DependencyIndexes: file_go_chatpb_chat_proto_depIdxs,
		MessageInfos:      file_go_chatpb_chat_proto_msgTypes,
	}.Build()
	File_go_chatpb_chat_proto = out.File
	file_go_chatpb_chat_proto_rawDesc = nil
	file_go_chatpb_chat_proto_goTypes = nil
	file_go_chatpb_chat_proto_depIdxs = nil
}
