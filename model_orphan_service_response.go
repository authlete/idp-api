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

// checks if the OrphanServiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrphanServiceResponse{}

// OrphanServiceResponse struct for OrphanServiceResponse
type OrphanServiceResponse struct {
	OnApiButNotOnIdp []OrphanService `json:"onApiButNotOnIdp,omitempty"`
	OnIdPButNotOnApi []OrphanService `json:"onIdPButNotOnApi,omitempty"`
}

// NewOrphanServiceResponse instantiates a new OrphanServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrphanServiceResponse() *OrphanServiceResponse {
	this := OrphanServiceResponse{}
	return &this
}

// NewOrphanServiceResponseWithDefaults instantiates a new OrphanServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrphanServiceResponseWithDefaults() *OrphanServiceResponse {
	this := OrphanServiceResponse{}
	return &this
}

// GetOnApiButNotOnIdp returns the OnApiButNotOnIdp field value if set, zero value otherwise.
func (o *OrphanServiceResponse) GetOnApiButNotOnIdp() []OrphanService {
	if o == nil || IsNil(o.OnApiButNotOnIdp) {
		var ret []OrphanService
		return ret
	}
	return o.OnApiButNotOnIdp
}

// GetOnApiButNotOnIdpOk returns a tuple with the OnApiButNotOnIdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrphanServiceResponse) GetOnApiButNotOnIdpOk() ([]OrphanService, bool) {
	if o == nil || IsNil(o.OnApiButNotOnIdp) {
		return nil, false
	}
	return o.OnApiButNotOnIdp, true
}

// HasOnApiButNotOnIdp returns a boolean if a field has been set.
func (o *OrphanServiceResponse) HasOnApiButNotOnIdp() bool {
	if o != nil && !IsNil(o.OnApiButNotOnIdp) {
		return true
	}

	return false
}

// SetOnApiButNotOnIdp gets a reference to the given []OrphanService and assigns it to the OnApiButNotOnIdp field.
func (o *OrphanServiceResponse) SetOnApiButNotOnIdp(v []OrphanService) {
	o.OnApiButNotOnIdp = v
}

// GetOnIdPButNotOnApi returns the OnIdPButNotOnApi field value if set, zero value otherwise.
func (o *OrphanServiceResponse) GetOnIdPButNotOnApi() []OrphanService {
	if o == nil || IsNil(o.OnIdPButNotOnApi) {
		var ret []OrphanService
		return ret
	}
	return o.OnIdPButNotOnApi
}

// GetOnIdPButNotOnApiOk returns a tuple with the OnIdPButNotOnApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrphanServiceResponse) GetOnIdPButNotOnApiOk() ([]OrphanService, bool) {
	if o == nil || IsNil(o.OnIdPButNotOnApi) {
		return nil, false
	}
	return o.OnIdPButNotOnApi, true
}

// HasOnIdPButNotOnApi returns a boolean if a field has been set.
func (o *OrphanServiceResponse) HasOnIdPButNotOnApi() bool {
	if o != nil && !IsNil(o.OnIdPButNotOnApi) {
		return true
	}

	return false
}

// SetOnIdPButNotOnApi gets a reference to the given []OrphanService and assigns it to the OnIdPButNotOnApi field.
func (o *OrphanServiceResponse) SetOnIdPButNotOnApi(v []OrphanService) {
	o.OnIdPButNotOnApi = v
}

func (o OrphanServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrphanServiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnApiButNotOnIdp) {
		toSerialize["onApiButNotOnIdp"] = o.OnApiButNotOnIdp
	}
	if !IsNil(o.OnIdPButNotOnApi) {
		toSerialize["onIdPButNotOnApi"] = o.OnIdPButNotOnApi
	}
	return toSerialize, nil
}

type NullableOrphanServiceResponse struct {
	value *OrphanServiceResponse
	isSet bool
}

func (v NullableOrphanServiceResponse) Get() *OrphanServiceResponse {
	return v.value
}

func (v *NullableOrphanServiceResponse) Set(val *OrphanServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrphanServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrphanServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrphanServiceResponse(val *OrphanServiceResponse) *NullableOrphanServiceResponse {
	return &NullableOrphanServiceResponse{value: val, isSet: true}
}

func (v NullableOrphanServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrphanServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
