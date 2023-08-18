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

// GrantScope struct for GrantScope
type GrantScope struct {
	Scope *string `json:"scope,omitempty"`
	Resource []string `json:"resource,omitempty"`
}

// NewGrantScope instantiates a new GrantScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantScope() *GrantScope {
	this := GrantScope{}
	return &this
}

// NewGrantScopeWithDefaults instantiates a new GrantScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantScopeWithDefaults() *GrantScope {
	this := GrantScope{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *GrantScope) GetScope() string {
	if o == nil || isNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantScope) GetScopeOk() (*string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GrantScope) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *GrantScope) SetScope(v string) {
	o.Scope = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *GrantScope) GetResource() []string {
	if o == nil || isNil(o.Resource) {
		var ret []string
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantScope) GetResourceOk() ([]string, bool) {
	if o == nil || isNil(o.Resource) {
    return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *GrantScope) HasResource() bool {
	if o != nil && !isNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given []string and assigns it to the Resource field.
func (o *GrantScope) SetResource(v []string) {
	o.Resource = v
}

func (o GrantScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !isNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableGrantScope struct {
	value *GrantScope
	isSet bool
}

func (v NullableGrantScope) Get() *GrantScope {
	return v.value
}

func (v *NullableGrantScope) Set(val *GrantScope) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantScope) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantScope(val *GrantScope) *NullableGrantScope {
	return &NullableGrantScope{value: val, isSet: true}
}

func (v NullableGrantScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

