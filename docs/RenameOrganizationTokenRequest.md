# RenameOrganizationTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **int64** |  | 
**TokenId** | **string** |  | 
**Description** | **string** |  | 

## Methods

### NewRenameOrganizationTokenRequest

`func NewRenameOrganizationTokenRequest(organizationId int64, tokenId string, description string, ) *RenameOrganizationTokenRequest`

NewRenameOrganizationTokenRequest instantiates a new RenameOrganizationTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameOrganizationTokenRequestWithDefaults

`func NewRenameOrganizationTokenRequestWithDefaults() *RenameOrganizationTokenRequest`

NewRenameOrganizationTokenRequestWithDefaults instantiates a new RenameOrganizationTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *RenameOrganizationTokenRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RenameOrganizationTokenRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RenameOrganizationTokenRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetTokenId

`func (o *RenameOrganizationTokenRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *RenameOrganizationTokenRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *RenameOrganizationTokenRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDescription

`func (o *RenameOrganizationTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RenameOrganizationTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RenameOrganizationTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


