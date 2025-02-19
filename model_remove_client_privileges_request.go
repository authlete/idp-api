/*
OpenAPI definition

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the RemoveClientPrivilegesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveClientPrivilegesRequest{}

// RemoveClientPrivilegesRequest struct for RemoveClientPrivilegesRequest
type RemoveClientPrivilegesRequest struct {
	ApiServerId    int64 `json:"apiServerId"`
	OrganizationId int64 `json:"organizationId"`
	UserId         int64 `json:"userId"`
	ServiceId      int64 `json:"serviceId"`
	ClientId       int64 `json:"clientId"`
}

type _RemoveClientPrivilegesRequest RemoveClientPrivilegesRequest

// NewRemoveClientPrivilegesRequest instantiates a new RemoveClientPrivilegesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveClientPrivilegesRequest(apiServerId int64, organizationId int64, userId int64, serviceId int64, clientId int64) *RemoveClientPrivilegesRequest {
	this := RemoveClientPrivilegesRequest{}
	this.ApiServerId = apiServerId
	this.OrganizationId = organizationId
	this.UserId = userId
	this.ServiceId = serviceId
	this.ClientId = clientId
	return &this
}

// NewRemoveClientPrivilegesRequestWithDefaults instantiates a new RemoveClientPrivilegesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveClientPrivilegesRequestWithDefaults() *RemoveClientPrivilegesRequest {
	this := RemoveClientPrivilegesRequest{}
	return &this
}

// GetApiServerId returns the ApiServerId field value
func (o *RemoveClientPrivilegesRequest) GetApiServerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApiServerId
}

// GetApiServerIdOk returns a tuple with the ApiServerId field value
// and a boolean to check if the value has been set.
func (o *RemoveClientPrivilegesRequest) GetApiServerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiServerId, true
}

// SetApiServerId sets field value
func (o *RemoveClientPrivilegesRequest) SetApiServerId(v int64) {
	o.ApiServerId = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *RemoveClientPrivilegesRequest) GetOrganizationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *RemoveClientPrivilegesRequest) GetOrganizationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *RemoveClientPrivilegesRequest) SetOrganizationId(v int64) {
	o.OrganizationId = v
}

// GetUserId returns the UserId field value
func (o *RemoveClientPrivilegesRequest) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *RemoveClientPrivilegesRequest) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *RemoveClientPrivilegesRequest) SetUserId(v int64) {
	o.UserId = v
}

// GetServiceId returns the ServiceId field value
func (o *RemoveClientPrivilegesRequest) GetServiceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *RemoveClientPrivilegesRequest) GetServiceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *RemoveClientPrivilegesRequest) SetServiceId(v int64) {
	o.ServiceId = v
}

// GetClientId returns the ClientId field value
func (o *RemoveClientPrivilegesRequest) GetClientId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *RemoveClientPrivilegesRequest) GetClientIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *RemoveClientPrivilegesRequest) SetClientId(v int64) {
	o.ClientId = v
}

func (o RemoveClientPrivilegesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveClientPrivilegesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiServerId"] = o.ApiServerId
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["userId"] = o.UserId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["clientId"] = o.ClientId
	return toSerialize, nil
}

func (o *RemoveClientPrivilegesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiServerId",
		"organizationId",
		"userId",
		"serviceId",
		"clientId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRemoveClientPrivilegesRequest := _RemoveClientPrivilegesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRemoveClientPrivilegesRequest)

	if err != nil {
		return err
	}

	*o = RemoveClientPrivilegesRequest(varRemoveClientPrivilegesRequest)

	return err
}

type NullableRemoveClientPrivilegesRequest struct {
	value *RemoveClientPrivilegesRequest
	isSet bool
}

func (v NullableRemoveClientPrivilegesRequest) Get() *RemoveClientPrivilegesRequest {
	return v.value
}

func (v *NullableRemoveClientPrivilegesRequest) Set(val *RemoveClientPrivilegesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveClientPrivilegesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveClientPrivilegesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveClientPrivilegesRequest(val *RemoveClientPrivilegesRequest) *NullableRemoveClientPrivilegesRequest {
	return &NullableRemoveClientPrivilegesRequest{value: val, isSet: true}
}

func (v NullableRemoveClientPrivilegesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveClientPrivilegesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
