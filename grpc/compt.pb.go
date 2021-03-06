// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0--rc2
// source: compt.proto

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

// Flow finish request.
type FlowFinishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Flow ID.
	FlowId string `protobuf:"bytes,1,opt,name=flowId,proto3" json:"flowId,omitempty"`
	// Flow version.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Last node of flow.
	Last *FlowNode `protobuf:"bytes,3,opt,name=last,proto3" json:"last,omitempty"`
}

func (x *FlowFinishRequest) Reset() {
	*x = FlowFinishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowFinishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowFinishRequest) ProtoMessage() {}

func (x *FlowFinishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_compt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowFinishRequest.ProtoReflect.Descriptor instead.
func (*FlowFinishRequest) Descriptor() ([]byte, []int) {
	return file_compt_proto_rawDescGZIP(), []int{0}
}

func (x *FlowFinishRequest) GetFlowId() string {
	if x != nil {
		return x.FlowId
	}
	return ""
}

func (x *FlowFinishRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *FlowFinishRequest) GetLast() *FlowNode {
	if x != nil {
		return x.Last
	}
	return nil
}

// Flow finish response.
type FlowFinishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Responsed code from component server.
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Responsed message from component server.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FlowFinishResponse) Reset() {
	*x = FlowFinishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowFinishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowFinishResponse) ProtoMessage() {}

func (x *FlowFinishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_compt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowFinishResponse.ProtoReflect.Descriptor instead.
func (*FlowFinishResponse) Descriptor() ([]byte, []int) {
	return file_compt_proto_rawDescGZIP(), []int{1}
}

func (x *FlowFinishResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FlowFinishResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Flow node.
type FlowNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Flow node ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Flow node name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Flow node description.
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	// Flow node server type, values are 'Golang','Java' or 'NodeJS'.
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// Flow context define.
	Context string `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *FlowNode) Reset() {
	*x = FlowNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowNode) ProtoMessage() {}

func (x *FlowNode) ProtoReflect() protoreflect.Message {
	mi := &file_compt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowNode.ProtoReflect.Descriptor instead.
func (*FlowNode) Descriptor() ([]byte, []int) {
	return file_compt_proto_rawDescGZIP(), []int{2}
}

func (x *FlowNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FlowNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlowNode) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *FlowNode) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FlowNode) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

var File_compt_proto protoreflect.FileDescriptor

var file_compt_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x6f, 0x6d, 0x70, 0x74, 0x22, 0x6a, 0x0a, 0x11, 0x46, 0x6c, 0x6f, 0x77, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c, 0x6f,
	0x77, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6c, 0x6f, 0x77, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x6c,
	0x61, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x74, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74,
	0x22, 0x42, 0x0a, 0x12, 0x46, 0x6c, 0x6f, 0x77, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x70, 0x0a, 0x08, 0x46, 0x6c, 0x6f, 0x77, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x32, 0x4c, 0x0a, 0x05, 0x43, 0x6f, 0x6d, 0x70, 0x74, 0x12,
	0x43, 0x0a, 0x0a, 0x46, 0x6c, 0x6f, 0x77, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x18, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x74, 0x2e,
	0x46, 0x6c, 0x6f, 0x77, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x74, 0x6d, 0x6f,
	0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x50, 0x01, 0x5a, 0x1b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x74, 0x6d, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_compt_proto_rawDescOnce sync.Once
	file_compt_proto_rawDescData = file_compt_proto_rawDesc
)

func file_compt_proto_rawDescGZIP() []byte {
	file_compt_proto_rawDescOnce.Do(func() {
		file_compt_proto_rawDescData = protoimpl.X.CompressGZIP(file_compt_proto_rawDescData)
	})
	return file_compt_proto_rawDescData
}

var file_compt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_compt_proto_goTypes = []interface{}{
	(*FlowFinishRequest)(nil),  // 0: compt.FlowFinishRequest
	(*FlowFinishResponse)(nil), // 1: compt.FlowFinishResponse
	(*FlowNode)(nil),           // 2: compt.FlowNode
}
var file_compt_proto_depIdxs = []int32{
	2, // 0: compt.FlowFinishRequest.last:type_name -> compt.FlowNode
	0, // 1: compt.Compt.FlowFinish:input_type -> compt.FlowFinishRequest
	1, // 2: compt.Compt.FlowFinish:output_type -> compt.FlowFinishResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_compt_proto_init() }
func file_compt_proto_init() {
	if File_compt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_compt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowFinishRequest); i {
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
		file_compt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowFinishResponse); i {
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
		file_compt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowNode); i {
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
			RawDescriptor: file_compt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_compt_proto_goTypes,
		DependencyIndexes: file_compt_proto_depIdxs,
		MessageInfos:      file_compt_proto_msgTypes,
	}.Build()
	File_compt_proto = out.File
	file_compt_proto_rawDesc = nil
	file_compt_proto_goTypes = nil
	file_compt_proto_depIdxs = nil
}
