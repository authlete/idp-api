# DeleteServiceTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **int64** |  | [optional] 
**OrganizationId** | Pointer to **int64** |  | [optional] 
**ApiServerId** | Pointer to **int64** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteServiceTokenRequest

`func NewDeleteServiceTokenRequest() *DeleteServiceTokenRequest`

NewDeleteServiceTokenRequest instantiates a new DeleteServiceTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteServiceTokenRequestWithDefaults

`func NewDeleteServiceTokenRequestWithDefaults() *DeleteServiceTokenRequest`

NewDeleteServiceTokenRequestWithDefaults instantiates a new DeleteServiceTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *DeleteServiceTokenRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeleteServiceTokenRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeleteServiceTokenRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DeleteServiceTokenRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *DeleteServiceTokenRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DeleteServiceTokenRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DeleteServiceTokenRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DeleteServiceTokenRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetApiServerId

`func (o *DeleteServiceTokenRequest) GetApiServerId() int64`

GetApiServerId returns the ApiServerId field if non-nil, zero value otherwise.

### GetApiServerIdOk

`func (o *DeleteServiceTokenRequest) GetApiServerIdOk() (*int64, bool)`

GetApiServerIdOk returns a tuple with the ApiServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerId

`func (o *DeleteServiceTokenRequest) SetApiServerId(v int64)`

SetApiServerId sets ApiServerId field to given value.

### HasApiServerId

`func (o *DeleteServiceTokenRequest) HasApiServerId() bool`

HasApiServerId returns a boolean if a field has been set.

### GetTokenId

`func (o *DeleteServiceTokenRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *DeleteServiceTokenRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *DeleteServiceTokenRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *DeleteServiceTokenRequest) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetDescription

`func (o *DeleteServiceTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeleteServiceTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeleteServiceTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeleteServiceTokenRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


