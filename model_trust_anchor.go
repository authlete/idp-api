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

// TrustAnchor struct for TrustAnchor
type TrustAnchor struct {
	EntityId *string `json:"entityId,omitempty"`
	Jwks *string `json:"jwks,omitempty"`
}

// NewTrustAnchor instantiates a new TrustAnchor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustAnchor() *TrustAnchor {
	this := TrustAnchor{}
	return &this
}

// NewTrustAnchorWithDefaults instantiates a new TrustAnchor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustAnchorWithDefaults() *TrustAnchor {
	this := TrustAnchor{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *TrustAnchor) GetEntityId() string {
	if o == nil || isNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustAnchor) GetEntityIdOk() (*string, bool) {
	if o == nil || isNil(o.EntityId) {
    return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *TrustAnchor) HasEntityId() bool {
	if o != nil && !isNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *TrustAnchor) SetEntityId(v string) {
	o.EntityId = &v
}

// GetJwks returns the Jwks field value if set, zero value otherwise.
func (o *TrustAnchor) GetJwks() string {
	if o == nil || isNil(o.Jwks) {
		var ret string
		return ret
	}
	return *o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustAnchor) GetJwksOk() (*string, bool) {
	if o == nil || isNil(o.Jwks) {
    return nil, false
	}
	return o.Jwks, true
}

// HasJwks returns a boolean if a field has been set.
func (o *TrustAnchor) HasJwks() bool {
	if o != nil && !isNil(o.Jwks) {
		return true
	}

	return false
}

// SetJwks gets a reference to the given string and assigns it to the Jwks field.
func (o *TrustAnchor) SetJwks(v string) {
	o.Jwks = &v
}

func (o TrustAnchor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !isNil(o.Jwks) {
		toSerialize["jwks"] = o.Jwks
	}
	return json.Marshal(toSerialize)
}

type NullableTrustAnchor struct {
	value *TrustAnchor
	isSet bool
}

func (v NullableTrustAnchor) Get() *TrustAnchor {
	return v.value
}

func (v *NullableTrustAnchor) Set(val *TrustAnchor) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustAnchor) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustAnchor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustAnchor(val *TrustAnchor) *NullableTrustAnchor {
	return &NullableTrustAnchor{value: val, isSet: true}
}

func (v NullableTrustAnchor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustAnchor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


