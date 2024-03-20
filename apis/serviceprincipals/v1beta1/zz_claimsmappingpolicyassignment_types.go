// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClaimsMappingPolicyAssignmentInitParameters struct {

	// The ID of the claims mapping policy to assign.
	// ID of the claims mapping policy to assign
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/policies/v1beta1.ClaimsMappingPolicy
	ClaimsMappingPolicyID *string `json:"claimsMappingPolicyId,omitempty" tf:"claims_mapping_policy_id,omitempty"`

	// Reference to a ClaimsMappingPolicy in policies to populate claimsMappingPolicyId.
	// +kubebuilder:validation:Optional
	ClaimsMappingPolicyIDRef *v1.Reference `json:"claimsMappingPolicyIdRef,omitempty" tf:"-"`

	// Selector for a ClaimsMappingPolicy in policies to populate claimsMappingPolicyId.
	// +kubebuilder:validation:Optional
	ClaimsMappingPolicyIDSelector *v1.Selector `json:"claimsMappingPolicyIdSelector,omitempty" tf:"-"`

	// The object ID of the service principal for the policy assignment.
	// Object ID of the service principal for which to assign the policy
	// +crossplane:generate:reference:type=Principal
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// Reference to a Principal to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDRef *v1.Reference `json:"servicePrincipalIdRef,omitempty" tf:"-"`

	// Selector for a Principal to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDSelector *v1.Selector `json:"servicePrincipalIdSelector,omitempty" tf:"-"`
}

type ClaimsMappingPolicyAssignmentObservation struct {

	// The ID of the claims mapping policy to assign.
	// ID of the claims mapping policy to assign
	ClaimsMappingPolicyID *string `json:"claimsMappingPolicyId,omitempty" tf:"claims_mapping_policy_id,omitempty"`

	// The ID of the Claims Mapping Policy Assignment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The object ID of the service principal for the policy assignment.
	// Object ID of the service principal for which to assign the policy
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`
}

type ClaimsMappingPolicyAssignmentParameters struct {

	// The ID of the claims mapping policy to assign.
	// ID of the claims mapping policy to assign
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/policies/v1beta1.ClaimsMappingPolicy
	// +kubebuilder:validation:Optional
	ClaimsMappingPolicyID *string `json:"claimsMappingPolicyId,omitempty" tf:"claims_mapping_policy_id,omitempty"`

	// Reference to a ClaimsMappingPolicy in policies to populate claimsMappingPolicyId.
	// +kubebuilder:validation:Optional
	ClaimsMappingPolicyIDRef *v1.Reference `json:"claimsMappingPolicyIdRef,omitempty" tf:"-"`

	// Selector for a ClaimsMappingPolicy in policies to populate claimsMappingPolicyId.
	// +kubebuilder:validation:Optional
	ClaimsMappingPolicyIDSelector *v1.Selector `json:"claimsMappingPolicyIdSelector,omitempty" tf:"-"`

	// The object ID of the service principal for the policy assignment.
	// Object ID of the service principal for which to assign the policy
	// +crossplane:generate:reference:type=Principal
	// +kubebuilder:validation:Optional
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// Reference to a Principal to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDRef *v1.Reference `json:"servicePrincipalIdRef,omitempty" tf:"-"`

	// Selector for a Principal to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDSelector *v1.Selector `json:"servicePrincipalIdSelector,omitempty" tf:"-"`
}

// ClaimsMappingPolicyAssignmentSpec defines the desired state of ClaimsMappingPolicyAssignment
type ClaimsMappingPolicyAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClaimsMappingPolicyAssignmentParameters `json:"forProvider"`
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
	InitProvider ClaimsMappingPolicyAssignmentInitParameters `json:"initProvider,omitempty"`
}

// ClaimsMappingPolicyAssignmentStatus defines the observed state of ClaimsMappingPolicyAssignment.
type ClaimsMappingPolicyAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClaimsMappingPolicyAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClaimsMappingPolicyAssignment is the Schema for the ClaimsMappingPolicyAssignments API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type ClaimsMappingPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClaimsMappingPolicyAssignmentSpec   `json:"spec"`
	Status            ClaimsMappingPolicyAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClaimsMappingPolicyAssignmentList contains a list of ClaimsMappingPolicyAssignments
type ClaimsMappingPolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClaimsMappingPolicyAssignment `json:"items"`
}

// Repository type metadata.
var (
	ClaimsMappingPolicyAssignment_Kind             = "ClaimsMappingPolicyAssignment"
	ClaimsMappingPolicyAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClaimsMappingPolicyAssignment_Kind}.String()
	ClaimsMappingPolicyAssignment_KindAPIVersion   = ClaimsMappingPolicyAssignment_Kind + "." + CRDGroupVersion.String()
	ClaimsMappingPolicyAssignment_GroupVersionKind = CRDGroupVersion.WithKind(ClaimsMappingPolicyAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&ClaimsMappingPolicyAssignment{}, &ClaimsMappingPolicyAssignmentList{})
}
