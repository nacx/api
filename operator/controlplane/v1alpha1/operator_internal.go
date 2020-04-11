// Copyright (c) Tetrate, Inc 2020 All Rights Reserved.

// This file was manually created.

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
)

// ControlPlaneConfigSpec is an internal struct that reflects the state of TSB control plane components. This is an overlay
// of ControlPlaneUserConfigSpec on Tsb pre built default profile. This struct is for internal use and not exposed to the user.
type ControlPlaneConfigSpec struct {
	// Default root for docker image paths e.g. docker.io/tetrate
	Hub string `protobuf:"bytes,11,opt,name=hub,proto3" json:"hub,omitempty"`
	// Default version tag for docker images e.g. 0.6.9
	Tag string `protobuf:"bytes,12,opt,name=tag,proto3" json:"tag,omitempty"`
	// Namespace for TSB
	Namespace string `protobuf:"bytes,13,opt,name=namespace,json=Namespace,proto3" json:"namespace,omitempty"`
	// Kubernetes resource settings, enablement and component-specific settings that are not internal
	// to the component.
	Components *ControlPlaneComponentSetSpec `protobuf:"bytes,50,opt,name=components,proto3" json:"components,omitempty"`
	// Overrides for default values.yaml. This is a validated pass-through to Helm templates.
	Values map[string]interface{} `protobuf:"bytes,100,opt,name=values,proto3" json:"values,omitempty"`
	// Unvalidated overrides for default values.yaml. Used for custom templates where new parameters
	// are added.
	UnvalidatedValues map[string]interface{} `protobuf:"bytes,101,opt,name=unvalidated_values,json=unvalidatedValues,proto3" json:"unvalidated_values,omitempty"`
}

func (m *ControlPlaneConfigSpec) Reset()         { *m = ControlPlaneConfigSpec{} }
func (m *ControlPlaneConfigSpec) String() string { return proto.CompactTextString(m) }
func (*ControlPlaneConfigSpec) ProtoMessage()    {}

func (m *ControlPlaneConfigSpec) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

func (m *ControlPlaneConfigSpec) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ControlPlaneConfigSpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ControlPlaneConfigSpec) GetComponents() *ControlPlaneComponentSetSpec {
	if m != nil {
		return m.Components
	}
	return nil
}

func init() {
	proto.RegisterType((*ControlPlaneConfigSpec)(nil), "tetrate.api.operator.controlplane.v1alpha1.ControlPlaneConfigSpec")

}
