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

// checks if the UserServiceMembershipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserServiceMembershipResponse{}

// UserServiceMembershipResponse struct for UserServiceMembershipResponse
type UserServiceMembershipResponse struct {
	Organizations []OrganizationMembershipResponse `json:"organizations,omitempty"`
	UserId        *int64                           `json:"userId,omitempty"`
	Admin         *bool                            `json:"admin,omitempty"`
}

// NewUserServiceMembershipResponse instantiates a new UserServiceMembershipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserServiceMembershipResponse() *UserServiceMembershipResponse {
	this := UserServiceMembershipResponse{}
	return &this
}

// NewUserServiceMembershipResponseWithDefaults instantiates a new UserServiceMembershipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserServiceMembershipResponseWithDefaults() *UserServiceMembershipResponse {
	this := UserServiceMembershipResponse{}
	return &this
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *UserServiceMembershipResponse) GetOrganizations() []OrganizationMembershipResponse {
	if o == nil || IsNil(o.Organizations) {
		var ret []OrganizationMembershipResponse
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceMembershipResponse) GetOrganizationsOk() ([]OrganizationMembershipResponse, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *UserServiceMembershipResponse) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []OrganizationMembershipResponse and assigns it to the Organizations field.
func (o *UserServiceMembershipResponse) SetOrganizations(v []OrganizationMembershipResponse) {
	o.Organizations = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserServiceMembershipResponse) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceMembershipResponse) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserServiceMembershipResponse) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UserServiceMembershipResponse) SetUserId(v int64) {
	o.UserId = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *UserServiceMembershipResponse) GetAdmin() bool {
	if o == nil || IsNil(o.Admin) {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceMembershipResponse) GetAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *UserServiceMembershipResponse) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *UserServiceMembershipResponse) SetAdmin(v bool) {
	o.Admin = &v
}

func (o UserServiceMembershipResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserServiceMembershipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	return toSerialize, nil
}

type NullableUserServiceMembershipResponse struct {
	value *UserServiceMembershipResponse
	isSet bool
}

func (v NullableUserServiceMembershipResponse) Get() *UserServiceMembershipResponse {
	return v.value
}

func (v *NullableUserServiceMembershipResponse) Set(val *UserServiceMembershipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserServiceMembershipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserServiceMembershipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserServiceMembershipResponse(val *UserServiceMembershipResponse) *NullableUserServiceMembershipResponse {
	return &NullableUserServiceMembershipResponse{value: val, isSet: true}
}

func (v NullableUserServiceMembershipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserServiceMembershipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
