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

// SnsCredentials struct for SnsCredentials
type SnsCredentials struct {
	Sns *string `json:"sns,omitempty"`
	ApiKey *string `json:"apiKey,omitempty"`
	ApiSecret *string `json:"apiSecret,omitempty"`
}

// NewSnsCredentials instantiates a new SnsCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnsCredentials() *SnsCredentials {
	this := SnsCredentials{}
	return &this
}

// NewSnsCredentialsWithDefaults instantiates a new SnsCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnsCredentialsWithDefaults() *SnsCredentials {
	this := SnsCredentials{}
	return &this
}

// GetSns returns the Sns field value if set, zero value otherwise.
func (o *SnsCredentials) GetSns() string {
	if o == nil || isNil(o.Sns) {
		var ret string
		return ret
	}
	return *o.Sns
}

// GetSnsOk returns a tuple with the Sns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnsCredentials) GetSnsOk() (*string, bool) {
	if o == nil || isNil(o.Sns) {
    return nil, false
	}
	return o.Sns, true
}

// HasSns returns a boolean if a field has been set.
func (o *SnsCredentials) HasSns() bool {
	if o != nil && !isNil(o.Sns) {
		return true
	}

	return false
}

// SetSns gets a reference to the given string and assigns it to the Sns field.
func (o *SnsCredentials) SetSns(v string) {
	o.Sns = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *SnsCredentials) GetApiKey() string {
	if o == nil || isNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnsCredentials) GetApiKeyOk() (*string, bool) {
	if o == nil || isNil(o.ApiKey) {
    return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *SnsCredentials) HasApiKey() bool {
	if o != nil && !isNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *SnsCredentials) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetApiSecret returns the ApiSecret field value if set, zero value otherwise.
func (o *SnsCredentials) GetApiSecret() string {
	if o == nil || isNil(o.ApiSecret) {
		var ret string
		return ret
	}
	return *o.ApiSecret
}

// GetApiSecretOk returns a tuple with the ApiSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnsCredentials) GetApiSecretOk() (*string, bool) {
	if o == nil || isNil(o.ApiSecret) {
    return nil, false
	}
	return o.ApiSecret, true
}

// HasApiSecret returns a boolean if a field has been set.
func (o *SnsCredentials) HasApiSecret() bool {
	if o != nil && !isNil(o.ApiSecret) {
		return true
	}

	return false
}

// SetApiSecret gets a reference to the given string and assigns it to the ApiSecret field.
func (o *SnsCredentials) SetApiSecret(v string) {
	o.ApiSecret = &v
}

func (o SnsCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sns) {
		toSerialize["sns"] = o.Sns
	}
	if !isNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !isNil(o.ApiSecret) {
		toSerialize["apiSecret"] = o.ApiSecret
	}
	return json.Marshal(toSerialize)
}

type NullableSnsCredentials struct {
	value *SnsCredentials
	isSet bool
}

func (v NullableSnsCredentials) Get() *SnsCredentials {
	return v.value
}

func (v *NullableSnsCredentials) Set(val *SnsCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableSnsCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableSnsCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnsCredentials(val *SnsCredentials) *NullableSnsCredentials {
	return &NullableSnsCredentials{value: val, isSet: true}
}

func (v NullableSnsCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnsCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

