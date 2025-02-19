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

// checks if the GoogleAuthenticatorKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleAuthenticatorKey{}

// GoogleAuthenticatorKey struct for GoogleAuthenticatorKey
type GoogleAuthenticatorKey struct {
	Config           *GoogleAuthenticatorConfig `json:"config,omitempty"`
	Key              *string                    `json:"key,omitempty"`
	VerificationCode *int32                     `json:"verificationCode,omitempty"`
	ScratchCodes     []int32                    `json:"scratchCodes,omitempty"`
}

// NewGoogleAuthenticatorKey instantiates a new GoogleAuthenticatorKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleAuthenticatorKey() *GoogleAuthenticatorKey {
	this := GoogleAuthenticatorKey{}
	return &this
}

// NewGoogleAuthenticatorKeyWithDefaults instantiates a new GoogleAuthenticatorKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleAuthenticatorKeyWithDefaults() *GoogleAuthenticatorKey {
	this := GoogleAuthenticatorKey{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *GoogleAuthenticatorKey) GetConfig() GoogleAuthenticatorConfig {
	if o == nil || IsNil(o.Config) {
		var ret GoogleAuthenticatorConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorKey) GetConfigOk() (*GoogleAuthenticatorConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *GoogleAuthenticatorKey) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given GoogleAuthenticatorConfig and assigns it to the Config field.
func (o *GoogleAuthenticatorKey) SetConfig(v GoogleAuthenticatorConfig) {
	o.Config = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GoogleAuthenticatorKey) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorKey) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GoogleAuthenticatorKey) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GoogleAuthenticatorKey) SetKey(v string) {
	o.Key = &v
}

// GetVerificationCode returns the VerificationCode field value if set, zero value otherwise.
func (o *GoogleAuthenticatorKey) GetVerificationCode() int32 {
	if o == nil || IsNil(o.VerificationCode) {
		var ret int32
		return ret
	}
	return *o.VerificationCode
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorKey) GetVerificationCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.VerificationCode) {
		return nil, false
	}
	return o.VerificationCode, true
}

// HasVerificationCode returns a boolean if a field has been set.
func (o *GoogleAuthenticatorKey) HasVerificationCode() bool {
	if o != nil && !IsNil(o.VerificationCode) {
		return true
	}

	return false
}

// SetVerificationCode gets a reference to the given int32 and assigns it to the VerificationCode field.
func (o *GoogleAuthenticatorKey) SetVerificationCode(v int32) {
	o.VerificationCode = &v
}

// GetScratchCodes returns the ScratchCodes field value if set, zero value otherwise.
func (o *GoogleAuthenticatorKey) GetScratchCodes() []int32 {
	if o == nil || IsNil(o.ScratchCodes) {
		var ret []int32
		return ret
	}
	return o.ScratchCodes
}

// GetScratchCodesOk returns a tuple with the ScratchCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorKey) GetScratchCodesOk() ([]int32, bool) {
	if o == nil || IsNil(o.ScratchCodes) {
		return nil, false
	}
	return o.ScratchCodes, true
}

// HasScratchCodes returns a boolean if a field has been set.
func (o *GoogleAuthenticatorKey) HasScratchCodes() bool {
	if o != nil && !IsNil(o.ScratchCodes) {
		return true
	}

	return false
}

// SetScratchCodes gets a reference to the given []int32 and assigns it to the ScratchCodes field.
func (o *GoogleAuthenticatorKey) SetScratchCodes(v []int32) {
	o.ScratchCodes = v
}

func (o GoogleAuthenticatorKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleAuthenticatorKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.VerificationCode) {
		toSerialize["verificationCode"] = o.VerificationCode
	}
	if !IsNil(o.ScratchCodes) {
		toSerialize["scratchCodes"] = o.ScratchCodes
	}
	return toSerialize, nil
}

type NullableGoogleAuthenticatorKey struct {
	value *GoogleAuthenticatorKey
	isSet bool
}

func (v NullableGoogleAuthenticatorKey) Get() *GoogleAuthenticatorKey {
	return v.value
}

func (v *NullableGoogleAuthenticatorKey) Set(val *GoogleAuthenticatorKey) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleAuthenticatorKey) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleAuthenticatorKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleAuthenticatorKey(val *GoogleAuthenticatorKey) *NullableGoogleAuthenticatorKey {
	return &NullableGoogleAuthenticatorKey{value: val, isSet: true}
}

func (v NullableGoogleAuthenticatorKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleAuthenticatorKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
