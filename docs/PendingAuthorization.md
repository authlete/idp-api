# PendingAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationResponse** | Pointer to [**AuthorizationResponse**](AuthorizationResponse.md) |  | [optional] 
**Audience** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 

## Methods

### NewPendingAuthorization

`func NewPendingAuthorization() *PendingAuthorization`

NewPendingAuthorization instantiates a new PendingAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingAuthorizationWithDefaults

`func NewPendingAuthorizationWithDefaults() *PendingAuthorization`

NewPendingAuthorizationWithDefaults instantiates a new PendingAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationResponse

`func (o *PendingAuthorization) GetAuthorizationResponse() AuthorizationResponse`

GetAuthorizationResponse returns the AuthorizationResponse field if non-nil, zero value otherwise.

### GetAuthorizationResponseOk

`func (o *PendingAuthorization) GetAuthorizationResponseOk() (*AuthorizationResponse, bool)`

GetAuthorizationResponseOk returns a tuple with the AuthorizationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationResponse

`func (o *PendingAuthorization) SetAuthorizationResponse(v AuthorizationResponse)`

SetAuthorizationResponse sets AuthorizationResponse field to given value.

### HasAuthorizationResponse

`func (o *PendingAuthorization) HasAuthorizationResponse() bool`

HasAuthorizationResponse returns a boolean if a field has been set.

### GetAudience

`func (o *PendingAuthorization) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *PendingAuthorization) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *PendingAuthorization) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *PendingAuthorization) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetResource

`func (o *PendingAuthorization) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PendingAuthorization) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PendingAuthorization) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PendingAuthorization) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


