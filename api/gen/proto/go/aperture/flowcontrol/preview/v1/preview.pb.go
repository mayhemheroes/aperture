// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aperture/flowcontrol/preview/v1/preview.proto

package previewv1

import (
	v1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/language/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PreviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of samples to collect. Defaults to 10.
	Samples int64 `protobuf:"varint,1,opt,name=samples,proto3" json:"samples,omitempty" validate:"min=1" default:"10"` // @gotags: validate:"min=1" default:"10"
	// Control point to preview.
	ControlPoint string `protobuf:"bytes,2,opt,name=control_point,json=controlPoint,proto3" json:"control_point,omitempty" validate:"required"` // @gotags: validate:"required"
	// Service to preview. Empty value implies catch all service.
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	// Flow labels to match. Optional parameter for advanced filtering.
	LabelMatcher *v1.LabelMatcher `protobuf:"bytes,4,opt,name=label_matcher,json=labelMatcher,proto3" json:"label_matcher,omitempty"`
}

func (x *PreviewRequest) Reset() {
	*x = PreviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviewRequest) ProtoMessage() {}

func (x *PreviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviewRequest.ProtoReflect.Descriptor instead.
func (*PreviewRequest) Descriptor() ([]byte, []int) {
	return file_aperture_flowcontrol_preview_v1_preview_proto_rawDescGZIP(), []int{0}
}

func (x *PreviewRequest) GetSamples() int64 {
	if x != nil {
		return x.Samples
	}
	return 0
}

func (x *PreviewRequest) GetControlPoint() string {
	if x != nil {
		return x.ControlPoint
	}
	return ""
}

func (x *PreviewRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *PreviewRequest) GetLabelMatcher() *v1.LabelMatcher {
	if x != nil {
		return x.LabelMatcher
	}
	return nil
}

type PreviewFlowLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Samples []*PreviewFlowLabelsResponse_FlowLabels `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *PreviewFlowLabelsResponse) Reset() {
	*x = PreviewFlowLabelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviewFlowLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviewFlowLabelsResponse) ProtoMessage() {}

func (x *PreviewFlowLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviewFlowLabelsResponse.ProtoReflect.Descriptor instead.
func (*PreviewFlowLabelsResponse) Descriptor() ([]byte, []int) {
	return file_aperture_flowcontrol_preview_v1_preview_proto_rawDescGZIP(), []int{1}
}

func (x *PreviewFlowLabelsResponse) GetSamples() []*PreviewFlowLabelsResponse_FlowLabels {
	if x != nil {
		return x.Samples
	}
	return nil
}

type PreviewHTTPRequestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Samples []*structpb.Struct `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *PreviewHTTPRequestsResponse) Reset() {
	*x = PreviewHTTPRequestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviewHTTPRequestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviewHTTPRequestsResponse) ProtoMessage() {}

func (x *PreviewHTTPRequestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviewHTTPRequestsResponse.ProtoReflect.Descriptor instead.
func (*PreviewHTTPRequestsResponse) Descriptor() ([]byte, []int) {
	return file_aperture_flowcontrol_preview_v1_preview_proto_rawDescGZIP(), []int{2}
}

func (x *PreviewHTTPRequestsResponse) GetSamples() []*structpb.Struct {
	if x != nil {
		return x.Samples
	}
	return nil
}

type PreviewFlowLabelsResponse_FlowLabels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PreviewFlowLabelsResponse_FlowLabels) Reset() {
	*x = PreviewFlowLabelsResponse_FlowLabels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviewFlowLabelsResponse_FlowLabels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviewFlowLabelsResponse_FlowLabels) ProtoMessage() {}

func (x *PreviewFlowLabelsResponse_FlowLabels) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviewFlowLabelsResponse_FlowLabels.ProtoReflect.Descriptor instead.
func (*PreviewFlowLabelsResponse_FlowLabels) Descriptor() ([]byte, []int) {
	return file_aperture_flowcontrol_preview_v1_preview_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PreviewFlowLabelsResponse_FlowLabels) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_aperture_flowcontrol_preview_v1_preview_proto protoreflect.FileDescriptor

var file_aperture_flowcontrol_preview_v1_preview_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31,
	0x1a, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0c, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0xb1, 0x02, 0x0a, 0x19, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52,
	0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x1a, 0xb2, 0x01, 0x0a, 0x0a, 0x46, 0x6c, 0x6f,
	0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x69, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x50, 0x0a,
	0x1b, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x32,
	0xfa, 0x04, 0x0a, 0x12, 0x46, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa7, 0x02, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2f, 0x2e, 0x61,
	0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e,
	0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa4, 0x01, 0x92, 0x41, 0x10, 0x0a,
	0x0e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x8a, 0x01, 0x3a, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x5a, 0x3f, 0x3a, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x6f, 0x77,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x7d, 0x22, 0x38, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x6f, 0x77,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x7d,
	0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x7d,
	0x12, 0xb9, 0x02, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x48, 0x54, 0x54, 0x50,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x2f, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x61, 0x70, 0x65, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb2, 0x01, 0x92, 0x41, 0x10, 0x0a, 0x0e, 0x61,
	0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x98, 0x01, 0x3a, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x5a, 0x46, 0x3a, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x22, 0x35, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x7d, 0x22, 0x3f, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x7d, 0x2f, 0x7b, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x7d, 0x42, 0xc0, 0x02, 0x0a,
	0x37, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2f, 0x61,
	0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x46, 0x50, 0xaa, 0x02, 0x1f, 0x41, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1f, 0x41, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x46, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5c, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2b, 0x41,
	0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x46, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x5c, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x22, 0x41, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x3a, 0x3a, 0x46, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x3a, 0x3a, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aperture_flowcontrol_preview_v1_preview_proto_rawDescOnce sync.Once
	file_aperture_flowcontrol_preview_v1_preview_proto_rawDescData = file_aperture_flowcontrol_preview_v1_preview_proto_rawDesc
)

func file_aperture_flowcontrol_preview_v1_preview_proto_rawDescGZIP() []byte {
	file_aperture_flowcontrol_preview_v1_preview_proto_rawDescOnce.Do(func() {
		file_aperture_flowcontrol_preview_v1_preview_proto_rawDescData = protoimpl.X.CompressGZIP(file_aperture_flowcontrol_preview_v1_preview_proto_rawDescData)
	})
	return file_aperture_flowcontrol_preview_v1_preview_proto_rawDescData
}

var file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_aperture_flowcontrol_preview_v1_preview_proto_goTypes = []interface{}{
	(*PreviewRequest)(nil),                       // 0: aperture.flowcontrol.preview.v1.PreviewRequest
	(*PreviewFlowLabelsResponse)(nil),            // 1: aperture.flowcontrol.preview.v1.PreviewFlowLabelsResponse
	(*PreviewHTTPRequestsResponse)(nil),          // 2: aperture.flowcontrol.preview.v1.PreviewHTTPRequestsResponse
	(*PreviewFlowLabelsResponse_FlowLabels)(nil), // 3: aperture.flowcontrol.preview.v1.PreviewFlowLabelsResponse.FlowLabels
	nil,                     // 4: aperture.flowcontrol.preview.v1.PreviewFlowLabelsResponse.FlowLabels.LabelsEntry
	(*v1.LabelMatcher)(nil), // 5: aperture.policy.language.v1.LabelMatcher
	(*structpb.Struct)(nil), // 6: google.protobuf.Struct
}
var file_aperture_flowcontrol_preview_v1_preview_proto_depIdxs = []int32{
	5, // 0: aperture.flowcontrol.preview.v1.PreviewRequest.label_matcher:type_name -> aperture.policy.language.v1.LabelMatcher
	3, // 1: aperture.flowcontrol.preview.v1.PreviewFlowLabelsResponse.samples:type_name -> aperture.flowcontrol.preview.v1.PreviewFlowLabelsResponse.FlowLabels
	6, // 2: aperture.flowcontrol.preview.v1.PreviewHTTPRequestsResponse.samples:type_name -> google.protobuf.Struct
	4, // 3: aperture.flowcontrol.preview.v1.PreviewFlowLabelsResponse.FlowLabels.labels:type_name -> aperture.flowcontrol.preview.v1.PreviewFlowLabelsResponse.FlowLabels.LabelsEntry
	0, // 4: aperture.flowcontrol.preview.v1.FlowPreviewService.PreviewFlowLabels:input_type -> aperture.flowcontrol.preview.v1.PreviewRequest
	0, // 5: aperture.flowcontrol.preview.v1.FlowPreviewService.PreviewHTTPRequests:input_type -> aperture.flowcontrol.preview.v1.PreviewRequest
	1, // 6: aperture.flowcontrol.preview.v1.FlowPreviewService.PreviewFlowLabels:output_type -> aperture.flowcontrol.preview.v1.PreviewFlowLabelsResponse
	2, // 7: aperture.flowcontrol.preview.v1.FlowPreviewService.PreviewHTTPRequests:output_type -> aperture.flowcontrol.preview.v1.PreviewHTTPRequestsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_aperture_flowcontrol_preview_v1_preview_proto_init() }
func file_aperture_flowcontrol_preview_v1_preview_proto_init() {
	if File_aperture_flowcontrol_preview_v1_preview_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviewRequest); i {
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
		file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviewFlowLabelsResponse); i {
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
		file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviewHTTPRequestsResponse); i {
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
		file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviewFlowLabelsResponse_FlowLabels); i {
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
			RawDescriptor: file_aperture_flowcontrol_preview_v1_preview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aperture_flowcontrol_preview_v1_preview_proto_goTypes,
		DependencyIndexes: file_aperture_flowcontrol_preview_v1_preview_proto_depIdxs,
		MessageInfos:      file_aperture_flowcontrol_preview_v1_preview_proto_msgTypes,
	}.Build()
	File_aperture_flowcontrol_preview_v1_preview_proto = out.File
	file_aperture_flowcontrol_preview_v1_preview_proto_rawDesc = nil
	file_aperture_flowcontrol_preview_v1_preview_proto_goTypes = nil
	file_aperture_flowcontrol_preview_v1_preview_proto_depIdxs = nil
}
