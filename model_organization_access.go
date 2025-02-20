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

// checks if the OrganizationAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationAccess{}

// OrganizationAccess struct for OrganizationAccess
type OrganizationAccess struct {
	Id                     *int64            `json:"id,omitempty"`
	Privileges             []string          `json:"privileges,omitempty"`
	AllApiServerPrivileges []string          `json:"allApiServerPrivileges,omitempty"`
	AllServicePrivileges   []string          `json:"allServicePrivileges,omitempty"`
	AllClientPrivileges    []string          `json:"allClientPrivileges,omitempty"`
	ApiServerAccess        []ApiServerAccess `json:"apiServerAccess,omitempty"`
	OrganizationId         *int64            `json:"organization_id,omitempty"`
	UserId                 *int64            `json:"user_id,omitempty"`
}

// NewOrganizationAccess instantiates a new OrganizationAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAccess() *OrganizationAccess {
	this := OrganizationAccess{}
	return &this
}

// NewOrganizationAccessWithDefaults instantiates a new OrganizationAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAccessWithDefaults() *OrganizationAccess {
	this := OrganizationAccess{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationAccess) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAccess) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationAccess) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OrganizationAccess) SetId(v int64) {
	o.Id = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *OrganizationAccess) GetPrivileges() []string {
	if o == nil || IsNil(o.Privileges) {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAccess) GetPrivilegesOk() ([]string, bool) {
	if o == nil || IsNil(o.Privileges) {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *OrganizationAccess) HasPrivileges() bool {
	if o != nil && !IsNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []string and assigns it to the Privileges field.
func (o *OrganizationAccess) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetAllApiServerPrivileges returns the AllApiServerPrivileges field value if set, zero value otherwise.
func (o *OrganizationAccess) GetAllApiServerPrivileges() []string {
	if o == nil || IsNil(o.AllApiServerPrivileges) {
		var ret []string
		return ret
	}
	return o.AllApiServerPrivileges
}

// GetAllApiServerPrivilegesOk returns a tuple with the AllApiServerPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAccess) GetAllApiServerPrivilegesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllApiServerPrivileges) {
		return nil, false
	}
	return o.AllApiServerPrivileges, true
}

// HasAllApiServerPrivileges returns a boolean if a field has been set.
func (o *OrganizationAccess) HasAllApiServerPrivileges() bool {
	if o != nil && !IsNil(o.AllApiServerPrivileges) {
		return true
	}

	return false
}

// SetAllApiServerPrivileges gets a reference to the given []string and assigns it to the AllApiServerPrivileges field.
func (o *OrganizationAccess) SetAllApiServerPrivileges(v []string) {
	o.AllApiServerPrivileges = v
}

// GetAllServicePrivileges returns the AllServicePrivileges field value if set, zero value otherwise.
func (o *OrganizationAccess) GetAllServicePrivileges() []string {
	if o == nil || IsNil(o.AllServicePrivileges) {
		var ret []string
		return ret
	}
	return o.AllServicePrivileges
}

// GetAllServicePrivilegesOk returns a tuple with the AllServicePrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAccess) GetAllServicePrivilegesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllServicePrivileges) {
		return nil, false
	}
	return o.AllServicePrivileges, true
}

// HasAllServicePrivileges returns a boolean if a field has been set.
func (o *OrganizationAccess) HasAllServicePrivileges() bool {
	if o != nil && !IsNil(o.AllServicePrivileges) {
		return true
	}

	return false
}

// SetAllServicePrivileges gets a reference to the given []string and assigns it to the AllServicePrivileges field.
func (o *OrganizationAccess) SetAllServicePrivileges(v []string) {
	o.AllServicePrivileges = v
}

// GetAllClientPrivileges returns the AllClientPrivileges field value if set, zero value otherwise.
func (o *OrganizationAccess) GetAllClientPrivileges() []string {
	if o == nil || IsNil(o.AllClientPrivileges) {
		var ret []string
		return ret
	}
	return o.AllClientPrivileges
}

// GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAccess) GetAllClientPrivilegesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllClientPrivileges) {
		return nil, false
	}
	return o.AllClientPrivileges, true
}

// HasAllClientPrivileges returns a boolean if a field has been set.
func (o *OrganizationAccess) HasAllClientPrivileges() bool {
	if o != nil && !IsNil(o.AllClientPrivileges) {
		return true
	}

	return false
}

// SetAllClientPrivileges gets a reference to the given []string and assigns it to the AllClientPrivileges field.
func (o *OrganizationAccess) SetAllClientPrivileges(v []string) {
	o.AllClientPrivileges = v
}

// GetApiServerAccess returns the ApiServerAccess field value if set, zero value otherwise.
func (o *OrganizationAccess) GetApiServerAccess() []ApiServerAccess {
	if o == nil || IsNil(o.ApiServerAccess) {
		var ret []ApiServerAccess
		return ret
	}
	return o.ApiServerAccess
}

// GetApiServerAccessOk returns a tuple with the ApiServerAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAccess) GetApiServerAccessOk() ([]ApiServerAccess, bool) {
	if o == nil || IsNil(o.ApiServerAccess) {
		return nil, false
	}
	return o.ApiServerAccess, true
}

// HasApiServerAccess returns a boolean if a field has been set.
func (o *OrganizationAccess) HasApiServerAccess() bool {
	if o != nil && !IsNil(o.ApiServerAccess) {
		return true
	}

	return false
}

// SetApiServerAccess gets a reference to the given []ApiServerAccess and assigns it to the ApiServerAccess field.
func (o *OrganizationAccess) SetApiServerAccess(v []ApiServerAccess) {
	o.ApiServerAccess = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *OrganizationAccess) GetOrganizationId() int64 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int64
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAccess) GetOrganizationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *OrganizationAccess) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int64 and assigns it to the OrganizationId field.
func (o *OrganizationAccess) SetOrganizationId(v int64) {
	o.OrganizationId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *OrganizationAccess) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAccess) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *OrganizationAccess) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *OrganizationAccess) SetUserId(v int64) {
	o.UserId = &v
}

func (o OrganizationAccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Privileges) {
		toSerialize["privileges"] = o.Privileges
	}
	if !IsNil(o.AllApiServerPrivileges) {
		toSerialize["allApiServerPrivileges"] = o.AllApiServerPrivileges
	}
	if !IsNil(o.AllServicePrivileges) {
		toSerialize["allServicePrivileges"] = o.AllServicePrivileges
	}
	if !IsNil(o.AllClientPrivileges) {
		toSerialize["allClientPrivileges"] = o.AllClientPrivileges
	}
	if !IsNil(o.ApiServerAccess) {
		toSerialize["apiServerAccess"] = o.ApiServerAccess
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableOrganizationAccess struct {
	value *OrganizationAccess
	isSet bool
}

func (v NullableOrganizationAccess) Get() *OrganizationAccess {
	return v.value
}

func (v *NullableOrganizationAccess) Set(val *OrganizationAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAccess(val *OrganizationAccess) *NullableOrganizationAccess {
	return &NullableOrganizationAccess{value: val, isSet: true}
}

func (v NullableOrganizationAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
