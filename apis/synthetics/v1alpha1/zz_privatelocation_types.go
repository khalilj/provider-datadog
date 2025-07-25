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

type MetadataInitParameters struct {

	// (Set of String) A list of role identifiers pulled from the Roles API to restrict read and write access.
	// A list of role identifiers pulled from the Roles API to restrict read and write access.
	// +listType=set
	RestrictedRoles []*string `json:"restrictedRoles,omitempty" tf:"restricted_roles,omitempty"`
}

type MetadataObservation struct {

	// (Set of String) A list of role identifiers pulled from the Roles API to restrict read and write access.
	// A list of role identifiers pulled from the Roles API to restrict read and write access.
	// +listType=set
	RestrictedRoles []*string `json:"restrictedRoles,omitempty" tf:"restricted_roles,omitempty"`
}

type MetadataParameters struct {

	// (Set of String) A list of role identifiers pulled from the Roles API to restrict read and write access.
	// A list of role identifiers pulled from the Roles API to restrict read and write access.
	// +kubebuilder:validation:Optional
	// +listType=set
	RestrictedRoles []*string `json:"restrictedRoles,omitempty" tf:"restricted_roles,omitempty"`
}

type PrivateLocationInitParameters struct {

	// (String) Description of the private location.
	// Description of the private location.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block List, Max: 1) The private location metadata (see below for nested schema)
	// The private location metadata
	Metadata []MetadataInitParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Synthetics private location name.
	// Synthetics private location name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String) A list of tags to associate with your synthetics private location.
	// A list of tags to associate with your synthetics private location.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateLocationObservation struct {

	// (String) Description of the private location.
	// Description of the private location.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Max: 1) The private location metadata (see below for nested schema)
	// The private location metadata
	Metadata []MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Synthetics private location name.
	// Synthetics private location name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String) A list of tags to associate with your synthetics private location.
	// A list of tags to associate with your synthetics private location.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateLocationParameters struct {

	// (String) Description of the private location.
	// Description of the private location.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block List, Max: 1) The private location metadata (see below for nested schema)
	// The private location metadata
	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String) Synthetics private location name.
	// Synthetics private location name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String) A list of tags to associate with your synthetics private location.
	// A list of tags to associate with your synthetics private location.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PrivateLocationSpec defines the desired state of PrivateLocation
type PrivateLocationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateLocationParameters `json:"forProvider"`
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
	InitProvider PrivateLocationInitParameters `json:"initProvider,omitempty"`
}

// PrivateLocationStatus defines the observed state of PrivateLocation.
type PrivateLocationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateLocationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PrivateLocation is the Schema for the PrivateLocations API. Provides a Datadog synthetics private location resource. This can be used to create and manage Datadog synthetics private locations.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type PrivateLocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   PrivateLocationSpec   `json:"spec"`
	Status PrivateLocationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLocationList contains a list of PrivateLocations
type PrivateLocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateLocation `json:"items"`
}

// Repository type metadata.
var (
	PrivateLocation_Kind             = "PrivateLocation"
	PrivateLocation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateLocation_Kind}.String()
	PrivateLocation_KindAPIVersion   = PrivateLocation_Kind + "." + CRDGroupVersion.String()
	PrivateLocation_GroupVersionKind = CRDGroupVersion.WithKind(PrivateLocation_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateLocation{}, &PrivateLocationList{})
}
