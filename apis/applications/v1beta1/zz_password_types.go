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

type PasswordInitParameters struct {

	// The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
	// The resource ID of the application for which this password should be created
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// The object ID of the application for which this password should be created
	// +crossplane:generate:reference:type=Application
	ApplicationObjectID *string `json:"applicationObjectId,omitempty" tf:"application_object_id,omitempty"`

	// Reference to a Application to populate applicationObjectId.
	// +kubebuilder:validation:Optional
	ApplicationObjectIDRef *v1.Reference `json:"applicationObjectIdRef,omitempty" tf:"-"`

	// Selector for a Application to populate applicationObjectId.
	// +kubebuilder:validation:Optional
	ApplicationObjectIDSelector *v1.Selector `json:"applicationObjectIdSelector,omitempty" tf:"-"`

	// A display name for the password. Changing this field forces a new resource to be created.
	// A display name for the password
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). Changing this field forces a new resource to be created.
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`)
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// A relative duration for which the password is valid until, for example 240h (10 days) or 2400h30m. Changing this field forces a new resource to be created.
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created
	EndDateRelative *string `json:"endDateRelative,omitempty" tf:"end_date_relative,omitempty"`

	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	// Arbitrary map of values that, when changed, will trigger rotation of the password
	// +mapType=granular
	RotateWhenChanged map[string]*string `json:"rotateWhenChanged,omitempty" tf:"rotate_when_changed,omitempty"`

	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

type PasswordObservation struct {

	// The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
	// The resource ID of the application for which this password should be created
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// The object ID of the application for which this password should be created
	ApplicationObjectID *string `json:"applicationObjectId,omitempty" tf:"application_object_id,omitempty"`

	// A display name for the password. Changing this field forces a new resource to be created.
	// A display name for the password
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). Changing this field forces a new resource to be created.
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`)
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// A relative duration for which the password is valid until, for example 240h (10 days) or 2400h30m. Changing this field forces a new resource to be created.
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created
	EndDateRelative *string `json:"endDateRelative,omitempty" tf:"end_date_relative,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A UUID used to uniquely identify this password credential.
	// A UUID used to uniquely identify this password credential
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	// Arbitrary map of values that, when changed, will trigger rotation of the password
	// +mapType=granular
	RotateWhenChanged map[string]*string `json:"rotateWhenChanged,omitempty" tf:"rotate_when_changed,omitempty"`

	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

type PasswordParameters struct {

	// The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
	// The resource ID of the application for which this password should be created
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// The object ID of the application for which this password should be created
	// +crossplane:generate:reference:type=Application
	// +kubebuilder:validation:Optional
	ApplicationObjectID *string `json:"applicationObjectId,omitempty" tf:"application_object_id,omitempty"`

	// Reference to a Application to populate applicationObjectId.
	// +kubebuilder:validation:Optional
	ApplicationObjectIDRef *v1.Reference `json:"applicationObjectIdRef,omitempty" tf:"-"`

	// Selector for a Application to populate applicationObjectId.
	// +kubebuilder:validation:Optional
	ApplicationObjectIDSelector *v1.Selector `json:"applicationObjectIdSelector,omitempty" tf:"-"`

	// A display name for the password. Changing this field forces a new resource to be created.
	// A display name for the password
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). Changing this field forces a new resource to be created.
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`)
	// +kubebuilder:validation:Optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// A relative duration for which the password is valid until, for example 240h (10 days) or 2400h30m. Changing this field forces a new resource to be created.
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created
	// +kubebuilder:validation:Optional
	EndDateRelative *string `json:"endDateRelative,omitempty" tf:"end_date_relative,omitempty"`

	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	// Arbitrary map of values that, when changed, will trigger rotation of the password
	// +kubebuilder:validation:Optional
	// +mapType=granular
	RotateWhenChanged map[string]*string `json:"rotateWhenChanged,omitempty" tf:"rotate_when_changed,omitempty"`

	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used
	// +kubebuilder:validation:Optional
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

// PasswordSpec defines the desired state of Password
type PasswordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PasswordParameters `json:"forProvider"`
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
	InitProvider PasswordInitParameters `json:"initProvider,omitempty"`
}

// PasswordStatus defines the observed state of Password.
type PasswordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PasswordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Password is the Schema for the Passwords API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type Password struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PasswordSpec   `json:"spec"`
	Status            PasswordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PasswordList contains a list of Passwords
type PasswordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Password `json:"items"`
}

// Repository type metadata.
var (
	Password_Kind             = "Password"
	Password_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Password_Kind}.String()
	Password_KindAPIVersion   = Password_Kind + "." + CRDGroupVersion.String()
	Password_GroupVersionKind = CRDGroupVersion.WithKind(Password_Kind)
)

func init() {
	SchemeBuilder.Register(&Password{}, &PasswordList{})
}
