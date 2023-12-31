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

// GetOrganizationTokensRequest struct for GetOrganizationTokensRequest
type GetOrganizationTokensRequest struct {
	OrganizationId int64 `json:"organizationId"`
}

// NewGetOrganizationTokensRequest instantiates a new GetOrganizationTokensRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationTokensRequest(organizationId int64) *GetOrganizationTokensRequest {
	this := GetOrganizationTokensRequest{}
	this.OrganizationId = organizationId
	return &this
}

// NewGetOrganizationTokensRequestWithDefaults instantiates a new GetOrganizationTokensRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationTokensRequestWithDefaults() *GetOrganizationTokensRequest {
	this := GetOrganizationTokensRequest{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value
func (o *GetOrganizationTokensRequest) GetOrganizationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationTokensRequest) GetOrganizationIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *GetOrganizationTokensRequest) SetOrganizationId(v int64) {
	o.OrganizationId = v
}

func (o GetOrganizationTokensRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["organizationId"] = o.OrganizationId
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationTokensRequest struct {
	value *GetOrganizationTokensRequest
	isSet bool
}

func (v NullableGetOrganizationTokensRequest) Get() *GetOrganizationTokensRequest {
	return v.value
}

func (v *NullableGetOrganizationTokensRequest) Set(val *GetOrganizationTokensRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationTokensRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationTokensRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationTokensRequest(val *GetOrganizationTokensRequest) *NullableGetOrganizationTokensRequest {
	return &NullableGetOrganizationTokensRequest{value: val, isSet: true}
}

func (v NullableGetOrganizationTokensRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationTokensRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


