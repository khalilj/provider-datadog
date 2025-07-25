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

type ApplicationKeyInitParameters struct {
}

type ApplicationKeyObservation struct {

	// (String)
	Hash *string `json:"hash,omitempty" tf:"hash,omitempty"`

	// (String) Name for Child Organization after creation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`
}

type ApplicationKeyParameters struct {
}

type ChildOrganizationAPIKeyInitParameters struct {
}

type ChildOrganizationAPIKeyObservation struct {

	// (String)
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) Name for Child Organization after creation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ChildOrganizationAPIKeyParameters struct {
}

type ChildOrganizationInitParameters struct {

	// (String) Name for Child Organization after creation.
	// Name for Child Organization after creation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ChildOrganizationObservation struct {

	// (List of Object) Datadog API key. (see below for nested schema)
	// Datadog API key.
	APIKey []ChildOrganizationAPIKeyObservation `json:"apiKey,omitempty" tf:"api_key,omitempty"`

	// (List of Object) An application key with its associated metadata. (see below for nested schema)
	// An application key with its associated metadata.
	ApplicationKey []ApplicationKeyObservation `json:"applicationKey,omitempty" tf:"application_key,omitempty"`

	// (String) Description of the organization.
	// Description of the organization.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Name for Child Organization after creation.
	// Name for Child Organization after creation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The public_id of the organization you are operating within.
	// The `public_id` of the organization you are operating within.
	PublicID *string `json:"publicId,omitempty" tf:"public_id,omitempty"`

	// (List of Object) Organization settings (see below for nested schema)
	// Organization settings
	Settings []SettingsObservation `json:"settings,omitempty" tf:"settings,omitempty"`

	// (List of Object) Information about a user (see below for nested schema)
	// Information about a user
	User []UserObservation `json:"user,omitempty" tf:"user,omitempty"`
}

type ChildOrganizationParameters struct {

	// (String) Name for Child Organization after creation.
	// Name for Child Organization after creation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SAMLAutocreateUsersDomainsInitParameters struct {
}

type SAMLAutocreateUsersDomainsObservation struct {

	// (List of String)
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// (Boolean)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SAMLAutocreateUsersDomainsParameters struct {
}

type SAMLIdpInitiatedLoginInitParameters struct {
}

type SAMLIdpInitiatedLoginObservation struct {

	// (Boolean)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SAMLIdpInitiatedLoginParameters struct {
}

type SAMLInitParameters struct {
}

type SAMLObservation struct {

	// (Boolean)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SAMLParameters struct {
}

type SAMLStrictModeInitParameters struct {
}

type SAMLStrictModeObservation struct {

	// (Boolean)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SAMLStrictModeParameters struct {
}

type SettingsInitParameters struct {
}

type SettingsObservation struct {

	// (Boolean)
	PrivateWidgetShare *bool `json:"privateWidgetShare,omitempty" tf:"private_widget_share,omitempty"`

	// (List of Object) (see below for nested schema)
	SAML []SAMLObservation `json:"saml,omitempty" tf:"saml,omitempty"`

	// (String)
	SAMLAutocreateAccessRole *string `json:"samlAutocreateAccessRole,omitempty" tf:"saml_autocreate_access_role,omitempty"`

	// (List of Object) (see below for nested schema)
	SAMLAutocreateUsersDomains []SAMLAutocreateUsersDomainsObservation `json:"samlAutocreateUsersDomains,omitempty" tf:"saml_autocreate_users_domains,omitempty"`

	// (Boolean)
	SAMLCanBeEnabled *bool `json:"samlCanBeEnabled,omitempty" tf:"saml_can_be_enabled,omitempty"`

	// (String)
	SAMLIdpEndpoint *string `json:"samlIdpEndpoint,omitempty" tf:"saml_idp_endpoint,omitempty"`

	// (List of Object) (see below for nested schema)
	SAMLIdpInitiatedLogin []SAMLIdpInitiatedLoginObservation `json:"samlIdpInitiatedLogin,omitempty" tf:"saml_idp_initiated_login,omitempty"`

	// (Boolean)
	SAMLIdpMetadataUploaded *bool `json:"samlIdpMetadataUploaded,omitempty" tf:"saml_idp_metadata_uploaded,omitempty"`

	// (String)
	SAMLLoginURL *string `json:"samlLoginUrl,omitempty" tf:"saml_login_url,omitempty"`

	// (List of Object) (see below for nested schema)
	SAMLStrictMode []SAMLStrictModeObservation `json:"samlStrictMode,omitempty" tf:"saml_strict_mode,omitempty"`
}

type SettingsParameters struct {
}

type UserInitParameters struct {
}

type UserObservation struct {

	// (String)
	AccessRole *string `json:"accessRole,omitempty" tf:"access_role,omitempty"`

	// (String)
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) Name for Child Organization after creation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type UserParameters struct {
}

// ChildOrganizationSpec defines the desired state of ChildOrganization
type ChildOrganizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ChildOrganizationParameters `json:"forProvider"`
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
	InitProvider ChildOrganizationInitParameters `json:"initProvider,omitempty"`
}

// ChildOrganizationStatus defines the observed state of ChildOrganization.
type ChildOrganizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ChildOrganizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ChildOrganization is the Schema for the ChildOrganizations API. Provides a Datadog Child Organization resource. This can be used to create Datadog Child Organizations. To manage created organization use datadog_organization_settings.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type ChildOrganization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ChildOrganizationSpec   `json:"spec"`
	Status ChildOrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChildOrganizationList contains a list of ChildOrganizations
type ChildOrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChildOrganization `json:"items"`
}

// Repository type metadata.
var (
	ChildOrganization_Kind             = "ChildOrganization"
	ChildOrganization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ChildOrganization_Kind}.String()
	ChildOrganization_KindAPIVersion   = ChildOrganization_Kind + "." + CRDGroupVersion.String()
	ChildOrganization_GroupVersionKind = CRDGroupVersion.WithKind(ChildOrganization_Kind)
)

func init() {
	SchemeBuilder.Register(&ChildOrganization{}, &ChildOrganizationList{})
}
