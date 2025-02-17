// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package languagev1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using Classifier within kubernetes types, where deepcopy-gen is used.
func (in *Classifier) DeepCopyInto(out *Classifier) {
	p := proto.Clone(in).(*Classifier)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Classifier. Required by controller-gen.
func (in *Classifier) DeepCopy() *Classifier {
	if in == nil {
		return nil
	}
	out := new(Classifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Classifier. Required by controller-gen.
func (in *Classifier) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rule within kubernetes types, where deepcopy-gen is used.
func (in *Rule) DeepCopyInto(out *Rule) {
	p := proto.Clone(in).(*Rule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule. Required by controller-gen.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rule. Required by controller-gen.
func (in *Rule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rule_Rego within kubernetes types, where deepcopy-gen is used.
func (in *Rule_Rego) DeepCopyInto(out *Rule_Rego) {
	p := proto.Clone(in).(*Rule_Rego)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule_Rego. Required by controller-gen.
func (in *Rule_Rego) DeepCopy() *Rule_Rego {
	if in == nil {
		return nil
	}
	out := new(Rule_Rego)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rule_Rego. Required by controller-gen.
func (in *Rule_Rego) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Extractor within kubernetes types, where deepcopy-gen is used.
func (in *Extractor) DeepCopyInto(out *Extractor) {
	p := proto.Clone(in).(*Extractor)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Extractor. Required by controller-gen.
func (in *Extractor) DeepCopy() *Extractor {
	if in == nil {
		return nil
	}
	out := new(Extractor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Extractor. Required by controller-gen.
func (in *Extractor) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using JSONExtractor within kubernetes types, where deepcopy-gen is used.
func (in *JSONExtractor) DeepCopyInto(out *JSONExtractor) {
	p := proto.Clone(in).(*JSONExtractor)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONExtractor. Required by controller-gen.
func (in *JSONExtractor) DeepCopy() *JSONExtractor {
	if in == nil {
		return nil
	}
	out := new(JSONExtractor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new JSONExtractor. Required by controller-gen.
func (in *JSONExtractor) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using AddressExtractor within kubernetes types, where deepcopy-gen is used.
func (in *AddressExtractor) DeepCopyInto(out *AddressExtractor) {
	p := proto.Clone(in).(*AddressExtractor)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressExtractor. Required by controller-gen.
func (in *AddressExtractor) DeepCopy() *AddressExtractor {
	if in == nil {
		return nil
	}
	out := new(AddressExtractor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AddressExtractor. Required by controller-gen.
func (in *AddressExtractor) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using JWTExtractor within kubernetes types, where deepcopy-gen is used.
func (in *JWTExtractor) DeepCopyInto(out *JWTExtractor) {
	p := proto.Clone(in).(*JWTExtractor)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTExtractor. Required by controller-gen.
func (in *JWTExtractor) DeepCopy() *JWTExtractor {
	if in == nil {
		return nil
	}
	out := new(JWTExtractor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new JWTExtractor. Required by controller-gen.
func (in *JWTExtractor) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PathTemplateMatcher within kubernetes types, where deepcopy-gen is used.
func (in *PathTemplateMatcher) DeepCopyInto(out *PathTemplateMatcher) {
	p := proto.Clone(in).(*PathTemplateMatcher)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PathTemplateMatcher. Required by controller-gen.
func (in *PathTemplateMatcher) DeepCopy() *PathTemplateMatcher {
	if in == nil {
		return nil
	}
	out := new(PathTemplateMatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PathTemplateMatcher. Required by controller-gen.
func (in *PathTemplateMatcher) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
