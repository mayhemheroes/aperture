/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// +kubebuilder:object:generate=true
// +groupName=fluxninja.com
package v1alpha1

import (
	"github.com/fluxninja/aperture/cmd/aperture-controller/controller"
	"github.com/fluxninja/aperture/operator/api"
	"github.com/fluxninja/aperture/operator/api/common"
	"github.com/fluxninja/aperture/pkg/jobs"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ControllerSpec defines the desired state for the Controller.
type ControllerSpec struct {
	// CommonSpec defines the common state between Agent and Controller
	common.CommonSpec `json:",inline"`

	// Image configuration
	//+kubebuilder:validation:Optional
	Image common.ControllerImage `json:"image"`

	// Pod's host aliases
	//+kubebuilder:validation:Optional
	HostAliases []corev1.HostAlias `json:"hostAliases,omitempty"`

	// Controller Configuration
	//+kubebuilder:validation:Optional
	ConfigSpec ControllerConfigSpec `json:"config"`
}

// ControllerConfigSpec holds controller configuration.
type ControllerConfigSpec struct {
	// CommonSpec
	//+kubebuilder:validation:Optional
	common.CommonConfigSpec `json:",inline"`

	// Policies configuration.
	//+kubebuilder:validation:Optional
	Policies PoliciesConfig `json:"policies"`

	// OTEL configuration.
	//+kubebuilder:validation:Optional
	OTEL controller.ControllerOTELConfig `json:"otel"`
}

// PoliciesConfig for policy engine.
type PoliciesConfig struct {
	// Scheduler for PromQL jobs.
	PromQLJobsScheduler jobs.JobGroupConfig `json:"promql_jobs_scheduler"`
}

// ControllerStatus defines the observed state of Controller.
type ControllerStatus struct {
	Resources string `json:"resources,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Resources",type=string,JSONPath=`.status.resources`
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// Controller is the Schema for the controllers API.
type Controller struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ControllerSpec   `json:"spec,omitempty"`
	Status ControllerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ControllerList contains a list of Controller.
type ControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Controller `json:"items"`
}

func init() {
	api.SchemeBuilder.Register(&Controller{}, &ControllerList{})
}
