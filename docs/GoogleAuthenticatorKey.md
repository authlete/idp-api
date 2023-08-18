# GoogleAuthenticatorKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**GoogleAuthenticatorConfig**](GoogleAuthenticatorConfig.md) |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**VerificationCode** | Pointer to **int32** |  | [optional] 
**ScratchCodes** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewGoogleAuthenticatorKey

`func NewGoogleAuthenticatorKey() *GoogleAuthenticatorKey`

NewGoogleAuthenticatorKey instantiates a new GoogleAuthenticatorKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleAuthenticatorKeyWithDefaults

`func NewGoogleAuthenticatorKeyWithDefaults() *GoogleAuthenticatorKey`

NewGoogleAuthenticatorKeyWithDefaults instantiates a new GoogleAuthenticatorKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *GoogleAuthenticatorKey) GetConfig() GoogleAuthenticatorConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GoogleAuthenticatorKey) GetConfigOk() (*GoogleAuthenticatorConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GoogleAuthenticatorKey) SetConfig(v GoogleAuthenticatorConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GoogleAuthenticatorKey) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetKey

`func (o *GoogleAuthenticatorKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GoogleAuthenticatorKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GoogleAuthenticatorKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GoogleAuthenticatorKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetVerificationCode

`func (o *GoogleAuthenticatorKey) GetVerificationCode() int32`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *GoogleAuthenticatorKey) GetVerificationCodeOk() (*int32, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *GoogleAuthenticatorKey) SetVerificationCode(v int32)`

SetVerificationCode sets VerificationCode field to given value.

### HasVerificationCode

`func (o *GoogleAuthenticatorKey) HasVerificationCode() bool`

HasVerificationCode returns a boolean if a field has been set.

### GetScratchCodes

`func (o *GoogleAuthenticatorKey) GetScratchCodes() []int32`

GetScratchCodes returns the ScratchCodes field if non-nil, zero value otherwise.

### GetScratchCodesOk

`func (o *GoogleAuthenticatorKey) GetScratchCodesOk() (*[]int32, bool)`

GetScratchCodesOk returns a tuple with the ScratchCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScratchCodes

`func (o *GoogleAuthenticatorKey) SetScratchCodes(v []int32)`

SetScratchCodes sets ScratchCodes field to given value.

### HasScratchCodes

`func (o *GoogleAuthenticatorKey) HasScratchCodes() bool`

HasScratchCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


