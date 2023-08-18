# AuthleteUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Xid** | Pointer to **int64** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**FamilyName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**LastAuthTime** | Pointer to **time.Time** |  | [optional] 
**Admin** | Pointer to **bool** |  | [optional] 
**WebAuthenticators** | Pointer to **map[string]string** |  | [optional] 
**LastAuthTimeAsLong** | Pointer to **int64** |  | [optional] 
**WebAuthenticatorIDs** | Pointer to **[]string** |  | [optional] 
**TotpEnabled** | Pointer to **bool** |  | [optional] 
**WebAuthnEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthleteUser

`func NewAuthleteUser() *AuthleteUser`

NewAuthleteUser instantiates a new AuthleteUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthleteUserWithDefaults

`func NewAuthleteUserWithDefaults() *AuthleteUser`

NewAuthleteUserWithDefaults instantiates a new AuthleteUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthleteUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthleteUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthleteUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AuthleteUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetXid

`func (o *AuthleteUser) GetXid() int64`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *AuthleteUser) GetXidOk() (*int64, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *AuthleteUser) SetXid(v int64)`

SetXid sets Xid field to given value.

### HasXid

`func (o *AuthleteUser) HasXid() bool`

HasXid returns a boolean if a field has been set.

### GetGivenName

`func (o *AuthleteUser) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *AuthleteUser) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *AuthleteUser) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *AuthleteUser) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetFamilyName

`func (o *AuthleteUser) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *AuthleteUser) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *AuthleteUser) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *AuthleteUser) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetEmail

`func (o *AuthleteUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthleteUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthleteUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthleteUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLastAuthTime

`func (o *AuthleteUser) GetLastAuthTime() time.Time`

GetLastAuthTime returns the LastAuthTime field if non-nil, zero value otherwise.

### GetLastAuthTimeOk

`func (o *AuthleteUser) GetLastAuthTimeOk() (*time.Time, bool)`

GetLastAuthTimeOk returns a tuple with the LastAuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthTime

`func (o *AuthleteUser) SetLastAuthTime(v time.Time)`

SetLastAuthTime sets LastAuthTime field to given value.

### HasLastAuthTime

`func (o *AuthleteUser) HasLastAuthTime() bool`

HasLastAuthTime returns a boolean if a field has been set.

### GetAdmin

`func (o *AuthleteUser) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *AuthleteUser) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *AuthleteUser) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *AuthleteUser) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetWebAuthenticators

`func (o *AuthleteUser) GetWebAuthenticators() map[string]string`

GetWebAuthenticators returns the WebAuthenticators field if non-nil, zero value otherwise.

### GetWebAuthenticatorsOk

`func (o *AuthleteUser) GetWebAuthenticatorsOk() (*map[string]string, bool)`

GetWebAuthenticatorsOk returns a tuple with the WebAuthenticators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthenticators

`func (o *AuthleteUser) SetWebAuthenticators(v map[string]string)`

SetWebAuthenticators sets WebAuthenticators field to given value.

### HasWebAuthenticators

`func (o *AuthleteUser) HasWebAuthenticators() bool`

HasWebAuthenticators returns a boolean if a field has been set.

### GetLastAuthTimeAsLong

`func (o *AuthleteUser) GetLastAuthTimeAsLong() int64`

GetLastAuthTimeAsLong returns the LastAuthTimeAsLong field if non-nil, zero value otherwise.

### GetLastAuthTimeAsLongOk

`func (o *AuthleteUser) GetLastAuthTimeAsLongOk() (*int64, bool)`

GetLastAuthTimeAsLongOk returns a tuple with the LastAuthTimeAsLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthTimeAsLong

`func (o *AuthleteUser) SetLastAuthTimeAsLong(v int64)`

SetLastAuthTimeAsLong sets LastAuthTimeAsLong field to given value.

### HasLastAuthTimeAsLong

`func (o *AuthleteUser) HasLastAuthTimeAsLong() bool`

HasLastAuthTimeAsLong returns a boolean if a field has been set.

### GetWebAuthenticatorIDs

`func (o *AuthleteUser) GetWebAuthenticatorIDs() []string`

GetWebAuthenticatorIDs returns the WebAuthenticatorIDs field if non-nil, zero value otherwise.

### GetWebAuthenticatorIDsOk

`func (o *AuthleteUser) GetWebAuthenticatorIDsOk() (*[]string, bool)`

GetWebAuthenticatorIDsOk returns a tuple with the WebAuthenticatorIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthenticatorIDs

`func (o *AuthleteUser) SetWebAuthenticatorIDs(v []string)`

SetWebAuthenticatorIDs sets WebAuthenticatorIDs field to given value.

### HasWebAuthenticatorIDs

`func (o *AuthleteUser) HasWebAuthenticatorIDs() bool`

HasWebAuthenticatorIDs returns a boolean if a field has been set.

### GetTotpEnabled

`func (o *AuthleteUser) GetTotpEnabled() bool`

GetTotpEnabled returns the TotpEnabled field if non-nil, zero value otherwise.

### GetTotpEnabledOk

`func (o *AuthleteUser) GetTotpEnabledOk() (*bool, bool)`

GetTotpEnabledOk returns a tuple with the TotpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpEnabled

`func (o *AuthleteUser) SetTotpEnabled(v bool)`

SetTotpEnabled sets TotpEnabled field to given value.

### HasTotpEnabled

`func (o *AuthleteUser) HasTotpEnabled() bool`

HasTotpEnabled returns a boolean if a field has been set.

### GetWebAuthnEnabled

`func (o *AuthleteUser) GetWebAuthnEnabled() bool`

GetWebAuthnEnabled returns the WebAuthnEnabled field if non-nil, zero value otherwise.

### GetWebAuthnEnabledOk

`func (o *AuthleteUser) GetWebAuthnEnabledOk() (*bool, bool)`

GetWebAuthnEnabledOk returns a tuple with the WebAuthnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnEnabled

`func (o *AuthleteUser) SetWebAuthnEnabled(v bool)`

SetWebAuthnEnabled sets WebAuthnEnabled field to given value.

### HasWebAuthnEnabled

`func (o *AuthleteUser) HasWebAuthnEnabled() bool`

HasWebAuthnEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


