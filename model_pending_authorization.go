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

// PendingAuthorization struct for PendingAuthorization
type PendingAuthorization struct {
	AuthorizationResponse *AuthorizationResponse `json:"authorizationResponse,omitempty"`
	Audience *string `json:"audience,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

// NewPendingAuthorization instantiates a new PendingAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingAuthorization() *PendingAuthorization {
	this := PendingAuthorization{}
	return &this
}

// NewPendingAuthorizationWithDefaults instantiates a new PendingAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingAuthorizationWithDefaults() *PendingAuthorization {
	this := PendingAuthorization{}
	return &this
}

// GetAuthorizationResponse returns the AuthorizationResponse field value if set, zero value otherwise.
func (o *PendingAuthorization) GetAuthorizationResponse() AuthorizationResponse {
	if o == nil || isNil(o.AuthorizationResponse) {
		var ret AuthorizationResponse
		return ret
	}
	return *o.AuthorizationResponse
}

// GetAuthorizationResponseOk returns a tuple with the AuthorizationResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingAuthorization) GetAuthorizationResponseOk() (*AuthorizationResponse, bool) {
	if o == nil || isNil(o.AuthorizationResponse) {
    return nil, false
	}
	return o.AuthorizationResponse, true
}

// HasAuthorizationResponse returns a boolean if a field has been set.
func (o *PendingAuthorization) HasAuthorizationResponse() bool {
	if o != nil && !isNil(o.AuthorizationResponse) {
		return true
	}

	return false
}

// SetAuthorizationResponse gets a reference to the given AuthorizationResponse and assigns it to the AuthorizationResponse field.
func (o *PendingAuthorization) SetAuthorizationResponse(v AuthorizationResponse) {
	o.AuthorizationResponse = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *PendingAuthorization) GetAudience() string {
	if o == nil || isNil(o.Audience) {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingAuthorization) GetAudienceOk() (*string, bool) {
	if o == nil || isNil(o.Audience) {
    return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *PendingAuthorization) HasAudience() bool {
	if o != nil && !isNil(o.Audience) {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *PendingAuthorization) SetAudience(v string) {
	o.Audience = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *PendingAuthorization) GetResource() string {
	if o == nil || isNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingAuthorization) GetResourceOk() (*string, bool) {
	if o == nil || isNil(o.Resource) {
    return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *PendingAuthorization) HasResource() bool {
	if o != nil && !isNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *PendingAuthorization) SetResource(v string) {
	o.Resource = &v
}

func (o PendingAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthorizationResponse) {
		toSerialize["authorizationResponse"] = o.AuthorizationResponse
	}
	if !isNil(o.Audience) {
		toSerialize["audience"] = o.Audience
	}
	if !isNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullablePendingAuthorization struct {
	value *PendingAuthorization
	isSet bool
}

func (v NullablePendingAuthorization) Get() *PendingAuthorization {
	return v.value
}

func (v *NullablePendingAuthorization) Set(val *PendingAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingAuthorization(val *PendingAuthorization) *NullablePendingAuthorization {
	return &NullablePendingAuthorization{value: val, isSet: true}
}

func (v NullablePendingAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


