# RenameServiceTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **int64** |  | 
**OrganizationId** | **int64** |  | 
**ApiServerId** | **int64** |  | 
**TokenId** | **string** |  | 
**Description** | **string** |  | 

## Methods

### NewRenameServiceTokenRequest

`func NewRenameServiceTokenRequest(serviceId int64, organizationId int64, apiServerId int64, tokenId string, description string, ) *RenameServiceTokenRequest`

NewRenameServiceTokenRequest instantiates a new RenameServiceTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameServiceTokenRequestWithDefaults

`func NewRenameServiceTokenRequestWithDefaults() *RenameServiceTokenRequest`

NewRenameServiceTokenRequestWithDefaults instantiates a new RenameServiceTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *RenameServiceTokenRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RenameServiceTokenRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RenameServiceTokenRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetOrganizationId

`func (o *RenameServiceTokenRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RenameServiceTokenRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RenameServiceTokenRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetApiServerId

`func (o *RenameServiceTokenRequest) GetApiServerId() int64`

GetApiServerId returns the ApiServerId field if non-nil, zero value otherwise.

### GetApiServerIdOk

`func (o *RenameServiceTokenRequest) GetApiServerIdOk() (*int64, bool)`

GetApiServerIdOk returns a tuple with the ApiServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerId

`func (o *RenameServiceTokenRequest) SetApiServerId(v int64)`

SetApiServerId sets ApiServerId field to given value.


### GetTokenId

`func (o *RenameServiceTokenRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *RenameServiceTokenRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *RenameServiceTokenRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDescription

`func (o *RenameServiceTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RenameServiceTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RenameServiceTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


