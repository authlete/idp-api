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

// UpdateApiServerPrivilegesRequest struct for UpdateApiServerPrivilegesRequest
type UpdateApiServerPrivilegesRequest struct {
	ApiServerId int64 `json:"apiServerId"`
	OrganizationId int64 `json:"organizationId"`
	UserId int64 `json:"userId"`
	Privileges []string `json:"privileges,omitempty"`
	AllServicePrivileges []string `json:"allServicePrivileges,omitempty"`
	AllClientPrivileges []string `json:"allClientPrivileges,omitempty"`
}

// NewUpdateApiServerPrivilegesRequest instantiates a new UpdateApiServerPrivilegesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApiServerPrivilegesRequest(apiServerId int64, organizationId int64, userId int64) *UpdateApiServerPrivilegesRequest {
	this := UpdateApiServerPrivilegesRequest{}
	this.ApiServerId = apiServerId
	this.OrganizationId = organizationId
	this.UserId = userId
	return &this
}

// NewUpdateApiServerPrivilegesRequestWithDefaults instantiates a new UpdateApiServerPrivilegesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApiServerPrivilegesRequestWithDefaults() *UpdateApiServerPrivilegesRequest {
	this := UpdateApiServerPrivilegesRequest{}
	return &this
}

// GetApiServerId returns the ApiServerId field value
func (o *UpdateApiServerPrivilegesRequest) GetApiServerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApiServerId
}

// GetApiServerIdOk returns a tuple with the ApiServerId field value
// and a boolean to check if the value has been set.
func (o *UpdateApiServerPrivilegesRequest) GetApiServerIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApiServerId, true
}

// SetApiServerId sets field value
func (o *UpdateApiServerPrivilegesRequest) SetApiServerId(v int64) {
	o.ApiServerId = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *UpdateApiServerPrivilegesRequest) GetOrganizationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *UpdateApiServerPrivilegesRequest) GetOrganizationIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *UpdateApiServerPrivilegesRequest) SetOrganizationId(v int64) {
	o.OrganizationId = v
}

// GetUserId returns the UserId field value
func (o *UpdateApiServerPrivilegesRequest) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UpdateApiServerPrivilegesRequest) GetUserIdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UpdateApiServerPrivilegesRequest) SetUserId(v int64) {
	o.UserId = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *UpdateApiServerPrivilegesRequest) GetPrivileges() []string {
	if o == nil || isNil(o.Privileges) {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiServerPrivilegesRequest) GetPrivilegesOk() ([]string, bool) {
	if o == nil || isNil(o.Privileges) {
    return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *UpdateApiServerPrivilegesRequest) HasPrivileges() bool {
	if o != nil && !isNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []string and assigns it to the Privileges field.
func (o *UpdateApiServerPrivilegesRequest) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetAllServicePrivileges returns the AllServicePrivileges field value if set, zero value otherwise.
func (o *UpdateApiServerPrivilegesRequest) GetAllServicePrivileges() []string {
	if o == nil || isNil(o.AllServicePrivileges) {
		var ret []string
		return ret
	}
	return o.AllServicePrivileges
}

// GetAllServicePrivilegesOk returns a tuple with the AllServicePrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiServerPrivilegesRequest) GetAllServicePrivilegesOk() ([]string, bool) {
	if o == nil || isNil(o.AllServicePrivileges) {
    return nil, false
	}
	return o.AllServicePrivileges, true
}

// HasAllServicePrivileges returns a boolean if a field has been set.
func (o *UpdateApiServerPrivilegesRequest) HasAllServicePrivileges() bool {
	if o != nil && !isNil(o.AllServicePrivileges) {
		return true
	}

	return false
}

// SetAllServicePrivileges gets a reference to the given []string and assigns it to the AllServicePrivileges field.
func (o *UpdateApiServerPrivilegesRequest) SetAllServicePrivileges(v []string) {
	o.AllServicePrivileges = v
}

// GetAllClientPrivileges returns the AllClientPrivileges field value if set, zero value otherwise.
func (o *UpdateApiServerPrivilegesRequest) GetAllClientPrivileges() []string {
	if o == nil || isNil(o.AllClientPrivileges) {
		var ret []string
		return ret
	}
	return o.AllClientPrivileges
}

// GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiServerPrivilegesRequest) GetAllClientPrivilegesOk() ([]string, bool) {
	if o == nil || isNil(o.AllClientPrivileges) {
    return nil, false
	}
	return o.AllClientPrivileges, true
}

// HasAllClientPrivileges returns a boolean if a field has been set.
func (o *UpdateApiServerPrivilegesRequest) HasAllClientPrivileges() bool {
	if o != nil && !isNil(o.AllClientPrivileges) {
		return true
	}

	return false
}

// SetAllClientPrivileges gets a reference to the given []string and assigns it to the AllClientPrivileges field.
func (o *UpdateApiServerPrivilegesRequest) SetAllClientPrivileges(v []string) {
	o.AllClientPrivileges = v
}

func (o UpdateApiServerPrivilegesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiServerId"] = o.ApiServerId
	}
	if true {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if !isNil(o.Privileges) {
		toSerialize["privileges"] = o.Privileges
	}
	if !isNil(o.AllServicePrivileges) {
		toSerialize["allServicePrivileges"] = o.AllServicePrivileges
	}
	if !isNil(o.AllClientPrivileges) {
		toSerialize["allClientPrivileges"] = o.AllClientPrivileges
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateApiServerPrivilegesRequest struct {
	value *UpdateApiServerPrivilegesRequest
	isSet bool
}

func (v NullableUpdateApiServerPrivilegesRequest) Get() *UpdateApiServerPrivilegesRequest {
	return v.value
}

func (v *NullableUpdateApiServerPrivilegesRequest) Set(val *UpdateApiServerPrivilegesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApiServerPrivilegesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApiServerPrivilegesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApiServerPrivilegesRequest(val *UpdateApiServerPrivilegesRequest) *NullableUpdateApiServerPrivilegesRequest {
	return &NullableUpdateApiServerPrivilegesRequest{value: val, isSet: true}
}

func (v NullableUpdateApiServerPrivilegesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApiServerPrivilegesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

