// Code generated by protoc-gen-go-json. DO NOT EDIT.
// source: aperture/policy/monitoring/v1/policy_view.proto

package monitoringv1

import (
	"google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON implements json.Marshaler
func (msg *PortView) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *PortView) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ComponentView) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ComponentView) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SourceTarget) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SourceTarget) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *Link) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *Link) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}
