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

// BindTotpCredentialsRequest struct for BindTotpCredentialsRequest
type BindTotpCredentialsRequest struct {
	Totp int32 `json:"totp"`
}

// NewBindTotpCredentialsRequest instantiates a new BindTotpCredentialsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindTotpCredentialsRequest(totp int32) *BindTotpCredentialsRequest {
	this := BindTotpCredentialsRequest{}
	this.Totp = totp
	return &this
}

// NewBindTotpCredentialsRequestWithDefaults instantiates a new BindTotpCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindTotpCredentialsRequestWithDefaults() *BindTotpCredentialsRequest {
	this := BindTotpCredentialsRequest{}
	return &this
}

// GetTotp returns the Totp field value
func (o *BindTotpCredentialsRequest) GetTotp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Totp
}

// GetTotpOk returns a tuple with the Totp field value
// and a boolean to check if the value has been set.
func (o *BindTotpCredentialsRequest) GetTotpOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Totp, true
}

// SetTotp sets field value
func (o *BindTotpCredentialsRequest) SetTotp(v int32) {
	o.Totp = v
}

func (o BindTotpCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totp"] = o.Totp
	}
	return json.Marshal(toSerialize)
}

type NullableBindTotpCredentialsRequest struct {
	value *BindTotpCredentialsRequest
	isSet bool
}

func (v NullableBindTotpCredentialsRequest) Get() *BindTotpCredentialsRequest {
	return v.value
}

func (v *NullableBindTotpCredentialsRequest) Set(val *BindTotpCredentialsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBindTotpCredentialsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBindTotpCredentialsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindTotpCredentialsRequest(val *BindTotpCredentialsRequest) *NullableBindTotpCredentialsRequest {
	return &NullableBindTotpCredentialsRequest{value: val, isSet: true}
}

func (v NullableBindTotpCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindTotpCredentialsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

