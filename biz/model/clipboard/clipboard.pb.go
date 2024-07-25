// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: clipboard.proto

package clipboard

import (
	_ "github.com/jhue/misgo/biz/model/api"
	base "github.com/jhue/misgo/biz/model/base"
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

type ClipResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty" form:"Message" query:"Message"`
	Board   map[int64]string `protobuf:"bytes,2,rep,name=Board,proto3" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" json:"Board,omitempty" form:"Board" query:"Board"`
}

func (x *ClipResp) Reset() {
	*x = ClipResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClipResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClipResp) ProtoMessage() {}

func (x *ClipResp) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClipResp.ProtoReflect.Descriptor instead.
func (*ClipResp) Descriptor() ([]byte, []int) {
	return file_clipboard_proto_rawDescGZIP(), []int{0}
}

func (x *ClipResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ClipResp) GetBoard() map[int64]string {
	if x != nil {
		return x.Board
	}
	return nil
}

type ClipReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID     string `protobuf:"bytes,1,opt,name=UID,proto3" form:"uid" json:"uid,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=Content,proto3" form:"content" json:"content,omitempty"`
}

func (x *ClipReq) Reset() {
	*x = ClipReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClipReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClipReq) ProtoMessage() {}

func (x *ClipReq) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClipReq.ProtoReflect.Descriptor instead.
func (*ClipReq) Descriptor() ([]byte, []int) {
	return file_clipboard_proto_rawDescGZIP(), []int{1}
}

func (x *ClipReq) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *ClipReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_clipboard_proto protoreflect.FileDescriptor

var file_clipboard_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x09, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x08, 0x43, 0x6c,
	0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x34, 0x0a, 0x05, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x6c, 0x69, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x38, 0x0a, 0x0a, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x4b, 0x0a, 0x07, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x03, 0x55,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x75, 0x69,
	0x64, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x25, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xbb, 0x18, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x9e, 0x01,
	0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a,
	0x0c, 0x43, 0x6c, 0x69, 0x70, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x11, 0xca, 0xc1, 0x18, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x12, 0x4a, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x70, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x50,
	0x75, 0x74, 0x12, 0x12, 0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43,
	0x6c, 0x69, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x22, 0x11, 0xda, 0xc1, 0x18,
	0x0d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x68, 0x75,
	0x65, 0x2f, 0x6d, 0x69, 0x73, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_clipboard_proto_rawDescOnce sync.Once
	file_clipboard_proto_rawDescData = file_clipboard_proto_rawDesc
)

func file_clipboard_proto_rawDescGZIP() []byte {
	file_clipboard_proto_rawDescOnce.Do(func() {
		file_clipboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_clipboard_proto_rawDescData)
	})
	return file_clipboard_proto_rawDescData
}

var file_clipboard_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_clipboard_proto_goTypes = []interface{}{
	(*ClipResp)(nil),   // 0: clipboard.ClipResp
	(*ClipReq)(nil),    // 1: clipboard.ClipReq
	nil,                // 2: clipboard.ClipResp.BoardEntry
	(*base.Empty)(nil), // 3: base.Empty
}
var file_clipboard_proto_depIdxs = []int32{
	2, // 0: clipboard.ClipResp.Board:type_name -> clipboard.ClipResp.BoardEntry
	3, // 1: clipboard.PingService.ClipBoardGet:input_type -> base.Empty
	1, // 2: clipboard.PingService.ClipBoardPut:input_type -> clipboard.ClipReq
	0, // 3: clipboard.PingService.ClipBoardGet:output_type -> clipboard.ClipResp
	0, // 4: clipboard.PingService.ClipBoardPut:output_type -> clipboard.ClipResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_clipboard_proto_init() }
func file_clipboard_proto_init() {
	if File_clipboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clipboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClipResp); i {
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
		file_clipboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClipReq); i {
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
			RawDescriptor: file_clipboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clipboard_proto_goTypes,
		DependencyIndexes: file_clipboard_proto_depIdxs,
		MessageInfos:      file_clipboard_proto_msgTypes,
	}.Build()
	File_clipboard_proto = out.File
	file_clipboard_proto_rawDesc = nil
	file_clipboard_proto_goTypes = nil
	file_clipboard_proto_depIdxs = nil
}
