/*
OpenAPI definition

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// AuthleteUser struct for AuthleteUser
type AuthleteUser struct {
	Id *int64 `json:"id,omitempty"`
	Xid *int64 `json:"xid,omitempty"`
	GivenName *string `json:"givenName,omitempty"`
	FamilyName *string `json:"familyName,omitempty"`
	Email *string `json:"email,omitempty"`
	LastAuthTime *time.Time `json:"lastAuthTime,omitempty"`
	Admin *bool `json:"admin,omitempty"`
	WebAuthenticators *map[string]string `json:"webAuthenticators,omitempty"`
	LastAuthTimeAsLong *int64 `json:"lastAuthTimeAsLong,omitempty"`
	WebAuthenticatorIDs []string `json:"webAuthenticatorIDs,omitempty"`
	TotpEnabled *bool `json:"totpEnabled,omitempty"`
	WebAuthnEnabled *bool `json:"webAuthnEnabled,omitempty"`
}

// NewAuthleteUser instantiates a new AuthleteUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthleteUser() *AuthleteUser {
	this := AuthleteUser{}
	return &this
}

// NewAuthleteUserWithDefaults instantiates a new AuthleteUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthleteUserWithDefaults() *AuthleteUser {
	this := AuthleteUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthleteUser) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthleteUser) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AuthleteUser) SetId(v int64) {
	o.Id = &v
}

// GetXid returns the Xid field value if set, zero value otherwise.
func (o *AuthleteUser) GetXid() int64 {
	if o == nil || isNil(o.Xid) {
		var ret int64
		return ret
	}
	return *o.Xid
}

// GetXidOk returns a tuple with the Xid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetXidOk() (*int64, bool) {
	if o == nil || isNil(o.Xid) {
    return nil, false
	}
	return o.Xid, true
}

// HasXid returns a boolean if a field has been set.
func (o *AuthleteUser) HasXid() bool {
	if o != nil && !isNil(o.Xid) {
		return true
	}

	return false
}

// SetXid gets a reference to the given int64 and assigns it to the Xid field.
func (o *AuthleteUser) SetXid(v int64) {
	o.Xid = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *AuthleteUser) GetGivenName() string {
	if o == nil || isNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetGivenNameOk() (*string, bool) {
	if o == nil || isNil(o.GivenName) {
    return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *AuthleteUser) HasGivenName() bool {
	if o != nil && !isNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *AuthleteUser) SetGivenName(v string) {
	o.GivenName = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *AuthleteUser) GetFamilyName() string {
	if o == nil || isNil(o.FamilyName) {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetFamilyNameOk() (*string, bool) {
	if o == nil || isNil(o.FamilyName) {
    return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *AuthleteUser) HasFamilyName() bool {
	if o != nil && !isNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *AuthleteUser) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AuthleteUser) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AuthleteUser) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AuthleteUser) SetEmail(v string) {
	o.Email = &v
}

// GetLastAuthTime returns the LastAuthTime field value if set, zero value otherwise.
func (o *AuthleteUser) GetLastAuthTime() time.Time {
	if o == nil || isNil(o.LastAuthTime) {
		var ret time.Time
		return ret
	}
	return *o.LastAuthTime
}

// GetLastAuthTimeOk returns a tuple with the LastAuthTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetLastAuthTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastAuthTime) {
    return nil, false
	}
	return o.LastAuthTime, true
}

// HasLastAuthTime returns a boolean if a field has been set.
func (o *AuthleteUser) HasLastAuthTime() bool {
	if o != nil && !isNil(o.LastAuthTime) {
		return true
	}

	return false
}

// SetLastAuthTime gets a reference to the given time.Time and assigns it to the LastAuthTime field.
func (o *AuthleteUser) SetLastAuthTime(v time.Time) {
	o.LastAuthTime = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *AuthleteUser) GetAdmin() bool {
	if o == nil || isNil(o.Admin) {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetAdminOk() (*bool, bool) {
	if o == nil || isNil(o.Admin) {
    return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *AuthleteUser) HasAdmin() bool {
	if o != nil && !isNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *AuthleteUser) SetAdmin(v bool) {
	o.Admin = &v
}

// GetWebAuthenticators returns the WebAuthenticators field value if set, zero value otherwise.
func (o *AuthleteUser) GetWebAuthenticators() map[string]string {
	if o == nil || isNil(o.WebAuthenticators) {
		var ret map[string]string
		return ret
	}
	return *o.WebAuthenticators
}

// GetWebAuthenticatorsOk returns a tuple with the WebAuthenticators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetWebAuthenticatorsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.WebAuthenticators) {
    return nil, false
	}
	return o.WebAuthenticators, true
}

// HasWebAuthenticators returns a boolean if a field has been set.
func (o *AuthleteUser) HasWebAuthenticators() bool {
	if o != nil && !isNil(o.WebAuthenticators) {
		return true
	}

	return false
}

// SetWebAuthenticators gets a reference to the given map[string]string and assigns it to the WebAuthenticators field.
func (o *AuthleteUser) SetWebAuthenticators(v map[string]string) {
	o.WebAuthenticators = &v
}

// GetLastAuthTimeAsLong returns the LastAuthTimeAsLong field value if set, zero value otherwise.
func (o *AuthleteUser) GetLastAuthTimeAsLong() int64 {
	if o == nil || isNil(o.LastAuthTimeAsLong) {
		var ret int64
		return ret
	}
	return *o.LastAuthTimeAsLong
}

// GetLastAuthTimeAsLongOk returns a tuple with the LastAuthTimeAsLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetLastAuthTimeAsLongOk() (*int64, bool) {
	if o == nil || isNil(o.LastAuthTimeAsLong) {
    return nil, false
	}
	return o.LastAuthTimeAsLong, true
}

// HasLastAuthTimeAsLong returns a boolean if a field has been set.
func (o *AuthleteUser) HasLastAuthTimeAsLong() bool {
	if o != nil && !isNil(o.LastAuthTimeAsLong) {
		return true
	}

	return false
}

// SetLastAuthTimeAsLong gets a reference to the given int64 and assigns it to the LastAuthTimeAsLong field.
func (o *AuthleteUser) SetLastAuthTimeAsLong(v int64) {
	o.LastAuthTimeAsLong = &v
}

// GetWebAuthenticatorIDs returns the WebAuthenticatorIDs field value if set, zero value otherwise.
func (o *AuthleteUser) GetWebAuthenticatorIDs() []string {
	if o == nil || isNil(o.WebAuthenticatorIDs) {
		var ret []string
		return ret
	}
	return o.WebAuthenticatorIDs
}

// GetWebAuthenticatorIDsOk returns a tuple with the WebAuthenticatorIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetWebAuthenticatorIDsOk() ([]string, bool) {
	if o == nil || isNil(o.WebAuthenticatorIDs) {
    return nil, false
	}
	return o.WebAuthenticatorIDs, true
}

// HasWebAuthenticatorIDs returns a boolean if a field has been set.
func (o *AuthleteUser) HasWebAuthenticatorIDs() bool {
	if o != nil && !isNil(o.WebAuthenticatorIDs) {
		return true
	}

	return false
}

// SetWebAuthenticatorIDs gets a reference to the given []string and assigns it to the WebAuthenticatorIDs field.
func (o *AuthleteUser) SetWebAuthenticatorIDs(v []string) {
	o.WebAuthenticatorIDs = v
}

// GetTotpEnabled returns the TotpEnabled field value if set, zero value otherwise.
func (o *AuthleteUser) GetTotpEnabled() bool {
	if o == nil || isNil(o.TotpEnabled) {
		var ret bool
		return ret
	}
	return *o.TotpEnabled
}

// GetTotpEnabledOk returns a tuple with the TotpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetTotpEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.TotpEnabled) {
    return nil, false
	}
	return o.TotpEnabled, true
}

// HasTotpEnabled returns a boolean if a field has been set.
func (o *AuthleteUser) HasTotpEnabled() bool {
	if o != nil && !isNil(o.TotpEnabled) {
		return true
	}

	return false
}

// SetTotpEnabled gets a reference to the given bool and assigns it to the TotpEnabled field.
func (o *AuthleteUser) SetTotpEnabled(v bool) {
	o.TotpEnabled = &v
}

// GetWebAuthnEnabled returns the WebAuthnEnabled field value if set, zero value otherwise.
func (o *AuthleteUser) GetWebAuthnEnabled() bool {
	if o == nil || isNil(o.WebAuthnEnabled) {
		var ret bool
		return ret
	}
	return *o.WebAuthnEnabled
}

// GetWebAuthnEnabledOk returns a tuple with the WebAuthnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthleteUser) GetWebAuthnEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.WebAuthnEnabled) {
    return nil, false
	}
	return o.WebAuthnEnabled, true
}

// HasWebAuthnEnabled returns a boolean if a field has been set.
func (o *AuthleteUser) HasWebAuthnEnabled() bool {
	if o != nil && !isNil(o.WebAuthnEnabled) {
		return true
	}

	return false
}

// SetWebAuthnEnabled gets a reference to the given bool and assigns it to the WebAuthnEnabled field.
func (o *AuthleteUser) SetWebAuthnEnabled(v bool) {
	o.WebAuthnEnabled = &v
}

func (o AuthleteUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Xid) {
		toSerialize["xid"] = o.Xid
	}
	if !isNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !isNil(o.FamilyName) {
		toSerialize["familyName"] = o.FamilyName
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.LastAuthTime) {
		toSerialize["lastAuthTime"] = o.LastAuthTime
	}
	if !isNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !isNil(o.WebAuthenticators) {
		toSerialize["webAuthenticators"] = o.WebAuthenticators
	}
	if !isNil(o.LastAuthTimeAsLong) {
		toSerialize["lastAuthTimeAsLong"] = o.LastAuthTimeAsLong
	}
	if !isNil(o.WebAuthenticatorIDs) {
		toSerialize["webAuthenticatorIDs"] = o.WebAuthenticatorIDs
	}
	if !isNil(o.TotpEnabled) {
		toSerialize["totpEnabled"] = o.TotpEnabled
	}
	if !isNil(o.WebAuthnEnabled) {
		toSerialize["webAuthnEnabled"] = o.WebAuthnEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableAuthleteUser struct {
	value *AuthleteUser
	isSet bool
}

func (v NullableAuthleteUser) Get() *AuthleteUser {
	return v.value
}

func (v *NullableAuthleteUser) Set(val *AuthleteUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthleteUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthleteUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthleteUser(val *AuthleteUser) *NullableAuthleteUser {
	return &NullableAuthleteUser{value: val, isSet: true}
}

func (v NullableAuthleteUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthleteUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


