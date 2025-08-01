/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MonitorJSONInitParameters struct {

	// (String) The JSON formatted definition of the monitor.
	// The JSON formatted definition of the monitor.
	Monitor *string `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// (String) The URL of the monitor.
	// The URL of the monitor.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type MonitorJSONObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The JSON formatted definition of the monitor.
	// The JSON formatted definition of the monitor.
	Monitor *string `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// (String) The URL of the monitor.
	// The URL of the monitor.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type MonitorJSONParameters struct {

	// (String) The JSON formatted definition of the monitor.
	// The JSON formatted definition of the monitor.
	// +kubebuilder:validation:Optional
	Monitor *string `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// (String) The URL of the monitor.
	// The URL of the monitor.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// MonitorJSONSpec defines the desired state of MonitorJSON
type MonitorJSONSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorJSONParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MonitorJSONInitParameters `json:"initProvider,omitempty"`
}

// MonitorJSONStatus defines the observed state of MonitorJSON.
type MonitorJSONStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorJSONObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MonitorJSON is the Schema for the MonitorJSONs API. Provides a Datadog monitor JSON resource. This can be used to create and manage Datadog monitors using the JSON definition.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type MonitorJSON struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.monitor) || (has(self.initProvider) && has(self.initProvider.monitor))",message="spec.forProvider.monitor is a required parameter"
	Spec   MonitorJSONSpec   `json:"spec"`
	Status MonitorJSONStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorJSONList contains a list of MonitorJSONs
type MonitorJSONList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorJSON `json:"items"`
}

// Repository type metadata.
var (
	MonitorJSON_Kind             = "MonitorJSON"
	MonitorJSON_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorJSON_Kind}.String()
	MonitorJSON_KindAPIVersion   = MonitorJSON_Kind + "." + CRDGroupVersion.String()
	MonitorJSON_GroupVersionKind = CRDGroupVersion.WithKind(MonitorJSON_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorJSON{}, &MonitorJSONList{})
}
