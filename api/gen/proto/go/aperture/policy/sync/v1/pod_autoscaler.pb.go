// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aperture/policy/sync/v1/pod_autoscaler.proto

package syncv1

import (
	v1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/language/v1"
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

type HorizontalPodScalerWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CommonAttributes
	CommonAttributes    *CommonAttributes       `protobuf:"bytes,1,opt,name=common_attributes,json=commonAttributes,proto3" json:"common_attributes,omitempty"`
	HorizontalPodScaler *v1.HorizontalPodScaler `protobuf:"bytes,2,opt,name=horizontal_pod_scaler,json=horizontalPodScaler,proto3" json:"horizontal_pod_scaler,omitempty"`
}

func (x *HorizontalPodScalerWrapper) Reset() {
	*x = HorizontalPodScalerWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HorizontalPodScalerWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HorizontalPodScalerWrapper) ProtoMessage() {}

func (x *HorizontalPodScalerWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HorizontalPodScalerWrapper.ProtoReflect.Descriptor instead.
func (*HorizontalPodScalerWrapper) Descriptor() ([]byte, []int) {
	return file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescGZIP(), []int{0}
}

func (x *HorizontalPodScalerWrapper) GetCommonAttributes() *CommonAttributes {
	if x != nil {
		return x.CommonAttributes
	}
	return nil
}

func (x *HorizontalPodScalerWrapper) GetHorizontalPodScaler() *v1.HorizontalPodScaler {
	if x != nil {
		return x.HorizontalPodScaler
	}
	return nil
}

type HorizontalPodScalerDynamicConfigWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CommonAttributes
	CommonAttributes *CommonAttributes `protobuf:"bytes,1,opt,name=common_attributes,json=commonAttributes,proto3" json:"common_attributes,omitempty"`
	// HorizontalPodScalerDynamicConfig is the dynamic configuration for the horizontal pod scaler.
	HorizontalPodScalerDynamicConfig *v1.HorizontalPodScaler_ScaleActuator_DynamicConfig `protobuf:"bytes,2,opt,name=horizontal_pod_scaler_dynamic_config,json=horizontalPodScalerDynamicConfig,proto3" json:"horizontal_pod_scaler_dynamic_config,omitempty"`
}

func (x *HorizontalPodScalerDynamicConfigWrapper) Reset() {
	*x = HorizontalPodScalerDynamicConfigWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HorizontalPodScalerDynamicConfigWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HorizontalPodScalerDynamicConfigWrapper) ProtoMessage() {}

func (x *HorizontalPodScalerDynamicConfigWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HorizontalPodScalerDynamicConfigWrapper.ProtoReflect.Descriptor instead.
func (*HorizontalPodScalerDynamicConfigWrapper) Descriptor() ([]byte, []int) {
	return file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescGZIP(), []int{1}
}

func (x *HorizontalPodScalerDynamicConfigWrapper) GetCommonAttributes() *CommonAttributes {
	if x != nil {
		return x.CommonAttributes
	}
	return nil
}

func (x *HorizontalPodScalerDynamicConfigWrapper) GetHorizontalPodScalerDynamicConfig() *v1.HorizontalPodScaler_ScaleActuator_DynamicConfig {
	if x != nil {
		return x.HorizontalPodScalerDynamicConfig
	}
	return nil
}

type ScaleStatusWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CommonAttributes
	CommonAttributes *CommonAttributes `protobuf:"bytes,1,opt,name=common_attributes,json=commonAttributes,proto3" json:"common_attributes,omitempty"`
	ScaleStatus      *ScaleStatus      `protobuf:"bytes,2,opt,name=scale_status,json=scaleStatus,proto3" json:"scale_status,omitempty"`
}

func (x *ScaleStatusWrapper) Reset() {
	*x = ScaleStatusWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScaleStatusWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScaleStatusWrapper) ProtoMessage() {}

func (x *ScaleStatusWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScaleStatusWrapper.ProtoReflect.Descriptor instead.
func (*ScaleStatusWrapper) Descriptor() ([]byte, []int) {
	return file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescGZIP(), []int{2}
}

func (x *ScaleStatusWrapper) GetCommonAttributes() *CommonAttributes {
	if x != nil {
		return x.CommonAttributes
	}
	return nil
}

func (x *ScaleStatusWrapper) GetScaleStatus() *ScaleStatus {
	if x != nil {
		return x.ScaleStatus
	}
	return nil
}

type ScaleStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfiguredReplicas int32 `protobuf:"varint,1,opt,name=configured_replicas,json=configuredReplicas,proto3" json:"configured_replicas,omitempty"`
	ActualReplicas     int32 `protobuf:"varint,2,opt,name=actual_replicas,json=actualReplicas,proto3" json:"actual_replicas,omitempty"`
}

func (x *ScaleStatus) Reset() {
	*x = ScaleStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScaleStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScaleStatus) ProtoMessage() {}

func (x *ScaleStatus) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScaleStatus.ProtoReflect.Descriptor instead.
func (*ScaleStatus) Descriptor() ([]byte, []int) {
	return file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescGZIP(), []int{3}
}

func (x *ScaleStatus) GetConfiguredReplicas() int32 {
	if x != nil {
		return x.ConfiguredReplicas
	}
	return 0
}

func (x *ScaleStatus) GetActualReplicas() int32 {
	if x != nil {
		return x.ActualReplicas
	}
	return 0
}

type ScaleDecisionWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonAttributes *CommonAttributes `protobuf:"bytes,1,opt,name=common_attributes,json=commonAttributes,proto3" json:"common_attributes,omitempty"`
	ScaleDecision    *ScaleDecision    `protobuf:"bytes,2,opt,name=scale_decision,json=scaleDecision,proto3" json:"scale_decision,omitempty"`
}

func (x *ScaleDecisionWrapper) Reset() {
	*x = ScaleDecisionWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScaleDecisionWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScaleDecisionWrapper) ProtoMessage() {}

func (x *ScaleDecisionWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScaleDecisionWrapper.ProtoReflect.Descriptor instead.
func (*ScaleDecisionWrapper) Descriptor() ([]byte, []int) {
	return file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescGZIP(), []int{4}
}

func (x *ScaleDecisionWrapper) GetCommonAttributes() *CommonAttributes {
	if x != nil {
		return x.CommonAttributes
	}
	return nil
}

func (x *ScaleDecisionWrapper) GetScaleDecision() *ScaleDecision {
	if x != nil {
		return x.ScaleDecision
	}
	return nil
}

type ScaleDecision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DesiredReplicas int32 `protobuf:"varint,1,opt,name=desired_replicas,json=desiredReplicas,proto3" json:"desired_replicas,omitempty"`
}

func (x *ScaleDecision) Reset() {
	*x = ScaleDecision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScaleDecision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScaleDecision) ProtoMessage() {}

func (x *ScaleDecision) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScaleDecision.ProtoReflect.Descriptor instead.
func (*ScaleDecision) Descriptor() ([]byte, []int) {
	return file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescGZIP(), []int{5}
}

func (x *ScaleDecision) GetDesiredReplicas() int32 {
	if x != nil {
		return x.DesiredReplicas
	}
	return 0
}

var File_aperture_policy_sync_v1_pod_autoscaler_proto protoreflect.FileDescriptor

var file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64, 0x5f, 0x61, 0x75,
	0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x1a, 0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61,
	0x6c, 0x50, 0x6f, 0x64, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x12, 0x56, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61,
	0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x64, 0x0a, 0x15, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x64, 0x5f, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61,
	0x6c, 0x50, 0x6f, 0x64, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x52, 0x13, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x64, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x22,
	0xa0, 0x02, 0x0a, 0x27, 0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x50, 0x6f,
	0x64, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x9c, 0x01, 0x0a, 0x24, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74,
	0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x64, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x5f, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x64, 0x53, 0x63,
	0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x20, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x64, 0x53,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0xb5, 0x01, 0x0a, 0x12, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52,
	0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x47, 0x0a, 0x0c, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x67, 0x0a, 0x0b, 0x53, 0x63,
	0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63,
	0x74, 0x75, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x14, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x44, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x5f, 0x64, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61,
	0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x0d, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x42,
	0x93, 0x02, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a,
	0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x65, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x2e, 0x76, 0x31, 0x42, 0x12, 0x50, 0x6f, 0x64, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2f,
	0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x79, 0x6e, 0x63, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x50, 0x53, 0xaa, 0x02,
	0x17, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x41, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x5c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c, 0x53, 0x79, 0x6e, 0x63, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5c, 0x53, 0x79, 0x6e, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x3a, 0x3a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x3a, 0x53, 0x79, 0x6e,
	0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescOnce sync.Once
	file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescData = file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDesc
)

func file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescGZIP() []byte {
	file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescOnce.Do(func() {
		file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescData = protoimpl.X.CompressGZIP(file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescData)
	})
	return file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDescData
}

var file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_aperture_policy_sync_v1_pod_autoscaler_proto_goTypes = []interface{}{
	(*HorizontalPodScalerWrapper)(nil),                         // 0: aperture.policy.sync.v1.HorizontalPodScalerWrapper
	(*HorizontalPodScalerDynamicConfigWrapper)(nil),            // 1: aperture.policy.sync.v1.HorizontalPodScalerDynamicConfigWrapper
	(*ScaleStatusWrapper)(nil),                                 // 2: aperture.policy.sync.v1.ScaleStatusWrapper
	(*ScaleStatus)(nil),                                        // 3: aperture.policy.sync.v1.ScaleStatus
	(*ScaleDecisionWrapper)(nil),                               // 4: aperture.policy.sync.v1.ScaleDecisionWrapper
	(*ScaleDecision)(nil),                                      // 5: aperture.policy.sync.v1.ScaleDecision
	(*CommonAttributes)(nil),                                   // 6: aperture.policy.sync.v1.CommonAttributes
	(*v1.HorizontalPodScaler)(nil),                             // 7: aperture.policy.language.v1.HorizontalPodScaler
	(*v1.HorizontalPodScaler_ScaleActuator_DynamicConfig)(nil), // 8: aperture.policy.language.v1.HorizontalPodScaler.ScaleActuator.DynamicConfig
}
var file_aperture_policy_sync_v1_pod_autoscaler_proto_depIdxs = []int32{
	6, // 0: aperture.policy.sync.v1.HorizontalPodScalerWrapper.common_attributes:type_name -> aperture.policy.sync.v1.CommonAttributes
	7, // 1: aperture.policy.sync.v1.HorizontalPodScalerWrapper.horizontal_pod_scaler:type_name -> aperture.policy.language.v1.HorizontalPodScaler
	6, // 2: aperture.policy.sync.v1.HorizontalPodScalerDynamicConfigWrapper.common_attributes:type_name -> aperture.policy.sync.v1.CommonAttributes
	8, // 3: aperture.policy.sync.v1.HorizontalPodScalerDynamicConfigWrapper.horizontal_pod_scaler_dynamic_config:type_name -> aperture.policy.language.v1.HorizontalPodScaler.ScaleActuator.DynamicConfig
	6, // 4: aperture.policy.sync.v1.ScaleStatusWrapper.common_attributes:type_name -> aperture.policy.sync.v1.CommonAttributes
	3, // 5: aperture.policy.sync.v1.ScaleStatusWrapper.scale_status:type_name -> aperture.policy.sync.v1.ScaleStatus
	6, // 6: aperture.policy.sync.v1.ScaleDecisionWrapper.common_attributes:type_name -> aperture.policy.sync.v1.CommonAttributes
	5, // 7: aperture.policy.sync.v1.ScaleDecisionWrapper.scale_decision:type_name -> aperture.policy.sync.v1.ScaleDecision
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_aperture_policy_sync_v1_pod_autoscaler_proto_init() }
func file_aperture_policy_sync_v1_pod_autoscaler_proto_init() {
	if File_aperture_policy_sync_v1_pod_autoscaler_proto != nil {
		return
	}
	file_aperture_policy_sync_v1_common_attributes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HorizontalPodScalerWrapper); i {
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
		file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HorizontalPodScalerDynamicConfigWrapper); i {
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
		file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScaleStatusWrapper); i {
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
		file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScaleStatus); i {
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
		file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScaleDecisionWrapper); i {
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
		file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScaleDecision); i {
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
			RawDescriptor: file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aperture_policy_sync_v1_pod_autoscaler_proto_goTypes,
		DependencyIndexes: file_aperture_policy_sync_v1_pod_autoscaler_proto_depIdxs,
		MessageInfos:      file_aperture_policy_sync_v1_pod_autoscaler_proto_msgTypes,
	}.Build()
	File_aperture_policy_sync_v1_pod_autoscaler_proto = out.File
	file_aperture_policy_sync_v1_pod_autoscaler_proto_rawDesc = nil
	file_aperture_policy_sync_v1_pod_autoscaler_proto_goTypes = nil
	file_aperture_policy_sync_v1_pod_autoscaler_proto_depIdxs = nil
}
