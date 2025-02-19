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

// checks if the ServiceTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTokenResponse{}

// ServiceTokenResponse struct for ServiceTokenResponse
type ServiceTokenResponse struct {
	ServiceId      *int64  `json:"serviceId,omitempty"`
	OrganizationId *int64  `json:"organizationId,omitempty"`
	ApiServerId    *int64  `json:"apiServerId,omitempty"`
	AccessToken    *string `json:"accessToken,omitempty"`
	TokenId        *string `json:"tokenId,omitempty"`
	Description    *string `json:"description,omitempty"`
}

// NewServiceTokenResponse instantiates a new ServiceTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTokenResponse() *ServiceTokenResponse {
	this := ServiceTokenResponse{}
	return &this
}

// NewServiceTokenResponseWithDefaults instantiates a new ServiceTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTokenResponseWithDefaults() *ServiceTokenResponse {
	this := ServiceTokenResponse{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ServiceTokenResponse) GetServiceId() int64 {
	if o == nil || IsNil(o.ServiceId) {
		var ret int64
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTokenResponse) GetServiceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ServiceTokenResponse) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given int64 and assigns it to the ServiceId field.
func (o *ServiceTokenResponse) SetServiceId(v int64) {
	o.ServiceId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ServiceTokenResponse) GetOrganizationId() int64 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int64
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTokenResponse) GetOrganizationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ServiceTokenResponse) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int64 and assigns it to the OrganizationId field.
func (o *ServiceTokenResponse) SetOrganizationId(v int64) {
	o.OrganizationId = &v
}

// GetApiServerId returns the ApiServerId field value if set, zero value otherwise.
func (o *ServiceTokenResponse) GetApiServerId() int64 {
	if o == nil || IsNil(o.ApiServerId) {
		var ret int64
		return ret
	}
	return *o.ApiServerId
}

// GetApiServerIdOk returns a tuple with the ApiServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTokenResponse) GetApiServerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ApiServerId) {
		return nil, false
	}
	return o.ApiServerId, true
}

// HasApiServerId returns a boolean if a field has been set.
func (o *ServiceTokenResponse) HasApiServerId() bool {
	if o != nil && !IsNil(o.ApiServerId) {
		return true
	}

	return false
}

// SetApiServerId gets a reference to the given int64 and assigns it to the ApiServerId field.
func (o *ServiceTokenResponse) SetApiServerId(v int64) {
	o.ApiServerId = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *ServiceTokenResponse) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *ServiceTokenResponse) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *ServiceTokenResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ServiceTokenResponse) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTokenResponse) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ServiceTokenResponse) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ServiceTokenResponse) SetTokenId(v string) {
	o.TokenId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceTokenResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTokenResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceTokenResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceTokenResponse) SetDescription(v string) {
	o.Description = &v
}

func (o ServiceTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.ApiServerId) {
		toSerialize["apiServerId"] = o.ApiServerId
	}
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableServiceTokenResponse struct {
	value *ServiceTokenResponse
	isSet bool
}

func (v NullableServiceTokenResponse) Get() *ServiceTokenResponse {
	return v.value
}

func (v *NullableServiceTokenResponse) Set(val *ServiceTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTokenResponse(val *ServiceTokenResponse) *NullableServiceTokenResponse {
	return &NullableServiceTokenResponse{value: val, isSet: true}
}

func (v NullableServiceTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
