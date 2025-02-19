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

// checks if the GoogleAuthenticatorConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleAuthenticatorConfig{}

// GoogleAuthenticatorConfig struct for GoogleAuthenticatorConfig
type GoogleAuthenticatorConfig struct {
	TimeStepSizeInMillis *int64  `json:"timeStepSizeInMillis,omitempty"`
	WindowSize           *int32  `json:"windowSize,omitempty"`
	CodeDigits           *int32  `json:"codeDigits,omitempty"`
	NumberOfScratchCodes *int32  `json:"numberOfScratchCodes,omitempty"`
	KeyModulus           *int32  `json:"keyModulus,omitempty"`
	SecretBits           *int32  `json:"secretBits,omitempty"`
	KeyRepresentation    *string `json:"keyRepresentation,omitempty"`
	HmacHashFunction     *string `json:"hmacHashFunction,omitempty"`
}

// NewGoogleAuthenticatorConfig instantiates a new GoogleAuthenticatorConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleAuthenticatorConfig() *GoogleAuthenticatorConfig {
	this := GoogleAuthenticatorConfig{}
	return &this
}

// NewGoogleAuthenticatorConfigWithDefaults instantiates a new GoogleAuthenticatorConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleAuthenticatorConfigWithDefaults() *GoogleAuthenticatorConfig {
	this := GoogleAuthenticatorConfig{}
	return &this
}

// GetTimeStepSizeInMillis returns the TimeStepSizeInMillis field value if set, zero value otherwise.
func (o *GoogleAuthenticatorConfig) GetTimeStepSizeInMillis() int64 {
	if o == nil || IsNil(o.TimeStepSizeInMillis) {
		var ret int64
		return ret
	}
	return *o.TimeStepSizeInMillis
}

// GetTimeStepSizeInMillisOk returns a tuple with the TimeStepSizeInMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorConfig) GetTimeStepSizeInMillisOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeStepSizeInMillis) {
		return nil, false
	}
	return o.TimeStepSizeInMillis, true
}

// HasTimeStepSizeInMillis returns a boolean if a field has been set.
func (o *GoogleAuthenticatorConfig) HasTimeStepSizeInMillis() bool {
	if o != nil && !IsNil(o.TimeStepSizeInMillis) {
		return true
	}

	return false
}

// SetTimeStepSizeInMillis gets a reference to the given int64 and assigns it to the TimeStepSizeInMillis field.
func (o *GoogleAuthenticatorConfig) SetTimeStepSizeInMillis(v int64) {
	o.TimeStepSizeInMillis = &v
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *GoogleAuthenticatorConfig) GetWindowSize() int32 {
	if o == nil || IsNil(o.WindowSize) {
		var ret int32
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorConfig) GetWindowSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.WindowSize) {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *GoogleAuthenticatorConfig) HasWindowSize() bool {
	if o != nil && !IsNil(o.WindowSize) {
		return true
	}

	return false
}

// SetWindowSize gets a reference to the given int32 and assigns it to the WindowSize field.
func (o *GoogleAuthenticatorConfig) SetWindowSize(v int32) {
	o.WindowSize = &v
}

// GetCodeDigits returns the CodeDigits field value if set, zero value otherwise.
func (o *GoogleAuthenticatorConfig) GetCodeDigits() int32 {
	if o == nil || IsNil(o.CodeDigits) {
		var ret int32
		return ret
	}
	return *o.CodeDigits
}

// GetCodeDigitsOk returns a tuple with the CodeDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorConfig) GetCodeDigitsOk() (*int32, bool) {
	if o == nil || IsNil(o.CodeDigits) {
		return nil, false
	}
	return o.CodeDigits, true
}

// HasCodeDigits returns a boolean if a field has been set.
func (o *GoogleAuthenticatorConfig) HasCodeDigits() bool {
	if o != nil && !IsNil(o.CodeDigits) {
		return true
	}

	return false
}

// SetCodeDigits gets a reference to the given int32 and assigns it to the CodeDigits field.
func (o *GoogleAuthenticatorConfig) SetCodeDigits(v int32) {
	o.CodeDigits = &v
}

// GetNumberOfScratchCodes returns the NumberOfScratchCodes field value if set, zero value otherwise.
func (o *GoogleAuthenticatorConfig) GetNumberOfScratchCodes() int32 {
	if o == nil || IsNil(o.NumberOfScratchCodes) {
		var ret int32
		return ret
	}
	return *o.NumberOfScratchCodes
}

// GetNumberOfScratchCodesOk returns a tuple with the NumberOfScratchCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorConfig) GetNumberOfScratchCodesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfScratchCodes) {
		return nil, false
	}
	return o.NumberOfScratchCodes, true
}

// HasNumberOfScratchCodes returns a boolean if a field has been set.
func (o *GoogleAuthenticatorConfig) HasNumberOfScratchCodes() bool {
	if o != nil && !IsNil(o.NumberOfScratchCodes) {
		return true
	}

	return false
}

// SetNumberOfScratchCodes gets a reference to the given int32 and assigns it to the NumberOfScratchCodes field.
func (o *GoogleAuthenticatorConfig) SetNumberOfScratchCodes(v int32) {
	o.NumberOfScratchCodes = &v
}

// GetKeyModulus returns the KeyModulus field value if set, zero value otherwise.
func (o *GoogleAuthenticatorConfig) GetKeyModulus() int32 {
	if o == nil || IsNil(o.KeyModulus) {
		var ret int32
		return ret
	}
	return *o.KeyModulus
}

// GetKeyModulusOk returns a tuple with the KeyModulus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorConfig) GetKeyModulusOk() (*int32, bool) {
	if o == nil || IsNil(o.KeyModulus) {
		return nil, false
	}
	return o.KeyModulus, true
}

// HasKeyModulus returns a boolean if a field has been set.
func (o *GoogleAuthenticatorConfig) HasKeyModulus() bool {
	if o != nil && !IsNil(o.KeyModulus) {
		return true
	}

	return false
}

// SetKeyModulus gets a reference to the given int32 and assigns it to the KeyModulus field.
func (o *GoogleAuthenticatorConfig) SetKeyModulus(v int32) {
	o.KeyModulus = &v
}

// GetSecretBits returns the SecretBits field value if set, zero value otherwise.
func (o *GoogleAuthenticatorConfig) GetSecretBits() int32 {
	if o == nil || IsNil(o.SecretBits) {
		var ret int32
		return ret
	}
	return *o.SecretBits
}

// GetSecretBitsOk returns a tuple with the SecretBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorConfig) GetSecretBitsOk() (*int32, bool) {
	if o == nil || IsNil(o.SecretBits) {
		return nil, false
	}
	return o.SecretBits, true
}

// HasSecretBits returns a boolean if a field has been set.
func (o *GoogleAuthenticatorConfig) HasSecretBits() bool {
	if o != nil && !IsNil(o.SecretBits) {
		return true
	}

	return false
}

// SetSecretBits gets a reference to the given int32 and assigns it to the SecretBits field.
func (o *GoogleAuthenticatorConfig) SetSecretBits(v int32) {
	o.SecretBits = &v
}

// GetKeyRepresentation returns the KeyRepresentation field value if set, zero value otherwise.
func (o *GoogleAuthenticatorConfig) GetKeyRepresentation() string {
	if o == nil || IsNil(o.KeyRepresentation) {
		var ret string
		return ret
	}
	return *o.KeyRepresentation
}

// GetKeyRepresentationOk returns a tuple with the KeyRepresentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorConfig) GetKeyRepresentationOk() (*string, bool) {
	if o == nil || IsNil(o.KeyRepresentation) {
		return nil, false
	}
	return o.KeyRepresentation, true
}

// HasKeyRepresentation returns a boolean if a field has been set.
func (o *GoogleAuthenticatorConfig) HasKeyRepresentation() bool {
	if o != nil && !IsNil(o.KeyRepresentation) {
		return true
	}

	return false
}

// SetKeyRepresentation gets a reference to the given string and assigns it to the KeyRepresentation field.
func (o *GoogleAuthenticatorConfig) SetKeyRepresentation(v string) {
	o.KeyRepresentation = &v
}

// GetHmacHashFunction returns the HmacHashFunction field value if set, zero value otherwise.
func (o *GoogleAuthenticatorConfig) GetHmacHashFunction() string {
	if o == nil || IsNil(o.HmacHashFunction) {
		var ret string
		return ret
	}
	return *o.HmacHashFunction
}

// GetHmacHashFunctionOk returns a tuple with the HmacHashFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleAuthenticatorConfig) GetHmacHashFunctionOk() (*string, bool) {
	if o == nil || IsNil(o.HmacHashFunction) {
		return nil, false
	}
	return o.HmacHashFunction, true
}

// HasHmacHashFunction returns a boolean if a field has been set.
func (o *GoogleAuthenticatorConfig) HasHmacHashFunction() bool {
	if o != nil && !IsNil(o.HmacHashFunction) {
		return true
	}

	return false
}

// SetHmacHashFunction gets a reference to the given string and assigns it to the HmacHashFunction field.
func (o *GoogleAuthenticatorConfig) SetHmacHashFunction(v string) {
	o.HmacHashFunction = &v
}

func (o GoogleAuthenticatorConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleAuthenticatorConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeStepSizeInMillis) {
		toSerialize["timeStepSizeInMillis"] = o.TimeStepSizeInMillis
	}
	if !IsNil(o.WindowSize) {
		toSerialize["windowSize"] = o.WindowSize
	}
	if !IsNil(o.CodeDigits) {
		toSerialize["codeDigits"] = o.CodeDigits
	}
	if !IsNil(o.NumberOfScratchCodes) {
		toSerialize["numberOfScratchCodes"] = o.NumberOfScratchCodes
	}
	if !IsNil(o.KeyModulus) {
		toSerialize["keyModulus"] = o.KeyModulus
	}
	if !IsNil(o.SecretBits) {
		toSerialize["secretBits"] = o.SecretBits
	}
	if !IsNil(o.KeyRepresentation) {
		toSerialize["keyRepresentation"] = o.KeyRepresentation
	}
	if !IsNil(o.HmacHashFunction) {
		toSerialize["hmacHashFunction"] = o.HmacHashFunction
	}
	return toSerialize, nil
}

type NullableGoogleAuthenticatorConfig struct {
	value *GoogleAuthenticatorConfig
	isSet bool
}

func (v NullableGoogleAuthenticatorConfig) Get() *GoogleAuthenticatorConfig {
	return v.value
}

func (v *NullableGoogleAuthenticatorConfig) Set(val *GoogleAuthenticatorConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleAuthenticatorConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleAuthenticatorConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleAuthenticatorConfig(val *GoogleAuthenticatorConfig) *NullableGoogleAuthenticatorConfig {
	return &NullableGoogleAuthenticatorConfig{value: val, isSet: true}
}

func (v NullableGoogleAuthenticatorConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleAuthenticatorConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
