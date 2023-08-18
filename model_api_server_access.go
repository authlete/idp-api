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

// ApiServerAccess struct for ApiServerAccess
type ApiServerAccess struct {
	Id *int64 `json:"id,omitempty"`
	Privileges []string `json:"privileges,omitempty"`
	AllServicePrivileges []string `json:"allServicePrivileges,omitempty"`
	AllClientPrivileges []string `json:"allClientPrivileges,omitempty"`
	ServiceAccess []ServiceAccess `json:"serviceAccess,omitempty"`
	OrganizationAccessId *int64 `json:"organization_access_id,omitempty"`
	ApiServerId *int64 `json:"api_server_id,omitempty"`
}

// NewApiServerAccess instantiates a new ApiServerAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiServerAccess() *ApiServerAccess {
	this := ApiServerAccess{}
	return &this
}

// NewApiServerAccessWithDefaults instantiates a new ApiServerAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiServerAccessWithDefaults() *ApiServerAccess {
	this := ApiServerAccess{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiServerAccess) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiServerAccess) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiServerAccess) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ApiServerAccess) SetId(v int64) {
	o.Id = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *ApiServerAccess) GetPrivileges() []string {
	if o == nil || isNil(o.Privileges) {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiServerAccess) GetPrivilegesOk() ([]string, bool) {
	if o == nil || isNil(o.Privileges) {
    return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *ApiServerAccess) HasPrivileges() bool {
	if o != nil && !isNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []string and assigns it to the Privileges field.
func (o *ApiServerAccess) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetAllServicePrivileges returns the AllServicePrivileges field value if set, zero value otherwise.
func (o *ApiServerAccess) GetAllServicePrivileges() []string {
	if o == nil || isNil(o.AllServicePrivileges) {
		var ret []string
		return ret
	}
	return o.AllServicePrivileges
}

// GetAllServicePrivilegesOk returns a tuple with the AllServicePrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiServerAccess) GetAllServicePrivilegesOk() ([]string, bool) {
	if o == nil || isNil(o.AllServicePrivileges) {
    return nil, false
	}
	return o.AllServicePrivileges, true
}

// HasAllServicePrivileges returns a boolean if a field has been set.
func (o *ApiServerAccess) HasAllServicePrivileges() bool {
	if o != nil && !isNil(o.AllServicePrivileges) {
		return true
	}

	return false
}

// SetAllServicePrivileges gets a reference to the given []string and assigns it to the AllServicePrivileges field.
func (o *ApiServerAccess) SetAllServicePrivileges(v []string) {
	o.AllServicePrivileges = v
}

// GetAllClientPrivileges returns the AllClientPrivileges field value if set, zero value otherwise.
func (o *ApiServerAccess) GetAllClientPrivileges() []string {
	if o == nil || isNil(o.AllClientPrivileges) {
		var ret []string
		return ret
	}
	return o.AllClientPrivileges
}

// GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiServerAccess) GetAllClientPrivilegesOk() ([]string, bool) {
	if o == nil || isNil(o.AllClientPrivileges) {
    return nil, false
	}
	return o.AllClientPrivileges, true
}

// HasAllClientPrivileges returns a boolean if a field has been set.
func (o *ApiServerAccess) HasAllClientPrivileges() bool {
	if o != nil && !isNil(o.AllClientPrivileges) {
		return true
	}

	return false
}

// SetAllClientPrivileges gets a reference to the given []string and assigns it to the AllClientPrivileges field.
func (o *ApiServerAccess) SetAllClientPrivileges(v []string) {
	o.AllClientPrivileges = v
}

// GetServiceAccess returns the ServiceAccess field value if set, zero value otherwise.
func (o *ApiServerAccess) GetServiceAccess() []ServiceAccess {
	if o == nil || isNil(o.ServiceAccess) {
		var ret []ServiceAccess
		return ret
	}
	return o.ServiceAccess
}

// GetServiceAccessOk returns a tuple with the ServiceAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiServerAccess) GetServiceAccessOk() ([]ServiceAccess, bool) {
	if o == nil || isNil(o.ServiceAccess) {
    return nil, false
	}
	return o.ServiceAccess, true
}

// HasServiceAccess returns a boolean if a field has been set.
func (o *ApiServerAccess) HasServiceAccess() bool {
	if o != nil && !isNil(o.ServiceAccess) {
		return true
	}

	return false
}

// SetServiceAccess gets a reference to the given []ServiceAccess and assigns it to the ServiceAccess field.
func (o *ApiServerAccess) SetServiceAccess(v []ServiceAccess) {
	o.ServiceAccess = v
}

// GetOrganizationAccessId returns the OrganizationAccessId field value if set, zero value otherwise.
func (o *ApiServerAccess) GetOrganizationAccessId() int64 {
	if o == nil || isNil(o.OrganizationAccessId) {
		var ret int64
		return ret
	}
	return *o.OrganizationAccessId
}

// GetOrganizationAccessIdOk returns a tuple with the OrganizationAccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiServerAccess) GetOrganizationAccessIdOk() (*int64, bool) {
	if o == nil || isNil(o.OrganizationAccessId) {
    return nil, false
	}
	return o.OrganizationAccessId, true
}

// HasOrganizationAccessId returns a boolean if a field has been set.
func (o *ApiServerAccess) HasOrganizationAccessId() bool {
	if o != nil && !isNil(o.OrganizationAccessId) {
		return true
	}

	return false
}

// SetOrganizationAccessId gets a reference to the given int64 and assigns it to the OrganizationAccessId field.
func (o *ApiServerAccess) SetOrganizationAccessId(v int64) {
	o.OrganizationAccessId = &v
}

// GetApiServerId returns the ApiServerId field value if set, zero value otherwise.
func (o *ApiServerAccess) GetApiServerId() int64 {
	if o == nil || isNil(o.ApiServerId) {
		var ret int64
		return ret
	}
	return *o.ApiServerId
}

// GetApiServerIdOk returns a tuple with the ApiServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiServerAccess) GetApiServerIdOk() (*int64, bool) {
	if o == nil || isNil(o.ApiServerId) {
    return nil, false
	}
	return o.ApiServerId, true
}

// HasApiServerId returns a boolean if a field has been set.
func (o *ApiServerAccess) HasApiServerId() bool {
	if o != nil && !isNil(o.ApiServerId) {
		return true
	}

	return false
}

// SetApiServerId gets a reference to the given int64 and assigns it to the ApiServerId field.
func (o *ApiServerAccess) SetApiServerId(v int64) {
	o.ApiServerId = &v
}

func (o ApiServerAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !isNil(o.ServiceAccess) {
		toSerialize["serviceAccess"] = o.ServiceAccess
	}
	if !isNil(o.OrganizationAccessId) {
		toSerialize["organization_access_id"] = o.OrganizationAccessId
	}
	if !isNil(o.ApiServerId) {
		toSerialize["api_server_id"] = o.ApiServerId
	}
	return json.Marshal(toSerialize)
}

type NullableApiServerAccess struct {
	value *ApiServerAccess
	isSet bool
}

func (v NullableApiServerAccess) Get() *ApiServerAccess {
	return v.value
}

func (v *NullableApiServerAccess) Set(val *ApiServerAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableApiServerAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableApiServerAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiServerAccess(val *ApiServerAccess) *NullableApiServerAccess {
	return &NullableApiServerAccess{value: val, isSet: true}
}

func (v NullableApiServerAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiServerAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


