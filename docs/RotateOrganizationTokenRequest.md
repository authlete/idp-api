# RotateOrganizationTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **int64** |  | 
**TokenId** | **string** |  | 

## Methods

### NewRotateOrganizationTokenRequest

`func NewRotateOrganizationTokenRequest(organizationId int64, tokenId string, ) *RotateOrganizationTokenRequest`

NewRotateOrganizationTokenRequest instantiates a new RotateOrganizationTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateOrganizationTokenRequestWithDefaults

`func NewRotateOrganizationTokenRequestWithDefaults() *RotateOrganizationTokenRequest`

NewRotateOrganizationTokenRequestWithDefaults instantiates a new RotateOrganizationTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *RotateOrganizationTokenRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RotateOrganizationTokenRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RotateOrganizationTokenRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetTokenId

`func (o *RotateOrganizationTokenRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *RotateOrganizationTokenRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *RotateOrganizationTokenRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


