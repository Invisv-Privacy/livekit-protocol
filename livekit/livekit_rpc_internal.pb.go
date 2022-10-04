// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: livekit_rpc_internal.proto

package livekit

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

type StartEgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// request metadata
	EgressId  string `protobuf:"bytes,1,opt,name=egress_id,json=egressId,proto3" json:"egress_id,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	SenderId  string `protobuf:"bytes,10,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	SentAt    int64  `protobuf:"varint,4,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	// request
	//
	// Types that are assignable to Request:
	//	*StartEgressRequest_RoomComposite
	//	*StartEgressRequest_TrackComposite
	//	*StartEgressRequest_Track
	Request isStartEgressRequest_Request `protobuf_oneof:"request"`
	// connection info
	RoomId string `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Token  string `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	WsUrl  string `protobuf:"bytes,9,opt,name=ws_url,json=wsUrl,proto3" json:"ws_url,omitempty"`
}

func (x *StartEgressRequest) Reset() {
	*x = StartEgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_rpc_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartEgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartEgressRequest) ProtoMessage() {}

func (x *StartEgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_rpc_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartEgressRequest.ProtoReflect.Descriptor instead.
func (*StartEgressRequest) Descriptor() ([]byte, []int) {
	return file_livekit_rpc_internal_proto_rawDescGZIP(), []int{0}
}

func (x *StartEgressRequest) GetEgressId() string {
	if x != nil {
		return x.EgressId
	}
	return ""
}

func (x *StartEgressRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *StartEgressRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *StartEgressRequest) GetSentAt() int64 {
	if x != nil {
		return x.SentAt
	}
	return 0
}

func (m *StartEgressRequest) GetRequest() isStartEgressRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *StartEgressRequest) GetRoomComposite() *RoomCompositeEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_RoomComposite); ok {
		return x.RoomComposite
	}
	return nil
}

func (x *StartEgressRequest) GetTrackComposite() *TrackCompositeEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_TrackComposite); ok {
		return x.TrackComposite
	}
	return nil
}

func (x *StartEgressRequest) GetTrack() *TrackEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_Track); ok {
		return x.Track
	}
	return nil
}

func (x *StartEgressRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *StartEgressRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *StartEgressRequest) GetWsUrl() string {
	if x != nil {
		return x.WsUrl
	}
	return ""
}

type isStartEgressRequest_Request interface {
	isStartEgressRequest_Request()
}

type StartEgressRequest_RoomComposite struct {
	RoomComposite *RoomCompositeEgressRequest `protobuf:"bytes,5,opt,name=room_composite,json=roomComposite,proto3,oneof"`
}

type StartEgressRequest_TrackComposite struct {
	TrackComposite *TrackCompositeEgressRequest `protobuf:"bytes,6,opt,name=track_composite,json=trackComposite,proto3,oneof"`
}

type StartEgressRequest_Track struct {
	Track *TrackEgressRequest `protobuf:"bytes,7,opt,name=track,proto3,oneof"`
}

func (*StartEgressRequest_RoomComposite) isStartEgressRequest_Request() {}

func (*StartEgressRequest_TrackComposite) isStartEgressRequest_Request() {}

func (*StartEgressRequest_Track) isStartEgressRequest_Request() {}

type EgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// request metadata
	EgressId  string `protobuf:"bytes,1,opt,name=egress_id,json=egressId,proto3" json:"egress_id,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	SenderId  string `protobuf:"bytes,5,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// request
	//
	// Types that are assignable to Request:
	//	*EgressRequest_UpdateStream
	//	*EgressRequest_Stop
	Request isEgressRequest_Request `protobuf_oneof:"request"`
}

func (x *EgressRequest) Reset() {
	*x = EgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_rpc_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressRequest) ProtoMessage() {}

func (x *EgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_rpc_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressRequest.ProtoReflect.Descriptor instead.
func (*EgressRequest) Descriptor() ([]byte, []int) {
	return file_livekit_rpc_internal_proto_rawDescGZIP(), []int{1}
}

func (x *EgressRequest) GetEgressId() string {
	if x != nil {
		return x.EgressId
	}
	return ""
}

func (x *EgressRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *EgressRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (m *EgressRequest) GetRequest() isEgressRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *EgressRequest) GetUpdateStream() *UpdateStreamRequest {
	if x, ok := x.GetRequest().(*EgressRequest_UpdateStream); ok {
		return x.UpdateStream
	}
	return nil
}

func (x *EgressRequest) GetStop() *StopEgressRequest {
	if x, ok := x.GetRequest().(*EgressRequest_Stop); ok {
		return x.Stop
	}
	return nil
}

type isEgressRequest_Request interface {
	isEgressRequest_Request()
}

type EgressRequest_UpdateStream struct {
	UpdateStream *UpdateStreamRequest `protobuf:"bytes,3,opt,name=update_stream,json=updateStream,proto3,oneof"`
}

type EgressRequest_Stop struct {
	Stop *StopEgressRequest `protobuf:"bytes,4,opt,name=stop,proto3,oneof"`
}

func (*EgressRequest_UpdateStream) isEgressRequest_Request() {}

func (*EgressRequest_Stop) isEgressRequest_Request() {}

type EgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *EgressInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Error     string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	RequestId string      `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *EgressResponse) Reset() {
	*x = EgressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_rpc_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressResponse) ProtoMessage() {}

func (x *EgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_rpc_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressResponse.ProtoReflect.Descriptor instead.
func (*EgressResponse) Descriptor() ([]byte, []int) {
	return file_livekit_rpc_internal_proto_rawDescGZIP(), []int{2}
}

func (x *EgressResponse) GetInfo() *EgressInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *EgressResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *EgressResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

// service -> ingress notification of a change in the ingress configuration
type IngressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// request metadata
	IngressId string `protobuf:"bytes,1,opt,name=ingress_id,json=ingressId,proto3" json:"ingress_id,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	SenderId  string `protobuf:"bytes,3,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// Types that are assignable to Request:
	//	*IngressRequest_Update
	//	*IngressRequest_Delete
	Request isIngressRequest_Request `protobuf_oneof:"request"`
}

func (x *IngressRequest) Reset() {
	*x = IngressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_rpc_internal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngressRequest) ProtoMessage() {}

func (x *IngressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_rpc_internal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngressRequest.ProtoReflect.Descriptor instead.
func (*IngressRequest) Descriptor() ([]byte, []int) {
	return file_livekit_rpc_internal_proto_rawDescGZIP(), []int{3}
}

func (x *IngressRequest) GetIngressId() string {
	if x != nil {
		return x.IngressId
	}
	return ""
}

func (x *IngressRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *IngressRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (m *IngressRequest) GetRequest() isIngressRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *IngressRequest) GetUpdate() *UpdateIngressRequest {
	if x, ok := x.GetRequest().(*IngressRequest_Update); ok {
		return x.Update
	}
	return nil
}

func (x *IngressRequest) GetDelete() *DeleteIngressRequest {
	if x, ok := x.GetRequest().(*IngressRequest_Delete); ok {
		return x.Delete
	}
	return nil
}

type isIngressRequest_Request interface {
	isIngressRequest_Request()
}

type IngressRequest_Update struct {
	Update *UpdateIngressRequest `protobuf:"bytes,4,opt,name=update,proto3,oneof"`
}

type IngressRequest_Delete struct {
	Delete *DeleteIngressRequest `protobuf:"bytes,5,opt,name=delete,proto3,oneof"`
}

func (*IngressRequest_Update) isIngressRequest_Request() {}

func (*IngressRequest_Delete) isIngressRequest_Request() {}

// Query an ingress info from an ingress ID or stream key
type GetIngressInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IngressId string `protobuf:"bytes,1,opt,name=ingress_id,json=ingressId,proto3" json:"ingress_id,omitempty"`
	StreamKey string `protobuf:"bytes,2,opt,name=stream_key,json=streamKey,proto3" json:"stream_key,omitempty"`
	RequestId string `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	SenderId  string `protobuf:"bytes,4,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	SentAt    int64  `protobuf:"varint,5,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *GetIngressInfoRequest) Reset() {
	*x = GetIngressInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_rpc_internal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIngressInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIngressInfoRequest) ProtoMessage() {}

func (x *GetIngressInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_rpc_internal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIngressInfoRequest.ProtoReflect.Descriptor instead.
func (*GetIngressInfoRequest) Descriptor() ([]byte, []int) {
	return file_livekit_rpc_internal_proto_rawDescGZIP(), []int{4}
}

func (x *GetIngressInfoRequest) GetIngressId() string {
	if x != nil {
		return x.IngressId
	}
	return ""
}

func (x *GetIngressInfoRequest) GetStreamKey() string {
	if x != nil {
		return x.StreamKey
	}
	return ""
}

func (x *GetIngressInfoRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *GetIngressInfoRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *GetIngressInfoRequest) GetSentAt() int64 {
	if x != nil {
		return x.SentAt
	}
	return 0
}

// Request to store an update to the ingress state ingress -> service
type UpdateIngressStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IngressId string        `protobuf:"bytes,1,opt,name=ingress_id,json=ingressId,proto3" json:"ingress_id,omitempty"`
	State     *IngressState `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *UpdateIngressStateRequest) Reset() {
	*x = UpdateIngressStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_rpc_internal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIngressStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIngressStateRequest) ProtoMessage() {}

func (x *UpdateIngressStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_rpc_internal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIngressStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateIngressStateRequest) Descriptor() ([]byte, []int) {
	return file_livekit_rpc_internal_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateIngressStateRequest) GetIngressId() string {
	if x != nil {
		return x.IngressId
	}
	return ""
}

func (x *UpdateIngressStateRequest) GetState() *IngressState {
	if x != nil {
		return x.State
	}
	return nil
}

type IngressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State     *IngressState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Error     string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	RequestId string        `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *IngressResponse) Reset() {
	*x = IngressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_rpc_internal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngressResponse) ProtoMessage() {}

func (x *IngressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_rpc_internal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngressResponse.ProtoReflect.Descriptor instead.
func (*IngressResponse) Descriptor() ([]byte, []int) {
	return file_livekit_rpc_internal_proto_rawDescGZIP(), []int{6}
}

func (x *IngressResponse) GetState() *IngressState {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *IngressResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *IngressResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type GetIngressInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *IngressInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Token     string       `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	WsUrl     string       `protobuf:"bytes,3,opt,name=ws_url,json=wsUrl,proto3" json:"ws_url,omitempty"`
	Error     string       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	RequestId string       `protobuf:"bytes,5,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *GetIngressInfoResponse) Reset() {
	*x = GetIngressInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_rpc_internal_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIngressInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIngressInfoResponse) ProtoMessage() {}

func (x *GetIngressInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_rpc_internal_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIngressInfoResponse.ProtoReflect.Descriptor instead.
func (*GetIngressInfoResponse) Descriptor() ([]byte, []int) {
	return file_livekit_rpc_internal_proto_rawDescGZIP(), []int{7}
}

func (x *GetIngressInfoResponse) GetInfo() *IngressInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *GetIngressInfoResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetIngressInfoResponse) GetWsUrl() string {
	if x != nil {
		return x.WsUrl
	}
	return ""
}

func (x *GetIngressInfoResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetIngressInfoResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_livekit_rpc_internal_proto protoreflect.FileDescriptor

var file_livekit_rpc_internal_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x1a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x65,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xab, 0x03, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x4c, 0x0a, 0x0e, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x72, 0x6f, 0x6f, 0x6d,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x6b, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15,
	0x0a, 0x06, 0x77, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x77, 0x73, 0x55, 0x72, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xea, 0x01, 0x0a, 0x0d, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0d, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x30, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x45, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x73, 0x74,
	0x6f, 0x70, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6e, 0x0a,
	0x0e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0xe8, 0x01,
	0x0a, 0x0e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x09, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xaa, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4b, 0x65, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x65, 0x6e, 0x74, 0x41, 0x74, 0x22, 0x67, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x73,
	0x0a, 0x0f, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15,
	0x0a, 0x06, 0x77, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x77, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x42, 0x46, 0x5a, 0x23, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0xaa, 0x02, 0x0d, 0x4c, 0x69, 0x76, 0x65, 0x4b, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0xea, 0x02, 0x0e, 0x4c, 0x69, 0x76, 0x65, 0x4b, 0x69, 0x74, 0x3a, 0x3a, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_livekit_rpc_internal_proto_rawDescOnce sync.Once
	file_livekit_rpc_internal_proto_rawDescData = file_livekit_rpc_internal_proto_rawDesc
)

func file_livekit_rpc_internal_proto_rawDescGZIP() []byte {
	file_livekit_rpc_internal_proto_rawDescOnce.Do(func() {
		file_livekit_rpc_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_livekit_rpc_internal_proto_rawDescData)
	})
	return file_livekit_rpc_internal_proto_rawDescData
}

var file_livekit_rpc_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_livekit_rpc_internal_proto_goTypes = []interface{}{
	(*StartEgressRequest)(nil),          // 0: livekit.StartEgressRequest
	(*EgressRequest)(nil),               // 1: livekit.EgressRequest
	(*EgressResponse)(nil),              // 2: livekit.EgressResponse
	(*IngressRequest)(nil),              // 3: livekit.IngressRequest
	(*GetIngressInfoRequest)(nil),       // 4: livekit.GetIngressInfoRequest
	(*UpdateIngressStateRequest)(nil),   // 5: livekit.UpdateIngressStateRequest
	(*IngressResponse)(nil),             // 6: livekit.IngressResponse
	(*GetIngressInfoResponse)(nil),      // 7: livekit.GetIngressInfoResponse
	(*RoomCompositeEgressRequest)(nil),  // 8: livekit.RoomCompositeEgressRequest
	(*TrackCompositeEgressRequest)(nil), // 9: livekit.TrackCompositeEgressRequest
	(*TrackEgressRequest)(nil),          // 10: livekit.TrackEgressRequest
	(*UpdateStreamRequest)(nil),         // 11: livekit.UpdateStreamRequest
	(*StopEgressRequest)(nil),           // 12: livekit.StopEgressRequest
	(*EgressInfo)(nil),                  // 13: livekit.EgressInfo
	(*UpdateIngressRequest)(nil),        // 14: livekit.UpdateIngressRequest
	(*DeleteIngressRequest)(nil),        // 15: livekit.DeleteIngressRequest
	(*IngressState)(nil),                // 16: livekit.IngressState
	(*IngressInfo)(nil),                 // 17: livekit.IngressInfo
}
var file_livekit_rpc_internal_proto_depIdxs = []int32{
	8,  // 0: livekit.StartEgressRequest.room_composite:type_name -> livekit.RoomCompositeEgressRequest
	9,  // 1: livekit.StartEgressRequest.track_composite:type_name -> livekit.TrackCompositeEgressRequest
	10, // 2: livekit.StartEgressRequest.track:type_name -> livekit.TrackEgressRequest
	11, // 3: livekit.EgressRequest.update_stream:type_name -> livekit.UpdateStreamRequest
	12, // 4: livekit.EgressRequest.stop:type_name -> livekit.StopEgressRequest
	13, // 5: livekit.EgressResponse.info:type_name -> livekit.EgressInfo
	14, // 6: livekit.IngressRequest.update:type_name -> livekit.UpdateIngressRequest
	15, // 7: livekit.IngressRequest.delete:type_name -> livekit.DeleteIngressRequest
	16, // 8: livekit.UpdateIngressStateRequest.state:type_name -> livekit.IngressState
	16, // 9: livekit.IngressResponse.state:type_name -> livekit.IngressState
	17, // 10: livekit.GetIngressInfoResponse.info:type_name -> livekit.IngressInfo
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_livekit_rpc_internal_proto_init() }
func file_livekit_rpc_internal_proto_init() {
	if File_livekit_rpc_internal_proto != nil {
		return
	}
	file_livekit_egress_proto_init()
	file_livekit_ingress_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_livekit_rpc_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartEgressRequest); i {
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
		file_livekit_rpc_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressRequest); i {
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
		file_livekit_rpc_internal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressResponse); i {
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
		file_livekit_rpc_internal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngressRequest); i {
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
		file_livekit_rpc_internal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIngressInfoRequest); i {
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
		file_livekit_rpc_internal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIngressStateRequest); i {
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
		file_livekit_rpc_internal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngressResponse); i {
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
		file_livekit_rpc_internal_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIngressInfoResponse); i {
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
	file_livekit_rpc_internal_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StartEgressRequest_RoomComposite)(nil),
		(*StartEgressRequest_TrackComposite)(nil),
		(*StartEgressRequest_Track)(nil),
	}
	file_livekit_rpc_internal_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*EgressRequest_UpdateStream)(nil),
		(*EgressRequest_Stop)(nil),
	}
	file_livekit_rpc_internal_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*IngressRequest_Update)(nil),
		(*IngressRequest_Delete)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_livekit_rpc_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_livekit_rpc_internal_proto_goTypes,
		DependencyIndexes: file_livekit_rpc_internal_proto_depIdxs,
		MessageInfos:      file_livekit_rpc_internal_proto_msgTypes,
	}.Build()
	File_livekit_rpc_internal_proto = out.File
	file_livekit_rpc_internal_proto_rawDesc = nil
	file_livekit_rpc_internal_proto_goTypes = nil
	file_livekit_rpc_internal_proto_depIdxs = nil
}
