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

// LogInRequest struct for LogInRequest
type LogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Totp *int32 `json:"totp,omitempty"`
	Webauthn *WebAuthnChallengeRequest `json:"webauthn,omitempty"`
}

// NewLogInRequest instantiates a new LogInRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogInRequest(username string, password string) *LogInRequest {
	this := LogInRequest{}
	this.Username = username
	this.Password = password
	return &this
}

// NewLogInRequestWithDefaults instantiates a new LogInRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogInRequestWithDefaults() *LogInRequest {
	this := LogInRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *LogInRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *LogInRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *LogInRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *LogInRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *LogInRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *LogInRequest) SetPassword(v string) {
	o.Password = v
}

// GetTotp returns the Totp field value if set, zero value otherwise.
func (o *LogInRequest) GetTotp() int32 {
	if o == nil || isNil(o.Totp) {
		var ret int32
		return ret
	}
	return *o.Totp
}

// GetTotpOk returns a tuple with the Totp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogInRequest) GetTotpOk() (*int32, bool) {
	if o == nil || isNil(o.Totp) {
    return nil, false
	}
	return o.Totp, true
}

// HasTotp returns a boolean if a field has been set.
func (o *LogInRequest) HasTotp() bool {
	if o != nil && !isNil(o.Totp) {
		return true
	}

	return false
}

// SetTotp gets a reference to the given int32 and assigns it to the Totp field.
func (o *LogInRequest) SetTotp(v int32) {
	o.Totp = &v
}

// GetWebauthn returns the Webauthn field value if set, zero value otherwise.
func (o *LogInRequest) GetWebauthn() WebAuthnChallengeRequest {
	if o == nil || isNil(o.Webauthn) {
		var ret WebAuthnChallengeRequest
		return ret
	}
	return *o.Webauthn
}

// GetWebauthnOk returns a tuple with the Webauthn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogInRequest) GetWebauthnOk() (*WebAuthnChallengeRequest, bool) {
	if o == nil || isNil(o.Webauthn) {
    return nil, false
	}
	return o.Webauthn, true
}

// HasWebauthn returns a boolean if a field has been set.
func (o *LogInRequest) HasWebauthn() bool {
	if o != nil && !isNil(o.Webauthn) {
		return true
	}

	return false
}

// SetWebauthn gets a reference to the given WebAuthnChallengeRequest and assigns it to the Webauthn field.
func (o *LogInRequest) SetWebauthn(v WebAuthnChallengeRequest) {
	o.Webauthn = &v
}

func (o LogInRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.Totp) {
		toSerialize["totp"] = o.Totp
	}
	if !isNil(o.Webauthn) {
		toSerialize["webauthn"] = o.Webauthn
	}
	return json.Marshal(toSerialize)
}

type NullableLogInRequest struct {
	value *LogInRequest
	isSet bool
}

func (v NullableLogInRequest) Get() *LogInRequest {
	return v.value
}

func (v *NullableLogInRequest) Set(val *LogInRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLogInRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLogInRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogInRequest(val *LogInRequest) *NullableLogInRequest {
	return &NullableLogInRequest{value: val, isSet: true}
}

func (v NullableLogInRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogInRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


