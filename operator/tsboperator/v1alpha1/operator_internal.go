// Copyright (c) Tetrate, Inc 2020 All Rights Reserved.

// This file was manually created.

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
)

// TsbOperatorSpec is an internal struct that reflects the state of TSB management plane components. This is an overlay
// of TsbUserOperatorSpec on Tsb pre built default state. This struct is for internal use and not exposed to the user.
type TsbOperatorSpec struct {
	// Default root for docker image paths e.g. docker.io/tetrate
	Hub string `protobuf:"bytes,11,opt,name=hub,proto3" json:"hub,omitempty"`
	// Default version tag for docker images e.g. 0.6.9
	Tag string `protobuf:"bytes,12,opt,name=tag,proto3" json:"tag,omitempty"`
	// Namespace for TSB
	TsbNamespace string `protobuf:"bytes,13,opt,name=tsb_namespace,json=tsbNamespace,proto3" json:"tsb_namespace,omitempty"`
	// Kubernetes resource settings, enablement and component-specific settings that are not internal
	// to the component.
	Components *TsbComponentSetSpec `protobuf:"bytes,50,opt,name=components,proto3" json:"components,omitempty"`
	// Overrides for default values.yaml. This is a validated pass-through to Helm templates.
	Values map[string]interface{} `protobuf:"bytes,100,opt,name=values,proto3" json:"values,omitempty"`
	// Unvalidated overrides for default values.yaml. Used for custom templates where new parameters
	// are added.
	UnvalidatedValues map[string]interface{} `protobuf:"bytes,101,opt,name=unvalidated_values,json=unvalidatedValues,proto3" json:"unvalidated_values,omitempty"`
}

func (m *TsbOperatorSpec) Reset()         { *m = TsbOperatorSpec{} }
func (m *TsbOperatorSpec) String() string { return proto.CompactTextString(m) }
func (*TsbOperatorSpec) ProtoMessage()    {}

func (m *TsbOperatorSpec) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

func (m *TsbOperatorSpec) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *TsbOperatorSpec) GetTsbNamespace() string {
	if m != nil {
		return m.TsbNamespace
	}
	return ""
}

func (m *TsbOperatorSpec) GetComponents() *TsbComponentSetSpec {
	if m != nil {
		return m.Components
	}
	return nil
}

func init() {
	proto.RegisterType((*TsbOperatorSpec)(nil), "tetrate.api.tsboperator.v1alpha1.TsbOperatorSpec")

}
