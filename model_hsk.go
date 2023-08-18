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

// Hsk struct for Hsk
type Hsk struct {
	Kty *string `json:"kty,omitempty"`
	Use *string `json:"use,omitempty"`
	Alg *string `json:"alg,omitempty"`
	Kid *string `json:"kid,omitempty"`
	HsmName *string `json:"hsmName,omitempty"`
	Handle *string `json:"handle,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
}

// NewHsk instantiates a new Hsk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHsk() *Hsk {
	this := Hsk{}
	return &this
}

// NewHskWithDefaults instantiates a new Hsk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHskWithDefaults() *Hsk {
	this := Hsk{}
	return &this
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *Hsk) GetKty() string {
	if o == nil || isNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hsk) GetKtyOk() (*string, bool) {
	if o == nil || isNil(o.Kty) {
    return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *Hsk) HasKty() bool {
	if o != nil && !isNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *Hsk) SetKty(v string) {
	o.Kty = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *Hsk) GetUse() string {
	if o == nil || isNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hsk) GetUseOk() (*string, bool) {
	if o == nil || isNil(o.Use) {
    return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *Hsk) HasUse() bool {
	if o != nil && !isNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *Hsk) SetUse(v string) {
	o.Use = &v
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *Hsk) GetAlg() string {
	if o == nil || isNil(o.Alg) {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hsk) GetAlgOk() (*string, bool) {
	if o == nil || isNil(o.Alg) {
    return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *Hsk) HasAlg() bool {
	if o != nil && !isNil(o.Alg) {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *Hsk) SetAlg(v string) {
	o.Alg = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *Hsk) GetKid() string {
	if o == nil || isNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hsk) GetKidOk() (*string, bool) {
	if o == nil || isNil(o.Kid) {
    return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *Hsk) HasKid() bool {
	if o != nil && !isNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *Hsk) SetKid(v string) {
	o.Kid = &v
}

// GetHsmName returns the HsmName field value if set, zero value otherwise.
func (o *Hsk) GetHsmName() string {
	if o == nil || isNil(o.HsmName) {
		var ret string
		return ret
	}
	return *o.HsmName
}

// GetHsmNameOk returns a tuple with the HsmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hsk) GetHsmNameOk() (*string, bool) {
	if o == nil || isNil(o.HsmName) {
    return nil, false
	}
	return o.HsmName, true
}

// HasHsmName returns a boolean if a field has been set.
func (o *Hsk) HasHsmName() bool {
	if o != nil && !isNil(o.HsmName) {
		return true
	}

	return false
}

// SetHsmName gets a reference to the given string and assigns it to the HsmName field.
func (o *Hsk) SetHsmName(v string) {
	o.HsmName = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *Hsk) GetHandle() string {
	if o == nil || isNil(o.Handle) {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hsk) GetHandleOk() (*string, bool) {
	if o == nil || isNil(o.Handle) {
    return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *Hsk) HasHandle() bool {
	if o != nil && !isNil(o.Handle) {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *Hsk) SetHandle(v string) {
	o.Handle = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *Hsk) GetPublicKey() string {
	if o == nil || isNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hsk) GetPublicKeyOk() (*string, bool) {
	if o == nil || isNil(o.PublicKey) {
    return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *Hsk) HasPublicKey() bool {
	if o != nil && !isNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *Hsk) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o Hsk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !isNil(o.Use) {
		toSerialize["use"] = o.Use
	}
	if !isNil(o.Alg) {
		toSerialize["alg"] = o.Alg
	}
	if !isNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}
	if !isNil(o.HsmName) {
		toSerialize["hsmName"] = o.HsmName
	}
	if !isNil(o.Handle) {
		toSerialize["handle"] = o.Handle
	}
	if !isNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableHsk struct {
	value *Hsk
	isSet bool
}

func (v NullableHsk) Get() *Hsk {
	return v.value
}

func (v *NullableHsk) Set(val *Hsk) {
	v.value = val
	v.isSet = true
}

func (v NullableHsk) IsSet() bool {
	return v.isSet
}

func (v *NullableHsk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHsk(val *Hsk) *NullableHsk {
	return &NullableHsk{value: val, isSet: true}
}

func (v NullableHsk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHsk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


