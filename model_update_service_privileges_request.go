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

// checks if the UpdateServicePrivilegesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServicePrivilegesRequest{}

// UpdateServicePrivilegesRequest struct for UpdateServicePrivilegesRequest
type UpdateServicePrivilegesRequest struct {
	ApiServerId         int64    `json:"apiServerId"`
	OrganizationId      int64    `json:"organizationId"`
	UserId              int64    `json:"userId"`
	ServiceId           int64    `json:"serviceId"`
	Privileges          []string `json:"privileges,omitempty"`
	AllClientPrivileges []string `json:"allClientPrivileges,omitempty"`
}

type _UpdateServicePrivilegesRequest UpdateServicePrivilegesRequest

// NewUpdateServicePrivilegesRequest instantiates a new UpdateServicePrivilegesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServicePrivilegesRequest(apiServerId int64, organizationId int64, userId int64, serviceId int64) *UpdateServicePrivilegesRequest {
	this := UpdateServicePrivilegesRequest{}
	this.ApiServerId = apiServerId
	this.OrganizationId = organizationId
	this.UserId = userId
	this.ServiceId = serviceId
	return &this
}

// NewUpdateServicePrivilegesRequestWithDefaults instantiates a new UpdateServicePrivilegesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServicePrivilegesRequestWithDefaults() *UpdateServicePrivilegesRequest {
	this := UpdateServicePrivilegesRequest{}
	return &this
}

// GetApiServerId returns the ApiServerId field value
func (o *UpdateServicePrivilegesRequest) GetApiServerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApiServerId
}

// GetApiServerIdOk returns a tuple with the ApiServerId field value
// and a boolean to check if the value has been set.
func (o *UpdateServicePrivilegesRequest) GetApiServerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiServerId, true
}

// SetApiServerId sets field value
func (o *UpdateServicePrivilegesRequest) SetApiServerId(v int64) {
	o.ApiServerId = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *UpdateServicePrivilegesRequest) GetOrganizationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *UpdateServicePrivilegesRequest) GetOrganizationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *UpdateServicePrivilegesRequest) SetOrganizationId(v int64) {
	o.OrganizationId = v
}

// GetUserId returns the UserId field value
func (o *UpdateServicePrivilegesRequest) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UpdateServicePrivilegesRequest) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UpdateServicePrivilegesRequest) SetUserId(v int64) {
	o.UserId = v
}

// GetServiceId returns the ServiceId field value
func (o *UpdateServicePrivilegesRequest) GetServiceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpdateServicePrivilegesRequest) GetServiceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpdateServicePrivilegesRequest) SetServiceId(v int64) {
	o.ServiceId = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *UpdateServicePrivilegesRequest) GetPrivileges() []string {
	if o == nil || IsNil(o.Privileges) {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServicePrivilegesRequest) GetPrivilegesOk() ([]string, bool) {
	if o == nil || IsNil(o.Privileges) {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *UpdateServicePrivilegesRequest) HasPrivileges() bool {
	if o != nil && !IsNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []string and assigns it to the Privileges field.
func (o *UpdateServicePrivilegesRequest) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetAllClientPrivileges returns the AllClientPrivileges field value if set, zero value otherwise.
func (o *UpdateServicePrivilegesRequest) GetAllClientPrivileges() []string {
	if o == nil || IsNil(o.AllClientPrivileges) {
		var ret []string
		return ret
	}
	return o.AllClientPrivileges
}

// GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServicePrivilegesRequest) GetAllClientPrivilegesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllClientPrivileges) {
		return nil, false
	}
	return o.AllClientPrivileges, true
}

// HasAllClientPrivileges returns a boolean if a field has been set.
func (o *UpdateServicePrivilegesRequest) HasAllClientPrivileges() bool {
	if o != nil && !IsNil(o.AllClientPrivileges) {
		return true
	}

	return false
}

// SetAllClientPrivileges gets a reference to the given []string and assigns it to the AllClientPrivileges field.
func (o *UpdateServicePrivilegesRequest) SetAllClientPrivileges(v []string) {
	o.AllClientPrivileges = v
}

func (o UpdateServicePrivilegesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServicePrivilegesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiServerId"] = o.ApiServerId
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["userId"] = o.UserId
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.Privileges) {
		toSerialize["privileges"] = o.Privileges
	}
	if !IsNil(o.AllClientPrivileges) {
		toSerialize["allClientPrivileges"] = o.AllClientPrivileges
	}
	return toSerialize, nil
}

func (o *UpdateServicePrivilegesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiServerId",
		"organizationId",
		"userId",
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

	varUpdateServicePrivilegesRequest := _UpdateServicePrivilegesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateServicePrivilegesRequest)

	if err != nil {
		return err
	}

	*o = UpdateServicePrivilegesRequest(varUpdateServicePrivilegesRequest)

	return err
}

type NullableUpdateServicePrivilegesRequest struct {
	value *UpdateServicePrivilegesRequest
	isSet bool
}

func (v NullableUpdateServicePrivilegesRequest) Get() *UpdateServicePrivilegesRequest {
	return v.value
}

func (v *NullableUpdateServicePrivilegesRequest) Set(val *UpdateServicePrivilegesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServicePrivilegesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServicePrivilegesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServicePrivilegesRequest(val *UpdateServicePrivilegesRequest) *NullableUpdateServicePrivilegesRequest {
	return &NullableUpdateServicePrivilegesRequest{value: val, isSet: true}
}

func (v NullableUpdateServicePrivilegesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServicePrivilegesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
