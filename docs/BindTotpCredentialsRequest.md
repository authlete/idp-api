# BindTotpCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Totp** | **int32** |  | 

## Methods

### NewBindTotpCredentialsRequest

`func NewBindTotpCredentialsRequest(totp int32, ) *BindTotpCredentialsRequest`

NewBindTotpCredentialsRequest instantiates a new BindTotpCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindTotpCredentialsRequestWithDefaults

`func NewBindTotpCredentialsRequestWithDefaults() *BindTotpCredentialsRequest`

NewBindTotpCredentialsRequestWithDefaults instantiates a new BindTotpCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotp

`func (o *BindTotpCredentialsRequest) GetTotp() int32`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *BindTotpCredentialsRequest) GetTotpOk() (*int32, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *BindTotpCredentialsRequest) SetTotp(v int32)`

SetTotp sets Totp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


