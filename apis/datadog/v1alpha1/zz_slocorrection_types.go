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

type SLOCorrectionInitParameters struct {

	// (String) Category the SLO correction belongs to. Valid values are Scheduled Maintenance, Outside Business Hours, Deployment, Other.
	// Category the SLO correction belongs to. Valid values are `Scheduled Maintenance`, `Outside Business Hours`, `Deployment`, `Other`.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// (String) Description of the correction being made.
	// Description of the correction being made.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Number) Length of time in seconds for a specified rrule recurring SLO correction
	// Length of time in seconds for a specified `rrule` recurring SLO correction (required if specifying `rrule`)
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (Number) Ending time of the correction in epoch seconds. Required for one time corrections, but optional if rrule is specified
	// Ending time of the correction in epoch seconds. Required for one time corrections, but optional if `rrule` is specified
	End *float64 `json:"end,omitempty" tf:"end,omitempty"`

	// (String) Recurrence rules as defined in the iCalendar RFC 5545. Supported rules for SLO corrections are FREQ, INTERVAL, COUNT and UNTIL.
	// Recurrence rules as defined in the iCalendar RFC 5545. Supported rules for SLO corrections are `FREQ`, `INTERVAL`, `COUNT` and `UNTIL`.
	Rrule *string `json:"rrule,omitempty" tf:"rrule,omitempty"`

	// (String) ID of the SLO that this correction will be applied to.
	// ID of the SLO that this correction will be applied to.
	SLOID *string `json:"sloId,omitempty" tf:"slo_id,omitempty"`

	// (Number) Starting time of the correction in epoch seconds.
	// Starting time of the correction in epoch seconds.
	Start *float64 `json:"start,omitempty" tf:"start,omitempty"`

	// (String) The timezone to display in the UI for the correction times (defaults to "UTC")
	// The timezone to display in the UI for the correction times (defaults to "UTC")
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type SLOCorrectionObservation struct {

	// (String) Category the SLO correction belongs to. Valid values are Scheduled Maintenance, Outside Business Hours, Deployment, Other.
	// Category the SLO correction belongs to. Valid values are `Scheduled Maintenance`, `Outside Business Hours`, `Deployment`, `Other`.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// (String) Description of the correction being made.
	// Description of the correction being made.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Number) Length of time in seconds for a specified rrule recurring SLO correction
	// Length of time in seconds for a specified `rrule` recurring SLO correction (required if specifying `rrule`)
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (Number) Ending time of the correction in epoch seconds. Required for one time corrections, but optional if rrule is specified
	// Ending time of the correction in epoch seconds. Required for one time corrections, but optional if `rrule` is specified
	End *float64 `json:"end,omitempty" tf:"end,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Recurrence rules as defined in the iCalendar RFC 5545. Supported rules for SLO corrections are FREQ, INTERVAL, COUNT and UNTIL.
	// Recurrence rules as defined in the iCalendar RFC 5545. Supported rules for SLO corrections are `FREQ`, `INTERVAL`, `COUNT` and `UNTIL`.
	Rrule *string `json:"rrule,omitempty" tf:"rrule,omitempty"`

	// (String) ID of the SLO that this correction will be applied to.
	// ID of the SLO that this correction will be applied to.
	SLOID *string `json:"sloId,omitempty" tf:"slo_id,omitempty"`

	// (Number) Starting time of the correction in epoch seconds.
	// Starting time of the correction in epoch seconds.
	Start *float64 `json:"start,omitempty" tf:"start,omitempty"`

	// (String) The timezone to display in the UI for the correction times (defaults to "UTC")
	// The timezone to display in the UI for the correction times (defaults to "UTC")
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type SLOCorrectionParameters struct {

	// (String) Category the SLO correction belongs to. Valid values are Scheduled Maintenance, Outside Business Hours, Deployment, Other.
	// Category the SLO correction belongs to. Valid values are `Scheduled Maintenance`, `Outside Business Hours`, `Deployment`, `Other`.
	// +kubebuilder:validation:Optional
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// (String) Description of the correction being made.
	// Description of the correction being made.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Number) Length of time in seconds for a specified rrule recurring SLO correction
	// Length of time in seconds for a specified `rrule` recurring SLO correction (required if specifying `rrule`)
	// +kubebuilder:validation:Optional
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (Number) Ending time of the correction in epoch seconds. Required for one time corrections, but optional if rrule is specified
	// Ending time of the correction in epoch seconds. Required for one time corrections, but optional if `rrule` is specified
	// +kubebuilder:validation:Optional
	End *float64 `json:"end,omitempty" tf:"end,omitempty"`

	// (String) Recurrence rules as defined in the iCalendar RFC 5545. Supported rules for SLO corrections are FREQ, INTERVAL, COUNT and UNTIL.
	// Recurrence rules as defined in the iCalendar RFC 5545. Supported rules for SLO corrections are `FREQ`, `INTERVAL`, `COUNT` and `UNTIL`.
	// +kubebuilder:validation:Optional
	Rrule *string `json:"rrule,omitempty" tf:"rrule,omitempty"`

	// (String) ID of the SLO that this correction will be applied to.
	// ID of the SLO that this correction will be applied to.
	// +kubebuilder:validation:Optional
	SLOID *string `json:"sloId,omitempty" tf:"slo_id,omitempty"`

	// (Number) Starting time of the correction in epoch seconds.
	// Starting time of the correction in epoch seconds.
	// +kubebuilder:validation:Optional
	Start *float64 `json:"start,omitempty" tf:"start,omitempty"`

	// (String) The timezone to display in the UI for the correction times (defaults to "UTC")
	// The timezone to display in the UI for the correction times (defaults to "UTC")
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

// SLOCorrectionSpec defines the desired state of SLOCorrection
type SLOCorrectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SLOCorrectionParameters `json:"forProvider"`
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
	InitProvider SLOCorrectionInitParameters `json:"initProvider,omitempty"`
}

// SLOCorrectionStatus defines the observed state of SLOCorrection.
type SLOCorrectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SLOCorrectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SLOCorrection is the Schema for the SLOCorrections API. Resource for interacting with the slo_correction API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type SLOCorrection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.category) || (has(self.initProvider) && has(self.initProvider.category))",message="spec.forProvider.category is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sloId) || (has(self.initProvider) && has(self.initProvider.sloId))",message="spec.forProvider.sloId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.start) || (has(self.initProvider) && has(self.initProvider.start))",message="spec.forProvider.start is a required parameter"
	Spec   SLOCorrectionSpec   `json:"spec"`
	Status SLOCorrectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SLOCorrectionList contains a list of SLOCorrections
type SLOCorrectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SLOCorrection `json:"items"`
}

// Repository type metadata.
var (
	SLOCorrection_Kind             = "SLOCorrection"
	SLOCorrection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SLOCorrection_Kind}.String()
	SLOCorrection_KindAPIVersion   = SLOCorrection_Kind + "." + CRDGroupVersion.String()
	SLOCorrection_GroupVersionKind = CRDGroupVersion.WithKind(SLOCorrection_Kind)
)

func init() {
	SchemeBuilder.Register(&SLOCorrection{}, &SLOCorrectionList{})
}
