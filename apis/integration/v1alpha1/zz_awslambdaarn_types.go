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

type AWSLambdaARNInitParameters struct {

	// (String) Your AWS Account ID without dashes.
	// Your AWS Account ID without dashes.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) The ARN of the Datadog forwarder Lambda.
	// The ARN of the Datadog forwarder Lambda.
	LambdaArn *string `json:"lambdaArn,omitempty" tf:"lambda_arn,omitempty"`
}

type AWSLambdaARNObservation struct {

	// (String) Your AWS Account ID without dashes.
	// Your AWS Account ID without dashes.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The ARN of the Datadog forwarder Lambda.
	// The ARN of the Datadog forwarder Lambda.
	LambdaArn *string `json:"lambdaArn,omitempty" tf:"lambda_arn,omitempty"`
}

type AWSLambdaARNParameters struct {

	// (String) Your AWS Account ID without dashes.
	// Your AWS Account ID without dashes.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) The ARN of the Datadog forwarder Lambda.
	// The ARN of the Datadog forwarder Lambda.
	// +kubebuilder:validation:Optional
	LambdaArn *string `json:"lambdaArn,omitempty" tf:"lambda_arn,omitempty"`
}

// AWSLambdaARNSpec defines the desired state of AWSLambdaARN
type AWSLambdaARNSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AWSLambdaARNParameters `json:"forProvider"`
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
	InitProvider AWSLambdaARNInitParameters `json:"initProvider,omitempty"`
}

// AWSLambdaARNStatus defines the observed state of AWSLambdaARN.
type AWSLambdaARNStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AWSLambdaARNObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AWSLambdaARN is the Schema for the AWSLambdaARNs API. Provides a Datadog - Amazon Web Services integration Lambda ARN resource. This can be used to create and manage the log collection Lambdas for an account. Update operations are currently not supported with datadog API so any change forces a new resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type AWSLambdaARN struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || (has(self.initProvider) && has(self.initProvider.accountId))",message="spec.forProvider.accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lambdaArn) || (has(self.initProvider) && has(self.initProvider.lambdaArn))",message="spec.forProvider.lambdaArn is a required parameter"
	Spec   AWSLambdaARNSpec   `json:"spec"`
	Status AWSLambdaARNStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AWSLambdaARNList contains a list of AWSLambdaARNs
type AWSLambdaARNList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSLambdaARN `json:"items"`
}

// Repository type metadata.
var (
	AWSLambdaARN_Kind             = "AWSLambdaARN"
	AWSLambdaARN_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AWSLambdaARN_Kind}.String()
	AWSLambdaARN_KindAPIVersion   = AWSLambdaARN_Kind + "." + CRDGroupVersion.String()
	AWSLambdaARN_GroupVersionKind = CRDGroupVersion.WithKind(AWSLambdaARN_Kind)
)

func init() {
	SchemeBuilder.Register(&AWSLambdaARN{}, &AWSLambdaARNList{})
}
