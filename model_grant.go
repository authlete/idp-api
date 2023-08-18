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

// Grant struct for Grant
type Grant struct {
	Scopes []GrantScope `json:"scopes,omitempty"`
	Claims []string `json:"claims,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
}

// NewGrant instantiates a new Grant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrant() *Grant {
	this := Grant{}
	return &this
}

// NewGrantWithDefaults instantiates a new Grant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantWithDefaults() *Grant {
	this := Grant{}
	return &this
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *Grant) GetScopes() []GrantScope {
	if o == nil || isNil(o.Scopes) {
		var ret []GrantScope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetScopesOk() ([]GrantScope, bool) {
	if o == nil || isNil(o.Scopes) {
    return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *Grant) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []GrantScope and assigns it to the Scopes field.
func (o *Grant) SetScopes(v []GrantScope) {
	o.Scopes = v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *Grant) GetClaims() []string {
	if o == nil || isNil(o.Claims) {
		var ret []string
		return ret
	}
	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetClaimsOk() ([]string, bool) {
	if o == nil || isNil(o.Claims) {
    return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *Grant) HasClaims() bool {
	if o != nil && !isNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given []string and assigns it to the Claims field.
func (o *Grant) SetClaims(v []string) {
	o.Claims = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *Grant) GetAuthorizationDetails() AuthzDetails {
	if o == nil || isNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || isNil(o.AuthorizationDetails) {
    return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *Grant) HasAuthorizationDetails() bool {
	if o != nil && !isNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *Grant) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

func (o Grant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !isNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !isNil(o.AuthorizationDetails) {
		toSerialize["authorizationDetails"] = o.AuthorizationDetails
	}
	return json.Marshal(toSerialize)
}

type NullableGrant struct {
	value *Grant
	isSet bool
}

func (v NullableGrant) Get() *Grant {
	return v.value
}

func (v *NullableGrant) Set(val *Grant) {
	v.value = val
	v.isSet = true
}

func (v NullableGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrant(val *Grant) *NullableGrant {
	return &NullableGrant{value: val, isSet: true}
}

func (v NullableGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


