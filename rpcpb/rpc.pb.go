// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: rpcpb/rpc.proto

package rpcpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Config      []byte `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Uri         string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	Bootstrapip string `protobuf:"bytes,4,opt,name=bootstrapip,proto3" json:"bootstrapip,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *NodeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeInfo) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *NodeInfo) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *NodeInfo) GetBootstrapip() string {
	if x != nil {
		return x.Bootstrapip
	}
	return ""
}

type CreateNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *CreateNetworkRequest) Reset() {
	*x = CreateNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNetworkRequest) ProtoMessage() {}

func (x *CreateNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNetworkRequest.ProtoReflect.Descriptor instead.
func (*CreateNetworkRequest) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNetworkRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

type CreateNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateNetworkResponse) Reset() {
	*x = CreateNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNetworkResponse) ProtoMessage() {}

func (x *CreateNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNetworkResponse.ProtoReflect.Descriptor instead.
func (*CreateNetworkResponse) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{4}
}

type GetNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *GetNodesRequest) Reset() {
	*x = GetNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesRequest) ProtoMessage() {}

func (x *GetNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesRequest.ProtoReflect.Descriptor instead.
func (*GetNodesRequest) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *GetNodesRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

type GetNodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*NodeInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *GetNodesResponse) Reset() {
	*x = GetNodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesResponse) ProtoMessage() {}

func (x *GetNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesResponse.ProtoReflect.Descriptor instead.
func (*GetNodesResponse) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *GetNodesResponse) GetNodes() []*NodeInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type GetNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetNodeRequest) Reset() {
	*x = GetNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeRequest.ProtoReflect.Descriptor instead.
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *GetNodeRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *GetNodeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *NodeInfo `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *GetNodeResponse) Reset() {
	*x = GetNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeResponse) ProtoMessage() {}

func (x *GetNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeResponse.ProtoReflect.Descriptor instead.
func (*GetNodeResponse) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *GetNodeResponse) GetNode() *NodeInfo {
	if x != nil {
		return x.Node
	}
	return nil
}

type AddNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Config  []byte `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *AddNodeRequest) Reset() {
	*x = AddNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNodeRequest) ProtoMessage() {}

func (x *AddNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNodeRequest.ProtoReflect.Descriptor instead.
func (*AddNodeRequest) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{9}
}

func (x *AddNodeRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *AddNodeRequest) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

type AddNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *NodeInfo `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *AddNodeResponse) Reset() {
	*x = AddNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNodeResponse) ProtoMessage() {}

func (x *AddNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNodeResponse.ProtoReflect.Descriptor instead.
func (*AddNodeResponse) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{10}
}

func (x *AddNodeResponse) GetNode() *NodeInfo {
	if x != nil {
		return x.Node
	}
	return nil
}

type TeardownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *TeardownRequest) Reset() {
	*x = TeardownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeardownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeardownRequest) ProtoMessage() {}

func (x *TeardownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeardownRequest.ProtoReflect.Descriptor instead.
func (*TeardownRequest) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{11}
}

func (x *TeardownRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

type TeardownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TeardownResponse) Reset() {
	*x = TeardownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeardownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeardownResponse) ProtoMessage() {}

func (x *TeardownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeardownResponse.ProtoReflect.Descriptor instead.
func (*TeardownResponse) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{12}
}

type NodeStopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Timeout int64  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *NodeStopRequest) Reset() {
	*x = NodeStopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStopRequest) ProtoMessage() {}

func (x *NodeStopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStopRequest.ProtoReflect.Descriptor instead.
func (*NodeStopRequest) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{13}
}

func (x *NodeStopRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *NodeStopRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeStopRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type NodeStopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NodeStopResponse) Reset() {
	*x = NodeStopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_rpc_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStopResponse) ProtoMessage() {}

func (x *NodeStopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_rpc_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStopResponse.ProtoReflect.Descriptor instead.
func (*NodeStopResponse) Descriptor() ([]byte, []int) {
	return file_rpcpb_rpc_proto_rawDescGZIP(), []int{14}
}

var File_rpcpb_rpc_proto protoreflect.FileDescriptor

var file_rpcpb_rpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x72, 0x70, 0x63, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x20, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x69, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x69, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x69, 0x70, 0x22, 0x30, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x39, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x42,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x36, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x54, 0x65,
	0x61, 0x72, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x12, 0x0a, 0x10, 0x54, 0x65, 0x61, 0x72, 0x64,
	0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0x0a, 0x0f, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74,
	0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x53, 0x0a, 0x0b, 0x50, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x32,
	0xcb, 0x04, 0x0a, 0x13, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1b, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x70,
	0x63, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x65, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x58, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x15, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x54, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x2e, 0x72, 0x70, 0x63,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61,
	0x64, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x72, 0x64, 0x6f, 0x77,
	0x6e, 0x12, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x61, 0x72, 0x64, 0x6f,
	0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x70,
	0x62, 0x2e, 0x54, 0x65, 0x61, 0x72, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x61, 0x72, 0x64, 0x6f, 0x77, 0x6e,
	0x3a, 0x01, 0x2a, 0x12, 0x58, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x12,
	0x16, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x74, 0x6f, 0x70, 0x3a, 0x01, 0x2a, 0x42, 0x39, 0x5a,
	0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x61, 0x72, 0x6f,
	0x6e, 0x62, 0x75, 0x63, 0x68, 0x77, 0x61, 0x6c, 0x64, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x72, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x3b, 0x72, 0x70, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpcpb_rpc_proto_rawDescOnce sync.Once
	file_rpcpb_rpc_proto_rawDescData = file_rpcpb_rpc_proto_rawDesc
)

func file_rpcpb_rpc_proto_rawDescGZIP() []byte {
	file_rpcpb_rpc_proto_rawDescOnce.Do(func() {
		file_rpcpb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpcpb_rpc_proto_rawDescData)
	})
	return file_rpcpb_rpc_proto_rawDescData
}

var file_rpcpb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_rpcpb_rpc_proto_goTypes = []interface{}{
	(*PingRequest)(nil),           // 0: rpcpb.PingRequest
	(*PingResponse)(nil),          // 1: rpcpb.PingResponse
	(*NodeInfo)(nil),              // 2: rpcpb.NodeInfo
	(*CreateNetworkRequest)(nil),  // 3: rpcpb.CreateNetworkRequest
	(*CreateNetworkResponse)(nil), // 4: rpcpb.CreateNetworkResponse
	(*GetNodesRequest)(nil),       // 5: rpcpb.GetNodesRequest
	(*GetNodesResponse)(nil),      // 6: rpcpb.GetNodesResponse
	(*GetNodeRequest)(nil),        // 7: rpcpb.GetNodeRequest
	(*GetNodeResponse)(nil),       // 8: rpcpb.GetNodeResponse
	(*AddNodeRequest)(nil),        // 9: rpcpb.AddNodeRequest
	(*AddNodeResponse)(nil),       // 10: rpcpb.AddNodeResponse
	(*TeardownRequest)(nil),       // 11: rpcpb.TeardownRequest
	(*TeardownResponse)(nil),      // 12: rpcpb.TeardownResponse
	(*NodeStopRequest)(nil),       // 13: rpcpb.NodeStopRequest
	(*NodeStopResponse)(nil),      // 14: rpcpb.NodeStopResponse
}
var file_rpcpb_rpc_proto_depIdxs = []int32{
	2,  // 0: rpcpb.GetNodesResponse.nodes:type_name -> rpcpb.NodeInfo
	2,  // 1: rpcpb.GetNodeResponse.node:type_name -> rpcpb.NodeInfo
	2,  // 2: rpcpb.AddNodeResponse.node:type_name -> rpcpb.NodeInfo
	0,  // 3: rpcpb.PingService.Ping:input_type -> rpcpb.PingRequest
	3,  // 4: rpcpb.OrchestratorService.CreateNetwork:input_type -> rpcpb.CreateNetworkRequest
	5,  // 5: rpcpb.OrchestratorService.GetNodes:input_type -> rpcpb.GetNodesRequest
	7,  // 6: rpcpb.OrchestratorService.GetNode:input_type -> rpcpb.GetNodeRequest
	9,  // 7: rpcpb.OrchestratorService.AddNode:input_type -> rpcpb.AddNodeRequest
	11, // 8: rpcpb.OrchestratorService.Teardown:input_type -> rpcpb.TeardownRequest
	13, // 9: rpcpb.OrchestratorService.NodeStop:input_type -> rpcpb.NodeStopRequest
	1,  // 10: rpcpb.PingService.Ping:output_type -> rpcpb.PingResponse
	4,  // 11: rpcpb.OrchestratorService.CreateNetwork:output_type -> rpcpb.CreateNetworkResponse
	6,  // 12: rpcpb.OrchestratorService.GetNodes:output_type -> rpcpb.GetNodesResponse
	8,  // 13: rpcpb.OrchestratorService.GetNode:output_type -> rpcpb.GetNodeResponse
	10, // 14: rpcpb.OrchestratorService.AddNode:output_type -> rpcpb.AddNodeResponse
	12, // 15: rpcpb.OrchestratorService.Teardown:output_type -> rpcpb.TeardownResponse
	14, // 16: rpcpb.OrchestratorService.NodeStop:output_type -> rpcpb.NodeStopResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_rpcpb_rpc_proto_init() }
func file_rpcpb_rpc_proto_init() {
	if File_rpcpb_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpcpb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_rpcpb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_rpcpb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
		file_rpcpb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNetworkRequest); i {
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
		file_rpcpb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNetworkResponse); i {
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
		file_rpcpb_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodesRequest); i {
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
		file_rpcpb_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodesResponse); i {
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
		file_rpcpb_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeRequest); i {
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
		file_rpcpb_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeResponse); i {
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
		file_rpcpb_rpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNodeRequest); i {
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
		file_rpcpb_rpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNodeResponse); i {
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
		file_rpcpb_rpc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeardownRequest); i {
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
		file_rpcpb_rpc_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeardownResponse); i {
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
		file_rpcpb_rpc_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStopRequest); i {
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
		file_rpcpb_rpc_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStopResponse); i {
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
			RawDescriptor: file_rpcpb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_rpcpb_rpc_proto_goTypes,
		DependencyIndexes: file_rpcpb_rpc_proto_depIdxs,
		MessageInfos:      file_rpcpb_rpc_proto_msgTypes,
	}.Build()
	File_rpcpb_rpc_proto = out.File
	file_rpcpb_rpc_proto_rawDesc = nil
	file_rpcpb_rpc_proto_goTypes = nil
	file_rpcpb_rpc_proto_depIdxs = nil
}