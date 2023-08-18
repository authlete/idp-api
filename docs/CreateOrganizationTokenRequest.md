# CreateOrganizationTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **int64** |  | 
**Description** | **string** |  | 

## Methods

### NewCreateOrganizationTokenRequest

`func NewCreateOrganizationTokenRequest(organizationId int64, description string, ) *CreateOrganizationTokenRequest`

NewCreateOrganizationTokenRequest instantiates a new CreateOrganizationTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationTokenRequestWithDefaults

`func NewCreateOrganizationTokenRequestWithDefaults() *CreateOrganizationTokenRequest`

NewCreateOrganizationTokenRequestWithDefaults instantiates a new CreateOrganizationTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *CreateOrganizationTokenRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateOrganizationTokenRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateOrganizationTokenRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetDescription

`func (o *CreateOrganizationTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrganizationTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrganizationTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


