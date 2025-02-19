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

// checks if the MoveServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveServiceRequest{}

// MoveServiceRequest struct for MoveServiceRequest
type MoveServiceRequest struct {
	OrganizationId    int64 `json:"organizationId"`
	ApiServerId       int64 `json:"apiServerId"`
	ServiceId         int64 `json:"serviceId"`
	NewOrganizationId int64 `json:"newOrganizationId"`
}

type _MoveServiceRequest MoveServiceRequest

// NewMoveServiceRequest instantiates a new MoveServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveServiceRequest(organizationId int64, apiServerId int64, serviceId int64, newOrganizationId int64) *MoveServiceRequest {
	this := MoveServiceRequest{}
	this.OrganizationId = organizationId
	this.ApiServerId = apiServerId
	this.ServiceId = serviceId
	this.NewOrganizationId = newOrganizationId
	return &this
}

// NewMoveServiceRequestWithDefaults instantiates a new MoveServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveServiceRequestWithDefaults() *MoveServiceRequest {
	this := MoveServiceRequest{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value
func (o *MoveServiceRequest) GetOrganizationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *MoveServiceRequest) GetOrganizationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *MoveServiceRequest) SetOrganizationId(v int64) {
	o.OrganizationId = v
}

// GetApiServerId returns the ApiServerId field value
func (o *MoveServiceRequest) GetApiServerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApiServerId
}

// GetApiServerIdOk returns a tuple with the ApiServerId field value
// and a boolean to check if the value has been set.
func (o *MoveServiceRequest) GetApiServerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiServerId, true
}

// SetApiServerId sets field value
func (o *MoveServiceRequest) SetApiServerId(v int64) {
	o.ApiServerId = v
}

// GetServiceId returns the ServiceId field value
func (o *MoveServiceRequest) GetServiceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *MoveServiceRequest) GetServiceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *MoveServiceRequest) SetServiceId(v int64) {
	o.ServiceId = v
}

// GetNewOrganizationId returns the NewOrganizationId field value
func (o *MoveServiceRequest) GetNewOrganizationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NewOrganizationId
}

// GetNewOrganizationIdOk returns a tuple with the NewOrganizationId field value
// and a boolean to check if the value has been set.
func (o *MoveServiceRequest) GetNewOrganizationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewOrganizationId, true
}

// SetNewOrganizationId sets field value
func (o *MoveServiceRequest) SetNewOrganizationId(v int64) {
	o.NewOrganizationId = v
}

func (o MoveServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["apiServerId"] = o.ApiServerId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["newOrganizationId"] = o.NewOrganizationId
	return toSerialize, nil
}

func (o *MoveServiceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"organizationId",
		"apiServerId",
		"serviceId",
		"newOrganizationId",
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

	varMoveServiceRequest := _MoveServiceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMoveServiceRequest)

	if err != nil {
		return err
	}

	*o = MoveServiceRequest(varMoveServiceRequest)

	return err
}

type NullableMoveServiceRequest struct {
	value *MoveServiceRequest
	isSet bool
}

func (v NullableMoveServiceRequest) Get() *MoveServiceRequest {
	return v.value
}

func (v *NullableMoveServiceRequest) Set(val *MoveServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveServiceRequest(val *MoveServiceRequest) *NullableMoveServiceRequest {
	return &NullableMoveServiceRequest{value: val, isSet: true}
}

func (v NullableMoveServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
