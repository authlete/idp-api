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

// GetServiceTokensRequest struct for GetServiceTokensRequest
type GetServiceTokensRequest struct {
	ServiceId int64 `json:"serviceId"`
	OrganizationId int64 `json:"organizationId"`
	ApiServerId int64 `json:"apiServerId"`
}

// NewGetServiceTokensRequest instantiates a new GetServiceTokensRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServiceTokensRequest(serviceId int64, organizationId int64, apiServerId int64) *GetServiceTokensRequest {
	this := GetServiceTokensRequest{}
	this.ServiceId = serviceId
	this.OrganizationId = organizationId
	this.ApiServerId = apiServerId
	return &this
}

// NewGetServiceTokensRequestWithDefaults instantiates a new GetServiceTokensRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServiceTokensRequestWithDefaults() *GetServiceTokensRequest {
	this := GetServiceTokensRequest{}
	return &this
}

// GetServiceId returns the ServiceId field value
func (o *GetServiceTokensRequest) GetServiceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *GetServiceTokensRequest) GetServiceIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *GetServiceTokensRequest) SetServiceId(v int64) {
	o.ServiceId = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *GetServiceTokensRequest) GetOrganizationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *GetServiceTokensRequest) GetOrganizationIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *GetServiceTokensRequest) SetOrganizationId(v int64) {
	o.OrganizationId = v
}

// GetApiServerId returns the ApiServerId field value
func (o *GetServiceTokensRequest) GetApiServerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApiServerId
}

// GetApiServerIdOk returns a tuple with the ApiServerId field value
// and a boolean to check if the value has been set.
func (o *GetServiceTokensRequest) GetApiServerIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApiServerId, true
}

// SetApiServerId sets field value
func (o *GetServiceTokensRequest) SetApiServerId(v int64) {
	o.ApiServerId = v
}

func (o GetServiceTokensRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serviceId"] = o.ServiceId
	}
	if true {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if true {
		toSerialize["apiServerId"] = o.ApiServerId
	}
	return json.Marshal(toSerialize)
}

type NullableGetServiceTokensRequest struct {
	value *GetServiceTokensRequest
	isSet bool
}

func (v NullableGetServiceTokensRequest) Get() *GetServiceTokensRequest {
	return v.value
}

func (v *NullableGetServiceTokensRequest) Set(val *GetServiceTokensRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServiceTokensRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServiceTokensRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServiceTokensRequest(val *GetServiceTokensRequest) *NullableGetServiceTokensRequest {
	return &NullableGetServiceTokensRequest{value: val, isSet: true}
}

func (v NullableGetServiceTokensRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServiceTokensRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

