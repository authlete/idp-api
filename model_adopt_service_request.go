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

// checks if the AdoptServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdoptServiceRequest{}

// AdoptServiceRequest struct for AdoptServiceRequest
type AdoptServiceRequest struct {
	OrganizationId int64 `json:"organizationId"`
	ApiServerId    int64 `json:"apiServerId"`
	ServiceId      int64 `json:"serviceId"`
}

type _AdoptServiceRequest AdoptServiceRequest

// NewAdoptServiceRequest instantiates a new AdoptServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdoptServiceRequest(organizationId int64, apiServerId int64, serviceId int64) *AdoptServiceRequest {
	this := AdoptServiceRequest{}
	this.OrganizationId = organizationId
	this.ApiServerId = apiServerId
	this.ServiceId = serviceId
	return &this
}

// NewAdoptServiceRequestWithDefaults instantiates a new AdoptServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdoptServiceRequestWithDefaults() *AdoptServiceRequest {
	this := AdoptServiceRequest{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value
func (o *AdoptServiceRequest) GetOrganizationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *AdoptServiceRequest) GetOrganizationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *AdoptServiceRequest) SetOrganizationId(v int64) {
	o.OrganizationId = v
}

// GetApiServerId returns the ApiServerId field value
func (o *AdoptServiceRequest) GetApiServerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApiServerId
}

// GetApiServerIdOk returns a tuple with the ApiServerId field value
// and a boolean to check if the value has been set.
func (o *AdoptServiceRequest) GetApiServerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiServerId, true
}

// SetApiServerId sets field value
func (o *AdoptServiceRequest) SetApiServerId(v int64) {
	o.ApiServerId = v
}

// GetServiceId returns the ServiceId field value
func (o *AdoptServiceRequest) GetServiceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *AdoptServiceRequest) GetServiceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *AdoptServiceRequest) SetServiceId(v int64) {
	o.ServiceId = v
}

func (o AdoptServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdoptServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["apiServerId"] = o.ApiServerId
	toSerialize["serviceId"] = o.ServiceId
	return toSerialize, nil
}

func (o *AdoptServiceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"organizationId",
		"apiServerId",
		"serviceId",
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

	varAdoptServiceRequest := _AdoptServiceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdoptServiceRequest)

	if err != nil {
		return err
	}

	*o = AdoptServiceRequest(varAdoptServiceRequest)

	return err
}

type NullableAdoptServiceRequest struct {
	value *AdoptServiceRequest
	isSet bool
}

func (v NullableAdoptServiceRequest) Get() *AdoptServiceRequest {
	return v.value
}

func (v *NullableAdoptServiceRequest) Set(val *AdoptServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAdoptServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAdoptServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdoptServiceRequest(val *AdoptServiceRequest) *NullableAdoptServiceRequest {
	return &NullableAdoptServiceRequest{value: val, isSet: true}
}

func (v NullableAdoptServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdoptServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
