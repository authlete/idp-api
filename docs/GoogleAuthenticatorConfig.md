# GoogleAuthenticatorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStepSizeInMillis** | Pointer to **int64** |  | [optional] 
**WindowSize** | Pointer to **int32** |  | [optional] 
**CodeDigits** | Pointer to **int32** |  | [optional] 
**NumberOfScratchCodes** | Pointer to **int32** |  | [optional] 
**KeyModulus** | Pointer to **int32** |  | [optional] 
**SecretBits** | Pointer to **int32** |  | [optional] 
**KeyRepresentation** | Pointer to **string** |  | [optional] 
**HmacHashFunction** | Pointer to **string** |  | [optional] 

## Methods

### NewGoogleAuthenticatorConfig

`func NewGoogleAuthenticatorConfig() *GoogleAuthenticatorConfig`

NewGoogleAuthenticatorConfig instantiates a new GoogleAuthenticatorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleAuthenticatorConfigWithDefaults

`func NewGoogleAuthenticatorConfigWithDefaults() *GoogleAuthenticatorConfig`

NewGoogleAuthenticatorConfigWithDefaults instantiates a new GoogleAuthenticatorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStepSizeInMillis

`func (o *GoogleAuthenticatorConfig) GetTimeStepSizeInMillis() int64`

GetTimeStepSizeInMillis returns the TimeStepSizeInMillis field if non-nil, zero value otherwise.

### GetTimeStepSizeInMillisOk

`func (o *GoogleAuthenticatorConfig) GetTimeStepSizeInMillisOk() (*int64, bool)`

GetTimeStepSizeInMillisOk returns a tuple with the TimeStepSizeInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStepSizeInMillis

`func (o *GoogleAuthenticatorConfig) SetTimeStepSizeInMillis(v int64)`

SetTimeStepSizeInMillis sets TimeStepSizeInMillis field to given value.

### HasTimeStepSizeInMillis

`func (o *GoogleAuthenticatorConfig) HasTimeStepSizeInMillis() bool`

HasTimeStepSizeInMillis returns a boolean if a field has been set.

### GetWindowSize

`func (o *GoogleAuthenticatorConfig) GetWindowSize() int32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *GoogleAuthenticatorConfig) GetWindowSizeOk() (*int32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *GoogleAuthenticatorConfig) SetWindowSize(v int32)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *GoogleAuthenticatorConfig) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetCodeDigits

`func (o *GoogleAuthenticatorConfig) GetCodeDigits() int32`

GetCodeDigits returns the CodeDigits field if non-nil, zero value otherwise.

### GetCodeDigitsOk

`func (o *GoogleAuthenticatorConfig) GetCodeDigitsOk() (*int32, bool)`

GetCodeDigitsOk returns a tuple with the CodeDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeDigits

`func (o *GoogleAuthenticatorConfig) SetCodeDigits(v int32)`

SetCodeDigits sets CodeDigits field to given value.

### HasCodeDigits

`func (o *GoogleAuthenticatorConfig) HasCodeDigits() bool`

HasCodeDigits returns a boolean if a field has been set.

### GetNumberOfScratchCodes

`func (o *GoogleAuthenticatorConfig) GetNumberOfScratchCodes() int32`

GetNumberOfScratchCodes returns the NumberOfScratchCodes field if non-nil, zero value otherwise.

### GetNumberOfScratchCodesOk

`func (o *GoogleAuthenticatorConfig) GetNumberOfScratchCodesOk() (*int32, bool)`

GetNumberOfScratchCodesOk returns a tuple with the NumberOfScratchCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfScratchCodes

`func (o *GoogleAuthenticatorConfig) SetNumberOfScratchCodes(v int32)`

SetNumberOfScratchCodes sets NumberOfScratchCodes field to given value.

### HasNumberOfScratchCodes

`func (o *GoogleAuthenticatorConfig) HasNumberOfScratchCodes() bool`

HasNumberOfScratchCodes returns a boolean if a field has been set.

### GetKeyModulus

`func (o *GoogleAuthenticatorConfig) GetKeyModulus() int32`

GetKeyModulus returns the KeyModulus field if non-nil, zero value otherwise.

### GetKeyModulusOk

`func (o *GoogleAuthenticatorConfig) GetKeyModulusOk() (*int32, bool)`

GetKeyModulusOk returns a tuple with the KeyModulus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyModulus

`func (o *GoogleAuthenticatorConfig) SetKeyModulus(v int32)`

SetKeyModulus sets KeyModulus field to given value.

### HasKeyModulus

`func (o *GoogleAuthenticatorConfig) HasKeyModulus() bool`

HasKeyModulus returns a boolean if a field has been set.

### GetSecretBits

`func (o *GoogleAuthenticatorConfig) GetSecretBits() int32`

GetSecretBits returns the SecretBits field if non-nil, zero value otherwise.

### GetSecretBitsOk

`func (o *GoogleAuthenticatorConfig) GetSecretBitsOk() (*int32, bool)`

GetSecretBitsOk returns a tuple with the SecretBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretBits

`func (o *GoogleAuthenticatorConfig) SetSecretBits(v int32)`

SetSecretBits sets SecretBits field to given value.

### HasSecretBits

`func (o *GoogleAuthenticatorConfig) HasSecretBits() bool`

HasSecretBits returns a boolean if a field has been set.

### GetKeyRepresentation

`func (o *GoogleAuthenticatorConfig) GetKeyRepresentation() string`

GetKeyRepresentation returns the KeyRepresentation field if non-nil, zero value otherwise.

### GetKeyRepresentationOk

`func (o *GoogleAuthenticatorConfig) GetKeyRepresentationOk() (*string, bool)`

GetKeyRepresentationOk returns a tuple with the KeyRepresentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRepresentation

`func (o *GoogleAuthenticatorConfig) SetKeyRepresentation(v string)`

SetKeyRepresentation sets KeyRepresentation field to given value.

### HasKeyRepresentation

`func (o *GoogleAuthenticatorConfig) HasKeyRepresentation() bool`

HasKeyRepresentation returns a boolean if a field has been set.

### GetHmacHashFunction

`func (o *GoogleAuthenticatorConfig) GetHmacHashFunction() string`

GetHmacHashFunction returns the HmacHashFunction field if non-nil, zero value otherwise.

### GetHmacHashFunctionOk

`func (o *GoogleAuthenticatorConfig) GetHmacHashFunctionOk() (*string, bool)`

GetHmacHashFunctionOk returns a tuple with the HmacHashFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacHashFunction

`func (o *GoogleAuthenticatorConfig) SetHmacHashFunction(v string)`

SetHmacHashFunction sets HmacHashFunction field to given value.

### HasHmacHashFunction

`func (o *GoogleAuthenticatorConfig) HasHmacHashFunction() bool`

HasHmacHashFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


