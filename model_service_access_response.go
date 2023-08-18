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

// ServiceAccessResponse struct for ServiceAccessResponse
type ServiceAccessResponse struct {
	UserId *int64 `json:"userId,omitempty"`
	OrganizationId *int64 `json:"organizationId,omitempty"`
	ApiServerId *int64 `json:"apiServerId,omitempty"`
	ServiceId *int64 `json:"serviceId,omitempty"`
	AllClientPrivileges []string `json:"allClientPrivileges,omitempty"`
	Privileges []string `json:"privileges,omitempty"`
	ClientAccess []ClientAccessResponse `json:"clientAccess,omitempty"`
}

// NewServiceAccessResponse instantiates a new ServiceAccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccessResponse() *ServiceAccessResponse {
	this := ServiceAccessResponse{}
	return &this
}

// NewServiceAccessResponseWithDefaults instantiates a new ServiceAccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccessResponseWithDefaults() *ServiceAccessResponse {
	this := ServiceAccessResponse{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ServiceAccessResponse) GetUserId() int64 {
	if o == nil || isNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessResponse) GetUserIdOk() (*int64, bool) {
	if o == nil || isNil(o.UserId) {
    return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ServiceAccessResponse) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *ServiceAccessResponse) SetUserId(v int64) {
	o.UserId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ServiceAccessResponse) GetOrganizationId() int64 {
	if o == nil || isNil(o.OrganizationId) {
		var ret int64
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessResponse) GetOrganizationIdOk() (*int64, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ServiceAccessResponse) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int64 and assigns it to the OrganizationId field.
func (o *ServiceAccessResponse) SetOrganizationId(v int64) {
	o.OrganizationId = &v
}

// GetApiServerId returns the ApiServerId field value if set, zero value otherwise.
func (o *ServiceAccessResponse) GetApiServerId() int64 {
	if o == nil || isNil(o.ApiServerId) {
		var ret int64
		return ret
	}
	return *o.ApiServerId
}

// GetApiServerIdOk returns a tuple with the ApiServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessResponse) GetApiServerIdOk() (*int64, bool) {
	if o == nil || isNil(o.ApiServerId) {
    return nil, false
	}
	return o.ApiServerId, true
}

// HasApiServerId returns a boolean if a field has been set.
func (o *ServiceAccessResponse) HasApiServerId() bool {
	if o != nil && !isNil(o.ApiServerId) {
		return true
	}

	return false
}

// SetApiServerId gets a reference to the given int64 and assigns it to the ApiServerId field.
func (o *ServiceAccessResponse) SetApiServerId(v int64) {
	o.ApiServerId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ServiceAccessResponse) GetServiceId() int64 {
	if o == nil || isNil(o.ServiceId) {
		var ret int64
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessResponse) GetServiceIdOk() (*int64, bool) {
	if o == nil || isNil(o.ServiceId) {
    return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ServiceAccessResponse) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given int64 and assigns it to the ServiceId field.
func (o *ServiceAccessResponse) SetServiceId(v int64) {
	o.ServiceId = &v
}

// GetAllClientPrivileges returns the AllClientPrivileges field value if set, zero value otherwise.
func (o *ServiceAccessResponse) GetAllClientPrivileges() []string {
	if o == nil || isNil(o.AllClientPrivileges) {
		var ret []string
		return ret
	}
	return o.AllClientPrivileges
}

// GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessResponse) GetAllClientPrivilegesOk() ([]string, bool) {
	if o == nil || isNil(o.AllClientPrivileges) {
    return nil, false
	}
	return o.AllClientPrivileges, true
}

// HasAllClientPrivileges returns a boolean if a field has been set.
func (o *ServiceAccessResponse) HasAllClientPrivileges() bool {
	if o != nil && !isNil(o.AllClientPrivileges) {
		return true
	}

	return false
}

// SetAllClientPrivileges gets a reference to the given []string and assigns it to the AllClientPrivileges field.
func (o *ServiceAccessResponse) SetAllClientPrivileges(v []string) {
	o.AllClientPrivileges = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *ServiceAccessResponse) GetPrivileges() []string {
	if o == nil || isNil(o.Privileges) {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessResponse) GetPrivilegesOk() ([]string, bool) {
	if o == nil || isNil(o.Privileges) {
    return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *ServiceAccessResponse) HasPrivileges() bool {
	if o != nil && !isNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []string and assigns it to the Privileges field.
func (o *ServiceAccessResponse) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetClientAccess returns the ClientAccess field value if set, zero value otherwise.
func (o *ServiceAccessResponse) GetClientAccess() []ClientAccessResponse {
	if o == nil || isNil(o.ClientAccess) {
		var ret []ClientAccessResponse
		return ret
	}
	return o.ClientAccess
}

// GetClientAccessOk returns a tuple with the ClientAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessResponse) GetClientAccessOk() ([]ClientAccessResponse, bool) {
	if o == nil || isNil(o.ClientAccess) {
    return nil, false
	}
	return o.ClientAccess, true
}

// HasClientAccess returns a boolean if a field has been set.
func (o *ServiceAccessResponse) HasClientAccess() bool {
	if o != nil && !isNil(o.ClientAccess) {
		return true
	}

	return false
}

// SetClientAccess gets a reference to the given []ClientAccessResponse and assigns it to the ClientAccess field.
func (o *ServiceAccessResponse) SetClientAccess(v []ClientAccessResponse) {
	o.ClientAccess = v
}

func (o ServiceAccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.ApiServerId) {
		toSerialize["apiServerId"] = o.ApiServerId
	}
	if !isNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !isNil(o.AllClientPrivileges) {
		toSerialize["allClientPrivileges"] = o.AllClientPrivileges
	}
	if !isNil(o.Privileges) {
		toSerialize["privileges"] = o.Privileges
	}
	if !isNil(o.ClientAccess) {
		toSerialize["clientAccess"] = o.ClientAccess
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccessResponse struct {
	value *ServiceAccessResponse
	isSet bool
}

func (v NullableServiceAccessResponse) Get() *ServiceAccessResponse {
	return v.value
}

func (v *NullableServiceAccessResponse) Set(val *ServiceAccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccessResponse(val *ServiceAccessResponse) *NullableServiceAccessResponse {
	return &NullableServiceAccessResponse{value: val, isSet: true}
}

func (v NullableServiceAccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


