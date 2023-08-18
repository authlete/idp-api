/*
OpenAPI definition

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AuthzDetails struct for AuthzDetails
type AuthzDetails struct {
	Elements []AuthzDetailsElement `json:"elements,omitempty"`
}

// NewAuthzDetails instantiates a new AuthzDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthzDetails() *AuthzDetails {
	this := AuthzDetails{}
	return &this
}

// NewAuthzDetailsWithDefaults instantiates a new AuthzDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthzDetailsWithDefaults() *AuthzDetails {
	this := AuthzDetails{}
	return &this
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *AuthzDetails) GetElements() []AuthzDetailsElement {
	if o == nil || isNil(o.Elements) {
		var ret []AuthzDetailsElement
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthzDetails) GetElementsOk() ([]AuthzDetailsElement, bool) {
	if o == nil || isNil(o.Elements) {
    return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *AuthzDetails) HasElements() bool {
	if o != nil && !isNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []AuthzDetailsElement and assigns it to the Elements field.
func (o *AuthzDetails) SetElements(v []AuthzDetailsElement) {
	o.Elements = v
}

func (o AuthzDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Elements) {
		toSerialize["elements"] = o.Elements
	}
	return json.Marshal(toSerialize)
}

type NullableAuthzDetails struct {
	value *AuthzDetails
	isSet bool
}

func (v NullableAuthzDetails) Get() *AuthzDetails {
	return v.value
}

func (v *NullableAuthzDetails) Set(val *AuthzDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthzDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthzDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthzDetails(val *AuthzDetails) *NullableAuthzDetails {
	return &NullableAuthzDetails{value: val, isSet: true}
}

func (v NullableAuthzDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthzDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

