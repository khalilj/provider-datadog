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

type AWSTagFilterInitParameters struct {

	// (String) Your AWS Account ID without dashes.
	// Your AWS Account ID without dashes.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) The namespace associated with the tag filter entry. Valid values are elb, application_elb, sqs, rds, custom, network_elb, lambda.
	// The namespace associated with the tag filter entry. Valid values are `elb`, `application_elb`, `sqs`, `rds`, `custom`, `network_elb`, `lambda`.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) The tag filter string.
	// The tag filter string.
	TagFilterStr *string `json:"tagFilterStr,omitempty" tf:"tag_filter_str,omitempty"`
}

type AWSTagFilterObservation struct {

	// (String) Your AWS Account ID without dashes.
	// Your AWS Account ID without dashes.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The namespace associated with the tag filter entry. Valid values are elb, application_elb, sqs, rds, custom, network_elb, lambda.
	// The namespace associated with the tag filter entry. Valid values are `elb`, `application_elb`, `sqs`, `rds`, `custom`, `network_elb`, `lambda`.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) The tag filter string.
	// The tag filter string.
	TagFilterStr *string `json:"tagFilterStr,omitempty" tf:"tag_filter_str,omitempty"`
}

type AWSTagFilterParameters struct {

	// (String) Your AWS Account ID without dashes.
	// Your AWS Account ID without dashes.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) The namespace associated with the tag filter entry. Valid values are elb, application_elb, sqs, rds, custom, network_elb, lambda.
	// The namespace associated with the tag filter entry. Valid values are `elb`, `application_elb`, `sqs`, `rds`, `custom`, `network_elb`, `lambda`.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) The tag filter string.
	// The tag filter string.
	// +kubebuilder:validation:Optional
	TagFilterStr *string `json:"tagFilterStr,omitempty" tf:"tag_filter_str,omitempty"`
}

// AWSTagFilterSpec defines the desired state of AWSTagFilter
type AWSTagFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AWSTagFilterParameters `json:"forProvider"`
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
	InitProvider AWSTagFilterInitParameters `json:"initProvider,omitempty"`
}

// AWSTagFilterStatus defines the observed state of AWSTagFilter.
type AWSTagFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AWSTagFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AWSTagFilter is the Schema for the AWSTagFilters API. Provides a Datadog AWS tag filter resource. This can be used to create and manage Datadog AWS tag filters.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type AWSTagFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || (has(self.initProvider) && has(self.initProvider.accountId))",message="spec.forProvider.accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.__namespace__) || (has(self.initProvider) && has(self.initProvider.__namespace__))",message="spec.forProvider.namespace is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tagFilterStr) || (has(self.initProvider) && has(self.initProvider.tagFilterStr))",message="spec.forProvider.tagFilterStr is a required parameter"
	Spec   AWSTagFilterSpec   `json:"spec"`
	Status AWSTagFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AWSTagFilterList contains a list of AWSTagFilters
type AWSTagFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSTagFilter `json:"items"`
}

// Repository type metadata.
var (
	AWSTagFilter_Kind             = "AWSTagFilter"
	AWSTagFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AWSTagFilter_Kind}.String()
	AWSTagFilter_KindAPIVersion   = AWSTagFilter_Kind + "." + CRDGroupVersion.String()
	AWSTagFilter_GroupVersionKind = CRDGroupVersion.WithKind(AWSTagFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&AWSTagFilter{}, &AWSTagFilterList{})
}
