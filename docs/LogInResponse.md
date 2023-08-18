# LogInResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**FamilyName** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**LastAuthTime** | Pointer to **time.Time** |  | [optional] 
**AuthnMethods** | Pointer to **[]string** |  | [optional] 
**Webauthn** | Pointer to [**WebAuthnChallengeResponse**](WebAuthnChallengeResponse.md) |  | [optional] 

## Methods

### NewLogInResponse

`func NewLogInResponse() *LogInResponse`

NewLogInResponse instantiates a new LogInResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogInResponseWithDefaults

`func NewLogInResponseWithDefaults() *LogInResponse`

NewLogInResponseWithDefaults instantiates a new LogInResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogInResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogInResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogInResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LogInResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGivenName

`func (o *LogInResponse) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *LogInResponse) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *LogInResponse) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *LogInResponse) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetFamilyName

`func (o *LogInResponse) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *LogInResponse) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *LogInResponse) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *LogInResponse) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetDisplayName

`func (o *LogInResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LogInResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LogInResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *LogInResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *LogInResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LogInResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LogInResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LogInResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLastAuthTime

`func (o *LogInResponse) GetLastAuthTime() time.Time`

GetLastAuthTime returns the LastAuthTime field if non-nil, zero value otherwise.

### GetLastAuthTimeOk

`func (o *LogInResponse) GetLastAuthTimeOk() (*time.Time, bool)`

GetLastAuthTimeOk returns a tuple with the LastAuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthTime

`func (o *LogInResponse) SetLastAuthTime(v time.Time)`

SetLastAuthTime sets LastAuthTime field to given value.

### HasLastAuthTime

`func (o *LogInResponse) HasLastAuthTime() bool`

HasLastAuthTime returns a boolean if a field has been set.

### GetAuthnMethods

`func (o *LogInResponse) GetAuthnMethods() []string`

GetAuthnMethods returns the AuthnMethods field if non-nil, zero value otherwise.

### GetAuthnMethodsOk

`func (o *LogInResponse) GetAuthnMethodsOk() (*[]string, bool)`

GetAuthnMethodsOk returns a tuple with the AuthnMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnMethods

`func (o *LogInResponse) SetAuthnMethods(v []string)`

SetAuthnMethods sets AuthnMethods field to given value.

### HasAuthnMethods

`func (o *LogInResponse) HasAuthnMethods() bool`

HasAuthnMethods returns a boolean if a field has been set.

### GetWebauthn

`func (o *LogInResponse) GetWebauthn() WebAuthnChallengeResponse`

GetWebauthn returns the Webauthn field if non-nil, zero value otherwise.

### GetWebauthnOk

`func (o *LogInResponse) GetWebauthnOk() (*WebAuthnChallengeResponse, bool)`

GetWebauthnOk returns a tuple with the Webauthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthn

`func (o *LogInResponse) SetWebauthn(v WebAuthnChallengeResponse)`

SetWebauthn sets Webauthn field to given value.

### HasWebauthn

`func (o *LogInResponse) HasWebauthn() bool`

HasWebauthn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


