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

// CreateApiServerRequest struct for CreateApiServerRequest
type CreateApiServerRequest struct {
	ApiServerUrl string `json:"apiServerUrl"`
	Description string `json:"description"`
	OwnedBy *int64 `json:"ownedBy,omitempty"`
}

// NewCreateApiServerRequest instantiates a new CreateApiServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiServerRequest(apiServerUrl string, description string) *CreateApiServerRequest {
	this := CreateApiServerRequest{}
	this.ApiServerUrl = apiServerUrl
	this.Description = description
	return &this
}

// NewCreateApiServerRequestWithDefaults instantiates a new CreateApiServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiServerRequestWithDefaults() *CreateApiServerRequest {
	this := CreateApiServerRequest{}
	return &this
}

// GetApiServerUrl returns the ApiServerUrl field value
func (o *CreateApiServerRequest) GetApiServerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiServerUrl
}

// GetApiServerUrlOk returns a tuple with the ApiServerUrl field value
// and a boolean to check if the value has been set.
func (o *CreateApiServerRequest) GetApiServerUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApiServerUrl, true
}

// SetApiServerUrl sets field value
func (o *CreateApiServerRequest) SetApiServerUrl(v string) {
	o.ApiServerUrl = v
}

// GetDescription returns the Description field value
func (o *CreateApiServerRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateApiServerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateApiServerRequest) SetDescription(v string) {
	o.Description = v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *CreateApiServerRequest) GetOwnedBy() int64 {
	if o == nil || isNil(o.OwnedBy) {
		var ret int64
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiServerRequest) GetOwnedByOk() (*int64, bool) {
	if o == nil || isNil(o.OwnedBy) {
    return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *CreateApiServerRequest) HasOwnedBy() bool {
	if o != nil && !isNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given int64 and assigns it to the OwnedBy field.
func (o *CreateApiServerRequest) SetOwnedBy(v int64) {
	o.OwnedBy = &v
}

func (o CreateApiServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiServerUrl"] = o.ApiServerUrl
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.OwnedBy) {
		toSerialize["ownedBy"] = o.OwnedBy
	}
	return json.Marshal(toSerialize)
}

type NullableCreateApiServerRequest struct {
	value *CreateApiServerRequest
	isSet bool
}

func (v NullableCreateApiServerRequest) Get() *CreateApiServerRequest {
	return v.value
}

func (v *NullableCreateApiServerRequest) Set(val *CreateApiServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiServerRequest(val *CreateApiServerRequest) *NullableCreateApiServerRequest {
	return &NullableCreateApiServerRequest{value: val, isSet: true}
}

func (v NullableCreateApiServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

