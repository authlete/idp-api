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

// UpdateUserRequest struct for UpdateUserRequest
type UpdateUserRequest struct {
	Email string `json:"email"`
	GivenName *string `json:"givenName,omitempty"`
	FamilyName *string `json:"familyName,omitempty"`
}

// NewUpdateUserRequest instantiates a new UpdateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequest(email string) *UpdateUserRequest {
	this := UpdateUserRequest{}
	this.Email = email
	return &this
}

// NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestWithDefaults() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *UpdateUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UpdateUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetGivenName() string {
	if o == nil || isNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetGivenNameOk() (*string, bool) {
	if o == nil || isNil(o.GivenName) {
    return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasGivenName() bool {
	if o != nil && !isNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *UpdateUserRequest) SetGivenName(v string) {
	o.GivenName = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetFamilyName() string {
	if o == nil || isNil(o.FamilyName) {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetFamilyNameOk() (*string, bool) {
	if o == nil || isNil(o.FamilyName) {
    return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasFamilyName() bool {
	if o != nil && !isNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *UpdateUserRequest) SetFamilyName(v string) {
	o.FamilyName = &v
}

func (o UpdateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !isNil(o.FamilyName) {
		toSerialize["familyName"] = o.FamilyName
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateUserRequest struct {
	value *UpdateUserRequest
	isSet bool
}

func (v NullableUpdateUserRequest) Get() *UpdateUserRequest {
	return v.value
}

func (v *NullableUpdateUserRequest) Set(val *UpdateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRequest(val *UpdateUserRequest) *NullableUpdateUserRequest {
	return &NullableUpdateUserRequest{value: val, isSet: true}
}

func (v NullableUpdateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

