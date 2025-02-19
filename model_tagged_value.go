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

// checks if the TaggedValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaggedValue{}

// TaggedValue struct for TaggedValue
type TaggedValue struct {
	Tag   *string `json:"tag,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewTaggedValue instantiates a new TaggedValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaggedValue() *TaggedValue {
	this := TaggedValue{}
	return &this
}

// NewTaggedValueWithDefaults instantiates a new TaggedValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaggedValueWithDefaults() *TaggedValue {
	this := TaggedValue{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *TaggedValue) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaggedValue) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *TaggedValue) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *TaggedValue) SetTag(v string) {
	o.Tag = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TaggedValue) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaggedValue) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TaggedValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TaggedValue) SetValue(v string) {
	o.Value = &v
}

func (o TaggedValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaggedValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTaggedValue struct {
	value *TaggedValue
	isSet bool
}

func (v NullableTaggedValue) Get() *TaggedValue {
	return v.value
}

func (v *NullableTaggedValue) Set(val *TaggedValue) {
	v.value = val
	v.isSet = true
}

func (v NullableTaggedValue) IsSet() bool {
	return v.isSet
}

func (v *NullableTaggedValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaggedValue(val *TaggedValue) *NullableTaggedValue {
	return &NullableTaggedValue{value: val, isSet: true}
}

func (v NullableTaggedValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaggedValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
