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

// checks if the OrphanService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrphanService{}

// OrphanService struct for OrphanService
type OrphanService struct {
	ApiServerId *int64 `json:"apiServerId,omitempty"`
	ServiceId   *int64 `json:"serviceId,omitempty"`
}

// NewOrphanService instantiates a new OrphanService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrphanService() *OrphanService {
	this := OrphanService{}
	return &this
}

// NewOrphanServiceWithDefaults instantiates a new OrphanService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrphanServiceWithDefaults() *OrphanService {
	this := OrphanService{}
	return &this
}

// GetApiServerId returns the ApiServerId field value if set, zero value otherwise.
func (o *OrphanService) GetApiServerId() int64 {
	if o == nil || IsNil(o.ApiServerId) {
		var ret int64
		return ret
	}
	return *o.ApiServerId
}

// GetApiServerIdOk returns a tuple with the ApiServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrphanService) GetApiServerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ApiServerId) {
		return nil, false
	}
	return o.ApiServerId, true
}

// HasApiServerId returns a boolean if a field has been set.
func (o *OrphanService) HasApiServerId() bool {
	if o != nil && !IsNil(o.ApiServerId) {
		return true
	}

	return false
}

// SetApiServerId gets a reference to the given int64 and assigns it to the ApiServerId field.
func (o *OrphanService) SetApiServerId(v int64) {
	o.ApiServerId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *OrphanService) GetServiceId() int64 {
	if o == nil || IsNil(o.ServiceId) {
		var ret int64
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrphanService) GetServiceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *OrphanService) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given int64 and assigns it to the ServiceId field.
func (o *OrphanService) SetServiceId(v int64) {
	o.ServiceId = &v
}

func (o OrphanService) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrphanService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiServerId) {
		toSerialize["apiServerId"] = o.ApiServerId
	}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	return toSerialize, nil
}

type NullableOrphanService struct {
	value *OrphanService
	isSet bool
}

func (v NullableOrphanService) Get() *OrphanService {
	return v.value
}

func (v *NullableOrphanService) Set(val *OrphanService) {
	v.value = val
	v.isSet = true
}

func (v NullableOrphanService) IsSet() bool {
	return v.isSet
}

func (v *NullableOrphanService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrphanService(val *OrphanService) *NullableOrphanService {
	return &NullableOrphanService{value: val, isSet: true}
}

func (v NullableOrphanService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrphanService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
