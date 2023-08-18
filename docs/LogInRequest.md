# LogInRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**Totp** | Pointer to **int32** |  | [optional] 
**Webauthn** | Pointer to [**WebAuthnChallengeRequest**](WebAuthnChallengeRequest.md) |  | [optional] 

## Methods

### NewLogInRequest

`func NewLogInRequest(username string, password string, ) *LogInRequest`

NewLogInRequest instantiates a new LogInRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogInRequestWithDefaults

`func NewLogInRequestWithDefaults() *LogInRequest`

NewLogInRequestWithDefaults instantiates a new LogInRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *LogInRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LogInRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LogInRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *LogInRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LogInRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LogInRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTotp

`func (o *LogInRequest) GetTotp() int32`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *LogInRequest) GetTotpOk() (*int32, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *LogInRequest) SetTotp(v int32)`

SetTotp sets Totp field to given value.

### HasTotp

`func (o *LogInRequest) HasTotp() bool`

HasTotp returns a boolean if a field has been set.

### GetWebauthn

`func (o *LogInRequest) GetWebauthn() WebAuthnChallengeRequest`

GetWebauthn returns the Webauthn field if non-nil, zero value otherwise.

### GetWebauthnOk

`func (o *LogInRequest) GetWebauthnOk() (*WebAuthnChallengeRequest, bool)`

GetWebauthnOk returns a tuple with the Webauthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthn

`func (o *LogInRequest) SetWebauthn(v WebAuthnChallengeRequest)`

SetWebauthn sets Webauthn field to given value.

### HasWebauthn

`func (o *LogInRequest) HasWebauthn() bool`

HasWebauthn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


