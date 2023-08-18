# DeleteOrganizationTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **int64** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteOrganizationTokenRequest

`func NewDeleteOrganizationTokenRequest() *DeleteOrganizationTokenRequest`

NewDeleteOrganizationTokenRequest instantiates a new DeleteOrganizationTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteOrganizationTokenRequestWithDefaults

`func NewDeleteOrganizationTokenRequestWithDefaults() *DeleteOrganizationTokenRequest`

NewDeleteOrganizationTokenRequestWithDefaults instantiates a new DeleteOrganizationTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *DeleteOrganizationTokenRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DeleteOrganizationTokenRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DeleteOrganizationTokenRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DeleteOrganizationTokenRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetTokenId

`func (o *DeleteOrganizationTokenRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *DeleteOrganizationTokenRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *DeleteOrganizationTokenRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *DeleteOrganizationTokenRequest) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetDescription

`func (o *DeleteOrganizationTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeleteOrganizationTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeleteOrganizationTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeleteOrganizationTokenRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


