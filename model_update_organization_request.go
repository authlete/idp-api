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

// checks if the UpdateOrganizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationRequest{}

// UpdateOrganizationRequest struct for UpdateOrganizationRequest
type UpdateOrganizationRequest struct {
	Plan *string `json:"plan,omitempty"`
	Name string  `json:"name"`
}

type _UpdateOrganizationRequest UpdateOrganizationRequest

// NewUpdateOrganizationRequest instantiates a new UpdateOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationRequest(name string) *UpdateOrganizationRequest {
	this := UpdateOrganizationRequest{}
	this.Name = name
	return &this
}

// NewUpdateOrganizationRequestWithDefaults instantiates a new UpdateOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationRequestWithDefaults() *UpdateOrganizationRequest {
	this := UpdateOrganizationRequest{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UpdateOrganizationRequest) GetPlan() string {
	if o == nil || IsNil(o.Plan) {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationRequest) GetPlanOk() (*string, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UpdateOrganizationRequest) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *UpdateOrganizationRequest) SetPlan(v string) {
	o.Plan = &v
}

// GetName returns the Name field value
func (o *UpdateOrganizationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateOrganizationRequest) SetName(v string) {
	o.Name = v
}

func (o UpdateOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *UpdateOrganizationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varUpdateOrganizationRequest := _UpdateOrganizationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateOrganizationRequest)

	if err != nil {
		return err
	}

	*o = UpdateOrganizationRequest(varUpdateOrganizationRequest)

	return err
}

type NullableUpdateOrganizationRequest struct {
	value *UpdateOrganizationRequest
	isSet bool
}

func (v NullableUpdateOrganizationRequest) Get() *UpdateOrganizationRequest {
	return v.value
}

func (v *NullableUpdateOrganizationRequest) Set(val *UpdateOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationRequest(val *UpdateOrganizationRequest) *NullableUpdateOrganizationRequest {
	return &NullableUpdateOrganizationRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
